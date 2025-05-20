package repository

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	db "github.com/nguyenhoang711/go-coin-tracking/database/sqlc"
	"github.com/nguyenhoang711/go-coin-tracking/global"
)

type IUserRepository interface {
	Create(ctx *gin.Context, email, hashedPassword string) (db.User, error)
	GetByEmail(ctx *gin.Context, email string) (db.User, error)
}

type userRepo struct{}

func NewUserRepo() IUserRepository {
	return &userRepo{}
}

// Create implements IUserRepository.
func (u *userRepo) Create(ctx *gin.Context, email string, hashedPassword string) (db.User, error) {
	return db.User{}, nil
}

// GetByEmail implements IUserRepository.
func (u *userRepo) GetByEmail(ctx *gin.Context, email string) (db.User, error) {
	user, err := global.Db.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		global.Logger.Error("not found user by email ",
			zap.String("email", email),
			zap.Error(err),
		)
		return db.User{}, err
	}
	return user, nil
}
