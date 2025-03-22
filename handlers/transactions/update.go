package transaction

import (
	"fmt"
	"github.com/obiewalker/nomada-api/pkg/database/models"
	"gorm.io/gorm"
)

func AddUserTransaction(db *gorm.DB, groupId string, userId string) (model.User, model.Group, error) {
	var user = model.User{ID: userId}
	var group model.Group
	db.Transaction(func(tx *gorm.DB)error {
		if err := tx.First(&user).Error; err != nil {
			return err
		}

		if err := tx.First(&group, "id = ?", groupId).Error; err != nil {
			return err
		}


		if err := tx.Model(&user).Association("Groups").Append(&group); err != nil {
			return err
		}

	fmt.Println("-------group id after-----------")
	fmt.Println(group.ID)
		// user.Groups = append(user.Groups, &group)
		// if err := tx.Save(&user).Error; err != nil {
		// 	return err
		// }
		// group.Members = append(group.Members, &user)
		// if err := tx.Save(&group).Error; err != nil {
		// 	return err
		// }
		return nil
	})

	return user, group, nil
}