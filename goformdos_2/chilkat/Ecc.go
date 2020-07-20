// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkEcc.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewEcc() *Ecc {
	return &Ecc{C.CkEcc_Create()}
}

func (z *Ecc) DisposeEcc() {
    if z != nil {
	C.CkEcc_Dispose(z.ckObj)
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
func (z *Ecc) DebugLogFilePath() string {
    return C.GoString(C.CkEcc_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Ecc) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkEcc_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Ecc) LastErrorHtml() string {
    return C.GoString(C.CkEcc_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Ecc) LastErrorText() string {
    return C.GoString(C.CkEcc_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Ecc) LastErrorXml() string {
    return C.GoString(C.CkEcc_lastErrorXml(z.ckObj))
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
func (z *Ecc) LastMethodSuccess() bool {
    v := int(C.CkEcc_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Ecc) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEcc_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Ecc) VerboseLogging() bool {
    v := int(C.CkEcc_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Ecc) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEcc_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Ecc) Version() string {
    return C.GoString(C.CkEcc_version(z.ckObj))
}
// Generates an ECC private key. The curveName specifies the curve name which determines
// the key size. The prng provides a source for generating the random private key.
// 
// The following curve names are accepted:
//     secp256r1 (also known as P-256 and prime256v1)
//     secp384r1 (also known as P-384)
//     secp521r1 (also known as P-521)
//     secp256k1 (This is the curve used for Bitcoin)
//
func (z *Ecc) GenEccKey(arg1 string, arg2 *Prng) *PrivateKey {
    cstr1 := C.CString(arg1)

    retObj := C.CkEcc_GenEccKey(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &PrivateKey{retObj}
}


// Generates an ECC private key using a specified value for K. The curveName specifies
// the curve name which determines the key size. The encodedK is the encoded value of
// the private key. The encoding is the encoding used for encodedK, which can be "hex",
// "base64", "decimal", etc.
// 
// Note: This method is typically used for testing -- such as when the same private
// key is desired to produce results identical from run to run.
// 
// The following curve names are accepted:
//     secp256r1 (also known as P-256 and prime256v1)
//     secp384r1 (also known as P-384)
//     secp521r1 (also known as P-521)
//     secp256k1 (This is the curve used for Bitcoin)
//
func (z *Ecc) GenEccKey2(arg1 string, arg2 string, arg3 string) *PrivateKey {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkEcc_GenEccKey2(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &PrivateKey{retObj}
}


// Computes a shared secret given a private and public key. For example, Alice and
// Bob can compute the identical shared secret by doing the following: Alice sends
// Bob her public key, and Bob calls SharedSecretENC with his private key and
// Alice's public key. Bob sends Alice his public key, and Alice calls
// SharedSecretENC with her private key and Bob's public key. Both calls to
// SharedSecretENC will produce the same result. The resulting bytes are returned
// in encoded string form (hex, base64, etc) as specified by encoding.
// 
// Note: The private and public keys must both be keys on the same ECC curve.
//
// return a string or nil if failed.
func (z *Ecc) SharedSecretENC(arg1 *PrivateKey, arg2 *PublicKey, arg3 string) *string {
    cstr3 := C.CString(arg3)

    retStr := C.CkEcc_sharedSecretENC(z.ckObj, arg1.ckObj, arg2.ckObj, cstr3)

    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Computes an ECC signature on a hash. ECC signatures are computed and verified on
// the hashes of data (such as SHA1, SHA256, etc.). The hash of the data is passed
// in encodedHash. The encoding, such as "base64", "hex", etc. is passed in encoding. The ECC
// private key is passed in the 3rd argument (privkey). Given that creating an ECC
// signature involves the generation of random numbers, a PRNG is passed in the 4th
// argument (prng). The signature is returned as an encoded string using the
// encoding specified by the encoding argument.
// return a string or nil if failed.
func (z *Ecc) SignHashENC(arg1 string, arg2 string, arg3 *PrivateKey, arg4 *Prng) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkEcc_signHashENC(z.ckObj, cstr1, cstr2, arg3.ckObj, arg4.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Verifies an ECC signature. ECC signatures are computed and verified on the
// hashes of data (such as SHA1, SHA256, etc.). The hash of the data is passed in
// encodedHash. The encoded signature is passed in encodedSig. The encoding of both the hash and
// signature, such as "base64", "hex", etc. is passed in encoding. The ECC public key
// is passed in the last argument (pubkey).
// 
// The method returns 1 for a valid signature, 0 for an invalid signature, and -1
// for any other failure.
//
func (z *Ecc) VerifyHashENC(arg1 string, arg2 string, arg3 string, arg4 *PublicKey) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkEcc_VerifyHashENC(z.ckObj, cstr1, cstr2, cstr3, arg4.ckObj)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return int(v)
}


