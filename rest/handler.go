package rest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"reveiw-go/dblayer"
	"review-go/models"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	db dblayer.DBLayer
}

func NewHandler() (*Handler, error) {
	return new(Handler), nil
}

func (h *Handler) GetAllReviews(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	reviews, err := h.db.GetAllReviews()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

func (h *Handler) CreateReview(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	var review models.Review
	err := c.ShouldBindJSON(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	review, err = h.db.CreateReview(review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, review)
}

func (h *Handler) FindReviewByVendorId(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}

	p := c.Param("vendorid")
	vendorid, err := strconv.Atoi(p)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	reviews, err := h.db.FindReviewByVendorId(vendorid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}