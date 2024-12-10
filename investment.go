// a program needs at least a package, it will be the 'main' package 
// main is a special name, it will be the first file that will be executed when the whole projet starts running. it's the main entry point  
// packages can be exported and imported so they can be used in other files 

package main 

// fmt is a go's standard library, more can be found at https://pkg.go.dev/std
// no comma in the import list 
import (
	"fmt" 
	"math"
	)
// DOUBLE QUOTE OR BACKTICKS FOR STRINGS
func main()  {
	var investmentAmount float64 = 1000 //automatically go will infer the data type as int, needs to define to float64. if not assigned properly, we have to use the float64() function to parse it
	var expectedReturnRate = 5.5 //float64
	var year float64 = 10

//	var futureValue = float64(investmentAmount)*math.Pow((1+expectedReturnRate/100),float64(year))

	futureValue:= math.Pow(1 + expectedReturnRate/100, year) * investmentAmount  	
	fmt.Print(futureValue)

}

// main() is executed by go 
