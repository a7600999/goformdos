// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkZip.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewZip() *Zip {
	return &Zip{C.CkZip_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Zip) DisposeZip() {
    if z != nil {
	C.CkZip_Dispose(z.ckObj)
	}
}




func (z *Zip) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkZip_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Zip) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkZip_setExternalProgress(z.ckObj,1)
}

func (z *Zip) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkZip_setExternalProgress(z.ckObj,1)
}

func (z *Zip) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkZip_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Zip) AbortCurrent() bool {
    v := int(C.CkZip_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Zip) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putAbortCurrent(z.ckObj,v)
}

// property: AppendFromDir
// When files are added to a Zip archive, they are appended from this directory.
// For example, to add all the files under c:/abc/123/myAppDir, this property could
// be set to "c:/abc/123", and "myAppDir/*" would be passed to AppendFiles. The
// path that is saved in the .zip would be "myAppDir/". (The value of the
// AppendFromDir property does not become part of the file path saved in the .zip.)
func (z *Zip) AppendFromDir() string {
    return C.GoString(C.CkZip_appendFromDir(z.ckObj))
}
// property setter: AppendFromDir
func (z *Zip) SetAppendFromDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putAppendFromDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CaseSensitive
// If true then all methods that get or search for zip entries by name will use
// case-sensitive filename matching. If false then filename matching will be case
// insensitive. Methods affected by this property include GetEntryByName,
// UnzipMatching, FirstMatchingEntry, etc.
// 
// The default value is false.
//
func (z *Zip) CaseSensitive() bool {
    v := int(C.CkZip_getCaseSensitive(z.ckObj))
    return v != 0
}
// property setter: CaseSensitive
func (z *Zip) SetCaseSensitive(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putCaseSensitive(z.ckObj,v)
}

// property: ClearArchiveAttribute
// Set this to true to clear the FILE_ATTRIBUTE_ARCHIVE file attribute of each
// file during a zipping operation.
// 
// The default value is false.
//
func (z *Zip) ClearArchiveAttribute() bool {
    v := int(C.CkZip_getClearArchiveAttribute(z.ckObj))
    return v != 0
}
// property setter: ClearArchiveAttribute
func (z *Zip) SetClearArchiveAttribute(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putClearArchiveAttribute(z.ckObj,v)
}

// property: ClearReadOnlyAttr
// If true, the read-only attribute is automatically cleared when unzipping. The
// default value of this property is false, which leaves the read-only attribute
// unchanged when unzipping.
func (z *Zip) ClearReadOnlyAttr() bool {
    v := int(C.CkZip_getClearReadOnlyAttr(z.ckObj))
    return v != 0
}
// property setter: ClearReadOnlyAttr
func (z *Zip) SetClearReadOnlyAttr(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putClearReadOnlyAttr(z.ckObj,v)
}

