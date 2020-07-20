// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkStringTable.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewStringTable() *StringTable {
	return &StringTable{C.CkStringTable_Create()}
}

func (z *StringTable) DisposeStringTable() {
    if z != nil {
	C.CkStringTable_Dispose(z.ckObj)
	}
}




// property: Count
// The number of strings in the table.
func (z *StringTable) Count() int {
    return int(C.CkStringTable_getCount(z.ckObj))
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
func (z *StringTable) DebugLogFilePath() string {
    return C.GoString(C.CkStringTable_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *StringTable) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkStringTable_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *StringTable) LastErrorHtml() string {
    return C.GoString(C.CkStringTable_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *StringTable) LastErrorText() string {
    return C.GoString(C.CkStringTable_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *StringTable) LastErrorXml() string {
    return C.GoString(C.CkStringTable_lastErrorXml(z.ckObj))
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
func (z *StringTable) LastMethodSuccess() bool {
    v := int(C.CkStringTable_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *StringTable) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStringTable_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *StringTable) VerboseLogging() bool {
    v := int(C.CkStringTable_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *StringTable) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStringTable_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *StringTable) Version() string {
    return C.GoString(C.CkStringTable_version(z.ckObj))
}
// Appends a string to the table.
func (z *StringTable) Append(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringTable_Append(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends strings, one per line, from a file. Each line in the path should be no
// longer than the length specified in maxLineLen. The charset indicates the character
// encoding of the contents of the file, such as "utf-8", "iso-8859-1",
// "Shift_JIS", etc.
func (z *StringTable) AppendFromFile(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkStringTable_AppendFromFile(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Appends strings, one per line, from the contents of a StringBuilder object.
func (z *StringTable) AppendFromSb(arg1 *StringBuilder) bool {

    v := C.CkStringTable_AppendFromSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Removes all the strings from the table.
func (z *StringTable) Clear()  {

    C.CkStringTable_Clear(z.ckObj)



}


// Return the index of the first string in the table containing substr. Begins
// searching strings starting at startIndex. If caseSensitive is true, then the search is case
// sensitive. If caseSensitive is false then the search is case insensitive. Returns -1 if
// the substr is not found.
func (z *StringTable) FindSubstring(arg1 int, arg2 string, arg3 bool) int {
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkStringTable_FindSubstring(z.ckObj, C.int(arg1), cstr2, b3)

    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Returns the Nth string in the table, converted to an integer value. The index is
// 0-based. (The first string is at index 0.) Returns -1 if no string is found at
// the specified index. Returns 0 if the string at the specified index exist, but
// is not an integer.
func (z *StringTable) IntAt(arg1 int) int {

    v := C.CkStringTable_IntAt(z.ckObj, C.int(arg1))


    return int(v)
}


// Saves the string table to a file. The charset is the character encoding to use,
// such as "utf-8", "iso-8859-1", "windows-1252", "Shift_JIS", "gb2312", etc. If
// bCrlf is true, then CRLF line endings are used, otherwise LF-only line endings
// are used.
func (z *StringTable) SaveToFile(arg1 string, arg2 bool, arg3 string) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    v := C.CkStringTable_SaveToFile(z.ckObj, cstr1, b2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Splits a string into parts based on a single character delimiterChar. If exceptDoubleQuoted is true,
// then the delimiter char found between double quotes is not treated as a
// delimiter. If exceptEscaped is true, then an escaped (with a backslash) delimiter char
// is not treated as a delimiter.
func (z *StringTable) SplitAndAppend(arg1 string, arg2 string, arg3 bool, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkStringTable_SplitAndAppend(z.ckObj, cstr1, cstr2, b3, b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the Nth string in the table. The index is 0-based. (The first string is
// at index 0.)
// return a string or nil if failed.
func (z *StringTable) StringAt(arg1 int) *string {

    retStr := C.CkStringTable_stringAt(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


