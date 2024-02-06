package cairo_platinum

/*
#cgo LDFLAGS: ./lib/libcairo_platinum.a
#include "./lib/cairo_platinum.h"
*/
import "C"
import "unsafe"

const MAX_PROOF_SIZE = 1024 * 1024

func VerifyCairoProof100Bits(proofBytes [MAX_PROOF_SIZE]byte, realLen uint) bool {
	proofPtr := (*C.uchar)(unsafe.Pointer(&proofBytes[0]))
	return (bool)(C.verify_cairo_proof_ffi_100_bits(proofPtr, (C.uint)(realLen)))
}
