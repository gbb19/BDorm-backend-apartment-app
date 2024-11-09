package datasources

import (
	"log"
	"onez19/config"
	"onez19/entities"
	"time"
)

func GetTransactionsByBillID(billID int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	// SQL คำสั่งเพื่อดึงข้อมูลจากตาราง transaction
	query := `SELECT transaction_id, payment_date_time, transaction_status, bill_id 
	          FROM transaction 
	          WHERE bill_id = ?`

	// ดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query, billID)
	if err != nil {
		log.Println("Error fetching transactions:", err)
		return nil, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	for rows.Next() {
		var transaction entities.Transaction
		var paymentDateTime []byte // ใช้ []byte ก่อนแปลงเป็น time.Time
		if err := rows.Scan(&transaction.TransactionID, &paymentDateTime, &transaction.TransactionStatus, &transaction.BillID); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		// แปลง paymentDateTime จาก []byte เป็น time.Time
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(paymentDateTime))
		if err != nil {
			log.Println("Error parsing payment_date_time:", err)
			return nil, err
		}

		// แปลง time.Time เป็น string
		transaction.PaymentDateTime = parsedTime.Format("2006-01-02 15:04:05")

		transactions = append(transactions, transaction)
	}

	// ตรวจสอบข้อผิดพลาดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return transactions, nil
}

func InsertTransaction(billID int) (int, error) {
	// คำสั่ง SQL สำหรับการแทรกข้อมูลลงในตาราง transaction
	query := `INSERT INTO transaction (bill_id)
	          VALUES (?)`

	// เตรียมการ query และการ execute
	result, err := config.DB.Exec(query, billID)
	if err != nil {
		log.Println("Error inserting transaction:", err)
		return 0, err
	}

	// ดึงค่า transaction_id ที่ถูกสร้างขึ้นจากการ insert
	transactionID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error fetching LastInsertId:", err)
		return 0, err
	}

	// คืนค่า transaction_id ที่ถูกสร้างขึ้น
	return int(transactionID), nil
}

func UpdateTransactionStatus(transactionID int, status int) error {
	// สร้างคำสั่ง SQL สำหรับอัปเดตสถานะของ transaction
	query := `UPDATE transaction SET transaction_status = ? WHERE transaction_id = ?`

	// เตรียม statement และ execute กับฐานข้อมูล
	_, err := config.DB.Exec(query, status, transactionID)
	if err != nil {
		log.Println("Error updating transaction status:", err)
		return err
	}

	return nil
}
