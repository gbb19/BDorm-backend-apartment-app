package datasources

import (
	"log"
	"onez19/config"
	"onez19/entities"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() ([]entities.UserResponse, error) {
	var users []entities.UserResponse

	// ดึงข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	rows, err := config.DB.Query("SELECT username, first_name, last_name FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.UserResponse
		if err := rows.Scan(&user.Username, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByUsername(username string) (int, error) {
	// เตรียมคำสั่ง SQL เพื่อค้นหาจำนวนผู้ใช้ที่มี username ตรงกับที่รับมา
	query := "SELECT COUNT(*) AS user_count FROM user WHERE username = ?"

	var userCount int

	// ใช้คำสั่ง QueryRow เพื่อดึงข้อมูลที่ตรงกับเงื่อนไข
	err := config.DB.QueryRow(query, username).Scan(&userCount)
	if err != nil {
		// หากเกิดข้อผิดพลาดในการ query จะคืนค่าผิดพลาด
		log.Println("Error fetching user count:", err)
		return 0, err
	}

	// คืนค่าจำนวนผู้ใช้
	return userCount, nil
}

func InsertUser(user entities.User) (bool, error) {
	// แฮชรหัสผ่านด้วย bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return false, err
	}

	// สั่งแทรกข้อมูลผู้ใช้ลงฐานข้อมูล
	_, err = config.DB.Exec("INSERT INTO user (username, password, first_name, last_name) VALUES (?, ?, ?, ?)",
		user.Username, hashedPassword, user.FirstName, user.LastName)
	if err != nil {
		log.Println("Error creating user:", err)
		return false, err
	}

	// คืนค่า true หากแทรกข้อมูลสำเร็จ
	return true, nil
}

func LoginUser(user entities.User) (entities.UserResponse, error) {
	var response entities.UserResponse

	// ค้นหาผู้ใช้ในฐานข้อมูล
	selectedUser := new(entities.User)
	row := config.DB.QueryRow("SELECT username, password, first_name, last_name FROM user WHERE username = ?", user.Username)
	err := row.Scan(&selectedUser.Username, &selectedUser.Password, &selectedUser.FirstName, &selectedUser.LastName)
	if err != nil {
		return response, err
	}

	// ตรวจสอบ password ที่ผู้ใช้กรอกกับ password ที่เก็บในฐานข้อมูล
	err = bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return response, err
	}

	// สร้าง JWT token
	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return response, err
	}

	// ตั้งค่าข้อมูลใน response
	response.Username = selectedUser.Username
	response.FirstName = selectedUser.FirstName
	response.LastName = selectedUser.LastName
	response.Token = t

	// ส่งคืนข้อมูลทั้งหมด
	return response, nil
}