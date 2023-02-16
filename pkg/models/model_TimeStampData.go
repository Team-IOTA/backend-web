package models

type Model_TimeStampData struct {
    ObjectID string `gorm:"column:objectid" json:"objectid" validate:"required"`
    VideoId string `json:"videoId" `
    UserId string `json:"userId" `
    Topic string `json:"topic" `
    TimeStamp string `json:"timeStamp" `
    }

