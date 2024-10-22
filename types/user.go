package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 8
	minEmailLen     = 6
	maxEmailLen     = 255
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() []string {
	validationErrors := []string{}
	if len(params.FirstName) < minFirstNameLen {
		validationErrors = append(validationErrors, fmt.Sprintf("firstName must be at least %d characters", minFirstNameLen))
	}
	if len(params.LastName) < minLastNameLen {
		validationErrors = append(validationErrors, fmt.Sprintf("lastName must be at least %d characters", minLastNameLen))
	}
	if !isEmailValid(params.Email) {
		validationErrors = append(validationErrors, fmt.Sprintf("email is not valid"))
	}
	if len(params.Password) < minPasswordLen {
		validationErrors = append(validationErrors, fmt.Sprintf("password must be at least %d characters", minPasswordLen))
	}

	if validationErrors != nil {
		return validationErrors
	}
	return nil
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  string(encpw),
	}, nil
}
