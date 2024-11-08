package entities

type BillItem struct {
	BillID         int     `json:"bill_id"`
	BillItemNumber int     `json:"bill_item_number"`
	BillItemName   string  `json:"bill_item_name"`
	Unit           int     `json:"unit"`
	UnitPrice      float64 `json:"unit_price"`
}
