package main

import (
	"time"
)

type Task struct {
	id          int32
	title       string
	description string
	created_at  time.Time
	updated_at  time.Time
	user_name   User
	status      Status
}

type User struct {
	id         int32
	first_name string
	last_name  string
	email      string
}

type Status struct {
	id    int32
	title string
}
