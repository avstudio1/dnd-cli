package dice

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// AllowedDice defines the standard dice types
var AllowedDice = map[int]bool{
	4:  true,
	6:  true,
	8:  true,
	10: true,
	12: true,
	20: true,
}

// Roll handles dice rolls for a given input, e.g., "2d6".
// It also supports the --percentage and --average flags.
func Roll(input string, percentage bool, average bool) (string, error) {
	// Initialize a new random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Parse input (e.g., "2d6")
	parts := strings.Split(input, "d")
	if len(parts) != 2 {
		return "", errors.New("invalid dice format, use NdX (e.g., 2d6, 1d20)")
	}

	if len(parts[0]) == 0 || len(parts[1]) == 0 {
		return "", errors.New("invalid dice format, use NdX (e.g., 2d6, 1d20)")
	}

	times, err1 := strconv.Atoi(parts[0])
	sides, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || times < 1 || sides < 1 {
		return "", errors.New("both N and X must be positive integers")
	}

	// Validate dice type
	if !AllowedDice[sides] {
		return "", fmt.Errorf("unsupported dice type: d%d. Allowed types are: d4, d6, d8, d10, d12, d20", sides)
	}

	// Special case for 2d10 with --percentage
	if percentage && times == 2 && sides == 10 {
		d1 := r.Intn(sides)
		d2 := r.Intn(sides)
		percentageResult := d1*10 + d2
		if percentageResult == 0 {
			percentageResult = 100 // 00 on 2d10 is treated as 100
		}
		return fmt.Sprintf("Rolled 2d10 [%v %v] (percentage): %d", d1, d2, percentageResult), nil
	}

	// Normal dice roll
	results := make([]int, times)
	total := 0
	for i := 0; i < times; i++ {
		results[i] = r.Intn(sides) + 1
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
