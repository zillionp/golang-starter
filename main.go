package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Password string

func (p Password) Validate() error {
	if len(p) < 5 {
		return errors.New("Password is too short")
	}
	return nil
}

type Person struct {
	Name string
	Age  int
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func greetPerson() {
	p := Person{
		Name: "Mandip",
		Age:  26,
	}
	password := Password("Name123")
	if err := password.Validate(); err != nil {
		panic(err)
	}

	fmt.Println("Hello " + p.Name)
}

func LongRunningOperation() {
	time.Sleep(time.Second * 2)
}

// Multithreading using golang
// One thing about multi-threaded logic is that
// you have to know when things are done
func main() {

	greetPerson()

	// It will take total of 6 seconds to complete the execution
	fmt.Println("Sync: Start")
	for i := 0; i < 3; i++ {
		LongRunningOperation()
	}
	fmt.Println("Sync: Stop")

	// It will take 2 seconds to complete
	// We can use go routines to run
	// these multiple long running processes concurrently
	// We have wait group and for each iteration we are adding
	// one to the wait group using Add and using Done for its completion
	wg := &sync.WaitGroup{}
	fmt.Println("Async: Start")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			LongRunningOperation()
			wg.Done()
		}()
	}
	// Wait until all operations are done
	wg.Wait()
	fmt.Println("Async: Stop")

}
