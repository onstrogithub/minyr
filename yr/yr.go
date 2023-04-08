package yr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/onstrogithub/funtemps/conv"
)

func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

func CelsiusToFahrenheitLine(line string) (string, error) {
	dividedString := strings.Split(line, ";")
	var err error

	if len(dividedString) != 4 {
		return "", nil // Return an empty string without an error
	}

	dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
	if err != nil {
		return "", err
	}

	return strings.Join(dividedString, ";"), nil
}

