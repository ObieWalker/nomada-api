package transaction

import (
	"gorm.io/gorm"
	"github.com/obiewalker/nomada-api/pkg/database/models"
)

func CreateGroupTransaction(db *gorm.DB, group *model.Group) (*model.Group, error) {
	db.Transaction(func(tx *gorm.DB)error {
		if err := tx.Create(&group).Error; err != nil {
			return err
		}
		// var owner model.User

		// if err := tx.First(&owner, "id = ?", group.OwnerID).Error; err != nil {
		// 	return err
		// }

		// group.Members = append(group.Members, owner)
		// if err := tx.Save(&group).Error; err != nil {
		// 	return err
		// }

		// // group.Owner = owner
		// if err := tx.Model(&group).Updates(model.Group{Owner: owner}).Error; err != nil {
		// 	return err
		// }



		// if err := tx.Save(&group).Error; err != nil {
		// 	return err
		// }
	
	
		newUserGroup := &model.UserGroup {
			UserID:       group.OwnerID,
			GroupID:      group.ID,
		}

		if err := tx.Create(&newUserGroup).Error; err != nil {
			return err
		}
		return nil
	})
	return group, nil
}