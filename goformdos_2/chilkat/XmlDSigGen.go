// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkXmlDSigGen.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewXmlDSigGen() *XmlDSigGen {
	return &XmlDSigGen{C.CkXmlDSigGen_Create()}
}

func (z *XmlDSigGen) DisposeXmlDSigGen() {
    if z != nil {
	C.CkXmlDSigGen_Dispose(z.ckObj)
	}
}




// property: Behaviors
// A comma-separated list of keywords to specify special behaviors to work around
// potential oddities or special requirements needed for providing signatures to
// particular systems. This is an open-ended property where new behaviors can be
// implemented depending on the needs encountered by Chilkat customers. The
// possible behaviors are listed below.
//     AttributeSortingBug (introduced in v9.5.0.79) Tells Chilkat to produce a
//     signature that duplicates a common XML canonicalization attribute sorting bug
//     found in some XML signature implementations (such as JPK VAT signed XML
//     documents for Polish government, i.e. mf.gov.pl, csioz.gov.pl, crd.gov.pl, etc).
//     SeeXML Signature Canonicalization Bug
//     <http://cknotes.com/xml-signature-canonicalization-bug-in-widely-used-softwar
//     e/> for details.
//     ForceAddEnvelopedSignatureTransform The "_LT_Transform
//     Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature" /> " is
//     normally only added when the Signature is contained within the XML fragment that
//     is signed. The meaning of this tranformation is to tell the verifier to remove
//     the Signature from the data prior to canonicalizing. If the Signature is not
//     contained within the XML fragment that was signed, then the signature was not
//     enveloped. There would be no need to remove the Signature because the Signature
//     is not contained in the XML fragment being verified. However.. some brain-dead
//     verifiying systems require this Transform to be present regardless of whether it
//     makes sense. This behavior will cause Chilkat to add the Transform regardless.
//     NoEnvelopedSignatureTransform (introduced in v9.5.0.82) Prevents the
//     "_LT_Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"
//     /> " from being added in all cases.
//     ebXmlTransform (introduced in v9.5.0.73) Causes the following tranform to be
//     added for ebXml messages:    
// _LT_Transform Algorithm="http://www.w3.org/TR/1999/REC-xpath-19991116">    
//     _LT_XPath xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/">not(ancestor-or-self::node()[@SOAP-ENV:actor="urn:oasis:names:tc:ebxml-msg:actor:nextMSH"]    
//          | ancestor-or-self::node()[@SOAP-ENV:actor="http://schemas.xmlsoap.org/soap/actor/next"])_LT_/XPath>    
// _LT_/Transform>    
//     TransformSignatureXPath (introduced in v9.5.0.75) Causes the following
//     tranform to be added:    
// _LT_ds:Transform Algorithm="http://www.w3.org/TR/1999/REC-xpath-19991116"_GT_    
//    _LT_ds:XPath_GT_not(ancestor-or-self::ds:Signature)_LT_/ds:XPath_GT_    
// _LT_/ds:Transform_GT_    
//     CompactSignedXml (introduced in v9.5.0.73) The passed-in XML to be signed is
//     first reformatted to a compact representation by removing all CR's, LF's, and
//     unnecessary whitespace so that the XML to be signed is on a single line. The
//     resulting XML (with signature) is also entirely contained on a single line. (If
//     an XML declarator is present, then it will remain on it's own line.)
//     IndentedSignature (introduced in v9.5.0.73) Causes the XML Signature to be
//     produced on multiple lines with indentation for easier human readability. The
//     CompactSignedXml behavior takes precedence over this behavior.
//     FullLocalSigningTime (introduced in v9.5.0.76) Causes the signing time to be
//     formatted like this: 2017-05-20T19:16:05.649+01:00.nnn, where the ".nnn" is
//     added to indicate milliseconds.
//     LocalSigningTime (introduced in v9.5.0.76) Causes the signing time to be
//     formatted using a local time (with a timezone offset such as "+01:00" rather
//     than "Z" to signify GMT).
//     DnReverseOrder (introduced in v9.5.0.77) Causes DN's (certificate
//     Distinguished Names) to be written in reverse order. Reverse order leads with
//     "CN", such as "CN=..., O=..., OU=..., C=...", whereas normal order ends with
//     "CN", such as "C=..., OU=..., O=..., CN=..."
//     IssuerSerialHex (introduced in v9.5.0.77) Causes the issuer serial number
//     located in SignedProperties.SignedSignatureProperties.SigningCertificate to be
//     emitted as uppercase hex instead of decimal. (Also, when signing XML for
//     e-dokumenty.mf.gov.pl, Chilkat automatically recognizes it and uses
//     IssuerSerialHex.)
//     IssuerSerialHexLower (introduced in v9.5.0.77) Causes the issuer serial
//     number located in SignedProperties.SignedSignatureProperties.SigningCertificate
//     to be emitted as lowercase hex instead of decimal.
//     SigningTimeAdjust-_LT_numSeconds_GT_ (introduced in v9.5.0.80) When Chilkat
//     automatically fills in the value for a SigningTime, it will use the current
//     system date/time. This behavior can be used to adjust the generate time to
//     numSeconds in the past. For example: "SigningTimeAdjust-60" will generate a
//     signing time 60 seconds prior to the current time.
func (z *XmlDSigGen) Behaviors() string {
    return C.GoString(C.CkXmlDSigGen_behaviors(z.ckObj))
}
// property setter: Behaviors
func (z *XmlDSigGen) SetBehaviors(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putBehaviors(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CustomKeyInfoXml
// Specifies custom XML to be inserted in the KeyInfo element of the Signature. A
// common use is to provide a wsse:SecurityTokenReference fragment of XML.
func (z *XmlDSigGen) CustomKeyInfoXml() string {
    return C.GoString(C.CkXmlDSigGen_customKeyInfoXml(z.ckObj))
}
// property setter: CustomKeyInfoXml
func (z *XmlDSigGen) SetCustomKeyInfoXml(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putCustomKeyInfoXml(z.ckObj,cStr)
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
func (z *XmlDSigGen) DebugLogFilePath() string {
    return C.GoString(C.CkXmlDSigGen_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *XmlDSigGen) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IncNamespacePrefix
// The namespace prefix to use for InclusiveNamespaces elements. The default value
// is "ec". Set this property to the empty string to omit an InclusiveNamespaces
// prefix. For example, given the default values of IncNamespaceUri and
// IncNamespacePrefix, generated InclusiveNamespaces elements will appear like
// this:
// ...
func (z *XmlDSigGen) IncNamespacePrefix() string {
    return C.GoString(C.CkXmlDSigGen_incNamespacePrefix(z.ckObj))
}
// property setter: IncNamespacePrefix
func (z *XmlDSigGen) SetIncNamespacePrefix(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putIncNamespacePrefix(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IncNamespaceUri
// The namespace URI for any InclusiveNamespaces elements that are created. The
// default value is "http://www.w3.org/2001/10/xml-exc-c14n#". For example, if the
// IncNamespacePrefix equals "ec" and this property remains at the default value,
// then the generated Signature element will be:
// ...
func (z *XmlDSigGen) IncNamespaceUri() string {
    return C.GoString(C.CkXmlDSigGen_incNamespaceUri(z.ckObj))
}
// property setter: IncNamespaceUri
func (z *XmlDSigGen) SetIncNamespaceUri(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putIncNamespaceUri(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: KeyInfoId
// If set, causes the generated KeyInfo element to include an Id attribute with
// this value. For example:
// ...
//    _LT_ds:KeyInfo Id="KeyInfo"_GT_
//       _LT_ds:X509Data_GT_
//          _LT_ds:X509SubjectName_GT_CERTIFICADO DE ABC_LT_/ds:X509SubjectName_GT_
//          _LT_ds:X509Certificate_GT_MIIITTCC....fIsIZeZOeQ=_LT_/ds:X509Certificate_GT_
//       _LT_/ds:X509Data_GT_
//    _LT_/ds:KeyInfo_GT_
// ...
func (z *XmlDSigGen) KeyInfoId() string {
    return C.GoString(C.CkXmlDSigGen_keyInfoId(z.ckObj))
}
// property setter: KeyInfoId
func (z *XmlDSigGen) SetKeyInfoId(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putKeyInfoId(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: KeyInfoKeyName
// Specifies the KeyName to be inserted in the KeyInfo element of the Signature if
// the KeyInfoType equals "KeyName".
func (z *XmlDSigGen) KeyInfoKeyName() string {
    return C.GoString(C.CkXmlDSigGen_keyInfoKeyName(z.ckObj))
}
// property setter: KeyInfoKeyName
func (z *XmlDSigGen) SetKeyInfoKeyName(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putKeyInfoKeyName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: KeyInfoType
// Specifies the type of information that will be included in the optional KeyInfo
// element of the Signature. Possible values are:
//     None
//     KeyName
//     KeyValue
//     X509Data
//     X509Data+KeyValue
//     Custom
// 
// The default value is "KeyValue". The "X509Data+KeyValue" option was added in
// Chilkat v9.5.0.73.
// 
// If None, then no KeyInfo element is added to the Signature when generated.
// 
// If KeyValue, then the KeyInfo will contain the public key (RSA, DSA, or ECDSA).
// 
// If X509Data, then the KeyInfo will contain information about an X.509
// certificate as specified by the X509Type property.
// 
// If Custom, then the KeyInfo will contain the custom XML contained in the
// CustomKeyInfoXml property.
//
func (z *XmlDSigGen) KeyInfoType() string {
    return C.GoString(C.CkXmlDSigGen_keyInfoType(z.ckObj))
}
// property setter: KeyInfoType
func (z *XmlDSigGen) SetKeyInfoType(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putKeyInfoType(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *XmlDSigGen) LastErrorHtml() string {
    return C.GoString(C.CkXmlDSigGen_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *XmlDSigGen) LastErrorText() string {
    return C.GoString(C.CkXmlDSigGen_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *XmlDSigGen) LastErrorXml() string {
    return C.GoString(C.CkXmlDSigGen_lastErrorXml(z.ckObj))
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
func (z *XmlDSigGen) LastMethodSuccess() bool {
    v := int(C.CkXmlDSigGen_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *XmlDSigGen) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmlDSigGen_putLastMethodSuccess(z.ckObj,v)
}

// property: SigId
// An option Id attribute value for the Signature element. The default value is the
// empty string, which generates a Signature element with no Id attribute. For
// example:
// If this property is set to "abc123", then the Signature element would be generated like this:
func (z *XmlDSigGen) SigId() string {
    return C.GoString(C.CkXmlDSigGen_sigId(z.ckObj))
}
// property setter: SigId
func (z *XmlDSigGen) SetSigId(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSigId(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SigLocation
// Indicates where the Signature is to be located within the XML that is signed.
// This is a path to the position in the XML where the Signature will be inserted,
// using Chilkat path syntax (using vertical bar characters to delimit tag names.
// If the Signature element is to be the root of XML document, then set this
// property equal to the empty string.
// 
// For example, if we have the following SOAP XML and wish to insert the Signature
// at the indicated location, then the SigLocation property should be set to
// "SOAP-ENV:Envelope|SOAP-ENV:Header|wsse:Security".
// ** The XML Signature is to be inserted here **
// 	...
//
func (z *XmlDSigGen) SigLocation() string {
    return C.GoString(C.CkXmlDSigGen_sigLocation(z.ckObj))
}
// property setter: SigLocation
func (z *XmlDSigGen) SetSigLocation(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSigLocation(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SigLocationMod
// Modifies the placement of the signature at the location specified by
// SigLocation. Possible values are:
//     0: Insert the Signature as the last child of the element at SigLocation.
//     This is the default.
//     1: Insert the Signature as a sibling directly after the element at
//     SigLocation.
//     2: Insert the Signature as a sibling directly before the element at
//     SigLocation.
func (z *XmlDSigGen) SigLocationMod() int {
    return int(C.CkXmlDSigGen_getSigLocationMod(z.ckObj))
}
// property setter: SigLocationMod
func (z *XmlDSigGen) SetSigLocationMod(value int) {
    C.CkXmlDSigGen_putSigLocationMod(z.ckObj,C.int(value))
}

// property: SigNamespacePrefix
// The namespace prefix of the Signature that is to be created. The default value
// is "ds". Set this property to the empty string to omit a Signature namespace URI
// and prefix. For example, given the default values of SigNamespaceUri and
// SigNamespacePrefix, the generated Signature element will be:
// ...
func (z *XmlDSigGen) SigNamespacePrefix() string {
    return C.GoString(C.CkXmlDSigGen_sigNamespacePrefix(z.ckObj))
}
// property setter: SigNamespacePrefix
func (z *XmlDSigGen) SetSigNamespacePrefix(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSigNamespacePrefix(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SigNamespaceUri
// The namespace URI of the Signature that is to be created. The default value is
// "http://www.w3.org/2000/09/xmldsig#". For example, if the SigNamespacePrefix
// equals "ds" and this property remains at the default value, then the generated
// Signature element will be:
// ...
func (z *XmlDSigGen) SigNamespaceUri() string {
    return C.GoString(C.CkXmlDSigGen_sigNamespaceUri(z.ckObj))
}
// property setter: SigNamespaceUri
func (z *XmlDSigGen) SetSigNamespaceUri(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSigNamespaceUri(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SignedInfoCanonAlg
// The canonicalization method to be used for the SignedInfo when creating the XML
// signature.
//     "C14N" -- for Inclusive Canonical XML (without comments)
//     "C14N_11" -- for Inclusive Canonical XML 1.1 (without comments)
//     "EXCL_C14N" -- for Exclusive Canonical XML (without comments)
//     "C14N_WithComments" -- for Inclusive Canonical XML (with comments)
//     "C14N_11_WithComments" -- for Inclusive Canonical XML 1.1 (with comments)
//     "EXCL_C14N_WithComments" -- for Exclusive Canonical XML (with comments)
//     Note: The WithComments options are available in Chilkat v9.5.0.71 and later.
// 
// The default value is "EXCL_C14N".
//
func (z *XmlDSigGen) SignedInfoCanonAlg() string {
    return C.GoString(C.CkXmlDSigGen_signedInfoCanonAlg(z.ckObj))
}
// property setter: SignedInfoCanonAlg
func (z *XmlDSigGen) SetSignedInfoCanonAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSignedInfoCanonAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SignedInfoDigestMethod
// The digest method to be used for signing the SignedInfo part of the Signature.
// Possible values are "sha1", "sha256", "sha384", and "sha512". The default is
// "sha256".
func (z *XmlDSigGen) SignedInfoDigestMethod() string {
    return C.GoString(C.CkXmlDSigGen_signedInfoDigestMethod(z.ckObj))
}
// property setter: SignedInfoDigestMethod
func (z *XmlDSigGen) SetSignedInfoDigestMethod(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSignedInfoDigestMethod(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SignedInfoId
// Optional Id attribute to be added to the SignedInfo element. The default value
// is the empty string, meaning that the SignedInfo is generated without an Id
// attribute.
func (z *XmlDSigGen) SignedInfoId() string {
    return C.GoString(C.CkXmlDSigGen_signedInfoId(z.ckObj))
}
// property setter: SignedInfoId
func (z *XmlDSigGen) SetSignedInfoId(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSignedInfoId(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SignedInfoPrefixList
// The inclusive namespace prefix list to be added, if any, when the
// SignedInfoCanonAlg is equal to "EXCL_C14N". The defautl value is the empty
// string. If namespaces are listed, they are separated by space characters.
// 
// If, for example, this property is set to "wsse SOAP-ENV", then the
// CanonicalizationMethod part of the SignedInfo that is generated would look like
// this:
// ...
//
func (z *XmlDSigGen) SignedInfoPrefixList() string {
    return C.GoString(C.CkXmlDSigGen_signedInfoPrefixList(z.ckObj))
}
// property setter: SignedInfoPrefixList
func (z *XmlDSigGen) SetSignedInfoPrefixList(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSignedInfoPrefixList(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SigningAlg
// Selects the signature algorithm to be used when using an RSA key to sign. The
// default value is "PKCS1-v1_5". This can be set to "RSASSA-PSS" (or simply "pss")
// to use the RSASSA-PSS signature scheme.
// 
// Note: This property only applies when signing with an RSA private key. It does
// not apply for ECC or DSA private keys.
//
func (z *XmlDSigGen) SigningAlg() string {
    return C.GoString(C.CkXmlDSigGen_signingAlg(z.ckObj))
}
// property setter: SigningAlg
func (z *XmlDSigGen) SetSigningAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSigningAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SigValueId
// An option Id attribute value for the SignatureValue element. The default value
// is the empty string, which generates a SignatureValue element with no Id
// attribute. For example:
// _LT_ds:SignatureValue_GT_
// If this property is set to "value-id-7d4a", then the Signature element would be
// generated like this:
// _LT_ds:SignatureValue  Id="value-id-7d4a"_GT_
func (z *XmlDSigGen) SigValueId() string {
    return C.GoString(C.CkXmlDSigGen_sigValueId(z.ckObj))
}
// property setter: SigValueId
func (z *XmlDSigGen) SetSigValueId(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putSigValueId(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *XmlDSigGen) VerboseLogging() bool {
    v := int(C.CkXmlDSigGen_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *XmlDSigGen) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmlDSigGen_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *XmlDSigGen) Version() string {
    return C.GoString(C.CkXmlDSigGen_version(z.ckObj))
}

// property: X509Type
// Specifies the kind of X.509 certificate information is provided in the KeyInfo
// element when the KeyInfoType equals "X509Data". Possible values are:
//     Certificate
//     CertChain
//     IssuerSerial
//     SubjectName
//     SKI
// 
// The default value is "Certificate".
// 
// Note: This property can be set to a comma-separated list of the keywords above.
// For example, If set to "SubjectName,Certificate", then both the X509SubjectName
// and X509Certificate parts will be added to the KeyInfo.
// 
// If Certificate, then the KeyInfo will contain the base64 encoded X.509v3
// certificate.
// 
// If CertChain, then the KeyInfo will contain the base64 encoded X.509v3
// certificate as well as any certificates available in the chain of authentication
// to the root cert.
// 
// If IssuerSerial, then the KeyInfo will contain the X.509 issuer's distinguished
// name and the signing certificate's serial number.
// 
// If SubjectName, then the KeyInfo will contain the X.509 subject distinguished
// name.
// 
// If SKI, then the KeyInfo will contain the base64 encoded value of the cert's
// X.509 SubjectKeyIdentifier extension.
//
func (z *XmlDSigGen) X509Type() string {
    return C.GoString(C.CkXmlDSigGen_x509Type(z.ckObj))
}
// property setter: X509Type
func (z *XmlDSigGen) SetX509Type(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSigGen_putX509Type(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}
// Specifies an enveloped Reference to be added to the Signature when generated. An
// enveloped Reference is for data contained within the Signature. (The Signature
// is to be an enveloping signature, and the data is enveloped by the Signature.)
// 
// The id is the value of the Id attribute of the Object element that is to be
// contained within the generated Signature. The content is the text content to be
// contained in the Object. Binary data can be signed by passing the bytes in content
// in an encoded format (such as base64 or hex).
// 
// The digestMethod is the digest method and can be one of the following: "sha1", "sha256",
// "sha384", "sha512", "ripemd160", or "md5".
// 
// The canonMethod is the canonicalization method, and can be one of the following.
//     "C14N" -- for Inclusive Canonical XML (without comments)
//     "C14N_11" -- for Inclusive Canonical XML 1.1 (without comments)
//     "EXCL_C14N" -- for Exclusive Canonical XML (without comments)
//     "C14N_WithComments" -- for Inclusive Canonical XML (with comments)
//     "C14N_11_WithComments" -- for Inclusive Canonical XML 1.1 (with comments)
//     "EXCL_C14N_WithComments" -- for Exclusive Canonical XML (with comments)
//     Note: The WithComments options are available in Chilkat v9.5.0.71 and later.
// 
// The refType is optional and is usually not needed. Set this to the empty string
// unless it is desired to add a Type attribute to the Reference that is advisory
// only.
//
func (z *XmlDSigGen) AddEnvelopedRef(arg1 string, arg2 *StringBuilder, arg3 string, arg4 string, arg5 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkXmlDSigGen_AddEnvelopedRef(z.ckObj, cstr1, arg2.ckObj, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Specifies an external non-XML binary data Reference to be added to the Signature
// when generated.
// 
// The uri is the value of the URI attribute of the Reference.
// 
// The content contains the binary data to be digested according to the digestMethod.
// 
// The digestMethod is the digest method and can be one of the following: "sha1", "sha256",
// "sha384", "sha512", "ripemd160", or "md5".
// 
// The refType is optional and is usually not needed. Set this to the empty string
// unless it is desired to add a Type attribute to the Reference that is advisory
// only.
//
func (z *XmlDSigGen) AddExternalBinaryRef(arg1 string, arg2 *BinData, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkXmlDSigGen_AddExternalBinaryRef(z.ckObj, cstr1, arg2.ckObj, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Specifies an external file Reference to be added to the Signature when
// generated.
// 
// The uri is the value of the URI attribute of the Reference. It can (and likely
// will) be different than the localFilePath which is the path to the local file to be
// added. (The local file is not read until the XML digital signature is actually
// created.)
// 
// The digestMethod is the digest method and can be one of the following: "sha1", "sha256",
// "sha384", "sha512", "ripemd160", or "md5".
// 
// The refType is optional and is usually not needed. Set this to the empty string
// unless it is desired to add a Type attribute to the Reference that is advisory
// only.
//
func (z *XmlDSigGen) AddExternalFileRef(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkXmlDSigGen_AddExternalFileRef(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Specifies an external non-XML text data Reference to be added to the Signature
// when generated.
// 
// The uri is the value of the URI attribute of the Reference.
// 
// The content contains the non-XML data to be digested according to the charset. The
// charset specifies the charset (such as "utf-8", "windows-1252", etc.) for the byte
// reprsentation of the text to be digested. The includeBom indicates whether the BOM
// (Byte Order Mark, also known as the preamble) is included in the byte
// representation that is digested.
// 
// The digestMethod is the digest method and can be one of the following: "sha1", "sha256",
// "sha384", "sha512", "ripemd160", or "md5".
// 
// The refType is optional and is usually not needed. Set this to the empty string
// unless it is desired to add a Type attribute to the Reference that is advisory
// only.
//
func (z *XmlDSigGen) AddExternalTextRef(arg1 string, arg2 *StringBuilder, arg3 string, arg4 bool, arg5 string, arg6 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    cstr5 := C.CString(arg5)
    cstr6 := C.CString(arg6)

    v := C.CkXmlDSigGen_AddExternalTextRef(z.ckObj, cstr1, arg2.ckObj, cstr3, b4, cstr5, cstr6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr5))
    C.free(unsafe.Pointer(cstr6))

    return v != 0
}


// Specifies an external XML Reference to be added to the Signature when generated.
// 
// The uri is the value of the URI attribute of the Reference.
// 
// The content contains the XML document to be referenced.
// 
// The digestMethod is the digest method and can be one of the following: "sha1", "sha256",
// "sha384", "sha512", "ripemd160", or "md5".
// 
// The canonMethod is the canonicalization method, and can be one of the following.
//     "C14N" -- for Inclusive Canonical XML (without comments)
//     "C14N_11" -- for Inclusive Canonical XML 1.1 (without comments)
//     "EXCL_C14N" -- for Exclusive Canonical XML (without comments)
//     "C14N_WithComments" -- for Inclusive Canonical XML (with comments)
//     "C14N_11_WithComments" -- for Inclusive Canonical XML 1.1 (with comments)
//     "EXCL_C14N_WithComments" -- for Exclusive Canonical XML (with comments)
//     "" -- An empty string indicates that no transformation should be included /
//     applied for this reference.
//     Note: The WithComments options are available in Chilkat v9.5.0.71 and later.
//     Note: The empty-string canonMethod is available in Chilkat v9.5.0.75 and
//     later.
// 
// The refType is optional and is usually not needed. Set this to the empty string
// unless it is desired to add a Type attribute to the Reference that is advisory
// only.
//
func (z *XmlDSigGen) AddExternalXmlRef(arg1 string, arg2 *StringBuilder, arg3 string, arg4 string, arg5 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkXmlDSigGen_AddExternalXmlRef(z.ckObj, cstr1, arg2.ckObj, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Specifies an Object to be added to the Signature.
//     The id is the value of the Object element's Id attribute.
//     The content contains the content of the Object element, which may be XML or
//     plain text.
//     The mimeType is the value of the Object element's MimeType attribute
//     The encoding is the value of the Object element's Encoding attribute
// In most cases, the mimeType and encoding are empty strings which cause the MimeType and
// Encoding attributes to be omitted.
func (z *XmlDSigGen) AddObject(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkXmlDSigGen_AddObject(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// This is the same as the AddSameDocRef method, except the reference is to content
// within an Object previously added via the AddObject method. The id must be an
// Id equal to the Id attribute of an Object, or the Id attribute of an element
// within the Object.
// 
// Note: The canonMethod can be set to "Base64" to use the
// http://www.w3.org/2000/09/xmldsig#base64 transform.
//
func (z *XmlDSigGen) AddObjectRef(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkXmlDSigGen_AddObjectRef(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Specifies a same document Reference to be added to the Signature when generated.
// A same document Reference can be the entire XML document, or a fragment of the
// XML document.
// 
// The id can be the empty string to sign the entire XML document, or it can be
// the fragment identifier to sign a portion of the XML document.
// 
// The digestMethod is the digest method and can be one of the following: "sha1", "sha256",
// "sha384", "sha512", "ripemd160", or "md5".
// 
// The canonMethod is the canonicalization method, and can be one of the following:
//     "C14N" -- for Inclusive Canonical XML (without comments)
//     "C14N_11" -- for Inclusive Canonical XML 1.1 (without comments)
//     "EXCL_C14N" -- for Exclusive Canonical XML (without comments)
//     "C14N_WithComments" -- for Inclusive Canonical XML (with comments)
//     "C14N_11_WithComments" -- for Inclusive Canonical XML 1.1 (with comments)
//     "EXCL_C14N_WithComments" -- for Exclusive Canonical XML (with comments)
//     "" -- An empty string indicates that no transformation should be included /
//     applied for this reference.
//     Note: The WithComments options are available in Chilkat v9.5.0.71 and later.
//     Note: The empty-string canonMethod is available in Chilkat v9.5.0.75 and
//     later.
// 
// If exclusive canonicalization is selected, then the prefixList can contain a space
// separated list of inclusive namespace prefixes. For inclusive canonicalization,
// this argument is ignored. In general, pass an empty string for this argument
// unless you have specific knowledge of namespace prefixes that need to be treated
// as inclusive when EXCL_C14N is used.
// 
// Starting in Chilkat v9.5.0.70, the prefixList can be set to the keyword "_EMPTY_" to
// force the generation of an empty PrefixList under the Transform. For example:
// 
// The refType is optional and is usually not needed. Set this to the empty string
// unless it is desired to add a Type attribute to the Reference that is advisory
// only.
//
func (z *XmlDSigGen) AddSameDocRef(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkXmlDSigGen_AddSameDocRef(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Can be called one or more times to add additional namespaces to the Signature
// element.
func (z *XmlDSigGen) AddSignatureNamespace(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXmlDSigGen_AddSignatureNamespace(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// This method will construct and return the canonicalized SignedInfo XML. The
// digests of each Reference are computed and included in the SignedInfo. This
// method is provided for certain special circumstances where one wants to get the
// exact canonicalized SignedInfo that would be signed using the private key.
// 
// Note: Properties such as SigLocation, SigningAlg, etc. and references must be
// set exactly as if an XML signature was to be actually generated because they
// determine the content of the SignedInfo.
// 
// Note, the sbXml is not signed by this method. It is not modified.
//
// return a string or nil if failed.
func (z *XmlDSigGen) ConstructSignedInfo(arg1 *StringBuilder) *string {

    retStr := C.CkXmlDSigGen_constructSignedInfo(z.ckObj, arg1.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Creates an XML Digital Signature. The application passes in the XML to be
// signed, and the signed XML is returned. If creating an enveloping signature
// where the Signature element is the root, then the inXml may be the empty string.
// 
//     Chilkat v9.5.0.76 or greater is required for XML signatures for
//     www.csioz.gov.pl
//
// return a string or nil if failed.
func (z *XmlDSigGen) CreateXmlDSig(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkXmlDSigGen_createXmlDSig(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Creates an XML Digital Signature. The application passes the XML to be signed in
// sbXml, and it is replaced with the signed XML if successful. (Thus, sbXml is both
// an input and output argument.) Note: If creating an enveloping signature where
// the Signature element is to be the root element, then the passed-in sbXml may be
// empty.
func (z *XmlDSigGen) CreateXmlDSigSb(arg1 *StringBuilder) bool {

    v := C.CkXmlDSigGen_CreateXmlDSigSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the HMAC key to be used if the Signature is to use an HMAC signing
// algorithm. The encoding specifies the encoding of key, and can be "hex", "base64",
// "ascii", or any of the binary encodings supported by Chilkat in the link below.
func (z *XmlDSigGen) SetHmacKey(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXmlDSigGen_SetHmacKey(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the private key to be used for creating the XML signature. The private key
// may be an RSA key, a DSA key, or an ECDSA key.
func (z *XmlDSigGen) SetPrivateKey(arg1 *PrivateKey) bool {

    v := C.CkXmlDSigGen_SetPrivateKey(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the "Id" attribute for a Reference.
func (z *XmlDSigGen) SetRefIdAttr(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXmlDSigGen_SetRefIdAttr(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Specifies the X.509 certificate to be used for the KeyInfo element when the
// KeyInfoType equals "X509Data". If usePrivateKey is true, then the private key will also
// be set using the certificate's private key. Thus, the SetPrivateKey method does
// not need to be called. If usePrivateKey is true, and the certificate does not have an
// associated private key available, then this method will return false.
// 
// Note: A certificate's private key is not stored within a certificate itself. If
// the certificate (cert) was obtained from a PFX, Java KeyStore, or other such
// source, which are containers for both certs and private keys, then Chilkat would
// have associated the cert with the private key when loading the PFX or JKS, and
// all is good. The same holds true if, on a Windows system, the certificate was
// obtained from a Windows-based registry certificate store where the private key
// was installed with the permission to export.
// 
// If, however, the certificate was loaded from a .cer file, or another type of
// file that contains only the certificate and not the private key, then the
// associated private key needs to be obtained by the application and provided by
// calling SetPrivateKey.
//
func (z *XmlDSigGen) SetX509Cert(arg1 *Cert, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkXmlDSigGen_SetX509Cert(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


