package vsrproto

func (ns25519 *Commit) SizeGOBE() uint64 {
	var ns25520 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,2,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,3,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""})(ns25519)

	// ZZ: (uint64)(ns25519.ViewNumber)
	ns25520 += 8

	// ZZ: (uint64)(ns25519.CommitNumber_MAX)
	ns25520 += 8

	// ZZ: (uint64)(ns25519.CommitNumber_MIN)
	ns25520 += 8

	return ns25520
}

func (ns25521 *Commit) MarshalGOBE(dst []byte) uint64 {
	var ns25522 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,2,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,3,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""})(ns25521)

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

	// ZZ: (uint64)(ns25521.CommitNumber_MAX)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.CommitNumber_MAX >> 0)
	dst[ns25522+1] = byte(ns25521.CommitNumber_MAX >> 8)
	dst[ns25522+2] = byte(ns25521.CommitNumber_MAX >> 16)
	dst[ns25522+3] = byte(ns25521.CommitNumber_MAX >> 24)
	dst[ns25522+4] = byte(ns25521.CommitNumber_MAX >> 32)
	dst[ns25522+5] = byte(ns25521.CommitNumber_MAX >> 40)
	dst[ns25522+6] = byte(ns25521.CommitNumber_MAX >> 48)
	dst[ns25522+7] = byte(ns25521.CommitNumber_MAX >> 56)
	ns25522 += 8

	// ZZ: (uint64)(ns25521.CommitNumber_MIN)
	_ = dst[ns25522+7]
	dst[ns25522+0] = byte(ns25521.CommitNumber_MIN >> 0)
	dst[ns25522+1] = byte(ns25521.CommitNumber_MIN >> 8)
	dst[ns25522+2] = byte(ns25521.CommitNumber_MIN >> 16)
	dst[ns25522+3] = byte(ns25521.CommitNumber_MIN >> 24)
	dst[ns25522+4] = byte(ns25521.CommitNumber_MIN >> 32)
	dst[ns25522+5] = byte(ns25521.CommitNumber_MIN >> 40)
	dst[ns25522+6] = byte(ns25521.CommitNumber_MIN >> 48)
	dst[ns25522+7] = byte(ns25521.CommitNumber_MIN >> 56)
	ns25522 += 8

	return ns25522
}

func (ns25523 *Commit) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,2,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,3,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""})(ns25523)

	// ZZ: (uint64)(ns25523.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25523.CommitNumber_MAX)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.CommitNumber_MAX = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25523.CommitNumber_MIN)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25523.CommitNumber_MIN = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25524 *DoViewChange) SizeGOBE() uint64 {
	var ns25525 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; EpochNumber uint64 "protobuf:\"varint,2,opt,name=EpochNumber,proto3\" json:\"EpochNumber,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,3,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,4,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""; PrepareOKs []*./replication/vsrproto.PrepareOK "protobuf:\"bytes,5,rep,name=PrepareOKs,proto3\" json:\"PrepareOKs,omitempty\""})(ns25524)

	// ZZ: (uint64)(ns25524.ViewNumber)
	ns25525 += 8

	// ZZ: (uint64)(ns25524.EpochNumber)
	ns25525 += 8

	// ZZ: (uint64)(ns25524.CommitNumber_MAX)
	ns25525 += 8

	// ZZ: (uint64)(ns25524.CommitNumber_MIN)
	ns25525 += 8

	// ZZ: ([]*./replication/vsrproto.PrepareOK)(ns25524.PrepareOKs)
	ns25525 += 8
	for ns25526 := 0; ns25526 < len(ns25524.PrepareOKs); ns25526++ {

		// ZZ: (*./replication/vsrproto.PrepareOK)(ns25524.PrepareOKs[ns25526])
		ns25525 += 1
		if ns25524.PrepareOKs[ns25526] != nil {

			// ZZ: (./replication/vsrproto.PrepareOK)((*ns25524.PrepareOKs[ns25526]))
			ns25525 += (*ns25524.PrepareOKs[ns25526]).SizeGOBE()
		}
	}

	return ns25525
}

func (ns25527 *DoViewChange) MarshalGOBE(dst []byte) uint64 {
	var ns25528 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; EpochNumber uint64 "protobuf:\"varint,2,opt,name=EpochNumber,proto3\" json:\"EpochNumber,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,3,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,4,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""; PrepareOKs []*./replication/vsrproto.PrepareOK "protobuf:\"bytes,5,rep,name=PrepareOKs,proto3\" json:\"PrepareOKs,omitempty\""})(ns25527)

	// ZZ: (uint64)(ns25527.ViewNumber)
	_ = dst[ns25528+7]
	dst[ns25528+0] = byte(ns25527.ViewNumber >> 0)
	dst[ns25528+1] = byte(ns25527.ViewNumber >> 8)
	dst[ns25528+2] = byte(ns25527.ViewNumber >> 16)
	dst[ns25528+3] = byte(ns25527.ViewNumber >> 24)
	dst[ns25528+4] = byte(ns25527.ViewNumber >> 32)
	dst[ns25528+5] = byte(ns25527.ViewNumber >> 40)
	dst[ns25528+6] = byte(ns25527.ViewNumber >> 48)
	dst[ns25528+7] = byte(ns25527.ViewNumber >> 56)
	ns25528 += 8

	// ZZ: (uint64)(ns25527.EpochNumber)
	_ = dst[ns25528+7]
	dst[ns25528+0] = byte(ns25527.EpochNumber >> 0)
	dst[ns25528+1] = byte(ns25527.EpochNumber >> 8)
	dst[ns25528+2] = byte(ns25527.EpochNumber >> 16)
	dst[ns25528+3] = byte(ns25527.EpochNumber >> 24)
	dst[ns25528+4] = byte(ns25527.EpochNumber >> 32)
	dst[ns25528+5] = byte(ns25527.EpochNumber >> 40)
	dst[ns25528+6] = byte(ns25527.EpochNumber >> 48)
	dst[ns25528+7] = byte(ns25527.EpochNumber >> 56)
	ns25528 += 8

	// ZZ: (uint64)(ns25527.CommitNumber_MAX)
	_ = dst[ns25528+7]
	dst[ns25528+0] = byte(ns25527.CommitNumber_MAX >> 0)
	dst[ns25528+1] = byte(ns25527.CommitNumber_MAX >> 8)
	dst[ns25528+2] = byte(ns25527.CommitNumber_MAX >> 16)
	dst[ns25528+3] = byte(ns25527.CommitNumber_MAX >> 24)
	dst[ns25528+4] = byte(ns25527.CommitNumber_MAX >> 32)
	dst[ns25528+5] = byte(ns25527.CommitNumber_MAX >> 40)
	dst[ns25528+6] = byte(ns25527.CommitNumber_MAX >> 48)
	dst[ns25528+7] = byte(ns25527.CommitNumber_MAX >> 56)
	ns25528 += 8

	// ZZ: (uint64)(ns25527.CommitNumber_MIN)
	_ = dst[ns25528+7]
	dst[ns25528+0] = byte(ns25527.CommitNumber_MIN >> 0)
	dst[ns25528+1] = byte(ns25527.CommitNumber_MIN >> 8)
	dst[ns25528+2] = byte(ns25527.CommitNumber_MIN >> 16)
	dst[ns25528+3] = byte(ns25527.CommitNumber_MIN >> 24)
	dst[ns25528+4] = byte(ns25527.CommitNumber_MIN >> 32)
	dst[ns25528+5] = byte(ns25527.CommitNumber_MIN >> 40)
	dst[ns25528+6] = byte(ns25527.CommitNumber_MIN >> 48)
	dst[ns25528+7] = byte(ns25527.CommitNumber_MIN >> 56)
	ns25528 += 8

	// ZZ: ([]*./replication/vsrproto.PrepareOK)(ns25527.PrepareOKs)
	var ns25529 uint64 = uint64(len(ns25527.PrepareOKs))
	_ = dst[ns25528+7]
	dst[ns25528+0] = byte(ns25529 >> 0)
	dst[ns25528+1] = byte(ns25529 >> 8)
	dst[ns25528+2] = byte(ns25529 >> 16)
	dst[ns25528+3] = byte(ns25529 >> 24)
	dst[ns25528+4] = byte(ns25529 >> 32)
	dst[ns25528+5] = byte(ns25529 >> 40)
	dst[ns25528+6] = byte(ns25529 >> 48)
	dst[ns25528+7] = byte(ns25529 >> 56)
	ns25528 += 8

	for ns25530 := 0; ns25530 < len(ns25527.PrepareOKs); ns25530++ {

		// ZZ: (*./replication/vsrproto.PrepareOK)(ns25527.PrepareOKs[ns25530])
		if ns25527.PrepareOKs[ns25530] != nil {
			dst[ns25528] = 1
			ns25528++

			// ZZ: (./replication/vsrproto.PrepareOK)((*ns25527.PrepareOKs[ns25530]))
			ns25528 += (*ns25527.PrepareOKs[ns25530]).MarshalGOBE(dst[ns25528:])
		} else {
			dst[ns25528] = 0
			ns25528++
		}

	}

	return ns25528
}

