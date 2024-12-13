package models

type Room struct {
	RoomID         string   `gorm:"primaryKey" json:"room_id"`
	RatePerNight   float64  `json:"rate_per_night"`
	MaxGuests      int      `json:"max_guests"`
	AvailableDates []string `gorm:"type:text[]" json:"available_dates"`
}