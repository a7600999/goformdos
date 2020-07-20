// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkTrustedRoots.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewTrustedRoots() *TrustedRoots {
	return &TrustedRoots{C.CkTrustedRoots_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *TrustedRoots) DisposeTrustedRoots() {
    if z != nil {
	C.CkTrustedRoots_Dispose(z.ckObj)
	}
}




func (z *TrustedRoots) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkTrustedRoots_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *TrustedRoots) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkTrustedRoots_setExternalProgress(z.ckObj,1)
}

func (z *TrustedRoots) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkTrustedRoots_setExternalProgress(z.ckObj,1)
}

func (z *TrustedRoots) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkTrustedRoots_setExternalProgress(z.ckObj,1)
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
func (z *TrustedRoots) DebugLogFilePath() string {
    return C.GoString(C.CkTrustedRoots_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *TrustedRoots) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkTrustedRoots_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *TrustedRoots) LastErrorHtml() string {
    return C.GoString(C.CkTrustedRoots_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *TrustedRoots) LastErrorText() string {
    return C.GoString(C.CkTrustedRoots_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *TrustedRoots) LastErrorXml() string {
    return C.GoString(C.CkTrustedRoots_lastErrorXml(z.ckObj))
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
func (z *TrustedRoots) LastMethodSuccess() bool {
    v := int(C.CkTrustedRoots_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *TrustedRoots) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTrustedRoots_putLastMethodSuccess(z.ckObj,v)
}

// property: NumCerts
// The number of certificates contained within this object.
// 
// This is the number of certificates explicitly added by the methods AddCert,
// AddJavaKeyStore, and LoadCaCertsPem.
//
func (z *TrustedRoots) NumCerts() int {
    return int(C.CkTrustedRoots_getNumCerts(z.ckObj))
}

// property: RejectSelfSignedCerts
// Indicates whether all self-signed certificates are to be rejected in SSL/TLS
// connections. The default value of this property is false.
// 
// Note: This is for the case where the server certificate chain of authentication
// is 1 certificate long (i.e. the TLS server certificate itself is self-signed).
//
func (z *TrustedRoots) RejectSelfSignedCerts() bool {
    v := int(C.CkTrustedRoots_getRejectSelfSignedCerts(z.ckObj))
    return v != 0
}
// property setter: RejectSelfSignedCerts
func (z *TrustedRoots) SetRejectSelfSignedCerts(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTrustedRoots_putRejectSelfSignedCerts(z.ckObj,v)
}

// property: TrustSystemCaRoots
// Indicates whether the operating system's CA root certificates are automatically
// trusted.
// 
// On a Windows operating system, this would be the registry-based CA certificate
// stores. On a Linux system, this could be /etc/ssl/certs/ca-certificates.crt, if
// it exists. The default value is true. Set this property equal to false to
// prevent Chilkat from automatically trusting system-provided root CA
// certificates.
//
func (z *TrustedRoots) TrustSystemCaRoots() bool {
    v := int(C.CkTrustedRoots_getTrustSystemCaRoots(z.ckObj))
    return v != 0
}
// property setter: TrustSystemCaRoots
func (z *TrustedRoots) SetTrustSystemCaRoots(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTrustedRoots_putTrustSystemCaRoots(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *TrustedRoots) VerboseLogging() bool {
    v := int(C.CkTrustedRoots_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *TrustedRoots) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkTrustedRoots_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *TrustedRoots) Version() string {
    return C.GoString(C.CkTrustedRoots_version(z.ckObj))
}
// Activates this collection of trusted roots as the set of CA and self-signed root
// certificates that are to be trusted Chilkat-wide for PKCS7 signature validation
// and SSL/TLS server certificate validation.
func (z *TrustedRoots) Activate() bool {

    v := C.CkTrustedRoots_Activate(z.ckObj)


    return v != 0
}


// Adds a certificate to the collection of trusted roots.
func (z *TrustedRoots) AddCert(arg1 *Cert) bool {

    v := C.CkTrustedRoots_AddCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds the trusted certificates from a Java key store to the collection of trusted
// roots.
func (z *TrustedRoots) AddJavaKeyStore(arg1 *JavaKeyStore) bool {

    v := C.CkTrustedRoots_AddJavaKeyStore(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the AddJavaKeyStore method
func (z *TrustedRoots) AddJavaKeyStoreAsync(arg1 *JavaKeyStore, c chan *Task) {

    hTask := C.CkTrustedRoots_AddJavaKeyStoreAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Deactivates a previously set of activated trusted roots so that all roots /
// self-signed certificates are implicitly trusted.
func (z *TrustedRoots) Deactivate() bool {

    v := C.CkTrustedRoots_Deactivate(z.ckObj)


    return v != 0
}


// Returns the Nth cert contained within this object. The 1st certificate is at
// index 0.
func (z *TrustedRoots) GetCert(arg1 int) *Cert {

    retObj := C.CkTrustedRoots_GetCert(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Loads a CA bundle in PEM format. This is a file containing CA root certificates
// that are to be trusted. An example of one such file is the CA certs from
// mozilla.org exported to a cacert.pem file by the mk-ca-bundle tool located here:
// http://curl.haxx.se/docs/caextract.html.
// 
// Note: This can also be called to load the /etc/ssl/certs/ca-certificates.crt
// file on Linux systems.
//
func (z *TrustedRoots) LoadCaCertsPem(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkTrustedRoots_LoadCaCertsPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the LoadCaCertsPem method
func (z *TrustedRoots) LoadCaCertsPemAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkTrustedRoots_LoadCaCertsPemAsync(z.ckObj, cstr1)

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


// Loads the caller of the task's async method.
func (z *TrustedRoots) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkTrustedRoots_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


