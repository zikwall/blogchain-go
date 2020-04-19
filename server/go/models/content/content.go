package content

import "github.com/zikwall/blogchain/models/user"

type Content struct {
	Id      int64
	UserId  int
	Title   string
	Content string

	User user.User
}
