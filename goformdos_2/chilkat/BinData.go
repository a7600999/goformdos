// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkBinData.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewBinData() *BinData {
	return &BinData{C.CkBinData_Create()}
}

func (z *BinData) DisposeBinData() {
    if z != nil {
	C.CkBinData_Dispose(z.ckObj)
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
func (z *BinData) LastMethodSuccess() bool {
    v := int(C.CkBinData_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *BinData) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkBinData_putLastMethodSuccess(z.ckObj,v)
}

// property: NumBytes
// The number of bytes contained within the object.
func (z *BinData) NumBytes() int {
    return int(C.CkBinData_getNumBytes(z.ckObj))
}
// Appends the contents of another BinData to this object.
func (z *BinData) AppendBd(arg1 *BinData) bool {

    v := C.CkBinData_AppendBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Appends binary data to the current contents, if any.
func (z *BinData) AppendBinary(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkBinData_AppendBinary(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Appends the appropriate BOM (byte order mark), also known as a "preamble", for
// the given charset. If the charset has no defined BOM, then nothing is appended. An
// application would typically call this to append the utf-8, utf-16, or utf-32
// BOM.
func (z *BinData) AppendBom(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkBinData_AppendBom(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends a single byte. The byteValue should be a value from 0 to 255.
func (z *BinData) AppendByte(arg1 int) bool {

    v := C.CkBinData_AppendByte(z.ckObj, C.int(arg1))


    return v != 0
}


// Appends encoded binary data to the current data. The encoding may be "Base64",
// "modBase64", "base64Url", "Base32", "Base58", "QP" (for quoted-printable), "URL"
// (for url-encoding), "Hex", or any of the encodings found atChilkat Binary
// Encodings List
// <http://cknotes.com/chilkat-binary-encoding-list/>.
func (z *BinData) AppendEncoded(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkBinData_AppendEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Decodes the contents of sb and appends the decoded bytes to this object. The
// encoding may be "Base64", "modBase64", "base64Url", "Base32", "Base58", "QP" (for
// quoted-printable), "URL" (for url-encoding), "Hex", or any of the encodings
// found atChilkat Binary Encodings List
// <http://cknotes.com/chilkat-binary-encoding-list/>.
func (z *BinData) AppendEncodedSb(arg1 *StringBuilder, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkBinData_AppendEncodedSb(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends a 16-bit integer (2 bytes). If littleEndian is true, then the integer bytes
// are appended in little-endian byte order, otherwise big-endian byte order is
// used.
func (z *BinData) AppendInt2(arg1 int, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkBinData_AppendInt2(z.ckObj, C.int(arg1), b2)


    return v != 0
}


// Appends a 32-bit integer (4 bytes). If littleEndian is true, then the integer bytes
// are appended in little-endian byte order, otherwise big-endian byte order is
// used.
func (z *BinData) AppendInt4(arg1 int, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkBinData_AppendInt4(z.ckObj, C.int(arg1), b2)


    return v != 0
}


// Appends a string to this object, padded to the fieldLen with NULL or SPACE chars. If
// padWithSpace is true, then SPACE chars are used and the string is not null-terminated.
// If fieldLen is false, then null bytes are used. The charset controls the byte
// representation to use, such as "utf-8".
// 
// Note: This call will always append a total number of bytes equal to fieldLen. If the
// str is longer than fieldLen, the method returns false to indicate failure and
// nothing is appended.
//
func (z *BinData) AppendPadded(arg1 string, arg2 string, arg3 bool, arg4 int) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkBinData_AppendPadded(z.ckObj, cstr1, cstr2, b3, C.int(arg4))

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends the contents of a StringBuilder to this object.
func (z *BinData) AppendSb(arg1 *StringBuilder, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkBinData_AppendSb(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends a string to this object. (This does not append the BOM. If a BOM is
// required, the AppendBom method can be called to append the appropriate BOM.)
func (z *BinData) AppendString(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkBinData_AppendString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Clears the contents.
func (z *BinData) Clear() bool {

    v := C.CkBinData_Clear(z.ckObj)


    return v != 0
}


// Return true if the contents of this object equals the contents of binData.
func (z *BinData) ContentsEqual(arg1 *BinData) bool {

    v := C.CkBinData_ContentsEqual(z.ckObj, arg1.ckObj)


    return v != 0
}


// Retrieves the binary data contained within the object.
func (z *BinData) GetBinary() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkBinData_GetBinary(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Retrieves a chunk of the binary data contained within the object.
func (z *BinData) GetBinaryChunk(arg1 int, arg2 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkBinData_GetBinaryChunk(z.ckObj, C.int(arg1), C.int(arg2), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Retrieves the binary data as an encoded string. The encoding may be "Base64",
// "modBase64", "base64Url", "Base32", "Base58", "QP" (for quoted-printable), "URL"
// (for url-encoding), "Hex", or any of the encodings found atChilkat Binary
// Encodings List
// <http://cknotes.com/chilkat-binary-encoding-list/>.
// return a string or nil if failed.
func (z *BinData) GetEncoded(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkBinData_getEncoded(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Retrieves a chunk of the binary data and returns it in encoded form. The encoding
// may be "Base64", "modBase64", "base64Url", "Base32", "Base58", "QP" (for
// quoted-printable), "URL" (for url-encoding), "Hex", or any of the encodings
// found atChilkat Binary Encodings List
// <http://cknotes.com/chilkat-binary-encoding-list/>.
// return a string or nil if failed.
func (z *BinData) GetEncodedChunk(arg1 int, arg2 int, arg3 string) *string {
    cstr3 := C.CString(arg3)

    retStr := C.CkBinData_getEncodedChunk(z.ckObj, C.int(arg1), C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Writes the encoded data to a StringBuilder. The encoding may be "Base64",
// "modBase64", "base64Url", "Base32", "Base58", "QP" (for quoted-printable), "URL"
// (for url-encoding), "Hex", or any of the encodings found atChilkat Binary
// Encodings List
// <http://cknotes.com/chilkat-binary-encoding-list/>.
func (z *BinData) GetEncodedSb(arg1 string, arg2 *StringBuilder) bool {
    cstr1 := C.CString(arg1)

    v := C.CkBinData_GetEncodedSb(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Interprets the bytes according to charset and returns the string. The charset can be
// "utf-8", "utf-16", "ansi", "iso-8859-*", "windows-125*", or any of the supported
// character encodings listed in the link below.
// return a string or nil if failed.
func (z *BinData) GetString(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkBinData_getString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads binary data and replaces the current contents, if any.
func (z *BinData) LoadBinary(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkBinData_LoadBinary(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Loads binary data from an encoded string, replacing the current contents, if
// any. The encoding may be "Base64", "modBase64", "base64Url", "Base32", "Base58",
// "QP" (for quoted-printable), "URL" (for url-encoding), "Hex", or any of the
// encodings found atChilkat Binary Encodings List
// <http://cknotes.com/chilkat-binary-encoding-list/>.
func (z *BinData) LoadEncoded(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkBinData_LoadEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads data from a file.
func (z *BinData) LoadFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkBinData_LoadFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Removes a chunk of bytes from the binary data.
func (z *BinData) RemoveChunk(arg1 int, arg2 int) bool {

    v := C.CkBinData_RemoveChunk(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Securely clears the contents by writing 0 bytes to the memory prior to
// deallocating the internal memory.
func (z *BinData) SecureClear() bool {

    v := C.CkBinData_SecureClear(z.ckObj)


    return v != 0
}


// Writes the contents to a file.
func (z *BinData) WriteFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkBinData_WriteFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


