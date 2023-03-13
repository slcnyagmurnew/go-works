package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

/*
Did not work
//type Locations struct {
//	Locations []Location `json:"locations"`
//}
*/

type Location struct {
	Name string `json:"name"`
	Code int64  `json:"code"`
}

func Printing(name string) string {
	message := fmt.Sprintln(name)
	return message
}

func PrintStruct(name Location) string {
	message := fmt.Sprintln(name)
	return message
}

func InitLocations(c *gin.Context) {
	var locations = []Location{{Name: "Turkey", Code: 90}, {Name: "France", Code: 33}, {Name: "Spain", Code: 34}}
	SaveDataToJson(locations, "data/locations.json")
	c.IndentedJSON(http.StatusOK, locations)
}

func GetLocations(c *gin.Context) {
	locations := GetDataFromFile("data/locations.json")
	c.IndentedJSON(http.StatusOK, locations)
}

func ReadJsonFileContent(filePath string) []byte {
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Printf("Bytes : %v \n", byteValue)
	return byteValue
}

func GetDataFromFile(filename string) []Location {
	var data []Location
	fileData := ReadJsonFileContent(filename)
	json.Unmarshal(fileData, &data)
	return data
}

func AddLocation(c *gin.Context) {
	var updatedLocations []Location
	existingData := GetDataFromFile("data/locations.json")
	fmt.Printf("Locations : %v \n", existingData)
	newLocationData := GetDataFromFile("data/newLocations.json")
	fmt.Printf("New locations : %v Len: %d \n", newLocationData, len(newLocationData))

	for i := 0; i < len(newLocationData); i++ {
		PrintStruct(newLocationData[i])
		updatedLocations = append(existingData, newLocationData[i])
	}
	fmt.Println(updatedLocations)
	SaveDataToJson(newLocationData, "data/locations.json")
	c.IndentedJSON(http.StatusCreated, updatedLocations)
}

func SaveDataToJson(data []Location, filename string) {
	if _, err := os.Stat(filename); err == nil {
		file, _ := json.MarshalIndent(data, "", "    ")
		_ = ioutil.WriteFile(filename, file, 0644) // permission rwx
	} else {
		Printing("Error occurred")
	}
}
