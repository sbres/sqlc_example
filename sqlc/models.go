// Code generated by sqlc. DO NOT EDIT.

package sqlc

import ()

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}