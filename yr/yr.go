package yr

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
"os"
"log"
"bufio"
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

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon 
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperaturaaling i grader celsius
func CelsiusToFahrenheitLine(line string) (string, error) {

        dividedString := strings.Split(line, ";")
	var err error

	if (len(dividedString) == 4) {
		dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil

	//return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
		}
		func CountLines(input string) int {
		const filename = input
		file, err := os.Open(filename)
                if err != 0 {
		log.Fatalln(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lines := 0
		for scanner.Scan() {
		lines++
		}
		if err := scanner.Err(); err != nil {
		panic(err)
		}
		return lines
		}

