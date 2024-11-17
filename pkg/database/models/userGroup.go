package model

import (
	"time"
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	UserID				string			`gorm:"primaryKey"`
	GroupID				string			`gorm:"primaryKey"`
	Stoppage			bool				`json:"stoppage" gorm:"default:false"`
	IsHome				bool				`json:"is_home" gorm:"default:false"`
	CreatedAt 		time.Time
  DeletedAt 		gorm.DeletedAt
}

type UserGroupRequest struct {
	UserID 				string
	GroupID				string
}

// func FilterUsersGroup(user *User) UserGroupResponse {
// 	return UserGroupResponse{
// 		ID:        		user.ID,
// 		Firstname:  	user.Firstname,
// 		Ridename:			user.Ridename,
// 		Thumbnail: 		user.Thumbnail,
// 		Groups:				filterGroups(user.Groups),
// 	}
// }

// func (userGroup *UserGroup) BeforeCreate(tx *gorm.DB) (err error) {
// 	fmt.Println("-----------------Inside before create for usergroup")

// 	userGroup.ID = uuid.New().String()
// 	return
// }

func MigrateUserGroups(db *gorm.DB) error {
	err := db.AutoMigrate(&UserGroup{})
	return err
}