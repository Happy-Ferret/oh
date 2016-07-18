// Released under an MIT license. See LICENSE.

// +build windows

package task

import (
	"os"
	"syscall"
)

func exitStatus(proc *os.Process) int {
	status, err := proc.Wait()
	if err != nil {
		return -1
	}

	return status.Sys().(syscall.WaitStatus).ExitStatus()
}