func (ns25531 *DoViewChange) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; EpochNumber uint64 "protobuf:\"varint,2,opt,name=EpochNumber,proto3\" json:\"EpochNumber,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,3,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,4,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""; PrepareOKs []*./replication/vsrproto.PrepareOK "protobuf:\"bytes,5,rep,name=PrepareOKs,proto3\" json:\"PrepareOKs,omitempty\""})(ns25531)

	// ZZ: (uint64)(ns25531.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25531.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25531.EpochNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25531.EpochNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25531.CommitNumber_MAX)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25531.CommitNumber_MAX = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25531.CommitNumber_MIN)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25531.CommitNumber_MIN = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: ([]*./replication/vsrproto.PrepareOK)(ns25531.PrepareOKs)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25532 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(cap(ns25531.PrepareOKs)) < ns25532 {
		if ns25532*uint64(8) <= 1<<15 {
			ns25531.PrepareOKs = make([]*PrepareOK, ns25532)
			for ns25533 := uint64(0); ns25533 < ns25532; ns25533++ {

				// ZZ: (*./replication/vsrproto.PrepareOK)(ns25531.PrepareOKs[ns25533])
				if uint64(len(src)) < offset+1 {
					return
				}
				if src[offset] == 1 {
					offset++
					if ns25531.PrepareOKs[ns25533] == nil {
						ns25531.PrepareOKs[ns25533] = new(PrepareOK)
					}

					// ZZ: (./replication/vsrproto.PrepareOK)((*ns25531.PrepareOKs[ns25533]))
					ns25535, ns25536 := (*ns25531.PrepareOKs[ns25533]).UnmarshalGOBE(src[offset:])
					offset += ns25535
					if !ns25536 {
						return
					}
				} else {
					offset++
					ns25531.PrepareOKs[ns25533] = nil
				}
			}
		} else {
			// Slice too large, Using Append
			ns25531.PrepareOKs = ns25531.PrepareOKs[:0]
			for ns25533 := uint64(0); ns25533 < ns25532; ns25533++ {
				var ns25534 *PrepareOK

				// ZZ: (*./replication/vsrproto.PrepareOK)(ns25534)
				if uint64(len(src)) < offset+1 {
					return
				}
				if src[offset] == 1 {
					offset++
					if ns25534 == nil {
						ns25534 = new(PrepareOK)
					}

					// ZZ: (./replication/vsrproto.PrepareOK)((*ns25534))
					ns25537, ns25538 := (*ns25534).UnmarshalGOBE(src[offset:])
					offset += ns25537
					if !ns25538 {
						return
					}
				} else {
					offset++
					ns25534 = nil
				}
				ns25531.PrepareOKs = append(ns25531.PrepareOKs, ns25534)
			}
		}
	} else {
		ns25531.PrepareOKs = ns25531.PrepareOKs[:ns25532]
		for ns25533 := uint64(0); ns25533 < ns25532; ns25533++ {

			// ZZ: (*./replication/vsrproto.PrepareOK)(ns25531.PrepareOKs[ns25533])
			if uint64(len(src)) < offset+1 {
				return
			}
			if src[offset] == 1 {
				offset++
				if ns25531.PrepareOKs[ns25533] == nil {
					ns25531.PrepareOKs[ns25533] = new(PrepareOK)
				}

				// ZZ: (./replication/vsrproto.PrepareOK)((*ns25531.PrepareOKs[ns25533]))
				ns25539, ns25540 := (*ns25531.PrepareOKs[ns25533]).UnmarshalGOBE(src[offset:])
				offset += ns25539
				if !ns25540 {
					return
				}
			} else {
				offset++
				ns25531.PrepareOKs[ns25533] = nil
			}
		}
	}

	ok = true
	return
}

func (ns25541 *Error) SizeGOBE() uint64 {
	var ns25542 uint64

	// ZZ: (./replication/vsrproto.Error)(*ns25541)

	// ZZ: (int32)(*ns25541)
	ns25542 += 4

	return ns25542
}

func (ns25543 *Error) MarshalGOBE(dst []byte) uint64 {
	var ns25544 uint64

	// ZZ: (./replication/vsrproto.Error)(*ns25543)

	// ZZ: (int32)(*ns25543)
	var ns25545 uint32 = uint32(*ns25543)
	_ = dst[ns25544+3]
	dst[ns25544+0] = byte(ns25545 >> 0)
	dst[ns25544+1] = byte(ns25545 >> 8)
	dst[ns25544+2] = byte(ns25545 >> 16)
	dst[ns25544+3] = byte(ns25545 >> 24)
	ns25544 += 4

	return ns25544
}

func (ns25546 *Error) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (./replication/vsrproto.Error)(*ns25546)

	// ZZ: (int32)(*ns25546)
	if uint64(len(src)) < offset+4 {
		return
	}
	_ = src[offset+3]
	*ns25546 = Error(
		uint32(src[offset+0])<<0 | uint32(src[offset+1])<<8 | uint32(src[offset+2])<<16 | uint32(src[offset+3])<<24)
	offset += 4

	ok = true
	return
}

func (ns25547 *Message) SizeGOBE() uint64 {
	var ns25548 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; GroupID uint64 "protobuf:\"varint,5,opt,name=GroupID,proto3\" json:\"GroupID,omitempty\""; Source uint64 "protobuf:\"varint,10,opt,name=Source,proto3\" json:\"Source,omitempty\""; Type ./replication/vsrproto.MessageType "protobuf:\"varint,20,opt,name=Type,proto3,enum=vsrproto.MessageType\" json:\"Type,omitempty\""; Propose *./replication/vsrproto.Propose "protobuf:\"bytes,30,opt,name=Propose,proto3,oneof\" json:\"Propose,omitempty\""; ProposeReject *./replication/vsrproto.ProposeReject "protobuf:\"bytes,31,opt,name=ProposeReject,proto3,oneof\" json:\"ProposeReject,omitempty\""; Prepare *./replication/vsrproto.Prepare "protobuf:\"bytes,32,opt,name=Prepare,proto3,oneof\" json:\"Prepare,omitempty\""; PrepareOK *./replication/vsrproto.PrepareOK "protobuf:\"bytes,33,opt,name=PrepareOK,proto3,oneof\" json:\"PrepareOK,omitempty\""; PrepareReject *./replication/vsrproto.PrepareReject "protobuf:\"bytes,34,opt,name=PrepareReject,proto3,oneof\" json:\"PrepareReject,omitempty\""; Commit *./replication/vsrproto.Commit "protobuf:\"bytes,35,opt,name=Commit,proto3,oneof\" json:\"Commit,omitempty\""; StartViewChange *./replication/vsrproto.StartViewChange "protobuf:\"bytes,36,opt,name=StartViewChange,proto3,oneof\" json:\"StartViewChange,omitempty\""; DoViewChange *./replication/vsrproto.DoViewChange "protobuf:\"bytes,37,opt,name=DoViewChange,proto3,oneof\" json:\"DoViewChange,omitempty\""; StartView *./replication/vsrproto.StartView "protobuf:\"bytes,38,opt,name=StartView,proto3,oneof\" json:\"StartView,omitempty\""})(ns25547)

	// ZZ: (uint64)(ns25547.GroupID)
	ns25548 += 8

	// ZZ: (uint64)(ns25547.Source)
	ns25548 += 8

	// ZZ: (./replication/vsrproto.MessageType)(ns25547.Type)

	// ZZ: (int32)(ns25547.Type)
	ns25548 += 4

	// ZZ: (*./replication/vsrproto.Propose)(ns25547.Propose)
	ns25548 += 1
	if ns25547.Propose != nil {

		// ZZ: (./replication/vsrproto.Propose)((*ns25547.Propose))
		ns25548 += (*ns25547.Propose).SizeGOBE()
	}

	// ZZ: (*./replication/vsrproto.ProposeReject)(ns25547.ProposeReject)
	ns25548 += 1
	if ns25547.ProposeReject != nil {

		// ZZ: (./replication/vsrproto.ProposeReject)((*ns25547.ProposeReject))
		ns25548 += (*ns25547.ProposeReject).SizeGOBE()
	}

	// ZZ: (*./replication/vsrproto.Prepare)(ns25547.Prepare)
	ns25548 += 1
	if ns25547.Prepare != nil {

		// ZZ: (./replication/vsrproto.Prepare)((*ns25547.Prepare))
		ns25548 += (*ns25547.Prepare).SizeGOBE()
	}

	// ZZ: (*./replication/vsrproto.PrepareOK)(ns25547.PrepareOK)
	ns25548 += 1
	if ns25547.PrepareOK != nil {

		// ZZ: (./replication/vsrproto.PrepareOK)((*ns25547.PrepareOK))
		ns25548 += (*ns25547.PrepareOK).SizeGOBE()
	}

	// ZZ: (*./replication/vsrproto.PrepareReject)(ns25547.PrepareReject)
	ns25548 += 1
	if ns25547.PrepareReject != nil {

		// ZZ: (./replication/vsrproto.PrepareReject)((*ns25547.PrepareReject))
		ns25548 += (*ns25547.PrepareReject).SizeGOBE()
	}

	// ZZ: (*./replication/vsrproto.Commit)(ns25547.Commit)
	ns25548 += 1
	if ns25547.Commit != nil {

		// ZZ: (./replication/vsrproto.Commit)((*ns25547.Commit))
		ns25548 += (*ns25547.Commit).SizeGOBE()
	}

	// ZZ: (*./replication/vsrproto.StartViewChange)(ns25547.StartViewChange)
	ns25548 += 1
	if ns25547.StartViewChange != nil {

		// ZZ: (./replication/vsrproto.StartViewChange)((*ns25547.StartViewChange))
		ns25548 += (*ns25547.StartViewChange).SizeGOBE()
	}

	// ZZ: (*./replication/vsrproto.DoViewChange)(ns25547.DoViewChange)
	ns25548 += 1
	if ns25547.DoViewChange != nil {

		// ZZ: (./replication/vsrproto.DoViewChange)((*ns25547.DoViewChange))
		ns25548 += (*ns25547.DoViewChange).SizeGOBE()
	}

	// ZZ: (*./replication/vsrproto.StartView)(ns25547.StartView)
	ns25548 += 1
	if ns25547.StartView != nil {

		// ZZ: (./replication/vsrproto.StartView)((*ns25547.StartView))
		ns25548 += (*ns25547.StartView).SizeGOBE()
	}

	return ns25548
}

