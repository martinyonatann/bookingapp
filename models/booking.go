package models

import "time"

type Booking struct {
	BookingId    int
	UserId       int
	HotelId      int
	CheckInStr   string
	CheckOutStr  string
	CardNumber   string
	NameOnCard   string
	CardExpMonth int
	CardExpYear  int
	Smoking      bool
	Beds         int

	// Transient
	CheckInDate  time.Time
	CheckOutDate time.Time
	User         *Users
	Hotel        *Hotel
}