// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkServerSentEvent.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewServerSentEvent() *ServerSentEvent {
	return &ServerSentEvent{C.CkServerSentEvent_Create()}
}

func (z *ServerSentEvent) DisposeServerSentEvent() {
    if z != nil {
	C.CkServerSentEvent_Dispose(z.ckObj)
	}
}




// property: Data
// The data for the server-side event. (If the "data" field was empty, then this
// will be empty.)
func (z *ServerSentEvent) Data() string {
    return C.GoString(C.CkServerSentEvent_data(z.ckObj))
}

// property: EventName
// The name of the server-side event. (If the "event" field was not present, then
// this will be empty.)
func (z *ServerSentEvent) EventName() string {
    return C.GoString(C.CkServerSentEvent_eventName(z.ckObj))
}

// property: LastEventId
// The content of the "id" field, if present.
func (z *ServerSentEvent) LastEventId() string {
    return C.GoString(C.CkServerSentEvent_lastEventId(z.ckObj))
}

// property: LastMethodSuccess
// Indicate whether the last method call succeeded or failed. A value of true
// indicates success, a value of false indicates failure. This property is
// automatically set for method calls. It is not modified by property accesses. The
// property is automatically set to indicate success for the following types of
// method calls:
//     Any method that returns a string.
//     Any method returning a Chilkat object, binary bytes, or a date/time.
//     Any method returning a standard boolean status value where success = true
//     and failure = false.
//     Any method returning an integer where failure is defined by a return value
//     less than zero.
// 
// Note: Methods that do not fit the above requirements will always set this
// property equal to true. For example, a method that returns no value (such as a
// "void" in C++) will technically always succeed.
//
func (z *ServerSentEvent) LastMethodSuccess() bool {
    v := int(C.CkServerSentEvent_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *ServerSentEvent) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkServerSentEvent_putLastMethodSuccess(z.ckObj,v)
}

// property: Retry
// The integer value of the "retry" field, if present. Otherwise 0.
func (z *ServerSentEvent) Retry() int {
    return int(C.CkServerSentEvent_getRetry(z.ckObj))
}
// Loads the multi-line event text into this object. For example, the eventText for a
// Firebase event might look like this:
// event: put
// data: {"path": "/c", "data": {"foo": true, "bar": false}}
func (z *ServerSentEvent) LoadEvent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkServerSentEvent_LoadEvent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


