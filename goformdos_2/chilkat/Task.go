// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkTask.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewTask() *Task {
	return &Task{C.CkTask_Create()}
}

func (z *Task) DisposeTask() {
    if z != nil {
	C.CkTask_Dispose(z.ckObj)
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
func (z *Task) DebugLogFilePath() string {
    return C.GoString(C.CkTask_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Task) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkTask_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Finished
// true if the task status is "canceled", "aborted", or "completed". A task can
// only reach the "canceled" status if it was activated via the Run method, made it
// onto the internal thread pool thread's queue, was waiting for a pool thread to
// become available, and was then canceled prior to the task actually starting.
func (z *Task) Finished() bool {
    v := int(C.CkTask_getFinished(z.ckObj))
    return v != 0
}

// property: HeartbeatMs
// The number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort the Wait method. If
// HeartbeatMs is 0 (the default), no AbortCheck event callbacks will fire. Note:
// An asynchronous task running in a background thread (in one of the thread pool
// threads) does not fire events. The task's event callbacks pertain only to the
// Wait method.
func (z *Task) HeartbeatMs() int {
    return int(C.CkTask_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Task) SetHeartbeatMs(value int) {
    C.CkTask_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: Inert
// true if the task status is "empty" or "loaded". When a task is inert, it has
// been loaded but is not scheduled to run yet.
func (z *Task) Inert() bool {
    v := int(C.CkTask_getInert(z.ckObj))
    return v != 0
}

// property: KeepProgressLog
// Determines if the in-memory progress info event log is kept. The default value
// is false and therefore no log is kept. To enable progress info logging, set
// this property equal to true (prior to running the task).
func (z *Task) KeepProgressLog() bool {
    v := int(C.CkTask_getKeepProgressLog(z.ckObj))
    return v != 0
}
// property setter: KeepProgressLog
func (z *Task) SetKeepProgressLog(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTask_putKeepProgressLog(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Task) LastErrorHtml() string {
    return C.GoString(C.CkTask_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Task) LastErrorText() string {
    return C.GoString(C.CkTask_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Task) LastErrorXml() string {
    return C.GoString(C.CkTask_lastErrorXml(z.ckObj))
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
func (z *Task) LastMethodSuccess() bool {
    v := int(C.CkTask_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Task) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTask_putLastMethodSuccess(z.ckObj,v)
}

// property: Live
// true if the task status is "queued" or "running". When a task is live, it is
// either already running, or is on the thread pool thread's queue waiting for a
// thread to become available.
func (z *Task) Live() bool {
    v := int(C.CkTask_getLive(z.ckObj))
    return v != 0
}

// property: PercentDone
// Indicates the percent completion while the task is running. The percent
// completed information is only available in cases where it is possible to know
// the percentage completed. For some methods, it is never possible to know, such
// as for methods that establish TCP or TLS connections. For other methods it is
// always possible to know -- such as for sending email (because the size of the
// email to be sent is already known). For some methods, it may or may not be
// possible to know the percent completed. For example, if an HTTP response is
// "chunked", there is no Content-Length header and therefore the receiver has no
// knowledge of the size of the forthcoming response body.
// 
// Also, the value of the PercentDoneScale property of the asynchronous method's
// object determines the scale, such as 0 to 100, or 0 to 1000, etc.
//
func (z *Task) PercentDone() int {
    return int(C.CkTask_getPercentDone(z.ckObj))
}

// property: ProgressLogSize
// What would normally be a ProgressInfo event callback (assuming Chilkat supports
// event callbacks for this language) is instead saved to an in-memory progress log
// that can be examined and pruned while the task is still running. This property
// returns the number of progress log entries that are currently available. (Note:
// the progress log is only kept if the KeepProgressLog property is turned on. By
// default, the KeepProgressLog is turned off.)
func (z *Task) ProgressLogSize() int {
    return int(C.CkTask_getProgressLogSize(z.ckObj))
}

// property: ResultErrorText
// The LastErrorText for the task's asynchronous method.
func (z *Task) ResultErrorText() string {
    return C.GoString(C.CkTask_resultErrorText(z.ckObj))
}

// property: ResultType
// Indicates the data type of the task's result. This property is only available
// after the task has completed. Possible values are "bool", "int", "string",
// "bytes", "object", and "void". For example, if the result data type is "bool",
// then call GetResultBool to get the boolean result of the underlying asynchronous
// method.
// 
// For example, if the synchronous version of the method returned a boolean, then
// in the asynchronous version of the method, the boolean return value is made
// available via the GetResultBool method.
//
func (z *Task) ResultType() string {
    return C.GoString(C.CkTask_resultType(z.ckObj))
}

// property: Status
// The current status of the task. Possible values are:
//     "empty" -- The method call and arguments are not yet loaded into the task
//     object. This can only happen if a task was explicitly created instead of being
//     returned by a method ending in "Async".
//     "loaded" -- The method call and arguments are loaded into the task object.
//     "queued" -- The task is in the thread pool's queue of tasks awaiting to be
//     run.
//     "running" -- The task is currently running.
//     "canceled" -- The task was canceled before it entered the "running" state.
//     "aborted" -- The task was canceled while it was in the running state.
//     "completed" -- The task completed. The success or failure depends on the
//     semantics of the method call and the value of the result.
func (z *Task) Status() string {
    return C.GoString(C.CkTask_status(z.ckObj))
}

// property: StatusInt
// The current status of the task as an integer value. Possible values are:
//     1 -- The method call and arguments are not yet loaded into the task object.
//     This can only happen if a task was explicitly created instead of being returned
//     by a method ending in "Async".
//     2 -- The method call and arguments are loaded into the task object.
//     3 -- The task is in the thread pool's queue of tasks awaiting to be run.
//     4 -- The task is currently running.
//     5 -- The task was canceled before it entered the "running" state.
//     6 -- The task was canceled while it was in the running state.
//     7 -- The task completed. The success or failure depends on the semantics of
//     the method call and the value of the result.
func (z *Task) StatusInt() int {
    return int(C.CkTask_getStatusInt(z.ckObj))
}

// property: TaskId
// A unique integer ID assigned to this task. The purpose of this property is to
// help an application identify the task if a TaskCompleted event callback is used.
func (z *Task) TaskId() int {
    return int(C.CkTask_getTaskId(z.ckObj))
}

// property: TaskSuccess
// This is the value of the LastMethodSuccess property of the underlying task
// object. This property is only valid for those methods where the
// LastMethodSuccess property would be valid had the method been called
// synchronously.
// 
// Important: This property is only meaningful for cases where the underlying
// method call has a non-boolean return value (such as for methods that return
// strings, other Chilkat objects, or integers). If the underlying method call
// returns a boolean, then call the GetResultBool() method instead to get the
// boolean return value.
//
func (z *Task) TaskSuccess() bool {
    v := int(C.CkTask_getTaskSuccess(z.ckObj))
    return v != 0
}

// property: UserData
// An application may use this property to attach some user-specific information
// with the task, which may be useful if a TaskCompleted event callback is used.
func (z *Task) UserData() string {
    return C.GoString(C.CkTask_userData(z.ckObj))
}
// property setter: UserData
func (z *Task) SetUserData(goStr string) {
    cStr := C.CString(goStr)
    C.CkTask_putUserData(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Task) VerboseLogging() bool {
    v := int(C.CkTask_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Task) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTask_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Task) Version() string {
    return C.GoString(C.CkTask_version(z.ckObj))
}
// Marks an asynchronous task for cancellation. The expected behavior depends on
// the current status of the task as described here:
//     "loaded" - If the task has been loaded but has not yet been queued to run in
//     the thread pool, then there is nothing to do. (There is nothing to cancel
//     because the task's Run method has not yet been called.) The task will remain in
//     the "loaded" state.
//     "queued" - The task is marked for cancellation, is dequeued, and will not
//     run. The task's status changes immediately to "canceled".
//     "running" - The already-running task is marked for cancellation. The task's
//     status will eventually change to "aborted" when the asynchronous method returns.
//     At that point in time, the ResultErrorText property will contain the
//     "LastErrorText" of the method call. In the case where a task is marked for
//     cancellation just at the time it's completing, the task status may instead
//     change to "completed".
//     "canceled", "aborted", "completed" - In these cases the task has already
//     finished, and there will be no change in status.
// Cancel returns true if the task was in the "queued" or "running" state when it
// was marked for cancellation. Cancel returns false if the task was in any other
// state.
// 
// Important: Calling the Cancel method marks a task for cancellation. It sets a
// flag in memory that the running task will soon notice and then abort. It is
// important to realize that your application is likely calling Cancel from the
// main UI thread, whereas the asynchronous task is running in a background thread.
// If the task was in the "running" state when Cancel was called, it will still be
// in the "running" state when Cancel returns. It will take a short amount of time
// until the task actually aborts. This is because operating systems schedule
// threads in time slices, and the thread needs one or more time slices to notice
// the cancellation flag and abort. After calling Cancel, your application might
// wish to call the Wait method to wait until the task has actually aborted, or it
// could periodically check the task's status and then react once the status
// changes to "aborted".
//
func (z *Task) Cancel() bool {

    v := C.CkTask_Cancel(z.ckObj)


    return v != 0
}


// Removes all entries from the progress info log.
func (z *Task) ClearProgressLog()  {

    C.CkTask_ClearProgressLog(z.ckObj)



}


// Returns the binary bytes result of the task. The bytes are copied to the caller.
func (z *Task) CopyResultBytes() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkTask_CopyResultBytes(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the boolean result of the task.
func (z *Task) GetResultBool() bool {

    v := C.CkTask_GetResultBool(z.ckObj)


    return v != 0
}


// Returns the binary bytes result of the task. The bytes are transferred to the
// caller, not copied. Call CopyResultBytes instead to copy the result bytes.
func (z *Task) GetResultBytes() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkTask_GetResultBytes(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the integer result of the task.
func (z *Task) GetResultInt() int {

    v := C.CkTask_GetResultInt(z.ckObj)


    return int(v)
}


// Returns the string result of the task.
// return a string or nil if failed.
func (z *Task) GetResultString() *string {

    retStr := C.CkTask_getResultString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the name of the Nth progress info event logged. The 1st entry is at
// index 0.
// return a string or nil if failed.
func (z *Task) ProgressInfoName(arg1 int) *string {

    retStr := C.CkTask_progressInfoName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of the Nth progress info event logged. The 1st entry is at
// index 0.
// return a string or nil if failed.
func (z *Task) ProgressInfoValue(arg1 int) *string {

    retStr := C.CkTask_progressInfoValue(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Removes the Nth progress info log entry.
func (z *Task) RemoveProgressInfo(arg1 int)  {

    C.CkTask_RemoveProgressInfo(z.ckObj, C.int(arg1))



}


// If a taskCompleted callback function is passed in, then the task is started on
// Node's internal thread pool. If no callback function is passed, then the task is
// run synchronously. Queues the task to run on the internal Chilkat thread pool.
func (z *Task) Run() bool {

    v := C.CkTask_Run(z.ckObj)


    return v != 0
}


// Runs the task synchronously. Then this method returns after the task has been
// run.
func (z *Task) RunSynchronously() bool {

    v := C.CkTask_RunSynchronously(z.ckObj)


    return v != 0
}


// Convenience method to force the calling thread to sleep for a number of
// milliseconds. (This does not cause the task's background thread to sleep.)
func (z *Task) SleepMs(arg1 int)  {

    C.CkTask_SleepMs(z.ckObj, C.int(arg1))



}


// Waits for the task to complete. Returns when task has completed, or after maxWaitMs
// milliseconds have elapsed. (A maxWaitMs value of 0 is to wait indefinitely.) Returns
// (false) if the task has not yet been started by calling the Run method, or if
// the maxWaitMs expired. If the task completed, was already completed, was canceled or
// aborted, then this method returns true.
func (z *Task) Wait(arg1 int) bool {

    v := C.CkTask_Wait(z.ckObj, C.int(arg1))


    return v != 0
}


