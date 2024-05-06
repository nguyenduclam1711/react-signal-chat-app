package models

import "time"

type UserModel struct {
	Id          int
	Username    string
	CreatedTime time.Time
}
