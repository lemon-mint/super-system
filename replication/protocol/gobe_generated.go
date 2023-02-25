// Code generated by "gobe ."; DO NOT EDIT.
// versions:
//     gobe: v0.2.1

package protocol

import (
	ns25563 "github.com/lemon-mint/super-system/replication/errno"
)

func (ns25519 *Message) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (struct{Source uint64; GroupID uint64; Type ..MessageType; ..Propose "gobe_enum:\"Type=MT_Propose\""; ..Prepare "gobe_enum:\"Type=MT_Prepare\""; ..PrepareAcceptance "gobe_enum:\"Type=MT_PrepareAcceptance\""; ..ProposeRejection "gobe_enum:\"Type=MT_ProposeRejection\""; ..PrepareRejection "gobe_enum:\"Type=MT_PrepareRejection\""})(ns25519)

	// ZZ: (uint64)(ns25519.Source)
	ns25520 += 8

	// ZZ: (uint64)(ns25519.GroupID)
	ns25520 += 8

	// ZZ: (..MessageType)(ns25519.Type)

	// ZZ: (uint8)(ns25519.Type)
	ns25520 += 1
	// ENUM: KEY=<Type>
	switch ns25519.Type {
	case MT_Prepare: // Type == MT_Prepare

		// ZZ: (..Prepare)(ns25519.Prepare)
		ns25520 += ns25519.Prepare.SizeGOBE()
	case MT_PrepareAcceptance: // Type == MT_PrepareAcceptance

		// ZZ: (..PrepareAcceptance)(ns25519.PrepareAcceptance)
		ns25520 += ns25519.PrepareAcceptance.SizeGOBE()
	case MT_PrepareRejection: // Type == MT_PrepareRejection

		// ZZ: (..PrepareRejection)(ns25519.PrepareRejection)
		ns25520 += ns25519.PrepareRejection.SizeGOBE()
	case MT_Propose: // Type == MT_Propose

		// ZZ: (..Propose)(ns25519.Propose)
		ns25520 += ns25519.Propose.SizeGOBE()
	case MT_ProposeRejection: // Type == MT_ProposeRejection

		// ZZ: (..ProposeRejection)(ns25519.ProposeRejection)
		ns25520 += ns25519.ProposeRejection.SizeGOBE()
	}

	return ns25520
}

func (ns25521 *Message) MarshalGOBE(dst []byte) uint64 {
	var ns25522 uint64

	// ZZ: (struct{Source uint64; GroupID uint64; Type ..MessageType; ..Propose "gobe_enum:\"Type=MT_Propose\""; ..Prepare "gobe_enum:\"Type=MT_Prepare\""; ..PrepareAcceptance "gobe_enum:\"Type=MT_PrepareAcceptance\""; ..ProposeRejection "gobe_enum:\"Type=MT_ProposeRejection\""; ..PrepareRejection "gobe_enum:\"Type=MT_PrepareRejection\""})(ns25521)

	// ZZ: (uint64)(ns25521.Source)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.Source >> 0)
	dst[ns25522+1] = byte(ns25521.Source >> 8)
	dst[ns25522+2] = byte(ns25521.Source >> 16)
	dst[ns25522+3] = byte(ns25521.Source >> 24)
	dst[ns25522+4] = byte(ns25521.Source >> 32)
	dst[ns25522+5] = byte(ns25521.Source >> 40)
	dst[ns25522+6] = byte(ns25521.Source >> 48)
	dst[ns25522+7] = byte(ns25521.Source >> 56)
	ns25522 += 8

	// ZZ: (uint64)(ns25521.GroupID)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.GroupID >> 0)
	dst[ns25522+1] = byte(ns25521.GroupID >> 8)
	dst[ns25522+2] = byte(ns25521.GroupID >> 16)
	dst[ns25522+3] = byte(ns25521.GroupID >> 24)
	dst[ns25522+4] = byte(ns25521.GroupID >> 32)
	dst[ns25522+5] = byte(ns25521.GroupID >> 40)
	dst[ns25522+6] = byte(ns25521.GroupID >> 48)
	dst[ns25522+7] = byte(ns25521.GroupID >> 56)
	ns25522 += 8

	// ZZ: (..MessageType)(ns25521.Type)

	// ZZ: (uint8)(ns25521.Type)
	dst[ns25522+0] = byte(ns25521.Type >> 0)
	ns25522++
	// ENUM: KEY=<Type>
	switch ns25521.Type {
	case MT_Prepare: // Type == MT_Prepare

		// ZZ: (..Prepare)(ns25521.Prepare)
		ns25522 += ns25521.Prepare.MarshalGOBE(dst[ns25522:])
	case MT_PrepareAcceptance: // Type == MT_PrepareAcceptance

		// ZZ: (..PrepareAcceptance)(ns25521.PrepareAcceptance)
		ns25522 += ns25521.PrepareAcceptance.MarshalGOBE(dst[ns25522:])
	case MT_PrepareRejection: // Type == MT_PrepareRejection

		// ZZ: (..PrepareRejection)(ns25521.PrepareRejection)
		ns25522 += ns25521.PrepareRejection.MarshalGOBE(dst[ns25522:])
	case MT_Propose: // Type == MT_Propose

		// ZZ: (..Propose)(ns25521.Propose)
		ns25522 += ns25521.Propose.MarshalGOBE(dst[ns25522:])
	case MT_ProposeRejection: // Type == MT_ProposeRejection

		// ZZ: (..ProposeRejection)(ns25521.ProposeRejection)
		ns25522 += ns25521.ProposeRejection.MarshalGOBE(dst[ns25522:])
	}

	return ns25522
}

