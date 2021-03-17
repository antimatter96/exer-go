package erratum

import (
	"fmt"
)

// Use tries to open and frobs a given resource
func Use(o ResourceOpener, input string) (err error) {
	var resc Resource

	for true {
		resc, err = o()
		if err == nil {
			break
		}
		_, ok := err.(TransientError)
		if !ok {
			return err
		}
	}

	defer resc.Close()

	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(FrobError); ok {
				err = v.inner
				resc.Defrob(v.defrobTag)
			}
			err = fmt.Errorf(fmt.Sprint(r))
		}
	}()

	resc.Frob(input)
	return
}
