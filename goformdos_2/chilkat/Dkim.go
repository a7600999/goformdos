// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkDkim.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewDkim() *Dkim {
	return &Dkim{C.CkDkim_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Dkim) DisposeDkim() {
    if z != nil {
	C.CkDkim_Dispose(z.ckObj)
	}
}




func (z *Dkim) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkDkim_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Dkim) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkDkim_setExternalProgress(z.ckObj,1)
}

func (z *Dkim) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkDkim_setExternalProgress(z.ckObj,1)
}

func (z *Dkim) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkDkim_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Dkim) AbortCurrent() bool {
    v := int(C.CkDkim_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Dkim) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDkim_putAbortCurrent(z.ckObj,v)
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
func (z *Dkim) DebugLogFilePath() string {
    return C.GoString(C.CkDkim_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Dkim) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DkimAlg
// The signing algorithm to be used in creating the DKIM-Signature. Possible values
// are "rsa-sha256" and "rsa-sha1". The default value is "rsa-sha256".
func (z *Dkim) DkimAlg() string {
    return C.GoString(C.CkDkim_dkimAlg(z.ckObj))
}
// property setter: DkimAlg
func (z *Dkim) SetDkimAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDkimAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DkimBodyLengthCount
// Optional body length count to set a maximum number of body bytes to be hashed
// when creating the DKIM-Signature field. The default value is 0, indicating that
// the entire body should be hashed.
func (z *Dkim) DkimBodyLengthCount() int {
    return int(C.CkDkim_getDkimBodyLengthCount(z.ckObj))
}
// property setter: DkimBodyLengthCount
func (z *Dkim) SetDkimBodyLengthCount(value int) {
    C.CkDkim_putDkimBodyLengthCount(z.ckObj,C.int(value))
}

// property: DkimCanon
// Canonicalization algorithm to be used for both header and body when creating the
// DKIM-Signature. Possible values are "simple" and "relaxed". The default value is
// "relaxed".
func (z *Dkim) DkimCanon() string {
    return C.GoString(C.CkDkim_dkimCanon(z.ckObj))
}
// property setter: DkimCanon
func (z *Dkim) SetDkimCanon(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDkimCanon(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DkimDomain
// The domain name of the signing domain when creating the DKIM-Signature.
func (z *Dkim) DkimDomain() string {
    return C.GoString(C.CkDkim_dkimDomain(z.ckObj))
}
// property setter: DkimDomain
func (z *Dkim) SetDkimDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDkimDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DkimHeaders
// A colon-separated list of header field names that identify headers presented to
// the signing algorithm when creating the DKIM-Signature. The default value is:
// "mime-version:date:message-id:subject:from:to:content-type".
func (z *Dkim) DkimHeaders() string {
    return C.GoString(C.CkDkim_dkimHeaders(z.ckObj))
}
// property setter: DkimHeaders
func (z *Dkim) SetDkimHeaders(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDkimHeaders(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DkimSelector
// The selector to be used to form the DNS query for the public key. This property
// applies to creating a DKIM-Signature. For example, if the selector is
// "reykjavik" and the domain is "example-code.com", then the DNS query would be
// for "reykjavik._domainkey.example-code.com".
func (z *Dkim) DkimSelector() string {
    return C.GoString(C.CkDkim_dkimSelector(z.ckObj))
}
// property setter: DkimSelector
func (z *Dkim) SetDkimSelector(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDkimSelector(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DomainKeyAlg
// The signing algorithm to be used in creating the DomainKey-Signature. The only
// possible value is "rsa-sha1". The default value is "rsa-sha1".
func (z *Dkim) DomainKeyAlg() string {
    return C.GoString(C.CkDkim_domainKeyAlg(z.ckObj))
}
// property setter: DomainKeyAlg
func (z *Dkim) SetDomainKeyAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDomainKeyAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DomainKeyCanon
// Canonicalization algorithm to be used for when creating the DomainKey-Signature.
// Possible values are "simple" and "nofws". The default value is "nofws".
func (z *Dkim) DomainKeyCanon() string {
    return C.GoString(C.CkDkim_domainKeyCanon(z.ckObj))
}
// property setter: DomainKeyCanon
func (z *Dkim) SetDomainKeyCanon(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDomainKeyCanon(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DomainKeyDomain
// The domain name of the signing domain when creating the DomainKey-Signature.
func (z *Dkim) DomainKeyDomain() string {
    return C.GoString(C.CkDkim_domainKeyDomain(z.ckObj))
}
// property setter: DomainKeyDomain
func (z *Dkim) SetDomainKeyDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDomainKeyDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DomainKeyHeaders
// A colon-separated list of header field names that identify headers presented to
// the signing algorithm when creating the DomainKey-Signature. The default value
// is: "mime-version:date:message-id:subject:from:to:content-type".
func (z *Dkim) DomainKeyHeaders() string {
    return C.GoString(C.CkDkim_domainKeyHeaders(z.ckObj))
}
// property setter: DomainKeyHeaders
func (z *Dkim) SetDomainKeyHeaders(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDomainKeyHeaders(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DomainKeySelector
// The selector to be used to form the DNS query for the public key. This property
// applies to creating a DomainKey-Signature. For example, if the selector is
// "reykjavik" and the domain is "example-code.com", then the DNS query would be
// for "reykjavik._domainkey.example-code.com".
func (z *Dkim) DomainKeySelector() string {
    return C.GoString(C.CkDkim_domainKeySelector(z.ckObj))
}
// property setter: DomainKeySelector
func (z *Dkim) SetDomainKeySelector(goStr string) {
    cStr := C.CString(goStr)
    C.CkDkim_putDomainKeySelector(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// The number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any method call prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
func (z *Dkim) HeartbeatMs() int {
    return int(C.CkDkim_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Dkim) SetHeartbeatMs(value int) {
    C.CkDkim_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Dkim) LastErrorHtml() string {
    return C.GoString(C.CkDkim_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Dkim) LastErrorText() string {
    return C.GoString(C.CkDkim_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Dkim) LastErrorXml() string {
    return C.GoString(C.CkDkim_lastErrorXml(z.ckObj))
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
func (z *Dkim) LastMethodSuccess() bool {
    v := int(C.CkDkim_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Dkim) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDkim_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Dkim) VerboseLogging() bool {
    v := int(C.CkDkim_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Dkim) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDkim_putVerboseLogging(z.ckObj,v)
}

// property: VerifyInfo
// Contains JSON with information about the last DKIM or DomainKey signature
// verified (or verification failure). The JSON will contain information like this:
// {
//   "domain": "amazonses.com",
//   "selector": "7v7vs6w47njt4pimodk5mmttbegzsi6n",
//   "publicKey": "MIGfMA0GCSq...z6uqeQIDAQAB",
//   "canonicalization": "relaxed/simple",
//   "algorithm": "rsa-sha256",
//   "signedHeaders": "Subject:From:To:Date:Mime-Version:Content-Type:References:Message-Id:Feedback-ID",
//   "verified": "yes"
// }
func (z *Dkim) VerifyInfo() string {
    return C.GoString(C.CkDkim_verifyInfo(z.ckObj))
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Dkim) Version() string {
    return C.GoString(C.CkDkim_version(z.ckObj))
}
// Constructs and prepends a DKIM-Signature header to the MIME passed in mimeData.
// Prior to calling this method, your program must set both the DkimDomain and
// DkimSelector properties, and it must load a private key by calling
// SetDkimPrivateKey.
func (z *Dkim) DkimSign(arg1 *BinData) bool {

    v := C.CkDkim_DkimSign(z.ckObj, arg1.ckObj)


    return v != 0
}


// Verifies the Nth DKIM-Signature header in the mimeData. (In most cases, there is
// only one signature.) The 1st signature is at sigIndex 0.
// 
// On return, the VerifyInfo property will contain additional information about the
// DKIM-Signature that was verified (or not verified).
//
func (z *Dkim) DkimVerify(arg1 int, arg2 *BinData) bool {

    v := C.CkDkim_DkimVerify(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Constructs and prepends a DomainKey-Signature header to the MIME passed in mimeData.
// Prior to calling this method, your program must set both the DomainKeyDomain and
// DomainKeySelector properties, and it must load a private key by calling
// SetDomainKeyPrivateKey.
func (z *Dkim) DomainKeySign(arg1 *BinData) bool {

    v := C.CkDkim_DomainKeySign(z.ckObj, arg1.ckObj)


    return v != 0
}


// Verifies the Nth DomainKey-Signature header in the mimeData. (In most cases, there
// is only one signature.) The 1st signature is at sigIndex 0.
// 
// On return, the VerifyInfo property will contain additional information about the
// DKIM-Signature that was verified (or not verified).
//
func (z *Dkim) DomainKeyVerify(arg1 int, arg2 *BinData) bool {

    v := C.CkDkim_DomainKeyVerify(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Caches a public key to be used for verifying DKIM and DomainKey signatures for a
// given selector and domain. The publicKey is a string containing an RSA public key in
// any text format, such as XML, PEM, etc. This method will automatically detect
// the format and load the public key correctly. This method is useful for testing
// DKIM and DomainKey verification when your public key has not yet been installed
// in DNS.
func (z *Dkim) LoadPublicKey(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkDkim_LoadPublicKey(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Caches a public key to be used for verifying DKIM and DomainKey signatures for a
// given selector and domain. The publicKeyFilepath is a filepath of an RSA public key in any
// format. This method will automatically detect the format and load the public key
// correctly. This method is useful for testing DKIM and DomainKey verification
// when your public key has not yet been installed in DNS.
func (z *Dkim) LoadPublicKeyFile(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkDkim_LoadPublicKeyFile(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Loads the caller of the task's async method.
func (z *Dkim) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkDkim_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the number of DKIM-Signature header fields found in mimeData.
func (z *Dkim) NumDkimSigs(arg1 *BinData) int {

    v := C.CkDkim_NumDkimSigs(z.ckObj, arg1.ckObj)


    return int(v)
}


// Returns the number of DomainKey-Signature header fields found in mimeData.
func (z *Dkim) NumDomainKeySigs(arg1 *BinData) int {

    v := C.CkDkim_NumDomainKeySigs(z.ckObj, arg1.ckObj)


    return int(v)
}


// Useful if your application is going to verify many emails from a single domain
// (or a few domains). This eliminates the need to do a DNS lookup to fetch the
// public key each time an email's DKIM or DomainKey signature is verified.
// 
// This method may be called multiple times -- once for each selector/domain to be
// pre-fetched. The verify methods (VerifyDkimSignature and
// VerifyDomainKeySignature) will use a pre-fetched key if the selector and domain
// match.
//
func (z *Dkim) PrefetchPublicKey(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkDkim_PrefetchPublicKey(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the PrefetchPublicKey method
func (z *Dkim) PrefetchPublicKeyAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkDkim_PrefetchPublicKeyAsync(z.ckObj, cstr1, cstr2)

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


// Sets the private key file to be used for creating a DKIM-Signature.
func (z *Dkim) SetDkimPrivateKey(arg1 *PrivateKey) bool {

    v := C.CkDkim_SetDkimPrivateKey(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the private key file to be used for creating a DomainKey-Signature.
func (z *Dkim) SetDomainKeyPrivateKey(arg1 *PrivateKey) bool {

    v := C.CkDkim_SetDomainKeyPrivateKey(z.ckObj, arg1.ckObj)


    return v != 0
}


// Unlocks the component allowing for the full functionality to be used. If this
// method unexpectedly returns false, examine the contents of the LastErrorText
// property to determine the reason for failure.
func (z *Dkim) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDkim_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


