package repository

import (
	"albumwebapp/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

type AlbumRepositoryError struct {
	error
}




type AlbumRepository struct {
	DB *gorm.DB
}

func GetAlbumRespository(db *gorm.DB) AlbumRepository{
	return AlbumRepository {DB: db}
}

func (u *AlbumRepository) GetAll(userid interface{}) ([]models.Album, error){
		var albums []models.Album
	u.DB.Debug().Where("user_id = ?", userid).Find(&albums)

	//rows, err := u.DB.Debug().Joins("inner join albums on albums.user_id = users.id").Rows()
	//u.DB.Debug().Find(&albums, models.Album{UserId: int(userid)})
	//rows, err := u.DB.Table("albums").Select("album_id, album_name, created_at, description").Joins("left join users on users.user_id = albums.user_id").Rows()
	//defer rows.Close()
	//
	//if err !=nil{
	//	msg := fmt.Errorf("error getting data")
	//	return nil, &UserRepositoryError{msg}
	//}
	//for rows.Next(){
	//	var album models.Album
	//	rows.Scan(&album.AlbumID, &album.AlbumName, &album.CreatedAt, &album.Description, &album.UserId)
	//	albums = append(albums, album)
	//}
	if len(albums) > 0 {
			return albums, nil
	}

	err := fmt.Errorf("Server error")
	return nil, &UserRepositoryError{err}
}

//func (u *AlbumRepository) GetByID(id int) (models.User, error){
//	var user models.User
//	d := u.DB.Find(user, id)
//	if d.RowsAffected == 0{
//		return  user, &UserRepositoryError{fmt.Errorf("can not create user")}
//	}else{
//		return user, nil
//	}
//}

func (u *AlbumRepository) Save(album models.Album) (models.Album,error){
	d := u.DB.Save(&album)
	if d.Error !=nil{
		return  album, &AlbumRepositoryError{fmt.Errorf(d.Error.Error())}
	}else{
		return album, nil
	}
	return  album, &AlbumRepositoryError{fmt.Errorf("can not create album")}
}

//func (u GetAlbumRespository) Delete(user models.User){
//	u.DB.Delete(&user)
//}
