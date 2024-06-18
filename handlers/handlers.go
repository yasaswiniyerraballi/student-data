package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yasaswiniyerraballi/studentdatabase-management/database"
	student "github.com/yasaswiniyerraballi/studentdatabase-management/models"
)

func GetStudents(c *fiber.Ctx) error {
	var students []student.Student
	result := database.DB.Find(&students)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(students)
}

func CreateStudent(c *fiber.Ctx) error {
	var newStudent student.Student
	if err := c.BodyParser(&newStudent); err != nil {
		return c.Status(400).SendString("Invalid student data")
	}
	// Validate student data (optional)
	result := database.DB.Create(&newStudent)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(newStudent) // Or return success message
}

func GetStudentByID(c *fiber.Ctx) error {
	id := c.Params("id") // Get student ID from URL parameter
	var studentByID student.Student
	result := database.DB.Find(&studentByID, "id = ?", id)
	if result.Error != nil {
		if result.RecordNotFound() { // Check for record not found error
			return c.Status(404).SendString("Student not found")
		}
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(studentByID)
}

func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var existingStudent student.Student
	result := database.DB.First(&existingStudent, "id = ?", id)
	if result.Error != nil {
		if result.RecordNotFound() {
			return c.Status(404).SendString("Student not found")
		}
		return c.Status(500).SendString(result.Error.Error())
	}
	// Update existingStudent based on request body data (optional)
	// Validate updated data (optional)
	result = database.DB.Save(&existingStudent)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	return c.JSON(existingStudent) // Or return success message
}

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	result := database.DB.Delete(&student.Student{}, "id = ?", id)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Student not found")
	}
	return c.SendString("Student deleted successfully") // Or return success message
}
