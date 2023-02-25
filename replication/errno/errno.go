package errno

//go:generate stringer -type=Errno
//go:generate gobe .

type Errno uint16

const (
	ERRNO_NOERROR Errno = iota
	ERRNO_UNKNOWN
	ERRNO_NOTIMPLEMENTED
	ERRNO_STATUSNOTNORMAL
	ERRNO_NOTINSTABLEVIEW
	ERRNO_VIEWMISMATCH
	ERRNO_DUPLICATE
	ERRNO_NOTLEADER
	ERRNO_REPLAY
	ERRNO_LATECOMMIT
	ERRNO_EARLYCOMMIT
)

func (e Errno) Error() string {
	return e.String()
}
