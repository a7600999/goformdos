// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkPublicKey.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewPublicKey() *PublicKey {
	return &PublicKey{C.CkPublicKey_Create()}
}

func (z *PublicKey) DisposePublicKey() {
    if z != nil {
	C.CkPublicKey_Dispose(z.ckObj)
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
func (z *PublicKey) DebugLogFilePath() string {
    return C.GoString(C.CkPublicKey_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *PublicKey) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkPublicKey_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: KeySize
// Gets the size (in bits) of the public key. For example: 1024, 2048, etc.
func (z *PublicKey) KeySize() int {
    return int(C.CkPublicKey_getKeySize(z.ckObj))
}

// property: KeyType
// The type of public key. Can be "empty", "rsa", "dsa", or "ecc".
func (z *PublicKey) KeyType() string {
    return C.GoString(C.CkPublicKey_keyType(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *PublicKey) LastErrorHtml() string {
    return C.GoString(C.CkPublicKey_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *PublicKey) LastErrorText() string {
    return C.GoString(C.CkPublicKey_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *PublicKey) LastErrorXml() string {
    return C.GoString(C.CkPublicKey_lastErrorXml(z.ckObj))
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
func (z *PublicKey) LastMethodSuccess() bool {
    v := int(C.CkPublicKey_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *PublicKey) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPublicKey_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *PublicKey) VerboseLogging() bool {
    v := int(C.CkPublicKey_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *PublicKey) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPublicKey_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *PublicKey) Version() string {
    return C.GoString(C.CkPublicKey_version(z.ckObj))
}
// Returns the public key in binary DER format. If the key type (such as RSA)
// supports both PKCS1 and PKCS8 formats, then preferPkcs1 determine which format is
// returned.
func (z *PublicKey) GetDer(arg1 bool) []byte {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    ckOutBytes := C.CkByteData_Create()

    v := C.CkPublicKey_GetDer(z.ckObj, b1, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the public key in DER format as an encoded string (such as base64 or
// hex). If the key type (such as RSA) supports both PKCS1 and PKCS8 formats, then
// preferPkcs1 determine which format is returned. The encoding specifies the encoding, which
// is typically "base64".
// return a string or nil if failed.
func (z *PublicKey) GetEncoded(arg1 bool, arg2 string) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr2 := C.CString(arg2)

    retStr := C.CkPublicKey_getEncoded(z.ckObj, b1, cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets the public key in JWK (JSON Web Key) format.
// 
// RSA public keys have this JWK format:
//          {"kty":"RSA",
//           "n": "0vx7agoebGcQSuuPiLJXZptN9 ... U8awapJzKnqDKgw",
//           "e":"AQAB"}
// 
// ECC public keys have this JWK format:
//          {"kty":"EC",
//           "crv":"P-256",
//           "x":"MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4",
//           "y":"4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM"}
// 
// Ed25519 public keys (added in v9.5.0.83) have this JWK format:
//          {"kty":"OKP",
//           "crv":"Ed25519",
//           "x": "SE2Kne5xt51z1eciMH2T2ftDQp96Gl6FhY6zSQujiP0"}
//
// return a string or nil if failed.
func (z *PublicKey) GetJwk() *string {

    retStr := C.CkPublicKey_getJwk(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the JWK thumbprint for the public key. This is the thumbprint of the
// JSON Web Key (JWK) as per RFC 7638.
// return a string or nil if failed.
func (z *PublicKey) GetJwkThumbprint(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkPublicKey_getJwkThumbprint(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the public key in PEM format. If the key type (such as RSA) supports
// both PKCS1 and PKCS8 formats, then preferPkcs1 determine which format is returned.
// return a string or nil if failed.
func (z *PublicKey) GetPem(arg1 bool) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    retStr := C.CkPublicKey_getPem(z.ckObj, b1)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets the public key in XML format. The format depends on the key type. The key
// parts indicated by "..." are base64 encoded.
// 
// RSA public keys have this XML format:
//   ...  ...
// 
// DSA public keys have this XML format:
// 
// ...
// 
// .........
// 
// ECC public keys have this XML format:
// ...
//
// return a string or nil if failed.
func (z *PublicKey) GetXml() *string {

    retStr := C.CkPublicKey_getXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads a public key from base64-encoded DER (can be PKCS1 or PKCS8).
func (z *PublicKey) LoadBase64(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPublicKey_LoadBase64(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a public key from any binary or string format where the data is contained
// in bd. Chilkat automatically recognizes the format and key type (RSA, EC, DSA,
// Ed25519, ..)
func (z *PublicKey) LoadBd(arg1 *BinData) bool {

    v := C.CkPublicKey_LoadBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads the public key object from a 32-byte ed25519 key specified as a hex
// string.
func (z *PublicKey) LoadEd25519(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPublicKey_LoadEd25519(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a public key from binary DER. Auto-recognizes both PKCS1 and PKCS8
// formats.
func (z *PublicKey) LoadFromBinary(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkPublicKey_LoadFromBinary(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Loads a public key from a file. The file can be in any string or binary format
// such as binary DER (PKCS1 or PKCS8), PEM, XML, or encoded binary DER (such as
// base64 encoded binary DER). The format of the contents of the file is
// auto-recognized.
// 
// Starting in version 9.5.0.66, this method also supports loading the JWK (JSON
// Web Key) format.
//
func (z *PublicKey) LoadFromFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPublicKey_LoadFromFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a public key from any string format, such as PEM, XML, or encoded binary
// DER (such as base64 encoded binary DER). The format of the keyString is
// auto-recognized.
// 
// Starting in version 9.5.0.66, this method also supports loading the JWK (JSON
// Web Key) format.
//
func (z *PublicKey) LoadFromString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPublicKey_LoadFromString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the public key to a file in binary DER format. If the key type (such as
// RSA) supports both PKCS1 and PKCS8 formats, then preferPkcs1 determine which format is
// returned.
func (z *PublicKey) SaveDerFile(arg1 bool, arg2 string) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr2 := C.CString(arg2)

    v := C.CkPublicKey_SaveDerFile(z.ckObj, b1, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Saves the public key to a file in PEM format. If the key type (such as RSA)
// supports both PKCS1 and PKCS8 formats, then preferPkcs1 determine which format is
// returned.
func (z *PublicKey) SavePemFile(arg1 bool, arg2 string) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr2 := C.CString(arg2)

    v := C.CkPublicKey_SavePemFile(z.ckObj, b1, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Saves the public key to an XML file.
func (z *PublicKey) SaveXmlFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPublicKey_SaveXmlFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


