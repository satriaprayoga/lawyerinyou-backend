package repo

import "lawyerinyou-backend/models"

type ILawUserRepository interface {
	Create(data *models.LawUser) error
	Update(ID int, data interface{}) error
	GetByAccount(account string, userType string) (result models.LawUser, err error)
	UpdatePasswordByEmail(Email string, Password string) error
	GetDataBy(ID int) (result *models.LawUser, err error)
	GetList(queryparam models.ParamList) (result []*models.LawUser, err error)
	Count(querparam models.ParamList) (result int, err error)
	Delete(ID int) error
}
