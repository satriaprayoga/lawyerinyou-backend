package usecases

import (
	"context"
	"errors"
	"lawyerinyou-backend/interfaces/repo"
	"lawyerinyou-backend/interfaces/services"
	"lawyerinyou-backend/models"
	"lawyerinyou-backend/pkg/email"
	"lawyerinyou-backend/pkg/redis"
	"lawyerinyou-backend/pkg/settings"
	"lawyerinyou-backend/pkg/utils"
	"lawyerinyou-backend/token"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type authService struct {
	lawUserRepo    repo.ILawUserRepository
	contextTimeOut time.Duration
}

func NewAuthService(lawUserRepo repo.ILawUserRepository, contextTimeOut time.Duration) services.IAuthService {
	return &authService{lawUserRepo: lawUserRepo, contextTimeOut: contextTimeOut}
}

func (a *authService) Logout(ctx context.Context, claims token.Claims, Token string) error {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	redis.TurncateList(Token)
	redis.TurncateList(claims.ID)

	return nil
}

func (a *authService) Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	var (
		expireToken = settings.AppConfigSetting.JWTExpired
	)

	DataUser, err := a.lawUserRepo.GetByAccount(dataLogin.Account, dataLogin.UserType)
	if err != nil {
		return nil, errors.New("email anda belum terdaftar")
	}

	if DataUser.UserType != "user" {
		return nil, errors.New("email anda belum terdaftar")
	}
	if !DataUser.IsActive {
		return nil, errors.New("account anda belum aktif, silahkan register ulang dengan email yang sama")
	}
	if !utils.ComparePassword(DataUser.Password, utils.GetPassword(dataLogin.Password)) {
		return nil, errors.New("password yang anda masukan salah")
	}
	sessionID := uuid.New().String()
	token, err := token.GenerateToken(sessionID, DataUser.UserID, dataLogin.Account, dataLogin.UserType)
	if err != nil {
		return nil, err
	}
	redis.AddSession(token, DataUser.UserID, time.Duration(expireToken)*time.Hour)
	restUser := map[string]interface{}{
		"user_id":   DataUser.UserID,
		"email":     DataUser.Email,
		"telp":      DataUser.Telp,
		"user_name": DataUser.Name,
		"user_type": DataUser.UserType,
		//"file_id":   DataUser.FileID,
		//"file_name": DataFile.FileName,
		//"file_path": DataFile.FilePath,
	}
	response := map[string]interface{}{
		"token":     token,
		"data_user": restUser,
	}

	return response, nil
}

func (a *authService) ForgotPassword(ctx context.Context, dataForgot *models.ForgotForm) (result string, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	DataUser, err := a.lawUserRepo.GetByAccount(dataForgot.Account, dataForgot.UserType)
	if err != nil {
		return "", errors.New("akun anda tidak valid")
	}
	if DataUser.Name == "" {
		return "", errors.New("akun anda tidak valid")
	}

	OTP := utils.GenerateNumber(4)

	//check redis
	data := redis.GetSession(dataForgot.Account + "_Forgot")
	if data != "" {
		redis.TurncateList(dataForgot.Account + "_Forgot")
	}
	//store to redis
	err = redis.AddSession(dataForgot.Account+"_Forgot", OTP, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return OTP, nil
}

func (a *authService) GenOTP(ctx context.Context, dataForgot *models.ForgotForm) (result interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	DataUser, err := a.lawUserRepo.GetByAccount(dataForgot.Account, dataForgot.UserType)
	if err != nil {
		return "", errors.New("akun anda tidak valid")
	}
	if DataUser.Name == "" {
		return "", errors.New("akun anda tidak valid")
	}

	OTP := utils.GenerateNumber(4)

	if DataUser.UserID > 0 {
		redis.TurncateList(dataForgot.Account + "_Register")
	}

	//store to redis
	err = redis.AddSession(dataForgot.Account+"_Register", OTP, 24*time.Hour)
	if err != nil {
		return "", err
	}

	out := map[string]interface{}{
		"otp":     OTP,
		"account": dataForgot.Account,
	}
	return out, nil
}

func (a *authService) ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	if dataReset.Passwd != dataReset.ConfirmPasswd {
		return errors.New("password dan confirm password harus sama")
	}

	DataUser, err := a.lawUserRepo.GetByAccount(dataReset.Account, dataReset.UserType)
	if err != nil {
		return err
	}

	if utils.ComparePassword(DataUser.Password, utils.GetPassword(dataReset.Passwd)) {
		return errors.New("password baru tidak boleh sama dengan yang lama")
	}

	DataUser.Password, _ = utils.Hash(dataReset.Passwd)
	err = a.lawUserRepo.UpdatePasswordByEmail(dataReset.Account, DataUser.Password)
	if err != nil {
		return err
	}
	return nil
}

func (a *authService) Register(ctx context.Context, dataRegister models.RegisterForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	var (
		User models.LawUser
	)

	CekData, err := a.lawUserRepo.GetByAccount(dataRegister.Account, dataRegister.UserType)
	if CekData.Email == dataRegister.Account {
		if CekData.IsActive {
			return output, errors.New("email anda sudah terdaftar")
		}
	}
	if dataRegister.Passwd != dataRegister.ConfirmPasswd {
		return output, errors.New("password dan confirm password tidak sama")
	}

	User.Name = dataRegister.Name
	User.Password, _ = utils.Hash(dataRegister.Passwd)
	User.JoinDate = time.Now()
	User.BirthOfDate = dataRegister.BirthOfDate
	User.UserType = dataRegister.UserType
	User.IsActive = false
	User.Email = dataRegister.Account

	if CekData.UserID > 0 && !CekData.IsActive {
		CekData.Name = User.Name
		CekData.Password = User.Password
		CekData.JoinDate = User.JoinDate
		CekData.UserType = User.UserType
		CekData.IsActive = User.IsActive
		CekData.Email = User.Email

		err = a.lawUserRepo.Update(CekData.UserID, CekData)
		if err != nil {
			return output, err
		}

	} else {
		User.UserEdit = dataRegister.Name
		User.UserInput = dataRegister.Name
		err = a.lawUserRepo.Create(&User)
		if err != nil {
			return output, models.ErrBadParamInput
		}
		mUser := map[string]interface{}{
			"user_input": strconv.Itoa(User.UserID),
			"user_edit":  strconv.Itoa(User.UserID),
		}
		err = a.lawUserRepo.Update(User.UserID, mUser)
		if err != nil {
			return output, err
		}
	}
	GenCode := utils.GenerateCode(4)

	go email.SendEmail(User.Email, "Register", "This is registration")

	if CekData.UserID > 0 {
		redis.TurncateList(dataRegister.Account + "_Register")
	}

	err = redis.AddSession(dataRegister.Account+"_Register", GenCode, 24*time.Hour)
	if err != nil {
		return output, err
	}
	out := map[string]interface{}{
		"otp":     GenCode,
		"account": User.Email,
	}
	return out, nil
}
