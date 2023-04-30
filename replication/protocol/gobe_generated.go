// Code generated by "gobe ."; DO NOT EDIT.
// versions:
//     gobe: v0.2.1

package protocol

import (
	ns25570 "github.com/lemon-mint/super-system/replication/errno"
)

func (ns25519 *Commit) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (struct{ViewNumber uint64; CommitNumber uint64})(ns25519)

	// ZZ: (uint64)(ns25519.ViewNumber)
	ns25520 += 8

	// ZZ: (uint64)(ns25519.CommitNumber)
	ns25520 += 8

	return ns25520
}

func (ns25521 *Commit) MarshalGOBE(dst []byte) uint64 {
	var ns25522 uint64

	// ZZ: (struct{ViewNumber uint64; CommitNumber uint64})(ns25521)

	// ZZ: (uint64)(ns25521.ViewNumber)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.ViewNumber >> 0)
	dst[ns25522+1] = byte(ns25521.ViewNumber >> 8)
	dst[ns25522+2] = byte(ns25521.ViewNumber >> 16)
	dst[ns25522+3] = byte(ns25521.ViewNumber >> 24)
	dst[ns25522+4] = byte(ns25521.ViewNumber >> 32)
	dst[ns25522+5] = byte(ns25521.ViewNumber >> 40)
	dst[ns25522+6] = byte(ns25521.ViewNumber >> 48)
	dst[ns25522+7] = byte(ns25521.ViewNumber >> 56)
	ns25522 += 8

	// ZZ: (uint64)(ns25521.CommitNumber)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.CommitNumber >> 0)
	dst[ns25522+1] = byte(ns25521.CommitNumber >> 8)
	dst[ns25522+2] = byte(ns25521.CommitNumber >> 16)
	dst[ns25522+3] = byte(ns25521.CommitNumber >> 24)
	dst[ns25522+4] = byte(ns25521.CommitNumber >> 32)
	dst[ns25522+5] = byte(ns25521.CommitNumber >> 40)
	dst[ns25522+6] = byte(ns25521.CommitNumber >> 48)
	dst[ns25522+7] = byte(ns25521.CommitNumber >> 56)
	ns25522 += 8

	return ns25522
}

func (ns25523 *Commit) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ViewNumber uint64; CommitNumber uint64})(ns25523)

	// ZZ: (uint64)(ns25523.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25523.CommitNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.CommitNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25524 *Message) SizeGOBE() uint64 {
	var ns25525 uint64

	// ZZ: (struct{Source uint64; GroupID uint64; Type ..MessageType; ..Propose "gobe_enum:\"Type=MT_Propose\""; ..Prepare "gobe_enum:\"Type=MT_Prepare\""; ..PrepareAcceptance "gobe_enum:\"Type=MT_PrepareAcceptance\""; ..ProposeRejection "gobe_enum:\"Type=MT_ProposeRejection\""; ..PrepareRejection "gobe_enum:\"Type=MT_PrepareRejection\""; ..Commit "gobe_enum:\"Type=MT_Commit\""})(ns25524)

	// ZZ: (uint64)(ns25524.Source)
	ns25525 += 8

	// ZZ: (uint64)(ns25524.GroupID)
	ns25525 += 8

	// ZZ: (..MessageType)(ns25524.Type)

	// ZZ: (uint8)(ns25524.Type)
	ns25525 += 1
	// ENUM: KEY=<Type>
	switch ns25524.Type {
	case MT_Commit: // Type == MT_Commit

		// ZZ: (..Commit)(ns25524.Commit)
		ns25525 += ns25524.Commit.SizeGOBE()
	case MT_Prepare: // Type == MT_Prepare

		// ZZ: (..Prepare)(ns25524.Prepare)
		ns25525 += ns25524.Prepare.SizeGOBE()
	case MT_PrepareAcceptance: // Type == MT_PrepareAcceptance

		// ZZ: (..PrepareAcceptance)(ns25524.PrepareAcceptance)
		ns25525 += ns25524.PrepareAcceptance.SizeGOBE()
	case MT_PrepareRejection: // Type == MT_PrepareRejection

		// ZZ: (..PrepareRejection)(ns25524.PrepareRejection)
		ns25525 += ns25524.PrepareRejection.SizeGOBE()
	case MT_Propose: // Type == MT_Propose

		// ZZ: (..Propose)(ns25524.Propose)
		ns25525 += ns25524.Propose.SizeGOBE()
	case MT_ProposeRejection: // Type == MT_ProposeRejection

		// ZZ: (..ProposeRejection)(ns25524.ProposeRejection)
		ns25525 += ns25524.ProposeRejection.SizeGOBE()
	}

	return ns25525
}

