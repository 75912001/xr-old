package ec

import (
	"fmt"
	"testing"
)

/*
 go test -v -count=1
*/
func TestExample(t *testing.T) {
	Register(100, "name-100", "description-100")
	for k, v := range errorInfoMap {
		fmt.Println(k, v)
	}

	fmt.Println(Description(100))
	fmt.Println(Detail(100))

	Register(999, "name-999", "description-999")

	var err error
	err = EC(999)
	fmt.Println(err)
}
