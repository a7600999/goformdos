// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkPfx.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewPfx() *Pfx {
	return &Pfx{C.CkPfx_Create()}
}

func (z *Pfx) DisposePfx() {
    if z != nil {
	C.CkPfx_Dispose(z.ckObj)
	}
}




// property: AlgorithmId
// The encryption algorithm to be used when writing a PFX. After loading a PFX,
// this property is set to the encryption algorithm used by the loaded PFX. (This
// is the algorithm used for the "shrouded key bag", which is internal to the PFX.)
// 
// The default value (for backward compatibility) is
// "pbeWithSHAAnd3_KeyTripleDES_CBC". Can be set to "pbes2", in which case the
// Pbes2CryptAlg and Pbes2HmacAlg properies will be set to the algorithms to be
// used when writing, or the algorithms used by the loaded PFX.
//
func (z *Pfx) AlgorithmId() string {
    return C.GoString(C.CkPfx_algorithmId(z.ckObj))
}
// property setter: AlgorithmId
func (z *Pfx) SetAlgorithmId(goStr string) {
    cStr := C.CString(goStr)
    C.CkPfx_putAlgorithmId(z.ckObj,cStr)
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
func (z *Pfx) DebugLogFilePath() string {
    return C.GoString(C.CkPfx_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Pfx) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkPfx_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Pfx) LastErrorHtml() string {
    return C.GoString(C.CkPfx_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Pfx) LastErrorText() string {
    return C.GoString(C.CkPfx_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Pfx) LastErrorXml() string {
    return C.GoString(C.CkPfx_lastErrorXml(z.ckObj))
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
func (z *Pfx) LastMethodSuccess() bool {
    v := int(C.CkPfx_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Pfx) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPfx_putLastMethodSuccess(z.ckObj,v)
}

// property: NumCerts
// The number of certificates contained in the PFX.
func (z *Pfx) NumCerts() int {
    return int(C.CkPfx_getNumCerts(z.ckObj))
}

// property: NumPrivateKeys
// The number of private keys contained in the PFX.
func (z *Pfx) NumPrivateKeys() int {
    return int(C.CkPfx_getNumPrivateKeys(z.ckObj))
}

// property: Pbes2CryptAlg
// If the AlgorithmId property equals "pbes2", then this is the encryption
// algorithm to be used when writing the PFX, or used by the PFX that was loaded.
// If the AlgorithmId is not equal to "pbes2", then the value of this property is
// meaningless.
// 
// Possible values are:
//     aes256-cbc
//     aes192-cbc
//     aes128-cbc
//     3des-cbc
// 
// The default value (for writing) is "aes256-cbc". Note: The algorithm specified
// by this property is only used when the Algorithmid = "pbes2".
//
func (z *Pfx) Pbes2CryptAlg() string {
    return C.GoString(C.CkPfx_pbes2CryptAlg(z.ckObj))
}
// property setter: Pbes2CryptAlg
func (z *Pfx) SetPbes2CryptAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkPfx_putPbes2CryptAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Pbes2HmacAlg
// If the AlgorithmId property equals "pbes2", then this is the HMAC hash algorithm
// to be used when writing the PFX, or used by the PFX that was loaded. If the
// AlgorithmId is not equal to "pbes2", then the value of this property is
// meaningless.
// 
// Possible values are:
//     hmacWithSha256
//     hmacWithSha384
//     hmacWithSha512
//     hmacWithSha1
// 
// The default value (for writing) is "hmacWithSha256". Note: The algorithm
// specified by this property is only used when the Algorithmid = "pbes2".
//
func (z *Pfx) Pbes2HmacAlg() string {
    return C.GoString(C.CkPfx_pbes2HmacAlg(z.ckObj))
}
// property setter: Pbes2HmacAlg
func (z *Pfx) SetPbes2HmacAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkPfx_putPbes2HmacAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. This property
// defaults to the empty string and should typically remain empty.
// 
// Can be set to a list of the following comma separated keywords:
//     "LegacyOrder" - Introduced in v9.5.0.83. Write the internal ContentInfos in
//     the order Chilkat traditionally used in previous versions.
//
func (z *Pfx) UncommonOptions() string {
    return C.GoString(C.CkPfx_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Pfx) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkPfx_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Pfx) VerboseLogging() bool {
    v := int(C.CkPfx_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Pfx) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPfx_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Pfx) Version() string {
    return C.GoString(C.CkPfx_version(z.ckObj))
}
// Adds a certificate, its private key (if it exists), and potentially its
// certificate chain to the PFX. If includeChain is true, then the certificate must have
// a private key. The certificate's private key is automatically obtained
// (internally) via the cert's ExportPrivateKey method. If the certificate's chain
// of authentication is to be added, it is automatically constructed and added
// using whatever resources are at hand (such as certs provided via the
// UseCertVault method, the trusted roots from Chilkat's TrustedRoots class, etc.
// If a certificate chain is to be added, which is the typical case, then the chain
// must be completed to the root to succeed.
func (z *Pfx) AddCert(arg1 *Cert, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkPfx_AddCert(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Adds a private key and certificate chain to the PFX. The private key should be
// such that it is associated with the 1st certificate in the chain. In other
// words, the 1st certificate in the chain has a public key (embedded within the
// X.509 structure of the cert itself) that is the counterpart to the private key.
func (z *Pfx) AddPrivateKey(arg1 *PrivateKey, arg2 *CertChain) bool {

    v := C.CkPfx_AddPrivateKey(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Finds and returns the certificate (in the PFX) that has a cert bag "localKeyId"
// attribute with the specified value. The localKeyId is specifid using the encoding
// (such as "decimal", "base64", "hex") specified by encoding.
func (z *Pfx) FindCertByLocalKeyId(arg1 string, arg2 string) *Cert {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkPfx_FindCertByLocalKeyId(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the Nth certificate in the PFX. (The 1st certificate is at index 0.)
func (z *Pfx) GetCert(arg1 int) *Cert {

    retObj := C.CkPfx_GetCert(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the Nth private key in the PFX. (The 1st private key is at index 0.)
func (z *Pfx) GetPrivateKey(arg1 int) *PrivateKey {

    retObj := C.CkPfx_GetPrivateKey(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &PrivateKey{retObj}
}


// Can be called to get one of the following safebag attributes for the Nth private
// key or certificate in the PFX. forPrivateKey should be true for a private key, and
// false for a certificate. The index is the index of the certificate or key in
// the PFX. The attrName can be one of the following:
// 
//     "localKeyId" : Returns the decimal representation of the local key ID. The
//     local key ID is used to associate the certificate contained in the PFX with this
//     private key. (The certificate will include a "localKeyId" attribute in its cert
//     bag of attributes within the PFX.)
//     "keyContainerName" : Returns the key container name (or key name) of the
//     private key. For more information about the directories where the Windows OS
//     stores private keys,
//     seehttps://docs.microsoft.com/en-us/windows/win32/seccng/key-storage-and-retrieva
//     l
//     <https://docs.microsoft.com/en-us/windows/win32/seccng/key-storage-and-retrieval>
//     "storageProvider" : Returns the name of the Cryptographic Storage Provider
//     to be used for the key.
// 
// Note: It is not required that any of the above attributes are present in the
// PFX.
//
// return a string or nil if failed.
func (z *Pfx) GetSafeBagAttr(arg1 bool, arg2 int, arg3 string) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    retStr := C.CkPfx_getSafeBagAttr(z.ckObj, b1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Provides information about what transpired in the last method called.
func (z *Pfx) LastJsonData() *JsonObject {

    retObj := C.CkPfx_LastJsonData(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Loads a PFX from a PEM formatted string. The PEM can contain the private key,
// the certificate, and certificates in the chain of authentication up to the CA
// root. For example:
//  -----BEGIN RSA PRIVATE KEY-----
// ...
// ... the private key associated with the main certificate.
// ...
// -----END RSA PRIVATE KEY-----
// -----BEGIN CERTIFICATE-----
// ...
// ... the main certificate
// ...
// -----END CERTIFICATE-----
// -----BEGIN CERTIFICATE-----
// ...
// ... an intermediate CA certificate (if present)
// ...
// -----END CERTIFICATE-----
// -----BEGIN CERTIFICATE-----
// ...
// ... the root CA certificate
// ...
// -----END CERTIFICATE-----
func (z *Pfx) LoadPem(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPfx_LoadPem(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a PFX from in-memory bytes.
// 
// If the .pfx/.p12 uses different passwords for integrity and private keys, then
// the password argument may contain JSON to specify the passwords. See the LoadPfxFile
// method (below) for details.
//
func (z *Pfx) LoadPfxBytes(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkPfx_LoadPfxBytes(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a PFX from encoded byte data. The encoding can by any encoding, such as
// "Base64", "modBase64", "Base32", "UU", "QP" (for quoted-printable), "URL" (for
// url-encoding), "Hex", "Q", "B", "url_oath", "url_rfc1738", "url_rfc2396", and
// "url_rfc3986".
// 
// If the .pfx/.p12 uses different passwords for integrity and private keys, then
// the encoding argument may contain JSON to specify the passwords. See the LoadPfxFile
// method (below) for details.
//
func (z *Pfx) LoadPfxEncoded(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkPfx_LoadPfxEncoded(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Loads a PFX from a file.
// 
// Starting in v9.5.0.75, a .pfx/.p12 file with different passwords for integrity
// and private keys can be loaded by passing the following JSON for the password.
//     {
//       "integrity": "password1",
//       "privKeys": "password2",
//      }
// If it is desired to open the .pfx/.p12 without access to the private keys, then
// add "skipPrivateKeys" like this:
//     {
//       "integrity": "password1",
//       "privKeys": "not used",
//        "skipPrivateKeys": true
//      }
//
func (z *Pfx) LoadPfxFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPfx_LoadPfxFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets a safe bag attribute for the Nth private key or certificate in the PFX.
// Safe bag attributes can be added by calling this method once for each attribute
// to be added to each certificate or key. forPrivateKey should be true for a private key,
// and false for a certificate. The index is the index of the certificate or key
// in the PFX. (The 1st item is at index 0.) See the example below for more
// information. The encoding indicates a binary encoding such as "base64", "hex",
// "decimal", "fingerprint", etc if the value contains binary (non-text) data.
// 
// A safe bag attribute can be removed by passing an empty string for the value.
//
func (z *Pfx) SetSafeBagAttr(arg1 bool, arg2 int, arg3 string, arg4 string, arg5 string) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkPfx_SetSafeBagAttr(z.ckObj, b1, C.int(arg2), cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Write the PFX to in-memory bytes.
func (z *Pfx) ToBinary(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkPfx_ToBinary(z.ckObj, cstr1, ckOutBytes)

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


// Write the PFX to an encoded string. The encoding can be any encoding such as
// "base64" or "hex".
// return a string or nil if failed.
func (z *Pfx) ToEncodedString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkPfx_toEncodedString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Write the PFX to a file. PFX and PKCS12 are essentially the same. Standard
// filename extensions are ".pfx" or ".p12".
func (z *Pfx) ToFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPfx_ToFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Converts the PFX (PKCS12) to a JavaKeyStore object. One JKS entry per private
// key found in the PKCS12 is added. The certs found within the PCKS12 are used to
// build the certificate chains for each private key. (A typical PFX file contains
// a single private key along with its associated certificate, and the certificates
// in the chain of authentication to the root CA cert.)
// 
// The specified alias is applied to the 1st private key found. If the alias is
// empty, then the alias is obtained from the cert/PFX in the following order of
// preference:
//     Certificate's subject common name
//     Certificate's subject email address
//     Certificate's friendly name found in the PKCS9 attributes of the PKCS12
//     Certificate's serial number
// 
// If multiple private keys are found in the PKCS12, then all but the first will
// automaticallly be assigned aliases using the preference just described.
//
func (z *Pfx) ToJavaKeyStore(arg1 string, arg2 string) *JavaKeyStore {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkPfx_ToJavaKeyStore(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &JavaKeyStore{retObj}
}


// Write the PFX to a PEM formatted string. The resultant PEM will contain the
// private key, as well as the certs in the chain of authentication (or whatever
// certs are available in the PFX). For example:
//  -----BEGIN RSA PRIVATE KEY-----
// ...
// ... the private key associated with the main certificate.
// ...
// -----END RSA PRIVATE KEY-----
// -----BEGIN CERTIFICATE-----
// ...
// ... the main certificate
// ...
// -----END CERTIFICATE-----
// -----BEGIN CERTIFICATE-----
// ...
// ... an intermediate CA certificate (if present)
// ...
// -----END CERTIFICATE-----
// -----BEGIN CERTIFICATE-----
// ...
// ... the root CA certificate
// ...
// -----END CERTIFICATE-----
// return a string or nil if failed.
func (z *Pfx) ToPem() *string {

    retStr := C.CkPfx_toPem(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Write the PFX to a PEM formatted string. If extendedAttrs is true, then extended
// properties (Bag Attributes and Key Attributes) are output. If noKeys is true,
// then no private keys are output. If noCerts is true, then no certificates are
// output. If noCaCerts is true, then no CA certs or intermediate CA certs are output.
// If encryptAlg is not empty, it indicates the encryption algorithm to be used for
// encrypting the private keys (otherwise the private keys are output unencrypted).
// The possible choices for the encryptAlg are "des3", "aes128", "aes192", and "aes256".
// (All encryption algorithm choices use CBC mode.) If the private keys are to be
// encrypted, then password is the password to be used. Otherwise, password may be left
// empty. For example:
// Bag Attributes
//     Microsoft Local Key set: localKeyID: 01 00 00 00 
//     friendlyName: le-2b09a3d2-9037-4a05-95cc-4d44518e8607
//     Microsoft CSP Name: Microsoft RSA SChannel Cryptographic Provider
// Key Attributes
//     X509v3 Key Usage: 10 
//  -----BEGIN RSA PRIVATE KEY-----
// ...
// ... the private key associated with the main certificate.
// ...
// -----END RSA PRIVATE KEY-----
// Bag Attributes
//     localKeyID: 01 00 00 00 
//     1.3.6.1.4.1.311.17.3.92: 00 08 00 00 
//     1.3.6.1.4.1.311.17.3.20: C2 53 54 F3 ...
//     1.3.6.1.4.1.311.17.3.71: 49 00 43 00 ...
//     1.3.6.1.4.1.311.17.3.75: 31 00 42 00 ...
// subject=/OU=Domain Control Validated/OU=PositiveSSL/CN=something.com
// issuer=/C=GB/ST=Greater Manchester/L=Salford/O=COMODO CA Limited/CN=COMODO RSA Domain Validation Secure Server CA
// -----BEGIN CERTIFICATE-----
// ...
// ... the main certificate
// ...
// -----END CERTIFICATE-----
// ...
// -----BEGIN CERTIFICATE-----
// ...
// ... an intermediate CA certificate (if present)
// ...
// -----END CERTIFICATE-----
// ...
// -----BEGIN CERTIFICATE-----
// ...
// ... the root CA certificate
// ...
// -----END CERTIFICATE-----
// return a string or nil if failed.
func (z *Pfx) ToPemEx(arg1 bool, arg2 bool, arg3 bool, arg4 bool, arg5 string, arg6 string) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    cstr5 := C.CString(arg5)
    cstr6 := C.CString(arg6)

    retStr := C.CkPfx_toPemEx(z.ckObj, b1, b2, b3, b4, cstr5, cstr6)

    C.free(unsafe.Pointer(cstr5))
    C.free(unsafe.Pointer(cstr6))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Adds an XML certificate vault to the object's internal list of sources to be
// searched for certificates for help in building certificate chains to a root
// certificate.
func (z *Pfx) UseCertVault(arg1 *XmlCertVault) bool {

    v := C.CkPfx_UseCertVault(z.ckObj, arg1.ckObj)


    return v != 0
}