func (ns25549 *Message) MarshalGOBE(dst []byte) uint64 {
	var ns25550 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; GroupID uint64 "protobuf:\"varint,5,opt,name=GroupID,proto3\" json:\"GroupID,omitempty\""; Source uint64 "protobuf:\"varint,10,opt,name=Source,proto3\" json:\"Source,omitempty\""; Type ./replication/vsrproto.MessageType "protobuf:\"varint,20,opt,name=Type,proto3,enum=vsrproto.MessageType\" json:\"Type,omitempty\""; Propose *./replication/vsrproto.Propose "protobuf:\"bytes,30,opt,name=Propose,proto3,oneof\" json:\"Propose,omitempty\""; ProposeReject *./replication/vsrproto.ProposeReject "protobuf:\"bytes,31,opt,name=ProposeReject,proto3,oneof\" json:\"ProposeReject,omitempty\""; Prepare *./replication/vsrproto.Prepare "protobuf:\"bytes,32,opt,name=Prepare,proto3,oneof\" json:\"Prepare,omitempty\""; PrepareOK *./replication/vsrproto.PrepareOK "protobuf:\"bytes,33,opt,name=PrepareOK,proto3,oneof\" json:\"PrepareOK,omitempty\""; PrepareReject *./replication/vsrproto.PrepareReject "protobuf:\"bytes,34,opt,name=PrepareReject,proto3,oneof\" json:\"PrepareReject,omitempty\""; Commit *./replication/vsrproto.Commit "protobuf:\"bytes,35,opt,name=Commit,proto3,oneof\" json:\"Commit,omitempty\""; StartViewChange *./replication/vsrproto.StartViewChange "protobuf:\"bytes,36,opt,name=StartViewChange,proto3,oneof\" json:\"StartViewChange,omitempty\""; DoViewChange *./replication/vsrproto.DoViewChange "protobuf:\"bytes,37,opt,name=DoViewChange,proto3,oneof\" json:\"DoViewChange,omitempty\""; StartView *./replication/vsrproto.StartView "protobuf:\"bytes,38,opt,name=StartView,proto3,oneof\" json:\"StartView,omitempty\""})(ns25549)

	// ZZ: (uint64)(ns25549.GroupID)
	_ = dst[ns25550+7]
	dst[ns25550+0] = byte(ns25549.GroupID >> 0)
	dst[ns25550+1] = byte(ns25549.GroupID >> 8)
	dst[ns25550+2] = byte(ns25549.GroupID >> 16)
	dst[ns25550+3] = byte(ns25549.GroupID >> 24)
	dst[ns25550+4] = byte(ns25549.GroupID >> 32)
	dst[ns25550+5] = byte(ns25549.GroupID >> 40)
	dst[ns25550+6] = byte(ns25549.GroupID >> 48)
	dst[ns25550+7] = byte(ns25549.GroupID >> 56)
	ns25550 += 8

	// ZZ: (uint64)(ns25549.Source)
	_ = dst[ns25550+7]
	dst[ns25550+0] = byte(ns25549.Source >> 0)
	dst[ns25550+1] = byte(ns25549.Source >> 8)
	dst[ns25550+2] = byte(ns25549.Source >> 16)
	dst[ns25550+3] = byte(ns25549.Source >> 24)
	dst[ns25550+4] = byte(ns25549.Source >> 32)
	dst[ns25550+5] = byte(ns25549.Source >> 40)
	dst[ns25550+6] = byte(ns25549.Source >> 48)
	dst[ns25550+7] = byte(ns25549.Source >> 56)
	ns25550 += 8

	// ZZ: (./replication/vsrproto.MessageType)(ns25549.Type)

	// ZZ: (int32)(ns25549.Type)
	var ns25551 uint32 = uint32(ns25549.Type)
	_ = dst[ns25550+3]
	dst[ns25550+0] = byte(ns25551 >> 0)
	dst[ns25550+1] = byte(ns25551 >> 8)
	dst[ns25550+2] = byte(ns25551 >> 16)
	dst[ns25550+3] = byte(ns25551 >> 24)
	ns25550 += 4

	// ZZ: (*./replication/vsrproto.Propose)(ns25549.Propose)
	if ns25549.Propose != nil {
		dst[ns25550] = 1
		ns25550++

		// ZZ: (./replication/vsrproto.Propose)((*ns25549.Propose))
		ns25550 += (*ns25549.Propose).MarshalGOBE(dst[ns25550:])
	} else {
		dst[ns25550] = 0
		ns25550++
	}

	// ZZ: (*./replication/vsrproto.ProposeReject)(ns25549.ProposeReject)
	if ns25549.ProposeReject != nil {
		dst[ns25550] = 1
		ns25550++

		// ZZ: (./replication/vsrproto.ProposeReject)((*ns25549.ProposeReject))
		ns25550 += (*ns25549.ProposeReject).MarshalGOBE(dst[ns25550:])
	} else {
		dst[ns25550] = 0
		ns25550++
	}

	// ZZ: (*./replication/vsrproto.Prepare)(ns25549.Prepare)
	if ns25549.Prepare != nil {
		dst[ns25550] = 1
		ns25550++

		// ZZ: (./replication/vsrproto.Prepare)((*ns25549.Prepare))
		ns25550 += (*ns25549.Prepare).MarshalGOBE(dst[ns25550:])
	} else {
		dst[ns25550] = 0
		ns25550++
	}

	// ZZ: (*./replication/vsrproto.PrepareOK)(ns25549.PrepareOK)
	if ns25549.PrepareOK != nil {
		dst[ns25550] = 1
		ns25550++

		// ZZ: (./replication/vsrproto.PrepareOK)((*ns25549.PrepareOK))
		ns25550 += (*ns25549.PrepareOK).MarshalGOBE(dst[ns25550:])
	} else {
		dst[ns25550] = 0
		ns25550++
	}

	// ZZ: (*./replication/vsrproto.PrepareReject)(ns25549.PrepareReject)
	if ns25549.PrepareReject != nil {
		dst[ns25550] = 1
		ns25550++

		// ZZ: (./replication/vsrproto.PrepareReject)((*ns25549.PrepareReject))
		ns25550 += (*ns25549.PrepareReject).MarshalGOBE(dst[ns25550:])
	} else {
		dst[ns25550] = 0
		ns25550++
	}

	// ZZ: (*./replication/vsrproto.Commit)(ns25549.Commit)
	if ns25549.Commit != nil {
		dst[ns25550] = 1
		ns25550++

		// ZZ: (./replication/vsrproto.Commit)((*ns25549.Commit))
		ns25550 += (*ns25549.Commit).MarshalGOBE(dst[ns25550:])
	} else {
		dst[ns25550] = 0
		ns25550++
	}

	// ZZ: (*./replication/vsrproto.StartViewChange)(ns25549.StartViewChange)
	if ns25549.StartViewChange != nil {
		dst[ns25550] = 1
		ns25550++

		// ZZ: (./replication/vsrproto.StartViewChange)((*ns25549.StartViewChange))
		ns25550 += (*ns25549.StartViewChange).MarshalGOBE(dst[ns25550:])
	} else {
		dst[ns25550] = 0
		ns25550++
	}

	// ZZ: (*./replication/vsrproto.DoViewChange)(ns25549.DoViewChange)
	if ns25549.DoViewChange != nil {
		dst[ns25550] = 1
		ns25550++

		// ZZ: (./replication/vsrproto.DoViewChange)((*ns25549.DoViewChange))
		ns25550 += (*ns25549.DoViewChange).MarshalGOBE(dst[ns25550:])
	} else {
		dst[ns25550] = 0
		ns25550++
	}

	// ZZ: (*./replication/vsrproto.StartView)(ns25549.StartView)
	if ns25549.StartView != nil {
		dst[ns25550] = 1
		ns25550++

		// ZZ: (./replication/vsrproto.StartView)((*ns25549.StartView))
		ns25550 += (*ns25549.StartView).MarshalGOBE(dst[ns25550:])
	} else {
		dst[ns25550] = 0
		ns25550++
	}

	return ns25550
}

