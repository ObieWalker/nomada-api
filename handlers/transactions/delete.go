package transaction

import (
	"gorm.io/gorm"
	"github.com/obiewalker/nomada-api/pkg/database/models"
)

func RemoveFromGroup(db *gorm.DB,  groupId string, currentUserId string, userId string) (error) {
	db.Transaction(func(tx *gorm.DB)error {
		var user = model.User{ID: userId}
		var group = model.Group{ID: groupId}
	
		if err := tx.First(&user).Error; err != nil {
			return err
		}
	
		if err := tx.First(&group).Error; err != nil {
			return err
		}

		if group.OwnerID == currentUserId && currentUserId == userId {
			return &model.LeaveGroupAsOwnerError{}
		}
		result := db.Model(&group).Association("Members").Delete(&user)
		if result != nil {
			return result
		}
		return nil
	})
	return nil
}