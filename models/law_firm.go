package models

type Firm struct {
	FirmID    int     `json:"firm_id" gorm:"primary_key;auto_increment:true"`
	FirmName  string  `json:"firm_name" gorm:"type:varchar(60)"`
	Address   string  `json:"address" gorm:"type:varchar(150)"`
	Telp      string  `json:"telp" gorm:"type:varchar(20)"`
	Email     string  `json:"email" gorm:"type:varchar(60)"`
	Latitude  float64 `json:"latitude" gorm:"type:float8"`
	Longitude float64 `json:"longitude" gorm:"type:float8"`
	IsActive  bool    `json:"is_active" gorm:"type:boolean"`
	Model
}
