package entities

type Contract struct {
	ContractNumber     int  `json:"contract_number"`
	ContractYear       int     `json:"contract_year"`
	ContractRoomNumber int  `json:"contract_room_number"`
	RentalPrice        float64 `json:"rental_price"`
	WaterRate          float64 `json:"water_rate"`
	ElectricityRate    float64 `json:"electricity_rate"`
	InternetServiceFee float64 `json:"internet_service_fee"`
	ContractStatus     int  `json:"contract_status"`
	Username           string  `json:"username"`
}
