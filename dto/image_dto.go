package dto

const (
	MESSAGE_FAILED_GET_IMAGE_ID   = "missing image id"
	MESSAGE_FAILED_OPEN_IMAGE     = "failed to open image"
	MESSAGE_FAILED_GET_IMAGE_PATH = "failed get image path"
	MESSAGE_FAILED_SEND_IMAGE     = "failed send image"
)

type (
	FilmImageResponse struct {
		ID     uint   `json:"id"`
		FilmID uint   `json:"film_id,omitempty"`
		Path   string `json:"path"`
		Status bool   `json:"status"`
	}
)
