// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkCsr.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCsr() *Csr {
	return &Csr{C.CkCsr_Create()}
}

func (z *Csr) DisposeCsr() {
    if z != nil {
	C.CkCsr_Dispose(z.ckObj)
	}
}




// property: CommonName
// The common name of the certificate to be generated. For SSL/TLS certificates,
// this would be the domain name. For email certificates this would be the email
// address.
// 
// It is the value for "CN" in the certificate's Subject's distinguished name (DN).
// (This is the value for OID "2.5.4.3")
// 
// This property is required for a CSR.
//
func (z *Csr) CommonName() string {
    return C.GoString(C.CkCsr_commonName(z.ckObj))
}
// property setter: CommonName
func (z *Csr) SetCommonName(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putCommonName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Company
// The company or organization name for the certificate to be generated.
// 
// It is the value for "O" in the certificate's Subject's distinguished name (DN).
// (This is the value for OID "2.5.4.10")
// 
// This property is optional. It may left empty.
//
func (z *Csr) Company() string {
    return C.GoString(C.CkCsr_company(z.ckObj))
}
// property setter: Company
func (z *Csr) SetCompany(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putCompany(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CompanyDivision
// The company division or organizational unit name for the certificate to be
// generated.
// 
// It is the value for "OU" in the certificate's Subject's distinguished name (DN).
// (This is the value for OID "2.5.4.11")
// 
// This property is optional. It may left empty.
//
func (z *Csr) CompanyDivision() string {
    return C.GoString(C.CkCsr_companyDivision(z.ckObj))
}
// property setter: CompanyDivision
func (z *Csr) SetCompanyDivision(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putCompanyDivision(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Country
// The two-letter uppercase country abbreviation, such as "US", for the certificate
// to be generated.
// 
// It is the value for "C" in the certificate's Subject's distinguished name (DN).
// (This is the value for OID "2.5.4.6")
// 
// This property is optional. It may left empty.
//
func (z *Csr) Country() string {
    return C.GoString(C.CkCsr_country(z.ckObj))
}
// property setter: Country
func (z *Csr) SetCountry(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putCountry(z.ckObj,cStr)
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
func (z *Csr) DebugLogFilePath() string {
    return C.GoString(C.CkCsr_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Csr) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EmailAddress
// The email address for the certificate to be generated.
// 
// It is the value for "E" in the certificate's Subject's distinguished name (DN).
// (This is the value for OID "1.2.840.113549.1.9.1")
// 
// This property is optional. It may left empty.
//
func (z *Csr) EmailAddress() string {
    return C.GoString(C.CkCsr_emailAddress(z.ckObj))
}
// property setter: EmailAddress
func (z *Csr) SetEmailAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putEmailAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HashAlgorithm
// The hash algorithm to be used when creating the CSR. The default is SHA256. Can
// be set to SHA1, SHA384, SHA256, or SHA512.
func (z *Csr) HashAlgorithm() string {
    return C.GoString(C.CkCsr_hashAlgorithm(z.ckObj))
}
// property setter: HashAlgorithm
func (z *Csr) SetHashAlgorithm(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putHashAlgorithm(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Csr) LastErrorHtml() string {
    return C.GoString(C.CkCsr_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Csr) LastErrorText() string {
    return C.GoString(C.CkCsr_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Csr) LastErrorXml() string {
    return C.GoString(C.CkCsr_lastErrorXml(z.ckObj))
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
func (z *Csr) LastMethodSuccess() bool {
    v := int(C.CkCsr_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Csr) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsr_putLastMethodSuccess(z.ckObj,v)
}

// property: Locality
// The locality (city or town) name for the certificate to be generated.
// 
// It is the value for "L" in the certificate's Subject's distinguished name (DN).
// (This is the value for OID "2.5.4.7")
// 
// This property is optional. It may left empty.
//
func (z *Csr) Locality() string {
    return C.GoString(C.CkCsr_locality(z.ckObj))
}
// property setter: Locality
func (z *Csr) SetLocality(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putLocality(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: MgfHashAlg
// If the private key is RSA and PssPadding equals true (RSASSA-PSS padding is
// used for the RSA signature), then this property controls the MGF hash algorithm
// used in the RSASSA-PSS padding. The default is "sha256". Can be set to "sha256",
// "sha384", or "sha512".
func (z *Csr) MgfHashAlg() string {
    return C.GoString(C.CkCsr_mgfHashAlg(z.ckObj))
}
// property setter: MgfHashAlg
func (z *Csr) SetMgfHashAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putMgfHashAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: PssPadding
// If _CKTRUE_, and if the private key is RSA, then uses RSASSA-PSS padding for the
// signature.
func (z *Csr) PssPadding() bool {
    v := int(C.CkCsr_getPssPadding(z.ckObj))
    return v != 0
}
// property setter: PssPadding
func (z *Csr) SetPssPadding(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsr_putPssPadding(z.ckObj,v)
}

// property: State
// The state or province for the certificate to be generated.
// 
// It is the value for "S" (or "ST") in the certificate's Subject's distinguished
// name (DN). (This is the value for OID "2.5.4.8")
// 
// This property is optional. It may left empty.
//
func (z *Csr) State() string {
    return C.GoString(C.CkCsr_state(z.ckObj))
}
// property setter: State
func (z *Csr) SetState(goStr string) {
    cStr := C.CString(goStr)
    C.CkCsr_putState(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Csr) VerboseLogging() bool {
    v := int(C.CkCsr_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Csr) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCsr_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Csr) Version() string {
    return C.GoString(C.CkCsr_version(z.ckObj))
}
// Generate a CSR and return the binary DER in csrData. The privKey can be an RSA or
// ECDSA private key.
func (z *Csr) GenCsrBd(arg1 *PrivateKey, arg2 *BinData) bool {

    v := C.CkCsr_GenCsrBd(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Generate a CSR and return it as a PEM string. The privKey can be an RSA or ECDSA
// private key.
// return a string or nil if failed.
func (z *Csr) GenCsrPem(arg1 *PrivateKey) *string {

    retStr := C.CkCsr_genCsrPem(z.ckObj, arg1.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the CSR's public key in the pubkey.
func (z *Csr) GetPublicKey(arg1 *PublicKey) bool {

    v := C.CkCsr_GetPublicKey(z.ckObj, arg1.ckObj)


    return v != 0
}


// Gets a subject field by OID, such as "2.5.4.9".
// Seehttp://www.alvestrand.no/objectid/2.5.4.html
// <http://www.alvestrand.no/objectid/2.5.4.html> for OID values and meanings.
// return a string or nil if failed.
func (z *Csr) GetSubjectField(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCsr_getSubjectField(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads this CSR object with a CSR PEM. All properties are set to the values found
// within the CSR.
func (z *Csr) LoadCsrPem(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCsr_LoadCsrPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets a subject field by OID, such as "2.5.4.9".
// Seehttp://www.alvestrand.no/objectid/2.5.4.html
// <http://www.alvestrand.no/objectid/2.5.4.html> for OID values and meanings.
// 
// The asnType can be "UTF8String", "IA5String", or "PrintableString". If you have no
// specific requirement, or don't know, choose "UTF8String".
//
func (z *Csr) SetSubjectField(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkCsr_SetSubjectField(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Verify the signature in the CSR. Return true if the signature is valid.
func (z *Csr) VerifyCsr() bool {

    v := C.CkCsr_VerifyCsr(z.ckObj)


    return v != 0
}


