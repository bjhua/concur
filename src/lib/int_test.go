package lib

import (
	"fmt"
	"testing"
)


func TestInt(t *testing.T){
	var zero, one, two, three, four int32 = 0, 1, 2, 3, 4


	fmt.Printf("%d, %d, %d, %d, %d\n", zero/2, one/2, two/2, three/2, four/2)
	fmt.Printf("%d, %d, %d, %d, %d\n", zero%2, one%2, two%2, three%2, four%2)

}
