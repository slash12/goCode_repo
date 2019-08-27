package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// fmt.Println("Hello World")

	// 26/08/19
	numSlice := []int{5, 4, 3, 2, 1} //array definition
	numSlice3 := make([]int, 5, 10)  //int array created, set to zero
	copy(numSlice3, numSlice)        //copy of array numSlice
	// append new item to array, to the last item
	numSlice3 = append(numSlice3, -1)
	// fmt.Println(numSlice3)

	// Map (a collection of data)
	presAge := make(map[string]int)
	presAge["test"] = 42
	presAge["test2"] = 422
	// fmt.Println(presAge["test"])
	// get the length of the map
	// fmt.Println(len(presAge))
	// delete an item from the map
	delete(presAge, "test2")
	// fmt.Println(len(presAge))

	// FUNCTION
	// listNums := []float64{1, 2, 3, 4, 5} //array of numbers
	// fmt.Println("Sum :", addThem(listNums)) //calling and printing the function result

	// create a function that return two output
	// num1, num2 := nxt2Values(5)
	// fmt.Println(num1, num2)

	// function that accept a range of parameters
	// fmt.Println(subtractThem(1, 2, 3, 4, 5))

	// function inside another function
	// closure definition
	// num3 := 3
	// doubleNum := func() int {
	// 	num3 *= 2
	// 	return num3
	// }
	// fmt.Println(doubleNum())

	// recursion (an inside function that called itself)
	// fmt.Println(factorial(5))

	// defer (prioritise another function over a function)
	// in this case, the function 'printOne' is going to be prioritised
	// defer printTwo()
	// printOne()

	// safe division function (try and catch model)
	// fmt.Println(safeDiv(3, 0))
	// fmt.Println(safeDiv(3, 2))

	// handle panic
	// demoPanic()

	// Pointer (main: 140)
	// x := 0
	// changeXVal(&x)
	// fmt.Println("x = ", x)
	// to print out the memory address of variable x
	// fmt.Println("Memory Address for x =", &x)

	// Make use of defined struct
	// rect1 := Rectangle{0, 50, 10, 10}
	// fmt.Println(rect1.width)
	// fmt.Println("Rectangle Area = ", rect1.area())

	// Make use of interface
	// rect := Rectangle{20, 50}
	// circ := Circle{4}
	// fmt.Println("Rectangle Area = ", getArea(rect))
	// fmt.Println("Circle Area = ", getArea(circ))

	// String import library
	// sampString := "Hello World"
	// fmt.Println(strings.Contains(sampString, "lo"))
	// fmt.Println(strings.Index(sampString, "lo"))
	// fmt.Println(strings.Count(sampString, "l"))
	// fmt.Println(strings.Replace(sampString, "l", "x", 3))
	// split
	// csvString := "1,2,3,4,5,6"
	// fmt.Println(strings.Split(csvString, ","))
	// listOfLetters := []string{"c", "a", "b"}
	// sort.Strings(listOfLetters) //sort the letters
	// fmt.Println("Letters: ", listOfLetters)

	// 270819
	// listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")
	// fmt.Println(listOfNums)

	// create a new file
	// file, err := os.Create("samp.txt")
	// record any error occuring
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// write a line to the file
	// file.WriteString("This is some text")
	// close the file
	// file.Close()
	// next part is to read the file
	// stream, err := ioutil.ReadFile("samp.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// we can convert the file read (output) into a string
	// readString := string(stream)
	// fmt.Println(readString)

	// convert values
	// randInt := 5
	// randFloat := 10.5
	// randString := "100"
	// randString2 := "205.5"
	// convert an int into a float
	// fmt.Println(float64(randInt))
	// convert a float into an int
	// fmt.Println(int(randFloat))
	// convert a string into an int
	// newInt, _ := strconv.ParseInt(randString, 0, 64)
	// fmt.Println(newInt)
	// convert a string into a float
	// newFloat, _ := strconv.ParseFloat(randString2, 64)
	// fmt.Println(newFloat)

	// create servers
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/earth", handler2)
	// what port to listen to
	// http.ListenAndServe(":8080", nil)

	// go routine (allow us to run different function in parallel)
	// for i := 0; i < 10; i++ {
	// 	go count(i)
	// }
	// time.Sleep(time.Millisecond * 11000)

	// make a channel
	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makePizza(stringChan)
		go makeSauce(stringChan)
		go addTopping(stringChan)
		time.Sleep(time.Millisecond * 5000)
	}
}

// channels (pass data to go routine)
var pizzaNum = 0
var pizzaName = ""

func makePizza(stringChan chan string) {
	// whenever the func is called, pizzaNum is inceased by 1
	pizzaNum++
	// build the pizza name, convert pizzaNum to string
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	fmt.Println("Make Pizza and send for sauce")
	// pass the pizzaName to the channel
	stringChan <- pizzaName
	// sleep the process
	time.Sleep(time.Millisecond * 10)
}

func makeSauce(stringChan chan string) {
	// pass the value that is passed in the channel
	pizza := <-stringChan
	fmt.Println("Add Sauce and Send", pizza, "for topping")
	// pass the pizzaName to the channel
	stringChan <- pizzaName
	// sleep the process
	time.Sleep(time.Millisecond * 10)
}

func addTopping(stringChan chan string) {
	// pass the value that is passed in the channel
	pizza := <-stringChan
	fmt.Println("Add Toppings to", pizza, "and ship")
	time.Sleep(time.Millisecond * 10)
}

func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)
		// sleep (in another word to pause) for 1k ms
		time.Sleep(time.Millisecond * 1000)
	}
}

// this is used to write a response when accessing the server
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\nTwo")
}

// new interface
type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

// // Define our own datatype with 'Struct'
// type Rectangle struct {
// 	leftX  float64
// 	topY   float64
// 	height float64
// 	width  float64
// }

// func (rect *Rectangle) area() float64 {
// 	return rect.width * rect.height
// }

// function name is 'addThem', param name is 'numbers' of type float64
// return type is float 64
func addThem(numbers []float64) float64 {
	// declaring a 'sum' variable
	sum := 0.0

	// for loop with no index (the purpose of '_')
	// range refers to all 'numbers'
	for _, val := range numbers {
		sum += val
	}

	return sum
}

// function that return two outputs
func nxt2Values(number int) (int, int) {
	return number + 1, number + 2
}

// a function that accept an infinite range of parameters
func subtractThem(args ...int) int {
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}

// recursion function
func factorial(num int) int {
	// if statement
	if num == 0 {
		return 1
	}

	// calling a function from inside of itself
	return num * factorial(num-1)
}

// sample functions used for the demo 'defer' keyword (main: 51)
func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

// safe way to do a division (similar to a try and catch structure)
func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	// If ever the num2 is equal to zero, this is triggered an error
	// which will be caught by 'recover()' function and '<nil>' is returned
	// basically what 'recover()' function does, is that it continue the operation
	// even when there is a runtime error

	solution := num1 / num2
	return solution
}

// this function shows how a voluntery error can be caused and then handled
// the function 'panic' is used to render an error and 'defer' keyword is used
// to handle that error
func demoPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("PANIC")
}

// demo to change the referred variable value (actual value)
func changeXVal(x *int) {
	*x = 2
}
