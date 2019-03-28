/*package main

import "fmt"

var p = fmt.Println

func main2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	p(sum)
}
*/

package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

type Struct struct {
	hansei   string
	cyber    string
	security string
}

func (s Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string/", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}
