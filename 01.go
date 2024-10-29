package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"reflect"
	"strconv"
	


)

//Create alias for long function name
var pl = fmt.Println

//    ---Functions---
func sayHello() {
	pl("hello")
}

//sum
func sum(a int, b int ) int{
	return a + b 
}
 
//multiple value 
func getTwo(x int ) (int , int) {
	return x + 1, x + 2
}

//Return Potential Error
func div(a int , b int) (ans int , err error){
	if b ==0 { 
		return 0 , fmt.Errorf("can't divided by 0")
	    
	}else {
		return a/b , nil
	}

}



//---Main---
func main() {
	fmt.Println("Hello World")
	sayHello()
	pl (sum(2, 3))
	pl(getTwo(2))
	pl (div (2, 2))

	pl("what is your name")
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)

	// ReadString will block until the delimiter is entered
	
	//name, _ := reader.ReadString('\n') // The blank identifier _ will get err and ignore it (Bad Practice)
	name, err := reader.ReadString('\n') // It is better to handle it
    if err == nil{
	   pl ("Hello", name)
    }else {
	   log.Fatal(err)
   }

// ----- VARIABLES -----

//    var v1, v2 = 1.2, 3
//    pl(v1, v2)
//    var vName string = "denish"
//    pl(vName)
// Short variable declaration (Type defined by data)
//    var v3 = "Hello"
//    pl(v3)	

// ----- DATA TYPES -----
// int, float64, bool, string, rune

pl(reflect.TypeOf(42))
pl(reflect.TypeOf(3.14))
pl(reflect.TypeOf(true))
pl(reflect.TypeOf("Hello"))
pl(reflect.TypeOf('a'))

// ----- CASTING -----
	// To cast type the type to convert to with the variable to
	// convert in parentheses
	// Doesn't work with bools or strings
	
	cV1 := 1.55
	cV2 := int(cV1)
    pl(cV2)


// ----- CONSTANTS -----
// Constants are variables that can't be changed
// They can be defined like variables but with const
// Constants can be character, string, boolean or numeric values
// Constants can't be declared using the := syntax
           
    const pi = 3.14159
    pl(pi) 

// Convert string to int (ASCII to Integer)
    cV3 := "55022"
	cV4, err := strconv.Atoi(cV3)
	if err == nil {
		pl(cV4)
	} else {
		log.Fatal(err)
	}

	
// Convert int to string (Integer to ASCII)
    cV5 := 511245
    cV6 := strconv.Itoa(cV5)
    pl(cV6)

//Convert string to float
	cV7 := "3.14159"
	cV8, err := strconv.ParseFloat(cV7, 64);
	pl(cV8)

// Use Sprintf to convert from float to string
    cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)

// ----- IF CONDITIONALS -----
// == != < <= > >= && || ! are the same as other languages

    iAge := 11

	if (iAge >=1) && (iAge <= 18){
		pl("You can go to school")
	}else if (iAge>=19) || (iAge <=60){
		pl("You can work")
	}else {
		pl("enjoy retirement")
	}

// ----- SWITCH -----
// break is automatic unless you add fallthrough
// You can switch on data types
// case can include multiple values
    iAge = 18 

	switch iAge {
	case 16 : pl("You can drive")
	case 18 : pl("You can vote")
	case 21 : pl("You can drink")
	default : pl("Enjoy your life")
	}

// ! turns bools into their opposite value
pl("!true =", !true)

// ----- STRINGS -----
// Strings are immutable
// Strings are a slice of bytes
// Strings can be created with backticks ` or double quotes "
// Strings can be added with +
// Strings can be converted to []byte

	s1 := "Hello World"
	pl(s1)
	s2 := `Hello
	World`
	pl(s2)

// Strings are a slice of bytes
	pl(s1[0])

// Convert string to []byte
	bArr := []byte(s1)
	pl(bArr)

// Convert back to string
	pl(string(bArr))

// Concatenate strings
	s3 := s1 + " again"
	pl(s3)

// Convert int to string
	s4 := strconv.Itoa(42)
	pl(s4)

sV1 := "A string"

// Get string length
pl(len(sV1))






}