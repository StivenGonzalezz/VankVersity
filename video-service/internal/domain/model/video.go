package model
type Video struct{
	ID uint `gorm:"primaryKey"`
	URL string
	Name string
	ProfessorID uint
	CreatedAt string
}