// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkFtp2.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewFtp2() *Ftp2 {
	return &Ftp2{C.CkFtp2_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Ftp2) DisposeFtp2() {
    if z != nil {
	C.CkFtp2_Dispose(z.ckObj)
	}
}




func (z *Ftp2) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkFtp2_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Ftp2) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkFtp2_setExternalProgress(z.ckObj,1)
}

func (z *Ftp2) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkFtp2_setExternalProgress(z.ckObj,1)
}

func (z *Ftp2) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkFtp2_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Ftp2) AbortCurrent() bool {
    v := int(C.CkFtp2_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Ftp2) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAbortCurrent(z.ckObj,v)
}

// property: Account
// Some FTP servers require an Account name in addition to login/password. This
// property can be set for those servers with this requirement.
func (z *Ftp2) Account() string {
    return C.GoString(C.CkFtp2_account(z.ckObj))
}
// property setter: Account
func (z *Ftp2) SetAccount(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putAccount(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ActivePortRangeEnd
// When Active (i.e. PORT) mode is used (opposite of Passive), the client-side is
// responsible for choosing a random port for each data connection. (Note: In the
// FTP protocol, each data transfer occurs on a separate TCP/IP connection.
// Commands are sent over the control channel (port 21 for non-SSL, port 990 for
// SSL).)
// 
// This property, along with ActivePortRangeStart, allows the client to specify a
// range of ports for data connections.
//
func (z *Ftp2) ActivePortRangeEnd() int {
    return int(C.CkFtp2_getActivePortRangeEnd(z.ckObj))
}
// property setter: ActivePortRangeEnd
func (z *Ftp2) SetActivePortRangeEnd(value int) {
    C.CkFtp2_putActivePortRangeEnd(z.ckObj,C.int(value))
}

// property: ActivePortRangeStart
// This property, along with ActivePortRangeEnd, allows the client to specify a
// range of ports for data connections when in Active mode.
func (z *Ftp2) ActivePortRangeStart() int {
    return int(C.CkFtp2_getActivePortRangeStart(z.ckObj))
}
// property setter: ActivePortRangeStart
func (z *Ftp2) SetActivePortRangeStart(value int) {
    C.CkFtp2_putActivePortRangeStart(z.ckObj,C.int(value))
}

// property: AllocateSize
// If set to a non-zero value, causes an ALLO command, with this size as the
// parameter, to be automatically sent when uploading files to an FTP server.
// 
// This command could be required by some servers to reserve sufficient storage
// space to accommodate the new file to be transferred.
//
func (z *Ftp2) AllocateSize() uint {
    return uint(C.CkFtp2_getAllocateSize(z.ckObj))
}
// property setter: AllocateSize
func (z *Ftp2) SetAllocateSize(value uint) {
    C.CkFtp2_putAllocateSize(z.ckObj,C.ulong(value))
}

// property: AllowMlsd
// If true, then uses the MLSD command to fetch directory listings when the FTP
// server supports MLSD. This property is true by default.
// 
// When MLSD is used, the GetPermissions method will return the "perm fact" for a
// given file or directory. This is a different format than the more commonly
// recognized UNIX permissions string. Note: MLSD provides more accurate and
// dependable file listings, especially for last-mod date/time information. If
// usage of the MLSD command is turned off, it may adversely affect the quality and
// availability of other information.
//
func (z *Ftp2) AllowMlsd() bool {
    v := int(C.CkFtp2_getAllowMlsd(z.ckObj))
    return v != 0
}
// property setter: AllowMlsd
func (z *Ftp2) SetAllowMlsd(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAllowMlsd(z.ckObj,v)
}

// property: AsyncBytesReceived
// The number of bytes received during an asynchronous FTP download. This property
// is updated in real-time and an application may periodically fetch and display
// it's value while the download is in progress.
func (z *Ftp2) AsyncBytesReceived() uint {
    return uint(C.CkFtp2_getAsyncBytesReceived(z.ckObj))
}

// property: AsyncBytesReceived64
// Same as AsyncBytesReceived, but returns the value as a 64-bit integer.
func (z *Ftp2) AsyncBytesReceived64() int64 {
    return int64(C.CkFtp2_getAsyncBytesReceived64(z.ckObj))
}

// property: AsyncBytesSent
// The number of bytes sent during an asynchronous FTP upload. This property is
// updated in real-time and an application may periodically fetch and display it's
// value while the upload is in progress.
func (z *Ftp2) AsyncBytesSent() uint {
    return uint(C.CkFtp2_getAsyncBytesSent(z.ckObj))
}

// property: AsyncBytesSent64
// Same as AsyncBytesSent, but returns the value as a 64-bit integer.
func (z *Ftp2) AsyncBytesSent64() int64 {
    return int64(C.CkFtp2_getAsyncBytesSent64(z.ckObj))
}

// property: AuthSsl
// Same as AuthTls, except the command sent to the FTP server is "AUTH SSL" instead
// of "AUTH TLS". Most FTP servers accept either. AuthTls is more commonly used. If
// a particular server has trouble with AuthTls, try AuthSsl instead.
func (z *Ftp2) AuthSsl() bool {
    v := int(C.CkFtp2_getAuthSsl(z.ckObj))
    return v != 0
}
// property setter: AuthSsl
func (z *Ftp2) SetAuthSsl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAuthSsl(z.ckObj,v)
}

// property: AuthTls
// Set this to true to switch to a TLS encrypted channel. This property should be
// set prior to connecting. If this property is set, the port typically remains at
// it's default (21) and the Ssl property should *not* be set. When AuthTls is
// used, all control and data transmissions are encrypted. If your FTP client is
// behind a network-address-translating router, you may need to call
// ClearControlChannel after connecting and authenticating (i.e. after calling the
// Connect method). This keeps all data transmissions encrypted, but clears the
// control channel so that commands are sent unencrypted, thus allowing the router
// to translate network IP numbers in FTP commands.
func (z *Ftp2) AuthTls() bool {
    v := int(C.CkFtp2_getAuthTls(z.ckObj))
    return v != 0
}
// property setter: AuthTls
func (z *Ftp2) SetAuthTls(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAuthTls(z.ckObj,v)
}

// property: AutoFeat
// When true (which is the default value), a "FEAT" command is automatically sent
// to the FTP server immediately after connecting. This allows the Chilkat FTP2
// component to know more about the server's capabilities and automatically adjust
// any applicable internal settings based on the response. In rare cases, some FTP
// servers reject the "FEAT" command and close the connection. Usually, if an FTP
// server does not implement FEAT, a harmless "command not understood" response is
// returned.
// 
// Set this property to false to prevent the FEAT command from being sent.
//
func (z *Ftp2) AutoFeat() bool {
    v := int(C.CkFtp2_getAutoFeat(z.ckObj))
    return v != 0
}
// property setter: AutoFeat
func (z *Ftp2) SetAutoFeat(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAutoFeat(z.ckObj,v)
}

// property: AutoFix
// If true, then the following will occur when a connection is made to an FTP
// server:
// 
// 1) If the Port property = 990, then sets AuthTls = false, AuthSsl = false,
// and Ssl = true
// 2) If the Port property = 21, sets Ssl = false
// 
// The default value of this property is true.
//
func (z *Ftp2) AutoFix() bool {
    v := int(C.CkFtp2_getAutoFix(z.ckObj))
    return v != 0
}
// property setter: AutoFix
func (z *Ftp2) SetAutoFix(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAutoFix(z.ckObj,v)
}

// property: AutoGetSizeForProgress
// Forces the component to retrieve each file's size prior to downloading for the
// purpose of monitoring percentage completion progress. For many FTP servers, this
// is not required and therefore for performance reasons this property defaults to
// false.
func (z *Ftp2) AutoGetSizeForProgress() bool {
    v := int(C.CkFtp2_getAutoGetSizeForProgress(z.ckObj))
    return v != 0
}
// property setter: AutoGetSizeForProgress
func (z *Ftp2) SetAutoGetSizeForProgress(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAutoGetSizeForProgress(z.ckObj,v)
}

// property: AutoOptsUtf8
// When true (which is the default value), then an "OPTS UTF8 ON" command is
// automatically sent when connecting/authenticating if it is discovered via the
// FEAT command that the UTF8 option is supported.
// 
// Set this property to false to prevent the "OPTS UTF8 ON" command from being
// sent.
//
func (z *Ftp2) AutoOptsUtf8() bool {
    v := int(C.CkFtp2_getAutoOptsUtf8(z.ckObj))
    return v != 0
}
// property setter: AutoOptsUtf8
func (z *Ftp2) SetAutoOptsUtf8(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAutoOptsUtf8(z.ckObj,v)
}

// property: AutoSetUseEpsv
// If true then the UseEpsv property is automatically set upon connecting to the
// FTP server. The default value of this property is false.
// 
// If the AutoFeat property is true, and if the AutoSetUseEpsv property is
// true, then the FTP server's features are automatically queried when
// connecting. In this case, the UseEpsv property is automatically set to true if
// the FTP server supports EPSV.
// 
// Important: EPSV can cause problems with some deep-inspection firewalls. If a
// passive data connection cannot be established, make sure to test with both the
// AutoSetUseEpsv and UseEpsv properties set equal to false.
//
func (z *Ftp2) AutoSetUseEpsv() bool {
    v := int(C.CkFtp2_getAutoSetUseEpsv(z.ckObj))
    return v != 0
}
// property setter: AutoSetUseEpsv
func (z *Ftp2) SetAutoSetUseEpsv(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAutoSetUseEpsv(z.ckObj,v)
}

// property: AutoSyst
// When true (which is the default value), a "SYST" command is automatically sent
// to the FTP server immediately after connecting. This allows the Chilkat FTP2
// component to know more about the server and automatically adjust any applicable
// internal settings based on the response. If the SYST command causes trouble
// (which is rare), this behavior can be turned off by setting this property equal
// to false.
func (z *Ftp2) AutoSyst() bool {
    v := int(C.CkFtp2_getAutoSyst(z.ckObj))
    return v != 0
}
// property setter: AutoSyst
func (z *Ftp2) SetAutoSyst(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAutoSyst(z.ckObj,v)
}

// property: AutoXcrc
// Many FTP servers support the XCRC command. The Chilkat FTP component will
// automatically know if XCRC is supported because it automatically sends a FEAT
// command to the server immediately after connecting.
// 
// If this property is set to true, then all uploads will be automatically
// verified by sending an XCRC command immediately after the transfer completes. If
// the CRC is not verified, the upload method (such as PutFile) will return a
// failed status.
// 
// To prevent XCRC checking, set this property to false.
//
func (z *Ftp2) AutoXcrc() bool {
    v := int(C.CkFtp2_getAutoXcrc(z.ckObj))
    return v != 0
}
// property setter: AutoXcrc
func (z *Ftp2) SetAutoXcrc(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putAutoXcrc(z.ckObj,v)
}

// property: BandwidthThrottleDown
// If set to a non-zero value, the FTP2 component will bandwidth throttle all
// downloads to this value.
// 
// The default value of this property is 0. The value should be specified in
// bytes/second.
// 
// Note: It is difficult to throttle very small downloads. (For example, how do you
// bandwidth throttle a 1-byte download???) As the downloaded file size gets
// larger, the transfer rate will better approximate this property's setting.
// 
// Also note: When downloading, the FTP server has no knowledge of the client's
// desire for throttling, and is always sending data as fast as possible. (There's
// nothing in the FTP protocol to request throttling.) Therefore, any throttling
// for a download on the client side is simply to allow system socket buffers
// (outgoing buffers on the sender, and incoming buffers on the client) to fill to
// 100% capacity, and this also poses the threat of causing a data connection to
// become broken. It's probably not worthwhile to attempt to throttle downloads. It
// may have been better that this property never existed.
//
func (z *Ftp2) BandwidthThrottleDown() int {
    return int(C.CkFtp2_getBandwidthThrottleDown(z.ckObj))
}
// property setter: BandwidthThrottleDown
func (z *Ftp2) SetBandwidthThrottleDown(value int) {
    C.CkFtp2_putBandwidthThrottleDown(z.ckObj,C.int(value))
}

// property: BandwidthThrottleUp
// If set to a non-zero value, the FTP2 component will bandwidth throttle all
// uploads to this value.
// 
// The default value of this property is 0. The value should be specified in
// bytes/second.
// 
// Note: It is difficult to throttle very small uploads. (For example, how do you
// bandwidth throttle a 1-byte upload???) As the uploaded file size gets larger,
// the transfer rate will better approximate this property's setting.
//
func (z *Ftp2) BandwidthThrottleUp() int {
    return int(C.CkFtp2_getBandwidthThrottleUp(z.ckObj))
}
// property setter: BandwidthThrottleUp
func (z *Ftp2) SetBandwidthThrottleUp(value int) {
    C.CkFtp2_putBandwidthThrottleUp(z.ckObj,C.int(value))
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
func (z *Ftp2) ClientIpAddress() string {
    return C.GoString(C.CkFtp2_clientIpAddress(z.ckObj))
}
// property setter: ClientIpAddress
func (z *Ftp2) SetClientIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putClientIpAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CommandCharset
// Indicates the charset to be used for commands sent to the FTP server. The
// command charset must match what the FTP server is expecting in order to
// communicate non-English characters correctly. The default value of this property
// is "ansi".
// 
// This property may be updated to "utf-8" after connecting because a "FEAT"
// command is automatically sent to get the features of the FTP server. If UTF8 is
// indicated as a feature, then this property is automatically changed to "utf-8".
//
func (z *Ftp2) CommandCharset() string {
    return C.GoString(C.CkFtp2_commandCharset(z.ckObj))
}
// property setter: CommandCharset
func (z *Ftp2) SetCommandCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putCommandCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConnectFailReason
// If the Connect method fails, this property can be checked to determine the
// reason for failure.
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
// FTP:
// 200 = Connected, but failed to receive greeting from FTP server.
// 201 = Failed to do AUTH TLS or AUTH SSL.
// Protocol/Component:
// 300 = asynch op in progress
// 301 = login failure.
//
func (z *Ftp2) ConnectFailReason() int {
    return int(C.CkFtp2_getConnectFailReason(z.ckObj))
}

// property: ConnectTimeout
// Maximum number of seconds to wait when connecting to an FTP server. The default
// is 30 seconds. A value of 0 indicates the willingness to wait forever.
func (z *Ftp2) ConnectTimeout() int {
    return int(C.CkFtp2_getConnectTimeout(z.ckObj))
}
// property setter: ConnectTimeout
func (z *Ftp2) SetConnectTimeout(value int) {
    C.CkFtp2_putConnectTimeout(z.ckObj,C.int(value))
}

// property: ConnectVerified
// True if the FTP2 component was able to establish a TCP/IP connection to the FTP
// server after calling Connect.
func (z *Ftp2) ConnectVerified() bool {
    v := int(C.CkFtp2_getConnectVerified(z.ckObj))
    return v != 0
}

// property: CrlfMode
// Used to control CRLF line endings when downloading text files in ASCII mode. The
// default value is 0.
// 
// Possible values are:
// 0 = Do nothing.  The line-endings are not modified as received from the FTP server.
// 1 = Convert all line-endings to CR+LF
// 2 = Convert all line-endings to bare LF's
// 3 = Convert all line-endings to bare CR's
//
func (z *Ftp2) CrlfMode() int {
    return int(C.CkFtp2_getCrlfMode(z.ckObj))
}
// property setter: CrlfMode
func (z *Ftp2) SetCrlfMode(value int) {
    C.CkFtp2_putCrlfMode(z.ckObj,C.int(value))
}

// property: DataProtection
// Controls the data protection level for the data connections. Possible values are
// "control", "clear", or "private".
//     "control" is the default, and the data connections will be the same as for
//     the control connection. If the control connection is SSL/TLS, then the data
//     connections are also SSL/TLS. If the control connection is unencrypted, then the
//     data connections will also be unencrypted.
//     "clear" means that the data connections will always be unencrypted (TCP
//     without TLS).
//     "private" means that the data connections will always be encrypted (TLS).
func (z *Ftp2) DataProtection() string {
    return C.GoString(C.CkFtp2_dataProtection(z.ckObj))
}
// property setter: DataProtection
func (z *Ftp2) SetDataProtection(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putDataProtection(z.ckObj,cStr)
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
func (z *Ftp2) DebugLogFilePath() string {
    return C.GoString(C.CkFtp2_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Ftp2) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DirListingCharset
// Indicates the charset of the directory listings received from the FTP server.
// The FTP2 client must interpret the directory listing bytes using the correct
// character encoding in order to correctly receive non-English characters. The
// default value of this property is "ansi".
// 
// This property may be updated to "utf-8" after connecting because a "FEAT"
// command is automatically sent to get the features of the FTP server. If UTF8 is
// indicated as a feature, then this property is automatically changed to "utf-8".
//
func (z *Ftp2) DirListingCharset() string {
    return C.GoString(C.CkFtp2_dirListingCharset(z.ckObj))
}
// property setter: DirListingCharset
func (z *Ftp2) SetDirListingCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putDirListingCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DownloadTransferRate
// The average download rate in bytes/second. This property is updated in real-time
// during any FTP download (asynchronous or synchronous).
func (z *Ftp2) DownloadTransferRate() int {
    return int(C.CkFtp2_getDownloadTransferRate(z.ckObj))
}

// property: ForcePortIpAddress
// If set, forces the IP address used in the PORT command for Active mode (i.e.
// non-passive) data transfers. This string property should be set to the IP
// address in dotted notation, such as "233.190.65.31".
// 
// Note: This property can also be set to the special keyword "control" to force
// the PORT IP address to be the address of the control connection's peer.
// 
// Starting in v9.5.0.58, the IP address can be prefixed with the string "bind-".
// For example, "bind-233.190.65.31". When "bind-" is specified, the local data
// socket will be bound to the IP address when created. Otherwise, the IP address
// is only used as the argument to the PORT command that is sent to the server.
//
func (z *Ftp2) ForcePortIpAddress() string {
    return C.GoString(C.CkFtp2_forcePortIpAddress(z.ckObj))
}
// property setter: ForcePortIpAddress
func (z *Ftp2) SetForcePortIpAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putForcePortIpAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Greeting
// The initial greeting received from the FTP server after connecting.
func (z *Ftp2) Greeting() string {
    return C.GoString(C.CkFtp2_greeting(z.ckObj))
}

// property: HasModeZ
// Chilkat FTP2 supports MODE Z, which is a transfer mode implemented by some FTP
// servers. It allows for files to be uploaded and downloaded using compressed
// streams (using the zlib deflate algorithm). This is a read-only property. It
// will be set to true if the FTP2 component detects that your FTP server
// supports MODE Z. Otherwise it is set to false.
func (z *Ftp2) HasModeZ() bool {
    v := int(C.CkFtp2_getHasModeZ(z.ckObj))
    return v != 0
}

// property: HeartbeatMs
// This is the number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any FTP operation prior to
// completion. If HeartbeatMs is 0, no AbortCheck event callbacks will occur. Also,
// AbortCheck callbacks do not occur when doing asynchronous transfers.
func (z *Ftp2) HeartbeatMs() int {
    return int(C.CkFtp2_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Ftp2) SetHeartbeatMs(value int) {
    C.CkFtp2_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: Hostname
// The domain name of the FTP server. May also use the IPv4 or IPv6 address in
// string format.
func (z *Ftp2) Hostname() string {
    return C.GoString(C.CkFtp2_hostname(z.ckObj))
}
// property setter: Hostname
func (z *Ftp2) SetHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyAuthMethod
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy authentication method name. Valid choices are "Basic" or "NTLM".
func (z *Ftp2) HttpProxyAuthMethod() string {
    return C.GoString(C.CkFtp2_httpProxyAuthMethod(z.ckObj))
}
// property setter: HttpProxyAuthMethod
func (z *Ftp2) SetHttpProxyAuthMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putHttpProxyAuthMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyDomain
// If an HTTP proxy is used, and it uses NTLM authentication, then this optional
// property is the NTLM authentication domain.
func (z *Ftp2) HttpProxyDomain() string {
    return C.GoString(C.CkFtp2_httpProxyDomain(z.ckObj))
}
// property setter: HttpProxyDomain
func (z *Ftp2) SetHttpProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putHttpProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyHostname
// If an HTTP proxy is to be used, set this property to the HTTP proxy hostname or
// IPv4 address (in dotted decimal notation).
func (z *Ftp2) HttpProxyHostname() string {
    return C.GoString(C.CkFtp2_httpProxyHostname(z.ckObj))
}
// property setter: HttpProxyHostname
func (z *Ftp2) SetHttpProxyHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putHttpProxyHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPassword
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy password.
func (z *Ftp2) HttpProxyPassword() string {
    return C.GoString(C.CkFtp2_httpProxyPassword(z.ckObj))
}
// property setter: HttpProxyPassword
func (z *Ftp2) SetHttpProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putHttpProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpProxyPort
// If an HTTP proxy is to be used, set this property to the HTTP proxy port number.
// (Two commonly used HTTP proxy ports are 8080 and 3128.)
func (z *Ftp2) HttpProxyPort() int {
    return int(C.CkFtp2_getHttpProxyPort(z.ckObj))
}
// property setter: HttpProxyPort
func (z *Ftp2) SetHttpProxyPort(value int) {
    C.CkFtp2_putHttpProxyPort(z.ckObj,C.int(value))
}

// property: HttpProxyUsername
// If an HTTP proxy requiring authentication is to be used, set this property to
// the HTTP proxy login name.
func (z *Ftp2) HttpProxyUsername() string {
    return C.GoString(C.CkFtp2_httpProxyUsername(z.ckObj))
}
// property setter: HttpProxyUsername
func (z *Ftp2) SetHttpProxyUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putHttpProxyUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IdleTimeoutMs
// Forces a timeout when a response is expected on the control channel, but no
// response arrives for this number of milliseconds. Setting IdleTimeoutMs = 0
// allows the application to wait indefinitely. The default value is 60000 (i.e. 60
// seconds).
func (z *Ftp2) IdleTimeoutMs() int {
    return int(C.CkFtp2_getIdleTimeoutMs(z.ckObj))
}
// property setter: IdleTimeoutMs
func (z *Ftp2) SetIdleTimeoutMs(value int) {
    C.CkFtp2_putIdleTimeoutMs(z.ckObj,C.int(value))
}

// property: IsConnected
// Important: This property is deprecated. Applications should instead call the
// CheckConnection method.
// 
// Returns true if currently connected and logged into an FTP server, otherwise
// returns false.
// 
// Note: Accessing this property may cause a NOOP command to be sent to the FTP
// server.
//
func (z *Ftp2) IsConnected() bool {
    v := int(C.CkFtp2_getIsConnected(z.ckObj))
    return v != 0
}

// property: KeepSessionLog
// Turns the in-memory session logging on or off. If on, the session log can be
// obtained via the SessionLog property.
func (z *Ftp2) KeepSessionLog() bool {
    v := int(C.CkFtp2_getKeepSessionLog(z.ckObj))
    return v != 0
}
// property setter: KeepSessionLog
func (z *Ftp2) SetKeepSessionLog(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putKeepSessionLog(z.ckObj,v)
}

// property: LargeFileMeasures
// Enables internal features that can help when downloading extremely large files.
// In some cases, if the time required to download a file is long, the control
// connection is closed by the server or other network infrastructure because it
// was idle for so long. Setting this property equal to true will keep the
// control connection very slightly used to prevent this from happening.
// 
// The default value of this property is false. This property should only be set
// to true if this sort of problem is encountered.
//
func (z *Ftp2) LargeFileMeasures() bool {
    v := int(C.CkFtp2_getLargeFileMeasures(z.ckObj))
    return v != 0
}
// property setter: LargeFileMeasures
func (z *Ftp2) SetLargeFileMeasures(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putLargeFileMeasures(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Ftp2) LastErrorHtml() string {
    return C.GoString(C.CkFtp2_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Ftp2) LastErrorText() string {
    return C.GoString(C.CkFtp2_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Ftp2) LastErrorXml() string {
    return C.GoString(C.CkFtp2_lastErrorXml(z.ckObj))
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
func (z *Ftp2) LastMethodSuccess() bool {
    v := int(C.CkFtp2_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Ftp2) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putLastMethodSuccess(z.ckObj,v)
}

// property: LastReply
// Contains the last control-channel reply. For example: "550 Failed to change
// directory." or "250 Directory successfully changed." The control channel reply
// is typically formatted as an integer status code followed by a one-line
// description.
func (z *Ftp2) LastReply() string {
    return C.GoString(C.CkFtp2_lastReply(z.ckObj))
}

// property: ListPattern
// A wildcard pattern, defaulting to "*" that determines the files and directories
// included in the following properties and methods: GetDirCount, GetCreateTime,
// GetFilename, GetIsDirectory, GetLastAccessTime, GetModifiedTime, GetSize.
// 
// Note: Do not include a directory path in the ListPattern. For example, do not
// set the ListPattern equal to a string such as this: "subdir/*.txt". The correct
// solution is to first change the remote directory to "subdir" by calling
// ChangeRemoteDir, and then set the ListPattern equal to "*.txt".
//
func (z *Ftp2) ListPattern() string {
    return C.GoString(C.CkFtp2_listPattern(z.ckObj))
}
// property setter: ListPattern
func (z *Ftp2) SetListPattern(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putListPattern(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LoginVerified
// True if the FTP2 component was able to login to the FTP server after calling
// Connect.
func (z *Ftp2) LoginVerified() bool {
    v := int(C.CkFtp2_getLoginVerified(z.ckObj))
    return v != 0
}

// property: PartialTransfer
// A read-only property that indicates whether a partial transfer was received in
// the last method call to download a file. Set to true if a partial transfer was
// received. Set to false if nothing was received, or if the full file was
// received.
func (z *Ftp2) PartialTransfer() bool {
    v := int(C.CkFtp2_getPartialTransfer(z.ckObj))
    return v != 0
}

// property: Passive
// Set to true for FTP to operate in passive mode, otherwise set to false for
// non-passive (.i.e. "active" or "port" mode). The default value of this property
// is true.
func (z *Ftp2) Passive() bool {
    v := int(C.CkFtp2_getPassive(z.ckObj))
    return v != 0
}
// property setter: Passive
func (z *Ftp2) SetPassive(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putPassive(z.ckObj,v)
}

// property: PassiveUseHostAddr
// This can handle problems that may arise when an FTP server is located behind a
// NAT router. FTP servers respond to the PASV command by sending the IP address
// and port where it will be listening for the data connection. If the control
// connection is SSL encrypted, the NAT router is not able to convert from an
// internal IP address (typically beginning with 192.168) to an external address.
// When set to true, PassiveUseHostAddr property tells the FTP client to discard
// the IP address part of the PASV response and replace it with the IP address of
// the already-established control connection. The default value of this property
// is false.
func (z *Ftp2) PassiveUseHostAddr() bool {
    v := int(C.CkFtp2_getPassiveUseHostAddr(z.ckObj))
    return v != 0
}
// property setter: PassiveUseHostAddr
func (z *Ftp2) SetPassiveUseHostAddr(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putPassiveUseHostAddr(z.ckObj,v)
}

// property: Password
// Password for logging into the FTP server.
func (z *Ftp2) Password() string {
    return C.GoString(C.CkFtp2_password(z.ckObj))
}
// property setter: Password
func (z *Ftp2) SetPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putPassword(z.ckObj,cStr)
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
func (z *Ftp2) PercentDoneScale() int {
    return int(C.CkFtp2_getPercentDoneScale(z.ckObj))
}
// property setter: PercentDoneScale
func (z *Ftp2) SetPercentDoneScale(value int) {
    C.CkFtp2_putPercentDoneScale(z.ckObj,C.int(value))
}

// property: Port
// Port number. Automatically defaults to the default port for the FTP service.
func (z *Ftp2) Port() int {
    return int(C.CkFtp2_getPort(z.ckObj))
}
// property setter: Port
func (z *Ftp2) SetPort(value int) {
    C.CkFtp2_putPort(z.ckObj,C.int(value))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *Ftp2) PreferIpv6() bool {
    v := int(C.CkFtp2_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *Ftp2) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putPreferIpv6(z.ckObj,v)
}

// property: PreferNlst
// If true, the NLST command is used instead of LIST when fetching a directory
// listing. This can help in very rare cases where the FTP server returns truncated
// filenames. The drawback to using NLST is that it won't return size or date/time
// info (but it should return the full filename).
// 
// The default value of this property is false.
//
func (z *Ftp2) PreferNlst() bool {
    v := int(C.CkFtp2_getPreferNlst(z.ckObj))
    return v != 0
}
// property setter: PreferNlst
func (z *Ftp2) SetPreferNlst(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putPreferNlst(z.ckObj,v)
}

// property: ProgressMonSize
// Progress monitoring for FTP downloads rely on the FTP server indicating the file
// size within the RETR response. Some FTP servers however, do not indicate the
// file size and therefore it is not possible to monitor progress based on
// percentage completion. This property allows the application to explicitly tell
// the FTP component the size of the file about to be downloaded for the next
// GetFile call.
func (z *Ftp2) ProgressMonSize() int {
    return int(C.CkFtp2_getProgressMonSize(z.ckObj))
}
// property setter: ProgressMonSize
func (z *Ftp2) SetProgressMonSize(value int) {
    C.CkFtp2_putProgressMonSize(z.ckObj,C.int(value))
}

// property: ProgressMonSize64
// Same as ProgressMonSize, but allows for sizes greater than the 32-bit integer
// limit.
func (z *Ftp2) ProgressMonSize64() int64 {
    return int64(C.CkFtp2_getProgressMonSize64(z.ckObj))
}
// property setter: ProgressMonSize64
func (z *Ftp2) SetProgressMonSize64(value int64) {
    C.CkFtp2_putProgressMonSize64(z.ckObj,C.__int64(value))
}

// property: ProxyHostname
// The hostname of your FTP proxy, if a proxy server is used.
func (z *Ftp2) ProxyHostname() string {
    return C.GoString(C.CkFtp2_proxyHostname(z.ckObj))
}
// property setter: ProxyHostname
func (z *Ftp2) SetProxyHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putProxyHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyMethod
// The proxy scheme used by your FTP proxy server. Valid values are 0 to 9. The
// default value is 0 which indicates that no proxy server is used. Supported proxy
// methods are as follows:
// 
// Note: The ProxyHostname is the hostname of the firewall, if the proxy is a
// firewall. Also, the ProxyUsername and ProxyPassword are the firewall
// username/password (if the proxy is a firewall).
// 
//     ProxyMethod = 1 (SITE site)
// 
//         USER ProxyUsername
//         PASS ProxyPassword
//         SITE Hostname
//         USER Username
//         PASS Password
// 
//     ProxyMethod = 2 (USER user@site)
// 
//         USER Username@Hostname:Port
//         PASS Password
// 
//     ProxyMethod = 3 (USER with login)
// 
//         USER ProxyUsername
//         PASS ProxyPassword
//         USER Username@Hostname:Port
//         PASS Password
// 
//     ProxyMethod = 4 (USER/PASS/ACCT)
// 
//         USER Username@Hostname:Port ProxyUsername
//         PASS Password
//         ACCT ProxyPassword
// 
//     ProxyMethod = 5 (OPEN site)
// 
//         USER ProxyUsername
//         PASS ProxyPassword
//         OPEN Hostname
//         USER Username
//         PASS Password
// 
//     ProxyMethod = 6 (firewallId@site)
// 
//         USER ProxyUsername@Hostname
//         USER Username
//         PASS Password
// 
//     ProxyMethod = 7
// 
//         USER ProxyUsername
//         USER ProxyPassword
//         SITE Hostname:Port USER Username
//         PASS Password
// 
//     ProxyMethod = 8
// 
//         USER Username@ProxyUsername@Hostname
//         PASS Password@ProxyPassword
// 
//     ProxyMethod = 9
// 
//         ProxyUsername ProxyPassword Username Password
//
func (z *Ftp2) ProxyMethod() int {
    return int(C.CkFtp2_getProxyMethod(z.ckObj))
}
// property setter: ProxyMethod
func (z *Ftp2) SetProxyMethod(value int) {
    C.CkFtp2_putProxyMethod(z.ckObj,C.int(value))
}

// property: ProxyPassword
// The password for authenticating with the FTP proxy server.
func (z *Ftp2) ProxyPassword() string {
    return C.GoString(C.CkFtp2_proxyPassword(z.ckObj))
}
// property setter: ProxyPassword
func (z *Ftp2) SetProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyPort
// If an FTP proxy server is used, this is the port number at which the proxy
// server is listening for connections.
func (z *Ftp2) ProxyPort() int {
    return int(C.CkFtp2_getProxyPort(z.ckObj))
}
// property setter: ProxyPort
func (z *Ftp2) SetProxyPort(value int) {
    C.CkFtp2_putProxyPort(z.ckObj,C.int(value))
}

// property: ProxyUsername
// The username for authenticating with the FTP proxy server.
func (z *Ftp2) ProxyUsername() string {
    return C.GoString(C.CkFtp2_proxyUsername(z.ckObj))
}
// property setter: ProxyUsername
func (z *Ftp2) SetProxyUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putProxyUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ReadTimeout
// Forces a timeout when incoming data is expected on a data channel, but no data
// arrives for this number of seconds. The ReadTimeout is the amount of time that
// needs to elapse while no additional data is forthcoming. During a long download,
// if the data stream halts for more than this amount, it will timeout. Otherwise,
// there is no limit on the length of time for the entire download.
// 
// The default value is 60.
//
func (z *Ftp2) ReadTimeout() int {
    return int(C.CkFtp2_getReadTimeout(z.ckObj))
}
// property setter: ReadTimeout
func (z *Ftp2) SetReadTimeout(value int) {
    C.CkFtp2_putReadTimeout(z.ckObj,C.int(value))
}

// property: RequireSslCertVerify
// If true, then the FTP2 client will verify the server's SSL certificate. The
// server's certificate signature is verified with its issuer, and the issuer's
// cert is verified with its issuer, etc. up to the root CA cert. If a signature
// verification fails, the connection is not allowed. Also, if the certificate is
// expired, or if the cert's signature is invalid, the connection is not allowed.
// The default value of this property is false.
func (z *Ftp2) RequireSslCertVerify() bool {
    v := int(C.CkFtp2_getRequireSslCertVerify(z.ckObj))
    return v != 0
}
// property setter: RequireSslCertVerify
func (z *Ftp2) SetRequireSslCertVerify(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putRequireSslCertVerify(z.ckObj,v)
}

// property: RestartNext
// Both uploads and downloads may be resumed by simply setting this property =
// true and re-calling the upload or download method.
func (z *Ftp2) RestartNext() bool {
    v := int(C.CkFtp2_getRestartNext(z.ckObj))
    return v != 0
}
// property setter: RestartNext
func (z *Ftp2) SetRestartNext(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putRestartNext(z.ckObj,v)
}

// property: SessionLog
// Contains the session log if KeepSessionLog is turned on.
func (z *Ftp2) SessionLog() string {
    return C.GoString(C.CkFtp2_sessionLog(z.ckObj))
}

// property: SocksHostname
// The SOCKS4/SOCKS5 hostname or IPv4 address (in dotted decimal notation). This
// property is only used if the SocksVersion property is set to 4 or 5).
func (z *Ftp2) SocksHostname() string {
    return C.GoString(C.CkFtp2_socksHostname(z.ckObj))
}
// property setter: SocksHostname
func (z *Ftp2) SetSocksHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSocksHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPassword
// The SOCKS5 password (if required). The SOCKS4 protocol does not include the use
// of a password, so this does not apply to SOCKS4.
func (z *Ftp2) SocksPassword() string {
    return C.GoString(C.CkFtp2_socksPassword(z.ckObj))
}
// property setter: SocksPassword
func (z *Ftp2) SetSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPort
// The SOCKS4/SOCKS5 proxy port. The default value is 1080. This property only
// applies if a SOCKS proxy is used (if the SocksVersion property is set to 4 or
// 5).
func (z *Ftp2) SocksPort() int {
    return int(C.CkFtp2_getSocksPort(z.ckObj))
}
// property setter: SocksPort
func (z *Ftp2) SetSocksPort(value int) {
    C.CkFtp2_putSocksPort(z.ckObj,C.int(value))
}

// property: SocksUsername
// The SOCKS4/SOCKS5 proxy username. This property is only used if the SocksVersion
// property is set to 4 or 5).
func (z *Ftp2) SocksUsername() string {
    return C.GoString(C.CkFtp2_socksUsername(z.ckObj))
}
// property setter: SocksUsername
func (z *Ftp2) SetSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksVersion
// SocksVersion May be set to one of the following integer values:
// 
// 0 - No SOCKS proxy is used. This is the default.
// 4 - Connect via a SOCKS4 proxy.
// 5 - Connect via a SOCKS5 proxy.
//
func (z *Ftp2) SocksVersion() int {
    return int(C.CkFtp2_getSocksVersion(z.ckObj))
}
// property setter: SocksVersion
func (z *Ftp2) SetSocksVersion(value int) {
    C.CkFtp2_putSocksVersion(z.ckObj,C.int(value))
}

// property: SoRcvBuf
// Sets the receive buffer size socket option. Normally, this property should be
// left unchanged. The default value is 4194304.
// 
// This property can be increased if download performance seems slow. It is
// recommended to be a multiple of 4096.
// 
// Note: This property only applies to FTP data connections. The FTP control
// connection is not used for uploading or downloading files, and is therefore not
// performance sensitive.
//
func (z *Ftp2) SoRcvBuf() int {
    return int(C.CkFtp2_getSoRcvBuf(z.ckObj))
}
// property setter: SoRcvBuf
func (z *Ftp2) SetSoRcvBuf(value int) {
    C.CkFtp2_putSoRcvBuf(z.ckObj,C.int(value))
}

// property: SoSndBuf
// Sets the send buffer size socket option. Normally, this property should be left
// unchanged. The default value is 262144.
// 
// This property can be increased if upload performance seems slow. It is
// recommended to be a multiple of 4096. Testing with sizes such as 512K and 1MB is
// reasonable.
// 
// Note: This property only applies to FTP data connections. The FTP control
// connection is not used for uploading or downloading files, and is therefore not
// performance sensitive.
//
func (z *Ftp2) SoSndBuf() int {
    return int(C.CkFtp2_getSoSndBuf(z.ckObj))
}
// property setter: SoSndBuf
func (z *Ftp2) SetSoSndBuf(value int) {
    C.CkFtp2_putSoSndBuf(z.ckObj,C.int(value))
}

// property: Ssl
// Use TLS/SSL for FTP connections. You would typically set Ssl = true when
// connecting to port 990 on FTP servers that support TLS/SSL mode. Note: It is
// more common to use AuthTls.
func (z *Ftp2) Ssl() bool {
    v := int(C.CkFtp2_getSsl(z.ckObj))
    return v != 0
}
// property setter: Ssl
func (z *Ftp2) SetSsl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putSsl(z.ckObj,v)
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
func (z *Ftp2) SslAllowedCiphers() string {
    return C.GoString(C.CkFtp2_sslAllowedCiphers(z.ckObj))
}
// property setter: SslAllowedCiphers
func (z *Ftp2) SetSslAllowedCiphers(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSslAllowedCiphers(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SslProtocol
// Selects the secure protocol to be used for secure (SSL/TLS) implicit and
// explicit (AUTH TLS / AUTH SSL) connections . Possible values are:
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
func (z *Ftp2) SslProtocol() string {
    return C.GoString(C.CkFtp2_sslProtocol(z.ckObj))
}
// property setter: SslProtocol
func (z *Ftp2) SetSslProtocol(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSslProtocol(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SslServerCertVerified
// Read-only property that returns true if the FTP server's digital certificate
// was verified when connecting via SSL / TLS.
func (z *Ftp2) SslServerCertVerified() bool {
    v := int(C.CkFtp2_getSslServerCertVerified(z.ckObj))
    return v != 0
}

// property: SyncCreateAllLocalDirs
// If true, then empty directories on the server are created locally when doing a
// download synchronization. If false, then only directories containing files
// that are downloaded are auto-created.
// 
// The default value of this property is true.
//
func (z *Ftp2) SyncCreateAllLocalDirs() bool {
    v := int(C.CkFtp2_getSyncCreateAllLocalDirs(z.ckObj))
    return v != 0
}
// property setter: SyncCreateAllLocalDirs
func (z *Ftp2) SetSyncCreateAllLocalDirs(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putSyncCreateAllLocalDirs(z.ckObj,v)
}

// property: SyncedFiles
// The paths of the files uploaded or downloaded in the last call to
// SyncDeleteTree, SyncLocalDir, SyncLocalTree, SyncRemoteTree, or SyncRemoteTree2.
// The paths are listed one per line. In both cases (for upload and download) each
// line contains the paths relative to the root synced directory.
func (z *Ftp2) SyncedFiles() string {
    return C.GoString(C.CkFtp2_syncedFiles(z.ckObj))
}
// property setter: SyncedFiles
func (z *Ftp2) SetSyncedFiles(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSyncedFiles(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustMatch
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "*.xml; *.txt; *.csv". If set, the Sync* upload and download methods
// will only transfer files that match any one of these patterns. Pattern matching
// is case-insensitive.
// 
// Note: Starting in version 9.5.0.47, this property also applies to the
// DownloadTree and DirTreeXml methods.
//
func (z *Ftp2) SyncMustMatch() string {
    return C.GoString(C.CkFtp2_syncMustMatch(z.ckObj))
}
// property setter: SyncMustMatch
func (z *Ftp2) SetSyncMustMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSyncMustMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustMatchDir
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "xml; txt; data_*". If set, the Sync* upload and download methods will
// only enter directories that match any one of these patterns. Pattern matching is
// case-insensitive.
func (z *Ftp2) SyncMustMatchDir() string {
    return C.GoString(C.CkFtp2_syncMustMatchDir(z.ckObj))
}
// property setter: SyncMustMatchDir
func (z *Ftp2) SetSyncMustMatchDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSyncMustMatchDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustNotMatch
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "*.xml; *.txt; *.csv". If set, the Sync* upload and download methods
// will not transfer files that match any one of these patterns. Pattern matching
// is case-insensitive.
// 
// Note: Starting in version 9.5.0.47, this property also applies to the
// DownloadTree and DirTreeXml methods.
//
func (z *Ftp2) SyncMustNotMatch() string {
    return C.GoString(C.CkFtp2_syncMustNotMatch(z.ckObj))
}
// property setter: SyncMustNotMatch
func (z *Ftp2) SetSyncMustNotMatch(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSyncMustNotMatch(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncMustNotMatchDir
// Can contain a wildcarded list of file patterns separated by semicolons. For
// example, "xml; txt; data_*". If set, the Sync* upload and download methods will
// enter directories that match any one of these patterns. Pattern matching is
// case-insensitive.
func (z *Ftp2) SyncMustNotMatchDir() string {
    return C.GoString(C.CkFtp2_syncMustNotMatchDir(z.ckObj))
}
// property setter: SyncMustNotMatchDir
func (z *Ftp2) SetSyncMustNotMatchDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putSyncMustNotMatchDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SyncPreview
// Contains the list of files that would be transferred in a call to
// SyncRemoteTree2 when the previewOnly argument is set to true. This string
// property contains one filepath per line, separated by CRLF line endings. After
// SyncRemoteTree2 is called, this property contains the filepaths of the local
// files that would be uploaded to the FTP server.
func (z *Ftp2) SyncPreview() string {
    return C.GoString(C.CkFtp2_syncPreview(z.ckObj))
}

// property: TlsCipherSuite
// Contains the current or last negotiated TLS cipher suite. If no TLS connection
// has yet to be established, or if a connection as attempted and failed, then this
// will be empty. A sample cipher suite string looks like this:
// TLS_DHE_RSA_WITH_AES_256_CBC_SHA256.
func (z *Ftp2) TlsCipherSuite() string {
    return C.GoString(C.CkFtp2_tlsCipherSuite(z.ckObj))
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
func (z *Ftp2) TlsPinSet() string {
    return C.GoString(C.CkFtp2_tlsPinSet(z.ckObj))
}
// property setter: TlsPinSet
func (z *Ftp2) SetTlsPinSet(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putTlsPinSet(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TlsVersion
// Contains the current or last negotiated TLS protocol version. If no TLS
// connection has yet to be established, or if a connection as attempted and
// failed, then this will be empty. Possible values are "SSL 3.0", "TLS 1.0", "TLS
// 1.1", and "TLS 1.2".
func (z *Ftp2) TlsVersion() string {
    return C.GoString(C.CkFtp2_tlsVersion(z.ckObj))
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. This property
// defaults to the empty string and should typically remain empty.
// 
// Can be set to a list of the following comma separated keywords:
//     "OpenNonExclusive" - Introduced in v9.5.0.78. When downloading files on
//     Windows systems, open the local file with non-exclusive access to allow other
//     programs the ability to access the file as it's being downloaded.
//     "ProtectFromVpn" - Introduced in v9.5.0.80. On Android systems, will bypass
//     any VPN that may be installed or active.
//     "EnableTls13" - Introduced in v9.5.0.82. Causes TLS 1.3 to be offered in the
//     ClientHello of the TLS protocol, allowing the server to select TLS 1.3 for the
//     session. Future versions of Chilkat will enable TLS 1.3 by default. This option
//     is only necessary in v9.5.0.82 if TLS 1.3 is desired.
//
func (z *Ftp2) UncommonOptions() string {
    return C.GoString(C.CkFtp2_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Ftp2) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UploadTransferRate
// The average upload rate in bytes/second. This property is updated in real-time
// during any FTP upload (asynchronous or synchronous).
func (z *Ftp2) UploadTransferRate() int {
    return int(C.CkFtp2_getUploadTransferRate(z.ckObj))
}

// property: UseEpsv
// If true, the FTP2 component will use the EPSV command instead of PASV for
// passive mode data transfers. The default value of this property is false. (It
// is somewhat uncommon for FTP servers to support EPSV.)
// 
// Note: If the AutoFeat property is true, then the FTP server's features are
// automatically queried after connecting. In this case, if the AutoSetUseEpsv
// property is also set to true, the UseEpsv property is automatically set to
// true if the FTP server supports EPSV.
// 
// Important: EPSV can cause problems with some deep-inspection firewalls. If a
// passive data connection cannot be established, make sure to test with both the
// AutoSetUseEpsv and UseEpsv properties set equal to false.
//
func (z *Ftp2) UseEpsv() bool {
    v := int(C.CkFtp2_getUseEpsv(z.ckObj))
    return v != 0
}
// property setter: UseEpsv
func (z *Ftp2) SetUseEpsv(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putUseEpsv(z.ckObj,v)
}

// property: Username
// Username for logging into the FTP server. Defaults to "anonymous".
func (z *Ftp2) Username() string {
    return C.GoString(C.CkFtp2_username(z.ckObj))
}
// property setter: Username
func (z *Ftp2) SetUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkFtp2_putUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Ftp2) VerboseLogging() bool {
    v := int(C.CkFtp2_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Ftp2) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkFtp2_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Ftp2) Version() string {
    return C.GoString(C.CkFtp2_version(z.ckObj))
}
// Same as PutFile but the file on the FTP server is appended.
// 
// If the remoteFilePath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) AppendFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_AppendFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the AppendFile method
func (z *Ftp2) AppendFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_AppendFileAsync(z.ckObj, cstr1, cstr2)

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


// Same as PutFileFromBinaryData, except the file on the FTP server is appended.
func (z *Ftp2) AppendFileFromBinaryData(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkFtp2_AppendFileFromBinaryData(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Asynchronous version of the AppendFileFromBinaryData method
func (z *Ftp2) AppendFileFromBinaryDataAsync(arg1 string, arg2 []byte, c chan *Task) {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    hTask := C.CkFtp2_AppendFileFromBinaryDataAsync(z.ckObj, cstr1, ckbyd2)

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


// Same as PutFileFromTextData, except the file on the FTP server is appended.
func (z *Ftp2) AppendFileFromTextData(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkFtp2_AppendFileFromTextData(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the AppendFileFromTextData method
func (z *Ftp2) AppendFileFromTextDataAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkFtp2_AppendFileFromTextDataAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Changes the current remote directory. The remoteDirPath should be relative to the current
// remote directory, which is initially the HOME directory of the FTP user account.
// 
// If the remoteDirPath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) ChangeRemoteDir(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_ChangeRemoteDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the ChangeRemoteDir method
func (z *Ftp2) ChangeRemoteDirAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_ChangeRemoteDirAsync(z.ckObj, cstr1)

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


// Returns true if currently connected and logged into an FTP server, otherwise
// returns false.
// 
// Note: This may cause a NOOP command to be sent to the FTP server.
//
func (z *Ftp2) CheckConnection() bool {

    v := C.CkFtp2_CheckConnection(z.ckObj)


    return v != 0
}


// Asynchronous version of the CheckConnection method
func (z *Ftp2) CheckConnectionAsync(c chan *Task) {

    hTask := C.CkFtp2_CheckConnectionAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Reverts the FTP control channel from SSL/TLS to an unencrypted channel. This may
// be required when using FTPS with AUTH TLS where the FTP client is behind a DSL
// or cable-modem router that performs NAT (network address translation). If the
// control channel is encrypted, the router is unable to translate the IP address
// sent in the PORT command for data transfers. By clearing the control channel,
// the data transfers will remain encrypted, but the FTP commands are passed
// unencrypted. Your program would typically clear the control channel after
// authenticating.
func (z *Ftp2) ClearControlChannel() bool {

    v := C.CkFtp2_ClearControlChannel(z.ckObj)


    return v != 0
}


// Asynchronous version of the ClearControlChannel method
func (z *Ftp2) ClearControlChannelAsync(c chan *Task) {

    hTask := C.CkFtp2_ClearControlChannelAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The GetDirCount method returns the count of files and sub-directories in the
// current remote FTP directory, according to the ListPattern property. For
// example, if ListPattern is set to "*.xml", then GetDirCount returns the count of
// XML files in the remote directory.
// 
// The 1st time it is accessed, the component will (behind the scenes) fetch the
// directory listing from the FTP server. This information is cached in the
// component until (1) the current remote directory is changed, or (2) the
// ListPattern is changed, or (3) the this method (ClearDirCache) is called.
//
func (z *Ftp2) ClearDirCache()  {

    C.CkFtp2_ClearDirCache(z.ckObj)



}


// Clears the in-memory session log.
func (z *Ftp2) ClearSessionLog()  {

    C.CkFtp2_ClearSessionLog(z.ckObj)



}


// Connects and logs in to the FTP server using the username/password provided in
// the component properties. Check the integer value of the ConnectFailReason if
// this method returns false (indicating failure).
// 
// Note: To separately establish the connection and then authenticate (in separate
// method calls), call ConnectOnly followed by LoginAfterConnectOnly.
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
//
func (z *Ftp2) Connect() bool {

    v := C.CkFtp2_Connect(z.ckObj)


    return v != 0
}


// Asynchronous version of the Connect method
func (z *Ftp2) ConnectAsync(c chan *Task) {

    hTask := C.CkFtp2_ConnectAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Connects to the FTP server, but does not authenticate. The combination of
// calling this method followed by LoginAfterConnectOnly is the equivalent of
// calling the Connect method (which both connects and authenticates).
// 
// Important: All TCP-based Internet communications, regardless of the protocol
// (such as HTTP, FTP, SSH, IMAP, POP3, SMTP, etc.), and regardless of SSL/TLS,
// begin with establishing a TCP connection to a remote host:port. External
// security-related infrastructure such as software firewalls (Windows Firewall),
// hardware firewalls, anti-virus, at either source or destination (or both) can
// block the connection. If the connection fails, make sure to check all potential
// external causes of blockage.
//
func (z *Ftp2) ConnectOnly() bool {

    v := C.CkFtp2_ConnectOnly(z.ckObj)


    return v != 0
}


// Asynchronous version of the ConnectOnly method
func (z *Ftp2) ConnectOnlyAsync(c chan *Task) {

    hTask := C.CkFtp2_ConnectOnlyAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Explicitly converts the control channel to a secure SSL/TLS connection.
// 
// Note: If you initially connect with either the AuthTls or AuthSsl property set
// to true, then DO NOT call ConvertToTls. The control channel is automatically
// converted to SSL/TLS from within the Connect method when these properties are
// set.
// 
// Note: It is very uncommon for this method to be needed.
//
func (z *Ftp2) ConvertToTls() bool {

    v := C.CkFtp2_ConvertToTls(z.ckObj)


    return v != 0
}


// Asynchronous version of the ConvertToTls method
func (z *Ftp2) ConvertToTlsAsync(c chan *Task) {

    hTask := C.CkFtp2_ConvertToTlsAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Creates an "FTP plan" that lists the FTP operations that would be performed when
// PutTree is called. Additionally, the PutPlan method executes an "FTP plan" and
// logs each successful operation to a plan log file. If a large-scale upload is
// interrupted, the PutPlan can be resumed, skipping over the operations already
// listed in the plan log file.
// return a string or nil if failed.
func (z *Ftp2) CreatePlan(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFtp2_createPlan(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the CreatePlan method
func (z *Ftp2) CreatePlanAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_CreatePlanAsync(z.ckObj, cstr1)

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


// Creates a directory on the FTP server. If the directory already exists, a new
// one is not created and false is returned.
// 
// If the remoteDirPath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) CreateRemoteDir(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_CreateRemoteDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the CreateRemoteDir method
func (z *Ftp2) CreateRemoteDirAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_CreateRemoteDirAsync(z.ckObj, cstr1)

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


// Deletes all the files in the current remote FTP directory matching the pattern.
// Returns the number of files deleted, or -1 for failure. The pattern is a string
// such as "*.txt", where any number of "*" wildcard characters can be used. "*"
// matches 0 or more of any character.
func (z *Ftp2) DeleteMatching(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_DeleteMatching(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the DeleteMatching method
func (z *Ftp2) DeleteMatchingAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_DeleteMatchingAsync(z.ckObj, cstr1)

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


// Deletes a file on the FTP server.
// 
// If the remoteFilePath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) DeleteRemoteFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_DeleteRemoteFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the DeleteRemoteFile method
func (z *Ftp2) DeleteRemoteFileAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_DeleteRemoteFileAsync(z.ckObj, cstr1)

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


// Deletes the entire subtree and all files from the current remote FTP directory.
// To delete a subtree on the FTP server, your program would first navigate to the
// root of the subtree to be deleted by calling ChangeRemoteDir, and then call
// DeleteTree. There are two event callbacks: VerifyDeleteFile and VerifyDeleteDir.
// Both are called prior to deleting each file or directory. The arguments to the
// callback include the full filepath of the file or directory, and an output-only
// "skip" flag. If your application sets the skip flag to true, the file or
// directory is NOT deleted. If a directory is not deleted, all files and
// sub-directories will remain. Example programs can be found at
// http://www.example-code.com/
func (z *Ftp2) DeleteTree() bool {

    v := C.CkFtp2_DeleteTree(z.ckObj)


    return v != 0
}


// Asynchronous version of the DeleteTree method
func (z *Ftp2) DeleteTreeAsync(c chan *Task) {

    hTask := C.CkFtp2_DeleteTreeAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Automatically determines the ProxyMethod that should be used with an FTP proxy
// server. Tries each of the five possible ProxyMethod settings and returns the
// value (1-5) of the ProxyMethod that succeeded.
// 
// This method may take a minute or two to complete. Returns 0 if no proxy methods
// were successful. Returns -1 to indicate an error (i.e. it was unable to test all
// proxy methods.)
//
func (z *Ftp2) DetermineProxyMethod() int {

    v := C.CkFtp2_DetermineProxyMethod(z.ckObj)


    return int(v)
}


// Asynchronous version of the DetermineProxyMethod method
func (z *Ftp2) DetermineProxyMethodAsync(c chan *Task) {

    hTask := C.CkFtp2_DetermineProxyMethodAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Discovers which combinations of FTP2 property settings result in successful data
// transfers.
// 
// DetermineSettings tries 13 different combinations of these properties:
// Ssl
// AuthTls
// AuthSsl
// Port
// Passive
// PassiveUseHostAddr
// Within the FTP protocol, the process of fetching a directory listing is also
// considered a "data transfer". The DetermineSettings method works by checking to
// see which combinations result in a successful directory listing download. The
// method takes no arguments and returns a string containing an XML report of the
// results. It is a blocking call that may take approximately a minute to run. If
// you are unsure about how to interpret the results, cut-and-paste it into an
// email and send it to support@chilkatsoft.com.
//
// return a string or nil if failed.
func (z *Ftp2) DetermineSettings() *string {

    retStr := C.CkFtp2_determineSettings(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the DetermineSettings method
func (z *Ftp2) DetermineSettingsAsync(c chan *Task) {

    hTask := C.CkFtp2_DetermineSettingsAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Recursively downloads the structure of a complete remote directory tree. Returns
// an XML document with the directory structure.
// 
// Note: Starting in version 9.5.0.47, the SyncMustMatch and SyncMustNotMatch
// properties apply to this method.
//
// return a string or nil if failed.
func (z *Ftp2) DirTreeXml() *string {

    retStr := C.CkFtp2_dirTreeXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the DirTreeXml method
func (z *Ftp2) DirTreeXmlAsync(c chan *Task) {

    hTask := C.CkFtp2_DirTreeXmlAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Disconnects from the FTP server, ending the current session.
func (z *Ftp2) Disconnect() bool {

    v := C.CkFtp2_Disconnect(z.ckObj)


    return v != 0
}


// Asynchronous version of the Disconnect method
func (z *Ftp2) DisconnectAsync(c chan *Task) {

    hTask := C.CkFtp2_DisconnectAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads an entire tree from the FTP server and recreates the directory tree on
// the local filesystem.
// 
// This method downloads all the files and subdirectories in the current remote
// directory. An application would first navigate to the directory to be downloaded
// via ChangeRemoteDir and then call this method.
// 
// Note: Starting in version 9.5.0.47, the SyncMustMatch and SyncMustNotMatch
// properties apply to this method.
//
func (z *Ftp2) DownloadTree(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_DownloadTree(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the DownloadTree method
func (z *Ftp2) DownloadTreeAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_DownloadTreeAsync(z.ckObj, cstr1)

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


// Sends a FEAT command to the FTP server and returns the response. Returns a
// zero-length string to indicate failure. Here is a typical response:
// 211-Features:
//  MDTM
//  REST STREAM
//  SIZE
//  MLST type*;size*;modify*;
//  MLSD
//  AUTH SSL
//  AUTH TLS
//  UTF8
//  CLNT
//  MFMT
// 211 End
// return a string or nil if failed.
func (z *Ftp2) Feat() *string {

    retStr := C.CkFtp2_feat(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the Feat method
func (z *Ftp2) FeatAsync(c chan *Task) {

    hTask := C.CkFtp2_FeatAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the create date/time for the Nth file or sub-directory in the current
// remote directory. The first file/dir is at index 0, and the last one is at index
// (GetDirCount()-1)
func (z *Ftp2) GetCreateDt(arg1 int) *CkDateTime {

    retObj := C.CkFtp2_GetCreateDt(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Asynchronous version of the GetCreateDt method
func (z *Ftp2) GetCreateDtAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetCreateDtAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the file-creation date/time for a remote file by filename.
// 
// Note: The filename passed to this method must NOT include a path. Prior to calling
// this method, make sure to set the current remote directory (via the
// ChangeRemoteDir method) to the remote directory where this file exists.
// 
// Note: Prior to calling this method, it should be ensured that the ListPattern
// property is set to a pattern that would match the requested filename. (The default
// value of ListPattern is "*", which will match all filenames.)
// 
// Note: Linux/Unix type filesystems do not store "create" date/times. Therefore,
// if the FTP server is on such as system, this method will return a date/time
// equal to the last-modified date/time.
//
func (z *Ftp2) GetCreateDtByName(arg1 string) *CkDateTime {
    cstr1 := C.CString(arg1)

    retObj := C.CkFtp2_GetCreateDtByName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Asynchronous version of the GetCreateDtByName method
func (z *Ftp2) GetCreateDtByNameAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetCreateDtByNameAsync(z.ckObj, cstr1)

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


// Returns the file-creation date/time (in RFC822 string format, such as "Tue, 25
// Sep 2012 12:25:32 -0500") for a remote file by filename.
// 
// Note: The filename passed to this method must NOT include a path. Prior to calling
// this method, make sure to set the current remote directory (via the
// ChangeRemoteDir method) to the remote directory where this file exists.
// 
// Note: Prior to calling this method, it should be ensured that the ListPattern
// property is set to a pattern that would match the requested filename. (The default
// value of ListPattern is "*", which will match all filenames.)
// 
// Note: Linux/Unix type filesystems do not store "create" date/times. If the FTP
// server is on such as system, this method will return a date/time equal to the
// last-modified date/time.
//
// return a string or nil if failed.
func (z *Ftp2) GetCreateTimeByNameStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFtp2_getCreateTimeByNameStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetCreateTimeByNameStr method
func (z *Ftp2) GetCreateTimeByNameStrAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetCreateTimeByNameStrAsync(z.ckObj, cstr1)

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


// Returns the create time (in RFC822 string format, such as "Tue, 25 Sep 2012
// 12:25:32 -0500") for the Nth file or sub-directory in the current remote
// directory. The first file/dir is at index 0, and the last one is at index
// (GetDirCount()-1)
// return a string or nil if failed.
func (z *Ftp2) GetCreateTimeStr(arg1 int) *string {

    retStr := C.CkFtp2_getCreateTimeStr(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetCreateTimeStr method
func (z *Ftp2) GetCreateTimeStrAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetCreateTimeStrAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the current remote directory.
// return a string or nil if failed.
func (z *Ftp2) GetCurrentRemoteDir() *string {

    retStr := C.CkFtp2_getCurrentRemoteDir(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetCurrentRemoteDir method
func (z *Ftp2) GetCurrentRemoteDirAsync(c chan *Task) {

    hTask := C.CkFtp2_GetCurrentRemoteDirAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the number of files and sub-directories in the current remote directory
// that match the ListPattern property.
// 
// Important: Calling this method may cause the directory listing to be retrieved
// from the FTP server. For FTP servers that do not support the MLST/MLSD commands,
// this is technically a data transfer that requires a temporary data connection to
// be established in the same way as when uploading or downloading files. If your
// program hangs while calling this method, it probably means that the data
// connection could not be established. The most common solution is to switch to
// using Passive mode by setting the Passive property = true, with the
// PassiveUseHostAddr property also set equal to true. If this does not help,
// examine the contents of the LastErrorText property after this method finally
// returns (after timing out). Also, seethis Chilkat blog post about FTP connection
// settings
// <http://www.cknotes.com/?p=282>.
//
func (z *Ftp2) GetDirCount() int {

    v := C.CkFtp2_GetDirCount(z.ckObj)


    return int(v)
}


// Asynchronous version of the GetDirCount method
func (z *Ftp2) GetDirCountAsync(c chan *Task) {

    hTask := C.CkFtp2_GetDirCountAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads a file from the FTP server to the local filesystem.
// 
// If the remoteFilePath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) GetFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_GetFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the GetFile method
func (z *Ftp2) GetFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_GetFileAsync(z.ckObj, cstr1, cstr2)

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


// Downloads a file from the FTP server into a BinData object.
// 
// If the remoteFilePath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) GetFileBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_GetFileBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the GetFileBd method
func (z *Ftp2) GetFileBdAsync(arg1 string, arg2 *BinData, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetFileBdAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Returns the filename for the Nth file or sub-directory in the current remote
// directory. The first file/dir is at index 0, and the last one is at index
// (GetDirCount()-1)
// return a string or nil if failed.
func (z *Ftp2) GetFilename(arg1 int) *string {

    retStr := C.CkFtp2_getFilename(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetFilename method
func (z *Ftp2) GetFilenameAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetFilenameAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads a file from the FTP server into a StringBuilder object.
// 
// If the remoteFilePath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) GetFileSb(arg1 string, arg2 string, arg3 *StringBuilder) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_GetFileSb(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the GetFileSb method
func (z *Ftp2) GetFileSbAsync(arg1 string, arg2 string, arg3 *StringBuilder, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_GetFileSbAsync(z.ckObj, cstr1, cstr2, arg3.ckObj)

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


// Downloads a file to a stream. If called synchronously, the remoteFilePath must have a
// sink, such as a file or another stream object. If called asynchronously, then
// the foreground thread can read the stream.
func (z *Ftp2) GetFileToStream(arg1 string, arg2 *Stream) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_GetFileToStream(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the GetFileToStream method
func (z *Ftp2) GetFileToStreamAsync(arg1 string, arg2 *Stream, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetFileToStreamAsync(z.ckObj, cstr1, arg2.ckObj)

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


// Returns group name, if available, for the Nth file. If empty, then no group
// information is available.
// 
// Note: When MLSD is used to get directory listings, it is likely that the owner
// and group information is not transmitted. In cases where the FTP server is on a
// UNIX/Linux system, the AllowMlsd property can be set to false to force UNIX
// directory listings instead of MLSD directory listings. This should result in
// being able to obtain owner/group information. However, it may sacrifice the
// quality and accuracy of the various date/time values that are returned.
//
// return a string or nil if failed.
func (z *Ftp2) GetGroup(arg1 int) *string {

    retStr := C.CkFtp2_getGroup(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetGroup method
func (z *Ftp2) GetGroupAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetGroupAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns true for a sub-directory and false for a file, for the Nth entry in
// the current remote directory. The first file/dir is at index 0, and the last one
// is at index (GetDirCount()-1)
func (z *Ftp2) GetIsDirectory(arg1 int) bool {

    v := C.CkFtp2_GetIsDirectory(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the GetIsDirectory method
func (z *Ftp2) GetIsDirectoryAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetIsDirectoryAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns true if the remote file is a symbolic link. (Symbolic links only exist
// on Unix/Linux systems, not on Windows filesystems.)
func (z *Ftp2) GetIsSymbolicLink(arg1 int) bool {

    v := C.CkFtp2_GetIsSymbolicLink(z.ckObj, C.int(arg1))


    return v != 0
}


// Asynchronous version of the GetIsSymbolicLink method
func (z *Ftp2) GetIsSymbolicLinkAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetIsSymbolicLinkAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the last modified date/time for the Nth file or sub-directory in the
// current remote directory. The first file/dir is at index 0, and the last one is
// at index (GetDirCount()-1)
func (z *Ftp2) GetLastModDt(arg1 int) *CkDateTime {

    retObj := C.CkFtp2_GetLastModDt(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Asynchronous version of the GetLastModDt method
func (z *Ftp2) GetLastModDtAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetLastModDtAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the last-modified date/time for a remote file.
// 
// Note: The filename passed to this method must NOT include a path. Prior to calling
// this method, make sure to set the current remote directory (via the
// ChangeRemoteDir method) to the remote directory where this file exists.
// 
// Note: Prior to calling this method, it should be ensured that the ListPattern
// property is set to a pattern that would match the requested filename. (The default
// value of ListPattern is "*", which will match all filenames.)
//
func (z *Ftp2) GetLastModDtByName(arg1 string) *CkDateTime {
    cstr1 := C.CString(arg1)

    retObj := C.CkFtp2_GetLastModDtByName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Asynchronous version of the GetLastModDtByName method
func (z *Ftp2) GetLastModDtByNameAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetLastModDtByNameAsync(z.ckObj, cstr1)

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


// Returns a remote file's last-modified date/time in RFC822 string format, such as
// "Tue, 25 Sep 2012 12:25:32 -0500".
// 
// Note: The filename passed to this method must NOT include a path. Prior to calling
// this method, make sure to set the current remote directory (via the
// ChangeRemoteDir method) to the remote directory where this file exists.
// 
// Note: Prior to calling this method, it should be ensured that the ListPattern
// property is set to a pattern that would match the requested filename. (The default
// value of ListPattern is "*", which will match all filenames.)
//
// return a string or nil if failed.
func (z *Ftp2) GetLastModifiedTimeByNameStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFtp2_getLastModifiedTimeByNameStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetLastModifiedTimeByNameStr method
func (z *Ftp2) GetLastModifiedTimeByNameStrAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetLastModifiedTimeByNameStrAsync(z.ckObj, cstr1)

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


// Returns the last modified date/time (in RFC822 string format, such as "Tue, 25
// Sep 2012 12:25:32 -0500") for the Nth file or sub-directory in the current
// remote directory. The first file/dir is at index 0, and the last one is at index
// (GetDirCount()-1)
// return a string or nil if failed.
func (z *Ftp2) GetLastModifiedTimeStr(arg1 int) *string {

    retStr := C.CkFtp2_getLastModifiedTimeStr(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetLastModifiedTimeStr method
func (z *Ftp2) GetLastModifiedTimeStrAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetLastModifiedTimeStrAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns owner name, if available, for the Nth file. If empty, then no owner
// information is available.
// 
// Note: When MLSD is used to get directory listings, it is likely that the owner
// and group information is not transmitted. In cases where the FTP server is on a
// UNIX/Linux system, the AllowMlsd property can be set to false to force UNIX
// directory listings instead of MLSD directory listings. This should result in
// being able to obtain owner/group information. However, it may sacrifice the
// quality and accuracy of the various date/time values that are returned.
//
// return a string or nil if failed.
func (z *Ftp2) GetOwner(arg1 int) *string {

    retStr := C.CkFtp2_getOwner(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetOwner method
func (z *Ftp2) GetOwnerAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetOwnerAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns permissions information, if available, for the Nth file. If empty, then
// no permissions information is available. The value returned by the GetPermType
// method defines the content and format of the permissions string returned by this
// method. Possible permission types are "mlsd", "unix", "netware", "openvms", and
// "batchStatusFlags". The format of each permission type is as follows:
// --------------------------------------------------------------------------------
// 
// PermType: mlsd:
// 
// A "perm fact" is returned. The format of the perm fact is defined in RFC 3659 as
// follows:
//   The perm fact is used to indicate access rights the current FTP user
//    has over the object listed.  Its value is always an unordered
//    sequence of alphabetic characters.
// 
//       perm-fact    = "Perm" "=" *pvals
//       pvals        = "a" / "c" / "d" / "e" / "f" /
//                      "l" / "m" / "p" / "r" / "w"
// 
//    There are ten permission indicators currently defined.  Many are
//    meaningful only when used with a particular type of object.  The
//    indicators are case independent, "d" and "D" are the same indicator.
// 
//    The "a" permission applies to objects of type=file, and indicates
//    that the APPE (append) command may be applied to the file named.
// 
//    The "c" permission applies to objects of type=dir (and type=pdir,
//    type=cdir).  It indicates that files may be created in the directory
//    named.  That is, that a STOU command is likely to succeed, and that
//    STOR and APPE commands might succeed if the file named did not
//    previously exist, but is to be created in the directory object that
//    has the "c" permission.  It also indicates that the RNTO command is
//    likely to succeed for names in the directory.
// 
//    The "d" permission applies to all types.  It indicates that the
//    object named may be deleted, that is, that the RMD command may be
//    applied to it if it is a directory, and otherwise that the DELE
//    command may be applied to it.
// 
//    The "e" permission applies to the directory types.  When set on an
//    object of type=dir, type=cdir, or type=pdir it indicates that a CWD
//    command naming the object should succeed, and the user should be able
//    to enter the directory named.  For type=pdir it also indicates that
//    the CDUP command may succeed (if this particular pathname is the one
//    to which a CDUP would apply.)
// 
//    The "f" permission for objects indicates that the object named may be
//    renamed - that is, may be the object of an RNFR command.
// 
//    The "l" permission applies to the directory file types, and indicates
//    that the listing commands, LIST, NLST, and MLSD may be applied to the
//    directory in question.
// 
//    The "m" permission applies to directory types, and indicates that the
//    MKD command may be used to create a new directory within the
//    directory under consideration.
// 
//    The "p" permission applies to directory types, and indicates that
//    objects in the directory may be deleted, or (stretching naming a
//    little) that the directory may be purged.  Note: it does not indicate
//    that the RMD command may be used to remove the directory named
//    itself, the "d" permission indicator indicates that.
// 
//    The "r" permission applies to type=file objects, and for some
//    systems, perhaps to other types of objects, and indicates that the
//    RETR command may be applied to that object.
// 
//    The "w" permission applies to type=file objects, and for some
//    systems, perhaps to other types of objects, and indicates that the
//    STOR command may be applied to the object named.
// 
//    Note: That a permission indicator is set can never imply that the
//       appropriate command is guaranteed to work -- just that it might.
//       Other system specific limitations, such as limitations on
//       available space for storing files, may cause an operation to fail,
//       where the permission flags may have indicated that it was likely
//       to succeed.  The permissions are a guide only.
// 
//    Implementation note: The permissions are described here as they apply
//       to FTP commands.  They may not map easily into particular
//       permissions available on the server's operating system.  Servers
//       are expected to synthesize these permission bits from the
//       permission information available from operating system.  For
//       example, to correctly determine whether the "D" permission bit
//       should be set on a directory for a server running on the UNIX(TM)
//       operating system, the server should check that the directory named
//       is empty, and that the user has write permission on both the
//       directory under consideration, and its parent directory.
// 
//       Some systems may have more specific permissions than those listed
//       here, such systems should map those to the flags defined as best
//       they are able.  Other systems may have only more broad access
//       controls.  They will generally have just a few possible
//       permutations of permission flags, however they should attempt to
//       correctly represent what is permitted.
// --------------------------------------------------------------------------------
// 
// PermType: unix:
// 
// A Unix/Linux permissions string is returned ( such as "drwxr-xr-x" or
// "-rw-r--r--")
//     The UNIX permissions string is 10 characters. Each character has a specific meaning. If the first character is:
//     d 	the entry is a directory.
//     b 	the entry is a block special file.
//     c 	the entry is a character special file.
//     l 	the entry is a symbolic link. Either the -N flag was specified, or the symbolic link did not point to an existing file.
//     p 	the entry is a first-in, first-out (FIFO) special file.
//     s 	the entry is a local socket.
//     - 	the entry is an ordinary file.
// 
//     The next nine characters are divided into three sets of three characters each. The first set of three characters show 
// the owner's permission. The next set of three characters show the permission of the other users in the group. The last
// set of three characters shows the permission of anyone else with access to the file. The three characters in each set 
// indicate, respectively, read, write, and execute permission of the file. With execute permission of a directory, you can search 
// a directory for a specified file. Permissions are indicated like this:
// 
//     r 	read
//     w 	write (edit)
//     x 	execute (search)
//     - 	corresponding permission not granted 
// --------------------------------------------------------------------------------
// 
// PermType: netware:
// 
// Contains the NetWare rights string from a NetWare FTP server directory listing
// format. For example "-WCE---S" or "RWCEAFMS".
// Directory Rights	Description
// ----------------	-------------------------------
// Read (R)		Read data from an existing file.
// Write (W)		Write data to an existing file.
// Create (C)		Create a new file or subdirectory.
// Erase (E)		Delete an existing files or directory.
// Modify (M)	Rename and change attributes of a file.
// File Scan (F)	List the contents of a directory.
// Access Control (A)	Control the rights of other users to access files or directories.
// Supervisor (S)	Automatically allowed all rights.
// --------------------------------------------------------------------------------
// 
// PermType: openvms:
// 
// Contains the OpenVMS permissions string. For example "(RWED,RWED,RWED,RWED)",
// "(RWED,RWED,,)", "(RWED,RWED,R,R)", etc.
// --------------------------------------------------------------------------------
// 
// PermType: batchStatusFlags:
// 
// Contains the batch status flags from a Connect:Enterprise Server. Such as
// "-CR--M----" or "-ART------".
// The Batch Status Flags  is a 10-character string where each character describes an attribute of the batch. 
// A dash indicates that flag is turned off and therefore has no meaning to the 
// batch in question. The flags are always displayed in the same order: 
// 
// 1) I  -- Incomplete batch which will NOT be processed. 
// 2) A or C -- Added or Collected
// 3) R -- Requestable by partner 
// 4) T -- Transmitted to partner 
// 5) E -- Extracted (inbound file processed by McLane) 
// 6) M -- Multi-transmittable 
// 7) U -- Un-extractable 
// 8) N -- Non-transmittable 
// 9) P -- In Progress 
// 10) - -- Always a dash.
//
// return a string or nil if failed.
func (z *Ftp2) GetPermissions(arg1 int) *string {

    retStr := C.CkFtp2_getPermissions(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetPermissions method
func (z *Ftp2) GetPermissionsAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetPermissionsAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the type of permissions information that is available for the Nth file.
// If empty, then no permissions information is available. The value returned by
// this method defines the content and format of the permissions string returned by
// the GetPermissions method. Possible values are "mlsd", "unix", "netware",
// "openvms", and "batchStatusFlags".
// return a string or nil if failed.
func (z *Ftp2) GetPermType(arg1 int) *string {

    retStr := C.CkFtp2_getPermType(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetPermType method
func (z *Ftp2) GetPermTypeAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetPermTypeAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Downloads the contents of a remote file into a byte array.
func (z *Ftp2) GetRemoteFileBinaryData(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkFtp2_GetRemoteFileBinaryData(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the GetRemoteFileBinaryData method
func (z *Ftp2) GetRemoteFileBinaryDataAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetRemoteFileBinaryDataAsync(z.ckObj, cstr1)

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


// Downloads a text file directly into a string variable. The character encoding of
// the text file is specified by the charset argument, which is a value such as utf-8,
// iso-8859-1, Shift_JIS, etc.
// return a string or nil if failed.
func (z *Ftp2) GetRemoteFileTextC(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkFtp2_getRemoteFileTextC(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetRemoteFileTextC method
func (z *Ftp2) GetRemoteFileTextCAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_GetRemoteFileTextCAsync(z.ckObj, cstr1, cstr2)

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


// Downloads the content of a remote text file directly into an in-memory string.
// 
// Note: If the remote text file does not use the ANSI character encoding, call
// GetRemoteFileTextC instead, which allows for the character encoding to be
// specified so that characters are properly interpreted.
//
// return a string or nil if failed.
func (z *Ftp2) GetRemoteFileTextData(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFtp2_getRemoteFileTextData(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetRemoteFileTextData method
func (z *Ftp2) GetRemoteFileTextDataAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetRemoteFileTextDataAsync(z.ckObj, cstr1)

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


// Returns the size of the Nth remote file in the current directory.
func (z *Ftp2) GetSize(arg1 int) int {

    v := C.CkFtp2_GetSize(z.ckObj, C.int(arg1))


    return int(v)
}


// Asynchronous version of the GetSize method
func (z *Ftp2) GetSizeAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetSizeAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the size of the Nth remote file in the current directory as a 64-bit
// integer. Returns -1 if the file does not exist.
func (z *Ftp2) GetSize64(arg1 int) int64 {

    v := C.CkFtp2_GetSize64(z.ckObj, C.int(arg1))


    return int64(v)
}


// Returns a remote file's size in bytes. Returns -1 if the file does not exist.
// 
// Note: The filename passed to this method must NOT include a path. Prior to calling
// this method, make sure to set the current remote directory (via the
// ChangeRemoteDir method) to the remote directory where this file exists.
// 
// Note: Prior to calling this method, it should be ensured that the ListPattern
// property is set to a pattern that would match the requested filename. (The default
// value of ListPattern is "*", which will match all filenames.)
//
func (z *Ftp2) GetSizeByName(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_GetSizeByName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the GetSizeByName method
func (z *Ftp2) GetSizeByNameAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetSizeByNameAsync(z.ckObj, cstr1)

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


// Returns a remote file's size in bytes as a 64-bit integer.
// 
// Note: The filename passed to this method must NOT include a path. Prior to calling
// this method, make sure to set the current remote directory (via the
// ChangeRemoteDir method) to the remote directory where this file exists.
// 
// Note: Prior to calling this method, it should be ensured that the ListPattern
// property is set to a pattern that would match the requested filename. (The default
// value of ListPattern is "*", which will match all filenames.)
//
func (z *Ftp2) GetSizeByName64(arg1 string) int64 {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_GetSizeByName64(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int64(v)
}


// Returns the size in decimal string format of the Nth remote file in the current
// directory. This is helpful for cases when the file size (in bytes) is greater
// than what can fit in a 32-bit integer.
// return a string or nil if failed.
func (z *Ftp2) GetSizeStr(arg1 int) *string {

    retStr := C.CkFtp2_getSizeStr(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetSizeStr method
func (z *Ftp2) GetSizeStrAsync(arg1 int, c chan *Task) {

    hTask := C.CkFtp2_GetSizeStrAsync(z.ckObj, C.int(arg1))


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the size of a remote file as a string. This is helpful when file a file
// size is greater than what can fit in a 32-bit integer.
// 
// Note: The filename passed to this method must NOT include a path. Prior to calling
// this method, make sure to set the current remote directory (via the
// ChangeRemoteDir method) to the remote directory where this file exists.
// 
// Note: Prior to calling this method, it should be ensured that the ListPattern
// property is set to a pattern that would match the requested filename. (The default
// value of ListPattern is "*", which will match all filenames.)
//
// return a string or nil if failed.
func (z *Ftp2) GetSizeStrByName(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFtp2_getSizeStrByName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetSizeStrByName method
func (z *Ftp2) GetSizeStrByNameAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetSizeStrByNameAsync(z.ckObj, cstr1)

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


// Returns the FTP server's digital certificate (for SSL / TLS connections).
func (z *Ftp2) GetSslServerCert() *Cert {

    retObj := C.CkFtp2_GetSslServerCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns a listing of the files and directories in the current directory matching
// the pattern. Passing "*.*" will return all the files and directories.
// return a string or nil if failed.
func (z *Ftp2) GetTextDirListing(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFtp2_getTextDirListing(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetTextDirListing method
func (z *Ftp2) GetTextDirListingAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetTextDirListingAsync(z.ckObj, cstr1)

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


// Returns (in XML format) the files and directories in the current directory
// matching the pattern. Passing "*.*" will return all the files and directories.
// 
// Note: The lastModTime XML elements contain date/time information in the local
// (client) timezone. However, it's possible based on the capabilities of an FTP
// server (or lack of capabilities) that the timezone information for the remote
// files is not available. In other words, in some cases, the timezone of an FTP
// server cannot be known, especially for older FTP server implementations.
//
// return a string or nil if failed.
func (z *Ftp2) GetXmlDirListing(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFtp2_getXmlDirListing(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetXmlDirListing method
func (z *Ftp2) GetXmlDirListingAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_GetXmlDirListingAsync(z.ckObj, cstr1)

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


// Return true if the component is already unlocked.
func (z *Ftp2) IsUnlocked() bool {

    v := C.CkFtp2_IsUnlocked(z.ckObj)


    return v != 0
}


// This is the same as PutFile, but designed to work around the following potential
// problem associated with an upload that is extremely large.
// 
// FTP uses two TCP (or TLS) connections: a control connection to submit commands
// and receive replies, and a data connection for actual file transfers. It is the
// nature of FTP that during a transfer the control connection stays completely
// idle. Many routers and firewalls automatically close idle connections after a
// certain period of time. Worse, they often don't notify the user, but just
// silently drop the connection.
// 
// For FTP, this means that during a long transfer the control connection can get
// dropped because it is detected as idle, but neither client nor server are
// notified. When all data has been transferred, the server assumes the control
// connection is alive and it sends the transfer confirmation reply.
// 
// Likewise, the client thinks the control connection is alive and it waits for the
// reply from the server. But since the control connection got dropped without
// notification, the reply never arrives and eventually the connection will
// timeout.
// 
// The Solution: This method uploads the file in chunks, where each chunk appends
// to the remote file. This way, each chunk is a separate FTP upload that does not
// take too long to complete. The chunkSize specifies the number of bytes to upload in
// each chunk. The size should be based on the amount of memory available (because
// each chunk will reside in memory as it's being uploaded), the transfer rate, and
// the total size of the file being uploaded. For example, if a 4GB file is
// uploaded, and the chunkSize is set to 1MB (1,048,576 bytes), then 4000 separate
// chunks would be required. This is likely not a good choice for chunkSize. A more
// appropriate chunkSize might be 20MB, in which case the upload would complete in 200
// separate chunks. The application would temporarily be using a 20MB buffer for
// uploading chunks. The tradeoff is between the number of chunks (the more chunks,
// the larger the overall time to upload), the amount of memory that is reasonable
// for the temporary buffer, and the amount of time required to upload each chunk
// (if the chunk size is too large, then the problem described above is not
// solved).
//
func (z *Ftp2) LargeFileUpload(arg1 string, arg2 string, arg3 int) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_LargeFileUpload(z.ckObj, cstr1, cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the LargeFileUpload method
func (z *Ftp2) LargeFileUploadAsync(arg1 string, arg2 string, arg3 int, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_LargeFileUploadAsync(z.ckObj, cstr1, cstr2, C.int(arg3))

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
func (z *Ftp2) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkFtp2_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Authenticates with the FTP server using the values provided in the Username,
// Password, and/or other properties. This can be called after establishing the
// connection via the ConnectOnly method. (The Connect method both connects and
// authenticates.) The combination of calling ConnectOnly followed by
// LoginAfterConnectOnly is the equivalent of calling the Connect method.
// 
// Note: After successful authentication, the FEAT and SYST commands are
// automatically sent to help the client understand what is supported by the FTP
// server. To prevent these commands from being sent, set the AutoFeat and/or
// AutoSyst properties equal to false.
//
func (z *Ftp2) LoginAfterConnectOnly() bool {

    v := C.CkFtp2_LoginAfterConnectOnly(z.ckObj)


    return v != 0
}


// Asynchronous version of the LoginAfterConnectOnly method
func (z *Ftp2) LoginAfterConnectOnlyAsync(c chan *Task) {

    hTask := C.CkFtp2_LoginAfterConnectOnlyAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Copies all the files in the current remote FTP directory to a local directory.
// To copy all the files in a remote directory, set remotePattern to "*.*" The
// pattern can contain any number of "*"characters, where "*" matches 0 or more of
// any character. The return value is the number of files transferred, and on
// error, a value of -1 is returned. Detailed information about the transfer can be
// obtained from the last-error information
// (LastErrorText/LastErrorHtml/LastErrorXml/SaveLastError).
// 
// About case sensitivity: The MGetFiles command works by sending the "LIST"
// command to the FTP server. For example: "LIST *.txt". The FTP server responds
// with a directory listing of the files matching the wildcarded pattern, and it is
// these files that are downloaded. Case sensitivity depends on the
// case-sensitivity of the remote file system. If the FTP server is running on a
// Windows-based computer, it is likely to be case insensitive. However, if the FTP
// server is running on Linux, MAC OS X, etc. it is likely to be case sensitive.
// There is no good way to force case-insensitivity if the remote filesystem is
// case-sensitive because it is not possible for the FTP client to send a LIST
// command indicating that it wants the matching to be case-insensitive.
//
func (z *Ftp2) MGetFiles(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_MGetFiles(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Asynchronous version of the MGetFiles method
func (z *Ftp2) MGetFilesAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_MGetFilesAsync(z.ckObj, cstr1, cstr2)

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


// Uploads all the files matching pattern on the local computer to the current
// remote FTP directory. The pattern parameter can include directory information,
// such as "C:/my_dir/*.txt" or it can simply be a pattern such as "*.*" that
// matches the files in the application's current directory. Subdirectories are not
// recursed. The return value is the number of files copied, with a value of -1
// returned for errors. Detailed information about the transfer can be obtained
// from the XML log.[
func (z *Ftp2) MPutFiles(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_MPutFiles(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Asynchronous version of the MPutFiles method
func (z *Ftp2) MPutFilesAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_MPutFilesAsync(z.ckObj, cstr1)

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


// Sends an NLST command to the FTP server and returns the results in XML format.
// The NLST command returns a list of filenames in the given directory (matching
// the pattern). The remoteDirPattern should be a pattern such as "*", "*.*", "*.txt",
// "subDir/*.xml", etc.
// 
// The format of the XML returned is:
// filename_or_dir_1filename_or_dir_2filename_or_dir_3filename_or_dir_4...
//
// return a string or nil if failed.
func (z *Ftp2) NlstXml(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFtp2_nlstXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the NlstXml method
func (z *Ftp2) NlstXmlAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_NlstXmlAsync(z.ckObj, cstr1)

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


// Issues a no-op command to the FTP server.
func (z *Ftp2) Noop() bool {

    v := C.CkFtp2_Noop(z.ckObj)


    return v != 0
}


// Asynchronous version of the Noop method
func (z *Ftp2) NoopAsync(c chan *Task) {

    hTask := C.CkFtp2_NoopAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Uploads a local file to the current directory on the FTP server.
// 
// If the remoteFilePath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) PutFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_PutFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the PutFile method
func (z *Ftp2) PutFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_PutFileAsync(z.ckObj, cstr1, cstr2)

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


// Uploads the contents of a BinData to a remote file.
// 
// If the remoteFilePath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) PutFileBd(arg1 *BinData, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_PutFileBd(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the PutFileBd method
func (z *Ftp2) PutFileBdAsync(arg1 *BinData, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_PutFileBdAsync(z.ckObj, arg1.ckObj, cstr2)

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


// Creates a file on the remote server containing the data passed in a byte array.
func (z *Ftp2) PutFileFromBinaryData(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkFtp2_PutFileFromBinaryData(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Asynchronous version of the PutFileFromBinaryData method
func (z *Ftp2) PutFileFromBinaryDataAsync(arg1 string, arg2 []byte, c chan *Task) {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    hTask := C.CkFtp2_PutFileFromBinaryDataAsync(z.ckObj, cstr1, ckbyd2)

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


// Creates a file on the remote server containing the data passed in a string.
func (z *Ftp2) PutFileFromTextData(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkFtp2_PutFileFromTextData(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the PutFileFromTextData method
func (z *Ftp2) PutFileFromTextDataAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkFtp2_PutFileFromTextDataAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Uploads the contents of a StringBuilder to a remote file.
// 
// If the charset contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) PutFileSb(arg1 *StringBuilder, arg2 string, arg3 bool, arg4 string) bool {
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    cstr4 := C.CString(arg4)

    v := C.CkFtp2_PutFileSb(z.ckObj, arg1.ckObj, cstr2, b3, cstr4)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Asynchronous version of the PutFileSb method
func (z *Ftp2) PutFileSbAsync(arg1 *StringBuilder, arg2 string, arg3 bool, arg4 string, c chan *Task) {
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    cstr4 := C.CString(arg4)

    hTask := C.CkFtp2_PutFileSbAsync(z.ckObj, arg1.ckObj, cstr2, b3, cstr4)

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


// Executes an "FTP plan" (created by the CreatePlan method) and logs each
// successful operation to a plan log file. If a large-scale upload is interrupted,
// the PutPlan can be resumed, skipping over the operations already listed in the
// plan log file. When resuming an interrupted PutPlan method, use the same log
// file. All completed operations found in the already-existing log will
// automatically be skipped.
func (z *Ftp2) PutPlan(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_PutPlan(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the PutPlan method
func (z *Ftp2) PutPlanAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_PutPlanAsync(z.ckObj, cstr1, cstr2)

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


// Uploads an entire directory tree from the local filesystem to the remote FTP
// server, recreating the directory tree on the server. The PutTree method copies a
// directory tree to the current remote directory on the FTP server.
func (z *Ftp2) PutTree(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_PutTree(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the PutTree method
func (z *Ftp2) PutTreeAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_PutTreeAsync(z.ckObj, cstr1)

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


// Sends an arbitrary (raw) command to the FTP server.
func (z *Ftp2) Quote(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_Quote(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Quote method
func (z *Ftp2) QuoteAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_QuoteAsync(z.ckObj, cstr1)

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


// Removes a directory from the FTP server.
// 
// If the remoteDirPath contains non-English characters, it may be necessary to set the
// DirListingCharset property equal to "utf-8". Please refer to the documentation
// for the DirListingCharset property.
//
func (z *Ftp2) RemoveRemoteDir(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_RemoveRemoteDir(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the RemoveRemoteDir method
func (z *Ftp2) RemoveRemoteDirAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_RemoveRemoteDirAsync(z.ckObj, cstr1)

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


// Renames a file or directory on the FTP server. To move a file from one directory
// to another on a remote FTP server, call this method and include the source and
// destination directory filepath.
// 
// If the existingRemoteFilePath or newRemoteFilePath contains non-English characters, it may be necessary to set
// the DirListingCharset property equal to "utf-8". Please refer to the
// documentation for the DirListingCharset property.
//
func (z *Ftp2) RenameRemoteFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_RenameRemoteFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the RenameRemoteFile method
func (z *Ftp2) RenameRemoteFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_RenameRemoteFileAsync(z.ckObj, cstr1, cstr2)

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


// Sends an raw command to the FTP server and returns the raw response.
// return a string or nil if failed.
func (z *Ftp2) SendCommand(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkFtp2_sendCommand(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the SendCommand method
func (z *Ftp2) SendCommandAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_SendCommandAsync(z.ckObj, cstr1)

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


// Chilkat FTP2 supports MODE Z, which is a transfer mode implemented by some FTP
// servers. It allows for files to be uploaded and downloaded using compressed
// streams (using the zlib deflate algorithm).
// 
// Call this method after connecting to enable Mode Z. Once enabled, all transfers
// (uploads, downloads, and directory listings) are compressed.
//
func (z *Ftp2) SetModeZ() bool {

    v := C.CkFtp2_SetModeZ(z.ckObj)


    return v != 0
}


// Asynchronous version of the SetModeZ method
func (z *Ftp2) SetModeZAsync(c chan *Task) {

    hTask := C.CkFtp2_SetModeZAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Used in conjunction with the DownloadTree method. Call this method prior to
// calling DownloadTree to set the oldest date for a file to be downloaded. When
// DownloadTree is called, any file older than this date will not be downloaded.
// 
// The oldestDateTimeStr should be a date/time string in RFC822 format, such as "Tue, 25 Sep
// 2012 12:25:32 -0500".
//
func (z *Ftp2) SetOldestDateStr(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkFtp2_SetOldestDateStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// This is a general purpose method to set miscellaneous options that might arise
// due to buggy or quirky FTP servers. The option is a string describing the option.
// The current list of possible options are:
//     "Microsoft-TLS-1.2-Workaround" -- This is to force the data connection to
//     use TLS 1.0 instead of the default. It works around the Microsoft FTP server bug
//     found here: https://support.microsoft.com/en-us/kb/2888853
// 
// To turn off an option, prepend the string "No-". For example
// "No-Microsoft-TLS-1.2-Workaround". All options are turned off by default.
//
func (z *Ftp2) SetOption(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_SetOption(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the last-modified date/time of a file on the FTP server. The dateTimeStr should be
// a date/time string in RFC822 format, such as "Tue, 25 Sep 2012 12:25:32 -0500".
// 
// Important: Not all FTP servers support this functionality. Please see the
// information at the Chilkat blog below:
//
func (z *Ftp2) SetRemoteFileDateTimeStr(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_SetRemoteFileDateTimeStr(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SetRemoteFileDateTimeStr method
func (z *Ftp2) SetRemoteFileDateTimeStrAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_SetRemoteFileDateTimeStrAsync(z.ckObj, cstr1, cstr2)

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


// Sets the last-modified date/time of a file on the FTP server. Important: Not all
// FTP servers support this functionality. Please see the information at the
// Chilkat blog below:
func (z *Ftp2) SetRemoteFileDt(arg1 *CkDateTime, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_SetRemoteFileDt(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the SetRemoteFileDt method
func (z *Ftp2) SetRemoteFileDtAsync(arg1 *CkDateTime, arg2 string, c chan *Task) {
    cstr2 := C.CString(arg2)

    hTask := C.CkFtp2_SetRemoteFileDtAsync(z.ckObj, arg1.ckObj, cstr2)

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


// Sets the password in a more secure way than setting the Password property.
// Calling this method is the equivalent of setting the Password property.
func (z *Ftp2) SetSecurePassword(arg1 *SecureString) bool {

    v := C.CkFtp2_SetSecurePassword(z.ckObj, arg1.ckObj)


    return v != 0
}


// Enforces a requirement on the FTP server's certificate. The reqName can be
// "SubjectDN", "SubjectCN", "IssuerDN", or "IssuerCN". The reqName specifies the part
// of the certificate, and the reqValue is the value that it must match (exactly). If
// the FTP server's certificate does not match, the SSL / TLS connection is
// aborted.
func (z *Ftp2) SetSslCertRequirement(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkFtp2_SetSslCertRequirement(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Allows for a client-side certificate to be used for the SSL / TLS connection.
func (z *Ftp2) SetSslClientCert(arg1 *Cert) bool {

    v := C.CkFtp2_SetSslClientCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Allows for a client-side certificate to be used for the SSL / TLS connection. If
// the PEM requires no password, pass an empty string in pemPassword. If the PEM is in a
// file, pass the path to the file in pemDataOrFilename. If the PEM is already loaded into a
// string variable, then pass the string containing the contents of the PEM in
// pemDataOrFilename.
func (z *Ftp2) SetSslClientCertPem(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_SetSslClientCertPem(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Allows for a client-side certificate to be used for the SSL / TLS connection.
func (z *Ftp2) SetSslClientCertPfx(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkFtp2_SetSslClientCertPfx(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Set the FTP transfer mode to us-ascii.
func (z *Ftp2) SetTypeAscii() bool {

    v := C.CkFtp2_SetTypeAscii(z.ckObj)


    return v != 0
}


// Asynchronous version of the SetTypeAscii method
func (z *Ftp2) SetTypeAsciiAsync(c chan *Task) {

    hTask := C.CkFtp2_SetTypeAsciiAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Set the FTP transfer mode to binary.
func (z *Ftp2) SetTypeBinary() bool {

    v := C.CkFtp2_SetTypeBinary(z.ckObj)


    return v != 0
}


// Asynchronous version of the SetTypeBinary method
func (z *Ftp2) SetTypeBinaryAsync(c chan *Task) {

    hTask := C.CkFtp2_SetTypeBinaryAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends an arbitrary "site" command to the FTP server. The params argument should
// contain the parameters to the site command as they would appear on a command
// line. For example: "recfm=fb lrecl=600".
func (z *Ftp2) Site(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_Site(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the Site method
func (z *Ftp2) SiteAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_SiteAsync(z.ckObj, cstr1)

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


// Causes the calling process to sleep for a number of milliseconds.
func (z *Ftp2) SleepMs(arg1 int)  {

    C.CkFtp2_SleepMs(z.ckObj, C.int(arg1))



}


// Sends a STAT command to the FTP server and returns the server's reply.
// return a string or nil if failed.
func (z *Ftp2) Stat() *string {

    retStr := C.CkFtp2_stat(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the Stat method
func (z *Ftp2) StatAsync(c chan *Task) {

    hTask := C.CkFtp2_StatAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Delete remote files that do not exist locally. The remote directory tree rooted
// at the current remote directory is traversed and remote files that have no
// corresponding local file are deleted.
// 
// Note: In v9.5.0.51 and higher, the list of deleted files is available in the
// SyncedFiles property.
//
func (z *Ftp2) SyncDeleteRemote(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_SyncDeleteRemote(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SyncDeleteRemote method
func (z *Ftp2) SyncDeleteRemoteAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_SyncDeleteRemoteAsync(z.ckObj, cstr1)

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


// The same as SyncLocalTree, except the sub-directories are not traversed. The
// files in the current remote directory are synchronized (downloaded) with the
// files in localRoot. For possible mode settings, see SyncLocalTree.
// 
// Note: In v9.5.0.51 and higher, the list of downloaded files is available in the
// SyncedFiles property.
//
func (z *Ftp2) SyncLocalDir(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_SyncLocalDir(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SyncLocalDir method
func (z *Ftp2) SyncLocalDirAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_SyncLocalDirAsync(z.ckObj, cstr1, C.int(arg2))

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


// Downloads files from the FTP server to a local directory tree. Synchronization
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
//     * There is no mode #4. It is a mode used internally by the DirTreeXml
//     method.
//     
// 
// Note: In v9.5.0.51 and higher, the list of downloaded (or deleted) files is
// available in the SyncedFiles property.
//
func (z *Ftp2) SyncLocalTree(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_SyncLocalTree(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SyncLocalTree method
func (z *Ftp2) SyncLocalTreeAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_SyncLocalTreeAsync(z.ckObj, cstr1, C.int(arg2))

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


// Uploads a directory tree from the local filesystem to the FTP server.
// Synchronization modes include:
// 
//     mode=0: Upload all files
//     mode=1: Upload all files that do not exist on the FTP server.
//     mode=2: Upload newer or non-existant files.
//     mode=3: Upload only newer files. If a file does not already exist on the FTP
//     server, it is not uploaded.
//     mode=4: transfer missing files or files with size differences.
//     mode=5: same as mode 4, but also newer files.
// 
// Note: In v9.5.0.51 and higher, the list of uploaded files is available in the
// SyncedFiles property.
//
func (z *Ftp2) SyncRemoteTree(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_SyncRemoteTree(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SyncRemoteTree method
func (z *Ftp2) SyncRemoteTreeAsync(arg1 string, arg2 int, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkFtp2_SyncRemoteTreeAsync(z.ckObj, cstr1, C.int(arg2))

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


// Same as SyncRemoteTree, except two extra arguments are added to allow for more
// flexibility. If bDescend is false, then the directory tree is not descended and
// only the files in localDirPath are synchronized. If bPreviewOnly is true then no files are
// transferred and instead the files that would've been transferred (had bPreviewOnly been
// set to false) are listed in the SyncPreview property.
// 
// Note: If bPreviewOnly is set to true, the remote directories (if they do not exist)
// are created. It is only the files that are not uploaded.
// 
// Note: In v9.5.0.51 and higher, the list of uploaded files is available in the
// SyncedFiles property.
//
func (z *Ftp2) SyncRemoteTree2(arg1 string, arg2 int, arg3 bool, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkFtp2_SyncRemoteTree2(z.ckObj, cstr1, C.int(arg2), b3, b4)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SyncRemoteTree2 method
func (z *Ftp2) SyncRemoteTree2Async(arg1 string, arg2 int, arg3 bool, arg4 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    hTask := C.CkFtp2_SyncRemoteTree2Async(z.ckObj, cstr1, C.int(arg2), b3, b4)

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


// Sends a SYST command to the FTP server to find out the type of operating system
// at the server. The method returns the FTP server's response string. Refer to RFC
// 959 for details.
// return a string or nil if failed.
func (z *Ftp2) Syst() *string {

    retStr := C.CkFtp2_syst(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the Syst method
func (z *Ftp2) SystAsync(c chan *Task) {

    hTask := C.CkFtp2_SystAsync(z.ckObj)


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
// method. A purchased unlock code for FTP2 should contain the substring "FTP", or
// can be a Bundle unlock code.
func (z *Ftp2) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkFtp2_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


