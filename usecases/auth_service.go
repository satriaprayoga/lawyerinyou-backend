package usecases

import (
	"context"
	"errors"
	"lawyerinyou-backend/interfaces/repo"
	"lawyerinyou-backend/interfaces/services"
	"lawyerinyou-backend/models"
	"lawyerinyou-backend/pkg/redis"
	"lawyerinyou-backend/pkg/settings"
	"lawyerinyou-backend/pkg/utils"
	"lawyerinyou-backend/token"
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
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeOut)
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
