package domain

//domain 層
//Entity の定義を行う
type User struct {
	ID       int    `gorm:"primary_key auto_increment"`
	Name     string `json:"name"`
	Mail     string `json:"mail" gorm:"unique"`
	Password string `json:"password"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt time.Time
}

type NameAndPassword struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
