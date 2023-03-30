package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"minyr/yr"
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

	var previousLine string

	for scanner.Scan() {
		line := scanner.Text()
		fahr, err := yr.CelsiusToFahrenheitLine(line)
		if err != nil {
			log.Fatal(err)
			continue
		}

		if previousLine != "" {
			fmt.Fprintln(writer, previousLine)
		}
		previousLine = fahr
	}

	if previousLine != "" {
		fmt.Fprintln(writer, "Data er basert paa gyldig data (per 18.03.2023)(CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Ole")
	}

	writer.Flush()
	fmt.Println("Conversions completed.")
}

