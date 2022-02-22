package repositories

type Tag struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

type ContentTag struct {
	ID        int64 `json:"id"`
	ContentID int64 `json:"content_id"`
	TagID     int64 `json:"tag_id"`
}
