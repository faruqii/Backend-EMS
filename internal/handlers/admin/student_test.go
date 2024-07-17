package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAdminHandler_CreateStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminService := mocks.NewMockAdminService(ctrl)
	mockMiddlewareManager := new(middleware.Middleware)

	handler := &AdminHandler{
		adminService:      mockAdminService,
		middlewareManager: *mockMiddlewareManager,
	}

	tests := []struct {
		name       string
		reqBody    string
		mockSetup  func()
		wantStatus int
		wantBody   string
	}{
		{
			name:    "Positive case: Successfully create student",
			reqBody: `{"username":"student1","password":"pass123","name":"John Doe","nisn":"123456","gender":"male","address":"123 Street","birthplace":"City","birthdate":"2000-01-01","province":"Province","city":"City","blood_type":"O","religion":"Religion","phone":"123456789","parent_phone":"987654321","email":"student@example.com"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().CreateStudent(gomock.Any()).Return(nil)
			},
			wantStatus: http.StatusCreated,
			wantBody:   `"message":"Student created successfully"`,
		},
		{
			name:    "Negative case: Body parsing error",
			reqBody: `invalid body`,
			mockSetup: func() {
				// No setup required as the body parsing itself will fail
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"invalid character 'i' looking for beginning of value"`,
		},
		{
			name:    "Negative case: Username already exists",
			reqBody: `{"username":"student1","password":"pass123","name":"John Doe","nisn":"123456","gender":"male","address":"123 Street","birthplace":"City","birthdate":"2000-01-01","province":"Province","city":"City","blood_type":"O","religion":"Religion","phone":"123456789","parent_phone":"987654321","email":"student@example.com"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().CreateStudent(gomock.Any()).Return(errors.New("username already exists"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"username already exists"`,
		},
		{
			name:    "Negative case: Service error",
			reqBody: `{"username":"student1","password":"pass123","name":"John Doe","nisn":"123456","gender":"male","address":"123 Street","birthplace":"City","birthdate":"2000-01-01","province":"Province","city":"City","blood_type":"O","religion":"Religion","phone":"123456789","parent_phone":"987654321","email":"student@example.com"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().CreateStudent(gomock.Any()).Return(errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/student", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.CreateStudent(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodPost, "/student", bytes.NewBufferString(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer test-token")

			resp, err := app.Test(req, -1)
			require.NoError(t, err)

			require.Equal(t, tt.wantStatus, resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			require.Contains(t, string(body), tt.wantBody)
		})
	}
}

func TestAdminHandler_GetAllStudents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminService := mocks.NewMockAdminService(ctrl)
	mockMiddlewareManager := new(middleware.Middleware)

	handler := &AdminHandler{
		adminService:      mockAdminService,
		middlewareManager: *mockMiddlewareManager,
	}

	tests := []struct {
		name        string
		classPrefix string
		mockSetup   func()
		wantStatus  int
		wantBody    string
	}{
		{
			name:        "Positive case: Successfully get all students",
			classPrefix: "",
			mockSetup: func() {
				students := []entities.Student{
					{
						ID:          "1",
						Name:        "Student1",
						NISN:        "123456",
						Address:     "Address1",
						Birthplace:  "City1",
						Birthdate:   "2000-01-01",
						Gender:      "Male",
						Province:    "Province1",
						City:        "City1",
						BloodType:   "O",
						Religion:    "Religion1",
						Phone:       "123456789",
						ParentPhone: "987654321",
						Email:       "student1@example.com",
						Class: entities.Class{
							Name: "Class1",
						},
					},
					{
						ID:          "2",
						Name:        "Student2",
						NISN:        "654321",
						Address:     "Address2",
						Birthplace:  "City2",
						Birthdate:   "2001-01-01",
						Gender:      "Female",
						Province:    "Province2",
						City:        "City2",
						BloodType:   "A",
						Religion:    "Religion2",
						Phone:       "987654321",
						ParentPhone: "123456789",
						Email:       "student2@example.com",
						Class: entities.Class{
							Name: "Class2",
						},
					},
				}
				mockAdminService.EXPECT().GetAllStudents().Return(students, nil)
			},
			wantStatus: http.StatusOK,
			wantBody: `{
				"data": [
					{
						"id": "1",
						"name": "Student1",
						"nisn": "123456",
						"address": "Address1",
						"birthplace": "City1",
						"birthdate": "2000-01-01",
						"gender": "Male",
						"province": "Province1",
						"city": "City1",
						"blood_type": "O",
						"religion": "Religion1",
						"phone": "123456789",
						"parent_phone": "987654321",
						"email": "student1@example.com",
						"class_name": "Class1"
					},
					{
						"id": "2",
						"name": "Student2",
						"nisn": "654321",
						"address": "Address2",
						"birthplace": "City2",
						"birthdate": "2001-01-01",
						"gender": "Female",
						"province": "Province2",
						"city": "City2",
						"blood_type": "A",
						"religion": "Religion2",
						"phone": "987654321",
						"parent_phone": "123456789",
						"email": "student2@example.com",
						"class_name": "Class2"
					}
				]
			}`,
		},
		{
			name:        "Positive case: Get students by class prefix",
			classPrefix: "prefix",
			mockSetup: func() {
				students := []entities.Student{
					{
						ID:          "1",
						Name:        "Student1",
						NISN:        "123456",
						Address:     "Address1",
						Birthplace:  "City1",
						Birthdate:   "2000-01-01",
						Gender:      "Male",
						Province:    "Province1",
						City:        "City1",
						BloodType:   "O",
						Religion:    "Religion1",
						Phone:       "123456789",
						ParentPhone: "987654321",
						Email:       "student1@example.com",
						Class: entities.Class{
							Name: "Class1",
						},
					},
				}
				mockAdminService.EXPECT().FindStudentByClassPrefix("prefix").Return(students, nil)
			},
			wantStatus: http.StatusOK,
			wantBody: `{
				"data": [
					{
						"id": "1",
						"name": "Student1",
						"nisn": "123456",
						"address": "Address1",
						"birthplace": "City1",
						"birthdate": "2000-01-01",
						"gender": "Male",
						"province": "Province1",
						"city": "City1",
						"blood_type": "O",
						"religion": "Religion1",
						"phone": "123456789",
						"parent_phone": "987654321",
						"email": "student1@example.com",
						"class_name": "Class1"
					}
				]
			}`,
		},
		{
			name:        "Negative case: Service error",
			classPrefix: "",
			mockSetup: func() {
				mockAdminService.EXPECT().GetAllStudents().Return(nil, errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody: `{
				"error": "service error"
			}`,
		},
		{
			name:        "Negative case: No students found",
			classPrefix: "",
			mockSetup: func() {
				mockAdminService.EXPECT().GetAllStudents().Return([]entities.Student{}, nil)
			},
			wantStatus: http.StatusNotFound,
			wantBody: `{
				"message": "No students found"
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/students", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.GetAllStudents(c)
			})

			tt.mockSetup()

			url := "/students"
			if tt.classPrefix != "" {
				url += "?class_prefix=" + tt.classPrefix
			}

			req := httptest.NewRequest(http.MethodGet, url, nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer test-token")

			resp, err := app.Test(req, -1)
			require.NoError(t, err)

			require.Equal(t, tt.wantStatus, resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			require.JSONEq(t, tt.wantBody, string(body))
		})
	}
}

