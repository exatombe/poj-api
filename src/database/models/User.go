package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	DiscordID string `json:"discord_id"`
	PojCoin   int    `json:"poj_coin"`
}

func CreateUser(db *gorm.DB, user User) {
	db.Create(&user)
}

func GetUser(db *gorm.DB, id string) User {
	var user User
	db.Where("discord_id = ?", id).First(&user)
	return user
}

func UpdateUser(db *gorm.DB, user User) {
	db.Save(&user)
}

func DeleteUser(db *gorm.DB, user User) {
	db.Delete(&user)
}

func GetAllUsers(db *gorm.DB, limit int, offset int) []User {
	var users []User
	db.Limit(limit).Offset(offset).Find(&users)
	return users
}
