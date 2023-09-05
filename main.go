package main

import (
	"fmt"
	"log"
	"net/http"
)

func IsEmptyFact(fact *string) int {
	if *fact == "" {
		return 0
	}
	return 1
}

func FactsCounter(fact1 *string, fact2 *string, fact3 *string) int {
	var sum int = 0
	sum += (IsEmptyFact(fact1) + IsEmptyFact(fact2) + IsEmptyFact(fact3))
	return sum
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	age := r.FormValue("age")
	var facts_count int = FactsCounter(&name, &age, &address)
	if facts_count == 0 {
		fmt.Fprintf(w, "I dont know anything about you :(\n")
	} else {
		if facts_count == 1 {
			fmt.Fprintf(w, "I know alittle about you\n")
		} else if facts_count == 2 {
			fmt.Fprintf(w, "I know almost everything about you\n")
		} else if facts_count == 3 {
			fmt.Fprintf(w, "I know everything about you\n")
		}
		if name != "" {
			fmt.Fprintf(w, "Your name is %s\n", name)
		}
		if address != "" {
			fmt.Fprintf(w, "Your address is %s\n", address)
		}
		if age != "" {
			fmt.Fprintf(w, "Your age is %s\n", age)
		}
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", FormHandler)
	http.HandleFunc("/hello", HelloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
