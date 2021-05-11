package ec_test

import (
	"fmt"
	"github.com/75912001/xr/lib/ec"
	"testing"
)

/*
 go test -v -count=1
*/
const (
	EC_TEST_ERROR_1 = ec.ECMAX + 1
)

func TestExample(t *testing.T) {
	ec.Register(EC_TEST_ERROR_1, "name-100", "description-100")
	fmt.Println(ec.Description(EC_TEST_ERROR_1))
	fmt.Println(ec.Detail(EC_TEST_ERROR_1))

	var err error
	err = ec.EC(EC_TEST_ERROR_1)
	fmt.Println(err)
}
