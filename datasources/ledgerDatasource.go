package datasources

import (
	"log"
	"onez19/config"
	"onez19/entities"
)

func CreateLedger(ledger entities.Ledger) error {
	// คำสั่ง SQL สำหรับการแทรกข้อมูลในตาราง ledger
	query := `INSERT INTO ledger (ledger_month, ledger_year, username)
	          VALUES (?, ?, ?)`

	// Execute คำสั่ง SQL
	_, err := config.DB.Exec(query, ledger.LedgerMonth, ledger.LedgerYear, ledger.Username)
	if err != nil {
		log.Println("Error inserting into ledger:", err)
		return err
	}

	// คืนค่า nil ถ้าการแทรกข้อมูลสำเร็จ
	return nil
}
