package repoimpl

import (
	"fmt"
	"lawyerinyou-backend/interfaces/repo"
	"lawyerinyou-backend/models"
	"lawyerinyou-backend/pkg/logging"
	"lawyerinyou-backend/pkg/settings"

	"gorm.io/gorm"
)

type lawUserRepo struct {
	Conn *gorm.DB
}

func NewLawUserRepo(Conn *gorm.DB) repo.ILawUserRepository {
	return &lawUserRepo{Conn}
}

func (r *lawUserRepo) Create(data *models.LawUser) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	q := r.Conn.Create(data)
	logger.Query(fmt.Sprintf("%v", q))
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *lawUserRepo) Update(ID int, data interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	q := r.Conn.Model(models.LawUser{}).Where("user_id=?", ID).Updates(data)
	logger.Query(fmt.Sprintf("%v", q))
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *lawUserRepo) GetByAccount(account string, userType string) (result models.LawUser, err error) {
	var (
		logger = logging.Logger{}
	)
	query := r.Conn.Where("email LIKE ? OR telp=? AND user_type=?", account, account, userType).Find(&result)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}
	return result, err
}
func (db *lawUserRepo) UpdatePasswordByEmail(Email string, Password string) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Exec(`UPDATE k_user set password = ? AND email = ?`, Password, Email)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *lawUserRepo) GetDataBy(ID int) (result *models.LawUser, err error) {
	var (
		logger  = logging.Logger{}
		LawUser = &models.LawUser{}
	)

	query := db.Conn.Where("user_id=?", ID).Find(&LawUser)
	logger.Query(fmt.Sprintf("%v", query))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}
	return LawUser, nil
}

func (db *lawUserRepo) GetList(queryparam models.ParamList) (result []*models.LawUser, err error) {
	var (
		pageNum  = 0
		pageSize = settings.AppConfigSetting.App.PageSize
		sWhere   = ""
		orderBy  = queryparam.SortField
		logger   = logging.Logger{}
	)
	// pagination
	if queryparam.Page > 0 {
		pageNum = (queryparam.Page - 1) * queryparam.PerPage
	}
	if queryparam.PerPage > 0 {
		pageSize = queryparam.PerPage
	}
	//end pagination

	// Order
	if queryparam.SortField != "" {
		orderBy = queryparam.SortField
	}
	//end Order by

	//WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and " + queryparam.Search
		} else {
			sWhere += queryparam.Search
		}
	}

	//end where

	if pageNum >= 0 && pageSize > 0 {
		query := db.Conn.Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
		logger.Query(fmt.Sprintf("%v", query))
		err = query.Error
	} else {
		query := db.Conn.Where(sWhere).Order(orderBy).Find(&result)
		logger.Query(fmt.Sprintf("%v", query))
		err = query.Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return result, nil
}

func (db *lawUserRepo) Count(querparam models.ParamList) (result int, err error) {
	var (
		sWhere        = ""
		logger        = logging.Logger{}
		_result int64 = 0
	)

	//WHERE
	if querparam.InitSearch == "" {
		sWhere = querparam.InitSearch
	}

	if querparam.Search != "" {
		if sWhere != "" {
			sWhere += " and " + querparam.Search
		}
	}

	query := db.Conn.Model(&models.LawUser{}).Where(sWhere).Count(&_result)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return 0, err
	}
	return int(_result), nil
}

func (db *lawUserRepo) Delete(ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Where("user_id=?", ID).Delete(&models.LawUser{})
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil

}
