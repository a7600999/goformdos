// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkImap.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewImap() *Imap {
	return &Imap{C.CkImap_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Imap) DisposeImap() {
    if z != nil {
	C.CkImap_Dispose(z.ckObj)
	}
}




func (z *Imap) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkImap_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Imap) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkImap_setExternalProgress(z.ckObj,1)
}

func (z *Imap) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkImap_setExternalProgress(z.ckObj,1)
}

func (z *Imap) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkImap_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Imap) AbortCurrent() bool {
    v := int(C.CkImap_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Imap) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putAbortCurrent(z.ckObj,v)
}

// property: AppendSeen
// When true (the default) the Append method will mark the email appended to a
// mailbox as already seen. Otherwise an appended email will be initialized to have
// a status of unseen.
func (z *Imap) AppendSeen() bool {
    v := int(C.CkImap_getAppendSeen(z.ckObj))
    return v != 0
}
// property setter: AppendSeen
func (z *Imap) SetAppendSeen(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putAppendSeen(z.ckObj,v)
}

// property: AppendUid
// The UID of the last email appended to a mailbox via an Append* method. (Not all
// IMAP servers report back the UID of the email appended.)
func (z *Imap) AppendUid() int {
    return int(C.CkImap_getAppendUid(z.ckObj))
}

