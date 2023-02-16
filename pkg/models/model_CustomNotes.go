package models

type Model_CustomNotes struct {
    ObjectID string `gorm:"column:objectid" json:"objectid" validate:"required"`
    NoteId string `json:"noteId" `
    VideoId string `json:"videoId" `
    UserId string `json:"userId" `
    Note string `json:"note" `
    Time string `json:"time" `
    }