// property: Comment
// The global Zip file comment.
func (z *Zip) Comment() string {
    return C.GoString(C.CkZip_comment(z.ckObj))
}
// property setter: Comment
func (z *Zip) SetComment(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putComment(z.ckObj,cStr)
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
func (z *Zip) DebugLogFilePath() string {
    return C.GoString(C.CkZip_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Zip) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DecryptPassword
// When opening a password-protected or AES encrypted Zip, this is the password to
// be used for decryption. Encrypted Zips may be opened without setting a password,
// but the contents cannot be unzipped without setting this password.
// 
// Note:The SetPassword method has the effect of setting both this property as well
// as the EncryptPassword property. The SetPassword method should no longer be
// used. It has been replaced by the DecryptPassword and EncryptPassword properties
// to make it possible to open an encrypted zip and re-write it with a new
// password.
//
func (z *Zip) DecryptPassword() string {
    return C.GoString(C.CkZip_decryptPassword(z.ckObj))
}
// property setter: DecryptPassword
func (z *Zip) SetDecryptPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putDecryptPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DiscardPaths
// If true, discards all file path information when zipping. The default value is
// false.
func (z *Zip) DiscardPaths() bool {
    v := int(C.CkZip_getDiscardPaths(z.ckObj))
    return v != 0
}
// property setter: DiscardPaths
func (z *Zip) SetDiscardPaths(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putDiscardPaths(z.ckObj,v)
}

// property: Encryption
// Indicate whether the Zip is to be strong encrypted or not. Valid values are 0
// (not encrypted) or 4 (AES encrypted). When this property is set to the value 4,
// WinZip AES compatible encrypted zip archives are produced.
// 
// Note: Prior to Chilkat v9.4.1, other possible values for this property were: 1
// (blowfish), 2 (twofish), and 3 (rijndael). These settings originally provided a
// way to produce strong encrypted zips prior to when the AES encrypted Zip
// standard existed. Using these legacy values (1, 2, or 3) produced encrypted zips
// that only applications using Chilkat could read. Chilkat no longer supports
// these custom modes of encryption. If using an older version of Chilkat with one
// of these deprecated encryption modes, make sure to decrypt using the old Chilkat
// version and re-encrypt using mode 4 (WinZip compatible AES encryption) prior to
// updating to the new Chilkat version.
// 
// Important:The Encryption and PasswordProtect properties are mutually exclusive.
// PasswordProtect corresponds to the older Zip 2.0 encryption, commonly referred
// to as a "password-protected" zip. If the PasswordProtect is set to true, the
// Encryption property should be set to 0. If the Encryption property is set to a
// non-zero value, then PasswordProtect should be set to false. A zip cannot be
// both password-protected and strong-encrypted.
//
func (z *Zip) Encryption() int {
    return int(C.CkZip_getEncryption(z.ckObj))
}
// property setter: Encryption
func (z *Zip) SetEncryption(value int) {
    C.CkZip_putEncryption(z.ckObj,C.int(value))
}

// property: EncryptKeyLength
// The encryption key length if AES, Blowfish, Twofish, or WinZip-compatible AES
// encryption is used. This value must be 128, 192, or 256. The default value is
// 128.
func (z *Zip) EncryptKeyLength() int {
    return int(C.CkZip_getEncryptKeyLength(z.ckObj))
}
// property setter: EncryptKeyLength
func (z *Zip) SetEncryptKeyLength(value int) {
    C.CkZip_putEncryptKeyLength(z.ckObj,C.int(value))
}

// property: EncryptPassword
// The password used when writing a password-protected or strong-encrytped Zip.
// 
// Note:The SetPassword method has the effect of setting both this property as well
// as the DecryptPassword property. The SetPassword method should no longer be
// used. It has been replaced by the DecryptPassword and EncryptPassword properties
// to make it possible to open an encrypted zip and re-write it with a new
// password.
//
func (z *Zip) EncryptPassword() string {
    return C.GoString(C.CkZip_encryptPassword(z.ckObj))
}
// property setter: EncryptPassword
func (z *Zip) SetEncryptPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putEncryptPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: FileCount
// The number of files (excluding directories) contained within the Zip.
func (z *Zip) FileCount() int {
    return int(C.CkZip_getFileCount(z.ckObj))
}

// property: FileName
// The path (absolute or relative) of the Zip archive. This is the path of the file
// that is created or overwritten when the zip is saved.
func (z *Zip) FileName() string {
    return C.GoString(C.CkZip_fileName(z.ckObj))
}
// property setter: FileName
func (z *Zip) SetFileName(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putFileName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HasZipFormatErrors
// true if the opened zip contained file format errors (that were not severe
// enough to prevent the zip from being opened and parsed).
func (z *Zip) HasZipFormatErrors() bool {
    v := int(C.CkZip_getHasZipFormatErrors(z.ckObj))
    return v != 0
}

// property: HeartbeatMs
// The number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any method call prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
func (z *Zip) HeartbeatMs() int {
    return int(C.CkZip_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Zip) SetHeartbeatMs(value int) {
    C.CkZip_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: IgnoreAccessDenied
// If true, then files that cannot be read due to "access denied" (i.e. a file
// permission error) will be ignored and the call to WriteZip, WriteZipAndClose,
// WriteExe, etc. will return a success status. If false, then the "access
// denied" filesystem errors are not ignored and any occurrence will cause the zip
// writing to fail. The default value is true.
func (z *Zip) IgnoreAccessDenied() bool {
    v := int(C.CkZip_getIgnoreAccessDenied(z.ckObj))
    return v != 0
}
// property setter: IgnoreAccessDenied
func (z *Zip) SetIgnoreAccessDenied(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putIgnoreAccessDenied(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Zip) LastErrorHtml() string {
    return C.GoString(C.CkZip_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Zip) LastErrorText() string {
    return C.GoString(C.CkZip_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Zip) LastErrorXml() string {
    return C.GoString(C.CkZip_lastErrorXml(z.ckObj))
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
func (z *Zip) LastMethodSuccess() bool {
    v := int(C.CkZip_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Zip) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putLastMethodSuccess(z.ckObj,v)
}

// property: NumEntries
// The number of entries in the Zip, including both files and directories.
func (z *Zip) NumEntries() int {
    return int(C.CkZip_getNumEntries(z.ckObj))
}

// property: OemCodePage
// Sets the OEM code page to be used for Unicode filenames. This property defaults
// to the OEM code page of the computer.
func (z *Zip) OemCodePage() int {
    return int(C.CkZip_getOemCodePage(z.ckObj))
}
// property setter: OemCodePage
func (z *Zip) SetOemCodePage(value int) {
    C.CkZip_putOemCodePage(z.ckObj,C.int(value))
}

// property: OverwriteExisting
// Determines whether existing files are overwritten during unzipping. The default
// is true, which means that already-existing files will be overwritten. Set this
// property = false to prevent existing files from being overwritten when
// unzipping.
func (z *Zip) OverwriteExisting() bool {
    v := int(C.CkZip_getOverwriteExisting(z.ckObj))
    return v != 0
}
// property setter: OverwriteExisting
func (z *Zip) SetOverwriteExisting(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putOverwriteExisting(z.ckObj,v)
}

// property: PasswordProtect
// true if the Zip should be password-protected using older Zip 2.0 encryption,
// commonly referred to as "password-protection".
// 
// This property is set when a zip archive is opened by any of the Open* methods,
// such as OpenZip, OpenFromMemory, etc.
//
func (z *Zip) PasswordProtect() bool {
    v := int(C.CkZip_getPasswordProtect(z.ckObj))
    return v != 0
}
// property setter: PasswordProtect
func (z *Zip) SetPasswordProtect(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putPasswordProtect(z.ckObj,v)
}

// property: PathPrefix
// A prefix that is added to each filename when zipping. One might set the
// PathPrefix to "subdir/" so that files are unzipped to a specified subdirectory
// when unzipping.
func (z *Zip) PathPrefix() string {
    return C.GoString(C.CkZip_pathPrefix(z.ckObj))
}
// property setter: PathPrefix
func (z *Zip) SetPathPrefix(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putPathPrefix(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
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
func (z *Zip) PercentDoneScale() int {
    return int(C.CkZip_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *Zip) SetPercentDoneScale(value int) {
    C.CkZip_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: PwdProtCharset
// For older password-protected Zip archives (Zip 2.0 encryption), specifies the
// charset used for the binary representation of the decrypt password. The default
// value is "ansi". Other possible choices are cp850, cp437, or any of the code
// pages listed at the link below.
func (z *Zip) PwdProtCharset() string {
    return C.GoString(C.CkZip_pwdProtCharset(z.ckObj))
}
// property setter: PwdProtCharset
func (z *Zip) SetPwdProtCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putPwdProtCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TempDir
// The temporary directory to use when unzipping files. When running in ASP or
// ASP.NET, the default value of TempDir is set to the directory where the .zip is
// being written. Set this property to override the default.
func (z *Zip) TempDir() string {
    return C.GoString(C.CkZip_tempDir(z.ckObj))
}
// property setter: TempDir
func (z *Zip) SetTempDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putTempDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TextFlag
// If set to true, the component will set the "text flag" for each file having
// these filename extensions: .txt, .xml, .htm, and .html. It will also preserve
// the "text flag" for existing zips that are opened and rewritten. By default,
// this property is set to false.
// 
// It is generally not necessary to set the text flag for a zip entry.
//
func (z *Zip) TextFlag() bool {
    v := int(C.CkZip_getTextFlag(z.ckObj))
    return v != 0
}
// property setter: TextFlag
func (z *Zip) SetTextFlag(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putTextFlag(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Zip) VerboseLogging() bool {
    v := int(C.CkZip_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Zip) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Zip) Version() string {
    return C.GoString(C.CkZip_version(z.ckObj))
}

// property: Zipx
// Starting in v9.4.1, Chilkat Zip will automatically unzip ZIPX files using any of
// the following compression methods: BZIP2, PPMd, LZMA, and Deflate64 ("Deflate64"
// is a trademark of PKWare, Inc.)
// 
// This property, however, controls whether or not a ZipX is automatically produced
// where the best compression algorithm for each file is automatically chosen based
// on file type. This property is for writing zip archives. It does not apply to
// when unzipping ZIPX archives, Chilkat Zip automatically handles the various
// compression algorithms when unzipping.
//
func (z *Zip) Zipx() bool {
    v := int(C.CkZip_getZipx(z.ckObj))
    return v != 0
}
// property setter: Zipx
func (z *Zip) SetZipx(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkZip_putZipx(z.ckObj,v)
}

// property: ZipxDefaultAlg
// The default compression algorithm to be used when creating ZIPX archives. The
// default value is "deflate". Other possible values are "ppmd", "lzma", "bzip2"
// and "deflate64". When writing a ZIPX archive, if the file extension does not
// indicate an obvious choice for the appropriate compression algorithm, then the
// ZipxDefaultAlg is used.
func (z *Zip) ZipxDefaultAlg() string {
    return C.GoString(C.CkZip_zipxDefaultAlg(z.ckObj))
}
// property setter: ZipxDefaultAlg
func (z *Zip) SetZipxDefaultAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkZip_putZipxDefaultAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}
// Attempting to compress already-compressed data is usually a waste of CPU cycles
// with little or no benefit. In fact, it is possible that attempting to compress
// already-compressed data results in a slightly increased size. The Zip file
// format allows for files to be "stored" rather than compressed. This allows the
// file data to be streamed directly into a .zip without compression.
// 
// An instance of the Zip object has an internal list of "no compress" extensions.
// A filename with a "no compress" extension is "stored" rather than compressed.
// Additional "no compress" extensions may be added by calling this method (once
// per file extension). You should pass the file extension, such as ".xyz" in fileExtension.
// 
// "no compress" extensions may be removed by calling RemoveNoCompressExtension.
// 
// The default "no compress" extensions are: .zip, .gif, .jpg, .gz, .rar, .jar,
// .tgz, .bz2, .z, .rpm, .msi, .png
//
func (z *Zip) AddNoCompressExtension(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkZip_AddNoCompressExtension(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Creates a new Zip entry and initializes it with already-compressed data that is
// Base64 encoded. (The ZipEntry.CopyBase64 method can be used to retrieve the
// compressed data in Base64 format.)
// 
// Note 1: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
// 
// Note 2: It is assumed that the compressed data is unencrypted deflated data.
// (Meaning data compressed using the "deflate" compression algorithm.)
//
func (z *Zip) AppendBase64(arg1 string, arg2 string) *ZipEntry {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkZip_AppendBase64(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Appends the contents of byteData as a new entry to this zip object. The zip entry
// object containing the data is returned.
func (z *Zip) AppendBd(arg1 string, arg2 *BinData) *ZipEntry {
    cstr1 := C.CString(arg1)

    retObj := C.CkZip_AppendBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Append memory data that is already Zip-compressed to the Zip object. The
// ZipEntry object containing the compressed data is returned. Note: This method
// appends the compressed data for a single zip entry. To load an entire in-memory
// .zip, call OpenFromMemory instead.
// 
// Note 1: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
// 
// Note 2: It is assumed that the compressed data is unencrypted deflated data.
// (Meaning data compressed using the "deflate" compression algorithm.)
//
func (z *Zip) AppendCompressed(arg1 string, arg2 []byte) *ZipEntry {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    retObj := C.CkZip_AppendCompressed(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Appends in-memory data as a new entry to a Zip object. The ZipEntry object
// containing the data is returned.
// 
// Note: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
//
func (z *Zip) AppendData(arg1 string, arg2 []byte) *ZipEntry {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    retObj := C.CkZip_AppendData(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Appends in-memory data as a new entry to a Zip object. The filename is the filename
// of the entry as it will appear within the zip. The encoding is the encoding of the
// data, such as "base64", "hex", etc. The full list of encodings is listed at the
// web page linked below.
// 
// Returns the zip entry object.
// 
// Note: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
//
func (z *Zip) AppendDataEncoded(arg1 string, arg2 string, arg3 string) *ZipEntry {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkZip_AppendDataEncoded(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Appends one or more files to the Zip object. The filePattern can use the "*"
// wildcard character for 0 or more of any characterSet recurse equal to True to
// recursively add all subdirectories, or False to only add files in the current
// directory.
// 
// Note: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
//
func (z *Zip) AppendFiles(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkZip_AppendFiles(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the AppendFiles method
func (z *Zip) AppendFilesAsync(arg1 string, arg2 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkZip_AppendFilesAsync(z.ckObj, cstr1, b2)

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


// Appends one or more files to the Zip object. The filePattern can use the "*" to mean 0
// or more of any character. The recurse controls whether directories are recursively
// traversed. Set recurse equal to true to append files and subdirectories in the
// directory tree. Set recurse equal to false to add files only from the indicated
// directory.
// 
// The saveExtraPath only applies when the filePattern is an absolute path pattern, such as
// "C:/temp/abc/*.txt". If saveExtraPath is true, then the absolute path will be included
// in the zip entry filenames as relative paths. For example, "temp/abc/xyz.txt".
// 
// The archiveOnly, includeHidden, and includeSystem flags only apply when on the Windows operating system.
// If archiveOnly is true, then only files that have the archive bit set will be
// included in the zip. If includeHidden is false, then hidden files are not included. If
// includeSystem is false, then files having the System attribute are not included.
// 
// Note: This method does not write the zip archive. It simply adds references to
// the files that will be included in the .zip when the WriteZip or
// WriteZipAndClose methods are eventually called. Files and/or data may be added
// to the zip object by calling any combination of the Append* methods before
// finally writing the zip via one of the Write* methods.
//
func (z *Zip) AppendFilesEx(arg1 string, arg2 bool, arg3 bool, arg4 bool, arg5 bool, arg6 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    v := C.CkZip_AppendFilesEx(z.ckObj, cstr1, b2, b3, b4, b5, b6)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the AppendFilesEx method
func (z *Zip) AppendFilesExAsync(arg1 string, arg2 bool, arg3 bool, arg4 bool, arg5 bool, arg6 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    hTask := C.CkZip_AppendFilesExAsync(z.ckObj, cstr1, b2, b3, b4, b5, b6)

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


// Creates a new Zip entry and initializes it with already-compressed data that is
// hexidecimal encoded. (The ZipEntry.CopyHex method can be used to retrieve the
// compressed data in Hex format.)
// 
// Note 1: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
// 
// Note 2: It is assumed that the compressed data is unencrypted deflated data.
// (Meaning data compressed using the "deflate" compression algorithm.)
//
func (z *Zip) AppendHex(arg1 string, arg2 string) *ZipEntry {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkZip_AppendHex(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// This method is the same as calling AppendFiles multiple times - once for each
// file pattern in fileSpecs
// 
// Note: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
//
func (z *Zip) AppendMultiple(arg1 *StringArray, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkZip_AppendMultiple(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Asynchronous version of the AppendMultiple method
func (z *Zip) AppendMultipleAsync(arg1 *StringArray, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkZip_AppendMultipleAsync(z.ckObj, arg1.ckObj, b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Appends a new and empty entry to the Zip object and returns the ZipEntry object.
// Data can be appended to the entry by calling ZipEntry.AppendData.
// 
// Important: To append an already-existing file, call the AppendOneFileOrDir
// method. The AppendNew method inserts a new and empty file entry within the Zip
// object. The purpose of AppendNew is to either create an empty file within the
// Zip, or to create a new file entry which can then be filled with data by calling
// the entry's AppendData method.
// 
// Note: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
//
func (z *Zip) AppendNew(arg1 string) *ZipEntry {
    cstr1 := C.CString(arg1)

    retObj := C.CkZip_AppendNew(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Adds an entry to the zip so that when it unzips, a new directory (with no files)
// is created. The directory does not need to exist on the local filesystem when
// calling this method. The dirName is simply a string that is used as the directory
// path for the entry added to the zip. The zip entry object is returned.
// 
// Note: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
//
func (z *Zip) AppendNewDir(arg1 string) *ZipEntry {
    cstr1 := C.CString(arg1)

    retObj := C.CkZip_AppendNewDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Appends a single file or directory to the Zip object. The saveExtraPath applies when fileOrDirPath
// is an absolute (non-relative) path. If saveExtraPath is true, then the absolute path is
// made relative and saved in the zip. For example, if the fileOrDirPath is
// "C:/temp/xyz/test.txt" and saveExtraPath is true, then the path in the zip will be
// "./temp/xyz/test.txt". If however, fileOrDirPath contains a relative path, then saveExtraPath has
// no effect.
func (z *Zip) AppendOneFileOrDir(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkZip_AppendOneFileOrDir(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the AppendOneFileOrDir method
func (z *Zip) AppendOneFileOrDirAsync(arg1 string, arg2 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkZip_AppendOneFileOrDirAsync(z.ckObj, cstr1, b2)

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


// Adds an in-memory string to the Zip object. The textData argument is converted to
// the ANSI charset before being added to the Zip. If the Zip were written to disk
// by calling WriteZip, and later unzipped, the entry would unzip to an ANSI text
// file.
// 
// Note: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
//
func (z *Zip) AppendString(arg1 string, arg2 string) *ZipEntry {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkZip_AppendString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Same as AppendString, but allows the charset to be specified. The textData is
// converted to charset before being added to the zip. The internalZipFilepath is the path of the
// file that will be stored within the zip.
// 
// Note: This method only updates the zip object. To update (rewrite) a zip file,
// either the WriteZip or WriteZipAndClose method would need to be called.
//
func (z *Zip) AppendString2(arg1 string, arg2 string, arg3 string) *ZipEntry {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkZip_AppendString2(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Adds the contents of another existing Zip file to this Zip object.
func (z *Zip) AppendZip(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_AppendZip(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Closes an open Zip file. This is identical to calling NewZip. (NewZip closes the
// current Zip file, if open, and initializes the Zip object to be empty. Zip files
// are only created when WriteZip is called.)
func (z *Zip) CloseZip()  {

    C.CkZip_CloseZip(z.ckObj)



}


// Removes a Zip entry from the calling Zip object.
func (z *Zip) DeleteEntry(arg1 *ZipEntry) bool {

    v := C.CkZip_DeleteEntry(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds a directory name to be excluded when AppendFiles is called to add an entire
// directory tree. All directories having a name equal to an excluded directory
// will not be included when AppendFiles (or AppendFileEx) is called. Multiple
// directories can be excluded by calling ExcludeDir multiple times. The name
// comparison is case-insensitive.
func (z *Zip) ExcludeDir(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkZip_ExcludeDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Unzip all the files into the specified directory. Subdirectories are
// automatically created as needed.
func (z *Zip) Extract(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_Extract(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Extract method
func (z *Zip) ExtractAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkZip_ExtractAsync(z.ckObj, cstr1)

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


// Unzips all the files in a Zip into a single directory regardless of the path
// stored in the Zip
func (z *Zip) ExtractInto(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_ExtractInto(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unzip all files matching a wildcard pattern.
func (z *Zip) ExtractMatching(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkZip_ExtractMatching(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Extracts only the files that have more recent last-modified-times than the files
// on disk. This allows you to easily refresh only the files that have been
// updated.
func (z *Zip) ExtractNewer(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_ExtractNewer(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Return the first entry in the Zip. Call ZipEntry.NextEntry to iterate over the
// entries in a Zip until a NULL is returned.
func (z *Zip) FirstEntry() *ZipEntry {

    retObj := C.CkZip_FirstEntry(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Returns the first entry having a filename matching a pattern. The "*" characters
// matches 0 or more of any character. The full filename, including path, is used
// when matching against the pattern. A NULL is returned if nothing matches.
func (z *Zip) FirstMatchingEntry(arg1 string) *ZipEntry {
    cstr1 := C.CString(arg1)

    retObj := C.CkZip_FirstMatchingEntry(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Return the contents of the Zip file directory in an XML formatted string
// return a string or nil if failed.
func (z *Zip) GetDirectoryAsXML() *string {

    retStr := C.CkZip_getDirectoryAsXML(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Retrieves a ZipEntry by ID. Chilkat Zip.NET automatically assigns a unique ID to
// each ZipEntry in the Zip. This feature makes it easy to associate an item in a
// UI control with a ZipEntry.
func (z *Zip) GetEntryByID(arg1 int) *ZipEntry {

    retObj := C.CkZip_GetEntryByID(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Retrieves a ZipEntry by index. The first entry is at index 0. This will return
// directory entries as well as files.
func (z *Zip) GetEntryByIndex(arg1 int) *ZipEntry {

    retObj := C.CkZip_GetEntryByIndex(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Returns a ZipEntry by filename. If a full or partial path is part of the
// filename, this must be included in the filename parameter.
func (z *Zip) GetEntryByName(arg1 string) *ZipEntry {
    cstr1 := C.CString(arg1)

    retObj := C.CkZip_GetEntryByName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Returns the current collection of exclusion patterns that have been set by
// SetExclusions.
func (z *Zip) GetExclusions() *StringArray {

    retObj := C.CkZip_GetExclusions(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Inserts a new and empty entry into the Zip object. To insert at the beginning of
// the Zip, beforeIndex should be 0. The ZipEntry's FileName property is
// initialized to fileName parameter.
func (z *Zip) InsertNew(arg1 string, arg2 int) *ZipEntry {
    cstr1 := C.CString(arg1)

    retObj := C.CkZip_InsertNew(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &ZipEntry{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Returns true if the fileExtension is contained in the set of "no compress" extensions,
// otherwise returns false. (See the documentation for the AddNoCompressExtension
// method.) The fileExtension may be passed with or without the ".". For example, both
// ".jpg" and "jpg" are acceptable.
func (z *Zip) IsNoCompressExtension(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_IsNoCompressExtension(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Return True if a Zip file is password protected
func (z *Zip) IsPasswordProtected(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_IsPasswordProtected(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns True if the class is already unlocked, otherwise returns False.
func (z *Zip) IsUnlocked() bool {

    v := C.CkZip_IsUnlocked(z.ckObj)


    return v != 0
}


// Loads the caller of the task's async method.
func (z *Zip) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkZip_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Clears and initializes the contents of the Zip object. If a Zip file was open,
// it is closed and all entries are removed from the object. The FileName property
// is set to the zipFilePath argument.
func (z *Zip) NewZip(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_NewZip(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Open a Zip contained in binData.
// 
// When a zip is opened, the PasswordProtect and Encryption properties will be
// appropriately set. If the zip is password protected (i.e. uses older Zip 2.0
// encrypion), then the PasswordProtect property will be set to true. If the zip
// is strong encrypted, the Encryption property will be set to a value 1 through 4,
// where 4 indicates WinZip compatible AES encryption.
//
func (z *Zip) OpenBd(arg1 *BinData) bool {

    v := C.CkZip_OpenBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Same as OpenFromMemory.
// 
// When a zip is opened, the PasswordProtect and Encryption properties will be
// appropriately set. If the zip is password protected (i.e. uses older Zip 2.0
// encrypion), then the PasswordProtect property will be set to true. If the zip
// is strong encrypted, the Encryption property will be set to a value 1 through 4,
// where 4 indicates WinZip compatible AES encryption.
//
func (z *Zip) OpenFromByteData(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkZip_OpenFromByteData(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Open a Zip that is completely in-memory. This allows for Zip files to be opened
// from non-filesystem sources, such as a database.
// 
// When a zip is opened, the PasswordProtect and Encryption properties will be
// appropriately set. If the zip is password protected (i.e. uses older Zip 2.0
// encrypion), then the PasswordProtect property will be set to true. If the zip
// is strong encrypted, the Encryption property will be set to a value 1 through 4,
// where 4 indicates WinZip compatible AES encryption.
//
func (z *Zip) OpenFromMemory(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkZip_OpenFromMemory(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Opens a Zip archive. Encrypted and password-protected zips may be opened without
// providing the password, but their contents may not be unzipped unless the
// correct password is provided via the DecryptPassword proprety, or the
// SetPassword method.
// 
// When a zip is opened, the PasswordProtect and Encryption properties will be
// appropriately set. If the zip is password protected (i.e. uses older Zip 2.0
// encrypion), then the PasswordProtect property will be set to true. If the zip
// is strong encrypted, the Encryption property will be set to a value 1 through 4,
// where 4 indicates WinZip compatible AES encryption.
//
func (z *Zip) OpenZip(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_OpenZip(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the OpenZip method
func (z *Zip) OpenZipAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkZip_OpenZipAsync(z.ckObj, cstr1)

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


// Efficiently appends additional files to an existing zip archive. QuickAppend
// leaves all entries in the existing .zip untouched. It operates by appending new
// files and updating the internal "central directory" of the zip archive.
func (z *Zip) QuickAppend(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_QuickAppend(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the QuickAppend method
func (z *Zip) QuickAppendAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkZip_QuickAppendAsync(z.ckObj, cstr1)

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


// Removes a file extension from the zip object's internal list of "no compress"
// extensions. (For more information, see AddNoCompressExtension.)
func (z *Zip) RemoveNoCompressExtension(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkZip_RemoveNoCompressExtension(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Sets the compression level for all file and data entries. The compression level
// for a mapped entry (i.e. an entry that is contained within an opened .zip,
// cannot be changed.) The default compression level is 6. A compression level of 0
// is equivalent to no compression. The maximum compression level is 9.
// 
// The zip.SetCompressionLevel method must be called after appending the files
// (i.e. after the calls to AppendFile*, AppendData, or AppendOneFileOrDir).
// 
// A single call to SetCompressionLevel will set the compression level for all
// existing file and data entries.
//
func (z *Zip) SetCompressionLevel(arg1 int)  {

    C.CkZip_SetCompressionLevel(z.ckObj, C.int(arg1))



}


// Specify a collection of exclusion patterns to be used when adding files to a
// Zip. Each pattern in the collection can use the "*" wildcard character, where
// "*" indicates 0 or more occurrences of any character.
func (z *Zip) SetExclusions(arg1 *StringArray)  {

    C.CkZip_SetExclusions(z.ckObj, arg1.ckObj)



}


// Set the password for an encrypted or password-protected Zip.
func (z *Zip) SetPassword(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkZip_SetPassword(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Unlocks the component allowing for the full functionality to be used. If a
// purchased unlock code is passed, there is no expiration. Any other string
// automatically begins a fully-functional 30-day trial the first time
// UnlockComponent is called.
func (z *Zip) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkZip_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unzips and returns the number of files unzipped, or -1 if a failure occurs.
// Subdirectories are automatically created during the unzipping process.
func (z *Zip) Unzip(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkZip_Unzip(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the Unzip method
func (z *Zip) UnzipAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkZip_UnzipAsync(z.ckObj, cstr1)

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


// Unzips and returns the number of files unzipped, or -1 if a failure occurs. All
// files in the Zip are unzipped into the specfied dirPath regardless of the
// directory path information contained in the Zip. This has the effect of
// collapsing all files into a single directory. If several files in the Zip have
// the same name, the files unzipped last will overwrite the files already
// unzipped.
func (z *Zip) UnzipInto(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkZip_UnzipInto(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the UnzipInto method
func (z *Zip) UnzipIntoAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkZip_UnzipIntoAsync(z.ckObj, cstr1)

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


// Same as Unzip, but only unzips files matching a pattern. If no wildcard
// characters ('*') are used, then only files that exactly match the pattern will
// be unzipped. The "*" characters matches 0 or more of any character.
func (z *Zip) UnzipMatching(arg1 string, arg2 string, arg3 bool) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkZip_UnzipMatching(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Asynchronous version of the UnzipMatching method
func (z *Zip) UnzipMatchingAsync(arg1 string, arg2 string, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkZip_UnzipMatchingAsync(z.ckObj, cstr1, cstr2, b3)

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


// Unzips matching files into a single directory, ignoring all path information
// stored in the Zip.
func (z *Zip) UnzipMatchingInto(arg1 string, arg2 string, arg3 bool) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkZip_UnzipMatchingInto(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Asynchronous version of the UnzipMatchingInto method
func (z *Zip) UnzipMatchingIntoAsync(arg1 string, arg2 string, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkZip_UnzipMatchingIntoAsync(z.ckObj, cstr1, cstr2, b3)

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


// Same as Unzip, but only files that don't already exist on disk, or have later
// file modification dates are unzipped.
func (z *Zip) UnzipNewer(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkZip_UnzipNewer(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the UnzipNewer method
func (z *Zip) UnzipNewerAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkZip_UnzipNewerAsync(z.ckObj, cstr1)

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


// Tests the current DecryptPassword setting against the currently opened zip.
// Returns true if the password is valid, otherwise returns false.
func (z *Zip) VerifyPassword() bool {

    v := C.CkZip_VerifyPassword(z.ckObj)


    return v != 0
}


// Same as WriteZip, but instead of writing the Zip to a file, it writes to binData.
// Zips that are written to binData can be opened by calling OpenBd. Note: Both
// WriteBd and OpenBd are added in Chilkat v9.5.0.66
func (z *Zip) WriteBd(arg1 *BinData) bool {

    v := C.CkZip_WriteBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the WriteBd method
func (z *Zip) WriteBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkZip_WriteBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Same as WriteZip, but instead of writing the Zip to a file, it writes to memory.
// Zips that are written to memory can also be opened from memory by calling
// OpenFromMemory.
func (z *Zip) WriteToMemory() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkZip_WriteToMemory(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the WriteToMemory method
func (z *Zip) WriteToMemoryAsync(c chan *Task) {

    hTask := C.CkZip_WriteToMemoryAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Saves the Zip to a file and implictly re-opens it so further operations can
// continue. Use WriteZipAndClose to write and close the Zip. There is no
// limitation on the size of files that may be contained within a .zip, the total
// number of files in a .zip, or the total size of a .zip. If necessary, WriteZip
// will use the ZIP64 file format extensions when 4GB or file count limitations of
// the old zip file format are exceeded.
func (z *Zip) WriteZip() bool {

    v := C.CkZip_WriteZip(z.ckObj)


    return v != 0
}


// Asynchronous version of the WriteZip method
func (z *Zip) WriteZipAsync(c chan *Task) {

    hTask := C.CkZip_WriteZipAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Saves the Zip to a file and closes it. On return, the Zip object will be in the
// state as if NewZip had been called. There is no limitation on the size of files
// that may be contained within a .zip, the total number of files in a .zip, or the
// total size of a .zip. If necessary, WriteZip will use the ZIP64 file format
// extensions when 4GB or file count limitations of the old zip file format are
// exceeded.
func (z *Zip) WriteZipAndClose() bool {

    v := C.CkZip_WriteZipAndClose(z.ckObj)


    return v != 0
}


// Asynchronous version of the WriteZipAndClose method
func (z *Zip) WriteZipAndCloseAsync(c chan *Task) {

    hTask := C.CkZip_WriteZipAndCloseAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


