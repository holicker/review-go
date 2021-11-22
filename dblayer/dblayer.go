package dblayer

import (
	"review-go/models"
)

type DBLayer interface {
	GetAllReviews() ([]models.Review, error)
	CreateReview(models.Review) (models.Review, error)
	FindReviewByVendorId(int) ([]models.Review, error)
}