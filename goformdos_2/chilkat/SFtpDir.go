// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkSFtpDir.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewSFtpDir() *SFtpDir {
	return &SFtpDir{C.CkSFtpDir_Create()}
}

func (z *SFtpDir) DisposeSFtpDir() {
    if z != nil {
	C.CkSFtpDir_Dispose(z.ckObj)
	}
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
func (z *SFtpDir) LastMethodSuccess() bool {
    v := int(C.CkSFtpDir_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *SFtpDir) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtpDir_putLastMethodSuccess(z.ckObj,v)
}

// property: NumFilesAndDirs
// The number of entries in this directory listing.
func (z *SFtpDir) NumFilesAndDirs() int {
    return int(C.CkSFtpDir_getNumFilesAndDirs(z.ckObj))
}

// property: OriginalPath
// The original path used to fetch this directory listing. This is the string that
// was originally passed to the OpenDir method when the directory was read.
func (z *SFtpDir) OriginalPath() string {
    return C.GoString(C.CkSFtpDir_originalPath(z.ckObj))
}
// Returns the Nth filename in the directory (indexing begins at 0).
// return a string or nil if failed.
func (z *SFtpDir) GetFilename(arg1 int) *string {

    retStr := C.CkSFtpDir_getFilename(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth entry in the directory. Indexing begins at 0.
func (z *SFtpDir) GetFileObject(arg1 int) *SFtpFile {

    retObj := C.CkSFtpDir_GetFileObject(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &SFtpFile{retObj}
}


// Loads the SFTP directory object from a completed asynchronous task.
func (z *SFtpDir) LoadTaskResult(arg1 *Task) bool {

    v := C.CkSFtpDir_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sorts the files and sub-directories in ascending or descending order based on
// the field. Possible values for field are "filename", "filenameNoCase",
// "lastModifiedTime", "lastAccessTime", "lastCreateTime", or "size". (For
// case-insensitive filename sorting, use "filenameNoCase".)
func (z *SFtpDir) Sort(arg1 string, arg2 bool)  {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    C.CkSFtpDir_Sort(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))


}


