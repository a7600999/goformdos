// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkCertChain.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCertChain() *CertChain {
	return &CertChain{C.CkCertChain_Create()}
}

func (z *CertChain) DisposeCertChain() {
    if z != nil {
	C.CkCertChain_Dispose(z.ckObj)
	}
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
func (z *CertChain) DebugLogFilePath() string {
    return C.GoString(C.CkCertChain_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *CertChain) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkCertChain_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *CertChain) LastErrorHtml() string {
    return C.GoString(C.CkCertChain_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *CertChain) LastErrorText() string {
    return C.GoString(C.CkCertChain_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *CertChain) LastErrorXml() string {
    return C.GoString(C.CkCertChain_lastErrorXml(z.ckObj))
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
func (z *CertChain) LastMethodSuccess() bool {
    v := int(C.CkCertChain_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *CertChain) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCertChain_putLastMethodSuccess(z.ckObj,v)
}

// property: NumCerts
// The number of certificates in the chain. The end-user subscriber certificate is
// at index 0. This is the certificate most removed from the trusted root.
// Intermediate certificates are at indices 1 to NumCerts - 2. The trusted root (or
// self-signed certificate) is at index NumCerts - 1.
func (z *CertChain) NumCerts() int {
    return int(C.CkCertChain_getNumCerts(z.ckObj))
}

// property: NumExpiredCerts
// Returns the number of certificates in the chain that have expired.
func (z *CertChain) NumExpiredCerts() int {
    return int(C.CkCertChain_getNumExpiredCerts(z.ckObj))
}

// property: ReachesRoot
// true if this certificate chain extends all the way to the root certificate.
// The root certificate is either a trusted CA root or a self-signed certificate.
// In both cases, the issuer of a root certificate is itself.
func (z *CertChain) ReachesRoot() bool {
    v := int(C.CkCertChain_getReachesRoot(z.ckObj))
    return v != 0
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *CertChain) VerboseLogging() bool {
    v := int(C.CkCertChain_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *CertChain) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCertChain_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *CertChain) Version() string {
    return C.GoString(C.CkCertChain_version(z.ckObj))
}
// Returns the Nth certificate in the chain.
func (z *CertChain) GetCert(arg1 int) *Cert {

    retObj := C.CkCertChain_GetCert(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns true if the root of the certificate chain is a certificate found in
// trustedRoots.
func (z *CertChain) IsRootTrusted(arg1 *TrustedRoots) bool {

    v := C.CkCertChain_IsRootTrusted(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads a certificate chain from the "x5c" parameter of a JWK (JSON Web Key).
func (z *CertChain) LoadX5C(arg1 *JsonObject) bool {

    v := C.CkCertChain_LoadX5C(z.ckObj, arg1.ckObj)


    return v != 0
}


// Verifies the certificate signatures to the root. Returns true if all
// certificate signatures are valid.
func (z *CertChain) VerifyCertSignatures() bool {

    v := C.CkCertChain_VerifyCertSignatures(z.ckObj)


    return v != 0
}


