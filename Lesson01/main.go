package main

import (
	"fmt"
	"os"
	"time"
)

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
			fmt.Println("There was a panic:", New("Error in Sum() function"), "\nOriginal error mesage:", e)
		}
	}()

	for i := 0; i <= len(n); i++ {
		s += n[i]
	}
	return s, e
}

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
	_, _ = fmt.Fprintln(f, time.Now().Format("2006/1/2 15:04"))
}

func main() {

	n := []int{1, 2, 3, 4}
	s, _ := sum(n)
	fmt.Printf("Sum of all the elements of the array equal: %d\n", s)

	createFile("file")

}
