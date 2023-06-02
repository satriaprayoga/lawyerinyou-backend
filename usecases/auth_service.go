package usecases

import (
	"lawyerinyou-backend/interfaces/repo"
	"lawyerinyou-backend/interfaces/services"
)

type authService struct {
	lawUserRepo *repo.ILawUserRepository
}

func NewAuthService(lawUserRepo *repo.ILawUserRepository) services.IAuthService {
	return &authService{lawUserRepo: lawUserRepo}
}