func (ns25552 *Message) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; GroupID uint64 "protobuf:\"varint,5,opt,name=GroupID,proto3\" json:\"GroupID,omitempty\""; Source uint64 "protobuf:\"varint,10,opt,name=Source,proto3\" json:\"Source,omitempty\""; Type ./replication/vsrproto.MessageType "protobuf:\"varint,20,opt,name=Type,proto3,enum=vsrproto.MessageType\" json:\"Type,omitempty\""; Propose *./replication/vsrproto.Propose "protobuf:\"bytes,30,opt,name=Propose,proto3,oneof\" json:\"Propose,omitempty\""; ProposeReject *./replication/vsrproto.ProposeReject "protobuf:\"bytes,31,opt,name=ProposeReject,proto3,oneof\" json:\"ProposeReject,omitempty\""; Prepare *./replication/vsrproto.Prepare "protobuf:\"bytes,32,opt,name=Prepare,proto3,oneof\" json:\"Prepare,omitempty\""; PrepareOK *./replication/vsrproto.PrepareOK "protobuf:\"bytes,33,opt,name=PrepareOK,proto3,oneof\" json:\"PrepareOK,omitempty\""; PrepareReject *./replication/vsrproto.PrepareReject "protobuf:\"bytes,34,opt,name=PrepareReject,proto3,oneof\" json:\"PrepareReject,omitempty\""; Commit *./replication/vsrproto.Commit "protobuf:\"bytes,35,opt,name=Commit,proto3,oneof\" json:\"Commit,omitempty\""; StartViewChange *./replication/vsrproto.StartViewChange "protobuf:\"bytes,36,opt,name=StartViewChange,proto3,oneof\" json:\"StartViewChange,omitempty\""; DoViewChange *./replication/vsrproto.DoViewChange "protobuf:\"bytes,37,opt,name=DoViewChange,proto3,oneof\" json:\"DoViewChange,omitempty\""; StartView *./replication/vsrproto.StartView "protobuf:\"bytes,38,opt,name=StartView,proto3,oneof\" json:\"StartView,omitempty\""})(ns25552)

	// ZZ: (uint64)(ns25552.GroupID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25552.GroupID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25552.Source)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25552.Source = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (./replication/vsrproto.MessageType)(ns25552.Type)

	// ZZ: (int32)(ns25552.Type)
	if uint64(len(src)) < offset+4 {
		return
	}
	_ = src[offset+3]
	ns25552.Type = MessageType(
		uint32(src[offset+0])<<0 | uint32(src[offset+1])<<8 | uint32(src[offset+2])<<16 | uint32(src[offset+3])<<24)
	offset += 4

	// ZZ: (*./replication/vsrproto.Propose)(ns25552.Propose)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25552.Propose == nil {
			ns25552.Propose = new(Propose)
		}

		// ZZ: (./replication/vsrproto.Propose)((*ns25552.Propose))
		ns25553, ns25554 := (*ns25552.Propose).UnmarshalGOBE(src[offset:])
		offset += ns25553
		if !ns25554 {
			return
		}
	} else {
		offset++
		ns25552.Propose = nil
	}

	// ZZ: (*./replication/vsrproto.ProposeReject)(ns25552.ProposeReject)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25552.ProposeReject == nil {
			ns25552.ProposeReject = new(ProposeReject)
		}

		// ZZ: (./replication/vsrproto.ProposeReject)((*ns25552.ProposeReject))
		ns25555, ns25556 := (*ns25552.ProposeReject).UnmarshalGOBE(src[offset:])
		offset += ns25555
		if !ns25556 {
			return
		}
	} else {
		offset++
		ns25552.ProposeReject = nil
	}

	// ZZ: (*./replication/vsrproto.Prepare)(ns25552.Prepare)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25552.Prepare == nil {
			ns25552.Prepare = new(Prepare)
		}

		// ZZ: (./replication/vsrproto.Prepare)((*ns25552.Prepare))
		ns25557, ns25558 := (*ns25552.Prepare).UnmarshalGOBE(src[offset:])
		offset += ns25557
		if !ns25558 {
			return
		}
	} else {
		offset++
		ns25552.Prepare = nil
	}

	// ZZ: (*./replication/vsrproto.PrepareOK)(ns25552.PrepareOK)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25552.PrepareOK == nil {
			ns25552.PrepareOK = new(PrepareOK)
		}

		// ZZ: (./replication/vsrproto.PrepareOK)((*ns25552.PrepareOK))
		ns25559, ns25560 := (*ns25552.PrepareOK).UnmarshalGOBE(src[offset:])
		offset += ns25559
		if !ns25560 {
			return
		}
	} else {
		offset++
		ns25552.PrepareOK = nil
	}

	// ZZ: (*./replication/vsrproto.PrepareReject)(ns25552.PrepareReject)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25552.PrepareReject == nil {
			ns25552.PrepareReject = new(PrepareReject)
		}

		// ZZ: (./replication/vsrproto.PrepareReject)((*ns25552.PrepareReject))
		ns25561, ns25562 := (*ns25552.PrepareReject).UnmarshalGOBE(src[offset:])
		offset += ns25561
		if !ns25562 {
			return
		}
	} else {
		offset++
		ns25552.PrepareReject = nil
	}

	// ZZ: (*./replication/vsrproto.Commit)(ns25552.Commit)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25552.Commit == nil {
			ns25552.Commit = new(Commit)
		}

		// ZZ: (./replication/vsrproto.Commit)((*ns25552.Commit))
		ns25563, ns25564 := (*ns25552.Commit).UnmarshalGOBE(src[offset:])
		offset += ns25563
		if !ns25564 {
			return
		}
	} else {
		offset++
		ns25552.Commit = nil
	}

	// ZZ: (*./replication/vsrproto.StartViewChange)(ns25552.StartViewChange)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25552.StartViewChange == nil {
			ns25552.StartViewChange = new(StartViewChange)
		}

		// ZZ: (./replication/vsrproto.StartViewChange)((*ns25552.StartViewChange))
		ns25565, ns25566 := (*ns25552.StartViewChange).UnmarshalGOBE(src[offset:])
		offset += ns25565
		if !ns25566 {
			return
		}
	} else {
		offset++
		ns25552.StartViewChange = nil
	}

	// ZZ: (*./replication/vsrproto.DoViewChange)(ns25552.DoViewChange)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25552.DoViewChange == nil {
			ns25552.DoViewChange = new(DoViewChange)
		}

		// ZZ: (./replication/vsrproto.DoViewChange)((*ns25552.DoViewChange))
		ns25567, ns25568 := (*ns25552.DoViewChange).UnmarshalGOBE(src[offset:])
		offset += ns25567
		if !ns25568 {
			return
		}
	} else {
		offset++
		ns25552.DoViewChange = nil
	}

	// ZZ: (*./replication/vsrproto.StartView)(ns25552.StartView)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25552.StartView == nil {
			ns25552.StartView = new(StartView)
		}

		// ZZ: (./replication/vsrproto.StartView)((*ns25552.StartView))
		ns25569, ns25570 := (*ns25552.StartView).UnmarshalGOBE(src[offset:])
		offset += ns25569
		if !ns25570 {
			return
		}
	} else {
		offset++
		ns25552.StartView = nil
	}

	ok = true
	return
}

func (ns25571 *MessageType) SizeGOBE() uint64 {
	var ns25572 uint64

	// ZZ: (./replication/vsrproto.MessageType)(*ns25571)

	// ZZ: (int32)(*ns25571)
	ns25572 += 4

	return ns25572
}

func (ns25573 *MessageType) MarshalGOBE(dst []byte) uint64 {
	var ns25574 uint64

	// ZZ: (./replication/vsrproto.MessageType)(*ns25573)

	// ZZ: (int32)(*ns25573)
	var ns25575 uint32 = uint32(*ns25573)
	_ = dst[ns25574+3]
	dst[ns25574+0] = byte(ns25575 >> 0)
	dst[ns25574+1] = byte(ns25575 >> 8)
	dst[ns25574+2] = byte(ns25575 >> 16)
	dst[ns25574+3] = byte(ns25575 >> 24)
	ns25574 += 4

	return ns25574
}

func (ns25576 *MessageType) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (./replication/vsrproto.MessageType)(*ns25576)

	// ZZ: (int32)(*ns25576)
	if uint64(len(src)) < offset+4 {
		return
	}
	_ = src[offset+3]
	*ns25576 = MessageType(
		uint32(src[offset+0])<<0 | uint32(src[offset+1])<<8 | uint32(src[offset+2])<<16 | uint32(src[offset+3])<<24)
	offset += 4

	ok = true
	return
}

func (ns25577 *Prepare) SizeGOBE() uint64 {
	var ns25578 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; OperationNumber uint64 "protobuf:\"varint,2,opt,name=OperationNumber,proto3\" json:\"OperationNumber,omitempty\""; Propose *./replication/vsrproto.Propose "protobuf:\"bytes,3,opt,name=Propose,proto3\" json:\"Propose,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,4,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,5,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""})(ns25577)

	// ZZ: (uint64)(ns25577.ViewNumber)
	ns25578 += 8

	// ZZ: (uint64)(ns25577.OperationNumber)
	ns25578 += 8

	// ZZ: (*./replication/vsrproto.Propose)(ns25577.Propose)
	ns25578 += 1
	if ns25577.Propose != nil {

		// ZZ: (./replication/vsrproto.Propose)((*ns25577.Propose))
		ns25578 += (*ns25577.Propose).SizeGOBE()
	}

	// ZZ: (uint64)(ns25577.CommitNumber_MAX)
	ns25578 += 8

	// ZZ: (uint64)(ns25577.CommitNumber_MIN)
	ns25578 += 8

	return ns25578
}

