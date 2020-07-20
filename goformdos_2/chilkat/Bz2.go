// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkBz2.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewBz2() *Bz2 {
	return &Bz2{C.CkBz2_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Bz2) DisposeBz2() {
    if z != nil {
	C.CkBz2_Dispose(z.ckObj)
	}
}




func (z *Bz2) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkBz2_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Bz2) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkBz2_setExternalProgress(z.ckObj,1)
}

func (z *Bz2) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkBz2_setExternalProgress(z.ckObj,1)
}

func (z *Bz2) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkBz2_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Bz2) AbortCurrent() bool {
    v := int(C.CkBz2_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Bz2) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkBz2_putAbortCurrent(z.ckObj,v)
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
func (z *Bz2) DebugLogFilePath() string {
    return C.GoString(C.CkBz2_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Bz2) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkBz2_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// This is the number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any compress or decompress
// operation prior to completion. If HeartbeatMs is 0, no AbortCheck event
// callbacks will occur.
func (z *Bz2) HeartbeatMs() int {
    return int(C.CkBz2_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Bz2) SetHeartbeatMs(value int) {
    C.CkBz2_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Bz2) LastErrorHtml() string {
    return C.GoString(C.CkBz2_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Bz2) LastErrorText() string {
    return C.GoString(C.CkBz2_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Bz2) LastErrorXml() string {
    return C.GoString(C.CkBz2_lastErrorXml(z.ckObj))
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
func (z *Bz2) LastMethodSuccess() bool {
    v := int(C.CkBz2_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Bz2) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkBz2_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Bz2) VerboseLogging() bool {
    v := int(C.CkBz2_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Bz2) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkBz2_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Bz2) Version() string {
    return C.GoString(C.CkBz2_version(z.ckObj))
}
// Compresses a file to create a BZip2 compressed file (.bz2).
// 
// Note: Both inFilename and toPath should be relative or absolute file paths (not just a
// path to a directory). For example "someDir1/someDir2/myFile.txt" or
// "c:/someDir1/myFile.bz2".
//
func (z *Bz2) CompressFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkBz2_CompressFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CompressFile method
func (z *Bz2) CompressFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkBz2_CompressFileAsync(z.ckObj, cstr1, cstr2)

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


// BZip2 compresses a file to an in-memory image of a .bz2 file.
func (z *Bz2) CompressFileToMem(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkBz2_CompressFileToMem(z.ckObj, cstr1, ckOutBytes)

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
func (z *Bz2) CompressFileToMemAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkBz2_CompressFileToMemAsync(z.ckObj, cstr1)

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


// Compresses in-memory data to an in-memory image of a .bz2 file.
func (z *Bz2) CompressMemory(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkBz2_CompressMemory(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the CompressMemory method
func (z *Bz2) CompressMemoryAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkBz2_CompressMemoryAsync(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// BZip2 compresses and creates a .bz2 file from in-memory data.
func (z *Bz2) CompressMemToFile(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkBz2_CompressMemToFile(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CompressMemToFile method
func (z *Bz2) CompressMemToFileAsync(arg1 []byte, arg2 string, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    hTask := C.CkBz2_CompressMemToFileAsync(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
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


// Loads the caller of the task's async method.
func (z *Bz2) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkBz2_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Unzips a .bz2 file.
func (z *Bz2) UncompressFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkBz2_UncompressFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UncompressFile method
func (z *Bz2) UncompressFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkBz2_UncompressFileAsync(z.ckObj, cstr1, cstr2)

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


// Unzips a .bz2 file directly to memory.
func (z *Bz2) UncompressFileToMem(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkBz2_UncompressFileToMem(z.ckObj, cstr1, ckOutBytes)

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
func (z *Bz2) UncompressFileToMemAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkBz2_UncompressFileToMemAsync(z.ckObj, cstr1)

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


// Unzips from an in-memory image of a .bz2 file directly into memory.
func (z *Bz2) UncompressMemory(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkBz2_UncompressMemory(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the UncompressMemory method
func (z *Bz2) UncompressMemoryAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkBz2_UncompressMemoryAsync(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Unzips from an in-memory image of a .bz2 file to a file.
func (z *Bz2) UncompressMemToFile(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkBz2_UncompressMemToFile(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UncompressMemToFile method
func (z *Bz2) UncompressMemToFileAsync(arg1 []byte, arg2 string, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    hTask := C.CkBz2_UncompressMemToFileAsync(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
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


// Unlocks the component allowing for the full functionality to be used. If a
// purchased unlock code is passed, there is no expiration. Any other string
// automatically begins a fully-functional 30-day trial the first time
// UnlockComponent is called.
func (z *Bz2) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkBz2_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


