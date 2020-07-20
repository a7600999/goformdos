// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkOAuth1.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewOAuth1() *OAuth1 {
	return &OAuth1{C.CkOAuth1_Create()}
}

func (z *OAuth1) DisposeOAuth1() {
    if z != nil {
	C.CkOAuth1_Dispose(z.ckObj)
	}
}




// property: AuthorizationHeader
// The authorization header. This is what would be included in the Authorization
// HTTP request header if it is going to be used as the means for providing the
// OAuth1 authorization information.
func (z *OAuth1) AuthorizationHeader() string {
    return C.GoString(C.CkOAuth1_authorizationHeader(z.ckObj))
}

// property: BaseString
// This is the exact string that was signed. For example, if the signature method
// is HMAC-SHA1, the BaseString is the exact string that passed to the HMAC-SHA1.
// An application does not set the BaseString property. The BaseString is exposed
// as a property to allow for debugging and to see the exact string that is signed.
func (z *OAuth1) BaseString() string {
    return C.GoString(C.CkOAuth1_baseString(z.ckObj))
}

// property: ConsumerKey
// The consumer key.
func (z *OAuth1) ConsumerKey() string {
    return C.GoString(C.CkOAuth1_consumerKey(z.ckObj))
}
// property setter: ConsumerKey
func (z *OAuth1) SetConsumerKey(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putConsumerKey(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConsumerSecret
// The consumer secret.
func (z *OAuth1) ConsumerSecret() string {
    return C.GoString(C.CkOAuth1_consumerSecret(z.ckObj))
}
// property setter: ConsumerSecret
func (z *OAuth1) SetConsumerSecret(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putConsumerSecret(z.ckObj,cStr)
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
func (z *OAuth1) DebugLogFilePath() string {
    return C.GoString(C.CkOAuth1_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *OAuth1) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EncodedSignature
// The URL encoded representation of the Signature property
func (z *OAuth1) EncodedSignature() string {
    return C.GoString(C.CkOAuth1_encodedSignature(z.ckObj))
}

// property: GeneratedUrl
// The URL that includes the OAuth1 query params.
func (z *OAuth1) GeneratedUrl() string {
    return C.GoString(C.CkOAuth1_generatedUrl(z.ckObj))
}

// property: HmacKey
// The exact HMAC key used to sign the BaseString. An application does not directly
// set the HmacKey. The HmacKey property is read-only and is provided for debugging
// to see the exact HMAC key used to sign the BaseString. The HMAC key is composed
// from the consumer secret (if it exists) and the token secret (if it exists).
func (z *OAuth1) HmacKey() string {
    return C.GoString(C.CkOAuth1_hmacKey(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *OAuth1) LastErrorHtml() string {
    return C.GoString(C.CkOAuth1_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *OAuth1) LastErrorText() string {
    return C.GoString(C.CkOAuth1_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *OAuth1) LastErrorXml() string {
    return C.GoString(C.CkOAuth1_lastErrorXml(z.ckObj))
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
func (z *OAuth1) LastMethodSuccess() bool {
    v := int(C.CkOAuth1_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *OAuth1) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkOAuth1_putLastMethodSuccess(z.ckObj,v)
}

// property: Nonce
// The nonce.
func (z *OAuth1) Nonce() string {
    return C.GoString(C.CkOAuth1_nonce(z.ckObj))
}
// property setter: Nonce
func (z *OAuth1) SetNonce(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putNonce(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OauthMethod
// The HTTP method, such as "GET", "POST", "PUT", "DELETE", or "HEAD". Defaults to
// "GET".
func (z *OAuth1) OauthMethod() string {
    return C.GoString(C.CkOAuth1_oauthMethod(z.ckObj))
}
// property setter: OauthMethod
func (z *OAuth1) SetOauthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putOauthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OauthUrl
// The OAuth URL, such as "http://echo.lab.madgex.com/echo.ashx". See
// http://bettiolo.github.io/oauth-reference-page/ to compare Chilkat results with
// another tool's results.
// 
// Note: The OAuthUrl should not include query parameters. For example, do not set
// the OAuthUrl equal
// tohttps://rest.sandbox.netsuite.com/app/site/hosting/restlet.nl?script=165&deploy
// =1 Instead, set OAuthUrl equal
// tohttps://rest.sandbox.netsuite.com/app/site/hosting/restlet.nl and then
// subsequently call AddParam for each query parameter.
//
func (z *OAuth1) OauthUrl() string {
    return C.GoString(C.CkOAuth1_oauthUrl(z.ckObj))
}
// property setter: OauthUrl
func (z *OAuth1) SetOauthUrl(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putOauthUrl(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OauthVersion
// The oauth_version. Defaults to "1.0". May be set to the empty string to exclude.
func (z *OAuth1) OauthVersion() string {
    return C.GoString(C.CkOAuth1_oauthVersion(z.ckObj))
}
// property setter: OauthVersion
func (z *OAuth1) SetOauthVersion(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putOauthVersion(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: QueryString
// Contains the normalized set of request parameters that are signed. This is a
// read-only property made available for debugging purposes.
func (z *OAuth1) QueryString() string {
    return C.GoString(C.CkOAuth1_queryString(z.ckObj))
}

// property: Realm
// The realm (optional).
func (z *OAuth1) Realm() string {
    return C.GoString(C.CkOAuth1_realm(z.ckObj))
}
// property setter: Realm
func (z *OAuth1) SetRealm(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putRealm(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Signature
// The generated base64 signature.
func (z *OAuth1) Signature() string {
    return C.GoString(C.CkOAuth1_signature(z.ckObj))
}

// property: SignatureMethod
// The signature method. Defaults to "HMAC-SHA1". Other possible choices are
// "HMAC1-SHA256", "RSA-SHA1", and "RSA-SHA2".
func (z *OAuth1) SignatureMethod() string {
    return C.GoString(C.CkOAuth1_signatureMethod(z.ckObj))
}
// property setter: SignatureMethod
func (z *OAuth1) SetSignatureMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putSignatureMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Timestamp
// The timestamp, such as "1441632569".
func (z *OAuth1) Timestamp() string {
    return C.GoString(C.CkOAuth1_timestamp(z.ckObj))
}
// property setter: Timestamp
func (z *OAuth1) SetTimestamp(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putTimestamp(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Token
// The token.
func (z *OAuth1) Token() string {
    return C.GoString(C.CkOAuth1_token(z.ckObj))
}
// property setter: Token
func (z *OAuth1) SetToken(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putToken(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TokenSecret
// The token secret
func (z *OAuth1) TokenSecret() string {
    return C.GoString(C.CkOAuth1_tokenSecret(z.ckObj))
}
// property setter: TokenSecret
func (z *OAuth1) SetTokenSecret(goStr string) {
    cStr := C.CString(goStr)
    C.CkOAuth1_putTokenSecret(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *OAuth1) VerboseLogging() bool {
    v := int(C.CkOAuth1_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *OAuth1) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkOAuth1_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *OAuth1) Version() string {
    return C.GoString(C.CkOAuth1_version(z.ckObj))
}
// Adds an extra name/value parameter to the OAuth1 signature.
func (z *OAuth1) AddParam(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkOAuth1_AddParam(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Generate the signature based on the property settings. Input properties are
// OauthVersion, OauthMethod, Url, ConsumerKey, ConsumerSecret, Token, TokenSecret,
// Nonce, and Timestamp. Properties set by this method include: BaseString,
// Signature, HmacKey, EncodedSignature, QueryString, GeneratedUrl,
// andAuthorizationHeader.
func (z *OAuth1) Generate() bool {

    v := C.CkOAuth1_Generate(z.ckObj)


    return v != 0
}


// Generates a random nonce numBytes in length and sets the Nonce property to the hex
// encoded value.
func (z *OAuth1) GenNonce(arg1 int) bool {

    v := C.CkOAuth1_GenNonce(z.ckObj, C.int(arg1))


    return v != 0
}


// Sets the Timestamp property to the current date/time.
func (z *OAuth1) GenTimestamp() bool {

    v := C.CkOAuth1_GenTimestamp(z.ckObj)


    return v != 0
}


// Removes a name/value parameter from the OAuth1 signature.
func (z *OAuth1) RemoveParam(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkOAuth1_RemoveParam(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the RSA key to be used when the SignatureMethod is set to "RSA-SHA1" or
// "RSA-SHA2".
func (z *OAuth1) SetRsaKey(arg1 *PrivateKey) bool {

    v := C.CkOAuth1_SetRsaKey(z.ckObj, arg1.ckObj)


    return v != 0
}


