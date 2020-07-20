// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkGlobal.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewGlobal() *Global {
	return &Global{C.CkGlobal_Create()}
}

func (z *Global) DisposeGlobal() {
    if z != nil {
	C.CkGlobal_Dispose(z.ckObj)
	}
}




// property: AnsiCodePage
// The default ANSI code page is determined at runtime based on the computer where
// the application happens to be running. For example, the ANSI code page for an
// application running on a Japanese computer is likely to be Shift_JIS (code page
// 932), whereas on a US-English computer it would be iso-8859-1 (or Windows-1252
// which is essentially a superset of iso-8859-1).
// 
// If there is a desire for the Chilkat library to use a specific ANSI code page
// regardless of locale, then this property should be set to the desired code page.
// The default value of this property is the ANSI code page of the local computer.
//
func (z *Global) AnsiCodePage() int {
    return int(C.CkGlobal_getAnsiCodePage(z.ckObj))
}
// property setter: AnsiCodePage
func (z *Global) SetAnsiCodePage(value int) {
    C.CkGlobal_putAnsiCodePage(z.ckObj,C.int(value))
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
func (z *Global) DebugLogFilePath() string {
    return C.GoString(C.CkGlobal_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Global) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkGlobal_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DefaultNtlmVersion
// Selects the default NTLM protocol version to use for NTLM authentication for
// HTTP, POP3, IMAP, SMTP, and HTTP proxies. The default value is 2. This property
// may optionally be set to 1.
func (z *Global) DefaultNtlmVersion() int {
    return int(C.CkGlobal_getDefaultNtlmVersion(z.ckObj))
}
// property setter: DefaultNtlmVersion
func (z *Global) SetDefaultNtlmVersion(value int) {
    C.CkGlobal_putDefaultNtlmVersion(z.ckObj,C.int(value))
}

// property: DefaultUtf8
// Applies only to programming languages where each class has the Utf8 property,
// and where strings are passed and returned as multibyte (null-terminated
// sequences of bytes). This includes the multibyte C/C++ API, Perl, Python 2.*
// (not Python 3.*), Ruby, and PHP. This does not include Java, Objective-C, or
// Python 3.* as all strings in these languages are utf-8. This property has no
// effect in programming languages or APIs that return Unicode or strings as
// objects (such as .NET).
// 
// A Chilkat class's Utf8 property controls whether strings are returned as utf-8
// or ANSI. It also controls how Chilkat is to interpret the bytes of passed-in
// arguments. It must be set to false if the application is passing ANSI strings
// (i.e. the byte representation is ANSI), and must be set to true if the
// application is passing string arguments using the utf-8 representation.
// 
// This global Utf8 property controls the default setting of the Utf8 property for
// all Chilkat objects. Thus it allows for an application to be entirely in "utf-8
// mode" or "ANSI mode" without needing to explicity set the Utf8 property of every
// Chilkat object instance.
//
func (z *Global) DefaultUtf8() bool {
    v := int(C.CkGlobal_getDefaultUtf8(z.ckObj))
    return v != 0
}
// property setter: DefaultUtf8
func (z *Global) SetDefaultUtf8(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGlobal_putDefaultUtf8(z.ckObj,v)
}

// property: DnsTimeToLive
// If DNS caching is enabled, this is the time-to-live (in seconds) for a cached
// DNS lookup. A DNS lookup result older than this expiration time is discarded,
// and causes a new DNS lookup to occur. A value of 0 indicates an infinite
// time-to-live. The default value of this property is 0.
func (z *Global) DnsTimeToLive() int {
    return int(C.CkGlobal_getDnsTimeToLive(z.ckObj))
}
// property setter: DnsTimeToLive
func (z *Global) SetDnsTimeToLive(value int) {
    C.CkGlobal_putDnsTimeToLive(z.ckObj,C.int(value))
}

// property: EnableDnsCaching
// Controls whether DNS domain lookups (to resolve to IP addresses) are cached in
// memory. The default value is false, meaning that DNS caching is disabled.
func (z *Global) EnableDnsCaching() bool {
    v := int(C.CkGlobal_getEnableDnsCaching(z.ckObj))
    return v != 0
}
// property setter: EnableDnsCaching
func (z *Global) SetEnableDnsCaching(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGlobal_putEnableDnsCaching(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Global) LastErrorHtml() string {
    return C.GoString(C.CkGlobal_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Global) LastErrorText() string {
    return C.GoString(C.CkGlobal_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Global) LastErrorXml() string {
    return C.GoString(C.CkGlobal_lastErrorXml(z.ckObj))
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
func (z *Global) LastMethodSuccess() bool {
    v := int(C.CkGlobal_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Global) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGlobal_putLastMethodSuccess(z.ckObj,v)
}

// property: MaxThreads
// The maximum number of thread pool threads. The default value is 100. The maximum
// value is 500. Note: Asynchronous worker threads are created on as needed up to
// the maximum.
func (z *Global) MaxThreads() int {
    return int(C.CkGlobal_getMaxThreads(z.ckObj))
}
// property setter: MaxThreads
func (z *Global) SetMaxThreads(value int) {
    C.CkGlobal_putMaxThreads(z.ckObj,C.int(value))
}

// property: PreferIpv6
// If true, then use IPv6 over IPv4 when both are supported for a particular
// domain. The default value of this property is false, which will choose IPv4
// over IPv6.
// 
// Note: Setting this property has the effect of also setting the default value of
// the PreferIpv6 property for other classes.
//
func (z *Global) PreferIpv6() bool {
    v := int(C.CkGlobal_getPreferIpv6(z.ckObj))
    return v != 0
}
// property setter: PreferIpv6
func (z *Global) SetPreferIpv6(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGlobal_putPreferIpv6(z.ckObj,v)
}

// property: ThreadPoolLogPath
// If set, indicates the path of a log file to be used by the thread pool thread
// and each of the pool worker threads for logging async activity.
func (z *Global) ThreadPoolLogPath() string {
    return C.GoString(C.CkGlobal_threadPoolLogPath(z.ckObj))
}
// property setter: ThreadPoolLogPath
func (z *Global) SetThreadPoolLogPath(goStr string) {
    cStr := C.CString(goStr)
    C.CkGlobal_putThreadPoolLogPath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UnlockStatus
// Indicates the unlocked status for the last call to UnlockBundle, or any
// UnlockComponent call. The possible values are:
//     Not unlocked. (Still in locked state.)
//     Unlocked with in fully-functional trial mode.
//     Unlocked using a valid purchased unlock code.
// 
// Note: If UnlockComponent or UnlockBundle is called with a purchased unlock code,
// the UnlockStatus is correctly set to the value 2. This value is intentionally
// sticky. If a subsequent and redundant call to UnlockComponent (or UnlockBundle)
// happens, it is effectively a "No-Op" because the library is already unlocked.
// The UnlockStatus will not change.
// 
// If however, if the 1st call resulted in UnlockStatus = 1, and THEN the unlock
// method is called again with a purchased unlock code, the UnlockStatus should
// change from 1 to 2.
//
func (z *Global) UnlockStatus() int {
    return int(C.CkGlobal_getUnlockStatus(z.ckObj))
}

// property: UsePkcsConstructedEncoding
// This property should typically be left at the default value of false. If set
// to true, then Chilkat will use a constructed ASN.1 encoding for PCKS7 data.
// (This is an internal implementation option that normally does not matter, and
// should not matter. Some PKCS7 receiving systems might be picky, and this option
// can be used to satisfy this requirement.)
func (z *Global) UsePkcsConstructedEncoding() bool {
    v := int(C.CkGlobal_getUsePkcsConstructedEncoding(z.ckObj))
    return v != 0
}
// property setter: UsePkcsConstructedEncoding
func (z *Global) SetUsePkcsConstructedEncoding(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGlobal_putUsePkcsConstructedEncoding(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Global) VerboseLogging() bool {
    v := int(C.CkGlobal_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Global) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGlobal_putVerboseLogging(z.ckObj,v)
}

// property: VerboseTls
// If set to true, then causes extremely verbose logging (in LastErrorText) all
// TLS connections in any Chilkat class. This property should only be used for
// troubleshooting TLS problems. The default value is false.
// 
// Note: This property only has effect on Chilkat objects not yet created. Set the
// property first, then instantiate the Chilkat object.
//
func (z *Global) VerboseTls() bool {
    v := int(C.CkGlobal_getVerboseTls(z.ckObj))
    return v != 0
}
// property setter: VerboseTls
func (z *Global) SetVerboseTls(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkGlobal_putVerboseTls(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Global) Version() string {
    return C.GoString(C.CkGlobal_version(z.ckObj))
}
// Clears the global DNS cache.
func (z *Global) DnsClearCache() bool {

    v := C.CkGlobal_DnsClearCache(z.ckObj)


    return v != 0
}


// Called to stop and finalize all threads in the thread pool, and causes the
// thread pool thread to exit.
// 
// The following behaviors exist in v9.5.0.64 and later:
//     All remaining asynchronous tasks are automatically canceled.
//     Restores the thread pool to it's pristine state where no background threads
//     are running.
// 
// It is a good idea to call this method at the very end of a program, just before
// it exits. This is especially true for programs written in VBScript and VB6.
//
func (z *Global) FinalizeThreadPool() bool {

    v := C.CkGlobal_FinalizeThreadPool(z.ckObj)


    return v != 0
}


// Logs a line to the thread pool log file.
func (z *Global) ThreadPoolLogLine(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkGlobal_ThreadPoolLogLine(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unlocks the entire Chilkat API for all classes. This should be called once at
// the beginning of a program. Once unlocked, objects of any Chilkat class may be
// instantiated and used. To unlock in fully-functional 30-day trial mode, pass any
// string, such as "Hello", in bundleUnlockCode. If a license is purchased, then replace the
// "Hello" with the purchased unlock code.
// 
// After calling UnlockBundle once, the instance of the CLASS_NAME object may be
// discarded/deleted (assuming the programming language requires explicit deletes).
// Multiple calls to UnlockComponent are harmless. If the Chilkat API is already
// unlocked, the duplicate calls to UnlockBundle are no-ops.
//
func (z *Global) UnlockBundle(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkGlobal_UnlockBundle(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


