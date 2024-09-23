package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func VerifyPathExists(path string) (string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		return e.Name(), nil
	}
	return "error", nil
}

func main() {
	path := "/sys/bus/w1/devices/"
	ds18b20, err := VerifyPathExists(path)
	if err != nil {
		log.Fatal(err)
	}

	for true {
		_, faren, err := GetSensorTemperature(path + ds18b20 + "/w1_slave")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Temp: ", faren)
		//time.Sleep(3 * time.Second)
	}

}

func GetSensorTemperature(fileName string) (float64, float64, error) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	breakLine := "\n"
	var farenheit float64 = 0.0
	var celcius float64 = 0.0

	if len(fileData) != 0 {
		secondLine := strings.Split(string(fileData[:]), breakLine)[1]
		temperatureData := strings.Split(secondLine, " ")[9]
		temperature, _ := strconv.ParseFloat(temperatureData[2:], 64)
		celcius := (temperature / 1000)
		farenheit := (celcius * 1.8) + 32
		return celcius, farenheit, err

	} else {
		fmt.Println("No Temp To Report, taking a small break...")
		time.Sleep(5 * time.Second)
		return celcius, farenheit, err
	}

}