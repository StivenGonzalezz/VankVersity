package model
type Video struct{
	ID uint `gorm:"primaryKey"`
	Name string
	Description string
	ProfessorID uint
	CreatedAt string
}