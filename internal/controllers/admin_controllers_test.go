package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Magetan-Boyz/Backend/internal/controllers"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// MockAdminService is a mock implementation of AdminService
type MockAdminService struct{}

// LogIn mocks the login functionality
func (m *MockAdminService) LogIn(username, password string) (*entities.Admin, error) {
	admin := &entities.Admin{
		User: entities.User{
			ID:       "123",
			Username: username,
			Password: password,
		},
	}
	return admin, nil
}

// CreateAdminToken mocks the token creation functionality
func (m *MockAdminService) CreateAdminToken(admin *entities.Admin) (string, error) {
	return "mocked-token", nil
}

// GetAdminByToken mocks the method to retrieve admin by token
func (m *MockAdminService) GetAdminByToken(token string) (*entities.Admin, error) {
	admin := &entities.Admin{
		User: entities.User{
			ID:       "123",
			Username: "testuser",
		},
	}
	return admin, nil
}

// CreateSubject mocks the method to create a subject
func (m *MockAdminService) CreateSubject(subject *entities.Subject) error {
	return nil
}

func TestLogin_Positive(t *testing.T) {
	app := fiber.New()

	mockAdminService := &MockAdminService{}

	controller := controllers.NewAdminController(mockAdminService)

	app.Post("/login", controller.Login)

	reqBody := dto.AdminLoginRequest{Username: "testuser", Password: "testpassword"}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response dto.AdminLoginResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "123", response.ID)
	assert.Equal(t, "testuser", response.Username)
	assert.Equal(t, "mocked-token", response.Token)
}

func TestLogin_Negative(t *testing.T) {
	app := fiber.New()

	mockAdminService := &MockAdminService{}

	controller := controllers.NewAdminController(mockAdminService)

	app.Post("/login", controller.Login)

	// Test with empty request body
	reqBody := []byte{}
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Add more negative test cases for error scenarios if needed
}

func TestCreateSubject_Success(t *testing.T) {
	app := fiber.New()

	mockAdminService := &MockAdminService{}

	controller := controllers.NewAdminController(mockAdminService)

	app.Post("/create-subject", controller.CreateSubject)

	reqBody := dto.SubjectRequest{Name: "Math", Description: "Mathematics subject", Semester: "Spring"}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/create-subject", bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Subject created successfully", response["message"])

	// Check if the data field is not nil
	assert.NotNil(t, response["data"])
	log.Println(response["data"])
}

func TestCreateSubject_BadRequest(t *testing.T) {
	app := fiber.New()

	mockAdminService := &MockAdminService{}

	controller := controllers.NewAdminController(mockAdminService)

	app.Post("/create-subject", controller.CreateSubject)

	// Test with empty request body
	reqBody := []byte{}
	req := httptest.NewRequest(http.MethodPost, "/create-subject", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Read the response body content
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the response body content bytes to a string and print it
	bodyString := string(bodyBytes)
	log.Println(bodyString)

	// Add more negative test cases for error scenarios if needed
}
