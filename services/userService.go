package services

import (
	"log"
	"onez19/datasources"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	// ดึงข้อมูลผู้ใช้ทั้งหมดจาก userService
	users, err := datasources.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	// ส่งข้อมูลผู้ใช้กลับไปยัง client
	return c.JSON(users)
}

func GetUsersWithTenant(c *fiber.Ctx) error {
	// ดึงข้อมูลผู้ใช้ทั้งหมดที่มี tenant จาก datasource
	users, err := datasources.GetAllUsersWithTenant()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users with tenant"})
	}

	// ส่งข้อมูลผู้ใช้ที่มี tenant กลับไปยัง client
	return c.JSON(fiber.Map{
		"tenants": users,
	})
}

func GetUserDetail(ctx *fiber.Ctx) error {
	// รับค่า username จาก path parameter
	username := ctx.Params("username")

	// ดึงรายละเอียดของผู้ใช้จาก datasource
	user, err := datasources.GetUserDetails(username)
	if err != nil {
		log.Println("Error fetching user details:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user details"})
	}

	// ตรวจสอบหากไม่พบข้อมูลผู้ใช้
	if user == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// ส่งรายละเอียดผู้ใช้กลับไปยัง client
	return ctx.JSON(fiber.Map{
		"username":  user.Username,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	})
}



