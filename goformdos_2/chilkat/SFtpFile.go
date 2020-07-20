// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkSFtpFile.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"


func NewSFtpFile() *SFtpFile {
	return &SFtpFile{C.CkSFtpFile_Create()}
}

func (z *SFtpFile) DisposeSFtpFile() {
    if z != nil {
	C.CkSFtpFile_Dispose(z.ckObj)
	}
}




// property: CreateTimeStr
// The same as the CreateTime property, but returns the date/time as an RFC822
// formatted string.
func (z *SFtpFile) CreateTimeStr() string {
    return C.GoString(C.CkSFtpFile_createTimeStr(z.ckObj))
}

// property: Filename
// The filename (or directory name, symbolic link name, etc.)
func (z *SFtpFile) Filename() string {
    return C.GoString(C.CkSFtpFile_filename(z.ckObj))
}

// property: FileType
// One of the following values:
//   regular
//   directory
//   symLink
//   special
//   unknown
//   socket
//   charDevice
//   blockDevice
//   fifo
func (z *SFtpFile) FileType() string {
    return C.GoString(C.CkSFtpFile_fileType(z.ckObj))
}

// property: Gid
// The integer Group ID of the file.
func (z *SFtpFile) Gid() int {
    return int(C.CkSFtpFile_getGid(z.ckObj))
}

// property: Group
// The group ownership of the file. This property is only supported by servers
// running SFTP v4 or later.
func (z *SFtpFile) Group() string {
    return C.GoString(C.CkSFtpFile_group(z.ckObj))
}

// property: IsAppendOnly
// If true, this file may only be appended. This property is only supported by
// servers running SFTP v6 or later.
func (z *SFtpFile) IsAppendOnly() bool {
    v := int(C.CkSFtpFile_getIsAppendOnly(z.ckObj))
    return v != 0
}

// property: IsArchive
// If true, the file should be included in backup / archive operations. This
// property is only supported by servers running SFTP v6 or later.
func (z *SFtpFile) IsArchive() bool {
    v := int(C.CkSFtpFile_getIsArchive(z.ckObj))
    return v != 0
}

// property: IsCaseInsensitive
// This attribute applies only to directories. This attribute means that files and
// directory names in this directory should be compared without regard to case.
// This property is only supported by servers running SFTP v6 or later.
func (z *SFtpFile) IsCaseInsensitive() bool {
    v := int(C.CkSFtpFile_getIsCaseInsensitive(z.ckObj))
    return v != 0
}

// property: IsCompressed
// The file is stored on disk using file-system level transparent compression. This
// flag does not affect the file data on the wire. This property is only supported
// by servers running SFTP v6 or later.
func (z *SFtpFile) IsCompressed() bool {
    v := int(C.CkSFtpFile_getIsCompressed(z.ckObj))
    return v != 0
}

// property: IsDirectory
// If true, this is a directory.
func (z *SFtpFile) IsDirectory() bool {
    v := int(C.CkSFtpFile_getIsDirectory(z.ckObj))
    return v != 0
}

// property: IsEncrypted
// The file is stored on disk using file-system level transparent encryption. This
// flag does not affect the file data on the wire (for either READ or WRITE
// requests.) This property is only supported by servers running SFTP v6 or later.
func (z *SFtpFile) IsEncrypted() bool {
    v := int(C.CkSFtpFile_getIsEncrypted(z.ckObj))
    return v != 0
}

// property: IsHidden
// If true, the file SHOULD NOT be shown to user unless specifically requested.
func (z *SFtpFile) IsHidden() bool {
    v := int(C.CkSFtpFile_getIsHidden(z.ckObj))
    return v != 0
}

// property: IsImmutable
// The file cannot be deleted or renamed, no hard link can be created to this file,
// and no data can be written to the file.
// 
// This bit implies a stronger level of protection than ReadOnly, the file
// permission mask or ACLs. Typically even the superuser cannot write to immutable
// files, and only the superuser can set or remove the bit.
// 
// This property is only supported by servers running SFTP v6 or later.
//
func (z *SFtpFile) IsImmutable() bool {
    v := int(C.CkSFtpFile_getIsImmutable(z.ckObj))
    return v != 0
}

// property: IsReadOnly
// If true, the file is read-only. This property is only supported by servers
// running SFTP v6 or later.
func (z *SFtpFile) IsReadOnly() bool {
    v := int(C.CkSFtpFile_getIsReadOnly(z.ckObj))
    return v != 0
}

// property: IsRegular
// true if this is a normal file (not a directory or any of the other non-file
// types).
func (z *SFtpFile) IsRegular() bool {
    v := int(C.CkSFtpFile_getIsRegular(z.ckObj))
    return v != 0
}

