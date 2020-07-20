// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkCompression.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCompression() *Compression {
	return &Compression{C.CkCompression_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Compression) DisposeCompression() {
    if z != nil {
	C.CkCompression_Dispose(z.ckObj)
	}
}




func (z *Compression) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkCompression_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Compression) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkCompression_setExternalProgress(z.ckObj,1)
}

func (z *Compression) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkCompression_setExternalProgress(z.ckObj,1)
}

func (z *Compression) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkCompression_setExternalProgress(z.ckObj,1)
}




// property: Algorithm
// The compression algorithm to be used. Should be set to either "ppmd", "deflate",
// "zlib", "bzip2", or "lzw".
// 
// Note: The PPMD compression algorithm is only available for 32-bit builds. It is
// not available for 64-bit implementations. Also, this PPMD implementation is the
// "J" variant.
//
func (z *Compression) Algorithm() string {
    return C.GoString(C.CkCompression_algorithm(z.ckObj))
}
// property setter: Algorithm
func (z *Compression) SetAlgorithm(goStr string) {
    cStr := C.CString(goStr)
    C.CkCompression_putAlgorithm(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Charset
// Only applies to methods that compress or decompress strings. This specifies the
// character encoding that the string should be converted to before compression.
// Many programming languages use Unicode (2 bytes per char) for representing
// characters. This property allows for the string to be converted to a 1-byte per
// char encoding prior to compression.
func (z *Compression) Charset() string {
    return C.GoString(C.CkCompression_charset(z.ckObj))
}
// property setter: Charset
func (z *Compression) SetCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkCompression_putCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
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
func (z *Compression) DebugLogFilePath() string {
    return C.GoString(C.CkCompression_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Compression) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkCompression_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DeflateLevel
// This property allows for customization of the compression level for the
// "deflate" and "zlib" compression algoirthms. ("zlib" is just the deflate
// algorithm with a zlib header.) A value of 0 = no compression, while 9 = maximum
// compression. The default is 6.
func (z *Compression) DeflateLevel() int {
    return int(C.CkCompression_getDeflateLevel(z.ckObj))
}
// property setter: DeflateLevel
func (z *Compression) SetDeflateLevel(value int) {
    C.CkCompression_putDeflateLevel(z.ckObj,C.int(value))
}

// property: EncodingMode
// Controls the encoding expected by methods ending in "ENC", such as
// CompressBytesENC. Valid values are "base64", "hex", "url", and
// "quoted-printable". Compression methods ending in ENC return the binary
// compressed data as an encoded string using this encoding. Decompress methods
// expect the input string to be this encoding.
func (z *Compression) EncodingMode() string {
    return C.GoString(C.CkCompression_encodingMode(z.ckObj))
}
// property setter: EncodingMode
func (z *Compression) SetEncodingMode(goStr string) {
    cStr := C.CString(goStr)
    C.CkCompression_putEncodingMode(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// The number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any method call prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
func (z *Compression) HeartbeatMs() int {
    return int(C.CkCompression_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Compression) SetHeartbeatMs(value int) {
    C.CkCompression_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Compression) LastErrorHtml() string {
    return C.GoString(C.CkCompression_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Compression) LastErrorText() string {
    return C.GoString(C.CkCompression_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Compression) LastErrorXml() string {
    return C.GoString(C.CkCompression_lastErrorXml(z.ckObj))
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
func (z *Compression) LastMethodSuccess() bool {
    v := int(C.CkCompression_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Compression) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCompression_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Compression) VerboseLogging() bool {
    v := int(C.CkCompression_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Compression) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCompression_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Compression) Version() string {
    return C.GoString(C.CkCompression_version(z.ckObj))
}
// Large amounts of binary byte data may be compressed in chunks by first calling
// BeginCompressBytes, followed by 0 or more calls to MoreCompressedBytes, and
// ending with a final call to EndCompressBytes. Each call returns 0 or more bytes
// of compressed data which may be output to a compressed data stream (such as a
// file, socket, etc.).
func (z *Compression) BeginCompressBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_BeginCompressBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the BeginCompressBytes method
func (z *Compression) BeginCompressBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_BeginCompressBytesAsync(z.ckObj, ckbyd1)

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


// Large amounts of binary byte data may be compressed in chunks by first calling
// BeginCompressBytesENC, followed by 0 or more calls to MoreCompressedBytesENC,
// and ending with a final call to EndCompressBytesENC. Each call returns 0 or more
// characters of compressed data (encoded as a string according to the EncodingMode
// property setting) which may be output to a compressed data stream (such as a
// file, socket, etc.).
// return a string or nil if failed.
func (z *Compression) BeginCompressBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCompression_beginCompressBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the BeginCompressBytesENC method
func (z *Compression) BeginCompressBytesENCAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_BeginCompressBytesENCAsync(z.ckObj, ckbyd1)

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


// Large amounts of string data may be compressed in chunks by first calling
// BeginCompressString, followed by 0 or more calls to MoreCompressedString, and
// ending with a final call to EndCompressString. Each call returns 0 or more bytes
// of compressed data which may be output to a compressed data stream (such as a
// file, socket, etc.).
func (z *Compression) BeginCompressString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_BeginCompressString(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the BeginCompressString method
func (z *Compression) BeginCompressStringAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_BeginCompressStringAsync(z.ckObj, cstr1)

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


// Large amounts of string data may be compressed in chunks by first calling
// BeginCompressStringENC, followed by 0 or more calls to MoreCompressedStringENC,
// and ending with a final call to EndCompressStringENC. Each call returns 0 or
// more characters of compressed data (encoded as a string according to the
// EncodingMode property setting) which may be output to a compressed data stream
// (such as a file, socket, etc.).
// return a string or nil if failed.
func (z *Compression) BeginCompressStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCompression_beginCompressStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the BeginCompressStringENC method
func (z *Compression) BeginCompressStringENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_BeginCompressStringENCAsync(z.ckObj, cstr1)

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


// A compressed data stream may be decompressed in chunks by first calling
// BeginDecompressBytes, followed by 0 or more calls to MoreDecompressedBytes, and
// ending with a final call to EndDecompressBytes. Each call returns 0 or more
// bytes of decompressed data.
func (z *Compression) BeginDecompressBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_BeginDecompressBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the BeginDecompressBytes method
func (z *Compression) BeginDecompressBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_BeginDecompressBytesAsync(z.ckObj, ckbyd1)

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


// The input to this method is an encoded string containing compressed data. The
// EncodingMode property should be set prior to calling this method. The input
// string is decoded according to the EncodingMode (hex, base64, etc.) and then
// decompressed.
// 
// A compressed data stream may be decompressed in chunks by first calling
// BeginDecompressBytesENC, followed by 0 or more calls to
// MoreDecompressedBytesENC, and ending with a final call to EndDecompressBytesENC.
// Each call returns 0 or more bytes of decompressed data.
//
func (z *Compression) BeginDecompressBytesENC(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_BeginDecompressBytesENC(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the BeginDecompressBytesENC method
func (z *Compression) BeginDecompressBytesENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_BeginDecompressBytesENCAsync(z.ckObj, cstr1)

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


// A compressed data stream may be decompressed in chunks by first calling
// BeginDecompressString, followed by 0 or more calls to MoreDecompressedString,
// and ending with a final call to EndDecompressString. Each call returns 0 or more
// characters of decompressed text.
// return a string or nil if failed.
func (z *Compression) BeginDecompressString(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCompression_beginDecompressString(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the BeginDecompressString method
func (z *Compression) BeginDecompressStringAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_BeginDecompressStringAsync(z.ckObj, ckbyd1)

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


// The input to this method is an encoded string containing compressed data. The
// EncodingMode property should be set prior to calling this method. The input
// string is decoded according to the EncodingMode (hex, base64, etc.) and then
// decompressed.
// 
// A compressed data stream may be decompressed in chunks by first calling
// BeginDecompressStringENC, followed by 0 or more calls to
// MoreDecompressedStringENC, and ending with a final call to
// EndDecompressStringENC. Each call returns 0 or more characters of decompressed
// text.
//
// return a string or nil if failed.
func (z *Compression) BeginDecompressStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCompression_beginDecompressStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the BeginDecompressStringENC method
func (z *Compression) BeginDecompressStringENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_BeginDecompressStringENCAsync(z.ckObj, cstr1)

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


// Compresses the data contained in a BinData object.
func (z *Compression) CompressBd(arg1 *BinData) bool {

    v := C.CkCompression_CompressBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the CompressBd method
func (z *Compression) CompressBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkCompression_CompressBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Compresses byte data.
func (z *Compression) CompressBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_CompressBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the CompressBytes method
func (z *Compression) CompressBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_CompressBytesAsync(z.ckObj, ckbyd1)

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


// Compresses bytes and returns the compressed data encoded to a string. The
// encoding (hex, base64, etc.) is determined by the EncodingMode property setting.
// return a string or nil if failed.
func (z *Compression) CompressBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCompression_compressBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the CompressBytesENC method
func (z *Compression) CompressBytesENCAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_CompressBytesENCAsync(z.ckObj, ckbyd1)

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


// Performs file-to-file compression. Files of any size may be compressed because
// the file is compressed internally in streaming mode.
func (z *Compression) CompressFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCompression_CompressFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CompressFile method
func (z *Compression) CompressFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkCompression_CompressFileAsync(z.ckObj, cstr1, cstr2)

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


// Compresses the contents of sb and appends the compressed bytes to binData.
func (z *Compression) CompressSb(arg1 *StringBuilder, arg2 *BinData) bool {

    v := C.CkCompression_CompressSb(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Asynchronous version of the CompressSb method
func (z *Compression) CompressSbAsync(arg1 *StringBuilder, arg2 *BinData, c chan *Task) {

    hTask := C.CkCompression_CompressSbAsync(z.ckObj, arg1.ckObj, arg2.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Compresses a stream. Internally, the strm's source is read, compressed, and the
// compressed data written to the strm's sink. It does this in streaming fashion.
// Extremely large or even infinite streams can be compressed with stable ungrowing
// memory usage.
func (z *Compression) CompressStream(arg1 *Stream) bool {

    v := C.CkCompression_CompressStream(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the CompressStream method
func (z *Compression) CompressStreamAsync(arg1 *Stream, c chan *Task) {

    hTask := C.CkCompression_CompressStreamAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Compresses a string.
func (z *Compression) CompressString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_CompressString(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the CompressString method
func (z *Compression) CompressStringAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_CompressStringAsync(z.ckObj, cstr1)

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


// Compresses a string and returns the compressed data encoded to a string. The
// output encoding (hex, base64, etc.) is determined by the EncodingMode property
// setting.
// return a string or nil if failed.
func (z *Compression) CompressStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCompression_compressStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the CompressStringENC method
func (z *Compression) CompressStringENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_CompressStringENCAsync(z.ckObj, cstr1)

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


// Decompresses the data contained in a BinData object.
func (z *Compression) DecompressBd(arg1 *BinData) bool {

    v := C.CkCompression_DecompressBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the DecompressBd method
func (z *Compression) DecompressBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkCompression_DecompressBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The opposite of CompressBytes.
func (z *Compression) DecompressBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_DecompressBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the DecompressBytes method
func (z *Compression) DecompressBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_DecompressBytesAsync(z.ckObj, ckbyd1)

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


// The opposite of CompressBytesENC. encodedCompressedData contains the compressed data as an
// encoded string (hex, base64, etc) as specified by the EncodingMode property
// setting.
func (z *Compression) DecompressBytesENC(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_DecompressBytesENC(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the DecompressBytesENC method
func (z *Compression) DecompressBytesENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_DecompressBytesENCAsync(z.ckObj, cstr1)

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


// Performs file-to-file decompression (the opposite of CompressFile). Internally
// the file is decompressed in streaming mode which allows files of any size to be
// decompressed.
func (z *Compression) DecompressFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCompression_DecompressFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the DecompressFile method
func (z *Compression) DecompressFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkCompression_DecompressFileAsync(z.ckObj, cstr1, cstr2)

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


// Decompresses the contents of binData and appends the decompressed string to sb.
func (z *Compression) DecompressSb(arg1 *BinData, arg2 *StringBuilder) bool {

    v := C.CkCompression_DecompressSb(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Asynchronous version of the DecompressSb method
func (z *Compression) DecompressSbAsync(arg1 *BinData, arg2 *StringBuilder, c chan *Task) {

    hTask := C.CkCompression_DecompressSbAsync(z.ckObj, arg1.ckObj, arg2.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Decompresses a stream. Internally, the strm's source is read, decompressed, and
// the decompressed data written to the strm's sink. It does this in streaming
// fashion. Extremely large or even infinite streams can be decompressed with
// stable ungrowing memory usage.
func (z *Compression) DecompressStream(arg1 *Stream) bool {

    v := C.CkCompression_DecompressStream(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the DecompressStream method
func (z *Compression) DecompressStreamAsync(arg1 *Stream, c chan *Task) {

    hTask := C.CkCompression_DecompressStreamAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Takes compressed bytes, decompresses, and returns the resulting string.
// return a string or nil if failed.
func (z *Compression) DecompressString(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCompression_decompressString(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the DecompressString method
func (z *Compression) DecompressStringAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_DecompressStringAsync(z.ckObj, ckbyd1)

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


// The opposite of CompressStringENC. encodedCompressedData contains the compressed data as an
// encoded string (hex, base64, etc) as specified by the EncodingMode property
// setting.
// return a string or nil if failed.
func (z *Compression) DecompressStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCompression_decompressStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the DecompressStringENC method
func (z *Compression) DecompressStringENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_DecompressStringENCAsync(z.ckObj, cstr1)

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


// Must be callled to finalize a compression stream. Returns any remaining
// (buffered) compressed data.
// 
// (See BeginCompressBytes)
//
func (z *Compression) EndCompressBytes() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_EndCompressBytes(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the EndCompressBytes method
func (z *Compression) EndCompressBytesAsync(c chan *Task) {

    hTask := C.CkCompression_EndCompressBytesAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Must be callled to finalize a compression stream. Returns any remaining
// (buffered) compressed data.
// 
// (See BeginCompressBytesENC)
//
// return a string or nil if failed.
func (z *Compression) EndCompressBytesENC() *string {

    retStr := C.CkCompression_endCompressBytesENC(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the EndCompressBytesENC method
func (z *Compression) EndCompressBytesENCAsync(c chan *Task) {

    hTask := C.CkCompression_EndCompressBytesENCAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Must be callled to finalize a compression stream. Returns any remaining
// (buffered) compressed data.
// 
// (See BeginCompressString)
//
func (z *Compression) EndCompressString() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_EndCompressString(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the EndCompressString method
func (z *Compression) EndCompressStringAsync(c chan *Task) {

    hTask := C.CkCompression_EndCompressStringAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Must be callled to finalize a compression stream. Returns any remaining
// (buffered) compressed data.
// 
// (See BeginCompressStringENC)
//
// return a string or nil if failed.
func (z *Compression) EndCompressStringENC() *string {

    retStr := C.CkCompression_endCompressStringENC(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the EndCompressStringENC method
func (z *Compression) EndCompressStringENCAsync(c chan *Task) {

    hTask := C.CkCompression_EndCompressStringENCAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Called to finalize the decompression stream and return any remaining (buffered)
// decompressed data.
// 
// (See BeginDecompressBytes)
//
func (z *Compression) EndDecompressBytes() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_EndDecompressBytes(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the EndDecompressBytes method
func (z *Compression) EndDecompressBytesAsync(c chan *Task) {

    hTask := C.CkCompression_EndDecompressBytesAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Called to finalize the decompression stream and return any remaining (buffered)
// decompressed data.
// 
// The input to this method is an encoded string containing compressed data. The
// EncodingMode property should be set prior to calling this method. The input
// string is decoded according to the EncodingMode (hex, base64, etc.) and then
// decompressed.
// 
// (See BeginDecompressBytesENC)
//
func (z *Compression) EndDecompressBytesENC() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_EndDecompressBytesENC(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the EndDecompressBytesENC method
func (z *Compression) EndDecompressBytesENCAsync(c chan *Task) {

    hTask := C.CkCompression_EndDecompressBytesENCAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Called to finalize the decompression stream and return any remaining (buffered)
// decompressed data.
// 
// (See BeginDecompressString)
//
// return a string or nil if failed.
func (z *Compression) EndDecompressString() *string {

    retStr := C.CkCompression_endDecompressString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the EndDecompressString method
func (z *Compression) EndDecompressStringAsync(c chan *Task) {

    hTask := C.CkCompression_EndDecompressStringAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Called to finalize the decompression stream and return any remaining (buffered)
// decompressed data.
// 
// The input to this method is an encoded string containing compressed data. The
// EncodingMode property should be set prior to calling this method. The input
// string is decoded according to the EncodingMode (hex, base64, etc.) and then
// decompressed.
// 
// (See BeginDecompressStringENC)
//
// return a string or nil if failed.
func (z *Compression) EndDecompressStringENC() *string {

    retStr := C.CkCompression_endDecompressStringENC(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the EndDecompressStringENC method
func (z *Compression) EndDecompressStringENCAsync(c chan *Task) {

    hTask := C.CkCompression_EndDecompressStringENCAsync(z.ckObj)


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
func (z *Compression) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkCompression_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// (See BeginCompressBytes)
func (z *Compression) MoreCompressBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_MoreCompressBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the MoreCompressBytes method
func (z *Compression) MoreCompressBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_MoreCompressBytesAsync(z.ckObj, ckbyd1)

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


// (See BeginCompressBytesENC)
// return a string or nil if failed.
func (z *Compression) MoreCompressBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCompression_moreCompressBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the MoreCompressBytesENC method
func (z *Compression) MoreCompressBytesENCAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_MoreCompressBytesENCAsync(z.ckObj, ckbyd1)

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


// (See BeginCompressString)
func (z *Compression) MoreCompressString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_MoreCompressString(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the MoreCompressString method
func (z *Compression) MoreCompressStringAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_MoreCompressStringAsync(z.ckObj, cstr1)

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


// (See BeginCompressStringENC)
// return a string or nil if failed.
func (z *Compression) MoreCompressStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCompression_moreCompressStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the MoreCompressStringENC method
func (z *Compression) MoreCompressStringENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_MoreCompressStringENCAsync(z.ckObj, cstr1)

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


// (See BeginDecompressBytes)
func (z *Compression) MoreDecompressBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_MoreDecompressBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the MoreDecompressBytes method
func (z *Compression) MoreDecompressBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_MoreDecompressBytesAsync(z.ckObj, ckbyd1)

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


// The input to this method is an encoded string containing compressed data. The
// EncodingMode property should be set prior to calling this method. The input
// string is decoded according to the EncodingMode (hex, base64, etc.) and then
// decompressed.
// 
// (See BeginDecompressBytesENC)
//
func (z *Compression) MoreDecompressBytesENC(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCompression_MoreDecompressBytesENC(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the MoreDecompressBytesENC method
func (z *Compression) MoreDecompressBytesENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_MoreDecompressBytesENCAsync(z.ckObj, cstr1)

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


// (See BeginDecompressString)
// return a string or nil if failed.
func (z *Compression) MoreDecompressString(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCompression_moreDecompressString(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the MoreDecompressString method
func (z *Compression) MoreDecompressStringAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCompression_MoreDecompressStringAsync(z.ckObj, ckbyd1)

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


// The input to this method is an encoded string containing compressed data. The
// EncodingMode property should be set prior to calling this method. The input
// string is decoded according to the EncodingMode (hex, base64, etc.) and then
// decompressed.
// 
// (See BeginDecompressStringENC)
//
// return a string or nil if failed.
func (z *Compression) MoreDecompressStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCompression_moreDecompressStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the MoreDecompressStringENC method
func (z *Compression) MoreDecompressStringENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCompression_MoreDecompressStringENCAsync(z.ckObj, cstr1)

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


// Unlocks the component allowing for the full functionality to be used. The
// component may be used fully-functional for the 1st 30-days after download by
// passing an arbitrary string to this method. If for some reason you do not
// receive the full 30-day trial, send email to support@chilkatsoft.com for a
// temporary unlock code w/ an explicit expiration date. Upon purchase, a purchased
// unlock code is provided which should replace the temporary/arbitrary string
// passed to this method.
func (z *Compression) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCompression_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


