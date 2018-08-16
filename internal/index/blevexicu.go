package index

//workaround for issue, https://github.com/blevesearch/blevex/issues/34


// #cgo LDFLAGS: -licuuc -licudata
// #include "unicode/ucnv.h"
import "C"

func init() {
	C.ucnv_setDefaultName(C.CString("UTF-8"))
}

