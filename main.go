package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
<<<<<<< HEAD
	"strconv"
	"strings"

	"github.com/onstrogithub/minyr/yr"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please choose 'convert', 'average', or 'exit'")

	for scanner.Scan() {
		input := scanner.Text()
=======
	"minyr/yr"
	"github.com/onstrogithub/funtemps/conv"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
>>>>>>> 508f6ebeafc5450e89a7385d08fc2bb41be52a67
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
<<<<<<< HEAD
			outputFile := "/minyr/kjevik-temp-fahr-20220318-20230318.csv"
			if _, err := os.Stat(outputFile); err == nil {
				fmt.Println("The file", outputFile, "already exists. Do you want to generate the file again? Enter 'j' for yes, or 'n' for no.")

				confirmScanner := bufio.NewScanner(os.Stdin)
				if confirmScanner.Scan() {
					confirmation := strings.ToLower(confirmScanner.Text())
					if confirmation == "j" {
						fmt.Println("Converting all measurements given in degrees Celsius to degrees Fahrenheit.")
						convertFahr()
					} else if confirmation == "n" {
						fmt.Println("Conversion canceled.")
					} else {
						fmt.Println("Invalid input. Please enter 'j' or 'n'.")
					}
				}
			} else {
				fmt.Println("Converting all measurements given in degrees Celsius to degrees Fahrenheit.")
				convertFahr()
			}
		} else if input == "average" {
			fmt.Println("Please choose the temperature unit: 'C' for Celsius or 'F' for Fahrenheit")
			unitScanner := bufio.NewScanner(os.Stdin)
			if unitScanner.Scan() {
				unitInput := strings.ToUpper(unitScanner.Text())
				if unitInput == "C" {
					fmt.Println("Calculating the average temperature in degrees Celsius.")
					avg, err := calculateAverage("/minyr/kjevik-temp-celsius-20220318-20230318.csv")
					if err != nil {
						fmt.Printf("Error calculating average temperature: %v\n", err)
					} else {
						fmt.Printf("Average temperature in degrees Celsius: %.2f\n", avg)
					}
				} else if unitInput == "F" {
					fmt.Println("Calculating the average temperature in degrees Fahrenheit.")
					avg, err := calculateAverage("/minyr/kjevik-temp-fahr-20220318-20230318.csv")
					if err != nil {
						fmt.Printf("Error calculating average temperature: %v\n", err)
					} else {
						fmt.Printf("Average temperature in degrees Fahrenheit: %.2f\n", avg)
					}
				} else {
					fmt.Println("Invalid input. Please enter 'C' or 'F'")
				}
			}
		}
	} // Add this closing brace
}




func convertFahr() {
    src, err := os.Open("/minyr/kjevik-temp-celsius-20220318-20230318.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer src.Close()
    dst, err := os.Create("/minyr/kjevik-temp-fahr-20220318-20230318.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer dst.Close()

    scanner := bufio.NewScanner(src)
    writer := bufio.NewWriter(dst)
    isFirstLine := true

    for scanner.Scan() {
        line := scanner.Text()
        if isFirstLine {
            fmt.Fprintln(writer, "Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur")
            isFirstLine = false
            continue
        }

        if strings.HasPrefix(line, "Data er gyldig per") || strings.TrimSpace(line) == "" {
            continue
        }

        fahr, err := yr.CelsiusToFahrenheitLine(line)
        if err != nil {
            fmt.Printf("Error processing line: %s\nError: %v\n", line, err)
        } else {
            fmt.Fprintln(writer, fahr)
        }
    }

    fmt.Fprintln(writer, "Data er basert på gyldig data (per 18.03.2023)(CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Ole")

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    writer.Flush()
    fmt.Println("Conversions completed.")
}

















func calculateAverage(filename string) (float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isFirstLine := true
	sum := 0.0
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if isFirstLine {
			isFirstLine = false
			continue
		}

		dividedString := strings.Split(line, ";")
		if len(dividedString) != 4 {
			continue
		}

		tempStr := strings.TrimSpace(dividedString[3])
		if tempStr == "" {
			continue
		}

		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			log.Printf("Error parsing temperature: %v\n", err)
			continue
		}

		sum += temp
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	if count > 0 {
		average := sum / float64(count)
		return average, nil
	}
	return 0, fmt.Errorf("temperature column not found")
=======
			fmt.Println("Konverterer alle maalingene gitt i grader Celsius til grader Fahrenheit.")
			convertFahr()
		} else if input == "average" {
		 averageTemp()
} else {
			fmt.Println("Vennligst velg convert, average eller exit")
		}
	}
}

func convertFahr() {
	src, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dst)

		if scanner.Scan() {
		line1 := (scanner.Text())
		fmt.Fprintln(writer, line1)
		}

	for scanner.Scan() {
		line := scanner.Text()
		fahr, err := yr.CelsiusToFahrenheitLine(line)
		if err != nil {
			log.Fatal(err)
			continue
		}
	fmt.Fprintln(writer, fahr)
}
	writer.Flush()
	fmt.Println("Convertion complete")
	
}
func averageTemp() {

fmt.Println("VEnligst velg c/f")

var input string
scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
input = scanner.Text()

sum := 0
count := 0

if input == "c" {
fmt.Println("Finding the average temp in Celsius")
avg := yr.AverageTemp(sum, float64(count))
fmt.Printf("Average: %.2f\n", avg)
} else if input == "f" {
fmt.Println("Finding the average temp in Fahrenheit")
avg := yr.AverageTemp(sum, float64(count))
avgFahr := conv.CelsiusToFahrenheit(avg)
fmt.Printf("Average: %.2f\n", avgFahr)
} else  {
fmt.Println("Please select (c/f)):")
}
}
>>>>>>> 508f6ebeafc5450e89a7385d08fc2bb41be52a67
}

