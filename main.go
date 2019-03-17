package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//https://techblog.xavient.com/golang-programming-for-fibonacci-numbers/
func fib(n int) []int {
	f := make([]int, n+1)
	var i int

	if n <= 0 {
		return []int{0}
	}

	if n < 2 {
		return []int{1}
	}
	/* 0th and 1st number of the series are 0 and 1*/
	f[0] = 0
	f[1] = 1

	for i = 2; i <= n; i++ {
		/* Add the previous 2 numbers in the series
		and store it */
		f[i] = f[i-1] + f[i-2]
	}

	return f
}

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "Welcome!\n")
// }
// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

func fibParm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// fmt.Fprintf(w, "A fib!! Welcome!", ps.ByName("num"))
	i, _ := strconv.Atoi(ps.ByName("num"))
	json.NewEncoder(w).Encode(fib(i))
}
func main() {
	router := httprouter.New()
	// router.GET("/api", Index)
	// router.GET("/api/hello/:name", Hello)
	router.GET("/api/fib/:num", fibParm)
	log.Fatal(http.ListenAndServe(":8080", router))
}
