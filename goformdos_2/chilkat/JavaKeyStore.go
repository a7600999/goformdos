// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkJavaKeyStore.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewJavaKeyStore() *JavaKeyStore {
	return &JavaKeyStore{C.CkJavaKeyStore_Create()}
}

func (z *JavaKeyStore) DisposeJavaKeyStore() {
    if z != nil {
	C.CkJavaKeyStore_Dispose(z.ckObj)
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
func (z *JavaKeyStore) DebugLogFilePath() string {
    return C.GoString(C.CkJavaKeyStore_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *JavaKeyStore) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkJavaKeyStore_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *JavaKeyStore) LastErrorHtml() string {
    return C.GoString(C.CkJavaKeyStore_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *JavaKeyStore) LastErrorText() string {
    return C.GoString(C.CkJavaKeyStore_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *JavaKeyStore) LastErrorXml() string {
    return C.GoString(C.CkJavaKeyStore_lastErrorXml(z.ckObj))
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
func (z *JavaKeyStore) LastMethodSuccess() bool {
    v := int(C.CkJavaKeyStore_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *JavaKeyStore) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJavaKeyStore_putLastMethodSuccess(z.ckObj,v)
}

// property: NumPrivateKeys
// The number of private keys contained within the keystore. Each private key has
// an alias and certificate chain associated with it.
func (z *JavaKeyStore) NumPrivateKeys() int {
    return int(C.CkJavaKeyStore_getNumPrivateKeys(z.ckObj))
}

// property: NumSecretKeys
// The number of secret keys (such as AES keys) contained within the keystore. Each
// secret key can have an alias associated with it.
func (z *JavaKeyStore) NumSecretKeys() int {
    return int(C.CkJavaKeyStore_getNumSecretKeys(z.ckObj))
}

// property: NumTrustedCerts
// The number of trusted certificates contained within the keystore. Each
// certificate has an alias (identifying string) associated with it.
func (z *JavaKeyStore) NumTrustedCerts() int {
    return int(C.CkJavaKeyStore_getNumTrustedCerts(z.ckObj))
}

// property: RequireCompleteChain
// If true, then adding a private key to the JKS only succeeds if the certificate
// chain can be completed to the root certificate. A root certificate is either a
// trusted CA root or a self-signed certificate. If false, then incomplete
// certificate chains are allowed. The default value is true.
func (z *JavaKeyStore) RequireCompleteChain() bool {
    v := int(C.CkJavaKeyStore_getRequireCompleteChain(z.ckObj))
    return v != 0
}
// property setter: RequireCompleteChain
func (z *JavaKeyStore) SetRequireCompleteChain(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJavaKeyStore_putRequireCompleteChain(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *JavaKeyStore) VerboseLogging() bool {
    v := int(C.CkJavaKeyStore_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *JavaKeyStore) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJavaKeyStore_putVerboseLogging(z.ckObj,v)
}

// property: VerifyKeyedDigest
// If true, then the keystore's keyed digest is required to pass validation
// (password required) for any of the load methods (such as LoadFile, LoadBinary,
// or LoadEncoded). If false, then a keystore may be loaded into memory without
// password validation (if a null or empty string is passed to the load method).
// The default value of this property is true.
func (z *JavaKeyStore) VerifyKeyedDigest() bool {
    v := int(C.CkJavaKeyStore_getVerifyKeyedDigest(z.ckObj))
    return v != 0
}
// property setter: VerifyKeyedDigest
func (z *JavaKeyStore) SetVerifyKeyedDigest(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJavaKeyStore_putVerifyKeyedDigest(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *JavaKeyStore) Version() string {
    return C.GoString(C.CkJavaKeyStore_version(z.ckObj))
}
// Adds the contents of a PFX or PKCS #12 (.p12) to the Java keystore object. One
// JKS entry per private key found in the PKCS12 is added. The certs found within
// the PCKS12 are used to build the certificate chains for each private key. (A
// typical PFX file contains a single private key along with its associated
// certificate, and the certificates in the chain of authentication to the root CA
// cert.)
// 
// This method does not add trusted certificate entries to the JKS.
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
// The UseCertVault method may be called to provide additional certificates for the
// automatic construction of the certificate chains. If the RequireCompleteChain
// property is set to true, then this method will fail if any certificate chain
// is not completed to the root. The TrustedRoots class may be used to provide a
// source for obtaining trusted CA roots if these are not already present within
// the PKCS12.
//
func (z *JavaKeyStore) AddPfx(arg1 *Pfx, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkJavaKeyStore_AddPfx(z.ckObj, arg1.ckObj, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds a private key entry to the JKS. Both the private key and certificate chain
// are obtained from the certificate object that is passed in the 1st argument.
// 
// If the alias is empty, then the alias is automatically chosen based on the
// certificate's information, in the following order of preference:
//     Certificate's subject common name
//     Certificate's subject email address
//     Certificate's serial number
// 
// The UseCertVault method may be called to provide additional certificates for the
// automatic construction of the certificate chains. If the RequireCompleteChain
// property is set to true, then this method will fail if the certificate chain
// is not completed to the root. The TrustedRoots class may be used to provide a
// source for obtaining trusted CA roots.
//
func (z *JavaKeyStore) AddPrivateKey(arg1 *Cert, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkJavaKeyStore_AddPrivateKey(z.ckObj, arg1.ckObj, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds a secret (symmetric) key entry to the JKS. This adds a symmetric key, which
// is simply a number of binary bytes (such as 16 bytes for a 128-bit AES key). The
// encodedKeyBytes provides the actual bytes of the symmetric key, in an encoded string form.
// The encoding indicates the encoding of encodedKeyBytes (such as "base64", "hex", "base64url",
// etc.) The algorithm describes the symmetric algorithm, such as "AES". The alias is the
// password used to seal (encrypt) the key bytes.
// 
// Note: The algorithm describes the usage of the encodedKeyBytes. For example, if encodedKeyBytes contains
// the 16 bytes of a 128-bit AES key, then algorithm should be set to "AES". The actual
// encryption algorithm used to seal the key within the JCEKS is
// PBEWithMD5AndTripleDES, which is part of the JCEKS specification.
//
func (z *JavaKeyStore) AddSecretKey(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkJavaKeyStore_AddSecretKey(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Adds a trusted certificate to the Java keystore object.
func (z *JavaKeyStore) AddTrustedCert(arg1 *Cert, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJavaKeyStore_AddTrustedCert(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Changes the password for a private key.
func (z *JavaKeyStore) ChangePassword(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkJavaKeyStore_ChangePassword(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Finds and returns the certificate chain for the private key with the specified
// alias.
func (z *JavaKeyStore) FindCertChain(arg1 string, arg2 bool) *CertChain {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retObj := C.CkJavaKeyStore_FindCertChain(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CertChain{retObj}
}


// Finds and returns the private key with the specified alias.
func (z *JavaKeyStore) FindPrivateKey(arg1 string, arg2 string, arg3 bool) *PrivateKey {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retObj := C.CkJavaKeyStore_FindPrivateKey(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &PrivateKey{retObj}
}


// Finds and returns the trusted certificate with the specified alias.
func (z *JavaKeyStore) FindTrustedCert(arg1 string, arg2 bool) *Cert {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retObj := C.CkJavaKeyStore_FindTrustedCert(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the certificate chain associated with the Nth private key contained
// within the keystore. The 1st private key is at index 0.
func (z *JavaKeyStore) GetCertChain(arg1 int) *CertChain {

    retObj := C.CkJavaKeyStore_GetCertChain(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &CertChain{retObj}
}


// Returns the Nth private key contained within the keystore. The 1st private key
// is at index 0.
func (z *JavaKeyStore) GetPrivateKey(arg1 string, arg2 int) *PrivateKey {
    cstr1 := C.CString(arg1)

    retObj := C.CkJavaKeyStore_GetPrivateKey(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &PrivateKey{retObj}
}


// Returns the Nth private key alias contained within the keystore. The 1st private
// key is at index 0.
// return a string or nil if failed.
func (z *JavaKeyStore) GetPrivateKeyAlias(arg1 int) *string {

    retStr := C.CkJavaKeyStore_getPrivateKeyAlias(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth secret key contained within the keystore. The 1st secret key is
// at index 0. The bytes of the secret key are returned in the specified encoding.
// (such as hex, base64, base64url, etc.)
// return a string or nil if failed.
func (z *JavaKeyStore) GetSecretKey(arg1 string, arg2 int, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    retStr := C.CkJavaKeyStore_getSecretKey(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth secret key alias contained within the keystore. The 1st secret
// key is at index 0.
// return a string or nil if failed.
func (z *JavaKeyStore) GetSecretKeyAlias(arg1 int) *string {

    retStr := C.CkJavaKeyStore_getSecretKeyAlias(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth trusted certificate contained within the keystore. The 1st
// certificate is at index 0.
func (z *JavaKeyStore) GetTrustedCert(arg1 int) *Cert {

    retObj := C.CkJavaKeyStore_GetTrustedCert(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the Nth trusted certificate alias contained within the keystore. The 1st
// certificate is at index 0.
// return a string or nil if failed.
func (z *JavaKeyStore) GetTrustedCertAlias(arg1 int) *string {

    retStr := C.CkJavaKeyStore_getTrustedCertAlias(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads a Java keystore from the contents of bd.
func (z *JavaKeyStore) LoadBd(arg1 string, arg2 *BinData) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJavaKeyStore_LoadBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a Java keystore from in-memory byte data.
func (z *JavaKeyStore) LoadBinary(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkJavaKeyStore_LoadBinary(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Loads a Java keystore from an encoded string (such as base64, hex, etc.)
func (z *JavaKeyStore) LoadEncoded(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkJavaKeyStore_LoadEncoded(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Loads a Java keystore from a file.
func (z *JavaKeyStore) LoadFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJavaKeyStore_LoadFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads the Java KeyStore from a JSON Web Key (JWK) Set.
func (z *JavaKeyStore) LoadJwkSet(arg1 string, arg2 *JsonObject) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJavaKeyStore_LoadJwkSet(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Removes the Nth trusted certificate or private key entry from the keystore. The
// entryType indicates whether it is a trusted root or private key entry (1 = trusted
// certificate entry, 2 = private key entry). The 1st entry is at index 0.
func (z *JavaKeyStore) RemoveEntry(arg1 int, arg2 int) bool {

    v := C.CkJavaKeyStore_RemoveEntry(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Sets the alias name for a trusted certificate or private key entry. The entryType
// indicates whether it is a trusted root or private key entry (1 = trusted
// certificate entry, 2 = private key entry). The 1st entry is at index 0.
func (z *JavaKeyStore) SetAlias(arg1 int, arg2 int, arg3 string) bool {
    cstr3 := C.CString(arg3)

    v := C.CkJavaKeyStore_SetAlias(z.ckObj, C.int(arg1), C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Writes the key store to in-memory bytes. The password is used for the keyed hash of
// the entire JKS file. (Each private key within the file may use different
// passwords, and these are provided when the private key is added via the
// AddPrivateKey method.)
func (z *JavaKeyStore) ToBinary(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkJavaKeyStore_ToBinary(z.ckObj, cstr1, ckOutBytes)

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


// Writes the key store to an encoded string. The encoding can be any encoding such as
// "base64" or "hex". The password is used for the keyed hash of the entire JKS file.
// (Each private key within the file may use different passwords, and these are
// provided when the private key is added via the AddPrivateKey method.)
// return a string or nil if failed.
func (z *JavaKeyStore) ToEncodedString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkJavaKeyStore_toEncodedString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Writes the key store to a file. The password is used for the keyed hash of the
// entire JKS file. (Each private key within the file may use different passwords,
// and these are provided when the private key is added via the AddPrivateKey
// method.)
func (z *JavaKeyStore) ToFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJavaKeyStore_ToFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the private keys in JSON JWK Set format. The JWK identifier (kid) will
// be set from the key alias in the store.
func (z *JavaKeyStore) ToJwkSet(arg1 string, arg2 *StringBuilder) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJavaKeyStore_ToJwkSet(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the Java KeyStore as a Pem object.
func (z *JavaKeyStore) ToPem(arg1 string) *Pem {
    cstr1 := C.CString(arg1)

    retObj := C.CkJavaKeyStore_ToPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Pem{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Returns the Java KeyStore as a Pfx object.
func (z *JavaKeyStore) ToPfx(arg1 string) *Pfx {
    cstr1 := C.CString(arg1)

    retObj := C.CkJavaKeyStore_ToPfx(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Pfx{retObj}
}


// Unlocks the component allowing for the full functionality to be used. If a
// purchased unlock code is passed, there is no expiration. Any other string
// automatically begins a fully-functional 30-day trial the first time
// UnlockComponent is called.
func (z *JavaKeyStore) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJavaKeyStore_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds an XML certificate vault to the object's internal list of sources to be
// searched for certificates for help in building certificate chains to a root
// certificate.
func (z *JavaKeyStore) UseCertVault(arg1 *XmlCertVault) bool {

    v := C.CkJavaKeyStore_UseCertVault(z.ckObj, arg1.ckObj)


    return v != 0
}


