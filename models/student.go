package models

type Student struct {
	ID    string  `json:"id" gorm:"primaryKey" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Major string  `json:"major"`
	GPA   float64 `json:"gpa" binding:"required,gte=0,lte=4"`
}