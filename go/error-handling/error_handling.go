package erratum

/*
import (
	"errors"
) */

func (e FrobError) String() string {
	return "FrobError " + e.defrobTag + "..."
}

func (r Resource) Close() error {
}

func (r Resource) Defrob(inp string) {
	e := new(FrobError)
	e.defrobTag = "input string is not 'hello' ... so i must quit..."
}

func (r Resource) Frob(inp string) {
	defer Defrob(inp)
	if inp != "hello" {
		panic("Panic on Frob....")
	}
}

// error handling example
func Use(o ResourceOpener, input string) error {

	// open resource
	resource, error := o()
	if error != nil {
		return error
	}

}
