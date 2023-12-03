package main

import (
	"fmt"
	_ "golang-dasar/database"
	_ "golang-dasar/helper"
	_ "golang-dasar/internal"
)

//* Number
// func main() {
// 	fmt.Println("satu =", 1)
// 	fmt.Println("dua =", 2)
// 	fmt.Println("tiga koma lima =", 3.5)
// }

//* boolean
// func main() {
// 	fmt.Println("benar =", true)
// 	fmt.Println("salah =", false)
// }

//* tipe data string
// func main() {
// 	fmt.Println(" ini tipe data string")
// 	fmt.Println(" ini tipe data string kedua")
// }
// func main() {
// 	fmt.Println(len("satu"))// 4
// 	fmt.Println(" ini tipe data string kedua"[0])// 32
// 	fmt.Println(" ini tipe data string kedua"[1])// 105
// }

//* Variable
// func main() {
// 	var name string

// 	name = "Variable"
// 	fmt.Println((name))

// 	name = "Variable yang baru"
// 	fmt.Println((name))
// }

// func main() {
// 	var name = "Variable"
// 	fmt.Println((name))

// 	name = "Variable yang baru"
// 	fmt.Println((name))
// }

// func main() {
// 	name := "Variable"
// 	fmt.Println((name))

// 	name = "Variable yang baru"
// 	fmt.Println((name))

// 	var (
// 		firstName ="Kwon"
// 		lastName = "yuli"
// 	)
// 	fmt.Println(firstName)
// 	fmt.Println(lastName)
// }

//* Constant
// func main() {
// 	const firstName ="kwon"
// 	const lastName = "yuli"

// 	fmt.Println(firstName)
// 	fmt.Println(lastName)
// }

//* konversi tipe data
// func main(){
// 	var nilai32 int32 = 32768
// 	var nilai64 int64 = int64(nilai32)
// 	var nilai16 int16 = int16(nilai32)

// 	fmt.Println(nilai32)// 32768
// 	fmt.Println(nilai64)// 32768
// 	fmt.Println(nilai16)// -32768

// 	var name = "kwon"
// 	var e = name[0]
// 	var eString = string(e)

// 	fmt.Println(name)
// 	fmt.Println(e)
// 	fmt.Println(eString)
// }

//* Type Declaration
// func main()  {
// 	type NoKTP string

// 	var ktpKwon NoKTP = "1234566"
// 	var contohKtp = "1234567890"
// 	fmt.Println(ktpKwon)
// 	fmt.Println(contohKtp)
// }

//* Operasi matematika
// func main() {
// 	var a = 10
// 	var b = 10
// 	var c = a + b
// 	fmt.Println(c)
// }

//* Augmented Assignments
// func main () {
// 	var i = 10
// 	i += 10 //10s
// 	fmt.Println(i)//20
// 	i += 5
// 	fmt.Println(i)//25
// }

//* Unary Operator
// func main () {
// 	var j = 1
// 	j++ // j = j + 1 // 2
// 	fmt.Println(j)
// 	j++ // j = j + 1 // 3
// 	fmt.Println(j)

// 	j-- // 2
// 	fmt.Println(j)
// 	j-- // 1
// 	fmt.Println(j)
// }

//* Operasi Perbandingan
// func main () {
// 	var name1 = "kwon"
// 	var name2 = "kwon"

// 	var result = name1 == name2
// 	var result2 = name1 != name2

// 	fmt.Println(result)// true
// 	fmt.Println(result2)// false

// 	var a = "a"
// 	var b = "b"

// 	var hasil = a < b

// 	fmt.Println(hasil)// true
// }

//* Operasi Boolean
// func main () {
// var nilaiAkhir = 90
// var absensi = 80

// var lulusNilaiAkhir = nilaiAkhir > 80
// var lulusAbsensi = absensi > 80

// var lulus = lulusNilaiAkhir && lulusAbsensi

// fmt.Println(lulus)// false

// 	var nilaiAkhir = 90
// 	var absensi = 81

// 	var lulusNilaiAkhir = nilaiAkhir > 80
// 	var lulusAbsensi = absensi > 80

// 	var lulus = lulusNilaiAkhir && lulusAbsensi

// 	fmt.Println(lulus)// true
// }

//* Tipe Data Array
// func main () {
// 	var names [3]string

