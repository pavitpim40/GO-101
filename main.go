package main

import ("fmt")


// type Book struct {
// 	Title string `json:"title"`
// 	Name string `json:"name"`
// 	Author string `json:"author"`
// 	Price float64 `json:"price,omitempty"`
// }

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
	// var point float64
	// fmt.Print("Input point: ")
	// fmt.Scanf("%f", &point);

	// if point >= 80 {
	// 	fmt.Println("A")	
	// } else if point >= 70 {
	// 	fmt.Println("B")
	// } else if point >= 60 {
	// 	fmt.Println("C")
	// } else if point >= 50 {
	// 	fmt.Println("D")
	// } else if point < 0 {
	// 	fmt.Println("Invalid point")
	// } else {
	// 	fmt.Println("F")
	// }

	// Loop
	// num := []string{"a","b","c"}
	// num = append(num, "z")
	// num2 := []string{"D","E","F"}
	// num = append(num, num2...)
	// for i, v := range num {
	// 	fmt.Println(i,v)
	// }

	// var num int64
	// fmt.Print("Enter Num")
	// fmt.Scanf("%d", &num)

	// var i int64 = 1 
	// for i <= 12{
	// 	fmt.Printf("%d * %d = %d\n", num, i, num*i)
	// 	i++
	// }
	
		// data := `{"Title" : "anime" , "Name" : "Harry Potter and the Philosopher's Stone", "Author" : "J.K. Rowling", "Price" : "10.00"}`
		// var book Book
		// data2 := Book{"","","",500}
		// data3 :=[]Book{{"","","",0},{"","","",0},{"","","",0}}
		// json.Unmarshal([]byte(data), &book)
		
		// // fmt.Println(book)
		// // fmt.Printf("%+v\n", book)
		// fmt.Println(data2)
		// fmt.Printf("%+v\n", data3)
		myFunction()
		myFunction2("Hello")
		myFunction3("test Function")
		data := myFunction3("return func3")
		fmt.Println(data)

}

func	myFunction() {
	fmt.Println("My function")
}

func	myFunction2(str string) {
	fmt.Println("My function2", str)
}

func myFunction3(str string) string {
	data := fmt.Sprintf("%s", str)
	return data
}