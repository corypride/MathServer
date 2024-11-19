package main

import (
	// "encoding/json"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	// "fmt"
	"log"
	"net/http"

	// "strconv"
	"unicode"
	// "slices"
)

// curl -X POST -d '{"aValue":"2+2"}' -│H "Content-Type: application/json" http://localhost:8080/equation

type server struct {
	router *http.ServeMux
}

type RawEquation struct {
	AValue string `json:"aValue"`
}

type ParsedEquation struct {
	left int
	right int
	op operator

}



func (op operator) perform (a , b int) int{
	switch op {
	case ADD:
		return a + b
	case Del:
		return a - b
	case Mult:
		return a * b
	case Div:
		return a / b
	default:
		return -1
	}
}
func fromStringToOp(a string) (operator, error){
	switch a {
	case "+":
		return ADD, nil
	case "-":
		return Del, nil
	case "*":
		return Mult, nil
	case "/":
		return Div, nil
	default:
		return -1, errors.New("Invalid operator")
	}
	
}
type operator int
const ( 
	ADD operator = iota
	Del 
	Mult	
	Div 	
)

type Calculation struct {
	Result string `json:"aValue"`
}


func indexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func includesOperatorAndNumbers(eq RawEquation) ParsedEquation {
	operators := []string{"+", "-", "*", "/"}
	var lft string
	var rgt string
	var strOp string


	var isLeft bool = true

	for _, char := range eq.AValue {
		// see if char is an operator
		if isOperator := contains(operators, string(char)); isOperator {
			// op, err := fromString(string(char))
			strOp = string(char)
			isLeft = false

		} else {
			// Check number
			if unicode.IsNumber(char) && isLeft {
				lft += string(char)
			} else if unicode.IsNumber(char) && ! isLeft{
				rgt += string(char)
			}
		}
	}

	if lft == "" || rgt == "" {
		// return nil, errors.New("Invalid string")
		fmt.Println("left or right side of equation is empty")
	}

	fmt.Println(lft)
	fmt.Println(rgt)
	fmt.Println(strOp)

	var lInt, lErr = strconv.Atoi(lft)
	var rInt, rErr = strconv.Atoi(rgt)
	 convertedOp, opErr := fromStringToOp(strOp)

	if lErr != nil {
		fmt.Println("Left side ERROR")
	}

	if rErr != nil {
		fmt.Println("Right side ERROR")
	}

	if opErr != nil {
		fmt.Println("Operator ERROR")

	}

	pe := ParsedEquation{left: lInt ,right: rInt, op: convertedOp}
	//TODO: run check here to see that numbers slice is length=2 for now but think in the 
	// future that there may be more than 2 groups of numbers and more than one operator
	return pe
}

func main() {
	
	// test := RawEquation{AValue: "33*5"}
	// fmt.Println("working now!!!")
	// includesOperatorAndNumbers(test)

	router := http.NewServeMux()
	srv := server{router: router}

	srv.router.HandleFunc("GET /", handleIndex)
	srv.router.HandleFunc("POST /equation", handleEquation)
	http.ListenAndServe(":8080", srv.router)
	
	
}

// TODO: handle num op num
// choose data type json or raw
// Handle calculation



func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func handleEquation(w http.ResponseWriter, r *http.Request) {
	// qryValues := r.URL.Query()["equation"]
	var eq RawEquation
	err := json.NewDecoder(r.Body).Decode(&eq)

	if err != nil {
		log.Fatal("Error decoding into struct")
	}
	defer r.Body.Close()

	pe := includesOperatorAndNumbers(eq)
	pe.op.perform(pe.left,pe.right)

	log.Println("Equation recieved:", pe.op.perform(pe.left,pe.right))
	// Split the equation into numbers and operators
	

	// if eq.AValue != "" {
	// 	for char := range eq.AValue {
	// 		fmt.Println(eq.AValue[char])
	// 	}
	// } else {
	// 	fmt.Println("The value is not present!")
	// }

}
