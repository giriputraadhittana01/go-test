package go_test

import (
	// "container/list"
	// "container/ring"
	"errors"
	"reflect"
	// "regexp"

	// "sort"
	// "time"

	// "strconv"

	// "math"
	// "strconv"

	// "strings"

	// "flag"
	"fmt"
	// "os"
	// "strings"
	// "strconv"
	// "github.com/giriputraadhittana01/go-test/helper"
	// "github.com/giriputraadhittana01/go-test/database"
)

// func printValueAndTypeString(item string) {
// 	fmt.Printf("%v, %T\n", item, item)
// 	// fmt.Println(item);
// }
// func printValueAndTypeInt(item int) {
// 	fmt.Printf("%v, %T\n", item, item)
// 	// fmt.Println(item);
// }

// type Student struct {
// 	name    string
// 	nim     string
// 	age     int
// 	courses []string
// }

// type Animal struct {
// 	name   string `required max : "100"`
// 	origin string
// }

// type Dog struct {
// 	Animal
// 	bark_strength int
// 	claw_strength int
// }

// func sum(x int,y int) int{
// 	return x+y
// }

// // func swap(x int,y int) (int,int){
// // 	return y,x
// // }
// func swap(x int,y int) (ax, ay int){
// 	ax = y
// 	ay = x
// 	return
// }

// func addArr(curr []string,arr ...string) ([]string){
// 	// append(curr,arr...)
// 	// tmp := make([]string,len(arr),cap(arr))
// 	// copy(tmp,arr)
// 	newArr := append(curr,arr...)
// 	return newArr
// }

// func displayArr(arr ...string){
// 	fmt.Println(arr)
// }

// func getName(str string) string{
// 	return "Hello "+ str
// }

// func spamFilter(text string) string{
// 	arr := strings.Split(text," ")
// 	var res []string
// 	for i:=0; i<len(arr); i++{
// 		str := ""
// 		if arr[i]=="Anjing"{
// 			str = "..."
// 		}else{
// 			str = arr[i]
// 		}
// 		res = append(res, str)
// 	}
// 	return strings.Join(res," ")
// }

// type Filter func(string) string
// func sendText(text string, filter Filter) string{
// 	filteredText := filter(text)
// 	return filteredText
// }

// func fib(n int) int{
// 	if n<2 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }

// func logging(){
// 	fmt.Println("Logging")
// }

// func runApplication(value int){
// 	defer logging()
// 	result := 10/value
// 	fmt.Println("Result :",result)
// }

// func endApp(){
// 	fmt.Println("End App")
// }

// func runApp(err bool){
// 	defer endApp()
// 	if err {
// 		panic("Something went wrong")
// 	}
// 	fmt.Println("Start App")
// }

// func end(){
// 	message := recover()
// 	fmt.Println("Err :",message)
// 	fmt.Println("End App")
// }

// func run(err bool){
// 	defer end()
// 	if err {
// 		panic("Something went wrong")
// 	}

// 	fmt.Println("Start App")
// }

// type Customer struct{
// 	Name string
// 	Address string
// 	age string
// }

// func (customer Customer) sayHello(name string) {
// 	fmt.Println("Hello "+customer.Name,", This is :",name)
// }

////////// Interface ////////
type HasName interface{
	GetName() string
	GetAge() int
}

func greeting(hasName HasName){
	fmt.Println("Hello : "+hasName.GetName())
	fmt.Println("Age : ",hasName.GetAge())
}
///////////////////////////////

type Person struct {
	Name string `required:"true" max:"100"`
	Age int
}

type Animal struct {
	AnimalName string
	Age int
}

func (person Person) GetName() string {
	return person.Name
}
func (person Person) GetAge() int {
	return person.Age
}

func (animal Animal) GetName() string{
	return animal.AnimalName
}
func (animal Animal) GetAge() int{
	return animal.Age
}

func ups(i int) interface{}{
	if i==1{
		return 1
	}else if i==2{
		return true
	}else{
		return "Ups"
	}
}