func (ns25523 *Message) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{Source uint64; GroupID uint64; Type ..MessageType; ..Propose "gobe_enum:\"Type=MT_Propose\""; ..Prepare "gobe_enum:\"Type=MT_Prepare\""; ..PrepareAcceptance "gobe_enum:\"Type=MT_PrepareAcceptance\""; ..ProposeRejection "gobe_enum:\"Type=MT_ProposeRejection\""; ..PrepareRejection "gobe_enum:\"Type=MT_PrepareRejection\""})(ns25523)

	// ZZ: (uint64)(ns25523.Source)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.Source = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25523.GroupID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.GroupID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (..MessageType)(ns25523.Type)

	// ZZ: (uint8)(ns25523.Type)
	if uint64(len(src)) < offset+1 {
		return
	}
	_ = src[offset+0]
	ns25523.Type = MessageType(
		uint8(src[offset+0]) << 0)
	offset += 1
	// ENUM: KEY=<Type>
	switch ns25523.Type {
	case MT_Prepare: // Type == MT_Prepare

		// ZZ: (..Prepare)(ns25523.Prepare)
		ns25524, ns25525 := ns25523.Prepare.UnmarshalGOBE(src[offset:])
		offset += ns25524
		if !ns25525 {
			return
		}
	case MT_PrepareAcceptance: // Type == MT_PrepareAcceptance

		// ZZ: (..PrepareAcceptance)(ns25523.PrepareAcceptance)
		ns25526, ns25527 := ns25523.PrepareAcceptance.UnmarshalGOBE(src[offset:])
		offset += ns25526
		if !ns25527 {
			return
		}
	case MT_PrepareRejection: // Type == MT_PrepareRejection

		// ZZ: (..PrepareRejection)(ns25523.PrepareRejection)
		ns25528, ns25529 := ns25523.PrepareRejection.UnmarshalGOBE(src[offset:])
		offset += ns25528
		if !ns25529 {
			return
		}
	case MT_Propose: // Type == MT_Propose

		// ZZ: (..Propose)(ns25523.Propose)
		ns25530, ns25531 := ns25523.Propose.UnmarshalGOBE(src[offset:])
		offset += ns25530
		if !ns25531 {
			return
		}
	case MT_ProposeRejection: // Type == MT_ProposeRejection

		// ZZ: (..ProposeRejection)(ns25523.ProposeRejection)
		ns25532, ns25533 := ns25523.ProposeRejection.UnmarshalGOBE(src[offset:])
		offset += ns25532
		if !ns25533 {
			return
		}
	}

	ok = true
	return
}

func (ns25534 *MessageType) SizeGOBE() uint64 {
	var ns25535 uint64

	// ZZ: (..MessageType)(*ns25534)

	// ZZ: (uint8)(*ns25534)
	ns25535 += 1

	return ns25535
}

