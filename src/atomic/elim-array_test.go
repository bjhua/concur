package atomic

import (
	"fmt"
	"sync"
	"testing"
)

func TestElimArray_(t *testing.T){
	a := NewElimArray(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(){
		n1, err1 = a.Exchange(0, Role_Yin, 66)
		wg.Done()
	}()

	go func(){
		n2, err2 = a.Exchange(0, Role_Yang, 77)
		wg.Done()
	}()

	/*
	//time.Sleep(1*time.Second)
	go func(){
		n3, err3 = e.Exchange(Role_Yin, 88, 1000)
		wg.Done()
	}()

	go func(){
		n4, err4 = e.Exchange(Role_Yang, 99, 1000)
		wg.Done()
	}()
*/

	wg.Wait()

	if err1!=nil{
		fmt.Printf("1 error: %s\n", err1.Error())
	}
	if err2 != nil{
		fmt.Printf("2 error: %s\n", err2.Error())
	}

/*
	if err3 != nil{
		fmt.Printf("3 error: %s\n", err3.Error())
	}
	if err4 != nil{
		fmt.Printf("4 error: %s\n", err4.Error())
	}

*/
	fmt.Printf("%d, %d, %d, %d\n", n1, n2, n3, n4)

	if n1!=77 || n2!=66 || n3!=99 || n4!=88{
		fmt.Printf("%d, %d, %d, %d\n", n1, n2, n3, n4)
		fmt.Printf("error: %s\n", "value not exchanged")
	}
}