func (ns25526 *Message) MarshalGOBE(dst []byte) uint64 {
	var ns25527 uint64

	// ZZ: (struct{Source uint64; GroupID uint64; Type ..MessageType; ..Propose "gobe_enum:\"Type=MT_Propose\""; ..Prepare "gobe_enum:\"Type=MT_Prepare\""; ..PrepareAcceptance "gobe_enum:\"Type=MT_PrepareAcceptance\""; ..ProposeRejection "gobe_enum:\"Type=MT_ProposeRejection\""; ..PrepareRejection "gobe_enum:\"Type=MT_PrepareRejection\""; ..Commit "gobe_enum:\"Type=MT_Commit\""})(ns25526)

	// ZZ: (uint64)(ns25526.Source)
	_ = dst[ns25527+7]
	dst[ns25527+0] = byte(ns25526.Source >> 0)
	dst[ns25527+1] = byte(ns25526.Source >> 8)
	dst[ns25527+2] = byte(ns25526.Source >> 16)
	dst[ns25527+3] = byte(ns25526.Source >> 24)
	dst[ns25527+4] = byte(ns25526.Source >> 32)
	dst[ns25527+5] = byte(ns25526.Source >> 40)
	dst[ns25527+6] = byte(ns25526.Source >> 48)
	dst[ns25527+7] = byte(ns25526.Source >> 56)
	ns25527 += 8

	// ZZ: (uint64)(ns25526.GroupID)
	_ = dst[ns25527+7]
	dst[ns25527+0] = byte(ns25526.GroupID >> 0)
	dst[ns25527+1] = byte(ns25526.GroupID >> 8)
	dst[ns25527+2] = byte(ns25526.GroupID >> 16)
	dst[ns25527+3] = byte(ns25526.GroupID >> 24)
	dst[ns25527+4] = byte(ns25526.GroupID >> 32)
	dst[ns25527+5] = byte(ns25526.GroupID >> 40)
	dst[ns25527+6] = byte(ns25526.GroupID >> 48)
	dst[ns25527+7] = byte(ns25526.GroupID >> 56)
	ns25527 += 8

	// ZZ: (..MessageType)(ns25526.Type)

	// ZZ: (uint8)(ns25526.Type)
	dst[ns25527+0] = byte(ns25526.Type >> 0)
	ns25527++
	// ENUM: KEY=<Type>
	switch ns25526.Type {
	case MT_Commit: // Type == MT_Commit

		// ZZ: (..Commit)(ns25526.Commit)
		ns25527 += ns25526.Commit.MarshalGOBE(dst[ns25527:])
	case MT_Prepare: // Type == MT_Prepare

		// ZZ: (..Prepare)(ns25526.Prepare)
		ns25527 += ns25526.Prepare.MarshalGOBE(dst[ns25527:])
	case MT_PrepareAcceptance: // Type == MT_PrepareAcceptance

		// ZZ: (..PrepareAcceptance)(ns25526.PrepareAcceptance)
		ns25527 += ns25526.PrepareAcceptance.MarshalGOBE(dst[ns25527:])
	case MT_PrepareRejection: // Type == MT_PrepareRejection

		// ZZ: (..PrepareRejection)(ns25526.PrepareRejection)
		ns25527 += ns25526.PrepareRejection.MarshalGOBE(dst[ns25527:])
	case MT_Propose: // Type == MT_Propose

		// ZZ: (..Propose)(ns25526.Propose)
		ns25527 += ns25526.Propose.MarshalGOBE(dst[ns25527:])
	case MT_ProposeRejection: // Type == MT_ProposeRejection

		// ZZ: (..ProposeRejection)(ns25526.ProposeRejection)
		ns25527 += ns25526.ProposeRejection.MarshalGOBE(dst[ns25527:])
	}

	return ns25527
}

