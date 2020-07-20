// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkZipEntry.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewZipEntry() *ZipEntry {
	return &ZipEntry{C.CkZipEntry_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *ZipEntry) DisposeZipEntry() {
    if z != nil {
	C.CkZipEntry_Dispose(z.ckObj)
	}
}




func (z *ZipEntry) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkZipEntry_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *ZipEntry) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkZipEntry_setExternalProgress(z.ckObj,1)
}

func (z *ZipEntry) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkZipEntry_setExternalProgress(z.ckObj,1)
}

func (z *ZipEntry) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkZipEntry_setExternalProgress(z.ckObj,1)
}




// property: Comment
// The comment stored within the Zip for this entry.
func (z *ZipEntry) Comment() string {
    return C.GoString(C.CkZipEntry_comment(z.ckObj))
}
// property setter: Comment
func (z *ZipEntry) SetComment(goStr string) {
    cStr := C.CString(goStr)
    C.CkZipEntry_putComment(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CompressedLength
// The size in bytes of this entry's file data when compression is applied
func (z *ZipEntry) CompressedLength() uint {
    return uint(C.CkZipEntry_getCompressedLength(z.ckObj))
}

// property: CompressedLength64
// The size in bytes of this entry's file data when compression is applied
func (z *ZipEntry) CompressedLength64() int64 {
    return int64(C.CkZipEntry_getCompressedLength64(z.ckObj))
}

// property: CompressionLevel
// The compression level. 0 = no compression, while 9 = maximum compression. The
// default is 6.
func (z *ZipEntry) CompressionLevel() int {
    return int(C.CkZipEntry_getCompressionLevel(z.ckObj))
}
// property setter: CompressionLevel
func (z *ZipEntry) SetCompressionLevel(value int) {
    C.CkZipEntry_putCompressionLevel(z.ckObj,C.int(value))
}

// property: CompressionMethod
// Set to 0 for no compression, or 8 for the Deflate algorithm. The Deflate
// algorithm is the default algorithm of the most popular Zip utility programs,
// such as WinZip
func (z *ZipEntry) CompressionMethod() int {
    return int(C.CkZipEntry_getCompressionMethod(z.ckObj))
}
// property setter: CompressionMethod
func (z *ZipEntry) SetCompressionMethod(value int) {
    C.CkZipEntry_putCompressionMethod(z.ckObj,C.int(value))
}

// property: Crc
// The CRC for the zip entry. For AES encrypted entries, the CRC value will be 0.
// (See http://www.winzip.com/aes_info.htm#CRC )
func (z *ZipEntry) Crc() int {
    return int(C.CkZipEntry_getCrc(z.ckObj))
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
func (z *ZipEntry) DebugLogFilePath() string {
    return C.GoString(C.CkZipEntry_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *ZipEntry) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkZipEntry_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EncryptionKeyLen
// If this entry is AES encrypted, then this property contains the AES key length
// (128, 192, or 256). If the entry is not AES encrypted, then the value is 0.
func (z *ZipEntry) EncryptionKeyLen() int {
    return int(C.CkZipEntry_getEncryptionKeyLen(z.ckObj))
}

// property: EntryID
// The unique ID assigned by Zip.NET while the object is instantiated in memory.
func (z *ZipEntry) EntryID() int {
    return int(C.CkZipEntry_getEntryID(z.ckObj))
}

// property: EntryType
// Indicates the origin of the entry. There are three possible values:
//     Mapped Entry: An entry in an existing Zip file.
//     File Entry: A file not yet in memory, but referenced. These entries are
//     added by calling AppendFiles, AppendFilesEx, AppendOneFileOrDir, etc.
//     Data Entry: An entry containing uncompressed data from memory. These entries
//     are added by calling AppendData, AppendString, etc.
//     Null Entry: An entry that no longer exists in the .zip.
//     New Directory Entry: A directory entry added by calling AppendNewDir.
// When the zip is written by calling WriteZip or WriteToMemory, all of the zip
// entries are transformed into mapped entries. They become entries that contain
// the compressed data within the .zip that was just created. (The WriteZip method
// call effectively writes the zip and then opens it, replacing all of the existing
// entries with mapped entries.)
func (z *ZipEntry) EntryType() int {
    return int(C.CkZipEntry_getEntryType(z.ckObj))
}

// property: FileDateTimeStr
// The local last-modified date/time in RFC822 string format.
func (z *ZipEntry) FileDateTimeStr() string {
    return C.GoString(C.CkZipEntry_fileDateTimeStr(z.ckObj))
}
// property setter: FileDateTimeStr
func (z *ZipEntry) SetFileDateTimeStr(goStr string) {
    cStr := C.CString(goStr)
    C.CkZipEntry_putFileDateTimeStr(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: FileName
// The file name of the Zip entry.
func (z *ZipEntry) FileName() string {
    return C.GoString(C.CkZipEntry_fileName(z.ckObj))
}
// property setter: FileName
func (z *ZipEntry) SetFileName(goStr string) {
    cStr := C.CString(goStr)
    C.CkZipEntry_putFileName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: FileNameHex
// A string containing the hex encoded raw filename bytes found in the Zip entry.
func (z *ZipEntry) FileNameHex() string {
    return C.GoString(C.CkZipEntry_fileNameHex(z.ckObj))
}

// property: HeartbeatMs
// The number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort inflate/extract/unzip calls
// prior to completion. If HeartbeatMs is 0 (the default), no AbortCheck event
// callbacks will fire.
func (z *ZipEntry) HeartbeatMs() int {
    return int(C.CkZipEntry_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *ZipEntry) SetHeartbeatMs(value int) {
    C.CkZipEntry_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: IsAesEncrypted
// true if the entry is AES encrypted. This property can only be true for
// entries already contained in a .zip (i.e. entries obtained from a zip archive
// that was opened via OpenZip, OpenBd, OpenFromMemory, etc.) The property is
// false if the entry contained in the zip is not AES encrypted.
func (z *ZipEntry) IsAesEncrypted() bool {
    v := int(C.CkZipEntry_getIsAesEncrypted(z.ckObj))
    return v != 0
}

// property: IsDirectory
// True if the Zip entry is a directory, false if it is a file.
func (z *ZipEntry) IsDirectory() bool {
    v := int(C.CkZipEntry_getIsDirectory(z.ckObj))
    return v != 0
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *ZipEntry) LastErrorHtml() string {
    return C.GoString(C.CkZipEntry_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *ZipEntry) LastErrorText() string {
    return C.GoString(C.CkZipEntry_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *ZipEntry) LastErrorXml() string {
    return C.GoString(C.CkZipEntry_lastErrorXml(z.ckObj))
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
func (z *ZipEntry) LastMethodSuccess() bool {
    v := int(C.CkZipEntry_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *ZipEntry) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZipEntry_putLastMethodSuccess(z.ckObj,v)
}

// property: TextFlag
// Controls whether the "text flag" of the internal file attributes for this entry
// within the zip is set. This is a bit flag that indicates whether the file
// contents are text or binary. It is for informational use and it is not
// imperative that this bit flag is accurately set. The ability to set this bit
// flag is only provided to satisfy certain cases where another application might
// be sensitive to the flag.
func (z *ZipEntry) TextFlag() bool {
    v := int(C.CkZipEntry_getTextFlag(z.ckObj))
    return v != 0
}
// property setter: TextFlag
func (z *ZipEntry) SetTextFlag(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZipEntry_putTextFlag(z.ckObj,v)
}

// property: UncompressedLength
// The size in bytes of this entry's file data when uncompressed.
func (z *ZipEntry) UncompressedLength() uint {
    return uint(C.CkZipEntry_getUncompressedLength(z.ckObj))
}

// property: UncompressedLength64
// The size in bytes of this entry's file data when uncompressed.
func (z *ZipEntry) UncompressedLength64() int64 {
    return int64(C.CkZipEntry_getUncompressedLength64(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *ZipEntry) VerboseLogging() bool {
    v := int(C.CkZipEntry_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *ZipEntry) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZipEntry_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *ZipEntry) Version() string {
    return C.GoString(C.CkZipEntry_version(z.ckObj))
}
// Appends binary data to the zip entry's file contents.
func (z *ZipEntry) AppendData(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkZipEntry_AppendData(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Asynchronous version of the AppendData method
func (z *ZipEntry) AppendDataAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkZipEntry_AppendDataAsync(z.ckObj, ckbyd1)

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


// Appends text data to the zip entry's file contents. The text is appended using
// the character encoding specified by the charset, which can be "utf-8", "ansi", etc.
func (z *ZipEntry) AppendString(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkZipEntry_AppendString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AppendString method
func (z *ZipEntry) AppendStringAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkZipEntry_AppendStringAsync(z.ckObj, cstr1, cstr2)

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


// Returns the compressed data as a byte array.
// 
// Note: The Copy method can only be called if the zip entry already contains
// compressed data (i.e. it is a "mapped entry"). This is the case when an existing
// .zip is opened (from memory or from a file), or after the .zip has been written
// (by calling WriteZip or WriteToMemory). If a zip entry is created via
// AppendData, AppendFiles, etc., then it does not yet contain compressed data.
// When the zip is written, each entry becomes a "mapped entry" (The value of the
// EntryType property becomes 0.)
//
func (z *ZipEntry) Copy() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkZipEntry_Copy(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the compressed data as a Base64-encoded string. It is only possible to
// retrieve the compressed data from a pre-existing .zip that has been opened or
// after writing the .zip but not closing it.
// 
// Note: The CopyToBase64 method can only be called if the zip entry already
// contains compressed data (i.e. it is a "mapped entry").
//
// return a string or nil if failed.
func (z *ZipEntry) CopyToBase64() *string {

    retStr := C.CkZipEntry_copyToBase64(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the compressed data as a hexidecimal encoded string. It is only possible
// to retrieve the compressed data from a pre-existing .zip that has been opened or
// after writing the .zip but not closing it.
// 
// Note: The CopyToBase64 method can only be called if the zip entry already
// contains compressed data (i.e. it is a "mapped entry").
//
// return a string or nil if failed.
func (z *ZipEntry) CopyToHex() *string {

    retStr := C.CkZipEntry_copyToHex(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Unzips this entry into the specified base directory. The file is unzipped to the
// subdirectory according to the relative path stored with the entry in the Zip.
// Use ExtractInto to unzip into a specific directory regardless of the path
// information stored in the Zip.
func (z *ZipEntry) Extract(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZipEntry_Extract(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Extract method
func (z *ZipEntry) ExtractAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkZipEntry_ExtractAsync(z.ckObj, cstr1)

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


// Unzip a file into a specific directory. If this entry is a directory, then
// nothing occurs. (An application can check the IsDirectory property and instead
// call Extract if it is desired to create the directory. )
func (z *ZipEntry) ExtractInto(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZipEntry_ExtractInto(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the ExtractInto method
func (z *ZipEntry) ExtractIntoAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkZipEntry_ExtractIntoAsync(z.ckObj, cstr1)

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


// Returns the last-modified date/time of this zip entry.
func (z *ZipEntry) GetDt() *CkDateTime {

    retObj := C.CkZipEntry_GetDt(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Inflate a file within a Zip directly into memory.
func (z *ZipEntry) Inflate() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkZipEntry_Inflate(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the Inflate method
func (z *ZipEntry) InflateAsync(c chan *Task) {

    hTask := C.CkZipEntry_InflateAsync(z.ckObj)


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
func (z *ZipEntry) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkZipEntry_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Return the next entry (file or directory) within the Zip
func (z *ZipEntry) NextEntry() *ZipEntry {

    retObj := C.CkZipEntry_NextEntry(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Returns the next entry having a filename matching the matchStr. The "*" characters
// matches 0 or more of any character. The full filename, including path, is used
// when matching against the pattern.
func (z *ZipEntry) NextMatchingEntry(arg1 string) *ZipEntry {
    cstr1 := C.CString(arg1)

    retObj := C.CkZipEntry_NextMatchingEntry(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Replaces the zip entry's existing contents with new data.
func (z *ZipEntry) ReplaceData(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkZipEntry_ReplaceData(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Replaces the zip entry's existing contents with new text data. The text will be
// stored using the character encoding as specified by charset, which can be "utf-8",
// "ansi", etc.
func (z *ZipEntry) ReplaceString(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkZipEntry_ReplaceString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the last-modified date/time for this zip entry.
func (z *ZipEntry) SetDt(arg1 *CkDateTime)  {

    C.CkZipEntry_SetDt(z.ckObj, arg1.ckObj)



}


// Unzips the entry contents into the binData.
func (z *ZipEntry) UnzipToBd(arg1 *BinData) bool {

    v := C.CkZipEntry_UnzipToBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the UnzipToBd method
func (z *ZipEntry) UnzipToBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkZipEntry_UnzipToBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Unzips a text file to the sb. The contents of sb are appended with the
// unzipped file. The lineEndingBehavior is as follows:
// 
// 0 = leave unchanged.
// 1 = convert all to bare LF's
// 2 = convert all to CRLF's
// 
// The srcCharset tells the component how to interpret the bytes of the uncompressed file
// -- i.e. as utf-8, utf-16, windows-1252, etc.
func (z *ZipEntry) UnzipToSb(arg1 int, arg2 string, arg3 *StringBuilder) bool {
    cstr2 := C.CString(arg2)

    v := C.CkZipEntry_UnzipToSb(z.ckObj, C.int(arg1), cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UnzipToSb method
func (z *ZipEntry) UnzipToSbAsync(arg1 int, arg2 string, arg3 *StringBuilder, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkZipEntry_UnzipToSbAsync(z.ckObj, C.int(arg1), cstr2, arg3.ckObj)

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


// Unzips a file within a Zip to a stream. If called synchronously, the toStream must
// have a sink, such as a file or another stream object. If called asynchronously,
// then the foreground thread can read the stream.
func (z *ZipEntry) UnzipToStream(arg1 *Stream) bool {

    v := C.CkZipEntry_UnzipToStream(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the UnzipToStream method
func (z *ZipEntry) UnzipToStreamAsync(arg1 *Stream, c chan *Task) {

    hTask := C.CkZipEntry_UnzipToStreamAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Inflate and return the uncompressed data as a string The lineEndingBehavior is as follows:
// 
// 0 = leave unchanged.
// 1 = convert all to bare LF's
// 2 = convert all to CRLF's
// 
// The srcCharset tells the component how to interpret the bytes of the uncompressed file
// -- i.e. as utf-8, utf-16, windows-1252, etc.
// return a string or nil if failed.
func (z *ZipEntry) UnzipToString(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkZipEntry_unzipToString(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the UnzipToString method
func (z *ZipEntry) UnzipToStringAsync(arg1 int, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkZipEntry_UnzipToStringAsync(z.ckObj, C.int(arg1), cstr2)

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


