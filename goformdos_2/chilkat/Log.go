// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkLog.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewLog() *Log {
	return &Log{C.CkLog_Create()}
}

func (z *Log) DisposeLog() {
    if z != nil {
	C.CkLog_Dispose(z.ckObj)
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
func (z *Log) DebugLogFilePath() string {
    return C.GoString(C.CkLog_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Log) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkLog_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Log) LastErrorHtml() string {
    return C.GoString(C.CkLog_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Log) LastErrorText() string {
    return C.GoString(C.CkLog_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Log) LastErrorXml() string {
    return C.GoString(C.CkLog_lastErrorXml(z.ckObj))
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
func (z *Log) LastMethodSuccess() bool {
    v := int(C.CkLog_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Log) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkLog_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Log) VerboseLogging() bool {
    v := int(C.CkLog_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Log) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkLog_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Log) Version() string {
    return C.GoString(C.CkLog_version(z.ckObj))
}
// Clears the log. The initialTag is the initial top-level context tag for the new log.
func (z *Log) Clear(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkLog_Clear(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Enters a new context labelled with the given tag. Must be paired with a matching
// call to LeaveContext.
func (z *Log) EnterContext(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkLog_EnterContext(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Leaves the current context. A context that is entered and exited without any
// logging within the context is automatically removed from the log. (To say it
// another way: Empty contexts are automaticallly removed from the log upon leaving
// the context.)
func (z *Log) LeaveContext()  {

    C.CkLog_LeaveContext(z.ckObj)



}


// Adds a tagged message to the log (i.e. a name/value pair).
func (z *Log) LogData(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkLog_LogData(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Logs binary data in base64 format.
func (z *Log) LogDataBase64(arg1 string, arg2 []byte)  {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    C.CkLog_LogDataBase64(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)


}


// Logs binary data in hex format.
func (z *Log) LogDataHex(arg1 string, arg2 []byte)  {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    C.CkLog_LogDataHex(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)


}


// Logs a string, but only up to the 1st maxNumChars characters of the string.
func (z *Log) LogDataMax(arg1 string, arg2 string, arg3 int)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkLog_LogDataMax(z.ckObj, cstr1, cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Logs the current date/time in RFC822 format. If gmt is true, then the GMT/UTC
// time is logged. Otherwise it is the local time.
func (z *Log) LogDateTime(arg1 string, arg2 bool)  {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    C.CkLog_LogDateTime(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))


}


// Logs an error within the current context.
func (z *Log) LogError(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkLog_LogError(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Logs an informational message within the current context.
func (z *Log) LogInfo(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkLog_LogInfo(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Logs an integer.
func (z *Log) LogInt(arg1 string, arg2 int)  {
    cstr1 := C.CString(arg1)

    C.CkLog_LogInt(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Logs a 64-bit integer.
func (z *Log) LogInt64(arg1 string, arg2 int64)  {
    cstr1 := C.CString(arg1)

    C.CkLog_LogInt64(z.ckObj, cstr1, C.__int64(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Logs the current time in HH:MM:SS:mmm format.
func (z *Log) LogTimestamp(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkLog_LogTimestamp(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


