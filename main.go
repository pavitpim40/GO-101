package main

import ("fmt")

func main(){
// var num1,num2 int64
// // var num3,num4 float64
// // var str string
// // str2 := ""
// num2 = 10000
// var isActive bool
// active := true
// // fmt.Print("Hello go ", isActive, " ",active, num1, num2)
// fmt.Printf("Hello go %t %t %d %d", isActive,active,num1, num2)

// Accept Input
// fmt.Print("Input number");
// var num3 float64
// fmt.Scanf("%f", &num3)
// fmt.Printf("%.2f", num3)

// Condition 

	var point float64
	fmt.Print("Input point: ")
	fmt.Scanf("%f", &point);

	if point >= 80 {
		fmt.Println("A")	
	} else if point >= 70 {
		fmt.Println("B")
	} else if point >= 60 {
		fmt.Println("C")
	} else if point >= 50 {
		fmt.Println("D")
	} else if point < 0 {
		fmt.Println("Invalid point")
	} else {
		fmt.Println("F")
	}
}