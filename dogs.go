package dogs

import (
	"strings"

	"github.com/ritvikrohan/puppy"
)

func WhenGrowsUp(s string) string {
	return "When grows up will:" + strings.ToUpper(puppy.Bark())
}
