package models

import "time"

type Image struct{
	ImageID int `gorm:"primary_key" json:"image_id"`
	ImageName string `json:"image_name"`
	Imagefile string `json:"imagefile"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	AlbumId int `json: album_id`
}

//type Image struct{
//	ImageID int `gorm:"primary_key" json:"image_id"`
//	ImageName string `json:"image_name"`
//	CreatedAt time.Time `json:"createdat"`
//	UpdatedAt time.Time `json:"updatedat"`
//	PulblicVisibility bool `sql:"DEFAULT:false" json:"image_pulblicvisibility"`
//	AlbumId int `json: album_id`
//}