func (ns25579 *Prepare) MarshalGOBE(dst []byte) uint64 {
	var ns25580 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; OperationNumber uint64 "protobuf:\"varint,2,opt,name=OperationNumber,proto3\" json:\"OperationNumber,omitempty\""; Propose *./replication/vsrproto.Propose "protobuf:\"bytes,3,opt,name=Propose,proto3\" json:\"Propose,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,4,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,5,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""})(ns25579)

	// ZZ: (uint64)(ns25579.ViewNumber)
	_ = dst[ns25580+7]
	dst[ns25580+0] = byte(ns25579.ViewNumber >> 0)
	dst[ns25580+1] = byte(ns25579.ViewNumber >> 8)
	dst[ns25580+2] = byte(ns25579.ViewNumber >> 16)
	dst[ns25580+3] = byte(ns25579.ViewNumber >> 24)
	dst[ns25580+4] = byte(ns25579.ViewNumber >> 32)
	dst[ns25580+5] = byte(ns25579.ViewNumber >> 40)
	dst[ns25580+6] = byte(ns25579.ViewNumber >> 48)
	dst[ns25580+7] = byte(ns25579.ViewNumber >> 56)
	ns25580 += 8

	// ZZ: (uint64)(ns25579.OperationNumber)
	_ = dst[ns25580+7]
	dst[ns25580+0] = byte(ns25579.OperationNumber >> 0)
	dst[ns25580+1] = byte(ns25579.OperationNumber >> 8)
	dst[ns25580+2] = byte(ns25579.OperationNumber >> 16)
	dst[ns25580+3] = byte(ns25579.OperationNumber >> 24)
	dst[ns25580+4] = byte(ns25579.OperationNumber >> 32)
	dst[ns25580+5] = byte(ns25579.OperationNumber >> 40)
	dst[ns25580+6] = byte(ns25579.OperationNumber >> 48)
	dst[ns25580+7] = byte(ns25579.OperationNumber >> 56)
	ns25580 += 8

	// ZZ: (*./replication/vsrproto.Propose)(ns25579.Propose)
	if ns25579.Propose != nil {
		dst[ns25580] = 1
		ns25580++

		// ZZ: (./replication/vsrproto.Propose)((*ns25579.Propose))
		ns25580 += (*ns25579.Propose).MarshalGOBE(dst[ns25580:])
	} else {
		dst[ns25580] = 0
		ns25580++
	}

	// ZZ: (uint64)(ns25579.CommitNumber_MAX)
	_ = dst[ns25580+7]
	dst[ns25580+0] = byte(ns25579.CommitNumber_MAX >> 0)
	dst[ns25580+1] = byte(ns25579.CommitNumber_MAX >> 8)
	dst[ns25580+2] = byte(ns25579.CommitNumber_MAX >> 16)
	dst[ns25580+3] = byte(ns25579.CommitNumber_MAX >> 24)
	dst[ns25580+4] = byte(ns25579.CommitNumber_MAX >> 32)
	dst[ns25580+5] = byte(ns25579.CommitNumber_MAX >> 40)
	dst[ns25580+6] = byte(ns25579.CommitNumber_MAX >> 48)
	dst[ns25580+7] = byte(ns25579.CommitNumber_MAX >> 56)
	ns25580 += 8

	// ZZ: (uint64)(ns25579.CommitNumber_MIN)
	_ = dst[ns25580+7]
	dst[ns25580+0] = byte(ns25579.CommitNumber_MIN >> 0)
	dst[ns25580+1] = byte(ns25579.CommitNumber_MIN >> 8)
	dst[ns25580+2] = byte(ns25579.CommitNumber_MIN >> 16)
	dst[ns25580+3] = byte(ns25579.CommitNumber_MIN >> 24)
	dst[ns25580+4] = byte(ns25579.CommitNumber_MIN >> 32)
	dst[ns25580+5] = byte(ns25579.CommitNumber_MIN >> 40)
	dst[ns25580+6] = byte(ns25579.CommitNumber_MIN >> 48)
	dst[ns25580+7] = byte(ns25579.CommitNumber_MIN >> 56)
	ns25580 += 8

	return ns25580
}

func (ns25581 *Prepare) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; OperationNumber uint64 "protobuf:\"varint,2,opt,name=OperationNumber,proto3\" json:\"OperationNumber,omitempty\""; Propose *./replication/vsrproto.Propose "protobuf:\"bytes,3,opt,name=Propose,proto3\" json:\"Propose,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,4,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,5,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""})(ns25581)

	// ZZ: (uint64)(ns25581.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25581.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25581.OperationNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25581.OperationNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (*./replication/vsrproto.Propose)(ns25581.Propose)
	if uint64(len(src)) < offset+1 {
		return
	}
	if src[offset] == 1 {
		offset++
		if ns25581.Propose == nil {
			ns25581.Propose = new(Propose)
		}

		// ZZ: (./replication/vsrproto.Propose)((*ns25581.Propose))
		ns25582, ns25583 := (*ns25581.Propose).UnmarshalGOBE(src[offset:])
		offset += ns25582
		if !ns25583 {
			return
		}
	} else {
		offset++
		ns25581.Propose = nil
	}

	// ZZ: (uint64)(ns25581.CommitNumber_MAX)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25581.CommitNumber_MAX = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25581.CommitNumber_MIN)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25581.CommitNumber_MIN = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25584 *PrepareOK) SizeGOBE() uint64 {
	var ns25585 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; OperationNumber uint64 "protobuf:\"varint,2,opt,name=OperationNumber,proto3\" json:\"OperationNumber,omitempty\""})(ns25584)

	// ZZ: (uint64)(ns25584.ViewNumber)
	ns25585 += 8

	// ZZ: (uint64)(ns25584.OperationNumber)
	ns25585 += 8

	return ns25585
}

func (ns25586 *PrepareOK) MarshalGOBE(dst []byte) uint64 {
	var ns25587 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; OperationNumber uint64 "protobuf:\"varint,2,opt,name=OperationNumber,proto3\" json:\"OperationNumber,omitempty\""})(ns25586)

	// ZZ: (uint64)(ns25586.ViewNumber)
	_ = dst[ns25587+7]
	dst[ns25587+0] = byte(ns25586.ViewNumber >> 0)
	dst[ns25587+1] = byte(ns25586.ViewNumber >> 8)
	dst[ns25587+2] = byte(ns25586.ViewNumber >> 16)
	dst[ns25587+3] = byte(ns25586.ViewNumber >> 24)
	dst[ns25587+4] = byte(ns25586.ViewNumber >> 32)
	dst[ns25587+5] = byte(ns25586.ViewNumber >> 40)
	dst[ns25587+6] = byte(ns25586.ViewNumber >> 48)
	dst[ns25587+7] = byte(ns25586.ViewNumber >> 56)
	ns25587 += 8

	// ZZ: (uint64)(ns25586.OperationNumber)
	_ = dst[ns25587+7]
	dst[ns25587+0] = byte(ns25586.OperationNumber >> 0)
	dst[ns25587+1] = byte(ns25586.OperationNumber >> 8)
	dst[ns25587+2] = byte(ns25586.OperationNumber >> 16)
	dst[ns25587+3] = byte(ns25586.OperationNumber >> 24)
	dst[ns25587+4] = byte(ns25586.OperationNumber >> 32)
	dst[ns25587+5] = byte(ns25586.OperationNumber >> 40)
	dst[ns25587+6] = byte(ns25586.OperationNumber >> 48)
	dst[ns25587+7] = byte(ns25586.OperationNumber >> 56)
	ns25587 += 8

	return ns25587
}

func (ns25588 *PrepareOK) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; OperationNumber uint64 "protobuf:\"varint,2,opt,name=OperationNumber,proto3\" json:\"OperationNumber,omitempty\""})(ns25588)

	// ZZ: (uint64)(ns25588.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25588.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25588.OperationNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25588.OperationNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}

func (ns25589 *PrepareReject) SizeGOBE() uint64 {
	var ns25590 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; OperationNumber uint64 "protobuf:\"varint,2,opt,name=OperationNumber,proto3\" json:\"OperationNumber,omitempty\""; Error ./replication/vsrproto.Error "protobuf:\"varint,3,opt,name=Error,proto3,enum=vsrproto.Error\" json:\"Error,omitempty\""})(ns25589)

	// ZZ: (uint64)(ns25589.ViewNumber)
	ns25590 += 8

	// ZZ: (uint64)(ns25589.OperationNumber)
	ns25590 += 8

	// ZZ: (./replication/vsrproto.Error)(ns25589.Error)

	// ZZ: (int32)(ns25589.Error)
	ns25590 += 4

	return ns25590
}

func (ns25591 *PrepareReject) MarshalGOBE(dst []byte) uint64 {
	var ns25592 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; OperationNumber uint64 "protobuf:\"varint,2,opt,name=OperationNumber,proto3\" json:\"OperationNumber,omitempty\""; Error ./replication/vsrproto.Error "protobuf:\"varint,3,opt,name=Error,proto3,enum=vsrproto.Error\" json:\"Error,omitempty\""})(ns25591)

	// ZZ: (uint64)(ns25591.ViewNumber)
	_ = dst[ns25592+7]
	dst[ns25592+0] = byte(ns25591.ViewNumber >> 0)
	dst[ns25592+1] = byte(ns25591.ViewNumber >> 8)
	dst[ns25592+2] = byte(ns25591.ViewNumber >> 16)
	dst[ns25592+3] = byte(ns25591.ViewNumber >> 24)
	dst[ns25592+4] = byte(ns25591.ViewNumber >> 32)
	dst[ns25592+5] = byte(ns25591.ViewNumber >> 40)
	dst[ns25592+6] = byte(ns25591.ViewNumber >> 48)
	dst[ns25592+7] = byte(ns25591.ViewNumber >> 56)
	ns25592 += 8

	// ZZ: (uint64)(ns25591.OperationNumber)
	_ = dst[ns25592+7]
	dst[ns25592+0] = byte(ns25591.OperationNumber >> 0)
	dst[ns25592+1] = byte(ns25591.OperationNumber >> 8)
	dst[ns25592+2] = byte(ns25591.OperationNumber >> 16)
	dst[ns25592+3] = byte(ns25591.OperationNumber >> 24)
	dst[ns25592+4] = byte(ns25591.OperationNumber >> 32)
	dst[ns25592+5] = byte(ns25591.OperationNumber >> 40)
	dst[ns25592+6] = byte(ns25591.OperationNumber >> 48)
	dst[ns25592+7] = byte(ns25591.OperationNumber >> 56)
	ns25592 += 8

	// ZZ: (./replication/vsrproto.Error)(ns25591.Error)

	// ZZ: (int32)(ns25591.Error)
	var ns25593 uint32 = uint32(ns25591.Error)
	_ = dst[ns25592+3]
	dst[ns25592+0] = byte(ns25593 >> 0)
	dst[ns25592+1] = byte(ns25593 >> 8)
	dst[ns25592+2] = byte(ns25593 >> 16)
	dst[ns25592+3] = byte(ns25593 >> 24)
	ns25592 += 4

	return ns25592
}

