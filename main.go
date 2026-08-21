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



// Read and write 


// package main 
// import (
// 	"fmt"
// 	"sync"
// )


// var number = 100

// var rw sync.RWMutex
// var wg sync.WaitGroup

// func Read()  {
// 	defer wg.Done()
// 	rw.RLock()
// 	fmt.Println(number)
// 	rw.RUnlock()
// }
// func Write(){
// 	defer wg.Done()
// 	rw.Lock()
// 	number++
// 	rw.Unlock()
// }
// func main()  {
// 	wg.Add(2)
// 	go Read()
// 	go Write()
// 	wg.Wait()
// }


// what is goroutine - It is a lightweight thread managed entirely by the Go runtime, rather than your computer's operating system.

// What is Mutex - it is a synchronization tool used in programming to prevent multiple goroutines from accessing or modifying the same piece of memory (like a variable , map or slice) at the exact same fraction of a second. 
// how does it works - if a worker (goroutine) wants to use the bathroom (modify data),  the must lock the door using the key (mu.Lock()).if another person worker arrives, they see the locked dooe and must wait in line , when the first worker is finished, they open the door and return the key (mu.Unlock()). , the next worker in line can now grab the key and enter.
 
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


// package main 
// import (
// 	"fmt"
// 	"sync"
// )


// var messagecount int 
// var mutex sync.Mutex 
// var wg sync.WaitGroup

// func service(name string) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	defer mutex.Unlock()
// 	messagecount++
// 	fmt.Println(name, "service completed")
	
// }

// func main()  {
// 	wg.Add(4)
// 	go service("Email")
// 	go service("Payment")
// 	go service("Notification")
// 	go service("Search Services")
// 	wg.Wait()
// 	fmt.Println("Final message count:", messagecount)
// }


// Adavance concurrency pattern


//what is concurrency -  a program can work on multiple tasks during the same period of time .
// what is context - it is use to control the lifetime of operations - espacially goroutines, API requests, databses queries etc. , WHY WE NEED context - 



// stop a goroutine using context create a worker that performs 10 itereations. The worer should stop immediately if the context is cancelled. cancel the context after approximately 3 seconds/

// package main 
// import (
// 	"fmt"
// 	"context"
// 	"time"
// )

// func worker(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("worker stopped")
// 			return
// 		default:
// 			fmt.Println("Worker is working..")
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go worker(ctx)
// 	time.Sleep(3*time.Second)
// 	cancel()
// 	time.Sleep(time.Second)
// }




// worker  pool - fixed number of goroutines , that process multiple tasks in a shared quese, Why do we use worker poo;; - limit Goroutines (suppose there as 1000 tasks and which is devided by three 5 workers) , and also we can reuse the workers. , it is better resource management.
// how does worker pool work ? - 1 create the job , 2. put into a queue process  3 worker take the tasks 



// package main 
// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for job := range jobs {
// 		fmt.Println("worker%d processing job %d\n", id, job,)
// 	}
// }

// func main(){
// 	jobs := make(chan int, 10)
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go worker(1, jobs, &wg)
// 	go worker(2, jobs, &wg)
// 	go worker(3, jobs, &wg)


// 	for i := 1; i <= 10; i++{
// 		jobs <- i
// 	}


// close(jobs)
// wg.Wait()
// }


//create a worker pool in go with 3 weker go routines , the program should , 1 creat a channel to send jobs , start 3 workers , send numbers from 1 to 10 as jobs . each worker should calculate the square of the received number, print the worker ID, job number adb calculated sqaure , use sync.WaitGroup  to wait for all workers. close the jobs channels after sending all jobs/


// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for job := range jobs {
// 		square := job * job
// 		fmt.Printf("Worker %d: Job %d, Square = %d\n", id, job, square)
// 	}
// }

// func main() {
// 	jobs := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(3)

// 	go worker(1, jobs, &wg)
// 	go worker(2, jobs, &wg)
// 	go worker(3, jobs, &wg)

// 	for i := 1; i <= 10; i++ {
// 		jobs <- i
// 	}
// 	close(jobs)
// 	wg.Wait()

// }



// a company need to sent 8 emails. instead of creatng a seperate goroutine for every email. create a worker pool with 2 workers each worker should: 1 recceive an email id from the job channel. 2, print htat it is sending thr email, process alll emails until the channel is closed use sunc.waitgroup to ensure all workers finish.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for emailID := range jobs {
// 		fmt.Printf("Worker %d is sending email %d\n", id, emailID)
// 	}
// }

// func main() {
// 	jobs := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go worker(1, jobs, &wg)
// 	go worker(2, jobs, &wg)
// 	for emailID := 1; emailID <= 8; emailID++ {
// 		jobs <- emailID
// 	}
// 	close(jobs)
// 	wg.Wait()
// }


