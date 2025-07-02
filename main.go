package main

import (
	"fmt"
	"net/http"
	//"runtime/trace"
)

var shortGolang = "Watch Go crash course"
var taskOne = "Watch TWN Golang Full Course"
var rewardDessert = "Reward myself with a cheesecake"

var taskItems = [20]string{shortGolang, taskOne, rewardDessert}

func main() {

	fmt.Println("##### Welcome to our Todolist App! #####")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil) //this starts the local server that will make the server available on the local machine

	//var shortGolang = "Watch Go crash course"
	//var taskOne = "Watch TWN Golang Full Course"
	//var rewardDessert = "Reward myself with a cheesecake"
	//
	//var taskItems = []string{shortGolang, taskOne, rewardDessert}
	//printTasks(taskItems)
	//
	//taskItems = addTask(taskItems, "Go for a run")
	//taskItems = addTask(taskItems, "Practice coding in Go")
	//fmt.Println("Updated List")
	//printTasks(taskItems)
	//printTasks() //empty because of function parameters
	////fmt.Println()

	//fmt.Println("List of my Todos")
	////fmt.Println("Tasks:", taskItems)
	//for index, task := range taskItems {
	//	//fmt.Println(index+1, ".", task)
	//	fmt.Printf("%d. %s \n", index+1, task)
	//
	//}

	//fmt.Println(shortGolang) // it adds a mini line at the end without adding it manually
	//fmt.Println(taskOne)
	//fmt.Println(rewardDessert)
	//
	//fmt.Println()
	//
	//fmt.Println("Tutorials")
	//fmt.Println(shortGolang)
	//fmt.Println(taskOne)
	//
	//fmt.Println()
	//
	//fmt.Println("Rewards")
	//fmt.Println(rewardDessert)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello User. Welcome to our Todolist App!"
	fmt.Fprintln(writer, greeting)
}
func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		//fmt.Println(index+1, ".", task)
		fmt.Printf("%d. %s \n", index+1, task)

	}

}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems

}
