package golzo

// #cgo LDFLAGS: -llzo2
// #include <lzo/lzoconf.h>
// #include <lzo/lzo1x.h>
// int golzo_init() { return lzo_init(); /* is a macro */ }
// int golzo1x_999_compress ( void *src, int src_len, void *dst, int *dst_len, void *wrkmem ) {
//     /* quick 'n dirty casting of pointers... */
//     return lzo1x_999_compress ( (const lzo_bytep) src, (lzo_uint) src_len,
//                               (lzo_bytep) dst, (lzo_uintp) dst_len,
//                               (lzo_voidp) wrkmem );
// }
import "C"
import "unsafe"

func Lzo_init() int {
	return int(C.golzo_init());
}

func Lzo1x_999_compress(input []byte) []byte {
	wrkmem := make([]byte, 512*1024) // rough guess, should be adapted to LZO consts
	outbuffer := make([]byte, len(input) + len(input)/16 + 64 + 3) // worst case
	output_length := 0
	C.golzo1x_999_compress(
		unsafe.Pointer(&input[0]),
		C.int(len(input)),
		unsafe.Pointer(&outbuffer[0]),
		(*C.int)(unsafe.Pointer(&output_length)),
		unsafe.Pointer(&wrkmem[0]))
	return outbuffer[0:output_length]
}

