package usecases

import (
	"context"
	"lawyerinyou-backend/interfaces/repo"
	"lawyerinyou-backend/interfaces/services"
	"lawyerinyou-backend/pkg/redis"
	"lawyerinyou-backend/token"
	"time"
)

type authService struct {
	lawUserRepo    *repo.ILawUserRepository
	contextTimeOut time.Duration
}

func NewAuthService(lawUserRepo *repo.ILawUserRepository, contextTimeOut time.Duration) services.IAuthService {
	return &authService{lawUserRepo: lawUserRepo, contextTimeOut: contextTimeOut}
}

func (a *authService) Logout(ctx context.Context, claims token.Claims, Token string) error {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	redis.TurncateList(Token)
	redis.TurncateList(claims.ID)

	return nil
}
