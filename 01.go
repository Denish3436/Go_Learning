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
    cV3 := " 5000005"
	cV4, err := strconv.Atoi(cV3)
	if err == nil {
		pl(cV4)
	} else {
		log.Fatal(err)
	}






}