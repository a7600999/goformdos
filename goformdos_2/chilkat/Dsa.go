// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkDsa.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewDsa() *Dsa {
	return &Dsa{C.CkDsa_Create()}
}

func (z *Dsa) DisposeDsa() {
    if z != nil {
	C.CkDsa_Dispose(z.ckObj)
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
func (z *Dsa) DebugLogFilePath() string {
    return C.GoString(C.CkDsa_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Dsa) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkDsa_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: GroupSize
// The group size (in bits) to be used in DSA key generation. The default value is
// 160 and corresponds to the fact that SHA-1 is used in signature generation. This
// property setting should not be changed. It exists for future expansion when
// additional underlying hash algorithms are supported.
func (z *Dsa) GroupSize() int {
    return int(C.CkDsa_getGroupSize(z.ckObj))
}
// property setter: GroupSize
func (z *Dsa) SetGroupSize(value int) {
    C.CkDsa_putGroupSize(z.ckObj,C.int(value))
}

// property: Hash
// The hash to be signed or verified via the SignHash and Verify methods. In both
// cases, the Hash should be set prior to calling SignHash or Verify. This property
// may also be set via the SetEncodedHash method.
func (z *Dsa) Hash() []byte {
    ckbyd := C.CkByteData_Create()
    C.CkDsa_getHash(z.ckObj,ckbyd)
    p := C.CkByteData_getData(ckbyd);
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckbyd)) )
    C.CkByteData_Dispose(ckbyd)
    return retBytes
}

// property setter: Hash
func (z *Dsa) SetHash(dataBytes []byte) {
    ckbyd := C.CkByteData_Create()
    pData := C.CBytes(dataBytes)
    C.CkByteData_borrowData(ckbyd, (*C.uchar)(pData), C.ulong(len(dataBytes)))
    C.CkDsa_putHash(z.ckObj,ckbyd)    
    C.free(pData)
    C.CkByteData_Dispose(ckbyd)
}


// property: HexG
// The "G" part of a public or private DSA key returned as a hex-encoded
// SSH1-format bignum. The "G" is the generator. DSA key params consist of G, P,
// and Q.
func (z *Dsa) HexG() string {
    return C.GoString(C.CkDsa_hexG(z.ckObj))
}

// property: HexP
// The "P" part of a public or private DSA key returned as a hex-encoded
// SSH1-format bignum. The "P" is a large prime. DSA key params consist of G, P,
// and Q.
func (z *Dsa) HexP() string {
    return C.GoString(C.CkDsa_hexP(z.ckObj))
}

// property: HexQ
// The "Q" part of a public or private DSA key returned as a hex-encoded
// SSH1-format bignum. DSA key params consist of G, P, and Q.
func (z *Dsa) HexQ() string {
    return C.GoString(C.CkDsa_hexQ(z.ckObj))
}

// property: HexX
// The "X" part of a DSA private key returned as a hex-encoded SSH1-format bignum.
func (z *Dsa) HexX() string {
    return C.GoString(C.CkDsa_hexX(z.ckObj))
}

