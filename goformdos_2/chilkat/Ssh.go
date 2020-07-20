// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkSsh.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewSsh() *Ssh {
	return &Ssh{C.CkSsh_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Ssh) DisposeSsh() {
    if z != nil {
	C.CkSsh_Dispose(z.ckObj)
	}
}




func (z *Ssh) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkSsh_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Ssh) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkSsh_setExternalProgress(z.ckObj,1)
}

func (z *Ssh) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkSsh_setExternalProgress(z.ckObj,1)
}

func (z *Ssh) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkSsh_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Ssh) AbortCurrent() bool {
    v := int(C.CkSsh_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Ssh) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putAbortCurrent(z.ckObj,v)
}

// property: AuthFailReason
// Set to one of the following values if a call to AuthenticatePw, AuthenticatePk,
// or AuthenticatePwPk returns a failed status.
//     1: Transport failure. This is a failure to communicate with the server (i.e.
//     the connection was lost, or a read or write failed or timed out).
//     2: Invalid key for public key authentication. The key was not a valid
//     format, not a valid key, not a private key, or not the right type of key.
//     3: No matching authentication methods were available.
//     4: SSH authentication protocol error - an unexpected or invalid message was
//     received.
//     5: The incorrect password or private key was provided.
//     6: Already authenticated. The SSH session is already authenticated.
//     7: Password change request: The server requires the password to be changed.
func (z *Ssh) AuthFailReason() int {
    return int(C.CkSsh_getAuthFailReason(z.ckObj))
}

// property: CaretControl
// Controls whether the caret character '^' is interpreted as indicating a control
// character. The default value of this property is false. If set to true, then
// the following sequences are interpreted as control characters in any string
// passed to SendReqExec or ChannelSendString.
// ^@ 	00 	00 	NUL  	Null
// ^A 	01 	01 	SOH  	Start of Heading
// ^B 	02 	02 	STX  	Start of Text
// ^C 	03 	03 	ETX  	End of Text
// ^D 	04 	04 	EOT  	End of Transmission
// ^E 	05 	05 	ENQ  	Enquiry
// ^F 	06 	06 	ACK  	Acknowledge
// ^G 	07 	07 	BEL  	Bell
// ^H 	08 	08 	BS  	Backspace
// ^I 	09 	09 	HT  	Character Tabulation, Horizontal Tabulation
// ^J 	10 	0A 	LF  	Line Feed
// ^K 	11 	0B 	VT  	Line Tabulation, Vertical Tabulation
// ^L 	12 	0C 	FF  	Form Feed
// ^M 	13 	0D 	CR  	Carriage Return
// ^N 	14 	0E 	SO  	Shift Out
// ^O 	15 	0F 	SI  	Shift In
// ^P 	16 	10 	DLE  	Data Link Escape
// ^Q 	17 	11 	DC1  	Device Control One (XON)
// ^R 	18 	12 	DC2  	Device Control Two
// ^S 	19 	13 	DC3  	Device Control Three (XOFF)
// ^T 	20 	14 	DC4  	Device Control Four
// ^U 	21 	15 	NAK  	Negative Acknowledge
// ^V 	22 	16 	SYN  	Synchronous Idle
// ^W 	23 	17 	ETB  	End of Transmission Block
// ^X 	24 	18 	CAN  	Cancel
// ^Y 	25 	19 	EM  	End of medium
// ^Z 	26 	1A 	SUB  	Substitute
// ^[ 	27 	1B 	ESC  	Escape
// ^\ 	28 	1C 	FS  	File Separator
// ^] 	29 	1D 	GS  	Group Separator
// ^^ 	30 	1E 	RS  	Record Separator
// ^_ 	31 	1F 	US  	Unit Separator
// ^? 	127 	7F 	DEL  	Delete
func (z *Ssh) CaretControl() bool {
    v := int(C.CkSsh_getCaretControl(z.ckObj))
    return v != 0
}
// property setter: CaretControl
func (z *Ssh) SetCaretControl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putCaretControl(z.ckObj,v)
}

// property: ChannelOpenFailCode
// If a request to open a channel fails, this property contains a code that
// identifies the reason for failure. The reason codes are defined in RFC 4254 and
// are reproduced here:
//              Symbolic name                           reason code
//              -------------                           -----------
//             SSH_OPEN_ADMINISTRATIVELY_PROHIBITED          1
//             SSH_OPEN_CONNECT_FAILED                       2
//             SSH_OPEN_UNKNOWN_CHANNEL_TYPE                 3
//             SSH_OPEN_RESOURCE_SHORTAGE                    4
func (z *Ssh) ChannelOpenFailCode() int {
    return int(C.CkSsh_getChannelOpenFailCode(z.ckObj))
}

// property: ChannelOpenFailReason
// The descriptive text corresponding to the ChannelOpenFailCode property.
func (z *Ssh) ChannelOpenFailReason() string {
    return C.GoString(C.CkSsh_channelOpenFailReason(z.ckObj))
}

