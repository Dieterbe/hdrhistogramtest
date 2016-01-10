package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Snap) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "LowestTrackableValue":
			z.LowestTrackableValue, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "HighestTrackableValue":
			z.HighestTrackableValue, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "SignificantFigures":
			z.SignificantFigures, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "Counts":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Counts) >= int(xsz) {
				z.Counts = z.Counts[:xsz]
			} else {
				z.Counts = make([]int64, xsz)
			}
			for xvk := range z.Counts {
				z.Counts[xvk], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Snap) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "LowestTrackableValue"
	err = en.Append(0x84, 0xb4, 0x4c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.LowestTrackableValue)
	if err != nil {
		return
	}
	// write "HighestTrackableValue"
	err = en.Append(0xb5, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.HighestTrackableValue)
	if err != nil {
		return
	}
	// write "SignificantFigures"
	err = en.Append(0xb2, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x66, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x67, 0x75, 0x72, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.SignificantFigures)
	if err != nil {
		return
	}
	// write "Counts"
	err = en.Append(0xa6, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Counts)))
	if err != nil {
		return
	}
	for xvk := range z.Counts {
		err = en.WriteInt64(z.Counts[xvk])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Snap) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "LowestTrackableValue"
	o = append(o, 0x84, 0xb4, 0x4c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendInt64(o, z.LowestTrackableValue)
	// string "HighestTrackableValue"
	o = append(o, 0xb5, 0x48, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendInt64(o, z.HighestTrackableValue)
	// string "SignificantFigures"
	o = append(o, 0xb2, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x66, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x46, 0x69, 0x67, 0x75, 0x72, 0x65, 0x73)
	o = msgp.AppendInt64(o, z.SignificantFigures)
	// string "Counts"
	o = append(o, 0xa6, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Counts)))
	for xvk := range z.Counts {
		o = msgp.AppendInt64(o, z.Counts[xvk])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Snap) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "LowestTrackableValue":
			z.LowestTrackableValue, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "HighestTrackableValue":
			z.HighestTrackableValue, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "SignificantFigures":
			z.SignificantFigures, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "Counts":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Counts) >= int(xsz) {
				z.Counts = z.Counts[:xsz]
			} else {
				z.Counts = make([]int64, xsz)
			}
			for xvk := range z.Counts {
				z.Counts[xvk], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Snap) Msgsize() (s int) {
	s = 1 + 21 + msgp.Int64Size + 22 + msgp.Int64Size + 19 + msgp.Int64Size + 7 + msgp.ArrayHeaderSize + (len(z.Counts) * (msgp.Int64Size))
	return
}
