package models

import "errors"

type SignupRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
}

func (r *SignupRequest) Validate() error {
	if r.Email == nil || r.Name == nil || r.Password == nil {
		return errors.New("email and name and password are required")
	}
	return nil
}
