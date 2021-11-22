package dblayer

import (
	"review-go/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con+"?charset=utf8&parseTime=True&loc=Local")
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllReviews() (reviews []models.Review, err error){
	return reviews, db.Find(&reviews).Error
}

func (db *DBORM) CreateReview(review models.Review) (models.Review, error){
	return review, db.Create(&review).Error
}

func (db *DBORM) FindReviewByVendorId(vendorid int) (reviews []models.Review, err error){
	return reviews, db.Where(&models.Review{VendorId:vendorid}).Find(&reviews).Error
}