func (ns25528 *Message) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Source uint64; GroupID uint64; Type ..MessageType; ..Propose "gobe_enum:\"Type=MT_Propose\""; ..Prepare "gobe_enum:\"Type=MT_Prepare\""; ..PrepareAcceptance "gobe_enum:\"Type=MT_PrepareAcceptance\""; ..ProposeRejection "gobe_enum:\"Type=MT_ProposeRejection\""; ..PrepareRejection "gobe_enum:\"Type=MT_PrepareRejection\""; ..Commit "gobe_enum:\"Type=MT_Commit\""})(ns25528)

	// ZZ: (uint64)(ns25528.Source)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25528.Source = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25528.GroupID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25528.GroupID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (..MessageType)(ns25528.Type)

	// ZZ: (uint8)(ns25528.Type)
	if uint64(len(src)) < offset+1 {
		return
	}
	_ = src[offset+0]
	ns25528.Type = MessageType(
		uint8(src[offset+0]) << 0)
	offset += 1
	// ENUM: KEY=<Type>
	switch ns25528.Type {
	case MT_Commit: // Type == MT_Commit

		// ZZ: (..Commit)(ns25528.Commit)
		ns25529, ns25530 := ns25528.Commit.UnmarshalGOBE(src[offset:])
		offset += ns25529
		if !ns25530 {
			return
		}
	case MT_Prepare: // Type == MT_Prepare

		// ZZ: (..Prepare)(ns25528.Prepare)
		ns25531, ns25532 := ns25528.Prepare.UnmarshalGOBE(src[offset:])
		offset += ns25531
		if !ns25532 {
			return
		}
	case MT_PrepareAcceptance: // Type == MT_PrepareAcceptance

		// ZZ: (..PrepareAcceptance)(ns25528.PrepareAcceptance)
		ns25533, ns25534 := ns25528.PrepareAcceptance.UnmarshalGOBE(src[offset:])
		offset += ns25533
		if !ns25534 {
			return
		}
	case MT_PrepareRejection: // Type == MT_PrepareRejection

		// ZZ: (..PrepareRejection)(ns25528.PrepareRejection)
		ns25535, ns25536 := ns25528.PrepareRejection.UnmarshalGOBE(src[offset:])
		offset += ns25535
		if !ns25536 {
			return
		}
	case MT_Propose: // Type == MT_Propose

		// ZZ: (..Propose)(ns25528.Propose)
		ns25537, ns25538 := ns25528.Propose.UnmarshalGOBE(src[offset:])
		offset += ns25537
		if !ns25538 {
			return
		}
	case MT_ProposeRejection: // Type == MT_ProposeRejection

		// ZZ: (..ProposeRejection)(ns25528.ProposeRejection)
		ns25539, ns25540 := ns25528.ProposeRejection.UnmarshalGOBE(src[offset:])
		offset += ns25539
		if !ns25540 {
			return
		}
	}

	ok = true
	return
}

func (ns25541 *MessageType) SizeGOBE() uint64 {
	var ns25542 uint64

	// ZZ: (..MessageType)(*ns25541)

	// ZZ: (uint8)(*ns25541)
	ns25542 += 1

	return ns25542
}

func (ns25543 *MessageType) MarshalGOBE(dst []byte) uint64 {
	var ns25544 uint64

	// ZZ: (..MessageType)(*ns25543)

	// ZZ: (uint8)(*ns25543)
	dst[ns25544+0] = byte(*ns25543 >> 0)
	ns25544++

	return ns25544
}

func (ns25545 *MessageType) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (..MessageType)(*ns25545)

	// ZZ: (uint8)(*ns25545)
	if uint64(len(src)) < offset+1 {
		return
	}
	_ = src[offset+0]
	*ns25545 = MessageType(
		uint8(src[offset+0]) << 0)
	offset += 1

	ok = true
	return
}

func (ns25546 *OperationEntry) SizeGOBE() uint64 {
	var ns25547 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64; Operation []byte})(ns25546)

	// ZZ: (uint64)(ns25546.ViewNumber)
	ns25547 += 8

	// ZZ: (uint64)(ns25546.OperationNumber)
	ns25547 += 8

	// ZZ: ([]byte)(ns25546.Operation)
	ns25547 += 8 + uint64(len(ns25546.Operation))

	return ns25547
}

