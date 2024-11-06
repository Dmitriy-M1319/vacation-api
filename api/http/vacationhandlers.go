package http

import (
	"net/http"
	"strconv"

	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	"github.com/gin-gonic/gin"
)

type VacationHandlers struct {
	repo interfaces.VacationRepository
}

func NewVacationHandlers(r interfaces.VacationRepository) *VacationHandlers {
	return &VacationHandlers{repo: r}
}

func (h *VacationHandlers) GetAll(c *gin.Context) {
	vacations, err := h.repo.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, vacations)
}

func (h *VacationHandlers) GetById(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	vacation, err := h.repo.GetById(uint64(idInt))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, vacation)
}

func (h *VacationHandlers) GetByEmployeeId(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	vacations, err := h.repo.GetByEmployeeId(uint64(idInt))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, vacations)
}

func (h *VacationHandlers) Insert(c *gin.Context) {
	var newVacation models.Vacation

	if err := c.BindJSON(&newVacation); err != nil {
		return
	}

	err := h.repo.Insert(&newVacation)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, newVacation)
}

func (h *VacationHandlers) Update(c *gin.Context) {
	var updatedVacation models.Vacation
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	if err := c.BindJSON(&updatedVacation); err != nil {
		return
	}

	err := h.repo.Update(uint64(idInt), &updatedVacation)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedVacation)
}

func (h *VacationHandlers) Delete(c *gin.Context) {
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
