package user

import (
	"gorm.io/gorm"
	"github.com/obiewalker/nomada-api/pkg/postgres"
	"github.com/obiewalker/nomada-api/pkg/database/models"
) 

func Create(db *gorm.DB, user *model.User) (string, error) {
	err := db.Create(user).Error
	if err != nil {
		if postgres.IsUniqueConstraintError(err, model.UniqueConstraintEmail) {
			return "", &model.EmailDuplicateError{Email: user.Email}
		}
		return "", err
	}
	return user.UUID, nil
}