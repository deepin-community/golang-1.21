// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_solaris.go

package lif

type sockaddrStorage struct {
	Family     uint16
	X_ss_pad1  [6]int8
	X_ss_align float64
	X_ss_pad2  [240]int8
}

const (
	sysLIFC_NOXMIT          = 0x1
	sysLIFC_EXTERNAL_SOURCE = 0x2
	sysLIFC_TEMPORARY       = 0x4
	sysLIFC_ALLZONES        = 0x8
	sysLIFC_UNDER_IPMP      = 0x10
	sysLIFC_ENABLED         = 0x20
)

const (
	sizeofLifnum       = 0xc
	sizeofLifreq       = 0x178
	sizeofLifconf      = 0x18
	sizeofLifIfinfoReq = 0x10
)

type lifnum struct {
	Family    uint16
	Pad_cgo_0 [2]byte
	Flags     int32
	Count     int32
}

type lifreq struct {
	Name   [32]int8
	Lifru1 [4]byte
	Type   uint32
	Lifru  [336]byte
}

type lifconf struct {
	Family    uint16
	Pad_cgo_0 [2]byte
	Flags     int32
	Len       int32
	Pad_cgo_1 [4]byte
	Lifcu     [8]byte
}

type lifIfinfoReq struct {
	Maxhops      uint8
	Pad_cgo_0    [3]byte
	Reachtime    uint32
	Reachretrans uint32
	Maxmtu       uint32
}