func (ns25594 *PrepareReject) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; OperationNumber uint64 "protobuf:\"varint,2,opt,name=OperationNumber,proto3\" json:\"OperationNumber,omitempty\""; Error ./replication/vsrproto.Error "protobuf:\"varint,3,opt,name=Error,proto3,enum=vsrproto.Error\" json:\"Error,omitempty\""})(ns25594)

	// ZZ: (uint64)(ns25594.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25594.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25594.OperationNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25594.OperationNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (./replication/vsrproto.Error)(ns25594.Error)

	// ZZ: (int32)(ns25594.Error)
	if uint64(len(src)) < offset+4 {
		return
	}
	_ = src[offset+3]
	ns25594.Error = Error(
		uint32(src[offset+0])<<0 | uint32(src[offset+1])<<8 | uint32(src[offset+2])<<16 | uint32(src[offset+3])<<24)
	offset += 4

	ok = true
	return
}

func (ns25595 *Propose) SizeGOBE() uint64 {
	var ns25596 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ClientID uint64 "protobuf:\"varint,1,opt,name=ClientID,proto3\" json:\"ClientID,omitempty\""; SequenceNumber uint64 "protobuf:\"varint,2,opt,name=SequenceNumber,proto3\" json:\"SequenceNumber,omitempty\""; Operation []byte "protobuf:\"bytes,3,opt,name=Operation,proto3\" json:\"Operation,omitempty\""})(ns25595)

	// ZZ: (uint64)(ns25595.ClientID)
	ns25596 += 8

	// ZZ: (uint64)(ns25595.SequenceNumber)
	ns25596 += 8

	// ZZ: ([]byte)(ns25595.Operation)
	ns25596 += 8 + uint64(len(ns25595.Operation))

	return ns25596
}

func (ns25597 *Propose) MarshalGOBE(dst []byte) uint64 {
	var ns25598 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ClientID uint64 "protobuf:\"varint,1,opt,name=ClientID,proto3\" json:\"ClientID,omitempty\""; SequenceNumber uint64 "protobuf:\"varint,2,opt,name=SequenceNumber,proto3\" json:\"SequenceNumber,omitempty\""; Operation []byte "protobuf:\"bytes,3,opt,name=Operation,proto3\" json:\"Operation,omitempty\""})(ns25597)

	// ZZ: (uint64)(ns25597.ClientID)
	_ = dst[ns25598+7]
	dst[ns25598+0] = byte(ns25597.ClientID >> 0)
	dst[ns25598+1] = byte(ns25597.ClientID >> 8)
	dst[ns25598+2] = byte(ns25597.ClientID >> 16)
	dst[ns25598+3] = byte(ns25597.ClientID >> 24)
	dst[ns25598+4] = byte(ns25597.ClientID >> 32)
	dst[ns25598+5] = byte(ns25597.ClientID >> 40)
	dst[ns25598+6] = byte(ns25597.ClientID >> 48)
	dst[ns25598+7] = byte(ns25597.ClientID >> 56)
	ns25598 += 8

	// ZZ: (uint64)(ns25597.SequenceNumber)
	_ = dst[ns25598+7]
	dst[ns25598+0] = byte(ns25597.SequenceNumber >> 0)
	dst[ns25598+1] = byte(ns25597.SequenceNumber >> 8)
	dst[ns25598+2] = byte(ns25597.SequenceNumber >> 16)
	dst[ns25598+3] = byte(ns25597.SequenceNumber >> 24)
	dst[ns25598+4] = byte(ns25597.SequenceNumber >> 32)
	dst[ns25598+5] = byte(ns25597.SequenceNumber >> 40)
	dst[ns25598+6] = byte(ns25597.SequenceNumber >> 48)
	dst[ns25598+7] = byte(ns25597.SequenceNumber >> 56)
	ns25598 += 8

	// ZZ: ([]byte)(ns25597.Operation)
	var ns25599 uint64 = uint64(len(ns25597.Operation))
	_ = dst[ns25598+7]
	dst[ns25598+0] = byte(ns25599 >> 0)
	dst[ns25598+1] = byte(ns25599 >> 8)
	dst[ns25598+2] = byte(ns25599 >> 16)
	dst[ns25598+3] = byte(ns25599 >> 24)
	dst[ns25598+4] = byte(ns25599 >> 32)
	dst[ns25598+5] = byte(ns25599 >> 40)
	dst[ns25598+6] = byte(ns25599 >> 48)
	dst[ns25598+7] = byte(ns25599 >> 56)
	copy(dst[ns25598+8:], ns25597.Operation)
	ns25598 += 8 + ns25599

	return ns25598
}

func (ns25600 *Propose) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ClientID uint64 "protobuf:\"varint,1,opt,name=ClientID,proto3\" json:\"ClientID,omitempty\""; SequenceNumber uint64 "protobuf:\"varint,2,opt,name=SequenceNumber,proto3\" json:\"SequenceNumber,omitempty\""; Operation []byte "protobuf:\"bytes,3,opt,name=Operation,proto3\" json:\"Operation,omitempty\""})(ns25600)

	// ZZ: (uint64)(ns25600.ClientID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25600.ClientID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25600.SequenceNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25600.SequenceNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: ([]byte)(ns25600.Operation)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25601 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25601 {
		return
	}
	ns25600.Operation = src[offset : offset+ns25601]
	offset += ns25601

	ok = true
	return
}

func (ns25602 *ProposeReject) SizeGOBE() uint64 {
	var ns25603 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ClientID uint64 "protobuf:\"varint,1,opt,name=ClientID,proto3\" json:\"ClientID,omitempty\""; SequenceNumber uint64 "protobuf:\"varint,2,opt,name=SequenceNumber,proto3\" json:\"SequenceNumber,omitempty\""; Error ./replication/vsrproto.Error "protobuf:\"varint,3,opt,name=Error,proto3,enum=vsrproto.Error\" json:\"Error,omitempty\""})(ns25602)

	// ZZ: (uint64)(ns25602.ClientID)
	ns25603 += 8

	// ZZ: (uint64)(ns25602.SequenceNumber)
	ns25603 += 8

	// ZZ: (./replication/vsrproto.Error)(ns25602.Error)

	// ZZ: (int32)(ns25602.Error)
	ns25603 += 4

	return ns25603
}

func (ns25604 *ProposeReject) MarshalGOBE(dst []byte) uint64 {
	var ns25605 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ClientID uint64 "protobuf:\"varint,1,opt,name=ClientID,proto3\" json:\"ClientID,omitempty\""; SequenceNumber uint64 "protobuf:\"varint,2,opt,name=SequenceNumber,proto3\" json:\"SequenceNumber,omitempty\""; Error ./replication/vsrproto.Error "protobuf:\"varint,3,opt,name=Error,proto3,enum=vsrproto.Error\" json:\"Error,omitempty\""})(ns25604)

	// ZZ: (uint64)(ns25604.ClientID)
	_ = dst[ns25605+7]
	dst[ns25605+0] = byte(ns25604.ClientID >> 0)
	dst[ns25605+1] = byte(ns25604.ClientID >> 8)
	dst[ns25605+2] = byte(ns25604.ClientID >> 16)
	dst[ns25605+3] = byte(ns25604.ClientID >> 24)
	dst[ns25605+4] = byte(ns25604.ClientID >> 32)
	dst[ns25605+5] = byte(ns25604.ClientID >> 40)
	dst[ns25605+6] = byte(ns25604.ClientID >> 48)
	dst[ns25605+7] = byte(ns25604.ClientID >> 56)
	ns25605 += 8

	// ZZ: (uint64)(ns25604.SequenceNumber)
	_ = dst[ns25605+7]
	dst[ns25605+0] = byte(ns25604.SequenceNumber >> 0)
	dst[ns25605+1] = byte(ns25604.SequenceNumber >> 8)
	dst[ns25605+2] = byte(ns25604.SequenceNumber >> 16)
	dst[ns25605+3] = byte(ns25604.SequenceNumber >> 24)
	dst[ns25605+4] = byte(ns25604.SequenceNumber >> 32)
	dst[ns25605+5] = byte(ns25604.SequenceNumber >> 40)
	dst[ns25605+6] = byte(ns25604.SequenceNumber >> 48)
	dst[ns25605+7] = byte(ns25604.SequenceNumber >> 56)
	ns25605 += 8

	// ZZ: (./replication/vsrproto.Error)(ns25604.Error)

	// ZZ: (int32)(ns25604.Error)
	var ns25606 uint32 = uint32(ns25604.Error)
	_ = dst[ns25605+3]
	dst[ns25605+0] = byte(ns25606 >> 0)
	dst[ns25605+1] = byte(ns25606 >> 8)
	dst[ns25605+2] = byte(ns25606 >> 16)
	dst[ns25605+3] = byte(ns25606 >> 24)
	ns25605 += 4

	return ns25605
}

