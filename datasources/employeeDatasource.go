package datasources

import (
	"onez19/config"
	"onez19/entities"
)

func GetEmployeeByUsername(username string) (*entities.Employee, error) {
	var employee entities.Employee

	query := `
		SELECT username 
		FROM employee 
		WHERE username = ?
	`

	row := config.DB.QueryRow(query, username)

	// Scan ผลลัพธ์ลงในตัวแปร employee
	err := row.Scan(&employee.Username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// ถ้าไม่พบข้อมูลให้คืนค่า nil
			return nil, nil
		}
		return nil, err
	}

	return &employee, nil
}
