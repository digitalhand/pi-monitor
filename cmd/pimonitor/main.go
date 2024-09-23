package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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


func captureVideo(filename string)(bool, error) {

	var timestampValue = time.Now().Format(time.RFC850)

	out, err := exec.Command(
	"rpicam-vid -b 9000000 -t 20000 --width 1920 --height 1080 --codec libav --libav-audio -o ", 
	fmt.Sprintf(filename), 
	"_", 
	fmt.Sprintf(timestampValue),
	).Output()
	if err != nil {
		fmt.Printf("%s", err)
		return false, err
  }

	output := string(out[:])
	fmt.Println("Video Request Sent. ", output)
	time.Sleep(40 * time.Second)
	return true, err

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

		video, err := captureVideo("videocapture")
			if err != nil {
			log.Fatal(err)
		}
		if video { fmt.Println("Video Captured!")}
		
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