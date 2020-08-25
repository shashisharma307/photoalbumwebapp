package dto


type AlbumRequest struct {
	AlbumName string `json:"album_name"`
	UserId int `json:"user_id"`
	Description string `json:"description"`
}
