// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkCsv.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCsv() *Csv {
	return &Csv{C.CkCsv_Create()}
}

func (z *Csv) DisposeCsv() {
    if z != nil {
	C.CkCsv_Dispose(z.ckObj)
	}
}




// property: AutoTrim
// If true, then the strings returned by GetCell and GetCellByName are
// auto-trimmed of whitespace from both ends.
func (z *Csv) AutoTrim() bool {
    v := int(C.CkCsv_getAutoTrim(z.ckObj))
    return v != 0
}
// property setter: AutoTrim
func (z *Csv) SetAutoTrim(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsv_putAutoTrim(z.ckObj,v)
}

// property: Crlf
// If true, then CRLF line endings are used when saving the CSV to a file or to a
// string (i.e. for the methods SaveFile, SaveFile2, SaveToString). If false then
// bare LF line-endings are used.
func (z *Csv) Crlf() bool {
    v := int(C.CkCsv_getCrlf(z.ckObj))
    return v != 0
}
// property setter: Crlf
func (z *Csv) SetCrlf(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsv_putCrlf(z.ckObj,v)
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
func (z *Csv) DebugLogFilePath() string {
    return C.GoString(C.CkCsv_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Csv) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsv_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Delimiter
// The character that separates fields in a record. It is a comma by default. If
// the Delimiter property is not explicitly set, the CSV component will detect the
// delimiter when loading a CSV. (Semicolons are typically used in locales where
// the comma is used as a decimal point.)
// 
// Note: If the default comma delimiter is not desired when creating a new CSV,
// make sure to set this property before adding rows/columns to the CSV.
//
func (z *Csv) Delimiter() string {
    return C.GoString(C.CkCsv_delimiter(z.ckObj))
}
// property setter: Delimiter
func (z *Csv) SetDelimiter(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsv_putDelimiter(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EnableQuotes
// If true, then the double-quote characters cause the quoted content, including
// CR's, LF's, and delimiter chars to be treated as normal text when
// reading/writing CSVs. The default is true.
func (z *Csv) EnableQuotes() bool {
    v := int(C.CkCsv_getEnableQuotes(z.ckObj))
    return v != 0
}
// property setter: EnableQuotes
func (z *Csv) SetEnableQuotes(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsv_putEnableQuotes(z.ckObj,v)
}

// property: EscapeBackslash
// If true, then the backslash character is treated as an escape character when
// reading/writing CSVs. The default is false.
func (z *Csv) EscapeBackslash() bool {
    v := int(C.CkCsv_getEscapeBackslash(z.ckObj))
    return v != 0
}
// property setter: EscapeBackslash
func (z *Csv) SetEscapeBackslash(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsv_putEscapeBackslash(z.ckObj,v)
}

// property: HasColumnNames
// Set to true prior to loading a CSV if the 1st record contains column names.
// This allows the CSV parser to correctly load the column names and not treat them
// as data.
func (z *Csv) HasColumnNames() bool {
    v := int(C.CkCsv_getHasColumnNames(z.ckObj))
    return v != 0
}
// property setter: HasColumnNames
func (z *Csv) SetHasColumnNames(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsv_putHasColumnNames(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Csv) LastErrorHtml() string {
    return C.GoString(C.CkCsv_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Csv) LastErrorText() string {
    return C.GoString(C.CkCsv_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Csv) LastErrorXml() string {
    return C.GoString(C.CkCsv_lastErrorXml(z.ckObj))
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
func (z *Csv) LastMethodSuccess() bool {
    v := int(C.CkCsv_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Csv) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsv_putLastMethodSuccess(z.ckObj,v)
}

// property: NumColumns
// The number of columns in the 1st row, which may be the row containing column
// names if HasColumnNames is true.
func (z *Csv) NumColumns() int {
    return int(C.CkCsv_getNumColumns(z.ckObj))
}

// property: NumRows
// The number of data rows. If the CSV has column names, the 1st row is not
// included in the count. Also, empty lines containing only whitespace characters
// that follow the last non-empty row are not included.
func (z *Csv) NumRows() int {
    return int(C.CkCsv_getNumRows(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Csv) VerboseLogging() bool {
    v := int(C.CkCsv_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Csv) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsv_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Csv) Version() string {
    return C.GoString(C.CkCsv_version(z.ckObj))
}
// Deletes the Nth column. (The 1st column is at index 0.)
func (z *Csv) DeleteColumn(arg1 int) bool {

    v := C.CkCsv_DeleteColumn(z.ckObj, C.int(arg1))


    return v != 0
}


// Deletes a column specified by name.
func (z *Csv) DeleteColumnByName(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCsv_DeleteColumnByName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Deletes the entire Nth row. (The 1st row is at index 0.)
func (z *Csv) DeleteRow(arg1 int) bool {

    v := C.CkCsv_DeleteRow(z.ckObj, C.int(arg1))


    return v != 0
}


// Returns the contents of the cell at row, col. Indexing begins at 0. (The
// topmost/leftmost cell is at 0,0)
// return a string or nil if failed.
func (z *Csv) GetCell(arg1 int, arg2 int) *string {

    retStr := C.CkCsv_getCell(z.ckObj, C.int(arg1), C.int(arg2))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The same as GetCell, but the column is specified by name instead of by index.
// return a string or nil if failed.
func (z *Csv) GetCellByName(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkCsv_getCellByName(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the name of the Nth column.
// return a string or nil if failed.
func (z *Csv) GetColumnName(arg1 int) *string {

    retStr := C.CkCsv_getColumnName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the column index for a given column.
func (z *Csv) GetIndex(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkCsv_GetIndex(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the number of columns for a specific row. If the row is larger than the
// number of rows in the CSV, a zero is returned.
func (z *Csv) GetNumCols(arg1 int) int {

    v := C.CkCsv_GetNumCols(z.ckObj, C.int(arg1))


    return int(v)
}


// Loads a CSV from a file. It is assumed that the CSV file contains ANSI
// characters.
func (z *Csv) LoadFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCsv_LoadFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a CSV from a file. The charset specifies the character encoding of the CSV
// file. A list of supported character encodings may be found on this
// page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
func (z *Csv) LoadFile2(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCsv_LoadFile2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a CSV document from an in-memory string variable.
func (z *Csv) LoadFromString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCsv_LoadFromString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Compares the contents of an entire row to a wildcarded match pattern where "*"
// can be used any number of times to match 0 or more of any character. Returns
// true if a match was found, otherwise returns false. If caseSensitive is true, then
// the pattern match is case sensitive, otherwise it is case insensitive.
func (z *Csv) RowMatches(arg1 int, arg2 string, arg3 bool) bool {
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkCsv_RowMatches(z.ckObj, C.int(arg1), cstr2, b3)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Saves a CSV to a file. The output file is written using the ANSI character
// encoding.
func (z *Csv) SaveFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCsv_SaveFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves a CSV to a file. The charset specifies the character encoding to use for the
// CSV file. The text data is converted to this charset when saving. A list of
// supported character encodings may be found on this page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
func (z *Csv) SaveFile2(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCsv_SaveFile2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Writes the entire CSV document to a string variable.
// return a string or nil if failed.
func (z *Csv) SaveToString() *string {

    retStr := C.CkCsv_saveToString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sets the contents for a single cell in the CSV. The content may include any
// characters including CRLF's, double-quotes, and the delimiter character. The
// Save* methods automatically double-quote fields with special chars when saving.
// The Load* methods automatically parse double-quoted and/or escaped fields
// correctly when loading.
func (z *Csv) SetCell(arg1 int, arg2 int, arg3 string) bool {
    cstr3 := C.CString(arg3)

    v := C.CkCsv_SetCell(z.ckObj, C.int(arg1), C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// The same as SetCell, except the column is specified by name instead of by index.
func (z *Csv) SetCellByName(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkCsv_SetCellByName(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Sets the name of the Nth column. The first column is at index 0. This method
// would only return false if an invalid index is passed (such as a negative
// number).
func (z *Csv) SetColumnName(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkCsv_SetColumnName(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sorts the rows in the CSV by the contents of a specific column. If ascending is
// true, the sort is in ascending order, otherwise descending order. If caseSensitive is
// true then the sorting is case sensitive.
func (z *Csv) SortByColumn(arg1 string, arg2 bool, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkCsv_SortByColumn(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sorts the rows in the CSV by the contents of a specific column index. If ascending is
// true, the sort is in ascending order, otherwise descending order. If caseSensitive is
// true then the sorting is case sensitive.
func (z *Csv) SortByColumnIndex(arg1 int, arg2 bool, arg3 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkCsv_SortByColumnIndex(z.ckObj, C.int(arg1), b2, b3)


    return v != 0
}


