// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkUpload.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewUpload() *Upload {
	return &Upload{C.CkUpload_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Upload) DisposeUpload() {
    if z != nil {
	C.CkUpload_Dispose(z.ckObj)
	}
}




func (z *Upload) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkUpload_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Upload) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkUpload_setExternalProgress(z.ckObj,1)
}

func (z *Upload) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkUpload_setExternalProgress(z.ckObj,1)
}

func (z *Upload) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkUpload_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Upload) AbortCurrent() bool {
    v := int(C.CkUpload_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Upload) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUpload_putAbortCurrent(z.ckObj,v)
}

// property: BandwidthThrottleUp
// If non-zero, limits (throttles) the upload bandwidth to approximately this
// maximum number of bytes per second. The default value of this property is 0.
func (z *Upload) BandwidthThrottleUp() int {
    return int(C.CkUpload_getBandwidthThrottleUp(z.ckObj))
}
// property setter: BandwidthThrottleUp
func (z *Upload) SetBandwidthThrottleUp(value int) {
    C.CkUpload_putBandwidthThrottleUp(z.ckObj,C.int(value))
}

// property: ChunkSize
// The chunk size (in bytes) used by the underlying TCP/IP sockets for uploading
// files. The default value is 65535.
func (z *Upload) ChunkSize() int {
    return int(C.CkUpload_getChunkSize(z.ckObj))
}
// property setter: ChunkSize
func (z *Upload) SetChunkSize(value int) {
    C.CkUpload_putChunkSize(z.ckObj,C.int(value))
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
func (z *Upload) ClientIpAddress() string {
    return C.GoString(C.CkUpload_clientIpAddress(z.ckObj))
}
// property setter: ClientIpAddress
func (z *Upload) SetClientIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putClientIpAddress(z.ckObj,cStr)
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
func (z *Upload) DebugLogFilePath() string {
    return C.GoString(C.CkUpload_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Upload) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Expect100Continue
// When true, the request header will included an "Expect: 100-continue" header
// field. This indicates that the server should respond with an intermediate
// response of "100 Continue" or "417 Expectation Failed" response based on the
// information available in the request header. This helps avoid situations such as
// limits on upload sizes. It allows the server to reject the upload, and then the
// client can abort prior to uploading the data.
// 
// The default value of this property is true.
//
func (z *Upload) Expect100Continue() bool {
    v := int(C.CkUpload_getExpect100Continue(z.ckObj))
    return v != 0
}
// property setter: Expect100Continue
func (z *Upload) SetExpect100Continue(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUpload_putExpect100Continue(z.ckObj,v)
}

// property: HeartbeatMs
// This property is only valid in programming environment and languages that allow
// for event callbacks.
// 
// Specifies the time interval in milliseconds between AbortCheck events. An Upload
// operation can be aborted via the AbortCheck event.
//
func (z *Upload) HeartbeatMs() int {
    return int(C.CkUpload_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Upload) SetHeartbeatMs(value int) {
    C.CkUpload_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: Hostname
// The hostname of the HTTP server that is the target of the upload. Do not include
// "http://" in the hostname. It can be a hostname, such as "www.chilkatsoft.com",
// or an IP address, such as "168.144.70.227".
func (z *Upload) Hostname() string {
    return C.GoString(C.CkUpload_hostname(z.ckObj))
}
// property setter: Hostname
func (z *Upload) SetHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IdleTimeoutMs
// A timeout in milliseconds. The default value is 30000. If the upload hangs (i.e.
// progress halts) for more than this time, the component will abort the upload.
// (It will timeout.)
func (z *Upload) IdleTimeoutMs() int {
    return int(C.CkUpload_getIdleTimeoutMs(z.ckObj))
}
// property setter: IdleTimeoutMs
func (z *Upload) SetIdleTimeoutMs(value int) {
    C.CkUpload_putIdleTimeoutMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Upload) LastErrorHtml() string {
    return C.GoString(C.CkUpload_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Upload) LastErrorText() string {
    return C.GoString(C.CkUpload_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Upload) LastErrorXml() string {
    return C.GoString(C.CkUpload_lastErrorXml(z.ckObj))
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
func (z *Upload) LastMethodSuccess() bool {
    v := int(C.CkUpload_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Upload) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUpload_putLastMethodSuccess(z.ckObj,v)
}

// property: Login
// The HTTP login for sites requiring authentication. Chilkat Upload supports Basic
// HTTP authentication.
func (z *Upload) Login() string {
    return C.GoString(C.CkUpload_login(z.ckObj))
}
// property setter: Login
func (z *Upload) SetLogin(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putLogin(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: NumBytesSent
// After an upload has completed, this property contains the number of bytes sent.
// During asynchronous uploads, this property contains the current number of bytes
// sent while the upload is in progress.
func (z *Upload) NumBytesSent() uint {
    return uint(C.CkUpload_getNumBytesSent(z.ckObj))
}

// property: Password
// The HTTP password for sites requiring authentication. Chilkat Upload supports
// Basic HTTP authentication.
func (z *Upload) Password() string {
    return C.GoString(C.CkUpload_password(z.ckObj))
}
// property setter: Password
func (z *Upload) SetPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Path
// The path part of the upload URL. Some examples:
// 
// If the upload target (i.e. consumer) URL is:
// http://www.freeaspupload.net/freeaspupload/testUpload.asp, then
// 
//     Hostname = "www.freeaspupload.net" Path = "/freeaspupload/testUpload.asp"
// 
// If the upload target URL is
// http://www.chilkatsoft.com/cgi-bin/ConsumeUpload.exe, then
// 
//     Hostname = "www.chilkatsoft.com" Path = "/cgi-bin/ConsumeUpload.exe"
//
func (z *Upload) Path() string {
    return C.GoString(C.CkUpload_path(z.ckObj))
}
// property setter: Path
func (z *Upload) SetPath(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putPath(z.ckObj,cStr)
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
func (z *Upload) PercentDoneScale() int {
    return int(C.CkUpload_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *Upload) SetPercentDoneScale(value int) {
    C.CkUpload_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: PercentUploaded
// Contains the current percentage completion (0 to 100) while an asynchronous
// upload is in progress.
func (z *Upload) PercentUploaded() uint {
    return uint(C.CkUpload_getPercentUploaded(z.ckObj))
}

// property: Port
// The port number of the upload target (i.e. consumer) URL. The default value is
// 80. If SSL is used, this should be set to 443 (typically).
func (z *Upload) Port() int {
    return int(C.CkUpload_getPort(z.ckObj))
}
// property setter: Port
func (z *Upload) SetPort(value int) {
    C.CkUpload_putPort(z.ckObj,C.int(value))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *Upload) PreferIpv6() bool {
    v := int(C.CkUpload_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *Upload) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUpload_putPreferIpv6(z.ckObj,v)
}

// property: ProxyDomain
// The domain name of a proxy host if an HTTP proxy is used. Do not include the
// "http://". The domain name may be a hostname, such as "www.chilkatsoft.com", or
// an IP address, such as "168.144.70.227".
func (z *Upload) ProxyDomain() string {
    return C.GoString(C.CkUpload_proxyDomain(z.ckObj))
}
// property setter: ProxyDomain
func (z *Upload) SetProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyLogin
// If an HTTP proxy is used and it requires authentication, this property specifies
// the HTTP proxy login.
func (z *Upload) ProxyLogin() string {
    return C.GoString(C.CkUpload_proxyLogin(z.ckObj))
}
// property setter: ProxyLogin
func (z *Upload) SetProxyLogin(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putProxyLogin(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyPassword
// If an HTTP proxy is used and it requires authentication, this property specifies
// the HTTP proxy password.
func (z *Upload) ProxyPassword() string {
    return C.GoString(C.CkUpload_proxyPassword(z.ckObj))
}
// property setter: ProxyPassword
func (z *Upload) SetProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyPort
// The port number of a proxy server if an HTTP proxy is used.
func (z *Upload) ProxyPort() int {
    return int(C.CkUpload_getProxyPort(z.ckObj))
}
// property setter: ProxyPort
func (z *Upload) SetProxyPort(value int) {
    C.CkUpload_putProxyPort(z.ckObj,C.int(value))
}

// property: ResponseBody
// An HTTP upload is nothing more than an HTTP POST that contains the content of
// the files being uploaded. Just as with any HTTP POST or GET, the server should
// send an HTTP response that consists of header and body.
// 
// This property contains the body part of the HTTP response.
//
func (z *Upload) ResponseBody() []byte {
    ckbyd := C.CkByteData_Create()
    C.CkUpload_getResponseBody(z.ckObj,ckbyd)
    p := C.CkByteData_getData(ckbyd);
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckbyd)) )
    C.CkByteData_Dispose(ckbyd)
    return retBytes
}


// property: ResponseBodyStr
// Returns the response body as a string.
func (z *Upload) ResponseBodyStr() string {
    return C.GoString(C.CkUpload_responseBodyStr(z.ckObj))
}

// property: ResponseHeader
// An HTTP upload is nothing more than an HTTP POST that contains the content of
// the files being uploaded. Just as with any HTTP POST or GET, the server should
// send an HTTP response that consists of header and body.
// 
// This property contains the header part of the HTTP response.
//
func (z *Upload) ResponseHeader() string {
    return C.GoString(C.CkUpload_responseHeader(z.ckObj))
}

// property: ResponseStatus
// The HTTP response status code of the HTTP response. A list of HTTP status codes
// can be found here:HTTP Response Status Codes
// <http://en.wikipedia.org/wiki/List_of_HTTP_status_codes>.
func (z *Upload) ResponseStatus() int {
    return int(C.CkUpload_getResponseStatus(z.ckObj))
}

// property: Ssl
// Set this to true if the upload is to HTTPS. For example, if the target of the
// upload is:
// 
//     https://www.myuploadtarget.com/consumeUpload.asp
// 
// then set:
// 
//     Ssl = true
//     Hostname = "www.myuploadtarget.com"
//     Path = "/consumeupload.asp"
//     Port = 443
//     
//
func (z *Upload) Ssl() bool {
    v := int(C.CkUpload_getSsl(z.ckObj))
    return v != 0
}
// property setter: Ssl
func (z *Upload) SetSsl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUpload_putSsl(z.ckObj,v)
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
func (z *Upload) SslAllowedCiphers() string {
    return C.GoString(C.CkUpload_sslAllowedCiphers(z.ckObj))
}
// property setter: SslAllowedCiphers
func (z *Upload) SetSslAllowedCiphers(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putSslAllowedCiphers(z.ckObj,cStr)
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
func (z *Upload) SslProtocol() string {
    return C.GoString(C.CkUpload_sslProtocol(z.ckObj))
}
// property setter: SslProtocol
func (z *Upload) SetSslProtocol(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putSslProtocol(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
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
func (z *Upload) TlsPinSet() string {
    return C.GoString(C.CkUpload_tlsPinSet(z.ckObj))
}
// property setter: TlsPinSet
func (z *Upload) SetTlsPinSet(goStr string) {
    cStr := C.CString(goStr)
    C.CkUpload_putTlsPinSet(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TotalUploadSize
// The total size of the upload (in bytes). This property will become set at the
// beginning of an asynchronous upload. A program may monitor asynchronous uploads
// by tracking both NumBytesSent and PercentUploaded.
// 
// This property is also set during synchronous uploads.
//
func (z *Upload) TotalUploadSize() uint {
    return uint(C.CkUpload_getTotalUploadSize(z.ckObj))
}

// property: UploadInProgress
// Set to true when an asynchronous upload is started. When the asynchronous
// upload is complete, this property becomes equal to false. A program will
// typically begin an asynchronous upload by calling BeginUpload, and then
// periodically checking the value of this property to determine when the upload is
// complete.
func (z *Upload) UploadInProgress() bool {
    v := int(C.CkUpload_getUploadInProgress(z.ckObj))
    return v != 0
}

// property: UploadSuccess
// Set to true (success) or false (failed) after an asynchronous upload
// completes or aborts due to failure. When a program does an asynchronous upload,
// it will wait until UploadInProgress becomes false. It will then check the
// value of this property to determine if the upload was successful or not.
func (z *Upload) UploadSuccess() bool {
    v := int(C.CkUpload_getUploadSuccess(z.ckObj))
    return v != 0
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Upload) VerboseLogging() bool {
    v := int(C.CkUpload_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Upload) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUpload_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Upload) Version() string {
    return C.GoString(C.CkUpload_version(z.ckObj))
}
// May be called during an asynchronous upload to abort.
func (z *Upload) AbortUpload()  {

    C.CkUpload_AbortUpload(z.ckObj)



}


// Adds a custom HTTP header to the HTTP upload.
func (z *Upload) AddCustomHeader(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkUpload_AddCustomHeader(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Adds a file to the list of files to be uploaded in the next call to
// BlockingUpload, BeginUpload, or UploadToMemory. To upload more than one file,
// call this method once for each file to be uploaded.
func (z *Upload) AddFileReference(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkUpload_AddFileReference(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Adds a custom HTTP request parameter to the upload.
func (z *Upload) AddParam(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkUpload_AddParam(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Starts an asynchronous upload. Only one asynchronous upload may be in progress
// at a time. To achieve multiple asynchronous uploads, use multiple instances of
// the Chilkat Upload object. Each object instance is capable of managing a single
// asynchronous upload.
// 
// When this method is called, a background thread is started and the asynchronous
// upload runs in the background. The upload may be aborted at any time by calling
// AbortUpload. The upload is completed (or failed) when UploadInProgress becomes
// false. At that point, the UploadSuccess property may be checked to determine
// success (true) or failure (false).
//
func (z *Upload) BeginUpload() bool {

    v := C.CkUpload_BeginUpload(z.ckObj)


    return v != 0
}


// Uploads files to a web server using HTTP. The files to be uploaded are indicated
// by calling AddFileReference once for each file (prior to calling this method).
func (z *Upload) BlockingUpload() bool {

    v := C.CkUpload_BlockingUpload(z.ckObj)


    return v != 0
}


// Asynchronous version of the BlockingUpload method
func (z *Upload) BlockingUploadAsync(c chan *Task) {

    hTask := C.CkUpload_BlockingUploadAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Clears the internal list of files created by calls to AddFileReference.
func (z *Upload) ClearFileReferences()  {

    C.CkUpload_ClearFileReferences(z.ckObj)



}


// Clears the internal list of params created by calls to AddParam.
func (z *Upload) ClearParams()  {

    C.CkUpload_ClearParams(z.ckObj)



}


// Loads the caller of the task's async method.
func (z *Upload) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkUpload_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// A convenience method for putting the calling process to sleep for N
// milliseconds. It is provided here because some programming languages do not
// provide this capability, and sleeping for short increments of time is helpful
// when doing asynchronous uploads.
func (z *Upload) SleepMs(arg1 int)  {

    C.CkUpload_SleepMs(z.ckObj, C.int(arg1))



}


// Writes the complete HTTP POST to memory. The POST contains the HTTP header, any
// custom params added by calling AddParam, and each file to be uploaded. This is
// helpful in debugging. It allows you to see the exact HTTP POST sent to the
// server if BlockingUpload or BeginUpload is called.
func (z *Upload) UploadToMemory() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkUpload_UploadToMemory(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


