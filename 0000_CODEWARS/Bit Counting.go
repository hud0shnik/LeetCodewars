package kata

import (
	"fmt"
	"strings"
)

func CountBits(n uint) int {
	return strings.Count(fmt.Sprintf("%b", n), "1")
}
