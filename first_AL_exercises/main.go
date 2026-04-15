package main

import (
	"bufio"
	"first_AL_exercises/vector"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(">> Linear Algebra calculator for vectors <<")

	fmt.Println("Enter your first vector:")
	v1, err := vector.CreateVectorFromInput(reader)
	if err != nil {
		fmt.Println("Error creating first vector:", err)
		return
	}
	fmt.Printf("\n First vector: %s\n", v1)

	for {
		printMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			v1 = scalarMultiplication(v1, reader)
		case "2":
			v1 = vectorAddition(v1, reader)
		case "3":
			v1 = vectorSubtraction(v1, reader)
		case "4":
			dotProduct(v1, reader)
		case "5":
			magnitude(v1)
		case "6":
			v1 = createNewVector(reader)
		case "7":
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option from the menu.")
		}
	}

}

func printMenu() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📊 MENU")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println(" 1. ✖️  Scalar Multiplication (v × scalar)")
	fmt.Println(" 2. ➕  Vector Addition (v1 + v2)")
	fmt.Println(" 3. ➖  Vector Subtraction (v1 - v2)")
	fmt.Println(" 4. ⚫  Dot Product (v1 · v2)")
	fmt.Println(" 5. 📏  Magnitude (|v|)")
	fmt.Println(" 6. 🆕  Create a new vector")
	fmt.Println(" 7. 🚪  Exit")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func scalarMultiplication(v vector.Vector, reader *bufio.Reader) vector.Vector {
	fmt.Printf("\n Current vector: %s\n", v)
	fmt.Println("\n Enter the scalar value:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	scalar, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return v
	}

	result := v.ScalarProduct(scalar)
	fmt.Printf("Result of scalar multiplication: %s\n", result)
	return result
}

func vectorAddition(v vector.Vector, reader *bufio.Reader) vector.Vector {
	fmt.Printf("\n Current vector: %s\n", v)
	fmt.Println("\n Enter the second vector for addition:")
	v2, err := vector.CreateVectorFromInput(reader)
	if err != nil {
		fmt.Println("Error creating second vector:", err)
		return v
	}

	result, err := v.AddVectors(v2)
	if err != nil {
		fmt.Println("Error occurred while adding vectors:", err)
		return v
	}
	fmt.Printf("Result of vector addition: %s\n", result)
	return result

}

func vectorSubtraction(v vector.Vector, reader *bufio.Reader) vector.Vector {

	fmt.Printf("\n Current vector: %s\n", v)
	fmt.Println("\n Enter the second vector for subtraction:")

	v2, err := vector.CreateVectorFromInput(reader)
	if err != nil {
		fmt.Println("Error creating second vector:", err)
		return v
	}

	result, err := v.SubtractVectors(v2)
	if err != nil {
		fmt.Println("Error occurred while subtracting vectors:", err)
		return v
	}
	fmt.Printf("Result of vector subtraction: %s\n", result)
	return result
}

func dotProduct(v vector.Vector, reader *bufio.Reader) {
	fmt.Printf("\n Current vector: %s\n", v)
	fmt.Println("\n Enter the second vector for dot product:")

	v2, err := vector.CreateVectorFromInput(reader)
	if err != nil {
		fmt.Println("Error creating second vector:", err)
		return
	}

	result, err := v.DotProduct(v2)
	if err != nil {
		fmt.Println("Error occurred while calculating dot product:", err)
		return
	}
	fmt.Printf("Result of dot product: %.4f\n", result)
}

func magnitude(v vector.Vector) {
	fmt.Printf("\n Current vector: %s\n", v)
	result := v.Magnitude()
	fmt.Printf("Magnitude of the vector: %.4f\n", result)
}

func createNewVector(reader *bufio.Reader) vector.Vector {
	fmt.Println("\n Creating a new vector:")
	v, err := vector.CreateVectorFromInput(reader)
	if err != nil {
		fmt.Println("Error creating new vector:", err)
		return nil
	}
	fmt.Printf("New vector created: %s\n", v)
	return v
}
