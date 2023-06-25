package executils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(program string, arguments []string) error {
	cmd := exec.Command(program, arguments...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	scannerOut := bufio.NewScanner(stdout)
	go func() {
		for scannerOut.Scan() {
			fmt.Println(scannerOut.Text())
		}
	}()

	scannerErr := bufio.NewScanner(stderr)
	go func() {
		for scannerErr.Scan() {
			fmt.Fprintln(os.Stderr, scannerErr.Text())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
