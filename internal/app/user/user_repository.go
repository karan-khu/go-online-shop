package user

import (
	"errors"
	"log/slog"

	"github.com/jinzhu/copier"

	"github.com/karan-khu/go-online-shop/config"
	"github.com/karan-khu/go-online-shop/pkg/database"
	"github.com/karan-khu/go-online-shop/pkg/database/models"
)

type UserRepositoryImpl struct {
	logger *slog.Logger
	db     database.Database
}

func NewUserRepository(logger *slog.Logger, conf *config.Config) UserRepository {
	return &UserRepositoryImpl{
		logger: logger,
		db:     conf.GetDb("main"),
	}
}

func (r *UserRepositoryImpl) Creating(user *UserEntity) (result *UserEntity, err error) {
	newUser := new(models.UserRecord)
	copier.Copy(newUser, user)

	userRecord := new(models.UserRecord)
	if err = r.db.Connect().Create(newUser).Scan(userRecord).Error; err != nil {
		r.logger.Error("failed to create user", "error", err)
		return nil, errors.New("failed to create user")
	}

	copier.Copy(result, userRecord)
	return result, nil
}

func (r *UserRepositoryImpl) FindById(userId string) (result *UserEntity, err error) {
	userRecord := new(models.UserRecord)
	if err = r.db.Connect().Where("UserId = ?", userId).First(userRecord).Error; err != nil {
		r.logger.Error("failed to find user by id", "error", err)
		return nil, errors.New("user not found")
	}

	copier.Copy(result, userRecord)
	return result, nil
}
