package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}


func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	r.GET("/reviewlist", h.GetAllReviews)

	r.GET("/list/:vendorid", h.FindReviewByVendorId)

	r.POST("/register", h.CreateReview)

	return r.Run(address)
}