func (ns25536 *MessageType) MarshalGOBE(dst []byte) uint64 {
	var ns25537 uint64

	// ZZ: (..MessageType)(*ns25536)

	// ZZ: (uint8)(*ns25536)
	dst[ns25537+0] = byte(*ns25536 >> 0)
	ns25537++

	return ns25537
}

func (ns25538 *MessageType) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (..MessageType)(*ns25538)

	// ZZ: (uint8)(*ns25538)
	if uint64(len(src)) < offset+1 {
		return
	}
	_ = src[offset+0]
	*ns25538 = MessageType(
		uint8(src[offset+0]) << 0)
	offset += 1

	ok = true
	return
}

func (ns25539 *OperationEntry) SizeGOBE() uint64 {
	var ns25540 uint64

	// ZZ: (struct{OperationNumber uint64; Operation []byte; ViewNumber uint64})(ns25539)

	// ZZ: (uint64)(ns25539.OperationNumber)
	ns25540 += 8

	// ZZ: ([]byte)(ns25539.Operation)
	ns25540 += 8 + uint64(len(ns25539.Operation))

	// ZZ: (uint64)(ns25539.ViewNumber)
	ns25540 += 8

	return ns25540
}

func (ns25541 *OperationEntry) MarshalGOBE(dst []byte) uint64 {
	var ns25542 uint64

	// ZZ: (struct{OperationNumber uint64; Operation []byte; ViewNumber uint64})(ns25541)

	// ZZ: (uint64)(ns25541.OperationNumber)
	_ = dst[ns25542+7]
	dst[ns25542+0] = byte(ns25541.OperationNumber >> 0)
	dst[ns25542+1] = byte(ns25541.OperationNumber >> 8)
	dst[ns25542+2] = byte(ns25541.OperationNumber >> 16)
	dst[ns25542+3] = byte(ns25541.OperationNumber >> 24)
	dst[ns25542+4] = byte(ns25541.OperationNumber >> 32)
	dst[ns25542+5] = byte(ns25541.OperationNumber >> 40)
	dst[ns25542+6] = byte(ns25541.OperationNumber >> 48)
	dst[ns25542+7] = byte(ns25541.OperationNumber >> 56)
	ns25542 += 8

	// ZZ: ([]byte)(ns25541.Operation)
	var ns25543 uint64 = uint64(len(ns25541.Operation))
	_ = dst[ns25542+7]
	dst[ns25542+0] = byte(ns25543 >> 0)
	dst[ns25542+1] = byte(ns25543 >> 8)
	dst[ns25542+2] = byte(ns25543 >> 16)
	dst[ns25542+3] = byte(ns25543 >> 24)
	dst[ns25542+4] = byte(ns25543 >> 32)
	dst[ns25542+5] = byte(ns25543 >> 40)
	dst[ns25542+6] = byte(ns25543 >> 48)
	dst[ns25542+7] = byte(ns25543 >> 56)
	copy(dst[ns25542+8:], ns25541.Operation)
	ns25542 += 8 + ns25543

	// ZZ: (uint64)(ns25541.ViewNumber)
	_ = dst[ns25542+7]
	dst[ns25542+0] = byte(ns25541.ViewNumber >> 0)
	dst[ns25542+1] = byte(ns25541.ViewNumber >> 8)
	dst[ns25542+2] = byte(ns25541.ViewNumber >> 16)
	dst[ns25542+3] = byte(ns25541.ViewNumber >> 24)
	dst[ns25542+4] = byte(ns25541.ViewNumber >> 32)
	dst[ns25542+5] = byte(ns25541.ViewNumber >> 40)
	dst[ns25542+6] = byte(ns25541.ViewNumber >> 48)
	dst[ns25542+7] = byte(ns25541.ViewNumber >> 56)
	ns25542 += 8

	return ns25542
}

