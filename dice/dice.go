package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// Roll handles dice rolls for a given input, e.g., "2d6".
// It also supports the --percentage and --average flags.
func Roll(input string, percentage bool, average bool) (string, error) {

	// Parse input (e.g., "2d6")
	parts := strings.Split(input, "d")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid dice format, use NdX (e.g., 2d6, 1d20)")
	}

	times, err1 := strconv.Atoi(parts[0])
	sides, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || times < 1 || sides < 1 {
		return "", fmt.Errorf("both N and X must be positive integers")
	}

	// Special case for 2d10 with --percentage
	if percentage && times == 2 && sides == 10 {
		d1 := rand.Intn(sides) + 1
		d2 := rand.Intn(sides) + 1
		percentageResult := d1*10 + d2
		if percentageResult == 0 {
			percentageResult = 100 // 00 on 2d10 is treated as 100
		}
		return fmt.Sprintf("Rolled 2d10 (percentage): %d", percentageResult), nil
	}

	// Normal dice roll
	results := make([]int, times)
	total := 0
	for i := 0; i < times; i++ {
		results[i] = rand.Intn(sides) + 1
		total += results[i]
	}

	// Calculate average if requested
	averageOutput := ""
	if average {
		avg := float64(times) * (float64(sides) + 1) / 2
		averageOutput = fmt.Sprintf(" (Average: %.2f)", avg)
	}

	return fmt.Sprintf("Rolled %s: %v (Total: %d)%s", input, results, total, averageOutput), nil
}
