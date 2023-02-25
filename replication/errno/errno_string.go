// Code generated by "stringer -type=Errno"; DO NOT EDIT.

package errno

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ERRNO_NOERROR-0]
	_ = x[ERRNO_UNKNOWN-1]
	_ = x[ERRNO_NOTIMPLEMENTED-2]
	_ = x[ERRNO_STATUSNOTNORMAL-3]
	_ = x[ERRNO_NOTINSTABLEVIEW-4]
	_ = x[ERRNO_VIEWMISMATCH-5]
	_ = x[ERRNO_DUPLICATE-6]
	_ = x[ERRNO_NOTLEADER-7]
	_ = x[ERRNO_REPLAY-8]
	_ = x[ERRNO_LATECOMMIT-9]
	_ = x[ERRNO_EARLYCOMMIT-10]
}

const _Errno_name = "ERRNO_NOERRORERRNO_UNKNOWNERRNO_NOTIMPLEMENTEDERRNO_STATUSNOTNORMALERRNO_NOTINSTABLEVIEWERRNO_VIEWMISMATCHERRNO_DUPLICATEERRNO_NOTLEADERERRNO_REPLAYERRNO_LATECOMMITERRNO_EARLYCOMMIT"

var _Errno_index = [...]uint8{0, 13, 26, 46, 67, 88, 106, 121, 136, 148, 164, 181}

func (i Errno) String() string {
	if i >= Errno(len(_Errno_index)-1) {
		return "Errno(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Errno_name[_Errno_index[i]:_Errno_index[i+1]]
}
