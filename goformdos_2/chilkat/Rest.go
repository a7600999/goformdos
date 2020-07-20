// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkRest.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewRest() *Rest {
	return &Rest{C.CkRest_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Rest) DisposeRest() {
    if z != nil {
	C.CkRest_Dispose(z.ckObj)
	}
}




func (z *Rest) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkRest_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Rest) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkRest_setExternalProgress(z.ckObj,1)
}

func (z *Rest) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkRest_setExternalProgress(z.ckObj,1)
}

func (z *Rest) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkRest_setExternalProgress(z.ckObj,1)
}




// property: AllowHeaderFolding
// If this property is set to false, then no MIME header folding will be
// automatically applied to any request header. The default is true. This
// property is provided to satisfy certain providers, such as Quickbooks, that
// require all MIME headers to be single unfolded lines regardless of length.
func (z *Rest) AllowHeaderFolding() bool {
    v := int(C.CkRest_getAllowHeaderFolding(z.ckObj))
    return v != 0
}
// property setter: AllowHeaderFolding
func (z *Rest) SetAllowHeaderFolding(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRest_putAllowHeaderFolding(z.ckObj,v)
}

// property: AllowHeaderQB
// Controls whether non us-ascii HTTP request headers are automatically Q/B
// encoded. The default value is true.
// 
// Q/B encoded headers explicitly indicate the charset and byte representation, and
// appear as such: =?utf-8?Q?...?= or =?utf-8?B?...?=, where the charset may be
// "utf-8" or any other possible charset.
// 
// If this property is set to false, then no Q/B encoding is applied to any
// request header.
//
func (z *Rest) AllowHeaderQB() bool {
    v := int(C.CkRest_getAllowHeaderQB(z.ckObj))
    return v != 0
}
// property setter: AllowHeaderQB
func (z *Rest) SetAllowHeaderQB(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRest_putAllowHeaderQB(z.ckObj,v)
}

