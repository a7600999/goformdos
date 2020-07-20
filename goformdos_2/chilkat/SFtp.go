// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkSFtp.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewSFtp() *SFtp {
	return &SFtp{C.CkSFtp_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *SFtp) DisposeSFtp() {
    if z != nil {
	C.CkSFtp_Dispose(z.ckObj)
	}
}




func (z *SFtp) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkSFtp_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *SFtp) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkSFtp_setExternalProgress(z.ckObj,1)
}

func (z *SFtp) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkSFtp_setExternalProgress(z.ckObj,1)
}

func (z *SFtp) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkSFtp_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *SFtp) AbortCurrent() bool {
    v := int(C.CkSFtp_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *SFtp) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putAbortCurrent(z.ckObj,v)
}

// property: AccumulateBuffer
// Contains the bytes downloaded from a remote file via the AccumulateBytes method
// call. Each call to AccumulateBytes appends to this buffer. To clear this buffer,
// call the ClearAccumulateBuffer method.
func (z *SFtp) AccumulateBuffer() []byte {
    ckbyd := C.CkByteData_Create()
    C.CkSFtp_getAccumulateBuffer(z.ckObj,ckbyd)
    p := C.CkByteData_getData(ckbyd);
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckbyd)) )
    C.CkByteData_Dispose(ckbyd)
    return retBytes
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
func (z *SFtp) AuthFailReason() int {
    return int(C.CkSFtp_getAuthFailReason(z.ckObj))
}

// property: BandwidthThrottleDown
// If non-zero, limits (throttles) the download bandwidth to approximately this
// maximum number of bytes per second. The default value of this property is 0.
func (z *SFtp) BandwidthThrottleDown() int {
    return int(C.CkSFtp_getBandwidthThrottleDown(z.ckObj))
}
// property setter: BandwidthThrottleDown
func (z *SFtp) SetBandwidthThrottleDown(value int) {
    C.CkSFtp_putBandwidthThrottleDown(z.ckObj,C.int(value))
}

// property: BandwidthThrottleUp
// If non-zero, limits (throttles) the upload bandwidth to approximately this
// maximum number of bytes per second. The default value of this property is 0.
func (z *SFtp) BandwidthThrottleUp() int {
    return int(C.CkSFtp_getBandwidthThrottleUp(z.ckObj))
}
// property setter: BandwidthThrottleUp
func (z *SFtp) SetBandwidthThrottleUp(value int) {
    C.CkSFtp_putBandwidthThrottleUp(z.ckObj,C.int(value))
}

