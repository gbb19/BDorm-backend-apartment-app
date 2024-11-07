package datasources

import (
	"log"
	"onez19/config"
)


func GetAllContractsByUsername(username string) ([]int, error) {
	var roomNumbers []int

	// เตรียมคำสั่ง SQL เพื่อดึงข้อมูลหมายเลขห้องจากตาราง contract
	query := "SELECT contract_room_number FROM contract WHERE username = ? AND contract_status = 0"

	// ใช้คำสั่ง Query เพื่อดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query, username)
	if err != nil {
		log.Println("Error fetching active contracts:", err)
		return nil, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	for rows.Next() {
		var roomNumber int
		if err := rows.Scan(&roomNumber); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		roomNumbers = append(roomNumbers, roomNumber)
	}

	// ตรวจสอบว่ามีข้อผิดพลาดเกิดขึ้นระหว่างการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	// คืนค่าหมายเลขห้องที่ใช้งาน
	if len(roomNumbers) > 0 {
		return roomNumbers, nil
	}
	return nil, nil
}