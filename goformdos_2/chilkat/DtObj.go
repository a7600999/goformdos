// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkDtObj.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewDtObj() *DtObj {
	return &DtObj{C.CkDtObj_Create()}
}

func (z *DtObj) DisposeDtObj() {
    if z != nil {
	C.CkDtObj_Dispose(z.ckObj)
	}
}




// property: Day
// The day of the month. The valid values for this member are 1 through 31.
func (z *DtObj) Day() int {
    return int(C.CkDtObj_getDay(z.ckObj))
}
// property setter: Day
func (z *DtObj) SetDay(value int) {
    C.CkDtObj_putDay(z.ckObj,C.int(value))
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
func (z *DtObj) DebugLogFilePath() string {
    return C.GoString(C.CkDtObj_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *DtObj) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkDtObj_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Hour
// The hour. The valid values for this member are 0 through 23.
func (z *DtObj) Hour() int {
    return int(C.CkDtObj_getHour(z.ckObj))
}
// property setter: Hour
func (z *DtObj) SetHour(value int) {
    C.CkDtObj_putHour(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *DtObj) LastErrorHtml() string {
    return C.GoString(C.CkDtObj_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *DtObj) LastErrorText() string {
    return C.GoString(C.CkDtObj_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *DtObj) LastErrorXml() string {
    return C.GoString(C.CkDtObj_lastErrorXml(z.ckObj))
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
func (z *DtObj) LastMethodSuccess() bool {
    v := int(C.CkDtObj_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *DtObj) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDtObj_putLastMethodSuccess(z.ckObj,v)
}

// property: Minute
// The minute. The valid values for this member are 0 through 59.
func (z *DtObj) Minute() int {
    return int(C.CkDtObj_getMinute(z.ckObj))
}
// property setter: Minute
func (z *DtObj) SetMinute(value int) {
    C.CkDtObj_putMinute(z.ckObj,C.int(value))
}

// property: Month
// The month. The valid values for this member are 1 through 12 where 1 = January
// and 12 = December.
func (z *DtObj) Month() int {
    return int(C.CkDtObj_getMonth(z.ckObj))
}
// property setter: Month
func (z *DtObj) SetMonth(value int) {
    C.CkDtObj_putMonth(z.ckObj,C.int(value))
}

// property: Second
// The second. The valid values for this member are 0 through 59.
func (z *DtObj) Second() int {
    return int(C.CkDtObj_getSecond(z.ckObj))
}
// property setter: Second
func (z *DtObj) SetSecond(value int) {
    C.CkDtObj_putSecond(z.ckObj,C.int(value))
}

// property: StructTmMonth
// The month. The valid values for this member are 0 through 11 where 0 = January
// and 11 = December.
func (z *DtObj) StructTmMonth() int {
    return int(C.CkDtObj_getStructTmMonth(z.ckObj))
}
// property setter: StructTmMonth
func (z *DtObj) SetStructTmMonth(value int) {
    C.CkDtObj_putStructTmMonth(z.ckObj,C.int(value))
}

// property: StructTmYear
// The year represented as the number of years since 1900.
func (z *DtObj) StructTmYear() int {
    return int(C.CkDtObj_getStructTmYear(z.ckObj))
}
// property setter: StructTmYear
func (z *DtObj) SetStructTmYear(value int) {
    C.CkDtObj_putStructTmYear(z.ckObj,C.int(value))
}

// property: Utc
// true if this is a UTC time, otherwise false if this is a local time.
func (z *DtObj) Utc() bool {
    v := int(C.CkDtObj_getUtc(z.ckObj))
    return v != 0
}
// property setter: Utc
func (z *DtObj) SetUtc(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDtObj_putUtc(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *DtObj) VerboseLogging() bool {
    v := int(C.CkDtObj_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *DtObj) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDtObj_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *DtObj) Version() string {
    return C.GoString(C.CkDtObj_version(z.ckObj))
}

// property: Year
// The year, such as 2012.
func (z *DtObj) Year() int {
    return int(C.CkDtObj_getYear(z.ckObj))
}
// property setter: Year
func (z *DtObj) SetYear(value int) {
    C.CkDtObj_putYear(z.ckObj,C.int(value))
}
// Loads the date/time with a string having the format as produced by the Serialize
// method, which is a string of SPACE separated integers containing (in this order)
// year, month, day, hour, minutes, seconds, and a UTC flag having the value of
// 1/0.
func (z *DtObj) DeSerialize(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkDtObj_DeSerialize(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Serializes the date/time to a us-ascii string that can be imported at a later
// time via the DeSerialize method. The format of the string returned by this
// method is not intended to match any published standard. It is formatted to a
// string with SPACE separated integers containing (in this order) year, month,
// day, hour, minutes, seconds, and a UTC flag having the value of 1/0.
// return a string or nil if failed.
func (z *DtObj) Serialize() *string {

    retStr := C.CkDtObj_serialize(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


