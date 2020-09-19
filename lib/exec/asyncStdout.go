package exec

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type FuncAsyncStd func(data string) int

// args:"chmod +x /xx/xx/x.sh"
// funcStdout nil:disable stdout
// funcStderr nil:disable stderr
func CommandAsyncStdout(args string, funcStdout FuncAsyncStd, funcStderr FuncAsyncStd) (err error) {
	cmd := exec.Command("sh", "-c", args)

	var stdout io.ReadCloser
	if stdout, err = cmd.StdoutPipe(); err != nil {
		return
	}
	var stderr io.ReadCloser
	if stderr, err = cmd.StderrPipe(); err != nil {
		return
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if nil != funcStdout {
		go asyncStdout(stdout, funcStdout)
	}
	if nil != funcStderr {
		go asyncStdout(stderr, funcStderr)
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return
}

// ScriptAsyncStdout(funcStdout, funcStderr, "/xxx/xxx.sh", "arg1", "arg2", "arg3", "arg4")
func ScriptAsyncStdout(funcStdout FuncAsyncStd, funcStderr FuncAsyncStd, scriptPathFile string, arg ...string) (err error) {
	var args string = scriptPathFile
	for _, v := range arg {
		if 0 == len(v) {
			v = "\"\""
		}
		args += " " + v
	}
	return CommandAsyncStdout(args, funcStdout, funcStderr)
}

func asyncStdout(reader io.ReadCloser, fun FuncAsyncStd) (err error) {
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if num > 0 {
			b := buf[:num]
			s := strings.Split(string(b), "\n")
			line := strings.Join(s[:len(s)-1], "\n")
			fun(fmt.Sprintf("%v\n", line))
		}
	}
}
