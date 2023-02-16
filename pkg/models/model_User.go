package models

type Model_User struct {
    ObjectID string `gorm:"column:objectid" json:"objectid" validate:"required"`
    UserId string `json:"userId" `
    UserName string `json:"userName" `
    Password string `json:"password" `
    Email string `json:"email" `
    }

