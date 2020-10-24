package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// DriverName is a name of compatible Linux kernel module
const DriverName = "i915"

func main() {
	if !checkModuleIsLoaded() {
		log.Fatalf("Driver %s was not found in loaded ones.\nExiting...\n", DriverName)
	}

	if len(os.Args) < 2 {
		log.Printf("Actual pwm frequency is: %d\n", getFrequency())
		log.Fatalln("Choose operating mode!")
	}
	var pwmChange string
	var blChange string

	flag.StringVar(&pwmChange, "pwm", "+0", "Specify PWM frequency in Hz") // A value is +100 like
	flag.StringVar(&blChange, "bl", "+0", "Specify Backlight level in %")  // A value is +100 like
	flag.Parse()

	switch string(pwmChange[0]) {
	case "+", "-":
		if change, err := strconv.ParseInt(pwmChange, 10, 64); err == nil {
			changeFrequency(int(change))
		} else {
			log.Fatal(err)
		}
	case "=":
		if set, err := strconv.ParseInt(pwmChange[1:], 10, 64); err == nil {
			setFrequency(int(set))
		} else {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Value should be like +<freq>/-<freq>/=<freq>")
	}

	switch string(blChange[0]) {
	case "+", "-":
		if change, err := strconv.ParseInt(blChange, 10, 64); err == nil {
			changeBacklightPercent(int(change))
		} else {
			log.Fatal(err)
		}
	case "=":
		if set, err := strconv.ParseInt(blChange[1:], 10, 64); err == nil {
			setBacklightPercent(int(set))
		} else {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Value should be like +<freq>/-<freq>/=<freq>")
	}

}

// All loaded modules are listed in procfs
func checkModuleIsLoaded() bool {
	content, err := ioutil.ReadFile("/proc/modules")
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Contains(string(content), DriverName)
}
