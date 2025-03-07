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

func CreateBillItem(billItem entities.BillItem) error {
	// คำสั่ง SQL สำหรับการแทรกข้อมูลในตาราง bill_item
	query := `INSERT INTO bill_item (bill_id, bill_item_number, bill_item_name, unit, unit_price)
	          VALUES (?, ?, ?, ?, ?)`

	// Execute คำสั่ง SQL
	_, err := config.DB.Exec(query, billItem.BillID, billItem.BillItemNumber, billItem.BillItemName, billItem.Unit, billItem.UnitPrice)
	if err != nil {
		log.Println("Error inserting into bill_item:", err)
		return err
	}

	// ไม่มีข้อผิดพลาดแสดงว่าการแทรกสำเร็จ
	return nil
}
