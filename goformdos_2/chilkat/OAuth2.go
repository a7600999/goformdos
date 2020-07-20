// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkOAuth2.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewOAuth2() *OAuth2 {
	return &OAuth2{C.CkOAuth2_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *OAuth2) DisposeOAuth2() {
    if z != nil {
	C.CkOAuth2_Dispose(z.ckObj)
	}
}




func (z *OAuth2) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkOAuth2_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *OAuth2) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkOAuth2_setExternalProgress(z.ckObj,1)
}

func (z *OAuth2) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkOAuth2_setExternalProgress(z.ckObj,1)
}

func (z *OAuth2) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkOAuth2_setExternalProgress(z.ckObj,1)
}




// property: AccessToken
// When the OAuth2 three-legged authorization has successfully completed in the
// background thread, this property contains the access_token.
// 
// For example, a successful Google API JSON response looks like this:
//  {
//              "access_token": "ya29.Ci9ZA-Z0Q7vtnch8xxxxxxxxxxxxxxgDVOOV97-IBvTt958xxxxxx1sasw",
//              "token_type": "Bearer",
// 
//             "expires_in": 3600,
// 
//             "refresh_token": "1/fYjEVR-3Oq9xxxxxxxxxxxxxxLzPtlNOeQ"
// }
//
func (z *OAuth2) AccessToken() string {
    return C.GoString(C.CkOAuth2_accessToken(z.ckObj))
}
// property setter: AccessToken
func (z *OAuth2) SetAccessToken(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putAccessToken(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AccessTokenResponse
// When the OAuth2 three-legged authorization has completed in the background
// thread, this property contains the response that contains the access_token, the
// optional refresh_token, and any other information included in the final
// response. If the authorization was denied, then this contains the error
// response.
// 
// For example, a successful JSON response for a Google API looks like this:
//  {
//              "access_token": "ya29.Ci9ZA-Z0Q7vtnch8xxxxxxxxxxxxxxgDVOOV97-IBvTt958xxxxxx1sasw",
//              "token_type": "Bearer",
// 
//             "expires_in": 3600,
// 
//             "refresh_token": "1/fYjEVR-3Oq9xxxxxxxxxxxxxxLzPtlNOeQ"
// }
// 
// Note: Not all responses are JSON. A successful Facebook response is plain text
// and looks like this:
// access_token=EAAZALuOC1wAwBAKH6FKnxOkjfEP ... UBZBhYD5hSVBETBx6AZD&expires=5134653
//
func (z *OAuth2) AccessTokenResponse() string {
    return C.GoString(C.CkOAuth2_accessTokenResponse(z.ckObj))
}

// property: AppCallbackUrl
// Some OAuth2 services, such as QuickBooks, do not allow for
// "http://localhost:port" callback URLs. When this is the case, a desktop app
// cannot pop up a browser and expect to get the final redirect callback. The
// workaround is to set this property to a URI on your web server, which sends a
// response to redirect back to "http://localhost:3017". Thus the callback becomes
// a double redirect, which ends at localhost:port, and thus completes the circuit.
// 
// If the OAuth2 service allows for "http://localhost:port" callback URLs, then
// leave this property empty.
// 
// As an example, one could set this property to
// "https://www.yourdomain.com/OAuth2.php", where the PHP source contains the
// following:
//
func (z *OAuth2) AppCallbackUrl() string {
    return C.GoString(C.CkOAuth2_appCallbackUrl(z.ckObj))
}
// property setter: AppCallbackUrl
func (z *OAuth2) SetAppCallbackUrl(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putAppCallbackUrl(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AuthFlowState
// Indicates the current progress of the OAuth2 three-legged authorization flow.
// Possible values are:
// 
// 0: Idle. No OAuth2 has yet been attempted.
// 1: Waiting for Redirect. The OAuth2 background thread is waiting to receive the
// redirect HTTP request from the browser.
// 2: Waiting for Final Response. The OAuth2 background thread is waiting for the
// final access token response.
// 3: Completed with Success. The OAuth2 flow has completed, the background thread
// exited, and the successful JSON response is available in AccessTokenResponse
// property.
// 4: Completed with Access Denied. The OAuth2 flow has completed, the background
// thread exited, and the error JSON is available in AccessTokenResponse property.
// 5: Failed Prior to Completion. The OAuth2 flow failed to complete, the
// background thread exited, and the error information is available in the
// FailureInfo property.
//
func (z *OAuth2) AuthFlowState() int {
    return int(C.CkOAuth2_getAuthFlowState(z.ckObj))
}

// property: AuthorizationEndpoint
// The URL used to obtain an authorization grant. For example, the Google APIs
// authorization endpoint is "https://accounts.google.com/o/oauth2/v2/auth". (In
// three-legged OAuth2, this is the very first point of contact that begins the
// OAuth2 authentication flow.)
func (z *OAuth2) AuthorizationEndpoint() string {
    return C.GoString(C.CkOAuth2_authorizationEndpoint(z.ckObj))
}
// property setter: AuthorizationEndpoint
func (z *OAuth2) SetAuthorizationEndpoint(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putAuthorizationEndpoint(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ClientId
// The "client_id" that identifies the application.
// 
// For example, if creating an app to use a Google API, one would create a client
// ID by:
//     Logging into the Google API Console (https://console.developers.google.com).
//     Navigate to "Credentials".
//     Click on "Create Credentials"
//     Choose "OAuth client ID"
//     Select the "Other" application type.
//     Name your app and click "Create", and a client_id and client_secret will be
//     generated.
// Other API's, such as Facebook, should have something similar for generating a
// client ID and client secret.
//
func (z *OAuth2) ClientId() string {
    return C.GoString(C.CkOAuth2_clientId(z.ckObj))
}
// property setter: ClientId
func (z *OAuth2) SetClientId(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putClientId(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ClientSecret
// The "client_secret" for the application. Application credentials (i.e. what
// identifies the application) consist of a client_id and client_secret. See the
// ClientId property for more information.
// 
// Is the Client Secret Really a Secret?
// 
// This deserves some explanation. For a web-based application (where the code is
// on the web server) and the user interacts with the application in a browser,
// then YES, the client secret MUST be kept secret at all times. One does not want
// to be interacting with a site that claims to be "Application XYZ" but is
// actually an impersonator. But the Chilkat OAuth2 class is for desktop
// applications and scripts (i.e. things that run on the local computer, not in a
// browser).
// 
// Consider Mozilla Thunderbird. It is an application installed on your computer.
// Thunderbird uses OAuth2 authentication for GMail accounts in the same way as
// this OAuth2 API. When you add a GMail account and need to authenticate for the
// 1st time, you'll get a popup window (a browser) where you interactively grant
// authorization to Thunderbird. You implicitly know the Thunderbird application is
// running because you started it. There can be no impersonation unless your
// computer has already been hacked and when you thought you started Thunderbird,
// you actually started some rogue app. But if you already started some rogue app,
// then all has already been lost.
// 
// It is essentially impossible for desktop applications to embed a secret key
// (such as the client secret) and assure confidentiality (i.e. that the key cannot
// be obtained by some hacker. An application can hide the secret, and can make it
// difficult to access, but in the end the secret cannot be assumed to be safe.
// Therefore, the client_secret, for desktop (installed) applications is not
// actually secret. One should still take care to shroud the client secret to some
// extent, but know that whatever is done cannot be deemed secure. But this is OK.
// The reason it is OK is that implicitly, when a person starts an application
// (such as Thunderbird), the identity of the application is known. If a fake
// Thunderbird was started, then all has already been lost. The security of the
// system is in preventing the fake/rogue applications in the 1st place. If that
// security has already been breached, then nothing else really matters.
//
func (z *OAuth2) ClientSecret() string {
    return C.GoString(C.CkOAuth2_clientSecret(z.ckObj))
}
// property setter: ClientSecret
func (z *OAuth2) SetClientSecret(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putClientSecret(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CodeChallenge
// Optional. Set this to true to send a code_challenge (as per RFC 7636) with the
// authorization request. The default value is false.
func (z *OAuth2) CodeChallenge() bool {
    v := int(C.CkOAuth2_getCodeChallenge(z.ckObj))
    return v != 0
}
// property setter: CodeChallenge
func (z *OAuth2) SetCodeChallenge(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkOAuth2_putCodeChallenge(z.ckObj,v)
}

// property: CodeChallengeMethod
// Optional. Only applies when the CodeChallenge property is set to true.
// Possible values are "plain" or "S256". The default is "S256".
func (z *OAuth2) CodeChallengeMethod() string {
    return C.GoString(C.CkOAuth2_codeChallengeMethod(z.ckObj))
}
// property setter: CodeChallengeMethod
func (z *OAuth2) SetCodeChallengeMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putCodeChallengeMethod(z.ckObj,cStr)
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
func (z *OAuth2) DebugLogFilePath() string {
    return C.GoString(C.CkOAuth2_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *OAuth2) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: FailureInfo
// If the OAuth2 three-legged authorization failed prior to completion (the
// AuthFlowState = 5), then information about the failure is contained in this
// property. This property is automatically cleared when OAuth2 authorization
// starts (i.e. when StartAuth is called).
func (z *OAuth2) FailureInfo() string {
    return C.GoString(C.CkOAuth2_failureInfo(z.ckObj))
}

// property: IncludeNonce
// Optional. Set this to true to send a nonce with the authorization request. The
// default value is false.
func (z *OAuth2) IncludeNonce() bool {
    v := int(C.CkOAuth2_getIncludeNonce(z.ckObj))
    return v != 0
}
// property setter: IncludeNonce
func (z *OAuth2) SetIncludeNonce(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkOAuth2_putIncludeNonce(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *OAuth2) LastErrorHtml() string {
    return C.GoString(C.CkOAuth2_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *OAuth2) LastErrorText() string {
    return C.GoString(C.CkOAuth2_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *OAuth2) LastErrorXml() string {
    return C.GoString(C.CkOAuth2_lastErrorXml(z.ckObj))
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
func (z *OAuth2) LastMethodSuccess() bool {
    v := int(C.CkOAuth2_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *OAuth2) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkOAuth2_putLastMethodSuccess(z.ckObj,v)
}

// property: ListenPort
// The port number to listen for the redirect URI request sent by the browser. If
// set to 0, then a random unused port is used. The default value of this property
// is 0.
// 
// In most cases, using a random unused port is the best choice. In some OAuth2
// situations, such as with Facebook, a specific port number must be chosen. This
// is due to the fact that Facebook requires an APP to have a Site URL, which must
// exactly match the redirect_uri used in OAuth2 authorization. For example, the
// Facebook Site URL might be "http://localhost:3017/" if port 3017 is the listen
// port.
//
func (z *OAuth2) ListenPort() int {
    return int(C.CkOAuth2_getListenPort(z.ckObj))
}
// property setter: ListenPort
func (z *OAuth2) SetListenPort(value int) {
    C.CkOAuth2_putListenPort(z.ckObj,C.int(value))
}

// property: ListenPortRangeEnd
// If set, then an unused port will be chosen in the range from the ListenPort
// property to this property. Some OAuth2 services, such as Google, require that
// callback URL's, including port numbers, be selected in advance. This feature
// allows for a range of callback URL's to be specified to cope with the
// possibility that another application on the same computer might be using a
// particular port.
// 
// For example, a Google ClientID might be configured with a set of authorized
// callback URI's such as:
//     http://localhost:55110/
//     http://localhost:55112/
//     http://localhost:55113/
//     http://localhost:55114/
//     http://localhost:55115/
//     http://localhost:55116/
//     http://localhost:55117/
// 
// In which case the ListenPort property would be set to 55110, and this property
// would be set to 55117.
//
func (z *OAuth2) ListenPortRangeEnd() int {
    return int(C.CkOAuth2_getListenPortRangeEnd(z.ckObj))
}
// property setter: ListenPortRangeEnd
func (z *OAuth2) SetListenPortRangeEnd(value int) {
    C.CkOAuth2_putListenPortRangeEnd(z.ckObj,C.int(value))
}

// property: LocalHost
// Defaults to "localhost". This should typically remain at the default value. It
// is the loopback domain or IP address used for the redirect_uri. For example,
// "http://localhost:2012/". (assuming 2012 was used or randomly chosen as the
// listen port number) If the desired redirect_uri is to be
// "http://127.0.0.1:2012/", then set this property equal to "127.0.0.1".
func (z *OAuth2) LocalHost() string {
    return C.GoString(C.CkOAuth2_localHost(z.ckObj))
}
// property setter: LocalHost
func (z *OAuth2) SetLocalHost(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putLocalHost(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: NonceLength
// Defines the length of the nonce in bytes. The nonce is only included if the
// IncludeNonce property = true. (The length of the nonce in characters will be
// twice the length in bytes, because the nonce is a hex string.)
// 
// The default nonce length is 4 bytes.
//
func (z *OAuth2) NonceLength() int {
    return int(C.CkOAuth2_getNonceLength(z.ckObj))
}
// property setter: NonceLength
func (z *OAuth2) SetNonceLength(value int) {
    C.CkOAuth2_putNonceLength(z.ckObj,C.int(value))
}

// property: RedirectAllowHtml
// The HTML returned to the browser when access is allowed by the end-user. The
// default value is HTMl that contains a META refresh to
// https://www.chilkatsoft.com/oauth2_allowed.html. Your application should set
// this property to display whatever HTML is desired when access is granted.
func (z *OAuth2) RedirectAllowHtml() string {
    return C.GoString(C.CkOAuth2_redirectAllowHtml(z.ckObj))
}
// property setter: RedirectAllowHtml
func (z *OAuth2) SetRedirectAllowHtml(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putRedirectAllowHtml(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: RedirectDenyHtml
// The HTML returned to the browser when access is denied by the end-user. The
// default value is HTMl that contains a META refresh to
// https://www.chilkatsoft.com/oauth2_denied.html. Your application should set this
// property to display whatever HTML is desired when access is denied.
func (z *OAuth2) RedirectDenyHtml() string {
    return C.GoString(C.CkOAuth2_redirectDenyHtml(z.ckObj))
}
// property setter: RedirectDenyHtml
func (z *OAuth2) SetRedirectDenyHtml(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putRedirectDenyHtml(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: RefreshToken
// When the OAuth2 three-legged authorization has successfully completed in the
// background thread, this property contains the refresh_token, if present.
// 
// For example, a successful Google API JSON response looks like this:
//  {
//              "access_token": "ya29.Ci9ZA-Z0Q7vtnch8xxxxxxxxxxxxxxgDVOOV97-IBvTt958xxxxxx1sasw",
//              "token_type": "Bearer",
// 
//             "expires_in": 3600,
// 
//             "refresh_token": "1/fYjEVR-3Oq9xxxxxxxxxxxxxxLzPtlNOeQ"
// }
//
func (z *OAuth2) RefreshToken() string {
    return C.GoString(C.CkOAuth2_refreshToken(z.ckObj))
}
// property setter: RefreshToken
func (z *OAuth2) SetRefreshToken(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putRefreshToken(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Resource
// This is an optional setting that defines the "resource" query parameter. For
// example, to call the Microsoft Graph API, set this property value to
// "https://graph.microsoft.com/". The Microsoft Dynamics CRM OAuth authentication
// also requires the Resource property.
func (z *OAuth2) Resource() string {
    return C.GoString(C.CkOAuth2_resource(z.ckObj))
}
// property setter: Resource
func (z *OAuth2) SetResource(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putResource(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ResponseMode
// Can be set to "form_post" to include a "response_mode=form_post" in the
// authorization request. The default value is the empty string to omit the
// "response_mode" query param.
func (z *OAuth2) ResponseMode() string {
    return C.GoString(C.CkOAuth2_responseMode(z.ckObj))
}
// property setter: ResponseMode
func (z *OAuth2) SetResponseMode(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putResponseMode(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ResponseType
// The default value is "code". Can be set to "id_token+code" for cases where
// "response_type=id_token+code" is required in the authorization request.
func (z *OAuth2) ResponseType() string {
    return C.GoString(C.CkOAuth2_responseType(z.ckObj))
}
// property setter: ResponseType
func (z *OAuth2) SetResponseType(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putResponseType(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Scope
// This is an optional setting that defines the scope of access. For example,
// Google API scopes are listed here:
// https://developers.google.com/identity/protocols/googlescopes
// 
// For example, if wishing to grant OAuth2 authorization for Google Drive, one
// would set this property to "https://www.googleapis.com/auth/drive".
//
func (z *OAuth2) Scope() string {
    return C.GoString(C.CkOAuth2_scope(z.ckObj))
}
// property setter: Scope
func (z *OAuth2) SetScope(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putScope(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TokenEndpoint
// The URL for exchanging an authorization grant for an access token. For example,
// the Google APIs token endpoint is "https://www.googleapis.com/oauth2/v4/token".
// (In three-legged OAuth2, this is the very last point of contact that ends the
// OAuth2 authentication flow.)
func (z *OAuth2) TokenEndpoint() string {
    return C.GoString(C.CkOAuth2_tokenEndpoint(z.ckObj))
}
// property setter: TokenEndpoint
func (z *OAuth2) SetTokenEndpoint(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putTokenEndpoint(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TokenType
// When the OAuth2 three-legged authorization has successfully completed in the
// background thread, this property contains the token_type, if present.
// 
// A successful Google API JSON response looks like this:
//  {
//              "access_token": "ya29.Ci9ZA-Z0Q7vtnch8xxxxxxxxxxxxxxgDVOOV97-IBvTt958xxxxxx1sasw",
//              "token_type": "Bearer",
// 
//             "expires_in": 3600,
// 
//             "refresh_token": "1/fYjEVR-3Oq9xxxxxxxxxxxxxxLzPtlNOeQ"
// }
// 
// Note: Some responses may not included a "token_type" param. In that case, this
// property will remain empty.
//
func (z *OAuth2) TokenType() string {
    return C.GoString(C.CkOAuth2_tokenType(z.ckObj))
}
// property setter: TokenType
func (z *OAuth2) SetTokenType(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth2_putTokenType(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UseBasicAuth
// If set to true, then the internal POST (on the background thread) that
// exchanges the code for an access token will send the client_id/client_secret in
// an "Authorization Basic ..." header where the client_id is the login and the
// client_secret is the password.
// 
// Some services, such as fitbit.com, require the client_id/client_secret to be
// passed in this way.
// 
// The default value of this property is false, which causes the
// client_id/client_secret to be sent as query params.
//
func (z *OAuth2) UseBasicAuth() bool {
    v := int(C.CkOAuth2_getUseBasicAuth(z.ckObj))
    return v != 0
}
// property setter: UseBasicAuth
func (z *OAuth2) SetUseBasicAuth(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkOAuth2_putUseBasicAuth(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *OAuth2) VerboseLogging() bool {
    v := int(C.CkOAuth2_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *OAuth2) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkOAuth2_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *OAuth2) Version() string {
    return C.GoString(C.CkOAuth2_version(z.ckObj))
}
// Cancels an OAuth2 authorization flow that is in progress.
func (z *OAuth2) Cancel() bool {

    v := C.CkOAuth2_Cancel(z.ckObj)


    return v != 0
}


// Some OAuth2 providers can provide additional parameters in the redirect request
// sent to the local listener (i.e. the Chilkat background thread). One such case
// is for QuickBooks, It contains a realmId parameter such as the following:
// http://localhost:55568/?state=xxxxxxxxxxxx&code=xxxxxxxxxxxx&realmId=1234567890
// 
// After the OAuth2 authentication is completed, an application can call this
// method to get any of the parameter values. For example, to get the realmId
// value, pass "realmId" in paramName.
//
// return a string or nil if failed.
func (z *OAuth2) GetRedirectRequestParam(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkOAuth2_getRedirectRequestParam(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads the caller of the task's async method.
func (z *OAuth2) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkOAuth2_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Monitors an already started OAuth2 authorization flow and returns when it is
// finished.
// 
// Note: It rarely makes sense to call this method. If this programming language
// supports callbacks, then MonitorAsync is a better choice. (See the Oauth2
// project repositories at https://github.com/chilkatsoft for samples.) If a
// programming language does not have callbacks, a better choice is to periodically
// check the AuthFlowState property for a value >= 3. If there is no response from
// the browser, the background thread (that is waiting on the browser) can be
// cancelled by calling the Cancel method.
//
func (z *OAuth2) Monitor() bool {

    v := C.CkOAuth2_Monitor(z.ckObj)


    return v != 0
}


// Asynchronous version of the Monitor method
func (z *OAuth2) MonitorAsync(c chan *Task) {

    hTask := C.CkOAuth2_MonitorAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a refresh request to the token endpoint to obtain a new access token.
// After a successful refresh request, the AccessToken and RefreshToken properties
// will be updated with new values.
// 
// Note: This method can only be called if the ClientId, ClientSecret, RefreshToken
// and TokenEndpoint properties contain valid values.
//
func (z *OAuth2) RefreshAccessToken() bool {

    v := C.CkOAuth2_RefreshAccessToken(z.ckObj)


    return v != 0
}


// Asynchronous version of the RefreshAccessToken method
func (z *OAuth2) RefreshAccessTokenAsync(c chan *Task) {

    hTask := C.CkOAuth2_RefreshAccessTokenAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Provides for the ability to add HTTP request headers for the request sent by the
// RefreshAccesToken method. For example, if the "Accept: application/json" header
// needs to be sent, then add it by calling this method with name = "Accept" and
// value = "application/json".
// 
// Multiple headers may be added by calling this method once for each. To remove a
// header, call this method with name equal to the header name, and with an empty
// string for value.
//
func (z *OAuth2) SetRefreshHeader(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkOAuth2_SetRefreshHeader(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SetRefreshHeader method
func (z *OAuth2) SetRefreshHeaderAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkOAuth2_SetRefreshHeaderAsync(z.ckObj, cstr1, cstr2)

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


// Convenience method to force the calling thread to sleep for a number of
// milliseconds.
func (z *OAuth2) SleepMs(arg1 int)  {

    C.CkOAuth2_SleepMs(z.ckObj, C.int(arg1))



}


// Initiates the three-legged OAuth2 flow. The various properties, such as
// ClientId, ClientSecret, Scope, CodeChallenge, AuthorizationEndpoint, and
// TokenEndpoint, should be set prior to calling this method.
// 
// This method does two things:
//     Forms and returns a URL that is to be loaded in a browser.
//     Starts a background thread that listens on a randomly selected unused port
//     to receive the redirect request from the browser. The receiving of the request
//     from the browser, and the sending of the HTTP request to complete the
//     three-legged OAuth2 flow is done entirely in the background thread. The
//     application controls this behavior by setting the various properties beforehand.
// The return value is the URL to be loaded (navigated to) in a popup or embedded
// browser.
// 
// Note: It's best not to call StartAuth if a previous call to StartAuth is in a
// non-completed state. However, starting in v9.5.0.76, if a background thread from
// a previous call to StartAuth is still running, it will be automatically
// canceled. However,rather than relying on this automatic behavior, your
// application should explicity Cancel the previous StartAuth before calling again.
//
// return a string or nil if failed.
func (z *OAuth2) StartAuth() *string {

    retStr := C.CkOAuth2_startAuth(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Calling this method is optional, and is only required if a proxy (HTTP or
// SOCKS), an SSH tunnel, or if special connection related socket options need to
// be used. When UseConnection is not called, the connection to the token endpoint
// is a direct connection using TLS (or not) based on the TokenEndpoint. (If the
// TokenEndpoint begins with "https://", then TLS is used.)
// 
// This method sets the socket object to be used for sending the requests to the
// token endpoint in the background thread. The sock can be an already-connected
// socket, or a socket object that is not yet connected. In some cases, such as for
// a connection through an SSH tunnel, the sock must already be connected. In other
// cases, an unconnected sock can be provided because the purpose for providing the
// socket object is to specify settings such as for HTTP or SOCKS proxies.
//
func (z *OAuth2) UseConnection(arg1 *Socket) bool {

    v := C.CkOAuth2_UseConnection(z.ckObj, arg1.ckObj)


    return v != 0
}


