package services

import (
	"log"
	"onez19/datasources"
	"onez19/entities"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error {
	var user entities.User

	// รับข้อมูลจาก request body
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// ตรวจสอบให้แน่ใจว่ามีข้อมูลครบถ้วน
	if user.Username == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
	}

	// ตรวจสอบว่า username มีอยู่ในระบบหรือไม่
	userExists, err := datasources.GetUserByUsername(user.Username)
	if err != nil {
		log.Println("Error checking user existence:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error checking user existence"})
	}

	if userExists {
		// หากมีผู้ใช้ในระบบแล้ว
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username already exists"})
	}

	// สร้างผู้ใช้ใหม่
	userInsertSuccess, err := datasources.InsertUser(user)
	if err != nil || !userInsertSuccess {
		log.Println("Failed to create user:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// สร้าง tenant ใหม่
	tenantInsertSuccess, err := datasources.InsertTenant(user.Username)
	if err != nil || !tenantInsertSuccess {
		log.Println("Failed to create tenant:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create tenant"})
	}

	// หากทุกขั้นตอนสำเร็จ
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func Login(ctx *fiber.Ctx) error {
	var user entities.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// ตรวจสอบข้อมูลผู้ใช้และรับผลลัพธ์
	resultLogin, err := datasources.LoginUser(user)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// ตรวจสอบข้อมูลพนักงาน
	resultEmployee, err := datasources.GetEmployeeByUsername(user.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// ดึงบทบาทพนักงาน
	var roles []string // ใช้ slice เพื่อเก็บบทบาทหลายๆ อัน
	if resultEmployee != nil {
		resultRoles, err := datasources.GetEmployeeRolesByUsername(user.Username)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		roles = resultRoles // รับค่าบทบาททั้งหมด
	} else {
		roles = append(roles, "tenant") // ถ้าไม่พบพนักงาน ก็ใช้ role เป็น "tenant"
	}

	// ตั้งค่า cookie สำหรับ JWT token
	ctx.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    resultLogin.Token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
	})

	// ส่งคืนข้อมูลผู้ใช้พร้อมกับ token และบทบาท
	return ctx.JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"username":   resultLogin.Username,
			"first_name": resultLogin.FirstName,
			"last_name":  resultLogin.LastName,
			"roles":      roles, // ส่งคืนบทบาททั้งหมดที่ผู้ใช้มี
			"token":      resultLogin.Token,
		},
	})

}
