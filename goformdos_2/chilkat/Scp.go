// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkScp.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewScp() *Scp {
	return &Scp{C.CkScp_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Scp) DisposeScp() {
    if z != nil {
	C.CkScp_Dispose(z.ckObj)
	}
}




func (z *Scp) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkScp_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Scp) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkScp_setExternalProgress(z.ckObj,1)
}

func (z *Scp) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkScp_setExternalProgress(z.ckObj,1)
}

func (z *Scp) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkScp_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Scp) AbortCurrent() bool {
    v := int(C.CkScp_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Scp) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkScp_putAbortCurrent(z.ckObj,v)
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
func (z *Scp) DebugLogFilePath() string {
    return C.GoString(C.CkScp_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Scp) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkScp_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// This is the number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any SSH operation prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
func (z *Scp) HeartbeatMs() int {
    return int(C.CkScp_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Scp) SetHeartbeatMs(value int) {
    C.CkScp_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Scp) LastErrorHtml() string {
    return C.GoString(C.CkScp_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Scp) LastErrorText() string {
    return C.GoString(C.CkScp_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Scp) LastErrorXml() string {
    return C.GoString(C.CkScp_lastErrorXml(z.ckObj))
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
func (z *Scp) LastMethodSuccess() bool {
    v := int(C.CkScp_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Scp) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkScp_putLastMethodSuccess(z.ckObj,v)
}

// property: PercentDoneScale
// This property is only valid in programming environment and languages that allow
// for event callbacks.
// 
// Sets the value to be defined as 100% complete for the purpose of PercentDone
// event callbacks. The defaut value of 100 means that at most 100 event
// PercentDone callbacks will occur in a method that (1) is event enabled and (2)
// is such that it is possible to measure progress as a percentage completed. This
// property may be set to larger numbers to get more fine-grained PercentDone
// callbacks. For example, setting this property equal to 1000 will provide
// callbacks with .1 percent granularity. For example, a value of 453 would
// indicate 45.3% competed. This property is clamped to a minimum value of 10, and
// a maximum value of 100000.
//
func (z *Scp) PercentDoneScale() int {
    return int(C.CkScp_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *Scp) SetPercentDoneScale(value int) {
    C.CkScp_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: SendEnv
// A JSON string specifying environment variables that are to be set for each SCP
// upload or download. For example:
// {
//     "LCS_PASSWORD": "myPassword",
//     "SOME_ENV_VAR": "some_value",
// ...
// }
func (z *Scp) SendEnv() string {
    return C.GoString(C.CkScp_sendEnv(z.ckObj))
}
// property setter: SendEnv
func (z *Scp) SetSendEnv(goStr string) {
    cStr := C.CString(goStr)
    C.CkScp_putSendEnv(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncedFiles
// The paths of the files uploaded or downloaded in the last call to SyncUploadTree
// or SyncDownloadTree. The paths are listed one per line. In both cases (for
// upload and download) each line contains the full local file path (not the remote
// path).
func (z *Scp) SyncedFiles() string {
    return C.GoString(C.CkScp_syncedFiles(z.ckObj))
}
// property setter: SyncedFiles
func (z *Scp) SetSyncedFiles(goStr string) {
    cStr := C.CString(goStr)
    C.CkScp_putSyncedFiles(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustMatch
// Can contain a wildcarded list of filename patterns separated by semicolons. For
// example, "*.xml; *.txt; *.csv". If set, the SyncTreeUpload and SyncTreeDownload
// methods will only transfer files having a filename that matches any one of these
// patterns.
func (z *Scp) SyncMustMatch() string {
    return C.GoString(C.CkScp_syncMustMatch(z.ckObj))
}
// property setter: SyncMustMatch
func (z *Scp) SetSyncMustMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkScp_putSyncMustMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustMatchDir
// Can contain a wildcarded list of directory name patterns separated by
// semicolons. For example, "a*; b*; c*". If set, the SyncTreeUpload and
// SyncTreeDownload methods will only traverse into a directory that matches any
// one of these patterns.
func (z *Scp) SyncMustMatchDir() string {
    return C.GoString(C.CkScp_syncMustMatchDir(z.ckObj))
}
// property setter: SyncMustMatchDir
func (z *Scp) SetSyncMustMatchDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkScp_putSyncMustMatchDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustNotMatch
// Can contain a wildcarded list of filename patterns separated by semicolons. For
// example, "*.xml; *.txt; *.csv". If set, the SyncTreeUpload and SyncTreeDownload
// methods will not transfer files having a filename that matches any one of these
// patterns.
func (z *Scp) SyncMustNotMatch() string {
    return C.GoString(C.CkScp_syncMustNotMatch(z.ckObj))
}
// property setter: SyncMustNotMatch
func (z *Scp) SetSyncMustNotMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkScp_putSyncMustNotMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustNotMatchDir
// Can contain a wildcarded list of directory name patterns separated by
// semicolons. For example, "a*; b*; c*". If set, the SyncTreeUpload and
// SyncTreeDownload methods will not traverse into a directory that matches any one
// of these patterns.
func (z *Scp) SyncMustNotMatchDir() string {
    return C.GoString(C.CkScp_syncMustNotMatchDir(z.ckObj))
}
// property setter: SyncMustNotMatchDir
func (z *Scp) SetSyncMustNotMatchDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkScp_putSyncMustNotMatchDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. The default value is
// the empty string.
// 
// Can be set to a list of the following comma separated keywords:
//     FilenameOnly - Introduced in v9.5.0.77. Set this property to the keyword
//     "FilenameOnly" if only the filename should be used in the "scp -t" command.
//     (LANCOM routers using SCP seem to need it.)
//     ProtectFromVpn - Introduced in v9.5.0.80. On Android systems, will bypass
//     any VPN that may be installed or active.
//
func (z *Scp) UncommonOptions() string {
    return C.GoString(C.CkScp_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Scp) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkScp_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UnixPermOverride
// When Chilkat uploads a file by SCP, the UNIX permissions of the remote file are
// set based on the permissions of the local file being uploaded. Usually this is
// OK, but in some cases the access permissions of the local file are not what is
// wanted for the remote file. This property can be set to an octal permissions
// string, such as "0644", to force the remote file permissions to this value.
// 
// The default value of this property is the empty string (remote files permissions
// mirror the permissions of the local file being uploaded).
//
func (z *Scp) UnixPermOverride() string {
    return C.GoString(C.CkScp_unixPermOverride(z.ckObj))
}
// property setter: UnixPermOverride
func (z *Scp) SetUnixPermOverride(goStr string) {
    cStr := C.CString(goStr)
    C.CkScp_putUnixPermOverride(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Scp) VerboseLogging() bool {
    v := int(C.CkScp_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Scp) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkScp_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Scp) Version() string {
    return C.GoString(C.CkScp_version(z.ckObj))
}
// Downloads a binary file from the SSH server and appends to the contents of bd.
func (z *Scp) DownloadBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkScp_DownloadBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the DownloadBd method
func (z *Scp) DownloadBdAsync(arg1 string, arg2 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkScp_DownloadBdAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Downloads a binary file from the SSH server and returns the contents.
func (z *Scp) DownloadBinary(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkScp_DownloadBinary(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the DownloadBinary method
func (z *Scp) DownloadBinaryAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkScp_DownloadBinaryAsync(z.ckObj, cstr1)

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


// Downloads a file from the SSH server returns the contents in an encoded string
// (using an encoding such as base64).
// return a string or nil if failed.
func (z *Scp) DownloadBinaryEncoded(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkScp_downloadBinaryEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the DownloadBinaryEncoded method
func (z *Scp) DownloadBinaryEncodedAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkScp_DownloadBinaryEncodedAsync(z.ckObj, cstr1, cstr2)

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


// Downloads a file from the remote SSH server to the local filesystem.
func (z *Scp) DownloadFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkScp_DownloadFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the DownloadFile method
func (z *Scp) DownloadFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkScp_DownloadFileAsync(z.ckObj, cstr1, cstr2)

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


// Downloads a text file from the SSH server and returns the contents as a string.
// return a string or nil if failed.
func (z *Scp) DownloadString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkScp_downloadString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the DownloadString method
func (z *Scp) DownloadStringAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkScp_DownloadStringAsync(z.ckObj, cstr1, cstr2)

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


// Loads the caller of the task's async method.
func (z *Scp) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkScp_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Downloads files from the SSH server to a local directory tree. Synchronization
// modes include:
// 
//     mode=0: Download all files
//     mode=1: Download all files that do not exist on the local filesystem.
//     mode=2: Download newer or non-existant files.
//     mode=3: Download only newer files. If a file does not already exist on the
//     local filesystem, it is not downloaded from the server.
//     mode=5: Download only missing files or files with size differences.
//     mode=6: Same as mode 5, but also download newer files.
//     
//
func (z *Scp) SyncTreeDownload(arg1 string, arg2 string, arg3 int, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkScp_SyncTreeDownload(z.ckObj, cstr1, cstr2, C.int(arg3), b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SyncTreeDownload method
func (z *Scp) SyncTreeDownloadAsync(arg1 string, arg2 string, arg3 int, arg4 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkScp_SyncTreeDownloadAsync(z.ckObj, cstr1, cstr2, C.int(arg3), b4)

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


// Uploads a directory tree from the local filesystem to the SSH server.
// Synchronization modes include:
// 
//     mode=0: Upload all files
//     mode=1: Upload all files that do not exist on the FTP server.
//     mode=2: Upload newer or non-existant files.
//     mode=3: Upload only newer files. If a file does not already exist on the FTP
//     server, it is not uploaded.
//     mode=4: transfer missing files or files with size differences.
//     mode=5: same as mode 4, but also newer files.
//
func (z *Scp) SyncTreeUpload(arg1 string, arg2 string, arg3 int, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkScp_SyncTreeUpload(z.ckObj, cstr1, cstr2, C.int(arg3), b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SyncTreeUpload method
func (z *Scp) SyncTreeUploadAsync(arg1 string, arg2 string, arg3 int, arg4 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkScp_SyncTreeUploadAsync(z.ckObj, cstr1, cstr2, C.int(arg3), b4)

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


// Uploads the contents of bd to a file on the SSH server.
func (z *Scp) UploadBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkScp_UploadBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the UploadBd method
func (z *Scp) UploadBdAsync(arg1 string, arg2 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkScp_UploadBdAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Uploads binary data to a file on the SSH server.
func (z *Scp) UploadBinary(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkScp_UploadBinary(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Asynchronous version of the UploadBinary method
func (z *Scp) UploadBinaryAsync(arg1 string, arg2 []byte, c chan *Task) {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    hTask := C.CkScp_UploadBinaryAsync(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Uploads the binary data to a file on the remote SSH server. The binary data is
// passed in encoded string representation (such as base64, or hex).
func (z *Scp) UploadBinaryEncoded(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkScp_UploadBinaryEncoded(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the UploadBinaryEncoded method
func (z *Scp) UploadBinaryEncodedAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkScp_UploadBinaryEncodedAsync(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Uploads a file from the local filesystem to the remote SSH server.
func (z *Scp) UploadFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkScp_UploadFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UploadFile method
func (z *Scp) UploadFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkScp_UploadFileAsync(z.ckObj, cstr1, cstr2)

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


// Uploads the contents of a string to a file on the remote SSH server.
func (z *Scp) UploadString(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkScp_UploadString(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the UploadString method
func (z *Scp) UploadStringAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkScp_UploadStringAsync(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Uses the SSH connection of sshConnection for the SCP transfers. All of the connection and
// socket related properties, proxy properites, timeout properties, session log,
// etc. set on the SSH object apply to the SCP methods (because internally it is
// the SSH object that is used to do the work of the file transfers).
// 
// Note: There is no UnlockComponent method in the SCP class because it is the SSH
// object that must be unlocked. When the SSH object is not unlocked, this method
// will return false to indicate failure.
//
func (z *Scp) UseSsh(arg1 *Ssh) bool {

    v := C.CkScp_UseSsh(z.ckObj, arg1.ckObj)


    return v != 0
}


