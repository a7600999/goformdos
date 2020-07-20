// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkAuthAzureStorage.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewAuthAzureStorage() *AuthAzureStorage {
	return &AuthAzureStorage{C.CkAuthAzureStorage_Create()}
}

func (z *AuthAzureStorage) DisposeAuthAzureStorage() {
    if z != nil {
	C.CkAuthAzureStorage_Dispose(z.ckObj)
	}
}




// property: AccessKey
// A valid base64 access key for the Azure storage account.
func (z *AuthAzureStorage) AccessKey() string {
    return C.GoString(C.CkAuthAzureStorage_accessKey(z.ckObj))
}
// property setter: AccessKey
func (z *AuthAzureStorage) SetAccessKey(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureStorage_putAccessKey(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Account
// The Azure storage account name. (A storage account can contain zero or more
// containers. A container contains properties, metadata, and zero or more blobs. A
// blob is any single entity comprised of binary data, properties, and metadata. )
func (z *AuthAzureStorage) Account() string {
    return C.GoString(C.CkAuthAzureStorage_account(z.ckObj))
}
// property setter: Account
func (z *AuthAzureStorage) SetAccount(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureStorage_putAccount(z.ckObj,cStr)
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
func (z *AuthAzureStorage) DebugLogFilePath() string {
    return C.GoString(C.CkAuthAzureStorage_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *AuthAzureStorage) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureStorage_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthAzureStorage) LastErrorHtml() string {
    return C.GoString(C.CkAuthAzureStorage_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *AuthAzureStorage) LastErrorText() string {
    return C.GoString(C.CkAuthAzureStorage_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthAzureStorage) LastErrorXml() string {
    return C.GoString(C.CkAuthAzureStorage_lastErrorXml(z.ckObj))
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
func (z *AuthAzureStorage) LastMethodSuccess() bool {
    v := int(C.CkAuthAzureStorage_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *AuthAzureStorage) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthAzureStorage_putLastMethodSuccess(z.ckObj,v)
}

// property: Scheme
// Can be "SharedKey" or "SharedKeyLite". The default value is "SharedKey".
func (z *AuthAzureStorage) Scheme() string {
    return C.GoString(C.CkAuthAzureStorage_scheme(z.ckObj))
}
// property setter: Scheme
func (z *AuthAzureStorage) SetScheme(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureStorage_putScheme(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Service
// Can be "Blob", "Queue", "File", or "Table". The default is "Blob".
// 
// Note: Authentication for the "Table" service did not work in versions prior to
// v9.5.0.83.
//
func (z *AuthAzureStorage) Service() string {
    return C.GoString(C.CkAuthAzureStorage_service(z.ckObj))
}
// property setter: Service
func (z *AuthAzureStorage) SetService(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureStorage_putService(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *AuthAzureStorage) VerboseLogging() bool {
    v := int(C.CkAuthAzureStorage_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *AuthAzureStorage) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthAzureStorage_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *AuthAzureStorage) Version() string {
    return C.GoString(C.CkAuthAzureStorage_version(z.ckObj))
}

// property: XMsVersion
// If set, automatically adds the "x-ms-version" HTTP request header to Azure
// Storage requests. The default value is "2014-02-14".
func (z *AuthAzureStorage) XMsVersion() string {
    return C.GoString(C.CkAuthAzureStorage_xMsVersion(z.ckObj))
}
// property setter: XMsVersion
func (z *AuthAzureStorage) SetXMsVersion(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureStorage_putXMsVersion(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}
