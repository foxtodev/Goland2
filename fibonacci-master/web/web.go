package web

import (
	"fmt"
	"net/http"
	"strconv"
)

func Serve(fiber Fiber) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		nSlice, ok := request.URL.Query()["n"] //http://localhost:8080/?n=10&n=20

		data := fmt.Sprintf("%s\n", fiber.Name())
		if ok {
			for _, nValue := range nSlice {
				n, err := strconv.Atoi(nValue)
				if err == nil {
					data += fmt.Sprintf("Fib(%d) = %d\n", n, fiber.Fib(n))
				}
			}
		}

		_, err := writer.Write([]byte(data))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
