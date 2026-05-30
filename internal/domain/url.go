package domain

type URL struct {
	ID          uint   `gorm:"primaryKey"`
	OriginalUrl string `gorm:"column:original_url"`
	Codigo      string `gorm:"column:codigo"`
}
