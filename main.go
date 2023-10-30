//// 2023/10/20 // 12:08 //

//// Arrays in Go

// // Arrays are fixed-size groups of variables of the same type.

// // The type [n]T is an array of n values of type T

// // To declare an array of 10 integers:
// var myInts [10]int

// // or to declare an initialized literal:
// primes := [6]int{2, 3, 5, 7, 11, 13}

//// Assignment

// When a message isn't responded to, we allos out clients to have up to
// 2 additional messages that are sent as nudging reminders.

// getMessageWithRetries return an array of 3 strings where index 0
// is the first message. If the first message isn't answered by the
// recipient, we send the second, if that one isn't answered then we
// send the third.

// // Update getMessageWithRetries to return the following 3 strings in
// // an array.
// // * click here to sign up
// // * pretty clease click here
// // * we beg you to sing up

// package main

// import "fmt"

// const (
// 	retry1 = "click here to sign up"
// 	retry2 = "pretty please click here"
// 	retry3 = "we beg you to sign up"
// )

// func getMessageWithRetries() [3]string {
// 	return [3]string{
// 		retry1,
// 		retry2,
// 		retry3,
// 	}
// }

// func testSend(name string, doneAt int) {
// 	fmt.Printf("sending to %v...", name)
// 	fmt.Println()

// 	messages := getMessageWithRetries()
// 	for i := 0; i < len(messages); i++ {
// 		msg := messages[i]
// 		fmt.Printf(`sending: "%v"`, msg)
// 		fmt.Println()
// 		if i == doneAt {
// 			fmt.Println("they responded!")
// 			break
// 		}
// 		if i == len(messages)-1 {
// 			fmt.Println("complete failure")
// 		}
// 	}
// }

// func main() {
// 	testSend("Bob", 0)
// 	testSend("Alice", 1)
// 	testSend("Mangalam", 2)
// 	testSend("Ozgur", 3)
// }

//// 12:27

//// Slices in Go

// // 99 times out of 100 you will use a slice instead of an array when
// // working with ordered lists.

// // Arrays are fixed in size. Once you make an array like [10]int you
// // can't add an 11th element.

// // A slice is a dynamically-sized, flexible view of the elements of an
// // array.

// // Slices always have an underlying array, through it isn't always
// // specified explicitly. To explicitly create a slice on top of an array we
// // can do:
// primes := [6]int{2, 3, 5, 7, 11, 13}
// mySlice := primes[1:4]
// // mySlice = {3, 5, 7}

// // The syntax is:
// arrayname[lowIndex:highIndex]
// arrayname[lowIndex:]
// arrayname[:highIndex]
// arrayname[:]
// // Where lowIndex is inclusive and highIndex is exclusive

// // Either lowIndex or highIndex or both can be omitted to use the
// // entire array on that side.

//// Assignmnet

// // Retries are a premium feature now! Textio's free users only get 1
// // retry message, while pro members get an unlimited amount.

// // Complete the getMessageWithRetriesForPlan function. It takes a
// // plan variable as input, and you've been provided with constants for
// // the plan types at the top of the program.
// // * If the plan is a pro plan, return all the strings from
// // getMessageWithRetries().
// // * If the plan is a free plan, retrun the first 2 strings from
// // getMessageWithRetries().
// // * If the plan isn't either of those, return an error that says
// // unsupported plan.

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// const (
// 	planFree = "free"
// 	planPro  = "pro"
// )

// func getMessageWithRetriesForPlan(plan string) ([]string, error) {
// 	allMessages := getMessageWithRetries()
// 	if plan == planPro {
// 		return allMessages[:], nil
// 	}
// 	if plan == planFree {
// 		return allMessages[0:2], nil
// 	}
// 	return nil, errors.New("unsopported plan")
// }

// func getMessageWithRetries() [3]string {
// 	return [3]string{
// 		"click here to sign up",
// 		"pretty please click here",
// 		"we beg you to sign up",
// 	}
// }

// func test(name string, doneAt int, plan string) {
// 	defer fmt.Println("====================================")
// 	fmt.Printf("sending to %v...", name)
// 	fmt.Println()

// 	messages, err := getMessageWithRetriesForPlan(plan)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	for i := 0; i < len(messages); i++ {
// 		msg := messages[i]
// 		fmt.Printf(`sneding: "%v"`, msg)
// 		fmt.Println()
// 		if i == doneAt {
// 			fmt.Println("they responded!")
// 			break
// 		}
// 		if i == len(messages)-1 {
// 			fmt.Println("no response")
// 		}
// 	}
// }

// func main() {
// 	test("Ozgur", 3, planFree)
// 	test("Jeff", 3, planPro)
// 	test("Sally", 2, planPro)
// 	test("Sally", 3, "no plan")
// }

