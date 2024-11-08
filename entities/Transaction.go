package entities

type Transaction struct {
	TransactionID     int    `json:"transaction_id"`
	PaymentDateTime   string `json:"payment_date_time"`
	TransactionStatus int    `json:"transaction_status"`
	BillID            int    `json:"bill_id"`
}
