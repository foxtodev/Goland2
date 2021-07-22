
#### 01
###### Напишите программу, в которой неявно будет срабатывать паника. Сделайте отложенную функцию, которая будет обрабатывать эту панику и печатать предупреждение в консоль. Критерий выполнения задания - программа не завершается аварийно. Дополните программу собственной ошибкой, хранящей время ее возникновения.

```go
type myError struct {
	msg  string
	time string
}

func New(text string) error {
	return &myError{text, time.Now().Format("2006/1/2 15:04")}
}

func (e *myError) Error() string {
	return fmt.Sprintf("%s at %s", e.msg, e.time)
}

func sum(n []int) (s int, e error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Panic:", New("Error in Sum() function"), "\nOriginal error mesage:", e)
		}
	}()

	for i := 0; i <= len(n); i++ {
		s += n[i]
	}
	return s, e
}

func main() {

	n := []int{1, 2, 3, 4}
	s, _ := sum(n)
	fmt.Printf("Sum of all the elements of the array equal: %d\n", s)

	createFile("file")

}
```

```
go run main.go
There was a panic: Error in Sum() function at 2021/7/16 02:21 
Original error mesage: runtime error: index out of range [4] with length 4
Sum of all the elements of the array equal: 10
```

#### 02
###### Напишите функцию которая создает файл в файловой системе и использует отложенный вызов функций для безопасного закрытия файла

```go
func createFile(fileName string) {
	_, err := os.Stat(fileName)
	if err == nil {
		fmt.Printf("File <%s> exists\n", fileName)
		return
	}
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("File not created:", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	_, _ = fmt.Fprintln(f, "File created")
}
```

<br />
