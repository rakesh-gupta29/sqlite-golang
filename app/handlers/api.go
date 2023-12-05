package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/rakesh-gupta29/sqlite-golang/app/database"
	"github.com/rakesh-gupta29/sqlite-golang/app/models"
	"github.com/rakesh-gupta29/sqlite-golang/pkg/validator"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("app is running fine")
}

func SeedData() error {
	seedData := []models.Form{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Message: "Hello, World!"},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Message: "Greetings from Jane!"},
	}

	for _, data := range seedData {
		_, err := database.Database.Exec(`
			INSERT INTO forms (name, email, message) VALUES (?, ?, ?);
		`, data.Name, data.Email, data.Message)
		if err != nil {
			return err
		}
	}

	return nil
}
func getAllData() (string, error) {
	forms, err := GetAllForms()
	if err != nil {
		return "", err
	}
	jsonData, err := formsToJSON(forms)
	if err != nil {
		return "", err
	}

	return jsonData, nil
}

func formsToJSON(forms []models.Form) (string, error) {
	jsonData, err := json.Marshal(forms)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func GetAllForms() ([]models.Form, error) {
	rows, err := database.Database.Query("SELECT * FROM forms")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forms []models.Form

	for rows.Next() {
		var form models.Form
		if err := rows.Scan(&form.ID, &form.Name, &form.Email, &form.Message); err != nil {
			return nil, err
		}
		forms = append(forms, form)
	}

	return forms, nil
}

func GetAllSubmissions(c *fiber.Ctx) error {
	jsonData, err := getAllData()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve forms"})
	}

	return c.SendString(jsonData)
}

func Submit(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	message := c.FormValue("message")

	validationErrors := fiber.Map{}

	ok, msg := validator.String(name)
	if !ok {
		validationErrors["name"] = msg
	}

	ok, msg = validator.Email(email)
	if !ok {
		validationErrors["email"] = msg
	}

	ok, msg = validator.Required(message, "message")

	if !ok {
		validationErrors["message"] = msg
	}
	if len(validationErrors) != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	// _, err := submitForm(name, email, message)

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to submit form"})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "submitted the form successfully"})

}

// func submitForm(name, email, message string) (int64, error) {
// 	result, err := database.Database.Exec(`
// 		INSERT INTO forms (name, email, message) VALUES (?, ?, ?);
// 	`, name, email, message)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return result.LastInsertId()
// }
