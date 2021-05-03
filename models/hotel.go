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

type Hotels struct {
	Hotels []Hotel
}

type ResponseModel struct {
	ResponseCode int     `json:"rc"`
	Message      string  `json:"message"`
	Detail       string  `json:"detail"`
	Data         []Hotel `json:"data"`
}

func RespSuccess(hotel []Hotel, detail string) ResponseModel {
	var response ResponseModel
	response.ResponseCode = 200
	response.Message = "Process Successfully"
	response.Detail = detail
	response.Data = hotel
	return response
}
func RespFailed(rc int, detail string) ResponseModel {
	var response ResponseModel
	response.ResponseCode = rc
	response.Message = "Process Failed"
	response.Detail = detail
	return response
}
