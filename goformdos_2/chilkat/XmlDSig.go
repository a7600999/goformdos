// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkXmlDSig.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewXmlDSig() *XmlDSig {
	return &XmlDSig{C.CkXmlDSig_Create()}
}

func (z *XmlDSig) DisposeXmlDSig() {
    if z != nil {
	C.CkXmlDSig_Dispose(z.ckObj)
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
func (z *XmlDSig) DebugLogFilePath() string {
    return C.GoString(C.CkXmlDSig_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *XmlDSig) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSig_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ExternalRefDirs
// May contain a set of directory paths specifying where referenced external files
// are located. Directory paths should be separated using a semicolon character.
// The default value of this property is the empty string which means no
// directories are automatically searched.
// 
// This property can be used if the external file referenced in the XML signature
// has the same filename as the file in the local filesystem.
//
func (z *XmlDSig) ExternalRefDirs() string {
    return C.GoString(C.CkXmlDSig_externalRefDirs(z.ckObj))
}
// property setter: ExternalRefDirs
func (z *XmlDSig) SetExternalRefDirs(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlDSig_putExternalRefDirs(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IgnoreExternalRefs
// If true, then ignore failures caused by external references not being
// available. This allows for the XML signature to be at least partially validated
// if the external referenced files are not available. The default value of this
// property is false.
func (z *XmlDSig) IgnoreExternalRefs() bool {
    v := int(C.CkXmlDSig_getIgnoreExternalRefs(z.ckObj))
    return v != 0
}
// property setter: IgnoreExternalRefs
func (z *XmlDSig) SetIgnoreExternalRefs(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmlDSig_putIgnoreExternalRefs(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *XmlDSig) LastErrorHtml() string {
    return C.GoString(C.CkXmlDSig_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *XmlDSig) LastErrorText() string {
    return C.GoString(C.CkXmlDSig_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *XmlDSig) LastErrorXml() string {
    return C.GoString(C.CkXmlDSig_lastErrorXml(z.ckObj))
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
func (z *XmlDSig) LastMethodSuccess() bool {
    v := int(C.CkXmlDSig_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *XmlDSig) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmlDSig_putLastMethodSuccess(z.ckObj,v)
}

// property: NumReferences
// The number of data objects referenced in the XML digital signature. A data
// object may be self-contained within the loaded XML signature, or it may be an
// external URI reference. An application can check each reference to see if it's
// external by calling IsReferenceExternal, and can get each reference URI by
// calling ReferenceUri.
func (z *XmlDSig) NumReferences() int {
    return int(C.CkXmlDSig_getNumReferences(z.ckObj))
}

// property: NumSignatures
// The number of digital signatures found within the loaded XML. Each digital
// signature is composed of XML having the following structure:
//   _LT_Signature ID?_GT_ 
//      _LT_SignedInfo_GT_
//        _LT_CanonicalizationMethod/_GT_
//        _LT_SignatureMethod/_GT_
//        (_LT_Reference URI? _GT_
//          (_LT_Transforms_GT_)?
//          _LT_DigestMethod_GT_
//          _LT_DigestValue_GT_
//        _LT_/Reference_GT_)+
//      _LT_/SignedInfo_GT_
//      _LT_SignatureValue_GT_ 
//     (_LT_KeyInfo_GT_)?
//     (_LT_Object ID?_GT_)*
//  _LT_/Signature_GT_
// Note: The "Signature" and other XML tags may be namespace prefixed.
// 
// The Selector property is used to select which XML signature is in effect when
// validating or calling other methods or properties.
//
func (z *XmlDSig) NumSignatures() int {
    return int(C.CkXmlDSig_getNumSignatures(z.ckObj))
}

// property: RefFailReason
// Indicates the failure reason for the last call to VerifyReferenceDigest.
// Possible values are:
//     0: No failure, the reference digest was valid.
//     1: The computed digest differs from the digest stored in the XML.
//     2: An external file is referenced, but it is unavailable for computing the
//     digest.
//     3: The index argument passed to VerifyReferenceDigest was out of range.
//     4: Unable to find the Signature.
//     5: A transformation specified some sort of XML canonicalization that is not
//     supported.
//     99: Unknown. (Should never get this value.)
func (z *XmlDSig) RefFailReason() int {
    return int(C.CkXmlDSig_getRefFailReason(z.ckObj))
}

// property: Selector
// If the loaded XML contains multiple signatures, this property can be set to
// specify which signature is in effect when calling other methods or properties.
// In most cases, the loaded XML contains a single signature and this property will
// remain at the default value of 0.
func (z *XmlDSig) Selector() int {
    return int(C.CkXmlDSig_getSelector(z.ckObj))
}
// property setter: Selector
func (z *XmlDSig) SetSelector(value int) {
    C.CkXmlDSig_putSelector(z.ckObj,C.int(value))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *XmlDSig) VerboseLogging() bool {
    v := int(C.CkXmlDSig_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *XmlDSig) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmlDSig_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *XmlDSig) Version() string {
    return C.GoString(C.CkXmlDSig_version(z.ckObj))
}
// Applies XML canonicalization to a fragment of the passed-in XML, and returns the
// canonicalized XML string. The fragmentId identifies the XML element where output
// begins. It is the XML element having an "id" attribute equal to fragmentId. The version
// may be one of the following:
//     "C14N" -- for Inclusive Canonical XML
//     "EXCL_C14N" -- for Exclusive Canonical XML
// 
// The prefixList only applies when the version is set to "EXCL_C14N". The prefixList may be an
// empty string, or a SPACE separated list of namespace prefixes. It is the
// InclusiveNamespaces PrefixList, which is the list of namespaces that are
// propagated from ancestor elements for canonicalization purposes.
// 
// If withComments is true, then XML comments are included in the output. If withComments is
// false, then XML comments are excluded from the output.
//
// return a string or nil if failed.
func (z *XmlDSig) CanonicalizeFragment(arg1 string, arg2 string, arg3 string, arg4 string, arg5 bool) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }

    retStr := C.CkXmlDSig_canonicalizeFragment(z.ckObj, cstr1, cstr2, cstr3, cstr4, b5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Applies XML canonicalization to the passed-in XML, and returns the canonicalized
// XML string. The version may be one of the following:
//     "C14N" -- for Inclusive Canonical XML
//     "EXCL_C14N" -- for Exclusive Canonical XML
// 
// If withComments is true, then XML comments are included in the output. If withComments is
// false, then XML comments are excluded from the output.
//
// return a string or nil if failed.
func (z *XmlDSig) CanonicalizeXml(arg1 string, arg2 string, arg3 bool) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retStr := C.CkXmlDSig_canonicalizeXml(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the certificates found in the signature indicated by the Selector
// property. The base64 representation of each certificate is returned.
func (z *XmlDSig) GetCerts(arg1 *StringArray) bool {

    v := C.CkXmlDSig_GetCerts(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the KeyInfo XML for the signature indicated by the Selector property.
// Returns _NULL_ if no KeyInfo exists.
func (z *XmlDSig) GetKeyInfo() *Xml {

    retObj := C.CkXmlDSig_GetKeyInfo(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Returns the public key from the KeyInfo XML for the signature indicated by the
// Selector property. Returns _NULL_ if no KeyInfo exists, or if no public key is
// actually contained in the KeyInfo.
func (z *XmlDSig) GetPublicKey() *PublicKey {

    retObj := C.CkXmlDSig_GetPublicKey(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &PublicKey{retObj}
}


// Returns true if the reference at index is external. Each external reference
// would first need to be provided by the application prior to validating the
// signature.
func (z *XmlDSig) IsReferenceExternal(arg1 int) bool {

    v := C.CkXmlDSig_IsReferenceExternal(z.ckObj, C.int(arg1))


    return v != 0
}


// Loads an XML document containing 1 or more XML digital signatures. An
// application would first load XML containing digital signature(s), then validate.
// After loading, it is also possible to use various methods and properties to get
// information about the signature, such as the key info, references, etc. If
// external data is referenced by the signature, it may be necessary to provide the
// referenced data prior to validating.
// 
// Note: When loading an XML document, the Selector property is automatically reset
// to 0, and the NumSignatures property is set to the number of XML digital
// signatures found.
//
func (z *XmlDSig) LoadSignature(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXmlDSig_LoadSignature(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads an XML document containing one or more XML digital signatures from the
// contents of a BinData object. An application would first load the XML, then
// validate. After loading, it is also possible to use various methods and
// properties to get information about the signature, such as the key info,
// references, etc. If external data is referenced by the signature, it may be
// necessary to provide the referenced data prior to validating.
// 
// Note: When loading an XML document, the Selector property is automatically reset
// to 0, and the NumSignatures property is set to the number of XML digital
// signatures found.
//
func (z *XmlDSig) LoadSignatureBd(arg1 *BinData) bool {

    v := C.CkXmlDSig_LoadSignatureBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads an XML document containing one or more XML digital signatures from the
// contents of a StringBuilder object. An application would first load the XML,
// then validate. After loading, it is also possible to use various methods and
// properties to get information about the signature, such as the key info,
// references, etc. If external data is referenced by the signature, it may be
// necessary to provide the referenced data prior to validating.
// 
// Note: When loading an XML document, the Selector property is automatically reset
// to 0, and the NumSignatures property is set to the number of XML digital
// signatures found.
//
func (z *XmlDSig) LoadSignatureSb(arg1 *StringBuilder) bool {

    v := C.CkXmlDSig_LoadSignatureSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the URI of the Nth reference specified by index. (The 1st reference is at
// index 0.) URI's beginning with a pound sign ('#') are internal references.
// return a string or nil if failed.
func (z *XmlDSig) ReferenceUri(arg1 int) *string {

    retStr := C.CkXmlDSig_referenceUri(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sets the HMAC key to be used if the Signature used an HMAC signing algorithm.
// The encoding specifies the encoding of key, and can be "hex", "base64", "ascii", or
// any of the binary encodings supported by Chilkat in the link below.
func (z *XmlDSig) SetHmacKey(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXmlDSig_SetHmacKey(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the public key to be used for verifying the signature indicated by the
// Selector property. A public key only needs to be explicitly provided by the
// application for those XML signatures where the public key is not already
// available within the KeyInfo of the Signature. In some cases, the KeyInfo within
// the Signature contains information about what public key was used for signing.
// The application can use this information to retrieve the matching public key and
// provide it via this method.
func (z *XmlDSig) SetPublicKey(arg1 *PublicKey) bool {

    v := C.CkXmlDSig_SetPublicKey(z.ckObj, arg1.ckObj)


    return v != 0
}


// Provides the binary data for the external reference at index.
func (z *XmlDSig) SetRefDataBd(arg1 int, arg2 *BinData) bool {

    v := C.CkXmlDSig_SetRefDataBd(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Provides the data for the external reference at index. When validating the
// signature, the data contained in path will be streamed to compute the digest for
// validation.
func (z *XmlDSig) SetRefDataFile(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkXmlDSig_SetRefDataFile(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Provides the text data for the external reference at index. The charset specifies
// the byte representation of the text, such as "utf-8", "utf-16", "windows-1252",
// etc. (If in doubt, try utf-8 first.)
func (z *XmlDSig) SetRefDataSb(arg1 int, arg2 *StringBuilder, arg3 string) bool {
    cstr3 := C.CString(arg3)

    v := C.CkXmlDSig_SetRefDataSb(z.ckObj, C.int(arg1), arg2.ckObj, cstr3)

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds an XML certificate vault to the object's internal list of sources to be
// searched for certificates having public keys when verifying an XML signature. A
// single XML certificate vault may be used. If UseCertVault is called multiple
// times, only the last certificate vault will be used, as each call to
// UseCertVault will replace the certificate vault provided in previous calls.
func (z *XmlDSig) UseCertVault(arg1 *XmlCertVault) bool {

    v := C.CkXmlDSig_UseCertVault(z.ckObj, arg1.ckObj)


    return v != 0
}


// This method allows for an application to verify the digest for each reference
// separately. This can be helpful if the full XMLDSIG validation fails, then one
// can test each referenced data's digest to see which, if any, fail to match.
func (z *XmlDSig) VerifyReferenceDigest(arg1 int) bool {

    v := C.CkXmlDSig_VerifyReferenceDigest(z.ckObj, C.int(arg1))


    return v != 0
}


// Verifies the signature indicated by the Selector property. If verifyReferenceDigests is true,
// then the digest of each Reference within the signature's SignedInfo is also
// validated.
func (z *XmlDSig) VerifySignature(arg1 bool) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkXmlDSig_VerifySignature(z.ckObj, b1)


    return v != 0
}


