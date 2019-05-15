package main 
import "fmt"
import "time"

func main() {
	i := 2 
	fmt.Println("write", i, "as")
	switch i {
	case 1: 
		fmt.Println("One")
	case 2: 
		fmt.Println("Two")
	case 3: 
		fmt.Println("three")
	}


	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default: 
		fmt.Println("its a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12: 
		fmt.Println("its before noon")
	default: 
		fmt.Println("its after noon")
	}



	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Ima a boolean")
		case int : 
			fmt.Println("Im an int")
		default: 
			fmt.Println("Type unknown", t)
		}	
	}
	whatAmI(true)
	whatAmI(4)
	whatAmI("hello world")
}