// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkAuthGoogle.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewAuthGoogle() *AuthGoogle {
	return &AuthGoogle{C.CkAuthGoogle_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *AuthGoogle) DisposeAuthGoogle() {
    if z != nil {
	C.CkAuthGoogle_Dispose(z.ckObj)
	}
}




func (z *AuthGoogle) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkAuthGoogle_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *AuthGoogle) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkAuthGoogle_setExternalProgress(z.ckObj,1)
}

func (z *AuthGoogle) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkAuthGoogle_setExternalProgress(z.ckObj,1)
}

func (z *AuthGoogle) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkAuthGoogle_setExternalProgress(z.ckObj,1)
}




// property: AccessToken
// The access token to be used in Google API requests. This property is set on a
// successful call to ObtainAccessToken.
// Important: This class is used for authenticating calls to the Google Cloud Platform API and Google Apps API using a service account.. 
// For 3-legged OAuth2, where a browser must be used to interactively get permission from the Google account owner, use the Chilkat OAuth2 class/object.
func (z *AuthGoogle) AccessToken() string {
    return C.GoString(C.CkAuthGoogle_accessToken(z.ckObj))
}
// property setter: AccessToken
func (z *AuthGoogle) SetAccessToken(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthGoogle_putAccessToken(z.ckObj,cStr)
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
func (z *AuthGoogle) DebugLogFilePath() string {
    return C.GoString(C.CkAuthGoogle_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *AuthGoogle) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthGoogle_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EmailAddress
// The client email address of the service account. If a JSON key is used, then the
// client_email should already be specified within the JSON key, and this property
// is unused. This property must be set if using a P12 key.
func (z *AuthGoogle) EmailAddress() string {
    return C.GoString(C.CkAuthGoogle_emailAddress(z.ckObj))
}
// property setter: EmailAddress
func (z *AuthGoogle) SetEmailAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthGoogle_putEmailAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ExpireNumSeconds
// The expiration time, in seconds, of the access token to be requested. The
// maximum value is 1 hour (3600 seconds). The default value is 3600.
func (z *AuthGoogle) ExpireNumSeconds() int {
    return int(C.CkAuthGoogle_getExpireNumSeconds(z.ckObj))
}
// property setter: ExpireNumSeconds
func (z *AuthGoogle) SetExpireNumSeconds(value int) {
    C.CkAuthGoogle_putExpireNumSeconds(z.ckObj,C.int(value))
}

// property: Iat
// This property can be set to override the default current date/time value for the
// "iat" claim of the JWT. It can be set to a value indicating the number of
// seconds from 1970-01-01T00:00:00Z UTC.
// 
// The default value is 0, which indicates to use the iat value for the current
// system date/time. Unless explicitly needed, always leave this property at the
// default value.
//
func (z *AuthGoogle) Iat() int {
    return int(C.CkAuthGoogle_getIat(z.ckObj))
}
// property setter: Iat
func (z *AuthGoogle) SetIat(value int) {
    C.CkAuthGoogle_putIat(z.ckObj,C.int(value))
}

// property: JsonKey
// The JSON key for obtaining an access token. An application must set either the
// P12 or JSON private key, but not both.
func (z *AuthGoogle) JsonKey() string {
    return C.GoString(C.CkAuthGoogle_jsonKey(z.ckObj))
}
// property setter: JsonKey
func (z *AuthGoogle) SetJsonKey(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthGoogle_putJsonKey(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthGoogle) LastErrorHtml() string {
    return C.GoString(C.CkAuthGoogle_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *AuthGoogle) LastErrorText() string {
    return C.GoString(C.CkAuthGoogle_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthGoogle) LastErrorXml() string {
    return C.GoString(C.CkAuthGoogle_lastErrorXml(z.ckObj))
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
func (z *AuthGoogle) LastMethodSuccess() bool {
    v := int(C.CkAuthGoogle_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *AuthGoogle) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthGoogle_putLastMethodSuccess(z.ckObj,v)
}

// property: NumSecondsRemaining
// If the access token is valid, contains the number of seconds remaining until it
// expires. A value of 0 indicates an invalid or expired access token.
func (z *AuthGoogle) NumSecondsRemaining() int {
    return int(C.CkAuthGoogle_getNumSecondsRemaining(z.ckObj))
}

// property: Scope
// A space-delimited list of the permissions that the application requests.
func (z *AuthGoogle) Scope() string {
    return C.GoString(C.CkAuthGoogle_scope(z.ckObj))
}
// property setter: Scope
func (z *AuthGoogle) SetScope(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthGoogle_putScope(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SubEmailAddress
// The email address of the user for which the application is requesting delegated
// access.
func (z *AuthGoogle) SubEmailAddress() string {
    return C.GoString(C.CkAuthGoogle_subEmailAddress(z.ckObj))
}
// property setter: SubEmailAddress
func (z *AuthGoogle) SetSubEmailAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthGoogle_putSubEmailAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Valid
// true if the AccessToken property contains a valid non-expired access token
// obtained via the call to ObtainAccessToken.
func (z *AuthGoogle) Valid() bool {
    v := int(C.CkAuthGoogle_getValid(z.ckObj))
    return v != 0
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *AuthGoogle) VerboseLogging() bool {
    v := int(C.CkAuthGoogle_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *AuthGoogle) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthGoogle_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *AuthGoogle) Version() string {
    return C.GoString(C.CkAuthGoogle_version(z.ckObj))
}
// Returns the private key in a PFX (P12) object. This is only possible if the
// private key was previously set by calling SetP12.
func (z *AuthGoogle) GetP12() *Pfx {

    retObj := C.CkAuthGoogle_GetP12(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Pfx{retObj}
}


// Loads the caller of the task's async method.
func (z *AuthGoogle) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkAuthGoogle_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sends the HTTP request to fetch the access token. When this method completes
// successfully, the access token is available in the AccessToken property. The
// connection is an existing connection to www.googleapis.com.
// 
// Important: Make sure your computer's date/time is accurately set to the current
// date/time, otherwise you'll get a 400 response status code with this error:
// "Invalid JWT: Token must be a short-lived token (60 minutes) and in a reasonable
// timeframe. Check your iat and exp values and use a clock with skew to account
// for clock differences between systems.".
//
func (z *AuthGoogle) ObtainAccessToken(arg1 *Socket) bool {

    v := C.CkAuthGoogle_ObtainAccessToken(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the ObtainAccessToken method
func (z *AuthGoogle) ObtainAccessTokenAsync(arg1 *Socket, c chan *Task) {

    hTask := C.CkAuthGoogle_ObtainAccessTokenAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sets the P12 private key to be used for obtaining an access token. An
// application must set either the P12 or JSON private key, but not both.
func (z *AuthGoogle) SetP12(arg1 *Pfx) bool {

    v := C.CkAuthGoogle_SetP12(z.ckObj, arg1.ckObj)


    return v != 0
}


