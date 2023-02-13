package main

import (
	"fmt"
	"sync"
)

var dogch = make(chan struct{})
var catch = make(chan struct{})
var fishch = make(chan struct{})

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	go dog( &wg)
	go cat( &wg)
	go fish( &wg)
	dogch <- struct{}{}
	
	wg.Wait()

}

func dog( wg *sync.WaitGroup) {
	count :=0
	for {

		<-dogch
		if count >= 2 {
			wg.Done()
			return
		}
		fmt.Println("dog")
		count++
		catch <- struct{}{}
	}
}
func cat( wg *sync.WaitGroup) {
	count :=0
	for {

		if count >= 2 {
			wg.Done()
			return
		}
		<-catch
		fmt.Println("cat")
		count++
		fishch <- struct{}{}
	}
}
func fish( wg *sync.WaitGroup) {
	count :=0
	for {

		if count >= 2 {
			wg.Done()
			return
		}
		<-fishch
		fmt.Println("fish")
		count++
		dogch <- struct{}{}
	}
}
