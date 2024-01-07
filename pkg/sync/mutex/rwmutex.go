package mutex

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.RWMutex

func readData(i int) {
	fmt.Printf("Try to get R lock: %v \n", i)
	mutexInfo("TryRLock: ")
	rw.RLock()
	fmt.Printf("Get R lock: %v \n", i)
	mutexInfo("GetRLock: ")
	defer rw.RUnlock()
	time.Sleep(1 * time.Second)
	innerRead(i)
}

func innerRead(i int) {
	mutexInfo("TryInnerRLock: ")
	fmt.Printf("Try to get Inner R lock: %v \n", i)
	rw.RLock()
	mutexInfo("GetInnerRLock: ")
	fmt.Printf("Get inner R lock: %v \n", i)
	defer rw.RUnlock()
	time.Sleep(5 * time.Second)
	fmt.Printf("Finish read something: %v \n", i)
}

func writeData(i int) {
	fmt.Printf("Try to get W lock: %v \n", i)
	mutexInfo("TryWLock: ")
	rw.Lock()
	fmt.Printf("Get W lock: %v \n", i)
	mutexInfo("GetWLock: ")
	defer rw.Unlock()
	time.Sleep(5 * time.Second)
	fmt.Printf("Finish write something: %v \n", i)
}

func mutexInfo(pre string) {
	fmt.Printf("%s: RWMutext stats: %#v \n", pre, rw)
}

func RWCompetion() {
	rw.Lock()
	fmt.Printf("Ready??!! \n")
	go readData(1)
	go writeData(2)
	fmt.Printf("Go ... \n")
	rw.Unlock()
	time.Sleep(20 * time.Second)
}
