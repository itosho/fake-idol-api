package main

import "time"

type Idol struct {
	Id        int64      `gorm:"type:int(10) unsigned auto_increment;primary_key" json:"id"`
	Name      string     `gorm:"type:varchar(32);not null" json:"name"`
	Age       int64      `gorm:"type:int(10) unsigned;not null;default:17" json:"age"`
	Profile   string     `gorm:"type:varchar(255);not null" json:"profile"`
	CreatedAt time.Time  `gorm:"type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:datetime;not null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:datetime;null" json:"-"`
}
