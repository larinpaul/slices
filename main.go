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

package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	messageCosts := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost := float64(len(messages[i])) * 0.01
		messageCosts[i] = cost
	}
	return messageCosts
}

func test(messages []string) {
	costs := getMessageCosts(message)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("==== END REPORT =====")
}

func main() {
	test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}
