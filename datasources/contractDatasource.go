package datasources

import (
	"log"
	"onez19/config"
	"onez19/entities"
)

func GetAllContractsByUsername(username string) ([]entities.ContractResponse, error) {
	var contracts []entities.ContractResponse

	// คำสั่ง SQL เพื่อดึงข้อมูลเฉพาะคอลัมน์ที่ต้องการ
	query := "SELECT contract_number, contract_year, contract_room_number FROM contract WHERE username = ? AND contract_status = 0"

	// ใช้ Query เพื่อดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query, username)
	if err != nil {
		log.Println("Error fetching contracts:", err)
		return nil, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	for rows.Next() {
		var contract entities.ContractResponse
		if err := rows.Scan(&contract.ContractNumber, &contract.ContractYear, &contract.ContractRoomNumber); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	// ตรวจสอบข้อผิดพลาดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return contracts, nil
}

func GetContractDetails(contractNumber, contractYear int) (*entities.Contract, error) {
	// คำสั่ง SQL เพื่อดึงข้อมูลสัญญาตาม contract_number และ contract_year
	query := `
		SELECT contract_number, contract_year, contract_room_number, rental_price, water_rate, electricity_rate, internet_service_fee
		FROM contract
		WHERE contract_number = ? AND contract_year = ?
	`

	var contract entities.Contract

	// ใช้คำสั่ง Query เพื่อดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query, contractNumber, contractYear)
	if err != nil {
		// หากเกิดข้อผิดพลาดในการ query
		log.Println("Error fetching contract details:", err)
		return nil, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	if rows.Next() {
		// สแกนข้อมูลเข้าใน struct
		if err := rows.Scan(
			&contract.ContractNumber,
			&contract.ContractYear,
			&contract.ContractRoomNumber,
			&contract.RentalPrice,
			&contract.WaterRate,
			&contract.ElectricityRate,
			&contract.InternetServiceFee,
		); err != nil {
			// หากเกิดข้อผิดพลาดในการแปลงข้อมูล
			log.Println("Error scanning row:", err)
			return nil, err
		}
	} else {
		// หากไม่พบแถวที่ตรงกับเงื่อนไข
		log.Println("No contract found with the specified contract number and year")
		return nil, nil // คืนค่า nil หากไม่พบข้อมูล
	}

	// ตรวจสอบข้อผิดพลาดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	// คืนค่า contract ที่พบ
	return &contract, nil
}

func GetAllContracts() ([]entities.ContractResponse, error) {
	var contracts []entities.ContractResponse

	// คำสั่ง SQL เพื่อดึงข้อมูลทั้งหมดจากตาราง contract ที่สถานะของสัญญาคือ 0 (active)
	query := "SELECT contract_number, contract_year, contract_room_number FROM contract WHERE contract_status = 0"

	// ใช้ Query เพื่อดึงข้อมูลจากฐานข้อมูล
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("Error fetching contracts:", err)
		return nil, err
	}
	defer rows.Close()

	// อ่านผลลัพธ์จากฐานข้อมูล
	for rows.Next() {
		var contract entities.ContractResponse
		if err := rows.Scan(&contract.ContractNumber, &contract.ContractYear, &contract.ContractRoomNumber); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	// ตรวจสอบข้อผิดพลาดจากการอ่านแถว
	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return contracts, nil
}

func UpdateContractStatus(contractRoomNumber int, status int) error {
	// สร้างคำสั่ง SQL เพื่ออัปเดตสถานะของ contract
	query := `UPDATE contract SET contract_status = ? WHERE contract_room_number = ?`

	// ใช้คำสั่ง SQL เพื่ออัปเดตฐานข้อมูล
	_, err := config.DB.Exec(query, status, contractRoomNumber)
	if err != nil {
		log.Println("Error updating contract status:", err)
		return err
	}

	return nil
}

// ฟังก์ชันตรวจสอบว่า contract_room_number นี้มี contract active อยู่หรือไม่
func CheckContractActive(contractRoomNumber int) (bool, error) {
	// สร้างคำสั่ง SQL เพื่อเช็คสถานะของ contract ที่เป็น active
	query := `SELECT COUNT(*) AS contract_count FROM contract WHERE contract_room_number = ? AND contract_status = 0`

	// ใช้คำสั่ง SQL เพื่อดึงข้อมูล
	var contractCount int
	err := config.DB.QueryRow(query, contractRoomNumber).Scan(&contractCount)
	if err != nil {
		log.Println("Error checking contract status:", err)
		return false, err
	}

	// หาก contract_count มากกว่า 0 หมายความว่า contract นี้ยัง active อยู่
	if contractCount > 0 {
		return true, nil
	}

	return false, nil
}

// ฟังก์ชันสำหรับการสร้าง contract ใหม่
func CreateContract(contractYear, contractRoomNumber int, rentalPrice, waterRate, electricityRate, internetServiceFee float64, username string) error {
	// สร้างคำสั่ง SQL เพื่อสร้าง contract ใหม่
	query := `INSERT INTO contract (contract_year, contract_room_number, rental_price, water_rate, electricity_rate, internet_service_fee, username) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)`

	// ใช้คำสั่ง SQL เพื่อเพิ่มข้อมูล contract ใหม่
	_, err := config.DB.Exec(query, contractYear, contractRoomNumber, rentalPrice, waterRate, electricityRate, internetServiceFee, username)
	if err != nil {
		log.Println("Error creating contract:", err)
		return err
	}

	return nil
}