//// 12:40

// //// 4-slices_no_array

// //// Make

// // Most of the time we don't need to think about the underlying array
// // of a slice. We can create a new slice using the make function:

// // func make([]T, len, cap) []T
// mySlice := make([]int, 5, 10)

// // the capacity argument is usually omitted and defaults
// mySlice := make([]int, 5)

// // Slices create with make will be filled with the zero value of the
// // type.

// // If we want to create a slice with a specific set of values, we can use a
// // slice literal:
// mySlice := []string{"I", "love", "go"}

// // Note that the array brackets ddo not have a 3 in them. If they did,
// // you'd have an array instead of a slice.

// // Length

// // The length of a slice is simply the number of elements it contains. It
// // is accessed using the built-in len() function:
// mySlice := []string{"I", "love", "go"}
// fmt.Println(len(mySlice)) // 3

// // Capacity

// // The capacity of a slice is the number of elements in the underlying
// // array, counting from the first element in the slice. It is accessed
// // using the build-in cap() function:
// mySlice := []string{"I", "love", "go"}
// fmt.Println(cap(mySlice)) // 3

// // Generally speaking, unless you're hyper-optimizing the memory
// // usage of your program, you don't need to worry about the capacity
// // of a slice because it will automatically grow as needed.

// // Assignment

// // We send a lot of text messages at Textio, and out API is getting slow
// // and unresponsive.

// // If we know the rough size of a slice before we fill it up, we can make
// // out program faster by creating the slice wit hthat size ahead of time
// // so that the Go runtime doesn't need to continuously allocate new
// // underlying arrays of larger and larger sizes. By setting the length,
// // the slice can still be resized later, but it means we can avoid all the
// // expensive resizing we know that we'll need.

// // Complete the getMessageCosts() function. It takes a slice of
// // messages and retruns a slcie of message costs.
// // 1. Preallocate a slice for the message costs of the same length as
// // the message slice.
// // 2. Fill the costs slice with costs for each message. The cost in the
// // cost slice should correspond to the message in the messages
// // slice at the same index. The cost of a message is the length of
// // the message multiplied by 0.01.

// package main

// import "fmt"

// func getMessageCosts(messages []string) []float64 {
// 	messageCosts := make([]float64, len(messages))
// 	for i := 0; i < len(messages); i++ {
// 		cost := float64(len(messages[i])) * 0.01
// 		messageCosts[i] = cost
// 	}
// 	return messageCosts
// }

// func test(messages []string) {
// 	costs := getMessageCosts(message)
// 	fmt.Println("Messages:")
// 	for i := 0; i < len(messages); i++ {
// 		fmt.Printf(" - %v\n", messages[i])
// 	}
// 	fmt.Println("Costs:")
// 	for i := 0; i < len(costs); i++ {
// 		fmt.Printf(" - %.2f\n", costs[i])
// 	}
// 	fmt.Println("==== END REPORT =====")
// }

// func main() {
// 	test([]string{
// 		"Welcome to the movies!",
// 		"Enjoy your popcorn!",
// 		"Please don't talk during the movie!",
// 	})
// 	test([]string{
// 		"I don't want to be here anymore",
// 		"Can we go home?",
// 		"I'm hungry",
// 		"I'm bored",
// 	})
// 	test([]string{
// 		"Hello",
// 		"Hi",
// 		"Hey",
// 		"Hi there",
// 		"Hey there",
// 		"Hi there",
// 		"Hello there",
// 		"Hey there",
// 		"Hello there",
// 		"General Kenobi",
// 	})
// }

//// 2023/10/27 // 20:57

// //// Variadic

// // Many function, especially those in the standard library, can take an
// // arbitrary number of final arguments. This is accomplished by using
// // the "..." syntax in the function signature.

// // A variadic function receives the variadic arguments as a slice.
// func sum(nums ...int) int {
// 	// nums is just a slice
// 	for i := 0; i < len(nums); i++ {
// 		num := nums[i]
// 	}
// }

// func main() {
// 	total := sum(1, 2, 3)
// 	fmt.Println(total)
// 	// prints "6"
// }

// // The familiar fmt.Println() and fmt.Sprintf() are variadic!
// // fmt.Println() prints each element with space delimiters and a
// // newline at the end.
// func Println(a ...interface{}) (n int, err error)

// //// Spread operator

// // The spread operator allows us to pass a slice into a variadic function.
// // The spread operator consists of three dots following the slice in the
// // function call.
// func printStrings(strings ...string) {
// 	for i := 0; i < len(strings); i++ {
// 		fmt.Println(strings[i])
// 	}
// }

