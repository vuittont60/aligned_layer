package cairo_platinum

/*
#cgo LDFLAGS: operator/cairo_platinum/lib/libcairo_platinum.a
#include "./lib/cairo_platinum.h"
*/
import "C"

func VerifyCairoProof100Bits(proofBytes *[]uint8, realLen uint) bool {
	return C.verify_cairo_proof_ffi_100_bits(proofBytes, realLen)
}
