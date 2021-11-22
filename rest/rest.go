package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	r := gin.Default()
	h, _ := NewHandler()

	r.GET("/reviewlist", h.GetAllReviews)

	r.GET("/list/:vendorid", h.FindReviewByVendorId)

	r.POST("/register", h.CreateReview)
}