// func main() {
// 	names := []string{"bob", "sue", "alice"}
// 	printStrings(names...)
// }

//// Assignment

// We need to sum up the costs of all individual messages so we can
// send an end-of-month bill to our customers.

// Complete the sum function so that returns the sum of all of its
// inputs.

// Note

// // Make note of how the variadic inputs and the spread operator are
// // used in the test suite.

// package main

// import "fmt"

// func sum(nums ...float64) float64 {
// 	sum := 0.0
// 	for i := 0; i < len(nums); i++ {
// 		num := nums[i]
// 		sum += num
// 	}
// 	return sum
// }

// // don't edit below this line

// func test(nums ...float64) {
// 	total := sum(nums...)
// 	fmt.Printf("Summing %v costs...\n", len(nums))
// 	fmt.Printf("Bill for the month: %.2f\n", total)
// 	fmt.Println("===== END REPORT =====")
// }

// func main() {
// 	test(1.0, 2.0, 3.0)
// 	test(1.0, 2.0, 3.0, 4.0, 5.0)
// 	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
// 	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
// }

//// 2023/10/30 // 20:25 //
//// Append

// // The built-in append function is used to dynamically add elements to
// // a slice:
// func append(slice []Type, elems ...Type) []Type

// // If the underlying array is not large enough, append() will create a
// // new underlying array and poitn the slice to it.

// // Notice that append() is variadic, the following are valid:
// slice = append(slice, oneThing)
// slice = append(slice, firstThing, secondThing)
// slice = append(slice, anotherSlice...)

// // Assignment

// // We've been asked to "bucket" costs for an entire month into the
// // cost occurring on each day of the month.

// // Complete the getCostsByDay function. It should return a slice of
// // float64s, where each element is the total cost for that day. The
// // length of the slice should be equal to the number of days
// // represented in the costs slice, including any days that have no
// // costs, up to the last day represented in the slice.

// // Days are numbered and start at 0.

// // Example

// // Input in day, cost format:
// []cost{
// 	{0, 4.0},
// 	{1, 2.1},
// 	{1, 3.1},
// 	{5, 2.5},
// }

// // // It would expect this resutl:
// // []float64{
// // 	4.0,
// // 	5.2,
// // 	0.0,
// // 	0.0,
// // 	0.0,
// // 	2.5,
// // }

// package main

// import "fmt"

// type cost struct {
// 	day   int
// 	value float64
// }

// func getCostsByDay(costs []cost) []float64 {
// 	costsByDay := []float64{}
// 	for i := 0; i < len(costs); i++ {
// 		cost := costs[i]
// 		for cost.day >= len(costsByDay) {
// 			costsByDay = append(costsByDay, 0.0)
// 		}
// 		costsByDay[cost.day] += cost.value
// 	}
// 	return costsByDay
// }

// func test(costs []cost) {
// 	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
// 	costsByDay := getCostsByDay(costs)
// 	fmt.Println("Costs by day:")
// 	for i := 0; i < len(costsByDay); i++ {
// 		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
// 	}
// 	fmt.Println("===== END REPORT =====")
// }

// func main() {
// 	test([]cost{
// 		{0, 1.0},
// 		{1, 2.0},
// 		{1, 3.1},
// 		{2, 2.5},
// 		{3, 3.6},
// 		{3, 2.7},
// 		{4, 3.34},
// 	})
// 	test([]cost{
// 		{0, 1.0},
// 		{10, 2.0},
// 		{3, 3.1},
// 		{2, 2.5},
// 		{1, 3.6},
// 		{2, 2.7},
// 		{4, 56.34},
// 		{13, 2.34},
// 		{28, 1.34},
// 		{25, 2.34},
// 		{30, 4.34},
// 	})
// }

//// 20:39
//// Slice of slices

// // Slices can hold otehr slices, effectively creating a matrix, or 2D slice.
// rows := [][]int{}

// Assignment

// We support various graphs and dashboards on Textio that display
// message analytics to our users. The UI for our graphs and charts is
// built on top of a grid system. Let's build some grid logic.

// Complete the createMatrix function. It takes a number of rows and
// columns and returns a 2D slice of integers where the value of each
// cell is i * j where i and j are the indexes of the row and
// column respectively.

// // For example, a 10x10 matrix would look like this:
// // [0 0 0 0 0 0 0 0 0 0]
// // [0 1 2 3 4 5 6 7 8 9]
// // [0 2 4 6 8 10 12 14 16 18]
// // [0 3 6 9 12 15 18 21 24 27]
// // [0 4 8 12 16 20 24 28 32 36]
// // [0 5 10 15 20 25 30 35 40 45]
// // [0 6 12 18 24 30 36 42 48 54]
// // [0 7 14 21 28 35 42 49 56 63]
// // [0 8 16 24 32 40 48 56 64 72]
// // [0 9 18 27 36 45 54 63 72 81]