func (ns25548 *OperationEntry) MarshalGOBE(dst []byte) uint64 {
	var ns25549 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64; Operation []byte})(ns25548)

	// ZZ: (uint64)(ns25548.ViewNumber)
	_ = dst[ns25549+7]
	dst[ns25549+0] = byte(ns25548.ViewNumber >> 0)
	dst[ns25549+1] = byte(ns25548.ViewNumber >> 8)
	dst[ns25549+2] = byte(ns25548.ViewNumber >> 16)
	dst[ns25549+3] = byte(ns25548.ViewNumber >> 24)
	dst[ns25549+4] = byte(ns25548.ViewNumber >> 32)
	dst[ns25549+5] = byte(ns25548.ViewNumber >> 40)
	dst[ns25549+6] = byte(ns25548.ViewNumber >> 48)
	dst[ns25549+7] = byte(ns25548.ViewNumber >> 56)
	ns25549 += 8

	// ZZ: (uint64)(ns25548.OperationNumber)
	_ = dst[ns25549+7]
	dst[ns25549+0] = byte(ns25548.OperationNumber >> 0)
	dst[ns25549+1] = byte(ns25548.OperationNumber >> 8)
	dst[ns25549+2] = byte(ns25548.OperationNumber >> 16)
	dst[ns25549+3] = byte(ns25548.OperationNumber >> 24)
	dst[ns25549+4] = byte(ns25548.OperationNumber >> 32)
	dst[ns25549+5] = byte(ns25548.OperationNumber >> 40)
	dst[ns25549+6] = byte(ns25548.OperationNumber >> 48)
	dst[ns25549+7] = byte(ns25548.OperationNumber >> 56)
	ns25549 += 8

	// ZZ: ([]byte)(ns25548.Operation)
	var ns25550 uint64 = uint64(len(ns25548.Operation))
	_ = dst[ns25549+7]
	dst[ns25549+0] = byte(ns25550 >> 0)
	dst[ns25549+1] = byte(ns25550 >> 8)
	dst[ns25549+2] = byte(ns25550 >> 16)
	dst[ns25549+3] = byte(ns25550 >> 24)
	dst[ns25549+4] = byte(ns25550 >> 32)
	dst[ns25549+5] = byte(ns25550 >> 40)
	dst[ns25549+6] = byte(ns25550 >> 48)
	dst[ns25549+7] = byte(ns25550 >> 56)
	copy(dst[ns25549+8:], ns25548.Operation)
	ns25549 += 8 + ns25550

	return ns25549
}

func (ns25551 *OperationEntry) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64; Operation []byte})(ns25551)

	// ZZ: (uint64)(ns25551.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25551.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25551.OperationNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25551.OperationNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: ([]byte)(ns25551.Operation)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25552 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25552 {
		return
	}
	ns25551.Operation = src[offset : offset+ns25552]
	offset += ns25552

	ok = true
	return
}

func (ns25553 *Prepare) SizeGOBE() uint64 {
	var ns25554 uint64

	// ZZ: (struct{..OperationEntry; CommitNumber uint64})(ns25553)

	// ZZ: (..OperationEntry)(ns25553.OperationEntry)
	ns25554 += ns25553.OperationEntry.SizeGOBE()

	// ZZ: (uint64)(ns25553.CommitNumber)
	ns25554 += 8

	return ns25554
}

func (ns25555 *Prepare) MarshalGOBE(dst []byte) uint64 {
	var ns25556 uint64

	// ZZ: (struct{..OperationEntry; CommitNumber uint64})(ns25555)

	// ZZ: (..OperationEntry)(ns25555.OperationEntry)
	ns25556 += ns25555.OperationEntry.MarshalGOBE(dst[ns25556:])

	// ZZ: (uint64)(ns25555.CommitNumber)
	_ = dst[ns25556+7]
	dst[ns25556+0] = byte(ns25555.CommitNumber >> 0)
	dst[ns25556+1] = byte(ns25555.CommitNumber >> 8)
	dst[ns25556+2] = byte(ns25555.CommitNumber >> 16)
	dst[ns25556+3] = byte(ns25555.CommitNumber >> 24)
	dst[ns25556+4] = byte(ns25555.CommitNumber >> 32)
	dst[ns25556+5] = byte(ns25555.CommitNumber >> 40)
	dst[ns25556+6] = byte(ns25555.CommitNumber >> 48)
	dst[ns25556+7] = byte(ns25555.CommitNumber >> 56)
	ns25556 += 8

	return ns25556
}

