package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var Original_price float64 = 87.90

var target_Price float64 = 92

const option_price float64 = 8.65

var delta_value float64 = 0.9670 //delta value

var theta_value float64 = -0.0360 //theta value

func delta() float64 {
	milestone_price := target_Price - Original_price

	increment_value := float64(milestone_price * delta_value)
	//fmt.Println(increment_value)
	return increment_value

}

func theta() float64 {

	year, month, day := time.Now().Date()
	//fmt.Println(f, year, month)

	y := year
	if len(os.Args) == 2 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			y = i
		}
	}
	m := month
	d := time.Date(y, m+1, 1, 0, 0, 0, 0, time.UTC).Add(-24 * time.Hour)
	d = d.Add(-time.Duration((d.Weekday()+7-time.Friday)%7) * 24 * time.Hour)
	last_thu := d.Day()

	last_thu -= 1
	counted_time := last_thu - day

	decreased_price := float64(float64(counted_time) * float64(theta_value))
	//fut_price := Original_price + decreased_price
	//fmt.Println(fut_price)

	return decreased_price

}

func gamma() {
}

func main() {
	r := theta()
	de := delta()
	Predicted := option_price + de + r
	fmt.Println("The predicted price is ", Predicted)
}
