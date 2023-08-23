package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// func main() {

// GOTO
// 	i := 0

// Here:
// 	fmt.Println(i)
// 	i++
// 	goto Here
// result := forLoop()
// fmt.Println("Sum:", result)
// forLoop()
// }

// FOR LOOPS

// func forLoop() int {
// 	//require the type, in this case int, to make the function able to return anything. If not, nothing is actually returned, aas such you could not use the sum
// 	// in the main function, as it could not be stored anywhere. It would have to be printed.
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum = sum + i
// 	}
// fmt.Println("Sum:", sum)
// 	return sum
// }

// for loops with break and goto

// func main() {
// for i := 0; i < 10; i++ {
// 	if i > 5 {
// 		break
// 	}
// 	fmt.Println(i)
// }

// A:
// 	for j := 0; j < 5; j++ {
// 		for i := 0; i < 10; i++ {
// 			if i > 5 {
// 				break A
// 			}
// 			fmt.Println(i)

// 		}
// 	}
// }

// func main() {

// RANGES

// list := []string{"a", "b", "c", "d", "e", "f"}
// for k, v := range list {
// 	fmt.Println(k, v)
// }
// for pos, char := range "Gõ!" {
// 	fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
// }

// SWITCH

// 	result := Switch('f')
// 	println(result)
// }

// func Switch(c rune) int {

// 	switch {
// 	case '0' <= c && c <= '9':
// 		return int(c - '0')
// 	case 'a' <= c && c <= 'f':
// 		return int(c - 'a' + 10)
// 	case 'A' <= c && c <= 'F':
// 		return int(-'A' + 10)
// 	}
// 	return 0
// }

//ARRAYS THEY HAVE A SIZE AND A TYPE. CANNOT HAVE ELEMENTS ADDED, COS THE SIZE IS PART OF THE TYPE
// var arr [10]int
// arr := [...]int{420, 391, 69}
// arr[0] = 42
// arr[1] = 13
// fmt.Printf("The first element is %d\n", arr[0])

// MULTIDIMENSIONAL ARRAYS
// a := [2][2]int{{1, 2}, {3, 4}}
// fmt.Printf("Number is %d\n", a[0][1])

//SLICES SIMILAR TO ARRAY BUT I CAN GROW WHEN NEW ELEMENTS ARE ADDED. POINTER TO AN ARRAY. REFERENCE TYPES

// a := [...]int{1, 2, 3, 4, 5}
// s1 := a[2:4]
// s2 := a[1:5]
// s3 := a[:]
// s4 := a[:4]
// s5 := s2[:]
// s6 := a[2:4:5]
// fmt.Printf("Slice 5 is %d\n", s5)
// fmt.Printf("Slice 2 is %d\n", s2)
// fmt.Printf("Slice 6 is %d\n", s6)
// fmt.Printf("Slice 2 is %d\n", s3)

// var array [100]int
// slice := array[0:99]

// slice[98] = 1
// slice[99] = 2 // RUNTIME ERROR: INDEX OUT OF RANGE [99] WITH LENGTH 99. WE ARE TRYING TO ASSIGN 2 TO THE INDEX 99 OF THE SLICE
// WHICH DOES NOT EXIST COS THE SLICE INDEXES ONLY GO FROM 0 TO 98

// APPEND FUNCTION APPENDS VALUES OR SLICES TO PREVIOUS SLICES
// s0 := []int{0, 0}
// s1 := append(s0, 2)
// s2 := append(s1, 3, 5, 7)
// s3 := append(s2, s0...)

// fmt.Println(s0)
// fmt.Println(s1)
// fmt.Println(s2)
// fmt.Println(s3)

// COPY FUNCTION COPIES SLICE ELEMENTS FROM A SOURCE TO A DESTINATION AND RETURN THE NUMBER OF ELEMENTS IT COPIED. MINIMUN OF THE LENGTH OF SOURCE AND LENGTH OF OF DESTINATION

// var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
// var s = make([]int, 6)
// n1 := copy(s, a[0:])
// n2 := copy(s, s[2:])

// fmt.Println(a)
// fmt.Println(s)
// fmt.Println(n1)
// fmt.Println(n2)

// USING SLICES, WE ARE USING A TYPE, PRETTY MUCH LIEK A TYPESCRIPT TYPE THING, WHERE YOU DEFINE THE TYPES
// type MonthDays struct {
// 	Month string
// 	Days  int
// }

