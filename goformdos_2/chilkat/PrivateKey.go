// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkPrivateKey.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewPrivateKey() *PrivateKey {
	return &PrivateKey{C.CkPrivateKey_Create()}
}

func (z *PrivateKey) DisposePrivateKey() {
    if z != nil {
	C.CkPrivateKey_Dispose(z.ckObj)
	}
}




// property: BitLength
// The bit length (strength) of the private key.
func (z *PrivateKey) BitLength() int {
    return int(C.CkPrivateKey_getBitLength(z.ckObj))
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
func (z *PrivateKey) DebugLogFilePath() string {
    return C.GoString(C.CkPrivateKey_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *PrivateKey) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkPrivateKey_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: KeyType
// The type of private key. Can be "empty", "rsa", "dsa", "ecc" (i.e. ECDSA), or
// "ed25519".
func (z *PrivateKey) KeyType() string {
    return C.GoString(C.CkPrivateKey_keyType(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *PrivateKey) LastErrorHtml() string {
    return C.GoString(C.CkPrivateKey_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *PrivateKey) LastErrorText() string {
    return C.GoString(C.CkPrivateKey_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *PrivateKey) LastErrorXml() string {
    return C.GoString(C.CkPrivateKey_lastErrorXml(z.ckObj))
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
func (z *PrivateKey) LastMethodSuccess() bool {
    v := int(C.CkPrivateKey_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *PrivateKey) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPrivateKey_putLastMethodSuccess(z.ckObj,v)
}

// property: Pkcs8EncryptAlg
// The encryption algorithm to be used when exporting the key to encrypted PKCS8.
// The default value is "3des". Possible choices also include "aes128", "aes192",
// and "aes256". All of the encryption algorithm choices use CBC mode.
func (z *PrivateKey) Pkcs8EncryptAlg() string {
    return C.GoString(C.CkPrivateKey_pkcs8EncryptAlg(z.ckObj))
}
// property setter: Pkcs8EncryptAlg
func (z *PrivateKey) SetPkcs8EncryptAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkPrivateKey_putPkcs8EncryptAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *PrivateKey) VerboseLogging() bool {
    v := int(C.CkPrivateKey_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *PrivateKey) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPrivateKey_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *PrivateKey) Version() string {
    return C.GoString(C.CkPrivateKey_version(z.ckObj))
}
// Gets the private key in JWK (JSON Web Key) format.
// 
// RSA keys have this JWK format:
//          {"kty":"RSA",
//           "n":"0vx7agoebGcQ ... JzKnqDKgw",
//           "e":"AQAB",
//           "d":"X4cTteJY_gn4F ... 4jfcKoAC8Q",
//           "p":"83i-7IvMGXoMX ... vn7O0nVbfs",
//           "q":"3dfOR9cuYq-0S ... 4vIcb6yelxk",
//           "dp":"G4sPXkc6Ya9 ... 8YeiKkTiBj0",
//           "dq":"s9lAH9fggBso ... w494Q_cgk",
//           "qi":"GyM_p6JrXySi ... zTKhAVRU"}
// 
// ECC keys have this JWK format.
//          {"kty":"EC",
//           "crv":"P-256",
//           "x":"MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4",
//           "y":"4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM",
//           "d":"870MB6gfuTJ4HtUnUvYMyJpr5eUZNP4Bk43bVdj3eAE"}
// 
// Ed25519 keys (added in v9.5.0.83) have this JWK format.
//          {"kty": "OKP",
//          "crv": "Ed25519",
//          "x": "SE2Kne5xt51z1eciMH2T2ftDQp96Gl6FhY6zSQujiP0",
//          "d": "O-eRXewadF0sNyB0U9omcnt8Qg2ZmeK3WSXPYgqe570",
//          "use": "sig"}
//
// return a string or nil if failed.
func (z *PrivateKey) GetJwk() *string {

    retStr := C.CkPrivateKey_getJwk(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the JWK thumbprint for the private key. This is the thumbprint of the
// JSON Web Key (JWK) as per RFC 7638.
// return a string or nil if failed.
func (z *PrivateKey) GetJwkThumbprint(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkPrivateKey_getJwkThumbprint(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets the private key in unencrypted binary DER format, preferring PKCS1 if
// possible.
// 
// RSA keys are returned in PKCS1 ASN.1 DER format:
// RSAPrivateKey ::= SEQUENCE {
//     version           Version,
//     modulus           INTEGER,  -- n
//     publicExponent    INTEGER,  -- e
//     privateExponent   INTEGER,  -- d
//     prime1            INTEGER,  -- p
//     prime2            INTEGER,  -- q
//     exponent1         INTEGER,  -- d mod (p-1)
//     exponent2         INTEGER,  -- d mod (q-1)
//     coefficient       INTEGER,  -- (inverse of q) mod p
//     otherPrimeInfos   OtherPrimeInfos OPTIONAL
// }
// 
// DSA keys are returned in this ASN.1 DER format:
// SEQUENCE(6 elem)
//     INTEGER 0
//     INTEGER(2048 bit) (p) 
//     INTEGER(160 bit) (q) 
//     INTEGER(2044 bit) (g) 
//     INTEGER(2040 bit) (y - public key) 
//     INTEGER(156 bit) (x - private key) 
// 
// ECC keys are returned in this ASN.1 DER format:
// (from RFC 5915)
// ECPrivateKey ::= SEQUENCE {
//     version        INTEGER { ecPrivkeyVer1(1) } (ecPrivkeyVer1),
//     privateKey     OCTET STRING,
//     parameters [0] ECParameters {{ NamedCurve }} OPTIONAL,
//     publicKey  [1] BIT STRING OPTIONAL (This is the ANSI X9.63 public key format.)
//
func (z *PrivateKey) GetPkcs1() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkPrivateKey_GetPkcs1(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Gets the private key in unencrypted binary DER format, preferring PKCS1 if
// possible, and returns in an encoded string, as specified by the encoding argument.
// 
// RSA keys are returned in PKCS1 ASN.1 DER format:
// RSAPrivateKey ::= SEQUENCE {
//     version           Version,
//     modulus           INTEGER,  -- n
//     publicExponent    INTEGER,  -- e
//     privateExponent   INTEGER,  -- d
//     prime1            INTEGER,  -- p
//     prime2            INTEGER,  -- q
//     exponent1         INTEGER,  -- d mod (p-1)
//     exponent2         INTEGER,  -- d mod (q-1)
//     coefficient       INTEGER,  -- (inverse of q) mod p
//     otherPrimeInfos   OtherPrimeInfos OPTIONAL
// }
// 
// DSA keys are returned in this ASN.1 DER format:
// SEQUENCE(6 elem)
//     INTEGER 0
//     INTEGER(2048 bit) (p) 
//     INTEGER(160 bit) (q) 
//     INTEGER(2044 bit) (g) 
//     INTEGER(2040 bit) (y - public key) 
//     INTEGER(156 bit) (x - private key) 
// 
// ECC keys are returned in this ASN.1 DER format:
// (from RFC 5915)
// ECPrivateKey ::= SEQUENCE {
//     version        INTEGER { ecPrivkeyVer1(1) } (ecPrivkeyVer1),
//     privateKey     OCTET STRING,
//     parameters [0] ECParameters {{ NamedCurve }} OPTIONAL,
//     publicKey  [1] BIT STRING OPTIONAL (This is the ANSI X9.63 public key format.)
//
// return a string or nil if failed.
func (z *PrivateKey) GetPkcs1ENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkPrivateKey_getPkcs1ENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets the private key in non-encrypted PEM format, preferring PKCS1 over PKCS8 if
// possible for the key type.
// return a string or nil if failed.
func (z *PrivateKey) GetPkcs1Pem() *string {

    retStr := C.CkPrivateKey_getPkcs1Pem(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets the private key in unencrypted PKCS8 format.
// 
// RSA keys are returned in PKCS8 ASN.1 DER format:
// SEQUENCE                  // PrivateKeyInfo
// +- INTEGER                // Version - 0 (v1998)
// +- SEQUENCE               // AlgorithmIdentifier
//    +- OID                 // 1.2.840.113549.1.1.1
//    +- NULL                // Optional Parameters
// +- OCTETSTRING            // PrivateKey
//    +- SEQUENCE            // RSAPrivateKey
//       +- INTEGER(0)       // Version - v1998(0)
//       +- INTEGER(N)       // N
//       +- INTEGER(E)       // E
//       +- INTEGER(D)       // D
//       +- INTEGER(P)       // P
//       +- INTEGER(Q)       // Q
//       +- INTEGER(DP)      // d mod p-1
//       +- INTEGER(DQ)      // d mod q-1
//       +- INTEGER(Inv Q)   // INV(q) mod p
// 
// DSA keys are returned in this ASN.1 DER format:
// SEQUENCE                 // PrivateKeyInfo
// +- INTEGER                 // Version
// +- SEQUENCE              // AlgorithmIdentifier
//     +- OID                       // 1.2.840.10040.4.1
//     +- SEQUENCE           // DSS-Params (Optional Parameters)
// 	+- INTEGER(P)      // P
// 	+- INTEGER(Q)      // Q
// 	+- INTEGER(G)      // G
//     +- OCTETSTRING             // PrivateKey
// 	+- INTEGER(X)      // DSAPrivateKey X
// 
// ECC keys are returned in this ASN.1 DER format:
// (from RFC 5915)
// ECPrivateKey ::= SEQUENCE {
//     version        INTEGER { ecPrivkeyVer1(1) } (ecPrivkeyVer1),
//     privateKey     OCTET STRING,
//     parameters [0] ECParameters {{ NamedCurve }} OPTIONAL,
//     publicKey  [1] BIT STRING OPTIONAL (This is the ANSI X9.63 public key format.)
//
func (z *PrivateKey) GetPkcs8() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkPrivateKey_GetPkcs8(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Gets the private key in unencrypted PKCS8 format and returned in an encoded
// string, as specified by the encoding argument.
// return a string or nil if failed.
func (z *PrivateKey) GetPkcs8ENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkPrivateKey_getPkcs8ENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Writes the private key to password-protected PKCS8 format. The Pkcs8EncryptAlg
// property controls the encryption algorithm used to encrypt.
func (z *PrivateKey) GetPkcs8Encrypted(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkPrivateKey_GetPkcs8Encrypted(z.ckObj, cstr1, ckOutBytes)

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


// Writes the private key to password-protected PKCS8 format and returns as an
// encoded string as specified by the encoding argument. The Pkcs8EncryptAlg property
// controls the encryption algorithm used to encrypt.
// return a string or nil if failed.
func (z *PrivateKey) GetPkcs8EncryptedENC(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkPrivateKey_getPkcs8EncryptedENC(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Writes the private key to password-protected PKCS8 PEM format. The
// Pkcs8EncryptAlg property controls the encryption algorithm used to encrypt.
// return a string or nil if failed.
func (z *PrivateKey) GetPkcs8EncryptedPem(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkPrivateKey_getPkcs8EncryptedPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets the private key in PKCS8 PEM format.
// return a string or nil if failed.
func (z *PrivateKey) GetPkcs8Pem() *string {

    retStr := C.CkPrivateKey_getPkcs8Pem(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the public key portion of the private key as a public key object.
func (z *PrivateKey) GetPublicKey() *PublicKey {

    retObj := C.CkPrivateKey_GetPublicKey(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &PublicKey{retObj}
}


// Returns the private key in raw hex format (lowercase). The public key is written
// to pubKey.
// 
// Ed25519 private and public keys are 32-byte each (64 chars in hex format).
// 
// The length of an EC key depends on the curve. The private key is a single hex
// string. The public key is a hex string composed of the "x" and "y" parts of the
// public key like this:
//     04||HEX(x)||HEX(y)
// 
// Note: This method is only applicable to Ed25519 and ECDSA keys. An RSA key
// cannot be returned in such as simple raw format because it is composed of
// multiple parts (modulus, exponent, and more).
//
// return a string or nil if failed.
func (z *PrivateKey) GetRawHex(arg1 *StringBuilder) *string {

    retStr := C.CkPrivateKey_getRawHex(z.ckObj, arg1.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the private key in XML format. The private key is returned unencrypted
// and the parts are base64 encoded.
// 
// RSA keys have this XML format:
//   ...  ...  
// 
// ...
// 
//   ...  ...  ...  ...  ...
// 
// DSA keys have this XML format:
// 
// ...
// 
// ............
// 
// ECC keys have this XML format. The CURVE_NAME could be one of secp256r1,
// secp384r1, secp521r1, secp256k1 (or others as new curves are supported.)
// ...
//
// return a string or nil if failed.
func (z *PrivateKey) GetXml() *string {

    retStr := C.CkPrivateKey_getXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads a private key from any format (PKCS1, PKCS8, PEM, JWK, PVK, etc.). The
// contents of the key (binary or text) is passed in privKeyData. The password is optional and
// should be specified if needed.
func (z *PrivateKey) LoadAnyFormat(arg1 *BinData, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkPrivateKey_LoadAnyFormat(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a private key from a file in any format (PKCS1, PKCS8, PEM, JWK, PVK,
// etc.). The password is optional and should be specified if needed.
func (z *PrivateKey) LoadAnyFormatFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPrivateKey_LoadAnyFormatFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads the private key object with an ed25519 key pair. The privKey is the 32-byte
// private key as a hex string. The pubKey is the 32-byte public key as a hex string.
// pubKey may be an empty string, in which case the public key is automatically
// computed from the private key.
func (z *PrivateKey) LoadEd25519(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPrivateKey_LoadEd25519(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads the private key from an in-memory encrypted PEM string. An encrypted PEM
// contains the private key in encrypted PKCS#8 format, where the data begins and
// ends with the following tags:
// -----BEGIN ENCRYPTED PRIVATE KEY-----
// BASE64 ENCODED DATA
// -----END ENCRYPTED PRIVATE KEY-----
// 
// For those requiring a deeper understanding: The base64 data contains ASN.1 DER
// with the following structure:
// EncryptedPrivateKeyInfo ::= SEQUENCE {
//   encryptionAlgorithm  EncryptionAlgorithmIdentifier,
//   encryptedData        EncryptedData
// }
// 
// EncryptionAlgorithmIdentifier ::= AlgorithmIdentifier
// 
// EncryptedData ::= OCTET STRING
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadEncryptedPem(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPrivateKey_LoadEncryptedPem(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a private key from an encrypted PEM file.
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadEncryptedPemFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPrivateKey_LoadEncryptedPemFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a private key from an JWK (JSON Web Key) string.
// 
// RSA keys have this JWK format:
//          {"kty":"RSA",
//           "n":"0vx7agoebGcQ ... JzKnqDKgw",
//           "e":"AQAB",
//           "d":"X4cTteJY_gn4F ... 4jfcKoAC8Q",
//           "p":"83i-7IvMGXoMX ... vn7O0nVbfs",
//           "q":"3dfOR9cuYq-0S ... 4vIcb6yelxk",
//           "dp":"G4sPXkc6Ya9 ... 8YeiKkTiBj0",
//           "dq":"s9lAH9fggBso ... w494Q_cgk",
//           "qi":"GyM_p6JrXySi ... zTKhAVRU"}
// 
// ECC keys have this JWK format.
//          {"kty":"EC",
//           "crv":"P-256",
//           "x":"MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4",
//           "y":"4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM",
//           "d":"870MB6gfuTJ4HtUnUvYMyJpr5eUZNP4Bk43bVdj3eAE"}
// 
// Ed25519 keys (added in v9.5.0.83) have this JWK format.
//          {"kty": "OKP",
//          "crv": "Ed25519",
//          "x": "SE2Kne5xt51z1eciMH2T2ftDQp96Gl6FhY6zSQujiP0",
//          "d": "O-eRXewadF0sNyB0U9omcnt8Qg2ZmeK3WSXPYgqe570",
//          "use": "sig"}
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadJwk(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_LoadJwk(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the private key from an in-memory PEM string. If the PEM contains an
// encrypted private key, then the LoadEncryptedPem method should instead be
// called. This method is for loading an unencrypted private key stored in PEM
// using PKCS#1 or PKCS#8.
// 
// A private key stored in PKCS#1 format begins and ends with the tags:
// -----BEGIN RSA PRIVATE KEY-----
// BASE64 ENCODED DATA
// -----END RSA PRIVATE KEY-----
// 
// For those requiring a deeper understanding, the PKCS1 base64 contains ASN.1 in
// DER encoding with the following structure:
// RSAPrivateKey ::= SEQUENCE {
//   version           Version,
//   modulus           INTEGER,  -- n
//   publicExponent    INTEGER,  -- e
//   privateExponent   INTEGER,  -- d
//   prime1            INTEGER,  -- p
//   prime2            INTEGER,  -- q
//   exponent1         INTEGER,  -- d mod (p-1)
//   exponent2         INTEGER,  -- d mod (q-1)
//   coefficient       INTEGER,  -- (inverse of q) mod p
//   otherPrimeInfos   OtherPrimeInfos OPTIONAL
// }
// 
// A private key stored in PKCS#8 format begins and ends with the tags:
// -----BEGIN PRIVATE KEY-----
// BASE64 ENCODED DATA
// -----END PRIVATE KEY-----
// 
// For those requiring a deeper understanding, the PKCS8 base64 contains ASN.1 in
// DER encoding with the following structure:
// PrivateKeyInfo ::= SEQUENCE {
//   version         Version,
//   algorithm       AlgorithmIdentifier,
//   PrivateKey      BIT STRING
// }
// 
// AlgorithmIdentifier ::= SEQUENCE {
//   algorithm       OBJECT IDENTIFIER,
//   parameters      ANY DEFINED BY algorithm OPTIONAL
// }
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadPem(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_LoadPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a private key from a PEM file.
func (z *PrivateKey) LoadPemFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_LoadPemFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads an RSA, ECC, or DSA private key from binary DER.
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadPkcs1(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkPrivateKey_LoadPkcs1(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Loads a private key from a PKCS1 file.
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadPkcs1File(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_LoadPkcs1File(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a private key from in-memory PKCS8 byte data.
// 
// For those requiring a deeper understanding, the PKCS8 contains ASN.1 in DER
// encoding with the following structure:
// PrivateKeyInfo ::= SEQUENCE {
//   version         Version,
//   algorithm       AlgorithmIdentifier,
//   PrivateKey      BIT STRING
// }
// 
// AlgorithmIdentifier ::= SEQUENCE {
//   algorithm       OBJECT IDENTIFIER,
//   parameters      ANY DEFINED BY algorithm OPTIONAL
// }
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadPkcs8(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkPrivateKey_LoadPkcs8(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Loads a private key from in-memory password-protected PKCS8 byte data.
// 
// For those requiring a deeper understanding, the encrypted PKCS8 contains ASN.1
// in DER encoding with the following structure:
// EncryptedPrivateKeyInfo ::= SEQUENCE {
//   encryptionAlgorithm  EncryptionAlgorithmIdentifier,
//   encryptedData        EncryptedData
// }
// 
// EncryptionAlgorithmIdentifier ::= AlgorithmIdentifier
// 
// EncryptedData ::= OCTET STRING
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadPkcs8Encrypted(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkPrivateKey_LoadPkcs8Encrypted(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a private key from an encrypted PKCS8 file.
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadPkcs8EncryptedFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPrivateKey_LoadPkcs8EncryptedFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a private key from a PKCS8 file.
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadPkcs8File(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_LoadPkcs8File(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a private key from an XML string.
// 
// RSA keys have this XML format:
//   ...  ...  
// 
// ...
// 
//   ...  ...  ...  ...  ...
// 
// DSA keys have this XML format:
// 
// ...
// 
// ............
// 
// ECC keys have this XML format. The CURVE_NAME could be one of secp256r1,
// secp384r1, secp521r1, secp256k1 (or others as new curves are supported.)
// ...
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_LoadXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a private key from an XML file.
// 
// Note: Each of the private key Load* methods willl auto-recognize the content and
// will parse appropriately. The private key should be successfully loaded even
// when the wrong format data is passed to the wrong method.
//
func (z *PrivateKey) LoadXmlFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_LoadXmlFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the private key to an unencrypted PKCS1 PEM format file.
func (z *PrivateKey) SavePemFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_SavePemFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the private key to an unencrypted binary PKCS1 format file.
func (z *PrivateKey) SavePkcs1File(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_SavePkcs1File(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the private key to a password-protected PKCS8 format file. The
// Pkcs8EncryptAlg property controls the encryption algorithm used to encrypt.
func (z *PrivateKey) SavePkcs8EncryptedFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPrivateKey_SavePkcs8EncryptedFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Saves the private key to a password-protected PKCS8 PEM format file. The
// Pkcs8EncryptAlg property controls the encryption algorithm used to encrypt.
func (z *PrivateKey) SavePkcs8EncryptedPemFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPrivateKey_SavePkcs8EncryptedPemFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Saves the private key to an unencrypted binary PKCS8 format file.
func (z *PrivateKey) SavePkcs8File(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_SavePkcs8File(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the private key to a PKCS8 PEM format file.
func (z *PrivateKey) SavePkcs8PemFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_SavePkcs8PemFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the private key to an XML file.
func (z *PrivateKey) SaveXmlFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrivateKey_SaveXmlFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