func (ns25607 *ProposeReject) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ClientID uint64 "protobuf:\"varint,1,opt,name=ClientID,proto3\" json:\"ClientID,omitempty\""; SequenceNumber uint64 "protobuf:\"varint,2,opt,name=SequenceNumber,proto3\" json:\"SequenceNumber,omitempty\""; Error ./replication/vsrproto.Error "protobuf:\"varint,3,opt,name=Error,proto3,enum=vsrproto.Error\" json:\"Error,omitempty\""})(ns25607)

	// ZZ: (uint64)(ns25607.ClientID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25607.ClientID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25607.SequenceNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25607.SequenceNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (./replication/vsrproto.Error)(ns25607.Error)

	// ZZ: (int32)(ns25607.Error)
	if uint64(len(src)) < offset+4 {
		return
	}
	_ = src[offset+3]
	ns25607.Error = Error(
		uint32(src[offset+0])<<0 | uint32(src[offset+1])<<8 | uint32(src[offset+2])<<16 | uint32(src[offset+3])<<24)
	offset += 4

	ok = true
	return
}

func (ns25608 *Response) SizeGOBE() uint64 {
	var ns25609 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ClientID uint64 "protobuf:\"varint,1,opt,name=ClientID,proto3\" json:\"ClientID,omitempty\""; SequenceNumber uint64 "protobuf:\"varint,2,opt,name=SequenceNumber,proto3\" json:\"SequenceNumber,omitempty\""; Result []byte "protobuf:\"bytes,3,opt,name=Result,proto3\" json:\"Result,omitempty\""})(ns25608)

	// ZZ: (uint64)(ns25608.ClientID)
	ns25609 += 8

	// ZZ: (uint64)(ns25608.SequenceNumber)
	ns25609 += 8

	// ZZ: ([]byte)(ns25608.Result)
	ns25609 += 8 + uint64(len(ns25608.Result))

	return ns25609
}

func (ns25610 *Response) MarshalGOBE(dst []byte) uint64 {
	var ns25611 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ClientID uint64 "protobuf:\"varint,1,opt,name=ClientID,proto3\" json:\"ClientID,omitempty\""; SequenceNumber uint64 "protobuf:\"varint,2,opt,name=SequenceNumber,proto3\" json:\"SequenceNumber,omitempty\""; Result []byte "protobuf:\"bytes,3,opt,name=Result,proto3\" json:\"Result,omitempty\""})(ns25610)

	// ZZ: (uint64)(ns25610.ClientID)
	_ = dst[ns25611+7]
	dst[ns25611+0] = byte(ns25610.ClientID >> 0)
	dst[ns25611+1] = byte(ns25610.ClientID >> 8)
	dst[ns25611+2] = byte(ns25610.ClientID >> 16)
	dst[ns25611+3] = byte(ns25610.ClientID >> 24)
	dst[ns25611+4] = byte(ns25610.ClientID >> 32)
	dst[ns25611+5] = byte(ns25610.ClientID >> 40)
	dst[ns25611+6] = byte(ns25610.ClientID >> 48)
	dst[ns25611+7] = byte(ns25610.ClientID >> 56)
	ns25611 += 8

	// ZZ: (uint64)(ns25610.SequenceNumber)
	_ = dst[ns25611+7]
	dst[ns25611+0] = byte(ns25610.SequenceNumber >> 0)
	dst[ns25611+1] = byte(ns25610.SequenceNumber >> 8)
	dst[ns25611+2] = byte(ns25610.SequenceNumber >> 16)
	dst[ns25611+3] = byte(ns25610.SequenceNumber >> 24)
	dst[ns25611+4] = byte(ns25610.SequenceNumber >> 32)
	dst[ns25611+5] = byte(ns25610.SequenceNumber >> 40)
	dst[ns25611+6] = byte(ns25610.SequenceNumber >> 48)
	dst[ns25611+7] = byte(ns25610.SequenceNumber >> 56)
	ns25611 += 8

	// ZZ: ([]byte)(ns25610.Result)
	var ns25612 uint64 = uint64(len(ns25610.Result))
	_ = dst[ns25611+7]
	dst[ns25611+0] = byte(ns25612 >> 0)
	dst[ns25611+1] = byte(ns25612 >> 8)
	dst[ns25611+2] = byte(ns25612 >> 16)
	dst[ns25611+3] = byte(ns25612 >> 24)
	dst[ns25611+4] = byte(ns25612 >> 32)
	dst[ns25611+5] = byte(ns25612 >> 40)
	dst[ns25611+6] = byte(ns25612 >> 48)
	dst[ns25611+7] = byte(ns25612 >> 56)
	copy(dst[ns25611+8:], ns25610.Result)
	ns25611 += 8 + ns25612

	return ns25611
}

func (ns25613 *Response) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ClientID uint64 "protobuf:\"varint,1,opt,name=ClientID,proto3\" json:\"ClientID,omitempty\""; SequenceNumber uint64 "protobuf:\"varint,2,opt,name=SequenceNumber,proto3\" json:\"SequenceNumber,omitempty\""; Result []byte "protobuf:\"bytes,3,opt,name=Result,proto3\" json:\"Result,omitempty\""})(ns25613)

	// ZZ: (uint64)(ns25613.ClientID)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25613.ClientID = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25613.SequenceNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25613.SequenceNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: ([]byte)(ns25613.Result)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25614 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(len(src)) < offset+ns25614 {
		return
	}
	ns25613.Result = src[offset : offset+ns25614]
	offset += ns25614

	ok = true
	return
}

func (ns25615 *StartView) SizeGOBE() uint64 {
	var ns25616 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; EpochNumber uint64 "protobuf:\"varint,2,opt,name=EpochNumber,proto3\" json:\"EpochNumber,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,3,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,4,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""; PrepareOKs []*./replication/vsrproto.PrepareOK "protobuf:\"bytes,5,rep,name=PrepareOKs,proto3\" json:\"PrepareOKs,omitempty\""})(ns25615)

	// ZZ: (uint64)(ns25615.ViewNumber)
	ns25616 += 8

	// ZZ: (uint64)(ns25615.EpochNumber)
	ns25616 += 8

	// ZZ: (uint64)(ns25615.CommitNumber_MAX)
	ns25616 += 8

	// ZZ: (uint64)(ns25615.CommitNumber_MIN)
	ns25616 += 8

	// ZZ: ([]*./replication/vsrproto.PrepareOK)(ns25615.PrepareOKs)
	ns25616 += 8
	for ns25617 := 0; ns25617 < len(ns25615.PrepareOKs); ns25617++ {

		// ZZ: (*./replication/vsrproto.PrepareOK)(ns25615.PrepareOKs[ns25617])
		ns25616 += 1
		if ns25615.PrepareOKs[ns25617] != nil {

			// ZZ: (./replication/vsrproto.PrepareOK)((*ns25615.PrepareOKs[ns25617]))
			ns25616 += (*ns25615.PrepareOKs[ns25617]).SizeGOBE()
		}
	}

	return ns25616
}

