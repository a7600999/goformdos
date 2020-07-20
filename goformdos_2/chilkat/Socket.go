// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkSocket.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewSocket() *Socket {
	return &Socket{C.CkSocket_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Socket) DisposeSocket() {
    if z != nil {
	C.CkSocket_Dispose(z.ckObj)
	}
}




func (z *Socket) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkSocket_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Socket) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkSocket_setExternalProgress(z.ckObj,1)
}

func (z *Socket) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkSocket_setExternalProgress(z.ckObj,1)
}

func (z *Socket) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkSocket_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Socket) AbortCurrent() bool {
    v := int(C.CkSocket_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Socket) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putAbortCurrent(z.ckObj,v)
}

// property: AcceptFailReason
// If a AcceptNextConnection method fails, this property can be checked to
// determine the reason for failure.
// 
// Note: If accepting a TLS connection, then this property can also have any of the
// values listed for the ReceiveFailReason and SendFailReason properties (because
// the TLS handshake involves sending/receiving on the initial TCP socket).
// 
// Possible values are:
// 
// 0 = Success
// 1 = An async operation is  in progress.
// 3 = An unspecified internal failure, perhaps out-of-memory, caused the failure.
// 5 = Timeout.  No connections were accepted in the amount of time alotted.
// 6 = The receive was aborted by the application in an event callback.
// 9 = An unspecified fatal socket error occurred (less common).
// 20 = Must first bind and listen on a port.
// 99 = The component is not unlocked.
// 
// Errors Relating to the SSL/TLS Handshake:
// 100 = TLS internal error.
// 102 = Unexpected handshake message.
// 109 = Failed to read handshake messages.
// 114 = Failed to send change cipher spec handshake message.
// 115 = Failed to send finished handshake message.
// 116 = Client's Finished message is invalid.
// 117 = Unable to agree on TLS protocol version.
// 118 = Unable to agree on a cipher spec.
// 119 = Failed to read the client's hello message.
// 120 = Failed to send handshake messages.
// 121 = Failed to process client cert message.
// 122 = Failed to process client cert URL message.
// 123 = Failed to process client key exchange message.
// 124 = Failed to process certificate verify message.
// 125 = Received and rejected an SSL 2.0 connection attempt.
func (z *Socket) AcceptFailReason() int {
    return int(C.CkSocket_getAcceptFailReason(z.ckObj))
}

