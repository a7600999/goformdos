// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkSecureString.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewSecureString() *SecureString {
	return &SecureString{C.CkSecureString_Create()}
}

func (z *SecureString) DisposeSecureString() {
    if z != nil {
	C.CkSecureString_Dispose(z.ckObj)
	}
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
func (z *SecureString) LastMethodSuccess() bool {
    v := int(C.CkSecureString_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *SecureString) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSecureString_putLastMethodSuccess(z.ckObj,v)
}

// property: MaintainHash
// If set to the name of a hash algorithm, then a hash of the current string value
// is maintained. This allows for the hash to be verified via the VerifyHash
// method. Possible hash algorithm names are "sha1", "sha256", "sha384", "sha512",
// "md5", "md2", "ripemd160", "ripemd128","ripemd256", and "ripemd320".
func (z *SecureString) MaintainHash() string {
    return C.GoString(C.CkSecureString_maintainHash(z.ckObj))
}
// property setter: MaintainHash
func (z *SecureString) SetMaintainHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkSecureString_putMaintainHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ReadOnly
// Can be set to true to make this secure string read-only (cannot be modified).
func (z *SecureString) ReadOnly() bool {
    v := int(C.CkSecureString_getReadOnly(z.ckObj))
    return v != 0
}
// property setter: ReadOnly
func (z *SecureString) SetReadOnly(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSecureString_putReadOnly(z.ckObj,v)
}
// Returns the clear-text string value.
// return a string or nil if failed.
func (z *SecureString) Access() *string {

    retStr := C.CkSecureString_access(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Appends a clear-text string to this secure string. The in-memory data will be
// decrypted, the string will be appended, and then it will be re-encrypted. Can
// return false if the string has been marked as read-only via the ReadOnly
// property.
func (z *SecureString) Append(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSecureString_Append(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends a clear-text string contained in a StringBuilder to this secure string.
// The in-memory data will be decrypted, the string will be appended, and then it
// will be re-encrypted. Can return false if the string has been marked as
// read-only via the ReadOnly property.
func (z *SecureString) AppendSb(arg1 *StringBuilder) bool {

    v := C.CkSecureString_AppendSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Appends the contents of a secure string to this secure string. The in-memory
// data will be decrypted, the secure string will be appended, and then it will be
// re-encrypted. Can return false if this string has been marked as read-only via
// the ReadOnly property.
func (z *SecureString) AppendSecure(arg1 *SecureString) bool {

    v := C.CkSecureString_AppendSecure(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the hash value for the current value of this secure string. The encoding
// specifies the encoding to be used. It can be any of the binary encoding
// algorithms, such as "base64", "hex", and many more listed atChilkat Binary
// Encodings
// <http://cknotes.com/chilkat-binary-encoding-list/>
// return a string or nil if failed.
func (z *SecureString) HashVal(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSecureString_hashVal(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads the contents of a file into this secure string. The current contents of
// this object are replaced with the new text from the file.
func (z *SecureString) LoadFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSecureString_LoadFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns true if the secStr equals the contents of this secure string.
func (z *SecureString) SecStrEquals(arg1 *SecureString) bool {

    v := C.CkSecureString_SecStrEquals(z.ckObj, arg1.ckObj)


    return v != 0
}


// Verifies the hashVal against the hash value stored for the current value of this
// secure string. The MaintainHash property must've previously been set for this
// secure string to maintain an internal hash. The encoding specifies the encoding of
// the hashVal. It can be any of the binary encoding algorithms, such as "base64",
// "hex", and many more listed atChilkat Binary Encodings
// <http://cknotes.com/chilkat-binary-encoding-list/>
func (z *SecureString) VerifyHash(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSecureString_VerifyHash(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