// property: IsSparse
// The file is a sparse file; this means that file blocks that have not been
// explicitly written are not stored on disk. For example, if a client writes a
// buffer at 10 M from the beginning of the file, the blocks between the previous
// EOF marker and the 10 M offset would not consume physical disk space.
// 
// Some servers may store all files as sparse files, in which case this bit will be
// unconditionally set. Other servers may not have a mechanism for determining if
// the file is sparse, and so the file MAY be stored sparse even if this flag is
// not set.
// 
// This property is only supported by servers running SFTP v6 or later.
//
func (z *SFtpFile) IsSparse() bool {
    v := int(C.CkSFtpFile_getIsSparse(z.ckObj))
    return v != 0
}

// property: IsSymLink
// true if this is a symbolic link.
func (z *SFtpFile) IsSymLink() bool {
    v := int(C.CkSFtpFile_getIsSymLink(z.ckObj))
    return v != 0
}

// property: IsSync
// When the file is modified, the changes are written synchronously to the disk.
// This property is only supported by servers running SFTP v6 or later.
func (z *SFtpFile) IsSync() bool {
    v := int(C.CkSFtpFile_getIsSync(z.ckObj))
    return v != 0
}

// property: IsSystem
// true if the file is part of the operating system. This property is only
// supported by servers running SFTP v6 or later.
func (z *SFtpFile) IsSystem() bool {
    v := int(C.CkSFtpFile_getIsSystem(z.ckObj))
    return v != 0
}

// property: LastAccessTimeStr
// The same as the LastAccessTime property, but returns the date/time as an RFC822
// formatted string.
func (z *SFtpFile) LastAccessTimeStr() string {
    return C.GoString(C.CkSFtpFile_lastAccessTimeStr(z.ckObj))
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
func (z *SFtpFile) LastMethodSuccess() bool {
    v := int(C.CkSFtpFile_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *SFtpFile) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtpFile_putLastMethodSuccess(z.ckObj,v)
}

// property: LastModifiedTimeStr
// The same as the LastModifiedTime property, but returns the date/time as an
// RFC822 formatted string.
func (z *SFtpFile) LastModifiedTimeStr() string {
    return C.GoString(C.CkSFtpFile_lastModifiedTimeStr(z.ckObj))
}

// property: Owner
// The owner of the file. This property is only supported by servers running SFTP
// v4 or later.
func (z *SFtpFile) Owner() string {
    return C.GoString(C.CkSFtpFile_owner(z.ckObj))
}

// property: Permissions
// The 'permissions' field contains a bit mask specifying file permissions. These
// permissions correspond to the st_mode field of the stat structure defined by
// POSIX [IEEE.1003-1.1996].
// 
// This protocol uses the following values for the symbols declared in the POSIX
// standard.
//        S_IRUSR  0000400 (octal)
//        S_IWUSR  0000200
//        S_IXUSR  0000100
//        S_IRGRP  0000040
//        S_IWGRP  0000020
//        S_IXGRP  0000010
//        S_IROTH  0000004
//        S_IWOTH  0000002
//        S_IXOTH  0000001
//        S_ISUID  0004000
//        S_ISGID  0002000
//        S_ISVTX  0001000
//
func (z *SFtpFile) Permissions() int {
    return int(C.CkSFtpFile_getPermissions(z.ckObj))
}

// property: Size32
// Size of the file in bytes. If the size is too large for 32-bits, a -1 is
// returned.
func (z *SFtpFile) Size32() int {
    return int(C.CkSFtpFile_getSize32(z.ckObj))
}

// property: Size64
// Size of the file in bytes. If the file size is a number too large for 64 bits,
// you have an AMAZINGLY large disk drive.
func (z *SFtpFile) Size64() int64 {
    return int64(C.CkSFtpFile_getSize64(z.ckObj))
}

// property: SizeStr
// Same as Size64, but the number is returned as a string in decimal format.
func (z *SFtpFile) SizeStr() string {
    return C.GoString(C.CkSFtpFile_sizeStr(z.ckObj))
}

// property: Uid
// The integer User ID of the file.
func (z *SFtpFile) Uid() int {
    return int(C.CkSFtpFile_getUid(z.ckObj))
}
// Returns the file creation date and time (GMT / UTC). This method is only
// supported by servers running SFTP v4 or later.
func (z *SFtpFile) GetCreateDt() *CkDateTime {

    retObj := C.CkSFtpFile_GetCreateDt(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Returns the last-access date and time (GMT / UTC).
func (z *SFtpFile) GetLastAccessDt() *CkDateTime {

    retObj := C.CkSFtpFile_GetLastAccessDt(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Returns the last-modified date and time (GMT / UTC).
func (z *SFtpFile) GetLastModifiedDt() *CkDateTime {

    retObj := C.CkSFtpFile_GetLastModifiedDt(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


