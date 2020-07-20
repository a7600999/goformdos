// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkMime.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewMime() *Mime {
	return &Mime{C.CkMime_Create()}
}

func (z *Mime) DisposeMime() {
    if z != nil {
	C.CkMime_Dispose(z.ckObj)
	}
}




// property: Boundary
// The boundary string for a multipart MIME message.
// 
// It is the value of the boundary attribute of the Content-Type header field. For
// example, if the Content-Type header is this:
// Content-Type: multipart/mixed; boundary="------------080707010302060306060800"
// then the value of the Boundary property is
// "------------080707010302060306060800".
// 
// When building multipart MIME messages, the boundary is automatically generated
// by methods such as NewMultipartMixed, to be a unique and random string, so
// explicitly setting the boundary is usually not necessary.
//
func (z *Mime) Boundary() string {
    return C.GoString(C.CkMime_boundary(z.ckObj))
}
// property setter: Boundary
func (z *Mime) SetBoundary(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putBoundary(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Charset
// The value of the "charset" attribute of the Content-Type header field. For
// example, if the Content-Type header is this:
// Content-Type: text/plain; charset="iso-8859-1"
// then the value of the Charset property is "iso-8859-1".
func (z *Mime) Charset() string {
    return C.GoString(C.CkMime_charset(z.ckObj))
}
// property setter: Charset
func (z *Mime) SetCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CmsOptions
// A JSON string for controlling extra CMS (PKCS7) signature and validation
// options.
func (z *Mime) CmsOptions() string {
    return C.GoString(C.CkMime_cmsOptions(z.ckObj))
}
// property setter: CmsOptions
func (z *Mime) SetCmsOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putCmsOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ContentType
// The MIME content type, such as "text/plain", "text/html", "image/gif",
// "multipart/alternative", "multipart/mixed", etc.
// 
// It is the value of the Content-Type header field, excluding any attributes. For
// example, if the Content-Type header is this:
// Content-Type: multipart/mixed; boundary="------------080707010302060306060800"
// then the value of the ContentType property is "multipart/mixed".
//
func (z *Mime) ContentType() string {
    return C.GoString(C.CkMime_contentType(z.ckObj))
}
// property setter: ContentType
func (z *Mime) SetContentType(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putContentType(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CurrentDateTime
// Returns the current date/time in RFC 822 format.
func (z *Mime) CurrentDateTime() string {
    return C.GoString(C.CkMime_currentDateTime(z.ckObj))
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
func (z *Mime) DebugLogFilePath() string {
    return C.GoString(C.CkMime_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Mime) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Disposition
// The value of the Content-Disposition header field, excluding any attributes. For
// example, if the Content-Disposition header is this:
// Content-Disposition: attachment; filename="starfish.gif"
// then the value of the Disposition property is "attachment".
func (z *Mime) Disposition() string {
    return C.GoString(C.CkMime_disposition(z.ckObj))
}
// property setter: Disposition
func (z *Mime) SetDisposition(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putDisposition(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Encoding
// The value of the Content-Transfer-Encoding header field. Typical values are
// "base64", "quoted-printable", "7bit", "8bit", "binary", etc. For example, if the
// Content-Transfer-Encoding header is this:
// Content-Transfer-Encoding: base64
// then the value of the Encoding property is "base64".
func (z *Mime) Encoding() string {
    return C.GoString(C.CkMime_encoding(z.ckObj))
}
// property setter: Encoding
func (z *Mime) SetEncoding(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putEncoding(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Filename
// The value of the "filename" attribute of the Content-Disposition header field.
// For example, if the Content-Disposition header is this:
// Content-Disposition: attachment; filename="starfish.gif"
// then the value of the Filename property is "starfish.gif".
func (z *Mime) Filename() string {
    return C.GoString(C.CkMime_filename(z.ckObj))
}
// property setter: Filename
func (z *Mime) SetFilename(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putFilename(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Mime) LastErrorHtml() string {
    return C.GoString(C.CkMime_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Mime) LastErrorText() string {
    return C.GoString(C.CkMime_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Mime) LastErrorXml() string {
    return C.GoString(C.CkMime_lastErrorXml(z.ckObj))
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
func (z *Mime) LastMethodSuccess() bool {
    v := int(C.CkMime_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Mime) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMime_putLastMethodSuccess(z.ckObj,v)
}

// property: Micalg
// The value of the "micalg" attribute of the Content-Type header field. For
// example, if the Content-Type header is this:
// Content-Type: multipart/signed; protocol="application/x-pkcs7-signature"; micalg=sha1; 
//   boundary="------------ms000908010507020408060303"
// then the value of the Micalg property is "sha".
// 
// Note: The micalg attribute is only present in PKCS7 signed MIME. Setting the
// Micalg property has the effect of choosing the hash algorithm used w/ signing.
// Possible choices are "sha1", "md5", "sha256", "sha384", and "sha512". However,
// it is preferable to set the signing hash algorithm by setting the SigningHashAlg
// property instead.
//
func (z *Mime) Micalg() string {
    return C.GoString(C.CkMime_micalg(z.ckObj))
}
// property setter: Micalg
func (z *Mime) SetMicalg(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putMicalg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Name
// The value of the "name" attribute of the Content-Type header field. For example,
// if the Content-Type header is this:
// Content-Type: image/gif; name="starfish.gif"
// then the value of the Name property is "starfish.gif".
func (z *Mime) Name() string {
    return C.GoString(C.CkMime_name(z.ckObj))
}
// property setter: Name
func (z *Mime) SetName(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: NumEncryptCerts
// The number of certificates found when decrypting S/MIME. This property is set
// after UnwrapSecurity is called.
func (z *Mime) NumEncryptCerts() int {
    return int(C.CkMime_getNumEncryptCerts(z.ckObj))
}

// property: NumHeaderFields
// The number of header fields. Header field names and values can be retrieved by
// index (starting at 0) by calling GetHeaderFieldName and GetHeaderFieldValue.
func (z *Mime) NumHeaderFields() int {
    return int(C.CkMime_getNumHeaderFields(z.ckObj))
}

// property: NumParts
// MIME messages are composed of parts in a tree structure. The NumParts property
// contains the number of direct children. To traverse an entire MIME tree, one
// would recursively descend the tree structure by iterating from 0 to NumParts-1,
// calling GetPart to get each direct child MIME object. The traversal would
// continue by iterating over each child's parts, and so on.
func (z *Mime) NumParts() int {
    return int(C.CkMime_getNumParts(z.ckObj))
}

// property: NumSignerCerts
// The number of certificates found when verifying signature(s). This property is
// set after UnwrapSecurity is called.
func (z *Mime) NumSignerCerts() int {
    return int(C.CkMime_getNumSignerCerts(z.ckObj))
}

// property: OaepHash
// Selects the hash algorithm for use within OAEP padding when encrypting MIME
// using RSAES-OAEP. The valid choices are "sha1", "sha256", "sha384", "sha512",
func (z *Mime) OaepHash() string {
    return C.GoString(C.CkMime_oaepHash(z.ckObj))
}
// property setter: OaepHash
func (z *Mime) SetOaepHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putOaepHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OaepMgfHash
// Selects the MGF hash algorithm for use within OAEP padding when encrypting MIME
// using RSAES-OAEP. The valid choices are "sha1", "sha256", "sha384", "sha512",
// The default is "sha1".
func (z *Mime) OaepMgfHash() string {
    return C.GoString(C.CkMime_oaepMgfHash(z.ckObj))
}
// property setter: OaepMgfHash
func (z *Mime) SetOaepMgfHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putOaepMgfHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OaepPadding
// Selects the RSA encryption scheme when encrypting MIME. The default value is
// false, which selects RSAES_PKCS1-V1_5. If set to true, then RSAES_OAEP is
// used.
func (z *Mime) OaepPadding() bool {
    v := int(C.CkMime_getOaepPadding(z.ckObj))
    return v != 0
}
// property setter: OaepPadding
func (z *Mime) SetOaepPadding(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMime_putOaepPadding(z.ckObj,v)
}

// property: Pkcs7CryptAlg
// When the MIME is encrypted (using PKCS7 public-key encryption), this selects the
// underlying symmetric encryption algorithm. Possible values are: "aes", "des",
// "3des", and "rc2". The default value is "aes".
func (z *Mime) Pkcs7CryptAlg() string {
    return C.GoString(C.CkMime_pkcs7CryptAlg(z.ckObj))
}
// property setter: Pkcs7CryptAlg
func (z *Mime) SetPkcs7CryptAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putPkcs7CryptAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Pkcs7KeyLength
// When the MIME is encrypted (using PKCS7 public-key encryption), this selects the
// key length of the underlying symmetric encryption algorithm. The possible values
// allowed depend on the Pkcs7CryptAlg property. For "aes", the key length may be
// 128, 192, or 256. For "3des" the key length must be 192. For "des" the key
// length must be 40. For "rc2" the key length can be 40, 56, 64, or 128.
func (z *Mime) Pkcs7KeyLength() int {
    return int(C.CkMime_getPkcs7KeyLength(z.ckObj))
}
// property setter: Pkcs7KeyLength
func (z *Mime) SetPkcs7KeyLength(value int) {
    C.CkMime_putPkcs7KeyLength(z.ckObj,C.int(value))
}

// property: Protocol
// The value of the "protocol" attribute of the Content-Type header field. For
// example, if the Content-Type header is this:
// Content-Type: multipart/signed; protocol="application/x-pkcs7-signature"; micalg=sha1; 
//   boundary="------------ms000908010507020408060303"
// then the value of the Protocol property is "application/x-pkcs7-signature".
func (z *Mime) Protocol() string {
    return C.GoString(C.CkMime_protocol(z.ckObj))
}
// property setter: Protocol
func (z *Mime) SetProtocol(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putProtocol(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SigningAlg
// Selects the signature algorithm to be used when creating signed (PKCS7) MIME.
// The default value is "PKCS1-v1_5". This can be set to "RSASSA-PSS" (or simply
// "pss") to use the RSASSA-PSS signature scheme.
// 
// Note: This property only applies when signing with an RSA private key. It does
// not apply for ECC or DSA private keys.
//
func (z *Mime) SigningAlg() string {
    return C.GoString(C.CkMime_signingAlg(z.ckObj))
}
// property setter: SigningAlg
func (z *Mime) SetSigningAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putSigningAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SigningHashAlg
// Selects the underlying hash algorithm used when creating signed (PKCS7) MIME.
// Possible values are "sha1", "sha256", "sha384", "sha512", "md5", and "md2".
func (z *Mime) SigningHashAlg() string {
    return C.GoString(C.CkMime_signingHashAlg(z.ckObj))
}
// property setter: SigningHashAlg
func (z *Mime) SetSigningHashAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkMime_putSigningHashAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UnwrapExtras
// Controls whether extra (informative) header fields are added to the MIME message
// when unwrapping security.
func (z *Mime) UnwrapExtras() bool {
    v := int(C.CkMime_getUnwrapExtras(z.ckObj))
    return v != 0
}
// property setter: UnwrapExtras
func (z *Mime) SetUnwrapExtras(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMime_putUnwrapExtras(z.ckObj,v)
}

// property: UseMmDescription
// Controls whether the boilerplate text "This is a multi-part message in MIME
// format." is used as the body content of a multipart MIME part.
func (z *Mime) UseMmDescription() bool {
    v := int(C.CkMime_getUseMmDescription(z.ckObj))
    return v != 0
}
// property setter: UseMmDescription
func (z *Mime) SetUseMmDescription(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMime_putUseMmDescription(z.ckObj,v)
}

// property: UseXPkcs7
// If true, then the Content-Type header fields created by Chilkat will use
// "x-pkcs7" instead of simply "pkcs7" . For example:
// Content-Type: multipart/signed;
// 	boundary="----=_NextPart_af8_0422_dbec3a60.7178e470";
// 	protocol="application/x-pkcs7-signature"; micalg=sha1
// 
// or
// 
// Content-Type: application/x-pkcs7-mime; name="smime.p7m"
// If false, then the "pcks7" is used. For example:
// Content-Type: multipart/signed;
// 	boundary="----=_NextPart_af8_0422_dbec3a60.7178e470";
// 	protocol="application/pkcs7-signature"; micalg=sha1
// 
// or
// 
// Content-Type: application/pkcs7-mime; name="smime.p7m"
// The default value of this property is true, meaning that "x-" is used by
// default.
func (z *Mime) UseXPkcs7() bool {
    v := int(C.CkMime_getUseXPkcs7(z.ckObj))
    return v != 0
}
// property setter: UseXPkcs7
func (z *Mime) SetUseXPkcs7(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMime_putUseXPkcs7(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Mime) VerboseLogging() bool {
    v := int(C.CkMime_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Mime) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkMime_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Mime) Version() string {
    return C.GoString(C.CkMime_version(z.ckObj))
}
// Computes the size of the MIME body and adds a Content-Length header field with
// the computed value. If the MIME body is non-multipart, the Content-Length is
// just the size of the content. If the MIME is multipart, then the Content-Length
// is the sum of all the sub-parts. Calling this method more than once causes the
// Content-Length header to be re-computed and updated.
func (z *Mime) AddContentLength()  {

    C.CkMime_AddContentLength(z.ckObj)



}


// Makes a certificate available for decrypting if needed by methods that decrypt,
// such as UnwrapSecurity. This method may be called multiple times to make more
// than one certificate (and it's private key) available. Alternative methods for
// making certificates available are UseCertVault, AddPfxSourceFile, and
// AddPfxSourceData.
func (z *Mime) AddDecryptCert(arg1 *Cert) bool {

    v := C.CkMime_AddDecryptCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Signs the message using the certificate provided. If successful, the message is
// converted to "multipart/signed" and the original message will be contained in
// the first sub-part.
func (z *Mime) AddDetachedSignature(arg1 *Cert) bool {

    v := C.CkMime_AddDetachedSignature(z.ckObj, arg1.ckObj)


    return v != 0
}


// Same as AddDetachedSignature, except an extra argument is provided to control
// whether header fields from the calling MIME object are transferred to the
// content part of the multipart/signed object. This method transforms the calling
// object into a multipart/signed MIME with two sub-parts. The first contains the
// original content of the calling object, and the second contains the digital
// signature.
func (z *Mime) AddDetachedSignature2(arg1 *Cert, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkMime_AddDetachedSignature2(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Adds a detached signature using a certificate and it's associated private key.
// This method would be used when the private key is external to the certificate --
// for example, if a PFX/P12 file is not used, but instead a pair of .cer and .pem
// files are used (one for the certificate and one for the associated private key).
func (z *Mime) AddDetachedSignaturePk(arg1 *Cert, arg2 *PrivateKey) bool {

    v := C.CkMime_AddDetachedSignaturePk(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Same as AddDetachedSignaturePk, except an extra argument is provided to control
// whether header fields from the calling MIME object are transferred to the
// content part of the multipart/signed object. This method transforms the calling
// object into a multipart/signed MIME with two sub-parts. The first contains the
// original content of the calling object, and the second contains the digital
// signature.
func (z *Mime) AddDetachedSignaturePk2(arg1 *Cert, arg2 *PrivateKey, arg3 bool) bool {
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkMime_AddDetachedSignaturePk2(z.ckObj, arg1.ckObj, arg2.ckObj, b3)


    return v != 0
}


// Adds a certificate to the object's internal list of certificates to be used when
// the EncryptN method is called. (See the EncryptN method for more information.)
// The internal list may be cleared by calling ClearEncryptCerts.
func (z *Mime) AddEncryptCert(arg1 *Cert) bool {

    v := C.CkMime_AddEncryptCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds a header field to the MIME.
func (z *Mime) AddHeaderField(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMime_AddHeaderField(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a PFX to the object's internal list of sources to be searched for
// certificates and private keys when decrypting . Multiple PFX sources can be
// added by calling this method once for each. (On the Windows operating system,
// the registry-based certificate stores are also automatically searched, so it is
// commonly not required to explicitly add PFX sources.)
// 
// The pfxFileData contains the bytes of a PFX file (also known as PKCS12 or .p12).
//
func (z *Mime) AddPfxSourceData(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkMime_AddPfxSourceData(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a PFX file to the object's internal list of sources to be searched for
// certificates and private keys when decrypting. Multiple PFX files can be added
// by calling this method once for each. (On the Windows operating system, the
// registry-based certificate stores are also automatically searched, so it is
// commonly not required to explicitly add PFX sources.)
// 
// The pfxFilePath contains the bytes of a PFX file (also known as PKCS12 or .p12).
//
func (z *Mime) AddPfxSourceFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMime_AddPfxSourceFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends a MIME message to the sub-parts of this message. Arbitrarily complex
// messages with unlimited nesting levels can be created. If the calling Mime
// object is not already multipart, it is automatically converted to
// multipart/mixed first.
func (z *Mime) AppendPart(arg1 *Mime) bool {

    v := C.CkMime_AppendPart(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads a file and creates a Mime message object using the file extension to
// determine the content type, and adds it as a sub-part to the calling object.
func (z *Mime) AppendPartFromFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_AppendPartFromFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// When the body of a MIME part contains PKCS7 (ASN.1 in DER format,
// base64-encoded), this method can be used to convert the ASN.1 to an XML format
// for inspection. Here is an example of how an ASN.1 body might look:
// Content-Type: application/x-pkcs7-mime;
// 	name="smime.p7m"; smime-type="signed-data"
// Content-Transfer-Encoding: base64
// Content-Disposition: attachment; filename="smime.p7m"
// 
// MIIXXAYJKoZIhvcNAQcCoIIXTTCCF0kCAQExCzAJBgUrDgMCGgUAMFoGCSqGSIb3DQEHAaBNBEtD
// b250ZW50LVR5cGU6IHRleHQvcGxhaW4NCkNvbnRlbnQtVHJhbnNmZXItRW5jb2Rpbmc6IDdiaXQN
// Cg0KdGhpcyBpcyBhIHRlc3SgghI/MIIE3jCCA8agAwIBAgICAwEwDQYJKoZIhvcNAQEFBQAwYzEL
// ...
// The XML produced would look something like this:
// _LT_?xml version="1.0" encoding="utf-8" ?>
// _LT_sequence>
//     _LT_oid>1.2.840.113549.1.7.2_LT_/oid>
//     _LT_contextSpecific tag="0" constructed="1">
//         _LT_sequence>
//             _LT_int>01_LT_/int>
//             _LT_set>
//                 _LT_sequence>
//                     _LT_oid>1.3.14.3.2.26_LT_/oid>
//                     _LT_null />
//                 _LT_/sequence>
//             _LT_/set>
//             _LT_sequence>
//                 _LT_oid>1.2.840.113549.1.7.1_LT_/oid>
//                 _LT_contextSpecific tag="0" constructed="1">
// ...
// return a string or nil if failed.
func (z *Mime) AsnBodyToXml() *string {

    retStr := C.CkMime_asnBodyToXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Clears the internal list of certificates added by previous calls to the
// AddEncryptCert method. (See the EncryptN method for information about encrypting
// using multiple certificates.)
func (z *Mime) ClearEncryptCerts()  {

    C.CkMime_ClearEncryptCerts(z.ckObj)



}


// Returns true if the MIME message contains encrypted parts.
// 
// Note: This method examines the MIME as-is. If UnwrapSecurity is called and it is
// successful, then the MIME should no longer contain encrypted parts, and this
// method would return 0.
// 
// Note: If a signed MIME message is then encrypted, then it is not possible to
// know that the MIME is both encrypted and signed until UnwrapSecurity is called.
// (In other words, it is not possible to know the contents of the encrypted MIME
// until it is decrypted.) Therefore, the ContainsSignedParts method would return
// false.
//
func (z *Mime) ContainsEncryptedParts() bool {

    v := C.CkMime_ContainsEncryptedParts(z.ckObj)


    return v != 0
}


// Returns true if the MIME message contains signed parts.
// 
// Note: This method examines the MIME as-is. If UnwrapSecurity is called and it is
// successful, then the MIME should no longer contain signed parts, and this method
// would return 0.
// 
// Note: If a signed MIME message is then encrypted, then it is not possible to
// know that the MIME is both encrypted and signed until UnwrapSecurity is called.
// (In other words, it is not possible to know the contents of the encrypted MIME
// until it is decrypted.) Therefore, the ContainsSignedParts method would return
// false.
// 
// Note: The same concept also applies to opaque signatures, such as with the MIME
// produced by calling ConvertToSigned.
//
func (z *Mime) ContainsSignedParts() bool {

    v := C.CkMime_ContainsSignedParts(z.ckObj)


    return v != 0
}


// Changes the content-transfer-encoding to "base64" for all 8bit or binary MIME
// subparts. This allows for the MIME to be exported as a string via the GetMime
// method.
func (z *Mime) Convert8Bit()  {

    C.CkMime_Convert8Bit(z.ckObj)



}


// Converts existing MIME to a multipart/alternative. This is accomplished by
// creating a new outermost multipart/alternative MIME part. The existing MIME is
// moved into the 1st (and only) sub-part of the new multipart/alternative
// enclosure. Header fields from the original top-level MIME part are transferred
// to the new top-level multipart/alternative header, except for Content-Type,
// Content-Transfer-Encoding, and Content-Disposition. For example, the following
// simple plain-text MIME is converted as follows:
// 
// Original:
// MIME-Version: 1.0
// Date: Sun, 11 Aug 2013 11:18:44 -0500
// Message-ID: Content-Type: text/plain
// Content-Transfer-Encoding: quoted-printable
// X-Priority: 3 (Normal)
// Subject: this is the subject.
// From: "Chilkat Software" 
// To: "Chilkat Sales" This is the plain-text body.
// 
// After Converting:
// MIME-Version: 1.0
// Date: Sun, 11 Aug 2013 11:18:44 -0500
// Message-ID: X-Priority: 3 (Normal)
// Subject: this is the subject.
// From: "Chilkat Software" 
// To: "Chilkat Sales" Content-Type: multipart/alternative;
// 	boundary="------------040101040804050401050400_.ALT"
// 
// --------------040101040804050401050400_.ALT
// Content-Type: text/plain
// Content-Transfer-Encoding: quoted-printable
// 
// This is the plain-text body.
// --------------040101040804050401050400_.ALT--
//
func (z *Mime) ConvertToMultipartAlt() bool {

    v := C.CkMime_ConvertToMultipartAlt(z.ckObj)


    return v != 0
}


// Converts existing MIME to a multipart/mixed. This is accomplished by creating a
// new outermost multipart/mixed MIME part. The existing MIME is moved into the 1st
// (and only) sub-part of the new multipart/mixed enclosure. Header fields from the
// original top-level MIME part are transferred to the new top-level
// multipart/mixed header, except for Content-Type, Content-Transfer-Encoding, and
// Content-Disposition. For example, the following simple plain-text MIME is
// converted as follows:
// 
// Original:
// MIME-Version: 1.0
// Date: Sun, 11 Aug 2013 11:27:04 -0500
// Message-ID: Content-Type: text/plain
// Content-Transfer-Encoding: quoted-printable
// X-Priority: 3 (Normal)
// Subject: this is the subject.
// From: "Chilkat Software" 
// To: "Chilkat Sales" This is the plain-text body.
// 
// After Converting:
// MIME-Version: 1.0
// Date: Sun, 11 Aug 2013 11:27:04 -0500
// Message-ID: X-Priority: 3 (Normal)
// Subject: this is the subject.
// From: "Chilkat Software" 
// To: "Chilkat Sales" Content-Type: multipart/mixed;
// 	boundary="------------050508060709030908040207"
// 
// --------------050508060709030908040207
// Content-Type: text/plain
// Content-Transfer-Encoding: quoted-printable
// 
// This is the plain-text body.
// --------------050508060709030908040207--
//
func (z *Mime) ConvertToMultipartMixed() bool {

    v := C.CkMime_ConvertToMultipartMixed(z.ckObj)


    return v != 0
}


// Digitally signs a MIME message. The MIME is converted to an
// application/x-pkcs7-mime which is a PKCS7 signature that includes both the
// original MIME message and the signature. This is different than
// AddDetachedSignature, where the signature is appended to the MIME.
// 
// Note: This is commonly referred to as an "opaque" signature.
//
func (z *Mime) ConvertToSigned(arg1 *Cert) bool {

    v := C.CkMime_ConvertToSigned(z.ckObj, arg1.ckObj)


    return v != 0
}


// Digitally signs the MIME to convert it to an "opaque" signed message using a
// certificate and it's associated private key. This method would be used when the
// private key is external to the certificate -- for example, if a PFX/P12 file is
// not used, but instead a pair of .cer and .pem files are used (one for the
// certificate and one for the associated private key).
func (z *Mime) ConvertToSignedPk(arg1 *Cert, arg2 *PrivateKey) bool {

    v := C.CkMime_ConvertToSignedPk(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Decrypts PKCS7 encrypted MIME (also known as S/MIME). Information about the
// certificates required for decryption is always embedded within PKCS7 encrypted
// MIME. This method will automatically find and use the certificate + private key
// required from three possible sources:
//     PFX files that were provided in one or more calls to AddPfxSourceData or
//     AddPfxSourceFile.
//     Certificates found in an XML certificate vault provided by calling the
//     UseCertVault method.
//     (On Windows systems) Certificates found in the system's registry-based
//     certificate stores.
func (z *Mime) Decrypt() bool {

    v := C.CkMime_Decrypt(z.ckObj)


    return v != 0
}


// The same as Decrypt, but useful when the certificate and private key are
// available in separate files (as opposed to a single file such as a .pfx/.p12).
func (z *Mime) Decrypt2(arg1 *Cert, arg2 *PrivateKey) bool {

    v := C.CkMime_Decrypt2(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Decrypts PKCS7 encrypted MIME (also known as S/MIME) using a specific
// certificate.
func (z *Mime) DecryptUsingCert(arg1 *Cert) bool {

    v := C.CkMime_DecryptUsingCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Decrypts MIME using a specific PFX ( also known as PKCS12, which is a file
// format commonly used to store private keys with accompanying public key
// certificates, protected with a password-based symmetric key). This method allows
// the bytes of the PKCS12 file to be passed directly, thus allowing PKCS12's to be
// persisted and retrieved from non-file-based locations, such as in LDAP or a
// database.
func (z *Mime) DecryptUsingPfxData(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkMime_DecryptUsingPfxData(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Decrypts MIME using a specific PFX file (also known as PKCS12) as the source for
// any required certificates and private keys. (Note: .pfx and .p12 files are both
// PKCS12 format.)
func (z *Mime) DecryptUsingPfxFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMime_DecryptUsingPfxFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Encrypts the MIME to create PKCS7 encrypted MIME. A digital certificate (which
// always contains a public-key) is used to encrypt.
func (z *Mime) Encrypt(arg1 *Cert) bool {

    v := C.CkMime_Encrypt(z.ckObj, arg1.ckObj)


    return v != 0
}


// Encrypt MIME using any number of digital certificates. Each certificate to be
// used must first be added by calling AddEncryptCert (once per certificate). See
// the example code below:
func (z *Mime) EncryptN() bool {

    v := C.CkMime_EncryptN(z.ckObj)


    return v != 0
}


// Recursively descends through the parts of a MIME message and extracts all parts
// having a filename to a file. The files are created in dirPath. Returns a
// (Ck)StringArray object containing the names of the files created. The filenames
// are obtained from the "filename" attribute of the content-disposition header. If
// a filename does not exist, then the MIME part is not saved to a file.
func (z *Mime) ExtractPartsToFiles(arg1 string) *StringArray {
    cstr1 := C.CString(arg1)

    retObj := C.CkMime_ExtractPartsToFiles(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Finds and returns the issuer certificate. If the certificate is a root or
// self-issued, then the certificate returned is a copy of the caller certificate.
func (z *Mime) FindIssuer(arg1 *Cert) *Cert {

    retObj := C.CkMime_FindIssuer(z.ckObj, arg1.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the body of the MIME message in a BinData object.
func (z *Mime) GetBodyBd(arg1 *BinData) bool {

    v := C.CkMime_GetBodyBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the body of the MIME message as a block of binary data. The body is
// automatically converted from its encoding type, such as base64 or
// quoted-printable, before being returned.
func (z *Mime) GetBodyBinary() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkMime_GetBodyBinary(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the body of the MIME message as a string. The body is automatically
// converted from its encoding type, such as base64 or quoted-printable, before
// being returned.
// return a string or nil if failed.
func (z *Mime) GetBodyDecoded() *string {

    retStr := C.CkMime_getBodyDecoded(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the body of the MIME message as a String. The body is explicitly not
// decoded from it's encoding type, so if it was represented in Base64, you will
// get the Base64 encoded body, as an example.
// return a string or nil if failed.
func (z *Mime) GetBodyEncoded() *string {

    retStr := C.CkMime_getBodyEncoded(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth certificate found when decrypting. The EncryptCerts property
// contains the number of certificates.
func (z *Mime) GetEncryptCert(arg1 int) *Cert {

    retObj := C.CkMime_GetEncryptCert(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the entire MIME body, including all sub-parts.
// return a string or nil if failed.
func (z *Mime) GetEntireBody() *string {

    retStr := C.CkMime_getEntireBody(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the MIME header.
// return a string or nil if failed.
func (z *Mime) GetEntireHead() *string {

    retStr := C.CkMime_getEntireHead(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of a MIME header field. fieldName is case-insensitive.
// return a string or nil if failed.
func (z *Mime) GetHeaderField(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkMime_getHeaderField(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Parses a MIME header field and returns the value of an attribute. MIME header
// fields w/ attributes are formatted like this:
// Header-Name:  value;  attrName1="value1"; attrName2="value2"; ....  attrNameN="valueN"
// Semi-colons separate attribute name=value pairs. The Content-Type header field
// often contains attributes. Here is an example:
// Content-Type: multipart/signed;
// 	protocol="application/x-pkcs7-signature";
// 	micalg=SHA1;
// 	boundary="----=_NextPart_000_0000_01CB03E4.D0BAF010"
// In the above example, to access the value of the "protocol" attribute, call
// GetHeaderFieldAttribute("Content-Type", "protocol");
// return a string or nil if failed.
func (z *Mime) GetHeaderFieldAttribute(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkMime_getHeaderFieldAttribute(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth MIME header field name.
// return a string or nil if failed.
func (z *Mime) GetHeaderFieldName(arg1 int) *string {

    retStr := C.CkMime_getHeaderFieldName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth MIME header field value.
// return a string or nil if failed.
func (z *Mime) GetHeaderFieldValue(arg1 int) *string {

    retStr := C.CkMime_getHeaderFieldValue(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a string containing the complete MIME message, including all sub-parts.
// return a string or nil if failed.
func (z *Mime) GetMime() *string {

    retStr := C.CkMime_getMime(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Appends the MIME to a BinData object.
func (z *Mime) GetMimeBd(arg1 *BinData) bool {

    v := C.CkMime_GetMimeBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns a byte array containing the complete MIME message, including all
// sub-parts.
func (z *Mime) GetMimeBytes() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkMime_GetMimeBytes(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Appends the MIME to a StringBuilder object.
func (z *Mime) GetMimeSb(arg1 *StringBuilder) bool {

    v := C.CkMime_GetMimeSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the Nth sub-part of the MIME message. Indexing begins at 0.
func (z *Mime) GetPart(arg1 int) *Mime {

    retObj := C.CkMime_GetPart(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Mime{retObj}
}


// The same as the GetSignatureSigningTime method, but returns tjhe date/time in
// RFC822 string format.
// return a string or nil if failed.
func (z *Mime) GetSignatureSigningTimeStr(arg1 int) *string {

    retStr := C.CkMime_getSignatureSigningTimeStr(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth digital certificate used to sign the MIME message. Indexing
// begins at 0.
func (z *Mime) GetSignerCert(arg1 int) *Cert {

    retObj := C.CkMime_GetSignerCert(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the full certificate chain for the Nth certificate used to sign the MIME
// message. Indexing begins at 0.
func (z *Mime) GetSignerCertChain(arg1 int) *CertChain {

    retObj := C.CkMime_GetSignerCertChain(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &CertChain{retObj}
}


// Returns a string summarizing the MIME structure. The output format is specified
// by fmt and can be "text" or "xml".
// return a string or nil if failed.
func (z *Mime) GetStructure(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkMime_getStructure(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Converts the MIME (or S/MIME) message to XML and returns the XML as a string.
// return a string or nil if failed.
func (z *Mime) GetXml() *string {

    retStr := C.CkMime_getXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if the Nth signature included a timestamp that recorded the
// signing time. The number of signatures (i.e. signer certs) is indicated by the
// NumSignerCerts property. (In most cases, the number of signer certs is 1.) The
// signing time can be obtained via the GetSignatureSigningTime or
// GetSignatureSigningTimeStr methods. The index of the 1st signature signing time
// is 0.
func (z *Mime) HasSignatureSigningTime(arg1 int) bool {

    v := C.CkMime_HasSignatureSigningTime(z.ckObj, C.int(arg1))


    return v != 0
}


// Return true if the MIME message contains application data, otherwise returns
// false.
func (z *Mime) IsApplicationData() bool {

    v := C.CkMime_IsApplicationData(z.ckObj)


    return v != 0
}


// Return true if this MIME message is an attachment, otherwise returns false.
// A MIME message is considered an attachment if the Content-Disposition header
// field contains the value "attachment".
func (z *Mime) IsAttachment() bool {

    v := C.CkMime_IsAttachment(z.ckObj)


    return v != 0
}


// Return true if the MIME message contains audio data, otherwise returns
// false.
func (z *Mime) IsAudio() bool {

    v := C.CkMime_IsAudio(z.ckObj)


    return v != 0
}


// Returns true if the MIME message is PKCS7 encrypted, otherwise returns
// false.
func (z *Mime) IsEncrypted() bool {

    v := C.CkMime_IsEncrypted(z.ckObj)


    return v != 0
}


// Return true if the MIME body is HTML, otherwise returns false.
func (z *Mime) IsHtml() bool {

    v := C.CkMime_IsHtml(z.ckObj)


    return v != 0
}


// Return true if the MIME message contains image data, otherwise returns
// false.
func (z *Mime) IsImage() bool {

    v := C.CkMime_IsImage(z.ckObj)


    return v != 0
}


// Return true if the MIME message is multipart (multipart/mixed,
// multipart/related, multipart/alternative, etc.), otherwise returns false.
func (z *Mime) IsMultipart() bool {

    v := C.CkMime_IsMultipart(z.ckObj)


    return v != 0
}


// Return true if the MIME message is multipart/alternative, otherwise returns
// false.
func (z *Mime) IsMultipartAlternative() bool {

    v := C.CkMime_IsMultipartAlternative(z.ckObj)


    return v != 0
}


// Return true if the MIME message is multipart/mixed, otherwise returns false.
func (z *Mime) IsMultipartMixed() bool {

    v := C.CkMime_IsMultipartMixed(z.ckObj)


    return v != 0
}


// Return true if the MIME message is multipart/related, otherwise returns
// false.
func (z *Mime) IsMultipartRelated() bool {

    v := C.CkMime_IsMultipartRelated(z.ckObj)


    return v != 0
}


// Return true if the MIME message body is plain text, otherwise returns false.
func (z *Mime) IsPlainText() bool {

    v := C.CkMime_IsPlainText(z.ckObj)


    return v != 0
}


// Return true if the MIME message is PKCS7 digitally signed, otherwise returns
// false.
func (z *Mime) IsSigned() bool {

    v := C.CkMime_IsSigned(z.ckObj)


    return v != 0
}


// Return true if the MIME message body is any text content type, such as
// text/plain, text/html, text/xml, etc., otherwise returns false.
func (z *Mime) IsText() bool {

    v := C.CkMime_IsText(z.ckObj)


    return v != 0
}


// Returns true if the component is already unlocked, otherwise returns false.
func (z *Mime) IsUnlocked() bool {

    v := C.CkMime_IsUnlocked(z.ckObj)


    return v != 0
}


// Return true if the MIME message contains video data, otherwise returns
// false.
func (z *Mime) IsVideo() bool {

    v := C.CkMime_IsVideo(z.ckObj)


    return v != 0
}


// Return true if the MIME message body is XML, otherwise returns false.
func (z *Mime) IsXml() bool {

    v := C.CkMime_IsXml(z.ckObj)


    return v != 0
}


// Provides information about what transpired in the last method called on this
// object instance. For many methods, there is no information. However, for some
// methods, details about what occurred can be obtained by getting the LastJsonData
// right after the method call returns. For example, after calling UnwrapSecurity,
// the LastJsonData will return JSON with details about the algorithms used for
// signature verification and decryption.
func (z *Mime) LastJsonData() *JsonObject {

    retObj := C.CkMime_LastJsonData(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Discards the current contents of the MIME object and loads a new MIME message
// from a string.
func (z *Mime) LoadMime(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_LoadMime(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Discards the current contents of the MIME object and loads a new MIME message
// from a BinData object.
func (z *Mime) LoadMimeBd(arg1 *BinData) bool {

    v := C.CkMime_LoadMimeBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads a MIME document from an in-memory byte array.
func (z *Mime) LoadMimeBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkMime_LoadMimeBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Discards the current contents of the MIME object and loads a new MIME message
// from a file.
func (z *Mime) LoadMimeFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_LoadMimeFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Discards the current contents of the MIME object and loads a new MIME message
// from a StringBuilder.
func (z *Mime) LoadMimeSb(arg1 *StringBuilder) bool {

    v := C.CkMime_LoadMimeSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Converts XML to MIME and replaces the MIME object's contents with the converted
// XML.
func (z *Mime) LoadXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_LoadXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Converts XML to MIME and replaces the MIME object's contents with the converted
// XML.
func (z *Mime) LoadXmlFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_LoadXmlFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Clears the Mime object and initializes it such that the header contains a
// "content-type: message/rfc822" line and the body is the MIME text of the Mime
// object passed to the method.
func (z *Mime) NewMessageRfc822(arg1 *Mime) bool {

    v := C.CkMime_NewMessageRfc822(z.ckObj, arg1.ckObj)


    return v != 0
}


// Discards the current MIME message header fields and contents, if any, an
// initializes the MIME object to be an empty mulipart/alternative message.
func (z *Mime) NewMultipartAlternative() bool {

    v := C.CkMime_NewMultipartAlternative(z.ckObj)


    return v != 0
}


// Discards the current MIME message header fields and contents, if any, an
// initializes the MIME object to be an empty mulipart/mixed message.
func (z *Mime) NewMultipartMixed() bool {

    v := C.CkMime_NewMultipartMixed(z.ckObj)


    return v != 0
}


// Discards the current MIME message header fields and contents, if any, an
// initializes the MIME object to be an empty mulipart/related message.
func (z *Mime) NewMultipartRelated() bool {

    v := C.CkMime_NewMultipartRelated(z.ckObj)


    return v != 0
}


// Removes a header field from the MIME header. If bAllOccurrences is true, then all
// occurrences of the header field are removed. Otherwise, only the 1st occurrence
// is removed.
func (z *Mime) RemoveHeaderField(arg1 string, arg2 bool)  {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    C.CkMime_RemoveHeaderField(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))


}


// Removes the Nth subpart from the MIME message.
func (z *Mime) RemovePart(arg1 int) bool {

    v := C.CkMime_RemovePart(z.ckObj, C.int(arg1))


    return v != 0
}


// Saves the MIME message body to a file. If the body is base64 or quoted-printable
// encoded, it is automatically decoded.
func (z *Mime) SaveBody(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_SaveBody(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the MIME message to a file, in MIME format. (This is the same as the .EML
// format used by Microsoft Outlook Express.)
func (z *Mime) SaveMime(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_SaveMime(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Converts the MIME message to XML and saves to an XML file.
func (z *Mime) SaveXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_SaveXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the MIME body content to a text string.
func (z *Mime) SetBody(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkMime_SetBody(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Sets the MIME message body from a byte array.
func (z *Mime) SetBodyFromBinary(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkMime_SetBodyFromBinary(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Sets the MIME message body from a Base64 or Quoted-Printable encoded string.
func (z *Mime) SetBodyFromEncoded(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMime_SetBodyFromEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the MIME message body from the contents of a file. Note: A MIME message
// consists of a header and a body. The body may itself be a MIME message that
// consists of a header and body, etc. This method loads the contents of a file
// into the body of a MIME message, without replacing the header.
// 
// The Content-Type and Content-Transfer-Encoding header fields are automatically
// updated to match the type of content loaded (based on file extension). If your
// application requires the MIME to have a specific Content-Type and/or
// Content-Transfer-Encoding, set the ContentType and Encoding properties after
// calling this method (not before).
//
func (z *Mime) SetBodyFromFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_SetBodyFromFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the MIME message body from a string containing HTML. The Content-Type
// header is added or updated to the value "text/html".
// 
// If 8bit (non-us-ascii) characters are present, and if the Charset property was
// not previously set, then the "charset" attribute is automatically added to the
// Content-Type header using the default value of "utf-8". This can be changed at
// any time by setting the Charset property.
// 
// If the Encoding property was not previously set, then the
// Content-Transfer-Encoding header is automatically added. It will be set to
// "7bit" or "8bit" depending on whether the HTML body contains 8-bit non-us-ascii
// characters.
// 
// To set the MIME body with no intentional side-effects, use SetBody instead.
//
func (z *Mime) SetBodyFromHtml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_SetBodyFromHtml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the MIME message body from a string containing plain-text. The Content-Type
// header is added or updated to the value "text/plain".
// 
// If 8bit (non-us-ascii) characters are present, and if the Charset property was
// not previously set, then the "charset" attribute is automatically added to the
// Content-Type header using the default value of "utf-8". This can be changed at
// any time by setting the Charset property.
// 
// If the Encoding property was not previously set, then the
// Content-Transfer-Encoding header is automatically added. It will be set to
// "7bit" or "8bit" depending on whether the plain-text body contains 8-bit
// non-us-ascii characters.
// 
// To set the MIME body with no intentional side-effects, use SetBody instead.
//
func (z *Mime) SetBodyFromPlainText(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_SetBodyFromPlainText(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the MIME message body from a string containing XML. The Content-Type header
// is added or updated to the value "text/xml".
// 
// If 8bit (non-us-ascii) characters are present, and if the Charset property was
// not previously set, then the "charset" attribute is automatically added to the
// Content-Type header using the default value of "utf-8". This can be changed at
// any time by setting the Charset property.
// 
// If the Encoding property was not previously set, then the
// Content-Transfer-Encoding header is automatically added. It will be set to
// "7bit" or "8bit" depending on whether the plain-text body contains 8-bit
// non-us-ascii characters.
// 
// To set the MIME body with no intentional side-effects, use SetBody instead.
//
func (z *Mime) SetBodyFromXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_SetBodyFromXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds or replaces a MIME message header field. If the field already exists, it is
// automatically replaced. Otherwise it is added. Pass zero-length value to remove
// the header field.
func (z *Mime) SetHeaderField(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkMime_SetHeaderField(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Allows a certificate to be explicitly specified for verifying a signature.
func (z *Mime) SetVerifyCert(arg1 *Cert) bool {

    v := C.CkMime_SetVerifyCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Unlocks the component allowing for the full functionality to be used. If this
// method unexpectedly returns false, examine the contents of the LastErrorText
// property to determine the reason for failure.
func (z *Mime) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkMime_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Decrypts and/or verifies all digital signatures contained within the MIME
// message, and returns true if all decryptions and verifications succeeded.
// Otherwise returns false. After unwrapping, the information regarding security
// and certificates can be obtained by the methods GetSignerCert and
// GetEncryptCert, and the properties NumEncryptCerts and NumSignerCerts.
// 
// The MIME is restored to the original structure/content prior to all signing
// and/or encryption.
// 
// The difference between UnwrapSecurity and methods such as Verify or Decrypt is
// that UnwrapSecurity will recursively traverse the MIME to decrypt and/or verify
// all parts. Also, UnwrapSecurity will unwrap layers until no further
// encrypted/signed content is found. For example, if a MIME message was encrypted
// and then subsequently signed, then UnwrapSecurity will verify and unwrap the
// detached signature/signed-data layer, and then decrypt the "enveloped data".
//
func (z *Mime) UnwrapSecurity() bool {

    v := C.CkMime_UnwrapSecurity(z.ckObj)


    return v != 0
}


// URL encodes the MIME body. The charset is important. For example, consider this
// MIME:
// Content-Type: text/plain
// Content-Transfer-Encoding: 8bit
// 
// Société
// If the charset is set to "utf-8", then the following is produced:
// Content-Type: text/plain
// Content-Transfer-Encoding: 8bit
// 
// Soci%C3%A9t%C3%A9
// However, if the charset is set to "ansi", then the following is the result:
// Content-Type: text/plain
// Content-Transfer-Encoding: 8bit
// 
// Soci%E9t%E9
func (z *Mime) UrlEncodeBody(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkMime_UrlEncodeBody(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Adds an XML certificate vault to the object's internal list of sources to be
// searched for certificates and private keys when encrypting/decrypting or
// signing/verifying. Unlike the AddPfxSourceData and AddPfxSourceFile methods,
// only a single XML certificate vault can be used. If UseCertVault is called
// multiple times, only the last certificate vault will be used, as each call to
// UseCertVault will replace the certificate vault provided in previous calls.
func (z *Mime) UseCertVault(arg1 *XmlCertVault) bool {

    v := C.CkMime_UseCertVault(z.ckObj, arg1.ckObj)


    return v != 0
}


// Verifies PKCS7 signed MIME and "unwraps" the signature. The MIME is restored to
// the original structure that it would have originally had prior to signing. The
// Verify method works with both detached signatures, as well as opaque/attached
// signatures.
// 
// A PKCS7 signature usually embeds both the signing certificate with its public
// key. Therefore, it is usually possible to verify a signature without the need to
// already have the certificate installed. If the signature does not embed the
// certificate, the Verify method will automatically locate and use the certificate
// if it was correctly pre-installed on the computer.
//
func (z *Mime) Verify() bool {

    v := C.CkMime_Verify(z.ckObj)


    return v != 0
}