// 	names[0] = "kwon"
// 	names[1] = "yuli"
// 	names[2] = "anto"

// 	fmt.Println(names[0])
// 	fmt.Println(names[1])
// 	fmt.Println(names[2])

// 	var value = [3]int{90, 80, 95}
// 	fmt.Println(value)

// 	var values = [...] int {
// 		90,
// 		80,
// 		95,
// 		100,
// 		110,
// 	}

// 	fmt.Println(values)
// 	fmt.Println(values[0])
// 	fmt.Println(values[1])
// 	fmt.Println(values[2])
// 	fmt.Println(len(values))
// }

//* Tipe Data Slice

// func main () {
// 	names := [...]string{"Kwon","yuli","anto","echo","echo1", "echo2"}

// 	slice1 := names[4:6]
// 	fmt.Println(slice1)

// 	slice2 := names[:3]
// 	fmt.Println(slice2)

// 	slice3 := names[3:]
// 	fmt.Println(slice3)

// 	var slice4 = names[:]
// 	fmt.Println(slice4)

// 	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

// 	daySlice1 := days[5:] // sabtu, minggu
// 	fmt.Println(daySlice1)

// 	daySlice1[0] = "sabtu baru"
// 	daySlice1[1] = "minggu baru"
// 	fmt.Println(daySlice1)
// 	fmt.Println(days)

// 	daySlice2 := append(daySlice1, "Libur Baru")
// 	daySlice2[0] = "Sabtu Lama"

// 	fmt.Println(daySlice1)
// 	fmt.Println(daySlice2)
// 	fmt.Println(days)

// 	var newSlice = make([]string, 2, 5)
// 	newSlice[0] = "echo1"
// 	newSlice[1] = "echo2"

// 	fmt.Println(newSlice)
// 	fmt.Println(len(newSlice))
// 	fmt.Println(cap(newSlice))

// 	// menambah array newslice
// 	newSlice2 := append(newSlice, "echo3")

// 	fmt.Println(newSlice2)
// 	fmt.Println(len(newSlice2))
// 	fmt.Println(cap(newSlice2))

// 	newSlice2[0] = "kwon"
// 	fmt.Println(newSlice2)
// 	fmt.Println(newSlice)

// 	fromSlice := days[:]
// 	toSlice := make([]string, len(fromSlice), cap(fromSlice))

// 	copy(toSlice, fromSlice)

// 	fmt.Println(fromSlice)
// 	fmt.Println(toSlice)

// 	iniArray := [...]int{1,2,3,4,5}
// 	iniSlice := []int{1,2,3,4,5}

// 	fmt.Println(iniArray)
// 	fmt.Println(iniSlice)
// }

//* Tipe data Map
// func main () {
// 	person := map[string]string{
// 		"name": "Kwon",
// 		"address" : "Guandong",
// 	}

// 	fmt.Println(person["name"])
// 	fmt.Println(person["address"])
// 	fmt.Println(person)

// 	book := make(map[string]string)
// 	book["title"] = "Buku Golang"
// 	book["author"] = "Kwon"
// 	book["ups"] = "salah"

// 	fmt.Println(book)

// 	delete(book, "ups")

// 	fmt.Println(book)
// }

//* If Expression
// func main () {
// name := "kwon"

// if name == "kwon" {
// 	fmt.Println("hello kwon")
// }

// if statemen
// name := "Kwon"// true
// name := "lee"// false
// 	name := "yuli"// true

// 	if name == "Kwon" {
// 		fmt.Println("Hello kwon")
// 	} else if name == "yuli" {
// 		fmt.Println("hello yuli")
// 	} else {
// 		fmt.Println("Hi, boleh kenalan")
// 	}

// 	if length := len(name); length > 5 {
// 		fmt.Println("nama terlalu panjang")
// 	} else {
// 		fmt.Println("Nama sudah benar")
// 	}

// }

//* Switch Expression
// func main ()  {
// 	name := "Kwon"

// 	switch name {
// 	case "Kwon":
// 		fmt.Println("Hello Kwon")
// 	case "Yuli" :
// 		fmt.Println("Hello yuli")
// 	default:
// 		fmt.Println("Hi, boleh kenalan?")
// 	}

// 	switch length := len(name); length > 5 {
// 	case true:
// 		fmt.Println("nama terlalu panjang")
// 	case false:
// 		fmt.Println("Nama sudah benar")
// 	}

