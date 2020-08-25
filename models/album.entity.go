package models

import "time"

type Album struct{
	AlbumID int `gorm:"primary_key" json:"album_id"`
	AlbumName string `json:"album_name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	Images []Image `gorm:"foreignkey:AlbumId" json:"images" `
	UserId int `json: "user_id"`
	AlbumThumbnail string `json:"album_thumbnail"`
}
//type Album struct{
//	AlbumID int `gorm:"primary_key" json:"album_id"`
//	AlbumName string `json:"album_name"`
//	CreatedAt time.Time `json:"createdat"`
//	UpdatedAt time.Time `json:"updatedat"`
//	PulblicVisibility bool `sql:"DEFAULT:false" json:"pulblicvisibility"`
//	Images []Image `gorm:"foreignkey:AlbumId" json:"images" `
//}