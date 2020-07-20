// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkRsa.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewRsa() *Rsa {
	return &Rsa{C.CkRsa_Create()}
}

func (z *Rsa) DisposeRsa() {
    if z != nil {
	C.CkRsa_Dispose(z.ckObj)
	}
}




// property: Charset
// This property only applies when encrypting, decrypting, signing, or verifying
// signatures for strings. When encrypting strings, the input string is first
// converted to this charset before encrypting.
// 
// When decrypting, the decrypted data is interpreted as a string with this charset
// encoding and converted to the appropriate return. For example, ActiveX's
// returning strings always return Unicode (2 bytes/char). Java strings are utf-8.
// Chilkat C++ strings are ANSI or utf-8. .NET strings are Unicode.
// 
// The default value of this property is the ANSI charset of the local computer.
// 
// When signing string data, the input string is first converted to this charset
// before being hashed and signed. When verifying the signature for string data,
// the input string is first converted to this charset before the verification
// process begins.
//
func (z *Rsa) Charset() string {
    return C.GoString(C.CkRsa_charset(z.ckObj))
}
// property setter: Charset
func (z *Rsa) SetCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkRsa_putCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
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
func (z *Rsa) DebugLogFilePath() string {
    return C.GoString(C.CkRsa_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Rsa) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkRsa_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EncodingMode
// Encoding mode to be used in methods ending in "ENC", such as EncryptStringENC.
// Valid EncodingModes are "base64", "hex", "url", or "quoted-printable" (or "qp").
// Encryption methods ending in "ENC" will return encrypted data as a string
// encoded according to this property's value. Decryption methods ending in "ENC"
// accept an encoded string as specified by this property. The string is first
// decoded and then decrypted. The default value is "base64".
// 
// This property also applies to the "ENC" methods for creating and verifying
// digital signatures.
//
func (z *Rsa) EncodingMode() string {
    return C.GoString(C.CkRsa_encodingMode(z.ckObj))
}
// property setter: EncodingMode
func (z *Rsa) SetEncodingMode(goStr string) {
    cStr := C.CString(goStr)
    C.CkRsa_putEncodingMode(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Rsa) LastErrorHtml() string {
    return C.GoString(C.CkRsa_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Rsa) LastErrorText() string {
    return C.GoString(C.CkRsa_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Rsa) LastErrorXml() string {
    return C.GoString(C.CkRsa_lastErrorXml(z.ckObj))
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
func (z *Rsa) LastMethodSuccess() bool {
    v := int(C.CkRsa_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Rsa) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRsa_putLastMethodSuccess(z.ckObj,v)
}

// property: LittleEndian
// The default value is false, which means that signatures and encrypted output
// will be created using the big endian byte ordering. A value of true will
// produce little-endian output, which is what Microsoft's Crypto API produces.
// 
// Important: Prior to v9.5.0.49, this property behaved the opposite as it should
// for encryption. When updating from an older version of Chilkat to v9.5.0.49 or
// greater, the following change is required:
//     If the application did NOT explicity set the LittleEndian property, then no
//     change is required for encryption/decryption. If signatures were being created
//     or verified, then explicitly set this property to true.
//     If the application explicitly set this property, then reverse the setting
//     ONLY if doing encryption/decryption. No changes are required if doing signature
//     creation/verification.
//
func (z *Rsa) LittleEndian() bool {
    v := int(C.CkRsa_getLittleEndian(z.ckObj))
    return v != 0
}
// property setter: LittleEndian
func (z *Rsa) SetLittleEndian(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRsa_putLittleEndian(z.ckObj,v)
}

// property: NoUnpad
// If true, skips unpadding when decrypting. The default is false. This
// property value is typically left unchanged.
func (z *Rsa) NoUnpad() bool {
    v := int(C.CkRsa_getNoUnpad(z.ckObj))
    return v != 0
}
// property setter: NoUnpad
func (z *Rsa) SetNoUnpad(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRsa_putNoUnpad(z.ckObj,v)
}

// property: NumBits
// The number of bits of the key generated or imported into this RSA encryption
// object. Keys ranging in size from 384 bits to 4096 bits can be generated by
// calling GenerateKey. A public or private key may be imported by calling
// ImportPublicKey or ImportPrivateKey. A key must be available either via
// GenerateKey or import before any of the encrypt/decrypt methods may be called.
func (z *Rsa) NumBits() int {
    return int(C.CkRsa_getNumBits(z.ckObj))
}

// property: OaepHash
// Selects the hash algorithm for use within OAEP padding. The valid choices are
// "sha1", "sha256", "sha384", "sha512", "md2", "md5", "haval", "ripemd128",
// "ripemd160","ripemd256", or "ripemd320". The default is "sha1".
func (z *Rsa) OaepHash() string {
    return C.GoString(C.CkRsa_oaepHash(z.ckObj))
}
// property setter: OaepHash
func (z *Rsa) SetOaepHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkRsa_putOaepHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OaepMgfHash
// Selects the MGF (mask generation) hash algorithm for use within OAEP padding.
// The valid choices are "sha1", "sha256", "sha384", "sha512", "md2", "md5",
// "haval", "ripemd128", "ripemd160","ripemd256", or "ripemd320". The default is
// "sha1".
func (z *Rsa) OaepMgfHash() string {
    return C.GoString(C.CkRsa_oaepMgfHash(z.ckObj))
}
// property setter: OaepMgfHash
func (z *Rsa) SetOaepMgfHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkRsa_putOaepMgfHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OaepPadding
// Controls whether Optimal Asymmetric Encryption Padding (OAEP) is used for the
// padding scheme (for encrypting/decrypting). If set to false, PKCS1 v1.5
// padding is used. If set to true, PKCS1 v2.0 (OAEP) padding is used.
// 
// Important: The OAEP padding algorithm uses randomly generated bytes. Therefore,
// the RSA result will be different each time, even if all of the other inputs are
// identical. For example, if you RSA encrypt or sign the same data using the same
// key 100 times, the output will appear different each time, but they are all
// valid.
// 
// When creating digital signatures, this property controls whether RSA-PSS or
// PKCS1 v1.5 is used. If true, then the RSA-PSS signature scheme is used. The
// default value of this property is false.
//
func (z *Rsa) OaepPadding() bool {
    v := int(C.CkRsa_getOaepPadding(z.ckObj))
    return v != 0
}
// property setter: OaepPadding
func (z *Rsa) SetOaepPadding(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRsa_putOaepPadding(z.ckObj,v)
}

// property: PssSaltLen
// Selects the PSS salt length when RSASSA-PSS padding is selected for signatures.
// The default value is -1 to indicate that the length of the hash function should
// be used. For example, if the hash function is SHA256, then the PSS salt length
// will be 32 bytes. Can be optionally set to a value such as 20 if a specific salt
// length is required. This property should normally remain at the default value.
func (z *Rsa) PssSaltLen() int {
    return int(C.CkRsa_getPssSaltLen(z.ckObj))
}
// property setter: PssSaltLen
func (z *Rsa) SetPssSaltLen(value int) {
    C.CkRsa_putPssSaltLen(z.ckObj,C.int(value))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Rsa) VerboseLogging() bool {
    v := int(C.CkRsa_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Rsa) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRsa_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Rsa) Version() string {
    return C.GoString(C.CkRsa_version(z.ckObj))
}
// RSA decrypts the contents of bd. usePrivateKey should be set to true if the private
// key is to be used for decrypting. Otherwise it should be set to false if the
// public key is to be used for decrypting.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
//
func (z *Rsa) DecryptBd(arg1 *BinData, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkRsa_DecryptBd(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Decrypts byte data using the RSA encryption algorithm. usePrivateKey should be set to
// true if the private key is to be used for decrypting. Otherwise it should be
// set to false if the public key is to be used for decrypting.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
//
func (z *Rsa) DecryptBytes(arg1 []byte, arg2 bool) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_DecryptBytes(z.ckObj, ckbyd1, b2, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Same as DecryptBytes, except the input is an encoded string. The encoding is
// specified by the EncodingMode property, which can have values such as "base64",
// "hex", "quoted-printable", "url", etc.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
//
func (z *Rsa) DecryptBytesENC(arg1 string, arg2 bool) []byte {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_DecryptBytesENC(z.ckObj, cstr1, b2, ckOutBytes)

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


// Decrypts encrypted string data and returns an unencrypted string. usePrivateKey should be
// set to true if the private key is to be used for decrypting. Otherwise it
// should be set to false if the public key is to be used. The Charset property
// controls how the component interprets the decrypted string. Depending on the
// programming language, strings are returned to the application as Unicode, utf-8,
// or ANSI. Internal to DecryptString, the decrypted string is automatically
// converted from the charset specified by the Charset property to the encoding
// required by the calling programming language.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
//
// return a string or nil if failed.
func (z *Rsa) DecryptString(arg1 []byte, arg2 bool) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkRsa_decryptString(z.ckObj, ckbyd1, b2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Same as DecryptString, except the input is an encoded string. The encoding is
// specified by the EncodingMode property, which can have values such as "base64",
// "hex", "quoted-printable", "url", etc.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
//
// return a string or nil if failed.
func (z *Rsa) DecryptStringENC(arg1 string, arg2 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkRsa_decryptStringENC(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// RSA encrypts the contents of bd. usePrivateKey should be set to true if the private
// key is to be used for encrypting. Otherwise it should be set to false if the
// public key is to be used for encrypting.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
// 
// Note: Except for special situations, the public key should always be used for
// encrypting, and the private key for decrypting. This makes sense because an
// encrypted message is sent to a recipient, and the recipient is the only one in
// possession of the private key, and therefore the only one that can decrypt and
// read the message.
//
func (z *Rsa) EncryptBd(arg1 *BinData, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkRsa_EncryptBd(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Encrypts byte data using the RSA encryption algorithm. usePrivateKey should be set to
// true if the private key is to be used for encrypting. Otherwise it should be
// set to false if the public key is to be used for encrypting.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
// 
// Note: Except for special situations, the public key should always be used for
// encrypting, and the private key for decrypting. This makes sense because an
// encrypted message is sent to a recipient, and the recipient is the only one in
// possession of the private key, and therefore the only one that can decrypt and
// read the message.
//
func (z *Rsa) EncryptBytes(arg1 []byte, arg2 bool) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_EncryptBytes(z.ckObj, ckbyd1, b2, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Same as EncryptBytes, except the output is an encoded string. The encoding is
// specified by the EncodingMode property, which can have values such as "base64",
// "hex", "quoted-printable", "url", etc.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
// 
// Note: Except for special situations, the public key should always be used for
// encrypting, and the private key for decrypting. This makes sense because an
// encrypted message is sent to a recipient, and the recipient is the only one in
// possession of the private key, and therefore the only one that can decrypt and
// read the message.
//
// return a string or nil if failed.
func (z *Rsa) EncryptBytesENC(arg1 []byte, arg2 bool) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkRsa_encryptBytesENC(z.ckObj, ckbyd1, b2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Encrypts a string using the RSA encryption algorithm. usePrivateKey should be set to
// true if the private key is to be used for encrypting. Otherwise it should be
// set to false if the public key is to be used for encrypting. The string is
// first converted (if necessary) to the character encoding specified by the
// Charset property before encrypting. The encrypted bytes are returned.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
// 
// Note: Except for special situations, the public key should always be used for
// encrypting, and the private key for decrypting. This makes sense because an
// encrypted message is sent to a recipient, and the recipient is the only one in
// possession of the private key, and therefore the only one that can decrypt and
// read the message.
//
func (z *Rsa) EncryptString(arg1 string, arg2 bool) []byte {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_EncryptString(z.ckObj, cstr1, b2, ckOutBytes)

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


// Same as EncryptString, except the output is an encoded string. The encoding is
// specified by the EncodingMode property, which can have values such as "base64",
// "hex", "quoted-printable", "url", etc.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false.
// 
// Note: Except for special situations, the public key should always be used for
// encrypting, and the private key for decrypting. This makes sense because an
// encrypted message is sent to a recipient, and the recipient is the only one in
// possession of the private key, and therefore the only one that can decrypt and
// read the message.
//
// return a string or nil if failed.
func (z *Rsa) EncryptStringENC(arg1 string, arg2 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkRsa_encryptStringENC(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Exports the private-key of an RSA key pair to XML format. This is typically
// called after generating a new RSA key via the GenerateKey method.
// return a string or nil if failed.
func (z *Rsa) ExportPrivateKey() *string {

    retStr := C.CkRsa_exportPrivateKey(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Exports the private-key to a private key object. This is typically called after
// generating a new RSA key via the GenerateKey method. Once the private key object
// is obtained, it may be saved in a variety of different formats.
func (z *Rsa) ExportPrivateKeyObj() *PrivateKey {

    retObj := C.CkRsa_ExportPrivateKeyObj(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &PrivateKey{retObj}
}


// Exports the public-key of an RSA key pair to XML format. This is typically
// called after generating a new RSA key via the GenerateKey method.
// return a string or nil if failed.
func (z *Rsa) ExportPublicKey() *string {

    retStr := C.CkRsa_exportPublicKey(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Exports the public key to a public key object. Once the public key object is
// obtained, it may be saved in a variety of different formats.
func (z *Rsa) ExportPublicKeyObj() *PublicKey {

    retObj := C.CkRsa_ExportPublicKeyObj(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &PublicKey{retObj}
}


// Generates a new RSA public/private key pair. The number of bits can range from
// 512 to 8192. Typical key lengths are 1024, 2048, or 4096 bits. After successful
// generation, the public/private parts of the key can be exported to XML via the
// ExportPrivateKey and ExportPublicKey methods.
// 
// Note: Prior to version 9.5.0.49, the max key size was 4096 bits. Generating an
// 8192-bit RSA key takes a considerable amount of time and CPU processing power.
// There are no event callbacks or progress monitoring for RSA key generation.
// Calling this will block the thread until it returns.
//
func (z *Rsa) GenerateKey(arg1 int) bool {

    v := C.CkRsa_GenerateKey(z.ckObj, C.int(arg1))


    return v != 0
}


// Imports a private key from XML format. After successful import, the private key
// can be used to encrypt or decrypt. A private key (by definition) contains both
// private and public parts. This is because the public key consist of modulus and
// exponent. The private key consists of modulus, exponent, P, Q, DP, DQ, InverseQ,
// and D using base64 representation:
// _LT_RSAKeyValue>
//   _LT_Modulus>..._LT_/Modulus>
//   _LT_Exponent>..._LT_/Exponent>
//   _LT_P>..._LT_/P>
//   _LT_Q>..._LT_/Q>
//   _LT_DP>..._LT_/DP>
//   _LT_DQ>..._LT_/DQ>
//   _LT_InverseQ>..._LT_/InverseQ>
//   _LT_D>..._LT_/D>
// _LT_/RSAKeyValue>
// 
// Important: The Rsa object can contain either a private key or a public key, but
// not both. Importing a private key overwrites the existing key regardless of
// whether the type of key is public or private.
//
func (z *Rsa) ImportPrivateKey(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRsa_ImportPrivateKey(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Imports a private key from a private key object. The imported private key is
// used in methods that sign or decrypt.
func (z *Rsa) ImportPrivateKeyObj(arg1 *PrivateKey) bool {

    v := C.CkRsa_ImportPrivateKeyObj(z.ckObj, arg1.ckObj)


    return v != 0
}


// Imports a public key from XML format. After successful import, the public key
// can be used to encrypt or decrypt.
// 
// Note: Importing a public key overwrites the key that is currently contained in
// this object - even if it's a private key.
// 
// A public key consists of modulus and exponent using base64 representation:
// _LT_RSAPublicKey>
//   _LT_Modulus>..._LT_/Modulus>
//   _LT_Exponent>..._LT_/Exponent>
// _LT_/RSAPublicKey>
// 
// Important: The Rsa object can contain either a private key or a public key, but
// not both. Importing a private key overwrites the existing key regardless of
// whether the type of key is public or private.
//
func (z *Rsa) ImportPublicKey(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRsa_ImportPublicKey(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Imports a public key from a public key object. The imported public key is used
// in methods that encrypt data or verify signatures.
func (z *Rsa) ImportPublicKeyObj(arg1 *PublicKey) bool {

    v := C.CkRsa_ImportPublicKeyObj(z.ckObj, arg1.ckObj)


    return v != 0
}


// Duplicates OpenSSL's rsautl utility for creating RSA signatures. The contents of
// bd are signed. If successful, the result is that bd contains the RSA
// signature that itself contains (embeds) the original data.
func (z *Rsa) OpenSslSignBd(arg1 *BinData) bool {

    v := C.CkRsa_OpenSslSignBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Duplicates OpenSSL's rsautl utility for creating RSA signatures. Input data
// consists of binary bytes, and returns the signature bytes.
func (z *Rsa) OpenSslSignBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_OpenSslSignBytes(z.ckObj, ckbyd1, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Duplicates OpenSSL's rsautl utility for creating RSA signatures. Input data
// consists of binary bytes, and returns the signature as a string encoded
// according to the EncodingMode property (base64, hex, etc.).
// return a string or nil if failed.
func (z *Rsa) OpenSslSignBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkRsa_openSslSignBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Duplicates OpenSSL's rsautl utility for creating RSA signatures. Input data is a
// string, and returns the signature bytes.
func (z *Rsa) OpenSslSignString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_OpenSslSignString(z.ckObj, cstr1, ckOutBytes)

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


// Duplicates OpenSSL's rsautl utility for creating RSA signatures. Input data is a
// string, and returns the signature as a string encoded according to the
// EncodingMode property (base64, hex, etc.).
// return a string or nil if failed.
func (z *Rsa) OpenSslSignStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkRsa_openSslSignStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Duplicates OpenSSL's rsautl utility for verifying RSA signatures and recovering
// the original data. On input, the bd contains the RSA signature that embeds the
// original data. If successful (i.e. the signature was verified), then the bd is
// transformed to contain just the original data.
func (z *Rsa) OpenSslVerifyBd(arg1 *BinData) bool {

    v := C.CkRsa_OpenSslVerifyBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Duplicates OpenSSL's rsautl utility for verifying RSA signatures and recovering
// the original data. Input data consists of the raw signature bytes and returns
// the original bytes.
func (z *Rsa) OpenSslVerifyBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_OpenSslVerifyBytes(z.ckObj, ckbyd1, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Duplicates OpenSSL's rsautl utility for verifying RSA signatures and recovering
// the original data. Input data is a signature string encoded according to the
// EncodingMode property (base64, hex, etc.). Returns the original bytes.
func (z *Rsa) OpenSslVerifyBytesENC(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_OpenSslVerifyBytesENC(z.ckObj, cstr1, ckOutBytes)

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


// Duplicates OpenSSL's rsautl utility for verifying RSA signatures and recovering
// the original data. Input data consists of the raw signature bytes and returns
// the original string.
// return a string or nil if failed.
func (z *Rsa) OpenSslVerifyString(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkRsa_openSslVerifyString(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Duplicates OpenSSL's rsautl utility for verifying RSA signatures and recovering
// the original data. Input data is a signature string encoded according to the
// EncodingMode property (base64, hex, etc.). Returns the original string.
// return a string or nil if failed.
func (z *Rsa) OpenSslVerifyStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkRsa_openSslVerifyStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Provides the private or public key indirectly through a certificate. This method
// is especially useful on Windows computers where the private key is installed as
// non-exportable (such as on a hardware token).
func (z *Rsa) SetX509Cert(arg1 *Cert, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkRsa_SetX509Cert(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Creates an RSA digital signature by hashing the contents of bdData and then
// signing the hash. The hash algorithm is specified by hashAlgorithm, which may be "SHA-1",
// "MD5", "MD2", "SHA-256", "SHA-384", or "SHA-512". The resulting signature is
// returned in bdSig.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false. (The LittleEndian property should also be set to false to match
// Amazon web services, such as CloudFront.)
// 
// A private key is required to create digital signatures.
//
func (z *Rsa) SignBd(arg1 *BinData, arg2 string, arg3 *BinData) bool {
    cstr2 := C.CString(arg2)

    v := C.CkRsa_SignBd(z.ckObj, arg1.ckObj, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Creates an RSA digital signature by hashing binaryData and then signing the hash. The
// hash algorithm is specified by hashAlgorithm, which may be "SHA-1", "MD5", "MD2",
// "SHA-256", "SHA-384", or "SHA-512". The recommended hash algorithm is "SHA-1".
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false. (The LittleEndian property should also be set to false to match
// Amazon web services, such as CloudFront.)
// 
// A private key is required to create a digital signature.
// 
// An error is indicated when a byte array of 0 length is returned.
//
func (z *Rsa) SignBytes(arg1 []byte, arg2 string) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_SignBytes(z.ckObj, ckbyd1, cstr2, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Creates an RSA digital signature by hashing binaryData and then signing the hash. The
// hash algorithm is specified by hashAlgorithm, which may be "SHA-1", "MD5", "MD2",
// "SHA-256", "SHA-384", or "SHA-512". The recommended hash algorithm is "SHA-1".
// The digital signature is returned as an encoded string, where the encoding is
// specified by the EncodingMode property.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false. (The LittleEndian property should also be set to false to match
// Amazon web services, such as CloudFront.)
// 
// A private key is required to create a digital signature.
// 
// An error is indicated when null reference is returned.
//
// return a string or nil if failed.
func (z *Rsa) SignBytesENC(arg1 []byte, arg2 string) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    retStr := C.CkRsa_signBytesENC(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The same as the SignBytes method, except the hash to be signed is passed
// directly.
func (z *Rsa) SignHash(arg1 []byte, arg2 string) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_SignHash(z.ckObj, ckbyd1, cstr2, ckOutBytes)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// The same as SignBytesENC except the hash is passed directly.
// return a string or nil if failed.
func (z *Rsa) SignHashENC(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkRsa_signHashENC(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Creates an RSA digital signature by hashing strToBeHashed and then signing the hash. The
// hash algorithm is specified by hashAlgorithm, which may be "SHA-1", "MD5", "MD2",
// "SHA-256", "SHA-384", or "SHA-512". The recommended hash algorithm is "SHA-1".
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false. (The LittleEndian property should also be set to false to match
// Amazon web services, such as CloudFront.)
// 
// A private key is required to create a digital signature.
// 
// An error is indicated when a byte array of 0 length is returned.
//
func (z *Rsa) SignString(arg1 string, arg2 string) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkRsa_SignString(z.ckObj, cstr1, cstr2, ckOutBytes)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Creates an RSA digital signature by hashing strToBeHashed and then signing the hash. The
// hash algorithm is specified by hashAlgorithm, which may be "SHA-1", "MD5", "MD2",
// "SHA-256", "SHA-384", or "SHA-512". The recommended hash algorithm is "SHA-1".
// The digital signature is returned as an encoded string, where the encoding is
// specified by the EncodingMode property.
// 
// Important: If trying to match OpenSSL results, set the LittleEndian property =
// false. (The LittleEndian property should also be set to false to match
// Amazon web services, such as CloudFront.)
// 
// A private key is required to create a digital signature.
// 
// An error is indicated when null reference is returned.
//
// return a string or nil if failed.
func (z *Rsa) SignStringENC(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkRsa_signStringENC(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Imports a .snk file to an XML document that can be imported via the
// ImportPrivateKey method.
// return a string or nil if failed.
func (z *Rsa) SnkToXml(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkRsa_snkToXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Unlocks the component. This must be called once prior to calling any other
// method.
func (z *Rsa) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRsa_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Verifies an RSA digital signature. Returns true if the signature in bdSig is
// valid an confirms that the original data in bdData has not been modified. The hashAlgorithm
// may be "SHA-1", "MD5", "MD2", "SHA-256", "SHA-384", or "SHA-512".
func (z *Rsa) VerifyBd(arg1 *BinData, arg2 string, arg3 *BinData) bool {
    cstr2 := C.CString(arg2)

    v := C.CkRsa_VerifyBd(z.ckObj, arg1.ckObj, cstr2, arg3.ckObj)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Verifies an RSA digital signature. Returns true if the signature is valid for
// the originalData. The hashAlgorithm may be "SHA-1", "MD5", "MD2", "SHA-256", "SHA-384", or
// "SHA-512". The recommended hash algorithm is "SHA-1".
func (z *Rsa) VerifyBytes(arg1 []byte, arg2 string, arg3 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkRsa_VerifyBytes(z.ckObj, ckbyd1, cstr2, ckbyd3)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// Verifies an RSA digital signature. Returns true if the signature is valid for
// the originalData. The hashAlgorithm may be "SHA-1", "MD5", "MD2", "SHA-256", "SHA-384", or
// "SHA-512". The recommended hash algorithm is "SHA-1".
// 
// The encodedSig is a digital signature encoded according to the EncodingMode property
// (i.e. base64, hex, etc.).
//
func (z *Rsa) VerifyBytesENC(arg1 []byte, arg2 string, arg3 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkRsa_VerifyBytesENC(z.ckObj, ckbyd1, cstr2, cstr3)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// The same as VerifyBytes except the hash of the original data is passed directly.
func (z *Rsa) VerifyHash(arg1 []byte, arg2 string, arg3 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkRsa_VerifyHash(z.ckObj, ckbyd1, cstr2, ckbyd3)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// The same as VerifyBytesENC except the hash of the original data is passed
// directly.
func (z *Rsa) VerifyHashENC(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkRsa_VerifyHashENC(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Returns true if the XML contains a valid RSA private key. Otherwise returns
// false.
func (z *Rsa) VerifyPrivateKey(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRsa_VerifyPrivateKey(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Verifies an RSA digital signature. Returns true if the signature is valid for
// the originalString. The hashAlgorithm may be "SHA-1", "MD5", "MD2", "SHA-256", "SHA-384", or
// "SHA-512". The recommended hash algorithm is "SHA-1".
func (z *Rsa) VerifyString(arg1 string, arg2 string, arg3 []byte) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkRsa_VerifyString(z.ckObj, cstr1, cstr2, ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// Verifies an RSA digital signature. Returns true if the signature is valid for
// the originalString. The hashAlgorithm may be "SHA-1", "MD5", "MD2", "SHA-256", "SHA-384", or
// "SHA-512". The recommended hash algorithm is "SHA-1".
// 
// The encodedSig is a digital signature encoded according to the EncodingMode property
// (i.e. base64, hex, etc.).
//
func (z *Rsa) VerifyStringENC(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkRsa_VerifyStringENC(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


