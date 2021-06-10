package repositories

type Tag struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

type ContentTag struct {
	Id        int64 `json:"id"`
	ContentId int64 `json:"content_id"`
	TagId     int64 `json:"tag_id"`
}
