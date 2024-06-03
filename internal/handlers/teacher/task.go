package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (t *TeacherHandler) CreateTask(ctx *fiber.Ctx) (err error) {
	token := ctx.Locals("user").(string)

	teacherID, err := t.teacherSvc.GetTeacherIDByUserID(token)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var req dto.TaskRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	task := entities.Task{
		ClassID:     req.ClassID,
		SubjectID:   req.SubjectID,
		TeacherID:   teacherID,
		Title:       req.Title,
		TypeOfTask:  req.TypeOfTask,
		Description: req.Description,
		Deadline:    req.Deadline,
		Link:        req.Link,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = t.teacherSvc.CreateTask(&task)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	taskDetails, err := t.teacherSvc.GetTask(task.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.TaskResponse{
		ID:          taskDetails.ID,
		ClassName:   taskDetails.Class.Name,
		SubjectName: taskDetails.Subject.Name,
		TeacherName: taskDetails.Teacher.Name,
		Title:       taskDetails.Title,
		TypeOfTask:  taskDetails.TypeOfTask,
		Description: taskDetails.Description,
		Deadline:    taskDetails.Deadline,
		Link:        taskDetails.Link,
		CreatedAt:   taskDetails.CreatedAt,
		UpdatedAt:   taskDetails.UpdatedAt,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
		"data":    response,
	})
}

func (t *TeacherHandler) GetAllTask(ctx *fiber.Ctx) (err error) {
	userID := ctx.Locals("user").(string)

	task, err := t.teacherSvc.GetAllTasks(userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var tasks []dto.TaskResponse
	for _, task := range task {
		taskRes := dto.TaskResponse{
			ID:          task.ID,
			ClassName:   task.Class.Name,
			SubjectName: task.Subject.Name,
			TeacherName: task.Teacher.Name,
			Title:       task.Title,
			TypeOfTask:  task.TypeOfTask,
			Description: task.Description,
			Deadline:    task.Deadline,
			Link:        task.Link,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		}
		tasks = append(tasks, taskRes)

	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Tasks fetched successfully",
		"data":    tasks,
	})
}

func (t *TeacherHandler) GetAllStudentAssignment(ctx *fiber.Ctx) (err error) {
	taskID := ctx.Params("taskID")

	studentAssignments, err := t.teacherSvc.GetStudentTaskAssignment(taskID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var studentAssignmentsRes []dto.StudentAssignmentResponse
	for _, studentAssignment := range studentAssignments {
		studentAssignmentRes := dto.StudentAssignmentResponse{
			ID:         studentAssignment.ID,
			Task:       studentAssignment.Task.Title,
			Student:    studentAssignment.Student.Name,
			Submission: studentAssignment.Submission,
			Grade:      studentAssignment.Grade,
			Feedback:   studentAssignment.Feedback,
			SubmitAt:   studentAssignment.SubmitAt,
		}
		studentAssignmentsRes = append(studentAssignmentsRes, studentAssignmentRes)
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student assignments fetched successfully",
		"data":    studentAssignmentsRes,
	})
}

func (t *TeacherHandler) UpdateStudentTaskAssignment(ctx *fiber.Ctx) (err error) {
	assignmentID := ctx.Params("assignmentID")

	var req dto.UpdateStudentTaskAssignmentRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = t.teacherSvc.UpdateStudentTaskAssignment(assignmentID, req.Grade, req.Feedback)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student assignment updated successfully",
	})
}
