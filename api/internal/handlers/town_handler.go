package handlers

import (
	"SportsClubs/internal/models"
	"SportsClubs/internal/services"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TownHandler struct {
	service *services.TownService
}

func NewTownHandler(service *services.TownService) *TownHandler {
	return &TownHandler{service: service}
}

func (h *TownHandler) GetTowns(c *gin.Context) {
	search := c.Query("search")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if limit == 0 {
		limit = 15
	}
	if page < 1 {
		page = 1
	}

	data, total, err := h.service.GetTowns(search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"towns": data,
		"total": total,
		"page":  page,
		"limit": limit,
		"pages": int(math.Ceil(float64(total) / float64(limit))),
	})
}

func (h *TownHandler) GetClubsByTownName(c *gin.Context) {
	search := c.Query("search")
	if search == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр поиска обязателен"})
		return
	}

	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if limit == 0 {
		limit = 15
	}
	if page < 1 {
		page = 1
	}

	data, total, err := h.service.GetClubsByTownName(search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"clubs": data,
		"total": total,
		"page":  page,
		"limit": limit,
		"pages": int(math.Ceil(float64(total) / float64(limit))),
	})
}

func (h *TownHandler) CreateTown(c *gin.Context) {
	var town models.Town
	if err := c.ShouldBindJSON(&town); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTown(&town); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, town)
}

func (h *TownHandler) UpdateTown(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный id города"})
		return
	}

	var updated models.Town
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	town, err := h.service.UpdateTown(uint(id), &updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, town)
}

func (h *TownHandler) DeleteTown(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный id города"})
		return
	}

	if err := h.service.DeleteTown(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *TownHandler) GetTownByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный id клуба"})
		return
	}

	town, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Город с данным id не найден"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, town)
}
