package pty

import (
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"io"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"unsafe"
)

type PTY struct {
	Master        *os.File
	MasterEncoded io.WriteCloser
	Slave         *os.File
}

func New() *PTY {
	// open master
	master, err := os.OpenFile("/dev/pts/ptmx", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}

	// unlock slave
	var unlock int32
	syscall.Syscall(syscall.SYS_IOCTL, master.Fd(), syscall.TIOCSPTLCK, uintptr(unsafe.Pointer(&unlock)))

	// find out slave name
	var ptyno uint32
	syscall.Syscall(syscall.SYS_IOCTL, master.Fd(), syscall.TIOCGPTN, uintptr(unsafe.Pointer(&ptyno)))
	if ptyno == 0 {
		panic("Failed to get ptyno")
	}

	// open slave
	slave, err := os.OpenFile("/dev/pts/"+strconv.Itoa(int(ptyno)), os.O_RDWR|syscall.O_NOCTTY, 0)
	if err != nil {
		panic(err)
	}

	// apply proper encoding
	encodedMaster, err := charset.NewWriter("ISO-8859-1", master)
	if err != nil {
		panic(err)
	}

	return &PTY{master, encodedMaster, slave}
}

func (pty *PTY) AdaptCommand(cmd *exec.Cmd) {
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = new(syscall.SysProcAttr)
	}
	//pty.Slave.Chown(int(cmd.SysProcAttr.Credential.Uid), -1)
	cmd.Stdin = pty.Slave
	cmd.Stdout = pty.Slave
	cmd.Stderr = pty.Slave
	cmd.SysProcAttr.Setsid = true
}

type winsize struct {
	ws_row, ws_col, ws_xpixel, ws_ypixel uint16
}

func (pty *PTY) SetSize(x, y uint16) {
	winsize := winsize{ws_col: x, ws_row: y}
	syscall.Syscall(syscall.SYS_IOCTL, pty.Slave.Fd(), syscall.TIOCSWINSZ, uintptr(unsafe.Pointer(&winsize)))
}
