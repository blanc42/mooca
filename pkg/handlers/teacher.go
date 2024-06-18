package handlers

import (
	"net/http"
	"strconv"

	"github.com/blanc42/mooca/pkg/models"
	"github.com/blanc42/mooca/pkg/usecases"
	"github.com/gin-gonic/gin"
)

type TeacherHandler struct {
	usecase usecases.TeacherUsecaseInterface
}

func NewTeacherHandler(usecase usecases.TeacherUsecaseInterface) *TeacherHandler {
	return &TeacherHandler{usecase: usecase}
}

func (th *TeacherHandler) CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := th.usecase.CreateTeacher(&teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, teacher)
}

func (th *TeacherHandler) UpdateTeacher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	var teacher models.Teacher
	if err := c.BindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	teacher.Id = uint(id)

	err = th.usecase.UpdateTeacher(&teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func (th *TeacherHandler) DeleteTeacher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	err = th.usecase.DeleteTeacher(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (th *TeacherHandler) GetTeacherByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	teacher, err := th.usecase.GetTeacherByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func (th *TeacherHandler) GetAllTeachers(c *gin.Context) {
	teachers, err := th.usecase.GetAllTeachers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teachers)
}
