package datasources

import (
	"log"
	"onez19/config"

)

func InsertTenant(username string) (bool, error) {
	// เตรียมคำสั่ง SQL สำหรับการแทรกข้อมูล
	query := "INSERT INTO tenant (username) VALUES (?)"

	// ใช้ Exec สำหรับการแทรกข้อมูล
	_, err := config.DB.Exec(query, username)
	if err != nil {
		// ถ้ามีข้อผิดพลาดในการแทรกข้อมูล
		log.Println("Error creating tenant:", err)
		return false, err
	}

	// ถ้าแทรกข้อมูลสำเร็จ คืนค่าผลลัพธ์เป็น true
	return true, nil
}