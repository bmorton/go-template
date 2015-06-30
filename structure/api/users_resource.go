package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/interval-io/backend/db"
	"github.com/jinzhu/gorm"
)

// UsersResource is the struct responsible for listing, detailing, creating,
// updating, and deleting users.
type UsersResource struct {
	db gorm.DB
}

// NewUsersResource creates a new UsersResource given a gorm database
func NewUsersResource(db gorm.DB) *UsersResource {
	return &UsersResource{db: db}
}

// Index is the action mounted at `GET /users` for listing all users
func (r *UsersResource) Index(c *gin.Context) {
	var records []*db.UserRecord

	if err := r.db.Find(&records).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var reps []*UserRepresentation
	for _, record := range records {
		reps = append(reps, NewUserRepresentation(record))
	}

	c.JSON(http.StatusOK, reps)
}

// Create is the action mounted at `POST /users` for creating a user
func (r *UsersResource) Create(c *gin.Context) {
	submission := &UserSubmission{}
	c.Bind(submission)
	record := submission.Record()

	if err := r.db.Create(record).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, NewUserRepresentation(record))
}

// Create is the action mounted at `GET /users/:userID` for getting a user
func (r *UsersResource) Show(c *gin.Context) {
	var record db.UserRecord
	if err := r.db.Where(&db.UserRecord{ID: c.Params.ByName("userID")}).First(&record).Error; err == gorm.RecordNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, NewUserRepresentation(&record))
}

// Update is the action mounted at `PATCH /users/:userID` for updating a user
func (r *UsersResource) Update(c *gin.Context) {
	id := c.Params.ByName("userID")

	var record db.UserRecord
	var submission UserSubmission
	c.Bind(&submission)
	if err := r.db.Where(&db.UserRecord{ID: id}).First(&record).Error; err == gorm.RecordNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	submission.Apply(&record)
	if err := r.db.Save(&record).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, NewUserRepresentation(&record))
}

// Destroy is the action mounted at `DELETE /users/:userID` for deleting a user
func (t *UsersResource) Destroy(c *gin.Context) {
	var record db.UserRecord
	if err := t.db.Where(&db.UserRecord{ID: c.Params.ByName("userIDxs")}).First(&record).Error; err == gorm.RecordNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := t.db.Delete(&record).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.String(http.StatusNoContent, "")
}
