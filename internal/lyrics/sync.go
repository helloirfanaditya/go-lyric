package lyrics

func CurrentLine(
	positionMs int64,
	lines []LyricLine,
) int {

	if len(lines) == 0 {
		return -1
	}

	current := 0

	for i, line := range lines {
		if positionMs >= line.TimeMs {
			current = i
		} else {
			break
		}
	}

	return current
}
