package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			//Make Body
			car := &Car{}
			car.Body = "Sport car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh chan *Car, paintCh chan *Car) {

	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter Tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		//Make a Body
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)

		fmt.Printf("%.2f Complete Car: %s %s %s \n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
