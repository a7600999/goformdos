// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkUnixCompress.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewUnixCompress() *UnixCompress {
	return &UnixCompress{C.CkUnixCompress_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *UnixCompress) DisposeUnixCompress() {
    if z != nil {
	C.CkUnixCompress_Dispose(z.ckObj)
	}
}




func (z *UnixCompress) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkUnixCompress_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *UnixCompress) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkUnixCompress_setExternalProgress(z.ckObj,1)
}

func (z *UnixCompress) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkUnixCompress_setExternalProgress(z.ckObj,1)
}

func (z *UnixCompress) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkUnixCompress_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *UnixCompress) AbortCurrent() bool {
    v := int(C.CkUnixCompress_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *UnixCompress) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUnixCompress_putAbortCurrent(z.ckObj,v)
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
func (z *UnixCompress) DebugLogFilePath() string {
    return C.GoString(C.CkUnixCompress_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *UnixCompress) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkUnixCompress_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// The number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any method call prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
func (z *UnixCompress) HeartbeatMs() int {
    return int(C.CkUnixCompress_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *UnixCompress) SetHeartbeatMs(value int) {
    C.CkUnixCompress_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *UnixCompress) LastErrorHtml() string {
    return C.GoString(C.CkUnixCompress_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *UnixCompress) LastErrorText() string {
    return C.GoString(C.CkUnixCompress_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *UnixCompress) LastErrorXml() string {
    return C.GoString(C.CkUnixCompress_lastErrorXml(z.ckObj))
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
func (z *UnixCompress) LastMethodSuccess() bool {
    v := int(C.CkUnixCompress_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *UnixCompress) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUnixCompress_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *UnixCompress) VerboseLogging() bool {
    v := int(C.CkUnixCompress_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *UnixCompress) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUnixCompress_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *UnixCompress) Version() string {
    return C.GoString(C.CkUnixCompress_version(z.ckObj))
}
// Compresses a file to create a UNIX compressed file (.Z). This compression uses
// the LZW (Lempel-Ziv-Welch) compression algorithm.
func (z *UnixCompress) CompressFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkUnixCompress_CompressFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CompressFile method
func (z *UnixCompress) CompressFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkUnixCompress_CompressFileAsync(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Unix compresses a file to an in-memory image of a .Z file. (This compression
// uses the LZW (Lempel-Ziv-Welch) compression algorithm.)
func (z *UnixCompress) CompressFileToMem(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkUnixCompress_CompressFileToMem(z.ckObj, cstr1, ckOutBytes)

    C.free(unsafe.Pointer(cstr1))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the CompressFileToMem method
func (z *UnixCompress) CompressFileToMemAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkUnixCompress_CompressFileToMemAsync(z.ckObj, cstr1)

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


// Compresses in-memory data to an in-memory image of a .Z file. (This compression
// uses the LZW (Lempel-Ziv-Welch) compression algorithm.)
func (z *UnixCompress) CompressMemory(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkUnixCompress_CompressMemory(z.ckObj, ckbyd1, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Unix compresses and creates a .Z file from in-memory data. (This compression
// uses the LZW (Lempel-Ziv-Welch) compression algorithm.)
func (z *UnixCompress) CompressMemToFile(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkUnixCompress_CompressMemToFile(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Compresses a string to an in-memory image of a .Z file. Prior to compression,
// the string is converted to the character encoding specified by charset. The charset
// can be any one of a large number of character encodings, such as "utf-8",
// "iso-8859-1", "Windows-1252", "ucs-2", etc.
func (z *UnixCompress) CompressString(arg1 string, arg2 string) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkUnixCompress_CompressString(z.ckObj, cstr1, cstr2, ckOutBytes)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Unix compresses and creates a .Z file from string data. The charset specifies the
// character encoding used for the byte representation of the characters when
// compressed. The charset can be any one of a large number of character encodings,
// such as "utf-8", "iso-8859-1", "Windows-1252", "ucs-2", etc.
func (z *UnixCompress) CompressStringToFile(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkUnixCompress_CompressStringToFile(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Returns true if the component has been unlocked. (For ActiveX, returns 1 if
// unlocked.)
func (z *UnixCompress) IsUnlocked() bool {

    v := C.CkUnixCompress_IsUnlocked(z.ckObj)


    return v != 0
}


// Loads the caller of the task's async method.
func (z *UnixCompress) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkUnixCompress_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Uncompresses a .Z file. (This compression uses the LZW (Lempel-Ziv-Welch)
// compression algorithm.)
func (z *UnixCompress) UncompressFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkUnixCompress_UncompressFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UncompressFile method
func (z *UnixCompress) UncompressFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkUnixCompress_UncompressFileAsync(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Uncompresses a .Z file directly to memory. (This compression uses the LZW
// (Lempel-Ziv-Welch) compression algorithm.)
func (z *UnixCompress) UncompressFileToMem(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkUnixCompress_UncompressFileToMem(z.ckObj, cstr1, ckOutBytes)

    C.free(unsafe.Pointer(cstr1))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the UncompressFileToMem method
func (z *UnixCompress) UncompressFileToMemAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkUnixCompress_UncompressFileToMemAsync(z.ckObj, cstr1)

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


// Uncompresses a .Z file that contains a text file. The contents of the text file
// are returned as a string. The character encoding of the text file is specified
// by charset. Typical charsets are "iso-8859-1", "utf-8", "windows-1252",
// "shift_JIS", "big5", etc.
// return a string or nil if failed.
func (z *UnixCompress) UncompressFileToString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkUnixCompress_uncompressFileToString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the UncompressFileToString method
func (z *UnixCompress) UncompressFileToStringAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkUnixCompress_UncompressFileToStringAsync(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Uncompresses from an in-memory image of a .Z file directly into memory. (This
// compression uses the LZW (Lempel-Ziv-Welch) compression algorithm.)
func (z *UnixCompress) UncompressMemory(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkUnixCompress_UncompressMemory(z.ckObj, ckbyd1, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Uncompresses from an in-memory image of a .Z file to a file. (This compression
// uses the LZW (Lempel-Ziv-Welch) compression algorithm.)
func (z *UnixCompress) UncompressMemToFile(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkUnixCompress_UncompressMemToFile(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Uncompresses from an in-memory image of a .Z file directly to a string. The charset
// specifies the character encoding used to interpret the bytes resulting from the
// decompression. The charset can be any one of a large number of character encodings,
// such as "utf-8", "iso-8859-1", "Windows-1252", "ucs-2", etc.
// return a string or nil if failed.
func (z *UnixCompress) UncompressString(arg1 []byte, arg2 string) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    retStr := C.CkUnixCompress_uncompressString(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Unlocks the component allowing for the full functionality to be used.
func (z *UnixCompress) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkUnixCompress_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unpacks a .tar.Z file. The decompress and untar occur in streaming mode. There
// are no temporary files and the memory footprint is constant (and small) while
// untarring.
func (z *UnixCompress) UnTarZ(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkUnixCompress_UnTarZ(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UnTarZ method
func (z *UnixCompress) UnTarZAsync(arg1 string, arg2 string, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkUnixCompress_UnTarZAsync(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


