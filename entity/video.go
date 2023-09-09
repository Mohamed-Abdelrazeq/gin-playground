package entity

type Video struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,max=1000"`
	URL         string `json:"url" binding:"required"`
}
