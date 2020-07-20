// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkSpider.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewSpider() *Spider {
	return &Spider{C.CkSpider_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Spider) DisposeSpider() {
    if z != nil {
	C.CkSpider_Dispose(z.ckObj)
	}
}




func (z *Spider) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkSpider_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Spider) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkSpider_setExternalProgress(z.ckObj,1)
}

func (z *Spider) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkSpider_setExternalProgress(z.ckObj,1)
}

func (z *Spider) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkSpider_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Spider) AbortCurrent() bool {
    v := int(C.CkSpider_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Spider) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSpider_putAbortCurrent(z.ckObj,v)
}

// property: AvoidHttps
// If set the 1 (true) the spider will avoid all HTTPS URLs. The default is 0
// (false).
func (z *Spider) AvoidHttps() bool {
    v := int(C.CkSpider_getAvoidHttps(z.ckObj))
    return v != 0
}
// property setter: AvoidHttps
func (z *Spider) SetAvoidHttps(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSpider_putAvoidHttps(z.ckObj,v)
}

// property: CacheDir
// Specifies a cache directory to use for spidering. If either of the
// FetchFromCache or UpdateCache properties are true, this is the location of the
// cache to be used. Note: the Internet Explorer, Netscape, and FireFox caches are
// completely separate from the Chilkat Spider cache directory. You should specify
// a new and empty directory.
func (z *Spider) CacheDir() string {
    return C.GoString(C.CkSpider_cacheDir(z.ckObj))
}
// property setter: CacheDir
func (z *Spider) SetCacheDir(goStr string) {
    cStr := C.CString(goStr)
    C.CkSpider_putCacheDir(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ChopAtQuery
// If equal to 1 (true), then the query portion of all URLs are automatically
// removed when adding to the unspidered list. The default value is 0 (false).
func (z *Spider) ChopAtQuery() bool {
    v := int(C.CkSpider_getChopAtQuery(z.ckObj))
    return v != 0
}
// property setter: ChopAtQuery
func (z *Spider) SetChopAtQuery(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSpider_putChopAtQuery(z.ckObj,v)
}

// property: ConnectTimeout
// The maximum number of seconds to wait while connecting to an HTTP server.
func (z *Spider) ConnectTimeout() int {
    return int(C.CkSpider_getConnectTimeout(z.ckObj))
}
// property setter: ConnectTimeout
func (z *Spider) SetConnectTimeout(value int) {
    C.CkSpider_putConnectTimeout(z.ckObj,C.int(value))
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
func (z *Spider) DebugLogFilePath() string {
    return C.GoString(C.CkSpider_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Spider) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkSpider_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Domain
// The domain name that is being spidered. This is the domain previously set in the
// Initialize method.
func (z *Spider) Domain() string {
    return C.GoString(C.CkSpider_domain(z.ckObj))
}

// property: FetchFromCache
// If equal to 1 (true) then pages are fetched from cache when possible. If 0, the
// cache is ignored. The default value is 1. Regardless, if no CacheDir is set then
// the cache is not used.
func (z *Spider) FetchFromCache() bool {
    v := int(C.CkSpider_getFetchFromCache(z.ckObj))
    return v != 0
}
// property setter: FetchFromCache
func (z *Spider) SetFetchFromCache(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSpider_putFetchFromCache(z.ckObj,v)
}

// property: HeartbeatMs
// The number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any method call prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
func (z *Spider) HeartbeatMs() int {
    return int(C.CkSpider_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Spider) SetHeartbeatMs(value int) {
    C.CkSpider_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Spider) LastErrorHtml() string {
    return C.GoString(C.CkSpider_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Spider) LastErrorText() string {
    return C.GoString(C.CkSpider_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Spider) LastErrorXml() string {
    return C.GoString(C.CkSpider_lastErrorXml(z.ckObj))
}

// property: LastFromCache
// Equal to 1 if the last page spidered was fetched from the cache. Otherwise equal
// to 0.
func (z *Spider) LastFromCache() bool {
    v := int(C.CkSpider_getLastFromCache(z.ckObj))
    return v != 0
}

// property: LastHtml
// The HTML text of the last paged fetched by the spider.
func (z *Spider) LastHtml() string {
    return C.GoString(C.CkSpider_lastHtml(z.ckObj))
}

// property: LastHtmlDescription
// The HTML META description from the last page fetched by the spider.
func (z *Spider) LastHtmlDescription() string {
    return C.GoString(C.CkSpider_lastHtmlDescription(z.ckObj))
}

// property: LastHtmlKeywords
// The HTML META keywords from the last page fetched by the spider.
func (z *Spider) LastHtmlKeywords() string {
    return C.GoString(C.CkSpider_lastHtmlKeywords(z.ckObj))
}

// property: LastHtmlTitle
// The HTML title from the last page fetched by the spider.
func (z *Spider) LastHtmlTitle() string {
    return C.GoString(C.CkSpider_lastHtmlTitle(z.ckObj))
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
func (z *Spider) LastMethodSuccess() bool {
    v := int(C.CkSpider_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Spider) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSpider_putLastMethodSuccess(z.ckObj,v)
}

// property: LastModDateStr
// The last modification date/time from the last page fetched by the spider.
func (z *Spider) LastModDateStr() string {
    return C.GoString(C.CkSpider_lastModDateStr(z.ckObj))
}

// property: LastUrl
// The URL of the last page spidered.
func (z *Spider) LastUrl() string {
    return C.GoString(C.CkSpider_lastUrl(z.ckObj))
}

// property: MaxResponseSize
// The maximum HTTP response size allowed. The spider will automatically fail any
// pages larger than this size. The default value is 250,000 bytes.
func (z *Spider) MaxResponseSize() int {
    return int(C.CkSpider_getMaxResponseSize(z.ckObj))
}
// property setter: MaxResponseSize
func (z *Spider) SetMaxResponseSize(value int) {
    C.CkSpider_putMaxResponseSize(z.ckObj,C.int(value))
}

// property: MaxUrlLen
// The maximum URL length allowed. URLs longer than this are not added to the
// unspidered list. The default value is 200.
func (z *Spider) MaxUrlLen() int {
    return int(C.CkSpider_getMaxUrlLen(z.ckObj))
}
// property setter: MaxUrlLen
func (z *Spider) SetMaxUrlLen(value int) {
    C.CkSpider_putMaxUrlLen(z.ckObj,C.int(value))
}

// property: NumAvoidPatterns
// The number of avoid patterns previously set by calling AddAvoidPattern.
func (z *Spider) NumAvoidPatterns() int {
    return int(C.CkSpider_getNumAvoidPatterns(z.ckObj))
}

// property: NumFailed
// The number of URLs in the component's failed URL list.
func (z *Spider) NumFailed() int {
    return int(C.CkSpider_getNumFailed(z.ckObj))
}

// property: NumOutboundLinks
// The number of URLs in the component's outbound links URL list.
func (z *Spider) NumOutboundLinks() int {
    return int(C.CkSpider_getNumOutboundLinks(z.ckObj))
}

// property: NumSpidered
// The number of URLs in the component's already-spidered URL list.
func (z *Spider) NumSpidered() int {
    return int(C.CkSpider_getNumSpidered(z.ckObj))
}

// property: NumUnspidered
// The number of URLs in the component's unspidered URL list.
func (z *Spider) NumUnspidered() int {
    return int(C.CkSpider_getNumUnspidered(z.ckObj))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
func (z *Spider) PreferIpv6() bool {
    v := int(C.CkSpider_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *Spider) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSpider_putPreferIpv6(z.ckObj,v)
}

// property: ProxyDomain
// The domain name of a proxy host if an HTTP proxy is used.
func (z *Spider) ProxyDomain() string {
    return C.GoString(C.CkSpider_proxyDomain(z.ckObj))
}
// property setter: ProxyDomain
func (z *Spider) SetProxyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkSpider_putProxyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyLogin
// If an HTTP proxy is used and it requires authentication, this property specifies
// the HTTP proxy login.
func (z *Spider) ProxyLogin() string {
    return C.GoString(C.CkSpider_proxyLogin(z.ckObj))
}
// property setter: ProxyLogin
func (z *Spider) SetProxyLogin(goStr string) {
    cStr := C.CString(goStr)
    C.CkSpider_putProxyLogin(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyPassword
// If an HTTP proxy is used and it requires authentication, this property specifies
// the HTTP proxy password.
func (z *Spider) ProxyPassword() string {
    return C.GoString(C.CkSpider_proxyPassword(z.ckObj))
}
// property setter: ProxyPassword
func (z *Spider) SetProxyPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkSpider_putProxyPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ProxyPort
// The port number of a proxy server if an HTTP proxy is used.
func (z *Spider) ProxyPort() int {
    return int(C.CkSpider_getProxyPort(z.ckObj))
}
// property setter: ProxyPort
func (z *Spider) SetProxyPort(value int) {
    C.CkSpider_putProxyPort(z.ckObj,C.int(value))
}

// property: ReadTimeout
// The maximum number of seconds to wait when reading from an HTTP server.
func (z *Spider) ReadTimeout() int {
    return int(C.CkSpider_getReadTimeout(z.ckObj))
}
// property setter: ReadTimeout
func (z *Spider) SetReadTimeout(value int) {
    C.CkSpider_putReadTimeout(z.ckObj,C.int(value))
}

// property: UpdateCache
// If equal to 1 (true) then pages saved to the cache. If 0, the cache is ignored.
// The default value is 1. Regardless, if no CacheDir is set then the cache is not
// used.
func (z *Spider) UpdateCache() bool {
    v := int(C.CkSpider_getUpdateCache(z.ckObj))
    return v != 0
}
// property setter: UpdateCache
func (z *Spider) SetUpdateCache(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSpider_putUpdateCache(z.ckObj,v)
}

// property: UserAgent
// The value of the HTTP user-agent header field to be sent with HTTP requests. The
// default value is "Chilkat/1.0.0 (+http://www.chilkatsoft.com/ChilkatHttpUA.asp)"
func (z *Spider) UserAgent() string {
    return C.GoString(C.CkSpider_userAgent(z.ckObj))
}
// property setter: UserAgent
func (z *Spider) SetUserAgent(goStr string) {
    cStr := C.CString(goStr)
    C.CkSpider_putUserAgent(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Spider) VerboseLogging() bool {
    v := int(C.CkSpider_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Spider) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkSpider_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Spider) Version() string {
    return C.GoString(C.CkSpider_version(z.ckObj))
}

// property: WindDownCount
// The "wind-down" phase begins when this number of URLs has been spidered. When in
// the wind-down phase, no new URLs are added to the unspidered list. The default
// value is 0 which means that there is NO wind-down phase.
func (z *Spider) WindDownCount() int {
    return int(C.CkSpider_getWindDownCount(z.ckObj))
}
// property setter: WindDownCount
func (z *Spider) SetWindDownCount(value int) {
    C.CkSpider_putWindDownCount(z.ckObj,C.int(value))
}
// Adds a wildcarded pattern to prevent collecting matching outbound link URLs. For
// example, if "*google*" is added, then any outbound links containing the word
// "google" will be ignored. The "*" character matches zero or more of any
// character.
func (z *Spider) AddAvoidOutboundLinkPattern(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkSpider_AddAvoidOutboundLinkPattern(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Adds a wildcarded pattern to prevent spidering matching URLs. For example, if
// "*register*" is added, then any url containing the word "register" is not
// spidered. The "*" character matches zero or more of any character.
func (z *Spider) AddAvoidPattern(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkSpider_AddAvoidPattern(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Adds a wildcarded pattern to limit spidering to only URLs that match the
// pattern. For example, if "*/products/*" is added, then only URLs containing
// "/products/" are spidered. This is helpful for only spidering a portion of a
// website. The "*" character matches zero or more of any character.
func (z *Spider) AddMustMatchPattern(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkSpider_AddMustMatchPattern(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// To begin spidering you must call this method one or more times to provide
// starting points. It adds a single URL to the component's internal queue of URLs
// to be spidered.
func (z *Spider) AddUnspidered(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkSpider_AddUnspidered(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Canonicalizes a URL by doing the following:
//     Drops username/password if present.
//     Drops fragment if present.
//     Converts domain to lowercase.
//     Removes port 80 or 443
//     Remove default.asp, index.html, index.htm, default.html, index.htm,
//     default.htm, index.php, index.asp, default.php, .cfm, .aspx, ,php3, .pl, .cgi,
//     .txt, .shtml, .phtml
//     Remove www. from the domain if present.
// return a string or nil if failed.
func (z *Spider) CanonicalizeUrl(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSpider_canonicalizeUrl(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Clears the component's internal list of URLs that could not be downloaded.
func (z *Spider) ClearFailedUrls()  {

    C.CkSpider_ClearFailedUrls(z.ckObj)



}


// Clears the component's internal list of outbound URLs that will automatically
// accumulate while spidering.
func (z *Spider) ClearOutboundLinks()  {

    C.CkSpider_ClearOutboundLinks(z.ckObj)



}


// Clears the component's internal list of already-spidered URLs that will
// automatically accumulate while spidering.
func (z *Spider) ClearSpideredUrls()  {

    C.CkSpider_ClearSpideredUrls(z.ckObj)



}


// Crawls the next URL in the internal list of unspidered URLs. The URL is moved
// from the unspidered list to the spidered list. Any new links within the same
// domain and not yet spidered are added to the unspidered list. (providing that
// they do not match "avoid" patterns, etc.) Any new outbound links are added to
// the outbound URL list. If successful, the HTML of the downloaded page is
// available in the LastHtml property. If there are no more URLs left unspidered,
// the method returns false. Information about the URL crawled is available in
// the properties LastUrl, LastFromCache, and LastModDate.
func (z *Spider) CrawlNext() bool {

    v := C.CkSpider_CrawlNext(z.ckObj)


    return v != 0
}


// Asynchronous version of the CrawlNext method
func (z *Spider) CrawlNextAsync(c chan *Task) {

    hTask := C.CkSpider_CrawlNextAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the contents of the robots.txt file from the domain being crawled. This
// spider component will not crawl URLs excluded by robots.txt. If you believe the
// spider is not behaving correctly, please notify us at support@chilkatsoft.com
// and provide information detailing a case that allows us to reproduce the
// problem.
// return a string or nil if failed.
func (z *Spider) FetchRobotsText() *string {

    retStr := C.CkSpider_fetchRobotsText(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the FetchRobotsText method
func (z *Spider) FetchRobotsTextAsync(c chan *Task) {

    hTask := C.CkSpider_FetchRobotsTextAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the Nth avoid pattern previously added by calling AddAvoidPattern.
// Indexing begins at 0.
// return a string or nil if failed.
func (z *Spider) GetAvoidPattern(arg1 int) *string {

    retStr := C.CkSpider_getAvoidPattern(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// To be documented soon.
// return a string or nil if failed.
func (z *Spider) GetBaseDomain(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSpider_getBaseDomain(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth URL in the failed URL list. Indexing begins at 0.
// return a string or nil if failed.
func (z *Spider) GetFailedUrl(arg1 int) *string {

    retStr := C.CkSpider_getFailedUrl(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth URL in the outbound link URL list. Indexing begins at 0.
// return a string or nil if failed.
func (z *Spider) GetOutboundLink(arg1 int) *string {

    retStr := C.CkSpider_getOutboundLink(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth URL in the already-spidered URL list. Indexing begins at 0.
// return a string or nil if failed.
func (z *Spider) GetSpideredUrl(arg1 int) *string {

    retStr := C.CkSpider_getSpideredUrl(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth URL in the unspidered URL list. Indexing begins at 0.
// return a string or nil if failed.
func (z *Spider) GetUnspideredUrl(arg1 int) *string {

    retStr := C.CkSpider_getUnspideredUrl(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the domain name part of a URL. For example, if the URL is
// "http://www.chilkatsoft.com/test.asp", then "www.chilkatsoft.com" is returned.
// return a string or nil if failed.
func (z *Spider) GetUrlDomain(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkSpider_getUrlDomain(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Initializes the component to begin spidering a domain. Calling Initialize clears
// any patterns added via the AddAvoidOutboundLinkPattern, AddAvoidPattern, and
// AddMustMatchPattern methods. The domain name passed to this method is what is
// returned by the Domain property. The spider only crawls URLs within the same
// domain.
func (z *Spider) Initialize(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkSpider_Initialize(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Loads the caller of the task's async method.
func (z *Spider) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkSpider_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Re-crawls the last URL spidered. This helpful when cookies set in a previous
// page load cause the page to be loaded differently the next time.
func (z *Spider) RecrawlLast() bool {

    v := C.CkSpider_RecrawlLast(z.ckObj)


    return v != 0
}


// Asynchronous version of the RecrawlLast method
func (z *Spider) RecrawlLastAsync(c chan *Task) {

    hTask := C.CkSpider_RecrawlLastAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Moves a URL from the unspidered list to the spidered list. This allows an
// application to skip a specific URL.
func (z *Spider) SkipUnspidered(arg1 int)  {

    C.CkSpider_SkipUnspidered(z.ckObj, C.int(arg1))



}


// Suspends the execution of the current thread until the time-out interval
// elapses.
func (z *Spider) SleepMs(arg1 int)  {

    C.CkSpider_SleepMs(z.ckObj, C.int(arg1))



}