// 	// sample 2
// 	name = "namasas"
// 	length := len(name)
// 	switch {
// 	case length > 10:
// 		fmt.Println("nama terlalu panjang")
// 	case length > 5:
// 		fmt.Println("nama lumayan panjang")
// 	default:
// 		fmt.Println("Nama sudah benar")
// 	}
// }

//* For Loops
// func main()  {
// counter := 1

// for counter <= 10 {
// 	fmt.Println("Perulangan ke", counter)
// 	counter++
// }

// fmt.Println("Selesai")

// cara lebih sederhana
// for counter := 1; counter <= 10; counter++ {
// 	fmt.Println("Perulangan ke", counter)
// }

// fmt.Println("Selesai")
// sample 2 cara manual
// 	names := []string{"kwon", "yuli", "anto"}
// 	for i := 0; i < len(names); i++ {
// 		fmt.Println(names[i])
// 	}

// 	// cara lebih praktis
// 	for index, name := range names {
// 		fmt.Println("Index", index, "=", name)
// 	}

// 	// tidak pakai index dan key-nya
// 	for _, name := range names {
// 		fmt.Println( name)
// 	}
// }

//* Break & Continue
// func main () {
// break
// for i := 0; i < 10; i++ {
// 	if i == 5 {
// 		break
// 	}
// 	fmt.Println("Perulangan ke", i)
// }

// continue
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Println("Perulangan ke", i)
// 	}
// }

//* Function
// func sayHello() {
// 	fmt.Println("Hello")
// }
// func main() {
// 	sayHello()
// 	sayHello()
// }

//* Function Parameter
// func sayHelloTo(firstName string, lastName string) {
// 	fmt.Println("Hello", firstName, lastName)
// }

// func main() {
// 	sayHelloTo("Kwon", "Yuli")
// 	sayHelloTo("Yuli", "Antos")
// }

//* Function Return Value
// func getHello(name string) string {
// 	hello := "Hello " + name
// 	return hello
// }

// func main() {
// 	result := getHello("kwon")
// 	fmt.Println(result)

// 	fmt.Println(getHello("echo1"))
// 	fmt.Println(getHello("echo2"))
// }

//* return multiple values
// func getFullName() (string, string) {
// 	return "echo1", "echo2"
// }

// func main () {
// firstName, lastName := getFullName()
// fmt.Println(firstName, lastName)

// 	firstName, _ := getFullName()
// 	fmt.Println(firstName)
// }

//* Named Return Values
// func getCompleteName() (firstName, middleName, lastName string) {
// 	firstName = "Kwon"
// 	middleName = "Yuli"
// 	lastName = "Anto"

// 	return firstName, middleName, lastName
// }

// func main() {
// 	a, b, c := getCompleteName()
// 	fmt.Println(a, b, c)
// }

//* Variadic Function
// func sumAll(numbers ...int) int {
// 	total := 0

// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }

// func main() {
// 	fmt.Println(sumAll(10, 10, 10))
// 	fmt.Println(sumAll(10, 10, 10, 10))
// 	fmt.Println((sumAll(10, 10, 10, 10, 10)))

// 	numbers := []int{10, 10}
// 	fmt.Println(sumAll(numbers...))
// }

//* Funtion Value
// func getGoodBye(name string) string {
// 	return "Good Bye " + name
// }

// func main() {
// 	sample1 := getGoodBye
// 	sample2 := getGoodBye

// 	fmt.Println(sample1("echo1"))
// 	fmt.Println(sample2("echo2"))
// }

//* Function as Parameter
// type Filter func(string) string
// func sayHelloWithFilter(name string, filter Filter) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello", filteredName)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// func main() {
// 	sayHelloWithFilter("echo1", spamFilter)

// 	filter := spamFilter("Anjing")
// 	sayHelloWithFilter(filter, spamFilter)
// }

//* Anonymous Function
// type Blacklist func(string) bool

// func registerUser(name string, blacklist Blacklist) {
// 	if blacklist(name) {
// 		fmt.Println("You are blocked", name)
// 	} else {
// 		fmt.Println("Welcome", name)
// 	}
// }

// func main() {
// 	blacklist := func(name string) bool {
// 		return name == "anjing"
// 	}
// 	registerUser("echo2", blacklist)

// 	registerUser("anjing", func(name string) bool {
// 		return name == "anjing"
// 	})
// }