func (ns25544 *OperationEntry) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{OperationNumber uint64; Operation []byte; ViewNumber uint64})(ns25544)

	// ZZ: (uint64)(ns25544.OperationNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25544.OperationNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: ([]byte)(ns25544.Operation)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25545 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25545 {
		return
	}
	ns25544.Operation = src[offset : offset+ns25545]
	offset += ns25545

	// ZZ: (uint64)(ns25544.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25544.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25546 *Prepare) SizeGOBE() uint64 {
	var ns25547 uint64

	// ZZ: (struct{..OperationEntry; CommitNumber uint64})(ns25546)

	// ZZ: (..OperationEntry)(ns25546.OperationEntry)
	ns25547 += ns25546.OperationEntry.SizeGOBE()

	// ZZ: (uint64)(ns25546.CommitNumber)
	ns25547 += 8

	return ns25547
}

func (ns25548 *Prepare) MarshalGOBE(dst []byte) uint64 {
	var ns25549 uint64

	// ZZ: (struct{..OperationEntry; CommitNumber uint64})(ns25548)

	// ZZ: (..OperationEntry)(ns25548.OperationEntry)
	ns25549 += ns25548.OperationEntry.MarshalGOBE(dst[ns25549:])

	// ZZ: (uint64)(ns25548.CommitNumber)
	_ = dst[ns25549+7]
	dst[ns25549+0] = byte(ns25548.CommitNumber >> 0)
	dst[ns25549+1] = byte(ns25548.CommitNumber >> 8)
	dst[ns25549+2] = byte(ns25548.CommitNumber >> 16)
	dst[ns25549+3] = byte(ns25548.CommitNumber >> 24)
	dst[ns25549+4] = byte(ns25548.CommitNumber >> 32)
	dst[ns25549+5] = byte(ns25548.CommitNumber >> 40)
	dst[ns25549+6] = byte(ns25548.CommitNumber >> 48)
	dst[ns25549+7] = byte(ns25548.CommitNumber >> 56)
	ns25549 += 8

	return ns25549
}

func (ns25550 *Prepare) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{..OperationEntry; CommitNumber uint64})(ns25550)

	// ZZ: (..OperationEntry)(ns25550.OperationEntry)
	ns25551, ns25552 := ns25550.OperationEntry.UnmarshalGOBE(src[offset:])
	offset += ns25551
	if !ns25552 {
		return
	}

	// ZZ: (uint64)(ns25550.CommitNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25550.CommitNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25553 *PrepareAcceptance) SizeGOBE() uint64 {
	var ns25554 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64})(ns25553)

	// ZZ: (uint64)(ns25553.ViewNumber)
	ns25554 += 8

	// ZZ: (uint64)(ns25553.OperationNumber)
	ns25554 += 8

	return ns25554
}

func (ns25555 *PrepareAcceptance) MarshalGOBE(dst []byte) uint64 {
	var ns25556 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64})(ns25555)

	// ZZ: (uint64)(ns25555.ViewNumber)
	_ = dst[ns25556+7]
	dst[ns25556+0] = byte(ns25555.ViewNumber >> 0)
	dst[ns25556+1] = byte(ns25555.ViewNumber >> 8)
	dst[ns25556+2] = byte(ns25555.ViewNumber >> 16)
	dst[ns25556+3] = byte(ns25555.ViewNumber >> 24)
	dst[ns25556+4] = byte(ns25555.ViewNumber >> 32)
	dst[ns25556+5] = byte(ns25555.ViewNumber >> 40)
	dst[ns25556+6] = byte(ns25555.ViewNumber >> 48)
	dst[ns25556+7] = byte(ns25555.ViewNumber >> 56)
	ns25556 += 8

	// ZZ: (uint64)(ns25555.OperationNumber)
	_ = dst[ns25556+7]
	dst[ns25556+0] = byte(ns25555.OperationNumber >> 0)
	dst[ns25556+1] = byte(ns25555.OperationNumber >> 8)
	dst[ns25556+2] = byte(ns25555.OperationNumber >> 16)
	dst[ns25556+3] = byte(ns25555.OperationNumber >> 24)
	dst[ns25556+4] = byte(ns25555.OperationNumber >> 32)
	dst[ns25556+5] = byte(ns25555.OperationNumber >> 40)
	dst[ns25556+6] = byte(ns25555.OperationNumber >> 48)
	dst[ns25556+7] = byte(ns25555.OperationNumber >> 56)
	ns25556 += 8

	return ns25556
}

