package exec

import (
	"fmt"
	"testing"
)

/*
 go test -v -count=1
*/
func funcStdout(data string) int {
	fmt.Printf("out:\n%v", data)
	return 0
}

func funcStderr(data string) int {
	fmt.Printf("err:\n%v", data)
	return 0
}

func TestAsyncStdout(t *testing.T) {
	err := CommandAsyncStdout("d:/test.sh", funcStdout, funcStderr)
	if err != nil {
		fmt.Printf("err:%v", err)
	}

	err = ScriptAsyncStdout(funcStdout, funcStderr, "d:/test.sh", "this-is-test.")
	if err != nil {
		fmt.Printf("err:%v", err)
	}
}
