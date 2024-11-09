package entities

type Bill struct {
	BillID         int    `json:"bill_id"`
	PaymentTerm    int    `json:"payment_term"`
	CreateDateTime string `json:"create_date_time"`
	BillStatus     int    `json:"bill_status"`
}

type BillCreate struct {
	PaymentTerm     int    `json:"payment_term"`     // ระยะเวลาการชำระเงิน
	TenantUsername  string `json:"tenant_username"`  // ชื่อผู้เช่า
	CashierUsername string `json:"cashier_username"` // ชื่อพนักงานที่รับเงิน
}
