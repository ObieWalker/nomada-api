package model

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Group struct {
	gorm.Model
	ID          			string 		`gorm:"type:uuid;primary_key;"`
	Name							string		`json:"name" gorm:"text;not null;default:null"`
	Members 				  []*User	  `gorm:"many2many:user_groups;"`
	OwnerID						string		`gorm:"type:uuid;column:user_foreign_key;not null;"`
	Owner							User			`gorm:"foreignKey:OwnerID"`
	Ride							Ride
}

type GroupRequest struct {
	ID				 string
  Name		   string
}

type GroupResponse struct {
	ID				 string
  Name		   string
	OwnerID		 string
}

// func FilterGroups(groups []Group) []GroupRequest {
// 	var filteredGroups []GroupRequest
// 	for _, group := range groups {
// 		currentGroup := GroupRequest{
// 			ID: 			group.ID,
// 			Name:			group.Name,
// 			Members:	group.Members,
// 			OwnerID:	group.OwnerID,
// 		}
// 		filteredGroups = append(filteredGroups, currentGroup)
// 	}
// 	return filteredGroups
// }

func FilterGroup(group *Group) GroupResponse {
		return GroupResponse{
			ID: 			group.ID,
			Name:			group.Name,
			OwnerID:	group.OwnerID,
		}
}

func (group *Group) BeforeCreate(tx *gorm.DB) (err error) {
	group.ID = uuid.New().String()
	return
}

func MigrateGroups(db *gorm.DB) error {
	err := db.AutoMigrate(&Group{})
	return err
}

type LeaveGroupAsOwnerError struct {}

func (e *LeaveGroupAsOwnerError) Error() string {
  return "You cannot leave this group as the Admin. Try deleting instead."
}