func (ns25557 *PrepareAcceptance) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64})(ns25557)

	// ZZ: (uint64)(ns25557.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25557.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25557.OperationNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25557.OperationNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25558 *PrepareRejection) SizeGOBE() uint64 {
	var ns25559 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25558)

	// ZZ: (uint64)(ns25558.ViewNumber)
	ns25559 += 8

	// ZZ: (uint64)(ns25558.OperationNumber)
	ns25559 += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25558.Reason)

	// ZZ: (uint16)(ns25558.Reason)
	ns25559 += 2

	return ns25559
}

func (ns25560 *PrepareRejection) MarshalGOBE(dst []byte) uint64 {
	var ns25561 uint64

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25560)

	// ZZ: (uint64)(ns25560.ViewNumber)
	_ = dst[ns25561+7]
	dst[ns25561+0] = byte(ns25560.ViewNumber >> 0)
	dst[ns25561+1] = byte(ns25560.ViewNumber >> 8)
	dst[ns25561+2] = byte(ns25560.ViewNumber >> 16)
	dst[ns25561+3] = byte(ns25560.ViewNumber >> 24)
	dst[ns25561+4] = byte(ns25560.ViewNumber >> 32)
	dst[ns25561+5] = byte(ns25560.ViewNumber >> 40)
	dst[ns25561+6] = byte(ns25560.ViewNumber >> 48)
	dst[ns25561+7] = byte(ns25560.ViewNumber >> 56)
	ns25561 += 8

	// ZZ: (uint64)(ns25560.OperationNumber)
	_ = dst[ns25561+7]
	dst[ns25561+0] = byte(ns25560.OperationNumber >> 0)
	dst[ns25561+1] = byte(ns25560.OperationNumber >> 8)
	dst[ns25561+2] = byte(ns25560.OperationNumber >> 16)
	dst[ns25561+3] = byte(ns25560.OperationNumber >> 24)
	dst[ns25561+4] = byte(ns25560.OperationNumber >> 32)
	dst[ns25561+5] = byte(ns25560.OperationNumber >> 40)
	dst[ns25561+6] = byte(ns25560.OperationNumber >> 48)
	dst[ns25561+7] = byte(ns25560.OperationNumber >> 56)
	ns25561 += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25560.Reason)

	// ZZ: (uint16)(ns25560.Reason)
	_ = dst[ns25561+1]
	dst[ns25561+0] = byte(ns25560.Reason >> 0)
	dst[ns25561+1] = byte(ns25560.Reason >> 8)
	ns25561 += 2

	return ns25561
}

func (ns25562 *PrepareRejection) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ViewNumber uint64; OperationNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25562)

	// ZZ: (uint64)(ns25562.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25562.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25562.OperationNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25562.OperationNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25562.Reason)

	// ZZ: (uint16)(ns25562.Reason)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25562.Reason = ns25563.Errno(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}

func (ns25564 *Propose) SizeGOBE() uint64 {
	var ns25565 uint64

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Operation []byte})(ns25564)

	// ZZ: (uint64)(ns25564.ClientID)
	ns25565 += 8

	// ZZ: (uint64)(ns25564.SequenceNumber)
	ns25565 += 8

	// ZZ: ([]byte)(ns25564.Operation)
	ns25565 += 8 + uint64(len(ns25564.Operation))

	return ns25565
}

