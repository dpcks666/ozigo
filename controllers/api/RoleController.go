package api

import (
	"ozigo/app"
	"ozigo/models"

	"github.com/gofiber/fiber/v2"
)

// Return all roles as JSON
func GetAllRoles(c *fiber.Ctx) error {
	db := app.Instance().DB

	var Role []models.Role
	if response := db.Find(&Role); response.Error != nil {
		panic("Error occurred while retrieving roles from the database: " + response.Error.Error())
	}
	err := c.JSON(Role)
	if err != nil {
		panic("Error occurred when returning JSON of roles: " + err.Error())
	}
	return err

}

// Return a single role as JSON
func GetRole(c *fiber.Ctx) error {
	db := app.Instance().DB
	Role := new(models.Role)
	id, _ := c.ParamsInt("id")
	if response := db.Find(&Role, "id = ?", id); response.Error != nil {
		panic("An error occurred when retrieving the role: " + response.Error.Error())
	}
	if Role.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"ID": id,
		})
	}
	return c.JSON(Role)
}

// Add a single role to the database
func AddRole(c *fiber.Ctx) error {
	db := app.Instance().DB
	Role := new(models.Role)
	if err := c.BodyParser(Role); err != nil {
		panic("An error occurred when parsing the new role: " + err.Error())
	}
	if response := db.Create(&Role); response.Error != nil {
		panic("An error occurred when storing the new role: " + response.Error.Error())
	}
	err := c.JSON(Role)
	if err != nil {
		panic("Error occurred when returning JSON of a role: " + err.Error())
	}
	return err
}

// Edit a single role
func EditRole(c *fiber.Ctx) error {
	db := app.Instance().DB
	id := c.Params("id")
	EditRole := new(models.Role)
	Role := new(models.Role)
	if err := c.BodyParser(EditRole); err != nil {
		panic("An error occurred when parsing the edited role: " + err.Error())
	}
	if response := db.Find(&Role, id); response.Error != nil {
		panic("An error occurred when retrieving the existing role: " + response.Error.Error())
	}
	// Role does not exist
	if Role.ID == 0 {
		err := c.SendStatus(fiber.StatusNotFound)
		if err != nil {
			panic("Cannot return status not found: " + err.Error())
		}
		err = c.JSON(fiber.Map{
			"ID": id,
		})
		if err != nil {
			panic("Error occurred when returning JSON of a role: " + err.Error())
		}
		return err
	}
	Role.Name = EditRole.Name
	Role.Description = EditRole.Description
	db.Save(&Role)

	err := c.JSON(Role)
	if err != nil {
		panic("Error occurred when returning JSON of a role: " + err.Error())
	}
	return err
}

// Delete a single role
func DeleteRole(c *fiber.Ctx) error {
	db := app.Instance().DB
	id := c.Params("id")
	var Role models.Role
	db.Find(&Role, id)
	if response := db.Find(&Role); response.Error != nil {
		panic("An error occurred when finding the role to be deleted: " + response.Error.Error())
	}
	db.Delete(&Role)

	err := c.JSON(fiber.Map{
		"ID":      id,
		"Deleted": true,
	})
	if err != nil {
		panic("Error occurred when returning JSON of a role: " + err.Error())
	}
	return err
}
