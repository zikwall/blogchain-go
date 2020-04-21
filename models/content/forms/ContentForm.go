package forms

type ContentForm struct {
	UserId     int64  `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Annotation string `json:"annotation"`
	Image      string `json:"image"`
}

// todo temporary
func (c *ContentForm) Validate() bool {
	return c.UserId > 0 && (c.Title != "" && len(c.Title) <= 200) && c.Content != ""
}
