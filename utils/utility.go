package utils

import (
	"albumwebapp/dto"
	"albumwebapp/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

var (
	ErrEmailNotFound = errors.New("Email não existe")
	ErrInvalidPassword = errors.New("Senha inválida")
	ErrEmptyFields = errors.New("Preencha todos os campos")
)

func ToUserEntity(r dto.UserRequest) models.User {
	return models.User{Fname: r.Fname, Lname: r.Lname,Email: r.Email,Contact: r.Contact, Address: r.Address, Password: r.Password, Create: time.Now(),}
}

func ToAlbumEntity(r dto.AlbumRequest) models.Album {
	return models.Album{AlbumName: r.AlbumName, Description: r.Description, CreatedAt: time.Now(), UpdatedAt: time.Now(), UserId: r.UserId}
}

//func ToAlbumDTO(r models.Album) dto.AlbumDTO {
//	return dto.AlbumDTO{AlbumName: r.AlbumName, Description: r.Description, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt, UserId: r.UserId, AlbumThumbnail: r.AlbumThumbnail}
//}

func ToAlbumDTOs(r []models.Album) []dto.AlbumDTO {
	dtos := make([]dto.AlbumDTO, 0)

	for _, v := range r{
		var albumdto dto.AlbumDTO
		//b := make([]byte, len(v.AlbumThumbnail))
		//albumdto.AlbumThumbnail = base64.StdEncoding.EncodeToString(b)

		albumdto.AlbumThumbnail = v.AlbumThumbnail
		albumdto.UpdatedAt = v.UpdatedAt
		albumdto.CreatedAt = v.CreatedAt
		albumdto.Description = v.Description
		albumdto.AlbumName = v.AlbumName
		albumdto.UserId = v.UserId
		albumdto.AlbumID = v.AlbumID
		for _, image := range v.Images{
			var imagedto dto.ImageDTO
			imagedto.ImageID = image.ImageID
			imagedto.AlbumId = image.AlbumId
			imagedto.ImageName = image.ImageName
			imagedto.CreatedAt = image.CreatedAt
			imagedto.UpdatedAt = image.UpdatedAt
			albumdto.Images = append(albumdto.Images, imagedto)
		}
		dtos = append(dtos, albumdto)
	}
	return dtos
}

func ToImageDTOs(r []models.Image) []dto.ImageDTO {
	dtos := make([]dto.ImageDTO, 0)

	for _, v := range r{
		var imageDto dto.ImageDTO
		//b := make([]byte, len(v.AlbumThumbnail))
		//albumdto.AlbumThumbnail = base64.StdEncoding.EncodeToString(b)

		imageDto.ImageID = v.ImageID
		imageDto.ImageName = v.ImageName
		imageDto.UpdatedAt = v.UpdatedAt
		imageDto.CreatedAt = v.CreatedAt
		imageDto.Imagefile = v.Imagefile
		imageDto.AlbumId = v.AlbumId

		dtos = append(dtos, imageDto)
	}
	return dtos
}


func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidateFields(email, password string) error {
	if IsEmpty(Trim(email)) || IsEmpty(password) {
		return ErrEmptyFields
	}
	return nil
}


func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	}
	return false
}

func Trim(attr string) string {
	return strings.TrimSpace(attr)
}