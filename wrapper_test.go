package gonix

import "fmt"
import "testing"

func TestHello(*testing.T) {
	fmt.Printf("stupid start valid test \n")
}

func TestHandle(*testing.T) {
	h := NixHandleCreate()
	defer NixHandleDestroy(h)

	s := NixStorePath(h)

	fmt.Printf(" Nix Store: %s \n", s)
}

func TestChainCreateDestroy(*testing.T) {
	h := NixHandleCreate()
	NixHandleDestroy(h)

	fmt.Printf(" Destroy One \n")

	g := NixHandleCreate()
	NixHandleDestroy(g)

	fmt.Printf(" Destroy Two \n")

	f := NixHandleCreate()
	NixHandleDestroy(f)

	fmt.Printf(" Destroy Three \n")
}
