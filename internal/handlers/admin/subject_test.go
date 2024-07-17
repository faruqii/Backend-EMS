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

func TestAdminHandler_CreateSubject(t *testing.T) {
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
			name:    "Positive case: Successfully create subject",
			reqBody: `{"name":"Math","description":"Mathematics","semester":"Spring"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().CreateSubject(gomock.Any()).Return(nil)
			},
			wantStatus: http.StatusCreated,
			wantBody:   `"message":"Subject created successfully"`,
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
			reqBody: `{"name":"Math","description":"Mathematics","semester":"Spring"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().CreateSubject(gomock.Any()).Return(errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/subject", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.CreateSubject(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodPost, "/subject", bytes.NewBufferString(tt.reqBody))
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

func TestAdminHandler_GetAllSubject(t *testing.T) {
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
			name: "Positive case: Successfully get all subjects",
			mockSetup: func() {
				subjects := []entities.Subject{
					{ID: "1", Name: "Math", Description: "Mathematics", Semester: "Spring"},
					{ID: "2", Name: "Science", Description: "Science Description", Semester: "Fall"},
				}
				mockAdminService.EXPECT().GetAllSubject().Return(subjects, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"data":[{"id":"1","name":"Math","description":"Mathematics","semester":"Spring"},{"id":"2","name":"Science","description":"Science Description","semester":"Fall"}]`,
		},
		{
			name: "Negative case: Service error",
			mockSetup: func() {
				mockAdminService.EXPECT().GetAllSubject().Return(nil, errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/subjects", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.GetAllSubject(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, "/subjects", nil)
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

func TestAdminHandler_AssignSubjectToClass(t *testing.T) {
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
			name:    "Positive case: Successfully assign subject to class",
			classID: "class123",
			reqBody: `{"subject_id":"subject456","teacher_id":"teacher789"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().AssignSubjectToClass("subject456", "teacher789", "class123").Return(nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `"message":"Subject assigned to class successfully"`,
		},
		{
			name:    "Negative case: Body parsing error",
			classID: "class123",
			reqBody: `invalid body`,
			mockSetup: func() {
				// No setup required as the body parsing itself will fail
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"invalid character 'i' looking for beginning of value"`,
		},
		{
			name:    "Negative case: Service error",
			classID: "class123",
			reqBody: `{"subject_id":"subject456","teacher_id":"teacher789"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().AssignSubjectToClass("subject456", "teacher789", "class123").Return(errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `"error":"service error"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/class/:id/assign-subject", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.AssignSubjectToClass(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodPost, "/class/"+tt.classID+"/assign-subject", bytes.NewBufferString(tt.reqBody))
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

func TestAdminHandler_GetClassesSubjectsAndTeachers(t *testing.T) {
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
		subjectID   string
		mockSetup   func()
		wantStatus  int
		wantBody    string
	}{
		{
			name:        "Positive case: Successfully fetch classes, subjects, and teachers",
			classPrefix: "classPrefix123",
			subjectID:   "subject456",
			mockSetup: func() {
				classSubjects := []entities.ClassSubject{
					{
						ClassID:   "class1",
						SubjectID: "subject456",
						TeacherID: "teacher1",
						Class:     entities.Class{ID: "class1", Name: "Class 1"},
						Subject:   entities.Subject{ID: "subject456", Name: "Subject 1"},
						Teacher:   entities.Teacher{ID: "teacher1", Name: "Teacher 1"},
					},
					{
						ClassID:   "class2",
						SubjectID: "subject456",
						TeacherID: "teacher2",
						Class:     entities.Class{ID: "class2", Name: "Class 2"},
						Subject:   entities.Subject{ID: "subject456", Name: "Subject 1"},
						Teacher:   entities.Teacher{ID: "teacher2", Name: "Teacher 2"},
					},
				}
				mockAdminService.EXPECT().GetClassSubjectsByPrefixAndSubject("classPrefix123", "subject456").Return(classSubjects, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"message":"Data fetched successfully","data":[{"class_id":"class1","class_name":"Class 1","subject_id":"subject456","subject":"Subject 1","teacher_id":"teacher1","teacher":"Teacher 1"},{"class_id":"class2","class_name":"Class 2","subject_id":"subject456","subject":"Subject 1","teacher_id":"teacher2","teacher":"Teacher 2"}]}`,
		},
		{
			name:        "Negative case: Service error",
			classPrefix: "classPrefix123",
			subjectID:   "subject456",
			mockSetup: func() {
				mockAdminService.EXPECT().GetClassSubjectsByPrefixAndSubject("classPrefix123", "subject456").Return(nil, errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"error":"service error"}`,
		},
		{
			name:        "Negative case: No data found",
			classPrefix: "classPrefix123",
			subjectID:   "subject456",
			mockSetup: func() {
				mockAdminService.EXPECT().GetClassSubjectsByPrefixAndSubject("classPrefix123", "subject456").Return([]entities.ClassSubject{}, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"message":"No data found","data":[]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/classes-subjects-teachers", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.GetClassesSubjectsAndTeachers(c)
			})

			tt.mockSetup()

			url := "/classes-subjects-teachers"
			if tt.classPrefix != "" {
				url += "?classPrefix=" + tt.classPrefix
			}
			if tt.subjectID != "" {
				url += "&subjectID=" + tt.subjectID
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

func TestAdminHandler_UpdateSubject(t *testing.T) {
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
		reqBody    string
		mockSetup  func()
		wantStatus int
		wantBody   string
	}{
		{
			name:      "Positive case: Successfully update subject",
			subjectID: "123",
			reqBody:   `{"name":"Updated Subject","description":"Updated Description","semester":"2"}`,
			mockSetup: func() {
				subject := &entities.Subject{
					ID:          "123",
					Name:        "Old Subject",
					Description: "Old Description",
					Semester:    "1",
				}
				mockAdminService.EXPECT().FindSubjectByID("123").Return(subject, nil)
				mockAdminService.EXPECT().UpdateSubject("123", gomock.Any()).DoAndReturn(func(subjectID string, updatedSubject *entities.Subject) error {
					require.Equal(t, "Updated Subject", updatedSubject.Name)
					require.Equal(t, "Updated Description", updatedSubject.Description)
					require.Equal(t, "2", updatedSubject.Semester)
					return nil
				})
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"message":"Subject updated successfully","data":{"id":"123","name":"Updated Subject","description":"Updated Description","semester":"2"}}`,
		},
		{
			name:      "Negative case: Body parsing error",
			subjectID: "123",
			reqBody:   `invalid body`,
			mockSetup: func() {
				// No setup required as the body parsing itself will fail
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"error":"invalid character 'i' looking for beginning of value"}`,
		},
		{
			name:      "Negative case: Subject not found",
			subjectID: "123",
			reqBody:   `{"name":"Updated Subject","description":"Updated Description","semester":"2"}`,
			mockSetup: func() {
				mockAdminService.EXPECT().FindSubjectByID("123").Return(nil, errors.New("subject not found"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"error":"subject not found"}`,
		},
		{
			name:      "Negative case: Service error on update",
			subjectID: "123",
			reqBody:   `{"name":"Updated Subject","description":"Updated Description","semester":"2"}`,
			mockSetup: func() {
				subject := &entities.Subject{
					ID:          "123",
					Name:        "Old Subject",
					Description: "Old Description",
					Semester:    "1",
				}
				mockAdminService.EXPECT().FindSubjectByID("123").Return(subject, nil)
				mockAdminService.EXPECT().UpdateSubject("123", gomock.Any()).Return(errors.New("update failed"))
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   `{"error":"update failed"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Put("/subjects/:subjectID", func(c *fiber.Ctx) error {
				c.Locals("testMode", false)
				return handler.UpdateSubject(c)
			})

			tt.mockSetup()

			req := httptest.NewRequest(http.MethodPut, "/subjects/"+tt.subjectID, bytes.NewBufferString(tt.reqBody))
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