// func main() {
//MAPS  MAPS ARE ACTUALLY NOT ORGANIZED. SO IF YOU LOOP THROUGH ONE, OYU WONT ACTUALLY GET THE SAME ORDER EVERYTIME

// monthdays := map[string]int{
// 	"Jan": 31, "Feb": 28, "Mar": 31,
// 	"Apr": 30, "May": 31, "Jun": 30,
// 	"Jul": 31, "Aug": 31, "Sep": 30,
// 	"Oct": 31, "Nov": 30, "Dec": 31,
// }

// fmt.Printf("%d\n", monthdays["Dec"])

// monthdays["Beans"] = 30 // ADDING THINGS TO THE MAP

// COMMA-OK FORM/IDIOM. TO CHECK IF SOMETHING EXISTS IN A MAP. IT CHECKS IF A VALUE IS RETURNED FROM THE KEY YOU PUT IN, IN THIS CASE "BEANS", IF IT DOES, OK WILL BE TRUE, IF NOT WILL BE FALSE

// if _, ok := monthdays["Beans"]; ok {
// 	fmt.Printf("It is %v\n", ok)
// } else {
// 	fmt.Println("Bagaço")
// }

// DOING THE SAME THING AS ABOVE BUT DIFFERENT WAY

// _, ok := monthdays["Beans"]
// if ok {
// 	fmt.Printf("It is %v\n", ok)
// } else {
// 	fmt.Printf("BAGAÇO\n")
// }

// delete(monthdays, "Beans") // DELETING THINGS FROM THE MAP

// daysInYear := 0
// for month, days := range monthdays {
// 	fmt.Printf("Month: %s, Days: %d\n", month, days)

// }

// for _, days := range monthdays {
// 	daysInYear += days
// }
// fmt.Printf("Numbers of days in a year: %d\n", daysInYear)

// SINCE MAPS ARE NOT ORGANIZED, IF YOU WANT THE KEYS IN A SPECIFIC ORDER, YOU CAN ACHIEVE ESSENTIALLY THE SAME THING HERE BUT USING A SLICE INSTEAD OF A MAP TO STORE THE MONTH/DAYS KEY:VALUE PAIRS

// monthdays := []MonthDays{
// 	{"Jan", 31}, {"Feb", 28}, {"Mar", 31},
// 	{"Apr", 30}, {"May", 31}, {"Jun", 30},
// 	{"Jul", 31}, {"Aug", 31}, {"Sep", 30},
// 	{"Oct", 31}, {"Nov", 30}, {"Dec", 31},
// }

// ADD A NEW MONTH FOR THE MEMES
// newMonth := MonthDays{"Beans", 30}
// monthdays = append(monthdays, newMonth)

// DELETE A MONTH. USING SLICES; YOU HAVE TO LOOP THE THE SLICE AND DELETE THE MONTH THAT MATCHES WITH THE ONE YOU WANT, WHICH IS STORED IN A VARIABLE
// monthToDelete := "Beans"
// for i, entry := range monthdays {
// 	if entry.Month == monthToDelete {
// 		monthdays = append(monthdays[:i], monthdays[i+1:]...)
// 	}
// }

// daysInYear := 0
// for _, things := range monthdays {
// 	daysInYear += things.Days
// 	fmt.Printf("Month: %s, Days: %d\n", things.Month, things.Days)
// }

// fmt.Printf("Numbers of days in a year: %d\n", daysInYear)

// EXERCISES

// FOR LOOPS

// 1. FOR

// for i := 0; i < 10; i++ {
// 	fmt.Printf("%d", i)
// }

// 2. GOTO KEYWORD

// 	i := 0
// Loop:
// 	if i < 10 {
// 		fmt.Printf("%d", i)
// 		i++
// 		goto Loop
// 	}

// 3. FOR. FILL ARRAY

// arr := [10]int{}

// for i := 0; i < 10; i++ {
// 	arr[i] = i
// }
// fmt.Printf("%v", arr)

// AVERAGE USING IF STATEMENTS AND FOR LOOPS

// xs := []float64{4, 5, 10, 21, 17, 5, 13}
// sum := 0.0
// avg := 0.0

// if len(xs) != 0 {
// 	for i := 0; i < len(xs); i++ {
// 		sum += xs[i]
// 	}
// 	avg = sum / float64(len(xs))
// }
// fmt.Printf("Sum: %f, Average: %.2f\n", sum, avg)

// sum := 0.0
// avg := 0.0
// switch len(xs) {
// case 0:
// 	avg = 0
// default:
// 	for _, v := range xs {
// 		sum += v
// 	}
// 	avg = sum / float64(len(xs))
// }
// fmt.Printf("Sum: %f, Average: %.2f\n", sum, avg)

