// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkJsonArray.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewJsonArray() *JsonArray {
	return &JsonArray{C.CkJsonArray_Create()}
}

func (z *JsonArray) DisposeJsonArray() {
    if z != nil {
	C.CkJsonArray_Dispose(z.ckObj)
	}
}




// property: DebugLogFilePath
// If set to a file path, causes each Chilkat method or property call to
// automatically append it's LastErrorText to the specified log file. The
// information is appended such that if a hang or crash occurs, it is possible to
// see the context in which the problem occurred, as well as a history of all
// Chilkat calls up to the point of the problem. The VerboseLogging property can be
// set to provide more detailed information.
// 
// This property is typically used for debugging the rare cases where a Chilkat
// method call hangs or generates an exception that halts program execution (i.e.
// crashes). A hang or crash should generally never happen. The typical causes of a
// hang are:
//     a timeout related property was set to 0 to explicitly indicate that an
//     infinite timeout is desired,
//     the hang is actually a hang within an event callback (i.e. it is a hang
//     within the application code), or
//     there is an internal problem (bug) in the Chilkat code that causes the hang.
//
func (z *JsonArray) DebugLogFilePath() string {
    return C.GoString(C.CkJsonArray_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *JsonArray) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkJsonArray_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EmitCompact
// If true then the Emit method outputs in the most compact form possible (a
// single-line with no extra whitespace). If false, then emits with whitespace
// and indentation to make the JSON human-readable.
// 
// The default value is true.
//
func (z *JsonArray) EmitCompact() bool {
    v := int(C.CkJsonArray_getEmitCompact(z.ckObj))
    return v != 0
}
// property setter: EmitCompact
func (z *JsonArray) SetEmitCompact(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJsonArray_putEmitCompact(z.ckObj,v)
}

// property: EmitCrlf
// If true then the Emit method uses CRLF line-endings when emitting the
// non-compact (pretty-print) format. If false, then bare-LF's are emitted. (The
// compact format emits to a single line with no end-of-line characters.) Windows
// systems traditionally use CRLF line-endings, whereas Linux, Mac OS X, and other
// systems traditionally use bare-LF line-endings.
// 
// The default value is true.
//
func (z *JsonArray) EmitCrlf() bool {
    v := int(C.CkJsonArray_getEmitCrlf(z.ckObj))
    return v != 0
}
// property setter: EmitCrlf
func (z *JsonArray) SetEmitCrlf(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJsonArray_putEmitCrlf(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *JsonArray) LastErrorHtml() string {
    return C.GoString(C.CkJsonArray_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *JsonArray) LastErrorText() string {
    return C.GoString(C.CkJsonArray_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *JsonArray) LastErrorXml() string {
    return C.GoString(C.CkJsonArray_lastErrorXml(z.ckObj))
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
func (z *JsonArray) LastMethodSuccess() bool {
    v := int(C.CkJsonArray_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *JsonArray) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJsonArray_putLastMethodSuccess(z.ckObj,v)
}

// property: Size
// The number of JSON values in the array.
func (z *JsonArray) Size() int {
    return int(C.CkJsonArray_getSize(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *JsonArray) VerboseLogging() bool {
    v := int(C.CkJsonArray_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *JsonArray) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJsonArray_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *JsonArray) Version() string {
    return C.GoString(C.CkJsonArray_version(z.ckObj))
}
// Inserts a new and empty JSON array member to the position indicated by index. To
// prepend, pass an index of 0. To append, pass an index of -1. Indexing is 0-based
// (the 1st member is at index 0).
func (z *JsonArray) AddArrayAt(arg1 int) bool {

    v := C.CkJsonArray_AddArrayAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Inserts a new boolean member to the position indicated by index. To prepend, pass
// an index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st member
// is at index 0).
func (z *JsonArray) AddBoolAt(arg1 int, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkJsonArray_AddBoolAt(z.ckObj, C.int(arg1), b2)


    return v != 0
}


// Inserts a new integer member to the position indicated by index. To prepend, pass
// an index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st member
// is at index 0).
func (z *JsonArray) AddIntAt(arg1 int, arg2 int) bool {

    v := C.CkJsonArray_AddIntAt(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Inserts a new null member to the position indicated by index. To prepend, pass an
// index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st member is
// at index 0).
func (z *JsonArray) AddNullAt(arg1 int) bool {

    v := C.CkJsonArray_AddNullAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Inserts a new numeric member to the position indicated by index. The numericStr is an
// integer, float, or double already converted to a string in the format desired by
// the application. To prepend, pass an index of 0. To append, pass an index of -1.
// Indexing is 0-based (the 1st member is at index 0).
func (z *JsonArray) AddNumberAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonArray_AddNumberAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Inserts a new and empty JSON object member to the position indicated by index. To
// prepend, pass an index of 0. To append, pass an index of -1. Indexing is 0-based
// (the 1st member is at index 0).
func (z *JsonArray) AddObjectAt(arg1 int) bool {

    v := C.CkJsonArray_AddObjectAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Inserts a copy of a JSON object to the position indicated by index. To prepend,
// pass an index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st
// member is at index 0).
func (z *JsonArray) AddObjectCopyAt(arg1 int, arg2 *JsonObject) bool {

    v := C.CkJsonArray_AddObjectCopyAt(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Inserts a new string at the position indicated by index. To prepend, pass an index
// of 0. To append, pass an index of -1. Indexing is 0-based (the 1st member is at
// index 0).
func (z *JsonArray) AddStringAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonArray_AddStringAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends the array items contained in jarr.
func (z *JsonArray) AppendArrayItems(arg1 *JsonArray) bool {

    v := C.CkJsonArray_AppendArrayItems(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the JSON array that is the value of the Nth array element. Indexing is
// 0-based (the 1st member is at index 0).
func (z *JsonArray) ArrayAt(arg1 int) *JsonArray {

    retObj := C.CkJsonArray_ArrayAt(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &JsonArray{retObj}
}


// Returns the boolean value of the Nth array element. Indexing is 0-based (the 1st
// member is at index 0).
func (z *JsonArray) BoolAt(arg1 int) bool {

    v := C.CkJsonArray_BoolAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Deletes all array elements.
func (z *JsonArray) Clear()  {

    C.CkJsonArray_Clear(z.ckObj)



}


// Fills the dateTime with the date/time string located in the Nth array element.
// Indexing is 0-based (the 1st member is at index 0). Auto-recognizes the
// following date/time string formats: ISO-8061 Timestamp (such as
// "2009-11-04T19:55:41Z"), RFC822 date/time format (such as "Wed, 18 Apr 2018
// 15:51:55 -0400"), or Unix timestamp integers.
func (z *JsonArray) DateAt(arg1 int, arg2 *CkDateTime) bool {

    v := C.CkJsonArray_DateAt(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Deletes the array element at the given index. Indexing is 0-based (the 1st member
// is at index 0).
func (z *JsonArray) DeleteAt(arg1 int) bool {

    v := C.CkJsonArray_DeleteAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Fills the dt with the date/time string located in the Nth array element. If
// bLocal is true, then dt is filled with the local date/time values, otherwise
// it is filled with the UTC/GMT values. Indexing is 0-based (the 1st member is at
// index 0). Auto-recognizes the following date/time string formats: ISO-8061
// Timestamp (such as "2009-11-04T19:55:41Z"), RFC822 date/time format (such as
// "Wed, 18 Apr 2018 15:51:55 -0400"), or Unix timestamp integers.
func (z *JsonArray) DtAt(arg1 int, arg2 bool, arg3 *DtObj) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkJsonArray_DtAt(z.ckObj, C.int(arg1), b2, arg3.ckObj)


    return v != 0
}


// Writes the JSON array (rooted at the caller) and returns as a string.
// 
// Note: To control the compact/non-compact format, and to control the LF/CRLF
// line-endings, set the EmitCompact and EmitCrlf properties.
//
// return a string or nil if failed.
func (z *JsonArray) Emit() *string {

    retStr := C.CkJsonArray_emit(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Writes the JSON array to the sb.
// 
// Note: To control the compact/non-compact format, and to control the LF/CRLF
// line-endings, set the EmitCompact and EmitCrlf properties.
//
func (z *JsonArray) EmitSb(arg1 *StringBuilder) bool {

    v := C.CkJsonArray_EmitSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Return the index of the first object in the array where value of the field at
// name matches value. name is an object member name. The value is a value pattern
// which can use "*" chars to indicate zero or more of any char. If caseSensitive is
// false, then the matching is case insenstive, otherwise it is case sensitive.
// Returns -1 if no matching string was found.
func (z *JsonArray) FindObject(arg1 string, arg2 string, arg3 bool) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkJsonArray_FindObject(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Return the index of the first matching string in the array. The value is a value
// pattern which can use "*" chars to indicate zero or more of any char. If caseSensitive is
// false, then the matching is case insenstive, otherwise it is case sensitive.
// Returns -1 if no matching string was found.
func (z *JsonArray) FindString(arg1 string, arg2 bool) int {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkJsonArray_FindString(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the integer value of the Nth array element. Indexing is 0-based (the 1st
// member is at index 0).
func (z *JsonArray) IntAt(arg1 int) int {

    v := C.CkJsonArray_IntAt(z.ckObj, C.int(arg1))


    return int(v)
}


// Returns the true if the Nth array element is null, otherwise returns false.
// Indexing is 0-based (the 1st member is at index 0).
func (z *JsonArray) IsNullAt(arg1 int) bool {

    v := C.CkJsonArray_IsNullAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Loads a JSON array from a string. A JSON array must begin with a "[" and end
// with a "]".
// 
// Note: The Load method causes the JsonArray to detach and become it's own JSON
// document. It should only be called on new instances of the JsonArray. See the
// example below.
//
func (z *JsonArray) Load(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonArray_Load(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a JSON array from a StringBuilder. A JSON array must begin with a "[" and
// end with a "]".
// 
// Note: The Load method causes the JsonArray to detach and become it's own JSON
// document. It should only be called on new instances of the JsonArray. See the
// example below.
//
func (z *JsonArray) LoadSb(arg1 *StringBuilder) bool {

    v := C.CkJsonArray_LoadSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the JSON object that is the value of the Nth array element. Indexing is
// 0-based (the 1st member is at index 0).
func (z *JsonArray) ObjectAt(arg1 int) *JsonObject {

    retObj := C.CkJsonArray_ObjectAt(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Sets the boolean value of the Nth array element. Indexing is 0-based (the 1st
// member is at index 0).
func (z *JsonArray) SetBoolAt(arg1 int, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkJsonArray_SetBoolAt(z.ckObj, C.int(arg1), b2)


    return v != 0
}


// Sets the integer value of the Nth array element. Indexing is 0-based (the 1st
// member is at index 0).
func (z *JsonArray) SetIntAt(arg1 int, arg2 int) bool {

    v := C.CkJsonArray_SetIntAt(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Sets the Nth array element to the value of null. Indexing is 0-based (the 1st
// member is at index 0).
func (z *JsonArray) SetNullAt(arg1 int) bool {

    v := C.CkJsonArray_SetNullAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Sets the numeric value of the Nth array element. The value is an integer, float,
// or double already converted to a string in the format desired by the
// application. Indexing is 0-based (the 1st member is at index 0).
func (z *JsonArray) SetNumberAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonArray_SetNumberAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the string value of the Nth array element. Indexing is 0-based (the 1st
// member is at index 0).
func (z *JsonArray) SetStringAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonArray_SetStringAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the string value of the Nth array element. Indexing is 0-based (the 1st
// member is at index 0).
// return a string or nil if failed.
func (z *JsonArray) StringAt(arg1 int) *string {

    retStr := C.CkJsonArray_stringAt(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Swaps the items at positions index1 and index2.
func (z *JsonArray) Swap(arg1 int, arg2 int) bool {

    v := C.CkJsonArray_Swap(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Returns the type of data at the given index. Possible return values are:
//     string
//     number
//     object
//     array
//     boolean
//     null
// Returns -1 if no member exists at the given index.
func (z *JsonArray) TypeAt(arg1 int) int {

    v := C.CkJsonArray_TypeAt(z.ckObj, C.int(arg1))


    return int(v)
}


