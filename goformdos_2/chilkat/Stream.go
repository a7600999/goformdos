// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkStream.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewStream() *Stream {
	return &Stream{C.CkStream_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Stream) DisposeStream() {
    if z != nil {
	C.CkStream_Dispose(z.ckObj)
	}
}




func (z *Stream) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkStream_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Stream) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkStream_setExternalProgress(z.ckObj,1)
}

func (z *Stream) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkStream_setExternalProgress(z.ckObj,1)
}

func (z *Stream) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkStream_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Stream) AbortCurrent() bool {
    v := int(C.CkStream_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Stream) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStream_putAbortCurrent(z.ckObj,v)
}

// property: CanRead
// true if the stream supports reading. Otherwise false.
// 
// Note: A stream that supports reading, which has already reached the
// end-of-stream, will still have a CanRead value of true. This property
// indicates the stream's inherent ability, and not whether or not the stream can
// be read at a particular moment in time.
//
func (z *Stream) CanRead() bool {
    v := int(C.CkStream_getCanRead(z.ckObj))
    return v != 0
}

// property: CanWrite
// true if the stream supports writing. Otherwise false.
// 
// Note: A stream that supports writing, which has already been closed for write,
// will still have a CanWrite value of true. This property indicates the stream's
// inherent ability, and not whether or not the stream can be written at a
// particular moment in time.
//
func (z *Stream) CanWrite() bool {
    v := int(C.CkStream_getCanWrite(z.ckObj))
    return v != 0
}