//* Recursive Function
// func factorialLoop(value int) int {
// 	result := 1

// 	for i := value; i > 0; i-- {
// 		result *= i
// 	}

// 	return result
// }

// func factorialRecursive(value int) int {
// 	if value == 1 {
// 		return 1
// 	} else  {
// 		return value * factorialRecursive(value-1)
// 	}
// }

// func main() {
// 	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
// 	fmt.Println(result)
// 	fmt.Println(factorialLoop(10))
// 	fmt.Println(factorialRecursive(10))
// }

//* Closure
// func main() {
// 	counter := 0

// 	increment := func ()  {
// 		fmt.Println("Increment")
// 		counter++
// 	}

// 	increment()
// 	increment()
// 	increment()

// 	fmt.Println(counter)
// }

//* Defer, Panic & Recover
// defer
// func logging() {
// 	fmt.Println("Selesai memanggil function")
// }

// func runApplication() {
// 	defer logging()
// 	fmt.Println("Run application")
// }

// func main() {
// 	runApplication()
// }

// panic
// func endApp() {
// 	fmt.Println("End app")
// 	message := recover()
// 	fmt.Println("terjadi panic", message)
// }

// func runApp(error bool) {
// 	defer endApp()
// 	if error {
// 		panic("Ups Error")
// 	}
// }

// func main() {
// 	runApp(true)
// 	fmt.Println("Echo1 echo2")
// }

//* Struct
// type Customer struct {
// 	Name, Address string
// 	Age		int
// }

// func main() {
// 	var yuli Customer
// 	fmt.Println(yuli)

// 	yuli.Name = "Kwon yuli"
// 	yuli.Address = "Indonesia"
// 	yuli.Age = 30
// 	fmt.Println(yuli)
// 	fmt.Println(yuli)
// 	fmt.Println(yuli)
// 	// akses satu persatu dengan cara dibawah ini :
// 	fmt.Println(yuli.Name)
// 	fmt.Println(yuli.Address)
// 	fmt.Println(yuli.Age)

// 	kwon := Customer{
// 		Name: "Kwon",
// 		Address: "Korea",
// 		Age: 30,
// 	}
// 	fmt.Println(kwon)

// 	zhao := Customer{"Zhao", "China", 30}
// 	fmt.Println(zhao)
// }

//* Struct Method
// type Customer struct {
// 	Name, Address string
// 	Age		int
// }

// func (customer Customer) sayHello(name string) {
// 	fmt.Println("hello", name, "my name is ", customer.Name)
// }

// func main() {
// 	var yuli Customer
// 	fmt.Println(yuli)

// 	yuli.Name = "Kwon yuli"
// 	yuli.Address = "Indonesia"
// 	yuli.Age = 30
// 	fmt.Println(yuli)
// 	fmt.Println(yuli)
// 	fmt.Println(yuli)
// 	// akses satu persatu dengan cara dibawah ini :
// 	fmt.Println(yuli.Name)
// 	fmt.Println(yuli.Address)
// 	fmt.Println(yuli.Age)

// 	kwon := Customer{
// 		Name: "Kwon",
// 		Address: "Korea",
// 		Age: 30,
// 	}
// 	fmt.Println(kwon)

// 	zhao := Customer{"Zhao", "China", 30}
// 	fmt.Println(zhao)

// 	zhao.sayHello("Anto")
// 	kwon.sayHello("Anto")
// 	yuli.sayHello("Anto")
// }

//* Interface
// type HashName interface {
// 	GetName() string
// }

// func SayHello(hashName HashName) {
// 	fmt.Println("Hello", hashName.GetName())
// }

// type Person struct {
// 	Name string
// }

// func (person Person) GetName() string {
// 	return person.Name
// }

// type Animal struct {
// 	Name string
// }

// func (animal Animal) GetName () string {
// 	return animal.Name
// }
// func main () {
// 	person := Person{Name: "Echo1"}
// 	SayHello(person)

// 	animal := Animal{Name: "Kucing"}
// 	SayHello(animal)
// }

//* Interface Kosong (any)
// func Ups() any {
// return 1 //1
// return true // true (boolean)
// 	return "Ups" //Ups
// }

// func main() {
// 	var Kosong = Ups()
// 	fmt.Println(Kosong)
// }

//* Nil
// func NewMap(name string) map[string]string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return map[string]string{
// 			"name": name,
// 		}
// 	}
// }

