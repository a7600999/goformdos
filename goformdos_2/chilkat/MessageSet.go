// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkMessageSet.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewMessageSet() *MessageSet {
	return &MessageSet{C.CkMessageSet_Create()}
}

func (z *MessageSet) DisposeMessageSet() {
    if z != nil {
	C.CkMessageSet_Dispose(z.ckObj)
	}
}




// property: Count
// The number of message UIDs (or sequence numbers) in this message set.
func (z *MessageSet) Count() int {
    return int(C.CkMessageSet_getCount(z.ckObj))
}

// property: HasUids
// If true then the message set contains UIDs, otherwise it contains sequence
// numbers.
func (z *MessageSet) HasUids() bool {
    v := int(C.CkMessageSet_getHasUids(z.ckObj))
    return v != 0
}
// property setter: HasUids
func (z *MessageSet) SetHasUids(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMessageSet_putHasUids(z.ckObj,v)
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
func (z *MessageSet) LastMethodSuccess() bool {
    v := int(C.CkMessageSet_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *MessageSet) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMessageSet_putLastMethodSuccess(z.ckObj,v)
}
// Returns true if the msgId is contained in the message set.
func (z *MessageSet) ContainsId(arg1 int) bool {

    v := C.CkMessageSet_ContainsId(z.ckObj, C.int(arg1))


    return v != 0
}


// Loads the message set from a compact-string representation. Here are some
// examples:
// 
// Non-Compact String
// 
// Compact String
// 
// 1,2,3,4,5
// 
// 1:5
// 
// 1,2,3,4,5,8,9,10
// 
// 1:5,8:10
// 
// 1,3,4,5,8,9,10
// 
// 1,3:5,8:10
//
func (z *MessageSet) FromCompactString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMessageSet_FromCompactString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the message ID of the Nth message in the set. (indexing begins at 0).
// Returns -1 if the index is out of range.
func (z *MessageSet) GetId(arg1 int) int {

    v := C.CkMessageSet_GetId(z.ckObj, C.int(arg1))


    return int(v)
}


// Inserts a message ID into the set. If the ID already exists, a duplicate is not
// inserted.
func (z *MessageSet) InsertId(arg1 int)  {

    C.CkMessageSet_InsertId(z.ckObj, C.int(arg1))



}


// Loads the message set from a completed asynchronous task.
func (z *MessageSet) LoadTaskResult(arg1 *Task) bool {

    v := C.CkMessageSet_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


// Removes a message ID from the set.
func (z *MessageSet) RemoveId(arg1 int)  {

    C.CkMessageSet_RemoveId(z.ckObj, C.int(arg1))



}


// Returns a string of comma-separated message IDs. (This is the non-compact string
// format.)
// return a string or nil if failed.
func (z *MessageSet) ToCommaSeparatedStr() *string {

    retStr := C.CkMessageSet_toCommaSeparatedStr(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the set of message IDs represented as a compact string. Here are some
// examples:
// 
// Non-Compact String
// 
// Compact String
// 
// 1,2,3,4,5
// 
// 1:5
// 
// 1,2,3,4,5,8,9,10
// 
// 1:5,8:10
// 
// 1,3,4,5,8,9,10
// 
// 1,3:5,8:10
//
// return a string or nil if failed.
func (z *MessageSet) ToCompactString() *string {

    retStr := C.CkMessageSet_toCompactString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