// func main() {
// FIZZ BUZZ

// for i := 1; i < 100; i++ {
// 	p := false
// 	if i%3 == 0 {
// 		fmt.Printf("Fizz")
// 		p = true
// 	}
// 	if i%5 == 0 {
// 		fmt.Printf("Buzz")
// 		p = true
// 	}
// 	if !p {
// 		fmt.Printf("%v", i)
// 	}
// 	fmt.Println()
// }

// FUNCTIONS

// rec(0)
// }

// RECURSION - PRINTS 9 to i, in decreasing order. It everytiem the function calls itself, it creates a stack, from i to 9, which then when base case in reached,
// runs the code after teh recursion, from the first on the stack, to the last. As a stack normally works.

// func rec(i int) {
// 	if i == 10 {
// 		return
// 	}
// 	rec(i + 1)
// 	fmt.Printf("%d", i)
// }

// SCOPE - VARIABLES OUTSIDE OF FUNCTIONS ARE GLOBAL. VARIABLES DEFINED INSIDE FUNCTIONS ARE LOCAL. IN THE EXAMPLE BELOW,  IT RETURNS 5, 6 and then 5. BECAUSE MAIN, WHERE a IS 5,
// PRINTS IT AS 5, NATURALLY, THE CALLS f(). INSIDE F, a IS NOW 6. SO PRINT(a), PRINTS 6. WHICH THEN CALLS g(). WHICH ONLY PRINTS a. SINCE a:= 6 WAS ONLY DEFINED INSIDE f, INSIDE THE SCOPE OF G, a IS 5 AND NOT 6-

// var a int

// func main() {
// a = 5
// print(a)
// f()
// }

// func f() {
// 	a := 6
// 	print(a)
// 	g()
// }

// func g() {
// 	print(a)
// }

// FUNCTIONS AS VALUES

// A MAP OF FUNCTIONS WITH INTS AS KEYS AND VALUES AS FUNCTION THAT RETURN AN INT. AS SEEN BELOW, IF xs[3]() IS CALLED, IT RETURN 30.
// var xs = map[int]func() int{
// 	1: func() int { return 10 },
// 	2: func() int { return 20 },
// 	3: func() int { return 30 },
// }

// func main() {

// fmt.Println(xs[1]())
// fmt.Println(xs[2]())
// fmt.Println(xs[3]())

// ANONYMOUS FUNCTION ASSIGNED TO A

// a:= func() {
// 	fmt.Println("Hello")
// }
// a()

// CALL BACK - The key concept here is that functions can be passed as arguments to other functions, allowing for flexible and reusable code patterns.
// PRINTIT FUNCTION, TAKES AN IT, AND THEN CALLBACK FUNCTION, WHICH TAKES AND INT AND A FUNCTION AS ARGUMENTS AND RUN THAT FUNCTION WITH THE INT AS AN ARGUMENT.
// IN THIS CALL, I CALLED CALLBACK(20, printit) WHICH RUNS THE FUNCTION PRINTIT WITH THE INT 20. PRINTING 20

// callback(20, printit)

// }

// func printit(x int) {
// 	fmt.Printf("%v\n", x)
// }

// func callback(y int, f func(int)) {
// 	f(y)
// }

// DEFERRED CODE

// func ReadWrite() bool {
// 	file.Open("file")
// 	//DO YOUR THING
// 	if failureX {
// 		file.Close()
// 		return false
// 	}

// 	if failureY {
// 		file.Close()
// 		return false
// 	}
// 	file.Close()
// 	return true
// }

// WRITING THIS CODE THE WAY ABOVE, REPEATS ALOT OF CODE. GO AS THE KEYWORD DEFER, WHICH LETS US SIMPLIFY THE CODE LIKE BELOW
// THIS PUTS CLOSE RIGHT ON TOP OF THE FUNCTION AND MAKES IT EASIER TO READ, COS LESS REPEATING CODE. CLOSE IS NOW AUTOMATICALLY DONE .

// func ReadWrite() bool {
// 	file.Open("filename")
// 	defer file.Close()
// 	//DO YOUR THING
// 	if failureX {
// 		return false
// 	}

// 	if failureY {
// 		return false
// 	}
// 	return true
// }

// THIS DOESNT PRINT 0,1,2,3,4. BECAUSE THE PRINT FUNCTION IS DEFERRED, IT IS PUT TO THE SIDE, IN A STACK, UNTIL ALL TEH LOGIC IS COMPLETE.
// SO IT ACTUALLY PRINTS OUT 4,3,2,10

