package handler

import (
	"console-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategoryH(c *gin.Context) {
	var input model.Object
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := h.service.ObjectService(input)
	c.JSON(http.StatusOK, res)
}
