package forms

import "mime/multipart"

type ContentForm struct {
	UserId     int64     `json:"user_id" form:"user_id"`
	Title      string    `json:"title" form:"title"`
	Content    string    `json:"content" form:"content"`
	Annotation string    `json:"annotation" form:"annotation"`
	image      FormImage `json:"image" form:"image"`
}

type FormImage struct {
	File *multipart.FileHeader
	Err  error
}

func (c *ContentForm) GetImage() FormImage {
	return c.image
}

func (c *ContentForm) SetImage(image FormImage) {
	c.image = image
}

// todo temporary
func (c *ContentForm) Validate() bool {
	return c.UserId > 0 && (c.Title != "" && len(c.Title) <= 200) && c.Content != ""
}
