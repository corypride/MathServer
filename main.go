package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	// "slices"
)

// curl -X POST -d '{"aValue":"2+2"}' -│H "Content-Type: application/json" http://localhost:8080/equation

type server struct {
	router *http.ServeMux
}

type Equation struct {
	AValue string `json:"aValue"`
}

type Calculation struct {
	Result string `json:"aValue"`
}

// Helper function to check if a slice contains a particular item
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func includesOperatorAndNumbers(letter string) (int, int, string) {
	operators := []string{"+", "-", "*", "/"}
	var numbers []int
	var operator string

	for _, char := range letter {
		// see if char is an operator
		if isOperator := contains(operators, string(char)); isOperator {
			operator = string(char)
		} else {
			// cast char to int
			if num, err := strconv.Atoi(string(char)); err == nil {
				numbers = append(numbers, num)
			}
		}
	}
	//TODO: run check here to see that numbers slice is length=2 for now but think in the 
	// future that there may be more than 2 groups of numbers and more than one operator
	return numbers[0], numbers[1], operator
}


func performCalculation(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			log.Println("Error: Division by zero")
			return 0
		}
		return num1 / num2
	default:
		log.Println("Error: Unsupported operator")
		return 0
	}
}
func main() {
	// router := http.NewServeMux()
	// srv := server{router: router}

	// srv.router.HandleFunc("GET /", handleIndex)
	// srv.router.HandleFunc("POST /equation", handleEquation)
	// http.ListenAndServe(":8080", srv.router)

	str := "9/0"
	num1, num2, op := includesOperatorAndNumbers(str)
	fmt.Println(performCalculation(num1, num2, op))
	
}

// TODO: handle num op num
// choose data type json or raw
// Handle calculation



// func handleIndex(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World"))
// }

// func handleEquation(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte ("Hello World"))
// 	// need to check the qry string for the num value, if it is there, store it to a storage structure

// 	// repeat the process, but this time checking for the operator. This will require storing a list of the possible operators to start with
// 	// symbols := []string{"*","/", "+","-","({)",")"}
// 	// qryValues := r.URL.Query()["equation"]
// 	eq := Equation{}

// 	err := json.NewDecoder(r.Body).Decode(&eq)
// 	// eq.printContents()
// 	if err != nil {
// 		log.Fatal("Error decoding into struct")
// 	}
// 	defer r.Body.Close()
// 	log.Println("Equation recieved", int( eq.AValue[0]))
// 	// Split the equation into numbers and operators
	

// 	// if eq.AValue != "" {
// 	// 	for char := range eq.AValue {
// 	// 		fmt.Println(eq.AValue[char])
// 	// 	}
// 	// } else {
// 	// 	fmt.Println("The value is not present!")
// 	// }

// }