func (ns25566 *Propose) MarshalGOBE(dst []byte) uint64 {
	var ns25567 uint64

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Operation []byte})(ns25566)

	// ZZ: (uint64)(ns25566.ClientID)
	_ = dst[ns25567+7]
	dst[ns25567+0] = byte(ns25566.ClientID >> 0)
	dst[ns25567+1] = byte(ns25566.ClientID >> 8)
	dst[ns25567+2] = byte(ns25566.ClientID >> 16)
	dst[ns25567+3] = byte(ns25566.ClientID >> 24)
	dst[ns25567+4] = byte(ns25566.ClientID >> 32)
	dst[ns25567+5] = byte(ns25566.ClientID >> 40)
	dst[ns25567+6] = byte(ns25566.ClientID >> 48)
	dst[ns25567+7] = byte(ns25566.ClientID >> 56)
	ns25567 += 8

	// ZZ: (uint64)(ns25566.SequenceNumber)
	_ = dst[ns25567+7]
	dst[ns25567+0] = byte(ns25566.SequenceNumber >> 0)
	dst[ns25567+1] = byte(ns25566.SequenceNumber >> 8)
	dst[ns25567+2] = byte(ns25566.SequenceNumber >> 16)
	dst[ns25567+3] = byte(ns25566.SequenceNumber >> 24)
	dst[ns25567+4] = byte(ns25566.SequenceNumber >> 32)
	dst[ns25567+5] = byte(ns25566.SequenceNumber >> 40)
	dst[ns25567+6] = byte(ns25566.SequenceNumber >> 48)
	dst[ns25567+7] = byte(ns25566.SequenceNumber >> 56)
	ns25567 += 8

	// ZZ: ([]byte)(ns25566.Operation)
	var ns25568 uint64 = uint64(len(ns25566.Operation))
	_ = dst[ns25567+7]
	dst[ns25567+0] = byte(ns25568 >> 0)
	dst[ns25567+1] = byte(ns25568 >> 8)
	dst[ns25567+2] = byte(ns25568 >> 16)
	dst[ns25567+3] = byte(ns25568 >> 24)
	dst[ns25567+4] = byte(ns25568 >> 32)
	dst[ns25567+5] = byte(ns25568 >> 40)
	dst[ns25567+6] = byte(ns25568 >> 48)
	dst[ns25567+7] = byte(ns25568 >> 56)
	copy(dst[ns25567+8:], ns25566.Operation)
	ns25567 += 8 + ns25568

	return ns25567
}

func (ns25569 *Propose) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Operation []byte})(ns25569)

	// ZZ: (uint64)(ns25569.ClientID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25569.ClientID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25569.SequenceNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25569.SequenceNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: ([]byte)(ns25569.Operation)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25570 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25570 {
		return
	}
	ns25569.Operation = src[offset : offset+ns25570]
	offset += ns25570

	ok = true
	return
}

func (ns25571 *ProposeRejection) SizeGOBE() uint64 {
	var ns25572 uint64

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25571)

	// ZZ: (uint64)(ns25571.ClientID)
	ns25572 += 8

	// ZZ: (uint64)(ns25571.SequenceNumber)
	ns25572 += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25571.Reason)

	// ZZ: (uint16)(ns25571.Reason)
	ns25572 += 2

	return ns25572
}

func (ns25573 *ProposeRejection) MarshalGOBE(dst []byte) uint64 {
	var ns25574 uint64

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25573)

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

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25573.Reason)

	// ZZ: (uint16)(ns25573.Reason)
	_ = dst[ns25574+1]
	dst[ns25574+0] = byte(ns25573.Reason >> 0)
	dst[ns25574+1] = byte(ns25573.Reason >> 8)
	ns25574 += 2

	return ns25574
}

func (ns25575 *ProposeRejection) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{ClientID uint64; SequenceNumber uint64; Reason github.com/lemon-mint/super-system/replication/errno.Errno})(ns25575)

	// ZZ: (uint64)(ns25575.ClientID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25575.ClientID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25575.SequenceNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25575.SequenceNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (github.com/lemon-mint/super-system/replication/errno.Errno)(ns25575.Reason)

	// ZZ: (uint16)(ns25575.Reason)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	ns25575.Reason = ns25563.Errno(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}

func (ns25576 *Status) SizeGOBE() uint64 {
	var ns25577 uint64

	// ZZ: (..Status)(*ns25576)

	// ZZ: (uint16)(*ns25576)
	ns25577 += 2

	return ns25577
}

func (ns25578 *Status) MarshalGOBE(dst []byte) uint64 {
	var ns25579 uint64

	// ZZ: (..Status)(*ns25578)

	// ZZ: (uint16)(*ns25578)
	_ = dst[ns25579+1]
	dst[ns25579+0] = byte(*ns25578 >> 0)
	dst[ns25579+1] = byte(*ns25578 >> 8)
	ns25579 += 2

	return ns25579
}

func (ns25580 *Status) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (..Status)(*ns25580)

	// ZZ: (uint16)(*ns25580)
	if uint64(len(src)) < offset+2 {
		return
	}
	_ = src[offset+1]
	*ns25580 = Status(
		uint16(src[offset+0])<<0 | uint16(src[offset+1])<<8)
	offset += 2

	ok = true
	return
}