func newMap(name string) map[string]string{
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"name" : name,
		}
	}
}

func pembagi(nilai int, pembagi int) (int,error){
	if pembagi==0{
		return 0,errors.New("Pembagi tidak boleh 0")
	}else{
		result := nilai/pembagi
		return result,nil
	}
}

func getData(n int) interface{}{
	if n==1{
		return "Hello"
	}else if(n==2){
		return 0
	}else{
		return true
	}
}


type Address struct{
	country string
}

func chgValue(n *int){
	*n = 20
}

type Man struct{
	name string
}

func (man *Man) Married(){
	man.name = "Hello "+man.name
	fmt.Println(man)
}


var boba = "Enak"

type User struct{
	Name string
	Age int
}

type UserSlice []User

func (value UserSlice) Len() int{
	return len(value)
}

func (value UserSlice) Less(i,j int) bool{
	return value[i].Age<value[j].Age
}

func (value UserSlice) Swap(i,j int){
	value[i],value[j] = value[j],value[i]
}

func isValid(data interface{}) bool{
	t := reflect.TypeOf(data)
	for i:=0;i<t.NumField();i++{
		field := t.Field(i)
		if field.Tag.Get("required") ==  "true"{
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == ""{
				return false
			}
		}
	}
	return true
}

func SayHello() string{
	return "Hello"
}

// func main() {
// 	// Declare a variable in golang
// 	// var number int = 20
// 	// number2 := 42
// 	// name := "Giri Putra Adhittana"

// 	// // Print in Golang
// 	// printValueAndTypeInt(number)
// 	// printValueAndTypeString(name)
// 	// printValueAndTypeInt(number2)

// 	// Convert golang
// 	// var i int = 42
// 	// j := float32(i)
// 	// fmt.Printf("%v %T", j, j)
// 	// j := string(i) // Not using strconv will be printing ascii
// 	// fmt.Printf("%v %T", j, j)
// 	// j := strconv.Itoa(i)
// 	// fmt.Printf("%v %T", j, j)

// 	// Declare an array
// 	// arr := [10]int{1,2,3,4,5,6,7,7,8,9}
// 	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	// fmt.Printf("Value : %v",arr)
// 	// Slice an element
// 	// a := arr[:]
// 	// b := arr[3:]
// 	// c := arr[:6]
// 	// d := arr[3:6]
// 	// arr[5] = 52
// 	// fmt.Printf("Value : %v\n", a)
// 	// fmt.Printf("Value : %v\n", b)
// 	// fmt.Printf("Value : %v\n", c)
// 	// fmt.Printf("Value : %v\n", d)

// 	// using make method
// 	// a := make([]int, 3, 100)
// 	// fmt.Println(a)
// 	// fmt.Printf("Length : %v\n", len(a))
// 	// fmt.Printf("Capacity : %v\n", cap(a))
	
// 	// Copy Slice
// 	// arr := [...]int{1,2,3,4,5,6}
// 	// a := make([]int, 3, 100)
// 	// a[0] = 120
// 	// a[1] = 200
// 	// a[2] = 300
// 	// copySlice := make([]int, len(a),cap(a))
// 	// copy(copySlice,a)
// 	// fmt.Println(a)
// 	// fmt.Println(copySlice)

// 	// Map and Struct
// 	// statePopulations := map[string]int{
// 	// 	"California": 123,
// 	// 	"Texas":      412,
// 	// 	"Florida":    751,
// 	// 	"Miami":      293,
// 	// 	"Ohio":       815,
// 	// }
// 	// statePopulations["Georgia"] = 1310
// 	// delete(statePopulations, "Texas") // Delete Entry
// 	// fmt.Println(statePopulations)
// 	// fmt.Println(statePopulations["Ohio"])

