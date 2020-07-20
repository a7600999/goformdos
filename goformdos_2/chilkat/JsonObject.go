// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkJsonObject.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewJsonObject() *JsonObject {
	return &JsonObject{C.CkJsonObject_Create()}
}

func (z *JsonObject) DisposeJsonObject() {
    if z != nil {
	C.CkJsonObject_Dispose(z.ckObj)
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
func (z *JsonObject) DebugLogFilePath() string {
    return C.GoString(C.CkJsonObject_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *JsonObject) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkJsonObject_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DelimiterChar
// Sets the delimiter char for JSON paths. The default value is ".". To use
// Firebase style paths, set this property to "/". (This is a string property that
// should have a length of 1 char.)
func (z *JsonObject) DelimiterChar() string {
    return C.GoString(C.CkJsonObject_delimiterChar(z.ckObj))
}
// property setter: DelimiterChar
func (z *JsonObject) SetDelimiterChar(goStr string) {
    cStr := C.CString(goStr)
    C.CkJsonObject_putDelimiterChar(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EmitCompact
// If true then the Emit method outputs in the most compact form possible (a
// single-line with no extra whitespace). If false, then emits with whitespace
// and indentation to make the JSON human-readable.
// 
// The default value is true.
//
func (z *JsonObject) EmitCompact() bool {
    v := int(C.CkJsonObject_getEmitCompact(z.ckObj))
    return v != 0
}
// property setter: EmitCompact
func (z *JsonObject) SetEmitCompact(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJsonObject_putEmitCompact(z.ckObj,v)
}

// property: EmitCrLf
// If true then the Emit method uses CRLF line-endings when emitting the
// non-compact (pretty-print) format. If false, then bare-LF's are emitted. (The
// compact format emits to a single line with no end-of-line characters.) Windows
// systems traditionally use CRLF line-endings, whereas Linux, Mac OS X, and other
// systems traditionally use bare-LF line-endings.
// 
// The default value is true.
//
func (z *JsonObject) EmitCrLf() bool {
    v := int(C.CkJsonObject_getEmitCrLf(z.ckObj))
    return v != 0
}
// property setter: EmitCrLf
func (z *JsonObject) SetEmitCrLf(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJsonObject_putEmitCrLf(z.ckObj,v)
}

// property: I
// The value of the "i" index to be used when evaluating a JSON path.
func (z *JsonObject) I() int {
    return int(C.CkJsonObject_getI(z.ckObj))
}
// property setter: I
func (z *JsonObject) SetI(value int) {
    C.CkJsonObject_putI(z.ckObj,C.int(value))
}

// property: J
// The value of the "j" index to be used when evaluating a JSON path.
func (z *JsonObject) J() int {
    return int(C.CkJsonObject_getJ(z.ckObj))
}
// property setter: J
func (z *JsonObject) SetJ(value int) {
    C.CkJsonObject_putJ(z.ckObj,C.int(value))
}

// property: K
// The value of the "k" index to be used when evaluating a JSON path.
func (z *JsonObject) K() int {
    return int(C.CkJsonObject_getK(z.ckObj))
}
// property setter: K
func (z *JsonObject) SetK(value int) {
    C.CkJsonObject_putK(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *JsonObject) LastErrorHtml() string {
    return C.GoString(C.CkJsonObject_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *JsonObject) LastErrorText() string {
    return C.GoString(C.CkJsonObject_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *JsonObject) LastErrorXml() string {
    return C.GoString(C.CkJsonObject_lastErrorXml(z.ckObj))
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
func (z *JsonObject) LastMethodSuccess() bool {
    v := int(C.CkJsonObject_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *JsonObject) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJsonObject_putLastMethodSuccess(z.ckObj,v)
}

// property: PathPrefix
// A prefix string that is automatically added to the JSON path passed in the first
// argument for other methods (such as StringOf, UpdateString, SetBoolOf,
// SizeOfArray, etc.)
// 
// The default value is the empty string.
//
func (z *JsonObject) PathPrefix() string {
    return C.GoString(C.CkJsonObject_pathPrefix(z.ckObj))
}
// property setter: PathPrefix
func (z *JsonObject) SetPathPrefix(goStr string) {
    cStr := C.CString(goStr)
    C.CkJsonObject_putPathPrefix(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Size
// The number of name/value members in this JSON object.
func (z *JsonObject) Size() int {
    return int(C.CkJsonObject_getSize(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *JsonObject) VerboseLogging() bool {
    v := int(C.CkJsonObject_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *JsonObject) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJsonObject_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *JsonObject) Version() string {
    return C.GoString(C.CkJsonObject_version(z.ckObj))
}
// Inserts a new and empty JSON array member to the position indicated by index. To
// prepend, pass an index of 0. To append, pass an index of -1. Indexing is 0-based
// (the 1st member is at index 0).
func (z *JsonObject) AddArrayAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_AddArrayAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Inserts a copy of a JSON array to the position indicated by index. To prepend,
// pass an index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st
// member is at index 0).
func (z *JsonObject) AddArrayCopyAt(arg1 int, arg2 string, arg3 *JsonArray) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_AddArrayCopyAt(z.ckObj, C.int(arg1), cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Inserts a new boolean member to the position indicated by index. To prepend, pass
// an index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st member
// is at index 0).
func (z *JsonObject) AddBoolAt(arg1 int, arg2 string, arg3 bool) bool {
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkJsonObject_AddBoolAt(z.ckObj, C.int(arg1), cstr2, b3)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Inserts a new integer member to the position indicated by index. To prepend, pass
// an index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st member
// is at index 0).
func (z *JsonObject) AddIntAt(arg1 int, arg2 string, arg3 int) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_AddIntAt(z.ckObj, C.int(arg1), cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Inserts a new null member to the position indicated by index. To prepend, pass an
// index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st member is
// at index 0).
func (z *JsonObject) AddNullAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_AddNullAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Inserts a new numeric member to the position indicated by index. The numericStr is an
// integer, float, or double already converted to a string in the format desired by
// the application. To prepend, pass an index of 0. To append, pass an index of -1.
// Indexing is 0-based (the 1st member is at index 0).
func (z *JsonObject) AddNumberAt(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkJsonObject_AddNumberAt(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Inserts a new and empty JSON object member to the position indicated by index. To
// prepend, pass an index of 0. To append, pass an index of -1. Indexing is 0-based
// (the 1st member is at index 0).
func (z *JsonObject) AddObjectAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_AddObjectAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Inserts a copy of a JSON object to the position indicated by index. To prepend,
// pass an index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st
// member is at index 0).
func (z *JsonObject) AddObjectCopyAt(arg1 int, arg2 string, arg3 *JsonObject) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_AddObjectCopyAt(z.ckObj, C.int(arg1), cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Inserts a new string member to the position indicated by index. To prepend, pass
// an index of 0. To append, pass an index of -1. Indexing is 0-based (the 1st member
// is at index 0).
func (z *JsonObject) AddStringAt(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkJsonObject_AddStringAt(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Appends a new and empty JSON array and returns it.
// 
// Important: The name is the member name, it is not a JSON path.
//
func (z *JsonObject) AppendArray(arg1 string) *JsonArray {
    cstr1 := C.CString(arg1)

    retObj := C.CkJsonObject_AppendArray(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &JsonArray{retObj}
}


// Appends a copy of a JSON array.
// 
// Important: The name is the member name, it is not a JSON path.
//
func (z *JsonObject) AppendArrayCopy(arg1 string, arg2 *JsonArray) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_AppendArrayCopy(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends a new boolean member. (This is the same as passing -1 to the AddBoolAt
// method.)
// 
// Important: The name is the member name. It is not a JSON path. To append (or
// update) using a JSON path, call UpdateBool instead.
//
func (z *JsonObject) AppendBool(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkJsonObject_AppendBool(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends a new integer member. (This is the same as passing an index of -1 to the
// AddIntAt method.)
// 
// Important: The name is the member name. It is not a JSON path. To append (or
// update) using a JSON path, call UpdateInt instead.
//
func (z *JsonObject) AppendInt(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_AppendInt(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends a new and empty JSON object and returns it.
// 
// Important: The name is the member name, it is not a JSON path.
//
func (z *JsonObject) AppendObject(arg1 string) *JsonObject {
    cstr1 := C.CString(arg1)

    retObj := C.CkJsonObject_AppendObject(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Appends a copy of a JSON object.
// 
// Important: The name is the member name, it is not a JSON path.
//
func (z *JsonObject) AppendObjectCopy(arg1 string, arg2 *JsonObject) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_AppendObjectCopy(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends a new string member. (This is the same as passing -1 to the AddStringAt
// method.)
// 
// Important: The name is the member name. It is not a JSON path. To append (or
// update) using a JSON path, call UpdateString instead.
//
func (z *JsonObject) AppendString(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_AppendString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends an array of string values.
// 
// Important: The name is the member name, it is not a JSON path.
//
func (z *JsonObject) AppendStringArray(arg1 string, arg2 *StringTable) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_AppendStringArray(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the JSON array that is the value of the Nth member. Indexing is 0-based
// (the 1st member is at index 0).
func (z *JsonObject) ArrayAt(arg1 int) *JsonArray {

    retObj := C.CkJsonObject_ArrayAt(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &JsonArray{retObj}
}


// Returns the JSON array at the specified jsonPath.
func (z *JsonObject) ArrayOf(arg1 string) *JsonArray {
    cstr1 := C.CString(arg1)

    retObj := C.CkJsonObject_ArrayOf(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &JsonArray{retObj}
}


// Returns the boolean value of the Nth member. Indexing is 0-based (the 1st member
// is at index 0).
func (z *JsonObject) BoolAt(arg1 int) bool {

    v := C.CkJsonObject_BoolAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Returns the boolean at the specified jsonPath.
func (z *JsonObject) BoolOf(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_BoolOf(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends the binary bytes at the specified jsonPath to bd. The encoding indicates the
// encoding of the bytes, such as "base64", "hex", etc.
func (z *JsonObject) BytesOf(arg1 string, arg2 string, arg3 *BinData) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_BytesOf(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Clears the contents of the JSON object. This is the equivalent of calling
// jsonObject.Load("{}")
func (z *JsonObject) Clear()  {

    C.CkJsonObject_Clear(z.ckObj)



}


// Returns a copy of this JSON object.
func (z *JsonObject) Clone() *JsonObject {

    retObj := C.CkJsonObject_Clone(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Fills the dateTime with the date/time string located at jsonPath. Auto-recognizes the
// following date/time string formats: ISO-8061 Timestamp (such as
// "2009-11-04T19:55:41Z"), RFC822 date/time format (such as "Wed, 18 Apr 2018
// 15:51:55 -0400"), or Unix timestamp integers.
func (z *JsonObject) DateOf(arg1 string, arg2 *CkDateTime) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_DateOf(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Deletes the member at having the name specified by name. Note: The name is not a
// tag path. It is the name of a member of this JSON object.
func (z *JsonObject) Delete(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_Delete(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Deletes the member at index index. Indexing is 0-based (the 1st member is at
// index 0).
func (z *JsonObject) DeleteAt(arg1 int) bool {

    v := C.CkJsonObject_DeleteAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Fills the dt with the date/time string located at jsonPath. If bLocal is true,
// then dt is filled with the local date/time values, otherwise it is filled with
// the UTC/GMT values. Auto-recognizes the following date/time string formats:
// ISO-8061 Timestamp (such as "2009-11-04T19:55:41Z"), RFC822 date/time format
// (such as "Wed, 18 Apr 2018 15:51:55 -0400"), or Unix timestamp integers.
func (z *JsonObject) DtOf(arg1 string, arg2 bool, arg3 *DtObj) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkJsonObject_DtOf(z.ckObj, cstr1, b2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Writes the JSON document (rooted at the caller) and returns as a string.
// return a string or nil if failed.
func (z *JsonObject) Emit() *string {

    retStr := C.CkJsonObject_emit(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Emits (appends) to the contents of bd.
func (z *JsonObject) EmitBd(arg1 *BinData) bool {

    v := C.CkJsonObject_EmitBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Appends the JSON to a StringBuilder object.
func (z *JsonObject) EmitSb(arg1 *StringBuilder) bool {

    v := C.CkJsonObject_EmitSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Emits the JSON document with variable substitutions applied. If omitEmpty is true,
// then members having empty strings or empty arrays are omitted.
// return a string or nil if failed.
func (z *JsonObject) EmitWithSubs(arg1 *Hashtable, arg2 bool) *string {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkJsonObject_emitWithSubs(z.ckObj, arg1.ckObj, b2)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Recursively searches the JSON subtree rooted at the caller's node for a JSON
// object containing a member having a specified name. (If the caller is the root
// node of the entire JSON document, then the entire JSON document is searched.)
// Returns the first match or _NULL_ if not found.
func (z *JsonObject) FindObjectWithMember(arg1 string) *JsonObject {
    cstr1 := C.CString(arg1)

    retObj := C.CkJsonObject_FindObjectWithMember(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Finds a JSON record in an array where a particular field equals or matches a
// value pattern. Reviewing the example below is the best way to understand this
// function.
func (z *JsonObject) FindRecord(arg1 string, arg2 string, arg3 string, arg4 bool) *JsonObject {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    retObj := C.CkJsonObject_FindRecord(z.ckObj, cstr1, cstr2, cstr3, b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Finds a JSON value in an record array where a particular field equals or matches
// a value pattern. Reviewing the example below is the best way to understand this
// function.
// return a string or nil if failed.
func (z *JsonObject) FindRecordString(arg1 string, arg2 string, arg3 string, arg4 bool, arg5 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    cstr5 := C.CString(arg5)

    retStr := C.CkJsonObject_findRecordString(z.ckObj, cstr1, cstr2, cstr3, b4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr5))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Applies a Firebase event to the JSON. The data contains JSON having a format
// such as
// {"path": "/", "data": {"a": 1, "b": 2}}
// The name should be "put" or "patch".
func (z *JsonObject) FirebaseApplyEvent(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_FirebaseApplyEvent(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// For each key in the jsonData, update (or add) the corresponding key in the JSON at
// the given jsonPath. The jsonPath is relative to this JSON object. (This is effectively
// applying a Firebase patch event.)
func (z *JsonObject) FirebasePatch(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_FirebasePatch(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Inserts or replaces the value at the jsonPath. The value can contain JSON text, an
// integer (in decimal string format), a boolean (true/false), the keyword "null",
// or a quoted string.
// 
// The jsonPath is relative to this JSON object. (This is effectively applying a
// Firebase put event.)
//
func (z *JsonObject) FirebasePut(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_FirebasePut(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the root of the JSON document. The root can be obtained from any JSON
// object within the JSON document. The entire JSON document remains in memory as
// long as at least one JSON object is referenced by the application. When the last
// reference is removed, the entire JSON document is automatically dellocated.
func (z *JsonObject) GetDocRoot() *JsonObject {

    retObj := C.CkJsonObject_GetDocRoot(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Returns true if the item at the jsonPath exists.
func (z *JsonObject) HasMember(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_HasMember(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the index of the member having the given name. Returns -1 if the name is
// not found.
func (z *JsonObject) IndexOf(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_IndexOf(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the integer value of the Nth member. Indexing is 0-based (the 1st member
// is at index 0).
func (z *JsonObject) IntAt(arg1 int) int {

    v := C.CkJsonObject_IntAt(z.ckObj, C.int(arg1))


    return int(v)
}


// Returns the integer at the specified jsonPath.
func (z *JsonObject) IntOf(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_IntOf(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the boolean value of the member having the specified index.
func (z *JsonObject) IsNullAt(arg1 int) bool {

    v := C.CkJsonObject_IsNullAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Returns true if the value at the specified jsonPath is null. Otherwise returns
// false.
func (z *JsonObject) IsNullOf(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_IsNullOf(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the type of data at the given jsonPath. Possible return values are:
//     1 - string
//     2- number
//     3- object
//     4- array
//     5- boolean
//     6- null
// Returns -1 if no member exists at the given jsonPath.
func (z *JsonObject) JsonTypeOf(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_JsonTypeOf(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Parses a JSON string and loads it into this JSON object to provide DOM-style
// access.
func (z *JsonObject) Load(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_Load(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the contents of bd.
func (z *JsonObject) LoadBd(arg1 *BinData) bool {

    v := C.CkJsonObject_LoadBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads a JSON file into this JSON object. The path is the file path to the JSON
// file.
func (z *JsonObject) LoadFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_LoadFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads this JSON object from a predefined document having a specified name.
func (z *JsonObject) LoadPredefined(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_LoadPredefined(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads JSON from the contents of a StringBuilder object.
func (z *JsonObject) LoadSb(arg1 *StringBuilder) bool {

    v := C.CkJsonObject_LoadSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Move the member from fromIndex to toIndex. If toIndex equals -1, then moves the member at
// position fromIndex to the last position. Set toIndex = 0 to move a member to the first
// position.
func (z *JsonObject) MoveMember(arg1 int, arg2 int) bool {

    v := C.CkJsonObject_MoveMember(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Returns the name of the Nth member. Indexing is 0-based (the 1st member is at
// index 0).
// return a string or nil if failed.
func (z *JsonObject) NameAt(arg1 int) *string {

    retStr := C.CkJsonObject_nameAt(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the JSON object that is the value of the Nth member. Indexing is 0-based
// (the 1st member is at index 0).
func (z *JsonObject) ObjectAt(arg1 int) *JsonObject {

    retObj := C.CkJsonObject_ObjectAt(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Returns the JSON object at the specified jsonPath.
func (z *JsonObject) ObjectOf(arg1 string) *JsonObject {
    cstr1 := C.CString(arg1)

    retObj := C.CkJsonObject_ObjectOf(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Adds or replaces this JSON to an internal global set of predefined JSON
// documents that can be subsequently loaded by name.
func (z *JsonObject) Predefine(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_Predefine(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Renames the member named oldName to newName.
func (z *JsonObject) Rename(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_Rename(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Renames the member at index to name.
func (z *JsonObject) RenameAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_RenameAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the boolean value of the Nth member. Indexing is 0-based (the 1st member is
// at index 0).
func (z *JsonObject) SetBoolAt(arg1 int, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkJsonObject_SetBoolAt(z.ckObj, C.int(arg1), b2)


    return v != 0
}


// Sets the boolean value at the specified jsonPath.
func (z *JsonObject) SetBoolOf(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkJsonObject_SetBoolOf(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the integer value of the Nth member. Indexing is 0-based (the 1st member is
// at index 0).
func (z *JsonObject) SetIntAt(arg1 int, arg2 int) bool {

    v := C.CkJsonObject_SetIntAt(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Sets the integer at the specified jsonPath.
func (z *JsonObject) SetIntOf(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_SetIntOf(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the value of the Nth member to null. Indexing is 0-based (the 1st member is
// at index 0).
func (z *JsonObject) SetNullAt(arg1 int) bool {

    v := C.CkJsonObject_SetNullAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Sets the value at the specified jsonPath to null.
func (z *JsonObject) SetNullOf(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_SetNullOf(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the numeric value of the Nth member. The value is an integer, float, or
// double already converted to a string in the format desired by the application.
// Indexing is 0-based (the 1st member is at index 0).
func (z *JsonObject) SetNumberAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_SetNumberAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the numeric value at the specified jsonPath. The value is an integer, float, or
// double already converted to a string in the format desired by the application.
func (z *JsonObject) SetNumberOf(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_SetNumberOf(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the string value of the Nth member. Indexing is 0-based (the 1st member is
// at index 0).
func (z *JsonObject) SetStringAt(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_SetStringAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the string value at the specified jsonPath.
func (z *JsonObject) SetStringOf(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_SetStringOf(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the size of the array at the given jsonPath. Returns -1 if the jsonPath does not
// evaluate to an existent JSON array.
func (z *JsonObject) SizeOfArray(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_SizeOfArray(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the string value of the Nth member. Indexing is 0-based (the 1st member
// is at index 0).
// return a string or nil if failed.
func (z *JsonObject) StringAt(arg1 int) *string {

    retStr := C.CkJsonObject_stringAt(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the string value at the specified jsonPath.
// return a string or nil if failed.
func (z *JsonObject) StringOf(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkJsonObject_stringOf(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Appends the string value at the specified jsonPath to sb.
func (z *JsonObject) StringOfSb(arg1 string, arg2 *StringBuilder) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_StringOfSb(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Swaps the positions of members at index1 and index2.
func (z *JsonObject) Swap(arg1 int, arg2 int) bool {

    v := C.CkJsonObject_Swap(z.ckObj, C.int(arg1), C.int(arg2))


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
func (z *JsonObject) TypeAt(arg1 int) int {

    v := C.CkJsonObject_TypeAt(z.ckObj, C.int(arg1))


    return int(v)
}


// Updates or appends a new string member with the encoded contents of bd. If the
// full path specified by jsonPath does not exist, it is automatically created as
// needed. The bytes contained in bd are encoded according to encoding (such as
// "base64", "hex", etc.)
func (z *JsonObject) UpdateBd(arg1 string, arg2 string, arg3 *BinData) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_UpdateBd(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Updates or appends a new boolean member. If the full path specified by jsonPath does
// not exist, it is automatically created as needed.
func (z *JsonObject) UpdateBool(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkJsonObject_UpdateBool(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Updates or appends a new integer member. If the full path specified by jsonPath does
// not exist, it is automatically created as needed.
func (z *JsonObject) UpdateInt(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_UpdateInt(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Updates or appends a new and empty array at the jsonPath. If the full path specified
// by jsonPath does not exist, it is automatically created as needed.
func (z *JsonObject) UpdateNewArray(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_UpdateNewArray(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Updates or appends a new and empty array at the jsonPath. If the full path specified
// by jsonPath does not exist, it is automatically created as needed.
func (z *JsonObject) UpdateNewObject(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_UpdateNewObject(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Updates or appends a null member. If the full path specified by jsonPath does not
// exist, it is automatically created as needed.
func (z *JsonObject) UpdateNull(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_UpdateNull(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Updates or appends a new numeric member. If the full path specified by jsonPath does
// not exist, it is automatically created as needed.
func (z *JsonObject) UpdateNumber(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_UpdateNumber(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Updates or appends a new string member with the contents of sb. If the full
// path specified by jsonPath does not exist, it is automatically created as needed.
func (z *JsonObject) UpdateSb(arg1 string, arg2 *StringBuilder) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_UpdateSb(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Updates or appends a new string member. If the full path specified by jsonPath does
// not exist, it is automatically created as needed.
// 
// Important: Prior to version 9.5.0.68, the string passed in to this method did
// not get properly JSON escaped. This could cause problems if non-us-ascii chars
// are present, or if certain special chars such as quotes, CR's, or LF's are
// present. Version 9.5.0.68 fixes the problem.
//
func (z *JsonObject) UpdateString(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJsonObject_UpdateString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Saves the JSON to a file.
func (z *JsonObject) WriteFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJsonObject_WriteFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