// property: DataAvailable
// true if it is known for sure that data is ready and waiting to be read.
// false if it is not known for sure (it may be that data is immediately
// available, but reading the stream with a ReadTimeoutMs of 0, which is to poll
// the stream, is the way to find out).
func (z *Stream) DataAvailable() bool {
    v := int(C.CkStream_getDataAvailable(z.ckObj))
    return v != 0
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
func (z *Stream) DebugLogFilePath() string {
    return C.GoString(C.CkStream_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Stream) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkStream_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DefaultChunkSize
// The default internal chunk size for reading or writing. The default value is
// 65536. If this property is set to 0, it will cause the default chunk size
// (65536) to be used. Note: The chunk size can have significant performance
// impact. If performance is an issue, be sure to experiment with different chunk
// sizes.
func (z *Stream) DefaultChunkSize() int {
    return int(C.CkStream_getDefaultChunkSize(z.ckObj))
}
// property setter: DefaultChunkSize
func (z *Stream) SetDefaultChunkSize(value int) {
    C.CkStream_putDefaultChunkSize(z.ckObj,C.int(value))
}

// property: EndOfStream
// true if the end-of-stream has already been reached. When the stream has
// already ended, all calls to Read* methods will return false with the
// ReadFailReason set to 3 (already at end-of-stream).
func (z *Stream) EndOfStream() bool {
    v := int(C.CkStream_getEndOfStream(z.ckObj))
    return v != 0
}

// property: IsWriteClosed
// true if the stream is closed for writing. Once closed, no more data may be
// written to the stream.
func (z *Stream) IsWriteClosed() bool {
    v := int(C.CkStream_getIsWriteClosed(z.ckObj))
    return v != 0
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Stream) LastErrorHtml() string {
    return C.GoString(C.CkStream_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Stream) LastErrorText() string {
    return C.GoString(C.CkStream_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Stream) LastErrorXml() string {
    return C.GoString(C.CkStream_lastErrorXml(z.ckObj))
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
func (z *Stream) LastMethodSuccess() bool {
    v := int(C.CkStream_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Stream) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStream_putLastMethodSuccess(z.ckObj,v)
}

// property: Length
// The length (in bytes) of the stream's source. If unknown, then this property
// will have a value of -1. This property may be set by the application if it knows
// in advance the length of the stream.
func (z *Stream) Length() int64 {
    return int64(C.CkStream_getLength(z.ckObj))
}
// property setter: Length
func (z *Stream) SetLength(value int64) {
    C.CkStream_putLength(z.ckObj,C.__int64(value))
}

// property: Length32
// The length (in bytes) of the stream's source. If unknown, then this property
// will have a value of -1. This property may be set by the application if it knows
// in advance the length of the stream.
// 
// Setting this property also sets the Length property (which is a 64-bit integer).
//
func (z *Stream) Length32() int {
    return int(C.CkStream_getLength32(z.ckObj))
}
// property setter: Length32
func (z *Stream) SetLength32(value int) {
    C.CkStream_putLength32(z.ckObj,C.int(value))
}

// property: NumReceived
// The number of bytes received by the stream.
func (z *Stream) NumReceived() int64 {
    return int64(C.CkStream_getNumReceived(z.ckObj))
}

// property: NumSent
// The number of bytes sent by the stream.
func (z *Stream) NumSent() int64 {
    return int64(C.CkStream_getNumSent(z.ckObj))
}

// property: ReadFailReason
// This property is automatically set when a Read* method is called. It indicates
// the reason for failure. Possible values are:
//     No failure (success)
//     Timeout, or no data is immediately available for a polling read.
//     Aborted by an application callback.
//     Already at end-of-stream.
//     Fatal stream error.
//     Out-of-memory error (this is very unlikely).
func (z *Stream) ReadFailReason() int {
    return int(C.CkStream_getReadFailReason(z.ckObj))
}

// property: ReadTimeoutMs
// The maximum number of seconds to wait while reading. The default value is 30
// seconds (i.e. 30000ms). A value of 0 indicates a poll. (A polling read is to
// return with a timeout if no data is immediately available.)
// 
// Important: For most Chilkat timeout related properties, a value of 0 indicates
// an infinite timeout. For this property, a value of 0 indicates a poll. If
// setting a timeout related property (or method argument) to zero, be sure to
// understand if 0 means "wait forever" or "poll".
// 
// The timeout value is not a total timeout. It is the maximum time to wait while
// no additional data is forthcoming.
//
func (z *Stream) ReadTimeoutMs() int {
    return int(C.CkStream_getReadTimeoutMs(z.ckObj))
}
// property setter: ReadTimeoutMs
func (z *Stream) SetReadTimeoutMs(value int) {
    C.CkStream_putReadTimeoutMs(z.ckObj,C.int(value))
}

// property: SinkFile
// Sets the sink to the path of a file. The file does not need to exist at the time
// of setting this property. The sink file will be automatically opened on demand,
// when the stream is first written.
// 
// Note: This property takes priority over other potential sinks. Make sure this
// property is set to an empty string if the stream's sink is to be something else.
//
func (z *Stream) SinkFile() string {
    return C.GoString(C.CkStream_sinkFile(z.ckObj))
}
// property setter: SinkFile
func (z *Stream) SetSinkFile(goStr string) {
    cStr := C.CString(goStr)
    C.CkStream_putSinkFile(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SinkFileAppend
// If true, the stream appends to the SinkFile rather than truncating and
// re-writing the sink file. The default value is false.
func (z *Stream) SinkFileAppend() bool {
    v := int(C.CkStream_getSinkFileAppend(z.ckObj))
    return v != 0
}
// property setter: SinkFileAppend
func (z *Stream) SetSinkFileAppend(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStream_putSinkFileAppend(z.ckObj,v)
}

// property: SourceFile
// Sets the source to the path of a file. The file does not need to exist at the
// time of setting this property. The source file will be automatically opened on
// demand, when the stream is first read.
// 
// Note: This property takes priority over other potential sources. Make sure this
// property is set to an empty string if the stream's source is to be something
// else.
//
func (z *Stream) SourceFile() string {
    return C.GoString(C.CkStream_sourceFile(z.ckObj))
}
// property setter: SourceFile
func (z *Stream) SetSourceFile(goStr string) {
    cStr := C.CString(goStr)
    C.CkStream_putSourceFile(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SourceFilePart
// If the source is a file, then this property can be used to stream one part of
// the file. The SourceFilePartSize property defines the size (in bytes) of each
// part. The SourceFilePart and SourceFilePartSize have default values of 0, which
// means the entire SourceFile is streamed.
// 
// This property is a 0-based index. For example, if the SourceFilePartSize is
// 1000, then part 0 is for bytes 0 to 999, part 1 is for bytes 1000 to 1999, etc.
//
func (z *Stream) SourceFilePart() int {
    return int(C.CkStream_getSourceFilePart(z.ckObj))
}
// property setter: SourceFilePart
func (z *Stream) SetSourceFilePart(value int) {
    C.CkStream_putSourceFilePart(z.ckObj,C.int(value))
}

// property: SourceFilePartSize
// If the source is a file, then this property, in conjunction with the
// SourceFilePart property, can be used to stream a single part of the file. This
// property defines the size (in bytes) of each part. The SourceFilePart and
// SourceFilePartSize have default values of 0, which means that by default, the
// entire SourceFile is streamed.
func (z *Stream) SourceFilePartSize() int {
    return int(C.CkStream_getSourceFilePartSize(z.ckObj))
}
// property setter: SourceFilePartSize
func (z *Stream) SetSourceFilePartSize(value int) {
    C.CkStream_putSourceFilePartSize(z.ckObj,C.int(value))
}

// property: StringBom
// If true, then include the BOM when creating a string source via
// SetSourceString where the charset is utf-8, utf-16, etc. (The term "BOM" stands
// for Byte Order Mark, also known as the preamble.) Also, if true, then include
// the BOM when writing a string via the WriteString method. The default value of
// this property is false.
func (z *Stream) StringBom() bool {
    v := int(C.CkStream_getStringBom(z.ckObj))
    return v != 0
}
// property setter: StringBom
func (z *Stream) SetStringBom(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStream_putStringBom(z.ckObj,v)
}

// property: StringCharset
// Indicates the expected character encoding, such as utf-8, windows-1256, utf-16,
// etc. for methods that read text such as: ReadString, ReadToCRLF, ReadUntilMatch.
// Also controls the character encoding when writing strings with the WriteString
// method. The supported charsets are indicated at the link below.
// 
// The default value is "utf-8".
//
func (z *Stream) StringCharset() string {
    return C.GoString(C.CkStream_stringCharset(z.ckObj))
}
// property setter: StringCharset
func (z *Stream) SetStringCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkStream_putStringCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Stream) VerboseLogging() bool {
    v := int(C.CkStream_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Stream) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStream_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Stream) Version() string {
    return C.GoString(C.CkStream_version(z.ckObj))
}

// property: WriteFailReason
// This property is automatically set when a Write* method is called. It indicates
// the reason for failure. Possible values are:
//     No failure (success)
//     Timeout, or unable to immediately write when the WriteTimeoutMs is 0.
//     Aborted by an application callback.
//     The stream has already ended.
//     Fatal stream error.
//     Out-of-memory error (this is very unlikely).
func (z *Stream) WriteFailReason() int {
    return int(C.CkStream_getWriteFailReason(z.ckObj))
}

// property: WriteTimeoutMs
// The maximum number of seconds to wait while writing. The default value is 30
// seconds (i.e. 30000ms). A value of 0 indicates to return immediately if it is
// not possible to write to the sink immediately.
func (z *Stream) WriteTimeoutMs() int {
    return int(C.CkStream_getWriteTimeoutMs(z.ckObj))
}
// property setter: WriteTimeoutMs
func (z *Stream) SetWriteTimeoutMs(value int) {
    C.CkStream_putWriteTimeoutMs(z.ckObj,C.int(value))
}
// Loads the caller of the task's async method.
func (z *Stream) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkStream_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Read as much data as is immediately available on the stream. If no data is
// immediately available, it waits up to ReadTimeoutMs milliseconds for data to
// arrive. The incoming data is appended to binData.
func (z *Stream) ReadBd(arg1 *BinData) bool {

    v := C.CkStream_ReadBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the ReadBd method
func (z *Stream) ReadBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkStream_ReadBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Read as much data as is immediately available on the stream. If no data is
// immediately available, it waits up to ReadTimeoutMs milliseconds for data to
// arrive.
func (z *Stream) ReadBytes() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkStream_ReadBytes(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the ReadBytes method
func (z *Stream) ReadBytesAsync(c chan *Task) {

    hTask := C.CkStream_ReadBytesAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as ReadBytes, except returns the received bytes in encoded string form.
// The encoding argument indicates the encoding, which can be "base64", "hex", or any
// of the multitude of encodings indicated in the link below.
// return a string or nil if failed.
func (z *Stream) ReadBytesENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkStream_readBytesENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReadBytesENC method
func (z *Stream) ReadBytesENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkStream_ReadBytesENCAsync(z.ckObj, cstr1)

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


// Reads exactly numBytes bytes from the stream. If no data is immediately available,
// it waits up to ReadTimeoutMs milliseconds for data to arrive.
func (z *Stream) ReadNBytes(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkStream_ReadNBytes(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the ReadNBytes method
func (z *Stream) ReadNBytesAsync(arg1 int, c chan *Task) {

    hTask := C.CkStream_ReadNBytesAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as ReadNBytes, except returns the received bytes in encoded string
// form. The encoding argument indicates the encoding, which can be "base64", "hex", or
// any of the multitude of encodings indicated in the link below.
// return a string or nil if failed.
func (z *Stream) ReadNBytesENC(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkStream_readNBytesENC(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReadNBytesENC method
func (z *Stream) ReadNBytesENCAsync(arg1 int, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkStream_ReadNBytesENCAsync(z.ckObj, C.int(arg1), cstr2)

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


// Read as much data as is immediately available on the stream. If no data is
// immediately available, it waits up to ReadTimeoutMs milliseconds for data to
// arrive. The data is appended to sb. The incoming bytes are interpreted
// according to the StringCharset property. For example, if utf-8 bytes are
// expected, then StringCharset should be set to "utf-8" prior to calling ReadSb.
func (z *Stream) ReadSb(arg1 *StringBuilder) bool {

    v := C.CkStream_ReadSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the ReadSb method
func (z *Stream) ReadSbAsync(arg1 *StringBuilder, c chan *Task) {

    hTask := C.CkStream_ReadSbAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Read as much data as is immediately available on the stream. If no data is
// immediately available, it waits up to ReadTimeoutMs milliseconds for data to
// arrive. The data is returned as a string. The incoming bytes are interpreted
// according to the StringCharset property. For example, if utf-8 bytes are
// expected, then StringCharset should be set to "utf-8" prior to calling
// ReadString.
// return a string or nil if failed.
func (z *Stream) ReadString() *string {

    retStr := C.CkStream_readString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReadString method
func (z *Stream) ReadStringAsync(c chan *Task) {

    hTask := C.CkStream_ReadStringAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads the stream until a CRLF is received. If no data is immediately available,
// it waits up to ReadTimeoutMs milliseconds for data to arrive. The data is
// returned as a string. The incoming bytes are interpreted according to the
// StringCharset property. For example, if utf-8 bytes are expected, then
// StringCharset should be set to "utf-8" prior to calling ReadString.
// 
// Note: If the end-of-stream is reached prior to receiving the CRLF, then the
// remaining data is returned, and the ReadFailReason property will be set to 3 (to
// indicate end-of-file). This is the only case where as string would be returned
// that does not end with CRLF.
//
// return a string or nil if failed.
func (z *Stream) ReadToCRLF() *string {

    retStr := C.CkStream_readToCRLF(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReadToCRLF method
func (z *Stream) ReadToCRLFAsync(c chan *Task) {

    hTask := C.CkStream_ReadToCRLFAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads the stream until the string indicated by matchStr is received. If no data is
// immediately available, it waits up to ReadTimeoutMs milliseconds for data to
// arrive. The data is returned as a string. The incoming bytes are interpreted
// according to the StringCharset property. For example, if utf-8 bytes are
// expected, then StringCharset should be set to "utf-8" prior to calling
// ReadString.
// 
// Note: If the end-of-stream is reached prior to receiving the match string, then
// the remaining data is returned, and the ReadFailReason property will be set to 3
// (to indicate end-of-file). This is the only case where as string would be
// returned that does not end with the desired match string.
//
// return a string or nil if failed.
func (z *Stream) ReadUntilMatch(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkStream_readUntilMatch(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReadUntilMatch method
func (z *Stream) ReadUntilMatchAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkStream_ReadUntilMatchAsync(z.ckObj, cstr1)

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


// Resets the stream. If a source or sink is open, then it is closed. Properties
// such as EndOfStream and IsWriteClose are reset to default values.
func (z *Stream) Reset()  {

    C.CkStream_Reset(z.ckObj)



}


// Runs the stream to completion. This method should only be called when the source
// of the string has been set by any of the following methods: SetSourceBytes,
// SetSourceString, or SetSourceStream, or when the SourceFile property has been
// set (giving the stream a file source).
// 
// This method will read the stream source and forward to the sink until the
// end-of-stream is reached, and all data has been written to the sink.
//
func (z *Stream) RunStream() bool {

    v := C.CkStream_RunStream(z.ckObj)


    return v != 0
}


// Asynchronous version of the RunStream method
func (z *Stream) RunStreamAsync(c chan *Task) {

    hTask := C.CkStream_RunStreamAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sets the stream's sink to strm. Any data written to this stream's sink will
// become available to strm on its source.
func (z *Stream) SetSinkStream(arg1 *Stream) bool {

    v := C.CkStream_SetSinkStream(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the stream's source to the contents of sourceData.
func (z *Stream) SetSourceBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkStream_SetSourceBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Sets the stream's source to be the sink of strm. Any data written to strm's sink
// will become available on this stream's source.
func (z *Stream) SetSourceStream(arg1 *Stream) bool {

    v := C.CkStream_SetSourceStream(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the stream's source to the contents of srcStr. The charset indicates the
// character encoding to be used for the byte representation of the srcStr.
func (z *Stream) SetSourceString(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStream_SetSourceString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Writes the contents of binData to the stream.
func (z *Stream) WriteBd(arg1 *BinData) bool {

    v := C.CkStream_WriteBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the WriteBd method
func (z *Stream) WriteBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkStream_WriteBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Writes a single byte to the stream. The byteVal must have a value from 0 to 255.
func (z *Stream) WriteByte(arg1 int) bool {

    v := C.CkStream_WriteByte(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the WriteByte method
func (z *Stream) WriteByteAsync(arg1 int, c chan *Task) {

    hTask := C.CkStream_WriteByteAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Writes binary bytes to a stream.
func (z *Stream) WriteBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkStream_WriteBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Asynchronous version of the WriteBytes method
func (z *Stream) WriteBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkStream_WriteBytesAsync(z.ckObj, ckbyd1)

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


// Writes binary bytes to a stream. The byte data is passed in encoded string form,
// where the encoding can be "base64", "hex", or any of the supported binary encodings
// listed at the link below.
func (z *Stream) WriteBytesENC(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStream_WriteBytesENC(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the WriteBytesENC method
func (z *Stream) WriteBytesENCAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkStream_WriteBytesENCAsync(z.ckObj, cstr1, cstr2)

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


// Indicates that no more data will be written to the stream.
func (z *Stream) WriteClose() bool {

    v := C.CkStream_WriteClose(z.ckObj)


    return v != 0
}


// Writes the contents of sb to the stream. The actual bytes written are the byte
// representation of the string as indicated by the StringCharset property. For
// example, to write utf-8 bytes, first set StringCharset equal to "utf-8" and then
// call WriteSb.
func (z *Stream) WriteSb(arg1 *StringBuilder) bool {

    v := C.CkStream_WriteSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the WriteSb method
func (z *Stream) WriteSbAsync(arg1 *StringBuilder, c chan *Task) {

    hTask := C.CkStream_WriteSbAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Writes a string to a stream. The actual bytes written are the byte
// representation of the string as indicated by the StringCharset property. For
// example, to write utf-8 bytes, first set StringCharset equal to "utf-8" and then
// call WriteString.
func (z *Stream) WriteString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStream_WriteString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the WriteString method
func (z *Stream) WriteStringAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkStream_WriteStringAsync(z.ckObj, cstr1)

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


