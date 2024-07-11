package handlers

import (
	"net/http"
	"time"

	"github.com/Magetan-Boyz/Backend/internal/domain/dto"
	"github.com/Magetan-Boyz/Backend/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

func (t *TeacherHandler) CreateTask(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
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

	// parse in location
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// parsing Deadline to DateTime
	deadline, err := time.ParseInLocation(time.DateTime, req.Deadline, loc)
	if err != nil {
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
		Deadline:    deadline,
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
		Deadline:    taskDetails.Deadline.Format(time.DateTime),
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
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
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
			Deadline:    task.Deadline.Format(time.DateTime),
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
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
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
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
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

func (t *TeacherHandler) GetStudentTaskAssignmentDetail(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	assignmentID := ctx.Params("assignmentID")

	studentAssignment, err := t.teacherSvc.GetStudentTaskAssignmentDetail(assignmentID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	studentAssignmentRes := dto.StudentAssignmentResponse{
		ID:         studentAssignment.ID,
		Task:       studentAssignment.Task.Title,
		Student:    studentAssignment.Student.Name,
		Submission: studentAssignment.Submission,
		Grade:      studentAssignment.Grade,
		Feedback:   studentAssignment.Feedback,
		SubmitAt:   studentAssignment.SubmitAt,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student assignment fetched successfully",
		"data":    studentAssignmentRes,
	})
}

func (t *TeacherHandler) GetTask(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	taskID := ctx.Params("taskID")

	task, err := t.teacherSvc.GetTask(taskID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	taskRes := dto.TaskResponse{
		ID:          task.ID,
		ClassName:   task.Class.Name,
		SubjectName: task.Subject.Name,
		TeacherName: task.Teacher.Name,
		Title:       task.Title,
		TypeOfTask:  task.TypeOfTask,
		Description: task.Description,
		Deadline:    task.Deadline.Format(time.DateTime),
		Link:        task.Link,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Task fetched successfully",
		"data":    taskRes,
	})
}

func (t *TeacherHandler) UpdateTask(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	taskID := ctx.Params("taskID")

	var req dto.TaskRequest
	if err = ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	loc, err := time.LoadLocation("Local")
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	deadline, err := time.ParseInLocation(time.DateTime, req.Deadline, loc)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	task := entities.Task{
		ClassID:     req.ClassID,
		SubjectID:   req.SubjectID,
		Title:       req.Title,
		TypeOfTask:  req.TypeOfTask,
		Description: req.Description,
		Deadline:    deadline,
		Link:        req.Link,
		UpdatedAt:   time.Now(),
	}

	err = t.teacherSvc.UpdateTask(taskID, &task)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	taskDetails, err := t.teacherSvc.GetTask(taskID)
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
		Deadline:    taskDetails.Deadline.Format(time.DateTime),
		Link:        taskDetails.Link,
		CreatedAt:   taskDetails.CreatedAt,
		UpdatedAt:   taskDetails.UpdatedAt,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Task updated successfully",
		"data":    response,
	})
}

func (t *TeacherHandler) DeleteTask(ctx *fiber.Ctx) (err error) {
	if ctx.Locals("testMode").(bool) {
		return ctx.JSON(fiber.Map{"message": "DB still the same"})
	}
	taskID := ctx.Params("taskID")

	err = t.teacherSvc.DeleteTask(taskID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
