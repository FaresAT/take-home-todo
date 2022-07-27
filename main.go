package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ToDos structure is an array of "ToDo" objects that contain a priority and a task
type ToDos struct {
	todos []ToDo
}

type ToDo struct {
	priority int
	task     string
}

func (t *ToDos) addToDo(prio int, tsk string) {
	for _, i := range t.todos {
		if i.priority == prio {
			fmt.Println("Cannot add todo, overlapping priority" + strconv.Itoa(prio) + " with " + i.task + "!")
		}
	}
	t.todos = append(t.todos, ToDo{priority: prio, task: tsk})
	fmt.Println("Successfully added todo " + tsk + " with priority " + strconv.Itoa(prio))
}

func (t *ToDos) viewToDos() {
	for _, i := range t.todos {
		fmt.Println(strconv.Itoa(i.priority) + ": " + i.task)
	}
}

func (t *ToDos) viewPriorities() {
	// this is terribly inefficient due to the unsorted nature of the todo struct
	var allPriorities []int
	for _, i := range t.todos {
		allPriorities = append(allPriorities, i.priority)
	}
	sort.Ints(allPriorities)

	missing := allPriorities[0]
	for i := range allPriorities {

		for missing < allPriorities[i]-i {
			fmt.Println(i + missing)
			missing += 1
		}
	}
}

func (t *ToDos) removeToDo(tsk string) bool {
	for index, i := range t.todos {
		if i.task == tsk {
			// creating a new todo list without the element that should be removed
			t.todos = append(t.todos[:index], t.todos[index+1:]...)
			return true
		}
	}
	return false
}

func main() {
	todoList := ToDos{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ToDo CLI Application")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")

		// golangs way of reading inputs from the terminal is kinda messy
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("add", text) == 0 {
			fmt.Println("What priority should this task be?")
			fmt.Print(" * ")
			prioStr, _ := reader.ReadString('\n')
			prioStr = strings.Replace(prioStr, "\n", "", -1)

			prioInt, err := strconv.Atoi(prioStr)

			if err != nil {
				log.Fatalf("Invalid number inputted to priority!")
				panic(err)
			}

			fmt.Println("What is this task?")
			fmt.Print(" * ")
			tsk, _ := reader.ReadString('\n')
			tsk = strings.Replace(tsk, "\n", "", -1)

			todoList.addToDo(prioInt, tsk)

		} else if strings.Compare("remove", text) == 0 {
			fmt.Println("What task should be removed?")
			fmt.Print(" * ")
			toRemove, _ := reader.ReadString('\n')
			toRemove = strings.Replace(toRemove, "\n", "", -1)

			if todoList.removeToDo(toRemove) {
				fmt.Println("Successfully removed item!")
			} else {
				fmt.Println("Couldn't remove item.")
			}

		} else if strings.Compare("view priorities", text) == 0 {
			fmt.Println("The unused priorities are:")
			todoList.viewPriorities()

		} else if strings.Compare("view tasks", text) == 0 {
			fmt.Println("The tasks are:")
			todoList.viewToDos()
		}
	}
}
