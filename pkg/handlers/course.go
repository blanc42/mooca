package handlers

import (
	"net/http"
	"strconv"

	"github.com/blanc42/mooca/pkg/models"
	"github.com/blanc42/mooca/pkg/usecases"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	usecase usecases.CourseUsecaseInterface
}

func NewCourseHandler(usecase usecases.CourseUsecaseInterface) *CourseHandler {
	return &CourseHandler{usecase: usecase}
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateCourse(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := h.usecase.GetCourseByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": course})
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.usecase.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.UpdateCourse(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
