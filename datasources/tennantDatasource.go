package datasources

import (
	"log"
	"onez19/config"
	"onez19/entities"
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

func GetAllUsersWithTenant() ([]entities.TenantResponse, error) {
	var users []entities.TenantResponse

	// คำสั่ง SQL เพื่อดึงข้อมูล username, first_name, last_name จาก user และ tenant
	query := `
		SELECT user.username, user.first_name, user.last_name
		FROM user
		JOIN tenant ON tenant.username = user.username
	`

	// ดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Error fetching users with tenant:", err)
		return nil, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	for rows.Next() {
		var user entities.TenantResponse
		if err := rows.Scan(&user.Username, &user.FirstName, &user.LastName); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		users = append(users, user)
	}

	// ตรวจสอบข้อผิดพลาดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return users, nil
}