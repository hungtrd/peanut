package domain

type Book struct {
	BaseModel
	Title       string
	Description string
	Author      string
}

type CreateBookReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
}

type CreateBookResp struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
