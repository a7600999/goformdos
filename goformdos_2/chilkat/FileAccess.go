// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkFileAccess.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewFileAccess() *FileAccess {
	return &FileAccess{C.CkFileAccess_Create()}
}

func (z *FileAccess) DisposeFileAccess() {
    if z != nil {
	C.CkFileAccess_Dispose(z.ckObj)
	}
}




// property: CurrentDir
// The current working directory of the calling process.
func (z *FileAccess) CurrentDir() string {
    return C.GoString(C.CkFileAccess_currentDir(z.ckObj))
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
func (z *FileAccess) DebugLogFilePath() string {
    return C.GoString(C.CkFileAccess_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *FileAccess) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkFileAccess_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EndOfFile
// Returns true if the current open file is at the end-of-file.
func (z *FileAccess) EndOfFile() bool {
    v := int(C.CkFileAccess_getEndOfFile(z.ckObj))
    return v != 0
}

// property: FileOpenError
// This property is set by the following methods: FileOpen, OpenForRead,
// OpenForWrite, OpenForReadWrite, and OpenForAppend. It provides an error code
// indicating the failure reason. Possible values are:
// 
//     Success (No error)
//     Access denied.
//     File not found.
//     General (non-specific) open error.
//     File aleady exists.
//     Path refers to a directory and the access requested involves writing.
//     Too many symbolic links were encountered in resolving path.
//     The process already has the maximum number of files open.
//     Pathname is too long.
//     The system limit on the total number of open files has been reached.
//     Pathname refers to a device special file and no corresponding device exists.
//     Insufficient kernel memory was available.
//     Pathname was to be created but the device containing pathname has no room
//     for the new file.
//     A component used as a directory in pathname is not, in fact, a directory.
//     Pathname refers to a regular file, too large to be opened (this would be a
//     limitation of the underlying operating system, not a limitation imposed by
//     Chilkat).
//     Pathname refers to a file on a read-only filesystem and write access was
//     requested.
//
func (z *FileAccess) FileOpenError() int {
    return int(C.CkFileAccess_getFileOpenError(z.ckObj))
}

// property: FileOpenErrorMsg
// The error message text associated with the FileOpenError code.
func (z *FileAccess) FileOpenErrorMsg() string {
    return C.GoString(C.CkFileAccess_fileOpenErrorMsg(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *FileAccess) LastErrorHtml() string {
    return C.GoString(C.CkFileAccess_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *FileAccess) LastErrorText() string {
    return C.GoString(C.CkFileAccess_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *FileAccess) LastErrorXml() string {
    return C.GoString(C.CkFileAccess_lastErrorXml(z.ckObj))
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
func (z *FileAccess) LastMethodSuccess() bool {
    v := int(C.CkFileAccess_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *FileAccess) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFileAccess_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *FileAccess) VerboseLogging() bool {
    v := int(C.CkFileAccess_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *FileAccess) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFileAccess_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *FileAccess) Version() string {
    return C.GoString(C.CkFileAccess_version(z.ckObj))
}
// Appends a string using the ANSI character encoding to the currently open file.
func (z *FileAccess) AppendAnsi(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_AppendAnsi(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends the contents of bd to the currently open file.
func (z *FileAccess) AppendBd(arg1 *BinData) bool {

    v := C.CkFileAccess_AppendBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Appends the contents of sb using the character encoding (such as "utf-8")
// specified by charset to the currently open file.
func (z *FileAccess) AppendSb(arg1 *StringBuilder, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkFileAccess_AppendSb(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends a string using the character encoding specified by str to the currently
// open file.
func (z *FileAccess) AppendText(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFileAccess_AppendText(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends the 2-byte Unicode BOM (little endian) to the currently open file.
func (z *FileAccess) AppendUnicodeBOM() bool {

    v := C.CkFileAccess_AppendUnicodeBOM(z.ckObj)


    return v != 0
}


// Appends the 3-byte utf-8 BOM to the currently open file.
func (z *FileAccess) AppendUtf8BOM() bool {

    v := C.CkFileAccess_AppendUtf8BOM(z.ckObj)


    return v != 0
}


// Same as DirEnsureExists, except the argument is a file path (the last part of
// the path is a filename and not a directory). Creates all missing directories
// such that filePath may be created.
func (z *FileAccess) DirAutoCreate(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_DirAutoCreate(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Creates a new directory specified by dirPath.
func (z *FileAccess) DirCreate(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_DirCreate(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Deletes the directory specified by dirPath. It is only possible to delete a
// directory if it contains no files or subdirectories.
func (z *FileAccess) DirDelete(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_DirDelete(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Creates all directories necessary such that the entire dirPath exists.
func (z *FileAccess) DirEnsureExists(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_DirEnsureExists(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Closes the currently open file.
func (z *FileAccess) FileClose()  {

    C.CkFileAccess_FileClose(z.ckObj)



}


// Compares the contents of two files and returns true if they are equal and
// otherwise returns false. The actual contents of the files are only compared if
// the sizes are equal. The files are not entirely loaded into memory. Instead,
// they are compared chunk by chunk. This allows for any size files to be compared,
// regardless of the memory capacity of the computer.
func (z *FileAccess) FileContentsEqual(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFileAccess_FileContentsEqual(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Copys existingFilepath to newFilepath. If failIfExists is true and newFilepath already exists, then an error is
// returned.
func (z *FileAccess) FileCopy(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkFileAccess_FileCopy(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Deletes the file specified by filePath.
func (z *FileAccess) FileDelete(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_FileDelete(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if filePath exists, otherwise returns false.
func (z *FileAccess) FileExists(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_FileExists(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns 1 if the file exists, 0 if the file does not exist, and -1 if unable to
// check because of directory permissions or some other error that prevents the
// ability to obtain the information.
func (z *FileAccess) FileExists3(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_FileExists3(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// This method should only be called on Windows operating systems. It's arguments
// are similar to the Windows Platform SDK function named CreateFile. For Linux,
// MAC OS X, and other operating system, use the OpenForRead, OpenForWrite,
// OpenForReadWrite, and OpenForAppend methods.
// 
// Opens a file for reading or writing. The arguments mirror the Windows CreateFile
// function:
// Access Modes:
// GENERIC_READ	(0x80000000)
// GENERIC_WRITE (0x40000000)
// 
// Share Modes:
// FILE_SHARE_READ(0x00000001)
// FILE_SHARE_WRITE(0x00000002)
// 
// Create Dispositions
// CREATE_NEW          1
// CREATE_ALWAYS       2
// OPEN_EXISTING       3
// OPEN_ALWAYS         4
// TRUNCATE_EXISTING   5
// 
// // Attributes:
// FILE_ATTRIBUTE_READONLY         0x00000001
// FILE_ATTRIBUTE_HIDDEN           0x00000002
// FILE_ATTRIBUTE_SYSTEM           0x00000004
// FILE_ATTRIBUTE_DIRECTORY        0x00000010
// FILE_ATTRIBUTE_ARCHIVE          0x00000020
// FILE_ATTRIBUTE_NORMAL           0x00000080
// FILE_ATTRIBUTE_TEMPORARY	   0x00000100
//
func (z *FileAccess) FileOpen(arg1 string, arg2 uint, arg3 uint, arg4 uint, arg5 uint) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_FileOpen(z.ckObj, cstr1, C.ulong(arg2), C.ulong(arg3), C.ulong(arg4), C.ulong(arg5))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Reads bytes from the currently open file. maxNumBytes specifies the maximum number of
// bytes to read. Returns an empty byte array on error.
func (z *FileAccess) FileRead(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkFileAccess_FileRead(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Reads bytes from the currently open file. maxNumBytes specifies the maximum number of
// bytes to read. Appends the bytes to the binData.
func (z *FileAccess) FileReadBd(arg1 int, arg2 *BinData) bool {

    v := C.CkFileAccess_FileReadBd(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Renames a file from existingFilepath to newFilepath.
func (z *FileAccess) FileRename(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFileAccess_FileRename(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the file pointer for the currently open file. The offset is an offset in
// bytes from the origin. The origin can be one of the following:
// 0 = Offset is from beginning of file.
// 1 = Offset is from current position of file pointer.
// 2 = Offset is from the end-of-file (offset may be negative).
func (z *FileAccess) FileSeek(arg1 int, arg2 int) bool {

    v := C.CkFileAccess_FileSeek(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Returns the size, in bytes, of a file. Returns -1 for failure.
func (z *FileAccess) FileSize(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_FileSize(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the size of the file in decimal string format.
// return a string or nil if failed.
func (z *FileAccess) FileSizeStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFileAccess_fileSizeStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Examines the file at path and returns one of the following values:
// 
// -1 = Unable to check because of directory permissions or some error preventing
// the ability to obtain the information.
// 0 = File does not exist.
// 1 = Regular file.
// 2 = Directory.
// 3 = Symbolic link.
// 4 = Windows Shortcut.
// 99 = Something else.
// 
// Additional file types may be added in the future as needed.
//
func (z *FileAccess) FileType(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_FileType(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Writes bytes to the currently open file.
func (z *FileAccess) FileWrite(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkFileAccess_FileWrite(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Writes the contents of binData to the currently open file. To specify the entire
// contents of binData, set both offset and numBytes equal to 0. To write all remaining data
// starting at offset, then set numBytes equal to 0.
func (z *FileAccess) FileWriteBd(arg1 *BinData, arg2 int, arg3 int) bool {

    v := C.CkFileAccess_FileWriteBd(z.ckObj, arg1.ckObj, C.int(arg2), C.int(arg3))


    return v != 0
}


// This is purely a utility/convenience method -- initially created to help with
// block file uploads to Azure Blob storage. It generates a block ID string that is
// the decimal representation of the index in length chars, and then encoded according
// to encoding (which can be an encoding such as "base64", "hex", "ascii", etc.) For
// example, if index = 8, length = 12, and encoding = "base64", then the string "00000012"
// is returned base64 encoded.
// return a string or nil if failed.
func (z *FileAccess) GenBlockId(arg1 int, arg2 int, arg3 string) *string {
    cstr3 := C.CString(arg3)

    retStr := C.CkFileAccess_genBlockId(z.ckObj, C.int(arg1), C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the directory information for the specified path string.
// GetDirectoryName('C:\MyDir\MySubDir\myfile.ext') returns 'C:\MyDir\MySubDir\'
// GetDirectoryName('C:\MyDir\MySubDir') returns 'C:\MyDir\'
// GetDirectoryName('C:\MyDir\') returns 'C:\MyDir\'
// GetDirectoryName('C:\MyDir') returns 'C:\'
// GetDirectoryName('C:\') returns 'C:\'
// return a string or nil if failed.
func (z *FileAccess) GetDirectoryName(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFileAccess_getDirectoryName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the extension of the specified path string.
// GetExtension('C:\mydir.old\myfile.ext') returns '.ext'
// GetExtension('C:\mydir.old\') returns ''
// return a string or nil if failed.
func (z *FileAccess) GetExtension(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFileAccess_getExtension(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the file name and extension of the specified path string.
// GetFileName('C:\mydir\myfile.ext') returns 'myfile.ext'
// GetFileName('C:\mydir\') returns ''
// return a string or nil if failed.
func (z *FileAccess) GetFileName(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFileAccess_getFileName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the file name of the specified path string without the extension.
// GetFileNameWithoutExtension('C:\mydir\myfile.ext') returns 'myfile'
// GetFileNameWithoutExtension('C:\mydir\') returns ''
// return a string or nil if failed.
func (z *FileAccess) GetFileNameWithoutExtension(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFileAccess_getFileNameWithoutExtension(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets one of the following date/times for a file:
// 0: Last-modified
// 1: Last-access
// 2: Creation
// The "path" argument indicates which time to return. The values can be 0, 1, or
// 2.
// 
// Note: Linux filesystems do not keep a file's creation date/time. In such a case,
// this method will return the last-modified time.
//
func (z *FileAccess) GetFileTime(arg1 string, arg2 int) *CkDateTime {
    cstr1 := C.CString(arg1)

    retObj := C.CkFileAccess_GetFileTime(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Gets the last-modified date/time for a file.
func (z *FileAccess) GetLastModified(arg1 string) *CkDateTime {
    cstr1 := C.CString(arg1)

    retObj := C.CkFileAccess_GetLastModified(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Returns the number of blocks in the currently open file. The number of bytes per
// block is specified by blockSize. The number of blocks is the file size divided by the
// blockSize, plus 1 if the file size is not evenly divisible by blockSize. For example, if
// the currently open file is 60500 bytes, and if the blockSize is 1000 bytes, then this
// method returns a count of 61 blocks.
// 
// Returns -1 if no file is open. Return 0 if the file is completely empty (0
// bytes).
//
func (z *FileAccess) GetNumBlocks(arg1 int) int {

    v := C.CkFileAccess_GetNumBlocks(z.ckObj, C.int(arg1))


    return int(v)
}


// Creates a temporary filepath of the form dirPath\prefix_xxxx.TMP Where "xxxx" are
// random alpha-numeric chars. The returned filepath is guaranteed to not already
// exist.
// return a string or nil if failed.
func (z *FileAccess) GetTempFilename(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkFileAccess_getTempFilename(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Opens a file for appending. If filePath did not already exists, it is created. When
// an existing file is opened with this method, the contents will not be
// overwritten and the file pointer is positioned at the end of the file.
// 
// If the open/create failed, then error information will be available in the
// FileOpenError and FileOpenErrorMsg properties.
//
func (z *FileAccess) OpenForAppend(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_OpenForAppend(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Opens a file for reading. The file may contain any type of data (binary or text)
// and it must already exist. If the open failed, then error information will be
// available in the FileOpenError and FileOpenErrorMsg properties.
func (z *FileAccess) OpenForRead(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_OpenForRead(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Opens a file for reading/writing. If filePath did not already exists, it is created.
// When an existing file is opened with this method, the contents will not be
// overwritten, but the file pointer is positioned at the beginning of the file.
// 
// If the open/create failed, then error information will be available in the
// FileOpenError and FileOpenErrorMsg properties.
//
func (z *FileAccess) OpenForReadWrite(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_OpenForReadWrite(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Opens a file for writing. If filePath did not already exists, it is created. When an
// existing file is opened with this method, the contents will be overwritten. (For
// example, calling OpenForWrite on an existing file and then immediately closing
// the file will result in an empty file.) If the open/create failed, then error
// information will be available in the FileOpenError and FileOpenErrorMsg
// properties.
func (z *FileAccess) OpenForWrite(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_OpenForWrite(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Reads the entire contents of a binary file and returns it as an encoded string
// (using an encoding such as Base64, Hex, etc.) The encoding may be one of the
// following strings: base64, hex, qp, or url.
// return a string or nil if failed.
func (z *FileAccess) ReadBinaryToEncoded(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkFileAccess_readBinaryToEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Reads the Nth block of a file, where the size of each block is specified by
// blockSize. The first block is at blockIndex 0. If the block to be read is the last in the
// file and there is not enough data to fill an entire block, then the partial
// block is returned.
func (z *FileAccess) ReadBlock(arg1 int, arg2 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkFileAccess_ReadBlock(z.ckObj, C.int(arg1), C.int(arg2), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Reads the Nth block of a file, where the size of each block is specified by
// blockSize. The first block is at blockIndex 0. If the block to be read is the last in the
// file and there is not enough data to fill an entire block, then the partial
// block is returned. The file data is appended to the contents of bd.
func (z *FileAccess) ReadBlockBd(arg1 int, arg2 int, arg3 *BinData) bool {

    v := C.CkFileAccess_ReadBlockBd(z.ckObj, C.int(arg1), C.int(arg2), arg3.ckObj)


    return v != 0
}


// Reads the entire contents of a binary file and returns the data.
func (z *FileAccess) ReadEntireFile(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkFileAccess_ReadEntireFile(z.ckObj, cstr1, ckOutBytes)

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


// Reads the entire contents of a text file, interprets the bytes according to the
// character encoding specified by charset, and returns the text file as a string.
// return a string or nil if failed.
func (z *FileAccess) ReadEntireTextFile(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkFileAccess_readEntireTextFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Scans the currently open file (opened by calling OpenForRead) for the next chunk
// of text delimited by beginMarker and endMarker. The matched text, including the beginMarker and
// endMarker are appended to sb. The bytes of the text file are interpreted according
// to charset. If startAtBeginning equals true, then scanning begins at the start of the file.
// Otherwise scanning begins starting at the byte following the last matched
// fragment.
// 
// The return value of this function is:
// 0: No match was found.
// 1: Found the next matching fragment and appended to sb.
// -1: Error reading the file.
// 
// To support a common need for use with XML files, the beginMarker is "XML tag aware". If
// the beginMarker is a string such as " ", then it will also match "
//
func (z *FileAccess) ReadNextFragment(arg1 bool, arg2 string, arg3 string, arg4 string, arg5 *StringBuilder) int {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkFileAccess_ReadNextFragment(z.ckObj, b1, cstr2, cstr3, cstr4, arg5.ckObj)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return int(v)
}


// Reassembles a file previously split by the SplitFile method.
func (z *FileAccess) ReassembleFile(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkFileAccess_ReassembleFile(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Replaces all occurrences of existingString with replacementString in a file. The character encoding,
// such as utf-8, ansi, etc. is specified by charset.
func (z *FileAccess) ReplaceStrings(arg1 string, arg2 string, arg3 string, arg4 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkFileAccess_ReplaceStrings(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return int(v)
}


// Sets the current working directory for the calling process to dirPath.
func (z *FileAccess) SetCurrentDir(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_SetCurrentDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the create date/time, the last-access date/time, and the last-modified
// date/time for a file. For non-Windows filesystems where create times are not
// implemented, the createTime is ignored.
func (z *FileAccess) SetFileTimes(arg1 string, arg2 *CkDateTime, arg3 *CkDateTime, arg4 *CkDateTime) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_SetFileTimes(z.ckObj, cstr1, arg2.ckObj, arg3.ckObj, arg4.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the last-modified date/time for a file.
func (z *FileAccess) SetLastModified(arg1 string, arg2 *CkDateTime) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_SetLastModified(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Splits a file into chunks. Please refer to the example below:
func (z *FileAccess) SplitFile(arg1 string, arg2 string, arg3 string, arg4 int, arg5 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr5 := C.CString(arg5)

    v := C.CkFileAccess_SplitFile(z.ckObj, cstr1, cstr2, cstr3, C.int(arg4), cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Creates a symbolic link.
// 
// Note: On Windows systems, this is not the same as creating a shortcut. A Windows
// symbolic link and a Windows shortcut are two different things. Shortcut files
// are common on Windows, but not symbolic links. Creating a symbolic link requires
// a special privilege, unless running as administrator. To be able to create
// symbolic links, your user account or group needs to be listed in secpol.msc →
// Security Settings → Local Policies → User Rights Assignment → Create symbolic
// links. However the special setting is not needed when running within the
// development environment, such as from Visual Studio.
//
func (z *FileAccess) SymlinkCreate(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFileAccess_SymlinkCreate(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the full pathname of the file at the end of the linkPath. Also handles
// Windows shortcut files by returning the absolute path of the target.
// return a string or nil if failed.
func (z *FileAccess) SymlinkTarget(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFileAccess_symlinkTarget(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Deletes an entire directory tree (all files and sub-directories).
func (z *FileAccess) TreeDelete(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFileAccess_TreeDelete(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Truncates the currently open file at the current file position.
func (z *FileAccess) Truncate() bool {

    v := C.CkFileAccess_Truncate(z.ckObj)


    return v != 0
}


// Opens/creates filePath, writes fileData, and closes the file.
func (z *FileAccess) WriteEntireFile(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkFileAccess_WriteEntireFile(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Opens filePath, writes textData using the character encoding specified by charset, and
// closes the file. If includedPreamble is true and the charset is Unicode or utf-8, then the
// BOM is included at the beginning of the file.
func (z *FileAccess) WriteEntireTextFile(arg1 string, arg2 string, arg3 string, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkFileAccess_WriteEntireTextFile(z.ckObj, cstr1, cstr2, cstr3, b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