// property: ClientIdentifier
// The client-identifier string to be used when connecting to an SSH/SFTP server.
// Defaults to "SSH-2.0-PuTTY_Release_0.63". (This string is used to mimic PuTTY
// because some servers are known to disconnect from clients with unknown client
// identifiers.)
func (z *Ssh) ClientIdentifier() string {
    return C.GoString(C.CkSsh_clientIdentifier(z.ckObj))
}
// property setter: ClientIdentifier
func (z *Ssh) SetClientIdentifier(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putClientIdentifier(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
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
func (z *Ssh) ClientIpAddress() string {
    return C.GoString(C.CkSsh_clientIpAddress(z.ckObj))
}
// property setter: ClientIpAddress
func (z *Ssh) SetClientIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putClientIpAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ClientPort
// Normally left at the default value of 0, in which case a unique port is assigned
// with a value between 1024 and 5000. This property would only be changed if it is
// specifically required. For example, one customer's requirements are as follows:
// 
//     "I have to connect to a Siemens PLC IP server on a technical network. This
//     machine expects that I connect to its server from a specific IP address using a
//     specific port otherwise the build in security disconnect the IP connection."
//
func (z *Ssh) ClientPort() int {
    return int(C.CkSsh_getClientPort(z.ckObj))
}
// property setter: ClientPort
func (z *Ssh) SetClientPort(value int) {
    C.CkSsh_putClientPort(z.ckObj,C.int(value))
}

// property: ConnectTimeoutMs
// Maximum number of milliseconds to wait when connecting to an SSH server.
func (z *Ssh) ConnectTimeoutMs() int {
    return int(C.CkSsh_getConnectTimeoutMs(z.ckObj))
}
// property setter: ConnectTimeoutMs
func (z *Ssh) SetConnectTimeoutMs(value int) {
    C.CkSsh_putConnectTimeoutMs(z.ckObj,C.int(value))
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
func (z *Ssh) DebugLogFilePath() string {
    return C.GoString(C.CkSsh_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Ssh) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DisconnectCode
// If the SSH server sent a DISCONNECT message when closing the connection, this
// property contains the "reason code" as specified in RFC 4253:
//            Symbolic name                                reason code
//            -------------                                -----------
//       SSH_DISCONNECT_HOST_NOT_ALLOWED_TO_CONNECT             1
//       SSH_DISCONNECT_PROTOCOL_ERROR                          2
//       SSH_DISCONNECT_KEY_EXCHANGE_FAILED                     3
//       SSH_DISCONNECT_RESERVED                                4
//       SSH_DISCONNECT_MAC_ERROR                               5
//       SSH_DISCONNECT_COMPRESSION_ERROR                       6
//       SSH_DISCONNECT_SERVICE_NOT_AVAILABLE                   7
//       SSH_DISCONNECT_PROTOCOL_VERSION_NOT_SUPPORTED          8
//       SSH_DISCONNECT_HOST_KEY_NOT_VERIFIABLE                 9
//       SSH_DISCONNECT_CONNECTION_LOST                        10
//       SSH_DISCONNECT_BY_APPLICATION                         11
//       SSH_DISCONNECT_TOO_MANY_CONNECTIONS                   12
//       SSH_DISCONNECT_AUTH_CANCELLED_BY_USER                 13
//       SSH_DISCONNECT_NO_MORE_AUTH_METHODS_AVAILABLE         14
//       SSH_DISCONNECT_ILLEGAL_USER_NAME                      15
func (z *Ssh) DisconnectCode() int {
    return int(C.CkSsh_getDisconnectCode(z.ckObj))
}

// property: DisconnectReason
// If the SSH/ server sent a DISCONNECT message when closing the connection, this
// property contains a descriptive string for the "reason code" as specified in RFC
// 4253.
func (z *Ssh) DisconnectReason() string {
    return C.GoString(C.CkSsh_disconnectReason(z.ckObj))
}

// property: EnableCompression
// Enables or disables the use of compression w/ the SSH connection. The default
// value is true, meaning that compression is used if the server supports it.
// 
// Some older SSH servers have been found that claim to support compression, but
// actually fail when compression is used. PuTTY does not enable compression by
// default. If trouble is encountered where the SSH server disconnects immediately
// after the connection is seemingly established (i.e. during authentication), then
// check to see if disabling compression resolves the problem.
//
func (z *Ssh) EnableCompression() bool {
    v := int(C.CkSsh_getEnableCompression(z.ckObj))
    return v != 0
}
// property setter: EnableCompression
func (z *Ssh) SetEnableCompression(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putEnableCompression(z.ckObj,v)
}

// property: ForceCipher
// Set to one of the following encryption algorithms to force that cipher to be
// used. By default, the component will automatically choose the first cipher
// supported by the server in the order listed here: "aes256-ctr", "aes128-ctr",
// "aes256-cbc", "aes128-cbc", "twofish256-cbc", "twofish128-cbc", "blowfish-cbc",
// "3des-cbc", "arcfour128", "arcfour256". (If blowfish is chosen, the encryption
// strength is 128 bits.)
// 
// Important: If this is property is set and the server does NOT support then
// encryption algorithm, then the Connect will fail.
//
func (z *Ssh) ForceCipher() string {
    return C.GoString(C.CkSsh_forceCipher(z.ckObj))
}
// property setter: ForceCipher
func (z *Ssh) SetForceCipher(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putForceCipher(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// This is the number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any SSH operation prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
func (z *Ssh) HeartbeatMs() int {
    return int(C.CkSsh_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Ssh) SetHeartbeatMs(value int) {
    C.CkSsh_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: HostKeyAlg
// Indicates the preferred host key algorithm to be used in establishing the SSH
// secure connection. The default is "DSS". It may be changed to "RSA" if needed.
// Chilkat recommends not changing this unless a problem warrants the change.
func (z *Ssh) HostKeyAlg() string {
    return C.GoString(C.CkSsh_hostKeyAlg(z.ckObj))
}
// property setter: HostKeyAlg
func (z *Ssh) SetHostKeyAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putHostKeyAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HostKeyFingerprint
// Set after connecting to an SSH server. The format of the fingerprint looks like
// this: "ssh-rsa 1024 68:ff:d1:4e:6c:ff:d7:b0:d6:58:73:85:07:bc:2e:d5"
func (z *Ssh) HostKeyFingerprint() string {
    return C.GoString(C.CkSsh_hostKeyFingerprint(z.ckObj))
}

// property: HttpProxyAuthMethod
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy authentication method name. Valid choices are "Basic" or "NTLM".
func (z *Ssh) HttpProxyAuthMethod() string {
    return C.GoString(C.CkSsh_httpProxyAuthMethod(z.ckObj))
}
// property setter: HttpProxyAuthMethod
func (z *Ssh) SetHttpProxyAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putHttpProxyAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyDomain
// The NTLM authentication domain (optional) if NTLM authentication is used w/ the
// HTTP proxy.
func (z *Ssh) HttpProxyDomain() string {
    return C.GoString(C.CkSsh_httpProxyDomain(z.ckObj))
}
// property setter: HttpProxyDomain
func (z *Ssh) SetHttpProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putHttpProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyHostname
// If an HTTP proxy is to be used, set this property to the HTTP proxy hostname or
// IPv4 address (in dotted decimal notation).
func (z *Ssh) HttpProxyHostname() string {
    return C.GoString(C.CkSsh_httpProxyHostname(z.ckObj))
}
// property setter: HttpProxyHostname
func (z *Ssh) SetHttpProxyHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putHttpProxyHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPassword
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy password.
func (z *Ssh) HttpProxyPassword() string {
    return C.GoString(C.CkSsh_httpProxyPassword(z.ckObj))
}
// property setter: HttpProxyPassword
func (z *Ssh) SetHttpProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putHttpProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPort
// If an HTTP proxy is to be used, set this property to the HTTP proxy port number.
// (Two commonly used HTTP proxy ports are 8080 and 3128.)
func (z *Ssh) HttpProxyPort() int {
    return int(C.CkSsh_getHttpProxyPort(z.ckObj))
}
// property setter: HttpProxyPort
func (z *Ssh) SetHttpProxyPort(value int) {
    C.CkSsh_putHttpProxyPort(z.ckObj,C.int(value))
}

// property: HttpProxyUsername
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy login name.
func (z *Ssh) HttpProxyUsername() string {
    return C.GoString(C.CkSsh_httpProxyUsername(z.ckObj))
}
// property setter: HttpProxyUsername
func (z *Ssh) SetHttpProxyUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putHttpProxyUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IdleTimeoutMs
// Causes SSH operations to fail when progress for sending or receiving data halts
// for more than this number of milliseconds. Setting IdleTimeoutMs = 0 (the
// default) allows the application to wait indefinitely.
func (z *Ssh) IdleTimeoutMs() int {
    return int(C.CkSsh_getIdleTimeoutMs(z.ckObj))
}
// property setter: IdleTimeoutMs
func (z *Ssh) SetIdleTimeoutMs(value int) {
    C.CkSsh_putIdleTimeoutMs(z.ckObj,C.int(value))
}

// property: IsConnected
// Returns true if the component is connected to an SSH server.
// 
// Note: The IsConnected property is set to true after successfully completing
// the Connect method call. The IsConnected property will only be set to false by
// calling Disconnect, or by the failure of another method call such that the
// disconnection is detected.
//
func (z *Ssh) IsConnected() bool {
    v := int(C.CkSsh_getIsConnected(z.ckObj))
    return v != 0
}

// property: KeepSessionLog
// Controls whether communications to/from the SSH server are saved to the
// SessionLog property. The default value is false. If this property is set to
// true, the contents of the SessionLog property will continuously grow as SSH
// activity transpires. The purpose of the KeepSessionLog / SessionLog properties
// is to help in debugging any future problems that may arise.
func (z *Ssh) KeepSessionLog() bool {
    v := int(C.CkSsh_getKeepSessionLog(z.ckObj))
    return v != 0
}
// property setter: KeepSessionLog
func (z *Ssh) SetKeepSessionLog(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putKeepSessionLog(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Ssh) LastErrorHtml() string {
    return C.GoString(C.CkSsh_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Ssh) LastErrorText() string {
    return C.GoString(C.CkSsh_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Ssh) LastErrorXml() string {
    return C.GoString(C.CkSsh_lastErrorXml(z.ckObj))
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
func (z *Ssh) LastMethodSuccess() bool {
    v := int(C.CkSsh_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Ssh) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putLastMethodSuccess(z.ckObj,v)
}

// property: MaxPacketSize
// The maximum packet length to be used in the SSH transport protocol. The default
// value is 8192.
// 
// Note: If a large amount of data is to be flowing to/from the SSH server, then
// setting the MaxPacketSize equal to 32768 may improve performance. For those
// familiar with the inner workings of the SSH protocol, this is the "maximum
// packet size" value that is sent in the SSH_MSG_CHANNEL_OPEN message as defined
// in RFC 4254.
//
func (z *Ssh) MaxPacketSize() int {
    return int(C.CkSsh_getMaxPacketSize(z.ckObj))
}
// property setter: MaxPacketSize
func (z *Ssh) SetMaxPacketSize(value int) {
    C.CkSsh_putMaxPacketSize(z.ckObj,C.int(value))
}

// property: NumOpenChannels
// The number of currently open channels.
func (z *Ssh) NumOpenChannels() int {
    return int(C.CkSsh_getNumOpenChannels(z.ckObj))
}

// property: PasswordChangeRequested
// Set by the AuthenticatePw and AuthenticatePwPk methods. If the authenticate
// method returns a failed status, and this property is set to true, then it
// indicates the server requested a password change. In this case, re-call the
// authenticate method, but provide both the old and new passwords in the following
// format, where vertical bar characters encapsulate the old and new passwords:
// 
//     |oldPassword|newPassword|
//
func (z *Ssh) PasswordChangeRequested() bool {
    v := int(C.CkSsh_getPasswordChangeRequested(z.ckObj))
    return v != 0
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *Ssh) PreferIpv6() bool {
    v := int(C.CkSsh_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *Ssh) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putPreferIpv6(z.ckObj,v)
}

// property: ReadTimeoutMs
// The maximum amount of time to allow for reading messages/data from the SSH
// server. This is different from the IdleTimeoutMs property. The IdleTimeoutMs is
// the maximum amount of time to wait while no incoming data is arriving. The
// ReadTimeoutMs is the maximum amount of time to allow for reading data even if
// data is continuing to arrive. The default value of 0 indicates an infinite
// timeout value.
func (z *Ssh) ReadTimeoutMs() int {
    return int(C.CkSsh_getReadTimeoutMs(z.ckObj))
}
// property setter: ReadTimeoutMs
func (z *Ssh) SetReadTimeoutMs(value int) {
    C.CkSsh_putReadTimeoutMs(z.ckObj,C.int(value))
}

// property: ReqExecCharset
// Indicates the charset to be used for the command sent via the SendReqExec
// method. The default is "ANSI". A likely alternate value would be "utf-8".
func (z *Ssh) ReqExecCharset() string {
    return C.GoString(C.CkSsh_reqExecCharset(z.ckObj))
}
// property setter: ReqExecCharset
func (z *Ssh) SetReqExecCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putReqExecCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ServerIdentifier
// The server-identifier string received from the server during connection
// establishment. For example, a typical value would be similar to
// "SSH-2.0-OpenSSH_7.2p2 Ubuntu-4ubuntu2.2".
func (z *Ssh) ServerIdentifier() string {
    return C.GoString(C.CkSsh_serverIdentifier(z.ckObj))
}

// property: SessionLog
// Contains a log of the messages sent to/from the SSH server. To enable session
// logging, set the KeepSessionLog property = true. Note: This property is not a
// filename -- it is a string property that contains the session log data.
func (z *Ssh) SessionLog() string {
    return C.GoString(C.CkSsh_sessionLog(z.ckObj))
}

// property: SocksHostname
// The SOCKS4/SOCKS5 hostname or IPv4 address (in dotted decimal notation). This
// property is only used if the SocksVersion property is set to 4 or 5).
func (z *Ssh) SocksHostname() string {
    return C.GoString(C.CkSsh_socksHostname(z.ckObj))
}
// property setter: SocksHostname
func (z *Ssh) SetSocksHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putSocksHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPassword
// The SOCKS5 password (if required). The SOCKS4 protocol does not include the use
// of a password, so this does not apply to SOCKS4.
func (z *Ssh) SocksPassword() string {
    return C.GoString(C.CkSsh_socksPassword(z.ckObj))
}
// property setter: SocksPassword
func (z *Ssh) SetSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPort
// The SOCKS4/SOCKS5 proxy port. The default value is 1080. This property only
// applies if a SOCKS proxy is used (if the SocksVersion property is set to 4 or
// 5).
func (z *Ssh) SocksPort() int {
    return int(C.CkSsh_getSocksPort(z.ckObj))
}
// property setter: SocksPort
func (z *Ssh) SetSocksPort(value int) {
    C.CkSsh_putSocksPort(z.ckObj,C.int(value))
}

// property: SocksUsername
// The SOCKS4/SOCKS5 proxy username. This property is only used if the SocksVersion
// property is set to 4 or 5).
func (z *Ssh) SocksUsername() string {
    return C.GoString(C.CkSsh_socksUsername(z.ckObj))
}
// property setter: SocksUsername
func (z *Ssh) SetSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksVersion
// SocksVersion May be set to one of the following integer values:
// 
// 0 - No SOCKS proxy is used. This is the default.
// 4 - Connect via a SOCKS4 proxy.
// 5 - Connect via a SOCKS5 proxy.
//
func (z *Ssh) SocksVersion() int {
    return int(C.CkSsh_getSocksVersion(z.ckObj))
}
// property setter: SocksVersion
func (z *Ssh) SetSocksVersion(value int) {
    C.CkSsh_putSocksVersion(z.ckObj,C.int(value))
}

// property: SoRcvBuf
// Sets the receive buffer size socket option. Normally, this property should be
// left unchanged. The default value is 4194304.
// 
// This property can be increased if download performance seems slow. It is
// recommended to be a multiple of 4096.
//
func (z *Ssh) SoRcvBuf() int {
    return int(C.CkSsh_getSoRcvBuf(z.ckObj))
}
// property setter: SoRcvBuf
func (z *Ssh) SetSoRcvBuf(value int) {
    C.CkSsh_putSoRcvBuf(z.ckObj,C.int(value))
}

// property: SoSndBuf
// Sets the send buffer size socket option. Normally, this property should be left
// unchanged. The default value is 262144.
// 
// This property can be increased if upload performance seems slow. It is
// recommended to be a multiple of 4096. Testing with sizes such as 512K and 1MB is
// reasonable.
//
func (z *Ssh) SoSndBuf() int {
    return int(C.CkSsh_getSoSndBuf(z.ckObj))
}
// property setter: SoSndBuf
func (z *Ssh) SetSoSndBuf(value int) {
    C.CkSsh_putSoSndBuf(z.ckObj,C.int(value))
}

// property: StderrToStdout
// If true, then stderr is redirected to stdout. In this case, channel output for
// both stdout and stderr is combined and retrievable via the following methods:
// GetReceivedData, GetReceivedDataN, GetReceivedText, GetReceivedTextS. If this
// property is false, then stderr is available separately via the
// GetReceivedStderr method.
// 
// The default value of this property is true.
// 
// Note: Most SSH servers do not send stderr output as "extended data" packets as
// specified in RFC 4254. The SessionLog may be examined to see if any
// CHANNEL_EXTENDED_DATA messages exist. If not, then all of the output (stdout +
// stderr) was sent via CHANNEL_DATA messages, and therefore it is not possible to
// differentiate stderr output from stdout. In summary: This feature will not work
// for most SSH servers.
//
func (z *Ssh) StderrToStdout() bool {
    v := int(C.CkSsh_getStderrToStdout(z.ckObj))
    return v != 0
}
// property setter: StderrToStdout
func (z *Ssh) SetStderrToStdout(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putStderrToStdout(z.ckObj,v)
}

// property: StripColorCodes
// If true, then terminal color codes are stripped from the received text. The
// default value of this property is true. (Color codes are non-printable escape
// sequences that provide information about color for text in a terminal.)
func (z *Ssh) StripColorCodes() bool {
    v := int(C.CkSsh_getStripColorCodes(z.ckObj))
    return v != 0
}
// property setter: StripColorCodes
func (z *Ssh) SetStripColorCodes(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putStripColorCodes(z.ckObj,v)
}

// property: TcpNoDelay
// Controls whether the TCP_NODELAY socket option is used for the underlying TCP/IP
// socket. The default value is true. This disables the Nagle algorithm and
// allows for better performance when small amounts of data are sent to/from the
// SSH server.
func (z *Ssh) TcpNoDelay() bool {
    v := int(C.CkSsh_getTcpNoDelay(z.ckObj))
    return v != 0
}
// property setter: TcpNoDelay
func (z *Ssh) SetTcpNoDelay(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putTcpNoDelay(z.ckObj,v)
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. This property
// defaults to the empty string, and should typically remain empty.
// 
// Can be set to a list of the following comma separated keywords:
//     "KEX_DH_GEX_REQUEST_OLD" - Introduced in v9.5.0.73. Force the old Group
//     Exchange message to be used. This would be used for very old SSH server
//     implementations that do not use the RFC standard for
//     diffie-hellman-group-exchange.
//     "ProtectFromVpn" - Introduced in v9.5.0.80. On Android systems, will bypass
//     any VPN that may be installed or active.
//
func (z *Ssh) UncommonOptions() string {
    return C.GoString(C.CkSsh_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Ssh) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UserAuthBanner
// If a user authentication banner message is received during authentication, it
// will be made available here. An application can check to see if this property
// contains a banner string after calling StartKeyboardAuth. It is only possible
// for an application to display this message if it is doing keyboard-interactive
// authentication via the StartKeyboardAuth and ContinueKeyboardAuth methods.
func (z *Ssh) UserAuthBanner() string {
    return C.GoString(C.CkSsh_userAuthBanner(z.ckObj))
}
// property setter: UserAuthBanner
func (z *Ssh) SetUserAuthBanner(goStr string) {
    cStr := C.CString(goStr)
    C.CkSsh_putUserAuthBanner(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Ssh) VerboseLogging() bool {
    v := int(C.CkSsh_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Ssh) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSsh_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Ssh) Version() string {
    return C.GoString(C.CkSsh_version(z.ckObj))
}
// Authenticates with the SSH server using public-key authentication. The
// corresponding public key must have been installed on the SSH server for the
// username. Authentication will succeed if the matching privateKey is provided.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *Ssh) AuthenticatePk(arg1 string, arg2 *SshKey) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSsh_AuthenticatePk(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the AuthenticatePk method
func (z *Ssh) AuthenticatePkAsync(arg1 string, arg2 *SshKey, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSsh_AuthenticatePkAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Authenticates with the SSH server using a login and password.
// 
// An SSH session always begins by first calling Connect to connect to the SSH
// server, and then calling either AuthenticatePw or AuthenticatePk to login.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
// Note: To learn about how to handle password change requests, see the
// PasswordChangeRequested property (above).
//
func (z *Ssh) AuthenticatePw(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSsh_AuthenticatePw(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AuthenticatePw method
func (z *Ssh) AuthenticatePwAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSsh_AuthenticatePwAsync(z.ckObj, cstr1, cstr2)

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


// Authentication for SSH servers that require both a password and private key.
// (Most SSH servers are configured to require one or the other, but not both.)
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *Ssh) AuthenticatePwPk(arg1 string, arg2 string, arg3 *SshKey) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSsh_AuthenticatePwPk(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AuthenticatePwPk method
func (z *Ssh) AuthenticatePwPkAsync(arg1 string, arg2 string, arg3 *SshKey, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSsh_AuthenticatePwPkAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// The same as AuthenticatePw, except the login and passwords strings are passed in
// secure string objects.
func (z *Ssh) AuthenticateSecPw(arg1 *SecureString, arg2 *SecureString) bool {

    v := C.CkSsh_AuthenticateSecPw(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Asynchronous version of the AuthenticateSecPw method
func (z *Ssh) AuthenticateSecPwAsync(arg1 *SecureString, arg2 *SecureString, c chan *Task) {

    hTask := C.CkSsh_AuthenticateSecPwAsync(z.ckObj, arg1.ckObj, arg2.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as AuthenticatePwPk, except the login and passwords strings are passed
// in secure string objects.
func (z *Ssh) AuthenticateSecPwPk(arg1 *SecureString, arg2 *SecureString, arg3 *SshKey) bool {

    v := C.CkSsh_AuthenticateSecPwPk(z.ckObj, arg1.ckObj, arg2.ckObj, arg3.ckObj)


    return v != 0
}


// Asynchronous version of the AuthenticateSecPwPk method
func (z *Ssh) AuthenticateSecPwPkAsync(arg1 *SecureString, arg2 *SecureString, arg3 *SshKey, c chan *Task) {

    hTask := C.CkSsh_AuthenticateSecPwPkAsync(z.ckObj, arg1.ckObj, arg2.ckObj, arg3.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns true if the channel indicated by channelNum is open. Otherwise returns
// false.
func (z *Ssh) ChannelIsOpen(arg1 int) bool {

    v := C.CkSsh_ChannelIsOpen(z.ckObj, C.int(arg1))


    return v != 0
}


// Polls for incoming data on an open channel. This method will read a channel,
// waiting at most pollTimeoutMs milliseconds for data to arrive. Return values are as
// follows:
// 
// -1 -- Error. Check the IsConnected property to see if the connection to the SSH
// server is still valid. Also, call ChannelIsOpen to see if the channel remains
// open. The LastErrorText property will contain more detailed information
// regarding the error.
// 
// -2 -- No additional data was received prior to the poll timeout.
// 
// >0 -- Additional data was received and the return value indicates how many bytes
// are available to be "picked up". Methods that read data on a channel do not
// return the received data directly. Instead, they return an integer to indicate
// how many bytes are available to be "picked up". An application picks up the
// available data by calling GetReceivedData or GetReceivedText.
// 
// =0 -- A zero can be returned if the channel EOF has already been received, or if
// the channel had already been closed.
//
func (z *Ssh) ChannelPoll(arg1 int, arg2 int) int {

    v := C.CkSsh_ChannelPoll(z.ckObj, C.int(arg1), C.int(arg2))


    return int(v)
}


// Asynchronous version of the ChannelPoll method
func (z *Ssh) ChannelPollAsync(arg1 int, arg2 int, c chan *Task) {

    hTask := C.CkSsh_ChannelPollAsync(z.ckObj, C.int(arg1), C.int(arg2))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads incoming data on an open channel. This method will read a channel, waiting
// at most IdleTimeoutMs milliseconds for data to arrive. Return values are as
// follows:
// 
// -1 -- Error. Check the IsConnected property to see if the connection to the SSH
// server is still valid. Also, call ChannelIsOpen to see if the channel remains
// open. The LastErrorText property will contain more detailed information
// regarding the error.
// 
// -2 -- No additional data was received prior to the IdleTimeoutMs timeout.
// 
// >0 -- Additional data was received and the return value indicates how many bytes
// are available to be "picked up". Methods that read data on a channel do not
// return the received data directly. Instead, they return an integer to indicate
// how many bytes are available to be "picked up". An application picks up the
// available data by calling GetReceivedData or GetReceivedText.
// 
// =0 -- A zero can be returned if the channel EOF has already been received, or if
// the channel had already been closed.
//
func (z *Ssh) ChannelRead(arg1 int) int {

    v := C.CkSsh_ChannelRead(z.ckObj, C.int(arg1))


    return int(v)
}


// Asynchronous version of the ChannelRead method
func (z *Ssh) ChannelReadAsync(arg1 int, c chan *Task) {

    hTask := C.CkSsh_ChannelReadAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads incoming data on an open channel and continues reading until no data
// arrives for pollTimeoutMs milliseconds. The first read will wait a max of IdleTimeoutMs
// milliseconds before timing out. Subsequent reads wait a max of pollTimeoutMs milliseconds
// before timing out.
// 
// The idea behind ChannelReadAndPoll is to capture the output of a shell command.
// One might imagine the typical sequence of events when executing a shell command
// to be like this: (1) client sends command to server, (2) server executes the
// command (i.e. it's computing...), potentially taking some amount of time, (3)
// output is streamed back to the client. It makes sense for the client to wait a
// longer period of time for the first data to arrive, but once it begins arriving,
// the timeout can be shortened. This is exactly what ChannelReadAndPoll does --
// the first timeout is controlled by the IdleTimeoutMs property, while the
// subsequent reads (once output starts flowing) is controlled by pollTimeoutMs.
// 
// Return values are as follows:
// -1 -- Error. Check the IsConnected property to see if the connection to the SSH
// server is still valid. Also, call ChannelIsOpen to see if the channel remains
// open. The LastErrorText property will contain more detailed information
// regarding the error.
// 
// -2 -- No additional data was received prior to the IdleTimeoutMs timeout.
// 
// >0 -- Additional data was received and the return value indicates how many bytes
// are available to be "picked up". Methods that read data on a channel do not
// return the received data directly. Instead, they return an integer to indicate
// how many bytes are available to be "picked up". An application picks up the
// available data by calling GetReceivedData or GetReceivedText.
// 
// =0 -- A zero can be returned if the channel EOF has already been received, or if
// the channel had already been closed.
//
func (z *Ssh) ChannelReadAndPoll(arg1 int, arg2 int) int {

    v := C.CkSsh_ChannelReadAndPoll(z.ckObj, C.int(arg1), C.int(arg2))


    return int(v)
}


// Asynchronous version of the ChannelReadAndPoll method
func (z *Ssh) ChannelReadAndPollAsync(arg1 int, arg2 int, c chan *Task) {

    hTask := C.CkSsh_ChannelReadAndPollAsync(z.ckObj, C.int(arg1), C.int(arg2))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as ChannelReadAndPoll, except this method will return as soon as maxNumBytes
// is exceeded, which may be as large as the MaxPacketSize property setting.
func (z *Ssh) ChannelReadAndPoll2(arg1 int, arg2 int, arg3 int) int {

    v := C.CkSsh_ChannelReadAndPoll2(z.ckObj, C.int(arg1), C.int(arg2), C.int(arg3))


    return int(v)
}


// Asynchronous version of the ChannelReadAndPoll2 method
func (z *Ssh) ChannelReadAndPoll2Async(arg1 int, arg2 int, arg3 int, c chan *Task) {

    hTask := C.CkSsh_ChannelReadAndPoll2Async(z.ckObj, C.int(arg1), C.int(arg2), C.int(arg3))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// true if a CLOSE message has been received on the channel indicated by channelNum.
// When a CLOSE is received, no subsequent data should be sent in either direction
// -- the channel is closed in both directions.
func (z *Ssh) ChannelReceivedClose(arg1 int) bool {

    v := C.CkSsh_ChannelReceivedClose(z.ckObj, C.int(arg1))


    return v != 0
}


// true if an EOF message has been received on the channel indicated by channelNum.
// When an EOF is received, no more data will be forthcoming on the channel.
// However, data may still be sent in the opposite direction.
func (z *Ssh) ChannelReceivedEof(arg1 int) bool {

    v := C.CkSsh_ChannelReceivedEof(z.ckObj, C.int(arg1))


    return v != 0
}


// true if an exit status code was received on the channel. Otherwise false.
func (z *Ssh) ChannelReceivedExitStatus(arg1 int) bool {

    v := C.CkSsh_ChannelReceivedExitStatus(z.ckObj, C.int(arg1))


    return v != 0
}


// Reads incoming data on an open channel until the channel is closed by the
// server. If successful, the number of bytes available to be "picked up" can be
// determined by calling GetReceivedNumBytes. The received data may be retrieved by
// calling GetReceivedData or GetReceivedText.
func (z *Ssh) ChannelReceiveToClose(arg1 int) bool {

    v := C.CkSsh_ChannelReceiveToClose(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the ChannelReceiveToClose method
func (z *Ssh) ChannelReceiveToCloseAsync(arg1 int, c chan *Task) {

    hTask := C.CkSsh_ChannelReceiveToCloseAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads incoming text data on an open channel until the received data matches the
// matchPattern. For example, to receive data until the string "Hello World" arrives, set
// matchPattern equal to "*Hello World*". charset indicates the character encoding of the text
// being received ("iso-8859-1" for example). caseSensitive may be set to true for case
// sensitive matching, or false for case insensitive matching.
// 
// Returns true if text data matching matchPattern was received and is available to be
// picked up by calling GetReceivedText (or GetReceivedTextS). IMPORTANT: This
// method may read beyond the matching text. Call GetReceivedTextS to extract only
// the data up-to and including the matching text.
// 
// Important Notes:
//     It's wise to set the ReadTimeoutMs property to a non-zero value to prevent
//     an infinite wait if the matchPattern never arrives.
//     If using a shell session and SendReqPty was called, set the termType =
//     "dumb". If terminal control codes get mixed into the output stream, it could
//     disrupt matching.
//     Be aware of the StderrToStdout property setting. The default value is true,
//     which means that stderr is mixed with stdout in the output stream. This could
//     disrupt matching. Set StderrToStdout to false to prevent this possibility.
//
func (z *Ssh) ChannelReceiveUntilMatch(arg1 int, arg2 string, arg3 string, arg4 bool) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkSsh_ChannelReceiveUntilMatch(z.ckObj, C.int(arg1), cstr2, cstr3, b4)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the ChannelReceiveUntilMatch method
func (z *Ssh) ChannelReceiveUntilMatchAsync(arg1 int, arg2 string, arg3 string, arg4 bool, c chan *Task) {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkSsh_ChannelReceiveUntilMatchAsync(z.ckObj, C.int(arg1), cstr2, cstr3, b4)

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


// The same as ChannelReceiveUntilMatch except that the method returns when any one
// of the match patterns specified in matchPatterns are received on the channel.
// 
// Important: It's wise to set the ReadTimeoutMs property to a non-zero value to
// prevent an infinite wait if of the matchPatterns ever arrives.
//
func (z *Ssh) ChannelReceiveUntilMatchN(arg1 int, arg2 *StringArray, arg3 string, arg4 bool) bool {
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkSsh_ChannelReceiveUntilMatchN(z.ckObj, C.int(arg1), arg2.ckObj, cstr3, b4)

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the ChannelReceiveUntilMatchN method
func (z *Ssh) ChannelReceiveUntilMatchNAsync(arg1 int, arg2 *StringArray, arg3 string, arg4 bool, c chan *Task) {
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkSsh_ChannelReceiveUntilMatchNAsync(z.ckObj, C.int(arg1), arg2.ckObj, cstr3, b4)

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


// Releases the internal memory resources for a channel previously opened by
// OpenSessionChannel, OpenCustomChannel, or OpenDirectTcpIpChannel. It is not
// absolutely necessary to call this method because the internal memory resources
// for all channels are automatically released when the SSH object instance is
// deleted/disposed. This method becomes necessary only when an extremely large
// number of channels within the same SSH object instance are opened, used, and
// closed over a long period of time.
func (z *Ssh) ChannelRelease(arg1 int)  {

    C.CkSsh_ChannelRelease(z.ckObj, C.int(arg1))



}


// Sends a CLOSE message to the server for the channel indicated by channelNum. This
// closes both directions of the bidirectional channel.
func (z *Ssh) ChannelSendClose(arg1 int) bool {

    v := C.CkSsh_ChannelSendClose(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the ChannelSendClose method
func (z *Ssh) ChannelSendCloseAsync(arg1 int, c chan *Task) {

    hTask := C.CkSsh_ChannelSendCloseAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends byte data on the channel indicated by channelNum.
func (z *Ssh) ChannelSendData(arg1 int, arg2 []byte) bool {
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkSsh_ChannelSendData(z.ckObj, C.int(arg1), ckbyd2)

    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Asynchronous version of the ChannelSendData method
func (z *Ssh) ChannelSendDataAsync(arg1 int, arg2 []byte, c chan *Task) {
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    hTask := C.CkSsh_ChannelSendDataAsync(z.ckObj, C.int(arg1), ckbyd2)

    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an EOF for the channel indicated by channelNum. Once an EOF is sent, no
// additional data may be sent on the channel. However, the channel remains open
// and additional data may still be received from the server.
func (z *Ssh) ChannelSendEof(arg1 int) bool {

    v := C.CkSsh_ChannelSendEof(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the ChannelSendEof method
func (z *Ssh) ChannelSendEofAsync(arg1 int, c chan *Task) {

    hTask := C.CkSsh_ChannelSendEofAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends character data on the channel indicated by channelNum. The text is converted to
// the charset indicated by charset prior to being sent. A list of supported charset
// values may be found on this page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
func (z *Ssh) ChannelSendString(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkSsh_ChannelSendString(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the ChannelSendString method
func (z *Ssh) ChannelSendStringAsync(arg1 int, arg2 string, arg3 string, c chan *Task) {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkSsh_ChannelSendStringAsync(z.ckObj, C.int(arg1), cstr2, cstr3)

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


// Returns true if the underlying TCP socket is connected to the SSH server.
func (z *Ssh) CheckConnection() bool {

    v := C.CkSsh_CheckConnection(z.ckObj)


    return v != 0
}


// Clears the collection of TTY modes that are sent with the SendReqPty method.
func (z *Ssh) ClearTtyModes()  {

    C.CkSsh_ClearTtyModes(z.ckObj)



}


// Connects to the SSH server at domainName:port
// 
// The domainName may be a domain name or an IPv4 or IPv6 address in string format.
// 
// Internally, the following SSH connection protocol algorithms are supported:
//     Hostkey Types: ssh-rsa, ssh-dsa, ecdsa-sha2-nistp256, rsa-sha2-256,
//     rsa-sha2-512, ssh-ed25519
//     Key Exchange Methods: curve25519-sha256@libssh.org, ecdh-sha2-nistp256,
//     ecdh-sha2-nistp384, ecdh-sha2-nistp521, diffie-hellman-group-exchange-sha256,
//     diffie-hellman-group-exchange-sha1, diffie-hellman-group14-sha1,
//     diffie-hellman-group1-sha1
//     Ciphers: chacha20-poly1305@openssh.com, aes256-ctr, aes192-ctr, aes128-ctr,
//     aes256-cbc, aes192-cbc, aes128-cbc, twofish256-cbc, twofish128-cbc,
//     blowfish-cbc, 3des-cbc, arcfour128, arcfour256
//     MAC Algorithms: hmac-sha2-256, hmac-sha2-512, hmac-sha1, hmac-md5,
//     hmac-ripemd160, hmac-sha1-96
//     Compression: none, zlib, zlib@openssh.com
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
//
func (z *Ssh) Connect(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSsh_Connect(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Connect method
func (z *Ssh) ConnectAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSsh_ConnectAsync(z.ckObj, cstr1, C.int(arg2))

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


// Connects to an SSH server through an existing SSH connection. The ssh is an
// existing connected and authenticated SSH object. The connection to hostname:port is
// made through the existing SSH connection via port-forwarding. If successful, the
// connection is as follows: application => ServerSSH1 => ServerSSH2. (where
// ServerSSH1 is the ssh and ServerSSH2 is the SSH server at hostname:port) Once
// connected in this way, all communications are routed through ServerSSH1 to
// ServerSSH2. This includes authentication -- which means the application must
// still call one of the Authenticate* methods to authenticate with ServerSSH2.
func (z *Ssh) ConnectThroughSsh(arg1 *Ssh, arg2 string, arg3 int) bool {
    cstr2 := C.CString(arg2)

    v := C.CkSsh_ConnectThroughSsh(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the ConnectThroughSsh method
func (z *Ssh) ConnectThroughSshAsync(arg1 *Ssh, arg2 string, arg3 int, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkSsh_ConnectThroughSshAsync(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

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


// Continues keyboard-interactive authentication with the SSH server. The response is
// typically the password. If multiple responses are required (because there were
// multiple prompts in the infoRequest XML returned by StartKeyboardAuth), then the
// response should be formatted as XML (as shown below) otherwise the response simply
// contains the single response string.
// _LT_response_GT_
//     _LT_response1_GT_response to first prompt_LT_/response1_GT_
//     _LT_response2_GT_response to second prompt_LT_/response2_GT_
//     ...
//     _LT_responseN_GT_response to Nth prompt_LT_/responseN_GT_
// _LT_/response_GT_
// 
// If the interactive authentication completed with success or failure, the XML
// response will be:
// _LT_success_GT_success_message_LT_/success_GT_
// 
// or
// 
// _LT_error_GT_error_message_LT_/error_GT_
// If additional steps are required to complete the interactive authentication,
// then an XML string that provides the name, instruction, and prompts is returned.
// The XML has the following format:
//  	_LT_infoRequest numPrompts="N"_GT_
// 	    _LT_name_GT_name_string_LT_/name_GT_
// 	    _LT_instruction_GT_instruction_string_LT_/instruction_GT_
// 	    _LT_prompt1 echo="1_or_0"_GT_prompt_string_LT_/prompt1_GT_
// 	    ...
// 	    _LT_promptN echo="1_or_0"_GT_prompt_string_LT_/promptN_GT_
// 	_LT_/infoRequest_GT_
//
// return a string or nil if failed.
func (z *Ssh) ContinueKeyboardAuth(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSsh_continueKeyboardAuth(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ContinueKeyboardAuth method
func (z *Ssh) ContinueKeyboardAuthAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSsh_ContinueKeyboardAuthAsync(z.ckObj, cstr1)

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


// Disconnects from the SSH server.
func (z *Ssh) Disconnect()  {

    C.CkSsh_Disconnect(z.ckObj)



}


// Queries the SSH server to find out which authentication methods it supports.
// Returns a string such as "publickey,password,keyboard-interactive".
// 
// This method should be called after connecting, but prior to authenticating. The
// method intentionally disconnects from the server after getting the
// authentication methods. An application may then connect again and authentication
// with a chosen method. (In most cases, an application knows in advance the type
// of authentication to be used, and this method is never called.)
//
// return a string or nil if failed.
func (z *Ssh) GetAuthMethods() *string {

    retStr := C.CkSsh_getAuthMethods(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetAuthMethods method
func (z *Ssh) GetAuthMethodsAsync(c chan *Task) {

    hTask := C.CkSsh_GetAuthMethodsAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the exit status code for a channel. This method should only be called if
// an exit status has been received. You may check to see if the exit status was
// received by calling ChannelReceivedExitStatus.
func (z *Ssh) GetChannelExitStatus(arg1 int) int {

    v := C.CkSsh_GetChannelExitStatus(z.ckObj, C.int(arg1))


    return int(v)
}


// Returns the channel number for the Nth open channel. Indexing begins at 0, and
// the number of currently open channels is indicated by the NumOpenChannels
// property. Returns -1 if the index is out of range.
func (z *Ssh) GetChannelNumber(arg1 int) int {

    v := C.CkSsh_GetChannelNumber(z.ckObj, C.int(arg1))


    return int(v)
}


// Returns a string describing the channel type for the Nth open channel. Channel
// types are: "session", "x11", "forwarded-tcpip", and "direct-tcpip".
// return a string or nil if failed.
func (z *Ssh) GetChannelType(arg1 int) *string {

    retStr := C.CkSsh_getChannelType(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the accumulated data received on the channel indicated by channelNum and
// clears the channel's internal receive buffer.
func (z *Ssh) GetReceivedData(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSsh_GetReceivedData(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Same as GetReceivedData, but a maximum of maxNumBytes bytes is returned.
func (z *Ssh) GetReceivedDataN(arg1 int, arg2 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSsh_GetReceivedDataN(z.ckObj, C.int(arg1), C.int(arg2), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the number of bytes available in the internal receive buffer for the
// specified channelNum. The received data may be retrieved by calling GetReceivedData or
// GetReceivedText.
func (z *Ssh) GetReceivedNumBytes(arg1 int) int {

    v := C.CkSsh_GetReceivedNumBytes(z.ckObj, C.int(arg1))


    return int(v)
}


// Returns the accumulated stderr bytes received on the channel indicated by channelNum
// and clears the channel's internal stderr receive buffer.
// 
// Note: If the StderrToStdout property is set to true, then stderr is
// automatically redirected to stdout. This is the default behavior. The following
// methods can be called to retrieve the channel's stdout: GetReceivedData,
// GetReceivedDataN, GetReceivedText, and GetReceivedTextS.
//
func (z *Ssh) GetReceivedStderr(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSsh_GetReceivedStderr(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the accumulated stderr text received on the channel indicated by channelNum
// and clears the channel's internal receive buffer. The charset indicates the charset
// of the character data in the internal receive buffer. A list of supported
// charset values may be found on this page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
// return a string or nil if failed.
func (z *Ssh) GetReceivedStderrText(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkSsh_getReceivedStderrText(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the accumulated text received on the channel indicated by channelNum and
// clears the channel's internal receive buffer. The charset indicates the charset of
// the character data in the internal receive buffer. A list of supported charset
// values may be found on this page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
// return a string or nil if failed.
func (z *Ssh) GetReceivedText(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkSsh_getReceivedText(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The same as GetReceivedText, except only the text up to and including substr is
// returned. The text returned is removed from the internal receive buffer. If the
// substr was not found in the internal receive buffer, an empty string is returned
// and the internal receive buffer is not modified.
// return a string or nil if failed.
func (z *Ssh) GetReceivedTextS(arg1 int, arg2 string, arg3 string) *string {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkSsh_getReceivedTextS(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Provides information about what transpired in the last method called. For many
// methods, there is no information. For some methods, details about what
// transpired can be obtained via LastJsonData.
func (z *Ssh) LastJsonData() *JsonObject {

    retObj := C.CkSsh_LastJsonData(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Loads the caller of the task's async method.
func (z *Ssh) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkSsh_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Opens a custom channel with a custom server that uses the SSH protocol. The channelType
// is application-defined.
// 
// If successful, the channel number is returned. This is the number that should be
// passed to any method requiring a channel number. A -1 is returned upon failure.
//
func (z *Ssh) OpenCustomChannel(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkSsh_OpenCustomChannel(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the OpenCustomChannel method
func (z *Ssh) OpenCustomChannelAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSsh_OpenCustomChannelAsync(z.ckObj, cstr1)

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


// Open a direct-tcpip channel for port forwarding. Data sent on the channel via
// ChannelSend* methods is sent to the SSH server and then forwarded to targetHostname:targetPort.
// The SSH server automatically forwards data received from targetHostname:targetPort to the SSH
// client. Therefore, calling ChannelRead* and ChannelReceive* methods is
// equivalent to reading directly from targetHostname:targetPort.
// 
// If successful, the channel number is returned. This is the number that should be
// passed to any method requiring a channel number. A -1 is returned upon failure.
//
func (z *Ssh) OpenDirectTcpIpChannel(arg1 string, arg2 int) int {
    cstr1 := C.CString(arg1)

    v := C.CkSsh_OpenDirectTcpIpChannel(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the OpenDirectTcpIpChannel method
func (z *Ssh) OpenDirectTcpIpChannelAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSsh_OpenDirectTcpIpChannelAsync(z.ckObj, cstr1, C.int(arg2))

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


// Opens a new session channel. Almost everything you will do with the Chilkat SSH
// component will involve opening a session channel. The normal sequence of
// operation is typically this: 1) Connect to the SSH server. 2) Authenticate. 3)
// Open a session channel. 4) do something on the channel such as opening a shell,
// execute a command, etc.
// 
// If successful, the channel number is returned. This is the number that should be
// passed to any method requiring a channel number. A -1 is returned upon failure.
//
func (z *Ssh) OpenSessionChannel() int {

    v := C.CkSsh_OpenSessionChannel(z.ckObj)


    return int(v)
}


// Asynchronous version of the OpenSessionChannel method
func (z *Ssh) OpenSessionChannelAsync(c chan *Task) {

    hTask := C.CkSsh_OpenSessionChannelAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// This is the same as GetReceivedText, except the internal receive buffer is not
// cleared.
// return a string or nil if failed.
func (z *Ssh) PeekReceivedText(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkSsh_peekReceivedText(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a channel number for a completed command that was previously sent via
// QuickCmdSend. Returns -1 if no commands have yet completed. The pollTimeoutMs indicates
// how long to wait (in milliseconds) for any command in progress (on any channel)
// to complete before returning -1.
// 
// Returns -2 if an error occurred (for example, if the connection to the SSH
// server was lost while checking for responses).
//
func (z *Ssh) QuickCmdCheck(arg1 int) int {

    v := C.CkSsh_QuickCmdCheck(z.ckObj, C.int(arg1))


    return int(v)
}


// Asynchronous version of the QuickCmdCheck method
func (z *Ssh) QuickCmdCheckAsync(arg1 int, c chan *Task) {

    hTask := C.CkSsh_QuickCmdCheckAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a command and returns the channel number for the command that has started.
// This is the equivalent of calling OpenSessionChannel, followed by SendReqExec. A
// value of -1 is returned on failure.
// 
// The ReqExecCharset property controls the charset used for the command that is
// sent.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *Ssh) QuickCmdSend(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkSsh_QuickCmdSend(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the QuickCmdSend method
func (z *Ssh) QuickCmdSendAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSsh_QuickCmdSendAsync(z.ckObj, cstr1)

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


// Simplified method for executing a remote command and getting the complete
// output. This is the equivalent of calling OpenSessionChannel, followed by
// SendReqExec, then ChannelReceiveToClose, and finally GetReceivedText.
// 
// The charset indicates the charset of the command's output (such as "utf-8" or
// "ansi"). A list of supported charset values may be found on this page:Supported
// Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
// 
// The ReqExecCharset property controls the charset used for the command that is
// sent.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
// return a string or nil if failed.
func (z *Ssh) QuickCommand(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkSsh_quickCommand(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the QuickCommand method
func (z *Ssh) QuickCommandAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSsh_QuickCommandAsync(z.ckObj, cstr1, cstr2)

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


// Simplified method for starting a remote shell session. It is the equivalent of
// calling OpenSessionChannel, followed by SendReqPty, and finally SendReqShell.
// 
// Returns the SSH channel number for the session, or -1 if not successful.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *Ssh) QuickShell() int {

    v := C.CkSsh_QuickShell(z.ckObj)


    return int(v)
}


// Asynchronous version of the QuickShell method
func (z *Ssh) QuickShellAsync(c chan *Task) {

    hTask := C.CkSsh_QuickShellAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Initiates a re-key with the SSH server. The ReKey method does not return until
// the key re-exchange is complete.
// 
// RFC 4253 (the SSH Transport Layer Protocol) recommends that keys be changed
// after each gigabyte of transmitted data or after each hour of connection time,
// whichever comes sooner. Key re-exchange is a public-key operation and requires a
// fair amount of processing power and should not be performed too often. Either
// side (client or server) may initiate a key re-exchange at any time.
// 
// In most cases, a server will automatically initiate key re-exchange whenever it
// deems necessary, and the Chilkat SSH component handles these transparently. For
// example, if the Chilkat SSH component receives a re-key message from the server
// while in the process of receiving data on a channel, it will automatically
// handle the key re-exchange and the application will not even realize that an
// underlying key re-exchange occurred.
//
func (z *Ssh) ReKey() bool {

    v := C.CkSsh_ReKey(z.ckObj)


    return v != 0
}


// Asynchronous version of the ReKey method
func (z *Ssh) ReKeyAsync(c chan *Task) {

    hTask := C.CkSsh_ReKeyAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an IGNORE message to the SSH server. This is one way of verifying that the
// connection to the SSH server is open and valid. The SSH server does not respond
// to an IGNORE message. It simply ignores it. IGNORE messages are not associated
// with a channel (in other words, you do not need to first open a channel prior to
// sending an IGNORE message).
func (z *Ssh) SendIgnore() bool {

    v := C.CkSsh_SendIgnore(z.ckObj)


    return v != 0
}


// Asynchronous version of the SendIgnore method
func (z *Ssh) SendIgnoreAsync(c chan *Task) {

    hTask := C.CkSsh_SendIgnoreAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Initiates execution of a command on the channel specified by channelNum. The commandLine
// contains the full command line including any command-line parameters (just as
// you would type the command at a shell prompt).
// 
// The user's default shell (typically defined in /etc/password in UNIX systems) is
// started on the SSH server to execute the command.
// 
// Important: A channel only exists for a single request. You may not call
// SendReqExec multiple times on the same open channel. The reason is that the SSH
// server automatically closes the channel at the end of the exec. The solution is
// to call OpenSessionChannel to get a new channel, and then call SendReqExec using
// the new channel. It is OK to have more than one channel open simultaneously.
// 
// Charset: The ReqExecCharset property has been added in version 9.5.0.47. This
// can be set to control the character encoding of the command sent to the server.
// The default is ANSI. A likely alternative value is "utf-8".
//
func (z *Ssh) SendReqExec(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkSsh_SendReqExec(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqExec method
func (z *Ssh) SendReqExecAsync(arg1 int, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkSsh_SendReqExecAsync(z.ckObj, C.int(arg1), cstr2)

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


// Requests a pseudo-terminal for a session channel. If the termType is a character
// oriented terminal ("vt100" for example), then widthInChars and heightInChars would be set to
// non-zero values, while widthInPixels and heightInPixels may be set to 0. If termType is pixel-oriented,
// such as "xterm", the reverse is true (i.e. set widthInPixels and heightInPixels, but set widthInChars and
// heightInChars equal to 0).
// 
// In most cases, you probably don't even want terminal emulation. In that case,
// try setting termType = "dumb". Terminal emulation causes terminal escape sequences
// to be included with shell command output. A "dumb" terminal should have no
// escape sequences.
// 
// Some SSH servers allow a shell to be started (via the SendReqShell method)
// without the need to first request a pseudo-terminal. The normal sequence for
// starting a remote shell is as follows:
// 1) Connect
// 2) Authenticate
// 3) OpenSessionChannel
// 4) Request a PTY via this method if necessary.
// 5) Start a shell by calling SendReqShell
//
func (z *Ssh) SendReqPty(arg1 int, arg2 string, arg3 int, arg4 int, arg5 int, arg6 int) bool {
    cstr2 := C.CString(arg2)

    v := C.CkSsh_SendReqPty(z.ckObj, C.int(arg1), cstr2, C.int(arg3), C.int(arg4), C.int(arg5), C.int(arg6))

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqPty method
func (z *Ssh) SendReqPtyAsync(arg1 int, arg2 string, arg3 int, arg4 int, arg5 int, arg6 int, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkSsh_SendReqPtyAsync(z.ckObj, C.int(arg1), cstr2, C.int(arg3), C.int(arg4), C.int(arg5), C.int(arg6))

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


// Sets an environment variable in the remote shell.
func (z *Ssh) SendReqSetEnv(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkSsh_SendReqSetEnv(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the SendReqSetEnv method
func (z *Ssh) SendReqSetEnvAsync(arg1 int, arg2 string, arg3 string, c chan *Task) {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkSsh_SendReqSetEnvAsync(z.ckObj, C.int(arg1), cstr2, cstr3)

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


// Starts a shell on an open session channel. Some SSH servers require that a PTY
// (pseudo-terminal) first be requested prior to starting a shell. In that case,
// call SendReqPty prior to calling this method. Once a shell is started, commands
// may be sent by calling ChannelSendString. (Don't forget to terminate commands
// with a CRLF).
func (z *Ssh) SendReqShell(arg1 int) bool {

    v := C.CkSsh_SendReqShell(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the SendReqShell method
func (z *Ssh) SendReqShellAsync(arg1 int, c chan *Task) {

    hTask := C.CkSsh_SendReqShellAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Delivers a signal to the remote process/service. signalName is one of the following:
// ABRT, ALRM, FPE, HUP, ILL, INT, KILL, PIPE, QUIT, SEGV, TERM, USR1, USR2.
// (Obviously, these are UNIX signals, so the remote SSH server would need to be a
// Unix/Linux system.)
func (z *Ssh) SendReqSignal(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkSsh_SendReqSignal(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqSignal method
func (z *Ssh) SendReqSignalAsync(arg1 int, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkSsh_SendReqSignalAsync(z.ckObj, C.int(arg1), cstr2)

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


// Executes a pre-defined subsystem. The SFTP protocol (Secure File Transfer
// Protocol) is started by the Chilkat SFTP component by starting the "sftp"
// subsystem.
func (z *Ssh) SendReqSubsystem(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkSsh_SendReqSubsystem(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendReqSubsystem method
func (z *Ssh) SendReqSubsystemAsync(arg1 int, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkSsh_SendReqSubsystemAsync(z.ckObj, C.int(arg1), cstr2)

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


// When the client-side window (terminal) size changes, this message may be sent to
// the server to inform it of the new size.
func (z *Ssh) SendReqWindowChange(arg1 int, arg2 int, arg3 int, arg4 int, arg5 int) bool {

    v := C.CkSsh_SendReqWindowChange(z.ckObj, C.int(arg1), C.int(arg2), C.int(arg3), C.int(arg4), C.int(arg5))


    return v != 0
}


// Asynchronous version of the SendReqWindowChange method
func (z *Ssh) SendReqWindowChangeAsync(arg1 int, arg2 int, arg3 int, arg4 int, arg5 int, c chan *Task) {

    hTask := C.CkSsh_SendReqWindowChangeAsync(z.ckObj, C.int(arg1), C.int(arg2), C.int(arg3), C.int(arg4), C.int(arg5))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Allows the client to send an X11 forwarding request to the server. Chilkat only
// provides this functionality because it is a message defined in the SSH
// connection protocol. Chilkat has no advice for when or why it would be needed.
func (z *Ssh) SendReqX11Forwarding(arg1 int, arg2 bool, arg3 string, arg4 string, arg5 int) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkSsh_SendReqX11Forwarding(z.ckObj, C.int(arg1), b2, cstr3, cstr4, C.int(arg5))

    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Asynchronous version of the SendReqX11Forwarding method
func (z *Ssh) SendReqX11ForwardingAsync(arg1 int, arg2 bool, arg3 string, arg4 string, arg5 int, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    hTask := C.CkSsh_SendReqX11ForwardingAsync(z.ckObj, C.int(arg1), b2, cstr3, cstr4, C.int(arg5))

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


// This method should be ignored and not used.
func (z *Ssh) SendReqXonXoff(arg1 int, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkSsh_SendReqXonXoff(z.ckObj, C.int(arg1), b2)


    return v != 0
}


// Asynchronous version of the SendReqXonXoff method
func (z *Ssh) SendReqXonXoffAsync(arg1 int, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkSsh_SendReqXonXoffAsync(z.ckObj, C.int(arg1), b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sets a TTY mode that is included in the SendReqPty method call. Most commonly,
// it is not necessary to call this method at all. Chilkat has no recommendations
// or expertise as to why or when a particular mode might be useful. This
// capability is provided because it is defined in the SSH connection protocol
// specification.
// 
// This method can be called multiple times to set many terminal mode flags (one
// per call).
// 
// The ttyValue is an integer, typically 0 or 1. Valid ttyName flag names include: VINTR,
// VQUIT, VERASE, VKILL, VEOF, VEOL, VEOL2, VSTART, VSTOP, VSUSP, VDSUSP, VREPRINT,
// VWERASE, VLNEXT, VFLUSH, VSWTCH, VSTATUS, VDISCARD, IGNPAR, PARMRK, INPCK,
// ISTRIP, INLCR, IGNCR, ICRNL, IUCLC, IXON, IXANY, IXOFF, IMAXBEL, ISIG, ICANON,
// XCASE, ECHO, ECHOE, ECHOK, ECHONL, NOFLSH, TOSTOP, IEXTEN, ECHOCTL, ECHOKE,
// PENDIN, OPOST, OLCUC, ONLCR, OCRNL, ONOCR, ONLRET, CS7, CS8, PARENB, PARODD,
// TTY_OP_ISPEED, TTY_OP_OSPEED
//
func (z *Ssh) SetTtyMode(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSsh_SetTtyMode(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Begins keyboard-interactive authentication with the SSH server. Returns an XML
// string providing the name, instruction, and prompts. The XML has the following
// format:
//  	_LT_infoRequest numPrompts="N"_GT_
// 	    _LT_name_GT_name_string_LT_/name_GT_
// 	    _LT_instruction_GT_instruction_string_LT_/instruction_GT_
// 	    _LT_prompt1 echo="1_or_0"_GT_prompt_string_LT_/prompt1_GT_
// 	    ...
// 	    _LT_promptN echo="1_or_0"_GT_prompt_string_LT_/promptN_GT_
// 	_LT_/infoRequest_GT_
// 
// If the authentication immediately succeeds because no password is required, or
// immediately fails, the XML response can be:
// _LT_success_GT_success_message_LT_/success_GT_
// 
// or
// 
// _LT_error_GT_error_message_LT_/error_GT_
//
// return a string or nil if failed.
func (z *Ssh) StartKeyboardAuth(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSsh_startKeyboardAuth(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the StartKeyboardAuth method
func (z *Ssh) StartKeyboardAuthAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSsh_StartKeyboardAuthAsync(z.ckObj, cstr1)

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


// Unlocks the component. This must be called once prior to calling any other
// method. A fully-functional 30-day trial is automatically started when an
// arbitrary string is passed to this method. For example, passing "Hello", or
// "abc123" will unlock the component for the 1st thirty days after the initial
// install.
func (z *Ssh) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSsh_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// The pollTimeoutMs is the number of milliseconds to wait. To poll, pass a value of 0 in
// pollTimeoutMs. Waits for an incoming message on any channel. This includes data, EOF,
// CLOSE, etc. If a message arrives in the alotted time, the channel number is
// returned. A value of -1 is returned for a timeout, and -2 for any other errors
// such as if the connection is lost.
// 
// Note: If a channel number is returned, the message must still be read by calling
// a method such as ChannelRead, ChannelReceiveUntilMatch, etc. Once the message is
// actually received, it may be collected by calling GetReceivedText,
// GetReceivedData, etc.
//
func (z *Ssh) WaitForChannelMessage(arg1 int) int {

    v := C.CkSsh_WaitForChannelMessage(z.ckObj, C.int(arg1))


    return int(v)
}


// Asynchronous version of the WaitForChannelMessage method
func (z *Ssh) WaitForChannelMessageAsync(arg1 int, c chan *Task) {

    hTask := C.CkSsh_WaitForChannelMessageAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


