// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkMailMan.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewMailMan() *MailMan {
	return &MailMan{C.CkMailMan_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *MailMan) DisposeMailMan() {
    if z != nil {
	C.CkMailMan_Dispose(z.ckObj)
	}
}




func (z *MailMan) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkMailMan_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *MailMan) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkMailMan_setExternalProgress(z.ckObj,1)
}

func (z *MailMan) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkMailMan_setExternalProgress(z.ckObj,1)
}

func (z *MailMan) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkMailMan_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *MailMan) AbortCurrent() bool {
    v := int(C.CkMailMan_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *MailMan) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putAbortCurrent(z.ckObj,v)
}

// property: AllOrNone
// Prevents sending any email if any of the addresses in the recipient list are
// rejected by the SMTP server. The default value is false, which indicates that
// the mail sending should continue even if some email addresses are invalid.
// (Note: Not all SMTP servers check the validity of email addresses, and even for
// those that do, it is not 100% accurate.)
// 
// Note: An SMTP server only knows the validity of email addresses within the
// domain it controls.
// 
// Important: The AllOrNone property only works if SMTP pipelining is turned off.
// By default, the SmtpPipelining property is turned on and has the value of
// true. If all-or-none behavior is desired, make sure to set SmtpPipelining
// equal to false.
//
func (z *MailMan) AllOrNone() bool {
    v := int(C.CkMailMan_getAllOrNone(z.ckObj))
    return v != 0
}
// property setter: AllOrNone
func (z *MailMan) SetAllOrNone(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putAllOrNone(z.ckObj,v)
}

// property: AutoFix
// If true, then the following will occur when a connection is made to an SMTP or
// POP3 server:
// 
// 1) If the SmtpPort property = 465, then sets StartTLS = false and SmtpSsl =
// true
// 2) If the SmtpPort property = 25, sets SmtpSsl = false
// 3) If the MailPort property = 995, sets PopSsl = true
// 4) If the MailPort property = 110, sets PopSsl = false
// 
// The default value of this property is true.
//
func (z *MailMan) AutoFix() bool {
    v := int(C.CkMailMan_getAutoFix(z.ckObj))
    return v != 0
}
// property setter: AutoFix
func (z *MailMan) SetAutoFix(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putAutoFix(z.ckObj,v)
}

// property: AutoGenMessageId
// Controls whether a unique Message-ID header is auto-generated for each email
// sent.
// 
// The Message-ID header field should contain a unique message ID for each email
// that is sent. The default behavior is to auto-generate this header field at the
// time the message is sent. This makes it easier for the same email object to be
// re-used. If the message ID is not unique, the SMTP server may consider the
// message to be a duplicate of one that has already been sent, and may discard it
// without sending. This property controls whether message IDs are automatically
// generated. If auto-generation is turned on (true), the value returned by
// GetHeaderField("Message-ID") will not reflect the actual message ID that gets
// sent with the email.
// 
// To turn off automatic Message-ID generation, set this property to false.
//
func (z *MailMan) AutoGenMessageId() bool {
    v := int(C.CkMailMan_getAutoGenMessageId(z.ckObj))
    return v != 0
}
// property setter: AutoGenMessageId
func (z *MailMan) SetAutoGenMessageId(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putAutoGenMessageId(z.ckObj,v)
}

// property: AutoSmtpRset
// If true, then the SMTP "RSET" command is automatically sent to ensure that the
// SMTP connection is in a valid state when a new email is about to be sent on an
// already established connection. The default value is false.
// 
// Important: This property only applies when an email is sent on an already-open
// SMTP connection.
//
func (z *MailMan) AutoSmtpRset() bool {
    v := int(C.CkMailMan_getAutoSmtpRset(z.ckObj))
    return v != 0
}
// property setter: AutoSmtpRset
func (z *MailMan) SetAutoSmtpRset(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putAutoSmtpRset(z.ckObj,v)
}

