// Code generated by "stringer -type=MessageType"; DO NOT EDIT.

package protocol

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MT_Unknown-0]
	_ = x[MT_Propose-1]
	_ = x[MT_Prepare-2]
	_ = x[MT_ProposeRejection-3]
	_ = x[MT_PrepareRejection-4]
	_ = x[MT_PrepareAcceptance-5]
	_ = x[MT_Commit-6]
}

const _MessageType_name = "MT_UnknownMT_ProposeMT_PrepareMT_ProposeRejectionMT_PrepareRejectionMT_PrepareAcceptanceMT_Commit"

var _MessageType_index = [...]uint8{0, 10, 20, 30, 49, 68, 88, 97}

func (i MessageType) String() string {
	if i >= MessageType(len(_MessageType_index)-1) {
		return "MessageType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}
