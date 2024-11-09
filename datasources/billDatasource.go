package datasources

import (
	"log"
	"onez19/config"
	"onez19/entities"
	"time"
)

func GetBillsByTenantUsername(username string) ([]entities.Bill, error) {
	var bills []entities.Bill

	// คำสั่ง SQL เพื่อดึงข้อมูลจากตาราง bill
	query := `SELECT bill_id, payment_term, create_date_time, bill_status 
	          FROM bill 
	          WHERE tenant_username = ?
						ORDER BY create_date_time DESC
						`

	// ใช้ Query เพื่อดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query, username)
	if err != nil {
		log.Println("Error fetching bills:", err)
		return nil, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	for rows.Next() {
		var bill entities.Bill
		var createDateTime []byte // ใช้เป็น []byte ก่อนแปลงเป็น time.Time
		if err := rows.Scan(&bill.BillID, &bill.PaymentTerm, &createDateTime, &bill.BillStatus); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		// แปลง createDateTime จาก []byte เป็น time.Time
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(createDateTime))
		if err != nil {
			log.Println("Error parsing create_date_time:", err)
			return nil, err
		}

		// แปลง time.Time เป็น string รูปแบบที่คุณต้องการ
		bill.CreateDateTime = parsedTime.Format("2006-01-02 15:04:05") // รูปแบบที่ต้องการ

		bills = append(bills, bill)
	}

	// ตรวจสอบข้อผิดพลาดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return bills, nil
}

func GetAllBills() ([]entities.Bill, error) {
	var bills []entities.Bill

	// คำสั่ง SQL เพื่อดึงข้อมูลทั้งหมดจากตาราง bill โดยเรียงลำดับจาก create_date_time ในแบบ DESC
	query := `SELECT bill_id, payment_term, create_date_time, bill_status 
	          FROM bill 
	          ORDER BY create_date_time DESC`

	// ใช้ Query เพื่อดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Error fetching bills:", err)
		return nil, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	for rows.Next() {
		var bill entities.Bill
		var createDateTime []byte // ใช้ []byte ชั่วคราวเพื่อแปลงเป็น time.Time
		if err := rows.Scan(&bill.BillID, &bill.PaymentTerm, &createDateTime, &bill.BillStatus); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		// แปลง createDateTime จาก []byte เป็น time.Time
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(createDateTime))
		if err != nil {
			log.Println("Error parsing create_date_time:", err)
			return nil, err
		}

		// แปลง time.Time เป็น string รูปแบบที่ต้องการ
		bill.CreateDateTime = parsedTime.Format("2006-01-02 15:04:05")

		bills = append(bills, bill)
	}

	// ตรวจสอบข้อผิดพลาดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return bills, nil
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

func UpdateBillStatus(billID string, status string) error {
	// สร้างคำสั่ง SQL เพื่ออัปเดตสถานะของใบแจ้งหนี้
	query := `UPDATE bill SET bill_status = ? WHERE bill_id = ?`

	// ใช้คำสั่ง SQL เพื่ออัปเดตฐานข้อมูล
	_, err := config.DB.Exec(query, status, billID)
	if err != nil {
		log.Println("Error updating bill status:", err)
		return err
	}

	return nil
}
