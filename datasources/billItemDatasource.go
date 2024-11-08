package datasources

import (
	"log"
	"onez19/config"
	"onez19/entities"
)

func GetBillItemsByBillID(billID int) ([]entities.BillItem, error) {
	var billItems []entities.BillItem

	query := `SELECT bill_id, bill_item_number, bill_item_name, unit, unit_price 
	          FROM bill_item 
	          WHERE bill_id = ?`

	rows, err := config.DB.Query(query, billID)
	if err != nil {
		log.Println("Error querying bill items:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var billItem entities.BillItem
		if err := rows.Scan(&billItem.BillID, &billItem.BillItemNumber, &billItem.BillItemName, &billItem.Unit, &billItem.UnitPrice); err != nil {
			log.Println("Error scanning bill item row:", err)
			return nil, err
		}
		billItems = append(billItems, billItem)
	}

	// ตรวจสอบข้อผิดพลาดที่อาจเกิดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with bill item rows:", err)
		return nil, err
	}

	// คืนค่า array ว่างแทนที่จะเป็นข้อผิดพลาดหากไม่พบข้อมูล
	return billItems, nil
}
