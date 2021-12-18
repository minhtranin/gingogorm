package entity

type Video struct {
	Title       string `json:"title"`
	Description string `json:"description" binding:"required,email"`
	Url         string `json:"url"`
}
