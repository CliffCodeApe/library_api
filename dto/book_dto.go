package dto

type BookRequest struct {
	Title       string `form:"title" validate:"required"`
	Author      string `form:"author" validate:"required"`
	Year        int    `form:"year" validate:"required"`
	Genre       string `form:"genre" validate:"required"`
	Stock       int    `form:"stock" validate:"required"`
	Description string `form:"description" validate:"required"`
	Publisher   string `form:"publisher" validate:"required"`
	ISBN        string `form:"isbn" validate:"required"`
	Language    string `form:"language" validate:"required"`
	Pages       int    `form:"pages" validate:"required"`
}

type BookListResponse struct {
	StatusCode int            `json:"status_code"`
	Message    string         `json:"message"`
	Data       []BookListData `json:"data"`
}

type BookListData struct {
	ID     uint64 `json:"id"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
}

type BookDetailResponse struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       BookData `json:"data"`
}

type BookData struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Publisher   string `json:"publisher"`
	ISBN        string `json:"isbn"`
	Language    string `json:"language"`
	Pages       int    `json:"pages"`
	Thumbnail   string `json:"thumbnail"`
	FilePath    string `json:"file_path"`
}
