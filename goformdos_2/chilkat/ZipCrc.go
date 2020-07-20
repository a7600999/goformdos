// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkZipCrc.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewZipCrc() *ZipCrc {
	return &ZipCrc{C.CkZipCrc_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *ZipCrc) DisposeZipCrc() {
    if z != nil {
	C.CkZipCrc_Dispose(z.ckObj)
	}
}




func (z *ZipCrc) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkZipCrc_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *ZipCrc) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkZipCrc_setExternalProgress(z.ckObj,1)
}

func (z *ZipCrc) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkZipCrc_setExternalProgress(z.ckObj,1)
}

func (z *ZipCrc) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkZipCrc_setExternalProgress(z.ckObj,1)
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
func (z *ZipCrc) DebugLogFilePath() string {
    return C.GoString(C.CkZipCrc_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *ZipCrc) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkZipCrc_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *ZipCrc) LastErrorHtml() string {
    return C.GoString(C.CkZipCrc_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *ZipCrc) LastErrorText() string {
    return C.GoString(C.CkZipCrc_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *ZipCrc) LastErrorXml() string {
    return C.GoString(C.CkZipCrc_lastErrorXml(z.ckObj))
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
func (z *ZipCrc) LastMethodSuccess() bool {
    v := int(C.CkZipCrc_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *ZipCrc) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZipCrc_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *ZipCrc) VerboseLogging() bool {
    v := int(C.CkZipCrc_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *ZipCrc) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZipCrc_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *ZipCrc) Version() string {
    return C.GoString(C.CkZipCrc_version(z.ckObj))
}
// Provides a way to calculate a CRC by streaming the data a chunk at a time. An
// application would start by calling BeginStream. Then it would add data by
// calling MoreData for each additional chunk. After the last chunk has been
// processed, the EndStream method is called to return the CRC.
func (z *ZipCrc) BeginStream()  {

    C.CkZipCrc_BeginStream(z.ckObj)



}


// Calculates a 32-bit CRC for in-memory byte data. This is the 32-bit CRC that
// would be found in a Zip file header if a file containing the data was added to a
// zip archive. Returns the CRC32 of the data.
func (z *ZipCrc) CalculateCrc(arg1 []byte) uint {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkZipCrc_CalculateCrc(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return uint(v)
}


// Calculates a CRC32 for the bytes contained in bd.
func (z *ZipCrc) CrcBd(arg1 *BinData) uint {

    v := C.CkZipCrc_CrcBd(z.ckObj, arg1.ckObj)


    return uint(v)
}


// Calculates a CRC32 for the string contained in sb. The charset is the byte
// representation to be used for the sb when calculating the CRC32. It can be
// utf-8, utf-16, windows-1252, iso-8859-1, or any of the character encodings
// (charsets) listed at the link below.
func (z *ZipCrc) CrcSb(arg1 *StringBuilder, arg2 string) uint {
    cstr2 := C.CString(arg2)

    v := C.CkZipCrc_CrcSb(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return uint(v)
}


// Calculates a CRC32 for a string. The charset is the byte representation to be used
// for the str when calculating the CRC32. It can be utf-8, utf-16, windows-1252,
// iso-8859-1, or any of the character encodings (charsets) listed at the link
// below.
func (z *ZipCrc) CrcString(arg1 string, arg2 string) uint {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkZipCrc_CrcString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return uint(v)
}


// Finalizes and returns the Zip CRC value calculated by calling BeginStream
// followed by multiple calls to MoreData.
func (z *ZipCrc) EndStream() uint {

    v := C.CkZipCrc_EndStream(z.ckObj)


    return uint(v)
}


// Calculates the CRC32 of a file. The data contained in the file is streamed for
// the calculation to keep the memory footprint small and constant. Returns the
// CRC32 of the file.
func (z *ZipCrc) FileCrc(arg1 string) uint {
    cstr1 := C.CString(arg1)

    v := C.CkZipCrc_FileCrc(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return uint(v)
}


// Asynchronous version of the FileCrc method
func (z *ZipCrc) FileCrcAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkZipCrc_FileCrcAsync(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Loads the caller of the task's async method.
func (z *ZipCrc) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkZipCrc_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds additional data to the CRC currently being calculated. (See BeginStream for
// more information.)
func (z *ZipCrc) MoreData(arg1 []byte)  {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    C.CkZipCrc_MoreData(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)


}


// Converts a 32-bit integer to a hex string.
// return a string or nil if failed.
func (z *ZipCrc) ToHex(arg1 uint) *string {

    retStr := C.CkZipCrc_toHex(z.ckObj, C.ulong(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


