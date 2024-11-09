package datasources

import (
	"fmt"
	"log"
	"onez19/config"
	"onez19/entities"
)

func InsertReservation(moveInDateTime string, roomNumber int, tenantUsername string, managerUsername *string, billID *int) (int, error) {
	// คำสั่ง SQL สำหรับแทรกข้อมูลการจอง
	query := `INSERT INTO reservation (move_in_date_time, reservation_room_number, tenant_username, manager_username, bill_id)
	          VALUES (?, ?, ?, ?, ?)`

	// เตรียมการ query และ execute
	result, err := config.DB.Exec(query, moveInDateTime, roomNumber, tenantUsername, managerUsername, billID)
	if err != nil {
		log.Println("Error inserting reservation:", err)
		return 0, err
	}

	// ดึงค่า reservation_id ที่ถูกสร้างขึ้นจากการ insert
	reservationID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error fetching LastInsertId for reservation:", err)
		return 0, err
	}

	// คืนค่า reservation_id ที่ถูกสร้างขึ้น
	return int(reservationID), nil
}

func GetReservationsByUsername(tenantUsername string) ([]entities.Reservation, error) {
	// SQL Query สำหรับดึงข้อมูล reservation ตาม tenant_username
	query := `SELECT reservation_id, reservation_room_number, reservation_status 
	          FROM reservation
	          WHERE tenant_username = ?`

	// เตรียมการ execute query
	rows, err := config.DB.Query(query, tenantUsername)
	if err != nil {
		log.Println("Error fetching reservations by username:", err)
		return nil, err
	}
	defer rows.Close()

	// สร้าง slice สำหรับเก็บผลลัพธ์ของ reservations
	var reservations []entities.Reservation

	// วน loop อ่านข้อมูลจาก rows
	for rows.Next() {
		var reservation entities.Reservation
		if err := rows.Scan(&reservation.ReservationID, &reservation.ReservationRoomNumber, &reservation.ReservationStatus); err != nil {
			log.Println("Error scanning reservation:", err)
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	// คืนค่า slice ของ reservations
	return reservations, nil
}

func UpdateReservationStatus(reservationID int, reservationStatus int) error {
	// SQL Query สำหรับอัปเดตสถานะของการจอง
	query := `UPDATE reservation SET reservation_status = ? WHERE reservation_id = ?`

	// Execute คำสั่ง SQL และรับผลลัพธ์
	result, err := config.DB.Exec(query, reservationStatus, reservationID)
	if err != nil {
		log.Println("Error updating reservation status:", err)
		return err
	}

	// ตรวจสอบจำนวนแถวที่ได้รับผลกระทบ
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error fetching rows affected:", err)
		return err
	}

	// ถ้าไม่มีแถวใดได้รับผลกระทบ แสดงว่าข้อมูลอาจไม่ตรงกับ reservationID ที่ระบุ
	if rowsAffected == 0 {
		log.Printf("No rows updated. Please check if reservationID %d exists.", reservationID)
		return fmt.Errorf("No rows updated")
	}

	log.Printf("Rows affected: %d", rowsAffected)
	return nil
}

func GetReservations() ([]entities.Reservation, error) {
	// SQL Query สำหรับดึงข้อมูล reservation ทั้งหมด
	query := `SELECT reservation_id, move_in_date_time, reservation_room_number, reservation_status, tenant_username, manager_username, bill_id
	          FROM reservation`

	// เตรียมการ execute query
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Error fetching reservations:", err)
		return nil, err
	}
	defer rows.Close()

	// สร้าง slice สำหรับเก็บผลลัพธ์ของ reservations
	var reservations []entities.Reservation

	// วน loop อ่านข้อมูลจาก rows
	for rows.Next() {
		var reservation entities.Reservation
		if err := rows.Scan(&reservation.ReservationID, &reservation.MoveInDateTime, &reservation.ReservationRoomNumber, &reservation.ReservationStatus, &reservation.TenantUsername, &reservation.ManagerUsername, &reservation.BillID); err != nil {
			log.Println("Error scanning reservation:", err)
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	// คืนค่า slice ของ reservations
	return reservations, nil
}

func GetReservationByID(reservationID int) (*entities.Reservation, error) {
	// คำสั่ง SQL เพื่อดึงข้อมูล reservation ตาม reservation_id
	query := `
		SELECT reservation_id, move_in_date_time, reservation_room_number, reservation_status, tenant_username, manager_username, bill_id
		FROM reservation
		WHERE reservation_id = ?
	`

	var reservation entities.Reservation

	// ใช้คำสั่ง Query เพื่อดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query, reservationID)
	if err != nil {
		// หากเกิดข้อผิดพลาดในการ query
		log.Println("Error fetching reservation by ID:", err)
		return nil, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	if rows.Next() {
		// สแกนข้อมูลเข้าใน struct
		if err := rows.Scan(
			&reservation.ReservationID,
			&reservation.MoveInDateTime,
			&reservation.ReservationRoomNumber,
			&reservation.ReservationStatus,
			&reservation.TenantUsername,
			&reservation.ManagerUsername,
			&reservation.BillID,
		); err != nil {
			// หากเกิดข้อผิดพลาดในการแปลงข้อมูล
			log.Println("Error scanning row:", err)
			return nil, err
		}
	} else {
		// หากไม่พบแถวที่ตรงกับเงื่อนไข
		log.Println("No reservation found with the specified reservation ID")
		return nil, nil // คืนค่า nil หากไม่พบข้อมูล
	}

	// ตรวจสอบข้อผิดพลาดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	// คืนค่า reservation ที่พบ
	return &reservation, nil
}

func UpdateReservationDetails(reservationID int, billID int, managerUsername string) error {
	// SQL Query สำหรับการอัปเดต bill_id และ manager_username
	query := `UPDATE reservation 
			  SET bill_id = ?, manager_username = ? 
			  WHERE reservation_id = ?`

	// เตรียมการ query และ execute
	_, err := config.DB.Exec(query, billID, managerUsername, reservationID)
	if err != nil {
		log.Println("Error executing update query:", err)
		return fmt.Errorf("Failed to update reservation details")
	}

	// คืนค่า nil หากอัปเดตสำเร็จ
	return nil
}

func GetReservationByBillID(billID int) (int, error) {
	// คำสั่ง SQL เพื่อดึง reservation_id ที่สัมพันธ์กับ bill_id
	query := `
		SELECT reservation_id 
		FROM reservation 
		WHERE bill_id = ?
	`

	var reservationID int

	// ใช้คำสั่ง Query เพื่อดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query, billID)
	if err != nil {
		log.Println("Error fetching reservation by bill ID:", err)
		return 0, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	if rows.Next() {
		// สแกนข้อมูล reservation_id
		if err := rows.Scan(&reservationID); err != nil {
			log.Println("Error scanning row:", err)
			return 0, err
		}
	} else {
		// หากไม่พบแถวที่ตรงกับเงื่อนไข
		log.Printf("No reservation found for bill_id %d", billID)
		return 0, nil // คืนค่า 0 และ nil หากไม่พบข้อมูล
	}

	// ตรวจสอบข้อผิดพลาดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return 0, err
	}

	// คืนค่า reservationID ที่พบ
	return reservationID, nil
}
