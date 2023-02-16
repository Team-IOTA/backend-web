package models

type Model_Summery struct {
    ObjectID string `gorm:"column:objectid" json:"objectid" validate:"required"`
    SummeryId string `json:"summeryId" `
    VideoId string `json:"videoId" `
    UserId string `json:"userId" `
    Summery string `json:"summery" `
    Time string `json:"time" `
    }

