// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkMailboxes.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewMailboxes() *Mailboxes {
	return &Mailboxes{C.CkMailboxes_Create()}
}

func (z *Mailboxes) DisposeMailboxes() {
    if z != nil {
	C.CkMailboxes_Dispose(z.ckObj)
	}
}




// property: Count
// The number of mailboxes in the collection.
// 
// Note: The Mailboxes class is for use with the Chilkat IMAP component.
//
func (z *Mailboxes) Count() int {
    return int(C.CkMailboxes_getCount(z.ckObj))
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
func (z *Mailboxes) LastMethodSuccess() bool {
    v := int(C.CkMailboxes_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Mailboxes) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailboxes_putLastMethodSuccess(z.ckObj,v)
}
// Returns a comma-separated list of flags for the Nth mailbox. For example,
// "\HasNoChildren,\Important".
// return a string or nil if failed.
func (z *Mailboxes) GetFlags(arg1 int) *string {

    retStr := C.CkMailboxes_getFlags(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the index of the mailbox having the specified name.
func (z *Mailboxes) GetMailboxIndex(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkMailboxes_GetMailboxIndex(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// The name of the Nth mailbox in this collection. Indexing begins at 0.
// return a string or nil if failed.
func (z *Mailboxes) GetName(arg1 int) *string {

    retStr := C.CkMailboxes_getName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the name of the Nth flag for the Mth mailbox. The index is the index of
// the mailbox. The flagIndex is the index of the flag.
// return a string or nil if failed.
func (z *Mailboxes) GetNthFlag(arg1 int, arg2 int) *string {

    retStr := C.CkMailboxes_getNthFlag(z.ckObj, C.int(arg1), C.int(arg2))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the number of flags for the Nth mailbox. Returns -1 if the index is out
// of range.
func (z *Mailboxes) GetNumFlags(arg1 int) int {

    v := C.CkMailboxes_GetNumFlags(z.ckObj, C.int(arg1))


    return int(v)
}


// Returns true if the Nth mailbox has the specified flag set. The flag name is
// case insensitive and should begin with a backslash character, such as
// "\Flagged". The index is the index of the Nth mailbox.
func (z *Mailboxes) HasFlag(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkMailboxes_HasFlag(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns true if the Nth mailbox has inferiors (i.e. sub-mailboxes), or if it
// is possible to create child mailboxes in the future.
// 
// Note: the HasNoChildren attribute/flag should not be confused with the IMAP4
// [RFC-2060] defined attribute Noinferiors which indicates that no child mailboxes
// exist now AND none can be created in the future.
//
func (z *Mailboxes) HasInferiors(arg1 int) bool {

    v := C.CkMailboxes_HasInferiors(z.ckObj, C.int(arg1))


    return v != 0
}


// Returns true if the Nth mailbox is marked.
func (z *Mailboxes) IsMarked(arg1 int) bool {

    v := C.CkMailboxes_IsMarked(z.ckObj, C.int(arg1))


    return v != 0
}


// Returns true if the Nth mailbox is selectable.
func (z *Mailboxes) IsSelectable(arg1 int) bool {

    v := C.CkMailboxes_IsSelectable(z.ckObj, C.int(arg1))


    return v != 0
}


// Loads the mailboxes object from a completed asynchronous task.
func (z *Mailboxes) LoadTaskResult(arg1 *Task) bool {

    v := C.CkMailboxes_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


