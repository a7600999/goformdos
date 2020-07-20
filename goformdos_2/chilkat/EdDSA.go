// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkEdDSA.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewEdDSA() *EdDSA {
	return &EdDSA{C.CkEdDSA_Create()}
}

func (z *EdDSA) DisposeEdDSA() {
    if z != nil {
	C.CkEdDSA_Dispose(z.ckObj)
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
func (z *EdDSA) DebugLogFilePath() string {
    return C.GoString(C.CkEdDSA_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *EdDSA) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkEdDSA_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *EdDSA) LastErrorHtml() string {
    return C.GoString(C.CkEdDSA_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *EdDSA) LastErrorText() string {
    return C.GoString(C.CkEdDSA_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *EdDSA) LastErrorXml() string {
    return C.GoString(C.CkEdDSA_lastErrorXml(z.ckObj))
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
func (z *EdDSA) LastMethodSuccess() bool {
    v := int(C.CkEdDSA_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *EdDSA) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEdDSA_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *EdDSA) VerboseLogging() bool {
    v := int(C.CkEdDSA_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *EdDSA) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEdDSA_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *EdDSA) Version() string {
    return C.GoString(C.CkEdDSA_version(z.ckObj))
}
// Generates an Ed25519 key. privKey is an output argument. The generated key is
// created in privKey.
func (z *EdDSA) GenEd25519Key(arg1 *Prng, arg2 *PrivateKey) bool {

    v := C.CkEdDSA_GenEd25519Key(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Computes a shared secret given a private and public key. For example, Alice and
// Bob can compute the identical shared secret by doing the following: Alice sends
// Bob her public key, and Bob calls SharedSecretENC with his private key and
// Alice's public key. Bob sends Alice his public key, and Alice calls
// SharedSecretENC with her private key and Bob's public key. Both calls to
// SharedSecretENC will produce the same result. The resulting bytes are returned
// in encoded string form (hex, base64, etc) as specified by encoding.
// return a string or nil if failed.
func (z *EdDSA) SharedSecretENC(arg1 *PrivateKey, arg2 *PublicKey, arg3 string) *string {
    cstr3 := C.CString(arg3)

    retStr := C.CkEdDSA_sharedSecretENC(z.ckObj, arg1.ckObj, arg2.ckObj, cstr3)

    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Signs the contents of bd and returns the signature according to encoding. The
// encoding can be any encoding supported by Chilkat, such as "hex", "base64", etc.
// return a string or nil if failed.
func (z *EdDSA) SignBdENC(arg1 *BinData, arg2 string, arg3 *PrivateKey) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkEdDSA_signBdENC(z.ckObj, arg1.ckObj, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Verifies the signature against the contents of bd. The encodedSig is passed as an
// encoded string (such as hex, base64, etc.) using the encoding specified by enocding.
// The pubkey contains the Ed25519 public key used to verify.
func (z *EdDSA) VerifyBdENC(arg1 *BinData, arg2 string, arg3 string, arg4 *PublicKey) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkEdDSA_VerifyBdENC(z.ckObj, arg1.ckObj, cstr2, cstr3, arg4.ckObj)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


