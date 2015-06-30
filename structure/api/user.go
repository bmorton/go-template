package api

import (
	"time"

	"github.com/guregu/null"
	"github.com/interval-io/backend/db"
)

// UserRepresentation is the struct that is encoded as JSON and returned for
// UsersResource requests.
type UserRepresentation struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUserRepresentation transforms a UserRecord into its JSON-encodeable
// representation.
func NewUserRepresentation(r *db.UserRecord) *UserRepresentation {
	return &UserRepresentation{
		ID:        r.ID,
		Name:      r.Name,
		Email:     r.Email,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

// UserSubmission is the struct that is used for user input in Create and Update.
type UserSubmission struct {
	Name  null.String `json:"name"`
	Email null.String `json:"email"`
}

// Record returns a database-friendly struct of the input
func (s *UserSubmission) Record() *db.UserRecord {
	return &db.UserRecord{
		Name:  s.Name,
		Email: s.Email,
	}
}

func (u *UserSubmission) Apply(r *db.UserRecord) {
	if u.Name.Valid {
		r.Name = u.Name
	}

	if u.Email.Valid {
		r.Email = u.Email
	}
}
