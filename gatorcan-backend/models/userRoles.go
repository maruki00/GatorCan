package models

type userRole string

const (
	Admin      userRole = "admin"
	Instructor userRole = "instructor"
	Student    userRole = "student"
	TA         userRole = "teaching_assistant"
)
