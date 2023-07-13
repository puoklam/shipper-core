package validation

import (
	"sync"

	"github.com/puoklam/shipper-core/db/models"
)

var userpool = sync.Pool{
	New: func() any {
		return &UserValidator{}
	},
}

type UserValidator struct {
	Username string `validate:"required"`
	Password string `validate:"required,min=8"`
	Email    string `validate:"required,email"`
}

func (v *UserValidator) From(u *models.User) *UserValidator {
	v.Username = u.Username
	v.Password = u.Password
	v.Email = u.Email
	return v
}

func (v *UserValidator) Done() {
	userpool.Put(v)
}

func NewUser() *UserValidator {
	return userpool.Get().(*UserValidator)
}
