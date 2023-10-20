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

// Retries are a premium feature now! Textio's free users only get 1
// retry message, while pro members get an unlimited amount.

// Complete the getMessageWithRetriesForPlan function. It takes a
// plan variable as input, and you've been provided with constants for
// the plan types at the top of the program.
// * If the plan is a pro plan, return all the strings from
// getMessageWithRetries().
// * If the plan is a free plan, retrun the first 2 strings from
// getMessageWithRetries().
// * If the plan isn't either of those, return an error that says
// unsupported plan.

package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsopported plan")
}

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func test(name string, doneAt int, plan string) {
	defer fmt.Println("====================================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sneding: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("no response")
		}
	}
}

func main() {
	test("Ozgur", 3, planFree)
	test("Jeff", 3, planPro)
	test("Sally", 2, planPro)
	test("Sally", 3, "no plan")
}
