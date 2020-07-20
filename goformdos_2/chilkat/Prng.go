// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkPrng.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewPrng() *Prng {
	return &Prng{C.CkPrng_Create()}
}

func (z *Prng) DisposePrng() {
    if z != nil {
	C.CkPrng_Dispose(z.ckObj)
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
func (z *Prng) DebugLogFilePath() string {
    return C.GoString(C.CkPrng_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Prng) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkPrng_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Prng) LastErrorHtml() string {
    return C.GoString(C.CkPrng_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Prng) LastErrorText() string {
    return C.GoString(C.CkPrng_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Prng) LastErrorXml() string {
    return C.GoString(C.CkPrng_lastErrorXml(z.ckObj))
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
func (z *Prng) LastMethodSuccess() bool {
    v := int(C.CkPrng_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Prng) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPrng_putLastMethodSuccess(z.ckObj,v)
}

// property: PrngName
// The name of the PRNG selected. Currently, the default and only possible value is
// "fortuna". See the links below for information about the Fortuna PRNG.
// 
// Note: Because "fortuna" is the only valid choice, assigning this property to a
// different value will always be ignored (until alternative PRNG algorithms are
// added in the future).
//
func (z *Prng) PrngName() string {
    return C.GoString(C.CkPrng_prngName(z.ckObj))
}
// property setter: PrngName
func (z *Prng) SetPrngName(goStr string) {
    cStr := C.CString(goStr)
    C.CkPrng_putPrngName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Prng) VerboseLogging() bool {
    v := int(C.CkPrng_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Prng) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPrng_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Prng) Version() string {
    return C.GoString(C.CkPrng_version(z.ckObj))
}
// Adds entropy to the PRNG (i.e. adds more seed material to the PRNG). Entropy can
// be obtained by calling GetEntropy, or the application might have its own sources
// for obtaining entropy. An application may continue to add entropy at desired
// intervals. How the entropy is used depends on the PRNG algorithm. For Fortuna,
// the entropy is added to the internal entropy pools and used when internal
// automatic reseeding occurs.
// 
// An application may add non-random entropy for testing purposes. This allows for
// the reproduction of the same pseudo-random number sequence for testing and
// debugging purposes.
// 
// The entropy bytes are passed in entropy using the binary encoding specified in
// encoding. Binary encodings can be "hex", "base64", etc. See the link below for
// supported binary encodings.
//
func (z *Prng) AddEntropy(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPrng_AddEntropy(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds entropy to the PRNG (i.e. adds more seed material to the PRNG). Entropy can
// be obtained by calling GetEntropy, or the application might have its own sources
// for obtaining entropy. An application may continue to add entropy at desired
// intervals. How the entropy is used depends on the PRNG algorithm. For Fortuna,
// the entropy is added to the internal entropy pools and used when internal
// automatic reseeding occurs.
// 
// An application may add non-random entropy for testing purposes. This allows for
// the reproduction of the same pseudo-random number sequence for testing and
// debugging purposes.
//
func (z *Prng) AddEntropyBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkPrng_AddEntropyBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Exports all accumulated entropy and returns it in a base64 encoded string.
// (Internally the entropy pools are re-hashed so that a hacker cannot determine
// the state of the PRNG even if the exported entropy was obtained.) When a system
// restarts it can import what was previously exported by calling ImportEntropy.
// This ensures an adequate amount of entropy is immediately available when first
// generating random bytes.
// 
// For example, an application could persist the exported entropy to a database or
// file. When the application starts again, it could import the persisted entropy,
// add some entropy from a system source (via the GetEntropy/AddEntropy methods),
// and then begin generating random data.
//
// return a string or nil if failed.
func (z *Prng) ExportEntropy() *string {

    retStr := C.CkPrng_exportEntropy(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Generates a random Firebase Push ID. SeeFirebase Unique Identifiers
// <https://www.firebase.com/blog/2015-02-11-firebase-unique-identifiers.html>.
// return a string or nil if failed.
func (z *Prng) FirebasePushId() *string {

    retStr := C.CkPrng_firebasePushId(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Generates and returns numBytes random bytes in encoded string form. The binary
// encoding is specified by encoding, and can be "hex", "base64", etc. (See the link
// below for supported binary encodings.)
// 
// Important: If no entropy was explicitly added prior to first call to generate
// random bytes, then 32 bytes of entropy (from the system source, such as
// /dev/random) are automatically added to seed the PRNG.
//
// return a string or nil if failed.
func (z *Prng) GenRandom(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkPrng_genRandom(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Appends numBytes random bytes to bd.
// 
// Important: If no entropy was explicitly added prior to first call to generate
// random bytes, then 32 bytes of entropy (from the system source, such as
// /dev/random) are automatically added to seed the PRNG.
//
func (z *Prng) GenRandomBd(arg1 int, arg2 *BinData) bool {

    v := C.CkPrng_GenRandomBd(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Generates and returns numBytes random bytes.
// 
// Important: If no entropy was explicitly added prior to first call to generate
// random bytes, then 32 bytes of entropy (from the system source, such as
// /dev/random) are automatically added to seed the PRNG.
//
func (z *Prng) GenRandomBytes(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkPrng_GenRandomBytes(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Reads real entropy bytes from a system entropy source and returns as an encoded
// string. On Linux/Unix based systems, including MAC OS X, this is accomplished by
// reading /dev/random. On Windows systems, it uses the Microsoft Cryptographic
// Service Provider's CryptGenRandom method.
// 
// It is recommended that no more than 32 bytes of entropy should be retrieved to
// initially seed a PRNG. Larger amounts of entropy are fairly useless. However, an
// app is free to periodically add bits of entropy to a long-running PRNG as it
// sees fit.
// 
// The encoding specifies the encoding to be used. It can be "hex", "base64", or many
// other possibilities. See the link below.
//
// return a string or nil if failed.
func (z *Prng) GetEntropy(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkPrng_getEntropy(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Reads and returns real entropy bytes from a system entropy source. On Linux/Unix
// based systems, including MAC OS X, this is accomplished by reading /dev/random.
// On Windows systems, it uses the Microsoft Cryptographic Service Provider's
// CryptGenRandom method.
// 
// It is recommended that no more than 32 bytes of entropy should be retrieved to
// initially seed a PRNG. Larger amounts of entropy are fairly useless. However, an
// app is free to periodically add bits of entropy to a long-running PRNG as it
// sees fit.
//
func (z *Prng) GetEntropyBytes(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkPrng_GetEntropyBytes(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Imports entropy from previously exported entropy. See the ExportEntropy method
// for more information.
func (z *Prng) ImportEntropy(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPrng_ImportEntropy(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Generates and returns a random integer between low and high (inclusive). For
// example, if low is 4 and high is 8, then random integers in the range 4, 5, 6,
// 7, 8 are returned.
func (z *Prng) RandomInt(arg1 int, arg2 int) int {

    v := C.CkPrng_RandomInt(z.ckObj, C.int(arg1), C.int(arg2))


    return int(v)
}


// Generates and returns a random password of a specified length. If mustIncludeDigit is
// true, the generated password will contain at least one digit (0-9). If upperAndLowercase is
// true, then generated password will contain both lowercase and uppercase
// USASCII chars (a-z and A-Z). If mustHaveOneOf is a non-empty string, it contains the set
// of non-alphanumeric characters, one of which must be included in the password.
// For example, mustHaveOneOf might be the string "!@#$%". If excludeChars is a non-empty string, it
// contains chars that should be excluded from the password. A typical need would
// be to exclude chars that appear similar to others, such as i, l, 1, L, o, 0, O.
// return a string or nil if failed.
func (z *Prng) RandomPassword(arg1 int, arg2 bool, arg3 bool, arg4 string, arg5 string) *string {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    retStr := C.CkPrng_randomPassword(z.ckObj, C.int(arg1), b2, b3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Generates and returns a random string that may contain digits (0-9), lowercase
// ASCII (a-z) , and uppercase ASCII (A-Z). To include numeric digits, set bDigits
// equal to true. To include lowercase ASCII, set bLower equal to true. To
// include uppercase ASCII, set bUpper equal to true. The length of the string to
// be generated is specified by length.
// return a string or nil if failed.
func (z *Prng) RandomString(arg1 int, arg2 bool, arg3 bool, arg4 bool) *string {
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

    retStr := C.CkPrng_randomString(z.ckObj, C.int(arg1), b2, b3, b4)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


