// +build linux aix
<<<<<<< HEAD
// +build !js
=======
>>>>>>> dev

package logrus

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS

func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}
<<<<<<< HEAD
=======

>>>>>>> dev