// 	// student := Student{
// 	// 	name: "Giri Putra Adhittana",
// 	// 	nim:  "2201767592",
// 	// 	age:  20,
// 	// 	courses: []string{
// 	// 		"BTC",
// 	// 		"XRP",
// 	// 		"ETH",
// 	// 	},
// 	// }
// 	// fmt.Println(student)

// 	// animal := Dog{
// 	// 	Animal: Animal{
// 	// 		name:   "Bobo",
// 	// 		origin: "Canada",
// 	// 	},
// 	// 	bark_strength: 200,
// 	// 	claw_strength: 141,
// 	// }
// 	// fmt.Println(animal)

// 	// t := reflect.TypeOf(Animal{})
// 	// field, _ := t.FieldByName("name")
// 	// fmt.Println(field.Tag)

// 	// Conditional Statement
// 	// number := 30
// 	// if true {
// 	// 	fmt.Println("Hello World")
// 	// }
// 	// if number > 50 {
// 	// 	fmt.Println("Number is bigger than 50")
// 	// } else if number > 25 {
// 	// 	fmt.Println("Number is bigger than 25")
// 	// }

// 	// number := 50
// 	// switch number {
// 	// case 50:
// 	// 	fmt.Println("50")
// 	// 	break
// 	// case 20:
// 	// 	fmt.Println("20")
// 	// 	break
// 	// }

// 	// Looping
// 	// for i := 0; i < 5; i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	// Defer
// 	// fmt.Println("Start")
// 	// defer fmt.Println("Thi was deffered")
// 	// panic("Something bad happened")
// 	// fmt.Println("End")

// 	// fmt.Println("Start")
// 	// defer func() {
// 	// 	if err := recover(); err != nil {
// 	// 		log.Println("Error : ", err)
// 	// 		panic(err)
// 	// 	}
// 	// }()
// 	// panic("Something bad happened")
// 	// fmt.Println("End")

// 	// type tipedata string
// 	// var name tipedata = "Giri Putra Adhittana"
// 	// var e byte = name[0]
// 	// fmt.Println(string(e))

// 	// var i = "Hello"
// 	// var j = 2
// 	// fmt.Println(i+j)

// 	// Map using make
// 	// var test map[string]string = make(map[string]string)
// 	// test["1"] = "Hello"
// 	// fmt.Println(test)

// 	// For Loop
// 	// arr := [...]string{
// 	// 	"Giri",
// 	// 	"Putra",
// 	// 	"Adhittana",
// 	// }
// 	// for i:=0; i<len(arr); i++ {
// 	// 	fmt.Println(arr[i])
// 	// }
// 	// for idx,value := range arr {
// 	// 	fmt.Println(idx," : ",value)
// 	// }

// 	// Append Array in Array
// 	// var arr []string 

// 	// arr = []string{
// 	// 	"1",
// 	// 	"2",
// 	// 	"3",
// 	// }

// 	// // arr = append(arr, "4")
// 	// // fmt.Println(arr)
// 	// tmp := []string{"4","5"}
// 	// arr = append(arr,tmp...)
// 	// fmt.Println(arr)

// 	// Append String
// 	// str := "Giri Putra Adhittana"
// 	// fmt.Println("Hello "+str+" Boba Enak")

// 	// Function
// 	// x := 2
// 	// y := 5
// 	// fmt.Println(sum(x,y))
// 	// x,y = swap(x,y)
// 	// fmt.Println(x,y)

// 	// arr := []string{
// 	// 	"Giri",
// 	// 	"Putra",
// 	// 	"Adhittana",
// 	// }
// 	// // arr = addArr(arr,"Hello","World","Panda")
// 	// // fmt.Println(arr)

// 	// displayArr(arr...)

// 	// Function Value
// 	// getname := getName
// 	// fmt.Println(getname("Panda"))

// 	// Function as Parameter
// 	// str := "Anjing kau"
// 	// filter := spamFilter

// 	// fmt.Println(sendText(str,filter))

// 	// Anonymous Function
// 	// userFunc := func(name string) string{
// 	// 	if name=="giri"{
// 	// 		return "Ientau"
// 	// 	}else{
// 	// 		return "Ane Bai"
// 	// 	}
// 	// }

