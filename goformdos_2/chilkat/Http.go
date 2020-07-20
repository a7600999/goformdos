// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkHttp.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewHttp() *Http {
	return &Http{C.CkHttp_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Http) DisposeHttp() {
    if z != nil {
	C.CkHttp_Dispose(z.ckObj)
	}
}




func (z *Http) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkHttp_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Http) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkHttp_setExternalProgress(z.ckObj,1)
}

func (z *Http) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkHttp_setExternalProgress(z.ckObj,1)
}

func (z *Http) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkHttp_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Http) AbortCurrent() bool {
    v := int(C.CkHttp_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Http) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putAbortCurrent(z.ckObj,v)
}

// property: Accept
// The Accept header field to be automatically included with GET requests issued by
// QuickGet or QuickGetStr. The default value is "*/*".
func (z *Http) Accept() string {
    return C.GoString(C.CkHttp_accept(z.ckObj))
}
// property setter: Accept
func (z *Http) SetAccept(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putAccept(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AcceptCharset
// The Accept-Charset header field to be automatically included with GET requests
// issued by QuickGet or QuickGetStr. The default value is
// "ISO-8859-1,utf-8;q=0.7,*;q=0.7".
func (z *Http) AcceptCharset() string {
    return C.GoString(C.CkHttp_acceptCharset(z.ckObj))
}
// property setter: AcceptCharset
func (z *Http) SetAcceptCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putAcceptCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AcceptLanguage
// The Accept-Language header field to be automatically included with GET requests
// issued by QuickGet or QuickGetStr. The default value is "en-us,en;q=0.5".
func (z *Http) AcceptLanguage() string {
    return C.GoString(C.CkHttp_acceptLanguage(z.ckObj))
}
// property setter: AcceptLanguage
func (z *Http) SetAcceptLanguage(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putAcceptLanguage(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AllowGzip
// Controls whether the "Accept-Encoding: gzip" header is added to HTTP requests
// sent via any method that sends an HTTP request without using the HttpRequest
// object (such as QuickGetStr). If false, then the empty Accept-Encoding header
// is added which means the server response should contain the uncompressed
// content. The default value is true, which means the server, if it chooses, may
// send a gzipped response.
func (z *Http) AllowGzip() bool {
    v := int(C.CkHttp_getAllowGzip(z.ckObj))
    return v != 0
}
// property setter: AllowGzip
func (z *Http) SetAllowGzip(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putAllowGzip(z.ckObj,v)
}

// property: AllowHeaderFolding
// If this property is set to false, then no MIME header folding will be
// automatically applied to any request header. The default is true.
func (z *Http) AllowHeaderFolding() bool {
    v := int(C.CkHttp_getAllowHeaderFolding(z.ckObj))
    return v != 0
}
// property setter: AllowHeaderFolding
func (z *Http) SetAllowHeaderFolding(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putAllowHeaderFolding(z.ckObj,v)
}

// property: AuthToken
// If set, then automatically adds the "Authorization: Bearer " header to all
// requests. (If you have an OAuth2 access token, set this property equal to the
// OAuth2 access token string. Note: For OAuth1 (older) tokens, use the OAuthTOken
// property. )
func (z *Http) AuthToken() string {
    return C.GoString(C.CkHttp_authToken(z.ckObj))
}
// property setter: AuthToken
func (z *Http) SetAuthToken(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putAuthToken(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AutoAddHostHeader
// If set to true, the "Host" header field will automatically be added to the
// request header for any QuickGet or QuickGetStr method calls. The value of the
// Host header field is taken from the hostname part of the URL passed to
// QuickGet/QuickGetStr.
func (z *Http) AutoAddHostHeader() bool {
    v := int(C.CkHttp_getAutoAddHostHeader(z.ckObj))
    return v != 0
}
// property setter: AutoAddHostHeader
func (z *Http) SetAutoAddHostHeader(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putAutoAddHostHeader(z.ckObj,v)
}

// property: AwsAccessKey
// The AWS Access Key to be used with the Amazon S3 methods listed below.
func (z *Http) AwsAccessKey() string {
    return C.GoString(C.CkHttp_awsAccessKey(z.ckObj))
}
// property setter: AwsAccessKey
func (z *Http) SetAwsAccessKey(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putAwsAccessKey(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AwsEndpoint
// The regional endpoint (domain) to be used for Amazon S3 method calls. The
// default value is "s3.amazonaws.com". This can be set to any valid Amazon S3
// endpoint, such as "s3-eu-west-1.amazonaws.com", or the endpoints for S3-API
// compatible services from other different providers.
func (z *Http) AwsEndpoint() string {
    return C.GoString(C.CkHttp_awsEndpoint(z.ckObj))
}
// property setter: AwsEndpoint
func (z *Http) SetAwsEndpoint(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putAwsEndpoint(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AwsRegion
// The AWS (S3) region, such as "us-east-1", "us-west-2", "eu-west-1",
// "eu-central-1", etc. This propery defaults to "us-east-1". It is only used when
// the AwsSignatureVersion property is set to 4. When the AwsSignatureVersion
// property is set to 2, then this property is unused.
func (z *Http) AwsRegion() string {
    return C.GoString(C.CkHttp_awsRegion(z.ckObj))
}
// property setter: AwsRegion
func (z *Http) SetAwsRegion(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putAwsRegion(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AwsSecretKey
// The AWS Secret Key to be used with the Amazon S3 methods listed below.
func (z *Http) AwsSecretKey() string {
    return C.GoString(C.CkHttp_awsSecretKey(z.ckObj))
}
// property setter: AwsSecretKey
func (z *Http) SetAwsSecretKey(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putAwsSecretKey(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AwsSignatureVersion
// Selects the AWS Signature Version algorithm. The default value is 4. May be set
// to 2 to select AWS Signature Version 2. (The only valid choices are 2 and 4.)
func (z *Http) AwsSignatureVersion() int {
    return int(C.CkHttp_getAwsSignatureVersion(z.ckObj))
}
// property setter: AwsSignatureVersion
func (z *Http) SetAwsSignatureVersion(value int) {
    C.CkHttp_putAwsSignatureVersion(z.ckObj,C.int(value))
}

// property: AwsSubResources
// The AWS sub-resources to be used with the Amazon S3 methods listed below.
// 
// If the S3 request needs to address a sub-resource, like ?versioning, ?policy,
// ?location, ?acl, or ?torrent, or ?versionid append the sub-resource and its
// value if it has one. Note that in case of multiple sub-resources, sub-resources
// must be lexicographically sorted by sub-resource name and separated by '&'. e.g.
// "acl&versionId=value"
// 
// The list of sub-resources that can be included are: acl, location, logging,
// notification, partNumber, policy, requestPayment, torrent, uploadId, uploads,
// versionId, versioning, versions and website.
//
func (z *Http) AwsSubResources() string {
    return C.GoString(C.CkHttp_awsSubResources(z.ckObj))
}
// property setter: AwsSubResources
func (z *Http) SetAwsSubResources(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putAwsSubResources(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: BandwidthThrottleDown
// If non-zero, limits (throttles) the download bandwidth to approximately this
// maximum number of bytes per second. The default value of this property is 0.
func (z *Http) BandwidthThrottleDown() int {
    return int(C.CkHttp_getBandwidthThrottleDown(z.ckObj))
}
// property setter: BandwidthThrottleDown
func (z *Http) SetBandwidthThrottleDown(value int) {
    C.CkHttp_putBandwidthThrottleDown(z.ckObj,C.int(value))
}

// property: BandwidthThrottleUp
// If non-zero, limits (throttles) the upload bandwidth to approximately this
// maximum number of bytes per second. The default value of this property is 0.
func (z *Http) BandwidthThrottleUp() int {
    return int(C.CkHttp_getBandwidthThrottleUp(z.ckObj))
}
// property setter: BandwidthThrottleUp
func (z *Http) SetBandwidthThrottleUp(value int) {
    C.CkHttp_putBandwidthThrottleUp(z.ckObj,C.int(value))
}

// property: BasicAuth
// If HTTP basic authentication is needed, this property must be set to true. The
// HTTP protocol allows for several different types of authentication schemes, such
// as NTLM, Digest, OAuth1, etc. A given server will support (or allow) certain
// authentication schemes (also known as authentication methods). Except for the
// "Basic" authentication method, the other forms of authentication do not involve
// sending the login and password in plain unencrypted text over the connection.
// The Basic authentication method is insecure in that it sends the login/password
// for all to see. If the connection is SSL/TLS, then this might be considered OK.
// Chilkat takes the safe approach and will not allow Basic authentication unless
// this property has been explicitly set to true. The default value of this
// property is false.
// 
// Note: It is not required to know the authentication methods accepted by the
// server beforehand (except for the case of Basic authentication). When
// authentication is required, Chilkat will first send the request without the
// Authorization header, receive back the 401 Authorization Required response along
// with information about what authentication methods are accepted, and then
// re-send with an accepted authentication method. If the authentication method is
// known in advance, then an application may set the appropriate property, such as
// NtlmAuth to true so that the extra (internal) round-trip is not required.
//
func (z *Http) BasicAuth() bool {
    v := int(C.CkHttp_getBasicAuth(z.ckObj))
    return v != 0
}
// property setter: BasicAuth
func (z *Http) SetBasicAuth(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putBasicAuth(z.ckObj,v)
}

// property: ClientIpAddress
// The IP address to use for computers with multiple network interfaces or IP
// addresses. For computers with a single network interface (i.e. most computers),
// this property should not be set. For multihoming computers, the default IP
// address is automatically used if this property is not set.
// 
// The IP address is a string such as in dotted notation using numbers, not domain
// names, such as "165.164.55.124".
//
func (z *Http) ClientIpAddress() string {
    return C.GoString(C.CkHttp_clientIpAddress(z.ckObj))
}
// property setter: ClientIpAddress
func (z *Http) SetClientIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putClientIpAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConnectFailReason
// This property will be set to the status of the last HTTP connection made (or
// failed to be made) by any HTTP method.
// 
// Possible values are:
// 0 = success
// 
// Normal (non-TLS) sockets:
// 1 = empty hostname
// 2 = DNS lookup failed
// 3 = DNS timeout
// 4 = Aborted by application.
// 5 = Internal failure.
// 6 = Connect Timed Out
// 7 = Connect Rejected (or failed for some other reason)
// 50 = HTTP proxy authentication failure.
// 98 = Async operation in progress.
// 99 = Product is not unlocked.
// 
// SSL/TLS:
// 100 = TLS internal error.
// 101 = Failed to send client hello.
// 102 = Unexpected handshake message.
// 103 = Failed to read server hello.
// 104 = No server certificate.
// 105 = Unexpected TLS protocol version.
// 106 = Server certificate verify failed (the server certificate is expired or the cert's signature verification failed).
// 107 = Unacceptable TLS protocol version.
// 108 = App-defined server certificate requirements failure.
// 109 = Failed to read handshake messages.
// 110 = Failed to send client certificate handshake message.
// 111 = Failed to send client key exchange handshake message.
// 112 = Client certificate's private key not accessible.
// 113 = Failed to send client cert verify handshake message.
// 114 = Failed to send change cipher spec handshake message.
// 115 = Failed to send finished handshake message.
// 116 = Server's Finished message is invalid.
// 125 = TLS Pin Set Mismatch.
//
func (z *Http) ConnectFailReason() int {
    return int(C.CkHttp_getConnectFailReason(z.ckObj))
}

// property: Connection
// The Connection header field to be automatically included with GET requests
// issued by QuickGet or QuickGetStr. The default value is "Keep-Alive". To prevent
// the Connection header from being added to the HTTP header, set this property to
// the empty string.
func (z *Http) Connection() string {
    return C.GoString(C.CkHttp_connection(z.ckObj))
}
// property setter: Connection
func (z *Http) SetConnection(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putConnection(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConnectTimeout
// The amount of time in seconds to wait before timing out when connecting to an
// HTTP server. The default ConnectTimeout is 30 seconds.
// 
// Note: This is the maximum number of seconds to wait for a server to accept a TCP
// connection. Once the connection is accepted, and bytes begin flowing
// back-and-forth, then it is the ReadTimeout property that applies. It is the
// ReadTimeout that applies when receiving data, which includes the reads that
// occur during a TLS handshake.
//
func (z *Http) ConnectTimeout() int {
    return int(C.CkHttp_getConnectTimeout(z.ckObj))
}
// property setter: ConnectTimeout
func (z *Http) SetConnectTimeout(value int) {
    C.CkHttp_putConnectTimeout(z.ckObj,C.int(value))
}

// property: CookieDir
// Specifies a directory where cookies are automatically persisted if the
// Http.SaveCookies property is turned on. Cookies are stored in XML formatted
// files, one per domain, to main it easy for other programs to understand and
// parse. May be set to the string "memory" to cache cookies in memory.
func (z *Http) CookieDir() string {
    return C.GoString(C.CkHttp_cookieDir(z.ckObj))
}
// property setter: CookieDir
func (z *Http) SetCookieDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putCookieDir(z.ckObj,cStr)
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
func (z *Http) DebugLogFilePath() string {
    return C.GoString(C.CkHttp_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Http) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DefaultFreshPeriod
// The default freshness period (in minutes) for cached documents when the
// FreshnessAlgorithm property is set to 0. The default value is 10080 (1 week).
func (z *Http) DefaultFreshPeriod() int {
    return int(C.CkHttp_getDefaultFreshPeriod(z.ckObj))
}
// property setter: DefaultFreshPeriod
func (z *Http) SetDefaultFreshPeriod(value int) {
    C.CkHttp_putDefaultFreshPeriod(z.ckObj,C.int(value))
}

// property: DigestAuth
// Setting this property to true causes the HTTP component to use digest
// authentication. The default value is false.
func (z *Http) DigestAuth() bool {
    v := int(C.CkHttp_getDigestAuth(z.ckObj))
    return v != 0
}
// property setter: DigestAuth
func (z *Http) SetDigestAuth(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putDigestAuth(z.ckObj,v)
}

// property: FetchFromCache
// Set to true if pages should be fetched from cache when possible. Only HTTP GET
// requests are cached. HTTP responses that include Set-Cookie headers are not
// cached. A page is fetched from the disk cache if it is present and it is "fresh"
// according to the FreshnessAlgorithm property. If a page exists in cache but is
// not fresh, the HTTP component will issue a revalidate request and update the
// cache appropriately according to the response.
func (z *Http) FetchFromCache() bool {
    v := int(C.CkHttp_getFetchFromCache(z.ckObj))
    return v != 0
}
// property setter: FetchFromCache
func (z *Http) SetFetchFromCache(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putFetchFromCache(z.ckObj,v)
}

// property: FinalRedirectUrl
// If an HTTP GET was redirected (as indicated by the WasRedirected property), this
// property will contain the final redirect URL, assuming the FollowRedirects
// property is true.
// 
// Note: Starting in v9.5.0.49, this property will contain the redirect URL for
// 301/302 responses even if FollowRedirects is not set to true.
//
func (z *Http) FinalRedirectUrl() string {
    return C.GoString(C.CkHttp_finalRedirectUrl(z.ckObj))
}

// property: FollowRedirects
// If true, then 301, 302, 303, 307, and 308 redirects are automatically followed
// when calling QuickGet and QuickGetStr. FollowRedirects is true by default.
func (z *Http) FollowRedirects() bool {
    v := int(C.CkHttp_getFollowRedirects(z.ckObj))
    return v != 0
}
// property setter: FollowRedirects
func (z *Http) SetFollowRedirects(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putFollowRedirects(z.ckObj,v)
}

// property: FreshnessAlgorithm
// The freshness algorithm to use when determining the freshness of a cached HTTP
// GET response. A value of 1 causes an LM-factor algorithm to be used. This is the
// default. The LMFactor property is a value between 1 and 100 indicating the
// percentage of time based on the last-modified date of the HTML page. For
// example, if the LMFactor is 50, and an HTML page was modified 10 days ago, then
// the page will expire (i.e. no longer be fresh) in 5 days (50% of 10 days). This
// only applies to HTTP responses that do not have page expiration information. If
// the FreshnessAlgorithm = 0, then a constant expire time period determined by the
// DefaultFreshPeriod property is used.
func (z *Http) FreshnessAlgorithm() int {
    return int(C.CkHttp_getFreshnessAlgorithm(z.ckObj))
}
// property setter: FreshnessAlgorithm
func (z *Http) SetFreshnessAlgorithm(value int) {
    C.CkHttp_putFreshnessAlgorithm(z.ckObj,C.int(value))
}

// property: HeartbeatMs
// This property is only valid in programming environment and languages that allow
// for event callbacks.
// 
// Specifies the time interval in milliseconds between AbortCheck events. A value
// of 0 (the default) indicate that no AbortCheck events will fire. Any HTTP
// operation can be aborted via the AbortCheck event.
//
func (z *Http) HeartbeatMs() int {
    return int(C.CkHttp_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Http) SetHeartbeatMs(value int) {
    C.CkHttp_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: IgnoreMustRevalidate
// Some HTTP responses contain a "Cache-Control: must-revalidate" header. If this
// is present, the server is requesting that the client always issue a revalidate
// HTTP request instead of serving the page directly from cache. If
// IgnoreMustRevalidate is set to true, then Chilkat HTTP will serve the page
// directly from cache without revalidating until the page is no longer fresh.
// 
// The default value of this property is false.
//
func (z *Http) IgnoreMustRevalidate() bool {
    v := int(C.CkHttp_getIgnoreMustRevalidate(z.ckObj))
    return v != 0
}
// property setter: IgnoreMustRevalidate
func (z *Http) SetIgnoreMustRevalidate(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putIgnoreMustRevalidate(z.ckObj,v)
}

// property: IgnoreNoCache
// Some HTTP responses contain headers of various types that indicate that the page
// should not be cached. Chilkat HTTP will adhere to this unless this property is
// set to true.
// 
// The default value of this property is false.
//
func (z *Http) IgnoreNoCache() bool {
    v := int(C.CkHttp_getIgnoreNoCache(z.ckObj))
    return v != 0
}
// property setter: IgnoreNoCache
func (z *Http) SetIgnoreNoCache(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putIgnoreNoCache(z.ckObj,v)
}

// property: KeepResponseBody
// If true, then the response body, if text, is saved to the LastResponseBody
// property for all methods that do not return an HttpResponse object. The default
// value of this property is false.
func (z *Http) KeepResponseBody() bool {
    v := int(C.CkHttp_getKeepResponseBody(z.ckObj))
    return v != 0
}
// property setter: KeepResponseBody
func (z *Http) SetKeepResponseBody(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putKeepResponseBody(z.ckObj,v)
}

// property: LastContentType
// The content-type of the last HTTP response received by the HTTP component.
func (z *Http) LastContentType() string {
    return C.GoString(C.CkHttp_lastContentType(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Http) LastErrorHtml() string {
    return C.GoString(C.CkHttp_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Http) LastErrorText() string {
    return C.GoString(C.CkHttp_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Http) LastErrorXml() string {
    return C.GoString(C.CkHttp_lastErrorXml(z.ckObj))
}

// property: LastHeader
// The text of the last HTTP header sent by any of this class's methods. The
// purpose of this property is to allow the developer to examine the exact HTTP
// header for debugging purposes.
func (z *Http) LastHeader() string {
    return C.GoString(C.CkHttp_lastHeader(z.ckObj))
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
func (z *Http) LastMethodSuccess() bool {
    v := int(C.CkHttp_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Http) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putLastMethodSuccess(z.ckObj,v)
}

// property: LastModDate
// The value of the Last-Modified header in the last HTTP response received by the
// HTTP component.
func (z *Http) LastModDate() string {
    return C.GoString(C.CkHttp_lastModDate(z.ckObj))
}

// property: LastResponseBody
// The response body of the last HTTP response received by the HTTP component (for
// methods that do not return an HttpResponse object). The last response body is
// only saved to this property IF the KeepResponseBody property is set to true.
func (z *Http) LastResponseBody() string {
    return C.GoString(C.CkHttp_lastResponseBody(z.ckObj))
}

// property: LastResponseHeader
// The entire response header for the last HTTP response received by the HTTP
// component (for methods that do not return an HttpResponse object).
func (z *Http) LastResponseHeader() string {
    return C.GoString(C.CkHttp_lastResponseHeader(z.ckObj))
}

// property: LastStatus
// The last HTTP status value received by the HTTP component. This only applies to
// methods that do not return an HTTP response object. For methods that return an
// HTTP response object, such as SynchronousRequest, the status code is found in
// the StatusCode property of the response object.
func (z *Http) LastStatus() int {
    return int(C.CkHttp_getLastStatus(z.ckObj))
}

// property: LastStatusText
// The last HTTP status text received by the HTTP component. This only applies to
// methods that do not return an HTTP response object. For methods that return an
// HTTP response object, such as SynchronousRequest, the status text is found in
// the StatusText property of the response object.
func (z *Http) LastStatusText() string {
    return C.GoString(C.CkHttp_lastStatusText(z.ckObj))
}

// property: LMFactor
// An integer between 1 and 100 that indicates the percentage of time from the HTTP
// page's last-modified date that will be used for the freshness period. The
// default value is 25. For example, if a page is fetched with a last-modified date
// of 4 weeks ago, and the LMFactor = 25, then the page will be considered fresh in
// the cache for 1 week (25% of 4 weeks).
func (z *Http) LMFactor() int {
    return int(C.CkHttp_getLMFactor(z.ckObj))
}
// property setter: LMFactor
func (z *Http) SetLMFactor(value int) {
    C.CkHttp_putLMFactor(z.ckObj,C.int(value))
}

// property: Login
// The HTTP login for pages requiring a login/password. Chilkat HTTP can do Basic,
// Digest, and NTLM HTTP authentication. (NTLM is also known as SPA (or Windows
// Integrated Authentication). To use Basic authentication, the BasicAuth property
// must be set equal to true. It is not necessary to set the NtlmAuth or
// DigestAuth properties beforehand if NTLM or Digest authentication is needed.
// However, it is most efficient to pre-set these properties when the type of
// authentication is known in advance.
// 
// Important: If NTLM authentication is used, it may be incorrect to set the Login
// property equal to "mydomain\mylogin". Instead, set the LoginDomain property
// equal to "mydomain", and set this property equal to "mylogin".
//
func (z *Http) Login() string {
    return C.GoString(C.CkHttp_login(z.ckObj))
}
// property setter: Login
func (z *Http) SetLogin(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putLogin(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LoginDomain
// The optional domain name to be used with NTLM / Kerberos / Negotiate
// authentication.
func (z *Http) LoginDomain() string {
    return C.GoString(C.CkHttp_loginDomain(z.ckObj))
}
// property setter: LoginDomain
func (z *Http) SetLoginDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putLoginDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: MaxConnections
// The maximum number of simultaneous open HTTP connections managed by the HTTP
// component. The Chilkat HTTP component automatically manages HTTP connections. If
// the number of open HTTP connections is about to be exceeded, the connection with
// the least recent activity is automatically closed.
func (z *Http) MaxConnections() int {
    return int(C.CkHttp_getMaxConnections(z.ckObj))
}
// property setter: MaxConnections
func (z *Http) SetMaxConnections(value int) {
    C.CkHttp_putMaxConnections(z.ckObj,C.int(value))
}

// property: MaxFreshPeriod
// Limits the amount of time a document can be kept "fresh" in the cache. The
// MaxFreshPeriod is specified in minutes, and the default value is 525600 which is
// equal to 1 year.
func (z *Http) MaxFreshPeriod() int {
    return int(C.CkHttp_getMaxFreshPeriod(z.ckObj))
}
// property setter: MaxFreshPeriod
func (z *Http) SetMaxFreshPeriod(value int) {
    C.CkHttp_putMaxFreshPeriod(z.ckObj,C.int(value))
}

// property: MaxResponseSize
// The maximum HTTP response size to be accepted by the calling program. A value of
// 0 (the default) indicates that there is no maximum value.
func (z *Http) MaxResponseSize() uint {
    return uint(C.CkHttp_getMaxResponseSize(z.ckObj))
}
// property setter: MaxResponseSize
func (z *Http) SetMaxResponseSize(value uint) {
    C.CkHttp_putMaxResponseSize(z.ckObj,C.ulong(value))
}

// property: MaxUrlLen
// The Http class will automatically fail any URL longer than this length. The
// default MaxUrlLen is 2000 characters.
func (z *Http) MaxUrlLen() int {
    return int(C.CkHttp_getMaxUrlLen(z.ckObj))
}
// property setter: MaxUrlLen
func (z *Http) SetMaxUrlLen(value int) {
    C.CkHttp_putMaxUrlLen(z.ckObj,C.int(value))
}

// property: MimicFireFox
// If set to true, then the appropriate headers to mimic Mozilla/FireFox are
// automatically added to requests sent via the QuickGet and QuickGetStr methods.
func (z *Http) MimicFireFox() bool {
    v := int(C.CkHttp_getMimicFireFox(z.ckObj))
    return v != 0
}
// property setter: MimicFireFox
func (z *Http) SetMimicFireFox(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putMimicFireFox(z.ckObj,v)
}

// property: MimicIE
// If set to true, then the appropriate headers to mimic Internet Explorer are
// automatically added to requests sent via the QuickGet and QuickGetStr methods.
func (z *Http) MimicIE() bool {
    v := int(C.CkHttp_getMimicIE(z.ckObj))
    return v != 0
}
// property setter: MimicIE
func (z *Http) SetMimicIE(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putMimicIE(z.ckObj,v)
}

// property: MinFreshPeriod
// The freshness period for a document in cache will not be less than this value
// (in minutes). The default value is 30.
func (z *Http) MinFreshPeriod() int {
    return int(C.CkHttp_getMinFreshPeriod(z.ckObj))
}
// property setter: MinFreshPeriod
func (z *Http) SetMinFreshPeriod(value int) {
    C.CkHttp_putMinFreshPeriod(z.ckObj,C.int(value))
}

// property: NegotiateAuth
// Set this property equal to true for Negotiate authentication. Negotiate
// authentication will dynamically select Kerberos or NTLM authentication depending
// on what the server requires.
// 
// Note: The NegotiateAuth property is only available for the Microsoft Windows
// operating system.
//
func (z *Http) NegotiateAuth() bool {
    v := int(C.CkHttp_getNegotiateAuth(z.ckObj))
    return v != 0
}
// property setter: NegotiateAuth
func (z *Http) SetNegotiateAuth(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putNegotiateAuth(z.ckObj,v)
}

// property: NtlmAuth
// Setting this property to true causes the HTTP component to use NTLM
// authentication (also known as IWA -- or Integrated Windows Authentication) when
// authentication with an HTTP server. The default value is false.
func (z *Http) NtlmAuth() bool {
    v := int(C.CkHttp_getNtlmAuth(z.ckObj))
    return v != 0
}
// property setter: NtlmAuth
func (z *Http) SetNtlmAuth(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putNtlmAuth(z.ckObj,v)
}

// property: NumCacheLevels
// The number of directory levels to be used under each cache root. The default is
// 0, meaning that each cached HTML page is stored in a cache root directory. A
// value of 1 causes each cached page to be stored in one of 255 subdirectories
// named "0","1", "2", ..."255" under a cache root. A value of 2 causes two levels
// of subdirectories ("0..255/0..255") under each cache root. The HTTP control
// automatically creates subdirectories as needed. The reason for mutliple levels
// is to alleviate problems that may arise with unrelated software when huge
// numbers of files are stored in a single directory. For example, Windows Explorer
// does not behave well when trying to display the contents of directories with
// thousands of files.
func (z *Http) NumCacheLevels() int {
    return int(C.CkHttp_getNumCacheLevels(z.ckObj))
}
// property setter: NumCacheLevels
func (z *Http) SetNumCacheLevels(value int) {
    C.CkHttp_putNumCacheLevels(z.ckObj,C.int(value))
}

// property: NumCacheRoots
// The number of cache roots to be used for the HTTP cache. This allows the disk
// cache spread out over multiple disk drives. Each cache root is a string
// indicating the drive letter and directory path. For example, "E:\Cache". An
// example of a very large low-cost cache might be four USB external drives. To
// create a cache with four roots, call AddCacheRoot once for each directory root.
func (z *Http) NumCacheRoots() int {
    return int(C.CkHttp_getNumCacheRoots(z.ckObj))
}

// property: OAuth1
// If true then causes an OAuth Authorization header to be added to any request
// sent by the HTTP object. For example:
// Authorization: OAuth realm="http://sp.example.com/",
//                 oauth_consumer_key="0685bd9184jfhq22",
//                 oauth_token="ad180jjd733klru7",
//                 oauth_signature_method="HMAC-SHA1",
//                 oauth_signature="wOJIO9A2W5mFwDgiDvZbTSMK%2FPY%3D",
//                 oauth_timestamp="137131200",
//                 oauth_nonce="4572616e48616d6d65724c61686176",
//                 oauth_version="1.0"
// The information used to compute the OAuth Authorization header is obtained from
// the other OAuth* properties, such as OAuthConsumerKey, OAuthConsumerSecret,
// OAuthRealm, etc.
func (z *Http) OAuth1() bool {
    v := int(C.CkHttp_getOAuth1(z.ckObj))
    return v != 0
}
// property setter: OAuth1
func (z *Http) SetOAuth1(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putOAuth1(z.ckObj,v)
}

// property: OAuthCallback
// The OAuth 1.0 callback URL. Defaults to "oob".
func (z *Http) OAuthCallback() string {
    return C.GoString(C.CkHttp_oAuthCallback(z.ckObj))
}
// property setter: OAuthCallback
func (z *Http) SetOAuthCallback(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putOAuthCallback(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OAuthConsumerKey
// The OAuth consumer key to be used in the Authorization header.
func (z *Http) OAuthConsumerKey() string {
    return C.GoString(C.CkHttp_oAuthConsumerKey(z.ckObj))
}
// property setter: OAuthConsumerKey
func (z *Http) SetOAuthConsumerKey(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putOAuthConsumerKey(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OAuthConsumerSecret
// The OAuth consumer secret to be used in computing the contents of the
// Authorization header.
func (z *Http) OAuthConsumerSecret() string {
    return C.GoString(C.CkHttp_oAuthConsumerSecret(z.ckObj))
}
// property setter: OAuthConsumerSecret
func (z *Http) SetOAuthConsumerSecret(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putOAuthConsumerSecret(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OAuthRealm
// The OAuth realm to be used in the Authorization header.
func (z *Http) OAuthRealm() string {
    return C.GoString(C.CkHttp_oAuthRealm(z.ckObj))
}
// property setter: OAuthRealm
func (z *Http) SetOAuthRealm(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putOAuthRealm(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OAuthSigMethod
// The OAuth signature method, such as "HMAC-SHA1" to be used in the Authorization
// header. The default is "HMAC-SHA1". It is also possible to choose "HMAC-SHA256",
// "RSA-SHA1" or "RSA-SHA2". For RSA algorithms, an RSA private key would need to
// be provided via the SetOAuthRsaKey method.
// 
// Note: RSA-SHA2 is supported starting in Chilkat v9.5.0.56
//
func (z *Http) OAuthSigMethod() string {
    return C.GoString(C.CkHttp_oAuthSigMethod(z.ckObj))
}
// property setter: OAuthSigMethod
func (z *Http) SetOAuthSigMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putOAuthSigMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OAuthToken
// The OAuth1 token to be used in the Authorization header. Note: This is for
// OAuth1. Use the AuthToken property for OAuth2.
func (z *Http) OAuthToken() string {
    return C.GoString(C.CkHttp_oAuthToken(z.ckObj))
}
// property setter: OAuthToken
func (z *Http) SetOAuthToken(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putOAuthToken(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OAuthTokenSecret
// The OAuth token secret to be used in computing the Authorization header.
func (z *Http) OAuthTokenSecret() string {
    return C.GoString(C.CkHttp_oAuthTokenSecret(z.ckObj))
}
// property setter: OAuthTokenSecret
func (z *Http) SetOAuthTokenSecret(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putOAuthTokenSecret(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OAuthVerifier
// The OAuth verifier to be used in the Authorization header.
func (z *Http) OAuthVerifier() string {
    return C.GoString(C.CkHttp_oAuthVerifier(z.ckObj))
}
// property setter: OAuthVerifier
func (z *Http) SetOAuthVerifier(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putOAuthVerifier(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Password
// The HTTP password for pages requiring a login/password. Chilkat HTTP can do
// Basic, Digest, and NTLM HTTP authentication. (NTLM is also known as SPA (or
// Windows Integrated Authentication). To use Basic authentication, the BasicAuth
// property must be set equal to true. It is not necessary to set the NtlmAuth or
// DigestAuth properties beforehand if NTLM or Digest authentication is needed.
// However, it is most efficient to pre-set these properties when the type of
// authentication is known in advance.
func (z *Http) Password() string {
    return C.GoString(C.CkHttp_password(z.ckObj))
}
// property setter: Password
func (z *Http) SetPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putPassword(z.ckObj,cStr)
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
func (z *Http) PercentDoneScale() int {
    return int(C.CkHttp_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *Http) SetPercentDoneScale(value int) {
    C.CkHttp_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *Http) PreferIpv6() bool {
    v := int(C.CkHttp_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *Http) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putPreferIpv6(z.ckObj,v)
}

// property: ProxyAuthMethod
// Set this to "basic" if you know in advance that Basic authentication is to be
// used for the HTTP proxy. Otherwise leave this property unset. Note: It is not
// necessary to set this property. The HTTP component will automatically handle
// proxy authentication for any of the supported authentication methods: NTLM,
// Digest, or Basic. Setting this property equal to "basic" prevents the 407
// response which is automatically handled internal to Chilkat and never seen by
// your application.
// 
// Note: If NTLM authentication does not succeed, set the Global.DefaultNtlmVersion
// property equal to 1 and then retry.
//
func (z *Http) ProxyAuthMethod() string {
    return C.GoString(C.CkHttp_proxyAuthMethod(z.ckObj))
}
// property setter: ProxyAuthMethod
func (z *Http) SetProxyAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putProxyAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyDirectTls
// Set to true if the proxy server expects a direct TLS connection. (This is
// where the initial connection to the HTTP proxy server is TLS. SeeSquid Direct
// TLS Connection
// <https://wiki.squid-cache.org/Features/HTTPS>. The default value of this
// property is false.
func (z *Http) ProxyDirectTls() bool {
    v := int(C.CkHttp_getProxyDirectTls(z.ckObj))
    return v != 0
}
// property setter: ProxyDirectTls
func (z *Http) SetProxyDirectTls(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putProxyDirectTls(z.ckObj,v)
}

// property: ProxyDomain
// The domain name of a proxy host if an HTTP proxy is used. This can also be set
// to an IP address.
func (z *Http) ProxyDomain() string {
    return C.GoString(C.CkHttp_proxyDomain(z.ckObj))
}
// property setter: ProxyDomain
func (z *Http) SetProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyLogin
// If an HTTP proxy is used and it requires authentication, this property specifies
// the HTTP proxy login.
func (z *Http) ProxyLogin() string {
    return C.GoString(C.CkHttp_proxyLogin(z.ckObj))
}
// property setter: ProxyLogin
func (z *Http) SetProxyLogin(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putProxyLogin(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyLoginDomain
// The NTLM authentication domain (optional) if NTLM authentication is used.
func (z *Http) ProxyLoginDomain() string {
    return C.GoString(C.CkHttp_proxyLoginDomain(z.ckObj))
}
// property setter: ProxyLoginDomain
func (z *Http) SetProxyLoginDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putProxyLoginDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyPassword
// If an HTTP proxy is used and it requires authentication, this property specifies
// the HTTP proxy password.
func (z *Http) ProxyPassword() string {
    return C.GoString(C.CkHttp_proxyPassword(z.ckObj))
}
// property setter: ProxyPassword
func (z *Http) SetProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyPort
// The port number of a proxy server if an HTTP proxy is used.
func (z *Http) ProxyPort() int {
    return int(C.CkHttp_getProxyPort(z.ckObj))
}
// property setter: ProxyPort
func (z *Http) SetProxyPort(value int) {
    C.CkHttp_putProxyPort(z.ckObj,C.int(value))
}

// property: ReadTimeout
// The amount of time in seconds to wait before timing out when reading from an
// HTTP server. The ReadTimeout is the amount of time that needs to elapse while no
// additional data is forthcoming. During a long download, if the data stream halts
// for more than this amount, it will timeout. Otherwise, there is no limit on the
// length of time for the entire download.
// 
// The default value is 60 seconds. Note: Prior to v9.5.0.76, the default was 20
// seconds.
//
func (z *Http) ReadTimeout() int {
    return int(C.CkHttp_getReadTimeout(z.ckObj))
}
// property setter: ReadTimeout
func (z *Http) SetReadTimeout(value int) {
    C.CkHttp_putReadTimeout(z.ckObj,C.int(value))
}

// property: RedirectVerb
// Indicates the HTTP verb, such as GET, POST, PUT, etc. to be used for a redirect
// when the FollowRedirects property is set to true. The default value of this
// property is "GET". This will produce the same behavior as a web browser (such as
// FireFox). If this property is set to the empty string, then it will cause the
// same verb as the original HTTP request to be used.
// 
// Note: Prior to version 9.5.0.44, the default value of this property was the
// empty string.
//
func (z *Http) RedirectVerb() string {
    return C.GoString(C.CkHttp_redirectVerb(z.ckObj))
}
// property setter: RedirectVerb
func (z *Http) SetRedirectVerb(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putRedirectVerb(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Referer
// The Referer header field to be automatically included with GET requests issued
// by QuickGet or QuickGetStr. The default value is the empty string which causes
// the Referer field to be omitted from the request header.
func (z *Http) Referer() string {
    return C.GoString(C.CkHttp_referer(z.ckObj))
}
// property setter: Referer
func (z *Http) SetReferer(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putReferer(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: RequiredContentType
// If set, then any HTTP response to any POST or GET, including downloads, will be
// rejected if the content-type in the response header does not match this setting.
// If the content-type does not match, only the header of the HTTP response is
// read, the connection to the HTTP server is closed, and the remainder of the
// response is never read.
// 
// This property is empty (zero-length string) by default.
// 
// Some typical content-types are "text/html", "text/xml", "image/gif",
// "image/jpeg", "application/zip", "application/msword", "application/pdf", etc.
//
func (z *Http) RequiredContentType() string {
    return C.GoString(C.CkHttp_requiredContentType(z.ckObj))
}
// property setter: RequiredContentType
func (z *Http) SetRequiredContentType(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putRequiredContentType(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: RequireSslCertVerify
// If true, then the HTTP client will verify the server's SSL certificate. The
// certificate is expired, or if the cert's signature is invalid, the connection is
// not allowed. The default value of this property is false.
func (z *Http) RequireSslCertVerify() bool {
    v := int(C.CkHttp_getRequireSslCertVerify(z.ckObj))
    return v != 0
}
// property setter: RequireSslCertVerify
func (z *Http) SetRequireSslCertVerify(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putRequireSslCertVerify(z.ckObj,v)
}

// property: S3Ssl
// If true, then all S3_* methods will use a secure SSL/TLS connection for
// communications. (If true, Chilkat uses TLS 1.2) The default value is true.
func (z *Http) S3Ssl() bool {
    v := int(C.CkHttp_getS3Ssl(z.ckObj))
    return v != 0
}
// property setter: S3Ssl
func (z *Http) SetS3Ssl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putS3Ssl(z.ckObj,v)
}

// property: SaveCookies
// If this property is true, cookies are automatically persisted to XML files in
// the directory specified by the CookiesDir property (or in memory if CookieDir =
// "memory"). Both CookiesDir and SaveCookies must be set for cookies to be
// persisted.
func (z *Http) SaveCookies() bool {
    v := int(C.CkHttp_getSaveCookies(z.ckObj))
    return v != 0
}
// property setter: SaveCookies
func (z *Http) SetSaveCookies(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putSaveCookies(z.ckObj,v)
}

// property: SendBufferSize
// The buffer size to be used with the underlying TCP/IP socket for sending. The
// default value is 65535.
func (z *Http) SendBufferSize() int {
    return int(C.CkHttp_getSendBufferSize(z.ckObj))
}
// property setter: SendBufferSize
func (z *Http) SetSendBufferSize(value int) {
    C.CkHttp_putSendBufferSize(z.ckObj,C.int(value))
}

// property: SendCookies
// If true, then cookies previously persisted to the CookiesDir are automatically
// added to all HTTP requests. Only cookies matching the domain and path are added.
func (z *Http) SendCookies() bool {
    v := int(C.CkHttp_getSendCookies(z.ckObj))
    return v != 0
}
// property setter: SendCookies
func (z *Http) SetSendCookies(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putSendCookies(z.ckObj,v)
}

// property: SessionLogFilename
// Enables file-based session logging. If set to a filename (or relative/absolute
// filepath), then the exact HTTP requests and responses are logged to a file. The
// file is created if it does not already exist, otherwise it is appended.
func (z *Http) SessionLogFilename() string {
    return C.GoString(C.CkHttp_sessionLogFilename(z.ckObj))
}
// property setter: SessionLogFilename
func (z *Http) SetSessionLogFilename(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putSessionLogFilename(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SniHostname
// Specifies the SNI hostname to be used in the TLS ClientHello. This property is
// only needed when the domain is specified via a dotted IP address and an SNI
// hostname is desired. (Normally, Chilkat automatically uses the domain name in
// the SNI hostname TLS ClientHello extension.)
func (z *Http) SniHostname() string {
    return C.GoString(C.CkHttp_sniHostname(z.ckObj))
}
// property setter: SniHostname
func (z *Http) SetSniHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putSniHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksHostname
// The SOCKS4/SOCKS5 hostname or IPv4 address (in dotted decimal notation). This
// property is only used if the SocksVersion property is set to 4 or 5).
func (z *Http) SocksHostname() string {
    return C.GoString(C.CkHttp_socksHostname(z.ckObj))
}
// property setter: SocksHostname
func (z *Http) SetSocksHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putSocksHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPassword
// The SOCKS5 password (if required). The SOCKS4 protocol does not include the use
// of a password, so this does not apply to SOCKS4.
func (z *Http) SocksPassword() string {
    return C.GoString(C.CkHttp_socksPassword(z.ckObj))
}
// property setter: SocksPassword
func (z *Http) SetSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPort
// The SOCKS4/SOCKS5 proxy port. The default value is 1080. This property only
// applies if a SOCKS proxy is used (if the SocksVersion property is set to 4 or
// 5).
func (z *Http) SocksPort() int {
    return int(C.CkHttp_getSocksPort(z.ckObj))
}
// property setter: SocksPort
func (z *Http) SetSocksPort(value int) {
    C.CkHttp_putSocksPort(z.ckObj,C.int(value))
}

// property: SocksUsername
// The SOCKS4/SOCKS5 proxy username. This property is only used if the SocksVersion
// property is set to 4 or 5).
func (z *Http) SocksUsername() string {
    return C.GoString(C.CkHttp_socksUsername(z.ckObj))
}
// property setter: SocksUsername
func (z *Http) SetSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksVersion
// SocksVersion May be set to one of the following integer values:
// 
// 0 - No SOCKS proxy is used. This is the default.
// 4 - Connect via a SOCKS4 proxy.
// 5 - Connect via a SOCKS5 proxy.
//
func (z *Http) SocksVersion() int {
    return int(C.CkHttp_getSocksVersion(z.ckObj))
}
// property setter: SocksVersion
func (z *Http) SetSocksVersion(value int) {
    C.CkHttp_putSocksVersion(z.ckObj,C.int(value))
}

// property: SoRcvBuf
// Sets the receive buffer size socket option. Normally, this property should be
// left unchanged. The default value is 4194304.
// 
// This property can be increased if download performance seems slow. It is
// recommended to be a multiple of 4096.
//
func (z *Http) SoRcvBuf() int {
    return int(C.CkHttp_getSoRcvBuf(z.ckObj))
}
// property setter: SoRcvBuf
func (z *Http) SetSoRcvBuf(value int) {
    C.CkHttp_putSoRcvBuf(z.ckObj,C.int(value))
}

// property: SoSndBuf
// Sets the send buffer size socket option. Normally, this property should be left
// unchanged. The default value is 262144.
// 
// This property can be increased if upload performance seems slow. It is
// recommended to be a multiple of 4096. Testing with sizes such as 512K and 1MB is
// reasonable.
//
func (z *Http) SoSndBuf() int {
    return int(C.CkHttp_getSoSndBuf(z.ckObj))
}
// property setter: SoSndBuf
func (z *Http) SetSoSndBuf(value int) {
    C.CkHttp_putSoSndBuf(z.ckObj,C.int(value))
}

// property: SslAllowedCiphers
// Provides a means for setting a list of ciphers that are allowed for SSL/TLS
// connections. The default (empty string) indicates that all implemented ciphers
// are possible. The TLS ciphers supported in Chilkat v9.5.0.55 and later are:
// TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
// TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
// TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256
// TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA
// TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256
// TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA
// TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384
// TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384
// TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
// TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA
// TLS_DHE_RSA_WITH_AES_256_CBC_SHA256
// TLS_DHE_RSA_WITH_AES_256_GCM_SHA384
// TLS_DHE_RSA_WITH_AES_256_CBC_SHA
// TLS_RSA_WITH_AES_256_CBC_SHA256
// TLS_RSA_WITH_AES_256_GCM_SHA384
// TLS_RSA_WITH_AES_256_CBC_SHA
// TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256
// TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
// TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA
// TLS_DHE_RSA_WITH_AES_128_CBC_SHA256
// TLS_DHE_RSA_WITH_AES_128_GCM_SHA256
// TLS_DHE_RSA_WITH_AES_128_CBC_SHA
// TLS_RSA_WITH_AES_128_CBC_SHA256
// TLS_RSA_WITH_AES_128_GCM_SHA256
// TLS_RSA_WITH_AES_128_CBC_SHA
// TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA
// TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA
// TLS_RSA_WITH_3DES_EDE_CBC_SHA
// TLS_ECDHE_RSA_WITH_RC4_128_SHA
// TLS_RSA_WITH_RC4_128_SHA
// TLS_RSA_WITH_RC4_128_MD5
// TLS_DHE_RSA_WITH_DES_CBC_SHA
// TLS_RSA_WITH_DES_CBC_SHA
// To restrict SSL/TLS connections to one or more specific ciphers, set this
// property to a comma-separated list of ciphers such as
// "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384".
// The order should be in terms of preference, with the preferred algorithms listed
// first. (Note that the client cannot specifically choose the algorithm is picked
// because it is the server that chooses. The client simply provides the server
// with a list from which to choose.)
// 
// The property can also disallow connections with servers having certificates with
// RSA keys less than a certain size. By default, server certificates having RSA
// keys of 512 bits or greater are allowed. Add the keyword "rsa1024" to disallow
// connections with servers having keys smaller than 1024 bits. Add the keyword
// "rsa2048" to disallow connections with servers having keys smaller than 2048
// bits.
// 
// Note: Prior to Chilkat v9.5.0.55, it was not possible to explicitly list allowed
// cipher suites. The deprecated means for indicating allowed ciphers was both
// incomplete and unprecise. For example, the following keywords could be listed to
// allow matching ciphers: "aes256-cbc", "aes128-cbc", "3des-cbc", and "rc4". These
// keywords will still be recognized, but programs should be updated to explicitly
// list the allowed ciphers.
// 
// secure-renegotiation: Starting in Chilkat v9.5.0.55, the keyword
// "secure-renegotiation" may be added to require that all renegotions be done
// securely (as per RFC 5746).
// 
// best-practices: Starting in Chilkat v9.5.0.55, this property may be set to the
// single keyword "best-practices". This will allow ciphers based on the current
// best practices. As new versions of Chilkat are released, the best practices may
// change. Changes will be noted here. The current best practices are:
// 
//     If the server uses an RSA key, it must be 1024 bits or greater.
//     All renegotations must be secure renegotiations.
//     All ciphers using RC4, DES, or 3DES are disallowed.
// 
// Example: The following string would restrict to 2 specific cipher suites,
// require RSA keys to be 1024 bits or greater, and require secure renegotiations:
// "TLS_DHE_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA, rsa1024,
// secure-renegotiation"
//
func (z *Http) SslAllowedCiphers() string {
    return C.GoString(C.CkHttp_sslAllowedCiphers(z.ckObj))
}
// property setter: SslAllowedCiphers
func (z *Http) SetSslAllowedCiphers(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putSslAllowedCiphers(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SslProtocol
// Selects the secure protocol to be used for secure (SSL/TLS) connections.
// Possible values are:
// 
//     default
//     TLS 1.2
//     TLS 1.1
//     TLS 1.0
//     SSL 3.0
//     TLS 1.2 or higher
//     TLS 1.1 or higher
//     TLS 1.0 or higher
//     
// 
// The default value is "default" which will choose the, which allows for the
// protocol to be selected dynamically at runtime based on the requirements of the
// server. Choosing an exact protocol will cause the connection to fail unless that
// exact protocol is negotiated. It is better to choose "X or higher" than an exact
// protocol. The "default" is effectively "SSL 3.0 or higher".
func (z *Http) SslProtocol() string {
    return C.GoString(C.CkHttp_sslProtocol(z.ckObj))
}
// property setter: SslProtocol
func (z *Http) SetSslProtocol(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putSslProtocol(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: StreamResponseBodyPath
// Allows for the HTTP response body to be streamed directly into a file. If this
// property is set, then any method returning an HTTP response object will stream
// the response body directly to the file path specified. The HTTP response object
// will still contain the response header. (This property is useful when the HTTP
// response is too large to fit into memory.)
func (z *Http) StreamResponseBodyPath() string {
    return C.GoString(C.CkHttp_streamResponseBodyPath(z.ckObj))
}
// property setter: StreamResponseBodyPath
func (z *Http) SetStreamResponseBodyPath(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putStreamResponseBodyPath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TlsCipherSuite
// Contains the current or last negotiated TLS cipher suite. If no TLS connection
// has yet to be established, or if a connection as attempted and failed, then this
// will be empty. A sample cipher suite string looks like this:
// TLS_DHE_RSA_WITH_AES_256_CBC_SHA256.
func (z *Http) TlsCipherSuite() string {
    return C.GoString(C.CkHttp_tlsCipherSuite(z.ckObj))
}

// property: TlsPinSet
// Specifies a set of pins for Public Key Pinning for TLS connections. This
// property lists the expected SPKI fingerprints for the server certificates. If
// the server's certificate (sent during the TLS handshake) does not match any of
// the SPKI fingerprints, then the TLS handshake is aborted and the connection
// fails. The format of this string property is as follows:
// hash_algorithm, encoding, SPKI_fingerprint_1, SPKI_fingerprint_2, ...
// For example, the following string specifies a single sha256 base64-encoded SPKI
// fingerprint:
// "sha256, base64, lKg1SIqyhPSK19tlPbjl8s02yChsVTDklQpkMCHvsTE="
// This example specifies two SPKI fingerprints:
// "sha256, base64, 4t37LpnGmrMEAG8HEz9yIrnvJV2euVRwCLb9EH5WZyI=, 68b0G5iqMvWVWvUCjMuhLEyekM5729PadtnU5tdXZKs="
// Any of the following hash algorithms are allowed:.sha1, sha256, sha384, sha512,
// md2, md5, haval, ripemd128, ripemd160,ripemd256, or ripemd320.
// 
// The following encodings are allowed: base64, hex, and any of the encodings
// indicated in the link below.
//
func (z *Http) TlsPinSet() string {
    return C.GoString(C.CkHttp_tlsPinSet(z.ckObj))
}
// property setter: TlsPinSet
func (z *Http) SetTlsPinSet(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putTlsPinSet(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TlsVersion
// Contains the current or last negotiated TLS protocol version. If no TLS
// connection has yet to be established, or if a connection as attempted and
// failed, then this will be empty. Possible values are "SSL 3.0", "TLS 1.0", "TLS
// 1.1", and "TLS 1.2".
func (z *Http) TlsVersion() string {
    return C.GoString(C.CkHttp_tlsVersion(z.ckObj))
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. This property
// defaults to the empty string and should typically remain empty. Can be set to a
// list of the following comma separated keywords:
//     "QuickDisconnect" - Introduced in v9.5.0.77. In the call to
//     CloseAllConnections, do not disconnect cleanly. Instead just disconnect as
//     quickly as possible.
//     "ProtectFromVpn" - Introduced in v9.5.0.80. On Android systems, will bypass
//     any VPN that may be installed or active.
//     "TlsNoClientRootCert" - Introduced in v9.5.0.82. Will exclude root CA certs
//     from being included in the client certificate chain that is sent to the server
//     for client-side authentication. This must be set prior to calling
//     SetSslClientCert.
//     "AllowEmptyHeaders" - Introduced in v9.5.0.82. If present, an empty value
//     string passed to SetHeaderField will cause the header to be added with an empty
//     value. Otherwise, for historical purposes and backward compatibility, the header
//     field is removed when an empty value string is passed.
//     "EnableTls13" - Introduced in v9.5.0.82. Causes TLS 1.3 to be offered in the
//     ClientHello of the TLS protocol. The server may then select TLS 1.3 for the
//     session. Future versions of Chilkat will enable TLS 1.3 by default. This option
//     is only necessary in v9.5.0.82 if TLS 1.3 is desired.
func (z *Http) UncommonOptions() string {
    return C.GoString(C.CkHttp_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Http) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UpdateCache
// Controls whether the cache is automatically updated with the responses from HTTP
// GET requests.
func (z *Http) UpdateCache() bool {
    v := int(C.CkHttp_getUpdateCache(z.ckObj))
    return v != 0
}
// property setter: UpdateCache
func (z *Http) SetUpdateCache(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putUpdateCache(z.ckObj,v)
}

// property: UseIEProxy
// If true, the proxy host/port used by Internet Explorer will also be used by
// Chilkat HTTP.
func (z *Http) UseIEProxy() bool {
    v := int(C.CkHttp_getUseIEProxy(z.ckObj))
    return v != 0
}
// property setter: UseIEProxy
func (z *Http) SetUseIEProxy(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putUseIEProxy(z.ckObj,v)
}

// property: UserAgent
// The UserAgent header field to be automatically included with GET requests issued
// by QuickGet or QuickGetStr. The default value is "Mozilla/5.0 (Windows NT 6.3;
// WOW64; rv:49.0) Gecko/20100101 Firefox/49.0". The reason for this default is to
// get the same server behavior for a recent version of a typical and popular
// browser. Some sites may respond differently depending on the User-Agent, and the
// goal is to avoid strange responses that are different than what a typical
// browser would receive.
func (z *Http) UserAgent() string {
    return C.GoString(C.CkHttp_userAgent(z.ckObj))
}
// property setter: UserAgent
func (z *Http) SetUserAgent(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttp_putUserAgent(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Http) VerboseLogging() bool {
    v := int(C.CkHttp_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Http) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttp_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Http) Version() string {
    return C.GoString(C.CkHttp_version(z.ckObj))
}

// property: WasRedirected
// Indicates whether the last HTTP GET was redirected.
func (z *Http) WasRedirected() bool {
    v := int(C.CkHttp_getWasRedirected(z.ckObj))
    return v != 0
}
// This method must be called at least once if disk caching is to be used. The file
// path (including drive letter) such as "E:\MyHttpCache\" is passed to
// AddCacheRoot to specify the root directory. The cache can be spread across
// multiple disk drives by calling AddCacheRoot multiple times, each with a
// directory path on a separate disk drive.
func (z *Http) AddCacheRoot(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkHttp_AddCacheRoot(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Removes all headers added via the SetRequestHeader method.
func (z *Http) ClearHeaders()  {

    C.CkHttp_ClearHeaders(z.ckObj)



}


// Clears all cookies cached in memory. Calling this only makes sense if the
// CookieDir property is set to the string "memory".
func (z *Http) ClearInMemoryCookies()  {

    C.CkHttp_ClearInMemoryCookies(z.ckObj)



}


// Clears all URL variable values previously set by one or more calls to SetUrlVar.
func (z *Http) ClearUrlVars()  {

    C.CkHttp_ClearUrlVars(z.ckObj)



}


// Closes all connections still open from previous HTTP requests.
// 
// An HTTP object instance will maintain up to 10 connections. If the HTTP server's
// response does not include a "Connection: Close" header, the connection will
// remain open and will be re-used if possible for the next HTTP request to the
// same hostname:port. (It uses the IP address (in string form) or the domain name,
// whichever is used in the URL provided by the application.) If 10 connections are
// already open and another is needed, the object will close the least recently
// used connection.
//
func (z *Http) CloseAllConnections() bool {

    v := C.CkHttp_CloseAllConnections(z.ckObj)


    return v != 0
}


// Asynchronous version of the CloseAllConnections method
func (z *Http) CloseAllConnectionsAsync(c chan *Task) {

    hTask := C.CkHttp_CloseAllConnectionsAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Creates an OCSP request for one or more certificates. The requestDetails is a JSON
// document describing the content of the OCSP request to be created. The examples
// in the links below show how to build the JSON request details.
// 
// Note: This method only creates an OCSP request. After creating, it may be sent
// to a server to get the OCSP response.
//
func (z *Http) CreateOcspRequest(arg1 *JsonObject, arg2 *BinData) bool {

    v := C.CkHttp_CreateOcspRequest(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Creates an RFC 3161 time-stamp request and returns the binary request token in
// tmestampToken. The hashAlg can be "sha1", "sha256", "sha384", "sha512", or "md5", The hashVal
// is the base64 hash of the data to be timestamped. The optional reqPolicyOid is the
// requested policy OID in a format such as "1.3.6.1.4.1.47272.1.2". The addNonce
// indicates whether to auto-generate and include a nonce in the request. It may be
// true or false. The reqTsaCert determines whether or not to request the TSA's
// certificate (true = Yes, false = No).
// 
// Note: This method only creates a timestamp request. After creating, it may be
// sent to a server to get the binary timestamp token.
//
func (z *Http) CreateTimestampRequest(arg1 string, arg2 string, arg3 string, arg4 bool, arg5 bool, arg6 *BinData) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }

    v := C.CkHttp_CreateTimestampRequest(z.ckObj, cstr1, cstr2, cstr3, b4, b5, arg6.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Clears the Chilkat-wide in-memory hostname-to-IP address DNS cache. Chilkat
// automatically maintains this in-memory cache to prevent redundant DNS lookups.
// If the TTL on the DNS A records being accessed are short and/or these DNS
// records change frequently, then this method can be called clear the internal
// cache. Note: The DNS cache is used/shared among all Chilkat objects in a
// program, and clearing the cache affects all Chilkat objects.
func (z *Http) DnsCacheClear()  {

    C.CkHttp_DnsCacheClear(z.ckObj)



}


// Retrieves the content at a URL and saves to a file. All content is saved in
// streaming mode such that the memory footprint is small and steady. HTTPS is
// fully supported, as it is with all the methods of this class.
func (z *Http) Download(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_Download(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the Download method
func (z *Http) DownloadAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_DownloadAsync(z.ckObj, cstr1, cstr2)

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


// Same as the Download method, but the output file is open for append.
func (z *Http) DownloadAppend(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_DownloadAppend(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the DownloadAppend method
func (z *Http) DownloadAppendAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_DownloadAppendAsync(z.ckObj, cstr1, cstr2)

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


// Downloads the content at the url into a BinData object.
func (z *Http) DownloadBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttp_DownloadBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the DownloadBd method
func (z *Http) DownloadBdAsync(arg1 string, arg2 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_DownloadBdAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Retrieves the content at a URL and computes and returns a hash of the content.
// The hash is returned as an encoded string according to the encoding, which may be
// "Base64", "modBase64", "Base32", "UU", "QP" (for quoted-printable), "URL" (for
// url-encoding), "Hex", "Q", "B", "url_oath", "url_rfc1738", "url_rfc2396", and
// "url_rfc3986". The hashAlgorithm may be "sha1", "sha256", "sha384", "sha512", "md2",
// "md5", "haval", "ripemd128", "ripemd160","ripemd256", or "ripemd320".
// return a string or nil if failed.
func (z *Http) DownloadHash(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkHttp_downloadHash(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the DownloadHash method
func (z *Http) DownloadHashAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkHttp_DownloadHashAsync(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads the content at the url into a Chilkat StringBuilder object. The charset
// tells Chilkat how to interpret the bytes received. The sb is appended with the
// downloaded text data.
func (z *Http) DownloadSb(arg1 string, arg2 string, arg3 *StringBuilder) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_DownloadSb(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the DownloadSb method
func (z *Http) DownloadSbAsync(arg1 string, arg2 string, arg3 *StringBuilder, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_DownloadSbAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// Convenience method for extracting the META refresh URL from HTML. For example,
// if the htmlContent contains a META refresh tag, such as:
// _LT_meta http-equiv="refresh" content="5;URL='http://example.com/'"_GT_
// Then the return value of this method would be "http://example.com/".
// return a string or nil if failed.
func (z *Http) ExtractMetaRefreshUrl(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_extractMetaRefreshUrl(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Makes an access token request to obtain a Google API OAuth2 access token for a
// service account. Access tokens issued by the Google OAuth 2.0 Authorization
// Server expire one hour after they are issued. When an access token expires, then
// the application should generate another JWT, sign it, and request another access
// token. The iss is the service account email address of the application making
// the access token request. The scope is a space-delimited list of the permissions
// that the application requests. (See
// https://developers.google.com/accounts/docs/OAuth2ServiceAccount )
// 
// The subEmail is the email address of the user for which the application is
// requesting delegated access. The subEmail may be left empty if there is no such
// email address.
//
// return a string or nil if failed.
func (z *Http) G_SvcOauthAccessToken(arg1 string, arg2 string, arg3 string, arg4 int, arg5 *Cert) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkHttp_g_SvcOauthAccessToken(z.ckObj, cstr1, cstr2, cstr3, C.int(arg4), arg5.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the G_SvcOauthAccessToken method
func (z *Http) G_SvcOauthAccessTokenAsync(arg1 string, arg2 string, arg3 string, arg4 int, arg5 *Cert, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkHttp_G_SvcOauthAccessTokenAsync(z.ckObj, cstr1, cstr2, cstr3, C.int(arg4), arg5.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as the G_SvcOauthAccessToken method, but with added flexibility for
// more customization. The 1st three args of the G_SvcOauthAccessToken are replaced
// with claimParams allowing for future expansion of name-value params. See the example
// below.
// return a string or nil if failed.
func (z *Http) G_SvcOauthAccessToken2(arg1 *Hashtable, arg2 int, arg3 *Cert) *string {

    retStr := C.CkHttp_g_SvcOauthAccessToken2(z.ckObj, arg1.ckObj, C.int(arg2), arg3.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the G_SvcOauthAccessToken2 method
func (z *Http) G_SvcOauthAccessToken2Async(arg1 *Hashtable, arg2 int, arg3 *Cert, c chan *Task) {

    hTask := C.CkHttp_G_SvcOauthAccessToken2Async(z.ckObj, arg1.ckObj, C.int(arg2), arg3.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the current GMT (also known as UTC) date/time in a string that is
// compliant with RFC 2616 format.
// return a string or nil if failed.
func (z *Http) GenTimeStamp() *string {

    retStr := C.CkHttp_genTimeStamp(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth cache root (indexing begins at 0). Cache roots are set by
// calling AddCacheRoot one or more times.
// return a string or nil if failed.
func (z *Http) GetCacheRoot(arg1 int) *string {

    retStr := C.CkHttp_getCacheRoot(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the cookies in XML format for a specific domain. Cookies are only
// persisted if the SaveCookies property is set to true. If the CookieDir
// property is set to the keyword "memory", then cookies are saved in-memory.
// return a string or nil if failed.
func (z *Http) GetCookieXml(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_getCookieXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Utility method for extracting the domain name from a full URL. For example, if
// "http://www.chilkatsoft.com/default.asp" is the URL passed in, then
// "www.chilkatsoft.com" is returned.
// return a string or nil if failed.
func (z *Http) GetDomain(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_getDomain(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sends an HTTP HEAD request for a URL and returns a response object. (Note: HEAD
// requests will never automatically follow redirects.)
func (z *Http) GetHead(arg1 string) *HttpResponse {
    cstr1 := C.CString(arg1)

    retObj := C.CkHttp_GetHead(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the GetHead method
func (z *Http) GetHeadAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_GetHeadAsync(z.ckObj, cstr1)

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


// Returns the value of a header field that has been pre-defined to be sent with
// all HTTP GET requests issued by the QuickGet and QuickGetStr methods. By
// default, this includes header fields such as Accept, AcceptCharset,
// AcceptLanguage, Connection, UserAgent, etc.
// return a string or nil if failed.
func (z *Http) GetRequestHeader(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_getRequestHeader(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Establishes an SSL/TLS connection with a web server for the purpose of
// retrieving the server's SSL certificate (public-key only of course...). Nothing
// is retrieved from the web server. This method simply makes a connection, gets
// the certificate information, and closes the connection.
func (z *Http) GetServerSslCert(arg1 string, arg2 int) *Cert {
    cstr1 := C.CString(arg1)

    retObj := C.CkHttp_GetServerSslCert(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Asynchronous version of the GetServerSslCert method
func (z *Http) GetServerSslCertAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_GetServerSslCertAsync(z.ckObj, cstr1, C.int(arg2))

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


// Returns the path part of a URL. The syntax of a URL is
// _LT_scheme>://_LT_user>:_LT_password>@_LT_host>:_LT_port>/_LT_path>;_LT_params>?_
// LT_query>#_LT_frag>. This method returns the "path" part.
// return a string or nil if failed.
func (z *Http) GetUrlPath(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_getUrlPath(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if the specified header field is defined such that it will be sent
// with all GET requests issued by the QuickGet and QuickGetStr methods.
func (z *Http) HasRequestHeader(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttp_HasRequestHeader(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if the Http class has been unlocked. It is necessary to call
// Http.UnlockComponent before calling any other methods. Passing any string to
// UnlockComponent will automatically activate a 30-day trial period.
func (z *Http) IsUnlocked() bool {

    v := C.CkHttp_IsUnlocked(z.ckObj)


    return v != 0
}


// Provides information about what transpired in the last method called on this
// object instance. For many methods, there is no information. However, for some
// methods, details about what occurred can be obtained by getting the LastJsonData
// right after the method call returns.
func (z *Http) LastJsonData() *JsonObject {

    retObj := C.CkHttp_LastJsonData(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Loads the caller of the task's async method.
func (z *Http) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkHttp_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Parses an OCSP reply. Returns the following possible integer values:
//     -1: The ocspReply does not contain a valid OCSP reply.
//     0: Successful - Response has valid confirmations..
//     1: Malformed request - Illegal confirmation request.
//     2: Internal error - Internal error in issuer.
//     3: Try later - Try again later.
//     4: Not used - This value is never returned.
//     5: Sig required - Must sign the request.
//     6: Unauthorized - Request unauthorized.
// 
// The binaryOCSP reply is provided in ocspReply. The replyData is populated with data parsed
// from ocspReply.
//
func (z *Http) ParseOcspReply(arg1 *BinData, arg2 *JsonObject) int {

    v := C.CkHttp_ParseOcspReply(z.ckObj, arg1.ckObj, arg2.ckObj)


    return int(v)
}


// Sends an HTTP request to the url. The verb can be "POST", "PUT", "PATCH", etc.
// The body of the HTTP request contains the bytes passed in byteData. The contentType is a
// content type such as "image/gif", "application/pdf", etc. If md5 is true,
// then a Content-MD5 header is added with the base64 MD5 hash of the byteData. Servers
// aware of the Content-MD5 header will perform a message integrity check to ensure
// that the data has not been corrupted. If gzip is true, the byteData is compressed
// using the gzip algorithm. The HTTP request body will contain the GZIP compressed
// data, and a "Content-Encoding: gzip" header is automatically added to indicate
// that the request data needs to be ungzipped when received (at the server).
func (z *Http) PBinary(arg1 string, arg2 string, arg3 []byte, arg4 string, arg5 bool, arg6 bool) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))
    cstr4 := C.CString(arg4)
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    retObj := C.CkHttp_PBinary(z.ckObj, cstr1, cstr2, ckbyd3, cstr4, b5, b6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)
    C.free(unsafe.Pointer(cstr4))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PBinary method
func (z *Http) PBinaryAsync(arg1 string, arg2 string, arg3 []byte, arg4 string, arg5 bool, arg6 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))
    cstr4 := C.CString(arg4)
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    hTask := C.CkHttp_PBinaryAsync(z.ckObj, cstr1, cstr2, ckbyd3, cstr4, b5, b6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)
    C.free(unsafe.Pointer(cstr4))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as PBinary, but the data to be uploaded is passed in data.
func (z *Http) PBinaryBd(arg1 string, arg2 string, arg3 *BinData, arg4 string, arg5 bool, arg6 bool) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr4 := C.CString(arg4)
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    retObj := C.CkHttp_PBinaryBd(z.ckObj, cstr1, cstr2, arg3.ckObj, cstr4, b5, b6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr4))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PBinaryBd method
func (z *Http) PBinaryBdAsync(arg1 string, arg2 string, arg3 *BinData, arg4 string, arg5 bool, arg6 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr4 := C.CString(arg4)
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    hTask := C.CkHttp_PBinaryBdAsync(z.ckObj, cstr1, cstr2, arg3.ckObj, cstr4, b5, b6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr4))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an HTTP/HTTPS request to the url. The verb can be "POST", "PUT", "PATCH",
// etc. The url can begin with "http://" or "https://" depending if TLS is
// desired. The body of the HTTP request is streamed directly from the localFilePath. The
// contentType is a content type such as "image/gif", "application/pdf", "text/xml", etc.
// If md5 is true, then a Content-MD5 header is added with the base64 MD5 hash
// of the localFilePath. Servers aware of the Content-MD5 header will perform a message
// integrity check to ensure that the data has not been corrupted. If gzip is
// true, the localFilePath is compressed using the gzip algorithm. The HTTP request body
// will contain the GZIP compressed data, and a "Content-Encoding: gzip" header is
// automatically added to indicate that the request data needs to be ungzipped when
// received (at the server).
func (z *Http) PFile(arg1 string, arg2 string, arg3 string, arg4 string, arg5 bool, arg6 bool) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    retObj := C.CkHttp_PFile(z.ckObj, cstr1, cstr2, cstr3, cstr4, b5, b6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PFile method
func (z *Http) PFileAsync(arg1 string, arg2 string, arg3 string, arg4 string, arg5 bool, arg6 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    hTask := C.CkHttp_PFileAsync(z.ckObj, cstr1, cstr2, cstr3, cstr4, b5, b6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an HTTP POST request to the url. The body of the HTTP request contains
// the bytes passed in byteData. The contentType is a content type such as "image/gif",
// "application/pdf", etc. If md5 is true, then a Content-MD5 header is added
// with the base64 MD5 hash of the byteData. Servers aware of the Content-MD5 header
// will perform a message integrity check to ensure that the data has not been
// corrupted. If gzip is true, the byteData is compressed using the gzip algorithm.
// The HTTP request body will contain the GZIP compressed data, and a
// "Content-Encoding: gzip" header is automatically added to indicate that the
// request data needs to be ungzipped when received (at the server).
// 
// Returns the text body of the HTTP response if the HTTP response has a success
// status code. Otherwise the method is considered to have failed. If more details
// of the HTTP response are required, call PBinary instead (which returns the HTTP
// response object).
// 
// Note: The HTTP response code is available in the LastStatus property. Other
// properties having information include LastResponseHeader, LastResponseBody,
// LastModDate, LastContentType, etc.
//
// return a string or nil if failed.
func (z *Http) PostBinary(arg1 string, arg2 []byte, arg3 string, arg4 bool, arg5 bool) *string {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }

    retStr := C.CkHttp_postBinary(z.ckObj, cstr1, ckbyd2, cstr3, b4, b5)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the PostBinary method
func (z *Http) PostBinaryAsync(arg1 string, arg2 []byte, arg3 string, arg4 bool, arg5 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }

    hTask := C.CkHttp_PostBinaryAsync(z.ckObj, cstr1, ckbyd2, cstr3, b4, b5)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// A simplified way of sending a JSON POST and receiving the JSON response. The
// HTTP response is returned in an HTTP response object. The content type of the
// HTTP request is "application/json". To send a JSON POST using a different
// content-type, call the PostJson2 method where the content type can be explicitly
// specified.
func (z *Http) PostJson(arg1 string, arg2 string) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkHttp_PostJson(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PostJson method
func (z *Http) PostJsonAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_PostJsonAsync(z.ckObj, cstr1, cstr2)

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


// The same as PostJson,except it allows for the content type to be explicitly
// provided. The PostJson method automatically uses "application/jsonrequest". If
// the application needs for the content type to be "application/json", or some
// other content type, then PostJson2 provides the means.
func (z *Http) PostJson2(arg1 string, arg2 string, arg3 string) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkHttp_PostJson2(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PostJson2 method
func (z *Http) PostJson2Async(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkHttp_PostJson2Async(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as PostJson2,except a JSON object is passed in for the request body.
func (z *Http) PostJson3(arg1 string, arg2 string, arg3 *JsonObject) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkHttp_PostJson3(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PostJson3 method
func (z *Http) PostJson3Async(arg1 string, arg2 string, arg3 *JsonObject, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_PostJson3Async(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// Sends a simple URL encoded POST. The form parameters are sent in the body of the
// HTTP request in x-www-form-urlencoded format. The content-type is
// "application/x-www-form-urlencoded".
func (z *Http) PostUrlEncoded(arg1 string, arg2 *HttpRequest) *HttpResponse {
    cstr1 := C.CString(arg1)

    retObj := C.CkHttp_PostUrlEncoded(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PostUrlEncoded method
func (z *Http) PostUrlEncodedAsync(arg1 string, arg2 *HttpRequest, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_PostUrlEncodedAsync(z.ckObj, cstr1, arg2.ckObj)

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


// A simplified way of posting XML content to a web server. This method is good for
// making SOAP calls using HTTP POST. The xmlCharset should match the character encoding
// used in the xmlContent, which is typically "utf-8". The HTTP response is returned in
// an HTTP response object.
// 
// Important: This method sends the POST with a "Content-Type" header value of
// "application/xml". In rare cases, a server might require the Content-Type header
// to be "text/xml". To use "text/xml" instead of the default "application/xml",
// call SetRequestHeader("Content-Type","text/xml") prior to calling this method.
// 
// To use HTTPS simply pass an endpointUrl beginning with "https://" instead of "http://".
// This applies to any Chilkat method where a URL is passed as an argument.
//
func (z *Http) PostXml(arg1 string, arg2 string, arg3 string) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkHttp_PostXml(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PostXml method
func (z *Http) PostXmlAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkHttp_PostXmlAsync(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an HTTP request to the url. The verb can be "POST", "PUT", "PATCH", etc.
// The body of the HTTP request contains the text passed in textData. The contentType is a
// content type such as "text/xml", "application/json", etc. If md5 is true,
// then a Content-MD5 header is added with the base64 MD5 hash of the textData. Servers
// aware of the Content-MD5 header will perform a message integrity check to ensure
// that the data has not been corrupted. If gzip is true, the contentType is compressed
// using the gzip algorithm. The HTTP request body will contain the GZIP compressed
// data, and a "Content-Encoding: gzip" header is automatically added to indicate
// that the request data needs to be ungzipped when received (at the server).
func (z *Http) PText(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 bool, arg7 bool) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }
    b7 := C.int(0)
    if arg7 {
        b7 = C.int(1)
    }

    retObj := C.CkHttp_PText(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5, b6, b7)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PText method
func (z *Http) PTextAsync(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 bool, arg7 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }
    b7 := C.int(0)
    if arg7 {
        b7 = C.int(1)
    }

    hTask := C.CkHttp_PTextAsync(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5, b6, b7)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as PText, but the data to be uploaded is passed in textData.
func (z *Http) PTextSb(arg1 string, arg2 string, arg3 *StringBuilder, arg4 string, arg5 string, arg6 bool, arg7 bool) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }
    b7 := C.int(0)
    if arg7 {
        b7 = C.int(1)
    }

    retObj := C.CkHttp_PTextSb(z.ckObj, cstr1, cstr2, arg3.ckObj, cstr4, cstr5, b6, b7)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the PTextSb method
func (z *Http) PTextSbAsync(arg1 string, arg2 string, arg3 *StringBuilder, arg4 string, arg5 string, arg6 bool, arg7 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }
    b7 := C.int(0)
    if arg7 {
        b7 = C.int(1)
    }

    hTask := C.CkHttp_PTextSbAsync(z.ckObj, cstr1, cstr2, arg3.ckObj, cstr4, cstr5, b6, b7)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an HTTP PUT request to the url. The body of the HTTP request is byteData. The
// contentType is a content type such as "image/gif", "application/pdf", etc. If md5 is
// true, then a Content-MD5 header is added with the base64 MD5 hash of the byteData.
// Servers aware of the Content-MD5 header will perform a message integrity check
// to ensure that the data has not been corrupted. If gzip is true, the byteData is
// compressed using the gzip algorithm. The HTTP request body will contain the GZIP
// compressed data, and a "Content-Encoding: gzip" header is automatically added to
// indicate that the request data needs to be ungzipped when received (at the
// server).
// 
// Returns the text body of the HTTP response if the HTTP response has a success
// status code. Otherwise the method is considered to have failed. If more details
// of the HTTP response are required, call PBinary instead (which returns the HTTP
// response object).
//
// return a string or nil if failed.
func (z *Http) PutBinary(arg1 string, arg2 []byte, arg3 string, arg4 bool, arg5 bool) *string {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }

    retStr := C.CkHttp_putBinary(z.ckObj, cstr1, ckbyd2, cstr3, b4, b5)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the PutBinary method
func (z *Http) PutBinaryAsync(arg1 string, arg2 []byte, arg3 string, arg4 bool, arg5 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }

    hTask := C.CkHttp_PutBinaryAsync(z.ckObj, cstr1, ckbyd2, cstr3, b4, b5)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an HTTP PUT request to the url. The body of the HTTP request is textData. The
// charset should be set to a charset name such as "iso-8859-1", "windows-1252",
// "Shift_JIS", "utf-8", etc. The string "ansi" may also be used as a charset name.
// The contentType is a content type such as "text/plain", "text/xml", etc. If md5 is
// true, then a Content-MD5 header is added with the base64 MD5 hash of the textData.
// Servers aware of the Content-MD5 header will perform a message integrity check
// to ensure that the data has not been corrupted. If gzip is true, the textData is
// compressed using the gzip algorithm. The HTTP request body will contain the GZIP
// compressed data, and a "Content-Encoding: gzip" header is automatically added to
// indicate that the request data needs to be ungzipped when received (at the
// server).
// 
// Returns the text body of the HTTP response if the HTTP response has a success
// status code. Otherwise the method is considered to have failed. If more details
// of the HTTP response are required, call PText instead (which returns the HTTP
// response object).
//
// return a string or nil if failed.
func (z *Http) PutText(arg1 string, arg2 string, arg3 string, arg4 string, arg5 bool, arg6 bool) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    retStr := C.CkHttp_putText(z.ckObj, cstr1, cstr2, cstr3, cstr4, b5, b6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the PutText method
func (z *Http) PutTextAsync(arg1 string, arg2 string, arg3 string, arg4 string, arg5 bool, arg6 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }
    b6 := C.int(0)
    if arg6 {
        b6 = C.int(1)
    }

    hTask := C.CkHttp_PutTextAsync(z.ckObj, cstr1, cstr2, cstr3, cstr4, b5, b6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Same as QuickGetStr, but uses the HTTP DELETE method instead of the GET method.
// 
// Note: The HTTP response code is available in the LastStatus property. Other
// properties having information include LastResponseHeader, LastResponseBody,
// LastModDate, LastContentType, etc.
//
// return a string or nil if failed.
func (z *Http) QuickDeleteStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_quickDeleteStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the QuickDeleteStr method
func (z *Http) QuickDeleteStrAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_QuickDeleteStrAsync(z.ckObj, cstr1)

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


// Sends an HTTP GET request for a URL and returns the response body as a byte
// array. The URL may contain query parameters. If the SendCookies property is
// true, matching cookies previously persisted to the CookiesDir are automatically
// included in the request. If the FetchFromCache property is true, the page may be
// fetched directly from cache. Because the URL can specify any type of resource
// (HTML page, GIF image, etc.) the return value is a byte array. If the resource
// is known to be a string, such as with an HTML page, you may call QuickGetStr
// instead. If the HTTP request fails, a zero-length byte array is returned and
// error information can be found in the LastErrorText, LastErrorXml, or
// LastErrorHtml properties.
// 
// Note: The HTTP response code is available in the LastStatus property. Other
// properties having information include LastResponseHeader, LastResponseBody,
// LastModDate, LastContentType, etc.
//
func (z *Http) QuickGet(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkHttp_QuickGet(z.ckObj, cstr1, ckOutBytes)

    C.free(unsafe.Pointer(cstr1))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the QuickGet method
func (z *Http) QuickGetAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_QuickGetAsync(z.ckObj, cstr1)

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


// The same as QuickGet, but returns the content in a Chilkat BinData object. The
// existing content of binData, if any, is cleared and replaced with the downloaded
// content.
func (z *Http) QuickGetBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttp_QuickGetBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the QuickGetBd method
func (z *Http) QuickGetBdAsync(arg1 string, arg2 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_QuickGetBdAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Sends an HTTP GET request for a URL and returns the response object. If the
// SendCookies property is true, matching cookies previously persisted to the
// CookiesDir are automatically included in the request. If the FetchFromCache
// property is true, the page could be fetched directly from cache.
func (z *Http) QuickGetObj(arg1 string) *HttpResponse {
    cstr1 := C.CString(arg1)

    retObj := C.CkHttp_QuickGetObj(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the QuickGetObj method
func (z *Http) QuickGetObjAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_QuickGetObjAsync(z.ckObj, cstr1)

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


// The same as QuickGetStr, but returns the content in a Chilkat StringBuilder
// object. The existing content of sbContent, if any, is cleared and replaced with the
// downloaded content.
func (z *Http) QuickGetSb(arg1 string, arg2 *StringBuilder) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttp_QuickGetSb(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the QuickGetSb method
func (z *Http) QuickGetSbAsync(arg1 string, arg2 *StringBuilder, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_QuickGetSbAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Sends an HTTP GET request for a URL and returns the response body as a string.
// The URL may contain query parameters. If the SendCookies property is true,
// matching cookies previously persisted to the CookiesDir are automatically
// included in the request. If the FetchFromCache property is true, the page
// could be fetched directly from cache. If the HTTP request fails, a _NULL_ value
// is returned and error information can be found in the LastErrorText,
// LastErrorXml, or LastErrorHtml properties.
// 
// Note: The HTTP response code is available in the LastStatus property. Other
// properties having information include LastResponseHeader, LastResponseBody,
// LastModDate, LastContentType, etc.
//
// return a string or nil if failed.
func (z *Http) QuickGetStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_quickGetStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the QuickGetStr method
func (z *Http) QuickGetStrAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_QuickGetStrAsync(z.ckObj, cstr1)

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


// Same as QuickGetStr, but uses the HTTP PUT method instead of the GET method.
// 
// Note: The HTTP response code is available in the LastStatus property. Other
// properties having information include LastResponseHeader, LastResponseBody,
// LastModDate, LastContentType, etc.
//
// return a string or nil if failed.
func (z *Http) QuickPutStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_quickPutStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the QuickPutStr method
func (z *Http) QuickPutStrAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_QuickPutStrAsync(z.ckObj, cstr1)

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


// Sends an HTTP GET request for a URL and returns the response object. If the
// SendCookies property is true, matching cookies previously persisted to the
// CookiesDir are automatically included in the request. If the FetchFromCache
// property is true, the page could be fetched directly from cache.
func (z *Http) QuickRequest(arg1 string, arg2 string) *HttpResponse {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkHttp_QuickRequest(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the QuickRequest method
func (z *Http) QuickRequestAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_QuickRequestAsync(z.ckObj, cstr1, cstr2)

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


// Removes a header from the internal list of custom header field name/value pairs
// to be automatically added when HTTP requests are sent via methods that do not
// use the HTTP request object. (The SetRequestHeader method is called to add
// custom header fields.)
func (z *Http) RemoveRequestHeader(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkHttp_RemoveRequestHeader(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Same as QuickGet, but does not send the HTTP GET. Instead, it builds the HTTP
// request that would've been sent and returns it.
// return a string or nil if failed.
func (z *Http) RenderGet(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_renderGet(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Same as the Download method, except a failed download may be resumed. The targetFilename
// is automatically checked and if it exists, the download will resume at the point
// where it previously failed. ResumeDownload may be called any number of times
// until the full download is complete.
func (z *Http) ResumeDownload(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_ResumeDownload(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the ResumeDownload method
func (z *Http) ResumeDownloadAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_ResumeDownloadAsync(z.ckObj, cstr1, cstr2)

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


// Same as the DownloadBd method, except a failed download may be resumed. The
// download will resume at a point based on the number of bytes already contained
// in binData. ResumeDownloadBd may be called any number of times until the full
// download is complete.
// 
// The incoming data is appended to binData.
//
func (z *Http) ResumeDownloadBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttp_ResumeDownloadBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the ResumeDownloadBd method
func (z *Http) ResumeDownloadBdAsync(arg1 string, arg2 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_ResumeDownloadBdAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Creates a new Amazon S3 bucket.
// 
// Note: x-amz-* headers, including metadata, can be added to any S3 request by
// adding each header with a call to SetRequestHeader. This applies to all S3
// methods, even if not explicitly stated.
//
func (z *Http) S3_CreateBucket(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttp_S3_CreateBucket(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the S3_CreateBucket method
func (z *Http) S3_CreateBucketAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_S3_CreateBucketAsync(z.ckObj, cstr1)

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


// Deletes an Amazon S3 bucket.
// 
// Note: If the bucket is in a region different than us-east-1, makes sure to set
// the AwsEndpoint property to the correct region, such as "eu-central-1". Also, if
// using an S3 compatible service such as Wasabi, always set the AwsEndpoint
// property. For example: "s3.wasabisys.com", "s3.eu-central-1.wasabisys.com".
//
func (z *Http) S3_DeleteBucket(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttp_S3_DeleteBucket(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the S3_DeleteBucket method
func (z *Http) S3_DeleteBucketAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_S3_DeleteBucketAsync(z.ckObj, cstr1)

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


// Deletes multiple objects from a bucket using a single HTTP request. The bucketName
// contains the names (also known as "keys") of the objects to be deleted. To
// delete a specific version of an object, append a versionId attribute to the
// object name. For example: "SampleDocument.txt;
// VersionId="OYcLXagmS.WaD..oyH4KRguB95_YhLs7""
func (z *Http) S3_DeleteMultipleObjects(arg1 string, arg2 *StringArray) *HttpResponse {
    cstr1 := C.CString(arg1)

    retObj := C.CkHttp_S3_DeleteMultipleObjects(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the S3_DeleteMultipleObjects method
func (z *Http) S3_DeleteMultipleObjectsAsync(arg1 string, arg2 *StringArray, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_S3_DeleteMultipleObjectsAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Deletes a remote file (object) on the Amazon S3 service.
func (z *Http) S3_DeleteObject(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_S3_DeleteObject(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the S3_DeleteObject method
func (z *Http) S3_DeleteObjectAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_S3_DeleteObjectAsync(z.ckObj, cstr1, cstr2)

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


// The same as DownloadFile, except the downloaded file data is appended to bd.
// 
// Note: If the bucket is in a region different than us-east-1, makes sure to set
// the AwsEndpoint property to the correct region, such as "eu-central-1". Also, if
// using an S3 compatible service such as Wasabi, always set the AwsEndpoint
// property. For example: "s3.wasabisys.com", "s3.eu-central-1.wasabisys.com".
//
func (z *Http) S3_DownloadBd(arg1 string, arg2 string, arg3 *BinData) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_S3_DownloadBd(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the S3_DownloadBd method
func (z *Http) S3_DownloadBdAsync(arg1 string, arg2 string, arg3 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_S3_DownloadBdAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// The same as DownloadFile, except the file data is returned directly in-memory
// instead of being written to a local file.
func (z *Http) S3_DownloadBytes(arg1 string, arg2 string) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkHttp_S3_DownloadBytes(z.ckObj, cstr1, cstr2, ckOutBytes)

    C.free(unsafe.Pointer(cstr1))
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


// Asynchronous version of the S3_DownloadBytes method
func (z *Http) S3_DownloadBytesAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_S3_DownloadBytesAsync(z.ckObj, cstr1, cstr2)

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


// Downloads a file from the Amazon S3 service.
// 
// Note: If the bucket is in a region different than us-east-1, makes sure to set
// the AwsEndpoint property to the correct region, such as "eu-central-1". Also, if
// using an S3 compatible service such as Wasabi, always set the AwsEndpoint
// property. For example: "s3.wasabisys.com", "s3.eu-central-1.wasabisys.com".
//
func (z *Http) S3_DownloadFile(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkHttp_S3_DownloadFile(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the S3_DownloadFile method
func (z *Http) S3_DownloadFileAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkHttp_S3_DownloadFileAsync(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads a text file (object) from the Amazon S3 service directly into a string
// variable. The charset specifies the character encoding, such as "utf-8", of the
// remote text object.
// 
// Note: If the bucket is in a region different than us-east-1, makes sure to set
// the AwsEndpoint property to the correct region, such as "eu-central-1". Also, if
// using an S3 compatible service such as Wasabi, always set the AwsEndpoint
// property. For example: "s3.wasabisys.com", "s3.eu-central-1.wasabisys.com".
//
// return a string or nil if failed.
func (z *Http) S3_DownloadString(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkHttp_s3_DownloadString(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the S3_DownloadString method
func (z *Http) S3_DownloadStringAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkHttp_S3_DownloadStringAsync(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Determines if a remote object (file) exists. Returns 1 if the file exists, 0 if
// it does not exist, -1 if there was a failure in checking, or 2 if using in
// asynchronous mode to indicate that the background task was successfully started.
// 
// Note: If the bucket is in a region different than us-east-1, makes sure to set
// the AwsEndpoint property to the correct region, such as "eu-central-1". Also, if
// using an S3 compatible service such as Wasabi, always set the AwsEndpoint
// property. For example: "s3.wasabisys.com", "s3.eu-central-1.wasabisys.com".
//
func (z *Http) S3_FileExists(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_S3_FileExists(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Asynchronous version of the S3_FileExists method
func (z *Http) S3_FileExistsAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_S3_FileExistsAsync(z.ckObj, cstr1, cstr2)

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


// Generates a temporary pre-signed URL for Amazon S3 using AWS Signature V2. (Call
// S3_GenerateUrlV4 to generate AWS Signature V4 pre-signed URLs.) Requires that
// the AwsSecretKey and AwsAccessKey be set to valid values prior to calling this
// method.
// 
// Note: This method can only generate URLs that are for HTTP GET requests (i.e.
// URLs you can paste into a browser address bar). This method does not generate
// URLs for POST, PUT, DELETE, etc.
//
// return a string or nil if failed.
func (z *Http) S3_GenerateUrl(arg1 string, arg2 string, arg3 *CkDateTime) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkHttp_s3_GenerateUrl(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Note: The S3_GenPresignedUrl method is new in Chilkat v9.5.0.83 and can create
// S3 pre-signed URLs for GET, POST, PUT, DELETE, or any other HTTP verb.
// 
// Generates a temporary pre-signed URL for Amazon S3 using AWS Signature V4. (Call
// S3_GenerateUrl to generate AWS Signature V2 pre-signed URLs.) Requires that the
// AwsSecretKey, AwsAccessKey, and AwsRegion properties be set to valid values
// prior to calling this method. Also requires the AwsEndpoint property to be set
// if the endpoint is different than "s3.amazonaws.com".
// 
// The URL that is generated has this format:
// https:////?X-Amz-Algorithm=AWS4-HMAC-SHA256
// &X-Amz-Credential=////aws4_request
// &X-Amz-Date=&X-Amz-Expires=&X-Amz-SignedHeaders=host
// &X-Amz-Signature=
// 
// The awsService is a string naming the AWS service, such as "s3".   If useHttps is true, then the URL begins with "https://", otherwise it begins with "http://".
// 
// Note: This method can only generate URLs that are for HTTP GET requests (i.e. URLs you can paste into a browser address bar).  This method does not generate URLs for POST, PUT, DELETE, etc.
//
// return a string or nil if failed.
func (z *Http) S3_GenerateUrlV4(arg1 bool, arg2 string, arg3 string, arg4 int, arg5 string) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr5 := C.CString(arg5)

    retStr := C.CkHttp_s3_GenerateUrlV4(z.ckObj, b1, cstr2, cstr3, C.int(arg4), cstr5)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr5))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Generates a temporary pre-signed URL for Amazon S3 using AWS Signature V4.
// Requires that the AwsSecretKey, AwsAccessKey, and AwsRegion properties be set to
// valid values prior to calling this method. Also requires the AwsEndpoint
// property to be set if the endpoint is different than "s3.amazonaws.com".
// 
// The URL that is generated has this format:
// https:////?X-Amz-Algorithm=AWS4-HMAC-SHA256
// &X-Amz-Credential=////aws4_request
// &X-Amz-Date=&X-Amz-Expires=&X-Amz-SignedHeaders=host
// &X-Amz-Signature=
// 
// The httpVerb is the HTTP verb such as "GET", "PUT", "POST", "DELETE", etc.  The awsService is a string naming the AWS service, such as "s3" or "s3-accelerate".   If useHttps is true, then the URL begins with "https://", otherwise it begins with "http://".
//
// return a string or nil if failed.
func (z *Http) S3_GenPresignedUrl(arg1 string, arg2 bool, arg3 string, arg4 string, arg5 int, arg6 string) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr6 := C.CString(arg6)

    retStr := C.CkHttp_s3_GenPresignedUrl(z.ckObj, cstr1, b2, cstr3, cstr4, C.int(arg5), cstr6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr6))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Retrieves the XML listing of the objects contained within an Amazon S3 bucket.
// (This is like a directory listing, but in XML format.)
// 
// The bucketPath name may be qualified with URL-encoded params. For example, to list the
// objects in a bucket named "ChilkatABC" with max-keys = 2000 and marker = "xyz",
// call S3_ListBucketObject passing the following string for bucketPath:
// "ChilkatABC?max-keys=2000&marker=xyz"
// 
// The S3_ListBucketObjects method recognized all params listed in the AWS
// documentation for listing objects in a bucket: delimiter, marker, max-keys, and
// prefix. See Amazon's AWS online documentation for more information.
//
// return a string or nil if failed.
func (z *Http) S3_ListBucketObjects(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_s3_ListBucketObjects(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the S3_ListBucketObjects method
func (z *Http) S3_ListBucketObjectsAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkHttp_S3_ListBucketObjectsAsync(z.ckObj, cstr1)

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


// Retrieves the XML listing of the buckets for an Amazon S3 account.
// return a string or nil if failed.
func (z *Http) S3_ListBuckets() *string {

    retStr := C.CkHttp_s3_ListBuckets(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the S3_ListBuckets method
func (z *Http) S3_ListBucketsAsync(c chan *Task) {

    hTask := C.CkHttp_S3_ListBucketsAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as S3_UploadFile, except the contents of the file come from bd
// instead of a local file.
// 
// Note: x-amz-* headers, including metadata, can be added to any S3 request by
// adding each header with a call to SetRequestHeader. This applies to all S3
// methods, even if not explicitly stated.
// 
// Note: If the bucket is in a region different than us-east-1, makes sure to set
// the AwsEndpoint property to the correct region, such as "eu-central-1". Also, if
// using an S3 compatible service such as Wasabi, always set the AwsEndpoint
// property. For example: "s3.wasabisys.com", "s3.eu-central-1.wasabisys.com".
//
func (z *Http) S3_UploadBd(arg1 *BinData, arg2 string, arg3 string, arg4 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkHttp_S3_UploadBd(z.ckObj, arg1.ckObj, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Asynchronous version of the S3_UploadBd method
func (z *Http) S3_UploadBdAsync(arg1 *BinData, arg2 string, arg3 string, arg4 string, c chan *Task) {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    hTask := C.CkHttp_S3_UploadBdAsync(z.ckObj, arg1.ckObj, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as S3_UploadFile, except the contents of the file come from contentBytes
// instead of a local file.
// 
// Note: x-amz-* headers, including metadata, can be added to any S3 request by
// adding each header with a call to SetRequestHeader. This applies to all S3
// methods, even if not explicitly stated.
//
func (z *Http) S3_UploadBytes(arg1 []byte, arg2 string, arg3 string, arg4 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkHttp_S3_UploadBytes(z.ckObj, ckbyd1, cstr2, cstr3, cstr4)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Asynchronous version of the S3_UploadBytes method
func (z *Http) S3_UploadBytesAsync(arg1 []byte, arg2 string, arg3 string, arg4 string, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    hTask := C.CkHttp_S3_UploadBytesAsync(z.ckObj, ckbyd1, cstr2, cstr3, cstr4)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Uploads a file to the Amazon S3 service.
// 
// Note: x-amz-* headers, including metadata, can be added to any S3 request by
// adding each header with a call to SetRequestHeader. This applies to all S3
// methods, even if not explicitly stated.
// 
// Note: If the bucket is in a region different than us-east-1, makes sure to set
// the AwsEndpoint property to the correct region, such as "eu-central-1". Also, if
// using an S3 compatible service such as Wasabi, always set the AwsEndpoint
// property. For example: "s3.wasabisys.com", "s3.eu-central-1.wasabisys.com".
//
func (z *Http) S3_UploadFile(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkHttp_S3_UploadFile(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Asynchronous version of the S3_UploadFile method
func (z *Http) S3_UploadFileAsync(arg1 string, arg2 string, arg3 string, arg4 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    hTask := C.CkHttp_S3_UploadFileAsync(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Uploads an in-memory string to the Amazon S3 service. This is the same as
// UploadFile, except that the file contents are from an in-memory string instead
// of a local file. Internal to this method, the objectContent is converted to the character
// encoding specified by charset prior to uploading.
// 
// Note: x-amz-* headers, including metadata, can be added to any S3 request by
// adding each header with a call to SetRequestHeader. This applies to all S3
// methods, even if not explicitly stated.
// 
// Note: If the bucket is in a region different than us-east-1, makes sure to set
// the AwsEndpoint property to the correct region, such as "eu-central-1". Also, if
// using an S3 compatible service such as Wasabi, always set the AwsEndpoint
// property. For example: "s3.wasabisys.com", "s3.eu-central-1.wasabisys.com".
//
func (z *Http) S3_UploadString(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkHttp_S3_UploadString(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Asynchronous version of the S3_UploadString method
func (z *Http) S3_UploadStringAsync(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    hTask := C.CkHttp_S3_UploadStringAsync(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Restores cookies for a particular domain. It is assumed that the cookie XML was
// previously retrieved via the GetCookieXml method, and saved to some sort of
// persistent storage, such as within a database table. It is then possible for an
// application to restore the cookies by calling this method.
func (z *Http) SetCookieXml(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_SetCookieXml(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the RSA key to be used with OAuth authentication when the RSA-SHA1 OAuth
// signature method is used (see the OAuthSigMethod property).
func (z *Http) SetOAuthRsaKey(arg1 *PrivateKey) bool {

    v := C.CkHttp_SetOAuthRsaKey(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds a custom header field to any HTTP request sent by a method that does not
// use the HTTP request object. These methods include Download, DownloadAppend,
// GetHead, PostBinary, PostXml, PutBinary, PutText, QuickDeleteStr, QuickGet,
// QuickGetObj, QuickGetStr, QuickPutStr, XmlRpc, and XmlRpcPut.
// 
// Cookies may be explictly added by calling this method passing "Cookie" for the
// headerFieldName.
// 
// The RemoveRequestHeader method can be called to remove a custom header.
// 
// Note 1: Never explicitly set the Content-Length header field. Chilkat will
// automatically compute the correct length and add the Content-Length header to
// all POST, PUT, or any other request where the Content-Length needs to be
// specified. (GET requests always have a 0 length body, and therefore never need a
// Content-Length header field.)
// 
// Note: To add more than one header, call this method once per header field.
// 
// Note 2:: Passing an empty value string causes the header to be removed. This was
// the unfortunate behavior of the method from the beginning, and cannot be changed
// for backward compatibility reasons. A workaround exists starting in v9.5.0.82 by
// adding the "AllowEmptyHeaders" keyword to the UncommonOptions property.
//
func (z *Http) SetRequestHeader(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkHttp_SetRequestHeader(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Equivalent to setting the Password property, but provides for a more secure way
// of passing the password in a secure string object.
func (z *Http) SetSecurePassword(arg1 *SecureString) bool {

    v := C.CkHttp_SetSecurePassword(z.ckObj, arg1.ckObj)


    return v != 0
}


// Allows for a client-side certificate to be used for an SSL connection.
func (z *Http) SetSslClientCert(arg1 *Cert) bool {

    v := C.CkHttp_SetSslClientCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Allows for a client-side certificate + private key to be used for the SSL / TLS
// connection (often called 2-way SSL).
func (z *Http) SetSslClientCertPem(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_SetSslClientCertPem(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Allows for a client-side certificate + private key to be used for the SSL / TLS
// connection (often called 2-way SSL).
func (z *Http) SetSslClientCertPfx(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_SetSslClientCertPfx(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the value of a variable for substitutions in URLs passed to any method.
// Variables can appear in URLs in the following format: {$varName}. For example:
// https://graph.microsoft.com/v1.0/users/{$id}
func (z *Http) SetUrlVar(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_SetUrlVar(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Authenticates with SharePoint Online, resulting in a cookie being set and used
// for subsequent SharePoint HTTP requests. Prior to calling this method, an
// application should set the CookieDir property to either "memory" or a directory
// path to persist the SharePoint authentication cookie to be automatically used in
// subsequent runs.
// 
// This method has the side-effect of setting the SaveCookies and SendCookies
// properties to true, because these settings are required for SharePoint Online
// authentication.
// 
// The siteUrl is a URL such as "https://yourdomain.sharepoint.com/". The username is an
// email address such as "username@yourdomain.com". The extraInfo is reserved for
// providing additional information as needed in the future.
//
func (z *Http) SharePointOnlineAuth(arg1 string, arg2 string, arg3 *SecureString, arg4 *JsonObject) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttp_SharePointOnlineAuth(z.ckObj, cstr1, cstr2, arg3.ckObj, arg4.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SharePointOnlineAuth method
func (z *Http) SharePointOnlineAuthAsync(arg1 string, arg2 string, arg3 *SecureString, arg4 *JsonObject, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_SharePointOnlineAuthAsync(z.ckObj, cstr1, cstr2, arg3.ckObj, arg4.ckObj)

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


// Convenience method to force the calling process to sleep for a number of
// milliseconds.
func (z *Http) SleepMs(arg1 int)  {

    C.CkHttp_SleepMs(z.ckObj, C.int(arg1))



}


// Sends an explicit HttpRequest to an HTTP server and returns an HttpResponse
// object. The HttpResponse object provides full access to the response including
// all headers and the response body. This method may be used to send POST
// requests, as well as GET, HEAD, file uploads, and XMLHTTP. To send via HTTPS
// (i.e. TLS), set the ssl property = true. Otherwise set it to false.
// 
// NOTE: The 1st argument of this method is a domain, not a URL. For example, DO
// NOT pass "https://www.somedomain.com/" in the 1st argument. Instead, pass
// "www.somedomain.com".
// 
// The Parts of a URL
// 
// http://example.com:8042/over/there?name=ferret#nose
// \__/   \______________/\_________/ \________/ \__/
//  |           |            |            |        |
// scheme   domain+port     path        query   fragment
func (z *Http) SynchronousRequest(arg1 string, arg2 int, arg3 bool, arg4 *HttpRequest) *HttpResponse {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retObj := C.CkHttp_SynchronousRequest(z.ckObj, cstr1, C.int(arg2), b3, arg4.ckObj)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &HttpResponse{retObj}
}


// Asynchronous version of the SynchronousRequest method
func (z *Http) SynchronousRequestAsync(arg1 string, arg2 int, arg3 bool, arg4 *HttpRequest, c chan *Task) {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkHttp_SynchronousRequestAsync(z.ckObj, cstr1, C.int(arg2), b3, arg4.ckObj)

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


// Unlocks the Http class/component. It is necessary to call Http.UnlockComponent
// before calling any other methods. Passing any string to UnlockComponent will
// automatically activate a 30-day trial period.
func (z *Http) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttp_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// URL decodes a string.
// return a string or nil if failed.
func (z *Http) UrlDecode(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_urlDecode(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// URL encodes a string.
// return a string or nil if failed.
func (z *Http) UrlEncode(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttp_urlEncode(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Verifies a timestamp reply received from a timestamp authority (TSA). Returns
// the following possible integer values:
//     -1: The timestampReply does not contain a valid timestamp reply.
//     -2: The timestampReply is a valid timestamp reply, but failed verification using the
//     public key of the tsaCert.
//     0: Granted and verified.
//     1: Granted and verified, with mods (see RFC 3161)
//     2: Rejected.
//     3: Waiting.
//     4: Revocation Warning
//     5: Revocation Notification
// 
// If the timestamp reply (timestampReply) is known to be from a trusted source, then the
// tsaCert may be empty. If tsaCert is empty (never loaded with a certificate), then the
// verification will use the certificate embedded in the timestamp reply.
//
func (z *Http) VerifyTimestampReply(arg1 *BinData, arg2 *Cert) int {

    v := C.CkHttp_VerifyTimestampReply(z.ckObj, arg1.ckObj, arg2.ckObj)


    return int(v)
}


// Makes an XML RPC call to a URL endpoint. The XML string is passed in an HTTP
// POST, and the XML response is returned.
// return a string or nil if failed.
func (z *Http) XmlRpc(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkHttp_xmlRpc(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the XmlRpc method
func (z *Http) XmlRpcAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_XmlRpcAsync(z.ckObj, cstr1, cstr2)

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


// Same as XmlRpc, but uses the HTTP PUT method instead of the POST method.
// return a string or nil if failed.
func (z *Http) XmlRpcPut(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkHttp_xmlRpcPut(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the XmlRpcPut method
func (z *Http) XmlRpcPutAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkHttp_XmlRpcPutAsync(z.ckObj, cstr1, cstr2)

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