func (ns25557 *Prepare) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{..OperationEntry; CommitNumber uint64})(ns25557)

	// ZZ: (..OperationEntry)(ns25557.OperationEntry)
	ns25558, ns25559 := ns25557.OperationEntry.UnmarshalGOBE(src[offset:])
	offset += ns25558
	if !ns25559 {
		return
	}

	// ZZ: (uint64)(ns25557.CommitNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25557.CommitNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25560 *PrepareAcceptance) SizeGOBE() uint64 {
	var ns25561 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64})(ns25560)

	// ZZ: (uint64)(ns25560.ViewNumber)
	ns25561 += 8

	// ZZ: (uint64)(ns25560.OperationNumber)
	ns25561 += 8

	return ns25561
}

func (ns25562 *PrepareAcceptance) MarshalGOBE(dst []byte) uint64 {
	var ns25563 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64})(ns25562)

	// ZZ: (uint64)(ns25562.ViewNumber)
	_ = dst[ns25563+7]
	dst[ns25563+0] = byte(ns25562.ViewNumber >> 0)
	dst[ns25563+1] = byte(ns25562.ViewNumber >> 8)
	dst[ns25563+2] = byte(ns25562.ViewNumber >> 16)
	dst[ns25563+3] = byte(ns25562.ViewNumber >> 24)
	dst[ns25563+4] = byte(ns25562.ViewNumber >> 32)
	dst[ns25563+5] = byte(ns25562.ViewNumber >> 40)
	dst[ns25563+6] = byte(ns25562.ViewNumber >> 48)
	dst[ns25563+7] = byte(ns25562.ViewNumber >> 56)
	ns25563 += 8

	// ZZ: (uint64)(ns25562.OperationNumber)
	_ = dst[ns25563+7]
	dst[ns25563+0] = byte(ns25562.OperationNumber >> 0)
	dst[ns25563+1] = byte(ns25562.OperationNumber >> 8)
	dst[ns25563+2] = byte(ns25562.OperationNumber >> 16)
	dst[ns25563+3] = byte(ns25562.OperationNumber >> 24)
	dst[ns25563+4] = byte(ns25562.OperationNumber >> 32)
	dst[ns25563+5] = byte(ns25562.OperationNumber >> 40)
	dst[ns25563+6] = byte(ns25562.OperationNumber >> 48)
	dst[ns25563+7] = byte(ns25562.OperationNumber >> 56)
	ns25563 += 8

	return ns25563
}

func (ns25564 *PrepareAcceptance) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64})(ns25564)

	// ZZ: (uint64)(ns25564.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25564.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25564.OperationNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25564.OperationNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25565 *PrepareRejection) SizeGOBE() uint64 {
	var ns25566 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25565)

	// ZZ: (uint64)(ns25565.ViewNumber)
	ns25566 += 8

	// ZZ: (uint64)(ns25565.OperationNumber)
	ns25566 += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25565.Reason)

	// ZZ: (uint16)(ns25565.Reason)
	ns25566 += 2

	return ns25566
}

