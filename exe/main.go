package main

import "fmt"

func main() {
	// gổuntine
	ch := make(chan int, 2) // Tạo một kênh đệm với dung lượng là 2

	ch <- 1 // Gửi giá trị 1 vào kênh
	ch <- 2 // Gửi giá trị 2 vào kênh

	x := <-ch      // Nhận giá trị từ kênh
	fmt.Println(x) // Kết quả: 1

	y := <-ch      // Nhận giá trị từ kênh
	fmt.Println(y) // Kết quả: 2

	// file, err := os.Open("file.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// fmt.Printf("đọc %d byte: %q\n", count, data[:count])
	// fmt.Println("done")
	// file.Close()
	// // interface
	// co := container{
	// 	base: base{num: 42},

	// 	str: "Hello, Go!",
	// }
	// fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// fmt.Println("also num :", co.base.num)
	// fmt.Println("describe :", co.describe())

	// type describer interface {
	// 	describe() string
	// }
	// var d describer = co
	// fmt.Println("describer:", d.describe())
	// fileLogger := FileLogger{file: createLogFile()}
	// userService := UserService{logger: fileLogger}
	// userService.CreateUser("Alice")

	// consoleLogger := ConsoleLogger{prefix: "USER-SERVICE"}
	// userService.logger = consoleLogger
	// userService.CreateUser("Bob")
	// Closure
	// counter1 := createCounter()
	// counter2 := createCounter()

	// fmt.Println(counter1()) // Đầu ra: 1
	// fmt.Println(counter1()) // Đầu ra: 2
	// fmt.Println(counter2()) // Đầu ra: 1
	// fmt.Println(counter2()) // Đầu ra: 2

	// hàm ẩn danh
	// func() {
	// fmt.Println("Inside anonymous function")
	// }()
	// f1 := func() func(v int) {
	// 	num := func(number int) {
	// 		fmt.Println("number : ", number)
	// 	}
	// 	return num
	// }

	// g := f1()
	// g(2411)
	// hàm variadio
	// var line string
	// names := []string{"Sammy", "Jessica", "Drew", "Jamie"}
	// line = join(",", names...)
	// fmt.Println(line)
	// numbers := []string{"1", "2", "3", "4"}
	// sum := sum(numbers...)
	// if sum != nil {
	// 	fmt.Println("tổng của ", join(",", numbers...), " : ", *sum)
	// }
}

// type base struct {
// 	num int
// }

// func (b base) describe() string {
// 	return fmt.Sprintf("base with num=%v", b.num)
// }

// type container struct {
// 	base
// 	str string
// }

// interface
// type Logger interface {
// 	Log(level string, message string)
// }

// type FileLogger struct {
// 	file *os.File
// }

// func (f FileLogger) Log(level, message string) {
// 	fmt.Fprintf(f.file, "[%s] %s: %s\n",
// 		time.Now().Format(time.RFC3339),
// 		level,
// 		message)
// }

// type ConsoleLogger struct {
// 	prefix string
// }

// func (c ConsoleLogger) Log(level, message string) {
// 	fmt.Printf("%s [%s] %s\n",
// 		c.prefix,
// 		level,
// 		message)
// }

// // service using the logger
// type UserService struct {
// 	logger Logger
// }

// func (s UserService) CreateUser(name string) error {
// 	s.logger.Log("INFO", fmt.Sprintf("Creating user: %s", name))
// 	return nil
// }

// closure
// func createCounter() func() int {
// 	count := 0
// 	increment := func() int {
// 		count++
// 		return count
// 	}
// 	return increment
// }

// hàm variadio
// func join(del string, values ...string) string {
// 	var line string
// 	for i, v := range values {
// 		line = line + v
// 		if i != len(values)-1 {
// 			line = line + del
// 		}
// 	}
// 	return line
// }

// func sum(numbers ...string) *float64 {
// 	var sum float64
// 	for _, v := range numbers {
// 		number, err := strconv.ParseFloat(v, 64)
// 		if err != nil {
// 			return nil
// 		}
// 		sum += number
// 	}
// 	return &sum
// }
