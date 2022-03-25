package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getEvents(c *gin.Context) {
	key := c.Query("key")

	events, err := h.service.GetEvents(key)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "no data")
		return
	}

	c.JSON(http.StatusOK, events)
}