func TestAdminHandler_InsertStudentToClass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminService := mocks.NewMockAdminService(ctrl)
	mockMiddlewareManager := new(middleware.Middleware)

	handler := &AdminHandler{
		adminService:      mockAdminService,
		middlewareManager: *mockMiddlewareManager,
	}

	tests := []struct {
		name       string
		classID    string
		reqBody    string
		mockSetup  func()
		wantStatus int
		wantBody   string
	}{
		{
			name:    "Positive case: Successfully insert student to class",
			classID: "class1",
			reqBody: `{"student_id":"student1"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().InsertStudentToClass("student1", "class1").Return(&entities.Student{ID: "student1"}, nil)
			},
			wantStatus: http.StatusCreated,
			wantBody:   `{"message":"Student inserted to class successfully"}`,
		},
		{
			name:    "Negative case: Body parsing error",
			classID: "class1",
			reqBody: `invalid body`,
			mockSetup: func() {
				// No setup required as the body parsing itself will fail
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"error":"invalid character 'i' looking for beginning of value"}`,
		},
		{
			name:    "Negative case: Student already in class",
			classID: "class1",
			reqBody: `{"student_id":"student1"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().InsertStudentToClass("student1", "class1").Return(nil, errors.New("student already in class"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"error":"student already in class"}`,
		},
		{
			name:    "Negative case: Service error",
			classID: "class1",
			reqBody: `{"student_id":"student1"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().InsertStudentToClass("student1", "class1").Return(nil, errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"error":"service error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/class/:id/student", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.InsertStudentToClass(c)
			})

			tt.mockSetup()

			url := "/class/" + tt.classID + "/student"

			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBufferString(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer test-token")

			resp, err := app.Test(req, -1)
			require.NoError(t, err)

			require.Equal(t, tt.wantStatus, resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			require.JSONEq(t, tt.wantBody, string(body))
		})
	}
}

func TestAdminHandler_RemoveStudentFromClass(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminService := mocks.NewMockAdminService(ctrl)
	mockMiddlewareManager := new(middleware.Middleware)

	handler := &AdminHandler{
		adminService:      mockAdminService,
		middlewareManager: *mockMiddlewareManager,
	}

	tests := []struct {
		name       string
		studentID  string
		mockSetup  func()
		wantStatus int
		wantBody   string
	}{
		{
			name:      "Positive case: Successfully remove student from class",
			studentID: "student1",
			mockSetup: func() {
				mockAdminService.EXPECT().RemoveStudentFromClass("student1").Return(nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"message":"Student removed from class successfully"}`,
		},
		{
			name:      "Negative case: Service error",
			studentID: "student1",
			mockSetup: func() {
				mockAdminService.EXPECT().RemoveStudentFromClass("student1").Return(errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"error":"service error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Delete("/student/:id/class", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.RemoveStudentFromClass(c)
			})

			tt.mockSetup()

			url := "/student/" + tt.studentID + "/class"

			req := httptest.NewRequest(http.MethodDelete, url, nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer test-token")

			resp, err := app.Test(req, -1)
			require.NoError(t, err)

			require.Equal(t, tt.wantStatus, resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			require.JSONEq(t, tt.wantBody, string(body))
		})
	}
}
