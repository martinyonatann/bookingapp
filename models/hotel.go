package models

type Hotel struct {
	HotelId int    `json:"hotel_id"`
	Name    string `json:"hotel_name"`
	Address string `json:"hotel_address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zipCode"`
	Country string `json:"country"`
	Price   int    `json:"price"`
}
