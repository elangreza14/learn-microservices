// Code generated by sqlc. DO NOT EDIT.

package db

import ()

type User struct {
	ID        int32  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
