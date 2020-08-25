package repository

import (
	"albumwebapp/models"
	"fmt"
	"github.com/jinzhu/gorm"
)


type ImageRepositoryError struct {
	error
}




type ImageRepository struct {
	DB *gorm.DB
}

func GetImageRespository(db *gorm.DB) ImageRepository{
	return ImageRepository {DB: db}
}

func (u *ImageRepository) Save(album models.Image) (models.Image,error){
	d := u.DB.Save(&album)
	if d.Error !=nil{
		return  album, &ImageRepositoryError{fmt.Errorf(d.Error.Error())}
	}else{
		return album, nil
	}
	return  album, &ImageRepositoryError{fmt.Errorf("can not create album")}
}

func (u *ImageRepository) GetAll(albumid interface{}) ([]models.Image, error){
	var images []models.Image
	u.DB.Debug().Where("album_id = ?", albumid).Find(&images)
	if len(images) > 0 {
		return images, nil
	}

	err := fmt.Errorf("Server error")
	return nil, &ImageRepositoryError{err}
}

func (u *ImageRepository) GetOne(imageid interface{}) (models.Image, error){
	var image models.Image
	db:= u.DB.Debug().Where("image_id = ?", imageid).Find(&image)
	if image.ImageID > 0 {
		return image, &ImageRepositoryError{db.Error}
	}

	err := fmt.Errorf("Server error")
	return image, &ImageRepositoryError{err}
}

func (u *ImageRepository) Delete(imageid interface{}) (bool, error){
	db := u.DB.Debug().Where("image_id = ?", imageid).Delete(&models.Image{})

	if db.Error!=nil{
		err := fmt.Errorf(db.Error.Error())
		return false, &ImageRepositoryError{err}
	}
	return true, nil
}
