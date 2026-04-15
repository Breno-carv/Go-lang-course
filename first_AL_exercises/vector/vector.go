package vector

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vector []float64

func createVector(values ...float64) Vector {
	return Vector(values)
}

// capitalize first letter to make it public (that's full bs tho)
func CreateVectorFromInput(reader *bufio.Reader) (Vector, error) {

	fmt.Println("Enter the values for the vector, one by one:")
	fmt.Println("Type OK when you're done, or type EXIT to quit.")

	var values []float64
	index := 1

	for {
		fmt.Printf("Value %d: ", index)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch strings.ToLower(input) {
		case "ok":
			if len(values) == 0 {
				fmt.Println("You must enter at least one value for the vector.")
				continue
			}
			fmt.Printf("Vector created with %d values.\n", len(values))
			return createVector(values...), nil
		case "exit":
			os.Exit(0)

		default:
			value, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number, or type OK to finish.")
				continue
			}

			values = append(values, value)
			index++
		}

	}

}

// Below this there's a new concept, a METHOD: basically a func that has it's receiver declared before it's name
// in Methods it's ok to start capitalizing
func (v Vector) ScalarProduct(scalar float64) Vector {
	result := make(Vector, len(v))
	for i, value := range v {
		result[i] = value * scalar
	}
	return result
}

func (v Vector) AddVectors(v2 Vector) (Vector, error) {
	if len(v) != len(v2) {
		return nil, fmt.Errorf("Vectors must have the same dimension: %d != %d", len(v), len(v2))
	}

	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] + v2[i]
	}
	return result, nil
}

func (v Vector) SubtractVectors(v2 Vector) (Vector, error) {
	if len(v) != len(v2) {
		return nil, fmt.Errorf("Vectors must have the same dimension: %d and %d", len(v), len(v2))
	}

	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] - v2[i]
	}
	return result, nil
}

func (v Vector) DotProduct(v2 Vector) (float64, error) {
	if len(v) != len(v2) {
		return 0, fmt.Errorf("Vectors must have the same dimension: %d and %d", len(v), len(v2))
	}

	var result float64
	for i := range v {
		result += v[i] * v2[i]
	}
	return result, nil
}

func (v Vector) Magnitude() float64 {
	var sumSquares float64 = 0.0
	for _, value := range v {
		sumSquares += value * value
	}
	return math.Sqrt(sumSquares)
}

func (v Vector) String() string {
	stringValues := make([]string, len(v))
	for i, value := range v {
		stringValues[i] = fmt.Sprintf("%.2f", value)
	}
	return fmt.Sprintf("Vector(%s)", strings.Join(stringValues, ", "))
}
