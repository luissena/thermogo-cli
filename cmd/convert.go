package cmd

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert temperatures between Celsius, Fahrenheit, and Kelvin",
	RunE: func(cmd *cobra.Command, args []string) error {
		units := []string{"Celsius", "Fahrenheit", "Kelvin"}

		fromPrompt := promptui.Select{
			Label: "Select input unit",
			Items: units,
		}
		fromIndex, _, err := fromPrompt.Run()
		if err != nil {
			return fmt.Errorf("failed to select input unit: %v", err)
		}

		toPrompt := promptui.Select{
			Label: "Select output unit",
			Items: units,
		}
		toIndex, _, err := toPrompt.Run()
		if err != nil {
			return fmt.Errorf("failed to select output unit: %v", err)
		}

		valuePrompt := promptui.Prompt{
			Label: "Enter temperature value",
			Validate: func(input string) error {
				_, err := strconv.ParseFloat(input, 64)
				if err != nil {
					return err
				}
				return nil
			},
		}
		inputStr, err := valuePrompt.Run()
		if err != nil {
			return fmt.Errorf("failed to read temperature value: %v", err)
		}

		inputValue, _ := strconv.ParseFloat(inputStr, 64)

		fromUnit := units[fromIndex]
		toUnit := units[toIndex]
		result := convertTemperature(inputValue, fromUnit, toUnit)

		fmt.Printf("\nResult: %.2f %s → %.2f %s\n", inputValue, fromUnit, result, toUnit)
		return nil
	},
}

func convertTemperature(value float64, fromUnit, toUnit string) float64 {
	if fromUnit == toUnit {
		return value
	}

	// Convert input to Celsius
	var celsius float64
	switch fromUnit {
	case "Celsius":
		celsius = value
	case "Fahrenheit":
		celsius = (value - 32) * 5 / 9
	case "Kelvin":
		celsius = value - 273.15
	}

	// Convert from Celsius to target unit
	switch toUnit {
	case "Celsius":
		return celsius
	case "Fahrenheit":
		return (celsius * 9 / 5) + 32
	case "Kelvin":
		return celsius + 273.15
	default:
		return value
	}
}
