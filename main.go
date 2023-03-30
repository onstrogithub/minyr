package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"minyr/yr"
	"github.com/onstrogithub/funtemps/conv"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
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
}