// property: AutoUnwrapSecurity
// If true, then digitally signed and/or encrypted email when downloaded from a
// mail server is automatically "unwrapped" and the results of the signature
// validation and decryption are available in various email object properties and
// methods. The default value of this property is true. Set this property to
// false to prevent unwrapping.
// 
// Note: A digitally signed or encrypted email can ONLY be verified and/or
// decrypted when initially loading the original MIME into the email object (i.e.
// when downloading from the server, or when loading from MIME). Once the MIME is
// parsed and stored in the internal email object format, the exactnes of the MIME
// has been lost and the signature can no longer be verified. This is why the
// signature is verified upon the intial loading of the MIME, and the results are
// made available through the various properties and methods. This property
// provides a means for downloading email where the .p7m (or .p7s) attachments are
// are to be treated as simple attachments and the desire is to access or save the
// original .p7m/.p7s files.
//
func (z *MailMan) AutoUnwrapSecurity() bool {
    v := int(C.CkMailMan_getAutoUnwrapSecurity(z.ckObj))
    return v != 0
}
// property setter: AutoUnwrapSecurity
func (z *MailMan) SetAutoUnwrapSecurity(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putAutoUnwrapSecurity(z.ckObj,v)
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
func (z *MailMan) ClientIpAddress() string {
    return C.GoString(C.CkMailMan_clientIpAddress(z.ckObj))
}
// property setter: ClientIpAddress
func (z *MailMan) SetClientIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putClientIpAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConnectFailReason
// This property will be set to the status of the last connection made (or failed
// to be made) by any method.
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
func (z *MailMan) ConnectFailReason() int {
    return int(C.CkMailMan_getConnectFailReason(z.ckObj))
}

// property: ConnectTimeout
// The time (in seconds) to wait before while trying to connect to a mail server
// (POP3 or SMTP). The default value is 30.
func (z *MailMan) ConnectTimeout() int {
    return int(C.CkMailMan_getConnectTimeout(z.ckObj))
}
// property setter: ConnectTimeout
func (z *MailMan) SetConnectTimeout(value int) {
    C.CkMailMan_putConnectTimeout(z.ckObj,C.int(value))
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
func (z *MailMan) DebugLogFilePath() string {
    return C.GoString(C.CkMailMan_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *MailMan) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DsnEnvid
// (An SMTP DSN service extension feature) An arbitrary string that will be used as
// the ENVID property when sending email. See RFC 3461 for more details.
func (z *MailMan) DsnEnvid() string {
    return C.GoString(C.CkMailMan_dsnEnvid(z.ckObj))
}
// property setter: DsnEnvid
func (z *MailMan) SetDsnEnvid(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putDsnEnvid(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DsnNotify
// (An SMTP DSN service extension feature) A string that will be used as the NOTIFY
// parameter when sending email. (See RFC 3461 for more details. ) This string can
// be left blank, or can be set to "NEVER", or any combination of a comma-separated
// list of "SUCCESS", "FAILURE", or "NOTIFY".
func (z *MailMan) DsnNotify() string {
    return C.GoString(C.CkMailMan_dsnNotify(z.ckObj))
}
// property setter: DsnNotify
func (z *MailMan) SetDsnNotify(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putDsnNotify(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DsnRet
// (An SMTP DSN service extension feature) A string that will be used as the RET
// parameter when sending email. (See RFC 3461 for more details. ) This string can
// be left blank, or can be set to "FULL" to receive entire-message DSN
// notifications, or "HDRS" to receive header-only DSN notifications.
func (z *MailMan) DsnRet() string {
    return C.GoString(C.CkMailMan_dsnRet(z.ckObj))
}
// property setter: DsnRet
func (z *MailMan) SetDsnRet(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putDsnRet(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EmbedCertChain
// If true, causes the digital certificate chain to be embedded in signed emails.
// The certificates in the chain of authentication are embedded up to but not
// including the root certificate. If the IncludeRootCert property is also true,
// then the root CA certificate is also included in the S/MIME signature.
func (z *MailMan) EmbedCertChain() bool {
    v := int(C.CkMailMan_getEmbedCertChain(z.ckObj))
    return v != 0
}
// property setter: EmbedCertChain
func (z *MailMan) SetEmbedCertChain(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putEmbedCertChain(z.ckObj,v)
}

// property: Filter
// An expression that is applied to any of the following method calls when present:
// LoadXmlFile, LoadXmlString, LoadMbx, CopyMail, and TransferMail. For these
// methods, only the emails that match the filter's expression are returned in the
// email bundle. In the case of TransferMail, only the matching emails are removed
// from the mail server. The filter allows any header field, or the body, to be
// checked.
// Here are some examples of expressions:
// 
// Body like "mortgage rates*". 
// Subject contains "update" and From contains "chilkat" 
// To = "info@chilkatsoft.com" 
// 
// Here are the general rules for forming filter expressions:
// 
// Any MIME header field name can be used, case is insensitive. 
// Literal strings are double-quoted, and case is insensitive. 
// The "*" wildcard matches 0 or more occurrences of any character. 
// Parentheses can be used to control precedence. 
// The logical operators are: AND, OR, NOT (case insensitive) 
// Comparison operators are: =, , =, String comparison operators are: CONTAINS, LIKE (case insensitive)
func (z *MailMan) Filter() string {
    return C.GoString(C.CkMailMan_filter(z.ckObj))
}
// property setter: Filter
func (z *MailMan) SetFilter(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putFilter(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// The time interval, in milliseconds, between AbortCheck event callbacks. The
// heartbeat provides a means for an application to monitor a mail-sending and/or
// mail-reading method call, and to abort it while in progress.
func (z *MailMan) HeartbeatMs() int {
    return int(C.CkMailMan_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *MailMan) SetHeartbeatMs(value int) {
    C.CkMailMan_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: HeloHostname
// Specifies the hostname to be used for the EHLO/HELO command sent to an SMTP
// server. By default, this property is an empty string which causes the local
// hostname to be used.
func (z *MailMan) HeloHostname() string {
    return C.GoString(C.CkMailMan_heloHostname(z.ckObj))
}
// property setter: HeloHostname
func (z *MailMan) SetHeloHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putHeloHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyAuthMethod
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy authentication method name. Valid choices are "Basic" or "NTLM".
func (z *MailMan) HttpProxyAuthMethod() string {
    return C.GoString(C.CkMailMan_httpProxyAuthMethod(z.ckObj))
}
// property setter: HttpProxyAuthMethod
func (z *MailMan) SetHttpProxyAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putHttpProxyAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyDomain
// The NTLM authentication domain (optional) if NTLM authentication is used.
func (z *MailMan) HttpProxyDomain() string {
    return C.GoString(C.CkMailMan_httpProxyDomain(z.ckObj))
}
// property setter: HttpProxyDomain
func (z *MailMan) SetHttpProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putHttpProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyHostname
// If an HTTP proxy is to be used, set this property to the HTTP proxy hostname or
// IPv4 address (in dotted decimal notation).
func (z *MailMan) HttpProxyHostname() string {
    return C.GoString(C.CkMailMan_httpProxyHostname(z.ckObj))
}
// property setter: HttpProxyHostname
func (z *MailMan) SetHttpProxyHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putHttpProxyHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPassword
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy password.
func (z *MailMan) HttpProxyPassword() string {
    return C.GoString(C.CkMailMan_httpProxyPassword(z.ckObj))
}
// property setter: HttpProxyPassword
func (z *MailMan) SetHttpProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putHttpProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPort
// If an HTTP proxy is to be used, set this property to the HTTP proxy port number.
// (Two commonly used HTTP proxy ports are 8080 and 3128.)
func (z *MailMan) HttpProxyPort() int {
    return int(C.CkMailMan_getHttpProxyPort(z.ckObj))
}
// property setter: HttpProxyPort
func (z *MailMan) SetHttpProxyPort(value int) {
    C.CkMailMan_putHttpProxyPort(z.ckObj,C.int(value))
}

// property: HttpProxyUsername
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy login name.
func (z *MailMan) HttpProxyUsername() string {
    return C.GoString(C.CkMailMan_httpProxyUsername(z.ckObj))
}
// property setter: HttpProxyUsername
func (z *MailMan) SetHttpProxyUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putHttpProxyUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ImmediateDelete
// If true (the default) then any method that deletes an email from the POP3
// server will also issue a QUIT command to close the session to ensure the message
// is deleted immediately.
// 
// The POP3 protocol is such that the DELE command marks a message for deletion. It
// is not actually deleted until the QUIT command is sent and the session is
// closed. If ImmediateDelete is true, then any Chilkat MailMan method that marks
// a message (or messages) for deletion will also followup with a QUIT command and
// close the session. If your program sets ImmediateDelete to false, it must make
// sure to call Pop3EndSession to ensure that messages marked for deletion are
// actually deleted.
//
func (z *MailMan) ImmediateDelete() bool {
    v := int(C.CkMailMan_getImmediateDelete(z.ckObj))
    return v != 0
}
// property setter: ImmediateDelete
func (z *MailMan) SetImmediateDelete(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putImmediateDelete(z.ckObj,v)
}

// property: IncludeRootCert
// Controls whether the root certificate in the chain of authentication (i.e. the
// CA root certificate) is included within the S/MIME signature of a signed email.
// Note: This property only applies if the EmbedCertChain property is also true.
func (z *MailMan) IncludeRootCert() bool {
    v := int(C.CkMailMan_getIncludeRootCert(z.ckObj))
    return v != 0
}
// property setter: IncludeRootCert
func (z *MailMan) SetIncludeRootCert(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putIncludeRootCert(z.ckObj,v)
}

// property: IsPop3Connected
// Returns true if still connected to the POP3 server. Otherwise returns false.
// 
// Note: Accessing this property does not trigger any communication with the POP3
// server. A connection to the POP3 server is established by explicitly calling
// Pop3BeginSession, or it is implicitly established as needed by any method that
// requires communication. A lost connection might only be detected when attempting
// to communicate with the server. To truly determine if a connection to the POP3
// server is open and valid, it may be necessary to call the Pop3Noop method
// instead. This property might return true if the server has disconnected, but
// the client has not attempted to communicate with the server since the
// disconnect.
//
func (z *MailMan) IsPop3Connected() bool {
    v := int(C.CkMailMan_getIsPop3Connected(z.ckObj))
    return v != 0
}

// property: IsSmtpConnected
// Returns true if still connected to the SMTP server. Otherwise returns false
// (if there was never a connection in the first place, or if the connection was
// lost).
// 
// Note: Accessing this property does not trigger any communication with the SMTP
// server. A connection to the SMTP server is established by explicitly calling
// OpenSmtpConnection, or it is implicitly established as needed by any method that
// requires communication. A lost connection might only be detected when attempting
// to communicate with the server. To truly determine if a connection to the SMTP
// server is open and valid, it may be necessary to call the SmtpNoop method
// instead. This property might return true if the server has disconnected, but
// the client has not attempted to communicate with the server since the
// disconnect.
//
func (z *MailMan) IsSmtpConnected() bool {
    v := int(C.CkMailMan_getIsSmtpConnected(z.ckObj))
    return v != 0
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *MailMan) LastErrorHtml() string {
    return C.GoString(C.CkMailMan_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *MailMan) LastErrorText() string {
    return C.GoString(C.CkMailMan_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *MailMan) LastErrorXml() string {
    return C.GoString(C.CkMailMan_lastErrorXml(z.ckObj))
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
func (z *MailMan) LastMethodSuccess() bool {
    v := int(C.CkMailMan_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *MailMan) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putLastMethodSuccess(z.ckObj,v)
}

// property: LastSmtpStatus
// Returns the last SMTP diagnostic status code. This can be checked after sending
// an email. SMTP reply codes are defined by RFC 821 - Simple Mail Transfer
// Protocol.
func (z *MailMan) LastSmtpStatus() int {
    return int(C.CkMailMan_getLastSmtpStatus(z.ckObj))
}

// property: LogMailReceivedFilename
// A log filename where the MailMan will log each message in the exact form it was
// received from a POP3 server. This property is provided for help in debugging.
func (z *MailMan) LogMailReceivedFilename() string {
    return C.GoString(C.CkMailMan_logMailReceivedFilename(z.ckObj))
}
// property setter: LogMailReceivedFilename
func (z *MailMan) SetLogMailReceivedFilename(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putLogMailReceivedFilename(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LogMailSentFilename
// A log filename where the MailMan will log the exact message sent to the SMTP
// server. This property is helpful in debugging.
func (z *MailMan) LogMailSentFilename() string {
    return C.GoString(C.CkMailMan_logMailSentFilename(z.ckObj))
}
// property setter: LogMailSentFilename
func (z *MailMan) SetLogMailSentFilename(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putLogMailSentFilename(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: MailHost
// The domain name of the POP3 server. Do not include "http://" in the domain name.
// This property may also be set to an IP address string, such as "168.144.70.227".
// Both IPv4 and IPv6 address formats are supported.
func (z *MailMan) MailHost() string {
    return C.GoString(C.CkMailMan_mailHost(z.ckObj))
}
// property setter: MailHost
func (z *MailMan) SetMailHost(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putMailHost(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: MailPort
// The port number of the POP3 server. Only needs to be set if the POP3 server is
// running on a non-standard port. The default value is 110. (If SSL/TLS is used by
// setting the PopSsl property = true, then this property should probably be set
// to 995, which is the standard SSL/TLS port for POP3.)
func (z *MailMan) MailPort() int {
    return int(C.CkMailMan_getMailPort(z.ckObj))
}
// property setter: MailPort
func (z *MailMan) SetMailPort(value int) {
    C.CkMailMan_putMailPort(z.ckObj,C.int(value))
}

// property: MaxCount
// Limits the number of messages the MailMan will try to retrieve from the POP3
// server in a single method call. If you are trying to read a large mailbox, you
// might set this to a value such as 100 to download 100 emails at a time.
func (z *MailMan) MaxCount() int {
    return int(C.CkMailMan_getMaxCount(z.ckObj))
}
// property setter: MaxCount
func (z *MailMan) SetMaxCount(value int) {
    C.CkMailMan_putMaxCount(z.ckObj,C.int(value))
}

// property: OAuth2AccessToken
// The OAUTH2 access token if OAUTH2 authentication is to be used for the
// authentication. For GMail, the Chilkat HTTP class/object's G_SvcOauthAccessToken
// method can be called to obtain an OAUTH2 access token for a GMail service
// account.
// 
// Starting in v9.5.0.83, this property can be set to also do XOAUTH2
// authentication for POP3 sessions (pop.gmail.com).
//
func (z *MailMan) OAuth2AccessToken() string {
    return C.GoString(C.CkMailMan_oAuth2AccessToken(z.ckObj))
}
// property setter: OAuth2AccessToken
func (z *MailMan) SetOAuth2AccessToken(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putOAuth2AccessToken(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OpaqueSigning
// When set to true, signed emails are sent using opaque signing. The default is
// to send clear-text (multipart/signed) emails.
func (z *MailMan) OpaqueSigning() bool {
    v := int(C.CkMailMan_getOpaqueSigning(z.ckObj))
    return v != 0
}
// property setter: OpaqueSigning
func (z *MailMan) SetOpaqueSigning(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putOpaqueSigning(z.ckObj,v)
}

// property: P7mEncryptAttachFilename
// The filename attribute to be used in the Content-Disposition header field when
// sending a PCKS7 encrypted email. The default value is "smime.p7m".
func (z *MailMan) P7mEncryptAttachFilename() string {
    return C.GoString(C.CkMailMan_p7mEncryptAttachFilename(z.ckObj))
}
// property setter: P7mEncryptAttachFilename
func (z *MailMan) SetP7mEncryptAttachFilename(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putP7mEncryptAttachFilename(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: P7mSigAttachFilename
// The filename attribute to be used in the Content-Disposition header field when
// sending a PCKS7 opaque signed email. The default value is "smime.p7m".
func (z *MailMan) P7mSigAttachFilename() string {
    return C.GoString(C.CkMailMan_p7mSigAttachFilename(z.ckObj))
}
// property setter: P7mSigAttachFilename
func (z *MailMan) SetP7mSigAttachFilename(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putP7mSigAttachFilename(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: P7sSigAttachFilename
// The filename attribute to be used in the Content-Disposition header field when
// sending a signed email with a detached PKCS7 signature. The default value is
// "smime.p7s".
func (z *MailMan) P7sSigAttachFilename() string {
    return C.GoString(C.CkMailMan_p7sSigAttachFilename(z.ckObj))
}
// property setter: P7sSigAttachFilename
func (z *MailMan) SetP7sSigAttachFilename(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putP7sSigAttachFilename(z.ckObj,cStr)
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
func (z *MailMan) PercentDoneScale() int {
    return int(C.CkMailMan_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *MailMan) SetPercentDoneScale(value int) {
    C.CkMailMan_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: Pop3SessionId
// 0 if no POP3 session is active. Otherwise a positive integer that is incremented
// with each new POP3 session. It may be used to determine if a new POP3 session
// has been established.
func (z *MailMan) Pop3SessionId() int {
    return int(C.CkMailMan_getPop3SessionId(z.ckObj))
}

// property: Pop3SessionLog
// This string property accumulates the raw commands sent to the POP3 server, and
// the raw responses received from the POP3 server. This property is read-only, but
// it may be cleared by calling ClearPop3SessionLog.
func (z *MailMan) Pop3SessionLog() string {
    return C.GoString(C.CkMailMan_pop3SessionLog(z.ckObj))
}

// property: Pop3SPA
// Controls whether SPA authentication for POP3 is used or not. To use SPA
// authentication, set this propoerty = true. No other programming changes are
// required. The default value is false.
// 
// Note: If SPA (i.e. NTLM) authentication does not succeed, set the
// Global.DefaultNtlmVersion property equal to 1 and then retry.
//
func (z *MailMan) Pop3SPA() bool {
    v := int(C.CkMailMan_getPop3SPA(z.ckObj))
    return v != 0
}
// property setter: Pop3SPA
func (z *MailMan) SetPop3SPA(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putPop3SPA(z.ckObj,v)
}

// property: Pop3SslServerCertVerified
// When connecting via SSL, this property is true if the POP3 server's SSL
// certificate was verified. Otherwise it is set to false.
func (z *MailMan) Pop3SslServerCertVerified() bool {
    v := int(C.CkMailMan_getPop3SslServerCertVerified(z.ckObj))
    return v != 0
}

// property: Pop3Stls
// If true, then an unencrypted connection (typically on port 110) is
// automatically converted to a secure TLS connection via the STLS command (see RFC
// 2595) when connecting. This should only be used with POP3 servers that support
// the STLS capability. If this property is set to true, then the PopSsl property
// should be set to false. (The PopSsl property controls whether the connection
// is SSL/TLS from the beginning. Setting the Pop3Stls property = true indicates
// that the POP3 client will initially connect unencrypted and then convert to
// TLS.)
func (z *MailMan) Pop3Stls() bool {
    v := int(C.CkMailMan_getPop3Stls(z.ckObj))
    return v != 0
}
// property setter: Pop3Stls
func (z *MailMan) SetPop3Stls(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putPop3Stls(z.ckObj,v)
}

// property: PopPassword
// The POP3 password.
// 
// If the Pop3SPA property is set, the PopUsername and PopPassword properties may
// be set to the string "default" to cause the component to use the current
// logged-on credentials (of the calling process) for authentication.
//
func (z *MailMan) PopPassword() string {
    return C.GoString(C.CkMailMan_popPassword(z.ckObj))
}
// property setter: PopPassword
func (z *MailMan) SetPopPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putPopPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: PopPasswordBase64
// Provides a way to specify the POP3 password from a Base64-encoded string.
func (z *MailMan) PopPasswordBase64() string {
    return C.GoString(C.CkMailMan_popPasswordBase64(z.ckObj))
}
// property setter: PopPasswordBase64
func (z *MailMan) SetPopPasswordBase64(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putPopPasswordBase64(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: PopSsl
// Controls whether TLS/SSL is used when reading email from a POP3 server. Note:
// Check first to determine if your POP3 server can accept TLS/SSL connections.
// Also, be sure to set the MailPort property to the TLS/SSL POP3 port number,
// which is typically 995.
func (z *MailMan) PopSsl() bool {
    v := int(C.CkMailMan_getPopSsl(z.ckObj))
    return v != 0
}
// property setter: PopSsl
func (z *MailMan) SetPopSsl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putPopSsl(z.ckObj,v)
}

// property: PopUsername
// The POP3 login name.
// 
// If the Pop3SPA property is set, the PopUsername and PopPassword properties may
// be set to the string "default" to cause the component to use the current
// logged-on credentials (of the calling process) for authentication.
//
func (z *MailMan) PopUsername() string {
    return C.GoString(C.CkMailMan_popUsername(z.ckObj))
}
// property setter: PopUsername
func (z *MailMan) SetPopUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putPopUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *MailMan) PreferIpv6() bool {
    v := int(C.CkMailMan_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *MailMan) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putPreferIpv6(z.ckObj,v)
}

// property: ReadTimeout
// The maximum time to wait, in seconds, if the POP3 or SMTP server stops
// responding. The default value is 30 seconds.
func (z *MailMan) ReadTimeout() int {
    return int(C.CkMailMan_getReadTimeout(z.ckObj))
}
// property setter: ReadTimeout
func (z *MailMan) SetReadTimeout(value int) {
    C.CkMailMan_putReadTimeout(z.ckObj,C.int(value))
}

// property: RequireSslCertVerify
// If true, then the mailman will verify the SMTP or POP3 server's SSL
// certificate when connecting. The certificate is expired, or if the cert's
// signature is invalid, the connection is not allowed. The default value of this
// property is false. (Obviously, this only applies to SSL/TLS connections.)
func (z *MailMan) RequireSslCertVerify() bool {
    v := int(C.CkMailMan_getRequireSslCertVerify(z.ckObj))
    return v != 0
}
// property setter: RequireSslCertVerify
func (z *MailMan) SetRequireSslCertVerify(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putRequireSslCertVerify(z.ckObj,v)
}

// property: ResetDateOnLoad
// Controls whether the Date header field is reset to the current date/time when an
// email is loaded from LoadMbx, LoadEml, LoadMime, LoadXml, or LoadXmlString. The
// default is false (to not reset the date). To automatically reset the date, set
// this property equal to true.
func (z *MailMan) ResetDateOnLoad() bool {
    v := int(C.CkMailMan_getResetDateOnLoad(z.ckObj))
    return v != 0
}
// property setter: ResetDateOnLoad
func (z *MailMan) SetResetDateOnLoad(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putResetDateOnLoad(z.ckObj,v)
}

// property: SendBufferSize
// The buffer size to be used with the underlying TCP/IP socket for sending. The
// default value is 32767.
func (z *MailMan) SendBufferSize() int {
    return int(C.CkMailMan_getSendBufferSize(z.ckObj))
}
// property setter: SendBufferSize
func (z *MailMan) SetSendBufferSize(value int) {
    C.CkMailMan_putSendBufferSize(z.ckObj,C.int(value))
}

// property: SendIndividual
// Determines how emails are sent to distribution lists. If true, emails are sent
// to each recipient in the list one at a time, with the "To"header field
// containing the email address of the recipient. If false, emails will contain
// in the "To"header field, and are sent to 100 BCC recipients at a time. As an
// example, if your distribution list contained 350 email addresses, 4 emails would
// be sent, the first 3 having 100 BCC recipients, and the last email with 50 BCC
// recipients.The default value of this property is true.
func (z *MailMan) SendIndividual() bool {
    v := int(C.CkMailMan_getSendIndividual(z.ckObj))
    return v != 0
}
// property setter: SendIndividual
func (z *MailMan) SetSendIndividual(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putSendIndividual(z.ckObj,v)
}

// property: SizeLimit
// The MailMan will not try to retrieve mail messages from a POP3 server that are
// greater than this size limit. The default value is 0 indicating no size limit.
// The SizeLimit is specified in number of bytes.
func (z *MailMan) SizeLimit() int {
    return int(C.CkMailMan_getSizeLimit(z.ckObj))
}
// property setter: SizeLimit
func (z *MailMan) SetSizeLimit(value int) {
    C.CkMailMan_putSizeLimit(z.ckObj,C.int(value))
}

// property: SmtpAuthMethod
// This property should usually be left empty. The MailMan will by default choose
// the most secure login method available to prevent unencrypted username and
// passwords from being transmitted if possible. However, some SMTP servers may not
// advertise the acceptable authorization methods, and therefore it is not possible
// to automatically determine the best authorization method. To force a particular
// auth method, or to prevent any authorization from being used, set this property
// to one of the following values: "NONE", "LOGIN", "PLAIN", "CRAM-MD5", or "NTLM".
// 
// Note: If NTLM authentication does not succeed, set the Global.DefaultNtlmVersion
// property equal to 1 and then retry.
//
func (z *MailMan) SmtpAuthMethod() string {
    return C.GoString(C.CkMailMan_smtpAuthMethod(z.ckObj))
}
// property setter: SmtpAuthMethod
func (z *MailMan) SetSmtpAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSmtpAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SmtpFailReason
// A keyword that indicates the cause of failure (or success) for the last SMTP
// related method called. Possible values are:
//     Success The method call was successful.
//     Failed A general failure not covered by any of the other possible keywords.
//     NoValidRecipients The SMTP server rejected all receipients.
//     NoRecipients The app failed to provide any recipients (TO, CC, or BCC).
//     SomeBadRecipients The AllOrNone property is true, and some recipients were
//     rejected by the SMTP server.
//     Aborted The application aborted the method.
//     NoFrom The failed to provide a FROM address.
//     FromFailure The SMTP replied with an error in response to the "MAIL FROM"
//     command.
//     NoCredentials The application did not provide the required credentials, such
//     as username or password.
//     AuthFailure The login (authentication) failed.
//     DataFailure The SMTP replied with an error in response to the "DATA"
//     command.
//     NoSmtpHostname The application failed to provide an SMTP hostname or IP
//     address.
//     StartTlsFailed Failed to convert the TCP connection to TLS via STARTTLS.
//     ConnectFailed Unable to establish a TCP or TLS connection to the SMTP
//     server.
//     GreetingError The SMTP server immediately responded with an error status in
//     the intial greeting.
//     ConnectionLost The connection to the SMTP server was lost at some point
//     during the method call.
//     Timeout A timeout occurred when reading or writing the socket connection.
//     RenderFailed A failure occurred when rendering the email. (Rendering the
//     email for sending includes tasks such as signing or encrypting.)
//     NotUnlocked The UnlockComponent method was not previously called on at least
//     one instance of the mailman object.
//     InternalFailure An internal failure that should be reported to Chilkat
//     support.
func (z *MailMan) SmtpFailReason() string {
    return C.GoString(C.CkMailMan_smtpFailReason(z.ckObj))
}

// property: SmtpHost
// The domain name of the SMTP server. Do not include "http://" in the domain name.
// This property may also be set to an IP address string, such as "168.144.70.227".
// Both IPv4 and IPv6 address formats are supported.
func (z *MailMan) SmtpHost() string {
    return C.GoString(C.CkMailMan_smtpHost(z.ckObj))
}
// property setter: SmtpHost
func (z *MailMan) SetSmtpHost(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSmtpHost(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SmtpLoginDomain
// The Windows domain for logging into the SMTP server. Use this only if your SMTP
// server requires NTLM authentication, which means your SMTP server uses
// Integrated Windows Authentication. If there is no domain, this can be left
// empty.
func (z *MailMan) SmtpLoginDomain() string {
    return C.GoString(C.CkMailMan_smtpLoginDomain(z.ckObj))
}
// property setter: SmtpLoginDomain
func (z *MailMan) SetSmtpLoginDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSmtpLoginDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SmtpPassword
// The password for logging into the SMTP server. Use this only if your SMTP server
// requires authentication. Chilkat Email.NET supports the LOGIN, PLAIN, CRAM-MD5,
// and NTLM login methods, and it will automatically choose the most secure method
// available. Additional login methods will be available in the future.
// 
// If NTLM (Windows-Integrated) authentication is used, the SmtpUsername and
// SmtpPassword properties may be set to the string "default" to cause the
// component to use the current logged-on credentials (of the calling process) for
// authentication.
//
func (z *MailMan) SmtpPassword() string {
    return C.GoString(C.CkMailMan_smtpPassword(z.ckObj))
}
// property setter: SmtpPassword
func (z *MailMan) SetSmtpPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSmtpPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SmtpPipelining
// Controls whether SMTP pipelining is automatically used when the SMTP server
// indicates support for it. The default is true. Setting this property equal to
// false will prevent the SMTP pipelining feature from being used.
func (z *MailMan) SmtpPipelining() bool {
    v := int(C.CkMailMan_getSmtpPipelining(z.ckObj))
    return v != 0
}
// property setter: SmtpPipelining
func (z *MailMan) SetSmtpPipelining(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putSmtpPipelining(z.ckObj,v)
}

// property: SmtpPort
// The port number of the SMTP server used to send email. Only needs to be set if
// the SMTP server is running on a non-standard port. The default value is 25. If
// SmtpSsl is set to true, this property should be set to 465. (TCP port 465 is
// reserved by common industry practice for secure SMTP communication using the SSL
// protocol.)
func (z *MailMan) SmtpPort() int {
    return int(C.CkMailMan_getSmtpPort(z.ckObj))
}
// property setter: SmtpPort
func (z *MailMan) SetSmtpPort(value int) {
    C.CkMailMan_putSmtpPort(z.ckObj,C.int(value))
}

// property: SmtpSessionLog
// This string property accumulates the raw commands sent to the SMTP server, and
// the raw responses received from the SMTP server. This property is read-only, but
// it may be cleared by calling ClearSmtpSessionLog.
func (z *MailMan) SmtpSessionLog() string {
    return C.GoString(C.CkMailMan_smtpSessionLog(z.ckObj))
}

// property: SmtpSsl
// When set to true, causes the mailman to connect to the SMTP server via the
// TLS/SSL protocol.
func (z *MailMan) SmtpSsl() bool {
    v := int(C.CkMailMan_getSmtpSsl(z.ckObj))
    return v != 0
}
// property setter: SmtpSsl
func (z *MailMan) SetSmtpSsl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putSmtpSsl(z.ckObj,v)
}

// property: SmtpSslServerCertVerified
// If using SSL, this property will be set to true if the SMTP server's SSL
// certificate was verified when establishing the connection. Otherwise it is set
// to false.
func (z *MailMan) SmtpSslServerCertVerified() bool {
    v := int(C.CkMailMan_getSmtpSslServerCertVerified(z.ckObj))
    return v != 0
}

// property: SmtpUsername
// The login for logging into the SMTP server. Use this only if your SMTP server
// requires authentication.
// 
// Note: In many cases, an SMTP server will not require authentication when sending
// to an email address local to it's domain. However, when sending email to an
// external domain, authentication is required (i.e. the SMTP server is being used
// as a relay).
// 
// If the SmtpAuthMethod property is set to "NTLM", the SmtpUsername and
// SmtpPassword properties may be set to the string "default" to use the current
// Windows logged-on user credentials.
// 
// smtp.office365.com: If SMTP authentication fails for your smtp.office365.com
// account, it may be that your account is configured to require MFA (multi-factor
// authentication). You may need to change settings to allow for legacy
// authentication (single-factor auth). See
// https://docs.microsoft.com/en-us/azure/active-directory/conditional-access/block-
// legacy-authentication Also, an app password may be required. See
// https://docs.microsoft.com/en-us/azure/active-directory/user-help/multi-factor-au
// thentication-end-user-app-passwords
//
func (z *MailMan) SmtpUsername() string {
    return C.GoString(C.CkMailMan_smtpUsername(z.ckObj))
}
// property setter: SmtpUsername
func (z *MailMan) SetSmtpUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSmtpUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksHostname
// The SOCKS4/SOCKS5 hostname or IPv4 address (in dotted decimal notation). This
// property is only used if the SocksVersion property is set to 4 or 5).
func (z *MailMan) SocksHostname() string {
    return C.GoString(C.CkMailMan_socksHostname(z.ckObj))
}
// property setter: SocksHostname
func (z *MailMan) SetSocksHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSocksHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPassword
// The SOCKS5 password (if required). The SOCKS4 protocol does not include the use
// of a password, so this does not apply to SOCKS4.
func (z *MailMan) SocksPassword() string {
    return C.GoString(C.CkMailMan_socksPassword(z.ckObj))
}
// property setter: SocksPassword
func (z *MailMan) SetSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPort
// The SOCKS4/SOCKS5 proxy port. The default value is 1080. This property only
// applies if a SOCKS proxy is used (if the SocksVersion property is set to 4 or
// 5).
func (z *MailMan) SocksPort() int {
    return int(C.CkMailMan_getSocksPort(z.ckObj))
}
// property setter: SocksPort
func (z *MailMan) SetSocksPort(value int) {
    C.CkMailMan_putSocksPort(z.ckObj,C.int(value))
}

// property: SocksUsername
// The SOCKS4/SOCKS5 proxy username. This property is only used if the SocksVersion
// property is set to 4 or 5).
func (z *MailMan) SocksUsername() string {
    return C.GoString(C.CkMailMan_socksUsername(z.ckObj))
}
// property setter: SocksUsername
func (z *MailMan) SetSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksVersion
// May be set to one of the following integer values:
// 
// 0 - No SOCKS proxy is used. This is the default.
// 4 - Connect via a SOCKS4 proxy.
// 5 - Connect via a SOCKS5 proxy.
//
func (z *MailMan) SocksVersion() int {
    return int(C.CkMailMan_getSocksVersion(z.ckObj))
}
// property setter: SocksVersion
func (z *MailMan) SetSocksVersion(value int) {
    C.CkMailMan_putSocksVersion(z.ckObj,C.int(value))
}

// property: SoRcvBuf
// Sets the receive buffer size socket option. Normally, this property should be
// left unchanged. The default value is 4194304.
// 
// This property can be increased if download performance seems slow. It is
// recommended to be a multiple of 4096.
//
func (z *MailMan) SoRcvBuf() int {
    return int(C.CkMailMan_getSoRcvBuf(z.ckObj))
}
// property setter: SoRcvBuf
func (z *MailMan) SetSoRcvBuf(value int) {
    C.CkMailMan_putSoRcvBuf(z.ckObj,C.int(value))
}

// property: SoSndBuf
// Sets the send buffer size socket option. Normally, this property should be left
// unchanged. The default value is 262144.
// 
// This property can be increased if upload performance seems slow. It is
// recommended to be a multiple of 4096. Testing with sizes such as 512K and 1MB is
// reasonable.
//
func (z *MailMan) SoSndBuf() int {
    return int(C.CkMailMan_getSoSndBuf(z.ckObj))
}
// property setter: SoSndBuf
func (z *MailMan) SetSoSndBuf(value int) {
    C.CkMailMan_putSoSndBuf(z.ckObj,C.int(value))
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
func (z *MailMan) SslAllowedCiphers() string {
    return C.GoString(C.CkMailMan_sslAllowedCiphers(z.ckObj))
}
// property setter: SslAllowedCiphers
func (z *MailMan) SetSslAllowedCiphers(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSslAllowedCiphers(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SslProtocol
// Selects the secure protocol to be used for secure (SSL/TLS) connections (for
// both SMTP and POP3). Possible values are:
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
func (z *MailMan) SslProtocol() string {
    return C.GoString(C.CkMailMan_sslProtocol(z.ckObj))
}
// property setter: SslProtocol
func (z *MailMan) SetSslProtocol(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putSslProtocol(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: StartTLS
// When set to true, causes the mailman to issue a STARTTLS command to switch
// over to a secure SSL/TLS connection prior to authenticating and sending email.
// The default value is false.
func (z *MailMan) StartTLS() bool {
    v := int(C.CkMailMan_getStartTLS(z.ckObj))
    return v != 0
}
// property setter: StartTLS
func (z *MailMan) SetStartTLS(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putStartTLS(z.ckObj,v)
}

// property: StartTLSifPossible
// When set to true, causes the mailman to do STARTTLS (if possible and supported
// by the server) to convert to a secure SSL/TLS connection prior to authenticating
// and sending email. The default value is true.
// 
// Note: Setting the StartTLS property = true causes STARTTLS to always be used,
// even if the SMTP server does not support it. This property allows for a
// non-encrypted connection, whereas the StartTLS property disallows non-encrypted
// connections.
//
func (z *MailMan) StartTLSifPossible() bool {
    v := int(C.CkMailMan_getStartTLSifPossible(z.ckObj))
    return v != 0
}
// property setter: StartTLSifPossible
func (z *MailMan) SetStartTLSifPossible(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putStartTLSifPossible(z.ckObj,v)
}

// property: TlsCipherSuite
// Contains the current or last negotiated TLS cipher suite. If no TLS connection
// has yet to be established, or if a connection as attempted and failed, then this
// will be empty. A sample cipher suite string looks like this:
// TLS_DHE_RSA_WITH_AES_256_CBC_SHA256.
func (z *MailMan) TlsCipherSuite() string {
    return C.GoString(C.CkMailMan_tlsCipherSuite(z.ckObj))
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
func (z *MailMan) TlsPinSet() string {
    return C.GoString(C.CkMailMan_tlsPinSet(z.ckObj))
}
// property setter: TlsPinSet
func (z *MailMan) SetTlsPinSet(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putTlsPinSet(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TlsVersion
// Contains the current or last negotiated TLS protocol version. If no TLS
// connection has yet to be established, or if a connection as attempted and
// failed, then this will be empty. Possible values are "SSL 3.0", "TLS 1.0", "TLS
// 1.1", and "TLS 1.2".
func (z *MailMan) TlsVersion() string {
    return C.GoString(C.CkMailMan_tlsVersion(z.ckObj))
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
func (z *MailMan) UncommonOptions() string {
    return C.GoString(C.CkMailMan_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *MailMan) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkMailMan_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UseApop
// If true, will automatically use APOP authentication if the POP3 server
// supports it. The default value of this property is false.
func (z *MailMan) UseApop() bool {
    v := int(C.CkMailMan_getUseApop(z.ckObj))
    return v != 0
}
// property setter: UseApop
func (z *MailMan) SetUseApop(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putUseApop(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *MailMan) VerboseLogging() bool {
    v := int(C.CkMailMan_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *MailMan) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMailMan_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *MailMan) Version() string {
    return C.GoString(C.CkMailMan_version(z.ckObj))
}
// Adds a PFX to the object's internal list of sources to be searched for
// certificates and private keys when decrypting or when creating signed email for
// sending. Multiple PFX sources can be added by calling this method once for each.
// (On the Windows operating system, the registry-based certificate stores are also
// automatically searched, so it is commonly not required to explicitly add PFX
// sources.)
// 
// The pfxData contains the bytes of a PFX file (also known as PKCS12 or .p12).
//
func (z *MailMan) AddPfxSourceData(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkMailMan_AddPfxSourceData(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a PFX file to the object's internal list of sources to be searched for
// certificates and private keys when decrypting or when sending signed email.
// Multiple PFX files can be added by calling this method once for each. (On the
// Windows operating system, the registry-based certificate stores are also
// automatically searched, so it is commonly not required to explicitly add PFX
// sources.)
// 
// The pfxFilePath contains the bytes of a PFX file (also known as PKCS12 or .p12).
//
func (z *MailMan) AddPfxSourceFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMailMan_AddPfxSourceFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the number of emails available on the POP3 server. Returns -1 on error.
// 
// The VerifyPopConnection method can be called to verify basic TCP/IP connectivity
// with the POP3 server. The VerifyPopLogin method can be called to verify the POP3
// login. The Verify* methods are intended to be called as a way of diagnosing the
// failure when a POP3 method returns an error status.
//
func (z *MailMan) CheckMail() int {

    v := C.CkMailMan_CheckMail(z.ckObj)


    return int(v)
}


// Asynchronous version of the CheckMail method
func (z *MailMan) CheckMailAsync(c chan *Task) {

    hTask := C.CkMailMan_CheckMailAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Clears the list of bad email addresses stored within the Mailman object. When an
// email-sending method is called, any email addresses rejected by the SMTP server
// will be cached within the Mailman object. These can be accessed by calling the
// GetBadEmailAddresses method. This method clears the Mailman's in-memory cache of
// bad addresses.
func (z *MailMan) ClearBadEmailAddresses()  {

    C.CkMailMan_ClearBadEmailAddresses(z.ckObj)



}


// Clears the contents of the Pop3SessionLog property.
func (z *MailMan) ClearPop3SessionLog()  {

    C.CkMailMan_ClearPop3SessionLog(z.ckObj)



}


// Clears the contents of the SmtpSessionLog property.
func (z *MailMan) ClearSmtpSessionLog()  {

    C.CkMailMan_ClearSmtpSessionLog(z.ckObj)



}


// The mailman object automatically opens an SMTP connection (if necessary)
// whenever an email-sending method is called. The connection is kept open until
// explicitly closed by this method. Calling this method is entirely optional. The
// SMTP connection is also automatically closed when the mailman object is
// destructed. Thus, if an application calls SendEmail 10 times to send 10 emails,
// the 1st call will open the SMTP connection, while the subsequent 9 will send
// over the existing connection (unless a property such as username, login,
// hostname, etc. is changed, which would force the connection to become closed and
// re-established with the next mail-sending method call).
// 
// Note: This method sends a QUIT command to the SMTP server prior to closing the
// connection.
//
func (z *MailMan) CloseSmtpConnection() bool {

    v := C.CkMailMan_CloseSmtpConnection(z.ckObj)


    return v != 0
}


// Asynchronous version of the CloseSmtpConnection method
func (z *MailMan) CloseSmtpConnectionAsync(c chan *Task) {

    hTask := C.CkMailMan_CloseSmtpConnectionAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Copy the email from a POP3 server into a EmailBundle. This does not remove the
// email from the POP3 server.
func (z *MailMan) CopyMail() *EmailBundle {

    retObj := C.CkMailMan_CopyMail(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the CopyMail method
func (z *MailMan) CopyMailAsync(c chan *Task) {

    hTask := C.CkMailMan_CopyMailAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Marks multiple emails on the POP3 server for deletion. (Each email in emailBundle that
// is also present on the server is marked for deletion.) To complete the deletion
// of the emails, a "QUIT" message must be sent and the POP3 session ended. This
// will happen automatically when the ImmediateDelete property equals true, which
// is the default. If ImmediateDelete equals false, then the Pop3EndSession
// method can be called to send the "QUIT" and end the session (i.e. disconnect.)
// 
// Note: When making multiple calls to a Delete* method, it's best to turn off
// ImmediateDelete, and then manually call Pop3EndSession to finalize the
// deletions.
// 
// Also, any method call requiring communication with the POP3 server will
// automatically re-establish a session based on the current property settings.
//
func (z *MailMan) DeleteBundle(arg1 *EmailBundle) bool {

    v := C.CkMailMan_DeleteBundle(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the DeleteBundle method
func (z *MailMan) DeleteBundleAsync(arg1 *EmailBundle, c chan *Task) {

    hTask := C.CkMailMan_DeleteBundleAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Marks an email for deletion by message number. WARNING: Be very careful if
// calling this method. Message numbers are specific to a POP3 session. If a
// maildrop has (for example) 10 messages, the message numbers will be 1, 2, 3, ...
// 10. If message number 1 is deleted and a new POP3 session is established, there
// will be 9 messages numbered 1, 2, 3, ... 9.
// 
// IMPORTANT: A POP3 must first be established by either calling Pop3BeginSession
// explicitly, or implicitly by calling some other method that automatically
// establishes the session. This method will not automatically establish a new POP3
// session (because if it did, the message numbers would potentially be different
// than what the application expects).
// 
// This method only marks an email for deletion. It is not actually removed from
// the maildrop until the POP3 session is explicitly ended by calling
// Pop3EndSession.
//
func (z *MailMan) DeleteByMsgnum(arg1 int) bool {

    v := C.CkMailMan_DeleteByMsgnum(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the DeleteByMsgnum method
func (z *MailMan) DeleteByMsgnumAsync(arg1 int, c chan *Task) {

    hTask := C.CkMailMan_DeleteByMsgnumAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Marks an email on the POP3 server for deletion. To complete the deletion of an
// email, a "QUIT" message must be sent and the POP3 session ended. This will
// happen automatically when the ImmediateDelete property equals true, which is
// the default. If ImmediateDelete equals false, then the Pop3EndSession method
// can be called to send the "QUIT" and end the session (i.e. disconnect.)
// 
// Note: When making multiple calls to a Delete* method, it's best to turn off
// ImmediateDelete, and then manually call Pop3EndSession to finalize the
// deletions.
// 
// Also, any method call requiring communication with the POP3 server will
// automatically re-establish a session based on the current property settings.
//
func (z *MailMan) DeleteByUidl(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMailMan_DeleteByUidl(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the DeleteByUidl method
func (z *MailMan) DeleteByUidlAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMailMan_DeleteByUidlAsync(z.ckObj, cstr1)

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


// Marks an email on the POP3 server for deletion. To complete the deletion of an
// email, a "QUIT" message must be sent and the POP3 session ended. This will
// happen automatically when the ImmediateDelete property equals true, which is
// the default. If ImmediateDelete equals false, then the Pop3EndSession method
// can be called to send the "QUIT" and end the session (i.e. disconnect.)
// 
// Note: When making multiple calls to a Delete* method, it's best to turn off
// ImmediateDelete, and then manually call Pop3EndSession to finalize the
// deletions.
// 
// Also, any method call requiring communication with the POP3 server will
// automatically re-establish a session based on the current property settings.
//
func (z *MailMan) DeleteEmail(arg1 *Email) bool {

    v := C.CkMailMan_DeleteEmail(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the DeleteEmail method
func (z *MailMan) DeleteEmailAsync(arg1 *Email, c chan *Task) {

    hTask := C.CkMailMan_DeleteEmailAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Marks multiple emails on the POP3 server for deletion. (Any email on the server
// having a UIDL equal to a UIDL found in uidlArray is marked for deletion.) To complete
// the deletion of the emails, a "QUIT" message must be sent and the POP3 session
// ended. This will happen automatically when the ImmediateDelete property equals
// true, which is the default. If ImmediateDelete equals false, then the
// Pop3EndSession method can be called to send the "QUIT" and end the session (i.e.
// disconnect.)
// 
// Note: When making multiple calls to a Delete* method, it's best to turn off
// ImmediateDelete, and then manually call Pop3EndSession to finalize the
// deletions.
// 
// Also, any method call requiring communication with the POP3 server will
// automatically re-establish a session based on the current property settings.
//
func (z *MailMan) DeleteMultiple(arg1 *StringArray) bool {

    v := C.CkMailMan_DeleteMultiple(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the DeleteMultiple method
func (z *MailMan) DeleteMultipleAsync(arg1 *StringArray, c chan *Task) {

    hTask := C.CkMailMan_DeleteMultipleAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Fetches an email by message number. WARNING: Be very careful if calling this
// method. Message numbers are specific to a POP3 session. If a maildrop has (for
// example) 10 messages, the message numbers will be 1, 2, 3, ... 10. If message
// number 1 is deleted and a new POP3 session is established, there will be 9
// messages numbered 1, 2, 3, ... 9.
// 
// IMPORTANT: A POP3 connection must first be established by either calling
// Pop3BeginSession explicitly, or implicitly by calling some other method that
// automatically establishes the session. This method will not automatically
// establish a new POP3 session (because if it did, the message numbers would
// potentially be different than what the application expects).
//
func (z *MailMan) FetchByMsgnum(arg1 int) *Email {

    retObj := C.CkMailMan_FetchByMsgnum(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Asynchronous version of the FetchByMsgnum method
func (z *MailMan) FetchByMsgnumAsync(arg1 int, c chan *Task) {

    hTask := C.CkMailMan_FetchByMsgnumAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Fetches an email from the POP3 mail server given its UIDL. Calling this method
// does not remove the email from the server. A typical program might get the email
// headers from the POP3 server by calling GetAllHeaders or GetHeaders, and then
// fetch individual emails by UIDL.
// 
// Returns a null reference on failure.
//
func (z *MailMan) FetchEmail(arg1 string) *Email {
    cstr1 := C.CString(arg1)

    retObj := C.CkMailMan_FetchEmail(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Asynchronous version of the FetchEmail method
func (z *MailMan) FetchEmailAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMailMan_FetchEmailAsync(z.ckObj, cstr1)

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


// Fetches an email by UIDL and returns the MIME source of the email in a byte
// array.
func (z *MailMan) FetchMime(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkMailMan_FetchMime(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the FetchMime method
func (z *MailMan) FetchMimeAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMailMan_FetchMimeAsync(z.ckObj, cstr1)

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


// Fetches an email by UIDL and returns the MIME source of the email in uidl.
func (z *MailMan) FetchMimeBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMailMan_FetchMimeBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the FetchMimeBd method
func (z *MailMan) FetchMimeBdAsync(arg1 string, arg2 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMailMan_FetchMimeBdAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Fetches an email by message number and returns the MIME source of the email in a
// byte array. WARNING: Message sequend numbers are specific to a POP3 session. If
// a maildrop has (for example) 10 messages, the message numbers will be 1, 2, 3,
// ... 10. If message number 1 is deleted and a new POP3 session is established,
// there will be 9 messages numbered 1, 2, 3, ... 9.
// 
// IMPORTANT: A POP3 connection must first be established by either calling
// Pop3BeginSession explicitly, or implicitly by calling some other method that
// automatically establishes the session. This method will not automatically
// establish a new POP3 session (because if it did, the message numbers would
// potentially be different than what the application expects).
//
func (z *MailMan) FetchMimeByMsgnum(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkMailMan_FetchMimeByMsgnum(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the FetchMimeByMsgnum method
func (z *MailMan) FetchMimeByMsgnumAsync(arg1 int, c chan *Task) {

    hTask := C.CkMailMan_FetchMimeByMsgnumAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Given an array of UIDL strings, fetchs all the emails from the POP3 server whose
// UIDL is present in the array, and returns the emails in a bundle.
func (z *MailMan) FetchMultiple(arg1 *StringArray) *EmailBundle {

    retObj := C.CkMailMan_FetchMultiple(z.ckObj, arg1.ckObj)


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the FetchMultiple method
func (z *MailMan) FetchMultipleAsync(arg1 *StringArray, c chan *Task) {

    hTask := C.CkMailMan_FetchMultipleAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Given an array of UIDL strings, fetchs all the email headers from the POP3
// server whose UIDL is present in the array.
// 
// Note: The email objects returned in the bundle contain only headers. The
// attachments will be missing, and the bodies will be mostly missing (only the 1st
// numBodyLines lines of either the plain-text or HTML body will be present).
//
func (z *MailMan) FetchMultipleHeaders(arg1 *StringArray, arg2 int) *EmailBundle {

    retObj := C.CkMailMan_FetchMultipleHeaders(z.ckObj, arg1.ckObj, C.int(arg2))


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the FetchMultipleHeaders method
func (z *MailMan) FetchMultipleHeadersAsync(arg1 *StringArray, arg2 int, c chan *Task) {

    hTask := C.CkMailMan_FetchMultipleHeadersAsync(z.ckObj, arg1.ckObj, C.int(arg2))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Given an array of UIDL strings, fetchs all the emails from the POP3 server whose
// UIDL is present in the array, and returns the MIME source of each email in an
// "stringarray" -- an object containing a collection of strings. Returns a null
// reference on failure.
func (z *MailMan) FetchMultipleMime(arg1 *StringArray) *StringArray {

    retObj := C.CkMailMan_FetchMultipleMime(z.ckObj, arg1.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Asynchronous version of the FetchMultipleMime method
func (z *MailMan) FetchMultipleMimeAsync(arg1 *StringArray, c chan *Task) {

    hTask := C.CkMailMan_FetchMultipleMimeAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Fetches a single header by message number. Returns an email object on success,
// or a null reference on failure.
// 
// Note: The email objects returned in the bundle contain only headers. The
// attachments will be missing, and the bodies will be mostly missing (only the 1st
// messageNumber lines of either the plain-text or HTML body will be present).
// 
// Also Important:Message numbers are specific to a POP3 session (whereas UIDLs are
// valid across sessions). Be very careful when using this method.
//
func (z *MailMan) FetchSingleHeader(arg1 int, arg2 int) *Email {

    retObj := C.CkMailMan_FetchSingleHeader(z.ckObj, C.int(arg1), C.int(arg2))


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Asynchronous version of the FetchSingleHeader method
func (z *MailMan) FetchSingleHeaderAsync(arg1 int, arg2 int, c chan *Task) {

    hTask := C.CkMailMan_FetchSingleHeaderAsync(z.ckObj, C.int(arg1), C.int(arg2))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Fetches a single header by UIDL. Returns an email object on success, or a null
// reference on failure.
// 
// Note: The email objects returned in the bundle contain only headers. The
// attachments will be missing, and the bodies will be mostly missing (only the 1st
// uidl lines of either the plain-text or HTML body will be present).
//
func (z *MailMan) FetchSingleHeaderByUidl(arg1 int, arg2 string) *Email {
    cstr2 := C.CString(arg2)

    retObj := C.CkMailMan_FetchSingleHeaderByUidl(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Asynchronous version of the FetchSingleHeaderByUidl method
func (z *MailMan) FetchSingleHeaderByUidlAsync(arg1 int, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkMailMan_FetchSingleHeaderByUidlAsync(z.ckObj, C.int(arg1), cstr2)

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


// Returns all the emails from the POP3 server, but only the first numBodyLines lines of
// the body. Attachments are not returned. The emails returned in the bundle are
// valid email objects, the only difference is that the body is truncated to
// include only the top numBodyLines lines, and the attachments will be missing.
func (z *MailMan) GetAllHeaders(arg1 int) *EmailBundle {

    retObj := C.CkMailMan_GetAllHeaders(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the GetAllHeaders method
func (z *MailMan) GetAllHeadersAsync(arg1 int, c chan *Task) {

    hTask := C.CkMailMan_GetAllHeadersAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns a string array object containing a list of failed and invalid email
// addresses that have accumulated during SMTP sends. The list will not contain
// duplicates. Also, this only works with some SMTP servers -- not all SMTP servers
// check the validity of each email address.
// 
// Note: An SMTP server can only validate the email addresses within it's own
// domain. External email address are not verifiable at the time of sending.
//
func (z *MailMan) GetBadEmailAddrs() *StringArray {

    retObj := C.CkMailMan_GetBadEmailAddrs(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// If a partial email was obtained using GetHeaders or GetAllHeaders, this method
// will take the partial email as an argument, and download the full email from the
// server. A new email object (separate from the partial email) is returned. A null
// reference is returned on failure.
func (z *MailMan) GetFullEmail(arg1 *Email) *Email {

    retObj := C.CkMailMan_GetFullEmail(z.ckObj, arg1.ckObj)


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Asynchronous version of the GetFullEmail method
func (z *MailMan) GetFullEmailAsync(arg1 *Email, c chan *Task) {

    hTask := C.CkMailMan_GetFullEmailAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as the GetAllHeaders method, except only the emails from fromIndex to toIndex
// on the POP3 server are returned. The first email on the server is at index 0.
func (z *MailMan) GetHeaders(arg1 int, arg2 int, arg3 int) *EmailBundle {

    retObj := C.CkMailMan_GetHeaders(z.ckObj, C.int(arg1), C.int(arg2), C.int(arg3))


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the GetHeaders method
func (z *MailMan) GetHeadersAsync(arg1 int, arg2 int, arg3 int, c chan *Task) {

    hTask := C.CkMailMan_GetHeadersAsync(z.ckObj, C.int(arg1), C.int(arg2), C.int(arg3))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the number of emails on the POP3 server, or -1 for failure.
// 
// This method is identical to CheckEmail. It was added for clarity.
//
func (z *MailMan) GetMailboxCount() int {

    v := C.CkMailMan_GetMailboxCount(z.ckObj)


    return int(v)
}


// Asynchronous version of the GetMailboxCount method
func (z *MailMan) GetMailboxCountAsync(c chan *Task) {

    hTask := C.CkMailMan_GetMailboxCountAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns an XML document with information about the emails in a POP3 mailbox. The
// XML contains the UIDL and size (in bytes) of each email in the mailbox.
// return a string or nil if failed.
func (z *MailMan) GetMailboxInfoXml() *string {

    retStr := C.CkMailMan_getMailboxInfoXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetMailboxInfoXml method
func (z *MailMan) GetMailboxInfoXmlAsync(c chan *Task) {

    hTask := C.CkMailMan_GetMailboxInfoXmlAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the total combined size in bytes of all the emails in the POP3 mailbox.
// This is also known as the "mail drop" size. Returns -1 on failure.
func (z *MailMan) GetMailboxSize() uint {

    v := C.CkMailMan_GetMailboxSize(z.ckObj)


    return uint(v)
}


// Asynchronous version of the GetMailboxSize method
func (z *MailMan) GetMailboxSizeAsync(c chan *Task) {

    hTask := C.CkMailMan_GetMailboxSizeAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the POP3 server's SSL certificate. This is available after connecting
// via SSL to a POP3 server. (To use POP3 SSL, set the PopSsl property = true.)
// 
// Returns a null reference if no POP3 SSL certificate is available.
//
func (z *MailMan) GetPop3SslServerCert() *Cert {

    retObj := C.CkMailMan_GetPop3SslServerCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the list of successful email addresses in the last call to a mail
// sending method, such as SendEmail.
// 
// When an email is sent, the email addresses that were flagged invalid by the SMTP
// server are saved in a "bad email addresses" list within the mailman object, and
// the acceptable email addresses are saved in a "good email addresses" list
// (within the mailman object). These internal lists are automatically reset at the
// start of the next mail-sending method call. This allows for a program to know
// which email addresses were accepted and which were not.
// 
// Note: The AllOrNone property controls whether the mail-sending method, such as
// SendEmail, returns false (to indicate failure) if any single email address is
// rejected.
// 
// Note: An SMTP server can only be aware of invalid email addresses that are of
// the same domain. For example, the comcast.net mail servers would only be aware
// of what comcast.net email addresses are valid. All external email addresses are
// implicitly accepted because the server is simply forwarding the email towards
// the mail server controlling that domain.
//
func (z *MailMan) GetSentToEmailAddrs() *StringArray {

    retObj := C.CkMailMan_GetSentToEmailAddrs(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Returns the size of an email (including attachments) given the UIDL of the email
// on the POP3 server. Returns -1 for failure.
func (z *MailMan) GetSizeByUidl(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkMailMan_GetSizeByUidl(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the GetSizeByUidl method
func (z *MailMan) GetSizeByUidlAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMailMan_GetSizeByUidlAsync(z.ckObj, cstr1)

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


// If using SSL/TLS, this method returns the SMTP server's digital certificate used
// with the secure connection.
func (z *MailMan) GetSmtpSslServerCert() *Cert {

    retObj := C.CkMailMan_GetSmtpSslServerCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the UIDLs of the emails currently stored on the POP3 server.
func (z *MailMan) GetUidls() *StringArray {

    retObj := C.CkMailMan_GetUidls(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Asynchronous version of the GetUidls method
func (z *MailMan) GetUidlsAsync(c chan *Task) {

    hTask := C.CkMailMan_GetUidlsAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Contacts the SMTP server and determines if it supports the DSN (Delivery Status
// Notification) features specified by RFC 3461 and supported by the DsnEnvid,
// DsnNotify, and DsnRet properties. Returns true if the SMTP server supports
// DSN, otherwise returns false.
func (z *MailMan) IsSmtpDsnCapable() bool {

    v := C.CkMailMan_IsSmtpDsnCapable(z.ckObj)


    return v != 0
}


// Asynchronous version of the IsSmtpDsnCapable method
func (z *MailMan) IsSmtpDsnCapableAsync(c chan *Task) {

    hTask := C.CkMailMan_IsSmtpDsnCapableAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns true if the mailman is already unlocked, otherwise returns false.
func (z *MailMan) IsUnlocked() bool {

    v := C.CkMailMan_IsUnlocked(z.ckObj)


    return v != 0
}


// Provides information about what transpired in the last method called on this
// object instance. For many methods, there is no information. However, for some
// methods, details about what occurred can be obtained by getting the LastJsonData
// right after the method call returns.
func (z *MailMan) LastJsonData() *JsonObject {

    retObj := C.CkMailMan_LastJsonData(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Loads an email from a .eml file. (EML files contain the MIME source of an
// email.) Returns a null reference on failure.
// 
// Note: MHT files can be loaded into an email object by calling this method.
//
func (z *MailMan) LoadEml(arg1 string) *Email {
    cstr1 := C.CString(arg1)

    retObj := C.CkMailMan_LoadEml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Loads a .mbx file containing emails and returns an email bundle. If a Filter is
// present, only emails matching the filter are returned.
func (z *MailMan) LoadMbx(arg1 string) *EmailBundle {
    cstr1 := C.CString(arg1)

    retObj := C.CkMailMan_LoadMbx(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Creates and loads an email from a MIME string. Returns a null reference on
// failure.
func (z *MailMan) LoadMime(arg1 string) *Email {
    cstr1 := C.CString(arg1)

    retObj := C.CkMailMan_LoadMime(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Loads the caller of the task's async method.
func (z *MailMan) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkMailMan_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads an XML file containing a single email and returns an email object. Returns
// a null reference on failure.
func (z *MailMan) LoadXmlEmail(arg1 string) *Email {
    cstr1 := C.CString(arg1)

    retObj := C.CkMailMan_LoadXmlEmail(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Loads an XML string containing a single email and returns an email object.
// Returns a null reference on failure.
func (z *MailMan) LoadXmlEmailString(arg1 string) *Email {
    cstr1 := C.CString(arg1)

    retObj := C.CkMailMan_LoadXmlEmailString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Loads an XML file containing one or more emails and returns an email bundle. If
// a Filter is present, only emails matching the filter are returned. Returns a
// null reference on failure.
func (z *MailMan) LoadXmlFile(arg1 string) *EmailBundle {
    cstr1 := C.CString(arg1)

    retObj := C.CkMailMan_LoadXmlFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Loads from an XML string containing emails and returns an email bundle. If a
// Filter is present, only emails matching the filter are returned.
func (z *MailMan) LoadXmlString(arg1 string) *EmailBundle {
    cstr1 := C.CString(arg1)

    retObj := C.CkMailMan_LoadXmlString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Performs a DNS MX lookup to return the mail server hostname based on an email
// address.
// return a string or nil if failed.
func (z *MailMan) MxLookup(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkMailMan_mxLookup(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Performs a DNS MX lookup to return the list of mail server hostnames based on an
// email address. The primary server is at index 0. In most cases, there is only
// one mail server for a given email address.
func (z *MailMan) MxLookupAll(arg1 string) *StringArray {
    cstr1 := C.CString(arg1)

    retObj := C.CkMailMan_MxLookupAll(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Explicitly opens a connection to the SMTP server and authenticates (if a
// username/password was specified). Calling this method is optional because the
// SendEmail method and other mail-sending methods will automatically open the
// connection to the SMTP server if one is not already established.
// 
// Note: This method is the equivalent of calling SmtpConnect followed by
// SmtpAuthenticate.
//
func (z *MailMan) OpenSmtpConnection() bool {

    v := C.CkMailMan_OpenSmtpConnection(z.ckObj)


    return v != 0
}


// Asynchronous version of the OpenSmtpConnection method
func (z *MailMan) OpenSmtpConnectionAsync(c chan *Task) {

    hTask := C.CkMailMan_OpenSmtpConnectionAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Authenticates with the POP3 server using the property settings such as
// PopUsername, PopPassword, etc. This method should only be called after a
// successful call to Pop3Connect.
// 
// Note 1: The Pop3BeginSession method both connects and authenticates. It is the
// equivalent of calling Pop3Connect followed by Pop3Authenticate.
// 
// Note 2: All methods that communicate with the POP3 server, such as FetchEmail,
// will automatically connect and authenticate if not already connected and
// authenticated.
//
func (z *MailMan) Pop3Authenticate() bool {

    v := C.CkMailMan_Pop3Authenticate(z.ckObj)


    return v != 0
}


// Asynchronous version of the Pop3Authenticate method
func (z *MailMan) Pop3AuthenticateAsync(c chan *Task) {

    hTask := C.CkMailMan_Pop3AuthenticateAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Call to explicitly begin a POP3 session. It is not necessary to call this method
// because any method requiring an established POP3 session will automatically
// connect and login if a session is not already open.
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
//
func (z *MailMan) Pop3BeginSession() bool {

    v := C.CkMailMan_Pop3BeginSession(z.ckObj)


    return v != 0
}


// Asynchronous version of the Pop3BeginSession method
func (z *MailMan) Pop3BeginSessionAsync(c chan *Task) {

    hTask := C.CkMailMan_Pop3BeginSessionAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Explicitly establishes a connection to the POP3 server, which includes
// establishing a secure TLS channel if required, and receives the initial
// greeting. This method stops short of authenticating. The Pop3Authenticate method
// should be called after a successful call to this method.
// 
// Note 1: The Pop3BeginSession method both connects and authenticates. It is the
// equivalent of calling Pop3Connect followed by Pop3Authenticate.
// 
// Note 2: All methods that communicate with the POP3 server, such as FetchEmail,
// will automatically connect and authenticate if not already connected and
// authenticated.
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
//
func (z *MailMan) Pop3Connect() bool {

    v := C.CkMailMan_Pop3Connect(z.ckObj)


    return v != 0
}


// Asynchronous version of the Pop3Connect method
func (z *MailMan) Pop3ConnectAsync(c chan *Task) {

    hTask := C.CkMailMan_Pop3ConnectAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Call to explicitly end a POP3 session. If the ImmediateDelete property is set to
// false, and emails marked for deletion will be deleted at this time.
func (z *MailMan) Pop3EndSession() bool {

    v := C.CkMailMan_Pop3EndSession(z.ckObj)


    return v != 0
}


// Asynchronous version of the Pop3EndSession method
func (z *MailMan) Pop3EndSessionAsync(c chan *Task) {

    hTask := C.CkMailMan_Pop3EndSessionAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// This method is identical to Pop3EndSession, but no "QUIT" command is sent. The
// client simply disconnects from the POP3 server.
// 
// This method should always return true.
//
func (z *MailMan) Pop3EndSessionNoQuit() bool {

    v := C.CkMailMan_Pop3EndSessionNoQuit(z.ckObj)


    return v != 0
}


// Asynchronous version of the Pop3EndSessionNoQuit method
func (z *MailMan) Pop3EndSessionNoQuitAsync(c chan *Task) {

    hTask := C.CkMailMan_Pop3EndSessionNoQuitAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a NOOP command to the POP3 server. This may be a useful method to call
// periodically to keep a connection open, or to verify that the POP3 connection
// (session) is open and functioning.
func (z *MailMan) Pop3Noop() bool {

    v := C.CkMailMan_Pop3Noop(z.ckObj)


    return v != 0
}


// Asynchronous version of the Pop3Noop method
func (z *MailMan) Pop3NoopAsync(c chan *Task) {

    hTask := C.CkMailMan_Pop3NoopAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a RSET command to the POP3 server. If any messages have been marked as
// deleted by the POP3 server, they are unmarked. Calling Pop3Reset resets the POP3
// session to a valid, known starting point.
func (z *MailMan) Pop3Reset() bool {

    v := C.CkMailMan_Pop3Reset(z.ckObj)


    return v != 0
}


// Asynchronous version of the Pop3Reset method
func (z *MailMan) Pop3ResetAsync(c chan *Task) {

    hTask := C.CkMailMan_Pop3ResetAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a raw command to the POP3 server and returns the POP3 server's response.
// If non-us-ascii characters are included in command, then charset indicates the charset
// to be used in sending the command (such as "utf-8", "ansi", "iso-8859-1",
// "Shift_JIS", etc.)
// return a string or nil if failed.
func (z *MailMan) Pop3SendRawCommand(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkMailMan_pop3SendRawCommand(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the Pop3SendRawCommand method
func (z *MailMan) Pop3SendRawCommandAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkMailMan_Pop3SendRawCommandAsync(z.ckObj, cstr1, cstr2)

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


// A quick way to send an email to a single recipient without having to explicitly
// create an email object.
func (z *MailMan) QuickSend(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkMailMan_QuickSend(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Asynchronous version of the QuickSend method
func (z *MailMan) QuickSendAsync(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    hTask := C.CkMailMan_QuickSendAsync(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5)

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


// When an email is sent by calling SendEmail, it is first "rendered" according to
// the values of the email properties and contents. It may be digitally signed,
// encrypted, values substituted for replacement patterns, and header fields "Q"or
// "B" encoded as needed based on the email. The RenderToMime method performs the
// rendering, but without the actual sending. The MIME text produced is exactly
// what would be sent to the SMTP server had SendEmail been called. (The SendEmail
// method is effectively the same as calling RenderToMime followed by a call to
// SendRendered.)
// 
// The rendered MIME string is returned on success.
//
// return a string or nil if failed.
func (z *MailMan) RenderToMime(arg1 *Email) *string {

    retStr := C.CkMailMan_renderToMime(z.ckObj, arg1.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The same as RenderToMimeBytes, except the MIME is rendered into renderedMime. The
// rendered MIME is appended to renderedMime.
func (z *MailMan) RenderToMimeBd(arg1 *Email, arg2 *BinData) bool {

    v := C.CkMailMan_RenderToMimeBd(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// This method is the same as RenderToMime, but the MIME is returned in a byte
// array. If an email uses an 8bit or binary MIME encoding, then calling
// RenderToMime may introduce errors because it is not possible to return non-text
// binary data as a string. Therefore, calling RenderToMimeBytes is recommended
// over RenderToMime, unless it is assured that the email (MIME) does not use a
// binary encoding for non-text data.
func (z *MailMan) RenderToMimeBytes(arg1 *Email) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkMailMan_RenderToMimeBytes(z.ckObj, arg1.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// The same as RenderToMime, except the MIME is rendered into renderedMime. The rendered
// MIME is appended to renderedMime.
func (z *MailMan) RenderToMimeSb(arg1 *Email, arg2 *StringBuilder) bool {

    v := C.CkMailMan_RenderToMimeSb(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Sends a bundle of emails. This is identical to calling SendEmail for each email
// in the bundle.
// 
// If an error occurs when sending one of the emails in the bundle, it will
// continue with each subsequent email until each email in the bundle has been
// attempted (unless a fatal error occurs, in which case the send is aborted).
// 
// Because it is difficult or impossible to programmatically identify which emails
// in the bundle failed and which succeeded, it is best to write a loop that sends
// each email separately (via the SendEmail method).
//
func (z *MailMan) SendBundle(arg1 *EmailBundle) bool {

    v := C.CkMailMan_SendBundle(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the SendBundle method
func (z *MailMan) SendBundleAsync(arg1 *EmailBundle, c chan *Task) {

    hTask := C.CkMailMan_SendBundleAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a single email. The connection to the SMTP server will remain open so that
// a subsequent call to SendEmail (or other email-sending methods) can re-use the
// same connection. If any properties relating to the SMTP server are changed, such
// as SmtpHost, SmtpUsername, etc., then the next call to an email-sending method
// will automatically close the connection and re-establish a connection using the
// updated property settings.
// 
// Important: Some SMTP servers do not actually send the email until the connection
// is closed. In these cases, it is necessary to call CloseSmtpConnection for the
// mail to be sent. Most SMTP servers send the email immediately, and it is not
// required to close the connection.
// 
// GMail: If sending via smtp.gmail.com, then send with OAuth2 authentication if
// possible. Otherwise you will need to change your GMail account settings to allow
// for sending by less secure apps. See the links below.
// 
// Note: After sending email, information about what transpired is available via
// the LastJsonData method.
//
func (z *MailMan) SendEmail(arg1 *Email) bool {

    v := C.CkMailMan_SendEmail(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the SendEmail method
func (z *MailMan) SendEmailAsync(arg1 *Email, c chan *Task) {

    hTask := C.CkMailMan_SendEmailAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Provides complete control over the email that is sent. The MIME text passed in
// mimeSource (the MIME source of an email) is passed exactly as-is to the SMTP server.
// The recipients is a comma separated list of recipient email addresses. The fromAddr is the
// reverse-path email address. This is where bounced email (non-delivery reports)
// will be delivered. It may be different than the "From" header field in the mimeSource.
// 
// To understand how the fromAddr and recipients relate to the email addresses found in the
// MIME headers (FROM, TO, CC), see the link below entitled "SMTP Protocol in a
// Nutshell". The fromAddr is what is passed to the SMTP server in the "MAIL FROM"
// command. The recipients are the email addresses passed in "RCPT TO" commands. These
// are usually the same email addresses found in the MIME headers, but need not be
// (unless the SMTP server enforces policies that require them to be the same).
//
func (z *MailMan) SendMime(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkMailMan_SendMime(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the SendMime method
func (z *MailMan) SendMimeAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkMailMan_SendMimeAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// This method is the same as SendMimeBytes, except the MIME is passed in an object
// (mimeData) rather than explicitly passing the bytes.
func (z *MailMan) SendMimeBd(arg1 string, arg2 string, arg3 *BinData) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMailMan_SendMimeBd(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendMimeBd method
func (z *MailMan) SendMimeBdAsync(arg1 string, arg2 string, arg3 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkMailMan_SendMimeBdAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// This method is the same as SendMime, except the MIME is passed in a byte array.
// This can be important if the MIME uses a binary encoding, or if a DKIM/DomainKey
// signature is included.
// 
// To understand how the fromAddr and recipients relate to the email addresses found in the
// MIME headers (FROM, TO, CC), see the link below entitled "SMTP Protocol in a
// Nutshell". The fromAddr is what is passed to the SMTP server in the "MAIL FROM"
// command. The recipients are the email addresses passed in "RCPT TO" commands. These
// are usually the same email addresses found in the MIME headers, but need not be
// (unless the SMTP server enforces policies that require them to be the same).
//
func (z *MailMan) SendMimeBytes(arg1 string, arg2 string, arg3 []byte) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkMailMan_SendMimeBytes(z.ckObj, cstr1, cstr2, ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// Asynchronous version of the SendMimeBytes method
func (z *MailMan) SendMimeBytesAsync(arg1 string, arg2 string, arg3 []byte, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    hTask := C.CkMailMan_SendMimeBytesAsync(z.ckObj, cstr1, cstr2, ckbyd3)

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


// Same as SendMime, but the recipient list is read from a text file (distListFilename)
// containing one email address per line.
func (z *MailMan) SendMimeToList(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkMailMan_SendMimeToList(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the SendMimeToList method
func (z *MailMan) SendMimeToListAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkMailMan_SendMimeToListAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Explicitly specifies the certificate to be used for decrypting encrypted email.
func (z *MailMan) SetDecryptCert(arg1 *Cert) bool {

    v := C.CkMailMan_SetDecryptCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Explicitly specifies the certificate and associated private key to be used for
// decrypting S/MIME encrypted email.
// 
// Note: In most cases, it is easier to call AddPfxSourceFile or AddPfxSourceData
// to provide the required cert and private key. On Windows systems where the
// certificate + private key has already been installed in the default certificate
// store, nothing needs to be done -- the mailman will automatically locate and use
// the required cert + private key.
//
func (z *MailMan) SetDecryptCert2(arg1 *Cert, arg2 *PrivateKey) bool {

    v := C.CkMailMan_SetDecryptCert2(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Provides a more secure way of setting either the POP3 or SMTP password. The protocol
// can be "pop3" or "smtp". When the protocol is "pop3", this is equivalent to setting
// the PopPassword property. When protocol is "smtp", this is equivalent to setting the
// SmtpPassword property.
func (z *MailMan) SetPassword(arg1 string, arg2 *SecureString) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMailMan_SetPassword(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the client-side certificate to be used with SSL connections. This is
// typically not required, as most SSL connections are such that only the server is
// authenticated while the client remains unauthenticated.
func (z *MailMan) SetSslClientCert(arg1 *Cert) bool {

    v := C.CkMailMan_SetSslClientCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Allows for a client-side certificate to be used for the SSL / TLS connection.
func (z *MailMan) SetSslClientCertPem(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMailMan_SetSslClientCertPem(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Allows for a client-side certificate to be used for the SSL / TLS connection.
func (z *MailMan) SetSslClientCertPfx(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMailMan_SetSslClientCertPfx(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Authenticates with the SMTP server using the property settings such as
// SmtpUsername, SmtpPassword, etc. This method should only be called after a
// successful call to SmtpConnect.
// 
// Note 1: The OpenSmtpConnection method both connects and authenticates. It is the
// equivalent of calling SmtpConnect followed by SmtpAuthenticate.
// 
// Note 2: All methods that communicate with the SMTP server, such as SendEmail,
// will automatically connect and authenticate if not already connected and
// authenticated.
//
func (z *MailMan) SmtpAuthenticate() bool {

    v := C.CkMailMan_SmtpAuthenticate(z.ckObj)


    return v != 0
}


// Asynchronous version of the SmtpAuthenticate method
func (z *MailMan) SmtpAuthenticateAsync(c chan *Task) {

    hTask := C.CkMailMan_SmtpAuthenticateAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Explicitly establishes a connection to the SMTP server, which includes
// establishing a secure TLS channel if required, and receives the initial
// greeting. This method stops short of authenticating. The SmtpAuthenticate method
// should be called after a successful call to this method.
// 
// Note 1: The OpenSmtpConnection method both connects and authenticates. It is the
// equivalent of calling SmtpConnect followed by SmtpAuthenticate.
// 
// Note 2: All methods that communicate with the SMTP server, such as SendEmail,
// will automatically connect and authenticate if not already connected and
// authenticated.
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
//
func (z *MailMan) SmtpConnect() bool {

    v := C.CkMailMan_SmtpConnect(z.ckObj)


    return v != 0
}


// Asynchronous version of the SmtpConnect method
func (z *MailMan) SmtpConnectAsync(c chan *Task) {

    hTask := C.CkMailMan_SmtpConnectAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a no-op to the SMTP server. Calling this method is good for testing to see
// if the connection to the SMTP server is working and valid. The SmtpNoop method
// will automatically establish the SMTP connection if it does not already exist.
func (z *MailMan) SmtpNoop() bool {

    v := C.CkMailMan_SmtpNoop(z.ckObj)


    return v != 0
}


// Asynchronous version of the SmtpNoop method
func (z *MailMan) SmtpNoopAsync(c chan *Task) {

    hTask := C.CkMailMan_SmtpNoopAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an RSET command to the SMTP server. This method is rarely needed. The RSET
// command resets the state of the connection to the SMTP server to the initial
// state (so that the component can proceed with sending a new email). The
// SmtpReset method would only be needed if a mail-sending method failed and left
// the connection with the SMTP server open and in a non-initial state. (A
// situation that is probably not even possible with the Chilkat mail component.)
func (z *MailMan) SmtpReset() bool {

    v := C.CkMailMan_SmtpReset(z.ckObj)


    return v != 0
}


// Asynchronous version of the SmtpReset method
func (z *MailMan) SmtpResetAsync(c chan *Task) {

    hTask := C.CkMailMan_SmtpResetAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a raw command to the SMTP server and returns the SMTP server's response.
// If non-us-ascii characters are included in command, then charset indicates the charset
// to be used in sending the command (such as "utf-8", "ansi", "iso-8859-1",
// "Shift_JIS", etc.)
// 
// If bEncodeBase64 is true, then the response is returned in Base64-encoded format.
// Otherwise the raw response is returned.
//
// return a string or nil if failed.
func (z *MailMan) SmtpSendRawCommand(arg1 string, arg2 string, arg3 bool) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retStr := C.CkMailMan_smtpSendRawCommand(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the SmtpSendRawCommand method
func (z *MailMan) SmtpSendRawCommandAsync(arg1 string, arg2 string, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkMailMan_SmtpSendRawCommandAsync(z.ckObj, cstr1, cstr2, b3)

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


// Authenticates with the SSH server using public-key authentication. The
// corresponding public key must have been installed on the SSH server for the
// sshLogin. Authentication will succeed if the matching sshUsername is provided.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *MailMan) SshAuthenticatePk(arg1 string, arg2 *SshKey) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMailMan_SshAuthenticatePk(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SshAuthenticatePk method
func (z *MailMan) SshAuthenticatePkAsync(arg1 string, arg2 *SshKey, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMailMan_SshAuthenticatePkAsync(z.ckObj, cstr1, arg2.ckObj)

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
// AuthenticatePk to authenticate.
// 
// Note: Once the SSH tunnel is setup by calling SshTunnel and SshAuthenticatePw
// (or SshAuthenticatePk), all underlying communcations with the POP3 or SMTP
// server use the SSH tunnel. No changes in programming are required other than
// making two initial calls to setup the tunnel.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *MailMan) SshAuthenticatePw(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMailMan_SshAuthenticatePw(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SshAuthenticatePw method
func (z *MailMan) SshAuthenticatePwAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkMailMan_SshAuthenticatePwAsync(z.ckObj, cstr1, cstr2)

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


// Closes the SSH tunnel for SMTP or POP3.
func (z *MailMan) SshCloseTunnel() bool {

    v := C.CkMailMan_SshCloseTunnel(z.ckObj)


    return v != 0
}


// Asynchronous version of the SshCloseTunnel method
func (z *MailMan) SshCloseTunnelAsync(c chan *Task) {

    hTask := C.CkMailMan_SshCloseTunnelAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Connects to an SSH server and creates a tunnel for SMTP or POP3. The sshHostname is the
// hostname (or IP address) of the SSH server. The sshPort is typically 22, which is
// the standard SSH port number.
// 
// An SSH tunneling (port forwarding) session always begins by first calling
// SshTunnel to connect to the SSH server, followed by calling either
// SshAuthenticatePw or SshAuthenticatePk to authenticate.
// 
// Note: Once the SSH tunnel is setup by calling SshOpenTunnel and
// SshAuthenticatePw (or SshAuthenticatePk), all underlying communcations with the
// SMTP or POP3 server use the SSH tunnel. No changes in programming are required
// other than making two initial calls to setup the tunnel.
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
//
func (z *MailMan) SshOpenTunnel(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMailMan_SshOpenTunnel(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SshOpenTunnel method
func (z *MailMan) SshOpenTunnelAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMailMan_SshOpenTunnelAsync(z.ckObj, cstr1, C.int(arg2))

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


// Downloads and removes all email from a POP3 server. A bundle containing the
// emails is returned. A null reference is returned on failure.
func (z *MailMan) TransferMail() *EmailBundle {

    retObj := C.CkMailMan_TransferMail(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &EmailBundle{retObj}
}


// Asynchronous version of the TransferMail method
func (z *MailMan) TransferMailAsync(c chan *Task) {

    hTask := C.CkMailMan_TransferMailAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Same as FetchMultipleMime except that the downloaded emails are also deleted
// from the server. Returns a null reference on failure.
func (z *MailMan) TransferMultipleMime(arg1 *StringArray) *StringArray {

    retObj := C.CkMailMan_TransferMultipleMime(z.ckObj, arg1.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Asynchronous version of the TransferMultipleMime method
func (z *MailMan) TransferMultipleMimeAsync(arg1 *StringArray, c chan *Task) {

    hTask := C.CkMailMan_TransferMultipleMimeAsync(z.ckObj, arg1.ckObj)


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
// (or ASP / ASP.NET page). It is very fast and has negligible overhead. An
// arbitrary string, such as "Hello World" may be passed to automatically begin a
// fully-functional 30-day trial.
// 
// A valid purchased unlock code for this object will always included the substring
// "MAIL", or can be a Bundle unlock code.
//
func (z *MailMan) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMailMan_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds an XML certificate vault to the object's internal list of sources to be
// searched for certificates and private keys when encrypting/decrypting or
// signing/verifying. Unlike the AddPfxSourceData and AddPfxSourceFile methods,
// only a single XML certificate vault can be used. If UseCertVault is called
// multiple times, only the last certificate vault will be used, as each call to
// UseCertVault will replace the certificate vault provided in previous calls.
func (z *MailMan) UseCertVault(arg1 *XmlCertVault) bool {

    v := C.CkMailMan_UseCertVault(z.ckObj, arg1.ckObj)


    return v != 0
}


// Uses an existing SSH tunnel for the connections to the POP3 andSMTP servers.
// This method is identical to the UseSshTunnel method, except the SSH connection
// is obtained from an SSH object instead of a Socket object.
// 
// Uses an existing SSH tunnel. This is useful for sharing an existing SSH tunnel
// connection wth other objects. (SSH is a protocol where the tunnel contains many
// logical channels. SMTP and POP3 connections can exist simultaneously within a
// single SSH tunnel as SSH channels.)
//
func (z *MailMan) UseSsh(arg1 *Ssh) bool {

    v := C.CkMailMan_UseSsh(z.ckObj, arg1.ckObj)


    return v != 0
}


// Uses an existing SSH tunnel. This is useful for sharing an existing SSH tunnel
// connection wth other objects. (SSH is a protocol where the tunnel contains many
// logical channels. SMTP and POP3 connections can exist simultaneously within a
// single SSH tunnel as SSH channels.)
func (z *MailMan) UseSshTunnel(arg1 *Socket) bool {

    v := C.CkMailMan_UseSshTunnel(z.ckObj, arg1.ckObj)


    return v != 0
}


// Return true if a TCP/IP connection can be established with the POP3 server,
// otherwise returns false.
func (z *MailMan) VerifyPopConnection() bool {

    v := C.CkMailMan_VerifyPopConnection(z.ckObj)


    return v != 0
}


// Asynchronous version of the VerifyPopConnection method
func (z *MailMan) VerifyPopConnectionAsync(c chan *Task) {

    hTask := C.CkMailMan_VerifyPopConnectionAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Return true if a TCP/IP connection and login is successful with the POP3
// server. Otherwise return false.
func (z *MailMan) VerifyPopLogin() bool {

    v := C.CkMailMan_VerifyPopLogin(z.ckObj)


    return v != 0
}


// Asynchronous version of the VerifyPopLogin method
func (z *MailMan) VerifyPopLoginAsync(c chan *Task) {

    hTask := C.CkMailMan_VerifyPopLoginAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Initiates sending an email, but aborts just after passing all recipients (TO,
// CC, BCC) to the SMTP server. This allows your program to collect email addresses
// flagged as invalid by the SMTP server.
// 
// Important: Please read this blog post before using this
// method:http://www.cknotes.com/?p=249
// <http://www.cknotes.com/?p=249>
//
func (z *MailMan) VerifyRecips(arg1 *Email, arg2 *StringArray) bool {

    v := C.CkMailMan_VerifyRecips(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Asynchronous version of the VerifyRecips method
func (z *MailMan) VerifyRecipsAsync(arg1 *Email, arg2 *StringArray, c chan *Task) {

    hTask := C.CkMailMan_VerifyRecipsAsync(z.ckObj, arg1.ckObj, arg2.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Return true if a TCP/IP connection can be established with the SMTP server,
// otherwise returns false.
func (z *MailMan) VerifySmtpConnection() bool {

    v := C.CkMailMan_VerifySmtpConnection(z.ckObj)


    return v != 0
}


// Asynchronous version of the VerifySmtpConnection method
func (z *MailMan) VerifySmtpConnectionAsync(c chan *Task) {

    hTask := C.CkMailMan_VerifySmtpConnectionAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Return true if a TCP/IP connection and login is successful with the SMTP
// server. Otherwise returns false.
func (z *MailMan) VerifySmtpLogin() bool {

    v := C.CkMailMan_VerifySmtpLogin(z.ckObj)


    return v != 0
}


// Asynchronous version of the VerifySmtpLogin method
func (z *MailMan) VerifySmtpLoginAsync(c chan *Task) {

    hTask := C.CkMailMan_VerifySmtpLoginAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