// 	// str := "giri"

// 	// fmt.Println(sendText(str,userFunc))

// 	// Recursive Function Fibo
// 	// fmt.Println(fib(5))

// 	// Closure
// 	// count := 0
// 	// increment := func(){
// 	// 	count++
// 	// 	fmt.Println("Increment : ",count)
// 	// }
// 	// increment()
// 	// increment()
	
// 	// Defer : Mengeksekusi program di akhir
// 	// runApplication(0)

// 	// Panic : Membuat program berhenti
// 	// runApp(true)

// 	// Recover : Recover panic (disimpan didalam defer)
// 	// run(true)

// 	// Struct

// 	// Struct normal
// 	// var customer Customer
// 	// customer.Name = "Giri Putra Adhittana"
// 	// customer.Address = "Thamrin"
// 	// customer.age = "20"

// 	// fmt.Println(customer)

// 	// Struct Literal
// 	// customer := Customer{
// 	// 	Name: "Bobo",
// 	// 	Address: "Anggrek",
// 	// 	age: "20",
// 	// }
// 	// fmt.Println(customer)

// 	// Struct Function
// 	// customer := Customer{
// 	// 	Name: "Giri",
// 	// 	Address: "Thamrin",
// 	// 	age: "20",
// 	// }
// 	// customer.sayHello("Binus")

// 	// Interface
// 	// var person Person
// 	// person.Name = "Giri"
// 	// greeting(person)

// 	// var animal Animal
// 	// animal.AnimalName = "Bear"
// 	// greeting(animal)
// 	// fmt.Println(animal)

// 	// Interface Kosong
// 	// var data interface{} = ups(3)
// 	// fmt.Println(data)

// 	// Nil
// 	// var person map[string]string = newMap("Giri")
// 	// fmt.Println(person)

// 	// Error Interface
// 	// var error error = errors.New("Ups Error")
// 	// fmt.Printf(error.Error())

// 	// hasil, err := pembagi(100,0)
// 	// if err!=nil{
// 	// 	fmt.Println(err)
// 	// }else{
// 	// 	fmt.Println(hasil)
// 	// }

// 	// Type Assertion
// 	// result := getData(3)
// 	// switch result.(type) {
// 	// 	case string :
// 	// 		fmt.Println("String")
// 	// 	case int :
// 	// 		fmt.Println("Integer")
// 	// 	default:
// 	// 		fmt.Println("Unknown")
// 	// }

// 	// Pointer
// 	// x := Address{
// 	// 	"Medan",
// 	// }
// 	// y := &x 
// 	// (*y).country = "Jakarta"
// 	// fmt.Println(x)
// 	// fmt.Println(y)

// 	// x := 1
// 	// y := &x
// 	// z := &x
// 	// *y = 2
// 	// *z = 4
// 	// fmt.Println(x)
// 	// fmt.Println(*y)
// 	// fmt.Println(*z)


// 	// Pointer function
// 	// x := 10
// 	// chgValue(&x)
// 	// fmt.Println(x)

// 	// Pointer di Method
// 	// man := Man{
// 	// 	name: "Giri Putra Adhittana",
// 	// }	
// 	// man.Married()
// 	// fmt.Println(man)

// 	// Access Modifier
// 	// helper.SayHello("Giri Putra Adhittana")
// 	// helper.SayHy("Giri Putra Adhittana")

// 	// fmt.Println(helper.Name)
// 	// // fmt.Println(helper.version)
// 	// fmt.Println(boba)

// 	// Package Initialization
// 	// fmt.Println(database.GetConnection())

// 	// Package OS
// 	// args := os.Args
// 	// fmt.Println(args)

// 	// hostname,err := os.Hostname()
// 	// if err == nil{
// 	// 	fmt.Println(hostname)
// 	// }else{
// 	// 	fmt.Println(err)
// 	// }
	
