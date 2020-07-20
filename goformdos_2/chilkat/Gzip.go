// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkGzip.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewGzip() *Gzip {
	return &Gzip{C.CkGzip_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Gzip) DisposeGzip() {
    if z != nil {
	C.CkGzip_Dispose(z.ckObj)
	}
}




func (z *Gzip) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkGzip_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Gzip) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkGzip_setExternalProgress(z.ckObj,1)
}

func (z *Gzip) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkGzip_setExternalProgress(z.ckObj,1)
}

func (z *Gzip) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkGzip_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Gzip) AbortCurrent() bool {
    v := int(C.CkGzip_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Gzip) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGzip_putAbortCurrent(z.ckObj,v)
}

// property: Comment
// Specifies an optional comment string that can be embedded within the .gz when
// any Compress* method is called.
func (z *Gzip) Comment() string {
    return C.GoString(C.CkGzip_comment(z.ckObj))
}
// property setter: Comment
func (z *Gzip) SetComment(goStr string) {
    cStr := C.CString(goStr)
    C.CkGzip_putComment(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CompressionLevel
// The compression level to be used when compressing. The default is 6, which is
// the typical value used for zip utility programs when compressing data. The
// compression level may range from 0 (no compression) to 9 (the most compression).
// The benefits of trying to increase compression may not be worth the added
// expense in compression time. In many cases (which is data dependent), the
// improvement in compression is minimal while the increase in computation time is
// significant.
func (z *Gzip) CompressionLevel() int {
    return int(C.CkGzip_getCompressionLevel(z.ckObj))
}
// property setter: CompressionLevel
func (z *Gzip) SetCompressionLevel(value int) {
    C.CkGzip_putCompressionLevel(z.ckObj,C.int(value))
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
func (z *Gzip) DebugLogFilePath() string {
    return C.GoString(C.CkGzip_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Gzip) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkGzip_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ExtraData
// Optional extra-data that can be included within a .gz when a Compress* method is
// called.
func (z *Gzip) ExtraData() []byte {
    ckbyd := C.CkByteData_Create()
    C.CkGzip_getExtraData(z.ckObj,ckbyd)
    p := C.CkByteData_getData(ckbyd);
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckbyd)) )
    C.CkByteData_Dispose(ckbyd)
    return retBytes
}

// property setter: ExtraData
func (z *Gzip) SetExtraData(dataBytes []byte) {
    ckbyd := C.CkByteData_Create()
    pData := C.CBytes(dataBytes)
    C.CkByteData_borrowData(ckbyd, (*C.uchar)(pData), C.ulong(len(dataBytes)))
    C.CkGzip_putExtraData(z.ckObj,ckbyd)    
    C.free(pData)
    C.CkByteData_Dispose(ckbyd)
}


// property: Filename
// The filename that is embedded within the .gz during any Compress* method call.
// When extracting from a .gz using applications such as WinZip, this will be the
// filename that is created.
func (z *Gzip) Filename() string {
    return C.GoString(C.CkGzip_filename(z.ckObj))
}
// property setter: Filename
func (z *Gzip) SetFilename(goStr string) {
    cStr := C.CString(goStr)
    C.CkGzip_putFilename(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// The number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any method call prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
func (z *Gzip) HeartbeatMs() int {
    return int(C.CkGzip_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Gzip) SetHeartbeatMs(value int) {
    C.CkGzip_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Gzip) LastErrorHtml() string {
    return C.GoString(C.CkGzip_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Gzip) LastErrorText() string {
    return C.GoString(C.CkGzip_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Gzip) LastErrorXml() string {
    return C.GoString(C.CkGzip_lastErrorXml(z.ckObj))
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
func (z *Gzip) LastMethodSuccess() bool {
    v := int(C.CkGzip_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Gzip) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGzip_putLastMethodSuccess(z.ckObj,v)
}

// property: LastModStr
// The same as the LastMod property, but allows the date/time to be get/set in
// RFC822 string format.
func (z *Gzip) LastModStr() string {
    return C.GoString(C.CkGzip_lastModStr(z.ckObj))
}
// property setter: LastModStr
func (z *Gzip) SetLastModStr(goStr string) {
    cStr := C.CString(goStr)
    C.CkGzip_putLastModStr(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UseCurrentDate
// If set to true, the file produced by an Uncompress* method will use the
// current date/time for the last-modification date instead of the date/time found
// within the Gzip format.
func (z *Gzip) UseCurrentDate() bool {
    v := int(C.CkGzip_getUseCurrentDate(z.ckObj))
    return v != 0
}
// property setter: UseCurrentDate
func (z *Gzip) SetUseCurrentDate(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGzip_putUseCurrentDate(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Gzip) VerboseLogging() bool {
    v := int(C.CkGzip_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Gzip) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGzip_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Gzip) Version() string {
    return C.GoString(C.CkGzip_version(z.ckObj))
}
// In-place gzip the contents of binDat.
func (z *Gzip) CompressBd(arg1 *BinData) bool {

    v := C.CkGzip_CompressBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the CompressBd method
func (z *Gzip) CompressBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkGzip_CompressBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Compresses a file to create a GZip compressed file (.gz).
func (z *Gzip) CompressFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkGzip_CompressFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CompressFile method
func (z *Gzip) CompressFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkGzip_CompressFileAsync(z.ckObj, cstr1, cstr2)

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


// Compresses a file to create a GZip compressed file (.gz). inFilename is the actual
// filename on disk. embeddedFilename is the filename to be embedded in the .gz such that when
// it is un-gzipped, this is the name of the file that will be created.
func (z *Gzip) CompressFile2(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkGzip_CompressFile2(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the CompressFile2 method
func (z *Gzip) CompressFile2Async(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkGzip_CompressFile2Async(z.ckObj, cstr1, cstr2, cstr3)

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


// Gzip compresses a file to an in-memory image of a .gz file.
// 
// Note: There is a 4GB size limitation. The compressed size of the file cannot be
// more than 4GB. Chilkat will be working to alleviate this limitation in the
// future.
//
func (z *Gzip) CompressFileToMem(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkGzip_CompressFileToMem(z.ckObj, cstr1, ckOutBytes)

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
func (z *Gzip) CompressFileToMemAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkGzip_CompressFileToMemAsync(z.ckObj, cstr1)

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


// Compresses in-memory data to an in-memory image of a .gz file.
// 
// Note: There is a 4GB uncompressed size limitation. The uncompressed size of the
// file cannot be more than 4GB. Chilkat will be working to alleviate this
// limitation in the future.
//
func (z *Gzip) CompressMemory(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkGzip_CompressMemory(z.ckObj, ckbyd1, ckOutBytes)

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
func (z *Gzip) CompressMemoryAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkGzip_CompressMemoryAsync(z.ckObj, ckbyd1)

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


// Gzip compresses and creates a .gz file from in-memory data.
func (z *Gzip) CompressMemToFile(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkGzip_CompressMemToFile(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CompressMemToFile method
func (z *Gzip) CompressMemToFileAsync(arg1 []byte, arg2 string, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    hTask := C.CkGzip_CompressMemToFileAsync(z.ckObj, ckbyd1, cstr2)

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


// Gzip compresses a string and writes the output to a byte array. The string is
// first converted to the charset specified by destCharset. Typical charsets are "utf-8",
// "iso-8859-1", "shift_JIS", etc.
func (z *Gzip) CompressString(arg1 string, arg2 string) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkGzip_CompressString(z.ckObj, cstr1, cstr2, ckOutBytes)

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


// Asynchronous version of the CompressString method
func (z *Gzip) CompressStringAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkGzip_CompressStringAsync(z.ckObj, cstr1, cstr2)

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


// The same as CompressString, except the binary output is returned in encoded
// string form according to the encoding. The encoding can be any of the following:
// "Base64", "modBase64", "Base32", "UU", "QP" (for quoted-printable), "URL" (for
// url-encoding), "Hex", "Q", "B", "url_oath", "url_rfc1738", "url_rfc2396", and
// "url_rfc3986".
// return a string or nil if failed.
func (z *Gzip) CompressStringENC(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkGzip_compressStringENC(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gzip compresses a string and writes the output to a .gz compressed file. The
// string is first converted to the charset specified by destCharset. Typical charsets are
// "utf-8", "iso-8859-1", "shift_JIS", etc.
func (z *Gzip) CompressStringToFile(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkGzip_CompressStringToFile(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the CompressStringToFile method
func (z *Gzip) CompressStringToFileAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkGzip_CompressStringToFileAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Decodes an encoded string and returns the original data. The encoding mode is
// determined by encoding. It may be "base64", "hex", "quoted-printable", or "url".
func (z *Gzip) Decode(arg1 string, arg2 string) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkGzip_Decode(z.ckObj, cstr1, cstr2, ckOutBytes)

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


// Provides the ability to use the zip/gzip's deflate algorithm directly to
// compress a string. Internal to this method, inString is first converted to the
// charset specified by charsetName. The data is then compressed using the deflate
// compression algorithm. The binary output is then encoded according to outputEncoding.
// Possible values for outputEncoding are "hex", "base64", "url", and "quoted-printable".
// 
// Note: The output of this method is compressed data with no Gzip file format. Use
// the Compress* methods to produce Gzip file format output.
//
// return a string or nil if failed.
func (z *Gzip) DeflateStringENC(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkGzip_deflateStringENC(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Encodes binary data to a printable string. The encoding mode is determined by
// encoding. It may be "base64", "hex", "quoted-printable", or "url".
// return a string or nil if failed.
func (z *Gzip) Encode(arg1 []byte, arg2 string) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    retStr := C.CkGzip_encode(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Determines if the inGzFilename is a Gzip formatted file. Returns true if it is a Gzip
// formatted file, otherwise returns false.
func (z *Gzip) ExamineFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkGzip_ExamineFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Determines if the in-memory bytes (inGzData) contain a Gzip formatted file. Returns
// true if it is Gzip format, otherwise returns false.
func (z *Gzip) ExamineMemory(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkGzip_ExamineMemory(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Gets the last-modification date/time to be embedded within the .gz.
func (z *Gzip) GetDt() *CkDateTime {

    retObj := C.CkGzip_GetDt(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// This the reverse of DeflateStringENC.
// 
// The input string is first decoded according to inputEncoding. (Possible values for inputEncoding
// are "hex", "base64", "url", and "quoted-printable".) The compressed data is then
// inflated, and the result is then converted from convertFromCharset (if necessary) to return a
// string.
//
// return a string or nil if failed.
func (z *Gzip) InflateStringENC(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkGzip_inflateStringENC(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if the component has been unlocked.
func (z *Gzip) IsUnlocked() bool {

    v := C.CkGzip_IsUnlocked(z.ckObj)


    return v != 0
}


// Loads the caller of the task's async method.
func (z *Gzip) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkGzip_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Reads a binary file into memory and returns the byte array. Note: This method
// does not parse a Gzip, it is only a convenience method for reading a binary file
// into memory.
func (z *Gzip) ReadFile(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkGzip_ReadFile(z.ckObj, cstr1, ckOutBytes)

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


// Sets the last-modification date/time to be embedded within the .gz when a
// Compress* method is called. If not explicitly set, the current system date/time
// is used.
func (z *Gzip) SetDt(arg1 *CkDateTime) bool {

    v := C.CkGzip_SetDt(z.ckObj, arg1.ckObj)


    return v != 0
}


// In-place ungzip the contents of binDat.
func (z *Gzip) UncompressBd(arg1 *BinData) bool {

    v := C.CkGzip_UncompressBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the UncompressBd method
func (z *Gzip) UncompressBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkGzip_UncompressBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Un-Gzips a .gz file. The output filename is specified by the 2nd argument and
// not by the filename embedded within the .gz.
func (z *Gzip) UncompressFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkGzip_UncompressFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UncompressFile method
func (z *Gzip) UncompressFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkGzip_UncompressFileAsync(z.ckObj, cstr1, cstr2)

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


// Un-Gzips a .gz file directly to memory.
// 
// Note: There is a 4GB uncompressed size limitation. The uncompressed size of the
// file cannot be more than 4GB. Chilkat will be working to alleviate this
// limitation in the future.
//
func (z *Gzip) UncompressFileToMem(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkGzip_UncompressFileToMem(z.ckObj, cstr1, ckOutBytes)

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
func (z *Gzip) UncompressFileToMemAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkGzip_UncompressFileToMemAsync(z.ckObj, cstr1)

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


// Uncompresses a .gz file that contains a text file. The contents of the text file
// are returned as a string. The character encoding of the text file is specified
// by charset. Typical charsets are "iso-8859-1", "utf-8", "windows-1252",
// "shift_JIS", "big5", etc.
// return a string or nil if failed.
func (z *Gzip) UncompressFileToString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkGzip_uncompressFileToString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the UncompressFileToString method
func (z *Gzip) UncompressFileToStringAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkGzip_UncompressFileToStringAsync(z.ckObj, cstr1, cstr2)

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


// Un-Gzips from an in-memory image of a .gz file directly into memory.
// 
// Note: There is a 4GB uncompressed size limitation. The uncompressed size of the
// file cannot be more than 4GB. Chilkat will be working to alleviate this
// limitation in the future.
//
func (z *Gzip) UncompressMemory(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkGzip_UncompressMemory(z.ckObj, ckbyd1, ckOutBytes)

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
func (z *Gzip) UncompressMemoryAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkGzip_UncompressMemoryAsync(z.ckObj, ckbyd1)

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


// Un-Gzips from an in-memory image of a .gz file to a file.
func (z *Gzip) UncompressMemToFile(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkGzip_UncompressMemToFile(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UncompressMemToFile method
func (z *Gzip) UncompressMemToFileAsync(arg1 []byte, arg2 string, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    hTask := C.CkGzip_UncompressMemToFileAsync(z.ckObj, ckbyd1, cstr2)

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


// The reverse of CompressString.
// 
// The bytes in inData are uncompressed, then converted from inCharset (if necessary) to
// return a string.
//
// return a string or nil if failed.
func (z *Gzip) UncompressString(arg1 []byte, arg2 string) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    retStr := C.CkGzip_uncompressString(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the UncompressString method
func (z *Gzip) UncompressStringAsync(arg1 []byte, arg2 string, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    hTask := C.CkGzip_UncompressStringAsync(z.ckObj, ckbyd1, cstr2)

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


// The same as UncompressString, except the compressed data is provided in encoded
// string form based on the value of encoding. The encoding can be "Base64", "modBase64",
// "Base32", "UU", "QP" (for quoted-printable), "URL" (for url-encoding), "Hex",
// "Q", "B", "url_oath", "url_rfc1738", "url_rfc2396", and "url_rfc3986".
// return a string or nil if failed.
func (z *Gzip) UncompressStringENC(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkGzip_uncompressStringENC(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Unlocks the component allowing for the full functionality to be used.
func (z *Gzip) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkGzip_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unpacks a .tar.gz file. The ungzip and untar occur in streaming mode. There are
// no temporary files and the memory footprint is constant (and small) while
// untarring. bNoAbsolute may be true or false. A value of true protects from
// untarring to absolute paths. (For example, imagine the trouble if the tar
// contains files with absolute paths beginning with /Windows/system32)
func (z *Gzip) UnTarGz(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkGzip_UnTarGz(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UnTarGz method
func (z *Gzip) UnTarGzAsync(arg1 string, arg2 string, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkGzip_UnTarGzAsync(z.ckObj, cstr1, cstr2, b3)

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


// A convenience method for writing a binary byte array to a file.
func (z *Gzip) WriteFile(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkGzip_WriteFile(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Converts base64-gzip .xfdl data to a decompressed XML string. The xfldData contains
// the base64 data. This method returns the decoded/decompressed XML string.
// return a string or nil if failed.
func (z *Gzip) XfdlToXml(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkGzip_xfdlToXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


