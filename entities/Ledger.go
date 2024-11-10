package entities

type Ledger struct {
	LedgerMonth int    `json:"ledger_month"` // เดือนของ Ledger
	LedgerYear  int    `json:"ledger_year"`  // ปีของ Ledger
	Username    string `json:"username"`     // ชื่อผู้ใช้งานที่เกี่ยวข้อง
}
