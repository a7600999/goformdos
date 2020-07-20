// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkStringArray.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewStringArray() *StringArray {
	return &StringArray{C.CkStringArray_Create()}
}

func (z *StringArray) DisposeStringArray() {
    if z != nil {
	C.CkStringArray_Dispose(z.ckObj)
	}
}




// property: Count
// The number of strings contained within the object's internal array (i.e. ordered
// collection).
// 
// Important: This is an object that contains a collection of strings. Although the
// class/object name includes the word "Array", it should not be confused with an
// array in the sense of the primitive array type used in a given programming
// language.
//
func (z *StringArray) Count() int {
    return int(C.CkStringArray_getCount(z.ckObj))
}

// property: Crlf
// If true, strings are always automatically converted to use CRLF line endings.
// If false, strings are automatically converted to use bare LF line endings.
func (z *StringArray) Crlf() bool {
    v := int(C.CkStringArray_getCrlf(z.ckObj))
    return v != 0
}
// property setter: Crlf
func (z *StringArray) SetCrlf(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStringArray_putCrlf(z.ckObj,v)
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
func (z *StringArray) DebugLogFilePath() string {
    return C.GoString(C.CkStringArray_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *StringArray) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkStringArray_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *StringArray) LastErrorHtml() string {
    return C.GoString(C.CkStringArray_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *StringArray) LastErrorText() string {
    return C.GoString(C.CkStringArray_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *StringArray) LastErrorXml() string {
    return C.GoString(C.CkStringArray_lastErrorXml(z.ckObj))
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
func (z *StringArray) LastMethodSuccess() bool {
    v := int(C.CkStringArray_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *StringArray) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStringArray_putLastMethodSuccess(z.ckObj,v)
}

// property: Length
// The number of strings contained within the internal collection. (Identical to
// the Count property.)
func (z *StringArray) Length() int {
    return int(C.CkStringArray_getLength(z.ckObj))
}

// property: Trim
// If true, whitespace, including carriage-returns and linefeeds, are
// automatically removed from the beginning and end of a string when added to the
// collection.
func (z *StringArray) Trim() bool {
    v := int(C.CkStringArray_getTrim(z.ckObj))
    return v != 0
}
// property setter: Trim
func (z *StringArray) SetTrim(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStringArray_putTrim(z.ckObj,v)
}

// property: Unique
// If true, then duplicates are not allowed. When an attempt is made to insert a
// string that already exists, the duplicate insertion is silently suppressed and
// no error is returned. The default value is false.
func (z *StringArray) Unique() bool {
    v := int(C.CkStringArray_getUnique(z.ckObj))
    return v != 0
}
// property setter: Unique
func (z *StringArray) SetUnique(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStringArray_putUnique(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *StringArray) VerboseLogging() bool {
    v := int(C.CkStringArray_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *StringArray) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStringArray_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *StringArray) Version() string {
    return C.GoString(C.CkStringArray_version(z.ckObj))
}
// Appends a string to the end of the internal ordered collection.
func (z *StringArray) Append(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringArray_Append(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends multiple strings to the end of the internal ordered collection. The encodedStr
// is what is returned from the Serialize method (see below).
func (z *StringArray) AppendSerialized(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringArray_AppendSerialized(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Remove all strings from the internal collection.
func (z *StringArray) Clear()  {

    C.CkStringArray_Clear(z.ckObj)



}


// Returns true if the string is present in the internal collection. The string
// comparisons are case sensitive.
func (z *StringArray) Contains(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringArray_Contains(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Finds the index of the first string equal to findStr. The search begins at startIndex. If
// the string is not found, -1 is returned. The first string in the array is at
// index 0.
func (z *StringArray) Find(arg1 string, arg2 int) int {
    cstr1 := C.CString(arg1)

    v := C.CkStringArray_Find(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Finds the first string that matches the matchPattern. The search begins at startIndex. If the
// string is not found, -1 is returned. The first string in the array is at index
// 0.
// 
// The matchPattern may contain zero or more asterisk chars, each of which matches 0 or
// more of any character.
//
func (z *StringArray) FindFirstMatch(arg1 string, arg2 int) int {
    cstr1 := C.CString(arg1)

    v := C.CkStringArray_FindFirstMatch(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the string at an indexed location within the internal collection. The
// first string is located at index 0.
// return a string or nil if failed.
func (z *StringArray) GetString(arg1 int) *string {

    retStr := C.CkStringArray_getString(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns Nth string's length, in characters. The first string is located at index
// 0.
func (z *StringArray) GetStringLen(arg1 int) int {

    v := C.CkStringArray_GetStringLen(z.ckObj, C.int(arg1))


    return int(v)
}


// Inserts a string into the internal collection at a specified index. Using index
// 0 will insert at the beginning.
func (z *StringArray) InsertAt(arg1 int, arg2 string)  {
    cstr2 := C.CString(arg2)

    C.CkStringArray_InsertAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))


}


// Returns the last string in the internal collection.
// return a string or nil if failed.
func (z *StringArray) LastString() *string {

    retStr := C.CkStringArray_lastString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads strings from a file (one per line) into the internal collection. It is
// assumed the file contains ANSI strings. To load from a file containing non-ANSI
// strings (such as utf-8), call LoadFromFile2 instead.
// 
// Note: This method appends the strings in path to the existing collection of
// strings contained in this object.
//
func (z *StringArray) LoadFromFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringArray_LoadFromFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads strings from a file (one per line) into the internal collection. The charset
// specifies the character encoding (such as utf-8) of the strings contained in the
// file.
// 
// Note: This method appends the strings in path to the existing collection of
// strings contained in this object.
//
func (z *StringArray) LoadFromFile2(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringArray_LoadFromFile2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads strings from an in-memory string (one per line) into the internal
// collection.
// 
// Note: This method appends the strings in str to the existing collection of
// strings contained in this object.
//
func (z *StringArray) LoadFromText(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkStringArray_LoadFromText(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Loads the string collection from a completed asynchronous task.
func (z *StringArray) LoadTaskResult(arg1 *Task) bool {

    v := C.CkStringArray_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the last string and removes it from the internal collection.
// return a string or nil if failed.
func (z *StringArray) Pop() *string {

    retStr := C.CkStringArray_pop(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Adds a string to the beginning of the internal collection.
func (z *StringArray) Prepend(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkStringArray_Prepend(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Removes all strings equal to the string argument from the internal collection.
func (z *StringArray) Remove(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkStringArray_Remove(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Removes the string at a particular index.
func (z *StringArray) RemoveAt(arg1 int) bool {

    v := C.CkStringArray_RemoveAt(z.ckObj, C.int(arg1))


    return v != 0
}


// Replaces the string at a specified index.
func (z *StringArray) ReplaceAt(arg1 int, arg2 string)  {
    cstr2 := C.CString(arg2)

    C.CkStringArray_ReplaceAt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))


}


// Saves the Nth string in the collection to a file.
func (z *StringArray) SaveNthToFile(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkStringArray_SaveNthToFile(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Saves the collection of strings to a file, one string per line. Strings are
// saved using the ANSI charset. (Call SaveToFile2 to specify a charset, such as
// "utf-8")
func (z *StringArray) SaveToFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringArray_SaveToFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the collection of strings to a file, one string per line. Strings are
// saved using the specified charset.
func (z *StringArray) SaveToFile2(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringArray_SaveToFile2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Writes the collection of strings to a single string, one per line (separated by
// CRLF line endings).
// return a string or nil if failed.
func (z *StringArray) SaveToText() *string {

    retStr := C.CkStringArray_saveToText(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns an string which is an encoded representation of all the strings in the
// collection. The string collection can be re-created by calling the
// AppendSerialized method.
// return a string or nil if failed.
func (z *StringArray) Serialize() *string {

    retStr := C.CkStringArray_serialize(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sorts the strings in the collection in ascending or descending order. To sort in
// ascending order, set ascending to true, otherwise set ascending equal to false.
func (z *StringArray) Sort(arg1 bool)  {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    C.CkStringArray_Sort(z.ckObj, b1)



}


// Splits a string at a character or substring boundary and adds each resulting
// string to the internal collection.
func (z *StringArray) SplitAndAppend(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkStringArray_SplitAndAppend(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Subtracts the strings contained within stringArrayObj from the caller's internal
// collection.
func (z *StringArray) Subtract(arg1 *StringArray)  {

    C.CkStringArray_Subtract(z.ckObj, arg1.ckObj)



}


// Performs the union set-operator. The result is that the caller will have a
// string collection that is the union of itself and stringArrayObj.
func (z *StringArray) Union(arg1 *StringArray)  {

    C.CkStringArray_Union(z.ckObj, arg1.ckObj)



}


