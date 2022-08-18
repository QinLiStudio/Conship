package schema

type Item struct {
	ID     int		`gorm:"primaryKey" json:"id"`
	Url    string	`json:"url"`
	Secret string	`json:"secret"`
}