// property: AuthMethod
// Can be set to "XOAUTH2", "CRAM-MD5", "NTLM", "PLAIN", or "LOGIN" to select the
// authentication method. NTLM is the most secure, and is a synonym for "Windows
// Integrated Authentication". The default is "LOGIN" (or the empty string) which
// is simple plain-text username/password authentication. Not all IMAP servers
// support all authentication methods.
// 
// The XOAUTH2 method was added in version 9.5.0.44.
// 
// Note: If SPA (i.e. NTLM) authentication does not succeed, set the
// Global.DefaultNtlmVersion property equal to 1 and then retry.
//
func (z *Imap) AuthMethod() string {
    return C.GoString(C.CkImap_authMethod(z.ckObj))
}
// property setter: AuthMethod
func (z *Imap) SetAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AuthzId
// Applies to the PLAIN authentication method. May be set to an authorization ID
// that is to be sent along with the Login and Password for authentication.
func (z *Imap) AuthzId() string {
    return C.GoString(C.CkImap_authzId(z.ckObj))
}
// property setter: AuthzId
func (z *Imap) SetAuthzId(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putAuthzId(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AutoDownloadAttachments
// If set to true, then all Fetch* methods will also automatically download
// attachments. If set to false, then the Fetch* methods download the email
// without attachments. The default value is true.
// 
// Note: Methods that download headers-only, such as FetchSingleHeader, ignore this
// property and never download attachments. Also, signed and/or encrypted emails
// will always be downloaded in full (with attachments) regardless of this property
// setting.
//
func (z *Imap) AutoDownloadAttachments() bool {
    v := int(C.CkImap_getAutoDownloadAttachments(z.ckObj))
    return v != 0
}
// property setter: AutoDownloadAttachments
func (z *Imap) SetAutoDownloadAttachments(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putAutoDownloadAttachments(z.ckObj,v)
}

// property: AutoFix
// If true, then the following will occur when a connection is made to an IMAP
// server:
// 
// 1) If the Port property = 993, then sets StartTls = false and Ssl = true
// 2) If the Port property = 143, sets Ssl = false
// 
// The default value of this property is true.
//
func (z *Imap) AutoFix() bool {
    v := int(C.CkImap_getAutoFix(z.ckObj))
    return v != 0
}
// property setter: AutoFix
func (z *Imap) SetAutoFix(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putAutoFix(z.ckObj,v)
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
func (z *Imap) ClientIpAddress() string {
    return C.GoString(C.CkImap_clientIpAddress(z.ckObj))
}
// property setter: ClientIpAddress
func (z *Imap) SetClientIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putClientIpAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConnectedToHost
// Contains the IMAP server's domain name (or IP address) if currently connected.
// Otherwise returns an empty string.
func (z *Imap) ConnectedToHost() string {
    return C.GoString(C.CkImap_connectedToHost(z.ckObj))
}

// property: ConnectTimeout
// Maximum number of seconds to wait when connecting to an IMAP server. The default
// value is 30 (units are in seconds).
func (z *Imap) ConnectTimeout() int {
    return int(C.CkImap_getConnectTimeout(z.ckObj))
}
// property setter: ConnectTimeout
func (z *Imap) SetConnectTimeout(value int) {
    C.CkImap_putConnectTimeout(z.ckObj,C.int(value))
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
func (z *Imap) DebugLogFilePath() string {
    return C.GoString(C.CkImap_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Imap) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Domain
// The Windows Domain to use for Windows Integrated Authentication (also known as
// NTLM). This may be empty.
func (z *Imap) Domain() string {
    return C.GoString(C.CkImap_domain(z.ckObj))
}
// property setter: Domain
func (z *Imap) SetDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// This is the number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any IMAP operation prior to
// completion. If HeartbeatMs is 0, no AbortCheck event callbacks will occur.
func (z *Imap) HeartbeatMs() int {
    return int(C.CkImap_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Imap) SetHeartbeatMs(value int) {
    C.CkImap_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: HttpProxyAuthMethod
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy authentication method name. Valid choices are "Basic" or "NTLM".
func (z *Imap) HttpProxyAuthMethod() string {
    return C.GoString(C.CkImap_httpProxyAuthMethod(z.ckObj))
}
// property setter: HttpProxyAuthMethod
func (z *Imap) SetHttpProxyAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putHttpProxyAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyDomain
// The NTLM authentication domain (optional) if NTLM authentication is used.
func (z *Imap) HttpProxyDomain() string {
    return C.GoString(C.CkImap_httpProxyDomain(z.ckObj))
}
// property setter: HttpProxyDomain
func (z *Imap) SetHttpProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putHttpProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyHostname
// If an HTTP proxy is to be used, set this property to the HTTP proxy hostname or
// IPv4 address (in dotted decimal notation).
func (z *Imap) HttpProxyHostname() string {
    return C.GoString(C.CkImap_httpProxyHostname(z.ckObj))
}
// property setter: HttpProxyHostname
func (z *Imap) SetHttpProxyHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putHttpProxyHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPassword
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy password.
func (z *Imap) HttpProxyPassword() string {
    return C.GoString(C.CkImap_httpProxyPassword(z.ckObj))
}
// property setter: HttpProxyPassword
func (z *Imap) SetHttpProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putHttpProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPort
// If an HTTP proxy is to be used, set this property to the HTTP proxy port number.
// (Two commonly used HTTP proxy ports are 8080 and 3128.)
func (z *Imap) HttpProxyPort() int {
    return int(C.CkImap_getHttpProxyPort(z.ckObj))
}
// property setter: HttpProxyPort
func (z *Imap) SetHttpProxyPort(value int) {
    C.CkImap_putHttpProxyPort(z.ckObj,C.int(value))
}

// property: HttpProxyUsername
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy login name.
func (z *Imap) HttpProxyUsername() string {
    return C.GoString(C.CkImap_httpProxyUsername(z.ckObj))
}
// property setter: HttpProxyUsername
func (z *Imap) SetHttpProxyUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putHttpProxyUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: KeepSessionLog
// Turns the in-memory session logging on or off. If on, the session log can be
// obtained via the SessionLog property. The default value is false.
// 
// The SessionLog contains the raw commands sent to the IMAP server, and the raw
// responses received from the IMAP server.
//
func (z *Imap) KeepSessionLog() bool {
    v := int(C.CkImap_getKeepSessionLog(z.ckObj))
    return v != 0
}
// property setter: KeepSessionLog
func (z *Imap) SetKeepSessionLog(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putKeepSessionLog(z.ckObj,v)
}

// property: LastAppendedMime
// The MIME source of the email last appended during a call to AppendMail, or
// AppendMime.
func (z *Imap) LastAppendedMime() string {
    return C.GoString(C.CkImap_lastAppendedMime(z.ckObj))
}

// property: LastCommand
// The last raw command sent to the IMAP server. (This information can be used for
// debugging if problems occur.)
func (z *Imap) LastCommand() string {
    return C.GoString(C.CkImap_lastCommand(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Imap) LastErrorHtml() string {
    return C.GoString(C.CkImap_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Imap) LastErrorText() string {
    return C.GoString(C.CkImap_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Imap) LastErrorXml() string {
    return C.GoString(C.CkImap_lastErrorXml(z.ckObj))
}

// property: LastIntermediateResponse
// The last intermediate response received from the IMAP server. (This information
// can be used for debugging if problems occur.)
func (z *Imap) LastIntermediateResponse() string {
    return C.GoString(C.CkImap_lastIntermediateResponse(z.ckObj))
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
func (z *Imap) LastMethodSuccess() bool {
    v := int(C.CkImap_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Imap) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putLastMethodSuccess(z.ckObj,v)
}

// property: LastResponse
// The raw data of the last response from the IMAP server. (Useful for debugging if
// problems occur.) This property is cleared whenever a command is sent to the IMAP
// server. If no response is received, then this property will remain empty.
// Otherwise, it will contain the last response received from the IMAP server.
func (z *Imap) LastResponse() string {
    return C.GoString(C.CkImap_lastResponse(z.ckObj))
}

// property: LastResponseCode
// The response code part of the last command response, if it exists. IMAP status
// responses MAY include an OPTIONAL "response code". A response code consists of
// data inside square brackets in the form of an atom, possibly followed by a space
// and arguments. The response code contains additional information or status codes
// for client software beyond the OK/NO/BAD condition, and are defined when there
// is a specific action that a client can take based upon the additional
// information. Examples of response codes are "NONEXISTENT" and
// "AUTHENTICATIONFAILED". The response code strings for a given failure condition
// may vary depending on the IMAP server implementation.
func (z *Imap) LastResponseCode() string {
    return C.GoString(C.CkImap_lastResponseCode(z.ckObj))
}

// property: LoggedInUser
// If logged into an IMAP server, the logged-in username.
func (z *Imap) LoggedInUser() string {
    return C.GoString(C.CkImap_loggedInUser(z.ckObj))
}

// property: NumMessages
// After selecting a mailbox (by calling SelectMailbox), this property will be
// updated to reflect the total number of emails in the mailbox.
func (z *Imap) NumMessages() int {
    return int(C.CkImap_getNumMessages(z.ckObj))
}

// property: PeekMode
// Set to true to prevent the mail flags (such as the "Seen" flag) from being set
// when email is retrieved. The default value of this property is false.
func (z *Imap) PeekMode() bool {
    v := int(C.CkImap_getPeekMode(z.ckObj))
    return v != 0
}
// property setter: PeekMode
func (z *Imap) SetPeekMode(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putPeekMode(z.ckObj,v)
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
func (z *Imap) PercentDoneScale() int {
    return int(C.CkImap_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *Imap) SetPercentDoneScale(value int) {
    C.CkImap_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: Port
// The IMAP port number. If using SSL, be sure to set this to the IMAP SSL port
// number, which is typically port 993. (If this is the case, make sure you also
// set the Ssl property = true.
func (z *Imap) Port() int {
    return int(C.CkImap_getPort(z.ckObj))
}
// property setter: Port
func (z *Imap) SetPort(value int) {
    C.CkImap_putPort(z.ckObj,C.int(value))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *Imap) PreferIpv6() bool {
    v := int(C.CkImap_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *Imap) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putPreferIpv6(z.ckObj,v)
}

// property: ReadTimeout
// The maximum amount of time (in seconds) that incoming data is allowed to stall
// while reading any kind of response from an IMAP server. This is the amount of
// time that needs to elapse while no additional response bytes are forthcoming.
// For the case of long responses, if the data stream halts for more than this
// amount, it will timeout. This property is not a maximum for the total response
// time, but only a maximum for the amount of time while no response arrives.
// 
// The default value is 30 seconds.
//
func (z *Imap) ReadTimeout() int {
    return int(C.CkImap_getReadTimeout(z.ckObj))
}
// property setter: ReadTimeout
func (z *Imap) SetReadTimeout(value int) {
    C.CkImap_putReadTimeout(z.ckObj,C.int(value))
}

// property: RequireSslCertVerify
// If true, then the IMAP client will verify the server's SSL certificate. The
// certificate is expired, or if the cert's signature is invalid, the connection is
// not allowed. The default value of this property is false.
func (z *Imap) RequireSslCertVerify() bool {
    v := int(C.CkImap_getRequireSslCertVerify(z.ckObj))
    return v != 0
}
// property setter: RequireSslCertVerify
func (z *Imap) SetRequireSslCertVerify(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putRequireSslCertVerify(z.ckObj,v)
}

// property: SearchCharset
// The "CHARSET" to be used in searches issued by the Search method. The default
// value is "UTF-8". (If no 8bit chars are found in the search criteria passed to
// the Search method, then no CHARSET is needed and this property doesn't apply.)
// The SearchCharset property can be set to "AUTO" to get the pre-v9.4.0 behavior,
// which is to examine the 8bit chars found in the search criteria and select an
// appropriate multibyte charset.
// 
// In summary, it is unlikely that this property needs to be changed. It should
// only be modified if trouble arises with some IMAP servers when non-English chars
// are used in the search criteria.
//
func (z *Imap) SearchCharset() string {
    return C.GoString(C.CkImap_searchCharset(z.ckObj))
}
// property setter: SearchCharset
func (z *Imap) SetSearchCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putSearchCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SelectedMailbox
// The currently selected mailbox, or an empty string if none.
func (z *Imap) SelectedMailbox() string {
    return C.GoString(C.CkImap_selectedMailbox(z.ckObj))
}

// property: SendBufferSize
// The buffer size to be used with the underlying TCP/IP socket for sending. The
// default value is 32767.
func (z *Imap) SendBufferSize() int {
    return int(C.CkImap_getSendBufferSize(z.ckObj))
}
// property setter: SendBufferSize
func (z *Imap) SetSendBufferSize(value int) {
    C.CkImap_putSendBufferSize(z.ckObj,C.int(value))
}

// property: SeparatorChar
// The separator character used by the IMAP server for the mailbox hierarchy. It is
// typically "/" or ".", but may vary depending on the IMAP server. The
// ListMailboxes method has the side-effect of setting this property to the correct
// value because the IMAP server's response when listing mailboxes includes
// information about the separator char.
// 
// Note: Starting in version 9.5.0.47, this property changed from a "char" type to
// a "string" type. The separator char property will always be a string of length 1
// character.
//
func (z *Imap) SeparatorChar() string {
    return C.GoString(C.CkImap_separatorChar(z.ckObj))
}
// property setter: SeparatorChar
func (z *Imap) SetSeparatorChar(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putSeparatorChar(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SessionLog
// Contains an in-memory log of the raw commands sent to the IMAP server, and the
// raw responses received from the IMAP server. The KeepSessionLog property must be
// set to true to enable session logging. Call ClearSessionLog to reset the log.
func (z *Imap) SessionLog() string {
    return C.GoString(C.CkImap_sessionLog(z.ckObj))
}

// property: SocksHostname
// The SOCKS4/SOCKS5 hostname or IPv4 address (in dotted decimal notation). This
// property is only used if the SocksVersion property is set to 4 or 5).
func (z *Imap) SocksHostname() string {
    return C.GoString(C.CkImap_socksHostname(z.ckObj))
}
// property setter: SocksHostname
func (z *Imap) SetSocksHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putSocksHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPassword
// The SOCKS5 password (if required). The SOCKS4 protocol does not include the use
// of a password, so this does not apply to SOCKS4.
func (z *Imap) SocksPassword() string {
    return C.GoString(C.CkImap_socksPassword(z.ckObj))
}
// property setter: SocksPassword
func (z *Imap) SetSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPort
// The SOCKS4/SOCKS5 proxy port. The default value is 1080. This property only
// applies if a SOCKS proxy is used (if the SocksVersion property is set to 4 or
// 5).
func (z *Imap) SocksPort() int {
    return int(C.CkImap_getSocksPort(z.ckObj))
}
// property setter: SocksPort
func (z *Imap) SetSocksPort(value int) {
    C.CkImap_putSocksPort(z.ckObj,C.int(value))
}

// property: SocksUsername
// The SOCKS4/SOCKS5 proxy username. This property is only used if the SocksVersion
// property is set to 4 or 5).
func (z *Imap) SocksUsername() string {
    return C.GoString(C.CkImap_socksUsername(z.ckObj))
}
// property setter: SocksUsername
func (z *Imap) SetSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksVersion
// SocksVersion May be set to one of the following integer values:
// 
// 0 - No SOCKS proxy is used. This is the default.
// 4 - Connect via a SOCKS4 proxy.
// 5 - Connect via a SOCKS5 proxy.
//
func (z *Imap) SocksVersion() int {
    return int(C.CkImap_getSocksVersion(z.ckObj))
}
// property setter: SocksVersion
func (z *Imap) SetSocksVersion(value int) {
    C.CkImap_putSocksVersion(z.ckObj,C.int(value))
}

// property: SoRcvBuf
// Sets the receive buffer size socket option. Normally, this property should be
// left unchanged. The default value is 4194304.
// 
// This property can be increased if download performance seems slow. It is
// recommended to be a multiple of 4096.
//
func (z *Imap) SoRcvBuf() int {
    return int(C.CkImap_getSoRcvBuf(z.ckObj))
}
// property setter: SoRcvBuf
func (z *Imap) SetSoRcvBuf(value int) {
    C.CkImap_putSoRcvBuf(z.ckObj,C.int(value))
}

// property: SoSndBuf
// Sets the send buffer size socket option. Normally, this property should be left
// unchanged. The default value is 262144.
// 
// This property can be increased if upload performance seems slow. It is
// recommended to be a multiple of 4096. Testing with sizes such as 512K and 1MB is
// reasonable.
//
func (z *Imap) SoSndBuf() int {
    return int(C.CkImap_getSoSndBuf(z.ckObj))
}
// property setter: SoSndBuf
func (z *Imap) SetSoSndBuf(value int) {
    C.CkImap_putSoSndBuf(z.ckObj,C.int(value))
}

// property: Ssl
// true if the IMAP connection should be TLS/SSL.
// 
// Note: The typical IMAP TLS/SSL port number is 993. If you set this property =
// true, it is likely that you should also set the Port property = 993.
//
func (z *Imap) Ssl() bool {
    v := int(C.CkImap_getSsl(z.ckObj))
    return v != 0
}
// property setter: Ssl
func (z *Imap) SetSsl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putSsl(z.ckObj,v)
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
func (z *Imap) SslAllowedCiphers() string {
    return C.GoString(C.CkImap_sslAllowedCiphers(z.ckObj))
}
// property setter: SslAllowedCiphers
func (z *Imap) SetSslAllowedCiphers(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putSslAllowedCiphers(z.ckObj,cStr)
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
func (z *Imap) SslProtocol() string {
    return C.GoString(C.CkImap_sslProtocol(z.ckObj))
}
// property setter: SslProtocol
func (z *Imap) SetSslProtocol(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putSslProtocol(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SslServerCertVerified
// Read-only property that returns true if the IMAP server's digital certificate
// was verified when connecting via SSL / TLS.
func (z *Imap) SslServerCertVerified() bool {
    v := int(C.CkImap_getSslServerCertVerified(z.ckObj))
    return v != 0
}

// property: StartTls
// If true, then the Connect method will (internallly) convert the connection to
// TLS/SSL via the STARTTLS IMAP command. This is called "explict SSL/TLS" because
// the client explicitly requests the connection be transformed into a TLS/SSL
// secure channel. The alternative is "implicit SSL/TLS" where the "Ssl" property
// is set to true and the IMAP client connects to the well-known TLS/SSL IMAP
// port of 993.
func (z *Imap) StartTls() bool {
    v := int(C.CkImap_getStartTls(z.ckObj))
    return v != 0
}
// property setter: StartTls
func (z *Imap) SetStartTls(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putStartTls(z.ckObj,v)
}

// property: TlsCipherSuite
// Contains the current or last negotiated TLS cipher suite. If no TLS connection
// has yet to be established, or if a connection as attempted and failed, then this
// will be empty. A sample cipher suite string looks like this:
// TLS_DHE_RSA_WITH_AES_256_CBC_SHA256.
func (z *Imap) TlsCipherSuite() string {
    return C.GoString(C.CkImap_tlsCipherSuite(z.ckObj))
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
func (z *Imap) TlsPinSet() string {
    return C.GoString(C.CkImap_tlsPinSet(z.ckObj))
}
// property setter: TlsPinSet
func (z *Imap) SetTlsPinSet(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putTlsPinSet(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TlsVersion
// Contains the current or last negotiated TLS protocol version. If no TLS
// connection has yet to be established, or if a connection as attempted and
// failed, then this will be empty. Possible values are "SSL 3.0", "TLS 1.0", "TLS
// 1.1", and "TLS 1.2".
func (z *Imap) TlsVersion() string {
    return C.GoString(C.CkImap_tlsVersion(z.ckObj))
}

// property: UidNext
// A positive integer value containing the UIDNEXT of the currently selected
// folder, or 0 if it's not available or no folder is selected.
func (z *Imap) UidNext() int {
    return int(C.CkImap_getUidNext(z.ckObj))
}

// property: UidValidity
// An integer value containing the UIDVALIDITY of the currently selected mailbox,
// or 0 if no mailbox is selected.
// 
// A client can save the UidValidity value for a mailbox and then compare it with
// the UidValidity on a subsequent session. If the new value is larger, the IMAP
// server is not keeping UID's unchanged between sessions. Most IMAP servers
// maintain UID's between sessions.
//
func (z *Imap) UidValidity() int {
    return int(C.CkImap_getUidValidity(z.ckObj))
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. This property
// defaults to the empty string and should typically remain empty. Can be set to a
// list of the following comma separated keywords:
//     "ProtectFromVpn" - Introduced in v9.5.0.80. On Android systems, will bypass
//     any VPN that may be installed or active.
//     "EnableTls13" - Introduced in v9.5.0.82. Causes TLS 1.3 to be offered in the
//     ClientHello of the TLS protocol, allowing the server to select TLS 1.3 for the
//     session. Future versions of Chilkat will enable TLS 1.3 by default. This option
//     is only necessary in v9.5.0.82 if TLS 1.3 is desired.
func (z *Imap) UncommonOptions() string {
    return C.GoString(C.CkImap_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Imap) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkImap_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Imap) VerboseLogging() bool {
    v := int(C.CkImap_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Imap) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkImap_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Imap) Version() string {
    return C.GoString(C.CkImap_version(z.ckObj))
}
// Returns true if the underlying TCP socket is connected to the IMAP server.
func (z *Imap) AddPfxSourceData(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkImap_AddPfxSourceData(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a PFX file to the object's internal list of sources to be searched for
// certificates and private keys when decrypting. Multiple PFX files can be added
// by calling this method once for each. (On the Windows operating system, the
// registry-based certificate stores are also automatically searched, so it is
// commonly not required to explicitly add PFX sources.)
// 
// The pfxFilePath contains the bytes of a PFX file (also known as PKCS12 or .p12).
//
func (z *Imap) AddPfxSourceFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkImap_AddPfxSourceFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends an email to an IMAP mailbox.
func (z *Imap) AppendMail(arg1 string, arg2 *Email) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_AppendMail(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the AppendMail method
func (z *Imap) AppendMailAsync(arg1 string, arg2 *Email, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_AppendMailAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Appends an email (represented as MIME text) to an IMAP mailbox.
func (z *Imap) AppendMime(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkImap_AppendMime(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AppendMime method
func (z *Imap) AppendMimeAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_AppendMimeAsync(z.ckObj, cstr1, cstr2)

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


// The same as AppendMimeWithDate, except the date/time is provided in RFC822
// string format, such as "Wed, 18 Oct 2017 09:08:21 GMT".
func (z *Imap) AppendMimeWithDateStr(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkImap_AppendMimeWithDateStr(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the AppendMimeWithDateStr method
func (z *Imap) AppendMimeWithDateStrAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkImap_AppendMimeWithDateStrAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Same as AppendMime, but allows the flags associated with the email to be set at
// the same time. A flag is on if true, and off if false.
func (z *Imap) AppendMimeWithFlags(arg1 string, arg2 string, arg3 bool, arg4 bool, arg5 bool, arg6 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
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

    v := C.CkImap_AppendMimeWithFlags(z.ckObj, cstr1, cstr2, b3, b4, b5, b6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AppendMimeWithFlags method
func (z *Imap) AppendMimeWithFlagsAsync(arg1 string, arg2 string, arg3 bool, arg4 bool, arg5 bool, arg6 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
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

    hTask := C.CkImap_AppendMimeWithFlagsAsync(z.ckObj, cstr1, cstr2, b3, b4, b5, b6)

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


// Same as AppendMimeWithFlags, but the MIME to be uploaded to the IMAP server is
// passed in a StringBuilder object.
func (z *Imap) AppendMimeWithFlagsSb(arg1 string, arg2 *StringBuilder, arg3 bool, arg4 bool, arg5 bool, arg6 bool) bool {
    cstr1 := C.CString(arg1)
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

    v := C.CkImap_AppendMimeWithFlagsSb(z.ckObj, cstr1, arg2.ckObj, b3, b4, b5, b6)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the AppendMimeWithFlagsSb method
func (z *Imap) AppendMimeWithFlagsSbAsync(arg1 string, arg2 *StringBuilder, arg3 bool, arg4 bool, arg5 bool, arg6 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
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

    hTask := C.CkImap_AppendMimeWithFlagsSbAsync(z.ckObj, cstr1, arg2.ckObj, b3, b4, b5, b6)

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


// Sends a CAPABILITY command to the IMAP server and returns the raw response.
// return a string or nil if failed.
func (z *Imap) Capability() *string {

    retStr := C.CkImap_capability(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the Capability method
func (z *Imap) CapabilityAsync(c chan *Task) {

    hTask := C.CkImap_CapabilityAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns true if the underlying TCP socket is connected to the IMAP server.
// 
// Internally, this method makes a lower-level socket system call to check if the
// TCP socket is still connected.
//
func (z *Imap) CheckConnection() bool {

    v := C.CkImap_CheckConnection(z.ckObj)


    return v != 0
}


// Checks for new email that has arrived since the mailbox was selected (via the
// SelectMailbox or ExamineMailbox methods), or since the last call to
// CheckForNewEmail (whichever was most recent). This method works by closing and
// re-opening the currently selected mailbox, and then sending a "SEARCH" command
// for either RECENT emails, or emails having a UID greater than the UIDNEXT value.
// A message set object containing the UID's of the new emails is returned, and
// this may be passed to methods such as FetchBundle to download the new emails.
func (z *Imap) CheckForNewEmail() *MessageSet {

    retObj := C.CkImap_CheckForNewEmail(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &MessageSet{retObj}
}


// Asynchronous version of the CheckForNewEmail method
func (z *Imap) CheckForNewEmailAsync(c chan *Task) {

    hTask := C.CkImap_CheckForNewEmailAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Clears the contents of the SessionLog property.
func (z *Imap) ClearSessionLog()  {

    C.CkImap_ClearSessionLog(z.ckObj)



}


// Closes the currently selected mailbox.
// 
// Note: The term "mailbox" and "folder" are synonymous. Whenever the word
// "mailbox" is used, it has the same meaning as "folder".
//
func (z *Imap) CloseMailbox(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_CloseMailbox(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the CloseMailbox method
func (z *Imap) CloseMailboxAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_CloseMailboxAsync(z.ckObj, cstr1)

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


// Connects to an IMAP server, but does not login. The domainName is the domain name of
// the IMAP server. (May also use the IPv4 or IPv6 address in string format.)
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
//
func (z *Imap) Connect(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_Connect(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Connect method
func (z *Imap) ConnectAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_ConnectAsync(z.ckObj, cstr1)

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


// Copies a message from the selected mailbox to copyToMailbox. If bUid is true, then msgId
// represents a UID. If bUid is false, then msgId represents a sequence number.
func (z *Imap) Copy(arg1 int, arg2 bool, arg3 string) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    v := C.CkImap_Copy(z.ckObj, C.int(arg1), b2, cstr3)

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the Copy method
func (z *Imap) CopyAsync(arg1 int, arg2 bool, arg3 string, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    hTask := C.CkImap_CopyAsync(z.ckObj, C.int(arg1), b2, cstr3)

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


// Same as the Copy method, except an entire set of emails is copied at once. The
// set of emails is specified in messageSet.
func (z *Imap) CopyMultiple(arg1 *MessageSet, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkImap_CopyMultiple(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CopyMultiple method
func (z *Imap) CopyMultipleAsync(arg1 *MessageSet, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_CopyMultipleAsync(z.ckObj, arg1.ckObj, cstr2)

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


// Copies one or more emails from one mailbox to another. The emails are specified
// as a range of sequence numbers. The 1st email in a mailbox is always at sequence
// number 1.
func (z *Imap) CopySequence(arg1 int, arg2 int, arg3 string) bool {
    cstr3 := C.CString(arg3)

    v := C.CkImap_CopySequence(z.ckObj, C.int(arg1), C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the CopySequence method
func (z *Imap) CopySequenceAsync(arg1 int, arg2 int, arg3 string, c chan *Task) {
    cstr3 := C.CString(arg3)

    hTask := C.CkImap_CopySequenceAsync(z.ckObj, C.int(arg1), C.int(arg2), cstr3)

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


// Creates a new mailbox.
// 
// Note: The term "mailbox" and "folder" are synonymous. Whenever the word
// "mailbox" is used, it has the same meaning as "folder".
//
func (z *Imap) CreateMailbox(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_CreateMailbox(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the CreateMailbox method
func (z *Imap) CreateMailboxAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_CreateMailboxAsync(z.ckObj, cstr1)

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


// Deletes an existing mailbox.
// 
// Note: The term "mailbox" and "folder" are synonymous. Whenever the word
// "mailbox" is used, it has the same meaning as "folder".
//
func (z *Imap) DeleteMailbox(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_DeleteMailbox(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the DeleteMailbox method
func (z *Imap) DeleteMailboxAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_DeleteMailboxAsync(z.ckObj, cstr1)

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


// Disconnects cleanly from the IMAP server. A non-success return from this method
// only indicates that the disconnect was not clean -- and this can typically be
// ignored.
func (z *Imap) Disconnect() bool {

    v := C.CkImap_Disconnect(z.ckObj)


    return v != 0
}


// Asynchronous version of the Disconnect method
func (z *Imap) DisconnectAsync(c chan *Task) {

    hTask := C.CkImap_DisconnectAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Selects a mailbox such that only read-only transactions are allowed. This method
// would be called instead of SelectMailbox if the logged-on user has read-only
// permission.
// 
// Note: The term "mailbox" and "folder" are synonymous. Whenever the word
// "mailbox" is used, it has the same meaning as "folder".
//
func (z *Imap) ExamineMailbox(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_ExamineMailbox(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the ExamineMailbox method
func (z *Imap) ExamineMailboxAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_ExamineMailboxAsync(z.ckObj, cstr1)

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


// Permanently removes from the currently selected mailbox all messages that have
// the Deleted flag set.
func (z *Imap) Expunge() bool {

    v := C.CkImap_Expunge(z.ckObj)


    return v != 0
}


// Asynchronous version of the Expunge method
func (z *Imap) ExpungeAsync(c chan *Task) {

    hTask := C.CkImap_ExpungeAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Permanently removes from the currently selected mailbox all messages that have
// the Deleted flag set, and closes the mailbox.
func (z *Imap) ExpungeAndClose() bool {

    v := C.CkImap_ExpungeAndClose(z.ckObj)


    return v != 0
}


// Asynchronous version of the ExpungeAndClose method
func (z *Imap) ExpungeAndCloseAsync(c chan *Task) {

    hTask := C.CkImap_ExpungeAndCloseAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads one of an email's attachments and saves it to a file. If the emailObject
// already contains the full email (including the attachments), then no
// communication with the IMAP server is necessary because the attachment data is
// already contained within the emailObject. In this case, the attachment is simply
// extracted and saved to saveToPath. (As with all Chilkat methods, indexing begins at 0.
// The 1st attachment is at attachmentIndex 0.)
// 
// Additional Notes:
// 
// If the AutoDownloadAttachments property is set to false, then emails
// downloaded via any of the Fetch* methods will not include attachments.
// 
// Note: "related" items are not considered attachments and are downloaded. These
// are images, style sheets, etc. that are embedded within the HTML body of an
// email.
// 
// Also: All signed and/or encrypted emails must be downloaded in full.
// 
// When an email is downloaded without attachments, the attachment information is
// included in header fields. The header fields have names beginning with
// "ckx-imap-". The attachment information can be obtained via the following
// methods:
// 
//     imap.GetMailNumAttach
//     imap.GetMailAttachFilename
//     imap.GetMailAttachSize
//     
//
func (z *Imap) FetchAttachment(arg1 *Email, arg2 int, arg3 string) bool {
    cstr3 := C.CString(arg3)

    v := C.CkImap_FetchAttachment(z.ckObj, arg1.ckObj, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the FetchAttachment method
func (z *Imap) FetchAttachmentAsync(arg1 *Email, arg2 int, arg3 string, c chan *Task) {
    cstr3 := C.CString(arg3)

    hTask := C.CkImap_FetchAttachmentAsync(z.ckObj, arg1.ckObj, C.int(arg2), cstr3)

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


// Downloads one of an email's attachments and returns the attachment data in a
// BinData object. ***See the FetchAttachment method description for more
// information about fetching attachments.
func (z *Imap) FetchAttachmentBd(arg1 *Email, arg2 int, arg3 *BinData) bool {

    v := C.CkImap_FetchAttachmentBd(z.ckObj, arg1.ckObj, C.int(arg2), arg3.ckObj)


    return v != 0
}


// Asynchronous version of the FetchAttachmentBd method
func (z *Imap) FetchAttachmentBdAsync(arg1 *Email, arg2 int, arg3 *BinData, c chan *Task) {

    hTask := C.CkImap_FetchAttachmentBdAsync(z.ckObj, arg1.ckObj, C.int(arg2), arg3.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads one of an email's attachments and returns the attachment data as
// in-memory bytes that may be accessed by an application. ***See the
// FetchAttachment method description for more information about fetching
// attachments.
func (z *Imap) FetchAttachmentBytes(arg1 *Email, arg2 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkImap_FetchAttachmentBytes(z.ckObj, arg1.ckObj, C.int(arg2), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the FetchAttachmentBytes method
func (z *Imap) FetchAttachmentBytesAsync(arg1 *Email, arg2 int, c chan *Task) {

    hTask := C.CkImap_FetchAttachmentBytesAsync(z.ckObj, arg1.ckObj, C.int(arg2))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads one of an email's attachments and returns the attachment data in a
// StringBuilder. It only makes sense to call this method for attachments that
// contain text data. The charset indicates the character encoding of the text, such
// as "utf-8" or "windows-1252". ***See the FetchAttachment method description for
// more information about fetching attachments.
func (z *Imap) FetchAttachmentSb(arg1 *Email, arg2 int, arg3 string, arg4 *StringBuilder) bool {
    cstr3 := C.CString(arg3)

    v := C.CkImap_FetchAttachmentSb(z.ckObj, arg1.ckObj, C.int(arg2), cstr3, arg4.ckObj)

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the FetchAttachmentSb method
func (z *Imap) FetchAttachmentSbAsync(arg1 *Email, arg2 int, arg3 string, arg4 *StringBuilder, c chan *Task) {
    cstr3 := C.CString(arg3)

    hTask := C.CkImap_FetchAttachmentSbAsync(z.ckObj, arg1.ckObj, C.int(arg2), cstr3, arg4.ckObj)

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


// Downloads one of an email's attachments and returns the attachment data as a
// string. It only makes sense to call this method for attachments that contain
// text data. The charset indicates the character encoding of the text, such as
// "utf-8" or "windows-1252". ***See the FetchAttachment method description for
// more information about fetching attachments.
// return a string or nil if failed.
func (z *Imap) FetchAttachmentString(arg1 *Email, arg2 int, arg3 string) *string {
    cstr3 := C.CString(arg3)

    retStr := C.CkImap_fetchAttachmentString(z.ckObj, arg1.ckObj, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FetchAttachmentString method
func (z *Imap) FetchAttachmentStringAsync(arg1 *Email, arg2 int, arg3 string, c chan *Task) {
    cstr3 := C.CString(arg3)

    hTask := C.CkImap_FetchAttachmentStringAsync(z.ckObj, arg1.ckObj, C.int(arg2), cstr3)

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


// Retrieves a set of messages from the IMAP server and returns them in an email
// bundle object. If the method fails, it may return a NULL reference.
func (z *Imap) FetchBundle(arg1 *MessageSet) *EmailBundle {

    retObj := C.CkImap_FetchBundle(z.ckObj, arg1.ckObj)


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the FetchBundle method
func (z *Imap) FetchBundleAsync(arg1 *MessageSet, c chan *Task) {

    hTask := C.CkImap_FetchBundleAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Retrieves a set of messages from the IMAP server and returns them in a string
// array object (NOTE: it does not return a string array, but an object that
// represents a string array.) Each string within the returned object is the
// complete MIME source of an email. On failure, a NULL object reference is
// returned.
func (z *Imap) FetchBundleAsMime(arg1 *MessageSet) *StringArray {

    retObj := C.CkImap_FetchBundleAsMime(z.ckObj, arg1.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Asynchronous version of the FetchBundleAsMime method
func (z *Imap) FetchBundleAsMimeAsync(arg1 *MessageSet, c chan *Task) {

    hTask := C.CkImap_FetchBundleAsMimeAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Fetches a chunk of emails starting at a specific sequence number. A bundle of
// fetched emails is returned. The last two arguments are message sets that are
// updated with the ids of messages successfully/unsuccessfully fetched.
func (z *Imap) FetchChunk(arg1 int, arg2 int, arg3 *MessageSet, arg4 *MessageSet) *EmailBundle {

    retObj := C.CkImap_FetchChunk(z.ckObj, C.int(arg1), C.int(arg2), arg3.ckObj, arg4.ckObj)


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the FetchChunk method
func (z *Imap) FetchChunkAsync(arg1 int, arg2 int, arg3 *MessageSet, arg4 *MessageSet, c chan *Task) {

    hTask := C.CkImap_FetchChunkAsync(z.ckObj, C.int(arg1), C.int(arg2), arg3.ckObj, arg4.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Fetches the flags for an email. The bUid argument determines whether the msgId is
// a UID or sequence number.
// 
// Returns the SPACE separated list of flags set for the email, such as "\Flagged
// \Seen $label1".
// 
// If an empty string is returned, then it could be that the email referenced by
// msgId does not exist in the currently selected mailbox, or it simply has no flags
// that are set. To determine the difference, examine the contents of the
// LastResponse property. For the case where the message does not exist, the
// LastResponse will contain a "NO" and will look something like this:
// aaah NO The specified message set is invalid.
// For the case where the message exists, but no flags are set, the LastResponse
// will contain an "OK" in the last response line. For example:
// ...
// aaah OK FETCH completed.
//
// return a string or nil if failed.
func (z *Imap) FetchFlags(arg1 int, arg2 bool) *string {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkImap_fetchFlags(z.ckObj, C.int(arg1), b2)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FetchFlags method
func (z *Imap) FetchFlagsAsync(arg1 int, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkImap_FetchFlagsAsync(z.ckObj, C.int(arg1), b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Retrieves a set of message headers from the IMAP server and returns them in an
// email bundle object. If the method fails, it may return a NULL reference. The
// following methods are useful for retrieving information about attachments and
// flags after email headers are retrieved: GetMailNumAttach, GetMailAttachSize,
// GetMailAttachFilename, GetMailFlag.
func (z *Imap) FetchHeaders(arg1 *MessageSet) *EmailBundle {

    retObj := C.CkImap_FetchHeaders(z.ckObj, arg1.ckObj)


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the FetchHeaders method
func (z *Imap) FetchHeadersAsync(arg1 *MessageSet, c chan *Task) {

    hTask := C.CkImap_FetchHeadersAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads email for a range of sequence numbers. The 1st email in a mailbox is
// always at sequence number 1. The total number of emails in the currently
// selected mailbox is available in the NumMessages property. If the numMessages is too
// large, the method will still succeed, but will return a bundle of emails from
// startSeqNum to the last email in the mailbox.
func (z *Imap) FetchSequence(arg1 int, arg2 int) *EmailBundle {

    retObj := C.CkImap_FetchSequence(z.ckObj, C.int(arg1), C.int(arg2))


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the FetchSequence method
func (z *Imap) FetchSequenceAsync(arg1 int, arg2 int, c chan *Task) {

    hTask := C.CkImap_FetchSequenceAsync(z.ckObj, C.int(arg1), C.int(arg2))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Same as FetchSequence, but instead of returning email objects in a bundle, the
// raw MIME of each email is returned.
func (z *Imap) FetchSequenceAsMime(arg1 int, arg2 int) *StringArray {

    retObj := C.CkImap_FetchSequenceAsMime(z.ckObj, C.int(arg1), C.int(arg2))


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Asynchronous version of the FetchSequenceAsMime method
func (z *Imap) FetchSequenceAsMimeAsync(arg1 int, arg2 int, c chan *Task) {

    hTask := C.CkImap_FetchSequenceAsMimeAsync(z.ckObj, C.int(arg1), C.int(arg2))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Same as FetchSequence, but only the email headers are returned. The email
// objects within the bundle will be lacking bodies and attachments.
// 
// Note: For any method call using sequence numbers, an application must make sure
// the sequence numbers are within the valid range. When a mailbox is selected, the
// NumMessages property will have been set, and the valid range of sequence numbers
// is from 1 to NumMessages. An attempt to fetch sequence numbers outside this
// range will result in an error.
//
func (z *Imap) FetchSequenceHeaders(arg1 int, arg2 int) *EmailBundle {

    retObj := C.CkImap_FetchSequenceHeaders(z.ckObj, C.int(arg1), C.int(arg2))


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the FetchSequenceHeaders method
func (z *Imap) FetchSequenceHeadersAsync(arg1 int, arg2 int, c chan *Task) {

    hTask := C.CkImap_FetchSequenceHeadersAsync(z.ckObj, C.int(arg1), C.int(arg2))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Retrieves a single message from the IMAP server, including attachments if the
// AutoDownloadAttachments property is true. If bUid is true, then msgId
// represents a UID. If bUid is false, then msgId represents a sequence number.
func (z *Imap) FetchSingle(arg1 int, arg2 bool) *Email {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retObj := C.CkImap_FetchSingle(z.ckObj, C.int(arg1), b2)


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Asynchronous version of the FetchSingle method
func (z *Imap) FetchSingleAsync(arg1 int, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkImap_FetchSingleAsync(z.ckObj, C.int(arg1), b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Retrieves a single message from the IMAP server and returns a string containing
// the complete MIME source of the email, including attachments if the
// AutoDownloadAttachments property is true. If bUid is true, then msgId
// represents a UID. If bUid is false, then msgId represents a sequence number.
// return a string or nil if failed.
func (z *Imap) FetchSingleAsMime(arg1 int, arg2 bool) *string {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkImap_fetchSingleAsMime(z.ckObj, C.int(arg1), b2)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FetchSingleAsMime method
func (z *Imap) FetchSingleAsMimeAsync(arg1 int, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkImap_FetchSingleAsMimeAsync(z.ckObj, C.int(arg1), b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Retrieves a single message from the IMAP server into the sbMime object. If bUid is
// true, then msgId represents a UID. If bUid is false, then msgId represents a
// sequence number. If successful, the sbMime will contain the complete MIME of the
// email, including attachments if the AutoDownloadAttachments property is true.
func (z *Imap) FetchSingleAsMimeSb(arg1 int, arg2 bool, arg3 *StringBuilder) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkImap_FetchSingleAsMimeSb(z.ckObj, C.int(arg1), b2, arg3.ckObj)


    return v != 0
}


// Asynchronous version of the FetchSingleAsMimeSb method
func (z *Imap) FetchSingleAsMimeSbAsync(arg1 int, arg2 bool, arg3 *StringBuilder, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkImap_FetchSingleAsMimeSbAsync(z.ckObj, C.int(arg1), b2, arg3.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Retrieves a single message from the IMAP server into the mimeData object.. If bUid
// is true, then msgId represents a UID. If bUid is false, then msgId represents
// a sequence number. If successful, the mimeData will contain the complete MIME of the
// email, including attachments if the AutoDownloadAttachments property is true.
func (z *Imap) FetchSingleBd(arg1 int, arg2 bool, arg3 *BinData) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkImap_FetchSingleBd(z.ckObj, C.int(arg1), b2, arg3.ckObj)


    return v != 0
}


// Asynchronous version of the FetchSingleBd method
func (z *Imap) FetchSingleBdAsync(arg1 int, arg2 bool, arg3 *BinData, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkImap_FetchSingleBdAsync(z.ckObj, C.int(arg1), b2, arg3.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Retrieves a single message header from the IMAP server. If the method fails, it
// may return a NULL reference. The following methods are useful for retrieving
// information about attachments and flags after an email header is retrieved:
// GetMailNumAttach, GetMailAttachSize, GetMailAttachFilename, GetMailFlag. If bUid
// is true, then msgID represents a UID. If bUid is false, then msgID represents a
// sequence number.
func (z *Imap) FetchSingleHeader(arg1 int, arg2 bool) *Email {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retObj := C.CkImap_FetchSingleHeader(z.ckObj, C.int(arg1), b2)


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Asynchronous version of the FetchSingleHeader method
func (z *Imap) FetchSingleHeaderAsync(arg1 int, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkImap_FetchSingleHeaderAsync(z.ckObj, C.int(arg1), b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Fetches and returns the MIME of a single email header.
// return a string or nil if failed.
func (z *Imap) FetchSingleHeaderAsMime(arg1 int, arg2 bool) *string {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkImap_fetchSingleHeaderAsMime(z.ckObj, C.int(arg1), b2)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FetchSingleHeaderAsMime method
func (z *Imap) FetchSingleHeaderAsMimeAsync(arg1 int, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkImap_FetchSingleHeaderAsMimeAsync(z.ckObj, C.int(arg1), b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns a message set object containing all the UIDs in the currently selected
// mailbox. A NULL object reference is returned on failure.
func (z *Imap) GetAllUids() *MessageSet {

    retObj := C.CkImap_GetAllUids(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &MessageSet{retObj}
}


// Asynchronous version of the GetAllUids method
func (z *Imap) GetAllUidsAsync(c chan *Task) {

    hTask := C.CkImap_GetAllUidsAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the Nth attachment filename. Indexing begins at 0.
// return a string or nil if failed.
func (z *Imap) GetMailAttachFilename(arg1 *Email, arg2 int) *string {

    retStr := C.CkImap_getMailAttachFilename(z.ckObj, arg1.ckObj, C.int(arg2))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth attachment size in bytes. Indexing begins at 0.
func (z *Imap) GetMailAttachSize(arg1 *Email, arg2 int) int {

    v := C.CkImap_GetMailAttachSize(z.ckObj, arg1.ckObj, C.int(arg2))


    return int(v)
}


// Sends a "Status" command to get the status of a mailbox. Returns an XML string
// containing the status values as named attributes. Possible status values are:
//     messages: The number of messages in the mailbox.
//     recent: The number of messages with the \Recent flag set.
//     uidnext: The next unique identifier value of the mailbox.
//     uidvalidity: The unique identifier validity value of the mailbox.
//     unseen: The number of messages which do not have the \Seen flag set.
// 
// An example of the string returned by this method is: _LT_status messages="240"
// recent="0" uidnext="3674" uidvalidity="3" unseen="213" /_GT_
//
// return a string or nil if failed.
func (z *Imap) GetMailboxStatus(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkImap_getMailboxStatus(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetMailboxStatus method
func (z *Imap) GetMailboxStatusAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_GetMailboxStatusAsync(z.ckObj, cstr1)

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


// Returns the value of a flag (1 = yes, 0 = no) for an email. Both standard system
// flags as well as custom flags may be set. Standard system flags typically begin
// with a backslash character, such as "\Seen", "\Answered", "\Flagged", "\Draft",
// "\Deleted", and "\Answered". Custom flags can be anything, such as "NonJunk",
// "$label1", "$MailFlagBit1", etc. .
func (z *Imap) GetMailFlag(arg1 *Email, arg2 string) int {
    cstr2 := C.CString(arg2)

    v := C.CkImap_GetMailFlag(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Returns the number of email attachments.
func (z *Imap) GetMailNumAttach(arg1 *Email) int {

    v := C.CkImap_GetMailNumAttach(z.ckObj, arg1.ckObj)


    return int(v)
}


// Returns the size (in bytes) of the entire email including attachments.
func (z *Imap) GetMailSize(arg1 *Email) int {

    v := C.CkImap_GetMailSize(z.ckObj, arg1.ckObj)


    return int(v)
}


// Sends the GETQUOTA command and returns the response in JSON format. This feature
// is only possible with IMAP servers that support the QUOTA extension/capability.
// return a string or nil if failed.
func (z *Imap) GetQuota(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkImap_getQuota(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetQuota method
func (z *Imap) GetQuotaAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_GetQuotaAsync(z.ckObj, cstr1)

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


// Sends the GETQUOTAROOT command and returns the response in JSON format. This
// feature is only possible with IMAP servers that support the QUOTA
// extension/capability.
// return a string or nil if failed.
func (z *Imap) GetQuotaRoot(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkImap_getQuotaRoot(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetQuotaRoot method
func (z *Imap) GetQuotaRootAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_GetQuotaRootAsync(z.ckObj, cstr1)

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


// Returns the IMAP server's digital certificate (for SSL / TLS connections).
func (z *Imap) GetSslServerCert() *Cert {

    retObj := C.CkImap_GetSslServerCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns true if the capability indicated by name is found in the capabilityResponse.
// Otherwise returns false.
func (z *Imap) HasCapability(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkImap_HasCapability(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Polls the connection to see if any real-time updates are available. The timeoutMs
// indicates how long to wait for incoming updates. This method does not send a
// command to the IMAP server, it simply checks the connection for already-arrived
// messages that the IMAP server sent. This method would only be called after IDLE
// has already been started via the IdleStart method.
// 
// If updates are available, they are returned in an XML string having the format
// as shown below. There is one child node for each notification. The possible
// notifcations are:
//     flags -- lists flags that have been set or unset for an email.
//     expunge -- provides the sequence number for an email that has been deleted.
//     exists -- reports the new number of messages in the currently selected
//     mailbox.
//     recent -- reports the new number of messages with the /RECENT flag set.
//     raw -- reports an unanticipated response line that was not parsed by
//     Chilkat. This should be reported to support@chilkatoft.com
// 
// A sample showing all possible notifications (except for "raw") is shown below.
// _LT_idle_GT_
//     _LT_flags seqnum="59" uid="11876"_GT_
//         _LT_flag_GT_\Deleted_LT_/flag_GT_
//         _LT_flag_GT_\Seen_LT_/flag_GT_
//     _LT_/flags_GT_
//     _LT_flags seqnum="69" uid="11889"_GT_
//         _LT_flag_GT_\Seen_LT_/flag_GT_
//     _LT_/flags_GT_
//     _LT_expunge_GT_58_LT_/expunge_GT_
//     _LT_expunge_GT_58_LT_/expunge_GT_
//     _LT_expunge_GT_67_LT_/expunge_GT_
//     _LT_exists_GT_115_LT_/exists_GT_
//     _LT_recent_GT_0_LT_/recent_GT_
// _LT_/idle_GT_
// 
// If no updates have been received, the returned XML string has the following
// format, as shown below. The
// _LT_idle_GT__LT_/idle_GT_
// 
// NOTE:Once IdleStart has been called, this method can and should be called
// frequently to see if any updates have arrived. This is NOT the same as polling
// the IMAP server because it does not send any requests to the IMAP server. It
// simply checks to see if any messages (i.e. updates) from the IMAP server are
// available and waiting to be read.
//
// return a string or nil if failed.
func (z *Imap) IdleCheck(arg1 int) *string {

    retStr := C.CkImap_idleCheck(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the IdleCheck method
func (z *Imap) IdleCheckAsync(arg1 int, c chan *Task) {

    hTask := C.CkImap_IdleCheckAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a command to the IMAP server to stop receiving real-time updates.
func (z *Imap) IdleDone() bool {

    v := C.CkImap_IdleDone(z.ckObj)


    return v != 0
}


// Asynchronous version of the IdleDone method
func (z *Imap) IdleDoneAsync(c chan *Task) {

    hTask := C.CkImap_IdleDoneAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an IDLE command to the IMAP server to begin receiving real-time updates.
func (z *Imap) IdleStart() bool {

    v := C.CkImap_IdleStart(z.ckObj)


    return v != 0
}


// Asynchronous version of the IdleStart method
func (z *Imap) IdleStartAsync(c chan *Task) {

    hTask := C.CkImap_IdleStartAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the last known "connected" state with the IMAP server. IsConnected does
// not send a message to the IMAP server to determine if it is still connected. The
// Noop method may be called to specifically send a no-operation message to
// determine actual connectivity.
// 
// The IsConnected method is useful for checking to see if the component is already
// in a known disconnected state.
//
func (z *Imap) IsConnected() bool {

    v := C.CkImap_IsConnected(z.ckObj)


    return v != 0
}


// Returns true if already logged into an IMAP server, otherwise returns false.
func (z *Imap) IsLoggedIn() bool {

    v := C.CkImap_IsLoggedIn(z.ckObj)


    return v != 0
}


// Returns true if the component is unlocked, false if not.
func (z *Imap) IsUnlocked() bool {

    v := C.CkImap_IsUnlocked(z.ckObj)


    return v != 0
}


// Returns a subset of the complete list of mailboxes available on the IMAP server.
// This method has the side-effect of setting the SeparatorChar property to the
// correct character used by the IMAP server, which is typically "/" or ".".
// 
// Note: The term "mailbox" and "folder" are synonymous. Whenever the word
// "mailbox" is used, it has the same meaning as "folder".
// 
// The reference and wildcardedMailbox parameters are passed unaltered to the IMAP
// LIST command:
// FROM RFC 3501 (IMAP Protocol)
// 
//       The LIST command returns a subset of names from the complete set
//       of all names available to the client.  Zero or more untagged LIST
//       replies are returned, containing the name attributes, hierarchy
//       delimiter, and name; see the description of the LIST reply for
//       more detail.
// 
//       An empty ("" string) reference name argument indicates that the
//       mailbox name is interpreted as by SELECT.  The returned mailbox
//       names MUST match the supplied mailbox name pattern.  A non-empty
//       reference name argument is the name of a mailbox or a level of
//       mailbox hierarchy, and indicates the context in which the mailbox
//       name is interpreted.
// 
//       An empty ("" string) mailbox name argument is a special request to
//       return the hierarchy delimiter and the root name of the name given
//       in the reference.  The value returned as the root MAY be the empty
//       string if the reference is non-rooted or is an empty string.  In
//       all cases, a hierarchy delimiter (or NIL if there is no hierarchy)
//       is returned.  This permits a client to get the hierarchy delimiter
//       (or find out that the mailbox names are flat) even when no
//       mailboxes by that name currently exist.
// 
//       The reference and mailbox name arguments are interpreted into a
//       canonical form that represents an unambiguous left-to-right
//       hierarchy.  The returned mailbox names will be in the interpreted
//       form.
// 
//            Note: The interpretation of the reference argument is
//            implementation-defined.  It depends upon whether the
//            server implementation has a concept of the "current
//            working directory" and leading "break out characters",
//            which override the current working directory.
// 
//            For example, on a server which exports a UNIX or NT
//            filesystem, the reference argument contains the current
//            working directory, and the mailbox name argument would
//            contain the name as interpreted in the current working
//            directory.
// 
//            If a server implementation has no concept of break out
//            characters, the canonical form is normally the reference
//            name appended with the mailbox name.  Note that if the
//            server implements the namespace convention (section
//            5.1.2), "#" is a break out character and must be treated
//            as such.
// 
//            If the reference argument is not a level of mailbox
//            hierarchy (that is, it is a \NoInferiors name), and/or
//            the reference argument does not end with the hierarchy
//            delimiter, it is implementation-dependent how this is
//            interpreted.  For example, a reference of "foo/bar" and
//            mailbox name of "rag/baz" could be interpreted as
//            "foo/bar/rag/baz", "foo/barrag/baz", or "foo/rag/baz".
//            A client SHOULD NOT use such a reference argument except
//            at the explicit request of the user.  A hierarchical
//            browser MUST NOT make any assumptions about server
//            interpretation of the reference unless the reference is
//            a level of mailbox hierarchy AND ends with the hierarchy
//            delimiter.
// 
//       Any part of the reference argument that is included in the
//       interpreted form SHOULD prefix the interpreted form.  It SHOULD
//       also be in the same form as the reference name argument.  This
//       rule permits the client to determine if the returned mailbox name
//       is in the context of the reference argument, or if something about
//       the mailbox argument overrode the reference argument.  Without
//       this rule, the client would have to have knowledge of the server's
//       naming semantics including what characters are "breakouts" that
//       override a naming context.
// 
//            For example, here are some examples of how references
//            and mailbox names might be interpreted on a UNIX-based
//            server:
// 
//                Reference     Mailbox Name  Interpretation
//                ------------  ------------  --------------
//                ~smith/Mail/  foo.*         ~smith/Mail/foo.*
//                archive/      %             archive/%
//                #news.        comp.mail.*   #news.comp.mail.*
//                ~smith/Mail/  /usr/doc/foo  /usr/doc/foo
//                archive/      ~fred/Mail/*  ~fred/Mail/*
// 
//            The first three examples demonstrate interpretations in
//            the context of the reference argument.  Note that
//            "~smith/Mail" SHOULD NOT be transformed into something
//            like "/u2/users/smith/Mail", or it would be impossible
//            for the client to determine that the interpretation was
//            in the context of the reference.
// 
//       The character "*" is a wildcard, and matches zero or more
//       characters at this position.  The character "%" is similar to "*",
//       but it does not match a hierarchy delimiter.  If the "%" wildcard
//       is the last character of a mailbox name argument, matching levels
//       of hierarchy are also returned.  If these levels of hierarchy are
//       not also selectable mailboxes, they are returned with the
//       \Noselect mailbox name attribute (see the description of the LIST
//       response for more details).
// 
//       Server implementations are permitted to "hide" otherwise
//       accessible mailboxes from the wildcard characters, by preventing
//       certain characters or names from matching a wildcard in certain
//       situations.  For example, a UNIX-based server might restrict the
//       interpretation of "*" so that an initial "/" character does not
//       match.
// 
//       The special name INBOX is included in the output from LIST, if
//       INBOX is supported by this server for this user and if the
//       uppercase string "INBOX" matches the interpreted reference and
//       mailbox name arguments with wildcards as described above.  The
//       criteria for omitting INBOX is whether SELECT INBOX will return
//       failure; it is not relevant whether the user's real INBOX resides
//       on this or some other server.
//
func (z *Imap) ListMailboxes(arg1 string, arg2 string) *Mailboxes {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkImap_ListMailboxes(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Mailboxes{retObj}
}


// Asynchronous version of the ListMailboxes method
func (z *Imap) ListMailboxesAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_ListMailboxesAsync(z.ckObj, cstr1, cstr2)

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


// The same as ListMailboxes, but returns only the subscribed mailboxes. (See
// ListMailboxes for more information.)
func (z *Imap) ListSubscribed(arg1 string, arg2 string) *Mailboxes {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkImap_ListSubscribed(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Mailboxes{retObj}
}


// Asynchronous version of the ListSubscribed method
func (z *Imap) ListSubscribedAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_ListSubscribedAsync(z.ckObj, cstr1, cstr2)

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


// Loads the caller of the task's async method.
func (z *Imap) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkImap_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Logs into the IMAP server. The component must first be connected to an IMAP
// server by calling Connect. If XOAUTH2 authentication is required, pass the
// XOAUTH2 access token in place of the password. (For GMail, the Chilkat HTTP
// class/object's G_SvcOauthAccessToken method can be called to obtain an XOAUTH2
// access token.)
// 
// To authenticate using XOAUTH2, make sure the AuthMethod property is set to
// "XOAUTH2". The XOAUTH2 authentication functionality was added in version
// 9.5.0.44.
//
func (z *Imap) Login(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkImap_Login(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the Login method
func (z *Imap) LoginAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_LoginAsync(z.ckObj, cstr1, cstr2)

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


// The same as Login, except the login name and password are passed as secure
// strings.
func (z *Imap) LoginSecure(arg1 *SecureString, arg2 *SecureString) bool {

    v := C.CkImap_LoginSecure(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Asynchronous version of the LoginSecure method
func (z *Imap) LoginSecureAsync(arg1 *SecureString, arg2 *SecureString, c chan *Task) {

    hTask := C.CkImap_LoginSecureAsync(z.ckObj, arg1.ckObj, arg2.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Logs out of the IMAP server.
func (z *Imap) Logout() bool {

    v := C.CkImap_Logout(z.ckObj)


    return v != 0
}


// Asynchronous version of the Logout method
func (z *Imap) LogoutAsync(c chan *Task) {

    hTask := C.CkImap_LogoutAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Moves a set of messages from one mailbox to another. Note: This is only possible
// if the IMAP server supports the "MOVE" extension. The messageSet contains message UIDs
// or sequence numbers for messages in the currently selected mailbox. The destFolder is
// the destination mailbox/folder.
func (z *Imap) MoveMessages(arg1 *MessageSet, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkImap_MoveMessages(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the MoveMessages method
func (z *Imap) MoveMessagesAsync(arg1 *MessageSet, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_MoveMessagesAsync(z.ckObj, arg1.ckObj, cstr2)

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


// Sends a NOOP command to the IMAP server and receives the response. The component
// must be connected and authenticated for this to succeed. Sending a NOOP is a
// good way of determining whether the connection to the IMAP server is up and
// active.
func (z *Imap) Noop() bool {

    v := C.CkImap_Noop(z.ckObj)


    return v != 0
}


// Asynchronous version of the Noop method
func (z *Imap) NoopAsync(c chan *Task) {

    hTask := C.CkImap_NoopAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Fetches the flags for an email and updates the flags in the email's header. When
// an email is retrieved from the IMAP server, it embeds the flags into the header
// in fields beginning with "ckx-". Methods such as GetMailFlag read these header
// fields.
func (z *Imap) RefetchMailFlags(arg1 *Email) bool {

    v := C.CkImap_RefetchMailFlags(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the RefetchMailFlags method
func (z *Imap) RefetchMailFlagsAsync(arg1 *Email, c chan *Task) {

    hTask := C.CkImap_RefetchMailFlagsAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Renames a mailbox. Can also be used to move a mailbox from one location to
// another. For example, from "Inbox.parent.test" to "Inbox.newParent.test", or
// from "abc.xyz" to "def.qrs".
func (z *Imap) RenameMailbox(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkImap_RenameMailbox(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the RenameMailbox method
func (z *Imap) RenameMailboxAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_RenameMailboxAsync(z.ckObj, cstr1, cstr2)

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


// Searches the already selected mailbox for messages that match criteria and returns a
// message set of all matching messages. If bUid is true, then UIDs are returned
// in the message set, otherwise sequence numbers are returned.
// 
// Note: It seems that Microsoft IMAP servers, such as outlook.office365.com and
// imap-mail.outlook.com do not support anything other than 7bit us-ascii chars in
// the search criteria string, regardless of the SEARCH charset that might be
// specified.
// 
// The criteria is passed through to the low-level IMAP protocol unmodified, so the
// rules for the IMAP SEARCH command (RFC 3501) apply and are reproduced here:
// FROM RFC 3501 (IMAP Protocol)
// 
//       The SEARCH command searches the mailbox for messages that match
//       the given searching criteria.  Searching criteria consist of one
//       or more search keys.  The untagged SEARCH response from the server
//       contains a listing of message sequence numbers corresponding to
//       those messages that match the searching criteria.
// 
//       When multiple keys are specified, the result is the intersection
//       (AND function) of all the messages that match those keys.  For
//       example, the criteria DELETED FROM "SMITH" SINCE 1-Feb-1994 refers
//       to all deleted messages from Smith that were placed in the mailbox
//       since February 1, 1994.  A search key can also be a parenthesized
//       list of one or more search keys (e.g., for use with the OR and NOT
//       keys).
// 
//       Server implementations MAY exclude [MIME-IMB] body parts with
//       terminal content media types other than TEXT and MESSAGE from
//       consideration in SEARCH matching.
// 
//       The OPTIONAL [CHARSET] specification consists of the word
//       "CHARSET" followed by a registered [CHARSET].  It indicates the
//       [CHARSET] of the strings that appear in the search criteria.
//       [MIME-IMB] content transfer encodings, and [MIME-HDRS] strings in
//       [RFC-2822]/[MIME-IMB] headers, MUST be decoded before comparing
//       text in a [CHARSET] other than US-ASCII.  US-ASCII MUST be
//       supported; other [CHARSET]s MAY be supported.
// 
//       If the server does not support the specified [CHARSET], it MUST
//       return a tagged NO response (not a BAD).  This response SHOULD
//       contain the BADCHARSET response code, which MAY list the
//       [CHARSET]s supported by the server.
// 
//       In all search keys that use strings, a message matches the key if
//       the string is a substring of the field.  The matching is
//       case-insensitive.
// 
//       The defined search keys are as follows.  Refer to the Formal
//       Syntax section for the precise syntactic definitions of the
//       arguments.
// 
//       
//          Messages with message sequence numbers corresponding to the
//          specified message sequence number set.
// 
//       ALL
//          All messages in the mailbox; the default initial key for
//          ANDing.
// 
//       ANSWERED
//          Messages with the \Answered flag set.
// 
//       BCC 
//          Messages that contain the specified string in the envelope
//          structure's BCC field.
// 
//       BEFORE 
//          Messages whose internal date (disregarding time and timezone)
//          is earlier than the specified date.
// 
//       BODY 
//          Messages that contain the specified string in the body of the
//          message.
// 
//       CC 
//          Messages that contain the specified string in the envelope
//          structure's CC field.
// 
//       DELETED
//          Messages with the \Deleted flag set.
// 
//       DRAFT
//          Messages with the \Draft flag set.
// 
//       FLAGGED
//          Messages with the \Flagged flag set.
// 
//       FROM 
//          Messages that contain the specified string in the envelope
//          structure's FROM field.
// 
//       HEADER  
//          Messages that have a header with the specified field-name (as
//          defined in [RFC-2822]) and that contains the specified string
//          in the text of the header (what comes after the colon).  If the
//          string to search is zero-length, this matches all messages that
//          have a header line with the specified field-name regardless of
//          the contents.
// 
//       KEYWORD 
//          Messages with the specified keyword flag set.
// 
//       LARGER 
//          Messages with an [RFC-2822] size larger than the specified
//          number of octets.
// 
//       NEW
//          Messages that have the \Recent flag set but not the \Seen flag.
//          This is functionally equivalent to "(RECENT UNSEEN)".
// 
//       NOT 
//          Messages that do not match the specified search key.
// 
//       OLD
//          Messages that do not have the \Recent flag set.  This is
//          functionally equivalent to "NOT RECENT" (as opposed to "NOT
//          NEW").
// 
//       ON 
//          Messages whose internal date (disregarding time and timezone)
//          is within the specified date.
// 
//       OR  
//          Messages that match either search key.
// 
//       RECENT
//          Messages that have the \Recent flag set.
// 
//       SEEN
//          Messages that have the \Seen flag set.
// 
//       SENTBEFORE 
//          Messages whose [RFC-2822] Date: header (disregarding time and
//          timezone) is earlier than the specified date.
// 
//       SENTON 
//          Messages whose [RFC-2822] Date: header (disregarding time and
//          timezone) is within the specified date.
// 
//       SENTSINCE 
//          Messages whose [RFC-2822] Date: header (disregarding time and
//          timezone) is within or later than the specified date.
// 
//       SINCE 
//          Messages whose internal date (disregarding time and timezone)
//          is within or later than the specified date.
// 
//       SMALLER 
//          Messages with an [RFC-2822] size smaller than the specified
//          number of octets.
// 
//       SUBJECT 
//          Messages that contain the specified string in the envelope
//          structure's SUBJECT field.
// 
//       TEXT 
//          Messages that contain the specified string in the header or
//          body of the message.
// 
//       TO 
//          Messages that contain the specified string in the envelope
//          structure's TO field.
// 
//       UID 
//          Messages with unique identifiers corresponding to the specified
//          unique identifier set.  Sequence set ranges are permitted.
// 
//       UNANSWERED
//          Messages that do not have the \Answered flag set.
// 
//       UNDELETED
//          Messages that do not have the \Deleted flag set.
// 
//       UNDRAFT
//          Messages that do not have the \Draft flag set.
// 
//       UNFLAGGED
//          Messages that do not have the \Flagged flag set.
// 
//       UNKEYWORD 
//          Messages that do not have the specified keyword flag set.
// 
//       UNSEEN
//          Messages that do not have the \Seen flag set.
//
func (z *Imap) Search(arg1 string, arg2 bool) *MessageSet {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retObj := C.CkImap_Search(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &MessageSet{retObj}
}


// Asynchronous version of the Search method
func (z *Imap) SearchAsync(arg1 string, arg2 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkImap_SearchAsync(z.ckObj, cstr1, b2)

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


// Selects a mailbox. A mailbox must be selected before some methods, such as
// Search or FetchSingle, can be called. If the logged-on user does not have
// write-access to the mailbox, call ExamineMailbox instead.
// 
// Calling this method updates the NumMessages property.
// 
// Note: The term "mailbox" and "folder" are synonymous. Whenever the word
// "mailbox" is used, it has the same meaning as "folder".
//
func (z *Imap) SelectMailbox(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_SelectMailbox(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SelectMailbox method
func (z *Imap) SelectMailboxAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_SelectMailboxAsync(z.ckObj, cstr1)

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


// Allows for the sending of arbitrary commands to the IMAP server.
// return a string or nil if failed.
func (z *Imap) SendRawCommand(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkImap_sendRawCommand(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the SendRawCommand method
func (z *Imap) SendRawCommandAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_SendRawCommandAsync(z.ckObj, cstr1)

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


// The same as SendRawCommand, but instead of returning the response as a string,
// the binary bytes of the response are returned.
func (z *Imap) SendRawCommandB(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkImap_SendRawCommandB(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the SendRawCommandB method
func (z *Imap) SendRawCommandBAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_SendRawCommandBAsync(z.ckObj, cstr1)

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


// The same as SendRawCommandB, except that the command is provided as binary bytes
// rather than a string.
func (z *Imap) SendRawCommandC(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkImap_SendRawCommandC(z.ckObj, ckbyd1, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the SendRawCommandC method
func (z *Imap) SendRawCommandCAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkImap_SendRawCommandCAsync(z.ckObj, ckbyd1)

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


// Explicitly specifies the certificate to be used for decrypting encrypted email.
func (z *Imap) SetDecryptCert(arg1 *Cert) bool {

    v := C.CkImap_SetDecryptCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Used to explicitly specify the certificate and associated private key to be used
// for decrypting S/MIME (PKCS7) email.
func (z *Imap) SetDecryptCert2(arg1 *Cert, arg2 *PrivateKey) bool {

    v := C.CkImap_SetDecryptCert2(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Sets a flag for a single message on the IMAP server. If value = 1, the flag is
// turned on, if value = 0, the flag is turned off. Standard system flags such as
// "\Deleted", "\Seen", "\Answered", "\Flagged", "\Draft", and "\Answered" may be
// set. Custom flags such as "NonJunk", "$label1", "$MailFlagBit1", etc. may also
// be set.
// 
// If bUid is true, then msgId represents a UID. If bUid is false, then msgId
// represents a sequence number.
//
func (z *Imap) SetFlag(arg1 int, arg2 bool, arg3 string, arg4 int) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    v := C.CkImap_SetFlag(z.ckObj, C.int(arg1), b2, cstr3, C.int(arg4))

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the SetFlag method
func (z *Imap) SetFlagAsync(arg1 int, arg2 bool, arg3 string, arg4 int, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    hTask := C.CkImap_SetFlagAsync(z.ckObj, C.int(arg1), b2, cstr3, C.int(arg4))

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


// Sets a flag for each message in the message set on the IMAP server. If value = 1,
// the flag is turned on, if value = 0, the flag is turned off. Standard system
// flags such as "\Deleted", "\Seen", "\Answered", "\Flagged", "\Draft", and
// "\Answered" may be set. Custom flags such as "NonJunk", "$label1",
// "$MailFlagBit1", etc. may also be set.
func (z *Imap) SetFlags(arg1 *MessageSet, arg2 string, arg3 int) bool {
    cstr2 := C.CString(arg2)

    v := C.CkImap_SetFlags(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SetFlags method
func (z *Imap) SetFlagsAsync(arg1 *MessageSet, arg2 string, arg3 int, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_SetFlagsAsync(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

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


// Sets a flag for a single message on the IMAP server. The UID of the email object
// is used to find the message on the IMAP server that is to be affected. If value =
// 1, the flag is turned on, if value = 0, the flag is turned off.
// 
// Both standard system flags as well as custom flags may be set. Standard system
// flags typically begin with a backslash character, such as "\Deleted", "\Seen",
// "\Answered", "\Flagged", "\Draft", and "\Answered". Custom flags can be
// anything, such as "NonJunk", "$label1", "$MailFlagBit1", etc. .
// 
// Note: When the Chilkat IMAP component downloads an email from an IMAP server, it
// inserts a "ckx-imap-uid" header field in the email object. This is subsequently
// used by this method to get the UID associated with the email. The "ckx-imap-uid"
// header must be present for this method to be successful.
// 
// Note: Calling this method is identical to calling the SetFlag method, except the
// UID is automatically obtained from the email object.
// 
// Important: Setting the "Deleted" flag does not remove the email from the
// mailbox. Emails marked "Deleted" are removed when the Expunge method is called.
//
func (z *Imap) SetMailFlag(arg1 *Email, arg2 string, arg3 int) bool {
    cstr2 := C.CString(arg2)

    v := C.CkImap_SetMailFlag(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SetMailFlag method
func (z *Imap) SetMailFlagAsync(arg1 *Email, arg2 string, arg3 int, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_SetMailFlagAsync(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

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


// Sets the quota for a quotaRoot. The resource should be one of two keywords:"STORAGE" or
// "MESSAGE". Use "STORAGE" to set the maximum capacity of the combined messages in
// quotaRoot. Use "MESSAGE" to set the maximum number of messages allowed.
// 
// If setting a STORAGE quota, the quota is in units of 1024 octets. For example, to
// specify a limit of 500,000,000 bytes, set quota equal to 500,000.
// 
// This feature is only possible with IMAP servers that support the QUOTA
// extension/capability. If an IMAP server supports the QUOTA extension, it likely
// supports the STORAGE resource. The MESSAGE resource is less commonly supported.
//
func (z *Imap) SetQuota(arg1 string, arg2 string, arg3 int) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkImap_SetQuota(z.ckObj, cstr1, cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SetQuota method
func (z *Imap) SetQuotaAsync(arg1 string, arg2 string, arg3 int, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_SetQuotaAsync(z.ckObj, cstr1, cstr2, C.int(arg3))

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


// Specifies a client-side certificate to be used for the SSL / TLS connection. In
// most cases, servers do not require client-side certificates for SSL/TLS. A
// client-side certificate is typically used in high-security situations where the
// certificate is an additional means to indentify the client to the server.
func (z *Imap) SetSslClientCert(arg1 *Cert) bool {

    v := C.CkImap_SetSslClientCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// (Same as SetSslClientCert, but allows a .pfx/.p12 file to be used directly)
// Specifies a client-side certificate to be used for the SSL / TLS connection. In
// most cases, servers do not require client-side certificates for SSL/TLS. A
// client-side certificate is typically used in high-security situations where the
// certificate is an additional means to indentify the client to the server.
// 
// The pemDataOrFilename may contain the actual PEM data, or it may contain the path of the PEM
// file. This method will automatically recognize whether it is a path or the PEM
// data itself.
//
func (z *Imap) SetSslClientCertPem(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkImap_SetSslClientCertPem(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// (Same as SetSslClientCert, but allows a .pfx/.p12 file to be used directly)
// Specifies a client-side certificate to be used for the SSL / TLS connection. In
// most cases, servers do not require client-side certificates for SSL/TLS. A
// client-side certificate is typically used in high-security situations where the
// certificate is an additional means to indentify the client to the server.
func (z *Imap) SetSslClientCertPfx(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkImap_SetSslClientCertPfx(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Searches the already selected mailbox for messages that match searchCriteria and returns a
// message set of all matching messages in the order specified by sortCriteria. If bUid is
// true, then UIDs are returned in the message set, otherwise sequence numbers
// are returned.
// 
// The sortCriteria is a string of SPACE separated keywords to indicate sort order (default
// is ascending). The keyword "REVERSE" can precede a keyword to reverse the sort
// order (i.e. make it descending). Possible sort keywords are:
//     ARRIVAL
//     CC
//     DATE
//     FROM
//     SIZE
//     SUBJECT
//     TO
// 
// Some examples of sortCriteria are:
//     "SUBJECT REVERSE DATE"
//     "REVERSE SIZE"
//     "ARRIVAL"
// 
// The searchCriteria is passed through to the low-level IMAP protocol unmodified, and
// therefore the rules for the IMAP SEARCH command (RFC 3501) apply. See the
// documentation for the Search method for more details (and also see RFC 3501).
//
func (z *Imap) Sort(arg1 string, arg2 string, arg3 string, arg4 bool) *MessageSet {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    retObj := C.CkImap_Sort(z.ckObj, cstr1, cstr2, cstr3, b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &MessageSet{retObj}
}


// Asynchronous version of the Sort method
func (z *Imap) SortAsync(arg1 string, arg2 string, arg3 string, arg4 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkImap_SortAsync(z.ckObj, cstr1, cstr2, cstr3, b4)

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


// Authenticates with the SSH server using public-key authentication. The
// corresponding public key must have been installed on the SSH server for the
// sshLogin. Authentication will succeed if the matching privateKey is provided.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *Imap) SshAuthenticatePk(arg1 string, arg2 *SshKey) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_SshAuthenticatePk(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SshAuthenticatePk method
func (z *Imap) SshAuthenticatePkAsync(arg1 string, arg2 *SshKey, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_SshAuthenticatePkAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Authenticates with the SSH server using a sshLogin and sshPassword.
// 
// An SSH tunneling (port forwarding) session always begins by first calling
// SshTunnel to connect to the SSH server, then calling either AuthenticatePw or
// AuthenticatePk to authenticate. Following this, your program should call Connect
// to connect with the IMAP server (via the SSH tunnel) and then Login to
// authenticate with the IMAP server.
// 
// Note: Once the SSH tunnel is setup by calling SshTunnel and SshAuthenticatePw
// (or SshAuthenticatePk), all underlying communcations with the IMAP server use
// the SSH tunnel. No changes in programming are required other than making two
// initial calls to setup the tunnel.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *Imap) SshAuthenticatePw(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkImap_SshAuthenticatePw(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SshAuthenticatePw method
func (z *Imap) SshAuthenticatePwAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkImap_SshAuthenticatePwAsync(z.ckObj, cstr1, cstr2)

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


// Closes the SSH tunnel previously opened by SshOpenTunnel.
func (z *Imap) SshCloseTunnel() bool {

    v := C.CkImap_SshCloseTunnel(z.ckObj)


    return v != 0
}


// Asynchronous version of the SshCloseTunnel method
func (z *Imap) SshCloseTunnelAsync(c chan *Task) {

    hTask := C.CkImap_SshCloseTunnelAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Connects to an SSH server and creates a tunnel for IMAP. The sshHostname is the
// hostname (or IP address) of the SSH server. The sshPort is typically 22, which is
// the standard SSH port number.
// 
// An SSH tunneling (port forwarding) session always begins by first calling
// SshOpenTunnel to connect to the SSH server, followed by calling either
// SshAuthenticatePw or SshAuthenticatePk to authenticate. Your program would then
// call Connect to connect with the IMAP server (via the SSH tunnel) and then Login
// to authenticate with the IMAP server.
// 
// Note: Once the SSH tunnel is setup by calling SshOpenTunnel and
// SshAuthenticatePw (or SshAuthenticatePk), all underlying communcations with the
// IMAP server use the SSH tunnel. No changes in programming are required other
// than making two initial calls to setup the tunnel.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *Imap) SshOpenTunnel(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_SshOpenTunnel(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SshOpenTunnel method
func (z *Imap) SshOpenTunnelAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_SshOpenTunnelAsync(z.ckObj, cstr1, C.int(arg2))

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


// Sets one or more flags to a specific value for an email. The email is indicated
// by either a UID or sequence number, depending on whether bUid is true (UID) or
// false (sequence number).
// 
// flagNames should be a space separated string of flag names. Both standard and
// customer flags may be set. Standard flag names typically begin with a backslash
// character. For example: "\Seen \Answered". Custom flag names may also be
// included. Custom flags often begin with a $ character, such as "$label1", or
// "$MailFlagBit0". Other customer flags may begin with any character, such as
// "NonJunk".
// 
// value should be 1 to turn the flags on, or 0 to turn the flags off.
//
func (z *Imap) StoreFlags(arg1 int, arg2 bool, arg3 string, arg4 int) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    v := C.CkImap_StoreFlags(z.ckObj, C.int(arg1), b2, cstr3, C.int(arg4))

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the StoreFlags method
func (z *Imap) StoreFlagsAsync(arg1 int, arg2 bool, arg3 string, arg4 int, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    hTask := C.CkImap_StoreFlagsAsync(z.ckObj, C.int(arg1), b2, cstr3, C.int(arg4))

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


// Subscribe to an IMAP mailbox.
// 
// Note: The term "mailbox" and "folder" are synonymous. Whenever the word
// "mailbox" is used, it has the same meaning as "folder".
//
func (z *Imap) Subscribe(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_Subscribe(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Subscribe method
func (z *Imap) SubscribeAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_SubscribeAsync(z.ckObj, cstr1)

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


// Sends the THREAD command to search the already selected mailbox for messages
// that match searchCriteria.
// 
// The following explanation is fromRFC 5256
// <https://tools.ietf.org/html/rfc5256>:
// 
// The THREAD command is a variant of SEARCH with threading semantics 
// for the results.  Thread has two arguments before the searching 
// criteria argument: a threading algorithm and the searching 
// charset.
// 
// The THREAD command first searches the mailbox for messages that
// match the given searching criteria using the charset argument for
// the interpretation of strings in the searching criteria.  It then
// returns the matching messages in an untagged THREAD response,
// threaded according to the specified threading algorithm.
// 
// All collation is in ascending order.  Earlier dates collate before
// later dates and strings are collated according to ascending values.
// 
// The defined threading algorithms are as follows:
// 
//       ORDEREDSUBJECT
// 
//          The ORDEREDSUBJECT threading algorithm is also referred to as
//          "poor man's threading".  The searched messages are sorted by
//          base subject and then by the sent date.  The messages are then
//          split into separate threads, with each thread containing
//          messages with the same base subject text.  Finally, the threads
//          are sorted by the sent date of the first message in the thread.
// 
//          The top level or "root" in ORDEREDSUBJECT threading contains
//          the first message of every thread.  All messages in the root
//          are siblings of each other.  The second message of a thread is
//          the child of the first message, and subsequent messages of the
//          thread are siblings of the second message and hence children of
//          the message at the root.  Hence, there are no grandchildren in
//          ORDEREDSUBJECT threading.
// 
//          Children in ORDEREDSUBJECT threading do not have descendents.
//          Client implementations SHOULD treat descendents of a child in a
//          server response as being siblings of that child.
// 
//       REFERENCES
// 
//          The REFERENCES threading algorithm threads the searched
//          messages by grouping them together in parent/child
//          relationships based on which messages are replies to others.
//          The parent/child relationships are built using two methods:
//          reconstructing a message's ancestry using the references
//          contained within it; and checking the original (not base)
//          subject of a message to see if it is a reply to (or forward of)
//          another message.
// 
// SeeRFC 5256
// <https://tools.ietf.org/html/rfc5256> for more details:
// 
// The searchCriteria is passed through to the low-level IMAP protocol unmodified, and
// therefore the rules for the IMAP SEARCH command (RFC 3501) apply. See the
// documentation for the Search method for more details (and also see RFC 3501).
// 
// The results are returned in a JSON object to make it easy to parse the
// parent/child relationships. See the example below for details.
//
func (z *Imap) ThreadCmd(arg1 string, arg2 string, arg3 string, arg4 bool) *JsonObject {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    retObj := C.CkImap_ThreadCmd(z.ckObj, cstr1, cstr2, cstr3, b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Asynchronous version of the ThreadCmd method
func (z *Imap) ThreadCmdAsync(arg1 string, arg2 string, arg3 string, arg4 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkImap_ThreadCmdAsync(z.ckObj, cstr1, cstr2, cstr3, b4)

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


// Unlocks the component. This must be called once at the beginning of your program
// to unlock the component. A purchased unlock code is provided when the IMAP
// component is licensed. Any string, such as "Hello World", may be passed to this
// method to automatically begin a fully-functional 30-day trial.
func (z *Imap) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unsubscribe from an IMAP mailbox.
func (z *Imap) Unsubscribe(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkImap_Unsubscribe(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Unsubscribe method
func (z *Imap) UnsubscribeAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkImap_UnsubscribeAsync(z.ckObj, cstr1)

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


// Adds an XML certificate vault to the object's internal list of sources to be
// searched for certificates and private keys when encrypting/decrypting or
// signing/verifying. Unlike the AddPfxSourceData and AddPfxSourceFile methods,
// only a single XML certificate vault can be used. If UseCertVault is called
// multiple times, only the last certificate vault will be used, as each call to
// UseCertVault will replace the certificate vault provided in previous calls.
func (z *Imap) UseCertVault(arg1 *XmlCertVault) bool {

    v := C.CkImap_UseCertVault(z.ckObj, arg1.ckObj)


    return v != 0
}


// Uses an existing SSH tunnel for the connection to the IMAP server. This method
// is identical to the UseSshTunnel method, except the SSH connection is obtained
// from an SSH object instead of a Socket object.
// 
// This is useful for sharing an existing SSH tunnel connection wth other objects.
// (SSH is a protocol where the tunnel contains many logical channels. IMAP
// connections can exist simultaneously with other connection within a single SSH
// tunnel as SSH channels.)
//
func (z *Imap) UseSsh(arg1 *Ssh) bool {

    v := C.CkImap_UseSsh(z.ckObj, arg1.ckObj)


    return v != 0
}


// Uses an existing SSH tunnel. This is useful for sharing an existing SSH tunnel
// connection wth other objects. (SSH is a protocol where the tunnel contains many
// logical channels. IMAP connections can exist simultaneously with other
// connection within a single SSH tunnel as SSH channels.)
func (z *Imap) UseSshTunnel(arg1 *Socket) bool {

    v := C.CkImap_UseSshTunnel(z.ckObj, arg1.ckObj)


    return v != 0
}


