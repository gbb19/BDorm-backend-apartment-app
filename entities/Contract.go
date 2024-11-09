package entities

type Contract struct {
	ContractNumber     int     `json:"contract_number"`
	ContractYear       int     `json:"contract_year"`
	ContractRoomNumber int     `json:"contract_room_number"`
	RentalPrice        float64 `json:"rental_price"`
	WaterRate          float64 `json:"water_rate"`
	ElectricityRate    float64 `json:"electricity_rate"`
	InternetServiceFee float64 `json:"internet_service_fee"`
}

type ContractResponse struct {
	ContractNumber     int `json:"contract_number"`
	ContractYear       int `json:"contract_year"`
	ContractRoomNumber int `json:"contract_room_number"`
}

type ContractCreate struct {
	ContractNumber     int     `json:"contract_number"`
	ContractYear       int     `json:"contract_year"`
	ContractRoomNumber int     `json:"contract_room_number"`
	RentalPrice        float64 `json:"rental_price"`
	WaterRate          float64 `json:"water_rate"`
	ElectricityRate    float64 `json:"electricity_rate"`
	InternetServiceFee float64 `json:"internet_service_fee"`
	Username           string  `json:"username"`
}