func (ns25618 *StartView) MarshalGOBE(dst []byte) uint64 {
	var ns25619 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; EpochNumber uint64 "protobuf:\"varint,2,opt,name=EpochNumber,proto3\" json:\"EpochNumber,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,3,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,4,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""; PrepareOKs []*./replication/vsrproto.PrepareOK "protobuf:\"bytes,5,rep,name=PrepareOKs,proto3\" json:\"PrepareOKs,omitempty\""})(ns25618)

	// ZZ: (uint64)(ns25618.ViewNumber)
	_ = dst[ns25619+7]
	dst[ns25619+0] = byte(ns25618.ViewNumber >> 0)
	dst[ns25619+1] = byte(ns25618.ViewNumber >> 8)
	dst[ns25619+2] = byte(ns25618.ViewNumber >> 16)
	dst[ns25619+3] = byte(ns25618.ViewNumber >> 24)
	dst[ns25619+4] = byte(ns25618.ViewNumber >> 32)
	dst[ns25619+5] = byte(ns25618.ViewNumber >> 40)
	dst[ns25619+6] = byte(ns25618.ViewNumber >> 48)
	dst[ns25619+7] = byte(ns25618.ViewNumber >> 56)
	ns25619 += 8

	// ZZ: (uint64)(ns25618.EpochNumber)
	_ = dst[ns25619+7]
	dst[ns25619+0] = byte(ns25618.EpochNumber >> 0)
	dst[ns25619+1] = byte(ns25618.EpochNumber >> 8)
	dst[ns25619+2] = byte(ns25618.EpochNumber >> 16)
	dst[ns25619+3] = byte(ns25618.EpochNumber >> 24)
	dst[ns25619+4] = byte(ns25618.EpochNumber >> 32)
	dst[ns25619+5] = byte(ns25618.EpochNumber >> 40)
	dst[ns25619+6] = byte(ns25618.EpochNumber >> 48)
	dst[ns25619+7] = byte(ns25618.EpochNumber >> 56)
	ns25619 += 8

	// ZZ: (uint64)(ns25618.CommitNumber_MAX)
	_ = dst[ns25619+7]
	dst[ns25619+0] = byte(ns25618.CommitNumber_MAX >> 0)
	dst[ns25619+1] = byte(ns25618.CommitNumber_MAX >> 8)
	dst[ns25619+2] = byte(ns25618.CommitNumber_MAX >> 16)
	dst[ns25619+3] = byte(ns25618.CommitNumber_MAX >> 24)
	dst[ns25619+4] = byte(ns25618.CommitNumber_MAX >> 32)
	dst[ns25619+5] = byte(ns25618.CommitNumber_MAX >> 40)
	dst[ns25619+6] = byte(ns25618.CommitNumber_MAX >> 48)
	dst[ns25619+7] = byte(ns25618.CommitNumber_MAX >> 56)
	ns25619 += 8

	// ZZ: (uint64)(ns25618.CommitNumber_MIN)
	_ = dst[ns25619+7]
	dst[ns25619+0] = byte(ns25618.CommitNumber_MIN >> 0)
	dst[ns25619+1] = byte(ns25618.CommitNumber_MIN >> 8)
	dst[ns25619+2] = byte(ns25618.CommitNumber_MIN >> 16)
	dst[ns25619+3] = byte(ns25618.CommitNumber_MIN >> 24)
	dst[ns25619+4] = byte(ns25618.CommitNumber_MIN >> 32)
	dst[ns25619+5] = byte(ns25618.CommitNumber_MIN >> 40)
	dst[ns25619+6] = byte(ns25618.CommitNumber_MIN >> 48)
	dst[ns25619+7] = byte(ns25618.CommitNumber_MIN >> 56)
	ns25619 += 8

	// ZZ: ([]*./replication/vsrproto.PrepareOK)(ns25618.PrepareOKs)
	var ns25620 uint64 = uint64(len(ns25618.PrepareOKs))
	_ = dst[ns25619+7]
	dst[ns25619+0] = byte(ns25620 >> 0)
	dst[ns25619+1] = byte(ns25620 >> 8)
	dst[ns25619+2] = byte(ns25620 >> 16)
	dst[ns25619+3] = byte(ns25620 >> 24)
	dst[ns25619+4] = byte(ns25620 >> 32)
	dst[ns25619+5] = byte(ns25620 >> 40)
	dst[ns25619+6] = byte(ns25620 >> 48)
	dst[ns25619+7] = byte(ns25620 >> 56)
	ns25619 += 8

	for ns25621 := 0; ns25621 < len(ns25618.PrepareOKs); ns25621++ {

		// ZZ: (*./replication/vsrproto.PrepareOK)(ns25618.PrepareOKs[ns25621])
		if ns25618.PrepareOKs[ns25621] != nil {
			dst[ns25619] = 1
			ns25619++

			// ZZ: (./replication/vsrproto.PrepareOK)((*ns25618.PrepareOKs[ns25621]))
			ns25619 += (*ns25618.PrepareOKs[ns25621]).MarshalGOBE(dst[ns25619:])
		} else {
			dst[ns25619] = 0
			ns25619++
		}

	}

	return ns25619
}

func (ns25622 *StartView) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""; EpochNumber uint64 "protobuf:\"varint,2,opt,name=EpochNumber,proto3\" json:\"EpochNumber,omitempty\""; CommitNumber_MAX uint64 "protobuf:\"varint,3,opt,name=CommitNumber_MAX,json=CommitNumberMAX,proto3\" json:\"CommitNumber_MAX,omitempty\""; CommitNumber_MIN uint64 "protobuf:\"varint,4,opt,name=CommitNumber_MIN,json=CommitNumberMIN,proto3\" json:\"CommitNumber_MIN,omitempty\""; PrepareOKs []*./replication/vsrproto.PrepareOK "protobuf:\"bytes,5,rep,name=PrepareOKs,proto3\" json:\"PrepareOKs,omitempty\""})(ns25622)

	// ZZ: (uint64)(ns25622.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25622.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25622.EpochNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25622.EpochNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25622.CommitNumber_MAX)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25622.CommitNumber_MAX = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: (uint64)(ns25622.CommitNumber_MIN)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25622.CommitNumber_MIN = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	// ZZ: ([]*./replication/vsrproto.PrepareOK)(ns25622.PrepareOKs)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	var ns25623 uint64 = uint64(src[offset]) | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56
	offset += 8
	if uint64(cap(ns25622.PrepareOKs)) < ns25623 {
		if ns25623*uint64(8) <= 1<<15 {
			ns25622.PrepareOKs = make([]*PrepareOK, ns25623)
			for ns25624 := uint64(0); ns25624 < ns25623; ns25624++ {

				// ZZ: (*./replication/vsrproto.PrepareOK)(ns25622.PrepareOKs[ns25624])
				if uint64(len(src)) < offset+1 {
					return
				}
				if src[offset] == 1 {
					offset++
					if ns25622.PrepareOKs[ns25624] == nil {
						ns25622.PrepareOKs[ns25624] = new(PrepareOK)
					}

					// ZZ: (./replication/vsrproto.PrepareOK)((*ns25622.PrepareOKs[ns25624]))
					ns25626, ns25627 := (*ns25622.PrepareOKs[ns25624]).UnmarshalGOBE(src[offset:])
					offset += ns25626
					if !ns25627 {
						return
					}
				} else {
					offset++
					ns25622.PrepareOKs[ns25624] = nil
				}
			}
		} else {
			// Slice too large, Using Append
			ns25622.PrepareOKs = ns25622.PrepareOKs[:0]
			for ns25624 := uint64(0); ns25624 < ns25623; ns25624++ {
				var ns25625 *PrepareOK

				// ZZ: (*./replication/vsrproto.PrepareOK)(ns25625)
				if uint64(len(src)) < offset+1 {
					return
				}
				if src[offset] == 1 {
					offset++
					if ns25625 == nil {
						ns25625 = new(PrepareOK)
					}

					// ZZ: (./replication/vsrproto.PrepareOK)((*ns25625))
					ns25628, ns25629 := (*ns25625).UnmarshalGOBE(src[offset:])
					offset += ns25628
					if !ns25629 {
						return
					}
				} else {
					offset++
					ns25625 = nil
				}
				ns25622.PrepareOKs = append(ns25622.PrepareOKs, ns25625)
			}
		}
	} else {
		ns25622.PrepareOKs = ns25622.PrepareOKs[:ns25623]
		for ns25624 := uint64(0); ns25624 < ns25623; ns25624++ {

			// ZZ: (*./replication/vsrproto.PrepareOK)(ns25622.PrepareOKs[ns25624])
			if uint64(len(src)) < offset+1 {
				return
			}
			if src[offset] == 1 {
				offset++
				if ns25622.PrepareOKs[ns25624] == nil {
					ns25622.PrepareOKs[ns25624] = new(PrepareOK)
				}

				// ZZ: (./replication/vsrproto.PrepareOK)((*ns25622.PrepareOKs[ns25624]))
				ns25630, ns25631 := (*ns25622.PrepareOKs[ns25624]).UnmarshalGOBE(src[offset:])
				offset += ns25630
				if !ns25631 {
					return
				}
			} else {
				offset++
				ns25622.PrepareOKs[ns25624] = nil
			}
		}
	}

	ok = true
	return
}

func (ns25632 *StartViewChange) SizeGOBE() uint64 {
	var ns25633 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""})(ns25632)

	// ZZ: (uint64)(ns25632.ViewNumber)
	ns25633 += 8

	return ns25633
}

func (ns25634 *StartViewChange) MarshalGOBE(dst []byte) uint64 {
	var ns25635 uint64

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""})(ns25634)

	// ZZ: (uint64)(ns25634.ViewNumber)
	_ = dst[ns25635+7]
	dst[ns25635+0] = byte(ns25634.ViewNumber >> 0)
	dst[ns25635+1] = byte(ns25634.ViewNumber >> 8)
	dst[ns25635+2] = byte(ns25634.ViewNumber >> 16)
	dst[ns25635+3] = byte(ns25634.ViewNumber >> 24)
	dst[ns25635+4] = byte(ns25634.ViewNumber >> 32)
	dst[ns25635+5] = byte(ns25634.ViewNumber >> 40)
	dst[ns25635+6] = byte(ns25634.ViewNumber >> 48)
	dst[ns25635+7] = byte(ns25634.ViewNumber >> 56)
	ns25635 += 8

	return ns25635
}

func (ns25636 *StartViewChange) UnmarshalGOBE(src []byte) (offset uint64, ok bool) {

	// ZZ: (struct{state google.golang.org/protobuf/internal/impl.MessageState; sizeCache int32; unknownFields []byte; ViewNumber uint64 "protobuf:\"varint,1,opt,name=ViewNumber,proto3\" json:\"ViewNumber,omitempty\""})(ns25636)

	// ZZ: (uint64)(ns25636.ViewNumber)
	if uint64(len(src)) < offset+8 {
		return
	}
	_ = src[offset+7]
	ns25636.ViewNumber = uint64(
		uint64(src[offset+0])<<0 | uint64(src[offset+1])<<8 | uint64(src[offset+2])<<16 | uint64(src[offset+3])<<24 | uint64(src[offset+4])<<32 | uint64(src[offset+5])<<40 | uint64(src[offset+6])<<48 | uint64(src[offset+7])<<56)
	offset += 8

	ok = true
	return
}
