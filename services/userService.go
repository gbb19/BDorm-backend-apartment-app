package services

import (
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
