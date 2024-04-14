package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// a waitGroup acts as a counter
var wg = sync.WaitGroup{}

var m = sync.Mutex{}

var m2 = sync.RWMutex{} //Same functionality as mutex, but adds a RLock() and a RUnlock() method

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	t0 = time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1) //adds one to the number of tasks to wait for
		go dbCallAsync(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time concurrently: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	//simulates db call delay
	//var delay float32 = rand.Float32() * 2000
	var delay float32 = 2000 //alters the result (several threads trying to write at the same time)  -->
	// To avoid this we should use a Mutex (mutual exclusion)
	time.Sleep(time.Duration(delay) * time.Millisecond)

	fmt.Println("\nTe result from de database is: ", dbData[i])

}

func dbCallAsync(i int) {
	//simulates db call delay
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	//fmt.Println("\nTe result from de database is: ", dbData[i])
	save(dbData[i])
	log()
	wg.Done() //sets a task as done (reduces the number of tasks to wait for in one)
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m2.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m2.RUnlock()
}

//45:50 performance improvement