// for i:= 0; i < 5; i++ {
// 	defer fmt.Printf("%d\n", i)
// }

// func main() {
// 	a := f()
// 	fmt.Printf("%d\n", a)
// }
// func f() (ret int) {
// 	defer func() {
// 		ret++
// 	}()
// 	return 0
// }

// VARIADIC PAREMETER
// func main() {
// 	arg := [...]int{20, 21, 23, 2, 4, 6}

// 	for i, n := range arg {
// 		fmt.Printf("And the number is: %d\n", n)
// 	}
// }

// PANIC AND RECOVER

// func Panic(f func()) (b bool) {
// 	defer func() {
// 		if x := recover(); x != nil {
// 			b = true
// 		}
// 	}()
// 	f()
// 	return
// }

// func panicy() {
// 	var a [4]int
// 	a[3] = 5
// }
// func main() {
// 	fmt.Println(Panic(panicy))
// }

//-------------------------------------------------------------------------------------------------------------------------------------

// EXERCISES

// 1. AVERAGE - WRITE A FUNCTION THAT CALCULATES THE AVERAGE OF A FLOAT64 slice

// func average(slice []float64) (avg float64) {
// 	sum := 0.0
// 	avg = 0.0
// 	if len(slice) == 0 {
// 		return avg
// 	}
// 	for i := 0; i < len(slice); i++ {
// 		sum += slice[i]
// 	}
// ANOTHER WAY OF WRITING THE FOR LOOP
// for _, v := range slice {
// 	sum += v
// }
// 	avg = sum / float64(len(slice))
// 	return
// }

// func main() {

// 	slice := []float64{32, 10, 15, 24, 42}
// 	fmt.Println("Average:", average(slice))

// }

// 2. BUBBLE SORT ALGORITHM

// 20, 3, 43, 27, 5, 39, 23, 54, 12 initial state (before loop starts)

// 3, 20, 43, 27, 5, 39, 23, 54, 12 i=1

// 3, 20, 43, 27, 5, 39, 23, 54, 12 i=2

// 3, 20, 27, 43, 5, 39, 23, 54, 12 i=3

// 3, 20, 27, 5, 43, 39, 23, 54, 12 i=4

// 3, 20, 27, 5, 39, 43, 23, 54, 12 i=5

// 3, 20, 27, 5, 39, 23, 43, 54, 12 i=6

// 3, 20, 27, 5, 39, 23, 43, 54, 12 i=7

// 3, 20, 27, 5, 39, 23, 43, 54, 12 i=8

// 3, 20, 27, 5, 39, 23, 43, 12, 54 end of first loop

// n -- length of the slice is now 7, instead of 8. so it doesnt reach 54,
// because it's already the max value of the list at the end of the list. 54 is sorted.

// the inner looped reaches the end, goes to the outer loop, Because there were swaps, swapped turns to true.
// since the outer loop only runs if swapped is true, it runs again. It turns swapped to false.
// The inner loop run again.
// And again, and again.

// Until there are no swaps. THe entire list is sorted. Swapped doesnt turn to true ever.
// Which makes the outer loop not run again.

// func bubbleSort(slice []int) {
// 	n := len(slice)
// 	swapped := true

// 	for swapped {
// 		swapped = false
// 		for i := 1; i < n; i++ {
// 			if slice[i-1] > slice[i] {
// 				slice[i-1], slice[i] = slice[i], slice[i-1]

// 				swapped = true
// 			}
// 		}
// 		n--
// 	}
// }

// func main() {
// 	slice := []int{20, 3, 43, 27, 5, 39, 23, 54, 12}
// 	fmt.Printf("Unsorted: %v\n", slice)
// 	bubbleSort(slice)
// 	fmt.Printf("Sorted: %v\n", slice)
// }

// 3. FOR LOOPS - Take what you did in exercise to write the for loop and extend it a bit. Put the body of the for loop - the fmt.Printf - in a separate function.

// func main() {
// 	for i := 0; i < 10; i++ {
// 		show(i)
// 	}
// }

// func show(j int) {
// 	fmt.Printf("%d\n", j)
// }

// 4. FIBONACCI SEQUENCE

// func main() {

// 	var n int
// 	fmt.Println(n)
// n := 42 // Change this to the desired number of terms
// result := fibonacci(n)

