package utils

import (
	"fmt"
)

func CreateImagePath(uuidv4 string) string {
	return fmt.Sprintf("%s.png", uuidv4)
}