// func main() {
// 	data := NewMap("") // data masih kosong
// data := NewMap("Ada data") // ada data

// 	if data == nil {
// 		fmt.Println("Data map masih kosong")
// 		} else {

// 			fmt.Println(data["name"])
// 	}
// }

//* Type Assertions
// func random() any {
// return "OK"
// return 1
// 	return true
// }

// func main() {
// 	var result = random()
// 	// var resultString = result.(string)
// 	// fmt.Println(resultString)
// 	// var resultInt = result.(Int)
// 	// fmt.Println(resultInt)

// 	switch value := result.(type) {
// 	case string:
// 		fmt.Println("String", value)
// 	case int:
// 		fmt.Println("Int", value)
// 	default:
// 		fmt.Println("Unknown", value)
// 	}

// }

//* Pointer
// type Address struct {
// 	City, Province, Country string
// }

// func main() {
// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
// address2 := address1 // copy value

// address2.City = "Bandung"
// fmt.Println(address1) // tidak berubah
// fmt.Println(address2) // berubah menjadi Bandung

// 	var address1 = Address{"Subang", "Jawa Barat", "Indonesia"}
// 	var address2 = &address1 // pointer

// 	address2.City = "Bandung"
// 	fmt.Println(address1) // ikut berubah
// 	fmt.Println(address2) // berubah menjadi Bandung
// }

//* Asterisk Opertor
// type Address struct {
// 	City, Province, Country string
// }

// func main() {
// 	var address1 = Address{"Subang", "Jawa Barat", "Indonesia"}
// 	var address2 = &address1 // pointer
// 	address2.City = "Bandung"
// 	fmt.Println(address1) // ikut berubah
// 	fmt.Println(address2) // berubah menjadi Bandung

// 	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
// 	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
// 	fmt.Println(address1)
// 	fmt.Println(address2)
// }

//* Operator New
// type Address struct {
// 	City, Province, Country string
// }

// func main () {
// 	var alamat1 = new(Address)
// 	var alamat2 = alamat1

// 	alamat2.Country = "Indonesia"

// 	fmt.Println(alamat1)
// 	fmt.Println(alamat2)
// }

//* Pointer di Function
// type Address struct {
// 	CIty, Province, Country string
// }

// func ChangeCountryToIndonesia(address *Address) {
// 	address.Country = "Indonesia"
// }

// func main() {
// 	var address Address = Address{}
// 	ChangeCountryToIndonesia(&address)

// 	fmt.Println(address)
// }

//* Pointer di Method
// type Man struct {
// 	Name string
// }

// func (man *Man) Married(){
// 	man.Name = "Mr. " + man.Name
// }

// func main() {
// 	echo1 := Man{"Echo1"}
// 	echo1.Married()

// 	fmt.Println(echo1.Name)
// }

//* Package dan import
// func main() {
// 	result := helper.SayHello("Echo1")
// 	fmt.Println(result)
// }

//* Access Modifier
// func main() {
// 	result := helper.SayHello("Echo1")
// 	fmt.Println(result)

// 	fmt.Println(helper.Application)
// fmt.Println(helper.version) // tidak bisa diakses
// fmt.Println(helper.sayGoodBye("Echo1")) // tidak bisa diakses
// }

//* Package Initialization
// func main() {
// 	fmt.Println(database.GetDatabase())
// }

//* Error
// func Pembagian(nilai int, pembagi int) (int, error) {
// 	if pembagi == 0 {
// 		return 0, errors.New("Pembagian Dengan NOL")
// 	} else {
// 		return nilai / pembagi, nil
// 	}
// }

// func main() {
// 	hasil, err := Pembagian(100, 0) // Error Pembagian Dengan NOL
// 	// hasil, err := Pembagian(100, 10) // Hasil 10
// 	if err == nil {
// 		fmt.Println("Hasil", hasil)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}
// }

//* Membuat Custom Error
type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, ) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "Echo1" {
		return &notFoundError{"data not found"}
	}
	// ok
	return nil
}

func main() {
	// err := SaveData("")// error
	err := SaveData("Echo1")// sukses
	if err != nil {
		// terjadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error", finalError.Error())
		default:
			fmt.Println("unknown error", finalError.Error())
		}
	} else {
		// sukses
		fmt.Println("Sukses")
	}
}