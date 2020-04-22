package forms

type ContentForm struct {
	UserId     int64  `json:"user_id" form:"user_id"`
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	Annotation string `json:"annotation" form:"annotation"`
	Image      string `json:"image" form:"image"`
}

// todo temporary
func (c *ContentForm) Validate() bool {
	return c.UserId > 0 && (c.Title != "" && len(c.Title) <= 200) && c.Content != ""
}
