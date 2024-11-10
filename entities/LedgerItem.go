package entities

type LedgerItem struct {
	LedgerItemRoomNumber int            `json:"ledger_item_room_number"` // หมายเลขห้อง
	LedgerMonth          int            `json:"ledger_month"`            // เดือนของ Ledger
	LedgerYear           int            `json:"ledger_year"`             // ปีของ Ledger
	WaterUnit            int            `json:"water_unit"`              // หน่วยน้ำ
	ElectricityUnit      int            `json:"electricity_unit"`        // หน่วยไฟฟ้า
	LedgerItemStatus     int            `json:"ledger_item_status"`      // สถานะของ Ledger Item (0, 1)
	Contract             ContractLedger // ฟิลด์ใหม่แสดงว่า room นี้ active หรือไม่
}

type LedgerItemUpdate struct {
	WaterUnit            int `json:"water_unit"`
	ElectricityUnit      int `json:"electricity_unit"`
	LedgerMonth          int `json:"ledger_month"`
	LedgerYear           int `json:"ledger_year"`
	LedgerItemRoomNumber int `json:"ledger_item_room_number"`
}
