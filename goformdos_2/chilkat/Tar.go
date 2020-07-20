// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkTar.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewTar() *Tar {
	return &Tar{C.CkTar_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Tar) DisposeTar() {
    if z != nil {
	C.CkTar_Dispose(z.ckObj)
	}
}




func (z *Tar) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkTar_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Tar) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkTar_setExternalProgress(z.ckObj,1)
}

func (z *Tar) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkTar_setExternalProgress(z.ckObj,1)
}

func (z *Tar) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkTar_setExternalProgress(z.ckObj,1)
}




// property: CaptureXmlListing
// If true, then untar methods, such as Untar, UntarGz, UntarBz2, and UntarZ,
// will also capture an XML listing of the contents in the XmlListing property. The
// format of the XML contained in XmlListing is identical to what is returned by
// the ListXml method. The default value is false.
func (z *Tar) CaptureXmlListing() bool {
    v := int(C.CkTar_getCaptureXmlListing(z.ckObj))
    return v != 0
}
// property setter: CaptureXmlListing
func (z *Tar) SetCaptureXmlListing(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTar_putCaptureXmlListing(z.ckObj,v)
}

// property: Charset
// Character encoding to be used when interpreting filenames within .tar archives
// for untar operations. The default is "utf-8", and this is typically not changed.
// (The WriteTar methods always uses utf-8 to store filenames within the TAR
// archive.)
func (z *Tar) Charset() string {
    return C.GoString(C.CkTar_charset(z.ckObj))
}
// property setter: Charset
func (z *Tar) SetCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putCharset(z.ckObj,cStr)
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
func (z *Tar) DebugLogFilePath() string {
    return C.GoString(C.CkTar_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Tar) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DirMode
// The directory permissions to used in WriteTar* methods. The default is octal
// 0755. This is the value to be stored in the "mode" field of each TAR header for
// a directory entries.
func (z *Tar) DirMode() int {
    return int(C.CkTar_getDirMode(z.ckObj))
}
// property setter: DirMode
func (z *Tar) SetDirMode(value int) {
    C.CkTar_putDirMode(z.ckObj,C.int(value))
}

// property: DirPrefix
// A prefix to be added to each file's path within the TAR archive as it is being
// created. For example, if this property is set to the string "subdir1", then
// "subdir1/" will be prepended to each file's path within the TAR.
// 
// Note: This property does not apply to files added using the AddFile2 method,
// which directly specifies the path-in-tar.
//
func (z *Tar) DirPrefix() string {
    return C.GoString(C.CkTar_dirPrefix(z.ckObj))
}
// property setter: DirPrefix
func (z *Tar) SetDirPrefix(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putDirPrefix(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: FileMode
// The file permissions to used in WriteTar* methods. The default is octal 0644.
// This is the value to be stored in the "mode" field of each TAR header for a file
// entries.
func (z *Tar) FileMode() int {
    return int(C.CkTar_getFileMode(z.ckObj))
}
// property setter: FileMode
func (z *Tar) SetFileMode(value int) {
    C.CkTar_putFileMode(z.ckObj,C.int(value))
}

// property: GroupId
// The default numerical GID to be stored in each TAR header when writing TAR
// archives. The default value is 1000.
func (z *Tar) GroupId() int {
    return int(C.CkTar_getGroupId(z.ckObj))
}
// property setter: GroupId
func (z *Tar) SetGroupId(value int) {
    C.CkTar_putGroupId(z.ckObj,C.int(value))
}

// property: GroupName
// The default group name to be stored in each TAR header when writing TAR
// archives. The default value is the logged-on username of the application's
// process.
func (z *Tar) GroupName() string {
    return C.GoString(C.CkTar_groupName(z.ckObj))
}
// property setter: GroupName
func (z *Tar) SetGroupName(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putGroupName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// This is the number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any TAR operation prior to
// completion. If HeartbeatMs is 0, no AbortCheck event callbacks will occur.
func (z *Tar) HeartbeatMs() int {
    return int(C.CkTar_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Tar) SetHeartbeatMs(value int) {
    C.CkTar_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Tar) LastErrorHtml() string {
    return C.GoString(C.CkTar_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Tar) LastErrorText() string {
    return C.GoString(C.CkTar_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Tar) LastErrorXml() string {
    return C.GoString(C.CkTar_lastErrorXml(z.ckObj))
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
func (z *Tar) LastMethodSuccess() bool {
    v := int(C.CkTar_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Tar) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTar_putLastMethodSuccess(z.ckObj,v)
}

// property: MatchCaseSensitive
// Determines whether pattern matching for the MustMatch and MustNotMatch
// properties is case-sensitive or not. The default value is false.
func (z *Tar) MatchCaseSensitive() bool {
    v := int(C.CkTar_getMatchCaseSensitive(z.ckObj))
    return v != 0
}
// property setter: MatchCaseSensitive
func (z *Tar) SetMatchCaseSensitive(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTar_putMatchCaseSensitive(z.ckObj,v)
}

// property: MustMatch
// If set, then file paths must match this pattern when creating TAR archives, or
// when extracting from TAR archives. If a file does not match, it will not be
// included when creating a TAR, or it will not be extracted when extracting from a
// TAR. This property also applies to methods that create or extract from
// compressed TAR archives.
// 
// The must-match pattern may include 0 or more asterisk characters, each of which
// represents 0 or more of any character. For example, the pattern "*.txt" causes
// only .txt files to be included or extracted. The default value is an empty
// string, indicating that all files are implicitly matched.
//
func (z *Tar) MustMatch() string {
    return C.GoString(C.CkTar_mustMatch(z.ckObj))
}
// property setter: MustMatch
func (z *Tar) SetMustMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putMustMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: MustNotMatch
// If set, then file paths must NOT match this pattern when creating TAR archives,
// or when extracting from TAR archives. If a file path matches, it will not be
// included when creating a TAR, or it will not be extracted when extracting from a
// TAR. This property also applies to methods that create or extract from
// compressed TAR archives.
// 
// The must-not-match pattern may include 0 or more asterisk characters, each of
// which represents 0 or more of any character. For example, the pattern "*.obj"
// causes all .obj files to be skipped. The default value is an empty string,
// indicating that no files are skipped.
//
func (z *Tar) MustNotMatch() string {
    return C.GoString(C.CkTar_mustNotMatch(z.ckObj))
}
// property setter: MustNotMatch
func (z *Tar) SetMustNotMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putMustNotMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: NoAbsolutePaths
// If true, then absolute paths are converted to relative paths by removing the
// leading "/" or "\" character when untarring. This protects your system from
// unknowingly untarring files into important system directories, such as
// C:\Windows\system32. The default value is true.
func (z *Tar) NoAbsolutePaths() bool {
    v := int(C.CkTar_getNoAbsolutePaths(z.ckObj))
    return v != 0
}
// property setter: NoAbsolutePaths
func (z *Tar) SetNoAbsolutePaths(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTar_putNoAbsolutePaths(z.ckObj,v)
}

// property: NumDirRoots
// The total number of directory roots set by calling AddDirRoot (i.e. the number
// of times AddDirRoot was called by the application). A TAR archive is created by
// calling AddDirRoot for one or more directory tree roots, followed by a single
// call to WriteTar (or WriteTarBz2, WriteTarGz, WriteTarZ). This allows for TAR
// archives containing multiple directory trees to be created.
func (z *Tar) NumDirRoots() int {
    return int(C.CkTar_getNumDirRoots(z.ckObj))
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
func (z *Tar) PercentDoneScale() int {
    return int(C.CkTar_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *Tar) SetPercentDoneScale(value int) {
    C.CkTar_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: ScriptFileMode
// The file permissions to used in WriteTar* methods for shell script files (.sh,
// .csh, .bash, .bsh). The default is octal 0755. This is the value to be stored in
// the "mode" field of each TAR header for a file entries.
func (z *Tar) ScriptFileMode() int {
    return int(C.CkTar_getScriptFileMode(z.ckObj))
}
// property setter: ScriptFileMode
func (z *Tar) SetScriptFileMode(value int) {
    C.CkTar_putScriptFileMode(z.ckObj,C.int(value))
}

// property: SuppressOutput
// If true, then untar methods, such as Untar, UntarGz, UntarBz2, and UntarZ, do
// not produce any output. Setting this value equal to true is useful when the
// CaptureXmlListing is also set to true, which enables an application to get the
// contents of a TAR archive without extracting. The default value is false.
func (z *Tar) SuppressOutput() bool {
    v := int(C.CkTar_getSuppressOutput(z.ckObj))
    return v != 0
}
// property setter: SuppressOutput
func (z *Tar) SetSuppressOutput(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTar_putSuppressOutput(z.ckObj,v)
}

// property: UntarCaseSensitive
// This property is deprecated. Applications should instead use the
// MatchCaseSensitive property. Until this property is officially removed, it will
// behave the same as the MatchCaseSensitive property.
func (z *Tar) UntarCaseSensitive() bool {
    v := int(C.CkTar_getUntarCaseSensitive(z.ckObj))
    return v != 0
}
// property setter: UntarCaseSensitive
func (z *Tar) SetUntarCaseSensitive(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTar_putUntarCaseSensitive(z.ckObj,v)
}

// property: UntarDebugLog
// Similar to the VerboseLogging property. If set to true, then information about
// each file/directory extracted in an untar method call is logged to LastErrorText
// (or LastErrorXml / LastErrorHtml). The default value is false.
func (z *Tar) UntarDebugLog() bool {
    v := int(C.CkTar_getUntarDebugLog(z.ckObj))
    return v != 0
}
// property setter: UntarDebugLog
func (z *Tar) SetUntarDebugLog(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTar_putUntarDebugLog(z.ckObj,v)
}

// property: UntarDiscardPaths
// If true, then discard all path information when untarring. This causes all
// files to be untarred into a single directory. The default value is false.
func (z *Tar) UntarDiscardPaths() bool {
    v := int(C.CkTar_getUntarDiscardPaths(z.ckObj))
    return v != 0
}
// property setter: UntarDiscardPaths
func (z *Tar) SetUntarDiscardPaths(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTar_putUntarDiscardPaths(z.ckObj,v)
}

// property: UntarFromDir
// The directory path where files are extracted when untarring. The default value
// is ".", meaning that the current working directory of the calling process is
// used. If UntarDiscardPaths is set, then all files are untarred into this
// directory. Otherwise, the untar operation will re-create a directory tree rooted
// in this directory.
func (z *Tar) UntarFromDir() string {
    return C.GoString(C.CkTar_untarFromDir(z.ckObj))
}
// property setter: UntarFromDir
func (z *Tar) SetUntarFromDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putUntarFromDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UntarMaxCount
// Limits the number of files extracted during an untar to this count. The default
// value is 0 to indicate no maximum. To untar a single file, one might set the
// UntarMatchPattern such that it will match only the file to be extracted, and
// also set UntarMaxCount equal to 1. This causes an untar to scan the TAR archive
// until it finds the matching file, extract it, and then return.
func (z *Tar) UntarMaxCount() int {
    return int(C.CkTar_getUntarMaxCount(z.ckObj))
}
// property setter: UntarMaxCount
func (z *Tar) SetUntarMaxCount(value int) {
    C.CkTar_putUntarMaxCount(z.ckObj,C.int(value))
}

// property: UserId
// The default numerical UID to be stored in each TAR header when writing TAR
// archives. The default value is 1000.
func (z *Tar) UserId() int {
    return int(C.CkTar_getUserId(z.ckObj))
}
// property setter: UserId
func (z *Tar) SetUserId(value int) {
    C.CkTar_putUserId(z.ckObj,C.int(value))
}

// property: UserName
// The default user name to be stored in each TAR header when writing TAR archives.
// The default value is the logged-on username of the application's process.
func (z *Tar) UserName() string {
    return C.GoString(C.CkTar_userName(z.ckObj))
}
// property setter: UserName
func (z *Tar) SetUserName(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putUserName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Tar) VerboseLogging() bool {
    v := int(C.CkTar_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Tar) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTar_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Tar) Version() string {
    return C.GoString(C.CkTar_version(z.ckObj))
}

// property: WriteFormat
// The TAR format to use when writing a TAR archive. Valid values are "gnu", "pax",
// and "ustar". The default value is "gnu".
func (z *Tar) WriteFormat() string {
    return C.GoString(C.CkTar_writeFormat(z.ckObj))
}
// property setter: WriteFormat
func (z *Tar) SetWriteFormat(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putWriteFormat(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: XmlListing
// Contains the XML listing of the contents of the TAR archive for the last untar
// method call (such as Untar, UntarGz, etc.) where the CaptureXmlListing property
// was set to true.
func (z *Tar) XmlListing() string {
    return C.GoString(C.CkTar_xmlListing(z.ckObj))
}
// property setter: XmlListing
func (z *Tar) SetXmlListing(goStr string) {
    cStr := C.CString(goStr)
    C.CkTar_putXmlListing(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}
// Adds a directory tree to be included in the next call to one of the WriteTar*
// methods. To include multiple directory trees in a .tar, call AddDirRoot multiple
// times followed by a single call to WriteTar.
func (z *Tar) AddDirRoot(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_AddDirRoot(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds a directory tree to be included in the next call to one of the WriteTar*
// methods. To include multiple directory trees in a .tar, call AddDirRoot2 (and/or
// AddDirRoot) multiple times followed by a single call to WriteTar.
// 
// The rootPrefix adds a prefix to the path in the TAR for all files added under this
// root. The rootPrefix should not end with a forward-slash char. For example: This is
// good: "abc/123", but this is not good: "abc/123/". If the DirPrefix property is
// also set, its prefix will added first.
//
func (z *Tar) AddDirRoot2(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkTar_AddDirRoot2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a local file to be included in the next call to one of the WriteTar*
// methods. To include multiple files or directory trees in a .tar, call
// AddFile/AddDirRoot multiple times followed by a single call to WriteTar (or
// WriteTarGz, or WriteTarBz2).
func (z *Tar) AddFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_AddFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds a local file to be included in the next call to one of the WriteTar*
// methods. Allows for the path within the TAR to be specified. To include multiple
// files or directory trees in a .tar, call AddFile/AddFile2/AddDirRoot multiple
// times followed by a single call to WriteTar (or WriteTarGz, or WriteTarBz2).
// 
// Note: The DirPrefix property does not apply to files added via this method
// because this method explicilty specifies the path-in-tar.
//
func (z *Tar) AddFile2(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkTar_AddFile2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Creates a .deb Debian binary package archive from a control.tar.gz and
// data.tar.gz. The controlPath is the path to the control.tar.gz file (or equivalent),
// and the dataPath is the path to the data.tar.gz file. The output file path (.deb) is
// specified in debPath.
func (z *Tar) CreateDeb(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkTar_CreateDeb(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Returns the value of the Nth directory root. For example, if an application
// calls AddDirRoot twice, then the NumDirRoots property would have a value of 2,
// and GetDirRoot(0) would return the path passed to AddDirRoot in the 1st call,
// and GetDirRoot(1) would return the directory path in the 2nd call to AddDirRoot.
// return a string or nil if failed.
func (z *Tar) GetDirRoot(arg1 int) *string {

    retStr := C.CkTar_getDirRoot(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Scans a TAR archive and returns XML detailing the files and directories found
// within the TAR.
// return a string or nil if failed.
func (z *Tar) ListXml(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkTar_listXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ListXml method
func (z *Tar) ListXmlAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTar_ListXmlAsync(z.ckObj, cstr1)

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


// Loads the caller of the task's async method.
func (z *Tar) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkTar_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Unlocks the component allowing for the full functionality to be used. If this
// method unexpectedly returns false, examine the contents of the LastErrorText
// property to determine the reason for failure.
func (z *Tar) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Extracts the files and directories from a TAR archive, reconstructing the
// directory tree(s) in the local filesystem. The files are extracted to the
// directory specified by the UntarFromDir property. Returns the number of files
// and directories extracted, or -1 for failure.
func (z *Tar) Untar(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkTar_Untar(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the Untar method
func (z *Tar) UntarAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTar_UntarAsync(z.ckObj, cstr1)

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


// Extracts the files and directories from a tar.bz2 (or tar.bzip2) archive,
// reconstructing the directory tree(s) in the local filesystem. The files are
// extracted to the directory specified by the UntarFromDir property.
func (z *Tar) UntarBz2(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_UntarBz2(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the UntarBz2 method
func (z *Tar) UntarBz2Async(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTar_UntarBz2Async(z.ckObj, cstr1)

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


// Untars the first file matching the matchPattern into bd.
func (z *Tar) UntarFirstMatchingToBd(arg1 string, arg2 string, arg3 *BinData) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkTar_UntarFirstMatchingToBd(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Memory-to-memory untar. The first file matching the matchPattern is extracted and
// returned.
func (z *Tar) UntarFirstMatchingToMemory(arg1 []byte, arg2 string) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkTar_UntarFirstMatchingToMemory(z.ckObj, ckbyd1, cstr2, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
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


// Extracts the files and directories from an in-memory TAR archive, reconstructing
// the directory tree(s) in the local filesystem. The files are extracted to the
// directory specified by the UntarFromDir property. Returns the number of files
// and directories extracted, or -1 for failure.
func (z *Tar) UntarFromMemory(arg1 []byte) int {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkTar_UntarFromMemory(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return int(v)
}


// Asynchronous version of the UntarFromMemory method
func (z *Tar) UntarFromMemoryAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkTar_UntarFromMemoryAsync(z.ckObj, ckbyd1)

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


// Extracts the files and directories from a tar.gz (or tar.gzip) archive,
// reconstructing the directory tree(s) in the local filesystem. The files are
// extracted to the directory specified by the UntarFromDir property.
func (z *Tar) UntarGz(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_UntarGz(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the UntarGz method
func (z *Tar) UntarGzAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTar_UntarGzAsync(z.ckObj, cstr1)

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


// Extracts the files and directories from a tar.Z archive, reconstructing the
// directory tree(s) in the local filesystem. The files are extracted to the
// directory specified by the UntarFromDir property.
func (z *Tar) UntarZ(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_UntarZ(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the UntarZ method
func (z *Tar) UntarZAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTar_UntarZAsync(z.ckObj, cstr1)

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


// Verifies that a TAR archive is valid. This method opens the TAR archive and
// scans the entire file by walking the TAR headers. Returns true if no errors
// were found. Otherwise returns false.
func (z *Tar) VerifyTar(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_VerifyTar(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the VerifyTar method
func (z *Tar) VerifyTarAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTar_VerifyTarAsync(z.ckObj, cstr1)

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


// Writes a TAR archive. The directory trees previously added by calling AddDirRoot
// one or more times are included in the output TAR archive.
func (z *Tar) WriteTar(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_WriteTar(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the WriteTar method
func (z *Tar) WriteTarAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTar_WriteTarAsync(z.ckObj, cstr1)

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


// Writes a .tar.bz2 compressed TAR archive. The directory trees previously added
// by calling AddDirRoot one or more times are included in the output file.
func (z *Tar) WriteTarBz2(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_WriteTarBz2(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the WriteTarBz2 method
func (z *Tar) WriteTarBz2Async(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTar_WriteTarBz2Async(z.ckObj, cstr1)

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


// Writes a .tar.gz (also known as .tgz) compressed TAR archive. The directory
// trees previously added by calling AddDirRoot one or more times are included in
// the output file.
func (z *Tar) WriteTarGz(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTar_WriteTarGz(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the WriteTarGz method
func (z *Tar) WriteTarGzAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTar_WriteTarGzAsync(z.ckObj, cstr1)

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