// property: Authorization
// The value of the Authorization HTTP request header (if needed).
func (z *Rest) Authorization() string {
    return C.GoString(C.CkRest_authorization(z.ckObj))
}
// property setter: Authorization
func (z *Rest) SetAuthorization(goStr string) {
    cStr := C.CString(goStr)
    C.CkRest_putAuthorization(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConnectFailReason
// If the Connect method fails, this property can be checked to determine the
// reason for failure.
// 
// Possible values are:
// 0 = success
// 
// Normal (non-SSL) sockets:
// 1 = empty hostname
// 2 = DNS lookup failed
// 3 = DNS timeout
// 4 = Aborted by application.
// 5 = Internal failure.
// 6 = Connect Timed Out
// 7 = Connect Rejected (or failed for some other reason)
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
// 109 = Failed to read handshake messages.
// 110 = Failed to send client certificate handshake message.
// 111 = Failed to send client key exchange handshake message.
// 112 = Client certificate's private key not accessible.
// 113 = Failed to send client cert verify handshake message.
// 114 = Failed to send change cipher spec handshake message.
// 115 = Failed to send finished handshake message.
// 116 = Server's Finished message is invalid.
//
func (z *Rest) ConnectFailReason() int {
    return int(C.CkRest_getConnectFailReason(z.ckObj))
}

// property: ConnectTimeoutMs
// The maximum amount of time to wait for the connection to be accepted by the HTTP
// server.
// 
// Note: Suprisingly, this property was forgotten and not added until Chilkat
// v9.5.0.71.
//
func (z *Rest) ConnectTimeoutMs() int {
    return int(C.CkRest_getConnectTimeoutMs(z.ckObj))
}
// property setter: ConnectTimeoutMs
func (z *Rest) SetConnectTimeoutMs(value int) {
    C.CkRest_putConnectTimeoutMs(z.ckObj,C.int(value))
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
func (z *Rest) DebugLogFilePath() string {
    return C.GoString(C.CkRest_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Rest) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkRest_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DebugMode
// If true then all calls to Send* or FullRequest* methods will not actually send
// a request. Instead, the request will be written to a memory buffer which can
// then be retrieved by calling GetLastDebugRequest.
func (z *Rest) DebugMode() bool {
    v := int(C.CkRest_getDebugMode(z.ckObj))
    return v != 0
}
// property setter: DebugMode
func (z *Rest) SetDebugMode(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRest_putDebugMode(z.ckObj,v)
}

// property: HeartbeatMs
// This property is only valid in programming environment and languages that allow
// for event callbacks.
// 
// Specifies the time interval in milliseconds between AbortCheck events. A value
// of 0 (the default) indicate that no AbortCheck events will fire. Any REST method
// can be aborted via the AbortCheck event.
//
func (z *Rest) HeartbeatMs() int {
    return int(C.CkRest_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Rest) SetHeartbeatMs(value int) {
    C.CkRest_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: Host
// The value of the Host HTTP request header.
func (z *Rest) Host() string {
    return C.GoString(C.CkRest_host(z.ckObj))
}
// property setter: Host
func (z *Rest) SetHost(goStr string) {
    cStr := C.CString(goStr)
    C.CkRest_putHost(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IdleTimeoutMs
// The maximum amount of time to wait for additional incoming data when receiving,
// or the max time to wait to send additional data. The default value is 30000 (30
// seconds). This is not an overall max timeout. Rather, it is the maximum time to
// wait when receiving or sending has halted.
func (z *Rest) IdleTimeoutMs() int {
    return int(C.CkRest_getIdleTimeoutMs(z.ckObj))
}
// property setter: IdleTimeoutMs
func (z *Rest) SetIdleTimeoutMs(value int) {
    C.CkRest_putIdleTimeoutMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Rest) LastErrorHtml() string {
    return C.GoString(C.CkRest_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Rest) LastErrorText() string {
    return C.GoString(C.CkRest_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Rest) LastErrorXml() string {
    return C.GoString(C.CkRest_lastErrorXml(z.ckObj))
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
func (z *Rest) LastMethodSuccess() bool {
    v := int(C.CkRest_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Rest) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRest_putLastMethodSuccess(z.ckObj,v)
}

// property: LastRequestHeader
// The full MIME header (not including the HTTP start line which contains the
// status code and status text), of the last request sent.
func (z *Rest) LastRequestHeader() string {
    return C.GoString(C.CkRest_lastRequestHeader(z.ckObj))
}

// property: LastRequestStartLine
// The full start line of the last request sent. (The start line begins with the
// HTTP verb, such as GET, POST, etc., is followed by the URI path, and ends with
// the HTTP version.)
func (z *Rest) LastRequestStartLine() string {
    return C.GoString(C.CkRest_lastRequestStartLine(z.ckObj))
}

// property: NumResponseHeaders
// The number of response header fields. The first response header field is at
// index 0.
func (z *Rest) NumResponseHeaders() int {
    return int(C.CkRest_getNumResponseHeaders(z.ckObj))
}

// property: PartSelector
// Only used for multipart requests. Selects the target MIME part for calls to
// AddHeader, RemoveHeader, SetMultipartBodyBinary, SetMultipartBodyString,
// SetMultipartBodyStream, etc. The default is an empty string and indicates the
// top-level header. A string value of "1" would indicate the 1st sub-part in a
// multipart request. A string value of "1.2" would indicate the 2nd sub-part under
// the 1st sub-part.
// 
// It is unlikely you'll ever encounter the need for nested multipart requests
// (i.e. part selectors such as "1.2") Also, most REST requests are NOT multipart,
// and therefore this feature is rarely used. An example of a multipart REST
// request would be for a Google Drive upload, where the top-level Content-Type is
// multipart/related, the1st sub-part contains the JSON meta-data, and the 2nd
// sub-part contains the file data.
//
func (z *Rest) PartSelector() string {
    return C.GoString(C.CkRest_partSelector(z.ckObj))
}
// property setter: PartSelector
func (z *Rest) SetPartSelector(goStr string) {
    cStr := C.CString(goStr)
    C.CkRest_putPartSelector(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: PercentDoneOnSend
// This property only applies to the FullRequest* methods, which are methods that
// both send an HTTP request and receive the response. (It also only applies to
// programming languages that support event callbacks.) It determines whether
// percentage completion is tracked for the sending of the HTTP request, or for the
// downloading the HTTP response. The default value is false, which is to measure
// the percent completion when receiving the response.
// 
// For example: If the REST request is to download a file, then this property
// should remain at the default value of false. If the REST request is to upload
// a file (using a Full* method), then set this property to true. Also note if a
// server sends an HTTP response in the chunked encoding, it is not possible to
// measure percent completion because the HTTP client has no way of knowing the
// total size of the HTTP response.
//
func (z *Rest) PercentDoneOnSend() bool {
    v := int(C.CkRest_getPercentDoneOnSend(z.ckObj))
    return v != 0
}
// property setter: PercentDoneOnSend
func (z *Rest) SetPercentDoneOnSend(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRest_putPercentDoneOnSend(z.ckObj,v)
}

// property: ResponseHeader
// The full response MIME header (not including the HTTP start line which contains
// the status code and status text).
func (z *Rest) ResponseHeader() string {
    return C.GoString(C.CkRest_responseHeader(z.ckObj))
}

// property: ResponseStatusCode
// The response status code.
func (z *Rest) ResponseStatusCode() int {
    return int(C.CkRest_getResponseStatusCode(z.ckObj))
}

// property: ResponseStatusText
// The status message corresponding to the response status code.
func (z *Rest) ResponseStatusText() string {
    return C.GoString(C.CkRest_responseStatusText(z.ckObj))
}

// property: StreamNonChunked
// If true, then methods that upload data are sent non-chunked if possible. For
// example, if the FullRequestStream method is called where the stream is a file
// stream, then the size of the content is known and the HTTP request will be sent
// using a Content-Length header instead of using a Transfer-Encoding: chunked
// upload. If false, then the chunked transfer encoding is used. The default
// value of this property is true.
func (z *Rest) StreamNonChunked() bool {
    v := int(C.CkRest_getStreamNonChunked(z.ckObj))
    return v != 0
}
// property setter: StreamNonChunked
func (z *Rest) SetStreamNonChunked(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRest_putStreamNonChunked(z.ckObj,v)
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. This property
// defaults to the empty string and should typically remain empty. Can be set to a
// list of the following comma separated keywords:
//     "ProtectFromVpn" - Introduced in v9.5.0.80. On Android systems, will bypass
//     any VPN that may be installed or active.
//     "AllowInsecureBasicAuth" - Introduced in v9.5.0.80. Allows HTTP Basic
//     authentication over non-TLS connections. (HTTP Basic authentication puts the
//     login:password in the Authorization header field in base64 encoding. If
//     transmitted over an insecure connection, it is potentially visible to anything
//     sniffing the traffic. By default, Chilkat disallows HTTP Basic authentication
//     except for SSL/TLS connections, SSH tunneled connections, or connections to
//     localhost.)
//     "EnableTls13" - Introduced in v9.5.0.82. Causes TLS 1.3 to be offered in the
//     ClientHello of the TLS protocol, allowing the server to select TLS 1.3 for the
//     session. Future versions of Chilkat will enable TLS 1.3 by default. This option
//     is only necessary in v9.5.0.82 if TLS 1.3 is desired.
func (z *Rest) UncommonOptions() string {
    return C.GoString(C.CkRest_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Rest) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkRest_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Rest) VerboseLogging() bool {
    v := int(C.CkRest_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Rest) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRest_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Rest) Version() string {
    return C.GoString(C.CkRest_version(z.ckObj))
}
// Adds an HTTP request header. If the header field already exists, then it is
// replaced.
func (z *Rest) AddHeader(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_AddHeader(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Computes the Amazon MWS signature using the mwsSecretKey and adds the "Signature"
// parameter to the request. This method should be called for all Amazon
// Marketplace Web Service (Amazon MWS) HTTP requests. It should be called after
// all request parameters have been added.
// 
// Important: The Chilkat v9.5.0.75 release accidentally breaks Amazon MWS (not
// AWS) authentication. If you need MWS with 9.5.0.75, send email to
// support@chilkatsoft.com for a hotfix, or revert back to v9.5.0.73, or update to
// a version after 9.5.0.75.
// 
// The domain should be the domain of the request, such as one of the following:
//     mws.amazonservices.com
//     mws-eu.amazonservices.com
//     mws.amazonservices.in
//     mws.amazonservices.com.cn
//     mws.amazonservices.jp
// 
// The httpVerb should be the HTTP verb, such as "GET", "POST", etc. The uriPath is the
// URI path, such as "/Feeds/2009-01-01". In general, the httpVerb and uriPath should be
// identical to the 1st two args passed to methods such as
// FullRequestFormUrlEncoded.
// 
// Note: This method also automatically adds or replaces the existing Timestamp
// parameter to the current system date/time.
//
func (z *Rest) AddMwsSignature(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkRest_AddMwsSignature(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Adds or replaces a path parameter. A path parameter is a string that will be
// replaced in any uriPath string passed to a method. For example, if name is
// "fileId" and value is "1R_70heIyzIAu1_u0prXbYcaIiJRVkgBl", then a uriPath
// argument of "/drive/v3/files/fileId" will be transformed to
// "/drive/v3/files/1R_70heIyzIAu1_u0prXbYcaIiJRVkgBl" in a method call.
func (z *Rest) AddPathParam(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_AddPathParam(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a query parameter. If the query parameter already exists, then it is
// replaced.
func (z *Rest) AddQueryParam(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_AddQueryParam(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds the query parameters from the queryString. The queryString is a query string of the
// format field1=value1&field2=value2&field3=value3... where each value is URL
// encoded.
func (z *Rest) AddQueryParams(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRest_AddQueryParams(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds a query parameter. If the query parameter already exists, then it is
// replaced. The parameter value is passed in a StringBuilder object.
func (z *Rest) AddQueryParamSb(arg1 string, arg2 *StringBuilder) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRest_AddQueryParamSb(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Removes all HTTP request headers.
func (z *Rest) ClearAllHeaders() bool {

    v := C.CkRest_ClearAllHeaders(z.ckObj)


    return v != 0
}


// Removes all sub-parts from a request. This is useful when preparing the REST
// object to send a new request after a multipart request has just been sent.
func (z *Rest) ClearAllParts() bool {

    v := C.CkRest_ClearAllParts(z.ckObj)


    return v != 0
}


// Clears all path parameters.
func (z *Rest) ClearAllPathParams() bool {

    v := C.CkRest_ClearAllPathParams(z.ckObj)


    return v != 0
}


// Clears all query parameters.
func (z *Rest) ClearAllQueryParams() bool {

    v := C.CkRest_ClearAllQueryParams(z.ckObj)


    return v != 0
}


// Clears all authentication settings.
func (z *Rest) ClearAuth() bool {

    v := C.CkRest_ClearAuth(z.ckObj)


    return v != 0
}


// Clears the response body stream.
func (z *Rest) ClearResponseBodyStream()  {

    C.CkRest_ClearResponseBodyStream(z.ckObj)



}


// Establishes an initial connection to a REST server. The hostname can be a domain
// name or an IP address. Both IPv4 and IPv6 addresses are supported. The port is
// the port, which is typically 80 or 443. If SSL/TLS is required, then tls should
// be set to true. The autoReconnect indicates whether connection should automatically be
// established as needed for subsequent REST requests.
// 
// Note: This method is for simple connections that do not require any proxies
// (HTTP or SOCKS), or SSH tunneling. If a proxy, SSH tunnel, or any other advanced
// socket feature is required, the Chilkat Socket API can be used to establish the
// connection. The UseConnection method can then be called to use the
// pre-established socket connection.
//
func (z *Rest) Connect(arg1 string, arg2 int, arg3 bool, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkRest_Connect(z.ckObj, cstr1, C.int(arg2), b3, b4)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Connect method
func (z *Rest) ConnectAsync(arg1 string, arg2 int, arg3 bool, arg4 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkRest_ConnectAsync(z.ckObj, cstr1, C.int(arg2), b3, b4)

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


// Closes the connection with the HTTP server if one is open. This method can be
// called to ensure the connection is closed. The maxWaitMs is the maximum time in
// milliseconds to wait for a clean close. If the connection is through an SSH
// tunnel, this closes the logical channel within the SSH tunnel, and not the
// connection with the SSH server itself.
func (z *Rest) Disconnect(arg1 int) bool {

    v := C.CkRest_Disconnect(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the Disconnect method
func (z *Rest) DisconnectAsync(arg1 int, c chan *Task) {

    hTask := C.CkRest_DisconnectAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a complete REST request (header + binary body) and receives the full
// response. The binary body of the request is passed in binData. The response body is
// returned in responseBody (replacing whatever contents responseBody may have previously held).
// 
// Note: If your application wishes to receive percent-done callbacks, make sure
// the PercentDoneOnSend property is set to indicate whether percent-done applies
// to sending or receiving.
//
func (z *Rest) FullRequestBd(arg1 string, arg2 string, arg3 *BinData, arg4 *StringBuilder) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_FullRequestBd(z.ckObj, cstr1, cstr2, arg3.ckObj, arg4.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the FullRequestBd method
func (z *Rest) FullRequestBdAsync(arg1 string, arg2 string, arg3 *BinData, arg4 *StringBuilder, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_FullRequestBdAsync(z.ckObj, cstr1, cstr2, arg3.ckObj, arg4.ckObj)

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


// Sends a complete REST request (header + body) and receives the full response. It
// is assumed that the request body is binary, and the response body is a string
// (such as JSON or XML). The response body is returned.
// 
// This method is the equivalent of making the following calls in sequence:
// SendReqBinaryBody, ReadResponseHeader, ReadRespBodyString.
// 
// Note: If your application wishes to receive percent-done callbacks, make sure
// the PercentDoneOnSend property is set to indicate whether percent-done applies
// to sending or receiving.
//
// return a string or nil if failed.
func (z *Rest) FullRequestBinary(arg1 string, arg2 string, arg3 []byte) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    retStr := C.CkRest_fullRequestBinary(z.ckObj, cstr1, cstr2, ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FullRequestBinary method
func (z *Rest) FullRequestBinaryAsync(arg1 string, arg2 string, arg3 []byte, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    hTask := C.CkRest_FullRequestBinaryAsync(z.ckObj, cstr1, cstr2, ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a complete application/x-www-form-urlencoded HTTP request and receives the
// full response. The request body is composed of the URL encoded query params. It
// is assumed that the response body is a string (such as JSON or XML). The
// response body is returned.
// 
// This method is the equivalent of making the following calls in sequence:
// SendReqFormUrlEncoded, ReadResponseHeader, ReadRespBodyString.
// 
// Note: If your application wishes to receive percent-done callbacks, make sure
// the PercentDoneOnSend property is set to indicate whether percent-done applies
// to sending or receiving.
//
// return a string or nil if failed.
func (z *Rest) FullRequestFormUrlEncoded(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkRest_fullRequestFormUrlEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FullRequestFormUrlEncoded method
func (z *Rest) FullRequestFormUrlEncodedAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_FullRequestFormUrlEncodedAsync(z.ckObj, cstr1, cstr2)

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


// Sends a complete multipart REST request (header + multipart body) and receives
// the full response. It is assumed that the response body is a string (such as
// JSON or XML). The response body is returned.
// 
// This method is the equivalent of making the following calls in sequence:
// SendReqMultipart, ReadResponseHeader, ReadRespBodyString.
// 
// Note: If your application wishes to receive percent-done callbacks, make sure
// the PercentDoneOnSend property is set to indicate whether percent-done applies
// to sending or receiving.
//
// return a string or nil if failed.
func (z *Rest) FullRequestMultipart(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkRest_fullRequestMultipart(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FullRequestMultipart method
func (z *Rest) FullRequestMultipartAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_FullRequestMultipartAsync(z.ckObj, cstr1, cstr2)

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


// Sends a complete REST request and receives the full response. It is assumed that
// the response body is a string (such as JSON or XML). The response body is
// returned.
// 
// This method is the equivalent of making the following calls in sequence:
// SendReqNoBody, ReadResponseHeader, ReadRespBodyString.
//
// return a string or nil if failed.
func (z *Rest) FullRequestNoBody(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkRest_fullRequestNoBody(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FullRequestNoBody method
func (z *Rest) FullRequestNoBodyAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_FullRequestNoBodyAsync(z.ckObj, cstr1, cstr2)

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


// The same as FullRequestNoBody, except returns the response body in the binData.
// This method is useful for downloading binary files.
func (z *Rest) FullRequestNoBodyBd(arg1 string, arg2 string, arg3 *BinData) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_FullRequestNoBodyBd(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the FullRequestNoBodyBd method
func (z *Rest) FullRequestNoBodyBdAsync(arg1 string, arg2 string, arg3 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_FullRequestNoBodyBdAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// The same as FullRequestNoBody, except returns the response body in the sb.
func (z *Rest) FullRequestNoBodySb(arg1 string, arg2 string, arg3 *StringBuilder) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_FullRequestNoBodySb(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the FullRequestNoBodySb method
func (z *Rest) FullRequestNoBodySbAsync(arg1 string, arg2 string, arg3 *StringBuilder, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_FullRequestNoBodySbAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// Sends a complete REST request (header + body string) and receives the full
// response. The body of the request is passed in requestBody. The response body is
// returned in responseBody (replacing whatever contents responseBody may have previously held).
// 
// Note: If your application wishes to receive percent-done callbacks, make sure
// the PercentDoneOnSend property is set to indicate whether percent-done applies
// to sending or receiving.
//
func (z *Rest) FullRequestSb(arg1 string, arg2 string, arg3 *StringBuilder, arg4 *StringBuilder) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_FullRequestSb(z.ckObj, cstr1, cstr2, arg3.ckObj, arg4.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the FullRequestSb method
func (z *Rest) FullRequestSbAsync(arg1 string, arg2 string, arg3 *StringBuilder, arg4 *StringBuilder, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_FullRequestSbAsync(z.ckObj, cstr1, cstr2, arg3.ckObj, arg4.ckObj)

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


// Sends a complete REST request and receives the full response. It is assumed that
// the response body is a string (such as JSON or XML). The response body is
// returned.
// 
// This method is the equivalent of making the following calls in sequence:
// SendReqStream, ReadResponseHeader, ReadRespBodyString.
// 
// Note: If your application wishes to receive percent-done callbacks, make sure
// the PercentDoneOnSend property is set to indicate whether percent-done applies
// to sending or receiving.
//
// return a string or nil if failed.
func (z *Rest) FullRequestStream(arg1 string, arg2 string, arg3 *Stream) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkRest_fullRequestStream(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FullRequestStream method
func (z *Rest) FullRequestStreamAsync(arg1 string, arg2 string, arg3 *Stream, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_FullRequestStreamAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// Sends a complete REST request (header + body string) and receives the full
// response. It is assumed that both the request and response bodies are strings
// (such as JSON or XML). The response body is returned.
// 
// This method is the equivalent of making the following calls in sequence:
// SendReqStringBody, ReadResponseHeader, ReadRespBodyString.
// 
// Note: If your application wishes to receive percent-done callbacks, make sure
// the PercentDoneOnSend property is set to indicate whether percent-done applies
// to sending or receiving.
//
// return a string or nil if failed.
func (z *Rest) FullRequestString(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkRest_fullRequestString(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FullRequestString method
func (z *Rest) FullRequestStringAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkRest_FullRequestStringAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Returns the fully composed HTTP request that would've been sent had the
// DebugMode been turned off. The request is returned in bd.
func (z *Rest) GetLastDebugRequest(arg1 *BinData) bool {

    v := C.CkRest_GetLastDebugRequest(z.ckObj, arg1.ckObj)


    return v != 0
}


// Provides information about what transpired in the last method called on this
// object instance. For many methods, there is no information. However, for some
// methods, details about what occurred can be obtained by getting the LastJsonData
// right after the method call returns.
func (z *Rest) LastJsonData() *JsonObject {

    retObj := C.CkRest_LastJsonData(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Loads the caller of the task's async method.
func (z *Rest) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkRest_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Reads the response body. Should only be called after ReadResponseHeader has been
// called, and should only be called when it is already known that the response
// body is binary, such as for JPG images or other non-text binary file types. The
// response body is received into responseBody (replacing whatever contents responseBody may have
// previously held).
func (z *Rest) ReadRespBd(arg1 *BinData) bool {

    v := C.CkRest_ReadRespBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the ReadRespBd method
func (z *Rest) ReadRespBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkRest_ReadRespBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads the response body. Should only be called after ReadResponseHeader has been
// called, and should only be called when it is already known that the response
// body is binary, such as for JPG images or other non-text binary file types.
func (z *Rest) ReadRespBodyBinary() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRest_ReadRespBodyBinary(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the ReadRespBodyBinary method
func (z *Rest) ReadRespBodyBinaryAsync(c chan *Task) {

    hTask := C.CkRest_ReadRespBodyBinaryAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads the response body to the stream. If autoSetStreamCharset is true, then the stream's
// StringCharset property will automatically get set based on the charset, if any,
// indicated in the response header. If the response is binary, then autoSetStreamCharset is
// ignored.
func (z *Rest) ReadRespBodyStream(arg1 *Stream, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkRest_ReadRespBodyStream(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Asynchronous version of the ReadRespBodyStream method
func (z *Rest) ReadRespBodyStreamAsync(arg1 *Stream, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkRest_ReadRespBodyStreamAsync(z.ckObj, arg1.ckObj, b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads the response body. Should only be called after ReadResponseHeader has been
// called, and should only be called when it is already known that the response
// body will be a string (such as XML, JSON, etc.)
// return a string or nil if failed.
func (z *Rest) ReadRespBodyString() *string {

    retStr := C.CkRest_readRespBodyString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReadRespBodyString method
func (z *Rest) ReadRespBodyStringAsync(c chan *Task) {

    hTask := C.CkRest_ReadRespBodyStringAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads the HTTP response header. If the HTTP response includes a body, then the
// application must call the desired method to read the response body. Otherwise,
// the HTTP request / response is finished after reading the response header (such
// as for a GET request). The contents of the response header are available in
// various properties and methods.
// 
// The HTTP response status code is returned (such as 200 for a typical success
// response). If an error occurred such that no response was received, then a value
// of -1 is returned.
//
func (z *Rest) ReadResponseHeader() int {

    v := C.CkRest_ReadResponseHeader(z.ckObj)


    return int(v)
}


// Asynchronous version of the ReadResponseHeader method
func (z *Rest) ReadResponseHeaderAsync(c chan *Task) {

    hTask := C.CkRest_ReadResponseHeaderAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads the response body. Should only be called after ReadResponseHeader has been
// called, and should only be called when it is already known that the response
// body will be a string (such as XML, JSON, etc.) The response body is stored in
// responseBody. (replacing whatever contents responseBody may have previously held).
func (z *Rest) ReadRespSb(arg1 *StringBuilder) bool {

    v := C.CkRest_ReadRespSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the ReadRespSb method
func (z *Rest) ReadRespSbAsync(arg1 *StringBuilder, c chan *Task) {

    hTask := C.CkRest_ReadRespSbAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// If the response was a redirect and contains a Location header field, this method
// returns the redirect URL.
func (z *Rest) RedirectUrl() *Url {

    retObj := C.CkRest_RedirectUrl(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Url{retObj}
}


// Removes all headers having the given name.
func (z *Rest) RemoveHeader(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRest_RemoveHeader(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Removes all query params having the given name.
func (z *Rest) RemoveQueryParam(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRest_RemoveQueryParam(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the value of the response header indicated by name.
// return a string or nil if failed.
func (z *Rest) ResponseHdrByName(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkRest_responseHdrByName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the name of the Nth response header field. (Chilkat always uses 0-based
// indexing. The first header field is at index 0.)
// return a string or nil if failed.
func (z *Rest) ResponseHdrName(arg1 int) *string {

    retStr := C.CkRest_responseHdrName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of the Nth response header field. (Chilkat always uses 0-based
// indexing. The first header field is at index 0.)
// return a string or nil if failed.
func (z *Rest) ResponseHdrValue(arg1 int) *string {

    retStr := C.CkRest_responseHdrValue(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sends a REST request that cotnains a binary body. The httpVerb is the HTTP verb
// (also known as the HTTP method), such as "PUT". The uriPath is the path of the
// resource URI. The body contains the bytes of the HTTP request body.
func (z *Rest) SendReqBd(arg1 string, arg2 string, arg3 *BinData) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_SendReqBd(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqBd method
func (z *Rest) SendReqBdAsync(arg1 string, arg2 string, arg3 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_SendReqBdAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// Sends a REST request that cotnains a binary body. The httpVerb is the HTTP verb
// (also known as the HTTP method), such as "PUT". The uriPath is the path of the
// resource URI. The body contains the bytes of the HTTP request body.
func (z *Rest) SendReqBinaryBody(arg1 string, arg2 string, arg3 []byte) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkRest_SendReqBinaryBody(z.ckObj, cstr1, cstr2, ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// Asynchronous version of the SendReqBinaryBody method
func (z *Rest) SendReqBinaryBodyAsync(arg1 string, arg2 string, arg3 []byte, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    hTask := C.CkRest_SendReqBinaryBodyAsync(z.ckObj, cstr1, cstr2, ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an application/x-www-form-urlencoded HTTP request where the body is
// composed of the URL encoded query params. The httpVerb is the HTTP verb (also known
// as the HTTP method), such as "POST". The uriPath is the path of the resource URI.
// If the Content-Type header was set, it is ignored and instead the Content-Type
// of the request will be "application/x-www-form-urlencoded".
func (z *Rest) SendReqFormUrlEncoded(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_SendReqFormUrlEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqFormUrlEncoded method
func (z *Rest) SendReqFormUrlEncodedAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_SendReqFormUrlEncodedAsync(z.ckObj, cstr1, cstr2)

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


// Sends a multipart REST request. The httpVerb is the HTTP verb (also known as the
// HTTP method), such as "GET". The uriPath is the path of the resource URI.
func (z *Rest) SendReqMultipart(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_SendReqMultipart(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqMultipart method
func (z *Rest) SendReqMultipartAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_SendReqMultipartAsync(z.ckObj, cstr1, cstr2)

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


// Sends a REST request that cotnains no body. The httpVerb is the HTTP verb (also
// known as the HTTP method), such as "GET". The uriPath is the path of the resource
// URI.
func (z *Rest) SendReqNoBody(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_SendReqNoBody(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqNoBody method
func (z *Rest) SendReqNoBodyAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_SendReqNoBodyAsync(z.ckObj, cstr1, cstr2)

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


// Sends a REST request that cotnains a text body, such as XML or JSON. The httpVerb is
// the HTTP verb (also known as the HTTP method), such as "PUT". The uriPath is the
// path of the resource URI. The bodySb contains the text of the HTTP request body.
func (z *Rest) SendReqSb(arg1 string, arg2 string, arg3 *StringBuilder) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_SendReqSb(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqSb method
func (z *Rest) SendReqSbAsync(arg1 string, arg2 string, arg3 *StringBuilder, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_SendReqSbAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// Sends a REST request that cotnains a binary or text body that is obtained by
// reading from the stream. The httpVerb is the HTTP verb (also known as the HTTP
// method), such as "PUT". The uriPath is the path of the resource URI.
func (z *Rest) SendReqStreamBody(arg1 string, arg2 string, arg3 *Stream) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_SendReqStreamBody(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqStreamBody method
func (z *Rest) SendReqStreamBodyAsync(arg1 string, arg2 string, arg3 *Stream, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkRest_SendReqStreamBodyAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// Sends a REST request that cotnains a text body, such as XML or JSON. The httpVerb is
// the HTTP verb (also known as the HTTP method), such as "PUT". The uriPath is the
// path of the resource URI. The bodyText contains the text of the HTTP request body.
func (z *Rest) SendReqStringBody(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkRest_SendReqStringBody(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the SendReqStringBody method
func (z *Rest) SendReqStringBodyAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkRest_SendReqStringBodyAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Sets the authorization service provider for Amazon AWS REST requests. An
// application that sets an AWS authentication provider need not explicitly set the
// Authorization property. Each REST request is automatically signed and
// authenticated using the authProvider.
func (z *Rest) SetAuthAws(arg1 *AuthAws) bool {

    v := C.CkRest_SetAuthAws(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the Azure AD (Active Directory) authorization service provider.
func (z *Rest) SetAuthAzureAD(arg1 *AuthAzureAD) bool {

    v := C.CkRest_SetAuthAzureAD(z.ckObj, arg1.ckObj)


    return v != 0
}


// Provides the information for Azure SAS (Shared Access Signature) authorization.
// Calling this method will cause the "Authorization: SharedAccessSignature ..."
// header to be automatically added to all requests.
func (z *Rest) SetAuthAzureSas(arg1 *AuthAzureSAS) bool {

    v := C.CkRest_SetAuthAzureSas(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the authorization service provider for Azure Storage Service requests.
func (z *Rest) SetAuthAzureStorage(arg1 *AuthAzureStorage) bool {

    v := C.CkRest_SetAuthAzureStorage(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the username and password to be used for Basic authentication. This method should
// be called when Basic authentication is required. It should only be used with
// secure SSL/TLS connections. Calling this method will cause the "Authorization:
// Basic ..." header to be automatically added to all requests. In many cases, a
// REST API will support Basic authentication where the username is a client ID or
// account ID, and the password is a client secret or token.
func (z *Rest) SetAuthBasic(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkRest_SetAuthBasic(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// The same as SetAuthBasic, but provides a more secure means for passing the
// arguments as secure string objects.
func (z *Rest) SetAuthBasicSecure(arg1 *SecureString, arg2 *SecureString) bool {

    v := C.CkRest_SetAuthBasicSecure(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Sets the authorization service provider for Google API requests.
func (z *Rest) SetAuthGoogle(arg1 *AuthGoogle) bool {

    v := C.CkRest_SetAuthGoogle(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the authentication provider for REST API requests needing OAuth 1.0 (and
// OAuth 1.0a) authentication. If useQueryParams is true, then the OAuth1 authentication
// information and signature is passed in query parameters. Otherwise it is passed
// in an Authorization header.
func (z *Rest) SetAuthOAuth1(arg1 *OAuth1, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkRest_SetAuthOAuth1(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Sets the authentication provider for REST API requests needing standards-based
// OAuth 2.0 authentication. This is for the case where a desktop/native/mobile
// application will be popping up a web browser, or embedding a web browser, to get
// authorization interactively from end-user of the application.
func (z *Rest) SetAuthOAuth2(arg1 *OAuth2) bool {

    v := C.CkRest_SetAuthOAuth2(z.ckObj, arg1.ckObj)


    return v != 0
}


// Only used for multipart requests. Sets the binary content of the multipart body
// indicated by the PartSelector.
func (z *Rest) SetMultipartBodyBd(arg1 *BinData) bool {

    v := C.CkRest_SetMultipartBodyBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Only used for multipart requests. Sets the binary content of the multipart body
// indicated by the PartSelector.
func (z *Rest) SetMultipartBodyBinary(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkRest_SetMultipartBodyBinary(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Only used for multipart requests. Sets the text content of the multipart body
// indicated by the PartSelector.
func (z *Rest) SetMultipartBodySb(arg1 *StringBuilder) bool {

    v := C.CkRest_SetMultipartBodySb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Only used for multipart requests. Sets the stream source of the multipart body
// indicated by the PartSelector.
func (z *Rest) SetMultipartBodyStream(arg1 *Stream) bool {

    v := C.CkRest_SetMultipartBodyStream(z.ckObj, arg1.ckObj)


    return v != 0
}


// Only used for multipart requests. Sets the text content of the multipart body
// indicated by the PartSelector.
func (z *Rest) SetMultipartBodyString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRest_SetMultipartBodyString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Only applies to the Full* methods, such as FullRequestNoBody, FullRequestBinary,
// FullRequestStream, etc. When set, the response body is streamed directly to
// responseStream, if (and only if) the HTTP response status code equals expectedStatus.
// 
// If the response body was streamed to responseStream, then the string return value of the
// Full* method instead becomes "OK" for success. If an attempt was made to stream
// the response body but it failed, then "FAIL" is returned. If the response body
// was not streamed because the response status code was not equal to expectedStatus, then
// the returned string contains the server's error response.
// 
// If autoSetStreamCharset is true, then the expectedStatus's StringCharset property will automatically get
// set based on the charset, if any, indicated in the response header. If the
// response is binary, then autoSetStreamCharset is ignored.
// 
// Starting in v9.5.0.75, the expectedStatus may be passed as a negative number to specify a
// range of expected (success) status codes. For example:
//     -200: Allow status codes 200 - 299
//     -210: Allow status codes 210 - 219
//     -220: Allow status codes 220 - 229
//     etc.
//
func (z *Rest) SetResponseBodyStream(arg1 int, arg2 bool, arg3 *Stream) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkRest_SetResponseBodyStream(z.ckObj, C.int(arg1), b2, arg3.ckObj)


    return v != 0
}


// Sets the connection to be used for sending the REST request. The connection should be
// an already-connected socket. It may be a TLS connection, an unencrypted
// connection, through an HTTP proxy, a SOCKS proxy, or even through SSH tunnels.
// All of the connection related functionality is handled by the Chilkat Socket
// API.
// 
// The autoReconnect indicates whether connection should automatically be established as
// needed for subsequent REST requests.
// 
// Important: The UseConnection method is provided as a means for handling more
// complicated connections -- such as connections through proxies, tunnels, etc. If
// your application is connecting directly to the HTTP server, then simply call
// this class's Connect method.
//
func (z *Rest) UseConnection(arg1 *Socket, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkRest_UseConnection(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


