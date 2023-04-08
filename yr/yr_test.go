package yr

import (
    "testing"
)

func TestCelsiusToFahrenheitString(t *testing.T) {
    type test struct {
        input string
        want  string
    }
    tests := []test{
        {input: "6", want: "42.8"},
        {input: "0", want: "32.0"},
    }

    for _, tc := range tests {
        got, err := CelsiusToFahrenheitString(tc.input)
        if err != nil {
            t.Errorf("Unexpected error: %v", err)
        }
        if tc.want != got {
            t.Errorf("Expected %s, got %s", tc.want, got)
        }
    }
}

func TestCelsiusToFahrenheitLine(t *testing.T) {
    type test struct {
        input string
        want  string
    }
    tests := []test{
        {input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
        {input: "Kjevik;SN39040;18.03.2022 01:50", want: ""},
        {input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
        {input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
    }
    for _, tc := range tests {
        got, err := CelsiusToFahrenheitLine(tc.input)
        if err != nil {
            t.Errorf("Unexpected error: %v", err)
        }
        if tc.want != got {
            t.Errorf("Expected %s, got %s", tc.want, got)
        }
    }
}

