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

package main

import "fmt"

func main() {
	var name string
	var mark1, mark2, mark3, mark4, mark5 float64

	fmt.Print("Enter Student Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter marks for Subject 1: ")
	fmt.Scan(&mark1)

	fmt.Print("Enter marks for Subject 2: ")
	fmt.Scan(&mark2)

	fmt.Print("Enter marks for Subject 3: ")
	fmt.Scan(&mark3)

	fmt.Print("Enter marks for Subject 4: ")
	fmt.Scan(&mark4)

	fmt.Print("Enter marks for Subject 5: ")
	fmt.Scan(&mark5)

	total := mark1 + mark2 + mark3 + mark4 + mark5
	percentage := total / 5

	fmt.Println("\n----- Student Result -----")
	fmt.Println("Student Name:", name)
	fmt.Println("Total Marks:", total)
	fmt.Println("Percentage:", percentage, "%")
}