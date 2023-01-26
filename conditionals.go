
package main

import (
	"fmt"
	
) 

var pl = fmt.Println
// Comment 
/* MultiLine Comments 
*/

func main() {  
	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !
	// ! -> not
	iAge := 5
	if(iAge >= 1) && (iAge <= 18){
		pl("important BirthDay")
	} else if (iAge == 21) || (iAge == 50){
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("important Birthday")

	} else {
		pl("Not Important Birthday")
	}
	pl("!true = ", !true)
}