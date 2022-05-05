// package main

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	countdown := "Launch in T minus " + "10 seconds"
// 	countdown = "Launch in T minus " + strconv.Itoa(10) + " seconds"
// 	age := 41
// 	margAge := float64(age)
// 	fmt.Println(countdown)
// 	fmt.Println(margAge)
// 	var bh float64 = 32768
// 	fmt.Println(int16(bh))
// 	fmt.Println(math.MaxInt16)
// 	v := 42
// 	if v >= 0 && v <= math.MaxUint8 {
// 		v8 := uint8(v)
// 		fmt.Println("converted:", v8)
// 	}

// 	fmt.Println(string(2340), string(2340), string(2340))
// 	str := fmt.Sprintf("%v", 12)
// 	newVal, err := strconv.Atoi("12")
// 	if err != nil {

// 	}
// 	fmt.Println(newVal)
// 	fmt.Println(str)

// 	launch := false
// 	launchText := fmt.Sprintf("%v", launch)
// 	fmt.Print(launchText)
// 	fmt.Print("a", "x")
// 	var yesNo string
// 	if launch {
// 		yesNo = "yes"
// 	} else {
// 		yesNo = "no"
// 	}
// 	fmt.Printf("Ready for launch: %v", yesNo)
// 	fmt.Println()
// 	rand.Seed(time.Now().UnixNano())
// 	fmt.Println(rand.Intn(100))
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// type myFloat64 float64

// func KelvinToCelsius(k float64) float64 {
// 	k -= 273.15
// 	return k
// }

// func main() {
// 	kelvin := 294.0
// 	celsius := KelvinToCelsius(kelvin)
// 	fmt.Print(kelvin, "度 k is ", celsius, "度 C")

// 	var temperature myFloat64 = 20
// 	fmt.Println(temperature)

// }
// type celsius float64
// type kelvin float64

// func A(k kelvin) celsius {
// 	return celsius(k - 273.15)
// }

// func (k kelvin) celsius() celsius {
// 	return celsius(k - 273.2)
// }
// func main() {

// }

// type kelvin float64

// func fakeSensor() kelvin {
// 	return kelvin(rand.Intn(151) + 150)
// }

// func realSensor() kelvin {
// 	return 0
// }

// func main() {
// 	sensor := fakeSensor
// 	fmt.Println(sensor())
// 	sensor = realSensor
// 	fmt.Println(sensor())

// 	measureTemperature(3, fakeSensor)
// }

// func measureTemperature(samples int, sensor func() kelvin) {
// 	for i := 0; i < samples; i++ {
// 		k := sensor()
// 		fmt.Printf("%v° K\n", k)
// 		time.Sleep(time.Second)
// 	}
// }

// func measureTemperatureN(samples int, s myFunc) {

// }
// func drawTable(rows myIntq, f myFuncTwo) {

// }

// type myFunc func() kelvin
// type myIntq int
// type myFuncTwo func(rows myIntq) (string, string)
// package main

// import (
// 	"fmt"
// 	"sort"
// 	//"sort"
// )

// type kelvin float64

// type sensor func() kelvin

// func calibrate(s sensor, offset kelvin) sensor {
// 	return func() kelvin {
// 		return s() + offset
// 	}
// }
// func realSensor() kelvin {
// 	return 9
// }

// // func main() {
// // 	sensor := calibrate(realSensor, 5)
// // 	fmt.Println(sensor())
// // }
// func main() {
// 	// var k kelvin = 294.0

// 	// sensor := func() kelvin {

// 	// 	return k
// 	// }
// 	// fmt.Println(sensor())
// 	// k++
// 	// fmt.Println(sensor())

// 	//planets := [...]string{"hello", "world", "my name", "is lili"}
// 	// for i := 0; i < len(planets); i++ {
// 	// 	planet := planets[i]
// 	// 	fmt.Println(i, planet)
// 	// }
// 	// for i, planet := range planets {
// 	// 	fmt.Println(i, planet)
// 	// }
// 	// fmt.Println(planets)
// 	//fmt.Println(planets[2:3])

// 	planets := []string{"hello", "world", "my name", "is lili", "1", "6", "dd"}
// 	planets = append(planets, "hehe")
// 	sort.StringSlice(planets).Sort()
// 	fmt.Println(planets)

// 	fmt.Println("容量", cap(planets), "长度", len(planets))
// 	dwarfs := make([]string, 10)
// 	fmt.Println(dwarfs)
// }

// type StringSlice []string

// func (p StringSlice) Sort() {}

package main

import (
	"fmt"
)

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	// twoWorlds := terraform("New", "Venus", "Mars")
	// fmt.Println(twoWorlds)

	// planes := []string{"Venus", "Mars", "Jupiter"}
	// newplanes := terraform("New", planes...)
	// fmt.Println(newplanes)

	// temperature := map[string]int{
	// 	"Earth": 15,
	// 	"Mars":  -65,
	// }

	// if moon, ok := temperature["Moon"]; ok {
	// 	fmt.Printf("......... is %v° C. \n", moon)
	// } else {
	// 	fmt.Println("where is the moon?")
	// }

	// temperatureNew := make(map[float64]int, 10)
	// fmt.Println(len(temperatureNew))

	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int)
	for _, t := range temperatures {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
}