// property: AlpnProtocol
// For TLS connections. Can be set to the name of an application layer protocol.
// This causes the ALPN extension to be added to the TLS ClientHello with the given
// ALPN protocol name.
func (z *Socket) AlpnProtocol() string {
    return C.GoString(C.CkSocket_alpnProtocol(z.ckObj))
}
// property setter: AlpnProtocol
func (z *Socket) SetAlpnProtocol(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putAlpnProtocol(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: BandwidthThrottleDown
// If non-zero, limits (throttles) the receiving bandwidth to approximately this
// maximum number of bytes per second. The default value of this property is 0.
func (z *Socket) BandwidthThrottleDown() int {
    return int(C.CkSocket_getBandwidthThrottleDown(z.ckObj))
}
// property setter: BandwidthThrottleDown
func (z *Socket) SetBandwidthThrottleDown(value int) {
    C.CkSocket_putBandwidthThrottleDown(z.ckObj,C.int(value))
}

// property: BandwidthThrottleUp
// If non-zero, limits (throttles) the sending bandwidth to approximately this
// maximum number of bytes per second. The default value of this property is 0.
func (z *Socket) BandwidthThrottleUp() int {
    return int(C.CkSocket_getBandwidthThrottleUp(z.ckObj))
}
// property setter: BandwidthThrottleUp
func (z *Socket) SetBandwidthThrottleUp(value int) {
    C.CkSocket_putBandwidthThrottleUp(z.ckObj,C.int(value))
}

// property: BigEndian
// Applies to the SendCount and ReceiveCount methods. If BigEndian is set to true
// (the default) then the 4-byte count is in big endian format. Otherwise it is
// little endian.
func (z *Socket) BigEndian() bool {
    v := int(C.CkSocket_getBigEndian(z.ckObj))
    return v != 0
}
// property setter: BigEndian
func (z *Socket) SetBigEndian(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putBigEndian(z.ckObj,v)
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
func (z *Socket) ClientIpAddress() string {
    return C.GoString(C.CkSocket_clientIpAddress(z.ckObj))
}
// property setter: ClientIpAddress
func (z *Socket) SetClientIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putClientIpAddress(z.ckObj,cStr)
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
func (z *Socket) ClientPort() int {
    return int(C.CkSocket_getClientPort(z.ckObj))
}
// property setter: ClientPort
func (z *Socket) SetClientPort(value int) {
    C.CkSocket_putClientPort(z.ckObj,C.int(value))
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
// 108 = App-defined server certificate requirements failure.
// 109 = Failed to read handshake messages.
// 110 = Failed to send client certificate handshake message.
// 111 = Failed to send client key exchange handshake message.
// 112 = Client certificate's private key not accessible.
// 113 = Failed to send client cert verify handshake message.
// 114 = Failed to send change cipher spec handshake message.
// 115 = Failed to send finished handshake message.
// 116 = Server's Finished message is invalid.
//
func (z *Socket) ConnectFailReason() int {
    return int(C.CkSocket_getConnectFailReason(z.ckObj))
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
func (z *Socket) DebugLogFilePath() string {
    return C.GoString(C.CkSocket_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Socket) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ElapsedSeconds
// Contains the number of seconds since the last call to StartTiming, otherwise
// contains 0. (The StartTiming method and ElapsedSeconds property is provided for
// convenience.)
func (z *Socket) ElapsedSeconds() int {
    return int(C.CkSocket_getElapsedSeconds(z.ckObj))
}

// property: HeartbeatMs
// The number of milliseconds between periodic heartbeat callbacks for blocking
// socket operations (connect, accept, dns query, send, receive). Set this to 0 to
// disable heartbeat events. The default value is 1000 (i.e. 1 heartbeat callback
// per second).
func (z *Socket) HeartbeatMs() int {
    return int(C.CkSocket_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Socket) SetHeartbeatMs(value int) {
    C.CkSocket_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: HttpProxyAuthMethod
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy authentication method name. Valid choices are "Basic" or "NTLM".
func (z *Socket) HttpProxyAuthMethod() string {
    return C.GoString(C.CkSocket_httpProxyAuthMethod(z.ckObj))
}
// property setter: HttpProxyAuthMethod
func (z *Socket) SetHttpProxyAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putHttpProxyAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyDomain
// The NTLM authentication domain (optional) if NTLM authentication is used.
func (z *Socket) HttpProxyDomain() string {
    return C.GoString(C.CkSocket_httpProxyDomain(z.ckObj))
}
// property setter: HttpProxyDomain
func (z *Socket) SetHttpProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putHttpProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyForHttp
// If this connection is effectively used to send HTTP requests, then set this
// property to true when using an HTTP proxy. The default value of this property
// is false.
// 
// This is because an HTTP proxy used for other protocols (IMAP, SMTP, SSH, FTP,
// etc.) can require some internal differences in behavior (i.e. how we do things).
// 
// For example, the Chilkat REST object can use this socket object's connection via
// the UseConnection method. This is a case where we know the proxied connection is
// for the HTTP protocol. Therefore we should set this property to true. (See the
// example below.)
//
func (z *Socket) HttpProxyForHttp() bool {
    v := int(C.CkSocket_getHttpProxyForHttp(z.ckObj))
    return v != 0
}
// property setter: HttpProxyForHttp
func (z *Socket) SetHttpProxyForHttp(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putHttpProxyForHttp(z.ckObj,v)
}

// property: HttpProxyHostname
// If an HTTP proxy is to be used, set this property to the HTTP proxy hostname or
// IPv4 address (in dotted decimal notation).
func (z *Socket) HttpProxyHostname() string {
    return C.GoString(C.CkSocket_httpProxyHostname(z.ckObj))
}
// property setter: HttpProxyHostname
func (z *Socket) SetHttpProxyHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putHttpProxyHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPassword
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy password.
func (z *Socket) HttpProxyPassword() string {
    return C.GoString(C.CkSocket_httpProxyPassword(z.ckObj))
}
// property setter: HttpProxyPassword
func (z *Socket) SetHttpProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putHttpProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPort
// If an HTTP proxy is to be used, set this property to the HTTP proxy port number.
// (Two commonly used HTTP proxy ports are 8080 and 3128.)
func (z *Socket) HttpProxyPort() int {
    return int(C.CkSocket_getHttpProxyPort(z.ckObj))
}
// property setter: HttpProxyPort
func (z *Socket) SetHttpProxyPort(value int) {
    C.CkSocket_putHttpProxyPort(z.ckObj,C.int(value))
}

// property: HttpProxyUsername
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy login name.
func (z *Socket) HttpProxyUsername() string {
    return C.GoString(C.CkSocket_httpProxyUsername(z.ckObj))
}
// property setter: HttpProxyUsername
func (z *Socket) SetHttpProxyUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putHttpProxyUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IsConnected
// Returns true if the socket is connected. Otherwise returns false.
// 
// Note: In general, this property indicates the last known state of the socket.
// For example, if the socket is connected, and your application does not read or
// write the socket, then IsConnected will remain true. This property is updated
// when your application tries to read or write and discovers that the socket is no
// longer connected. It is also updated if your application explicitly closes the
// socket.
//
func (z *Socket) IsConnected() bool {
    v := int(C.CkSocket_getIsConnected(z.ckObj))
    return v != 0
}

// property: KeepAlive
// Controls whether the SO_KEEPALIVE socket option is used for the underlying
// TCP/IP socket. The default value is true.
func (z *Socket) KeepAlive() bool {
    v := int(C.CkSocket_getKeepAlive(z.ckObj))
    return v != 0
}
// property setter: KeepAlive
func (z *Socket) SetKeepAlive(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putKeepAlive(z.ckObj,v)
}

// property: KeepSessionLog
// Controls whether socket (or SSL) communications are logged to the SessionLog
// string property. To turn on session logging, set this property = true,
// otherwise set to false (which is the default value).
func (z *Socket) KeepSessionLog() bool {
    v := int(C.CkSocket_getKeepSessionLog(z.ckObj))
    return v != 0
}
// property setter: KeepSessionLog
func (z *Socket) SetKeepSessionLog(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putKeepSessionLog(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Socket) LastErrorHtml() string {
    return C.GoString(C.CkSocket_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Socket) LastErrorText() string {
    return C.GoString(C.CkSocket_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Socket) LastErrorXml() string {
    return C.GoString(C.CkSocket_lastErrorXml(z.ckObj))
}

// property: LastMethodFailed
// true if the last method called on this object failed. This provides an easier
// (less confusing) way of determining whether a method such as ReceiveBytes
// succeeded or failed.
func (z *Socket) LastMethodFailed() bool {
    v := int(C.CkSocket_getLastMethodFailed(z.ckObj))
    return v != 0
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
func (z *Socket) LastMethodSuccess() bool {
    v := int(C.CkSocket_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Socket) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putLastMethodSuccess(z.ckObj,v)
}

// property: ListenIpv6
// If set to true, then a socket that listens for incoming connections (via the
// BindAndList and AcceptNextConnection method calls) will use IPv6 and not IPv4.
// The default value is false for IPv4.
func (z *Socket) ListenIpv6() bool {
    v := int(C.CkSocket_getListenIpv6(z.ckObj))
    return v != 0
}
// property setter: ListenIpv6
func (z *Socket) SetListenIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putListenIpv6(z.ckObj,v)
}

// property: ListenPort
// The BindAndListen method will find a random unused port to listen on if you bind
// to port 0. This chosen listen port is available via this property.
func (z *Socket) ListenPort() int {
    return int(C.CkSocket_getListenPort(z.ckObj))
}

// property: LocalIpAddress
// The local IP address for a bound or connected socket.
func (z *Socket) LocalIpAddress() string {
    return C.GoString(C.CkSocket_localIpAddress(z.ckObj))
}

// property: LocalPort
// The local port for a bound or connected socket.
func (z *Socket) LocalPort() int {
    return int(C.CkSocket_getLocalPort(z.ckObj))
}

// property: MaxReadIdleMs
// The maximum number of milliseconds to wait on a socket read operation while no
// additional data is forthcoming. To wait indefinitely, set this property to 0.
// The default value is 0.
func (z *Socket) MaxReadIdleMs() int {
    return int(C.CkSocket_getMaxReadIdleMs(z.ckObj))
}
// property setter: MaxReadIdleMs
func (z *Socket) SetMaxReadIdleMs(value int) {
    C.CkSocket_putMaxReadIdleMs(z.ckObj,C.int(value))
}

// property: MaxSendIdleMs
// The maximum number of milliseconds to wait for the socket to become writeable on
// a socket write operation. To wait indefinitely, set this property to 0. The
// default value is 0.
func (z *Socket) MaxSendIdleMs() int {
    return int(C.CkSocket_getMaxSendIdleMs(z.ckObj))
}
// property setter: MaxSendIdleMs
func (z *Socket) SetMaxSendIdleMs(value int) {
    C.CkSocket_putMaxSendIdleMs(z.ckObj,C.int(value))
}

// property: MyIpAddress
// The local IP address of the local computer. For multi-homed computers (i.e.
// computers with multiple IP adapters) this property returns the default IP
// address.
// 
// Note: This will be the internal IP address, not an external IP address. (For
// example, if your computer is on a LAN, it is likely to be an IP address
// beginning with "192.168.".
// 
// Important: Use LocalIpAddress and LocalIpPort to get the local IP/port for a
// bound or connected socket.
//
func (z *Socket) MyIpAddress() string {
    return C.GoString(C.CkSocket_myIpAddress(z.ckObj))
}

// property: NumReceivedClientCerts
// If the socket is the server-side of an SSL/TLS connection, the property
// represents the number of client-side certificates received during the SSL/TLS
// handshake (i.e. connection process). Each client-side cert may be retrieved by
// calling the GetReceivedClientCert method and passing an integer index value from
// 0 to N-1, where N is the number of client certs received.
// 
// Note: A client only sends a certificate if 2-way SSL/TLS is required. In other
// words, if the server demands a certificate from the client.
// 
// Important: This property should be examined on the socket object that is
// returned by AcceptNextConnection.
//
func (z *Socket) NumReceivedClientCerts() int {
    return int(C.CkSocket_getNumReceivedClientCerts(z.ckObj))
}

// property: NumSocketsInSet
// If this socket is a "socket set", then NumSocketsInSet returns the number of
// sockets contained in the set. A socket object can become a "socket set" by
// calling the TakeSocket method on one or more connected sockets. This makes it
// possible to select for reading on the set (i.e. wait for data to arrive from any
// one of multiple sockets). See the following methods and properties for more
// information: TakeSocket, SelectorIndex, SelectorReadIndex, SelectorWriteIndex,
// SelectForReading, SelectForWriting.
func (z *Socket) NumSocketsInSet() int {
    return int(C.CkSocket_getNumSocketsInSet(z.ckObj))
}

// property: NumSslAcceptableClientCAs
// If connected as an SSL/TLS client to an SSL/TLS server where the server requires
// a client-side certificate for authentication, then this property contains the
// number of acceptable certificate authorities sent by the server during
// connection establishment handshake. The GetSslAcceptableClientCaDn method may be
// called to get the Distinguished Name (DN) of each acceptable CA.
func (z *Socket) NumSslAcceptableClientCAs() int {
    return int(C.CkSocket_getNumSslAcceptableClientCAs(z.ckObj))
}

// property: ObjectId
// Each socket object is assigned a unique object ID. This ID is passed in event
// callbacks to allow your application to associate the event with the socket
// object.
func (z *Socket) ObjectId() int {
    return int(C.CkSocket_getObjectId(z.ckObj))
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
func (z *Socket) PercentDoneScale() int {
    return int(C.CkSocket_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *Socket) SetPercentDoneScale(value int) {
    C.CkSocket_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *Socket) PreferIpv6() bool {
    v := int(C.CkSocket_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *Socket) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putPreferIpv6(z.ckObj,v)
}

// property: RcvBytesPerSec
// Returns the cumulative receive rate in bytes per second. The measurement
// includes the overhead bytes for protocols such as TLS or SSH tunneling. For
// example, if 1000 application bytes are received, the actual number of raw bytes
// received on a TLS connection is greater. This property measures the actual
// number of raw bytes received in a given time period. The ResetPerf method can be
// called to reset this property value and to begin the performance measurement
// afresh.
func (z *Socket) RcvBytesPerSec() int {
    return int(C.CkSocket_getRcvBytesPerSec(z.ckObj))
}

// property: ReceivedCount
// Any method that receives data will increase the value of this property by the
// number of bytes received. The application may reset this property to 0 at any
// point. It is provided as a way to keep count of the total number of bytes
// received on a socket connection, regardless of which method calls are used to
// receive the data.
// 
// Note: The ReceivedCount may be larger than the number of bytes returned by some
// methods. For methods such as ReceiveUntilMatch, the excess received on the
// socket (beyond the match), is buffered by Chilkat for subsequent method calls.
// The ReceivedCount is updated based on the actual number of bytes received on the
// underlying socket in real-time. (The ReceivedCount does not include the overhead
// bytes associated with the TLS and/or SSH protocols.
//
func (z *Socket) ReceivedCount() int {
    return int(C.CkSocket_getReceivedCount(z.ckObj))
}
// property setter: ReceivedCount
func (z *Socket) SetReceivedCount(value int) {
    C.CkSocket_putReceivedCount(z.ckObj,C.int(value))
}

// property: ReceivedInt
// Contains the last integer received via a call to ReceiveByte, ReceiveInt16, or
// ReceiveInt32.
func (z *Socket) ReceivedInt() int {
    return int(C.CkSocket_getReceivedInt(z.ckObj))
}
// property setter: ReceivedInt
func (z *Socket) SetReceivedInt(value int) {
    C.CkSocket_putReceivedInt(z.ckObj,C.int(value))
}

// property: ReceiveFailReason
// If a Receive method fails, this property can be checked to determine the reason
// for failure.
// 
// Possible values are:
// 0 = Success
// 1 = An async receive operation is already in progress.
// 2 = The socket is not connected, such as if it was never connected, or if the connection was previously lost.
// 3 = An unspecified internal failure, perhaps out-of-memory, caused the failure.
// 4 = Invalid parameters were passed to the receive method call.
// 5 = Timeout.  Data stopped arriving for more than the amount of time specified by the MaxReadIdleMs property.
// 6 = The receive was aborted by the application in an event callback.
// 7 = The connection was lost -- the remote peer reset the connection. (The connection was forcibly closed by the peer.)
// 8 = An established connection was aborted by the software in your host machine. (See https://www.chilkatsoft.com/p/p_299.asp )
// 9 = An unspecified fatal socket error occurred (less common).
// 10 = The connection was closed by the peer.
//
func (z *Socket) ReceiveFailReason() int {
    return int(C.CkSocket_getReceiveFailReason(z.ckObj))
}

// property: ReceivePacketSize
// The number of bytes to receive at a time (internally). This setting has an
// effect on methods such as ReadBytes and ReadString where the number of bytes to
// read is not explicitly specified. The default value is 4096.
func (z *Socket) ReceivePacketSize() int {
    return int(C.CkSocket_getReceivePacketSize(z.ckObj))
}
// property setter: ReceivePacketSize
func (z *Socket) SetReceivePacketSize(value int) {
    C.CkSocket_putReceivePacketSize(z.ckObj,C.int(value))
}

// property: RemoteIpAddress
// When a socket is connected, the remote IP address of the connected peer is
// available in this property.
func (z *Socket) RemoteIpAddress() string {
    return C.GoString(C.CkSocket_remoteIpAddress(z.ckObj))
}

// property: RemotePort
// When a socket is connected, the remote port of the connected peer is available
// in this property.
func (z *Socket) RemotePort() int {
    return int(C.CkSocket_getRemotePort(z.ckObj))
}

// property: RequireSslCertVerify
// If true, then the SSL/TLS client will verify the server's SSL certificate. The
// certificate is expired, or if the cert's signature is invalid, the connection is
// not allowed. The default value of this property is false.
func (z *Socket) RequireSslCertVerify() bool {
    v := int(C.CkSocket_getRequireSslCertVerify(z.ckObj))
    return v != 0
}
// property setter: RequireSslCertVerify
func (z *Socket) SetRequireSslCertVerify(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putRequireSslCertVerify(z.ckObj,v)
}

// property: SelectorIndex
// If this socket contains a collection of connected sockets (i.e. it is a "socket
// set") then method calls and property gets/sets are routed to the contained
// socket indicated by this property. Indexing begins at 0. See the TakeSocket
// method and SelectForReading method for more information.
func (z *Socket) SelectorIndex() int {
    return int(C.CkSocket_getSelectorIndex(z.ckObj))
}
// property setter: SelectorIndex
func (z *Socket) SetSelectorIndex(value int) {
    C.CkSocket_putSelectorIndex(z.ckObj,C.int(value))
}

// property: SelectorReadIndex
// When SelectForReading returns a number greater than 0 indicating that 1 or more
// sockets are ready for reading, this property is used to select the socket in the
// "ready set" for reading. See the example below:
func (z *Socket) SelectorReadIndex() int {
    return int(C.CkSocket_getSelectorReadIndex(z.ckObj))
}
// property setter: SelectorReadIndex
func (z *Socket) SetSelectorReadIndex(value int) {
    C.CkSocket_putSelectorReadIndex(z.ckObj,C.int(value))
}

// property: SelectorWriteIndex
// When SelectForWriting returns a number greater than 0 indicating that one or
// more sockets are ready for writing, this property is used to select the socket
// in the "ready set" for writing.
func (z *Socket) SelectorWriteIndex() int {
    return int(C.CkSocket_getSelectorWriteIndex(z.ckObj))
}
// property setter: SelectorWriteIndex
func (z *Socket) SetSelectorWriteIndex(value int) {
    C.CkSocket_putSelectorWriteIndex(z.ckObj,C.int(value))
}

// property: SendBytesPerSec
// Returns the cumulative send rate in bytes per second. The measurement includes
// the overhead bytes for protocols such as TLS or SSH tunneling. For example, if
// 1000 application bytes are sent, the actual number of raw bytes sent on a TLS
// connection is greater. This property measures the actual number of raw bytes
// sent in a given time period. The ResetPerf method can be called to reset this
// property value and to begin the performance measurement afresh.
func (z *Socket) SendBytesPerSec() int {
    return int(C.CkSocket_getSendBytesPerSec(z.ckObj))
}

// property: SendFailReason
// If a Send method fails, this property can be checked to determine the reason for
// failure.
// 
// Possible values are:
// 0 = Success
// 1 = An async receive operation is already in progress.
// 2 = The socket is not connected, such as if it was never connected, or if the connection was previously lost.
// 3 = An unspecified internal failure, perhaps out-of-memory, caused the failure.
// 4 = Invalid parameters were passed to the receive method call.
// 5 = Timeout.  Data stopped arriving for more than the amount of time specified by the MaxReadIdleMs property.
// 6 = The receive was aborted by the application in an event callback.
// 7 = The connection was lost -- the remote peer reset the connection. (The connection was forcibly closed by the peer.)
// 8 = An established connection was aborted by the software in your host machine. (See https://www.chilkatsoft.com/p/p_299.asp )
// 9 = An unspecified fatal socket error occurred (less common).
// 10 = The connection was closed by the peer.
// 11 = Decoding error (possible in SendString when coverting to the StringCharset, or in SendBytesENC).
//
func (z *Socket) SendFailReason() int {
    return int(C.CkSocket_getSendFailReason(z.ckObj))
}

// property: SendPacketSize
// The number of bytes to send at a time (internally). This can also be though of
// as the "chunk size". If a large amount of data is to be sent, the data is sent
// in chunks equal to this size in bytes. The default value is 65535. (Note: This
// only applies to non-SSL/TLS connections. SSL and TLS have their own pre-defined
// packet sizes.)
func (z *Socket) SendPacketSize() int {
    return int(C.CkSocket_getSendPacketSize(z.ckObj))
}
// property setter: SendPacketSize
func (z *Socket) SetSendPacketSize(value int) {
    C.CkSocket_putSendPacketSize(z.ckObj,C.int(value))
}

// property: SessionLog
// Contains a log of the bytes sent and received on this socket. The KeepSessionLog
// property must be set to true for logging to occur.
func (z *Socket) SessionLog() string {
    return C.GoString(C.CkSocket_sessionLog(z.ckObj))
}

// property: SessionLogEncoding
// Controls how the data is encoded in the SessionLog. Possible values are "esc"
// and "hex". The default value is "esc".
// 
// When set to "hex", the bytes are encoded as a hexidecimalized string. The "esc"
// encoding is a C-string like encoding, and is more compact than hex if most of
// the data to be logged is text. Printable us-ascii chars are unmodified. Common
// "C" control chars are represented as "\r", "\n", "\t", etc. Non-printable and
// byte values greater than 0x80 are escaped using a backslash and hex encoding:
// \xHH. Certain printable chars are backslashed: SPACE, double-quote,
// single-quote, etc.
//
func (z *Socket) SessionLogEncoding() string {
    return C.GoString(C.CkSocket_sessionLogEncoding(z.ckObj))
}
// property setter: SessionLogEncoding
func (z *Socket) SetSessionLogEncoding(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putSessionLogEncoding(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SniHostname
// Specifies the SNI hostname to be used in the TLS ClientHello. This property is
// only needed when the domain is specified via a dotted IP address and an SNI
// hostname is desired. (Normally, Chilkat automatically uses the domain name in
// the SNI hostname TLS ClientHello extension.)
func (z *Socket) SniHostname() string {
    return C.GoString(C.CkSocket_sniHostname(z.ckObj))
}
// property setter: SniHostname
func (z *Socket) SetSniHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putSniHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksHostname
// The SOCKS4/SOCKS5 hostname or IPv4 address (in dotted decimal notation). This
// property is only used if the SocksVersion property is set to 4 or 5).
func (z *Socket) SocksHostname() string {
    return C.GoString(C.CkSocket_socksHostname(z.ckObj))
}
// property setter: SocksHostname
func (z *Socket) SetSocksHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putSocksHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPassword
// The SOCKS5 password (if required). The SOCKS4 protocol does not include the use
// of a password, so this does not apply to SOCKS4.
func (z *Socket) SocksPassword() string {
    return C.GoString(C.CkSocket_socksPassword(z.ckObj))
}
// property setter: SocksPassword
func (z *Socket) SetSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPort
// The SOCKS4/SOCKS5 proxy port. The default value is 1080. This property only
// applies if a SOCKS proxy is used (if the SocksVersion property is set to 4 or
// 5).
func (z *Socket) SocksPort() int {
    return int(C.CkSocket_getSocksPort(z.ckObj))
}
// property setter: SocksPort
func (z *Socket) SetSocksPort(value int) {
    C.CkSocket_putSocksPort(z.ckObj,C.int(value))
}

// property: SocksUsername
// The SOCKS4/SOCKS5 proxy username. This property is only used if the SocksVersion
// property is set to 4 or 5).
func (z *Socket) SocksUsername() string {
    return C.GoString(C.CkSocket_socksUsername(z.ckObj))
}
// property setter: SocksUsername
func (z *Socket) SetSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksVersion
// SocksVersion May be set to one of the following integer values:
// 
// 0 - No SOCKS proxy is used. This is the default.
// 4 - Connect via a SOCKS4 proxy.
// 5 - Connect via a SOCKS5 proxy.
//
func (z *Socket) SocksVersion() int {
    return int(C.CkSocket_getSocksVersion(z.ckObj))
}
// property setter: SocksVersion
func (z *Socket) SetSocksVersion(value int) {
    C.CkSocket_putSocksVersion(z.ckObj,C.int(value))
}

// property: SoRcvBuf
// Sets the receive buffer size socket option. Normally, this property should be
// left unchanged. The default value is 4194304.
// 
// This property can be increased if download performance seems slow. It is
// recommended to be a multiple of 4096.
//
func (z *Socket) SoRcvBuf() int {
    return int(C.CkSocket_getSoRcvBuf(z.ckObj))
}
// property setter: SoRcvBuf
func (z *Socket) SetSoRcvBuf(value int) {
    C.CkSocket_putSoRcvBuf(z.ckObj,C.int(value))
}

// property: SoReuseAddr
// Sets the SO_REUSEADDR socket option for a socket that will bind to a port and
// listen for incoming connections. The default value is true, meaning that the
// SO_REUSEADDR socket option is set. If the socket option must be unset, set this
// property equal to false prior to calling BindAndListen or InitSslServer.
func (z *Socket) SoReuseAddr() bool {
    v := int(C.CkSocket_getSoReuseAddr(z.ckObj))
    return v != 0
}
// property setter: SoReuseAddr
func (z *Socket) SetSoReuseAddr(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putSoReuseAddr(z.ckObj,v)
}

// property: SoSndBuf
// Sets the send buffer size socket option. Normally, this property should be left
// unchanged. The default value is 262144.
// 
// This property can be increased if upload performance seems slow. It is
// recommended to be a multiple of 4096. Testing with sizes such as 512K and 1MB is
// reasonable.
//
func (z *Socket) SoSndBuf() int {
    return int(C.CkSocket_getSoSndBuf(z.ckObj))
}
// property setter: SoSndBuf
func (z *Socket) SetSoSndBuf(value int) {
    C.CkSocket_putSoSndBuf(z.ckObj,C.int(value))
}

// property: Ssl
// Set this property to true if SSL/TLS is required for accepted connections
// (AcceptNextConnection). The default value is false.
// 
// Note: This property should have been more precisely named "RequireSslClient". It
// is a property that if set to true, requires all accepted connections use
// SSL/TLS. If a client attempts to connect but cannot establish the TLS
// connection, then it is not accepted. This property is not meant to reflect the
// current state of the connection.
// 
// The TlsVersion property shows the current or last negotiated TLS version of the
// connection. The TlsVersion will be empty for a non-SSL/TLS connection.
//
func (z *Socket) Ssl() bool {
    v := int(C.CkSocket_getSsl(z.ckObj))
    return v != 0
}
// property setter: Ssl
func (z *Socket) SetSsl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putSsl(z.ckObj,v)
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
func (z *Socket) SslAllowedCiphers() string {
    return C.GoString(C.CkSocket_sslAllowedCiphers(z.ckObj))
}
// property setter: SslAllowedCiphers
func (z *Socket) SetSslAllowedCiphers(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putSslAllowedCiphers(z.ckObj,cStr)
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
func (z *Socket) SslProtocol() string {
    return C.GoString(C.CkSocket_sslProtocol(z.ckObj))
}
// property setter: SslProtocol
func (z *Socket) SetSslProtocol(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putSslProtocol(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: StringCharset
// A charset such as "utf-8", "windows-1252", "Shift_JIS", "iso-8859-1", etc.
// Methods for sending and receiving strings will use this charset as the encoding.
// Strings sent on the socket are first converted (if necessary) to this encoding.
// When reading, it is assumed that the bytes received are converted FROM this
// charset if necessary. This ONLY APPLIES TO THE SendString and ReceiveString
// methods. The default value is "ansi".
func (z *Socket) StringCharset() string {
    return C.GoString(C.CkSocket_stringCharset(z.ckObj))
}
// property setter: StringCharset
func (z *Socket) SetStringCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putStringCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TcpNoDelay
// Controls whether the TCP_NODELAY socket option is used for the underlying TCP/IP
// socket. The default value is false. Setting the value to true disables the
// Nagle algorithm and allows for better performance when small amounts of data are
// sent on the socket connection.
func (z *Socket) TcpNoDelay() bool {
    v := int(C.CkSocket_getTcpNoDelay(z.ckObj))
    return v != 0
}
// property setter: TcpNoDelay
func (z *Socket) SetTcpNoDelay(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putTcpNoDelay(z.ckObj,v)
}

// property: TlsCipherSuite
// Contains the current or last negotiated TLS cipher suite. If no TLS connection
// has yet to be established, or if a connection as attempted and failed, then this
// will be empty. A sample cipher suite string looks like this:
// TLS_DHE_RSA_WITH_AES_256_CBC_SHA256.
func (z *Socket) TlsCipherSuite() string {
    return C.GoString(C.CkSocket_tlsCipherSuite(z.ckObj))
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
func (z *Socket) TlsPinSet() string {
    return C.GoString(C.CkSocket_tlsPinSet(z.ckObj))
}
// property setter: TlsPinSet
func (z *Socket) SetTlsPinSet(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putTlsPinSet(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TlsVersion
// Contains the current or last negotiated TLS protocol version. If no TLS
// connection has yet to be established, or if a connection as attempted and
// failed, then this will be empty. Possible values are "SSL 3.0", "TLS 1.0", "TLS
// 1.1", and "TLS 1.2".
func (z *Socket) TlsVersion() string {
    return C.GoString(C.CkSocket_tlsVersion(z.ckObj))
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
func (z *Socket) UncommonOptions() string {
    return C.GoString(C.CkSocket_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Socket) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UserData
// Provides a way to store text data with the socket object. The UserData is purely
// for convenience and is not involved in the socket communications in any way. An
// application might use this property to keep extra information associated with
// the socket.
func (z *Socket) UserData() string {
    return C.GoString(C.CkSocket_userData(z.ckObj))
}
// property setter: UserData
func (z *Socket) SetUserData(goStr string) {
    cStr := C.CString(goStr)
    C.CkSocket_putUserData(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Socket) VerboseLogging() bool {
    v := int(C.CkSocket_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Socket) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSocket_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Socket) Version() string {
    return C.GoString(C.CkSocket_version(z.ckObj))
}
// Blocking call to accept the next incoming connection on the socket. maxWaitMs
// specifies the maximum time to wait (in milliseconds). Set this to 0 to wait
// indefinitely. If successful, a new socket object is returned.
// 
// Important: If accepting an SSL/TLS connection, the SSL handshake is part of the
// connection establishment process. This involves a few back-and-forth messages
// between the client and server to establish algorithms and a shared key to create
// the secure channel. The sending and receiving of these messages are governed by
// the MaxReadIdleMs and MaxSendIdleMs properties. If these properties are set to 0
// (and this is the default unless changed by your application), then the
// AcceptNextConnection can hang indefinitely during the SSL handshake process.
// Make sure these properties are set to appropriate values before calling this
// method.
//
func (z *Socket) AcceptNextConnection(arg1 int) *Socket {

    retObj := C.CkSocket_AcceptNextConnection(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Socket{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Asynchronous version of the AcceptNextConnection method
func (z *Socket) AcceptNextConnectionAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_AcceptNextConnectionAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// If this object is a server-side socket accepting SSL/TLS connections, and wishes
// to require a client-side certificate for authentication, then it should make one
// or more calls to this method to identify the CA's it will accept for client-side
// certificates.
// 
// If no CA DN's are added by this method, then client certificates from any root
// CA are accepted.
// 
// Important: If calling this method, it must be called before calling
// InitSslServer.
//
func (z *Socket) AddSslAcceptableClientCaDn(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSocket_AddSslAcceptableClientCaDn(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Binds a TCP socket to a port and configures it to listen for incoming
// connections. The size of the backlog is passed in backLog. The backLog is necessary
// when multiple connections arrive at the same time, or close enough in time such
// that they cannot be serviced immediately. (A typical value to use for backLog is
// 5.) This method should be called once prior to receiving incoming connection
// requests via the AcceptNextConnection or AsyncAcceptStart methods.
// 
// Note:This method will find a random unused port to listen on if you bind to port
// 0. The chosen port is available via the read-only ListenPort property after this
// method returns successful.
// 
// To bind and listen using IPv6, set the ListenIpv6 property = true prior to
// calling this method.
// 
// What is a reasonable value for backLog? The answer depends on how many simultaneous
// incoming connections could be expected, and how quickly your application can
// process an incoming connection and then return to accept the next connection.
//
func (z *Socket) BindAndListen(arg1 int, arg2 int) bool {

    v := C.CkSocket_BindAndListen(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Asynchronous version of the BindAndListen method
func (z *Socket) BindAndListenAsync(arg1 int, arg2 int, c chan *Task) {

    hTask := C.CkSocket_BindAndListenAsync(z.ckObj, C.int(arg1), C.int(arg2))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Binds a TCP socket to an unused port within a port range (beginPort to endPort) and
// configures it to listen for incoming connections. The size of the backlog is
// passed in endPort. The endPort is necessary when multiple connections arrive at the
// same time, or close enough in time such that they cannot be serviced
// immediately. (A typical value to use for endPort is 5.) This method should be
// called once prior to receiving incoming connection requests via the
// AcceptNextConnection method.
// 
// To bind and listen using IPv6, set the ListenIpv6 property = true prior to
// calling this method.
// 
// Returns the port number that was bound, or -1 if no port was available or if it
// failed for some other reason.
//
func (z *Socket) BindAndListenPortRange(arg1 int, arg2 int, arg3 int) int {

    v := C.CkSocket_BindAndListenPortRange(z.ckObj, C.int(arg1), C.int(arg2), C.int(arg3))


    return int(v)
}


// Asynchronous version of the BindAndListenPortRange method
func (z *Socket) BindAndListenPortRangeAsync(arg1 int, arg2 int, arg3 int, c chan *Task) {

    hTask := C.CkSocket_BindAndListenPortRangeAsync(z.ckObj, C.int(arg1), C.int(arg2), C.int(arg3))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Convenience method for building a simple HTTP GET request from a URL.
// return a string or nil if failed.
func (z *Socket) BuildHttpGetRequest(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSocket_buildHttpGetRequest(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Determines if the socket is writeable. Returns one of the following integer
// values:
// 
// 1: If the socket is connected and ready for writing.
// 0: If a timeout occurred or if the application aborted the method during an
// event callback.
// -1: The socket is not connected.
// 
// A maxWaitMs value of 0 indicates a poll.
//
func (z *Socket) CheckWriteable(arg1 int) int {

    v := C.CkSocket_CheckWriteable(z.ckObj, C.int(arg1))


    return int(v)
}


// Asynchronous version of the CheckWriteable method
func (z *Socket) CheckWriteableAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_CheckWriteableAsync(z.ckObj, C.int(arg1))


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
func (z *Socket) ClearSessionLog()  {

    C.CkSocket_ClearSessionLog(z.ckObj)



}


// Creates a copy that shares the same underlying TCP (or SSL/TLS) connection. This
// allows for simultaneous reading/writing by different threads on the socket. When
// using asynchronous reading/writing, it is not necessary to clone the socket.
// However, if separate background threads are making synchronous calls to
// read/write, then one thread may use the original socket, and the other should
// use a clone.
func (z *Socket) CloneSocket() *Socket {

    retObj := C.CkSocket_CloneSocket(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Socket{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Cleanly terminates and closes a TCP, TLS, or SSH channel connection. The maxWaitMs
// applies to SSL/TLS connections because there is a handshake that occurs during
// secure channel shutdown.
func (z *Socket) Close(arg1 int) bool {

    v := C.CkSocket_Close(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the Close method
func (z *Socket) CloseAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_CloseAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Establishes a secure SSL/TLS or a plain non-secure TCP connection with a remote
// host:port. This is a blocking call. The maximum wait time (in milliseconds) is
// passed in maxWaitMs. This is the amount of time the app is willing to wait for the
// TCP connection to be accepted.
// 
// To establish an SSL/TLS connection, set ssl = true, otherwise set ssl =
// false for a normal TCP connection. Note: The timeouts that apply to the
// internal SSL/TLS handshaking messages are the MaxReadIdleMs and MaxSendIdleMs
// properties.
// 
// Note: Connections do not automatically close because of inactivity. A connection
// will remain open indefinitely even if there is no activity.
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
// 
// Question: How do I Choose the TLS version, such as 1.2? Answer: The client does
// not specifically choose the TLS version. In the TLS handshake (which is what
// occurs internally in this method), the client tells the server the version of
// the TLS protocol it wishes to use, which should be the highest version is
// supports. In this case, (at the time of this writing on 22-June-2017) it is TLS
// 1.2. The server then chooses the TLS version that will actually be used. In most
// cases it will be TLS 1.2. The client can then choose to accept or reject the
// connection based on the TLS version chosen by the server. By default, Chilkat
// will reject anything lower than SSL 3.0 (i.e. SSL 2.0 or lower is rejected). The
// SslProtocol property can be set to change what is accepted by Chilkat. For
// example, it can be set to "TLS 1.0 or higher".
//
func (z *Socket) Connect(arg1 string, arg2 int, arg3 bool, arg4 int) bool {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkSocket_Connect(z.ckObj, cstr1, C.int(arg2), b3, C.int(arg4))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Connect method
func (z *Socket) ConnectAsync(arg1 string, arg2 int, arg3 bool, arg4 int, c chan *Task) {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSocket_ConnectAsync(z.ckObj, cstr1, C.int(arg2), b3, C.int(arg4))

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


// Closes the secure (TLS/SSL) channel leaving the socket in a connected state
// where data sent and received is unencrypted.
func (z *Socket) ConvertFromSsl() bool {

    v := C.CkSocket_ConvertFromSsl(z.ckObj)


    return v != 0
}


// Asynchronous version of the ConvertFromSsl method
func (z *Socket) ConvertFromSslAsync(c chan *Task) {

    hTask := C.CkSocket_ConvertFromSslAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Converts a non-SSL/TLS connected socket to a secure channel using TLS/SSL.
func (z *Socket) ConvertToSsl() bool {

    v := C.CkSocket_ConvertToSsl(z.ckObj)


    return v != 0
}


// Asynchronous version of the ConvertToSsl method
func (z *Socket) ConvertToSslAsync(c chan *Task) {

    hTask := C.CkSocket_ConvertToSslAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Clears the Chilkat-wide in-memory hostname-to-IP address DNS cache. Chilkat
// automatically maintains this in-memory cache to prevent redundant DNS lookups.
// If the TTL on the DNS A records being accessed are short and/or these DNS
// records change frequently, then this method can be called clear the internal
// cache. Note: The DNS cache is used/shared among all Chilkat objects in a
// program, and clearing the cache affects all Chilkat objects.
func (z *Socket) DnsCacheClear()  {

    C.CkSocket_DnsCacheClear(z.ckObj)



}


// Performs a DNS query to resolve a hostname to an IP address. The IP address is
// returned if successful. The maximum time to wait (in milliseconds) is passed in
// maxWaitMs. To wait indefinitely, set maxWaitMs = 0.
// return a string or nil if failed.
func (z *Socket) DnsLookup(arg1 string, arg2 int) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSocket_dnsLookup(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the DnsLookup method
func (z *Socket) DnsLookupAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSocket_DnsLookupAsync(z.ckObj, cstr1, C.int(arg2))

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


// Returns the digital certificate to be used for SSL connections. This method
// would only be called by an SSL server application. The SSL certificate is
// initially specified by calling InitSslServer.
//
func (z *Socket) GetMyCert() *Cert {

    retObj := C.CkSocket_GetMyCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the Nth client certificate received during an SSL/TLS handshake. This
// method only applies to the server-side of an SSL/TLS connection. The 1st client
// certificate is at index 0. The NumReceivedClientCerts property indicates the
// number of client certificates received during the SSL/TLS connection
// establishment.
// 
// Client certificates are customarily only sent when the server demands
// client-side authentication, as in 2-way SSL/TLS. This method provides the
// ability for the server to access and examine the client-side certs immediately
// after a connection is established. (Of course, if the client-side certs are
// inadequate for authentication, then the application can choose to immediately
// disconnect.)
// 
// Important: This method should be called from the socket object that is returned
// by AcceptNextConnection.
//
func (z *Socket) GetReceivedClientCert(arg1 int) *Cert {

    retObj := C.CkSocket_GetReceivedClientCert(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// If connected as an SSL/TLS client to an SSL/TLS server where the server requires
// a client-side certificate for authentication, then the NumSslAcceptableClientCAs
// property contains the number of acceptable certificate authorities sent by the
// server during connection establishment handshake. This method may be called to
// get the Distinguished Name (DN) of each acceptable CA. The index should range
// from 0 to NumSslAcceptableClientCAs - 1.
// return a string or nil if failed.
func (z *Socket) GetSslAcceptableClientCaDn(arg1 int) *string {

    retStr := C.CkSocket_getSslAcceptableClientCaDn(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the SSL server's digital certificate. This method would only be called
// by the client-side of an SSL connection. It returns the certificate of the
// remote SSL server for the current SSL connection. If the socket is not
// connected, or is not connected via SSL, then a NULL reference is returned.
func (z *Socket) GetSslServerCert() *Cert {

    retObj := C.CkSocket_GetSslServerCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// SSL Server applications should call this method with the SSL server certificate
// to be used for SSL connections. It should be called prior to accepting
// connections. This method has an intended side-effect: If not already connected,
// then the Ssl property is set to true.
func (z *Socket) InitSslServer(arg1 *Cert) bool {

    v := C.CkSocket_InitSslServer(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns true if the component is unlocked.
func (z *Socket) IsUnlocked() bool {

    v := C.CkSocket_IsUnlocked(z.ckObj)


    return v != 0
}


// Provides information about what transpired in the last method called on this
// object instance. For many methods, there is no information. However, for some
// methods, details about what occurred can be obtained by getting the LastJsonData
// right after the method call returns.
func (z *Socket) LastJsonData() *JsonObject {

    retObj := C.CkSocket_LastJsonData(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Loads the caller of the task's async method.
func (z *Socket) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkSocket_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads the socket object from a completed asynchronous task.
func (z *Socket) LoadTaskResult(arg1 *Task) bool {

    v := C.CkSocket_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


// Check to see if data is available for reading on the socket. Returns true if
// data is waiting and false if no data is waiting to be read.
func (z *Socket) PollDataAvailable() bool {

    v := C.CkSocket_PollDataAvailable(z.ckObj)


    return v != 0
}


// Asynchronous version of the PollDataAvailable method
func (z *Socket) PollDataAvailableAsync(c chan *Task) {

    hTask := C.CkSocket_PollDataAvailableAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives as much data as is immediately available on a connected TCP socket and
// appends the incoming data to binData. If no data is immediately available, it waits
// up to MaxReadIdleMs milliseconds for data to arrive.
func (z *Socket) ReceiveBd(arg1 *BinData) bool {

    v := C.CkSocket_ReceiveBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the ReceiveBd method
func (z *Socket) ReceiveBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkSocket_ReceiveBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads exactly numBytes bytes from the connection. This method blocks until numBytes
// bytes are read or the read times out. The timeout is specified by the
// MaxReadIdleMs property (in milliseconds).
func (z *Socket) ReceiveBdN(arg1 uint, arg2 *BinData) bool {

    v := C.CkSocket_ReceiveBdN(z.ckObj, C.ulong(arg1), arg2.ckObj)


    return v != 0
}


// Asynchronous version of the ReceiveBdN method
func (z *Socket) ReceiveBdNAsync(arg1 uint, arg2 *BinData, c chan *Task) {

    hTask := C.CkSocket_ReceiveBdNAsync(z.ckObj, C.ulong(arg1), arg2.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives a single byte. The received byte will be available in the ReceivedInt
// property. If bUnsigned is true, then a value from 0 to 255 is returned in
// ReceivedInt. If bUnsigned is false, then a value from -128 to +127 is returned.
func (z *Socket) ReceiveByte(arg1 bool) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkSocket_ReceiveByte(z.ckObj, b1)


    return v != 0
}


// Asynchronous version of the ReceiveByte method
func (z *Socket) ReceiveByteAsync(arg1 bool, c chan *Task) {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    hTask := C.CkSocket_ReceiveByteAsync(z.ckObj, b1)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives as much data as is immediately available on a connected TCP socket. If
// no data is immediately available, it waits up to MaxReadIdleMs milliseconds for
// data to arrive.
func (z *Socket) ReceiveBytes() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSocket_ReceiveBytes(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the ReceiveBytes method
func (z *Socket) ReceiveBytesAsync(c chan *Task) {

    hTask := C.CkSocket_ReceiveBytesAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as ReceiveBytes, except the bytes are returned in encoded string form
// according to encodingAlg. The encodingAlg can be "Base64", "modBase64", "Base32", "UU", "QP"
// (for quoted-printable), "URL" (for url-encoding), "Hex", "Q", "B", "url_oath",
// "url_rfc1738", "url_rfc2396", or "url_rfc3986".
// return a string or nil if failed.
func (z *Socket) ReceiveBytesENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSocket_receiveBytesENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReceiveBytesENC method
func (z *Socket) ReceiveBytesENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSocket_ReceiveBytesENCAsync(z.ckObj, cstr1)

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


// Reads exactly numBytes bytes from a connected SSL or non-SSL socket. This method
// blocks until numBytes bytes are read or the read times out. The timeout is specified
// by the MaxReadIdleMs property (in milliseconds).
func (z *Socket) ReceiveBytesN(arg1 uint) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSocket_ReceiveBytesN(z.ckObj, C.ulong(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the ReceiveBytesN method
func (z *Socket) ReceiveBytesNAsync(arg1 uint, c chan *Task) {

    hTask := C.CkSocket_ReceiveBytesNAsync(z.ckObj, C.ulong(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives as much data as is immediately available on a connected TCP socket. If
// no data is immediately available, it waits up to MaxReadIdleMs milliseconds for
// data to arrive.
// 
// The received data is appended to the file specified by appendFilename.
//
func (z *Socket) ReceiveBytesToFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSocket_ReceiveBytesToFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the ReceiveBytesToFile method
func (z *Socket) ReceiveBytesToFileAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSocket_ReceiveBytesToFileAsync(z.ckObj, cstr1)

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


// Receives a 4-byte signed integer and returns the value received. Returns -1 on
// error.
func (z *Socket) ReceiveCount() int {

    v := C.CkSocket_ReceiveCount(z.ckObj)


    return int(v)
}


// Asynchronous version of the ReceiveCount method
func (z *Socket) ReceiveCountAsync(c chan *Task) {

    hTask := C.CkSocket_ReceiveCountAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives a 16-bit integer (2 bytes). The received integer will be available in
// the ReceivedInt property. Set bigEndian equal to true if the incoming 16-bit
// integer is in big-endian byte order. Otherwise set bigEndian equal to false for
// receving a little-endian integer. If bUnsigned is true, the ReceivedInt will range
// from 0 to 65,535. If bUnsigned is false, the ReceivedInt will range from -32,768
// through 32,767.
func (z *Socket) ReceiveInt16(arg1 bool, arg2 bool) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkSocket_ReceiveInt16(z.ckObj, b1, b2)


    return v != 0
}


// Asynchronous version of the ReceiveInt16 method
func (z *Socket) ReceiveInt16Async(arg1 bool, arg2 bool, c chan *Task) {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkSocket_ReceiveInt16Async(z.ckObj, b1, b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives a 32-bit integer (4 bytes). The received integer will be available in
// the ReceivedInt property. Set bigEndian equal to true if the incoming 32-bit
// integer is in big-endian byte order. Otherwise set bigEndian equal to false for
// receving a little-endian integer.
func (z *Socket) ReceiveInt32(arg1 bool) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkSocket_ReceiveInt32(z.ckObj, b1)


    return v != 0
}


// Asynchronous version of the ReceiveInt32 method
func (z *Socket) ReceiveInt32Async(arg1 bool, c chan *Task) {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    hTask := C.CkSocket_ReceiveInt32Async(z.ckObj, b1)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as ReceiveBytesN, except the bytes are returned in encoded string form
// using the encoding specified by numBytes. The numBytes can be "Base64", "modBase64",
// "Base32", "UU", "QP" (for quoted-printable), "URL" (for url-encoding), "Hex",
// "Q", "B", "url_oath", "url_rfc1738", "url_rfc2396", or "url_rfc3986".
// return a string or nil if failed.
func (z *Socket) ReceiveNBytesENC(arg1 uint, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkSocket_receiveNBytesENC(z.ckObj, C.ulong(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReceiveNBytesENC method
func (z *Socket) ReceiveNBytesENCAsync(arg1 uint, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkSocket_ReceiveNBytesENCAsync(z.ckObj, C.ulong(arg1), cstr2)

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


// Receives as much data as is immediately available on the connection. If no data
// is immediately available, it waits up to MaxReadIdleMs milliseconds for data to
// arrive. The incoming bytes are interpreted according to the StringCharset
// property and appended to sb.
func (z *Socket) ReceiveSb(arg1 *StringBuilder) bool {

    v := C.CkSocket_ReceiveSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the ReceiveSb method
func (z *Socket) ReceiveSbAsync(arg1 *StringBuilder, c chan *Task) {

    hTask := C.CkSocket_ReceiveSbAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives as much data as is immediately available on a TCP/IP or SSL socket. If
// no data is immediately available, it waits up to MaxReadIdleMs milliseconds for
// data to arrive. The incoming bytes are interpreted according to the
// StringCharset property and returned as a string.
// return a string or nil if failed.
func (z *Socket) ReceiveString() *string {

    retStr := C.CkSocket_receiveString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReceiveString method
func (z *Socket) ReceiveStringAsync(c chan *Task) {

    hTask := C.CkSocket_ReceiveStringAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Same as ReceiveString, but limits the amount of data returned to a maximum of
// maxByteCount bytes.
// 
// (Receives as much data as is immediately available on the TCP/IP or SSL socket.
// If no data is immediately available, it waits up to MaxReadIdleMs milliseconds
// for data to arrive. The incoming bytes are interpreted according to the
// StringCharset property and returned as a string.)
//
// return a string or nil if failed.
func (z *Socket) ReceiveStringMaxN(arg1 int) *string {

    retStr := C.CkSocket_receiveStringMaxN(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReceiveStringMaxN method
func (z *Socket) ReceiveStringMaxNAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_ReceiveStringMaxNAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives bytes on a connected SSL or non-SSL socket until a specific 1-byte
// value is read. Returns a string containing all the bytes up to but excluding the
// lookForByte.
// return a string or nil if failed.
func (z *Socket) ReceiveStringUntilByte(arg1 int) *string {

    retStr := C.CkSocket_receiveStringUntilByte(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReceiveStringUntilByte method
func (z *Socket) ReceiveStringUntilByteAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_ReceiveStringUntilByteAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads text from the connected TCP/IP or SSL socket until a CRLF is received.
// Returns the text up to and including the CRLF. The incoming bytes are
// interpreted according to the charset specified by the StringCharset property.
// return a string or nil if failed.
func (z *Socket) ReceiveToCRLF() *string {

    retStr := C.CkSocket_receiveToCRLF(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReceiveToCRLF method
func (z *Socket) ReceiveToCRLFAsync(c chan *Task) {

    hTask := C.CkSocket_ReceiveToCRLFAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives bytes on the TCP/IP or SSL socket until a specific 1-byte value is
// read. Returns all the bytes up to and including the lookForByte.
func (z *Socket) ReceiveUntilByte(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSocket_ReceiveUntilByte(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Asynchronous version of the ReceiveUntilByte method
func (z *Socket) ReceiveUntilByteAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_ReceiveUntilByteAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Receives bytes on the TCP/IP or SSL socket until a specific 1-byte value is
// read. Returns all the bytes up to and including the lookForByte. The received bytes are
// appended to bd.
func (z *Socket) ReceiveUntilByteBd(arg1 int, arg2 *BinData) bool {

    v := C.CkSocket_ReceiveUntilByteBd(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Asynchronous version of the ReceiveUntilByteBd method
func (z *Socket) ReceiveUntilByteBdAsync(arg1 int, arg2 *BinData, c chan *Task) {

    hTask := C.CkSocket_ReceiveUntilByteBdAsync(z.ckObj, C.int(arg1), arg2.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reads text from the connected TCP/IP or SSL socket until a matching string
// (matchStr) is received. Returns the text up to and including the matching string. As
// an example, to one might read the header of an HTTP request or a MIME message by
// reading up to the first double CRLF ("\r\n\r\n"). The incoming bytes are
// interpreted according to the charset specified by the StringCharset property.
// return a string or nil if failed.
func (z *Socket) ReceiveUntilMatch(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSocket_receiveUntilMatch(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReceiveUntilMatch method
func (z *Socket) ReceiveUntilMatchAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSocket_ReceiveUntilMatchAsync(z.ckObj, cstr1)

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


// Resets the performance measurements for either receiving or sending. If rcvPerf is
// true, then the receive performance monitoring is reset. If rcvPerf is false,
// then the sending performance monitoring is reset.
func (z *Socket) ResetPerf(arg1 bool)  {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    C.CkSocket_ResetPerf(z.ckObj, b1)



}


// Wait for data to arrive on this socket, or any of the contained sockets if the
// caller is a "socket set". (If the socket is a listener socket, then waits for an
// incoming connect. Listener sockets can be added to the "socket set" just like
// connected sockets.)
// 
// (see the example at the link below for more detailed information)
// 
// Waits a maximum of timeoutMs milliseconds. If maxWaitMs = 0, then it is effectively a
// poll. Returns the number of sockets with data available for reading. If no
// sockets have data available for reading, then a value of 0 is returned. A value
// of -1 indicates an error condition. Note: when the remote peer (in this case the
// web server) disconnects, the socket will appear as if it has data available. A
// "ready" socket is one where either data is available for reading or the socket
// has become disconnected.
// 
// If the peer closed the connection, it will not be discovered until an attempt is
// made to read the socket. If the read fails, then the IsConnected property may be
// checked to see if the connection was closed.
//
func (z *Socket) SelectForReading(arg1 int) int {

    v := C.CkSocket_SelectForReading(z.ckObj, C.int(arg1))


    return int(v)
}


// Asynchronous version of the SelectForReading method
func (z *Socket) SelectForReadingAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_SelectForReadingAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Waits until it is known that data can be written to one or more sockets without
// it blocking.
// 
// Socket writes are typically buffered by the operating system. When an
// application writes data to a socket, the operating system appends it to the
// socket's outgoing send buffers and returns immediately. However, if the OS send
// buffers become filled up (because the sender is sending data faster than the
// remote receiver can read it), then a socket write can block (until outgoing send
// buffer space becomes available).
// 
// Waits a maximum of timeoutMs milliseconds. If maxWaitMs = 0, then it is effectively a
// poll. Returns the number of sockets such that data can be written without
// blocking. A value of -1 indicates an error condition.
//
func (z *Socket) SelectForWriting(arg1 int) int {

    v := C.CkSocket_SelectForWriting(z.ckObj, C.int(arg1))


    return int(v)
}


// Asynchronous version of the SelectForWriting method
func (z *Socket) SelectForWritingAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_SelectForWritingAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends bytes from binData over a connected SSL or non-SSL socket. If transmission
// halts for more than MaxSendIdleMs milliseconds, the send is aborted. This is a
// blocking (synchronous) method. It returns only after the bytes have been sent.
// 
// Set offset and/or numBytes to non-zero values to send a portion of the binData. If offset
// and numBytes are both 0, then the entire binData is sent. If offset is non-zero and numBytes
// is zero, then the bytes starting at offset until the end are sent.
//
func (z *Socket) SendBd(arg1 *BinData, arg2 uint, arg3 uint) bool {

    v := C.CkSocket_SendBd(z.ckObj, arg1.ckObj, C.ulong(arg2), C.ulong(arg3))


    return v != 0
}


// Asynchronous version of the SendBd method
func (z *Socket) SendBdAsync(arg1 *BinData, arg2 uint, arg3 uint, c chan *Task) {

    hTask := C.CkSocket_SendBdAsync(z.ckObj, arg1.ckObj, C.ulong(arg2), C.ulong(arg3))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a single byte. The integer must have a value from 0 to 255.
func (z *Socket) SendByte(arg1 int) bool {

    v := C.CkSocket_SendByte(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the SendByte method
func (z *Socket) SendByteAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_SendByteAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends bytes over a connected SSL or non-SSL socket. If transmission halts for
// more than MaxSendIdleMs milliseconds, the send is aborted. This is a blocking
// (synchronous) method. It returns only after the bytes have been sent.
func (z *Socket) SendBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkSocket_SendBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Asynchronous version of the SendBytes method
func (z *Socket) SendBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkSocket_SendBytesAsync(z.ckObj, ckbyd1)

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


// The same as SendBytes, except the bytes are provided in encoded string form as
// specified by encodingAlg. The encodingAlg can be "Base64", "modBase64", "Base32", "Base58",
// "UU", "QP" (for quoted-printable), "URL" (for url-encoding), "Hex", "Q", "B",
// "url_oauth", "url_rfc1738", "url_rfc2396", and "url_rfc3986".
func (z *Socket) SendBytesENC(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSocket_SendBytesENC(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SendBytesENC method
func (z *Socket) SendBytesENCAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSocket_SendBytesENCAsync(z.ckObj, cstr1, cstr2)

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


// Sends a 4-byte signed integer on the connection. The receiver may call
// ReceiveCount to receive the integer. The SendCount and ReceiveCount methods are
// handy for sending byte counts prior to sending data. The sender would send a
// count followed by the data, and the receiver would receive the count first, and
// then knows how many data bytes it should expect to receive.
func (z *Socket) SendCount(arg1 int) bool {

    v := C.CkSocket_SendCount(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the SendCount method
func (z *Socket) SendCountAsync(arg1 int, c chan *Task) {

    hTask := C.CkSocket_SendCountAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a 16-bit integer (2 bytes). Set bigEndian equal to true to send the integer
// in big-endian byte order (this is the standard network byte order). Otherwise
// set bigEndian equal to false to send in little-endian byte order.
func (z *Socket) SendInt16(arg1 int, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkSocket_SendInt16(z.ckObj, C.int(arg1), b2)


    return v != 0
}


// Asynchronous version of the SendInt16 method
func (z *Socket) SendInt16Async(arg1 int, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkSocket_SendInt16Async(z.ckObj, C.int(arg1), b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a 32-bit integer (4 bytes). Set bigEndian equal to true to send the integer
// in big-endian byte order (this is the standard network byte order). Otherwise
// set bigEndian equal to false to send in little-endian byte order.
func (z *Socket) SendInt32(arg1 int, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkSocket_SendInt32(z.ckObj, C.int(arg1), b2)


    return v != 0
}


// Asynchronous version of the SendInt32 method
func (z *Socket) SendInt32Async(arg1 int, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkSocket_SendInt32Async(z.ckObj, C.int(arg1), b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends the contents of sb over the connection. If transmission halts for more
// than MaxSendIdleMs milliseconds, the send is aborted. The string is sent in the
// charset encoding specified by the StringCharset property.
// 
// This is a blocking (synchronous) method. It returns after the string has been
// sent.
//
func (z *Socket) SendSb(arg1 *StringBuilder) bool {

    v := C.CkSocket_SendSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the SendSb method
func (z *Socket) SendSbAsync(arg1 *StringBuilder, c chan *Task) {

    hTask := C.CkSocket_SendSbAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a string over a connected SSL or non-SSL (TCP/IP) socket. If transmission
// halts for more than MaxSendIdleMs milliseconds, the send is aborted. The string
// is sent in the charset encoding specified by the StringCharset property.
// 
// This is a blocking (synchronous) method. It returns after the string has been
// sent.
//
func (z *Socket) SendString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSocket_SendString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SendString method
func (z *Socket) SendStringAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSocket_SendStringAsync(z.ckObj, cstr1)

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


// Sends a "Wake on Lan" magic packet to a computer. A Wake on Lan is a way to
// power on a computer remotely by sending a data packet known as a magic packet.
// For this to work, the network card must have enabled the feature: “Power on Lan”
// or “Power on PCI Device“, which is done by accessing the BIOS of the machine.
// 
// The macAddress is the MAC address (in hex) of the computer to wake. A MAC address
// should be 6 bytes in length. For example, "000102030405". The port is the port
// which should be 7 or 9. (Port number 9 is more commonly used.) The ipBroadcastAddr is the
// broadcast address of your network, which usually ends with *.255. For example:
// "192.168.1.255".
// 
// Your application does not call Connect prior to calling SendWakeOnLan. To use
// this method, it's just a matter of instantiating an instance of this socket
// object and then call SendWakeOnLan.
//
func (z *Socket) SendWakeOnLan(arg1 string, arg2 int, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    v := C.CkSocket_SendWakeOnLan(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// The same as SendWakeOnLan, but includes an additional argument to specify a
// SecureOn password. The password should be a hexidecimal string representing 4 or 6
// bytes. (See https://wiki.wireshark.org/WakeOnLAN) Sending a WakeOnLAN (WOL) to
// an IPv4 address would need a 4-byte SecureOn password, whereas an IPv6 address
// would need a 6-byte SecureOn password.
func (z *Socket) SendWakeOnLan2(arg1 string, arg2 int, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkSocket_SendWakeOnLan2(z.ckObj, cstr1, C.int(arg2), cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// A client-side certificate for SSL/TLS connections is optional. It should be used
// only if the server demands it. This method allows the certificate to be
// specified using a certificate object.
func (z *Socket) SetSslClientCert(arg1 *Cert) bool {

    v := C.CkSocket_SetSslClientCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// A client-side certificate for SSL/TLS connections is optional. It should be used
// only if the server demands it. This method allows the certificate to be
// specified using a PEM file.
func (z *Socket) SetSslClientCertPem(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSocket_SetSslClientCertPem(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// A client-side certificate for SSL/TLS connections is optional. It should be used
// only if the server demands it. This method allows the certificate to be
// specified using a PFX file.
func (z *Socket) SetSslClientCertPfx(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSocket_SetSslClientCertPfx(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Convenience method to force the calling thread to sleep for a number of
// milliseconds.
func (z *Socket) SleepMs(arg1 int)  {

    C.CkSocket_SleepMs(z.ckObj, C.int(arg1))



}


// Authenticates with the SSH server using public-key authentication. The
// corresponding public key must have been installed on the SSH server for the
// sshLogin. Authentication will succeed if the matching privateKey is provided.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *Socket) SshAuthenticatePk(arg1 string, arg2 *SshKey) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSocket_SshAuthenticatePk(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SshAuthenticatePk method
func (z *Socket) SshAuthenticatePkAsync(arg1 string, arg2 *SshKey, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSocket_SshAuthenticatePkAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Authenticates with the SSH server using a sshLogin and sshPassword. This method is only
// used for SSH tunneling. The tunnel is established by calling SshOpenTunnel, then
// (if necessary) authenticated by calling SshAuthenticatePw or SshAuthenticatePk.
func (z *Socket) SshAuthenticatePw(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSocket_SshAuthenticatePw(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SshAuthenticatePw method
func (z *Socket) SshAuthenticatePwAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSocket_SshAuthenticatePwAsync(z.ckObj, cstr1, cstr2)

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
func (z *Socket) SshCloseTunnel() bool {

    v := C.CkSocket_SshCloseTunnel(z.ckObj)


    return v != 0
}


// Asynchronous version of the SshCloseTunnel method
func (z *Socket) SshCloseTunnelAsync(c chan *Task) {

    hTask := C.CkSocket_SshCloseTunnelAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Opens a new channel within an SSH tunnel. Returns the socket that is connected
// to the destination host:port through the SSH tunnel via port forwarding. If ssl
// is true, the connection is TLS (i.e. TLS inside the SSH tunnel). Returns the
// socket object that is the port-forwarded tunneled connection. Any number of
// channels may be opened within a single SSH tunnel, and may be port-forwarded to
// different remote host:port endpoints.
func (z *Socket) SshOpenChannel(arg1 string, arg2 int, arg3 bool, arg4 int) *Socket {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retObj := C.CkSocket_SshOpenChannel(z.ckObj, cstr1, C.int(arg2), b3, C.int(arg4))

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Socket{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Asynchronous version of the SshOpenChannel method
func (z *Socket) SshOpenChannelAsync(arg1 string, arg2 int, arg3 bool, arg4 int, c chan *Task) {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSocket_SshOpenChannelAsync(z.ckObj, cstr1, C.int(arg2), b3, C.int(arg4))

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


// Connects to an SSH server and creates a tunnel for port forwarding. The sshHostname is
// the hostname (or IP address) of the SSH server. The sshPort is typically 22, which
// is the standard SSH port number.
// 
// An SSH tunneling (port forwarding) session always begins by first calling
// SshOpenTunnel to connect to the SSH server, followed by calling either
// SshAuthenticatePw or SshAuthenticatePk to authenticate. A program would then
// call SshOpenChannel to connect to the destination server (via the SSH tunnel).
// Any number of channels can be opened over the same SSH tunnel.
//
func (z *Socket) SshOpenTunnel(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSocket_SshOpenTunnel(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SshOpenTunnel method
func (z *Socket) SshOpenTunnelAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSocket_SshOpenTunnelAsync(z.ckObj, cstr1, C.int(arg2))

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


// Used in combination with the ElapsedSeconds property, which will contain the
// number of seconds since the last call to this method. (The StartTiming method
// and ElapsedSeconds property is provided for convenience.)
func (z *Socket) StartTiming()  {

    C.CkSocket_StartTiming(z.ckObj)



}


// Takes the connection from sock. If the caller of this method had an open
// connection, then it will be closed. This method is different than the TakeSocket
// method because the caller does not become a "socket set".
func (z *Socket) TakeConnection(arg1 *Socket) bool {

    v := C.CkSocket_TakeConnection(z.ckObj, arg1.ckObj)


    return v != 0
}


// Takes ownership of the sock. sock is added to the internal set of connected
// sockets. The caller object is now effectively a "socket set", i.e. a collection
// of connected and/or listener sockets. Method calls are routed to the internal
// sockets based on the value of the SelectorIndex property. For example, if
// SelectorIndex equals 2, then a call to SendBytes is actually a call to SendBytes
// on the 3rd socket in the set. (Indexing begins at 0.) Likewise, getting and
// setting properties are also routed to the contained socket based on
// SelectorIndex. It is possible to wait on a set of sockets for data to arrive on
// any of them by calling SelectForReading. See the example link below.
func (z *Socket) TakeSocket(arg1 *Socket) bool {

    v := C.CkSocket_TakeSocket(z.ckObj, arg1.ckObj)


    return v != 0
}


// Initiates a renegotiation of the TLS security parameters. This sends a
// ClientHello to re-do the TLS handshake to establish new TLS security params.
func (z *Socket) TlsRenegotiate() bool {

    v := C.CkSocket_TlsRenegotiate(z.ckObj)


    return v != 0
}


// Asynchronous version of the TlsRenegotiate method
func (z *Socket) TlsRenegotiateAsync(c chan *Task) {

    hTask := C.CkSocket_TlsRenegotiateAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Unlocks the component allowing for the full functionality to be used. An
// arbitrary string can be passed to initiate a fully-functional 30-day trial.
func (z *Socket) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSocket_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Uses an existing SSH tunnel for the connection. This is an alternative way of
// establishing a socket connection through an SSH tunnel. There are four ways of
// running a TCP or SSL/TLS connection through an SSH tunnel:
//     UseSsh
//         Establish the SSH connection and authenticate using the Chilkat SSH
//         object.
//         Call UseSsh to indicate that the connections should be made through the
//         SSH tunnel.
//         Call the Connect method to establish the TCP or SSL/TLS connection with
//         a destination host:port. The connection is not direct, but will instead be
//         routed through the SSH tunnel and then port-forwarded (from the SSH server) to
//         the destination host:port. (Had UseSsh not been called, the connection would be
//         direct.)
//     SshOpenTunnel
//         Call the Socket object's SshOpenTunnel method to connect to an SSH
//         server.
//         Call SshAuthenticatePw to authenticate with the SSH server.
//         Instead of calling Connect to connect with the destination host:port,
//         the SshOpenChannel method is called to connect via port-forwarding through the
//         SSH tunnel.
//     SshTunnel object with dynamic port forwarding
//         The Chilkat SSH Tunnel object is utilized to run in a background thread.
//         It connects and authenticates with an SSH server, and then listens at a port
//         chosen by the application, and behaves as a SOCKS5 proxy server.
//         The Socket object sets the SOCKS5 proxy host:port to
//         localhost:_LT_port_GT_,
//         The Socket's Connect method is called to connect via the SSH Tunnel. The
//         connection is routed through the SSH tunnel via dynamic port forwarding.
//         Once the background SSH Tunnel thread is running, it can handle any
//         number of incoming connections from the foreground thread, other threads, or
//         even other programs that are local or remote. Each incoming connection is routed
//         via dynamic port forwarding to it's chosen destnation host:port on it's own
//         logical SSH channel.
//     SshTunnel object with hard-coded port forwarding
//         The Chilkat SSH Tunnel object is utilized to run in a background thread.
//         It connects and authenticates with an SSH server, and then listens at a port
//         chosen by the application. It does not behave as a SOCKS5 proxy server, but
//         instead has a hard-coded destination host:port.
//         The Socket's Connect method is called to connect to
//         localhost:_LT_port_GT_. The connection is automatically port-forwarded through
//         the SSH tunnel to the hard-coded destination host:port.
// In all cases, the SSH tunnels can hold both unencrypted TCP connections and
// SSL/TLS connections.
func (z *Socket) UseSsh(arg1 *Ssh) bool {

    v := C.CkSocket_UseSsh(z.ckObj, arg1.ckObj)


    return v != 0
}


