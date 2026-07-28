package utils

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateTgUsername(prefix string) string {
	id := gonanoid.MustGenerate("ABCDEFGHJKLMNPQRSTUVWXYZ23456789", 5)
	return fmt.Sprintf("%s_%s_bot", prefix, id)
}
