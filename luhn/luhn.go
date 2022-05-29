package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	ids := strings.Split(id, "")
	sum := 0

	for i := 1; i <= len(ids); i++ {
		n, e := strconv.Atoi(ids[len(ids)-i])

		if e != nil {
			return false
		}

		if i%2 == 0 {
			n *= 2

			if 9 < n {
				sum += n - 9
			} else {
				sum += n
			}
		} else {
			sum += n
		}
	}

	return sum%10 == 0
}