// property: ClientIdentifier
// The client-identifier string to be used when connecting to an SSH/SFTP server.
// Defaults to "SSH-2.0-PuTTY_Release_0.63". (This string is used to mimic PuTTY
// because some servers are known to disconnect from clients with unknown client
// identifiers.)
func (z *SFtp) ClientIdentifier() string {
    return C.GoString(C.CkSFtp_clientIdentifier(z.ckObj))
}
// property setter: ClientIdentifier
func (z *SFtp) SetClientIdentifier(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putClientIdentifier(z.ckObj,cStr)
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
func (z *SFtp) ClientIpAddress() string {
    return C.GoString(C.CkSFtp_clientIpAddress(z.ckObj))
}
// property setter: ClientIpAddress
func (z *SFtp) SetClientIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putClientIpAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConnectTimeoutMs
// Maximum number of milliseconds to wait when connecting to an SSH server.
// 
// To clarify: This property determines how long to wait for the SSH server to
// accept the TCP connection. Once the connection is made, it is the IdleTimeoutMs
// property that applies to receiving data and responses.
//
func (z *SFtp) ConnectTimeoutMs() int {
    return int(C.CkSFtp_getConnectTimeoutMs(z.ckObj))
}
// property setter: ConnectTimeoutMs
func (z *SFtp) SetConnectTimeoutMs(value int) {
    C.CkSFtp_putConnectTimeoutMs(z.ckObj,C.int(value))
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
func (z *SFtp) DebugLogFilePath() string {
    return C.GoString(C.CkSFtp_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *SFtp) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DisconnectCode
// If the SSH/SFTP server sent a DISCONNECT message when closing the connection,
// this property contains the "reason code" as specified in RFC 4253:
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
func (z *SFtp) DisconnectCode() int {
    return int(C.CkSFtp_getDisconnectCode(z.ckObj))
}

// property: DisconnectReason
// If the SSH/SFTP server sent a DISCONNECT message when closing the connection,
// this property contains a descriptive string for the "reason code" as specified
// in RFC 4253.
func (z *SFtp) DisconnectReason() string {
    return C.GoString(C.CkSFtp_disconnectReason(z.ckObj))
}

// property: EnableCache
// Controls whether the component keeps an internal file size & attribute cache.
// The cache affects the following methods: GetFileSize32, GetFileSize64,
// GetFileSizeStr, GetFileCreateTime, GetFileLastAccess, GetFileLastModified,
// GetFileOwner, GetFileGroup, and GetFilePermissions.
// 
// The file attribute cache exists to minimize communications with the SFTP server.
// If the cache is enabled, then a request for any single attribute will cause all
// of the attributes to be cached. A subsequent request for a different attribute
// of the same file will be fulfilled from cache, eliminating the need for a
// message to be sent to the SFTP server.
// 
// Note: Caching only occurs when filenames are used, not handles.
//
func (z *SFtp) EnableCache() bool {
    v := int(C.CkSFtp_getEnableCache(z.ckObj))
    return v != 0
}
// property setter: EnableCache
func (z *SFtp) SetEnableCache(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putEnableCache(z.ckObj,v)
}

// property: EnableCompression
// Enables or disables the use of compression w/ the SSH connection. The default
// value is false.
// 
// Some older SSH servers have been found that claim to support compression, but
// actually fail when compression is used. PuTTY does not enable compression by
// default. If trouble is encountered where the SSH server disconnects immediately
// after the connection is seemingly established (i.e. during authentication), then
// check to see if disabling compression resolves the problem.
//
func (z *SFtp) EnableCompression() bool {
    v := int(C.CkSFtp_getEnableCompression(z.ckObj))
    return v != 0
}
// property setter: EnableCompression
func (z *SFtp) SetEnableCompression(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putEnableCompression(z.ckObj,v)
}

// property: FilenameCharset
// Automatically set during the InitializeSftp method if the server sends a
// filename-charset name-value extension. Otherwise, may be set directly to the
// name of a charset, such as "utf-8", "iso-8859-1", "windows-1252", etc. ("ansi"
// is also a valid choice.) Incoming filenames are interpreted as utf-8 if no
// FilenameCharset is set. Outgoing filenames are sent using utf-8 by default.
// Otherwise, incoming and outgoing filenames use the charset specified by this
// property.
func (z *SFtp) FilenameCharset() string {
    return C.GoString(C.CkSFtp_filenameCharset(z.ckObj))
}
// property setter: FilenameCharset
func (z *SFtp) SetFilenameCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putFilenameCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
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
func (z *SFtp) ForceCipher() string {
    return C.GoString(C.CkSFtp_forceCipher(z.ckObj))
}
// property setter: ForceCipher
func (z *SFtp) SetForceCipher(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putForceCipher(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ForceV3
// If set to true, forces the client to choose version 3 of the SFTP protocol,
// even if the server supports a higher version. The default value of this property
// is true.
func (z *SFtp) ForceV3() bool {
    v := int(C.CkSFtp_getForceV3(z.ckObj))
    return v != 0
}
// property setter: ForceV3
func (z *SFtp) SetForceV3(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putForceV3(z.ckObj,v)
}

// property: HeartbeatMs
// This is the number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any SFTP operation prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
func (z *SFtp) HeartbeatMs() int {
    return int(C.CkSFtp_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *SFtp) SetHeartbeatMs(value int) {
    C.CkSFtp_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: HostKeyAlg
// Indicates the preferred host key algorithm to be used in establishing the SSH
// secure connection. The default is "DSS". It may be changed to "RSA" if needed.
// Chilkat recommends not changing this unless a problem warrants the change.
func (z *SFtp) HostKeyAlg() string {
    return C.GoString(C.CkSFtp_hostKeyAlg(z.ckObj))
}
// property setter: HostKeyAlg
func (z *SFtp) SetHostKeyAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putHostKeyAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HostKeyFingerprint
// Set after connecting to an SSH/SFTP server. The format of the fingerprint looks
// like this: "ssh-rsa 1024 68:ff:d1:4e:6c:ff:d7:b0:d6:58:73:85:07:bc:2e:d5"
func (z *SFtp) HostKeyFingerprint() string {
    return C.GoString(C.CkSFtp_hostKeyFingerprint(z.ckObj))
}

// property: HttpProxyAuthMethod
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy authentication method name. Valid choices are "Basic" or "NTLM".
func (z *SFtp) HttpProxyAuthMethod() string {
    return C.GoString(C.CkSFtp_httpProxyAuthMethod(z.ckObj))
}
// property setter: HttpProxyAuthMethod
func (z *SFtp) SetHttpProxyAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putHttpProxyAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyDomain
// The NTLM authentication domain (optional) if NTLM authentication is used w/ the
// HTTP proxy.
func (z *SFtp) HttpProxyDomain() string {
    return C.GoString(C.CkSFtp_httpProxyDomain(z.ckObj))
}
// property setter: HttpProxyDomain
func (z *SFtp) SetHttpProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putHttpProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyHostname
// If an HTTP proxy is to be used, set this property to the HTTP proxy hostname or
// IPv4 address (in dotted decimal notation).
func (z *SFtp) HttpProxyHostname() string {
    return C.GoString(C.CkSFtp_httpProxyHostname(z.ckObj))
}
// property setter: HttpProxyHostname
func (z *SFtp) SetHttpProxyHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putHttpProxyHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPassword
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy password.
func (z *SFtp) HttpProxyPassword() string {
    return C.GoString(C.CkSFtp_httpProxyPassword(z.ckObj))
}
// property setter: HttpProxyPassword
func (z *SFtp) SetHttpProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putHttpProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPort
// If an HTTP proxy is to be used, set this property to the HTTP proxy port number.
// (Two commonly used HTTP proxy ports are 8080 and 3128.)
func (z *SFtp) HttpProxyPort() int {
    return int(C.CkSFtp_getHttpProxyPort(z.ckObj))
}
// property setter: HttpProxyPort
func (z *SFtp) SetHttpProxyPort(value int) {
    C.CkSFtp_putHttpProxyPort(z.ckObj,C.int(value))
}

// property: HttpProxyUsername
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy login name.
func (z *SFtp) HttpProxyUsername() string {
    return C.GoString(C.CkSFtp_httpProxyUsername(z.ckObj))
}
// property setter: HttpProxyUsername
func (z *SFtp) SetHttpProxyUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putHttpProxyUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IdleTimeoutMs
// Causes SFTP operations to fail when progress for sending or receiving data halts
// for more than this number of milliseconds. Setting IdleTimeoutMs = 0 allows the
// application to wait indefinitely. The default value of this property is 30000
// (which equals 30 seconds).
func (z *SFtp) IdleTimeoutMs() int {
    return int(C.CkSFtp_getIdleTimeoutMs(z.ckObj))
}
// property setter: IdleTimeoutMs
func (z *SFtp) SetIdleTimeoutMs(value int) {
    C.CkSFtp_putIdleTimeoutMs(z.ckObj,C.int(value))
}

// property: IncludeDotDirs
// If true, then the ReadDir method will include the "." and ".." directories in
// its results. The default value of this property is false.
func (z *SFtp) IncludeDotDirs() bool {
    v := int(C.CkSFtp_getIncludeDotDirs(z.ckObj))
    return v != 0
}
// property setter: IncludeDotDirs
func (z *SFtp) SetIncludeDotDirs(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putIncludeDotDirs(z.ckObj,v)
}

// property: InitializeFailCode
// The InitializeSftp method call opens a channel for the SFTP session. If the
// request to open a channel fails, this property contains a code that identifies
// the reason for failure. The reason codes are defined in RFC 4254 and are
// reproduced here:
//              Symbolic name                           reason code
//              -------------                           -----------
//             SSH_OPEN_ADMINISTRATIVELY_PROHIBITED          1
//             SSH_OPEN_CONNECT_FAILED                       2
//             SSH_OPEN_UNKNOWN_CHANNEL_TYPE                 3
//             SSH_OPEN_RESOURCE_SHORTAGE                    4
func (z *SFtp) InitializeFailCode() int {
    return int(C.CkSFtp_getInitializeFailCode(z.ckObj))
}

// property: InitializeFailReason
// The InitializeSftp method call opens a channel for the SFTP session. If the
// request to open a channel fails, this property contains a description of the
// reason for failure. (It contains descriptive text matching the
// InitializeFailCode property.)
func (z *SFtp) InitializeFailReason() string {
    return C.GoString(C.CkSFtp_initializeFailReason(z.ckObj))
}

// property: IsConnected
// Returns true if connected to the SSH server. Note: This does not mean
// authentication has happened or InitializeSftp has already succeeded. It only
// means that the connection has been established by calling Connect.
// 
// Understanding the IsConnected property: The IsConnected property is the last
// known state of the TCP connection (either connected or disconnected). This
// requires some explanation because most developer have incorrect assumptions
// about TCP connections.
//     If a TCP connection is established, and neither side is reading or writing
//     the socket (i.e. both sides are doing nothing), then you can disconnect the
//     network cable from the computer for any length of time, and then re-connect, and
//     the TCP connection is not affected.
//     A TCP connection only becomes disconnected when an attempt is made to
//     read/write while a network problem exists. If no attempts to read/write occur, a
//     network problem may arise and then become resolved without affecting the TCP
//     connection.
//     If the peer chooses to close its side of the TCP connection, your
//     application won't magically know about it until you try to do something with the
//     TCP socket (such as read or write).
//     A Chilkat API property (as opposed to a method) CANNOT and should not do
//     something that would cause an application to timeout, hang, etc. Therefore, it
//     is not appropriate for the IsConnected property to attempt any kind of socket
//     operation (read/write/peek) on the socket. It simply returns the last known
//     state of the connection. It may very well be that your network cable is
//     unplugged and IsConnected returns true because technically, if neither peer is
//     trying to read/write, the network cable could be plugged back in without
//     affecting the connection. IsConnected could also return true if the peer has
//     already closed its side of the connection, because the state of the connection
//     is only updated after trying to read/write/peek.
//     To truly know the current connected state (as opposed to the last known
//     connection state), your application should attempt a network operation that is
//     appropriate to the protocol. For SFTP, an application could call SendIgnore, and
//     then check IsConnected.
//
func (z *SFtp) IsConnected() bool {
    v := int(C.CkSFtp_getIsConnected(z.ckObj))
    return v != 0
}

// property: KeepSessionLog
// Controls whether communications to/from the SFTP server are saved to the
// SessionLog property. The default value is false. If this property is set to
// true, the contents of the SessionLog property will continuously grow as SFTP
// activity transpires. The purpose of the KeepSessionLog / SessionLog properties
// is to help in debugging any future problems that may arise.
func (z *SFtp) KeepSessionLog() bool {
    v := int(C.CkSFtp_getKeepSessionLog(z.ckObj))
    return v != 0
}
// property setter: KeepSessionLog
func (z *SFtp) SetKeepSessionLog(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putKeepSessionLog(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *SFtp) LastErrorHtml() string {
    return C.GoString(C.CkSFtp_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *SFtp) LastErrorText() string {
    return C.GoString(C.CkSFtp_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *SFtp) LastErrorXml() string {
    return C.GoString(C.CkSFtp_lastErrorXml(z.ckObj))
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
func (z *SFtp) LastMethodSuccess() bool {
    v := int(C.CkSFtp_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *SFtp) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putLastMethodSuccess(z.ckObj,v)
}

// property: MaxPacketSize
// The maximum packet length to be used in the underlying SSH transport protocol.
// The default value is 32768. (This should generally be left unchanged.)
func (z *SFtp) MaxPacketSize() int {
    return int(C.CkSFtp_getMaxPacketSize(z.ckObj))
}
// property setter: MaxPacketSize
func (z *SFtp) SetMaxPacketSize(value int) {
    C.CkSFtp_putMaxPacketSize(z.ckObj,C.int(value))
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
func (z *SFtp) PasswordChangeRequested() bool {
    v := int(C.CkSFtp_getPasswordChangeRequested(z.ckObj))
    return v != 0
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
func (z *SFtp) PercentDoneScale() int {
    return int(C.CkSFtp_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *SFtp) SetPercentDoneScale(value int) {
    C.CkSFtp_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *SFtp) PreferIpv6() bool {
    v := int(C.CkSFtp_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *SFtp) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putPreferIpv6(z.ckObj,v)
}

// property: PreserveDate
// If true, then the file last-modified and create date/time will be preserved
// for downloaded and uploaded files. The default value is false.
// 
// Note: Prior to version 9.5.0.40, this property only applied to downloads.
// Starting in v9.5.0.40, it also applies to the UploadFileByName method.
//     It does not apply to uploads or downloads where the remote file is opened to
//     obtain a handle, the data is uploaded/downloaded, and then the handle is closed.
//     The last-mod date/times are always preserved ini the SyncTreeDownload and
//     SyncTreeUpload methods.
//
func (z *SFtp) PreserveDate() bool {
    v := int(C.CkSFtp_getPreserveDate(z.ckObj))
    return v != 0
}
// property setter: PreserveDate
func (z *SFtp) SetPreserveDate(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putPreserveDate(z.ckObj,v)
}

// property: ProtocolVersion
// The negotiated SFTP protocol version, which should be a value between 3 and 6
// inclusive. When the InitializeSftp method is called, the Chilkat SFTP client
// sends it's highest supported protocol version to the server, and the server
// sends it's SFTP protocol version in response. The negotiated protocol (i.e. the
// protocol version used for the session) is the lower of the two numbers. If the
// SFTP server's protocol version is lower than 6, some file date/attributes are
// not supported because they are not supported by the earlier protocol version.
// These exceptions are noted throughout the reference documentation.
func (z *SFtp) ProtocolVersion() int {
    return int(C.CkSFtp_getProtocolVersion(z.ckObj))
}

// property: ReadDirMustMatch
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "*.xml; *.txt; *.csv". If set, the ReadDir method will only return
// files that match any one of these patterns.
func (z *SFtp) ReadDirMustMatch() string {
    return C.GoString(C.CkSFtp_readDirMustMatch(z.ckObj))
}
// property setter: ReadDirMustMatch
func (z *SFtp) SetReadDirMustMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putReadDirMustMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ReadDirMustNotMatch
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "*.xml; *.txt; *.csv". If set, the ReadDir method will only return
// files that do not match any of these patterns.
func (z *SFtp) ReadDirMustNotMatch() string {
    return C.GoString(C.CkSFtp_readDirMustNotMatch(z.ckObj))
}
// property setter: ReadDirMustNotMatch
func (z *SFtp) SetReadDirMustNotMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putReadDirMustNotMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ServerIdentifier
// The server-identifier string received from the server during connection
// establishment. For example, a typical value would be similar to
// "SSH-2.0-OpenSSH_7.2p2 Ubuntu-4ubuntu2.2".
func (z *SFtp) ServerIdentifier() string {
    return C.GoString(C.CkSFtp_serverIdentifier(z.ckObj))
}

// property: SessionLog
// Contains a log of the messages sent to/from the SFTP server. To enable session
// logging, set the KeepSessionLog property = true. Note: This property is not a
// filename -- it is a string property that contains the session log data.
func (z *SFtp) SessionLog() string {
    return C.GoString(C.CkSFtp_sessionLog(z.ckObj))
}

// property: SocksHostname
// The SOCKS4/SOCKS5 hostname or IPv4 address (in dotted decimal notation). This
// property is only used if the SocksVersion property is set to 4 or 5).
func (z *SFtp) SocksHostname() string {
    return C.GoString(C.CkSFtp_socksHostname(z.ckObj))
}
// property setter: SocksHostname
func (z *SFtp) SetSocksHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putSocksHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPassword
// The SOCKS5 password (if required). The SOCKS4 protocol does not include the use
// of a password, so this does not apply to SOCKS4.
func (z *SFtp) SocksPassword() string {
    return C.GoString(C.CkSFtp_socksPassword(z.ckObj))
}
// property setter: SocksPassword
func (z *SFtp) SetSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPort
// The SOCKS4/SOCKS5 proxy port. The default value is 1080. This property only
// applies if a SOCKS proxy is used (if the SocksVersion property is set to 4 or
// 5).
func (z *SFtp) SocksPort() int {
    return int(C.CkSFtp_getSocksPort(z.ckObj))
}
// property setter: SocksPort
func (z *SFtp) SetSocksPort(value int) {
    C.CkSFtp_putSocksPort(z.ckObj,C.int(value))
}

// property: SocksUsername
// The SOCKS4/SOCKS5 proxy username. This property is only used if the SocksVersion
// property is set to 4 or 5).
func (z *SFtp) SocksUsername() string {
    return C.GoString(C.CkSFtp_socksUsername(z.ckObj))
}
// property setter: SocksUsername
func (z *SFtp) SetSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksVersion
// SocksVersion May be set to one of the following integer values:
// 
// 0 - No SOCKS proxy is used. This is the default.
// 4 - Connect via a SOCKS4 proxy.
// 5 - Connect via a SOCKS5 proxy.
//
func (z *SFtp) SocksVersion() int {
    return int(C.CkSFtp_getSocksVersion(z.ckObj))
}
// property setter: SocksVersion
func (z *SFtp) SetSocksVersion(value int) {
    C.CkSFtp_putSocksVersion(z.ckObj,C.int(value))
}

// property: SoRcvBuf
// Sets the receive buffer size socket option. Normally, this property should be
// left unchanged. The default value is 4194304.
// 
// This property can be increased if download performance seems slow. It is
// recommended to be a multiple of 4096.
//
func (z *SFtp) SoRcvBuf() int {
    return int(C.CkSFtp_getSoRcvBuf(z.ckObj))
}
// property setter: SoRcvBuf
func (z *SFtp) SetSoRcvBuf(value int) {
    C.CkSFtp_putSoRcvBuf(z.ckObj,C.int(value))
}

// property: SoSndBuf
// Sets the send buffer size socket option. Normally, this property should be left
// unchanged. The default value is 262144.
// 
// This property can be increased if upload performance seems slow. It is
// recommended to be a multiple of 4096. Testing with sizes such as 512K and 1MB is
// reasonable.
//
func (z *SFtp) SoSndBuf() int {
    return int(C.CkSFtp_getSoSndBuf(z.ckObj))
}
// property setter: SoSndBuf
func (z *SFtp) SetSoSndBuf(value int) {
    C.CkSFtp_putSoSndBuf(z.ckObj,C.int(value))
}

// property: SyncCreateAllLocalDirs
// If true, then empty directories on the server are created locally when doing a
// download synchronization. If false, then only directories containing files
// that are downloaded are auto-created.
// 
// The default value of this property is true.
//
func (z *SFtp) SyncCreateAllLocalDirs() bool {
    v := int(C.CkSFtp_getSyncCreateAllLocalDirs(z.ckObj))
    return v != 0
}
// property setter: SyncCreateAllLocalDirs
func (z *SFtp) SetSyncCreateAllLocalDirs(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putSyncCreateAllLocalDirs(z.ckObj,v)
}

// property: SyncDirectives
// A property that can contain a list of comma-separated keywords to control
// certain aspects of an upload or download synchronization (for the SyncTreeUpload
// and SyncTreeDownload methods). At this time there is only one possible
// directive, but others may be added in the future.
// 
// Set this property to "UploadIgnoreLocalOpenFailures" to skip local files that
// cannot be opened. A common reason for this would be if another process on the
// system has the file open for exclusive access.
//
func (z *SFtp) SyncDirectives() string {
    return C.GoString(C.CkSFtp_syncDirectives(z.ckObj))
}
// property setter: SyncDirectives
func (z *SFtp) SetSyncDirectives(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putSyncDirectives(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncedFiles
// The paths of the files uploaded or downloaded in the last call to SyncUploadTree
// or SyncDownloadTree. The paths are listed one per line. In both cases (for
// upload and download) each line contains the paths relative to the root synced
// directory.
// 
// Note: For SyncTreeDownload, some of entires can be the paths of local
// directories that were created. Starting in v9.5.0.77, local directory paths will
// be terminated with a "/" char (to disinguish a directory from an actual file).
//
func (z *SFtp) SyncedFiles() string {
    return C.GoString(C.CkSFtp_syncedFiles(z.ckObj))
}
// property setter: SyncedFiles
func (z *SFtp) SetSyncedFiles(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putSyncedFiles(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustMatch
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "*.xml; *.txt; *.csv". If set, the SyncTreeUpload and SyncTreeDownload
// methods will only transfer files that match any one of these patterns. This
// property only applies to files. It does not apply to sub-directory names when
// recursively traversing a directory tree.
func (z *SFtp) SyncMustMatch() string {
    return C.GoString(C.CkSFtp_syncMustMatch(z.ckObj))
}
// property setter: SyncMustMatch
func (z *SFtp) SetSyncMustMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putSyncMustMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustMatchDir
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "xml; txt; data_*". If set, the SyncTreeUpload and SyncTreeDownload
// methods will only enter directories that match any one of these patterns.
func (z *SFtp) SyncMustMatchDir() string {
    return C.GoString(C.CkSFtp_syncMustMatchDir(z.ckObj))
}
// property setter: SyncMustMatchDir
func (z *SFtp) SetSyncMustMatchDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putSyncMustMatchDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustNotMatch
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "*.xml; *.txt; *.csv". If set, the SyncTreeUpload and SyncTreeDownload
// methods will not transfer files that match any one of these patterns. This
// property only applies to files. It does not apply to sub-directory names when
// recursively traversing a directory tree.
func (z *SFtp) SyncMustNotMatch() string {
    return C.GoString(C.CkSFtp_syncMustNotMatch(z.ckObj))
}
// property setter: SyncMustNotMatch
func (z *SFtp) SetSyncMustNotMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putSyncMustNotMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustNotMatchDir
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "xml; txt; data_*". If set, the SyncTreeUpload and SyncTreeDownload
// methods will not enter directories that match any one of these patterns.
func (z *SFtp) SyncMustNotMatchDir() string {
    return C.GoString(C.CkSFtp_syncMustNotMatchDir(z.ckObj))
}
// property setter: SyncMustNotMatchDir
func (z *SFtp) SetSyncMustNotMatchDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putSyncMustNotMatchDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TcpNoDelay
// This property controls the use of the internal TCP_NODELAY socket option (which
// disables the Nagle algorithm). The default value of this property is false.
// Setting this value to true disables the delay of sending successive small
// packets on the network.
func (z *SFtp) TcpNoDelay() bool {
    v := int(C.CkSFtp_getTcpNoDelay(z.ckObj))
    return v != 0
}
// property setter: TcpNoDelay
func (z *SFtp) SetTcpNoDelay(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putTcpNoDelay(z.ckObj,v)
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
//     "NoSyncDownloadEmptyFiles" - Introduced in v9.5.0.80. Prevents empty files
//     from being downloaded in SyncTreeDownload.
//
func (z *SFtp) UncommonOptions() string {
    return C.GoString(C.CkSFtp_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *SFtp) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkSFtp_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UploadChunkSize
// The chunk size to use when uploading files via the UploadFile or
// UploadFileByName methods. The default value is 32000. If an upload fails because
// "WSAECONNABORTED An established connection was aborted by the software in your
// host machine.", then try setting this property to a smaller value, such as 4096.
// A smaller value will result in slower transfer rates, but may help avoid this
// problem.
func (z *SFtp) UploadChunkSize() int {
    return int(C.CkSFtp_getUploadChunkSize(z.ckObj))
}
// property setter: UploadChunkSize
func (z *SFtp) SetUploadChunkSize(value int) {
    C.CkSFtp_putUploadChunkSize(z.ckObj,C.int(value))
}

// property: UtcMode
// If true, then date/times are returned as UTC times. If false (the default)
// then date/times are returned as local times.
func (z *SFtp) UtcMode() bool {
    v := int(C.CkSFtp_getUtcMode(z.ckObj))
    return v != 0
}
// property setter: UtcMode
func (z *SFtp) SetUtcMode(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putUtcMode(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *SFtp) VerboseLogging() bool {
    v := int(C.CkSFtp_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *SFtp) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSFtp_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *SFtp) Version() string {
    return C.GoString(C.CkSFtp_version(z.ckObj))
}

// property: XferByteCount
// The current transfer byte count for an ongoing upload or download. Programs
// doing asynchronous uploads or downloads can read this property in real time
// during the upload. For SyncTreeUpload and SyncTreeDownload operations, this is
// the real-time cumulative number of bytes for all files uploaded or downloaded.
func (z *SFtp) XferByteCount() uint {
    return uint(C.CkSFtp_getXferByteCount(z.ckObj))
}

// property: XferByteCount64
// The current transfer byte count for an ongoing upload or download. Programs
// doing asynchronous uploads or downloads can read this property in real time
// during the upload. For SyncTreeUpload and SyncTreeDownload operations, this is
// the real-time cumulative number of bytes for all files uploaded or downloaded.
func (z *SFtp) XferByteCount64() int64 {
    return int64(C.CkSFtp_getXferByteCount64(z.ckObj))
}
// Downloads bytes from an open file and appends them to the AccumulateBuffer. The
// handle is a file handle returned by the OpenFile method. The maxBytes is the maximum
// number of bytes to read. If the end-of-file is reached prior to reading the
// number of requested bytes, then fewer bytes may be returned.
// 
// Returns the number of bytes downloaded and appended to AccumulateBuffer. Returns
// -1 on error.
//
func (z *SFtp) AccumulateBytes(arg1 string, arg2 int) int {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_AccumulateBytes(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the AccumulateBytes method
func (z *SFtp) AccumulateBytesAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_AccumulateBytesAsync(z.ckObj, cstr1, C.int(arg2))

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


// Convenience method for 64-bit addition. Allows for two 64-bit integers to be
// passed as decimal strings and returns the sum in a decimal sting.
// return a string or nil if failed.
func (z *SFtp) Add64(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkSFtp_add64(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Authenticates with the SSH server using public-key authentication. The
// corresponding public key must have been installed on the SSH server for the
// username. Authentication will succeed if the matching privateKey is provided.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
//
func (z *SFtp) AuthenticatePk(arg1 string, arg2 *SshKey) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_AuthenticatePk(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the AuthenticatePk method
func (z *SFtp) AuthenticatePkAsync(arg1 string, arg2 *SshKey, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_AuthenticatePkAsync(z.ckObj, cstr1, arg2.ckObj)

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
// An SFTP session always begins by first calling Connect to connect to the SSH
// server, then calling either AuthenticatePw or AuthenticatePk to login, and
// finally calling InitializeSftp.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
// 
// Note: To learn about how to handle password change requests, see the
// PasswordChangeRequested property (above).
//
func (z *SFtp) AuthenticatePw(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_AuthenticatePw(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AuthenticatePw method
func (z *SFtp) AuthenticatePwAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_AuthenticatePwAsync(z.ckObj, cstr1, cstr2)

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
func (z *SFtp) AuthenticatePwPk(arg1 string, arg2 string, arg3 *SshKey) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_AuthenticatePwPk(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AuthenticatePwPk method
func (z *SFtp) AuthenticatePwPkAsync(arg1 string, arg2 string, arg3 *SshKey, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_AuthenticatePwPkAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// The same as AuthenticatePw, but the login and password are passed in secure
// string objects.
func (z *SFtp) AuthenticateSecPw(arg1 *SecureString, arg2 *SecureString) bool {

    v := C.CkSFtp_AuthenticateSecPw(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Asynchronous version of the AuthenticateSecPw method
func (z *SFtp) AuthenticateSecPwAsync(arg1 *SecureString, arg2 *SecureString, c chan *Task) {

    hTask := C.CkSFtp_AuthenticateSecPwAsync(z.ckObj, arg1.ckObj, arg2.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The same as AuthenticatePwPk, but the login and password are passed in secure
// string objects.
func (z *SFtp) AuthenticateSecPwPk(arg1 *SecureString, arg2 *SecureString, arg3 *SshKey) bool {

    v := C.CkSFtp_AuthenticateSecPwPk(z.ckObj, arg1.ckObj, arg2.ckObj, arg3.ckObj)


    return v != 0
}


// Asynchronous version of the AuthenticateSecPwPk method
func (z *SFtp) AuthenticateSecPwPkAsync(arg1 *SecureString, arg2 *SecureString, arg3 *SshKey, c chan *Task) {

    hTask := C.CkSFtp_AuthenticateSecPwPkAsync(z.ckObj, arg1.ckObj, arg2.ckObj, arg3.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Clears the contents of the AccumulateBuffer property.
func (z *SFtp) ClearAccumulateBuffer()  {

    C.CkSFtp_ClearAccumulateBuffer(z.ckObj)



}


// Clears the internal file attribute cache. (Please refer to the EnableCache
// property for more information about the file attribute cache.)
func (z *SFtp) ClearCache()  {

    C.CkSFtp_ClearCache(z.ckObj)



}


// Clears the contents of the SessionLog property.
func (z *SFtp) ClearSessionLog()  {

    C.CkSFtp_ClearSessionLog(z.ckObj)



}


// Closes a file on the SSH/SFTP server. handle is a file handle returned from a
// previous call to OpenFile or OpenDir.
func (z *SFtp) CloseHandle(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_CloseHandle(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the CloseHandle method
func (z *SFtp) CloseHandleAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_CloseHandleAsync(z.ckObj, cstr1)

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


// Connects to an SSH/SFTP server. The domainName may be a domain name or an IP address
// (example: 192.168.1.10). Both IPv4 and IPv6 addresses are supported. The port is
// typically 22, which is the standard port for SSH servers.
// 
// An SFTP session always begins by first calling Connect to connect to the SSH
// server, then calling either AuthenticatePw or AuthenticatePk to login, and
// finally calling InitializeSftp.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
//
func (z *SFtp) Connect(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_Connect(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Connect method
func (z *SFtp) ConnectAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_ConnectAsync(z.ckObj, cstr1, C.int(arg2))

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


// Connects to an SSH/SFTP server through an existing SSH connection. The sshConn is
// an existing connected and authenticated SSH object. The connection to hostname:port
// is made through the existing SSH connection via port-forwarding. If successful,
// the connection is as follows: application => ServerSSH1 => ServerSSH2. (where
// ServerSSH1 is the sshConn and ServerSSH2 is the SSH server at hostname:port) Once
// connected in this way, all communications are routed through ServerSSH1 to
// ServerSSH2. This includes authentication -- which means the application must
// still call one of the Authenticate* methods to authenticate with ServerSSH2.
func (z *SFtp) ConnectThroughSsh(arg1 *Ssh, arg2 string, arg3 int) bool {
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_ConnectThroughSsh(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the ConnectThroughSsh method
func (z *SFtp) ConnectThroughSshAsync(arg1 *Ssh, arg2 string, arg3 int, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_ConnectThroughSshAsync(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

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


// Sets the date/time and other attributes of a remote file to be equal to that of
// a local file.
// 
// The attributes copied depend on the SFTP version of the server:
// SFTP v3 (and below)
//     Last-Modified Date/Time
//     Last-Access Date/Time
// 
// SFTP v4, v5
//     Last-Modified Date/Time
//     Last-Access Date/Time
//     Create Date/Time
// 
// SFTP v6 (and above)
//     Last-Modified Date/Time
//     Last-Access Date/Time
//     Create Date/Time
//     Read-Only Flag
//     Hidden Flag
//     Archive Flag
//     Compressed Flag
//     Encrypted Flag
// 
// Notes:
// (1) The Last-Access date/time may or may not be set. Chilkat has found that the
// Last-Access time is set to the current date/time, which is probably a result of
// the operating system setting it based on when the SFTP server is touching the
// file.
// (2) At the time of this writing, it is unknown whether the compressed/encryption
// settings for a local file are transferred to the remote file. For example, does
// the remote file become compressed and/or encrypted just by setting the flags? It
// may depend on the SFTP server and/or the remote filesystem.
// (3) Dates/times are sent in GMT. SFTP servers should convert GMT times to local
// time zones.
//
func (z *SFtp) CopyFileAttr(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkSFtp_CopyFileAttr(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CopyFileAttr method
func (z *SFtp) CopyFileAttrAsync(arg1 string, arg2 string, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_CopyFileAttrAsync(z.ckObj, cstr1, cstr2, b3)

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


// Creates a directory on the SFTP server.
func (z *SFtp) CreateDir(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_CreateDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the CreateDir method
func (z *SFtp) CreateDirAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_CreateDirAsync(z.ckObj, cstr1)

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
func (z *SFtp) Disconnect()  {

    C.CkSFtp_Disconnect(z.ckObj)



}


// Downloads the contents of a remote file to a BinData object. (Appends to the
// BinData.)
func (z *SFtp) DownloadBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_DownloadBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the DownloadBd method
func (z *SFtp) DownloadBdAsync(arg1 string, arg2 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_DownloadBdAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Downloads a file from the SSH server to the local filesystem. There are no
// limitations on file size and the data is streamed from SSH server to the local
// file. handle is a file handle returned by a previous call to OpenFile.
func (z *SFtp) DownloadFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_DownloadFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the DownloadFile method
func (z *SFtp) DownloadFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_DownloadFileAsync(z.ckObj, cstr1, cstr2)

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


// Simplified method for downloading files.
// 
// The last-modified date/time is only preserved when the PreserveDate property is
// set to true. (The default value of PreserveDate is false.)
//
func (z *SFtp) DownloadFileByName(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_DownloadFileByName(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the DownloadFileByName method
func (z *SFtp) DownloadFileByNameAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_DownloadFileByNameAsync(z.ckObj, cstr1, cstr2)

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


// Downloads the contents of a remote file to a StringBuilder object. (Appends to
// the StringBuilder.)
func (z *SFtp) DownloadSb(arg1 string, arg2 string, arg3 *StringBuilder) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_DownloadSb(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the DownloadSb method
func (z *SFtp) DownloadSbAsync(arg1 string, arg2 string, arg3 *StringBuilder, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_DownloadSbAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// Returns true if the last read operation for a handle reached the end of file.
// Otherwise returns false. If an invalid handle is passed, a value of true is
// returned.
func (z *SFtp) Eof(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_Eof(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns a value to indicate whether the remote file exists or not. remotePath is the
// path of the remote file. If followLinks is true, then symbolic links will be followed
// on the server.
// 
// This method returns one of the following possible values:
//     -1: Unable to check. Examine the LastErrorText to determine the reason for
//     failure.
//     0: File does not exist.
//     1: The regular file exists.
//     2: It exists, but it is a directory.
//     3: It exists, but it is a symlink (only possible if followLinks is false)
//     4: It exists, but it is a special filesystem entry type.
//     5: It exists, but it is an unkown filesystem entry type.
//     6: It exists, but it is an socket filesystem entry type.
//     7: It exists, but it is an char device entry type.
//     8: It exists, but it is an block device entry type.
//     9: It exists, but it is an FIFO entry type.
// 
// Note: The values greater than zero correspond to the possible values as
// specified in the SFTP protocol specification. A given value may not make sense
// on all operating systems.
//
func (z *SFtp) FileExists(arg1 string, arg2 bool) int {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkSFtp_FileExists(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the FileExists method
func (z *SFtp) FileExistsAsync(arg1 string, arg2 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkSFtp_FileExistsAsync(z.ckObj, cstr1, b2)

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


// Causes the SFTP server to do an fsync on the open file. Specifically, this is
// directing the SFTP server to call fsync (https://linux.die.net/man/2/fsync) on
// the open file.
// 
// This method uses the fsync@openssh.com and only works for servers supporting the
// fsync@openssh.com extension.
//
func (z *SFtp) Fsync(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_Fsync(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Fsync method
func (z *SFtp) FsyncAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_FsyncAsync(z.ckObj, cstr1)

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


// Returns the create date/time for a file. pathOrHandle may be a remote filepath or an
// open handle string as returned by OpenFile. If pathOrHandle is a handle, then bIsHandle must
// be set to true, otherwise it should be false. If bFollowLinks is true, then
// symbolic links will be followed on the server.
// 
// Note: Servers running the SFTP v3 protocol or lower do not have the ability to
// return a file's creation date/time.
//
func (z *SFtp) GetFileCreateDt(arg1 string, arg2 bool, arg3 bool) *CkDateTime {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retObj := C.CkSFtp_GetFileCreateDt(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Asynchronous version of the GetFileCreateDt method
func (z *SFtp) GetFileCreateDtAsync(arg1 string, arg2 bool, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_GetFileCreateDtAsync(z.ckObj, cstr1, b2, b3)

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


// The same as GetFileCreateTime, except the date/time is returned as an RFC822
// formatted string.
// return a string or nil if failed.
func (z *SFtp) GetFileCreateTimeStr(arg1 string, arg2 bool, arg3 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retStr := C.CkSFtp_getFileCreateTimeStr(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetFileCreateTimeStr method
func (z *SFtp) GetFileCreateTimeStrAsync(arg1 string, arg2 bool, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_GetFileCreateTimeStrAsync(z.ckObj, cstr1, b2, b3)

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


// Returns the group of a file. pathOrHandle may be a remote filepath or an open handle
// string as returned by OpenFile. If pathOrHandle is a handle, then bIsHandle must be set to
// true, otherwise it should be false. If bFollowLinks is true, then symbolic links
// will be followed on the server.
// 
// Note: Servers running the SFTP v3 protocol or lower do not have the ability to
// return a file's group name. Instead, the decimal GID of the file is returned.
//
// return a string or nil if failed.
func (z *SFtp) GetFileGroup(arg1 string, arg2 bool, arg3 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retStr := C.CkSFtp_getFileGroup(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetFileGroup method
func (z *SFtp) GetFileGroupAsync(arg1 string, arg2 bool, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_GetFileGroupAsync(z.ckObj, cstr1, b2, b3)

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


// Returns the last-access date/time for a file. pathOrHandle may be a remote filepath or
// an open handle string as returned by OpenFile. If pathOrHandle is a handle, then bIsHandle
// must be set to true, otherwise it should be false. If bFollowLinks is true, then
// symbolic links will be followed on the server.
func (z *SFtp) GetFileLastAccessDt(arg1 string, arg2 bool, arg3 bool) *CkDateTime {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retObj := C.CkSFtp_GetFileLastAccessDt(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Asynchronous version of the GetFileLastAccessDt method
func (z *SFtp) GetFileLastAccessDtAsync(arg1 string, arg2 bool, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_GetFileLastAccessDtAsync(z.ckObj, cstr1, b2, b3)

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


// The same as GetFileLastAccess, except the date/time is returned as an RFC822
// formatted string.
// return a string or nil if failed.
func (z *SFtp) GetFileLastAccessStr(arg1 string, arg2 bool, arg3 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retStr := C.CkSFtp_getFileLastAccessStr(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetFileLastAccessStr method
func (z *SFtp) GetFileLastAccessStrAsync(arg1 string, arg2 bool, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_GetFileLastAccessStrAsync(z.ckObj, cstr1, b2, b3)

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


// Returns the last-modified date/time for a file. pathOrHandle may be a remote filepath or
// an open handle string as returned by OpenFile. If pathOrHandle is a handle, then bIsHandle
// must be set to true, otherwise it should be false. If bFollowLinks is true, then
// symbolic links will be followed on the server.
func (z *SFtp) GetFileLastModifiedDt(arg1 string, arg2 bool, arg3 bool) *CkDateTime {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retObj := C.CkSFtp_GetFileLastModifiedDt(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Asynchronous version of the GetFileLastModifiedDt method
func (z *SFtp) GetFileLastModifiedDtAsync(arg1 string, arg2 bool, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_GetFileLastModifiedDtAsync(z.ckObj, cstr1, b2, b3)

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


// The same as GetFileLastModified, except the date/time is returned as an RFC822
// formatted string.
// return a string or nil if failed.
func (z *SFtp) GetFileLastModifiedStr(arg1 string, arg2 bool, arg3 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retStr := C.CkSFtp_getFileLastModifiedStr(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetFileLastModifiedStr method
func (z *SFtp) GetFileLastModifiedStrAsync(arg1 string, arg2 bool, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_GetFileLastModifiedStrAsync(z.ckObj, cstr1, b2, b3)

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


// Returns the owner of a file. pathOrHandle may be a remote filepath or an open handle
// string as returned by OpenFile. If pathOrHandle is a handle, then bIsHandle must be set to
// true, otherwise it should be false. If bFollowLinks is true, then symbolic links
// will be followed on the server.
// 
// Note: Servers running the SFTP v3 protocol or lower do not have the ability to
// return a file's owner name. Instead, the decimal UID of the file is returned.
//
// return a string or nil if failed.
func (z *SFtp) GetFileOwner(arg1 string, arg2 bool, arg3 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retStr := C.CkSFtp_getFileOwner(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetFileOwner method
func (z *SFtp) GetFileOwnerAsync(arg1 string, arg2 bool, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_GetFileOwnerAsync(z.ckObj, cstr1, b2, b3)

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


// Returns the access permisssions flags of a file. pathOrHandle may be a remote filepath
// or an open handle string as returned by OpenFile. If pathOrHandle is a handle, then bIsHandle
// must be set to true, otherwise it should be false. If bFollowLinks is true, then
// symbolic links will be followed on the server.
func (z *SFtp) GetFilePermissions(arg1 string, arg2 bool, arg3 bool) int {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkSFtp_GetFilePermissions(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the GetFilePermissions method
func (z *SFtp) GetFilePermissionsAsync(arg1 string, arg2 bool, arg3 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    hTask := C.CkSFtp_GetFilePermissionsAsync(z.ckObj, cstr1, b2, b3)

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


// Returns the size in bytes of a file on the SSH server. If the file size exceeds
// what can be represented in 32-bits, a value of -1 is returned. pathOrHandle may be a
// remote filepath or an open handle string as returned by OpenFile. If pathOrHandle is a
// handle, then bIsHandle must be set to true, otherwise it should be false. If bFollowLinks
// is true, then symbolic links will be followed on the server.
func (z *SFtp) GetFileSize32(arg1 string, arg2 bool, arg3 bool) int {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkSFtp_GetFileSize32(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns a 64-bit integer containing the size (in bytes) of a file on the SSH
// server. pathOrHandle may be a remote filepath or an open handle string as returned by
// OpenFile. If pathOrHandle is a handle, then bIsHandle must be set to true, otherwise it
// should be false. If bFollowLinks is true, then symbolic links will be followed on
// the server.
func (z *SFtp) GetFileSize64(arg1 string, arg2 bool, arg3 bool) int64 {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkSFtp_GetFileSize64(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

    return int64(v)
}


// Returns the size in bytes (in decimal string form) of a file on the SSH server.
// pathOrHandle may be a remote filepath or an open handle string as returned by OpenFile.
// If pathOrHandle is a handle, then bIsHandle must be set to true, otherwise it should be
// false. If bFollowLinks is true, then symbolic links will be followed on the server.
// 
// Note: This method exists for environments that do not have 64-bit integer
// support. The Add64 method is provided for 64-bit addition, and other methods
// such as ReadFileBytes64s allow for 64-bit values to be passed as strings.
//
// return a string or nil if failed.
func (z *SFtp) GetFileSizeStr(arg1 string, arg2 bool, arg3 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retStr := C.CkSFtp_getFileSizeStr(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Creates a hard link on the server using the hardlink@openssh.com extension. This
// only works for SFTP servers that support the hardlink@openssh.com extension.
func (z *SFtp) HardLink(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_HardLink(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the HardLink method
func (z *SFtp) HardLinkAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_HardLinkAsync(z.ckObj, cstr1, cstr2)

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


// Intializes the SFTP subsystem. This should be called after connecting and
// authenticating. An SFTP session always begins by first calling Connect to
// connect to the SSH server, then calling either AuthenticatePw or AuthenticatePk
// to login, and finally calling InitializeSftp.
// 
// Important: When reporting problems, please send the full contents of the
// LastErrorText property to support@chilkatsoft.com.
// 
// If this method fails, the reason may be present in the InitializeFailCode and
// InitializeFailReason properties (assuming the failure occurred when trying to
// open the SFTP session channel).
//
func (z *SFtp) InitializeSftp() bool {

    v := C.CkSFtp_InitializeSftp(z.ckObj)


    return v != 0
}


// Asynchronous version of the InitializeSftp method
func (z *SFtp) InitializeSftpAsync(c chan *Task) {

    hTask := C.CkSFtp_InitializeSftpAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Provides information about what transpired in the last method called. For many
// methods, there is no information. For some methods, details about what
// transpired can be obtained via LastJsonData.
func (z *SFtp) LastJsonData() *JsonObject {

    retObj := C.CkSFtp_LastJsonData(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Returns true if the last read on the specified handle failed. Otherwise
// returns false.
func (z *SFtp) LastReadFailed(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_LastReadFailed(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the number of bytes received by the last read on a specified channel.
func (z *SFtp) LastReadNumBytes(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_LastReadNumBytes(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Loads the caller of the task's async method.
func (z *SFtp) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkSFtp_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Opens a directory for reading. To get a directory listing, first open the
// directory by calling this method, then call ReadDir to read the directory, and
// finally call CloseHandle to close the directory.
// 
// The SFTP protocol represents file names as strings. File names are assumed to
// use the slash ('/') character as a directory separator.
// 
// File names starting with a slash are "absolute", and are relative to the root of
// the file system. Names starting with any other character are relative to the
// user's default directory (home directory). Note that identifying the user is
// assumed to take place outside of this protocol.
// 
// Servers SHOULD interpret a path name component ".." as referring to the parent
// directory, and "." as referring to the current directory.
// 
// An empty path name is valid, and it refers to the user's default directory
// (usually the user's home directory).
// 
// Please note: This method does NOT "change" the remote working directory. It is
// only a method for opening a directory for the purpose of reading the directory
// listing.
// 
// SFTP is Secure File Transfer over SSH. It is not the FTP protocol. There is no
// similarity or relationship between FTP and SFTP. Therefore, concepts such as
// "current remote directory" that exist in FTP do not exist with SFTP. With the
// SFTP protocol, the current directory will always be the home directory of the
// user account used during SSH/SFTP authentication. You may pass relative or
// absolute directory/file paths. A relative path is always relative to the home
// directory of the SSH user account.
//
// return a string or nil if failed.
func (z *SFtp) OpenDir(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSFtp_openDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the OpenDir method
func (z *SFtp) OpenDirAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_OpenDirAsync(z.ckObj, cstr1)

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


// Opens or creates a file on the remote system. Returns a handle which may be
// passed to methods for reading and/or writing the file. The remotePath is the remote
// file path (the path to the file on the server). When the application is finished
// with the handle, it should call CloseHandle(remotePath).
// 
// access should be one of the following strings: "readOnly", "writeOnly", or
// "readWrite".
// 
// createDisposition is a comma-separated list of keywords to provide more control over how the
// file is opened or created. One of the following keywords must be present:
// "createNew", "createTruncate", "openExisting", "openOrCreate", or
// "truncateExisting". All other keywords are optional. The list of keywords and
// their meanings are shown here:
// 
// createNew
// A new file is created; if the file already exists the method fails.
// 
// createTruncate
// A new file is created; if the file already exists, it is opened and truncated.
// 
// openExisting
// An existing file is opened. If the file does not exist the method fails.
// 
// openOrCreate
// If the file exists, it is opened. If the file does not exist, it is created.
// 
// truncateExisting
// An existing file is opened and truncated. If the file does not exist the method
// fails.
// 
// appendData
// Data is always written at the end of the file. Data is not required to be
// appended atomically. This means that if multiple writers attempt to append data
// simultaneously, data from the first may be lost.
// 
// appendDataAtomic
// Data is always written at the end of the file. Data MUST be written atomically
// so that there is no chance that multiple appenders can collide and result in
// data being lost.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// textMode
// Indicates that the server should treat the file as text and convert it to the
// canonical newline convention in use. When a file is opened with this flag, data
// is always appended to the end of the file. Servers MUST process multiple,
// parallel reads and writes correctly in this mode.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// blockRead
// The server MUST guarantee that no other handle has been opened with read access,
// and that no other handle will be opened with read access until the client closes
// the handle. (This MUST apply both to other clients and to other processes on the
// server.) In a nutshell, this opens the file in non-sharing mode.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// blockWrite
// The server MUST guarantee that no other handle has been opened with write
// access, and that no other handle will be opened with write access until the
// client closes the handle. (This MUST apply both to other clients and to other
// processes on the server.) In a nutshell, this opens the file in non-sharing
// mode.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// blockDelete
// The server MUST guarantee that the file itself is not deleted in any other way
// until the client closes the handle. No other client or process is allowed to
// open the file with delete access.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// blockAdvisory
// If set, the above "block" modes are advisory. In advisory mode, only other
// accesses that specify a "block" mode need be considered when determining whether
// the "block" can be granted, and the server need not prevent I/O operations that
// violate the block mode. The server MAY perform mandatory locking even if the
// blockAdvisory flag is set.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// noFollow
// If the final component of the path is a symlink, then the open MUST fail.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// deleteOnClose
// The file should be deleted when the last handle to it is closed. (The last
// handle may not be an sftp-handle.) This MAY be emulated by a server if the OS
// doesn't support it by deleting the file when this handle is closed.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// accessAuditAlarmInfo
// The client wishes the server to enable any privileges or extra capabilities that
// the user may have in to allow the reading and writing of AUDIT or ALARM access
// control entries.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// accessBackup
// The client wishes the server to enable any privileges or extra capabilities that
// the user may have in order to bypass normal access checks for the purpose of
// backing up or restoring files.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// backupStream
// This flag indicates that the client wishes to read or write a backup stream. A
// backup stream is a system dependent structured data stream that encodes all the
// information that must be preserved in order to restore the file from backup
// medium. The only well defined use for backup stream data read in this fashion is
// to write it to the same server to a file also opened using the backupStream
// flag. However, if the server has a well defined backup stream format, there may
// be other uses for this data outside the scope of this protocol.
// (Only supported in SFTP protocol versions 5 and later. See the note below.)
// 
// IMPORANT: If remotePath is a filename with no path, such as "test.txt", and the server
// responds with a "Folder not found" error, then try prepending "./" to the remotePath.
// For example, instead of passing "test.txt", try "./test.txt".
// 
// IMPORTANT note about createDisposition: Many of the options, such as textMode, are not
// implemented in the SFTP protocol versions 3 and 4. Only SFTP servers at protocol
// version 5 or later support these options. You can find out the protocol version
// of your server by examining the value of the ProtocolVersion property after
// calling InitializeSftp. Also, make sure the ForceV3 property is set to false
// (the default value is true)
//
// return a string or nil if failed.
func (z *SFtp) OpenFile(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkSFtp_openFile(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the OpenFile method
func (z *SFtp) OpenFileAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkSFtp_OpenFileAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Reads the contents of a directory and returns the directory listing (as an
// object). The handle returned by OpenDir should be passed to this method.
func (z *SFtp) ReadDir(arg1 string) *SFtpDir {
    cstr1 := C.CString(arg1)

    retObj := C.CkSFtp_ReadDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &SFtpDir{retObj}
}


// Asynchronous version of the ReadDir method
func (z *SFtp) ReadDirAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_ReadDirAsync(z.ckObj, cstr1)

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


// Reads file data from a remote file on the SSH server. The handle is a file handle
// returned by the OpenFile method. The numBytes is the maximum number of bytes to
// read. If the end-of-file is reached prior to reading the number of requested
// bytes, then fewer bytes may be returned. The received bytes are appended to the
// contents of bd.
// 
// To read an entire file, one may call ReadFileBd repeatedly until Eof(handle)
// returns true.
//
func (z *SFtp) ReadFileBd(arg1 string, arg2 int, arg3 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_ReadFileBd(z.ckObj, cstr1, C.int(arg2), arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the ReadFileBd method
func (z *SFtp) ReadFileBdAsync(arg1 string, arg2 int, arg3 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_ReadFileBdAsync(z.ckObj, cstr1, C.int(arg2), arg3.ckObj)

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


// Reads file data from a remote file on the SSH server. The handle is a file handle
// returned by the OpenFile method. The numBytes is the maximum number of bytes to
// read. If the end-of-file is reached prior to reading the number of requested
// bytes, then fewer bytes may be returned.
// 
// To read an entire file, one may call ReadFileBytes repeatedly until Eof(handle)
// returns true.
//
func (z *SFtp) ReadFileBytes(arg1 string, arg2 int) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSFtp_ReadFileBytes(z.ckObj, cstr1, C.int(arg2), ckOutBytes)

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


// Asynchronous version of the ReadFileBytes method
func (z *SFtp) ReadFileBytesAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_ReadFileBytesAsync(z.ckObj, cstr1, C.int(arg2))

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


// Reads file data from a remote file on the SSH server. The handle is a file handle
// returned by the OpenFile method. The offset is measured in bytes relative to the
// beginning of the file. (64-bit offsets are supported via the ReadFileBytes64 and
// ReadFileBytes64s methods.) The offset is ignored if the "textMode" flag was
// specified during the OpenFile. The numBytes is the maximum number of bytes to read.
// If the end-of-file is reached prior to reading the number of requested bytes,
// then fewer bytes may be returned.
func (z *SFtp) ReadFileBytes32(arg1 string, arg2 int, arg3 int) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSFtp_ReadFileBytes32(z.ckObj, cstr1, C.int(arg2), C.int(arg3), ckOutBytes)

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


// Reads file data from a remote file on the SSH server. The handle is a file handle
// returned by the OpenFile method. The offset is a 64-bit integer measured in bytes
// relative to the beginning of the file. The offset is ignored if the "textMode"
// flag was specified during the OpenFile. The numBytes is the maximum number of bytes
// to read. If the end-of-file is reached prior to reading the number of requested
// bytes, then fewer bytes may be returned.
func (z *SFtp) ReadFileBytes64(arg1 string, arg2 int64, arg3 int) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSFtp_ReadFileBytes64(z.ckObj, cstr1, C.__int64(arg2), C.int(arg3), ckOutBytes)

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


// (This method exists for systems that do not support 64-bit integers. The 64-bit
// integer offset is passed as a decimal string instead.)
// 
// Reads file data from a remote file on the SSH server. The handle is a file handle
// returned by the OpenFile method. The offset is a 64-bit integer represented as a
// decimal string. It represents an offset in bytes from the beginning of the file.
// The offset is ignored if the "textMode" flag was specified during the OpenFile.
// The numBytes is the maximum number of bytes to read. If the end-of-file is reached
// prior to reading the number of requested bytes, then fewer bytes may be
// returned.
//
func (z *SFtp) ReadFileBytes64s(arg1 string, arg2 string, arg3 int) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkSFtp_ReadFileBytes64s(z.ckObj, cstr1, cstr2, C.int(arg3), ckOutBytes)

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


// This method is identical to ReadFileBytes except for one thing: The bytes are
// interpreted according to the specified charset (i.e. the character encoding) and
// returned as a string. A list of supported charset values may be found on this
// page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
// 
// Note: If the charset is an encoding where a single character might be represented
// in multiple bytes (such as utf-8, Shift_JIS, etc.) then there is a risk that the
// very last character may be partially read. This is because the method specifies
// the number of bytes to read, not the number of characters. This is never a
// problem with character encodings that use a single byte per character, such as
// all of the iso-8859-* encodings, or the Windows-* encodings.
// 
// To read an entire file, one may call ReadFileText repeatedly until Eof(handle)
// returns true.
//
// return a string or nil if failed.
func (z *SFtp) ReadFileText(arg1 string, arg2 int, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    retStr := C.CkSFtp_readFileText(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReadFileText method
func (z *SFtp) ReadFileTextAsync(arg1 string, arg2 int, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    hTask := C.CkSFtp_ReadFileTextAsync(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
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


// This method is identical to ReadFileBytes32 except for one thing: The bytes are
// interpreted according to the specified charset (i.e. the character encoding) and
// returned as a string. A list of supported charset values may be found on this
// page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
// 
// Note: If the charset is an encoding where a single character might be represented
// in multiple bytes (such as utf-8, Shift_JIS, etc.) then there is a risk that the
// very last character may be partially read. This is because the method specifies
// the number of bytes to read, not the number of characters. This is never a
// problem with character encodings that use a single byte per character, such as
// all of the iso-8859-* encodings, or the Windows-* encodings.
//
// return a string or nil if failed.
func (z *SFtp) ReadFileText32(arg1 string, arg2 int, arg3 int, arg4 string) *string {
    cstr1 := C.CString(arg1)
    cstr4 := C.CString(arg4)

    retStr := C.CkSFtp_readFileText32(z.ckObj, cstr1, C.int(arg2), C.int(arg3), cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr4))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// This method is identical to ReadFileBytes64 except for one thing: The bytes are
// interpreted according to the specified charset (i.e. the character encoding) and
// returned as a string. A list of supported charset values may be found on this
// page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
// 
// Note: If the charset is an encoding where a single character might be represented
// in multiple bytes (such as utf-8, Shift_JIS, etc.) then there is a risk that the
// very last character may be partially read. This is because the method specifies
// the number of bytes to read, not the number of characters. This is never a
// problem with character encodings that use a single byte per character, such as
// all of the iso-8859-* encodings, or the Windows-* encodings.
//
// return a string or nil if failed.
func (z *SFtp) ReadFileText64(arg1 string, arg2 int64, arg3 int, arg4 string) *string {
    cstr1 := C.CString(arg1)
    cstr4 := C.CString(arg4)

    retStr := C.CkSFtp_readFileText64(z.ckObj, cstr1, C.__int64(arg2), C.int(arg3), cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr4))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// This method is identical to ReadFileBytes64s except for one thing: The bytes are
// interpreted according to the specified charset (i.e. the character encoding) and
// returned as a string. A list of supported charset values may be found on this
// page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
// 
// Note: If the charset is an encoding where a single character might be represented
// in multiple bytes (such as utf-8, Shift_JIS, etc.) then there is a risk that the
// very last character may be partially read. This is because the method specifies
// the number of bytes to read, not the number of characters. This is never a
// problem with character encodings that use a single byte per character, such as
// all of the iso-8859-* encodings, or the Windows-* encodings.
//
// return a string or nil if failed.
func (z *SFtp) ReadFileText64s(arg1 string, arg2 string, arg3 int, arg4 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr4 := C.CString(arg4)

    retStr := C.CkSFtp_readFileText64s(z.ckObj, cstr1, cstr2, C.int(arg3), cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr4))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the target of a symbolic link on the server. The path is the path of the
// symbolic link on the server.
// return a string or nil if failed.
func (z *SFtp) ReadLink(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSFtp_readLink(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the ReadLink method
func (z *SFtp) ReadLinkAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_ReadLinkAsync(z.ckObj, cstr1)

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


// This method can be used to have the server canonicalize any given path name to
// an absolute path. This is useful for converting path names containing ".."
// components or relative pathnames without a leading slash into absolute paths.
// The absolute path is returned by this method.
// 
// originalPath is the first component of the path which the client wishes resolved into a
// absolute canonical path. This may be the entire path.
// 
// The composePath is a path which the client wishes the server to compose with the
// original path to form the new path. This field is optional and may be set to a
// zero-length string.
// 
// The server will take the originalPath and apply the composePath as a modification to it. composePath
// may be relative to originalPath or may be an absolute path, in which case originalPath will be
// discarded. The composePath may be zero length.
// 
// Note: Servers running SFTP v4 and below do not support composePath.
//
// return a string or nil if failed.
func (z *SFtp) RealPath(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkSFtp_realPath(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the RealPath method
func (z *SFtp) RealPathAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_RealPathAsync(z.ckObj, cstr1, cstr2)

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


// Deletes a directory on the remote server. Most (if not all) SFTP servers require
// that the directorybe empty of files before it may be deleted.
func (z *SFtp) RemoveDir(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_RemoveDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the RemoveDir method
func (z *SFtp) RemoveDirAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_RemoveDirAsync(z.ckObj, cstr1)

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


// Deletes a file on the SFTP server.
func (z *SFtp) RemoveFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_RemoveFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the RemoveFile method
func (z *SFtp) RemoveFileAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_RemoveFileAsync(z.ckObj, cstr1)

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


// Renames a file or directory on the SFTP server.
func (z *SFtp) RenameFileOrDir(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_RenameFileOrDir(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the RenameFileOrDir method
func (z *SFtp) RenameFileOrDirAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_RenameFileOrDirAsync(z.ckObj, cstr1, cstr2)

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


// Resumes an SFTP download. The size of the localFilePath is checked and the download
// begins at the appropriate position in the remoteFilePath. If localFilePath is empty or
// non-existent, then this method is identical to DownloadFileByName. If the localFilePath
// is already fully downloaded, then no additional data is downloaded and the
// method will return true.
func (z *SFtp) ResumeDownloadFileByName(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_ResumeDownloadFileByName(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the ResumeDownloadFileByName method
func (z *SFtp) ResumeDownloadFileByNameAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_ResumeDownloadFileByNameAsync(z.ckObj, cstr1, cstr2)

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


// Resumes a file upload to the SFTP/SSH server. The size of the remoteFilePath is first
// checked to determine the starting offset for the upload. If remoteFilePath is empty or
// does not exist, this method is equivalent to UploadFileByName. If remoteFilePath is
// already fully uploaded (i.e. it's size is equal to localFilePath), then no additional
// bytes are uploaded and true is returned.
func (z *SFtp) ResumeUploadFileByName(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_ResumeUploadFileByName(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the ResumeUploadFileByName method
func (z *SFtp) ResumeUploadFileByNameAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_ResumeUploadFileByNameAsync(z.ckObj, cstr1, cstr2)

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


// Sends an IGNORE message to the SSH server. This is one way of verifying that the
// connection to the SSH server is open and valid. The SSH server does not respond
// to an IGNORE message. It simply ignores it. IGNORE messages are not associated
// with a channel (i.e., you do not need to first open a channel prior to sending
// an IGNORE message).
func (z *SFtp) SendIgnore() bool {

    v := C.CkSFtp_SendIgnore(z.ckObj)


    return v != 0
}


// Asynchronous version of the SendIgnore method
func (z *SFtp) SendIgnoreAsync(c chan *Task) {

    hTask := C.CkSFtp_SendIgnoreAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sets the create date/time for a file on the server. The pathOrHandle may be a filepath
// or the handle of a currently open file. isHandle should be set to true if the pathOrHandle
// is a handle, otherwise set isHandle to false.
// 
// Note: Servers running version 3 or lower of the SFTP protocol do not support
// setting the create date/time.
//
func (z *SFtp) SetCreateDt(arg1 string, arg2 bool, arg3 *CkDateTime) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkSFtp_SetCreateDt(z.ckObj, cstr1, b2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SetCreateDt method
func (z *SFtp) SetCreateDtAsync(arg1 string, arg2 bool, arg3 *CkDateTime, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkSFtp_SetCreateDtAsync(z.ckObj, cstr1, b2, arg3.ckObj)

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


// The same as SetCreateTime, except the date/time is passed as an RFC822 formatted
// string.
func (z *SFtp) SetCreateTimeStr(arg1 string, arg2 bool, arg3 string) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    v := C.CkSFtp_SetCreateTimeStr(z.ckObj, cstr1, b2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the SetCreateTimeStr method
func (z *SFtp) SetCreateTimeStrAsync(arg1 string, arg2 bool, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    hTask := C.CkSFtp_SetCreateTimeStrAsync(z.ckObj, cstr1, b2, cstr3)

    C.free(unsafe.Pointer(cstr1))
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


// Sets the last-access date/time for a file on the server. The pathOrHandle may be a
// filepath or the handle of a currently open file. isHandle should be set to true if
// the pathOrHandle is a handle, otherwise set isHandle to false.
func (z *SFtp) SetLastAccessDt(arg1 string, arg2 bool, arg3 *CkDateTime) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkSFtp_SetLastAccessDt(z.ckObj, cstr1, b2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SetLastAccessDt method
func (z *SFtp) SetLastAccessDtAsync(arg1 string, arg2 bool, arg3 *CkDateTime, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkSFtp_SetLastAccessDtAsync(z.ckObj, cstr1, b2, arg3.ckObj)

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


// The same as SetLastAccessTime, except the date/time is passed as an RFC822
// formatted string.
func (z *SFtp) SetLastAccessTimeStr(arg1 string, arg2 bool, arg3 string) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    v := C.CkSFtp_SetLastAccessTimeStr(z.ckObj, cstr1, b2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the SetLastAccessTimeStr method
func (z *SFtp) SetLastAccessTimeStrAsync(arg1 string, arg2 bool, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    hTask := C.CkSFtp_SetLastAccessTimeStrAsync(z.ckObj, cstr1, b2, cstr3)

    C.free(unsafe.Pointer(cstr1))
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


// Sets the last-modified date/time for a file on the server. The pathOrHandle may be a
// filepath or the handle of a currently open file. isHandle should be set to true if
// the pathOrHandle is a handle, otherwise set isHandle to false.
func (z *SFtp) SetLastModifiedDt(arg1 string, arg2 bool, arg3 *CkDateTime) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkSFtp_SetLastModifiedDt(z.ckObj, cstr1, b2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SetLastModifiedDt method
func (z *SFtp) SetLastModifiedDtAsync(arg1 string, arg2 bool, arg3 *CkDateTime, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkSFtp_SetLastModifiedDtAsync(z.ckObj, cstr1, b2, arg3.ckObj)

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


// The same as SetLastModifiedTime, except the date/time is passed as an RFC822
// formatted string.
func (z *SFtp) SetLastModifiedTimeStr(arg1 string, arg2 bool, arg3 string) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    v := C.CkSFtp_SetLastModifiedTimeStr(z.ckObj, cstr1, b2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the SetLastModifiedTimeStr method
func (z *SFtp) SetLastModifiedTimeStrAsync(arg1 string, arg2 bool, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    hTask := C.CkSFtp_SetLastModifiedTimeStrAsync(z.ckObj, cstr1, b2, cstr3)

    C.free(unsafe.Pointer(cstr1))
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


// Sets the owner and group for a file on the server. The pathOrHandle may be a filepath or
// the handle of a currently open file. isHandle should be set to true if the pathOrHandle is
// a handle, otherwise set isHandle to false.
// 
// Note: Servers running version 3 or lower of the SFTP protocol do not support
// setting the owner and group.
//
func (z *SFtp) SetOwnerAndGroup(arg1 string, arg2 bool, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkSFtp_SetOwnerAndGroup(z.ckObj, cstr1, b2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Asynchronous version of the SetOwnerAndGroup method
func (z *SFtp) SetOwnerAndGroupAsync(arg1 string, arg2 bool, arg3 string, arg4 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    hTask := C.CkSFtp_SetOwnerAndGroupAsync(z.ckObj, cstr1, b2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
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


// Sets the permissions for a file on the server. The pathOrHandle may be a filepath or the
// handle of a currently open file. isHandle should be set to true if the pathOrHandle is a
// handle, otherwise set isHandle to false.
func (z *SFtp) SetPermissions(arg1 string, arg2 bool, arg3 int) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkSFtp_SetPermissions(z.ckObj, cstr1, b2, C.int(arg3))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SetPermissions method
func (z *SFtp) SetPermissionsAsync(arg1 string, arg2 bool, arg3 int, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkSFtp_SetPermissionsAsync(z.ckObj, cstr1, b2, C.int(arg3))

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


// Create a symbolic link from oldpath to newpath on the server filesystem.
func (z *SFtp) SymLink(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_SymLink(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SymLink method
func (z *SFtp) SymLinkAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_SymLinkAsync(z.ckObj, cstr1, cstr2)

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


// Downloads files from the SFTP server to a local directory tree. Synchronization
// modes include:
// 
//     mode=0: Download all files
//     mode=1: Download all files that do not exist on the local filesystem.
//     mode=2: Download newer or non-existant files.
//     mode=3: Download only newer files. If a file does not already exist on the
//     local filesystem, it is not downloaded from the server.
//     mode=5: Download only missing files or files with size differences.
//     mode=6: Same as mode 5, but also download newer files.
//     mode=99: Do not download files, but instead delete remote files that do not
//     exist locally.
//     
// 
// If recurse is false, then the remotel directory tree is not recursively
// descended.
//
func (z *SFtp) SyncTreeDownload(arg1 string, arg2 string, arg3 int, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkSFtp_SyncTreeDownload(z.ckObj, cstr1, cstr2, C.int(arg3), b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SyncTreeDownload method
func (z *SFtp) SyncTreeDownloadAsync(arg1 string, arg2 string, arg3 int, arg4 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkSFtp_SyncTreeDownloadAsync(z.ckObj, cstr1, cstr2, C.int(arg3), b4)

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


// Uploads a directory tree from the local filesystem to the SFTP server.
// Synchronization modes include:
// 
//     mode=0: Upload all files
//     mode=1: Upload all files that do not exist on the server.
//     mode=2: Upload newer or non-existant files.
//     mode=3: Upload only newer files. If a file does not already exist on the
//     server, it is not uploaded.
//     mode=4: transfer missing files or files with size differences.
//     mode=5: same as mode 4, but also newer files.
// 
// If bRecurse is false, then the local directory tree is not recursively descended.
func (z *SFtp) SyncTreeUpload(arg1 string, arg2 string, arg3 int, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkSFtp_SyncTreeUpload(z.ckObj, cstr1, cstr2, C.int(arg3), b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SyncTreeUpload method
func (z *SFtp) SyncTreeUploadAsync(arg1 string, arg2 string, arg3 int, arg4 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkSFtp_SyncTreeUploadAsync(z.ckObj, cstr1, cstr2, C.int(arg3), b4)

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


// Unlocks the component. This must be called once prior to calling any other
// method. A fully-functional 30-day trial is automatically started when an
// arbitrary string is passed to this method. For example, passing "Hello", or
// "abc123" will unlock the component for the 1st thirty days after the initial
// install.
// 
// A purchased unlock code for SFTP should contain the substring ".SS" or "SSH" (or
// it can be a Bundle unlock code) because SFTP is the Secure File Transfer
// protocol over SSH. It is a sub-system of the SSH protocol. It is not the FTP
// protocol. If the Chilkat FTP2 component/library should be used for the FTP
// protocol.
//
func (z *SFtp) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Uploads the contents of a BinData to a remote file.
func (z *SFtp) UploadBd(arg1 *BinData, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_UploadBd(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UploadBd method
func (z *SFtp) UploadBdAsync(arg1 *BinData, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_UploadBdAsync(z.ckObj, arg1.ckObj, cstr2)

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


// Uploads a file from the local filesystem to the SFTP server. handle is a handle of
// a currently open file (obtained by calling the OpenFile method). fromLocalFilePath is the
// local file path of the file to be uploaded.
func (z *SFtp) UploadFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_UploadFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UploadFile method
func (z *SFtp) UploadFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_UploadFileAsync(z.ckObj, cstr1, cstr2)

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


// Simplified method for uploading a file to the SFTP/SSH server.
// 
// The last-modified date/time is only preserved if the PreserveDate property is
// set to true. This behavior of maintaining the last-mod date/time was
// introduced in v9.5.0.40.
//
func (z *SFtp) UploadFileByName(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkSFtp_UploadFileByName(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the UploadFileByName method
func (z *SFtp) UploadFileByNameAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkSFtp_UploadFileByNameAsync(z.ckObj, cstr1, cstr2)

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


// Uploads the contents of a StringBuilder to a remote file.
func (z *SFtp) UploadSb(arg1 *StringBuilder, arg2 string, arg3 string, arg4 bool) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkSFtp_UploadSb(z.ckObj, arg1.ckObj, cstr2, cstr3, b4)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the UploadSb method
func (z *SFtp) UploadSbAsync(arg1 *StringBuilder, arg2 string, arg3 string, arg4 bool, c chan *Task) {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkSFtp_UploadSbAsync(z.ckObj, arg1.ckObj, cstr2, cstr3, b4)

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


// Appends the contents of bd to an open file. The handle is a file handle returned
// by the OpenFile method.
func (z *SFtp) WriteFileBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkSFtp_WriteFileBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the WriteFileBd method
func (z *SFtp) WriteFileBdAsync(arg1 string, arg2 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkSFtp_WriteFileBdAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Appends byte data to an open file. The handle is a file handle returned by the
// OpenFile method.
func (z *SFtp) WriteFileBytes(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkSFtp_WriteFileBytes(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Asynchronous version of the WriteFileBytes method
func (z *SFtp) WriteFileBytesAsync(arg1 string, arg2 []byte, c chan *Task) {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    hTask := C.CkSFtp_WriteFileBytesAsync(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
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


// Writes data to an open file at a specific offset from the beginning of the file.
// The handle is a file handle returned by the OpenFile method. The offset is an offset
// from the beginning of the file.
func (z *SFtp) WriteFileBytes32(arg1 string, arg2 int, arg3 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkSFtp_WriteFileBytes32(z.ckObj, cstr1, C.int(arg2), ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// Writes data to an open file at a specific offset from the beginning of the file.
// The handle is a file handle returned by the OpenFile method. The offset64 is an offset
// from the beginning of the file.
func (z *SFtp) WriteFileBytes64(arg1 string, arg2 int64, arg3 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkSFtp_WriteFileBytes64(z.ckObj, cstr1, C.__int64(arg2), ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// Writes data to an open file at a specific offset from the beginning of the file.
// The handle is a file handle returned by the OpenFile method. The offset64 is an offset
// (in decimal string format) from the beginning of the file.
func (z *SFtp) WriteFileBytes64s(arg1 string, arg2 string, arg3 []byte) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkSFtp_WriteFileBytes64s(z.ckObj, cstr1, cstr2, ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// Appends character data to an open file. The handle is a file handle returned by
// the OpenFile method. charset is a character encoding and is typically set to values
// such as "ansi", "utf-8", "windows-1252", etc. A list of supported character
// encodings is found on this page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
// 
// Note: It is necessary to specify the character encoding because in many
// programming languages, strings are represented as Unicode (2 bytes/char) and in
// most cases one does not wish to write Unicode chars to a text file (although it
// is possible by setting charset = "Unicode").
//
func (z *SFtp) WriteFileText(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkSFtp_WriteFileText(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the WriteFileText method
func (z *SFtp) WriteFileTextAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkSFtp_WriteFileTextAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Writes character data to an open file at a specific offset from the beginning of
// the file. The handle is a file handle returned by the OpenFile method. charset is a
// character encoding and is typically set to values such as "ansi", "utf-8",
// "windows-1252", etc. A list of supported character encodings is found on this
// page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
func (z *SFtp) WriteFileText32(arg1 string, arg2 int, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkSFtp_WriteFileText32(z.ckObj, cstr1, C.int(arg2), cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Writes character data to an open file at a specific offset from the beginning of
// the file. The handle is a file handle returned by the OpenFile method. charset is a
// character encoding and is typically set to values such as "ansi", "utf-8",
// "windows-1252", etc. A list of supported character encodings is found on this
// page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
func (z *SFtp) WriteFileText64(arg1 string, arg2 int64, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkSFtp_WriteFileText64(z.ckObj, cstr1, C.__int64(arg2), cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Writes character data to an open file at a specific offset from the beginning of
// the file. The handle is a file handle returned by the OpenFile method. The offset64 is
// an offset (in decimal string format) from the beginning of the file. charset is a
// character encoding and is typically set to values such as "ansi", "utf-8",
// "windows-1252", etc. A list of supported character encodings is found on this
// page:Supported Charsets
// <http://www.chilkatsoft.com/p/p_463.asp>.
func (z *SFtp) WriteFileText64s(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkSFtp_WriteFileText64s(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


