// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkCert.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCert() *Cert {
	return &Cert{C.CkCert_Create()}
}

func (z *Cert) DisposeCert() {
    if z != nil {
	C.CkCert_Dispose(z.ckObj)
	}
}




// property: AuthorityKeyId
// The authority key identifier of the certificate in base64 string format. This is
// only present if the certificate contains the extension OID 2.5.29.35.
func (z *Cert) AuthorityKeyId() string {
    return C.GoString(C.CkCert_authorityKeyId(z.ckObj))
}

// property: CertVersion
// The version of the certificate (1, 2, or 3). A value of 0 indicates an error --
// the most likely cause being that the certificate object is empty (i.e. was never
// loaded with a certificate). Note: This is not the version of the software, it is
// the version of the X.509 certificate object. The version of the Chilkat
// certificate software is indicated by the Version property.
func (z *Cert) CertVersion() int {
    return int(C.CkCert_getCertVersion(z.ckObj))
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
func (z *Cert) DebugLogFilePath() string {
    return C.GoString(C.CkCert_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Cert) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkCert_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Expired
// Has a value of true if the certificate or any certificate in the chain of
// authority has expired. (This information is not available when running on
// Windows 95/98 computers.)
func (z *Cert) Expired() bool {
    v := int(C.CkCert_getExpired(z.ckObj))
    return v != 0
}

// property: ForClientAuthentication
// true if this certificate can be used for client authentication, otherwise
// false.
func (z *Cert) ForClientAuthentication() bool {
    v := int(C.CkCert_getForClientAuthentication(z.ckObj))
    return v != 0
}

// property: ForCodeSigning
// true if this certificate can be used for code signing, otherwise false.
func (z *Cert) ForCodeSigning() bool {
    v := int(C.CkCert_getForCodeSigning(z.ckObj))
    return v != 0
}

// property: ForSecureEmail
// true if this certificate can be used for sending secure email, otherwise
// false.
func (z *Cert) ForSecureEmail() bool {
    v := int(C.CkCert_getForSecureEmail(z.ckObj))
    return v != 0
}

// property: ForServerAuthentication
// true if this certificate can be used for server authentication, otherwise
// false.
func (z *Cert) ForServerAuthentication() bool {
    v := int(C.CkCert_getForServerAuthentication(z.ckObj))
    return v != 0
}

// property: ForTimeStamping
// true if this certificate can be used for time stamping, otherwise false.
func (z *Cert) ForTimeStamping() bool {
    v := int(C.CkCert_getForTimeStamping(z.ckObj))
    return v != 0
}

// property: IntendedKeyUsage
// Bitflags indicating the intended usages of the certificate. The flags are:
// Digital Signature: 0x80
// Non-Repudiation: 0x40
// Key Encipherment: 0x20
// Data Encipherment: 0x10
// Key Agreement: 0x08
// Certificate Signing: 0x04
// CRL Signing: 0x02
// Encipher-Only: 0x01
func (z *Cert) IntendedKeyUsage() uint {
    return uint(C.CkCert_getIntendedKeyUsage(z.ckObj))
}

// property: IsRoot
// true if this is the root certificate, otherwise false.
func (z *Cert) IsRoot() bool {
    v := int(C.CkCert_getIsRoot(z.ckObj))
    return v != 0
}

// property: IssuerC
// The certificate issuer's country.
func (z *Cert) IssuerC() string {
    return C.GoString(C.CkCert_issuerC(z.ckObj))
}

// property: IssuerCN
// The certificate issuer's common name.
func (z *Cert) IssuerCN() string {
    return C.GoString(C.CkCert_issuerCN(z.ckObj))
}

// property: IssuerDN
// The issuer's full distinguished name.
func (z *Cert) IssuerDN() string {
    return C.GoString(C.CkCert_issuerDN(z.ckObj))
}

// property: IssuerE
// The certificate issuer's email address.
func (z *Cert) IssuerE() string {
    return C.GoString(C.CkCert_issuerE(z.ckObj))
}

// property: IssuerL
// The certificate issuer's locality, which could be a city, count, township, or
// other geographic region.
func (z *Cert) IssuerL() string {
    return C.GoString(C.CkCert_issuerL(z.ckObj))
}

// property: IssuerO
// The certificate issuer's organization, which is typically the company name.
func (z *Cert) IssuerO() string {
    return C.GoString(C.CkCert_issuerO(z.ckObj))
}

// property: IssuerOU
// The certificate issuer's organizational unit, which is the unit within the
// organization.
func (z *Cert) IssuerOU() string {
    return C.GoString(C.CkCert_issuerOU(z.ckObj))
}

// property: IssuerS
// The certificate issuer's state or province.
func (z *Cert) IssuerS() string {
    return C.GoString(C.CkCert_issuerS(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Cert) LastErrorHtml() string {
    return C.GoString(C.CkCert_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Cert) LastErrorText() string {
    return C.GoString(C.CkCert_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Cert) LastErrorXml() string {
    return C.GoString(C.CkCert_lastErrorXml(z.ckObj))
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
func (z *Cert) LastMethodSuccess() bool {
    v := int(C.CkCert_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Cert) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCert_putLastMethodSuccess(z.ckObj,v)
}

// property: OcspUrl
// If present in the certificate's extensions, returns the OCSP URL of the
// certificate. (The Online Certificate Status Protocol (OCSP) is an Internet
// protocol used for obtaining the revocation status of an X.509 digital
// certificate.)
func (z *Cert) OcspUrl() string {
    return C.GoString(C.CkCert_ocspUrl(z.ckObj))
}

// property: Revoked
// true if the certificate or any certificate in the chain of authority has been
// revoked. This information is not available when running on Windows 95/98
// computers. Note: If this property is false, it could mean that it was not able
// to check the revocation status. Because of this uncertainty, a CheckRevoked
// method has been added. It returns an integer indicating one of three possible
// states: 1 (revoked) , 0 (not revoked), -1 (unable to check revocation status).
func (z *Cert) Revoked() bool {
    v := int(C.CkCert_getRevoked(z.ckObj))
    return v != 0
}

// property: Rfc822Name
// The RFC-822 name of the certificate. (Also known as the Subject Alternative
// Name.)
// 
// If the certificate contains a list of Subject Alternative Names, such as a list
// of host names to be protected by a single SSL certificate, then this property
// will contain the comma separated list of names.
//
func (z *Cert) Rfc822Name() string {
    return C.GoString(C.CkCert_rfc822Name(z.ckObj))
}

// property: SelfSigned
// true if this is a self-signed certificate, otherwise false.
func (z *Cert) SelfSigned() bool {
    v := int(C.CkCert_getSelfSigned(z.ckObj))
    return v != 0
}

// property: SerialDecimal
// The certificate's serial number as a decimal string.
func (z *Cert) SerialDecimal() string {
    return C.GoString(C.CkCert_serialDecimal(z.ckObj))
}

// property: SerialNumber
// The certificate's serial number as a hexidecimal string.
func (z *Cert) SerialNumber() string {
    return C.GoString(C.CkCert_serialNumber(z.ckObj))
}

// property: Sha1Thumbprint
// Hexidecimal string of the SHA-1 thumbprint for the certificate. (This is the
// SHA1 hash of the binary DER representation of the entire X.509 certificate.)
func (z *Cert) Sha1Thumbprint() string {
    return C.GoString(C.CkCert_sha1Thumbprint(z.ckObj))
}

// property: SignatureVerified
// Returns true if the certificate and all certificates in the chain of authority
// have valid signatures, otherwise returns false.
func (z *Cert) SignatureVerified() bool {
    v := int(C.CkCert_getSignatureVerified(z.ckObj))
    return v != 0
}

// property: SubjectC
// The certificate subject's country.
func (z *Cert) SubjectC() string {
    return C.GoString(C.CkCert_subjectC(z.ckObj))
}

// property: SubjectCN
// The certificate subject's common name.
func (z *Cert) SubjectCN() string {
    return C.GoString(C.CkCert_subjectCN(z.ckObj))
}

// property: SubjectDN
// The certificate subject's full distinguished name.
func (z *Cert) SubjectDN() string {
    return C.GoString(C.CkCert_subjectDN(z.ckObj))
}

// property: SubjectE
// The certificate subject's email address.
func (z *Cert) SubjectE() string {
    return C.GoString(C.CkCert_subjectE(z.ckObj))
}

// property: SubjectKeyId
// The subject key identifier of the certificate in base64 string format. This is
// only present if the certificate contains the extension OID 2.5.29.14.
func (z *Cert) SubjectKeyId() string {
    return C.GoString(C.CkCert_subjectKeyId(z.ckObj))
}

// property: SubjectL
// The certificate subject's locality, which could be a city, count, township, or
// other geographic region.
func (z *Cert) SubjectL() string {
    return C.GoString(C.CkCert_subjectL(z.ckObj))
}

// property: SubjectO
// The certificate subject's organization, which is typically the company name.
func (z *Cert) SubjectO() string {
    return C.GoString(C.CkCert_subjectO(z.ckObj))
}

// property: SubjectOU
// The certificate subject's organizational unit, which is the unit within the
// organization.
func (z *Cert) SubjectOU() string {
    return C.GoString(C.CkCert_subjectOU(z.ckObj))
}

// property: SubjectS
// The certificate subject's state or province.
func (z *Cert) SubjectS() string {
    return C.GoString(C.CkCert_subjectS(z.ckObj))
}

// property: TrustedRoot
// Returns true if the certificate has a trusted root authority, otherwise
// returns false.
// 
// Note: As of version 9.5.0.41, the notion of what your application deems as
// trusted becomes more specific. The TrustedRoots class/object was added in
// v9.5.0.0. Prior to this, a certificate was considered to be anchored by a
// trusted root if the certificate chain could be established to a root
// (self-signed) certificate, AND if the root certificate was located somewhere in
// the Windows registry-based certificate stores. There are two problems with this:
// (1) it's a Windows-only solution. This property would always return false on
// non-Windows systems, and (2) it might be considered not a strong enough set of
// conditions for trusting a root certificate.
// 
// As of version 9.5.0.41, this property pays attention to the new TrustedRoots
// class/object, which allows for an application to specificallly indicate which
// root certificates are to be trusted. Certificates may be added to the
// TrustedRoots object via the LoadCaCertsPem or AddCert methods, and then
// activated by calling the TrustedRoots.Activate method. The activated trusted
// roots are deemed to be trusted in any Chilkat API method/property that needs to
// make this determination. In addition, the TrustedRoots object has a property
// named TrustSystemCaRoots, which defaults to true, which allows for backward
// compatibility. It will trust CA certificates stored in the Windows
// registry-based certificate stores, or if on Linux, will trust certificates found
// in /etc/ssl/certs/ca-certificates.crt.
//
func (z *Cert) TrustedRoot() bool {
    v := int(C.CkCert_getTrustedRoot(z.ckObj))
    return v != 0
}

// property: ValidFromStr
// The date (in RFC822 string format) that this certificate becomes (or became)
// valid. It is a GMT/UTC date that is returned.
func (z *Cert) ValidFromStr() string {
    return C.GoString(C.CkCert_validFromStr(z.ckObj))
}

// property: ValidToStr
// The date (in RFC822 string format) that this certificate becomes (or became)
// invalid. It is a GMT/UTC date that is returned.
func (z *Cert) ValidToStr() string {
    return C.GoString(C.CkCert_validToStr(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Cert) VerboseLogging() bool {
    v := int(C.CkCert_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Cert) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCert_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Cert) Version() string {
    return C.GoString(C.CkCert_version(z.ckObj))
}
// Returns 1 if the certificate has been revoked, 0 if not revoked, and -1 if
// unable to check the revocation status.
// 
// Note: This method is only implemented on Windows systems. It uses the underlying
// Microsoft CertVerifyRevocation Platform SDK function to check the revocation
// status of a certificate. (Search "CertVerifyRevocation" to get information about
// it.)
// 
// Non-Windows (and Windows) applications can send an OCSP request as shown in the
// example below.
//
func (z *Cert) CheckRevoked() int {

    v := C.CkCert_CheckRevoked(z.ckObj)


    return int(v)
}


// Verifies that the SmartCardPin property setting is correct. Returns 1 if
// correct, 0 if incorrect, and -1 if unable to check because the underlying CSP
// does not support the functionality.
func (z *Cert) CheckSmartCardPin() int {

    v := C.CkCert_CheckSmartCardPin(z.ckObj)


    return int(v)
}


// Exports the digital certificate to ASN.1 DER format.
func (z *Cert) ExportCertDer() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCert_ExportCertDer(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Exports the digital certificate in ASN.1 DER format to a BinData object.
func (z *Cert) ExportCertDerBd(arg1 *BinData) bool {

    v := C.CkCert_ExportCertDerBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Exports the digital certificate to ASN.1 DER format binary file.
func (z *Cert) ExportCertDerFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_ExportCertDerFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Exports the digital certificate to an unencrypted PEM formatted string.
// return a string or nil if failed.
func (z *Cert) ExportCertPem() *string {

    retStr := C.CkCert_exportCertPem(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Exports the digital certificate to an unencrypted PEM formatted file.
func (z *Cert) ExportCertPemFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_ExportCertPemFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Exports a certificate to an XML format where the XML tags are the names of the
// ASN.1 objects that compose the X.509 certificate. Binary data is either hex or
// base64 encoded. (The binary data for a "bits" ASN.1 tag is hex encoded, whereas
// for all other ASN.1 tags, such as "octets", it is base64.)
// return a string or nil if failed.
func (z *Cert) ExportCertXml() *string {

    retStr := C.CkCert_exportCertXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Exports the certificate's private key.
func (z *Cert) ExportPrivateKey() *PrivateKey {

    retObj := C.CkCert_ExportPrivateKey(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &PrivateKey{retObj}
}


// Exports the certificate's public key.
func (z *Cert) ExportPublicKey() *PublicKey {

    retObj := C.CkCert_ExportPublicKey(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &PublicKey{retObj}
}


// Exports the certificate and private key (if available) to pfxData. The password is what
// will be required to access the PFX contents at a later time. If includeCertChain is true,
// then the certificates in the chain of authority are also included in the PFX.
func (z *Cert) ExportToPfxBd(arg1 string, arg2 bool, arg3 *BinData) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkCert_ExportToPfxBd(z.ckObj, cstr1, b2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Exports the certificate and private key (if available) to an in-memory PFX
// image. The password is what will be required to access the PFX contents at a later
// time. If includeCertChain is true, then the certificates in the chain of authority are
// also included in the PFX.
func (z *Cert) ExportToPfxData(arg1 string, arg2 bool) []byte {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCert_ExportToPfxData(z.ckObj, cstr1, b2, ckOutBytes)

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


// Exports the certificate and private key (if available) to a PFX (.pfx or .p12)
// file. The output PFX is secured using the pfxPassword. If bIncludeCertChain is true, then the
// certificates in the chain of authority are also included in the PFX output file.
func (z *Cert) ExportToPfxFile(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkCert_ExportToPfxFile(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Finds and returns the issuer certificate. If the certificate is a root or
// self-issued, then the certificate returned is a copy of the caller certificate.
// (The IsRoot property can be check to see if the certificate is a root (or
// self-issued) certificate.)
func (z *Cert) FindIssuer() *Cert {

    retObj := C.CkCert_FindIssuer(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns a certficate chain object containing all the certificates (including
// this one), in the chain of authentication to the trusted root (if possible). If
// this certificate object was loaded from a PFX, then the certiicates contained in
// the PFX are automatically available for building the certificate chain. The
// UseCertVault method can be called to provide additional certificates that might
// be required to build the cert chain. Finally, the TrustedRoots object can be
// used to provide a way of making trusted root certificates available.
// 
// Note: Prior to v9.5.0.50, this method would fail if the certificate chain could
// not be completed to the root. Starting in v9.5.0.50, the incomplete certificate
// chain will be returned. The certificate chain's ReachesRoot property can be
// examined to see if the chain was completed to the root.
// 
// On Windows systems, the registry-based certificate stores are automatically
// consulted if needed to locate intermediate or root certificates in the chain.
// Chilkat searches certificate stores in the following order. SeeSystem Store
// Locations
// <https://docs.microsoft.com/en-us/windows/desktop/seccrypto/system-store-location
// s> for more information.
//     Current-User "CA" Certificate Store
//     Local-Machine "CA" Certificate Store
//     Current-User "Root" Certificate Store
//     Local-Machine "Root" Certificate Store
//     Current-User "MY" Certificate Store
//     Local-Machine "MY" Certificate Store
//     Current-User "ADDRESSBOOK" Certificate Store (if it exists)
//     Local-Machine "ADDRESSBOOK" Certificate Store (if it exists)
//
func (z *Cert) GetCertChain() *CertChain {

    retObj := C.CkCert_GetCertChain(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CertChain{retObj}
}


// Returns a base64 encoded string representation of the certificate's binary DER
// format, which can be passed to SetFromEncoded to recreate the certificate
// object.
// return a string or nil if failed.
func (z *Cert) GetEncoded() *string {

    retStr := C.CkCert_getEncoded(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the certificate extension data in XML format (converted from ASN.1). The
// oid is an OID, such as the ones listed here:
// http://www.alvestrand.no/objectid/2.5.29.html
// 
// Note: In many cases, the data within the XML is returned base64 encoded. An
// application may need to take one further step to base64 decode the information
// contained within the XML.
//
// return a string or nil if failed.
func (z *Cert) GetExtensionAsXml(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCert_getExtensionAsXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Exports the certificate's private key to a PEM string (if the private key is
// available).
// return a string or nil if failed.
func (z *Cert) GetPrivateKeyPem() *string {

    retStr := C.CkCert_getPrivateKeyPem(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the SPKI Fingerprint suitable for use in pinning. (See RFC 7469.) An
// SPKI Fingerprint is defined as the output of a known cryptographic hash
// algorithm whose input is the DER-encoded ASN.1 representation of the Subject
// Public Key Info (SPKI) of an X.509 certificate. The hashAlg specifies the hash
// algorithm and may be "sha256", "sha384", "sha512", "sha1", "md2", "md5",
// "haval", "ripemd128", "ripemd160","ripemd256", or "ripemd320". The encoding
// specifies the encoding, and may be "base64", "hex", or any of the encoding modes
// specified in the article at the link below.
// return a string or nil if failed.
func (z *Cert) GetSpkiFingerprint(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkCert_getSpkiFingerprint(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the date/time this certificate becomes (or became) valid.
func (z *Cert) GetValidFromDt() *CkDateTime {

    retObj := C.CkCert_GetValidFromDt(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Returns the date/time this certificate becomes (or became) invalid.
func (z *Cert) GetValidToDt() *CkDateTime {

    retObj := C.CkCert_GetValidToDt(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Returns an encoded hash of a particular part of the certificate. The part may be
// one of the following:
//     IssuerDN
//     IssuerPublicKey
//     SubjectDN
//     SubjectPublicKey
// 
// The hashAlg is the name of the hash algorithm, such as "sha1", "sha256", "sha384",
// "sha512", "md5", etc. The encoding is the format to return, such as "hex", "base64",
// etc.
//
// return a string or nil if failed.
func (z *Cert) HashOf(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkCert_hashOf(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if a private key associated with the certificate is available.
func (z *Cert) HasPrivateKey() bool {

    v := C.CkCert_HasPrivateKey(z.ckObj)


    return v != 0
}


// (Relevant only when running on a Microsoft Windows operating system.) Searches
// the Windows Local Machine and Current User registry-based certificate stores for
// a certificate having the common name specified. If found, the certificate is
// loaded and ready for use.
func (z *Cert) LoadByCommonName(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_LoadByCommonName(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// (Relevant only when running on a Microsoft Windows operating system.) Searches
// the Windows Local Machine and Current User registry-based certificate stores for
// a certificate containing the email address specified. If found, the certificate
// is loaded and ready for use.
func (z *Cert) LoadByEmailAddress(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_LoadByEmailAddress(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// (Relevant only when running on a Microsoft Windows operating system.) Searches
// the Windows Local Machine and Current User registry-based certificate stores for
// a certificate matching the issuerCN and having an issuer matching the serialNumber. If
// found, the certificate is loaded and ready for use.
func (z *Cert) LoadByIssuerAndSerialNumber(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCert_LoadByIssuerAndSerialNumber(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// (Relevant only when running on a Microsoft Windows operating system.) Searches
// the Windows Local Machine and Current User registry-based certificate stores for
// a certificate having an MD5 or SHA1 thumbprint equal to the thumbprint. The hash (i.e.
// thumbprint) is passed as a string using the encoding specified by encoding (such as
// "base64", "hex", etc.).
func (z *Cert) LoadByThumbprint(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCert_LoadByThumbprint(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads an ASN.1 or DER encoded certificate represented in a Base64 string.
func (z *Cert) LoadFromBase64(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_LoadFromBase64(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads an X.509 certificate from data contained in certBytes.
// 
// Note: The certBytes may contain the certificate in any format. It can be binary DER
// (ASN.1), PEM, Base64, etc. Chilkat will automatically detect the format.
//
func (z *Cert) LoadFromBd(arg1 *BinData) bool {

    v := C.CkCert_LoadFromBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads an X.509 certificate from ASN.1 DER encoded bytes.
// 
// Note: The data may contain the certificate in any format. It can be binary DER
// (ASN.1), PEM, Base64, etc. Chilkat will automatically detect the format.
//
func (z *Cert) LoadFromBinary(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkCert_LoadFromBinary(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Loads a certificate from a .cer, .crt, .p7b, or .pem file. This method accepts
// certificates from files in any of the following formats:
// 1. DER encoded binary X.509 (.CER)
// 2. Base-64 encoded X.509 (.CER)
// 3. Cryptographic Message Syntax Standard - PKCS #7 Certificates (.P7B)
// 4. PEM format
// This method decodes the certificate based on the contents if finds within the
// file, and not based on the file extension. If your certificate is in a file
// having a different extension, try loading it using this method before assuming
// it won't work. This method does not load .p12 or .pfx (PKCS #12) files.
func (z *Cert) LoadFromFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_LoadFromFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the X.509 certificate from the smartcard currently in the reader, or from
// a USB token.
// 
// If the smartcard contains multiple certificates, this method arbitrarily picks
// one. To load a specific certificate, call CertStore.OpenSmartcard and then
// iterate over the certs to find the desired certificate.
// 
// The csp can be set to the name of the CSP (Cryptographic Service Provider) that
// should be used. If csp is an empty string, then the 1st CSP found matching one
// of the following names will be used:
// 
//     Microsoft Smart Card Key Storage Provider
//     Microsoft Base Smart Card Crypto Provider
//     Bit4id Universal Middleware Provider
//     YubiHSM Key Storage Provider (starting in v9.5.0.83)
//     eToken Base Cryptographic Provider
//     FTSafe ePass1000 RSA Cryptographic Service Provider
//     SecureStoreCSP
//     EnterSafe ePass2003 CSP v2.0
//     Gemalto Classic Card CSP
//     PROXKey CSP India V1.0
//     PROXKey CSP India V2.0
//     TRUST KEY CSP V1.0
//     Watchdata Brazil CSP V1.0
//     Luna Cryptographic Services for Microsoft Windows
//     Luna SChannel Cryptographic Services for Microsoft Windows
//     Safenet RSA Full Cryptographic Provider
//     nCipher Enhanced Cryptographic Provider
//     SafeSign Standard Cryptographic Service Provider
//     SafeSign Standard RSA and AES Cryptographic Service Provider
//     MySmartLogon NFC CSP
//     NFC Connector Enterprise
//     ActivClient Cryptographic Service Provider
//     EnterSafe ePass2003 CSP v1.0
//     Oberthur Card Systems Cryptographic Provider
//     Athena ASECard Crypto CSP"
func (z *Cert) LoadFromSmartcard(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_LoadFromSmartcard(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the certificate from a PEM string.
func (z *Cert) LoadPem(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_LoadPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the certificate from the PFX contained in pfxData. Note: If the PFX contains
// multiple certificates, the 1st certificate in the PFX is loaded.
func (z *Cert) LoadPfxBd(arg1 *BinData, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkCert_LoadPfxBd(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a PFX from an in-memory image of a PFX file. Note: If the PFX contains
// multiple certificates, the 1st certificate in the PFX is loaded.
func (z *Cert) LoadPfxData(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkCert_LoadPfxData(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a PFX file. Note: If the PFX contains multiple certificates, the 1st
// certificate in the PFX is loaded.
func (z *Cert) LoadPfxFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCert_LoadPfxFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads the certificate from a completed asynchronous task.
func (z *Cert) LoadTaskResult(arg1 *Task) bool {

    v := C.CkCert_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


// Converts a PEM file to a DER file.
func (z *Cert) PemFileToDerFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCert_PemFileToDerFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Saves a certificate object to a .cer file.
func (z *Cert) SaveToFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_SaveToFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Initializes the certificate object from a base64 encoded string representation
// of the certificate's binary DER format.
func (z *Cert) SetFromEncoded(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_SetFromEncoded(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Used to associate a private key with the certificate for subsequent (PKCS7)
// signature creation or decryption.
func (z *Cert) SetPrivateKey(arg1 *PrivateKey) bool {

    v := C.CkCert_SetPrivateKey(z.ckObj, arg1.ckObj)


    return v != 0
}


// Same as SetPrivateKey, but the key is provided in unencrypted PEM format. (Note:
// The privKeyPem is not a file path, it is the actual PEM text.)
func (z *Cert) SetPrivateKeyPem(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCert_SetPrivateKeyPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds an XML certificate vault to the object's internal list of sources to be
// searched for certificates for help in building certificate chains and verifying
// the certificate signature to the trusted root.
func (z *Cert) UseCertVault(arg1 *XmlCertVault) bool {

    v := C.CkCert_UseCertVault(z.ckObj, arg1.ckObj)


    return v != 0
}


// Verifies the certificate signature, as well as the signatures of all
// certificates in the chain of authentication to the trusted root. Returns true
// if all signatures are verified to the trusted root. Otherwise returns false.
func (z *Cert) VerifySignature() bool {

    v := C.CkCert_VerifySignature(z.ckObj)


    return v != 0
}


// Returns the base64 representation of an X509PKIPathv1 containing just the
// calling certificate. This is typically used in an X.509 Binary Security Token.
// It is a PKIPath that contains an ordered list of X.509 public certificates
// packaged in a PKIPath. The X509PKIPathv1 token type may be used to represent a
// certificate path. (This is sometimes used in XAdES signatures.)
// return a string or nil if failed.
func (z *Cert) X509PKIPathv1() *string {

    retStr := C.CkCert_x509PKIPathv1(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


