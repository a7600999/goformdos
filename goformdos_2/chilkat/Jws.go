// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkJws.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewJws() *Jws {
	return &Jws{C.CkJws_Create()}
}

func (z *Jws) DisposeJws() {
    if z != nil {
	C.CkJws_Dispose(z.ckObj)
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
func (z *Jws) DebugLogFilePath() string {
    return C.GoString(C.CkJws_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Jws) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkJws_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Jws) LastErrorHtml() string {
    return C.GoString(C.CkJws_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Jws) LastErrorText() string {
    return C.GoString(C.CkJws_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Jws) LastErrorXml() string {
    return C.GoString(C.CkJws_lastErrorXml(z.ckObj))
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
func (z *Jws) LastMethodSuccess() bool {
    v := int(C.CkJws_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Jws) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJws_putLastMethodSuccess(z.ckObj,v)
}

// property: NumSignatures
// The number of signatures contained in this JWS.
func (z *Jws) NumSignatures() int {
    return int(C.CkJws_getNumSignatures(z.ckObj))
}

// property: PreferCompact
// Controls whether to use the JWS Compact Serialization or JWS JSON Serialization
// when creating JWSs. The default value is true, which is to choose the compact
// serialization when possible. If multiple signatures exist, or if any unprotected
// headers exist, then JWS JSON Serialization is used regardless of this property
// setting.
func (z *Jws) PreferCompact() bool {
    v := int(C.CkJws_getPreferCompact(z.ckObj))
    return v != 0
}
// property setter: PreferCompact
func (z *Jws) SetPreferCompact(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJws_putPreferCompact(z.ckObj,v)
}

// property: PreferFlattened
// Controls whether to use the flattened or general JWE JSON Serialization when
// creating JWSs. The default value is true, which is to choose the flattened
// serialization when possible. If multiple signatures exist, then the general
// (non-flattened) JWS JSON Serialization is used regardless of this property
// setting.
func (z *Jws) PreferFlattened() bool {
    v := int(C.CkJws_getPreferFlattened(z.ckObj))
    return v != 0
}
// property setter: PreferFlattened
func (z *Jws) SetPreferFlattened(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJws_putPreferFlattened(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Jws) VerboseLogging() bool {
    v := int(C.CkJws_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Jws) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJws_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Jws) Version() string {
    return C.GoString(C.CkJws_version(z.ckObj))
}
// Creates and returns the JWS containing one or more signatures according to the
// previously set payload, headers, and keys.
// return a string or nil if failed.
func (z *Jws) CreateJws() *string {

    retStr := C.CkJws_createJws(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Creates and returns the JWS by appending it to the contents of sbJws.
func (z *Jws) CreateJwsSb(arg1 *StringBuilder) bool {

    v := C.CkJws_CreateJwsSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the JWS payload. The charset specifies the byte representation to be used
// in interpreting the bytes of the payload. (For example, "utf-8", "windows-1252",
// "utf-16", etc.)
// return a string or nil if failed.
func (z *Jws) GetPayload(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkJws_getPayload(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the JWS payload by appending it to the binData.
func (z *Jws) GetPayloadBd(arg1 *BinData) bool {

    v := C.CkJws_GetPayloadBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the JWS payload by appending it to the sbPayload. The charset specifies the byte
// representation to be used in interpreting the bytes of the payload. (For
// example, "utf-8", "windows-1252", "utf-16", etc.)
func (z *Jws) GetPayloadSb(arg1 string, arg2 *StringBuilder) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJws_GetPayloadSb(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the protected header for the Nth signature. The index is the index of the
// signature. The 1st signature is at index 0. In most cases, a JWS contains a
// single signature at index 0.
func (z *Jws) GetProtectedHeader(arg1 int) *JsonObject {

    retObj := C.CkJws_GetProtectedHeader(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Returns the optional unprotected header for the Nth signature. The index is the
// index of the signature. The 1st signature is at index 0. In most cases, a JWS
// contains a single signature at index 0.
func (z *Jws) GetUnprotectedHeader(arg1 int) *JsonObject {

    retObj := C.CkJws_GetUnprotectedHeader(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Loads a JWS contained in jwsStr.
func (z *Jws) LoadJws(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJws_LoadJws(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a JWS from the contents of sbJws.
func (z *Jws) LoadJwsSb(arg1 *StringBuilder) bool {

    v := C.CkJws_LoadJwsSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the MAC key to be used for creating or validating a signature for the
// algorithms "HS256", "HS384", or "HS512". The key is an encoded string
// representation of the MAC key bytes. The encoding is an encoding, such as
// "base64url", "base64", "hex", etc. The 1st signature is at index 0. (Most JWSs
// only contain a single signature.)
func (z *Jws) SetMacKey(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkJws_SetMacKey(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Sets the MAC key to be used for creating or validating a signature for the
// algorithms "HS256", "HS384", or "HS512". The key contains the binary bytes of
// the MAC key for the Nth signature. The 1st signature is at index 0.
func (z *Jws) SetMacKeyBd(arg1 int, arg2 *BinData) bool {

    v := C.CkJws_SetMacKeyBd(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Sets the JWS payload from the contents of payload. The charset specifies the byte
// representation to be used for the string that is to be the payload. (For
// example, "utf-8", "windows-1252", "utf-16", etc.). If includeBom is true, then the
// byte-order-mark (BOM), also known as preamble, if one exists for the given charset,
// is included in the payload. Normally, includeBom should be set to false.
func (z *Jws) SetPayload(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkJws_SetPayload(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the JWS payload from the contents of binData.
func (z *Jws) SetPayloadBd(arg1 *BinData) bool {

    v := C.CkJws_SetPayloadBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the JWS payload from the contents of sbPayload. The charset specifies the byte
// representation to be used for the string that is to be the payload. (For
// example, "utf-8", "windows-1252", "utf-16", etc.). If includeBom is true, then the
// byte-order-mark (BOM), also known as preamble, if one exists for the given charset,
// is included in the payload. Normally, includeBom should be set to false.
func (z *Jws) SetPayloadSb(arg1 *StringBuilder, arg2 string, arg3 bool) bool {
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkJws_SetPayloadSb(z.ckObj, arg1.ckObj, cstr2, b3)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the private key (ECC or RSA) to be used in creating a signature for the
// following algorithms:
// RS256
// RS384
// RS512
// ES256
// ES384
// ES512
// PS256
// PS384
// PS512
// The JWS algorithms are described in RFC 7518, section 3.1. Here is a summary:
//    +--------------+-------------------------------+--------------------+
//    | "alg" Param  | Digital Signature or MAC      | Implementation     |
//    | Value        | Algorithm                     | Requirements       |
//    +--------------+-------------------------------+--------------------+
//    | HS256        | HMAC using SHA-256            | Required           |
//    | HS384        | HMAC using SHA-384            | Optional           |
//    | HS512        | HMAC using SHA-512            | Optional           |
//    | RS256        | RSASSA-PKCS1-v1_5 using       | Recommended        |
//    |              | SHA-256                       |                    |
//    | RS384        | RSASSA-PKCS1-v1_5 using       | Optional           |
//    |              | SHA-384                       |                    |
//    | RS512        | RSASSA-PKCS1-v1_5 using       | Optional           |
//    |              | SHA-512                       |                    |
//    | ES256        | ECDSA using P-256 and SHA-256 | Recommended+       |
//    | ES384        | ECDSA using P-384 and SHA-384 | Optional           |
//    | ES512        | ECDSA using P-521 and SHA-512 | Optional           |
//    | PS256        | RSASSA-PSS using SHA-256 and  | Optional           |
//    |              | MGF1 with SHA-256             |                    |
//    | PS384        | RSASSA-PSS using SHA-384 and  | Optional           |
//    |              | MGF1 with SHA-384             |                    |
//    | PS512        | RSASSA-PSS using SHA-512 and  | Optional           |
//    |              | MGF1 with SHA-512             |                    |
//    | none         | No digital signature or MAC   | Optional           |
//    |              | performed                     |                    |
//    +--------------+-------------------------------+--------------------+
func (z *Jws) SetPrivateKey(arg1 int, arg2 *PrivateKey) bool {

    v := C.CkJws_SetPrivateKey(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Sets a signature's protected header. The index is the index of the signature. The
// 1st signature is at index 0. In most cases, a JWS contains a single signature at
// index 0.
func (z *Jws) SetProtectedHeader(arg1 int, arg2 *JsonObject) bool {

    v := C.CkJws_SetProtectedHeader(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Sets the public key (ECC or RSA) to be used in validating a signature for the
// following algorithms:
// RS256
// RS384
// RS512
// ES256
// ES384
// ES512
// PS256
// PS384
// PS512
func (z *Jws) SetPublicKey(arg1 int, arg2 *PublicKey) bool {

    v := C.CkJws_SetPublicKey(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Sets a signature's optional unprotected header. The index is the index of the
// signature. The 1st signature is at index 0. In most cases, a JWS contains a
// single signature at index 0.
func (z *Jws) SetUnprotectedHeader(arg1 int, arg2 *JsonObject) bool {

    v := C.CkJws_SetUnprotectedHeader(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Validates the Nth signature or MAC using the Nth public key or MAC key.
// 
// Returns 1 if the signature was validated, thus verifying that the signer used
// the same MAC key or the matching private key.
// Returns 0 if the incorrect public key or MAC key was provided.
// Returns -1 for any other error, such as if no public key or MAC key was
// previously set for the given index.
// 
// Prior to calling this method, a program should provide the public key or MAC key
// via the SetPublicKey, SetMacKey, or SetMacKeyBd methods.
//
func (z *Jws) Validate(arg1 int) int {

    v := C.CkJws_Validate(z.ckObj, C.int(arg1))


    return int(v)
}