func (ns25567 *PrepareRejection) MarshalGOBE(dst []byte) uint64 {
	var ns25568 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25567)

	// ZZ: (uint64)(ns25567.ViewNumber)
	_ = dst[ns25568+7]
	dst[ns25568+0] = byte(ns25567.ViewNumber >> 0)
	dst[ns25568+1] = byte(ns25567.ViewNumber >> 8)
	dst[ns25568+2] = byte(ns25567.ViewNumber >> 16)
	dst[ns25568+3] = byte(ns25567.ViewNumber >> 24)
	dst[ns25568+4] = byte(ns25567.ViewNumber >> 32)
	dst[ns25568+5] = byte(ns25567.ViewNumber >> 40)
	dst[ns25568+6] = byte(ns25567.ViewNumber >> 48)
	dst[ns25568+7] = byte(ns25567.ViewNumber >> 56)
	ns25568 += 8

	// ZZ: (uint64)(ns25567.OperationNumber)
	_ = dst[ns25568+7]
	dst[ns25568+0] = byte(ns25567.OperationNumber >> 0)
	dst[ns25568+1] = byte(ns25567.OperationNumber >> 8)
	dst[ns25568+2] = byte(ns25567.OperationNumber >> 16)
	dst[ns25568+3] = byte(ns25567.OperationNumber >> 24)
	dst[ns25568+4] = byte(ns25567.OperationNumber >> 32)
	dst[ns25568+5] = byte(ns25567.OperationNumber >> 40)
	dst[ns25568+6] = byte(ns25567.OperationNumber >> 48)
	dst[ns25568+7] = byte(ns25567.OperationNumber >> 56)
	ns25568 += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25567.Reason)

	// ZZ: (uint16)(ns25567.Reason)
	_ = dst[ns25568+1]
	dst[ns25568+0] = byte(ns25567.Reason >> 0)
	dst[ns25568+1] = byte(ns25567.Reason >> 8)
	ns25568 += 2

	return ns25568
}

func (ns25569 *PrepareRejection) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25569)

	// ZZ: (uint64)(ns25569.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25569.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25569.OperationNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25569.OperationNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25569.Reason)

	// ZZ: (uint16)(ns25569.Reason)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25569.Reason = ns25570.Errno(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}

func (ns25571 *Propose) SizeGOBE() uint64 {
	var ns25572 uint64

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Operation []byte})(ns25571)

	// ZZ: (uint64)(ns25571.ClientID)
	ns25572 += 8

	// ZZ: (uint64)(ns25571.SequenceNumber)
	ns25572 += 8

	// ZZ: ([]byte)(ns25571.Operation)
	ns25572 += 8 + uint64(len(ns25571.Operation))

	return ns25572
}

func (ns25573 *Propose) MarshalGOBE(dst []byte) uint64 {
	var ns25574 uint64

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Operation []byte})(ns25573)

	// ZZ: (uint64)(ns25573.ClientID)
	_ = dst[ns25574+7]
	dst[ns25574+0] = byte(ns25573.ClientID >> 0)
	dst[ns25574+1] = byte(ns25573.ClientID >> 8)
	dst[ns25574+2] = byte(ns25573.ClientID >> 16)
	dst[ns25574+3] = byte(ns25573.ClientID >> 24)
	dst[ns25574+4] = byte(ns25573.ClientID >> 32)
	dst[ns25574+5] = byte(ns25573.ClientID >> 40)
	dst[ns25574+6] = byte(ns25573.ClientID >> 48)
	dst[ns25574+7] = byte(ns25573.ClientID >> 56)
	ns25574 += 8

	// ZZ: (uint64)(ns25573.SequenceNumber)
	_ = dst[ns25574+7]
	dst[ns25574+0] = byte(ns25573.SequenceNumber >> 0)
	dst[ns25574+1] = byte(ns25573.SequenceNumber >> 8)
	dst[ns25574+2] = byte(ns25573.SequenceNumber >> 16)
	dst[ns25574+3] = byte(ns25573.SequenceNumber >> 24)
	dst[ns25574+4] = byte(ns25573.SequenceNumber >> 32)
	dst[ns25574+5] = byte(ns25573.SequenceNumber >> 40)
	dst[ns25574+6] = byte(ns25573.SequenceNumber >> 48)
	dst[ns25574+7] = byte(ns25573.SequenceNumber >> 56)
	ns25574 += 8

	// ZZ: ([]byte)(ns25573.Operation)
	var ns25575 uint64 = uint64(len(ns25573.Operation))
	_ = dst[ns25574+7]
	dst[ns25574+0] = byte(ns25575 >> 0)
	dst[ns25574+1] = byte(ns25575 >> 8)
	dst[ns25574+2] = byte(ns25575 >> 16)
	dst[ns25574+3] = byte(ns25575 >> 24)
	dst[ns25574+4] = byte(ns25575 >> 32)
	dst[ns25574+5] = byte(ns25575 >> 40)
	dst[ns25574+6] = byte(ns25575 >> 48)
	dst[ns25574+7] = byte(ns25575 >> 56)
	copy(dst[ns25574+8:], ns25573.Operation)
	ns25574 += 8 + ns25575

	return ns25574
}

