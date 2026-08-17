// package main

// func main() {
// 	demoFunctions()
// }

// package main

// import "fmt"

// func main() {
//     name := "Vishnu"
//     age := 20
//     Employee_Name := "Kunal"

//     fmt.Println(name)
//     fmt.Println(age)
//     fmt.Println(Employee_Name)
// }

//find the area of rectangle
// package main
// import "fmt"

// func main(){
//     length := 10
//     breath := 20
//     area_of_rectangle := length * breath

//     fmt.Println((area_of_rectangle))
// }

// find grosssalary
// package main
// import "fmt"
// func main(){
//     var basicsalary float64 = 50000
//     HRA := 0.20
//     DA := 0.10
//     grosssalary := basicsalary + HRA + DA

//     fmt.Println(basicsalary)
//     fmt.Println(HRA)
//     fmt.Println(DA)
//     fmt.Println(grosssalary)
// }

// student percentage calculator

// package main

// import "fmt"

// func main() {
// 	var name string
// 	var mark1, mark2, mark3, mark4, mark5 float64

// 	fmt.Print("Enter Student Name: ")
// 	fmt.Scan(&name)

// 	fmt.Print("Enter marks for Subject 1: ")
// 	fmt.Scan(&mark1)

// 	fmt.Print("Enter marks for Subject 2: ")
// 	fmt.Scan(&mark2)

// 	fmt.Print("Enter marks for Subject 3: ")
// 	fmt.Scan(&mark3)

// 	fmt.Print("Enter marks for Subject 4: ")
// 	fmt.Scan(&mark4)

// 	fmt.Print("Enter marks for Subject 5: ")
// 	fmt.Scan(&mark5)

// 	total := mark1 + mark2 + mark3 + mark4 + mark5
// 	percentage := total / 5

// 	fmt.Println("\n----- Student Result -----")
// 	fmt.Println("Student Name:", name)
// 	fmt.Println("Total Marks:", total)
// 	fmt.Println("Percentage:", percentage, "%")
// }






// package main
// import "fmt"
// type Movie struct {
// 	Title       string
// 	Director    string
// 	ReleaseYear int
// 	Rating      float32
// }

// func main() {
// 	movie := Movie{
// 		Title:       "3 Idiots",
// 		Director:    "Rajkumar Hirani",
// 		ReleaseYear: 2009,
// 		Rating:      8.4,
// 	}
// 	fmt.Println("Title:", movie.Title)
// 	fmt.Println("Director:", movie.Director)
// 	fmt.Println("Release Year:", movie.ReleaseYear)
// 	fmt.Println("Rating:", movie.Rating)
// }


// package main

// import "fmt"

// // Define the Laptop struct
// type Laptop struct {
// 	Brand     string
// 	RAM       int
// 	Storage   int
// 	Processor string
// }

// func main() {

// 	// Create a Laptop object
// 	laptop := Laptop{
// 		Brand:     "HP",
// 		RAM:       8,
// 		Storage:   512,
// 		Processor: "Intel Core i5",
// 	}

// 	// Display information before updating RAM
// 	fmt.Println("Laptop Information Before Update")
// 	fmt.Println("--------------------------------")
// 	fmt.Println("Brand      :", laptop.Brand)
// 	fmt.Println("RAM        :", laptop.RAM, "GB")
// 	fmt.Println("Storage    :", laptop.Storage, "GB")
// 	fmt.Println("Processor  :", laptop.Processor)

// 	// Update RAM from 8 GB to 16 GB
// 	laptop.RAM = 16

// 	// Display information after updating RAM
// 	fmt.Println("\nLaptop Information After Update")
// 	fmt.Println("-------------------------------")
// 	fmt.Println("Brand      :", laptop.Brand)
// 	fmt.Println("RAM        :", laptop.RAM, "GB")
// 	fmt.Println("Storage    :", laptop.Storage, "GB")
// 	fmt.Println("Processor  :", laptop.Processor)
// }


// package main 
// import "fmt"
// type Employee struct {
// 	Name string
// 	basicsalary float64
// 	Bonus float64
// }

// func calculateSalary(emp Employee) float64 {
// 	return emp.basicsalary + emp.Bonus
// }

// func main() {
// 	employee := Employee{
// 		Name:        "Vishnu",
// 		basicsalary: 50000,
// 		Bonus:       10000,
// 	}
// 	finalSalary := calculateSalary(employee)

// 	fmt.Println(employee.Name)
// 	fmt.Println(employee.basicsalary)
// 	fmt.Println(employee.Bonus)
// 	fmt.Println(finalSalary)
// }



// package main
// import "fmt"
// type Movie struct {
// 	Title   string
// 	Hours   int
// 	Minutes int
// }

// func calculateMinutes(movie Movie) int {
// 	return (movie.Hours * 60) + movie.Minutes
// }

// func main() {
// 	movie := Movie{
// 		Title:   "3 Idiots",
// 		Hours:   2,
// 		Minutes: 50,
// 	}
// 	totalMinutes := calculateMinutes(movie)
// 	fmt.Println( movie.Title)
// 	fmt.Println( movie.Hours)
// 	fmt.Println( movie.Minutes)
// 	fmt.Println( totalMinutes, "minutes")
// }



// package main
// import "fmt"

