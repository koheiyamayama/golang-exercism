package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	boolean, _ := regexp.MatchString(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`, text)
	return boolean
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<\W*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	pattern := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0

	for _, v := range lines {
		if pattern.MatchString(v) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	pattern := regexp.MustCompile(`end-of-line\d*`)
	return pattern.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := regexp.MustCompile(`User\s+(\S*)`)

	for i, line := range lines {
		if pattern.MatchString(line) {
			userName := pattern.FindAllStringSubmatch(line, 1)[0][1]
			lines[i] = fmt.Sprintf("[USR] %s", userName) + " " + line
		}
	}

	return lines
}
