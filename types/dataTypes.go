package types

import "fmt"

// Make provides allocation in memory (like malloc, calloc)
// CreateSlices /*

func CreateSlices() {
	// create empty slice with make
	emptySlice := make([]string, 4)
	fmt.Println(emptySlice)
	// initialize slice with definition of values
	initializedSlice := []int{6, 0, 1}
	fmt.Println(initializedSlice)
	// create 2D slices
	twoDSlice := make([][]string, 4) // # of rows: 4
	const innerSliceLen = 3          // # of columns: 3
	for i := 0; i < 4; i++ {
		// create columns over rows
		twoDSlice[i] = make([]string, innerSliceLen)
		for j := 0; j < innerSliceLen; j++ {
			// fill every column for one row
			twoDSlice[i][j] = "kk"
		}
	}
	fmt.Println(twoDSlice)
}

func CopyData() {
	// copy slice with builtin 'copy' function
	copiedData := make([]string, 5)
	realData := []string{"y", "a", "g"} // different lengths not causes an error
	copy(copiedData, realData)
	fmt.Println(copiedData)
}

func GetSliceData() {
	exampleData := []float64{7, 3.8, 8, 5.2, 6.32, 0.74}
	fmt.Println(exampleData[3])
	fmt.Println(exampleData[2:5])
}

func MapOperations() {
	// map[key type] value type
	emptyMap := make(map[string]int)
	// add key-value pair to map
	emptyMap["yagmur"] = 24
	fmt.Println(emptyMap)
	// initialize map with definition (very similar to Python)
	initializedMap := map[string]int{"second": 4, "third": 59}
	fmt.Println(initializedMap)
	// get value from key
	value := initializedMap["third"]
	fmt.Println(value)
	// get length
	fmt.Println(len(initializedMap))
	// delete value from map
	delete(initializedMap, "third")
	// check if key exists in map
	_, checkout := initializedMap["third"]
	fmt.Println(checkout)
}

func IterateWithRanges() {
	// range for slice
	iterationSlice := []int{5, 2, 6, 8, 1}
	var sum = 0
	// find average of slice
	for _, num := range iterationSlice {
		sum += num
	}
	fmt.Printf("Average: %.2f ", float64(sum/len(iterationSlice)))
	// range for map
	iterationMap := map[string]string{"a": "b", "c": "d", "e": "f"}
	for k, v := range iterationMap {
		fmt.Printf("Key: %s -> Value: %s \n", k, v)
	}
	// range for string (ASCII values)
	for index, character := range "free" {
		fmt.Printf("Index: %d for char %d \n", index, character)
	}
}

// CreateArray fixed length slice
func CreateArray() {
	// empty array
	var values [4]int
	fmt.Println(values)
	// change value with index
	values[3] = 10
	fmt.Println(values)
	// create array with initialized values
	newArray := [5]float64{3.4, 6.3, 0.7}
	fmt.Println(newArray)
	// 2D array has same logic with slices' 2D
}