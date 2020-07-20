// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkJwe.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewJwe() *Jwe {
	return &Jwe{C.CkJwe_Create()}
}

func (z *Jwe) DisposeJwe() {
    if z != nil {
	C.CkJwe_Dispose(z.ckObj)
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
func (z *Jwe) DebugLogFilePath() string {
    return C.GoString(C.CkJwe_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Jwe) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkJwe_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Jwe) LastErrorHtml() string {
    return C.GoString(C.CkJwe_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Jwe) LastErrorText() string {
    return C.GoString(C.CkJwe_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Jwe) LastErrorXml() string {
    return C.GoString(C.CkJwe_lastErrorXml(z.ckObj))
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
func (z *Jwe) LastMethodSuccess() bool {
    v := int(C.CkJwe_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Jwe) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJwe_putLastMethodSuccess(z.ckObj,v)
}

// property: NumRecipients
// The number of recipients for this JWE.
func (z *Jwe) NumRecipients() int {
    return int(C.CkJwe_getNumRecipients(z.ckObj))
}

// property: PreferCompact
// Controls whether the JWE Compact Serialization or JWE JSON Serialization is
// preferred when creating JWEs. The default value is true, which is to use
// compact serialization when possible. If multiple recipients exist, or if any
// unprotected headers exist, then JWE JSON Serialization is used regardless of
// this property setting.
func (z *Jwe) PreferCompact() bool {
    v := int(C.CkJwe_getPreferCompact(z.ckObj))
    return v != 0
}
// property setter: PreferCompact
func (z *Jwe) SetPreferCompact(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJwe_putPreferCompact(z.ckObj,v)
}

// property: PreferFlattened
// Controls whether the flattened serialization is preferred when JWE JSON
// Serialization is used. The default value is true, which is to use the
// flattened serialization when possible. If multiple recipients exist, then the
// general (non-flattened) JWE JSON Serialization is used regardless of this
// property setting.
func (z *Jwe) PreferFlattened() bool {
    v := int(C.CkJwe_getPreferFlattened(z.ckObj))
    return v != 0
}
// property setter: PreferFlattened
func (z *Jwe) SetPreferFlattened(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJwe_putPreferFlattened(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Jwe) VerboseLogging() bool {
    v := int(C.CkJwe_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Jwe) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkJwe_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Jwe) Version() string {
    return C.GoString(C.CkJwe_version(z.ckObj))
}
// Decrypts a JWE and returns the original (decrypted) string content. The byte
// representation of the decrypted bytes is indicated by charset (such as "utf-8").
// (The charset tells Chilkat how to intepret the decrypted bytes as characters.)
// 
// The index specifies which recipient key is used for decryption. (Most JWEs have
// only a single recipent, and thus the index is typically 0.)
// 
// Supported Algorithms:
//     RSAES OAEP 256 (using SHA-256 and MGF1 with SHA-256) encryption with
//     A128CBC-HS256, A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     RSAES OAEP (using SHA-1 and MGF1 with SHA-1) encryption with A128CBC-HS256,
//     A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     RSAES-PKCS1-V1_5 encryption with A128CBC-HS256, A192CBC-HS384,
//     A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     Direct symmetric key encryption with pre-shared key A128CBC-HS256,
//     A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM and A256GCM
//     A128KW, A192KW, A256KW encryption with A128CBC-HS256, A192CBC-HS384,
//     A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     A128GCMKW, A192GCMKW, A256GCMKW encryption with A128CBC-HS256,
//     A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     PBES2-HS256+A128KW, PBES2-HS384+A192KW, PBES2-HS512+A256KW with
//     A128CBC-HS256, A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, A256GCM
//
// return a string or nil if failed.
func (z *Jwe) Decrypt(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkJwe_decrypt(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Decrypts the loaded JWE and appends the decrypted bytes to the contents of bd.
// The index specifies which recipient key is used for decryption. (Most JWEs have
// only a single recipent, and thus the index is typically 0.)
func (z *Jwe) DecryptBd(arg1 int, arg2 *BinData) bool {

    v := C.CkJwe_DecryptBd(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Decrypts the loaded JWE and appends the decrypted content to contentSb. The byte
// representation of the decrypted bytes is indicated by charset (such as "utf-8").
// (This tells Chilkat how to interpret the bytes as characters.)
// 
// The index specifies which recipient key is used for decryption. (Most JWEs have
// only a single recipent, and thus the index is typically 0.)
//
func (z *Jwe) DecryptSb(arg1 int, arg2 string, arg3 *StringBuilder) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJwe_DecryptSb(z.ckObj, C.int(arg1), cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Encrypts string content to produce a JWE. The byte representation of the content is
// indicated by charset (such as "utf-8").
// 
// Supported Algorithms:
//     RSAES OAEP 256 (using SHA-256 and MGF1 with SHA-256) encryption with
//     A128CBC-HS256, A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     RSAES OAEP (using SHA-1 and MGF1 with SHA-1) encryption with A128CBC-HS256,
//     A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     RSAES-PKCS1-V1_5 encryption with A128CBC-HS256, A192CBC-HS384,
//     A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     Direct symmetric key encryption with pre-shared key A128CBC-HS256,
//     A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM and A256GCM
//     A128KW, A192KW, A256KW encryption with A128CBC-HS256, A192CBC-HS384,
//     A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     A128GCMKW, A192GCMKW, A256GCMKW encryption with A128CBC-HS256,
//     A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, A256GCM
//     PBES2-HS256+A128KW, PBES2-HS384+A192KW, PBES2-HS512+A256KW with
//     A128CBC-HS256, A192CBC-HS384, A256CBC-HS512, A128GCM, A192GCM, A256GCM
//
// return a string or nil if failed.
func (z *Jwe) Encrypt(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkJwe_encrypt(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Encrypts the contents of contentBd to produce a JWE that is appended to the contents
// of jweSb. (This method provides the means to produce a JWE from binary content.)
func (z *Jwe) EncryptBd(arg1 *BinData, arg2 *StringBuilder) bool {

    v := C.CkJwe_EncryptBd(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Encrypts the contents of contentSb to produce a JWE that is appended to the contents
// of jweSb. The byte representation of the string to be encrypted is indicated by
// charset (such as "utf-8").
func (z *Jwe) EncryptSb(arg1 *StringBuilder, arg2 string, arg3 *StringBuilder) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJwe_EncryptSb(z.ckObj, arg1.ckObj, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Finds the index of the recipient with a header parameter (paramName) equal to a
// specified value (paramValue). Returns -1 if no recipient contains a header with the
// given name/value. If caseSensitive is true, then the header param name/value
// comparisons are case sensitive. Otherwise it is case insensitive.
// 
// The procedure for decrypting a JWE with multiple recipients is the following:
//     Load the JWE via one of the Load* methods.
//     Find the recipient index by some identifying header paramter. The typical
//     case is via the "kid" header parameter. ("kid" is an arbitrary key ID
//     applications can assign to identify keys.)
//     Set the key for decryption at the found index by calling SetPrivateKey,
//     SetWrappingKey, or SetPassword, depending on the type of key wrapping that is
//     employed.
//     Call Decrypt, DecryptSb, or DecryptBd to decrypt for the recipient (and key)
//     at the given index.
//
func (z *Jwe) FindRecipient(arg1 string, arg2 string, arg3 bool) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkJwe_FindRecipient(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Loads the contents of a JWE.
func (z *Jwe) LoadJwe(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkJwe_LoadJwe(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the contents of a JWE from a StringBuilder object.
func (z *Jwe) LoadJweSb(arg1 *StringBuilder) bool {

    v := C.CkJwe_LoadJweSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the optional Additional Authenticated Data. This is only used for
// non-compact serializations. The charset specifies the character encoding (such as
// "utf-8") to be used for the byte representation for the additional authenticated
// data.
func (z *Jwe) SetAad(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkJwe_SetAad(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the optional Additional Authenticated Data. This is only used for
// non-compact serializations. This method provides a way for binary (non-text)
// additional authenticated data to be used.
func (z *Jwe) SetAadBd(arg1 *BinData) bool {

    v := C.CkJwe_SetAadBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the PBES2 password for key encryption or decryption. This is for the case
// where the content encryption key (CEK) is encrypted using PBES2. An PBES2
// password should be used in the cases where the "alg" header parameter value is
// equal to one of the following:
// PBES2-HS256+A128KW
// PBES2-HS384+A192KW
// PBES2-HS512+A256KW
// The index is the index of the recipient, where the 1st recipient is at index 0.
// (The typical use case for JWEs is for a single recipient.)
func (z *Jwe) SetPassword(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkJwe_SetPassword(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets a private key for RSA key unwrapping/decryption. This is for the case where
// the content encryption key (CEK) is encrypted using RSA. An RSA private key
// should be used for decrypting in the cases where the "alg" header parameter
// value is equal to one of the following:
// RSA1_5
// RSA-OAEP
// RSA-OAEP-256
// RSA-OAEP-384  (added in Chilkat v9.5.0.71)
// RSA-OAEP-512  (added in Chilkat v9.5.0.71)
// The index is the index of the recipient, where the 1st recipient is at index 0.
// (The typical use case for JWEs is for a single recipient.)
func (z *Jwe) SetPrivateKey(arg1 int, arg2 *PrivateKey) bool {

    v := C.CkJwe_SetPrivateKey(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Sets the JWE Protected Header.
func (z *Jwe) SetProtectedHeader(arg1 *JsonObject) bool {

    v := C.CkJwe_SetProtectedHeader(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets a public key for RSA key wrapping encryption. This is for the case where
// the content encryption key (CEK) is encrypted using RSA. An RSA public key
// should be used when encrypting for the cases where the "alg" header parameter
// value is equal to one of the following:
// RSA1_5
// RSA-OAEP
// RSA-OAEP-256
// The index is the index of the recipient, where the 1st recipient is at index 0.
// (The typical use case for JWEs is for a single recipient.)
func (z *Jwe) SetPublicKey(arg1 int, arg2 *PublicKey) bool {

    v := C.CkJwe_SetPublicKey(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Sets a per-recipient unprotected header. This method would only be called if the
// JWE is for multiple recipients. The 1st recipient is at index 0.
func (z *Jwe) SetRecipientHeader(arg1 int, arg2 *JsonObject) bool {

    v := C.CkJwe_SetRecipientHeader(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Sets the JWE Shared Unprotected Header.
func (z *Jwe) SetUnprotectedHeader(arg1 *JsonObject) bool {

    v := C.CkJwe_SetUnprotectedHeader(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the AES wrapping key for encryption or decryption. This is for the case
// where the content encryption key (CEK) is encrypted using AES Key Wrap or AES
// GCM. An AES key should be used in the cases where the "alg" header parameter
// value is equal to one of the following:
// A128KW
// A192KW
// A256KW
// A128GCMKW
// A192GCMKW
// A256GCMKW
// dir
// The index is the index of the recipient, where the 1st recipient is at index 0.
// (The typical use case for JWEs is for a single recipient.)
// 
// Note: This method also sets the shared direct symmetric key for the case when
// the "alg" is equal to "dir". In this case, the key specified is not actualy a
// key encryption key, but is the direct content encryption key.
// 
// The encoding indicates the representation, such as "base64", "hex", "base64url",
// etc. of the encodedKey.
//
func (z *Jwe) SetWrappingKey(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkJwe_SetWrappingKey(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


