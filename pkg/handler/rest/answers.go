package rest

import (
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createAnswer(c *gin.Context) {
	var input models.Answer

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	answer, err := h.service.CreateAnswer(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, answer)
}

func (h *Handler) getAnswer(c *gin.Context) {
	key := c.Query("key")

	answer, err := h.service.GetAnswer(key)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, answer)
}

func (h *Handler) updateAnswer(c *gin.Context) {
	key := c.Query("key")

	var input models.AnswerUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	answer, err := h.service.UpdateAnswer(key, input.Value)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, answer)
}

func (h *Handler) deleteAnswer(c *gin.Context) {
	key := c.Query("key")

	err := h.service.DeleteAnswer(key)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}