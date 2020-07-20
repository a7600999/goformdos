// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkJwt.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewJwt() *Jwt {
	return &Jwt{C.CkJwt_Create()}
}

func (z *Jwt) DisposeJwt() {
    if z != nil {
	C.CkJwt_Dispose(z.ckObj)
	}
}




// property: AutoCompact
// If true, the JSON passed to CreateJwt and CreateJwtPk will be compacted to
// remove unnecessary whitespace. This will result in the smallest possible JWT.
// The default value is true.
func (z *Jwt) AutoCompact() bool {
    v := int(C.CkJwt_getAutoCompact(z.ckObj))
    return v != 0
}
// property setter: AutoCompact
func (z *Jwt) SetAutoCompact(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJwt_putAutoCompact(z.ckObj,v)
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
func (z *Jwt) DebugLogFilePath() string {
    return C.GoString(C.CkJwt_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Jwt) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkJwt_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Jwt) LastErrorHtml() string {
    return C.GoString(C.CkJwt_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Jwt) LastErrorText() string {
    return C.GoString(C.CkJwt_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Jwt) LastErrorXml() string {
    return C.GoString(C.CkJwt_lastErrorXml(z.ckObj))
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
func (z *Jwt) LastMethodSuccess() bool {
    v := int(C.CkJwt_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Jwt) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJwt_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Jwt) VerboseLogging() bool {
    v := int(C.CkJwt_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Jwt) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJwt_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Jwt) Version() string {
    return C.GoString(C.CkJwt_version(z.ckObj))
}
// Creates a JWT. The header is the JOSE JSON header. It can be the full JOSE JSON,
// or it can be a shorthand string such as "HS256", "HS384", or "HS512", in which
// case the standard JOSE header for the given algorithm will be used.
// 
// The payload is the JSON payload that contains the claims. The password is the secret.
// Given that the secret is a shared passwod string, this method should only be
// called for creating JWT's where the JOSE header's "alg" is HS256, HS384, or
// HS512. For RS256, RS384, RS512, ES256, ES384, and ES512, call CreateJwtPk
// instead.
// 
// When successful, this method returns a JWT with the format xxxxx.yyyyy.zzzzz,
// where xxxxx is the base64url encoded JOSE header, yyyyy is the base64url encoded
// payload, and zzzzz is the base64url signature.
//
// return a string or nil if failed.
func (z *Jwt) CreateJwt(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkJwt_createJwt(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Creates a JWT using an RSA or ECC private key. The header is the JOSE JSON header.
// It can be the full JOSE JSON, or it can be a shorthand string such as "RS256",
// "RS384", "RS512", "ES256", "ES384", or "ES512", in which case the standard JOSE
// header for the given algorithm will be used.
// 
// The payload is the JSON payload that contains the claims. The key is the private
// key. This method should only be called for creating JWT's where the JOSE
// header's "alg" is RS256, RS384, RS512, ES256, ES384, and ES512. If the secret is
// a shared password string, then call CreateJwt instead.
// 
// When successful, this method returns a JWT with the format xxxxx.yyyyy.zzzzz,
// where xxxxx is the base64url encoded JOSE header, yyyyy is the base64url encoded
// payload, and zzzzz is the base64url signature.
//
// return a string or nil if failed.
func (z *Jwt) CreateJwtPk(arg1 string, arg2 string, arg3 *PrivateKey) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkJwt_createJwtPk(z.ckObj, cstr1, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Generates a JSON numeric value representing the number of seconds from
// 1970-01-01T00:00:00Z UTC until the specified UTC date/time, ignoring leap
// seconds. The date/time generated is equal to the current system time plus the
// number of seconds specified by numSecOffset. The numSecOffset can be negative.
func (z *Jwt) GenNumericDate(arg1 int) int {

    v := C.CkJwt_GenNumericDate(z.ckObj, C.int(arg1))


    return int(v)
}


// Decodes the first part of a JWT (the "xxxxx" part of the "xxxxx.yyyyy.zzzzz"
// JWT) and returns the JSON string. This is the JOSE header of the JWT.
// return a string or nil if failed.
func (z *Jwt) GetHeader(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkJwt_getHeader(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Decodes the second part of a JWT (the "yyyyy" part of the "xxxxx.yyyyy.zzzzz"
// JWT) and returns the JSON string. This is the claims payload of the JWT.
// return a string or nil if failed.
func (z *Jwt) GetPayload(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkJwt_getPayload(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Verifies the "exp" and/or "nbf" claims and returns true if the current system
// date/time is within range. Returns false if the current system date/time is
// outside the allowed range of time. The leeway may be set to a non-zero number of
// seconds to allow for some small leeway (usually no more than a few minutes) to
// account for clock skew.
func (z *Jwt) IsTimeValid(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJwt_IsTimeValid(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Verifies a JWT that requires a shared password string for verification. The token
// should be a JWT with the format xxxxx.yyyyy.zzzzz. This method should only be
// called for JWT's using the HS256, HS384, or HS512 algorithms. The VerifyJwtPk
// method should be called for verifying JWT's requiring an RSA or ECC key.
// 
// Returns true if the signature was verified. Returns false if the signature
// was not successfully verified.
// 
// Note: This method will return false if the "alg" in the JOSE header is
// anything other than the algorithms specifically for HMAC, namely "hs256,
// "hs384", and "hs512". For example, if the "alg" is "none", then this method
// immediately returns a failed status.
// 
// Further Explanation: This method calculates the signature using the password
// provided by the application, and compares it against the signature found in the
// JWT. If the signatures are equal, then the password is correct, and true is
// returned.
//
func (z *Jwt) VerifyJwt(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJwt_VerifyJwt(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Verifies a JWT that requires an RSA or ECC public key for verification. The token
// should be a JWT with the format xxxxx.yyyyy.zzzzz. This method should only be
// called for JWT's using the RS256, RS384, RS512, ES256, ES384, or ES512
// algorithms.
// 
// Returns true if the signature was verified. Returns false if the signature
// was not successfully verified.
// 
// Note: This method will return false if the "alg" in the JOSE header is
// anything other than the algorithms specifically for RSA and ECC. For example, if
// the "alg" is "none", then this method immediately returns a failed status.
// 
// Further Explanation: This method calculates the signature using the key
// provided by the application, and compares it against the signature found in the
// JWT. If the signatures are equal, then the key corresponds to the private key
// used to sign, and true is returned.
//
func (z *Jwt) VerifyJwtPk(arg1 string, arg2 *PublicKey) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJwt_VerifyJwtPk(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


