// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkAuthAzureAD.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewAuthAzureAD() *AuthAzureAD {
	return &AuthAzureAD{C.CkAuthAzureAD_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *AuthAzureAD) DisposeAuthAzureAD() {
    if z != nil {
	C.CkAuthAzureAD_Dispose(z.ckObj)
	}
}




func (z *AuthAzureAD) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkAuthAzureAD_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *AuthAzureAD) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkAuthAzureAD_setExternalProgress(z.ckObj,1)
}

func (z *AuthAzureAD) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkAuthAzureAD_setExternalProgress(z.ckObj,1)
}

func (z *AuthAzureAD) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkAuthAzureAD_setExternalProgress(z.ckObj,1)
}




// property: AccessToken
// The access token to be used in Azure AD REST API requests. This property is set
// on a successful call to ObtainAccessToken.
func (z *AuthAzureAD) AccessToken() string {
    return C.GoString(C.CkAuthAzureAD_accessToken(z.ckObj))
}
// property setter: AccessToken
func (z *AuthAzureAD) SetAccessToken(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureAD_putAccessToken(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ClientId
// Specifies the Azure AD client id of the calling web service. To find the calling
// application's client ID, in the Azure Management Portal, click Active Directory,
// click the directory, click the application, and then click Configure.
func (z *AuthAzureAD) ClientId() string {
    return C.GoString(C.CkAuthAzureAD_clientId(z.ckObj))
}
// property setter: ClientId
func (z *AuthAzureAD) SetClientId(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureAD_putClientId(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ClientSecret
// A key registered for the calling web service in Azure AD. To create a key, in
// the Azure Management Portal, click Active Directory, click the directory, click
// the application, and then click Configure.
func (z *AuthAzureAD) ClientSecret() string {
    return C.GoString(C.CkAuthAzureAD_clientSecret(z.ckObj))
}
// property setter: ClientSecret
func (z *AuthAzureAD) SetClientSecret(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureAD_putClientSecret(z.ckObj,cStr)
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
func (z *AuthAzureAD) DebugLogFilePath() string {
    return C.GoString(C.CkAuthAzureAD_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *AuthAzureAD) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureAD_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthAzureAD) LastErrorHtml() string {
    return C.GoString(C.CkAuthAzureAD_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *AuthAzureAD) LastErrorText() string {
    return C.GoString(C.CkAuthAzureAD_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthAzureAD) LastErrorXml() string {
    return C.GoString(C.CkAuthAzureAD_lastErrorXml(z.ckObj))
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
func (z *AuthAzureAD) LastMethodSuccess() bool {
    v := int(C.CkAuthAzureAD_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *AuthAzureAD) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthAzureAD_putLastMethodSuccess(z.ckObj,v)
}

// property: NumSecondsRemaining
// If the access token is valid, contains the number of seconds remaining until it
// expires. A value of 0 indicates an invalid or expired access token.
func (z *AuthAzureAD) NumSecondsRemaining() int {
    return int(C.CkAuthAzureAD_getNumSecondsRemaining(z.ckObj))
}

// property: Resource
// The App ID URI of the receiving web service. To find the App ID URI, in the
// Azure Management Portal, click Active Directory, click the directory, click the
// application, and then click Configure.
func (z *AuthAzureAD) Resource() string {
    return C.GoString(C.CkAuthAzureAD_resource(z.ckObj))
}
// property setter: Resource
func (z *AuthAzureAD) SetResource(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureAD_putResource(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TenantId
// true if the AccessToken property contains a valid non-expired access token
// obtained via the call to ObtainAccessToken.
func (z *AuthAzureAD) TenantId() string {
    return C.GoString(C.CkAuthAzureAD_tenantId(z.ckObj))
}
// property setter: TenantId
func (z *AuthAzureAD) SetTenantId(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureAD_putTenantId(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Valid
// true if the AccessToken property contains a valid non-expired access token
// obtained via the call to ObtainAccessToken.
func (z *AuthAzureAD) Valid() bool {
    v := int(C.CkAuthAzureAD_getValid(z.ckObj))
    return v != 0
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *AuthAzureAD) VerboseLogging() bool {
    v := int(C.CkAuthAzureAD_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *AuthAzureAD) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthAzureAD_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *AuthAzureAD) Version() string {
    return C.GoString(C.CkAuthAzureAD_version(z.ckObj))
}
// Loads the caller of the task's async method.
func (z *AuthAzureAD) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkAuthAzureAD_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sends the HTTP request to fetch the access token. When this method completes
// successfully, the access token is available in the AccessToken property. The
// connection is an existing connection to login.microsoftonline.com.
func (z *AuthAzureAD) ObtainAccessToken(arg1 *Socket) bool {

    v := C.CkAuthAzureAD_ObtainAccessToken(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the ObtainAccessToken method
func (z *AuthAzureAD) ObtainAccessTokenAsync(arg1 *Socket, c chan *Task) {

    hTask := C.CkAuthAzureAD_ObtainAccessTokenAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


