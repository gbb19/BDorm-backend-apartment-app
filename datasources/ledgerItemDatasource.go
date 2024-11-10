package datasources

import (
	"log"
	"onez19/config"
	"onez19/entities"
)

func CreateLedgerItem(ledgerItem entities.LedgerItem) error {
	// คำสั่ง SQL สำหรับการแทรกข้อมูลในตาราง ledger_item
	query := `INSERT INTO ledger_item (ledger_item_room_number, ledger_month, ledger_year)
	          VALUES (?, ?, ?)`

	// Execute คำสั่ง SQL
	_, err := config.DB.Exec(query, ledgerItem.LedgerItemRoomNumber, ledgerItem.LedgerMonth, ledgerItem.LedgerYear)
	if err != nil {
		log.Println("Error inserting into ledger_item:", err)
		return err
	}

	// คืนค่า nil ถ้าการแทรกข้อมูลสำเร็จ
	return nil
}

func UpdateLedgerItem(waterUnit int, electricityUnit int, month int, year int, roomNumber int) error {
	// คำสั่ง SQL สำหรับการอัปเดตข้อมูลในตาราง ledger_item
	query := `UPDATE ledger_item
	          SET water_unit = ?, electricity_unit = ?
	          WHERE ledger_month = ? AND ledger_year = ? AND ledger_item_room_number = ?`

	// Execute คำสั่ง SQL
	_, err := config.DB.Exec(query, waterUnit, electricityUnit, month, year, roomNumber)
	if err != nil {
		log.Println("Error updating ledger_item:", err)
		return err
	}

	// คืนค่า nil ถ้าการอัปเดตสำเร็จ
	return nil
}

func UpdateLedgerItemStatus(month int, year int, roomNumber int) error {
	// คำสั่ง SQL สำหรับการอัปเดต ledger_item_status เป็น 1
	query := `UPDATE ledger_item
	          SET ledger_item_status = 1
	          WHERE ledger_month = ? AND ledger_year = ? AND ledger_item_room_number = ?`

	// Execute คำสั่ง SQL
	_, err := config.DB.Exec(query, month, year, roomNumber)
	if err != nil {
		log.Println("Error updating ledger_item_status:", err)
		return err
	}

	// คืนค่า nil ถ้าการอัปเดตสำเร็จ
	return nil
}

func GetLedgerItemsByMonthAndYear(month int, year int) ([]entities.LedgerItem, error) {
	var ledgerItems []entities.LedgerItem

	query := `SELECT ledger_month, ledger_year, ledger_item_room_number, water_unit, electricity_unit, ledger_item_status
	          FROM ledger_item
	          WHERE ledger_month = ? AND ledger_year = ?`

	rows, err := config.DB.Query(query, month, year)
	if err != nil {
		log.Println("Error querying ledger items:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ledgerItem entities.LedgerItem
		if err := rows.Scan(&ledgerItem.LedgerMonth, &ledgerItem.LedgerYear, &ledgerItem.LedgerItemRoomNumber,
			&ledgerItem.WaterUnit, &ledgerItem.ElectricityUnit, &ledgerItem.LedgerItemStatus); err != nil {
			log.Println("Error scanning ledger item row:", err)
			return nil, err
		}

		// เรียกใช้ GetContractDetails เพื่อดึงข้อมูลเพิ่มเติมจากสัญญา
		contractDetails, err := GetContractDetailsLedger(ledgerItem.LedgerItemRoomNumber)
		if err != nil {
			log.Println("Error fetching contract details:", err)
			return nil, err
		}

		// เพิ่มข้อมูล Contract ให้กับ LedgerItem
		if contractDetails != nil {
			ledgerItem.Contract = *contractDetails // ซ้อนข้อมูล Contract เข้าไปใน LedgerItem
		}

		ledgerItems = append(ledgerItems, ledgerItem)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error with ledger item rows:", err)
		return nil, err
	}

	return ledgerItems, nil
}
