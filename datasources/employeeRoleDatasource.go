package datasources

import (
	"onez19/config"
)

func GetEmployeeRolesByUsername(username string) ([]string, error) {
	var roleNames []string

	query := `
		SELECT role_name 
		FROM employee_role 
		WHERE username = ?
	`

	// ดึงข้อมูลหลายแถว
	rows, err := config.DB.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// วนลูปเพื่อดึง role_name ทุกอันที่เกี่ยวข้อง
	for rows.Next() {
		var roleName string
		if err := rows.Scan(&roleName); err != nil {
			return nil, err
		}
		roleNames = append(roleNames, roleName)
	}

	// ตรวจสอบข้อผิดพลาดจาก rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// คืนค่า slice ที่เก็บ role_name ทั้งหมด
	return roleNames, nil
}