// // Join the sequence elements using a comma and space. Just to make the output cleaner
// // Sprint, for instance, turns [1, 1, 2, 3, 5] into "1 1 2 3 5"
// // strings.Fields turns "1 1 2 3 5" into ["1", "1", "2", "3", "5"]
// // strings.Join turns ["1", "1", "2", "3", "5"] into ["1, 1, 2, 3, 5"]
// // strings.Trim turns ["1, 1, 2, 3, 5"] into "1, 1, 2, 3, 5"
// sequenceStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), ", "), "[]")

// fmt.Printf("First %d terms of the Fibonacci sequence: %s\n", n, sequenceStr)
// }

// func fibonacci(x int) []int {
// 	if x <= 0 {
// 		return nil
// 	} else if x == 1 {
// 		return []int{1}
// 	} else if x == 2 {
// 		return []int{1, 1}
// 	}

// 	sequence := []int{1, 1}
// 	// another way to create the same thing
// 	// sequence := make([]int, x)
// 	// sequence[0] = 1
// 	// sequence[1] = 1

// 	for i := 2; i < x; i++ {
// 		sequence[i] = sequence[i-1] + sequence[i-2]
// 	}

// 	return sequence

// }

// 5.VAR ARGUMENTS

// func main() {
// 	print(1, 4, 5, 6, 7)
// 	print(1, 4, 5, 6, 7)
// }

// func print(numbers ...int) {
// 	for _, d := range numbers {
// 		fmt.Printf("%d\n", d)
// 	}
// }

// 5.2JUST ANOTHER EXERCISE WITH VAR ARGUMENTS

// func sum(numbers ...int) int {
// 	total := 0
// 	for _, num := range numbers {
// 		total += num
// 	}
// 	return total
// }

// func main() {
// 	result1 := sum(1, 2, 3)
// 	result2 := sum(4, 5, 6, 7, 8)
// 	fmt.Println("Sum 1: ", result1)
// 	fmt.Println("Sum 2: ", result2)
// }

// 6. FUNCTIONS THAT RETURN FUNCTIONS

// 6.1. Write a function that returns a function that performs a +2
// on integers. Name the function plusTwo. You should then be able do the following:

// func main() {
// 	fmt.Printf("%v\n", plusTwo()(2))
// }

// func plusTwo() func(int) int {
// 	return func(x int) int { return x + 2 }
// }

// 6.2. USING CLOSURE

// func plusX(x int) func(int) int {
// 	return func(y int) int { return x + y }
// }

// func main() {
// 	fmt.Println("Beans", plusX(3)(7))

// }

// 7. MAXIMUM

// func maximum(numbers ...int) int {
// 	max := numbers[0]
// 	for _, v := range numbers {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	return max
// }

// func main() {
// 	fmt.Printf("Max is %d\n", maximum(2, 4, 6, 10, 21, 7, 23))
// }

// 8.

// MAP FUNCTION - BASICALLY HAVE A FUNCTION THAT TAKES a fUNCTION AND A SLICE AS AN ARGUMENT AND RETURN A slICE. YOU LOOP THROUGH THE SLICE AND APPLY THAT FUNCTION:
// IN THIS CASE  IT SQUARES THE ELEMENT, TO EACH ELEMENT OF THE SLICE

// func Map(f func(int) int, l []int) []int {
// 	j := make([]int, len(l))
// 	for k, v := range l {
// 		j[k] = f(v)
// 	}
// 	return j
// }

// func main() {
// 	m := []int{1, 3, 4}
// 	f := func(i int) int {
// 		return i * i
// 	}
// 	fmt.Printf("%v", (Map(f, m)))
// }

// 9.
// STACK - * is a pointer

// type stack struct {
// 	i    int
// 	data [10]int
// }

// func (s *stack) push(k int) {
// 	s.data[s.i] = k
// 	s.i++
// }

// func (s *stack) pop() int {
// 	s.i--
// 	ret := s.data[s.i]
// 	s.data[s.i] = 0
// 	return ret
// }

// func main() {
// 	var s stack
// 	s.push(25)
// 	s.push(14)
// 	s.pop()
// 	fmt.Printf("stack %v\n", s)
// }

// func (s stack) String() string {
// 	var str string
// 	for i := 0; i <= s.i; i++ {
// 		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
// 	}
// 	return str
// }

// STACK - CHATGPT
// type Stack struct {
// 	data    []int
// 	maxSize int
// 	size    int
// }

// func NewStack(maxSize int) *Stack {
// 	return &Stack{
// 		data:    make([]int, maxSize),
// 		maxSize: maxSize,
// 		size:    0,
// 	}
// }

