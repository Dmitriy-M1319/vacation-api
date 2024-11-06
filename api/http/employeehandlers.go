package http

import (
	"net/http"
	"strconv"

	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	"github.com/gin-gonic/gin"
)

type EmployeeHandlers struct {
	repo interfaces.EmployeeRepository
}

func NewEmployeeHandlers(r interfaces.EmployeeRepository) *EmployeeHandlers {
	return &EmployeeHandlers{repo: r}
}

func (h *EmployeeHandlers) GetAll(c *gin.Context) {
	employeers, err := h.repo.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, employeers)
}

func (h *EmployeeHandlers) GetById(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	employee, err := h.repo.GetById(uint64(idInt))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, employee)
}

func (h *EmployeeHandlers) Insert(c *gin.Context) {
	var newEmployee models.Employee

	if err := c.BindJSON(&newEmployee); err != nil {
		return
	}

	err := h.repo.Insert(&newEmployee)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, newEmployee)
}

func (h *EmployeeHandlers) Update(c *gin.Context) {
	var updatedEmployee models.Employee
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	if err := c.BindJSON(&updatedEmployee); err != nil {
		return
	}

	err := h.repo.Update(uint64(idInt), &updatedEmployee)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedEmployee)
}

func (h *EmployeeHandlers) Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	err := h.repo.Delete(uint64(idInt))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{})
}
