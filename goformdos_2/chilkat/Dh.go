// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkDh.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewDh() *Dh {
	return &Dh{C.CkDh_Create()}
}

func (z *Dh) DisposeDh() {
    if z != nil {
	C.CkDh_Dispose(z.ckObj)
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
func (z *Dh) DebugLogFilePath() string {
    return C.GoString(C.CkDh_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Dh) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkDh_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: G
// The generator. The value of G should be either 2 or 5.
func (z *Dh) G() int {
    return int(C.CkDh_getG(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Dh) LastErrorHtml() string {
    return C.GoString(C.CkDh_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Dh) LastErrorText() string {
    return C.GoString(C.CkDh_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Dh) LastErrorXml() string {
    return C.GoString(C.CkDh_lastErrorXml(z.ckObj))
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
func (z *Dh) LastMethodSuccess() bool {
    v := int(C.CkDh_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Dh) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDh_putLastMethodSuccess(z.ckObj,v)
}

// property: P
// A "safe" large prime returned as a hex string. The hex string represent a bignum
// in SSH1 format.
func (z *Dh) P() string {
    return C.GoString(C.CkDh_p(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Dh) VerboseLogging() bool {
    v := int(C.CkDh_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Dh) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDh_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Dh) Version() string {
    return C.GoString(C.CkDh_version(z.ckObj))
}
// The 1st step in Diffie-Hellman key exchange (to generate a shared-secret). The
// numBits should be twice the size (in bits) of the shared secret to be generated.
// For example, if you are using DH to create a 128-bit AES session key, then numBits
// should be set to 256. Returns E as a bignum in SSH-format as a hex string.
// return a string or nil if failed.
func (z *Dh) CreateE(arg1 int) *string {

    retStr := C.CkDh_createE(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The 2nd and final step in Diffie-Hellman (DH) key exchange. E is the E
// created by the other party. Returns the shared secret (K) as an SSH1-format
// bignum encoded as a hex string.
// return a string or nil if failed.
func (z *Dh) FindK(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkDh_findK(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Generates a large safe prime that is numBits bits in size using the generator G.
// Generating a new (random) P is expensive in both time and CPU cycles. A prime
// should be 1024 or more bits in length.
func (z *Dh) GenPG(arg1 int, arg2 int) bool {

    v := C.CkDh_GenPG(z.ckObj, C.int(arg1), C.int(arg2))


    return v != 0
}


// Sets explicit values for P and G. Returns true if P and G conform to the
// requirements for Diffie-Hellman. P is an SSH1-format bignum passed as a
// hexidecimalized string.
func (z *Dh) SetPG(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDh_SetPG(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unlocks the component. This must be called once prior to calling any other
// method. A fully-functional 30-day trial is automatically started when an
// arbitrary string is passed to this method. For example, passing "Hello", or
// "abc123" will unlock the component for the 1st thirty days after the initial
// install.
func (z *Dh) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDh_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets P and G to a known safe prime. The index may have the following values:
// 
// 1: First Oakley Default Group from RFC2409, section 6.1. Generator is 2. The
// prime is: 2^768 - 2 ^704 - 1 + 2^64 * { [2^638 pi] + 149686 }
// 
// 2: Prime for 2nd Oakley Group (RFC 2409) -- 1024-bit MODP Group. Generator is 2.
// The prime is: 2^1024 - 2^960 - 1 + 2^64 * { [2^894 pi] + 129093 }.
// 
// 3: 1536-bit MODP Group from RFC3526, Section 2. Generator is 2. The prime is:
// 2^1536 - 2^1472 - 1 + 2^64 * { [2^1406 pi] + 741804 }
// 
// 4: Prime for 14th Oakley Group (RFC 3526) -- 2048-bit MODP Group. Generator is
// 2. The prime is: 2^2048 - 2^1984 - 1 + 2^64 * { [2^1918 pi] + 124476 }
// 
// 5: 3072-bit MODP Group from RFC3526, Section 4. Generator is 2. The prime is:
// 2^3072 - 2^3008 - 1 + 2^64 * { [2^2942 pi] + 1690314 }
// 
// 6: 4096-bit MODP Group from RFC3526, Section 5. Generator is 2. The prime is:
// 2^4096 - 2^4032 - 1 + 2^64 * { [2^3966 pi] + 240904 }
// 
// 7: 6144-bit MODP Group from RFC3526, Section 6. Generator is 2. The prime is:
// 2^6144 - 2^6080 - 1 + 2^64 * { [2^6014 pi] + 929484 }
// 
// 8: 8192-bit MODP Group from RFC3526, Section 7. Generator is 2. The prime is:
// 2^8192 - 2^8128 - 1 + 2^64 * { [2^8062 pi] + 4743158 }
//
func (z *Dh) UseKnownPrime(arg1 int)  {

    C.CkDh_UseKnownPrime(z.ckObj, C.int(arg1))



}


