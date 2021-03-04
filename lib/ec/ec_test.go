package ec

import (
	"fmt"
	"testing"
)

/*
 go test -v -count=1
*/
const (
	EC_TEST_ERROR_1 = ECMAX + 1
)

func TestExample(t *testing.T) {
	Register(EC_TEST_ERROR_1, "name-100", "description-100")
	fmt.Println(Description(EC_TEST_ERROR_1))
	fmt.Println(Detail(EC_TEST_ERROR_1))

	var err error
	err = EC(EC_TEST_ERROR_1)
	fmt.Println(err)
}