// package main

// import "fmt"

// func createMatrix(rows, cols int) [][]int {
// 	matrix := [][]int{}
// 	for i := 0; i < rows; i++ {
// 		row := []int{}
// 		for j := 0; j < cols; j++ {
// 			row = append(row, i*j)
// 		}
// 		matrix = append(matrix, row)
// 	}
// 	return matrix
// }

// func test(rows, cols int) {
// 	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
// 	matrix := createMatrix(rows, cols)
// 	for i := 0; i < len(matrix); i++ {
// 		fmt.Println(matrix[i])
// 	}
// 	fmt.Println("===== END REPORT =====")
// }

// func main() {
// 	test(3, 3)
// 	test(5, 5)
// 	test(10, 10)
// 	test(15, 15)
// }

//// 20:54
//// Tricky Slices

// // The append() function changes the underlying array of its
// // parameter AND returns a new slice. This means that using append()
// // on anything other than itself is usually a BAD idea.

// // dont do this!
// someSlice = append(otherSlice, element)

// // Take a look at these head-scratchers:

// // Example 1: Works as expected
// a := make([]int, 3)
// fmt.Println("len of a:", len(a))
// // len of a: 3
// fmt.Println("cap of a:", cap(a))
// // cap of a: 3
// fmt.Println("appending 4 to b from a")
// // appending 4 to b from a
// b := append(a, 4)
// fmt.Println("b:", b)
// // b: [0 0 0 4]
// fmt.Println("addr of b:", &b[0])
// // addr of b: 0x44a0c0
// fmt.Println("appending 5 to c from a")
// // appending 5 to c from a
// c := append(a, 5)
// fmt.Println("addr of c:", &c[0])
// // addr of c: 0x44a180
// fmt.Println("a:", a)
// // a: [0 0 0]
// fmt.Println("b:", b)
// // b: [0 0 0 4]
// fmt.Println("c:", c)
// // c: [0 0 0 5]

// // With slices a, b, and c, 4 and 5 seem to be appended as we
// // would expect. We can even check the memory addresses and
// // confirm that b and c point to different underlying arrays.

// // Example 2: Something fishy
// i := make([]int, 3, 8)
// fmt.Println("len of i:", len(i))
// // len of i: 3
// fmt.Println("cap of i:", cap(i))
// // cap of i: 8
// fmt.Println("appending 4 to j from i")
// // appending 4 to j from i
// j := append(i, 4)
// fmt.Println("j:", j)
// // j: [0 0 0 4]
// fmt.Println("addr of j:", &j[0])
// // addr of j: 0x454000
// fmt.Println("appending 5 to g from i")
// // appending 5 to g from i
// g := append(i, 5)
// fmt.Println("addr of g:", &g[0])
// // addr of g: 0x454000
// fmt.Println("i:", i)
// // i: [0 0 0]
// fmt.Println("j:", j)
// // j: [0 0 0 5]
// fmt.Println("g:", g)
// // g: [0 0 0 5]

// // In this example howeer, when 5 is appended to g it overwrites
// // j's fourth index because j and g point to the same underlying array.
// // The append() function only creates a new array when there
// // isn't any capacity left. We created i with a length of 3 and a
// // capacity of 8, which means we can append 5 items before a new
// // array is automatically allocated.

// // Again, to avoid bugs like this, you should always use the append
// // function on the same slice the result is assigned to:
// mySlice := []int{1, 2, 3}
// mySlcie = append(mySlice, 4)

//// 13:09
//// Range

// Range

// // Go provides syntactic sugar to iterate easily over elements of a slice:
// for INDEX, ELEMENT := range SLICE {
// }

// // For example:
// fruits := []string{"apple", "banana", "grape"}
// for i, fruit := range fruits {
// 	fmt.Println(i, fruit)
// }
// // 0 apple
// // 1 banana
// // 2 grape

// Assignment

// We need to be able to quickly detect bad words in the messages our
// system sends.

// Complete the indexOfFirstBadWord function. If it finds any bad
// words in the message it should return the index of the first bad
// word in the msg slice. This will help us filter out naighty words from
// out messaging system. If no bad words are found, return -1 instead

// Use the range keyword.
package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func test(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("===================================")
}

func main() {
	badWords := []string{"crap", "shoot", "dang", "frick"}
	message := []string{"hey", "there", "john"}
	test(message, badWords)

	message = []string{"ugh", "oh", "my", "frick"}
	test(message, badWords)

	message = []string{"what", "the", "shoot", "I", "hate", "that", "crap"}
	test(message, badWords)
}