// 	// GOPATH := os.Getenv("GOPATH")
// 	// fmt.Println(GOPATH)

// 	// Package Flag
// 	// go run playground.go -host=Giri -username=Putra -password=Adhittana
// 	// host := flag.String("host","host","Put Your Database Host")
// 	// username := flag.String("username","username","Put Your Database Username")
// 	// password := flag.String("password","password","Put Your Database Password")
// 	// flag.Parse()
// 	// fmt.Println(*host)
// 	// fmt.Println(*username)
// 	// fmt.Println(*password)

// 	// Package String
// 	// fmt.Println(strings.Contains("Giri Putra Adhittana","Putra"))
// 	// fmt.Println(strings.Split("Giri Putra Adhittana"," "))
// 	// fmt.Println(strings.ToLower("GIRI PUTRA ADHITTANA"))
// 	// fmt.Println(strings.ToUpper("giri putra adhittana"))

// 	// Package strconv
// 	// boolean, err := strconv.ParseBool("True")
// 	// if err==nil{
// 	// 	fmt.Println(boolean)
// 	// }else{
// 	// 	fmt.Println(err)
// 	// }

// 	// Convert from string to int
// 	// number, err := strconv.Atoi("1000")
// 	// if err == nil{
// 	// 	fmt.Println(number)
// 	// }else{
// 	// 	fmt.Println(err)
// 	// }

// 	// Package Math
// 	// fmt.Println(math.Round(1.2))
// 	// fmt.Println(math.Ceil(1.2))

// 	// Package Container / List
// 	// data := list.New()
// 	// data.PushBack(1)
// 	// data.PushBack(2)
// 	// data.PushFront(3)
// 	// fmt.Println(data.Back().Prev().Value)
// 	// fmt.Println(data.Front().Next().Value)
// 	// for temp := data.Front();temp!= nil;temp = temp.Next(){
// 	// 	fmt.Println(temp.Value)
// 	// }

// 	// Package Container / Ring
// 	// data := ring.New(5)
// 	// for i:=0; i<data.Len();i++{
// 	// 	data.Value = "Value : "+strconv.FormatInt(int64(i),10)
// 	// 	data = data.Next()
// 	// }
// 	// for i:=0; i<data.Len();i++{
// 	// 	fmt.Println(data.Value)
// 	// 	data = data.Next()
// 	// }
// 	// data.Do(func(value interface{}){
// 	// 	fmt.Println(value)
// 	// })

// 	// Package Sprt
// 	// data := []User{
// 	// 	{"Adhittana",22},
// 	// 	{"Giri",20},
// 	// 	{"Putra",21},
// 	// }
// 	// fmt.Println(data)
// 	// sort.Sort(UserSlice(data))
// 	// fmt.Println(data)

// 	// Package Time
// 	// now := time.Now()
// 	// fmt.Println(now)
// 	// fmt.Println(now.Year(),now.Month(),now.Day())

// 	// layout := "2006-01-02"
// 	// parse,_ := time.Parse(layout,"2020-10-02")
// 	// fmt.Println(parse)

// 	// Package Reflect
// 	// person := Person{"Giri Putra Adhittana",20}
// 	// pType := reflect.TypeOf(person)
// 	// fmt.Println(pType.NumField())
// 	// fmt.Println(pType.Field(0).Name)
// 	// fmt.Println(pType.Field(1).Name)
// 	// fmt.Println(pType.Field(0).Tag.Get("required"))
// 	// fmt.Println(pType.Field(0).Tag.Get("max"))
// 	// fmt.Println(pType.Field(1).Tag.Get("required"))
// 	// person.Name = ""
// 	// fmt.Println(isValid(person))

// 	// regex := regexp.MustCompile("g[a-z]r[a-z]")
// 	// fmt.Println(regex.MatchString("giri"))
// 	// fmt.Println(regex.MatchString("gir"))
// 	// fmt.Println(regex.MatchString("gara"))
// 	fmt.Println("Hello")
// }