// type ElectricityBill struct{
// 	CustomerName string 
// 	Units int
// 	RatePerUnit float64
// }
// func calculateBill(bill ElectricityBill) float64 {
// 	return float64(bill.Units) * bill.RatePerUnit
// }

// func main() {
// 	bill := ElectricityBill{
// 		CustomerName:  "Rahul",
// 		Units: 250,
// 	}

// 	totalBill := calculateBill(bill)
// 	fmt.Println("Customer Name :", bill.CustomerName)
// 	fmt.Println("Units         :", bill.Units)
// 	fmt.Println("Rate Per Unit :", bill.RatePerUnit)
// 	fmt.Println("Total Bill    :", totalBill)
// }





// create a fucntion print number() that print number 1 to 5 . execute the fucntion using goroutine

// package main

// import "fmt"

// func printNumber() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	go printNumber()
// 	fmt.Scanln()
// }



//create two fucntion: processOrders() and then print "processing order 1" to "processingOrder to 5" , sendNotification() and then print "sendNotification 1" to "sendNotification 5" ,run both functuon using goroutine

// package main 
// import (
// 	"fmt"
// 	"sync"      
// )

// func processOrders(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i:=1; i<=5; i++ {
// 		fmt.Println("Processing order", i)
// 	}
// } 


// func sendNotification(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Send notification", i)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go processOrders(&wg)
// 	go sendNotification(&wg)

// 	wg.Wait()
// }



// updateInventory() → prints "Inventory Updated"
// sendEmail() → prints "Email Sent"
// sendNotification() → prints "Notification Sent"


// package main 
// import "fmt"
// func updateInventory() {
// 	fmt.Println("Inventory Updated")
// }

// func sendEmail() {
// 	fmt.Println("Email Sent")
// }

// func sendNotification() {
// 	fmt.Println("Notification Sent")
// }

// func main() {
// 	go updateInventory()
// 	go sendEmail()
// 	go sendNotification()

//	fmt.Scanln()    // Scanln() is used to wait for inout from the user




// CHANNEL - commmunication mechanish used by goroutine to send & receive the date, allows go routine to communicate without directly sharing the data.




// buffered and unbuffered channels.

// buffered - ch := make(chan int) - it has no storage space -


// Goroutine A                 Goroutine B
//      |                           |
//      | ---- 10 ---------------->|
//      |                           |
//    WAIT                       RECEIVE

// package main
// import "fmt"
// func main() {
// 	ch := make(chan int)

// go func() {
// 	ch <- 10
// }()

// fmt.Println(<-ch)

// }


// create a channel of type int. create a goroutine that sends the number 100 through the channel . recieves the number in main and print the recieved 100:


// package main 
// import "fmt"
// func main(){
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 100
// 	}()
// 	num := <-ch
// 	fmt.Println("recieved:", num)
// }


// craett a function calculationsquare() that recieves a number and se ds its square thorugh a channedl


// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter int
// var wg sync.WaitGroup

// func Increment() {
// 	defer wg.Done()

// 	for i := 0; i < 1000; i++ {
// 		counter++
// 	}
// }

// func main() {
// 	wg.Add(2)

// 	go Increment()
// 	go Increment()

// 	wg.Wait()

// 	fmt.Println(counter)
// }



//write a go programe with start go routine . the go routine should print "hello from goroutine use sunc.waitgroup so main() waits for the goroutine to finish


// package main
// import (
//     "fmt"
//     "sync"
// )

// var wg sync.WaitGroup
// func Hello() {
//     defer wg.Done()
//     fmt.Println("Hello from goroutine")

// }

// func main() {
//     wg.Add(1)
//     go Hello()
//     wg.Wait()
// }


// package main 
// import (
// 	"fmt"
// 	"sync"
// )


// var number = 100

// var rw sync.RWMutex
// func Read()  {
// 	rw.RLock()
// 	fmt.Println(number)
// 	rw.RUnlock()
// }


// func Write(){
// 	rw.Lock()
// 	number++
// 	rw.Unlock()
// }

// func ()  {
	
// }




// Create a shared slice called message. Start three goroutines. Each goroutine should add one message : "hello", "welcome", and "Good Mornig". Use a Mutex to safely modify the shared slice.


// package main 
// import (
// 	"fmt"
// 	"sync"
// )

// var messages []string
// var mutex sync.Mutex
// var wg sync.WaitGroup

// func addMesage(message string) {
// 	defer wg.Done()

// 	mutex.Lock()
// 	messages = append(messages, message)
// 	mutex.Unlock()
// }

// func main() {
// 	wg.Add(3)

// 	go addMesage("Hello")
// 	go addMesage("Welcome")
// 	go addMesage("Good Morning")
// 	wg.Wait()
// 	fmt.Println(messages)
// }

//the project is i want to build an app 

//create four goroutines representing email, payment, notification, and search services . Each service should increase a shared messagecount by 1. use a mutex


package main 
import (
	"fmt"
	"sync"
)


var messagecount int 
var mutex sync.Mutex 
var wg sync.WaitGroup

func service(name string) {
	defer wg.Done()
	mutex.Lock()
	messagecount++
	mutex.Unlock()
	fmt.Println(name, "service completed")
	
}

func main()  {
	wg.Add(4)

	go service("Email")
	go service("Payment")
	go service("notification")
	go service("Search Services")
	wg.Wait()
	fmt.Println("Final message count:", messagecount)
}
