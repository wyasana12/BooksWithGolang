package bookcontroller

import (
	"encoding/json"
	"errors"
	"golang_belajar/config"
	"golang_belajar/helper"
	"golang_belajar/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	var BookResponse []models.BookResponse

	if err := config.DB.Joins("Author").Find(&books).Find(&BookResponse).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List Books", &BookResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var books models.Book

	if err := json.NewDecoder(r.Body).Decode(&books); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	//Check Author
	var author models.Author
	if err := config.DB.First(&author, books.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&books).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Create Book", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book
	var BookResponse models.BookResponse

	if err := config.DB.Joins("Author").First(&book, id).First(&BookResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail Book", BookResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book Not Found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var bookPayload models.Book
	if err := json.NewDecoder(r.Body).Decode(&bookPayload); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var author models.Author
	if bookPayload.AuthorID != 0 {
		if err := config.DB.First(&author, bookPayload.AuthorID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Author Not Found", nil)
				return
			}

			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Where("id = ?", id).Updates(&bookPayload).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Updated Book", nil)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book
	res := config.DB.Delete(&book, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Book Not Found", nil)
		return
	}

	helper.Response(w, 200, "Success Delete Book", nil)
}
