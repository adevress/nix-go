package nix

import "fmt"
import "testing"

func TestHello(*testing.T) {
	fmt.Printf("stupid start valid test \n")
}

func TestHandle(t *testing.T) {
	h, err := New()
    if(err != nil){
        t.Fail()
    }
    
	defer Delete(h)

	s := h.StorePath()

	fmt.Printf(" Nix Store: %s \n", s)
}

func TestChainCreateDestroy(t *testing.T) {
	h, err := New()
    if(err != nil){
        t.Fail()
    }
    
	Delete(h)

	fmt.Printf(" Destroy One \n")

	g, err := New()
    if(err != nil){
        t.Fail()
    }
       
	Delete(g)

	fmt.Printf(" Destroy Two \n")

	f, err := New()
    if(err != nil){
        t.Fail()
    }
        
	Delete(f)

	fmt.Printf(" Destroy Three \n")
}
