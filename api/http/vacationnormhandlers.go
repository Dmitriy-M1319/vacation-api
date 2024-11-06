package http

import (
	"net/http"
	"strconv"

	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	"github.com/gin-gonic/gin"
)

type VacationNormHandlers struct {
	repo interfaces.VacationNormRepository
}

func NewVacationNormHandlers(r interfaces.VacationNormRepository) *VacationNormHandlers {
	return &VacationNormHandlers{repo: r}
}

func (h *VacationNormHandlers) GetAll(c *gin.Context) {
	norms, err := h.repo.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, norms)
}

func (h *VacationNormHandlers) GetById(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	norm, err := h.repo.GetById(uint64(idInt))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, norm)
}

func (h *VacationNormHandlers) Insert(c *gin.Context) {
	var newVacationNorm models.VacationNorm

	if err := c.BindJSON(&newVacationNorm); err != nil {
		return
	}

	err := h.repo.Insert(&newVacationNorm)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, newVacationNorm)
}

func (h *VacationNormHandlers) Update(c *gin.Context) {
	var updatedVacationNorm models.VacationNorm
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	if err := c.BindJSON(&updatedVacationNorm); err != nil {
		return
	}

	err := h.repo.Update(uint64(idInt), &updatedVacationNorm)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedVacationNorm)
}

func (h *VacationNormHandlers) Delete(c *gin.Context) {
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
