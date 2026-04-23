package main

import "fmt"

// variadic fn's are just functions that can take an arbitrary number of arguments

// Assignment:
// 	Complete the getCostsByDay function. It should return a slice of float64, where each element is the total Cost of the day.
/* The length of the slice should be equal as the number of days represented in the costs slice,
inlcuding days that have no cost, up to the last day in the costs.
Days are numbered and start at 0

e.g: []costs{
{0, 4.0}
{1, 2.1}
... day / cost
 format
}
*/

type TCosts struct {
	day   int
	value float64
}

func getCostsByDay(costs []TCosts) []float64 {
	var lastDay = 0
	for _, c := range costs {
		if c.day > lastDay {
			lastDay = c.day
		}

	}

	values := make([]float64, lastDay+1)

	for _, c := range costs {
		values[c.day] += c.value
	}

	return values

}

// another easy wasy to solve this was simply by checking wether the cost.day >= len(costs), if so, just append 0.0 on costsByDay
// then just add at costsByDay[costs.day] += costs.value

func main() {
	costsSlice1 := []TCosts{
		{0, 1.5}, {1, 2.3}, {2, 0}, {3, 3.1}, {4, 2.8},
		{5, 0}, {6, 1.9}, {7, 2.2}, {8, 2.7}, {9, 0},
		{10, 3.4}, {11, 2.6}, {12, 0}, {13, 3.0}, {14, 2.1},
		{15, 1.8}, {16, 0}, {17, 2.9}, {18, 3.3}, {19, 2.5},
		{20, 0}, {21, 2.0}, {22, 1.7},
	}

	costsSlice2 := []TCosts{
		{0, 5.0}, {1, 4.8}, {2, 4.5}, {3, 0}, {4, 4.2},
		{5, 3.9}, {6, 0}, {7, 3.7}, {8, 3.5}, {9, 3.3},
		{10, 0}, {11, 3.0}, {12, 2.8}, {13, 2.6}, {14, 0},
		{15, 2.4}, {16, 2.2}, {17, 2.0}, {18, 0}, {19, 1.8},
		{20, 1.6}, {21, 1.4}, {22, 0}, {23, 1.2}, {24, 1.0},
		{25, 0}, {26, 0.8}, {27, 0.6},
	}

	costsSlice3 := []TCosts{
		{0, 0}, {1, 1.1}, {2, 1.3}, {3, 0}, {4, 1.5},
		{5, 1.7}, {6, 1.9}, {7, 0}, {8, 2.1}, {9, 2.3},
		{10, 2.5}, {11, 0}, {12, 2.7}, {13, 2.9}, {14, 3.1},
		{15, 0}, {16, 3.3}, {17, 3.5}, {18, 3.7}, {19, 0},
		{20, 3.9}, {21, 4.1}, {22, 0}, {23, 4.3}, {24, 4.5},
	}

	costsSlice4 := []TCosts{
		{0, 12.0}, {0, 1.9}, // repeated day 0 (total should be 13.9)
		{3, 2.5},         // gap: days 1,2 missing
		{4, 0}, {4, 3.2}, // repeated day 4 (includes zero)
		{6, 1.1},                   // gap: day 5 missing
		{7, 2.0}, {7, 2.3}, {7, 0}, // repeated day 7
		{10, 5.5},          // gap: days 8,9 missing
		{12, 0}, {12, 4.4}, // repeated day 12
		{15, 3.3},            // gap: 13,14 missing
		{18, 6.0}, {18, 1.0}, // repeated day 18
		{21, 0},              // standalone zero
		{25, 2.2}, {25, 2.8}, // repeated day 25
		{29, 4.0}, // gap before
	}

	fmt.Println(getCostsByDay(costsSlice1))
	fmt.Println(getCostsByDay(costsSlice2))
	fmt.Println(getCostsByDay(costsSlice3))
	fmt.Println(getCostsByDay(costsSlice4))
}