// property: HexY
// The "Y" part of a DSA public key returned as a hex-encoded SSH1-format bignum.
// (The "Y" value is also accessible w/ a private key, but the "X" value is not
// available in a DSA public key.)
func (z *Dsa) HexY() string {
    return C.GoString(C.CkDsa_hexY(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Dsa) LastErrorHtml() string {
    return C.GoString(C.CkDsa_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Dsa) LastErrorText() string {
    return C.GoString(C.CkDsa_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Dsa) LastErrorXml() string {
    return C.GoString(C.CkDsa_lastErrorXml(z.ckObj))
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
func (z *Dsa) LastMethodSuccess() bool {
    v := int(C.CkDsa_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Dsa) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDsa_putLastMethodSuccess(z.ckObj,v)
}

// property: Signature
// The signature created by calling SignHash. If verifying a signature, this may be
// set prior to calling the Verify method.
func (z *Dsa) Signature() []byte {
    ckbyd := C.CkByteData_Create()
    C.CkDsa_getSignature(z.ckObj,ckbyd)
    p := C.CkByteData_getData(ckbyd);
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckbyd)) )
    C.CkByteData_Dispose(ckbyd)
    return retBytes
}

// property setter: Signature
func (z *Dsa) SetSignature(dataBytes []byte) {
    ckbyd := C.CkByteData_Create()
    pData := C.CBytes(dataBytes)
    C.CkByteData_borrowData(ckbyd, (*C.uchar)(pData), C.ulong(len(dataBytes)))
    C.CkDsa_putSignature(z.ckObj,ckbyd)    
    C.free(pData)
    C.CkByteData_Dispose(ckbyd)
}


// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Dsa) VerboseLogging() bool {
    v := int(C.CkDsa_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Dsa) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDsa_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Dsa) Version() string {
    return C.GoString(C.CkDsa_version(z.ckObj))
}
// Loads a DSA private key from in-memory DSA DER encoded bytes.
func (z *Dsa) FromDer(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkDsa_FromDer(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Loads a DSA private key from a DER-encoded file.
func (z *Dsa) FromDerFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_FromDerFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a DSA private key from an in-memory encrypted PEM-formatted string. If the
// PEM passed to this method is unencrypted, the password is ignored and the PEM is
// simply loaded.
func (z *Dsa) FromEncryptedPem(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkDsa_FromEncryptedPem(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a DSA private key from an in-memory unencrypted PEM-formatted string.
func (z *Dsa) FromPem(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_FromPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a DSA public key from an in-memory DER-encoded byte array.
func (z *Dsa) FromPublicDer(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkDsa_FromPublicDer(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Loads a DSA public key from a DER-encoded file.
func (z *Dsa) FromPublicDerFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_FromPublicDerFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a DSA public-key from an in-memory PEM string.
func (z *Dsa) FromPublicPem(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_FromPublicPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a DSA public or private key from an in-memory XML string.
func (z *Dsa) FromXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_FromXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Generates a new DSA key that is numBits bits in length. The numBits should be at least
// 1024 bits and a multiple of 64. Typical values are 1024 and 2048. The newly
// generated key may be exported by calling one of the To* methods.
func (z *Dsa) GenKey(arg1 int) bool {

    v := C.CkDsa_GenKey(z.ckObj, C.int(arg1))


    return v != 0
}


// Generates a new DSA key from in-memory DER parameters created by OpenSSL. The
// newly generated key may be exported by calling one of the To* methods.
func (z *Dsa) GenKeyFromParamsDer(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkDsa_GenKeyFromParamsDer(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Generates a new DSA key from a DER-format parameters file created by OpenSSL. An
// example of using OpenSSL to generate DSA parameters in DER format is:
// openssl dsaparam -outform DER 1024 dsaparam.der
// The newly generated key may be exported by calling one of the To* methods.
func (z *Dsa) GenKeyFromParamsDerFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_GenKeyFromParamsDerFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Generates a new DSA key from parameters (PEM formatted string) created by
// OpenSSL. The newly generated key may be exported by calling one of the To*
// methods.
func (z *Dsa) GenKeyFromParamsPem(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_GenKeyFromParamsPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Generates a new DSA key from a PEM parameters file created by OpenSSL. The newly
// generated key may be exported by calling one of the To* methods.
// 
// OpenSSL provides the ability to generate DSA key parameters. For example:
// openssl dsaparam 1024 dsaparam.pem
// 
// Here is a sample DSA parameters PEM:
// -----BEGIN DSA PARAMETERS-----
// MIIBHgKBgQCXIAx4XzLVZ5ZqOFzdsYWVyH/6E/mVPw4TgMZS6Wxajnbdn1/CUBzE
// RWTUp8SguTSDpjC1Q/nyno0G6Q96VoW+PUXv8qUph8vbSaEdsjYO/8jSfzkGfvsa
// cucr1ythfNyk63aZAKzxeutOmsVe77l6pZI96ROjWF5iizuUB4WgmwIVANxM70wH
// 8iPPYVzPZqtXCB66I2SnAoGAIbW2VYRjRdoA7trJgmnfWakghKwV1WyaYrotqeDE
// 07/dipp0cNuY0IAJgSmqLHlAkNa2ZNI/c1mNxcwhYzZrnn8CXIqrYmtI33w0PYCx
// KHPqj7puhddFwYS/rFiyWAN0jtCMHlfCVzFGbSzach5QQraPV9YApJXy+ORJ8VPU
// /zo=
// -----END DSA PARAMETERS-----
//
func (z *Dsa) GenKeyFromParamsPemFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_GenKeyFromParamsPemFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the bytes of the Hash property as a hex or base64 encoded string. The
// encoding should be set to either "hex" or "base64".
// return a string or nil if failed.
func (z *Dsa) GetEncodedHash(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkDsa_getEncodedHash(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the bytes of the Signature property as a hex or base64 encoded string.
// The encoding should be set to either "hex" or "base64". The Signature property is
// set when SignHash is called.
// return a string or nil if failed.
func (z *Dsa) GetEncodedSignature(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkDsa_getEncodedSignature(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Convenience method to load the entire contents of a text file into a string. It
// is assumed that the text contains ANSI encoded character data.
// return a string or nil if failed.
func (z *Dsa) LoadText(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkDsa_loadText(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Convenience method to save a string to a text file. The text is saved using the
// ANSI character encoding.
func (z *Dsa) SaveText(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkDsa_SaveText(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Provides a way to set the Hash property by passing an encoded string. The encoding
// can be "hex" or "base64". The encodedHash contains the encoded bytes of the hash that
// will be signed or verified via the SignHash and Verify methods.
func (z *Dsa) SetEncodedHash(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkDsa_SetEncodedHash(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the bytes of the Signature property. (The signature must be specified prior
// to calling the Verify method.) This method allows for the signature to be set
// via a hex or base64 encoded string. The encoding should be set to either "hex" or
// "base64".
func (z *Dsa) SetEncodedSignature(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkDsa_SetEncodedSignature(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the bytes of the Signature property by providing the R and S values in
// encoded form. (The signature must be specified prior to calling the Verify
// method.) The R and S values may be set via a hex or base64 encoded string. The
// encoding should be set to either "hex" or "base64".
func (z *Dsa) SetEncodedSignatureRS(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkDsa_SetEncodedSignatureRS(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Sets the DSA object's private key from explicitly provided pHex, qHex, gHex, and
// xHex values. The groupSizeInBytes specifies the group size (in bytes). It is typically equal
// to 20, which is the length of the underlying hash function (SHA-1) for signing.
// The pHex, qHex, gHex, and xHex values are hex-encoded SSH1-format bignums.
func (z *Dsa) SetKeyExplicit(arg1 int, arg2 string, arg3 string, arg4 string, arg5 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkDsa_SetKeyExplicit(z.ckObj, C.int(arg1), cstr2, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Sets the DSA object's private key from explicitly provided pHex, qHex, gHex, and
// yHex values. The groupSizeInBytes specifies the group size (in bytes). It is typically equal
// to 20, which is the length of the underlying hash function (SHA-1) for signing.
// The pHex, qHex, gHex, and yHex values are hex-encoded SSH1-format bignums.
func (z *Dsa) SetPubKeyExplicit(arg1 int, arg2 string, arg3 string, arg4 string, arg5 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkDsa_SetPubKeyExplicit(z.ckObj, C.int(arg1), cstr2, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Signs a hash using the digital signature algorithm. Before calling this method,
// set the hash to be signed by either calling SetEncodedHash or by setting the
// Hash property. If SignHash returns true, the signature may be retrieved by
// either calling GetEncodedHash, or by accessing the Signature property.
func (z *Dsa) SignHash() bool {

    v := C.CkDsa_SignHash(z.ckObj)


    return v != 0
}


// Writes the DSA private key to a DER-encoded byte array.
func (z *Dsa) ToDer() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkDsa_ToDer(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Writes the DSA private key to a DER-format file.
func (z *Dsa) ToDerFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_ToDerFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Writes the DSA private key to an in-memory encrypted PEM string.
// return a string or nil if failed.
func (z *Dsa) ToEncryptedPem(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkDsa_toEncryptedPem(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Writes the DSA private key to an in-memory PEM string.
// return a string or nil if failed.
func (z *Dsa) ToPem() *string {

    retStr := C.CkDsa_toPem(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Writes the DSA public key to a DER-encoded byte array.
func (z *Dsa) ToPublicDer() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkDsa_ToPublicDer(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Writes the DSA public key to a DER-format file.
func (z *Dsa) ToPublicDerFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_ToPublicDerFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Writes the DSA public key to an in-memory PEM string.
// return a string or nil if failed.
func (z *Dsa) ToPublicPem() *string {

    retStr := C.CkDsa_toPublicPem(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Writes the DSA private or public key to an in-memory XML string. The bPublicOnly
// determines whether the public or private key is written.
// return a string or nil if failed.
func (z *Dsa) ToXml(arg1 bool) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    retStr := C.CkDsa_toXml(z.ckObj, b1)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Unlocks the component. This must be called once prior to calling any other
// method.
func (z *Dsa) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDsa_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Verifies a DSA signature. The Hash and Signature properties must be set prior to
// calling this method. (These properties may also be set via the SetEncodedHash
// and SetEncodedSignature methods.)
func (z *Dsa) Verify() bool {

    v := C.CkDsa_Verify(z.ckObj)


    return v != 0
}


// Verifies that the public or private key contained in the calling Dsa object is
// valid. Returns true if valid, otherwise returns false.
func (z *Dsa) VerifyKey() bool {

    v := C.CkDsa_VerifyKey(z.ckObj)


    return v != 0
}