func (ns25576 *Propose) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Operation []byte})(ns25576)

	// ZZ: (uint64)(ns25576.ClientID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25576.ClientID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25576.SequenceNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25576.SequenceNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: ([]byte)(ns25576.Operation)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25577 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25577 {
		return
	}
	ns25576.Operation = src[offset : offset+ns25577]
	offset += ns25577

	ok = true
	return
}

func (ns25578 *ProposeRejection) SizeGOBE() uint64 {
	var ns25579 uint64

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25578)

	// ZZ: (uint64)(ns25578.ClientID)
	ns25579 += 8

	// ZZ: (uint64)(ns25578.SequenceNumber)
	ns25579 += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25578.Reason)

	// ZZ: (uint16)(ns25578.Reason)
	ns25579 += 2

	return ns25579
}

func (ns25580 *ProposeRejection) MarshalGOBE(dst []byte) uint64 {
	var ns25581 uint64

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25580)

	// ZZ: (uint64)(ns25580.ClientID)
	_ = dst[ns25581+7]
	dst[ns25581+0] = byte(ns25580.ClientID >> 0)
	dst[ns25581+1] = byte(ns25580.ClientID >> 8)
	dst[ns25581+2] = byte(ns25580.ClientID >> 16)
	dst[ns25581+3] = byte(ns25580.ClientID >> 24)
	dst[ns25581+4] = byte(ns25580.ClientID >> 32)
	dst[ns25581+5] = byte(ns25580.ClientID >> 40)
	dst[ns25581+6] = byte(ns25580.ClientID >> 48)
	dst[ns25581+7] = byte(ns25580.ClientID >> 56)
	ns25581 += 8

	// ZZ: (uint64)(ns25580.SequenceNumber)
	_ = dst[ns25581+7]
	dst[ns25581+0] = byte(ns25580.SequenceNumber >> 0)
	dst[ns25581+1] = byte(ns25580.SequenceNumber >> 8)
	dst[ns25581+2] = byte(ns25580.SequenceNumber >> 16)
	dst[ns25581+3] = byte(ns25580.SequenceNumber >> 24)
	dst[ns25581+4] = byte(ns25580.SequenceNumber >> 32)
	dst[ns25581+5] = byte(ns25580.SequenceNumber >> 40)
	dst[ns25581+6] = byte(ns25580.SequenceNumber >> 48)
	dst[ns25581+7] = byte(ns25580.SequenceNumber >> 56)
	ns25581 += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25580.Reason)

	// ZZ: (uint16)(ns25580.Reason)
	_ = dst[ns25581+1]
	dst[ns25581+0] = byte(ns25580.Reason >> 0)
	dst[ns25581+1] = byte(ns25580.Reason >> 8)
	ns25581 += 2

	return ns25581
}

func (ns25582 *ProposeRejection) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25582)

	// ZZ: (uint64)(ns25582.ClientID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25582.ClientID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25582.SequenceNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25582.SequenceNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25582.Reason)

	// ZZ: (uint16)(ns25582.Reason)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25582.Reason = ns25570.Errno(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}

func (ns25583 *Status) SizeGOBE() uint64 {
	var ns25584 uint64

	// ZZ: (..Status)(*ns25583)

	// ZZ: (uint16)(*ns25583)
	ns25584 += 2

	return ns25584
}

func (ns25585 *Status) MarshalGOBE(dst []byte) uint64 {
	var ns25586 uint64

	// ZZ: (..Status)(*ns25585)

	// ZZ: (uint16)(*ns25585)
	_ = dst[ns25586+1]
	dst[ns25586+0] = byte(*ns25585 >> 0)
	dst[ns25586+1] = byte(*ns25585 >> 8)
	ns25586 += 2

	return ns25586
}

func (ns25587 *Status) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (..Status)(*ns25587)

	// ZZ: (uint16)(*ns25587)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	*ns25587 = Status(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}