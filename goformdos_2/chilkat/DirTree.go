// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkDirTree.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewDirTree() *DirTree {
	return &DirTree{C.CkDirTree_Create()}
}

func (z *DirTree) DisposeDirTree() {
    if z != nil {
	C.CkDirTree_Dispose(z.ckObj)
	}
}




// property: BaseDir
// Begin iterating from this directory.
func (z *DirTree) BaseDir() string {
    return C.GoString(C.CkDirTree_baseDir(z.ckObj))
}
// property setter: BaseDir
func (z *DirTree) SetBaseDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkDirTree_putBaseDir(z.ckObj,cStr)
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
func (z *DirTree) DebugLogFilePath() string {
    return C.GoString(C.CkDirTree_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *DirTree) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkDirTree_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DoneIterating
// Set to true when the last file or sub-directory has been iterated.
func (z *DirTree) DoneIterating() bool {
    v := int(C.CkDirTree_getDoneIterating(z.ckObj))
    return v != 0
}

// property: FileSize32
// The file size of the current file in the iteration. (0 if it is a directory.)
func (z *DirTree) FileSize32() int {
    return int(C.CkDirTree_getFileSize32(z.ckObj))
}

// property: FileSize64
// The file size as a 64-bit integer of the current file in the iteration. (0 if it
// is a directory.)
func (z *DirTree) FileSize64() int64 {
    return int64(C.CkDirTree_getFileSize64(z.ckObj))
}

// property: FullPath
// The absolute directory path of the current file or sub-directory.
func (z *DirTree) FullPath() string {
    return C.GoString(C.CkDirTree_fullPath(z.ckObj))
}

// property: FullUncPath
// The absolute UNC directory path of the current file or sub-directory.
func (z *DirTree) FullUncPath() string {
    return C.GoString(C.CkDirTree_fullUncPath(z.ckObj))
}

// property: IsDirectory
// true if the current position is a sub-directory, false if it is a file.
func (z *DirTree) IsDirectory() bool {
    v := int(C.CkDirTree_getIsDirectory(z.ckObj))
    return v != 0
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *DirTree) LastErrorHtml() string {
    return C.GoString(C.CkDirTree_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *DirTree) LastErrorText() string {
    return C.GoString(C.CkDirTree_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *DirTree) LastErrorXml() string {
    return C.GoString(C.CkDirTree_lastErrorXml(z.ckObj))
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
func (z *DirTree) LastMethodSuccess() bool {
    v := int(C.CkDirTree_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *DirTree) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDirTree_putLastMethodSuccess(z.ckObj,v)
}

// property: Recurse
// If true, the iteration will be recursive. If false the iteration is
// non-recursive. The default value is true.
func (z *DirTree) Recurse() bool {
    v := int(C.CkDirTree_getRecurse(z.ckObj))
    return v != 0
}
// property setter: Recurse
func (z *DirTree) SetRecurse(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDirTree_putRecurse(z.ckObj,v)
}

// property: RelativePath
// The relative directory path of the current file or sub-directory. (Relative to
// the BaseDir)
func (z *DirTree) RelativePath() string {
    return C.GoString(C.CkDirTree_relativePath(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *DirTree) VerboseLogging() bool {
    v := int(C.CkDirTree_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *DirTree) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDirTree_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *DirTree) Version() string {
    return C.GoString(C.CkDirTree_version(z.ckObj))
}
// Advances the current position in the directory tree traversal to the next file
// or sub-directory.
// 
// Important: If AdvancePosition returns false, it can be an error, or it could
// be that there are no more files and directories. To distinguish between the two
// cases, examine the DoneIterating property. If DoneIterating is true, then the
// false return value is not an error, but instead indicates that the end has
// been reached.
//
func (z *DirTree) AdvancePosition() bool {

    v := C.CkDirTree_AdvancePosition(z.ckObj)


    return v != 0
}


// Begins a directory tree traversal. After calling this method, the various
// property values such as Fullpath, FileSize32, etc. can be retrieved for the 1st
// file / sub-directory in the traversal.
// 
// Important: If BeginIterate returns false, it can be an error, or it could be
// that there are 0 files and directories. To distinguish between the two cases,
// examine the DoneIterating property. If DoneIterating is true, then the false
// return value is not an error, but instead indicates 0 files/directories.
//
func (z *DirTree) BeginIterate() bool {

    v := C.CkDirTree_BeginIterate(z.ckObj)


    return v != 0
}


