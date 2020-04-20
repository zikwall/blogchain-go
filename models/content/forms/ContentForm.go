package forms

type ContentForm struct {
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// todo temporary
func (c *ContentForm) Validate() bool {
	return c.UserId > 0 && (c.Title != "" && len(c.Title) <= 200) && c.Content != ""
}