// func (s *Stack) Push(value int) error {
// 	if s.size >= s.maxSize {
// 		return fmt.Errorf("stack is full")
// 	}
// 	s.data[s.size] = value
// 	s.size++
// 	return nil
// }

// func (s *Stack) Pop() (int, error) {
// 	if s.size == 0 {
// 		return 0, fmt.Errorf("stack is empty")
// 	}
// 	value := s.data[s.size-1]
// 	s.size--
// 	return value, nil
// }

// func (s *Stack) String() string {
// 	elements := make([]string, s.size)
// 	for i := 0; i < s.size; i++ {
// 		elements[i] = fmt.Sprintf("%d", s.data[i])
// 	}
// 	return "[" + strings.Join(elements, " ") + "]"
// }

// func main() {
// 	stack := NewStack(5)

// 	// Push elements onto the stack
// 	stack.Push(1)
// 	fmt.Println("Pushed 1:", stack)
// 	stack.Push(2)
// 	fmt.Println("Pushed 2:", stack)
// 	stack.Push(3)
// 	fmt.Println("Pushed 3:", stack)

// 	// Pop elements from the stack
// 	for i := 0; i < 3; i++ {
// 		value, err := stack.Pop()
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 		} else {
// 			fmt.Println("Popped:", value)
// 			fmt.Println("Current Stack:", stack)
// 		}
// 	}
// }

// CHAPTER 4 - PACKAGES

// func main() {
// 	i := 5
// 	fmt.Printf("Is %d even ? %v\n", i, even.Even(i))
// }

// CHAPTER 5. BEYOND THE BASICS

// POINTERS

// var p *int

// func main() {

// 	fmt.Printf("%v\n", p)
// 	var i int
// 	p = &i
// 	fmt.Printf("%v", p)
// }

// ALLOCATION - NEW AND MAKE - MAKE ONLY APPLIES TO MAPS, SLICES AND CHANNELS. DOES NOT RETURN A POINTER.
// TO OBTAIN AN EXPLICIT POINT ALLOCATE WITH NEW
// NEW ALLOCATES, MAKE INITIALIZES

// - NEW(T) RETURNS *T POINTING TO A ZEROED T
// - MAKE(T) RETURNS AND INITIALIZED T

// var p *[]int = new([]int)
// var v []int = make([]int, 100)

// var p *[]int = new([]int) // COMPLEX FOR THE SAEK OF COMPLEXITY. NEXT SYNTAX IS UST BETTER AND SIMPLER
// *p = make([]int, 100, 100) // CREATES AND ARRAY WITH 100 INTS. THEN CREATES A SLICE WITH LENGTH 10, CAPACITY 100. POINTING AT THE FIRST 10 ELEMENTS OF THE ARRAY

// v := make([]int, 100) // NEW SLICE OF 100 INTS

// CONSTRUCTORS AND COMPOSITE LITERALS

// func NewFile(fd int, name string) *File {
// 	if fd < 0 {
// 		return nil
// 	}
// 	f := new(File)
// 	f.fd = fd
// 	f.name = name
// 	f.dirinfo = nil
// 	f.nepipe = 0
// 	return f
// }

// ALOT OF BOILER PLATE. IT CAN bE SIMPLIFIED USING A COMPOSITE LITERAL. IT CREATES A NEW INSTANCE EACH TIME IT IS EVALUATED

// func NewFile(fd int, name string) *File {
// 	if fd < 0 {
// 		return nil
// 	}
// 	return &File{fd, name, nil, 0} // THE ITEMS CALLED FIELDS, MUST BE LAID OUT IN ORDER AND ALL BE PRESENT

// 	// HOWEVER, IF WE USE FIELDS:VALUE PAIRS, INITIALIZERS CAN APPEAR IN ANY ORDER, AND MISSING ONES SIMPLY GET THEIR VALUE AS ZEROo
// 	return &File{fd; fd, name: name}
// }

// COMPOSITE LITERALS CAN ALSO BE USE FOR ARRAYS, MAPS, SLICES. FIELDS LABELS ARE INDICES OR MAP KEYS.

// ar := [...]string{Enone: "no error", Einval: "invalid argument"}
// sl := []string {Enone: "no error", Einval: "invalid argument"}
// ma .= map[int]string{Enone: "no error", Einval: "invalid argument"}

//-----------------------------------------------------------------

// DEFINING YOUR OWN TYPES - TYPE KEYWORD

// FOR INSTANCE, type foo int , CREATES A NEW TPYE NAMED FOO WHICH ACTS LIKE AN INT
// MORE COMPLEX/SophisTICATED TYPES IS DONE WITH STRUCT KEYWORD

// type NameAge struct {
// 	name string
// 	age  int
// }

// func main() {
// 	a := new(NameAge)
// 	a.name = "Pete"
// 	a.age = 42
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%s\n", a.name)
// 	fmt.Printf("%d\n", a.age)
// }

// EACH ITEM IN A STRUCT IS CALLED A FILED

// struct {
// 	x,y int
// 	A *[]int
// 	F func()
// }

// IF YOU OMIT THE NAME, YOU CREATE AND ANONYMOUS FIELD
// FIELDS THAT START WITH A CAPITAL LETTER ARE EXPORTED ( CAN BE READ FROM OTHER PACKAGES, JUST LIKE FUNCTIONS DEFINED IN PACKAGES)
// WHILE LOWER CASE ARE PRIVATE

// METHODS

//1. FUNCTION THAT TAKES THE TYPE YOU CREATED AS AN ARGUMENT

// func doSomething(n1 *NameAge, n2 int) {}

//2. WORKS ON THE TYPE

// func (n1 *NameAge) doSomething(n2 int) {}

// METHOD CALL

// var n *NameAge
// n.doSomething(2)

// There is a subtle but major difference between the following type declarations.

// type Mutex struct{}

// func (m *Mutex) Lock()   {}
// func (m *Mutex) Unlock() {}

// type NewMutex Mutex
// type PrintableMutex struct{ Mutex }

//NEWMUTEX IS EQUAL TO MUTEX BUT DOES NOT HAVE ITS METHODS ( LOCK AND UNLOCK)
// ON THE OTHER HAND, PRINTABLEMUTEX HAS INHERITED THE METHODS FROM MUTEX

// -------------------------------------------------------

// CONVERSIONS - CONVERTING A TYPE TO ANOTHER

// NOT EVERY CONVERSION IS ALLOWED

// From	| b []byte  |	i []int	| r []rune  |	s string |	f float32 | i int
// To      |           |           |           |            |            |
// ________|___________|___________|___________|____________|____________|_______________
// []byte	|   ·		| []byte(s)	|	        |            |            |
// ________|___________|___________|___________|____________|____________|__________
// []int	|	·		| []int(s)	|	        |            |            |
// ________|___________|___________|___________|____________|____________|___________
// []rune	|			| []rune(s)	|	        |            |            |
// ________|___________|___________|___________|____________|____________|___________
// string	| string(b)	| string(i)	|string(r)  |	  ·	     |	          |
// ________|___________|___________|___________|____________|____________|___________
// float32	|			|	·	    |float32(i) |            |            |
// ________|___________|___________|___________|____________|____________|__________
// int		|			| int(f)	|   ·       |            |            |

// TABLE OF VALID CONVERSIONS. FLOAT64 WORKS SAME AS FLOAT32

// FROM STRING TO A SLICE OF BYTES

// func main() {
// 	// mystring := "hello this is string"
// 	// byteslice := []byte(mystring)

// 	// // STRING TO SLICE OF RUNES

// 	// runeslice := []rune(mystring)

// 	// // STRINGS IN GO ARE ENCODED IN UTF-8. SOME CHARACTERS IN THE MAY END UP in 1,2 ,3 or 4 bytes
// 	// fmt.Printf("Beans: %d\n", byteslice)

// 	// // EACH RUNE CONTAINS A UNICODE CODE POINT.
// 	// // EVERY CHRACTER CORRESPONDS TO ONE RUNE

// 	// fmt.Printf("RUNEBEANS: %d", runeslice)

// 	// FROM A SLICE OF BYTES TO A STRING

// 	b := []byte{'h', 'e', 'l', 'l', 'o'} // COMPOSITE LITERAL
// 	s := string(b)
// 	fmt.Printf("BYTETOSTRING: %s\n", s)

// 	i := []rune{257, 1024, 65}
// 	r := string(i)
// 	fmt.Printf("RUNESTOSTRING: %s\n", r)
// }

// FOR NUMERICS VALUES THE FOLLOWING CONVERSIONS ARE DEFINED
// Convert to an integer with a specific (bit) length: uint8(int)
//From floating point to an integer value: int(float32). This discards the fraction part from the floating point value.
//And the other way around: float32(int).

// USER DEFINED TYPES AND CONVERSIONS - HOW TO CONVERT BETWEEN TYPES YOU HAVE DEFINED YOURSELF

// type foo struct{ int } // ANONYMOUS STRUCT FIELD
// type bar foo           // BAR IS AN ALIAS OF FOO

// var b bar = bar{1} // DECLARE B TO BE A BAR
// var f foo = foo(b) // ASSIGN B TO F

// ---------------------------------------

// EXERCISES

// 1. MAP FUNCTION WITH INTERFACES

// BEFORE INTERFACE

// func Map(f func(int) int, l []int) []int {
// 	j := make([]int, len(l))
// 	for k, v := range l {
// 		j[k] = f(v)
// 	}
// 	return j
// }

// func main() {
// 	m := []int{1, 3, 4}
// 	f := func(i int) int {
// 		return i * i
// 	}
// 	fmt.Printf("%v", (Map(f, m)))
// }

// WITH INTERFACE - WORKING WITH INT AND STRING

type beans interface{}

// func beanDealer(input beans) beans {

// 	switch input.(type) {
// 	case int:
// 		return input.(int) * input.(int)
// 	case string:
// 		return input.(string) + input.(string) + input.(string) + input.(string)
// 	}
// 	return input
// }

// func Map(f func(beans) beans, l []beans) []beans {
// 	j := make([]beans, len(l))
// 	for k, v := range l {
// 		j[k] = f(v)
// 	}
// 	return j
// }

// func main() {
// 	m := []beans{1, 3, 4}
// 	s := []beans{"a", "b", "c", "d"}

// 	mf := Map(beanDealer, m)
// 	sf := Map(beanDealer, s)

// 	fmt.Printf("%v\n", mf)
// 	fmt.Printf("%v\n", sf)
// }

// POINTERS

// type Person struct {
// 	name string
// 	age  int
// }

// 1.WHAT IS THE DIFFERENCE BETWEEN THE 2 LINES

// var p1 Person
// p2 := new(Person)

// ANSWER:
// FIRST ONE ASSIGNS THE TYPE PERSON TO P1. THE TYPE OF P1 IS PERSON.
// SECOND ONE ALLOCATES MEMORY AND ASSINGS A POINTER TO P2. THE TYPE OF P2 IS *PERSON

// 2. WHAT IS THE DIFFERENCE BETWEEN TEH FOLLOWING TWO ALLOCATIONS

// func set(t *T) {
// 	x = t
// }

// func Set(t T) {
// 	x = &t
// }

//ANSWER:
// FIRST FUNTION, X POINTS TO THE SAME THING T DOES. WHICH IS THE SAME THING THE ARGUMENT POINTS TO.
// SECOND FUNCTION, X POINTS TO A NEW(HEAP-ALLOCATED) VARiABLE T WHICH CONTAINS A COPY OF WHATEVER THE ARGUMENT VALUE IS

// LINKED LIST - DOUBLY
// 1. USING CONTAINER/LIST PACkAGE

// func main() {
// 	l := list.New()
// 	l.PushBack(1)
// 	l.PushBack(2)
// 	l.PushBack(4)

// 	for e := l.Front(); e != nil; e = e.Next() {
// 		fmt.Printf("%v\n", e.Value)
// 	}

// }

// 2. DOUBLY LINKED LIST - RAW - NO pACKAGES

// type Node struct {
// 	prev  *Node
// 	next  *Node
// 	value interface{}
// }

// type DoublyLinkedList struct {
// 	head *Node
// 	tail *Node
// }

// // PUSHBACK ADDS A NEW NODe WITH GIVEN VALUE TO THE END OF HTE LIST
// func (list *DoublyLinkedList) PushBack(value interface{}) {
// 	newNode := &Node{
// 		prev:  list.tail,
// 		next:  nil,
// 		value: value,
// 	}

// 	if list.head == nil {
// 		list.head = newNode
// 	} else {
// 		list.tail.next = newNode
// 	}

// 	list.tail = newNode
// }

// // PRINT THE CONTENTS OF THE LSIT

// func (list *DoublyLinkedList) Print() {
// 	fmt.Println("Doubly Linked list:")
// 	current := list.head
// 	for current != nil {
// 		fmt.Println(current.value)
// 		current = current.next
// 	}
// }

// func main() {
// 	// CREATE NEW DOUBLY LINKED LSIT
// 	l := DoublyLinkedList{}

// 	// PUSH VALUES INTO THE LSIT
// 	l.PushBack(1)
// 	l.PushBack(2)
// 	l.PushBack(4)

// 	l.Print()
// }

var numberFlag = flag.Bool("n", false, "number each line") // <<2>>

func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d  %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
