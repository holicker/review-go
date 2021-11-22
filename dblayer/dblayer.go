package dblayer

import (
	"errors"
	"github.com/holicker/review-backend/models"
)

type DBLayer interface {
	GetAllReviews() ([]models.Review, error)
}