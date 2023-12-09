package model

// Model User
type User struct {
    ID            int64  `gorm:"primaryKey" json:"id"`
    Email         string `gorm:"type:varchar(100)" json:"email"`
    Password      string `json:"-"`
    HousingAreaID int64  `json:"housingAreaId"`
    HousingArea   HousingArea `gorm:"foreignKey:HousingAreaID;references:ID"`
}