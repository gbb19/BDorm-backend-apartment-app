package entities


type Bill struct {
	BillID          int       `json:"bill_id"`
	PaymentTerm     int       `json:"payment_term"`
	CreateDateTime  string `json:"create_date_time"`
	BillStatus      int       `json:"bill_status"`
}