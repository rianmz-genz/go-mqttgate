package model

type HousingArea struct {
    ID      int64  `gorm:"primaryKey" json:"id"`
    Name    string `gorm:"type:varchar(100)" json:"name"`
    Code    string `gorm:"type:varchar(6)" json:"code"`
    Address string `gorm:"type:varchar(300)" json:"address"`
    Users   []User `gorm:"foreignKey:HousingAreaID"`
}