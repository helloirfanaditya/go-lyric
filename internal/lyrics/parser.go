package lyrics

import (
	"regexp"
	"strconv"
	"strings"
)

type LyricLine struct {
	TimeMs int64
	Text   string
}

var re = regexp.MustCompile(`\[(\d+):(\d+)\.(\d+)\](.*)`)

func ParseLRC(content string) []LyricLine {
	lines := strings.Split(content, "\n")

	var result []LyricLine

	for _, line := range lines {
		match := re.FindStringSubmatch(line)

		if len(match) != 5 {
			continue
		}

		minute, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		centis, _ := strconv.Atoi(match[3])

		totalMs := int64(minute*60000 + second*1000 + centis*10)

		result = append(result, LyricLine{
			TimeMs: totalMs,
			Text:   strings.TrimSpace(match[4]),
		})
	}

	return result
}
