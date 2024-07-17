package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/Magetan-Boyz/Backend/internal/middleware"
	"github.com/Magetan-Boyz/Backend/internal/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAdminHandler_CreateTeacher(t *testing.T) {
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
			name:    "Positive case: Successfully create teacher",
			reqBody: `{"username":"testuser","password":"testpass","name":"Test Name","email":"test@example.com"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().CreateTeacher(gomock.Any()).Return(nil)
			},
			wantStatus: http.StatusCreated,
			wantBody:   `"message":"Teacher created successfully"`,
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
			name:    "Negative case: Service error",
			reqBody: `{"username":"testuser","password":"testpass","name":"Test Name","email":"test@example.com"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().CreateTeacher(gomock.Any()).Return(errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/teacher", func(c *fiber.Ctx) error {
				// Set the testMode local to avoid the panic
				c.Locals("testMode", false)
				return handler.CreateTeacher(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodPost, "/teacher", bytes.NewBufferString(tt.reqBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer test-token") // Set the Authorization header

			resp, err := app.Test(req, -1)
			require.NoError(t, err)

			require.Equal(t, tt.wantStatus, resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			require.Contains(t, string(body), tt.wantBody)
		})
	}
}

func TestAdminHandler_GetAllTeacher(t *testing.T) {
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
		mockSetup  func()
		wantStatus int
		wantBody   string
	}{
		{
			name: "Positive case: Successfully get all teachers",
			mockSetup: func() {
				mockAdminService.EXPECT().GetAllTeacher().Return([]entities.Teacher{
					{User: entities.User{ID: "1", Username: "user1"}, ID: "1", Name: "Teacher1", Email: "teacher1@example.com"},
					{User: entities.User{ID: "2", Username: "user2"}, ID: "2", Name: "Teacher2", Email: "teacher2@example.com"},
				}, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"data":[{"id":"1","name":"Teacher1","email":"teacher1@example.com","is_homeroom_teacher":false,"is_councelor":false},{"id":"2","name":"Teacher2","email":"teacher2@example.com","is_homeroom_teacher":false,"is_councelor":false}]`,
		},
		{
			name: "Negative case: Service error",
			mockSetup: func() {
				mockAdminService.EXPECT().GetAllTeacher().Return(nil, errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/teachers", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.GetAllTeacher(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, "/teachers", nil)
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

func TestAdminHandler_AssignTeacherToSubject(t *testing.T) {
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
			name:    "Positive case: Successfully assign teacher to subject",
			reqBody: `{"teacher_id":["1"]}`,
			mockSetup: func() {
				mockAdminService.EXPECT().AssignTeacherToSubject([]string{"1"}, "123").Return(nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"Teacher assigned to subject successfully"`,
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
			name:    "Negative case: Service error",
			reqBody: `{"teacher_id":["1"]}`,
			mockSetup: func() {
				mockAdminService.EXPECT().AssignTeacherToSubject([]string{"1"}, "123").Return(errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/subject/:id/assign", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.AssignTeacherToSubject(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodPost, "/subject/123/assign", bytes.NewBufferString(tt.reqBody))
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

func TestAdminHandler_GetTeachersBySubjectID(t *testing.T) {
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
		subjectID  string
		mockSetup  func()
		wantStatus int
		wantBody   string
	}{
		{
			name:      "Positive case: Successfully get teachers by subject ID",
			subjectID: "123",
			mockSetup: func() {
				teachers := []dto.TeacherSubjectResponse{
					{TeacherID: "1", TeacherName: "Teacher1", SubjectID: "123", SubjectName: "Subject1"},
					{TeacherID: "2", TeacherName: "Teacher2", SubjectID: "123", SubjectName: "Subject2"},
				}
				mockAdminService.EXPECT().GetTeachersBySubjectID("123").Return(teachers, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"Teachers fetched successfully"`,
		},
		{
			name:      "Negative case: Service error",
			subjectID: "123",
			mockSetup: func() {
				mockAdminService.EXPECT().GetTeachersBySubjectID("123").Return(nil, errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/subject/:id/teachers", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.GetTeachersBySubjectID(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, "/subject/"+tt.subjectID+"/teachers", nil)
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

func TestAdminHandler_GetTeacherSubjects(t *testing.T) {
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
		teacherID  string
		mockSetup  func()
		wantStatus int
		wantBody   string
	}{
		{
			name:      "Positive case: Successfully get teacher subjects",
			teacherID: "123",
			mockSetup: func() {
				subjects := []dto.TeacherSubjectsResponse{
					{
						TeacherID:   "123",
						TeacherName: "Teacher1",
						Subjects: []dto.SubjectResponse{
							{ID: "1", Name: "Math", Description: "Mathematics", Semester: "Spring"},
						},
					},
				}
				mockAdminService.EXPECT().GetTeacherSubjects("123").Return(subjects, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"Subjects fetched successfully"`,
		},
		{
			name:      "Negative case: Service error",
			teacherID: "123",
			mockSetup: func() {
				mockAdminService.EXPECT().GetTeacherSubjects("123").Return(nil, errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/teacher/:id/subjects", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.GetTeacherSubjects(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, "/teacher/"+tt.teacherID+"/subjects", nil)
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

func TestAdminHandler_GetTeachersByClassAndSubject(t *testing.T) {
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
		subjectID  string
		mockSetup  func()
		wantStatus int
		wantBody   string
	}{
		{
			name:      "Positive case: Successfully get teachers by class and subject ID",
			classID:   "class123",
			subjectID: "subject456",
			mockSetup: func() {
				teachers := []dto.TeacherSubjectResponse{
					{TeacherID: "1", TeacherName: "Teacher1", SubjectID: "subject456", SubjectName: "Subject1"},
					{TeacherID: "2", TeacherName: "Teacher2", SubjectID: "subject456", SubjectName: "Subject2"},
				}
				mockAdminService.EXPECT().GetTeachersByClassAndSubject("class123", "subject456").Return(teachers, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"Teachers fetched successfully"`,
		},
		{
			name:      "Negative case: Missing query parameters",
			classID:   "",
			subjectID: "",
			mockSetup: func() {
				// No setup required as the query parameter validation will fail
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"classID and subjectID query parameters are required"`,
		},
		{
			name:      "Negative case: Service error",
			classID:   "class123",
			subjectID: "subject456",
			mockSetup: func() {
				mockAdminService.EXPECT().GetTeachersByClassAndSubject("class123", "subject456").Return(nil, errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/teachers", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.GetTeachersByClassAndSubject(c)
			})

			tt.mockSetup()

			url := "/teachers"
			if tt.classID != "" && tt.subjectID != "" {
				url += "?classID=" + tt.classID + "&subjectID=" + tt.subjectID
			}
			req := httptest.NewRequest(http.MethodGet, url, nil)
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
