package forms

import (
	"errors"
)

type ContentForm struct {
	UserId     int64  `json:"user_id" form:"user_id"`
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	Annotation string `json:"annotation" form:"annotation"`
	Tags       string `json:"tags" form:"tags"`
	ImageName  string
	UUID       string
}

// todo temporary
func (c *ContentForm) Validate() error {
	if c.UserId > 0 && (c.Title != "" && len(c.Title) <= 200) && c.Content != "" {
		return nil
	}

	return errors.New("Invalid data entered")
}
