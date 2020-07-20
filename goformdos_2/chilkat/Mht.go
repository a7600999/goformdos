// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkMht.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewMht() *Mht {
	return &Mht{C.CkMht_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Mht) DisposeMht() {
    if z != nil {
	C.CkMht_Dispose(z.ckObj)
	}
}




func (z *Mht) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkMht_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Mht) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkMht_setExternalProgress(z.ckObj,1)
}

func (z *Mht) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkMht_setExternalProgress(z.ckObj,1)
}

func (z *Mht) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkMht_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Mht) AbortCurrent() bool {
    v := int(C.CkMht_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Mht) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putAbortCurrent(z.ckObj,v)
}

// property: BaseUrl
// When processing an HTML file or string (not a website URL), this defines the
// base URL to be used when converting relative HREFs to absolute HREFs.
func (z *Mht) BaseUrl() string {
    return C.GoString(C.CkMht_baseUrl(z.ckObj))
}
// property setter: BaseUrl
func (z *Mht) SetBaseUrl(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putBaseUrl(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ConnectTimeout
// The amount of time in seconds to wait before timing out when connecting to an
// HTTP server. The default value is 10 seconds.
func (z *Mht) ConnectTimeout() int {
    return int(C.CkMht_getConnectTimeout(z.ckObj))
}
// property setter: ConnectTimeout
func (z *Mht) SetConnectTimeout(value int) {
    C.CkMht_putConnectTimeout(z.ckObj,C.int(value))
}

// property: DebugHtmlAfter
// A filename to save the result HTML when converting a URL, file, or HTML string.
// If problems are experienced, the before/after HTML can be analyzed to help
// determine the cause.
func (z *Mht) DebugHtmlAfter() string {
    return C.GoString(C.CkMht_debugHtmlAfter(z.ckObj))
}
// property setter: DebugHtmlAfter
func (z *Mht) SetDebugHtmlAfter(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putDebugHtmlAfter(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DebugHtmlBefore
// A filename to save the input HTML when converting a URL, file, or HTML string.
// If problems are experienced, the before/after HTML can be analyzed to help
// determine the cause.
func (z *Mht) DebugHtmlBefore() string {
    return C.GoString(C.CkMht_debugHtmlBefore(z.ckObj))
}
// property setter: DebugHtmlBefore
func (z *Mht) SetDebugHtmlBefore(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putDebugHtmlBefore(z.ckObj,cStr)
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
func (z *Mht) DebugLogFilePath() string {
    return C.GoString(C.CkMht_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Mht) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DebugTagCleaning
// When true causes the Mht class to be much more verbose in its logging. The
// default is false.
func (z *Mht) DebugTagCleaning() bool {
    v := int(C.CkMht_getDebugTagCleaning(z.ckObj))
    return v != 0
}
// property setter: DebugTagCleaning
func (z *Mht) SetDebugTagCleaning(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putDebugTagCleaning(z.ckObj,v)
}

// property: EmbedImages
// Controls whether images are embedded in the MHT/EML, or whether the IMG SRC
// attributes are left as external URL references. If false, the IMG SRC tags are
// converted to absolute URLs (if necessary) and the images are not embedded within
// the MHT/EML.
func (z *Mht) EmbedImages() bool {
    v := int(C.CkMht_getEmbedImages(z.ckObj))
    return v != 0
}
// property setter: EmbedImages
func (z *Mht) SetEmbedImages(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putEmbedImages(z.ckObj,v)
}

// property: EmbedLocalOnly
// If true, only images found on the local filesystem (i.e. links to files) will
// be embedded within the MHT.
func (z *Mht) EmbedLocalOnly() bool {
    v := int(C.CkMht_getEmbedLocalOnly(z.ckObj))
    return v != 0
}
// property setter: EmbedLocalOnly
func (z *Mht) SetEmbedLocalOnly(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putEmbedLocalOnly(z.ckObj,v)
}

// property: FetchFromCache
// If true, page parts such as images, style sheets, etc. will be fetched from
// the disk cache if possible. The disk cache root may be defined by calling
// AddCacheRoot. The default value is false.
func (z *Mht) FetchFromCache() bool {
    v := int(C.CkMht_getFetchFromCache(z.ckObj))
    return v != 0
}
// property setter: FetchFromCache
func (z *Mht) SetFetchFromCache(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putFetchFromCache(z.ckObj,v)
}

// property: HeartbeatMs
// The time interval, in milliseconds, between AbortCheck event callbacks. The
// heartbeat/AbortCheck provides a means for an application to abort any MHT method
// before completion.
// 
// The default value is 0, which means that no AbortCheck events will be fired.
//
func (z *Mht) HeartbeatMs() int {
    return int(C.CkMht_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Mht) SetHeartbeatMs(value int) {
    C.CkMht_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: IgnoreMustRevalidate
// Some HTTP responses contain a "Cache-Control: must-revalidate" header. If this
// is present, the server is requesting that the client always issue a revalidate
// HTTP request instead of serving the page directly from cache. If
// IgnoreMustRevalidate is set to true, then Chilkat MHT will serve the page
// directly from cache without revalidating until the page is no longer fresh.
// (assuming that FetchFromCache is set to true)
// 
// The default value of this property is false.
//
func (z *Mht) IgnoreMustRevalidate() bool {
    v := int(C.CkMht_getIgnoreMustRevalidate(z.ckObj))
    return v != 0
}
// property setter: IgnoreMustRevalidate
func (z *Mht) SetIgnoreMustRevalidate(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putIgnoreMustRevalidate(z.ckObj,v)
}

// property: IgnoreNoCache
// Some HTTP responses contain headers of various types that indicate that the page
// should not be cached. Chilkat MHT will adhere to this unless this property is
// set to true.
// 
// The default value of this property is false.
//
func (z *Mht) IgnoreNoCache() bool {
    v := int(C.CkMht_getIgnoreNoCache(z.ckObj))
    return v != 0
}
// property setter: IgnoreNoCache
func (z *Mht) SetIgnoreNoCache(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putIgnoreNoCache(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Mht) LastErrorHtml() string {
    return C.GoString(C.CkMht_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Mht) LastErrorText() string {
    return C.GoString(C.CkMht_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Mht) LastErrorXml() string {
    return C.GoString(C.CkMht_lastErrorXml(z.ckObj))
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
func (z *Mht) LastMethodSuccess() bool {
    v := int(C.CkMht_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Mht) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putLastMethodSuccess(z.ckObj,v)
}

// property: NoScripts
// Only applies when creating MHT files. Scripts are always removed when creating
// EML or emails from HTML. If set to true, then all scripts are removed, if set
// to false (the default) then scripts are not removed.
func (z *Mht) NoScripts() bool {
    v := int(C.CkMht_getNoScripts(z.ckObj))
    return v != 0
}
// property setter: NoScripts
func (z *Mht) SetNoScripts(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putNoScripts(z.ckObj,v)
}

// property: NtlmAuth
// Setting this property to true causes the MHT component to use NTLM
// authentication (also known as IWA -- or Integrated Windows Authentication) when
// authentication with an HTTP server.
// 
// The default value of this property is false.
//
func (z *Mht) NtlmAuth() bool {
    v := int(C.CkMht_getNtlmAuth(z.ckObj))
    return v != 0
}
// property setter: NtlmAuth
func (z *Mht) SetNtlmAuth(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putNtlmAuth(z.ckObj,v)
}

// property: NumCacheLevels
// The number of directory levels to be used under each cache root. The default is
// 0, meaning that each cached item is stored in a cache root directory. A value of
// 1 causes each cached page to be stored in one of 255 subdirectories named
// "0","1", "2", ..."255" under a cache root. A value of 2 causes two levels of
// subdirectories ("0..255/0..255") under each cache root. The MHT control
// automatically creates subdirectories as needed. The reason for mutliple levels
// is to alleviate problems that may arise when huge numbers of files are stored in
// a single directory. For example, Windows Explorer does not behave well when
// trying to display the contents of directories with thousands of files.
func (z *Mht) NumCacheLevels() int {
    return int(C.CkMht_getNumCacheLevels(z.ckObj))
}
// property setter: NumCacheLevels
func (z *Mht) SetNumCacheLevels(value int) {
    C.CkMht_putNumCacheLevels(z.ckObj,C.int(value))
}

// property: NumCacheRoots
// The number of cache roots to be used for the disk cache. This allows the disk
// cache spread out over multiple disk drives. Each cache root is a string
// indicating the drive letter and directory path. For example, "E:\Cache". To
// create a cache with four roots, call AddCacheRoot once for each directory root.
func (z *Mht) NumCacheRoots() int {
    return int(C.CkMht_getNumCacheRoots(z.ckObj))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *Mht) PreferIpv6() bool {
    v := int(C.CkMht_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *Mht) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putPreferIpv6(z.ckObj,v)
}

// property: PreferMHTScripts
// This property provides a means for the noscript option to be selected when
// possible. If PreferMHTScripts = false, then scripts with noscript alternatives
// are removed and the noscript content is kept. If true (the default), then
// scripts are preserved and the noscript options are discarded.
func (z *Mht) PreferMHTScripts() bool {
    v := int(C.CkMht_getPreferMHTScripts(z.ckObj))
    return v != 0
}
// property setter: PreferMHTScripts
func (z *Mht) SetPreferMHTScripts(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putPreferMHTScripts(z.ckObj,v)
}

// property: Proxy
// (Optional) A proxy host:port if a proxy is necessary to access the Internet. The
// proxy string should be formatted as "hostname:port", such as
// "www.chilkatsoft.com:100".
func (z *Mht) Proxy() string {
    return C.GoString(C.CkMht_proxy(z.ckObj))
}
// property setter: Proxy
func (z *Mht) SetProxy(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putProxy(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyLogin
// If an HTTP proxy is used and it requires authentication, this property specifies
// the HTTP proxy login.
func (z *Mht) ProxyLogin() string {
    return C.GoString(C.CkMht_proxyLogin(z.ckObj))
}
// property setter: ProxyLogin
func (z *Mht) SetProxyLogin(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putProxyLogin(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyPassword
// If an HTTP proxy is used and it requires authentication, this property specifies
// the HTTP proxy password.
func (z *Mht) ProxyPassword() string {
    return C.GoString(C.CkMht_proxyPassword(z.ckObj))
}
// property setter: ProxyPassword
func (z *Mht) SetProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ReadTimeout
// The amount of time in seconds to wait before timing out when reading from an
// HTTP server. The ReadTimeout is the amount of time that needs to elapse while no
// additional data is forthcoming. During a long data transfer, if the data stream
// halts for more than this amount, it will timeout. Otherwise, there is no limit
// on the length of time for the entire data transfer.
// 
// The default value is 20 seconds.
//
func (z *Mht) ReadTimeout() int {
    return int(C.CkMht_getReadTimeout(z.ckObj))
}
// property setter: ReadTimeout
func (z *Mht) SetReadTimeout(value int) {
    C.CkMht_putReadTimeout(z.ckObj,C.int(value))
}

// property: RequireSslCertVerify
// If true, then the HTTP client will verify the server's SSL certificate. The
// certificate is expired, or if the cert's signature is invalid, the connection is
// not allowed. The default value of this property is false.
func (z *Mht) RequireSslCertVerify() bool {
    v := int(C.CkMht_getRequireSslCertVerify(z.ckObj))
    return v != 0
}
// property setter: RequireSslCertVerify
func (z *Mht) SetRequireSslCertVerify(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putRequireSslCertVerify(z.ckObj,v)
}

// property: SocksHostname
// The SOCKS4/SOCKS5 hostname or IPv4 address (in dotted decimal notation). This
// property is only used if the SocksVersion property is set to 4 or 5).
func (z *Mht) SocksHostname() string {
    return C.GoString(C.CkMht_socksHostname(z.ckObj))
}
// property setter: SocksHostname
func (z *Mht) SetSocksHostname(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putSocksHostname(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPassword
// The SOCKS5 password (if required). The SOCKS4 protocol does not include the use
// of a password, so this does not apply to SOCKS4.
func (z *Mht) SocksPassword() string {
    return C.GoString(C.CkMht_socksPassword(z.ckObj))
}
// property setter: SocksPassword
func (z *Mht) SetSocksPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putSocksPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksPort
// The SOCKS4/SOCKS5 proxy port. The default value is 1080. This property only
// applies if a SOCKS proxy is used (if the SocksVersion property is set to 4 or
// 5).
func (z *Mht) SocksPort() int {
    return int(C.CkMht_getSocksPort(z.ckObj))
}
// property setter: SocksPort
func (z *Mht) SetSocksPort(value int) {
    C.CkMht_putSocksPort(z.ckObj,C.int(value))
}

// property: SocksUsername
// The SOCKS4/SOCKS5 proxy username. This property is only used if the SocksVersion
// property is set to 4 or 5).
func (z *Mht) SocksUsername() string {
    return C.GoString(C.CkMht_socksUsername(z.ckObj))
}
// property setter: SocksUsername
func (z *Mht) SetSocksUsername(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putSocksUsername(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SocksVersion
// SocksVersion May be set to one of the following integer values:
// 
// 0 - No SOCKS proxy is used. This is the default.
// 4 - Connect via a SOCKS4 proxy.
// 5 - Connect via a SOCKS5 proxy.
//
func (z *Mht) SocksVersion() int {
    return int(C.CkMht_getSocksVersion(z.ckObj))
}
// property setter: SocksVersion
func (z *Mht) SetSocksVersion(value int) {
    C.CkMht_putSocksVersion(z.ckObj,C.int(value))
}

// property: UnpackDirect
// If true, then the UnpackMHT and UnpackMHTString methods will unpack the MHT
// directly with no transformations. Normally, the related parts are unpacked to a
// "parts" sub-directory, and the unpacked HTML is edited to update references to
// point to the unpacked image and script files. When unpacking direct, the HTML is
// not edited, and the related parts are unpacked to sub-directories rooted in the
// directory where HTML file is created (i.e. the unpack directory). When unpacking
// direct, the "partsSubDir" argument of the UnpackMHT* methods is unused.
// 
// Note: It is only possible to directly unpack MHT files where the
// Content-Location headers DO NOT contain URLs. The MHT must be such that the
// Content-Location headers of the related items contain relative paths.
// 
// Note: The default value of this property is false.
//
func (z *Mht) UnpackDirect() bool {
    v := int(C.CkMht_getUnpackDirect(z.ckObj))
    return v != 0
}
// property setter: UnpackDirect
func (z *Mht) SetUnpackDirect(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putUnpackDirect(z.ckObj,v)
}

// property: UnpackUseRelPaths
// Controls whether absolute or relative paths are used when referencing images in
// the unpacked HTML. The default value is true indicating that relative paths
// will be used. To use absolute paths, set this property value equal to false.
func (z *Mht) UnpackUseRelPaths() bool {
    v := int(C.CkMht_getUnpackUseRelPaths(z.ckObj))
    return v != 0
}
// property setter: UnpackUseRelPaths
func (z *Mht) SetUnpackUseRelPaths(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putUnpackUseRelPaths(z.ckObj,v)
}

// property: UpdateCache
// Controls whether the cache is automatically updated with the responses from HTTP
// GET requests. If true, the disk cache is updated, if false (the default),
// the cache is not updated.
func (z *Mht) UpdateCache() bool {
    v := int(C.CkMht_getUpdateCache(z.ckObj))
    return v != 0
}
// property setter: UpdateCache
func (z *Mht) SetUpdateCache(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putUpdateCache(z.ckObj,v)
}

// property: UseCids
// Controls whether CID URLs are used for embedded references when generating MHT
// or EML documents. If UseCids is false, then URLs are left unchanged and the
// embedded items will contain "content-location" headers that match the URLs in
// the HTML. If true, CIDs are generated and the URLs within the HTML are
// replaced with "CID:" links.
// 
// The default value of this property is true.
//
func (z *Mht) UseCids() bool {
    v := int(C.CkMht_getUseCids(z.ckObj))
    return v != 0
}
// property setter: UseCids
func (z *Mht) SetUseCids(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putUseCids(z.ckObj,v)
}

// property: UseFilename
// If true, a "filename" attribute is added to each Content-Disposition MIME
// header field for each embedded item (image, style sheet, etc.). If false, then
// no filename attribute is added.
// 
// The default value of this property is true.
//
func (z *Mht) UseFilename() bool {
    v := int(C.CkMht_getUseFilename(z.ckObj))
    return v != 0
}
// property setter: UseFilename
func (z *Mht) SetUseFilename(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putUseFilename(z.ckObj,v)
}

// property: UseIEProxy
// If true, the proxy host/port used by Internet Explorer will also be used by
// Chilkat MHT.
func (z *Mht) UseIEProxy() bool {
    v := int(C.CkMht_getUseIEProxy(z.ckObj))
    return v != 0
}
// property setter: UseIEProxy
func (z *Mht) SetUseIEProxy(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putUseIEProxy(z.ckObj,v)
}

// property: UseInline
// If true, an "inline" attribute is added to each Content-Disposition MIME
// header field for each embedded item (image, style sheet, etc.). If false, then
// no inline attribute is added.
// 
// The default value of this property is true.
//
func (z *Mht) UseInline() bool {
    v := int(C.CkMht_getUseInline(z.ckObj))
    return v != 0
}
// property setter: UseInline
func (z *Mht) SetUseInline(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putUseInline(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Mht) VerboseLogging() bool {
    v := int(C.CkMht_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Mht) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMht_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Mht) Version() string {
    return C.GoString(C.CkMht_version(z.ckObj))
}

// property: WebSiteLogin
// (Optional) Specifies the login if a a Web page is accessed that requires a login
func (z *Mht) WebSiteLogin() string {
    return C.GoString(C.CkMht_webSiteLogin(z.ckObj))
}
// property setter: WebSiteLogin
func (z *Mht) SetWebSiteLogin(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putWebSiteLogin(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: WebSiteLoginDomain
// The optional domain name to be used with NTLM / Kerberos / Negotiate
// authentication.
func (z *Mht) WebSiteLoginDomain() string {
    return C.GoString(C.CkMht_webSiteLoginDomain(z.ckObj))
}
// property setter: WebSiteLoginDomain
func (z *Mht) SetWebSiteLoginDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putWebSiteLoginDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: WebSitePassword
// Optional) Specifies the password if a a Web page is accessed that requires a
// login and password
func (z *Mht) WebSitePassword() string {
    return C.GoString(C.CkMht_webSitePassword(z.ckObj))
}
// property setter: WebSitePassword
func (z *Mht) SetWebSitePassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkMht_putWebSitePassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}
// If disk caching is used, this must be called once for each cache root. For
// example, if the cache is spread across D:\cacheRoot, E:\cacheRoot, and
// F:\cacheRoot, an application would setup the cache object by calling AddRoot
// three times -- once with "D:\cacheRoot", once with "E:\cacheRoot", and once with
// "F:\cacheRoot".
func (z *Mht) AddCacheRoot(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkMht_AddCacheRoot(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Adds a custom HTTP header to all HTTP requests sent by the MHT component. To add
// multiple header fields, call this method once for each custom header.
func (z *Mht) AddCustomHeader(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkMht_AddCustomHeader(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// (This method rarely needs to be called.) Includes an additional style sheet that
// would not normally be included with the HTML. This method is provided for cases
// when style sheet names are constructed and dynamically included in Javascript
// such that MHT .NET cannot know beforehand what stylesheet to embed. MHT .NET by
// default downloads and embeds all stylesheets externally referenced by the HTML
func (z *Mht) AddExternalStyleSheet(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkMht_AddExternalStyleSheet(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Removes all custom headers that may have accumulated from previous calls to
// AddCustomHeader.
func (z *Mht) ClearCustomHeaders()  {

    C.CkMht_ClearCustomHeaders(z.ckObj)



}


// (This method rarely needs to be called.) Tells Chilkat MHT .NET to not embed any
// images whose URL matches a pattern. Sometimes images can be referenced within
// style sheets and not actually used when rendering the page. In cases like those,
// the image will appear as an attachment in the HTML email. This feature allows
// you to explicitly remove those images from the email so no attachments appear.
func (z *Mht) ExcludeImagesMatching(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkMht_ExcludeImagesMatching(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Creates an EML file from a web page or HTML file. All external images and style
// sheets are downloaded and embedded in the EML file.
func (z *Mht) GetAndSaveEML(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMht_GetAndSaveEML(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the GetAndSaveEML method
func (z *Mht) GetAndSaveEMLAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkMht_GetAndSaveEMLAsync(z.ckObj, cstr1, cstr2)

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


// Creates an MHT file from a web page or local HTML file. All external images,
// scripts, and style sheets are downloaded and embedded in the MHT file.
func (z *Mht) GetAndSaveMHT(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMht_GetAndSaveMHT(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the GetAndSaveMHT method
func (z *Mht) GetAndSaveMHTAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkMht_GetAndSaveMHTAsync(z.ckObj, cstr1, cstr2)

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


// Creates an EML file from a web page or HTML file, compresses, and appends to a
// new or existing Zip file. All external images and style sheets are downloaded
// and embedded in the EML.
func (z *Mht) GetAndZipEML(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkMht_GetAndZipEML(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the GetAndZipEML method
func (z *Mht) GetAndZipEMLAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkMht_GetAndZipEMLAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Creates an MHT file from a web page or HTML file, compresses, and appends to a
// new or existing Zip file. All external images and style sheets are downloaded
// and embedded in the MHT.
func (z *Mht) GetAndZipMHT(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkMht_GetAndZipMHT(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the GetAndZipMHT method
func (z *Mht) GetAndZipMHTAsync(arg1 string, arg2 string, arg3 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    hTask := C.CkMht_GetAndZipMHTAsync(z.ckObj, cstr1, cstr2, cstr3)

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


// Returns the Nth cache root (indexing begins at 0). Cache roots are set by
// calling AddCacheRoot one or more times.
// return a string or nil if failed.
func (z *Mht) GetCacheRoot(arg1 int) *string {

    retStr := C.CkMht_getCacheRoot(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Creates EML from a web page or HTML file, and returns the EML (MIME) message
// data as a string.
// return a string or nil if failed.
func (z *Mht) GetEML(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkMht_getEML(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetEML method
func (z *Mht) GetEMLAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMht_GetEMLAsync(z.ckObj, cstr1)

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


// Creates MHT from a web page or local HTML file, and returns the MHT (MIME)
// message data as a string
// return a string or nil if failed.
func (z *Mht) GetMHT(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkMht_getMHT(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the GetMHT method
func (z *Mht) GetMHTAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMht_GetMHTAsync(z.ckObj, cstr1)

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


// Creates an in-memory EML string from an in-memory HTML string. All external
// images and style sheets are downloaded and embedded in the EML string that is
// returned.
// return a string or nil if failed.
func (z *Mht) HtmlToEML(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkMht_htmlToEML(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the HtmlToEML method
func (z *Mht) HtmlToEMLAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMht_HtmlToEMLAsync(z.ckObj, cstr1)

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


// Creates an EML file from an in-memory HTML string. All external images and style
// sheets are downloaded and embedded in the EML file.
func (z *Mht) HtmlToEMLFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMht_HtmlToEMLFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the HtmlToEMLFile method
func (z *Mht) HtmlToEMLFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkMht_HtmlToEMLFileAsync(z.ckObj, cstr1, cstr2)

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


// Creates an in-memory MHT web archive from an in-memory HTML string. All external
// images and style sheets are downloaded and embedded in the MHT string.
// return a string or nil if failed.
func (z *Mht) HtmlToMHT(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkMht_htmlToMHT(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the HtmlToMHT method
func (z *Mht) HtmlToMHTAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkMht_HtmlToMHTAsync(z.ckObj, cstr1)

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


// Creates an MHT file from an in-memory HTML string. All external images and style
// sheets are downloaded and embedded in the MHT file.
func (z *Mht) HtmlToMHTFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMht_HtmlToMHTFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the HtmlToMHTFile method
func (z *Mht) HtmlToMHTFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkMht_HtmlToMHTFileAsync(z.ckObj, cstr1, cstr2)

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


// Returns true if the MHT component is unlocked.
func (z *Mht) IsUnlocked() bool {

    v := C.CkMht_IsUnlocked(z.ckObj)


    return v != 0
}


// Loads the caller of the task's async method.
func (z *Mht) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkMht_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Removes a custom header by header field name.
func (z *Mht) RemoveCustomHeader(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkMht_RemoveCustomHeader(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Restores the default property settings.
func (z *Mht) RestoreDefaults()  {

    C.CkMht_RestoreDefaults(z.ckObj)



}


// Unlocks the component allowing for the full functionality to be used. Returns
// true if the unlock code is valid.
func (z *Mht) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMht_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unpacks the contents of a MHT file. The destination directory is specified by
// unpackDir. The name of the HTML file created is specified by htmlFilename, and supporting
// files (images, javascripts, etc.) are created in partsSubDir, which is automatically
// created if it does not already exist.
func (z *Mht) UnpackMHT(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkMht_UnpackMHT(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Same as UnpackMHT, except the MHT is passed in as an in-memory string.
func (z *Mht) UnpackMHTString(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkMht_UnpackMHTString(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


