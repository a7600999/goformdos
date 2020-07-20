// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkSshTunnel.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewSshTunnel() *SshTunnel {
	return &SshTunnel{C.CkSshTunnel_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *SshTunnel) DisposeSshTunnel() {
    if z != nil {
	C.CkSshTunnel_Dispose(z.ckObj)
	}
}




func (z *SshTunnel) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkSshTunnel_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *SshTunnel) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkSshTunnel_setExternalProgress(z.ckObj,1)
}

func (z *SshTunnel) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkSshTunnel_setExternalProgress(z.ckObj,1)
}

func (z *SshTunnel) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkSshTunnel_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *SshTunnel) AbortCurrent() bool {
    v := int(C.CkSshTunnel_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *SshTunnel) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSshTunnel_putAbortCurrent(z.ckObj,v)
}

// property: AcceptLog
// Contains an in-memory log of the listen thread. This will only contain content
// if the KeepAcceptLog property is true.
func (z *SshTunnel) AcceptLog() string {
    return C.GoString(C.CkSshTunnel_acceptLog(z.ckObj))
}
// property setter: AcceptLog
func (z *SshTunnel) SetAcceptLog(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putAcceptLog(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: AcceptLogPath
// Specifies a log file to be kept for the activity in the listen thread.
func (z *SshTunnel) AcceptLogPath() string {
    return C.GoString(C.CkSshTunnel_acceptLogPath(z.ckObj))
}
// property setter: AcceptLogPath
func (z *SshTunnel) SetAcceptLogPath(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putAcceptLogPath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConnectTimeoutMs
// Maximum number of milliseconds to wait when connecting to an SSH server. The
// default value is 10000 (i.e. 10 seconds). A value of 0 indicates no timeout
// (wait forever).
func (z *SshTunnel) ConnectTimeoutMs() int {
    return int(C.CkSshTunnel_getConnectTimeoutMs(z.ckObj))
}
// property setter: ConnectTimeoutMs
func (z *SshTunnel) SetConnectTimeoutMs(value int) {
    C.CkSshTunnel_putConnectTimeoutMs(z.ckObj,C.int(value))
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
func (z *SshTunnel) DebugLogFilePath() string {
    return C.GoString(C.CkSshTunnel_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *SshTunnel) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DestHostname
// The destination hostname or IP address (in dotted decimal notation) of the
// service (such as a database server). Data sent through the SSH tunnel is
// forwarded by the SSH server to this destination. Data received from the
// destination (by the SSH server) is forwarded back to the client through the SSH
// tunnel.
func (z *SshTunnel) DestHostname() string {
    return C.GoString(C.CkSshTunnel_destHostname(z.ckObj))
}
// property setter: DestHostname
func (z *SshTunnel) SetDestHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putDestHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DestPort
// The destination port of the service (such as a database server).
func (z *SshTunnel) DestPort() int {
    return int(C.CkSshTunnel_getDestPort(z.ckObj))
}
// property setter: DestPort
func (z *SshTunnel) SetDestPort(value int) {
    C.CkSshTunnel_putDestPort(z.ckObj,C.int(value))
}

// property: DynamicPortForwarding
// If true, then this behaves as a SOCKS proxy server for inbound connections.
// When this property is true, the DestHostname and DestPort properties are
// unused because the destination IP:port is dynamically provided by the SOCKS
// client. The default value of this property is false.
// 
// When dynamic port forwarding is used, the InboundSocksVersion property must be
// set to 4 or 5. If inbound SOCKS5 is used, then the InboundSocksUsername and
// InboundSocksPassword may be set to the required login/password for a client to
// gain access.
//
func (z *SshTunnel) DynamicPortForwarding() bool {
    v := int(C.CkSshTunnel_getDynamicPortForwarding(z.ckObj))
    return v != 0
}
// property setter: DynamicPortForwarding
func (z *SshTunnel) SetDynamicPortForwarding(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSshTunnel_putDynamicPortForwarding(z.ckObj,v)
}

// property: HostKeyFingerprint
// Set after connecting to an SSH server. The format of the fingerprint looks like
// this: "ssh-rsa 1024 68:ff:d1:4e:6c:ff:d7:b0:d6:58:73:85:07:bc:2e:d5"
func (z *SshTunnel) HostKeyFingerprint() string {
    return C.GoString(C.CkSshTunnel_hostKeyFingerprint(z.ckObj))
}

// property: HttpProxyAuthMethod
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy authentication method name. Valid choices are "Basic" or "NTLM".
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through an HTTP proxy.
//
func (z *SshTunnel) HttpProxyAuthMethod() string {
    return C.GoString(C.CkSshTunnel_httpProxyAuthMethod(z.ckObj))
}
// property setter: HttpProxyAuthMethod
func (z *SshTunnel) SetHttpProxyAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putHttpProxyAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyDomain
// The NTLM authentication domain (optional) if NTLM authentication is used w/ the
// HTTP proxy.
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through an HTTP proxy.
//
func (z *SshTunnel) HttpProxyDomain() string {
    return C.GoString(C.CkSshTunnel_httpProxyDomain(z.ckObj))
}
// property setter: HttpProxyDomain
func (z *SshTunnel) SetHttpProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putHttpProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyHostname
// If an HTTP proxy is to be used, set this property to the HTTP proxy hostname or
// IPv4 address (in dotted decimal notation).
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through an HTTP proxy.
//
func (z *SshTunnel) HttpProxyHostname() string {
    return C.GoString(C.CkSshTunnel_httpProxyHostname(z.ckObj))
}
// property setter: HttpProxyHostname
func (z *SshTunnel) SetHttpProxyHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putHttpProxyHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPassword
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy password.
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through an HTTP proxy.
//
func (z *SshTunnel) HttpProxyPassword() string {
    return C.GoString(C.CkSshTunnel_httpProxyPassword(z.ckObj))
}
// property setter: HttpProxyPassword
func (z *SshTunnel) SetHttpProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putHttpProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPort
// If an HTTP proxy is to be used, set this property to the HTTP proxy port number.
// (Two commonly used HTTP proxy ports are 8080 and 3128.)
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through an HTTP proxy.
//
func (z *SshTunnel) HttpProxyPort() int {
    return int(C.CkSshTunnel_getHttpProxyPort(z.ckObj))
}
// property setter: HttpProxyPort
func (z *SshTunnel) SetHttpProxyPort(value int) {
    C.CkSshTunnel_putHttpProxyPort(z.ckObj,C.int(value))
}

// property: HttpProxyUsername
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy login name.
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through an HTTP proxy.
//
func (z *SshTunnel) HttpProxyUsername() string {
    return C.GoString(C.CkSshTunnel_httpProxyUsername(z.ckObj))
}
// property setter: HttpProxyUsername
func (z *SshTunnel) SetHttpProxyUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putHttpProxyUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IdleTimeoutMs
// A tunnel will fail when progress for sending or receiving data halts for more
// than this number of milliseconds. The default value is 10,000 (which is 10
// seconds). A timeout of 0 indicates an infinite wait time (i.e. no timeout).
func (z *SshTunnel) IdleTimeoutMs() int {
    return int(C.CkSshTunnel_getIdleTimeoutMs(z.ckObj))
}
// property setter: IdleTimeoutMs
func (z *SshTunnel) SetIdleTimeoutMs(value int) {
    C.CkSshTunnel_putIdleTimeoutMs(z.ckObj,C.int(value))
}

// property: InboundSocksPassword
// If dynamic port forwarding is used, then this may be set to the password
// required for authenticating inbound connections.
func (z *SshTunnel) InboundSocksPassword() string {
    return C.GoString(C.CkSshTunnel_inboundSocksPassword(z.ckObj))
}
// property setter: InboundSocksPassword
func (z *SshTunnel) SetInboundSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putInboundSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: InboundSocksUsername
// If dynamic port forwarding is used, then this may be set to the username
// required for authenticating inbound connections. If no username is set, then the
// inbound connection needs no authentication.
func (z *SshTunnel) InboundSocksUsername() string {
    return C.GoString(C.CkSshTunnel_inboundSocksUsername(z.ckObj))
}
// property setter: InboundSocksUsername
func (z *SshTunnel) SetInboundSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putInboundSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IsAccepting
// true if a background thread is running and accepting connections.
func (z *SshTunnel) IsAccepting() bool {
    v := int(C.CkSshTunnel_getIsAccepting(z.ckObj))
    return v != 0
}

// property: KeepAcceptLog
// If true, then an in-memory log of the listen thread is kept. The default value
// is false.
func (z *SshTunnel) KeepAcceptLog() bool {
    v := int(C.CkSshTunnel_getKeepAcceptLog(z.ckObj))
    return v != 0
}
// property setter: KeepAcceptLog
func (z *SshTunnel) SetKeepAcceptLog(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSshTunnel_putKeepAcceptLog(z.ckObj,v)
}

// property: KeepTunnelLog
// If true, then a log of the SSH tunnel thread activity is kept in memory. The
// default value is false.
func (z *SshTunnel) KeepTunnelLog() bool {
    v := int(C.CkSshTunnel_getKeepTunnelLog(z.ckObj))
    return v != 0
}
// property setter: KeepTunnelLog
func (z *SshTunnel) SetKeepTunnelLog(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSshTunnel_putKeepTunnelLog(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *SshTunnel) LastErrorHtml() string {
    return C.GoString(C.CkSshTunnel_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *SshTunnel) LastErrorText() string {
    return C.GoString(C.CkSshTunnel_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *SshTunnel) LastErrorXml() string {
    return C.GoString(C.CkSshTunnel_lastErrorXml(z.ckObj))
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
func (z *SshTunnel) LastMethodSuccess() bool {
    v := int(C.CkSshTunnel_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *SshTunnel) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSshTunnel_putLastMethodSuccess(z.ckObj,v)
}

// property: ListenBindIpAddress
// In most cases, this property does not need to be set. It is provided for cases
// where it is required to bind the listen socket to a specific IP address (usually
// for computers with multiple network interfaces or IP addresses). For computers
// with a single network interface (i.e. most computers), this property should not
// be set. For multihoming computers, the default IP address is automatically used
// if this property is not set.
func (z *SshTunnel) ListenBindIpAddress() string {
    return C.GoString(C.CkSshTunnel_listenBindIpAddress(z.ckObj))
}
// property setter: ListenBindIpAddress
func (z *SshTunnel) SetListenBindIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putListenBindIpAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ListenPort
// If a port number equal to 0 is passed to BeginAccepting, then this property will
// contain the actual allocated port number used. Otherwise it is equal to the port
// number passed to BeginAccepting, or 0 if BeginAccepting was never called.
func (z *SshTunnel) ListenPort() int {
    return int(C.CkSshTunnel_getListenPort(z.ckObj))
}

// property: OutboundBindIpAddress
// In most cases, this property does not need to be set. It is provided for cases
// where it is required to bind the socket that is to connect to the SSH server (in
// the background thread) to a specific IP address (usually for computers with
// multiple network interfaces or IP addresses). For computers with a single
// network interface (i.e. most computers), this property should not be set. For
// multihoming computers, the default IP address is automatically used if this
// property is not set.
func (z *SshTunnel) OutboundBindIpAddress() string {
    return C.GoString(C.CkSshTunnel_outboundBindIpAddress(z.ckObj))
}
// property setter: OutboundBindIpAddress
func (z *SshTunnel) SetOutboundBindIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putOutboundBindIpAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OutboundBindPort
// Unless there is a specific requirement for binding to a specific port, leave
// this unset (at the default value of 0). (99.9% of all users should never need to
// set this property.)
func (z *SshTunnel) OutboundBindPort() int {
    return int(C.CkSshTunnel_getOutboundBindPort(z.ckObj))
}
// property setter: OutboundBindPort
func (z *SshTunnel) SetOutboundBindPort(value int) {
    C.CkSshTunnel_putOutboundBindPort(z.ckObj,C.int(value))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *SshTunnel) PreferIpv6() bool {
    v := int(C.CkSshTunnel_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *SshTunnel) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSshTunnel_putPreferIpv6(z.ckObj,v)
}

// property: SocksHostname
// The SOCKS4/SOCKS5 hostname or IPv4 address (in dotted decimal notation). This
// property is only used if the SocksVersion property is set to 4 or 5).
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through a SOCKS4 or SOCKS5 proxy.
//
func (z *SshTunnel) SocksHostname() string {
    return C.GoString(C.CkSshTunnel_socksHostname(z.ckObj))
}
// property setter: SocksHostname
func (z *SshTunnel) SetSocksHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putSocksHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPassword
// The SOCKS5 password (if required). The SOCKS4 protocol does not include the use
// of a password, so this does not apply to SOCKS4.
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through a SOCKS4 or SOCKS5 proxy.
//
func (z *SshTunnel) SocksPassword() string {
    return C.GoString(C.CkSshTunnel_socksPassword(z.ckObj))
}
// property setter: SocksPassword
func (z *SshTunnel) SetSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPort
// The SOCKS4/SOCKS5 proxy port. The default value is 1080. This property only
// applies if a SOCKS proxy is used (if the SocksVersion property is set to 4 or
// 5).
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through a SOCKS4 or SOCKS5 proxy.
//
func (z *SshTunnel) SocksPort() int {
    return int(C.CkSshTunnel_getSocksPort(z.ckObj))
}
// property setter: SocksPort
func (z *SshTunnel) SetSocksPort(value int) {
    C.CkSshTunnel_putSocksPort(z.ckObj,C.int(value))
}

// property: SocksUsername
// The SOCKS4/SOCKS5 proxy username. This property is only used if the SocksVersion
// property is set to 4 or 5).
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through a SOCKS4 or SOCKS5 proxy.
//
func (z *SshTunnel) SocksUsername() string {
    return C.GoString(C.CkSshTunnel_socksUsername(z.ckObj))
}
// property setter: SocksUsername
func (z *SshTunnel) SetSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksVersion
// SocksVersion May be set to one of the following integer values:
// 
// 0 - No SOCKS proxy is used. This is the default.
// 4 - Connect via a SOCKS4 proxy.
// 5 - Connect via a SOCKS5 proxy.
// 
// Note: This is for the outbound connection to the SSH server. It is used when the
// outbound connection to the SSH server must go through a SOCKS4 or SOCKS5 proxy.
//
func (z *SshTunnel) SocksVersion() int {
    return int(C.CkSshTunnel_getSocksVersion(z.ckObj))
}
// property setter: SocksVersion
func (z *SshTunnel) SetSocksVersion(value int) {
    C.CkSshTunnel_putSocksVersion(z.ckObj,C.int(value))
}

// property: SoRcvBuf
// Sets the receive buffer size socket option. Normally, this property should be
// left unchanged. The default value is 4194304.
// 
// This property can be increased if download performance seems slow. It is
// recommended to be a multiple of 4096.
//
func (z *SshTunnel) SoRcvBuf() int {
    return int(C.CkSshTunnel_getSoRcvBuf(z.ckObj))
}
// property setter: SoRcvBuf
func (z *SshTunnel) SetSoRcvBuf(value int) {
    C.CkSshTunnel_putSoRcvBuf(z.ckObj,C.int(value))
}

// property: SoSndBuf
// Sets the send buffer size socket option. Normally, this property should be left
// unchanged. The default value is 262144.
// 
// This property can be increased if upload performance seems slow. It is
// recommended to be a multiple of 4096. Testing with sizes such as 512K and 1MB is
// reasonable.
//
func (z *SshTunnel) SoSndBuf() int {
    return int(C.CkSshTunnel_getSoSndBuf(z.ckObj))
}
// property setter: SoSndBuf
func (z *SshTunnel) SetSoSndBuf(value int) {
    C.CkSshTunnel_putSoSndBuf(z.ckObj,C.int(value))
}

// property: TcpNoDelay
// Controls whether the TCP_NODELAY socket option is used for the underlying TCP/IP
// socket. The default value is false. Setting this property equal to true
// disables the Nagle algorithm and allows for better performance when small
// amounts of data are sent.
func (z *SshTunnel) TcpNoDelay() bool {
    v := int(C.CkSshTunnel_getTcpNoDelay(z.ckObj))
    return v != 0
}
// property setter: TcpNoDelay
func (z *SshTunnel) SetTcpNoDelay(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSshTunnel_putTcpNoDelay(z.ckObj,v)
}

// property: TunnelLog
// Contains an in-memory log of the SSH tunnel thread. This will only contain
// content if the KeepTunnelLog property is true.
func (z *SshTunnel) TunnelLog() string {
    return C.GoString(C.CkSshTunnel_tunnelLog(z.ckObj))
}
// property setter: TunnelLog
func (z *SshTunnel) SetTunnelLog(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putTunnelLog(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TunnelLogPath
// Set to keep a log file of the SSH tunnel thread.
func (z *SshTunnel) TunnelLogPath() string {
    return C.GoString(C.CkSshTunnel_tunnelLogPath(z.ckObj))
}
// property setter: TunnelLogPath
func (z *SshTunnel) SetTunnelLogPath(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putTunnelLogPath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
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
//     "NoKeepAliveIgnoreMsg" - (Added in v9.5.0.76) Prevents the default behavior
//     of the SSH tunnel sending an "ignore" message every 20 seconds to keep an unused
//     connection alive.
//     "ProtectFromVpn" - Introduced in v9.5.0.80. On Android systems, will bypass
//     any VPN that may be installed or active.
//
func (z *SshTunnel) UncommonOptions() string {
    return C.GoString(C.CkSshTunnel_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *SshTunnel) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkSshTunnel_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *SshTunnel) VerboseLogging() bool {
    v := int(C.CkSshTunnel_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *SshTunnel) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSshTunnel_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *SshTunnel) Version() string {
    return C.GoString(C.CkSshTunnel_version(z.ckObj))
}
// Authenticates with the SSH server using public-key authentication. The
// corresponding public key must have been installed on the SSH server for the
// username. Authentication will succeed if the matching privateKey is provided.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *SshTunnel) AuthenticatePk(arg1 string, arg2 *SshKey) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSshTunnel_AuthenticatePk(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the AuthenticatePk method
func (z *SshTunnel) AuthenticatePkAsync(arg1 string, arg2 *SshKey, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSshTunnel_AuthenticatePkAsync(z.ckObj, cstr1, arg2.ckObj)

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
func (z *SshTunnel) AuthenticatePw(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSshTunnel_AuthenticatePw(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AuthenticatePw method
func (z *SshTunnel) AuthenticatePwAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSshTunnel_AuthenticatePwAsync(z.ckObj, cstr1, cstr2)

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
func (z *SshTunnel) AuthenticatePwPk(arg1 string, arg2 string, arg3 *SshKey) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSshTunnel_AuthenticatePwPk(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AuthenticatePwPk method
func (z *SshTunnel) AuthenticatePwPkAsync(arg1 string, arg2 string, arg3 *SshKey, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSshTunnel_AuthenticatePwPkAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// The same as AuthenticatePw, except the login and password strings are passed in
// secure string objects.
func (z *SshTunnel) AuthenticateSecPw(arg1 *SecureString, arg2 *SecureString) bool {

    v := C.CkSshTunnel_AuthenticateSecPw(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Asynchronous version of the AuthenticateSecPw method
func (z *SshTunnel) AuthenticateSecPwAsync(arg1 *SecureString, arg2 *SecureString, c chan *Task) {

    hTask := C.CkSshTunnel_AuthenticateSecPwAsync(z.ckObj, arg1.ckObj, arg2.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as AuthenticatePwPk, except the login and password strings are passed
// in secure string objects.
func (z *SshTunnel) AuthenticateSecPwPk(arg1 *SecureString, arg2 *SecureString, arg3 *SshKey) bool {

    v := C.CkSshTunnel_AuthenticateSecPwPk(z.ckObj, arg1.ckObj, arg2.ckObj, arg3.ckObj)


    return v != 0
}


// Asynchronous version of the AuthenticateSecPwPk method
func (z *SshTunnel) AuthenticateSecPwPkAsync(arg1 *SecureString, arg2 *SecureString, arg3 *SshKey, c chan *Task) {

    hTask := C.CkSshTunnel_AuthenticateSecPwPkAsync(z.ckObj, arg1.ckObj, arg2.ckObj, arg3.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Starts a background thread for listening on listenPort. A new SSH tunnel is created
// and managed for each accepted connection. SSH tunnels are managed in a 2nd
// background thread: the SSH tunnel pool thread.
// 
// BeginAccepting starts a background thread that creates a socket, binds to the
// port, and begins listening. If the bind fails (meaning something else may have
// already bound to the same port), then the background thread exits. You may check
// to see if BeginAccepting succeeds by waiting a short time (perhaps 50 millisec)
// and then examine the IsAccepting property. If it is false, then BeginAccepting
// failed.
// 
// Important: The listenPort must be a port number that nothing else on the local
// computer is listening on.
//
func (z *SshTunnel) BeginAccepting(arg1 int) bool {

    v := C.CkSshTunnel_BeginAccepting(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the BeginAccepting method
func (z *SshTunnel) BeginAcceptingAsync(arg1 int, c chan *Task) {

    hTask := C.CkSshTunnel_BeginAcceptingAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Closes the SSH tunnel and disconnects all existing clients. If waitForThreads is true,
// the method will wait for the tunnel and client threads to exit before returning.
func (z *SshTunnel) CloseTunnel(arg1 bool) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkSshTunnel_CloseTunnel(z.ckObj, b1)


    return v != 0
}


// Connects to the SSH server to be used for SSH tunneling.
// 
// Important: Make sure to call CloseTunnel when finished with the tunnel, such as
// before exiting your program.
//
func (z *SshTunnel) Connect(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSshTunnel_Connect(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Connect method
func (z *SshTunnel) ConnectAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSshTunnel_ConnectAsync(z.ckObj, cstr1, C.int(arg2))

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
func (z *SshTunnel) ConnectThroughSsh(arg1 *Ssh, arg2 string, arg3 int) bool {
    cstr2 := C.CString(arg2)

    v := C.CkSshTunnel_ConnectThroughSsh(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the ConnectThroughSsh method
func (z *SshTunnel) ConnectThroughSshAsync(arg1 *Ssh, arg2 string, arg3 int, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkSshTunnel_ConnectThroughSshAsync(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

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
func (z *SshTunnel) ContinueKeyboardAuth(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSshTunnel_continueKeyboardAuth(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ContinueKeyboardAuth method
func (z *SshTunnel) ContinueKeyboardAuthAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSshTunnel_ContinueKeyboardAuthAsync(z.ckObj, cstr1)

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


// Disconnects all clients, keeping the SSH tunnel open. If waitForThreads is true, the
// method will wait for the client threads to exit before returning.
func (z *SshTunnel) DisconnectAllClients(arg1 bool) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkSshTunnel_DisconnectAllClients(z.ckObj, b1)


    return v != 0
}


// Returns the current state of existing tunnels in an XML string.
// return a string or nil if failed.
func (z *SshTunnel) GetCurrentState() *string {

    retStr := C.CkSshTunnel_getCurrentState(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if connected to the SSH server. Returns false if the connection
// has been lost (or was never established).
func (z *SshTunnel) IsSshConnected() bool {

    v := C.CkSshTunnel_IsSshConnected(z.ckObj)


    return v != 0
}


// Loads the caller of the task's async method.
func (z *SshTunnel) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkSshTunnel_LoadTaskCaller(z.ckObj, arg1.ckObj)


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
func (z *SshTunnel) StartKeyboardAuth(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSshTunnel_startKeyboardAuth(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the StartKeyboardAuth method
func (z *SshTunnel) StartKeyboardAuthAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSshTunnel_StartKeyboardAuthAsync(z.ckObj, cstr1)

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


// Stops the listen background thread. It is possible to continue accepting
// connections by re-calling BeginAccepting. If waitForThread is true, the method will
// wait for the listen thread to exit before returning.
func (z *SshTunnel) StopAccepting(arg1 bool) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkSshTunnel_StopAccepting(z.ckObj, b1)


    return v != 0
}


// Unlocks the component. This must be called once prior to calling any other
// method. A fully-functional 30-day trial is automatically started when an
// arbitrary string is passed to this method. For example, passing "Hello", or
// "abc123" will unlock the component for the 1st thirty days after the initial
// install.
func (z *SshTunnel) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSshTunnel_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


