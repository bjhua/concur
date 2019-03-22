package atomic

import (
	"fmt"
	"sync"
	"testing"
)

const nanoTimeOut = 10

var n1, n2, n3, n4 int
var err1, err2, err3, err4 error


func TestExchanger_(t *testing.T) {
	e := NewExchanger()
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(){
		n1, err1 = e.Exchange(Role_Yin, 66, nanoTimeOut)
		wg.Done()
	}()

	go func(){
		n2, err2 = e.Exchange(Role_Yang, 77, nanoTimeOut)
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
	fmt.Printf("%d, %d\n", n1, n2)


}


