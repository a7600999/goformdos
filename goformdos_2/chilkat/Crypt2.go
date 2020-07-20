// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkCrypt2.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCrypt2() *Crypt2 {
	return &Crypt2{C.CkCrypt2_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Crypt2) DisposeCrypt2() {
    if z != nil {
	C.CkCrypt2_Dispose(z.ckObj)
	}
}




func (z *Crypt2) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkCrypt2_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Crypt2) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkCrypt2_setExternalProgress(z.ckObj,1)
}

func (z *Crypt2) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkCrypt2_setExternalProgress(z.ckObj,1)
}

func (z *Crypt2) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkCrypt2_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Crypt2) AbortCurrent() bool {
    v := int(C.CkCrypt2_getAbortCurrent(z.ckObj))
    return v != 0
}
// property setter: AbortCurrent
func (z *Crypt2) SetAbortCurrent(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCrypt2_putAbortCurrent(z.ckObj,v)
}

// property: BCryptWorkFactor
// The BCrypt work factor to be used for the BCryptHash and BCryptVerify. This is
// the log2 of the number of rounds of hashing to apply. For example, if the work
// (cost) factor is 12, then 2^12 rounds of hashing are applied. The purpose of
// this cost factor is to make the BCrypt computation expensive enought to prevent
// brute-force attacks. (Any complaints about BCrypt "not being fast enough" will
// be ignored.)
// 
// This property must have a value ranging from 4 to 31 inclusive.
// 
// The default value is 10.
//
func (z *Crypt2) BCryptWorkFactor() int {
    return int(C.CkCrypt2_getBCryptWorkFactor(z.ckObj))
}
// property setter: BCryptWorkFactor
func (z *Crypt2) SetBCryptWorkFactor(value int) {
    C.CkCrypt2_putBCryptWorkFactor(z.ckObj,C.int(value))
}

// property: BlockSize
// The block-size (in bytes) of the selected encryption algorithm. For example, if
// the CryptAlgorithm property is set to "aes", the BlockSize property is
// automatically set to 16. The block-size for the ARC4 streaming encryption
// algorithm is 1.
func (z *Crypt2) BlockSize() int {
    return int(C.CkCrypt2_getBlockSize(z.ckObj))
}

// property: CadesEnabled
// Applies to all methods that create PKCS7 signatures. To create a CAdES-BES
// signature, set this property equal to true. The default value of this property
// is false.
func (z *Crypt2) CadesEnabled() bool {
    v := int(C.CkCrypt2_getCadesEnabled(z.ckObj))
    return v != 0
}
// property setter: CadesEnabled
func (z *Crypt2) SetCadesEnabled(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCrypt2_putCadesEnabled(z.ckObj,v)
}

// property: CadesSigPolicyHash
// This is the base64 hash of the policy document located at the CadesSigPolicyUri.
// You can use either the SHA256 or SHA1 hash. You may use this online tool to
// compute the base64 hash:Compute Base64 Hash for CaDES Signature Policy URL
// <http://tools.chilkat.io/hashFileAtUrl.cshtml>
// 
// Note: This property applies to all methods that create PKCS7 signatures. To
// create a CAdES-EPES signature, set the CadesEnabled property = true, and also
// provide values for each of the following properties: CadesSigPolicyHash,
// CadesSigPolicyId, and CadesSigPolicyUri. For example (in pseudo-code):
// cryptObj.CadesSigPolicyId = "2.16.76.1.7.1.1.1"
// cryptObj.CadesSigPolicyUri = "http://politicas.icpbrasil.gov.br/PA_AD_RB.der"
// cryptObj.CadesSigPolicyHash = "rySugyKaMhiMR8Y/o5yuU2A2bF0="
// Note: Do NOT use the values above. They are only provided as an example to show
// valid values. For example, the Policy ID is an OID. The Policy URI is a
// typically a URL to a DER encoded policy file, and the Policy Hash is a base64
// encoded hash.
//
func (z *Crypt2) CadesSigPolicyHash() string {
    return C.GoString(C.CkCrypt2_cadesSigPolicyHash(z.ckObj))
}
// property setter: CadesSigPolicyHash
func (z *Crypt2) SetCadesSigPolicyHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putCadesSigPolicyHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CadesSigPolicyId
// See the description for the CadesSigPolicyHash property above.
func (z *Crypt2) CadesSigPolicyId() string {
    return C.GoString(C.CkCrypt2_cadesSigPolicyId(z.ckObj))
}
// property setter: CadesSigPolicyId
func (z *Crypt2) SetCadesSigPolicyId(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putCadesSigPolicyId(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CadesSigPolicyUri
// See the description for the CadesSigPolicyHash property above.
func (z *Crypt2) CadesSigPolicyUri() string {
    return C.GoString(C.CkCrypt2_cadesSigPolicyUri(z.ckObj))
}
// property setter: CadesSigPolicyUri
func (z *Crypt2) SetCadesSigPolicyUri(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putCadesSigPolicyUri(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Charset
// Controls the character encoding of the text encrypted, signed, hashed or
// compressed. This property is relevant wherever strings are used as inputs or
// outputs.
// 
// This property defaults to the ANSI charset of the computer. For example, the
// default ANSI code page on Windows computers in the USA and Western Europe would
// be "windows-1252".
// 
// When working with strings, it is important to know the exact bytes that are
// being encrypted/hashed/signed/compressed. This is critical when interoperating
// with other systems. If your application is sending an encrypted string to
// another system that will decrypt it, you will need to know the encoding of the
// string that is expected on the receiving end (after decryption). If you pass
// Unicode data (2 byte per character) to the encryptor, subsequent decryption will
// reproduce the original Unicode. However, it may be that your program works with
// Unicode strings, but the recipient of the encrypted data works with iso-8859-1
// strings. In such a case, setting the Charset property to "iso-8859-1" causes the
// character data to be automatically converted to the Charset before being
// encrypted (or compressed, or hashed, or signed). The set of valid charsets is
// listed below:
// hex
// base64
//     * "hex" and "base64" are special values that allow for binary (non-text) encoded data to be passed to any method where the input data is a string.
//        Rather than converting to an actual charset (such as utf-8, iso-8859-1), the binary data is decoded, and the decoded bytes are passed
//         to the underlying encryptor, hashing, signing, etc.
// ANSI
// us-ascii
// unicode
// unicodefffe
// iso-8859-1
// iso-8859-2
// iso-8859-3
// iso-8859-4
// iso-8859-5
// iso-8859-6
// iso-8859-7
// iso-8859-8
// iso-8859-9
// iso-8859-13
// iso-8859-15
// windows-874
// windows-1250
// windows-1251
// windows-1252
// windows-1253
// windows-1254
// windows-1255
// windows-1256
// windows-1257
// windows-1258
// utf-7
// utf-8
// utf-32
// utf-32be
// shift_jis
// gb2312
// ks_c_5601-1987
// big5
// iso-2022-jp
// iso-2022-kr
// euc-jp
// euc-kr
// macintosh
// x-mac-japanese
// x-mac-chinesetrad
// x-mac-korean
// x-mac-arabic
// x-mac-hebrew
// x-mac-greek
// x-mac-cyrillic
// x-mac-chinesesimp
// x-mac-romanian
// x-mac-ukrainian
// x-mac-thai
// x-mac-ce
// x-mac-icelandic
// x-mac-turkish
// x-mac-croatian
// asmo-708
// dos-720
// dos-862
// ibm037
// ibm437
// ibm500
// ibm737
// ibm775
// ibm850
// ibm852
// ibm855
// ibm857
// ibm00858
// ibm860
// ibm861
// ibm863
// ibm864
// ibm865
// cp866
// ibm869
// ibm870
// cp875
// koi8-r
// koi8-u
//
func (z *Crypt2) Charset() string {
    return C.GoString(C.CkCrypt2_charset(z.ckObj))
}
// property setter: Charset
func (z *Crypt2) SetCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CipherMode
// Controls the cipher mode for block encryption algorithms (AES, Blowfish,TwoFish,
// DES, 3DES, RC2). Possible values are "CBC" (the default) , "ECB", "CTR", "OFB",
// "GCM", and "CFB". These acronyms have the following meanings:
// 
//     CBC: Cipher Block Chaining,
//     ECB: Electronic CookBook
//     CTR: Counter Mode
//     CFB: Cipher Feedback
//     OFB: Output Feedback
//     GCM: Galois/Counter Mode
// 
// (see http://en.wikipedia.org/wiki/Block_cipher_modes_of_operation )
// 
// Note: Prior to Chilkat v9.5.0.55, the CFB mode is only implemented for AES,
// Blowfish, and DES/3DES, and the CTR mode is only implemented for AES.
// 
// Starting in v9.5.0.55 CFB and OFB modes are useable with all encryption
// algorithms, and GCM (Galois/Counter Mode) is available with any cipher having a
// 16-byte block size, such as AES and Twofish. CFB, OFB, CTR, and GCM modes
// convert block ciphers into stream ciphers. In these modes of operation, the
// PaddingScheme property is unused because no padding occurs.
//
func (z *Crypt2) CipherMode() string {
    return C.GoString(C.CkCrypt2_cipherMode(z.ckObj))
}
// property setter: CipherMode
func (z *Crypt2) SetCipherMode(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putCipherMode(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CmsOptions
// A JSON string for controlling extra CMS (PKCS7) signature and validation
// options.
func (z *Crypt2) CmsOptions() string {
    return C.GoString(C.CkCrypt2_cmsOptions(z.ckObj))
}
// property setter: CmsOptions
func (z *Crypt2) SetCmsOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putCmsOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CryptAlgorithm
// Selects the encryption algorithm for encrypting and decrypting. Possible values
// are: "chacha20", "pki", "aes", "blowfish2", "des", "3des", "rc2", "arc4",
// "twofish", "pbes1" and "pbes2". The "pki" encryption algorithm isn't a specific
// algorithm, but instead tells the component to encrypt/decrypt using public-key
// encryption with digital certificates. The other choices are symmetric encryption
// algorithms that do not involve digital certificates and public/private keys.
// 
// The default value is "aes"
// 
// The original Chilkat implementation of Blowfish (in 2004) has a 4321
// byte-swapping issue (the results are 4321 byte-swapped). The newer
// implementation (in 2006 and named "blowfish2") does not byte swap. This should
// be used for compatibility with other Blowfish software. If an application needs
// to decrypt something encrypted with the old 4321 byte-swapped blowfish, set the
// property to "blowfish_old".
// 
// Password-based encryption (PBE) is selected by setting this property to "pbes1"
// or "pbes2". Password-based encryption is defined in the PKCS5 Password-Based
// Cryptography Standard at https://tools.ietf.org/html/rfc2898. If PBE is used,
// the underlying encryption algorithm is specified by the PbesAlgorithm property.
// The underlying encryption (PbesAlgorithm) for PBES1 is limited to 56-bit DES or
// 64-bit RC2.
// 
// Note:The chacha20 algorithm is introduced in Chilkat v9.5.0.55.
//
func (z *Crypt2) CryptAlgorithm() string {
    return C.GoString(C.CkCrypt2_cryptAlgorithm(z.ckObj))
}
// property setter: CryptAlgorithm
func (z *Crypt2) SetCryptAlgorithm(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putCryptAlgorithm(z.ckObj,cStr)
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
func (z *Crypt2) DebugLogFilePath() string {
    return C.GoString(C.CkCrypt2_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Crypt2) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EncodingMode
// Controls the encoding of binary data to a printable string for many methods. The
// valid modes are "Base64", "modBase64", "base64url", "Base32", "Base58", "UU",
// "QP" (for quoted-printable), "URL" (for url-encoding), "Hex", "Q", "B",
// "url_oauth", "url_rfc1738", "url_rfc2396", "url_rfc3986", "fingerprint", or
// "decimal".
// 
// The default value is "base64"
// 
// The "fingerprint" and"decimal" encodings are introduced in Chilkat v9.5.0.55.
// 
// The "fingerprint" encoding is a lowercase hex encoding where each hex digit is
// separated by a colon character. For example:
// 6a:de:e0:af:56:f8:0c:04:11:5b:ef:4d:49:ad:09:23
// 
// The "decimal" encoding is for converting large decimal integers to/from a
// big-endian binary representation. For example, the decimal string
// "72623859790382856" converts to the bytes 0x01 0x02 0x03 0x04 0x05 0x06 0x07
// 0x08.
//
func (z *Crypt2) EncodingMode() string {
    return C.GoString(C.CkCrypt2_encodingMode(z.ckObj))
}
// property setter: EncodingMode
func (z *Crypt2) SetEncodingMode(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putEncodingMode(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: FirstChunk
// Chilkat Crypt2 provides the ability to feed the encryption/decryption methods
// with chunks of data. This allows a large amount of data, or a data stream, to be
// fed piecemeal for encrypting or decrypting. It applies to all symmetric
// algorithms currently supported (AES, Blowfish, Twofish, 3DES, RC2, DES, ARC4),
// and all algorithms supported in the future.
// 
// The default value for both FirstChunk and LastChunk is true. This means when
// an Encrypt* or Decrypt* method is called, it is both the first and last chunk
// (i.e. it's the entire amount of data to be encrypted or decrypted).
// 
// If you wish to feed the data piecemeal, do this:
// 
//     Set FirstChunk = true, LastChunk = false for the first chunk of data.
//     For all "middle" chunks (i.e. all chunks except for the final chunk) set
//     FirstChunk = false and LastChunk = false.
//     For the final chunk, set FirstChunk = false and LastChunk = true
// 
// There is no need to worry about feeding data according to the block size of the
// encryption algorithm. For example, AES has a block size of 16 bytes. Data may be
// fed in chunks of any size. The Chilkat Crypt2 component will buffer the data.
// When the final chunk is passed, the output is padded to the algorithm's block
// size according to the PaddingScheme.
//
func (z *Crypt2) FirstChunk() bool {
    v := int(C.CkCrypt2_getFirstChunk(z.ckObj))
    return v != 0
}
// property setter: FirstChunk
func (z *Crypt2) SetFirstChunk(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCrypt2_putFirstChunk(z.ckObj,v)
}

// property: HashAlgorithm
// Selects the hash algorithm used by methods that create hashes. The valid choices
// are "sha1", "sha256", "sha384", "sha512", "sha3-224", "sha3-256", "sha3-384",
// "sha3-512", "md2", "md5", "haval", "ripemd128", "ripemd160","ripemd256", or
// "ripemd320".
// 
// Note: SHA-2 designates a set of cryptographic hash functions that includes
// SHA-256, SHA-384, and SHA-512. Chilkat by definition supports "SHA-2" because it
// supports these algorithms.
// 
// The default value is "sha1".
// 
// Note: The HAVAL hash algorithm is affected by two other properties: HavalRounds
// and KeyLength.
// 
//     The HavalRounds may have values of 3, 4, or 5.
//     The KeyLength may have values of 128, 160, 192, 224, or 256.
// 
// Note: The "sha3-224", "sha3-256", "sha3-384", "sha3-512" algorithms are added in
// Chilkat v9.5.0.83.
//
func (z *Crypt2) HashAlgorithm() string {
    return C.GoString(C.CkCrypt2_hashAlgorithm(z.ckObj))
}
// property setter: HashAlgorithm
func (z *Crypt2) SetHashAlgorithm(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putHashAlgorithm(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HavalRounds
// Applies to the HAVAL hash algorithm only and must be set to the integer value 3,
// 4, or 5. The default value is 3.
func (z *Crypt2) HavalRounds() int {
    return int(C.CkCrypt2_getHavalRounds(z.ckObj))
}
// property setter: HavalRounds
func (z *Crypt2) SetHavalRounds(value int) {
    C.CkCrypt2_putHavalRounds(z.ckObj,C.int(value))
}

// property: HeartbeatMs
// The number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort some methods call prior to
// completion. If HeartbeatMs is 0 (the default), no AbortCheck event callbacks
// will fire.
// 
// The methods with event callbacks are: CkDecryptFile, CkEncryptFile, HashFile,
// and HashFileENC.
//
func (z *Crypt2) HeartbeatMs() int {
    return int(C.CkCrypt2_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Crypt2) SetHeartbeatMs(value int) {
    C.CkCrypt2_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: IncludeCertChain
// Only applies when creating digital signatures. If true (the default), then
// additional certificates (if any) in the chain of authentication are included in
// the PKCS7 digital signature.
func (z *Crypt2) IncludeCertChain() bool {
    v := int(C.CkCrypt2_getIncludeCertChain(z.ckObj))
    return v != 0
}
// property setter: IncludeCertChain
func (z *Crypt2) SetIncludeCertChain(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCrypt2_putIncludeCertChain(z.ckObj,v)
}

// property: InitialCount
// The initial counter for the ChaCha20 encryption algorithm. The default value is
// 0.
func (z *Crypt2) InitialCount() int {
    return int(C.CkCrypt2_getInitialCount(z.ckObj))
}
// property setter: InitialCount
func (z *Crypt2) SetInitialCount(value int) {
    C.CkCrypt2_putInitialCount(z.ckObj,C.int(value))
}

// property: IterationCount
// Iteration count to be used with password-based encryption (PBE). Password-based
// encryption is defined in the PKCS5 Password-Based Cryptography Standard at
// http://www.rsa.com/rsalabs/node.asp?id=2127
// 
// The purpose of the iteration count is to increase the computation required to
// encrypt and decrypt. A larger iteration count makes cracking via exhaustive
// search more difficult. The default value is 1024.
//
func (z *Crypt2) IterationCount() int {
    return int(C.CkCrypt2_getIterationCount(z.ckObj))
}
// property setter: IterationCount
func (z *Crypt2) SetIterationCount(value int) {
    C.CkCrypt2_putIterationCount(z.ckObj,C.int(value))
}

// property: IV
// The initialization vector to be used with symmetric encryption algorithms (AES,
// Blowfish, Twofish, etc.). If left unset, no initialization vector is used.
func (z *Crypt2) IV() []byte {
    ckbyd := C.CkByteData_Create()
    C.CkCrypt2_getIV(z.ckObj,ckbyd)
    p := C.CkByteData_getData(ckbyd);
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckbyd)) )
    C.CkByteData_Dispose(ckbyd)
    return retBytes
}

// property setter: IV
func (z *Crypt2) SetIV(dataBytes []byte) {
    ckbyd := C.CkByteData_Create()
    pData := C.CBytes(dataBytes)
    C.CkByteData_borrowData(ckbyd, (*C.uchar)(pData), C.ulong(len(dataBytes)))
    C.CkCrypt2_putIV(z.ckObj,ckbyd)    
    C.free(pData)
    C.CkByteData_Dispose(ckbyd)
}


// property: KeyLength
// The key length in bits for symmetric encryption algorithms. The default value is
// 256.
func (z *Crypt2) KeyLength() int {
    return int(C.CkCrypt2_getKeyLength(z.ckObj))
}
// property setter: KeyLength
func (z *Crypt2) SetKeyLength(value int) {
    C.CkCrypt2_putKeyLength(z.ckObj,C.int(value))
}

// property: LastChunk
// (See the description for the FirstChunk property.)
func (z *Crypt2) LastChunk() bool {
    v := int(C.CkCrypt2_getLastChunk(z.ckObj))
    return v != 0
}
// property setter: LastChunk
func (z *Crypt2) SetLastChunk(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCrypt2_putLastChunk(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Crypt2) LastErrorHtml() string {
    return C.GoString(C.CkCrypt2_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Crypt2) LastErrorText() string {
    return C.GoString(C.CkCrypt2_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Crypt2) LastErrorXml() string {
    return C.GoString(C.CkCrypt2_lastErrorXml(z.ckObj))
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
func (z *Crypt2) LastMethodSuccess() bool {
    v := int(C.CkCrypt2_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Crypt2) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCrypt2_putLastMethodSuccess(z.ckObj,v)
}

// property: MacAlgorithm
// Selects the MAC algorithm to be used for any of the Mac methods, such as
// MacStringENC, MacBytes, etc. The default value is "hmac". Possible values are
// "hmac" and "poly1305".
func (z *Crypt2) MacAlgorithm() string {
    return C.GoString(C.CkCrypt2_macAlgorithm(z.ckObj))
}
// property setter: MacAlgorithm
func (z *Crypt2) SetMacAlgorithm(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putMacAlgorithm(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: NumSignerCerts
// This property is set when a digital signature is verified. It contains the
// number of signer certificates. Each signing certificate can be retrieved by
// calling the GetSignerCert method, passing an index from 0 to NumSignerCerts-1.
func (z *Crypt2) NumSignerCerts() int {
    return int(C.CkCrypt2_getNumSignerCerts(z.ckObj))
}

// property: OaepHash
// Selects the hash algorithm for use within OAEP padding when encrypting using
// "pki" with RSAES-OAEP. The valid choices are "sha1", "sha256", "sha384",
// "sha512",
// 
// The default value is "sha256"
//
func (z *Crypt2) OaepHash() string {
    return C.GoString(C.CkCrypt2_oaepHash(z.ckObj))
}
// property setter: OaepHash
func (z *Crypt2) SetOaepHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putOaepHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OaepMgfHash
// Selects the MGF hash algorithm for use within OAEP padding when encrypting using
// "pki" with RSAES-OAEP. The valid choices are "sha1", "sha256", "sha384",
// "sha512", The default is "sha1".
func (z *Crypt2) OaepMgfHash() string {
    return C.GoString(C.CkCrypt2_oaepMgfHash(z.ckObj))
}
// property setter: OaepMgfHash
func (z *Crypt2) SetOaepMgfHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putOaepMgfHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OaepPadding
// Selects the RSA encryption scheme when encrypting using "pki" (with a
// certificate and private key). The default value is false, which selects
// RSAES_PKCS1-V1_5. If set to true, then RSAES_OAEP is used.
func (z *Crypt2) OaepPadding() bool {
    v := int(C.CkCrypt2_getOaepPadding(z.ckObj))
    return v != 0
}
// property setter: OaepPadding
func (z *Crypt2) SetOaepPadding(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCrypt2_putOaepPadding(z.ckObj,v)
}

// property: PaddingScheme
// The padding scheme used by block encryption algorithms such as AES (Rijndael),
// Blowfish, Twofish, RC2, DES, 3DES, etc. Block encryption algorithms pad
// encrypted data to a multiple of algorithm's block size. The default value of
// this property is 0.
// 
// Possible values are:
// 
// 0 = RFC 1423 padding scheme: Each padding byte is set to the number of padding
// bytes. If the data is already a multiple of algorithm's block size bytes, an
// extra block is appended each having a value equal to the block size. (for
// example, if the algorithm's block size is 16, then 16 bytes having the value
// 0x10 are added.). (This is also known as PKCS5 padding: PKCS #5 padding string
// consists of a sequence of bytes, each of which is equal to the total number of
// padding bytes added. )
// 
// 1 = FIPS81 (Federal Information Processing Standards 81) where the last byte
// contains the number of padding bytes, including itself, and the other padding
// bytes are set to random values.
// 
// 2 = Each padding byte is set to a random value. The decryptor must know how many
// bytes are in the original unencrypted data.
// 
// 3 = Pad with NULLs. (If already a multiple of the algorithm's block size, no
// padding is added).
// 
// 4 = Pad with SPACE chars(0x20). (If already a multiple of algorithm's block
// size, no padding is added).
//
func (z *Crypt2) PaddingScheme() int {
    return int(C.CkCrypt2_getPaddingScheme(z.ckObj))
}
// property setter: PaddingScheme
func (z *Crypt2) SetPaddingScheme(value int) {
    C.CkCrypt2_putPaddingScheme(z.ckObj,C.int(value))
}

// property: PbesAlgorithm
// If the CryptAlgorithm property is set to "pbes1" or "pbes2", this property
// specifies the underlying encryption algorithm to be used with password-based
// encryption (PBE). Password-based encryption is defined in the PKCS5
// Password-Based Cryptography Standard at
// http://www.rsa.com/rsalabs/node.asp?id=2127
func (z *Crypt2) PbesAlgorithm() string {
    return C.GoString(C.CkCrypt2_pbesAlgorithm(z.ckObj))
}
// property setter: PbesAlgorithm
func (z *Crypt2) SetPbesAlgorithm(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putPbesAlgorithm(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: PbesPassword
// The password to be used with password-based encryption (PBE). Password-based
// encryption is defined in the PKCS5 Password-Based Cryptography Standard at
// http://www.rsa.com/rsalabs/node.asp?id=2127
func (z *Crypt2) PbesPassword() string {
    return C.GoString(C.CkCrypt2_pbesPassword(z.ckObj))
}
// property setter: PbesPassword
func (z *Crypt2) SetPbesPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putPbesPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Pkcs7CryptAlg
// When the CryptAlgorithm property is "PKI" to select PKCS7 public-key encryption,
// this selects the underlying symmetric encryption algorithm. Possible values are:
// "aes", "des", "3des", and "rc2". The default value is "aes".
func (z *Crypt2) Pkcs7CryptAlg() string {
    return C.GoString(C.CkCrypt2_pkcs7CryptAlg(z.ckObj))
}
// property setter: Pkcs7CryptAlg
func (z *Crypt2) SetPkcs7CryptAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putPkcs7CryptAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Rc2EffectiveKeyLength
// The effective key length (in bits) for the RC2 encryption algorithm. When RC2 is
// used, both the KeyLength and Rc2EffectiveKeyLength properties should be set. For
// RC2, both should be between 8 and 1024 (inclusive).
// 
// The default value is 128
//
func (z *Crypt2) Rc2EffectiveKeyLength() int {
    return int(C.CkCrypt2_getRc2EffectiveKeyLength(z.ckObj))
}
// property setter: Rc2EffectiveKeyLength
func (z *Crypt2) SetRc2EffectiveKeyLength(value int) {
    C.CkCrypt2_putRc2EffectiveKeyLength(z.ckObj,C.int(value))
}

// property: Salt
// The salt to be used with password-based encryption (PBE). Password-based
// encryption is defined in the PKCS5 Password-Based Cryptography Standard at
// http://www.rsa.com/rsalabs/node.asp?id=2127
// 
// To clarify: This property is used in encryption when the CryptAlgorithm is set
// to "pbes1" or "pbes2". Also note that it is not used by the Pbkdf1 or Pbkdf2
// methods, as the salt is passed in an argument to those methods.
//
func (z *Crypt2) Salt() []byte {
    ckbyd := C.CkByteData_Create()
    C.CkCrypt2_getSalt(z.ckObj,ckbyd)
    p := C.CkByteData_getData(ckbyd);
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckbyd)) )
    C.CkByteData_Dispose(ckbyd)
    return retBytes
}

// property setter: Salt
func (z *Crypt2) SetSalt(dataBytes []byte) {
    ckbyd := C.CkByteData_Create()
    pData := C.CBytes(dataBytes)
    C.CkByteData_borrowData(ckbyd, (*C.uchar)(pData), C.ulong(len(dataBytes)))
    C.CkCrypt2_putSalt(z.ckObj,ckbyd)    
    C.free(pData)
    C.CkByteData_Dispose(ckbyd)
}


// property: SecretKey
// The binary secret key used for symmetric encryption (Aes, Blowfish, Twofish,
// ChaCha20, ARC4, 3DES, RC2, etc.). The secret key must be identical for
// decryption to succeed. The length in bytes of the SecretKey must equal the
// KeyLength/8.
func (z *Crypt2) SecretKey() []byte {
    ckbyd := C.CkByteData_Create()
    C.CkCrypt2_getSecretKey(z.ckObj,ckbyd)
    p := C.CkByteData_getData(ckbyd);
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckbyd)) )
    C.CkByteData_Dispose(ckbyd)
    return retBytes
}

// property setter: SecretKey
func (z *Crypt2) SetSecretKey(dataBytes []byte) {
    ckbyd := C.CkByteData_Create()
    pData := C.CBytes(dataBytes)
    C.CkByteData_borrowData(ckbyd, (*C.uchar)(pData), C.ulong(len(dataBytes)))
    C.CkCrypt2_putSecretKey(z.ckObj,ckbyd)    
    C.free(pData)
    C.CkByteData_Dispose(ckbyd)
}


// property: SigningAlg
// This property selects the signature algorithm for the OpaqueSign*, Sign*, and
// CreateDetachedSignature, CreateP7M, and CreateP7S methods. The default value is
// "PKCS1-v1_5". This can be set to "RSASSA-PSS" (or simply "pss") to use the
// RSASSA-PSS signature scheme.
// 
// Note: This property only applies when the private key is an RSA private key. It
// does not apply for ECC or DSA private keys.
//
func (z *Crypt2) SigningAlg() string {
    return C.GoString(C.CkCrypt2_signingAlg(z.ckObj))
}
// property setter: SigningAlg
func (z *Crypt2) SetSigningAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putSigningAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SigningAttributes
// Contains JSON to specify the authenticated (signed) attributes or
// unauthenticated (unsigned) attributes that are to be included in CMS signatures.
// The default value is:
// {
//     "contentType": 1,
//     "signingTime": 1,
//     "messageDigest": 1
// }
// 
// Other possible values that can be added are:
//     signingCertificateV2
//     signingCertificate
//     sMIMECapabilities
//     microsoftRecipientInfo
//     encrypKeyPref
// Contact Chilkat (support@chilkatsoft.com) about other signed/unsigned attributes
// that may be needed for CAdES signatures.
//
func (z *Crypt2) SigningAttributes() string {
    return C.GoString(C.CkCrypt2_signingAttributes(z.ckObj))
}
// property setter: SigningAttributes
func (z *Crypt2) SetSigningAttributes(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putSigningAttributes(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. This property
// defaults to the empty string and should typically remain empty.
// 
// Can be set to a list of the following comma separated keywords:
//     "UseConstructedOctets" - Introduced in v9.5.0.83. When creating opaque CMS
//     signatures (signatures that embed the data being signed), will use the
//     "constructed octets" form of the ASN.1 that holds the data. This is to satify
//     some validators that are brittle/fragile/picky and require a particular format,
//     such as for the ICP-Brazil Digital Signature Standard.
//
func (z *Crypt2) UncommonOptions() string {
    return C.GoString(C.CkCrypt2_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Crypt2) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UuFilename
// When UU encoding, this is the filename to be embedded in UU encoded output. The
// default is "file.dat". When UU decoding, this is the filename found in the UU
// encoded input.
func (z *Crypt2) UuFilename() string {
    return C.GoString(C.CkCrypt2_uuFilename(z.ckObj))
}
// property setter: UuFilename
func (z *Crypt2) SetUuFilename(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putUuFilename(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UuMode
// When UU encoding, this is the file permissions mode to be embedded in UU encoded
// output. The default is "644". When UU decoding, this property is set to the mode
// found in the UU encoded input.
func (z *Crypt2) UuMode() string {
    return C.GoString(C.CkCrypt2_uuMode(z.ckObj))
}
// property setter: UuMode
func (z *Crypt2) SetUuMode(goStr string) {
    cStr := C.CString(goStr)
    C.CkCrypt2_putUuMode(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Crypt2) VerboseLogging() bool {
    v := int(C.CkCrypt2_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Crypt2) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCrypt2_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Crypt2) Version() string {
    return C.GoString(C.CkCrypt2_version(z.ckObj))
}
// Adds a certificate to be used for public-key encryption. (To use public-key
// encryption with digital certificates, set the CryptAlgorithm property = "pki".)
// To encrypt with more than one certificate , call AddEncryptCert once per
// certificate.
func (z *Crypt2) AddEncryptCert(arg1 *Cert)  {

    C.CkCrypt2_AddEncryptCert(z.ckObj, arg1.ckObj)



}


// Adds a PFX to the object's internal list of sources to be searched for
// certificates and private keys when decrypting. Multiple PFX sources can be added
// by calling this method once for each. (On the Windows operating system, the
// registry-based certificate stores are also automatically searched, so it is
// commonly not required to explicitly add PFX sources.)
// 
// The pfxBytes contains the bytes of a PFX file (also known as PKCS12 or .p12).
//
func (z *Crypt2) AddPfxSourceData(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_AddPfxSourceData(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a PFX file to the object's internal list of sources to be searched for
// certificates and private keys when decrypting. Multiple PFX files can be added
// by calling this method once for each. (On the Windows operating system, the
// registry-based certificate stores are also automatically searched, so it is
// commonly not required to explicitly add PFX sources.)
// 
// The pfxFilePath contains the bytes of a PFX file (also known as PKCS12 or .p12).
//
func (z *Crypt2) AddPfxSourceFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_AddPfxSourceFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a certificate to be used for signing. To sign with more than one
// certificate, call AddSigningCert once per certificate.
// 
// Note: This method was added in v9.5.0.83. The SetSigningCert and SetSigningCert2
// methods are used to set the signing certificate for signatures with one signer.
//
func (z *Crypt2) AddSigningCert(arg1 *Cert) bool {

    v := C.CkCrypt2_AddSigningCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Implements the AES Key Wrap Algorithm (RFC 3394) for unwrapping. The kek is the
// Key Encryption Key (the AES key used to unwrap the wrappedKeyData). The arguments and
// return value are binary encoded strings using the encoding specified by encoding
// (which can be "base64", "hex", "base64url", etc.) The full list of supported
// encodings is available at the link below.
// 
// The kek should be an AES key of 16 bytes, 24 bytes, or 32 bytes (i.e. 128-bits,
// 192- bits, or 256-bits). For example, if passed as a hex string, then the kek
// should be 32 chars in length, 48 chars, or 64 chars (because each byte is
// represented as 2 chars in hex).
// 
// The wrappedKeyData contains the data to be unwrapped. The result, if decoded, is 8 bytes
// less than the wrapped key data. For example, if a 256-bit AES key (32 bytes) is
// wrapped, the size of the wrapped key data is 40 bytes. Unwrapping restores it to
// the original 32 bytes.
//
// return a string or nil if failed.
func (z *Crypt2) AesKeyUnwrap(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkCrypt2_aesKeyUnwrap(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Implements the AES Key Wrap Algorithm (RFC 3394). The kek is the Key Encryption
// Key (the AES key used to encrypt the keyData). The arguments and return value are
// binary encoded strings using the encoding specified by encoding (which can be
// "base64", "hex", "base64url", etc.) The full list of supported encodings is
// available at the link below.
// 
// The kek should be an AES key of 16 bytes, 24 bytes, or 32 bytes (i.e. 128-bits,
// 192- bits, or 256-bits). For example, if passed as a hex string, then the kek
// should be 32 chars in length, 48 chars, or 64 chars (because each byte is
// represented as 2 chars in hex).
// 
// The keyData contains the data to be key wrapped. It must be a multiple of 64-bits
// in length. In other words, if the keyData is decoded to binary, it should be a
// number of bytes that is a multiple of 8.
// 
// The return string, if decoded to binary bytes, is equal to the size of the key
// data + 8 additional bytes.
//
// return a string or nil if failed.
func (z *Crypt2) AesKeyWrap(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkCrypt2_aesKeyWrap(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Computes and returns a bcrypt hash of the password. The number of rounds of hashing
// is determined by the BCryptWorkFactor property.
// 
// Starting in v9.5.0.76, if the password is prefixed with "$2b$" then the output will
// use the $2b version of bcrypt. For example, to create a "$2b$" bcrypt has for
// the password "secret", pass in the string "$2b$secret" for password.
//
// return a string or nil if failed.
func (z *Crypt2) BCryptHash(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_bCryptHash(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Verifies the password against a previously computed BCrypt hash. Returns true if
// the password matches the bcryptHash. Returns false if the password does not match.
func (z *Crypt2) BCryptVerify(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_BCryptVerify(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Utility method to convert bytes to a string -- interpreting the bytes according
// to the charset specified.
// return a string or nil if failed.
func (z *Crypt2) BytesToString(arg1 []byte, arg2 string) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    retStr := C.CkCrypt2_bytesToString(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Convenience method for byte swapping between little-endian byte ordering and
// big-endian byte ordering.
func (z *Crypt2) ByteSwap4321(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_ByteSwap4321(z.ckObj, ckbyd1, ckOutBytes)

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


// File-to-file decryption. There is no limit to the size of the file that can be
// decrypted because the component will operate in streaming mode internally.
func (z *Crypt2) CkDecryptFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_CkDecryptFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CkDecryptFile method
func (z *Crypt2) CkDecryptFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkCrypt2_CkDecryptFileAsync(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// File-to-file encryption. There is no limit to the size of the file that can be
// encrypted because the component will operate in streaming mode internally.
func (z *Crypt2) CkEncryptFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_CkEncryptFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CkEncryptFile method
func (z *Crypt2) CkEncryptFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkCrypt2_CkEncryptFileAsync(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Clears the internal list of digital certificates to be used for public-key
// encryption.
func (z *Crypt2) ClearEncryptCerts()  {

    C.CkCrypt2_ClearEncryptCerts(z.ckObj)



}


// Clears the set of certificates to be used in signing.
func (z *Crypt2) ClearSigningCerts()  {

    C.CkCrypt2_ClearSigningCerts(z.ckObj)



}


// Bzip2 compresses a byte array and returns the compressed bytes.
// 
// This is a legacy method that should not be used in new development. It will not
// be marked as deprecated or removed from future APIs because existing
// applications may have data already compressed using this method.
// 
// The output of this method includes an 8-byte header composed of a 4-byte magic
// number (0xB394A7E1) and the 4-byte length of the uncompressed data.
//
func (z *Crypt2) CompressBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_CompressBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Same as CompressBytes, except an encoded string is returned. The output encoding
// is specified by the EncodingMode property.
// return a string or nil if failed.
func (z *Crypt2) CompressBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCrypt2_compressBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Compresses a string and returns the compressed bytes. Prior to compressing, the
// string is converted to a byte representation such as utf-8, utf-16, etc. as
// determined by the Charset property. Otherwise, this method is the same as the
// CompressBytes method.
func (z *Crypt2) CompressString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_CompressString(z.ckObj, cstr1, ckOutBytes)

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


// Compresses a string and returns the encoded compressed bytes. Prior to
// compressing, the string is converted to a byte representation such as utf-8,
// utf-16, etc. as determined by the Charset property. The output encoding is
// specified by the EncodingMode property. Otherwise, this method is the same as
// the CompressBytes method.
// return a string or nil if failed.
func (z *Crypt2) CompressStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_compressStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Calculates a CRC for in-memory byte data. To compute the CRC used in the Zip
// file format, pass "CRC-32" for the crcAlg. (The crcAlg argument provides the
// flexibility to add additional CRC algorithms on an as-needed basis in the
// future.)
func (z *Crypt2) CrcBytes(arg1 string, arg2 []byte) uint {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkCrypt2_CrcBytes(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return uint(v)
}


// Calculates a CRC for the contents of a file. To compute the CRC used in the Zip
// file format, pass "CRC-32" for the crcAlg. (The crcAlg argument provides the
// flexibility to add additional CRC algorithms on an as-needed basis in the
// future.) A value of 0 is returned if the file is unable to be read. Given that
// there is a 1 in 4 billion chance of having an actual CRC of 0, an application
// might choose to react to a 0 return value by testing to see if the file can be
// opened and read.
func (z *Crypt2) CrcFile(arg1 string, arg2 string) uint {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_CrcFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return uint(v)
}


// Asynchronous version of the CrcFile method
func (z *Crypt2) CrcFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkCrypt2_CrcFileAsync(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a file and writes the digital signature to a separate output
// file (a PKCS#7 signature file). The input file (inFilePath) is unmodified. A
// certificate for signing must be specified by calling SetSigningCert or
// SetSigningCert2 prior to calling this method.
// 
// This method is equivalent to CreateP7S. The CreateP7S method was added to
// clarify the format of the signature file that is created.
//
func (z *Crypt2) CreateDetachedSignature(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_CreateDetachedSignature(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Digitally signs a file and creates a .p7m (PKCS #7 Message) file that contains
// both the signature and original file content. The input file (inFilename) is
// unmodified. A certificate for signing must be specified by calling
// SetSigningCert or SetSigningCert2 prior to calling this method.
// 
// To sign with a particular hash algorithm, set the HashAlgorithm property. Valid
// hash algorithms for signing are "sha256", "sha1", "sha384", "sha512", "md5", and
// "md2".
// 
// Note: The CreateP7M method creates an opaque signature. To do the same thing
// entirely in memory, your application would call any of the OpaqueSign* methods,
// such as OpaqueSignBd, OpaqueSignString, OpaqueSignStringENC, etc.
//
func (z *Crypt2) CreateP7M(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_CreateP7M(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CreateP7M method
func (z *Crypt2) CreateP7MAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkCrypt2_CreateP7MAsync(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a file and creates a .p7s (PKCS #7 Signature) signature file.
// The input file (inFilename) is unmodified. The output file (p7sPath) contains only the
// signature and not the original data. A certificate for signing must be specified
// by calling SetSigningCert or SetSigningCert2 prior to calling this method.
// 
// To sign with a particular hash algorithm, set the HashAlgorithm property. Valid
// hash algorithms for signing are "sha256", "sha1", "sha384", "sha512", "md5", and
// "md2".
// 
// Note: The CreateP7S method creates a detached signature. To do the same thing
// entirely in memory, your application would call any of the Sign* methods, such
// as SignBdENC, SignString, SignStringENC, SignSbENC, etc.
//
func (z *Crypt2) CreateP7S(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_CreateP7S(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the CreateP7S method
func (z *Crypt2) CreateP7SAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkCrypt2_CreateP7SAsync(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Decode binary data from an encoded string. The encoding can be set to any of the
// following strings: "base64", "hex", "quoted-printable", "url", "base32", "Q",
// "B", "url_rc1738", "url_rfc2396", "url_rfc3986", "url_oauth", "uu", "modBase64",
// or "html" (for HTML entity encoding).
func (z *Crypt2) Decode(arg1 string, arg2 string) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_Decode(z.ckObj, cstr1, cstr2, ckOutBytes)

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


// Decodes from an encoding back to the original string. The encoding can be set to any
// of the following strings: "base64", "hex", "quoted-printable", "url", "base32",
// "Q", "B", "url_rc1738", "url_rfc2396", "url_rfc3986", "url_oauth", "uu",
// "modBase64", or "html" (for HTML entity encoding).
// return a string or nil if failed.
func (z *Crypt2) DecodeString(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkCrypt2_decodeString(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// In-place decrypts the contents of bd. The minimal set of properties that
// should be set before decrypting are: CryptAlgorithm, SecretKey. Other properties
// that control encryption are: CipherMode, PaddingScheme, KeyLength, IV.
func (z *Crypt2) DecryptBd(arg1 *BinData) bool {

    v := C.CkCrypt2_DecryptBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Decrypts a byte array and returns the unencrypted byte array. The property
// settings used when encrypting the data must match the settings when decrypting.
// Specifically, the CryptAlgorithm, CipherMode, PaddingScheme, KeyLength, IV, and
// SecretKey properties must match.
func (z *Crypt2) DecryptBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_DecryptBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Decrypts string-encoded encrypted data and returns the unencrypted byte array.
// Data encrypted with EncryptBytesENC can be decrypted with this method. The
// property settings used when encrypting the data must match the settings when
// decrypting. Specifically, the EncodingMode, CryptAlgorithm, CipherMode,
// PaddingScheme, KeyLength, IV, and SecretKey properties must match.
func (z *Crypt2) DecryptBytesENC(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_DecryptBytesENC(z.ckObj, cstr1, ckOutBytes)

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


// Encrypted data is passed to this method as an encoded string (base64, hex,
// etc.). This method first decodes the input data according to the EncodingMode
// property setting. It then decrypts and re-encodes using the EncodingMode
// setting, and returns the decrypted data in encoded string form.
// return a string or nil if failed.
func (z *Crypt2) DecryptEncoded(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_decryptEncoded(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Decrypts the contents of bdIn to sbOut. The decrypted string is appended to sbOut.
// The minimal set of properties that should be set before ecrypting are:
// CryptAlgorithm, SecretKey. Other properties that control encryption are:
// CipherMode, PaddingScheme, KeyLength, IV.
func (z *Crypt2) DecryptSb(arg1 *BinData, arg2 *StringBuilder) bool {

    v := C.CkCrypt2_DecryptSb(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Identical to DecryptStringENC, except the decrypts the cipherText and appends the
// decrypted string to the secureStr.
func (z *Crypt2) DecryptSecureENC(arg1 string, arg2 *SecureString) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCrypt2_DecryptSecureENC(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Decrypts a stream. Internally, the strm's source is read, decrypted, and the
// decrypted data written to the strm's sink. It does this in streaming fashion.
// Extremely large or even infinite streams can be decrypted with stable ungrowing
// memory usage.
func (z *Crypt2) DecryptStream(arg1 *Stream) bool {

    v := C.CkCrypt2_DecryptStream(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the DecryptStream method
func (z *Crypt2) DecryptStreamAsync(arg1 *Stream, c chan *Task) {

    hTask := C.CkCrypt2_DecryptStreamAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// The reverse of EncryptString.
// 
// Decrypts encrypted byte data and returns the original string. The property
// settings used when encrypting the string must match the settings when
// decrypting. Specifically, the Charset, CryptAlgorithm, CipherMode,
// PaddingScheme, KeyLength, IV, and SecretKey properties must match.
//
// return a string or nil if failed.
func (z *Crypt2) DecryptString(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCrypt2_decryptString(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The reverse of EncryptStringENC.
// 
// Decrypts string-encoded encrypted data and returns the original string. The
// property settings used when encrypting the string must match the settings when
// decrypting. Specifically, the Charset, EncodingMode, CryptAlgorithm, CipherMode,
// PaddingScheme, KeyLength, IV, and SecretKey properties must match.
//
// return a string or nil if failed.
func (z *Crypt2) DecryptStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_decryptStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Encode binary data to base64, hex, quoted-printable, or URL-encoding. The encoding
// can be set to any of the following strings: "base64", "hex", "quoted-printable"
// (or "qp"), "url", "base32", "Q", "B", "url_rc1738", "url_rfc2396",
// "url_rfc3986", "url_oauth", "uu", "modBase64", or "html" (for HTML entity
// encoding).
// return a string or nil if failed.
func (z *Crypt2) Encode(arg1 []byte, arg2 string) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    retStr := C.CkCrypt2_encode(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Encodes an integer to N bytes and returns in the specified encoding. If littleEndian is
// true, then little endian byte ordering is used. Otherwise big-endian byte
// order is used.
// return a string or nil if failed.
func (z *Crypt2) EncodeInt(arg1 int, arg2 int, arg3 bool, arg4 string) *string {
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    cstr4 := C.CString(arg4)

    retStr := C.CkCrypt2_encodeInt(z.ckObj, C.int(arg1), C.int(arg2), b3, cstr4)

    C.free(unsafe.Pointer(cstr4))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Encodes a string. The toEncodingName can be set to any of the following strings: "base64",
// "hex", "quoted-printable", "url", "base32", "Q", "B", "url_rc1738",
// "url_rfc2396", "url_rfc3986", "url_oauth", "uu", "modBase64", or "html" (for
// HTML entity encoding). The charsetName is important, and usually you'll want to specify
// "ansi". For example, if the string "ABC" is to be encoded to "hex" using ANSI,
// the result will be "414243". However, if "unicode" is used, the result is
// "410042004300".
// return a string or nil if failed.
func (z *Crypt2) EncodeString(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkCrypt2_encodeString(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// In-place encrypts the contents of bd. The minimal set of properties that
// should be set before encrypting are: CryptAlgorithm, SecretKey. Other properties
// that control encryption are: CipherMode, PaddingScheme, KeyLength, IV. When
// decrypting, all property settings must match otherwise the result is garbled
// data.
func (z *Crypt2) EncryptBd(arg1 *BinData) bool {

    v := C.CkCrypt2_EncryptBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Encrypts a byte array. The minimal set of properties that should be set before
// encrypting are: CryptAlgorithm, SecretKey. Other properties that control
// encryption are: CipherMode, PaddingScheme, KeyLength, IV. When decrypting, all
// property settings must match otherwise garbled data is returned.
func (z *Crypt2) EncryptBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_EncryptBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Encrypts a byte array and returns the encrypted data as an encoded (printable)
// string. The minimal set of properties that should be set before encrypting are:
// CryptAlgorithm, SecretKey, EncodingMode. Other properties that control
// encryption are: CipherMode, PaddingScheme, KeyLength, IV. When decrypting, all
// property settings must match otherwise garbled data is returned. The encoding of
// the string that is returned is controlled by the EncodingMode property, which
// can be set to "Base64", "QP", or "Hex".
// return a string or nil if failed.
func (z *Crypt2) EncryptBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCrypt2_encryptBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The input string is first decoded according to the encoding algorithm specified
// by the EncodingMode property (such as base64, hex, etc.) It is then encrypted
// according to the encryption algorithm specified by CryptAlgorithm. The resulting
// encrypted data is encoded (using EncodingMode) and returned.
// return a string or nil if failed.
func (z *Crypt2) EncryptEncoded(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_encryptEncoded(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Encrypts the contents of sbIn to bdOut. The minimal set of properties that should
// be set before ecrypting are: CryptAlgorithm, SecretKey. Other properties that
// control encryption are: CipherMode, PaddingScheme, KeyLength, IV.
func (z *Crypt2) EncryptSb(arg1 *StringBuilder, arg2 *BinData) bool {

    v := C.CkCrypt2_EncryptSb(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Identical to EncryptStringENC, except the clear-text contained within the secureStr
// is encrypted and returned.
// return a string or nil if failed.
func (z *Crypt2) EncryptSecureENC(arg1 *SecureString) *string {

    retStr := C.CkCrypt2_encryptSecureENC(z.ckObj, arg1.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Encrypts a stream. Internally, the strm's source is read, encrypted, and the
// encrypted data written to the strm's sink. It does this in streaming fashion.
// Extremely large or even infinite streams can be encrypted with stable ungrowing
// memory usage.
func (z *Crypt2) EncryptStream(arg1 *Stream) bool {

    v := C.CkCrypt2_EncryptStream(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the EncryptStream method
func (z *Crypt2) EncryptStreamAsync(arg1 *Stream, c chan *Task) {

    hTask := C.CkCrypt2_EncryptStreamAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Encrypts a string and returns the encrypted data as a byte array. The minimal
// set of properties that should be set before encrypting are: CryptAlgorithm,
// SecretKey, Charset. Other properties that control encryption are: CipherMode,
// PaddingScheme, KeyLength, IV. When decrypting, all property settings must match
// otherwise garbled data is returned. The Charset property controls the exact
// bytes that get encrypted. Languages such as VB.NET, C#, and Visual Basic work
// with Unicode strings, thus the input string is Unicode. If Unicode is to be
// encrypted (i.e. 2 bytes per character) then set the Charset property to
// "Unicode". To implicitly convert the string to another charset before the
// encryption is applied, set the Charset property to something else, such as
// "iso-8859-1", "Shift_JIS", "big5", "windows-1252", etc. The complete list of
// possible charsets is listed here:
// us-ascii
// unicode
// unicodefffe
// iso-8859-1
// iso-8859-2
// iso-8859-3
// iso-8859-4
// iso-8859-5
// iso-8859-6
// iso-8859-7
// iso-8859-8
// iso-8859-9
// iso-8859-13
// iso-8859-15
// windows-874
// windows-1250
// windows-1251
// windows-1252
// windows-1253
// windows-1254
// windows-1255
// windows-1256
// windows-1257
// windows-1258
// utf-7
// utf-8
// utf-32
// utf-32be
// shift_jis
// gb2312
// ks_c_5601-1987
// big5
// iso-2022-jp
// iso-2022-kr
// euc-jp
// euc-kr
// macintosh
// x-mac-japanese
// x-mac-chinesetrad
// x-mac-korean
// x-mac-arabic
// x-mac-hebrew
// x-mac-greek
// x-mac-cyrillic
// x-mac-chinesesimp
// x-mac-romanian
// x-mac-ukrainian
// x-mac-thai
// x-mac-ce
// x-mac-icelandic
// x-mac-turkish
// x-mac-croatian
// asmo-708
// dos-720
// dos-862
// ibm037
// ibm437
// ibm500
// ibm737
// ibm775
// ibm850
// ibm852
// ibm855
// ibm857
// ibm00858
// ibm860
// ibm861
// ibm863
// ibm864
// ibm865
// cp866
// ibm869
// ibm870
// cp875
// koi8-r
// koi8-u
func (z *Crypt2) EncryptString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_EncryptString(z.ckObj, cstr1, ckOutBytes)

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


// Encrypts a string and returns the encrypted data as an encoded (printable)
// string. The minimal set of properties that should be set before encrypting are:
// CryptAlgorithm, SecretKey, Charset, and EncodingMode. Other properties that
// control encryption are: CipherMode, PaddingScheme, KeyLength, IV. When
// decrypting (with DecryptStringENC), all property settings must match otherwise
// garbled data is returned. The Charset property controls the exact bytes that get
// encrypted. Languages such as VB.NET, C#, and Visual Basic work with Unicode
// strings, thus the input string is Unicode. If Unicode is to be encrypted (i.e. 2
// bytes per character) then set the Charset property to "Unicode". To implicitly
// convert the string to another charset before the encryption is applied, set the
// Charset property to something else, such as "iso-8859-1", "Shift_JIS", "big5",
// "windows-1252", etc. (Refer to EncryptString for the complete list of charsets.)
// 
// The EncodingMode property controls the encoding of the string that is returned.
// It can be set to "Base64", "QP", or "Hex".
//
// return a string or nil if failed.
func (z *Crypt2) EncryptStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_encryptStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Important: In the v9.5.0.49 release, a bug involving this method was introduced:
// The encoding is ignored and instead the encoding used is the current value of the
// EncodingMode property. The workaround is to make sure the EncodingMode property
// is set to the value of the desired output encoding. This problem will be fixed
// in v9.5.0.50.
// 
// Identical to the GenerateSecretKey method, except it returns the binary secret
// key as a string encoded according to encoding, which may be "base64", "hex", "url",
// etc. Please see the documentation for GenerateSecretKey for more information.
//
// return a string or nil if failed.
func (z *Crypt2) GenEncodedSecretKey(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkCrypt2_genEncodedSecretKey(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Hashes a string to a byte array that has the same number of bits as the current
// value of the KeyLength property. For example, if KeyLength is equal to 128 bits,
// then a 16-byte array is returned. This can be used to set the SecretKey
// property. In order to decrypt, the SecretKey must match exactly. To use
// "password-based" encryption, the password is passed to this method to generate a
// binary secret key that can then be assigned to the SecretKey property.
// 
// IMPORTANT: If you are trying to decrypt something encrypted by another party
// such that the other party provided you with the secret key, DO NOT use this
// method. This method is for transforming an arbitrary-length password into a
// binary secret key of the proper length. Please see this Chilkat blog
// post:Getting Started with AES Decryption
// <http://www.cknotes.com/?p=290>
//
func (z *Crypt2) GenerateSecretKey(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_GenerateSecretKey(z.ckObj, cstr1, ckOutBytes)

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


// Generates a random UUID string having standard UUID format, such as
// "de305d54-75b4-431b-adb2-eb6b9e546014".
// 
// Note: This generates a "version 4 UUID" using random byte values. See RFC 4122.
//
// return a string or nil if failed.
func (z *Crypt2) GenerateUuid() *string {

    retStr := C.CkCrypt2_generateUuid(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Generates numBytes random bytes and returns them as an encoded string. The encoding,
// such as base64, hex, etc. is controlled by the EncodingMode property.
// return a string or nil if failed.
func (z *Crypt2) GenRandomBytesENC(arg1 int) *string {

    retStr := C.CkCrypt2_genRandomBytesENC(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the last certificate used for public-key decryption.
func (z *Crypt2) GetDecryptCert() *Cert {

    retObj := C.CkCrypt2_GetDecryptCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the authenticated additional data as an encoded string. The encoding
// argument can be set to any of the following strings: "base64", "hex",
// "quoted-printable", or "url".
// 
// The Aad is used when the CipherMode is "gcm" (Galois/Counter Mode), which is a
// mode valid for symmetric ciphers that have a block size of 16 bytes, such as AES
// or Twofish.
//
// return a string or nil if failed.
func (z *Crypt2) GetEncodedAad(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_getEncodedAad(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the authentication tag as an encoded string. The encoding argument may be
// set to any of the following strings: "base64", "hex", "quoted-printable", or
// "url". The authentication tag is an output of authenticated encryption modes
// such as GCM when encrypting. When GCM mode decrypting, the authenticate tag is
// set by the application and is the expected result.
// 
// The authenticated tag plays a role when the CipherMode is "gcm" (Galois/Counter
// Mode), which is a mode valid for symmetric block ciphers that have a block size
// of 16 bytes, such as AES or Twofish.
//
// return a string or nil if failed.
func (z *Crypt2) GetEncodedAuthTag(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_getEncodedAuthTag(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the initialization vector as an encoded string. The encoding argument
// can be set to any of the following strings: "base64", "hex", "quoted-printable",
// or "url".
// return a string or nil if failed.
func (z *Crypt2) GetEncodedIV(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_getEncodedIV(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the secret key as an encoded string. The encoding argument can be set to
// any of the following strings: "base64", "hex", "quoted-printable", or "url".
// return a string or nil if failed.
func (z *Crypt2) GetEncodedKey(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_getEncodedKey(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the password-based encryption (PBE) salt bytes as an encoded string. The
// encoding argument can be set to any of the following strings: "base64", "hex",
// "quoted-printable", or "url".
// return a string or nil if failed.
func (z *Crypt2) GetEncodedSalt(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_getEncodedSalt(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the last certificate used when verifying a signature. This method is
// deprecated. Applications should instead call GetSignerCert with an index of 0.
func (z *Crypt2) GetLastCert() *Cert {

    retObj := C.CkCrypt2_GetLastCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// The same as GetSignatureSigningTime, except the date/time is returned in RFC822
// string format.
// return a string or nil if failed.
func (z *Crypt2) GetSignatureSigningTimeStr(arg1 int) *string {

    retStr := C.CkCrypt2_getSignatureSigningTimeStr(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Extracts the signed (authenticated) attributes for the Nth signer. In most
// cases, a signature has only one signer, and the signerIndex should equal 0 to specify
// the 1st (and only) signer.
// 
// The binary PKCS7 is passed in pkcs7Der. On success, the sbJson will contain the signed
// attributes in JSON format.
// 
// Sample JSON output:
// {
//   "signedAttributes": [
//     {
//       "oid": "1.2.840.113549.1.9.3",
//       "name": "Content Type"
//     },
//     {
//       "oid": "1.2.840.113549.1.9.5",
//       "name": "Signing Time"
//     },
//     {
//       "oid": "1.2.840.113549.1.9.4",
//       "name": "Message Digest"
//     },
//     {
//       "oid": "1.2.840.113549.1.9.16.2.47",
//       "name": "Signing Certificate V2"
//     }
//   ]
// }
//
func (z *Crypt2) GetSignedAttributes(arg1 int, arg2 *BinData, arg3 *StringBuilder) bool {

    v := C.CkCrypt2_GetSignedAttributes(z.ckObj, C.int(arg1), arg2.ckObj, arg3.ckObj)


    return v != 0
}


// Gets the Nth certificate used for signing. This method can be called after
// verifying a digital signature to get the signer certs. The 1st certificate is at
// index 0. The NumSignerCerts property contains the total number of signing
// certificates. (Typically, a single certificate is used in creating a digital
// signature.)
func (z *Crypt2) GetSignerCert(arg1 int) *Cert {

    retObj := C.CkCrypt2_GetSignerCert(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the full certificate chain for the Nth certificate used to for signing.
// Indexing begins at 0.
func (z *Crypt2) GetSignerCertChain(arg1 int) *CertChain {

    retObj := C.CkCrypt2_GetSignerCertChain(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &CertChain{retObj}
}


// Hashes the the bytes contained in bd and returns the hash as an encoded
// string.
// 
// The hash algorithm is specified by the HashAlgorithm property, The encoding is
// controlled by the EncodingMode property, which can be set to "base64", "hex",
// "base64url", or any of the encodings listed at the link below.
//
// return a string or nil if failed.
func (z *Crypt2) HashBdENC(arg1 *BinData) *string {

    retStr := C.CkCrypt2_hashBdENC(z.ckObj, arg1.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Begin hashing a byte stream. Call this method to hash the 1st chunk. Additional
// chunks are hashed by calling HashMoreBytes 0 or more times followed by a final
// call to HashFinal (or HashFinalENC) to retrieve the result. The hash algorithm
// is selected by the HashAlgorithm property setting.
func (z *Crypt2) HashBeginBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkCrypt2_HashBeginBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Begin hashing a text stream. Call this method to hash the 1st chunk. Additional
// chunks are hashed by calling HashMoreString 0 or more times followed by a final
// call to HashFinal (or HashFinalENC) to retrieve the result. The hash algorithm
// is selected by the HashAlgorithm property setting.
func (z *Crypt2) HashBeginString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCrypt2_HashBeginString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Hashes a byte array.
// 
// The hash algorithm is specified by the HashAlgorithm property, The encoding is
// controlled by the EncodingMode property, which can be set to "base64", "hex",
// "base64url", or any of the encodings listed at the link below.
//
func (z *Crypt2) HashBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_HashBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Hashes a byte array and returns the hash as an encoded string.
// 
// The hash algorithm is specified by the HashAlgorithm property, The encoding is
// controlled by the EncodingMode property, which can be set to "base64", "hex",
// "base64url", or any of the encodings listed at the link below.
//
// return a string or nil if failed.
func (z *Crypt2) HashBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCrypt2_hashBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Hashes a file and returns the hash bytes.
// 
// The hash algorithm is specified by the HashAlgorithm property,
// 
// Any size file may be hashed because the file is hashed internally in streaming
// mode (keeping memory usage low and constant).
//
func (z *Crypt2) HashFile(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_HashFile(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the HashFile method
func (z *Crypt2) HashFileAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCrypt2_HashFileAsync(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Hashes a file and returns the hash as an encoded string.
// 
// The hash algorithm is specified by the HashAlgorithm property, The encoding is
// controlled by the EncodingMode property, which can be set to "base64", "hex",
// "base64url", or any of the encodings listed at the link below.
// 
// Any size file is supported because the file is hashed internally in streaming
// mode (keeping memory usage low and constant).
//
// return a string or nil if failed.
func (z *Crypt2) HashFileENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_hashFileENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the HashFileENC method
func (z *Crypt2) HashFileENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCrypt2_HashFileENCAsync(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Finalizes a multi-step hash computation and returns the hash bytes.
func (z *Crypt2) HashFinal() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_HashFinal(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Finalizes a multi-step hash computation and returns the hash bytes encoded
// according to the EncodingMode property setting.
// return a string or nil if failed.
func (z *Crypt2) HashFinalENC() *string {

    retStr := C.CkCrypt2_hashFinalENC(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Adds more bytes to the hash currently under computation. (See HashBeginBytes)
func (z *Crypt2) HashMoreBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkCrypt2_HashMoreBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Adds more text to the hash currently under computation. (See HashBeginString)
func (z *Crypt2) HashMoreString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCrypt2_HashMoreString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Hashes a string and returns a binary hash. The hash algorithm is specified by
// the HashAlgorithm property,
// 
// The Charset property controls the character encoding of the string that is
// hashed. Languages such as VB.NET, C#, and Visual Basic work with Unicode
// strings. If it is desired to hash Unicode directly (2 bytes/char) then set the
// Charset property to "Unicode". To implicitly convert to another charset before
// hashing, set the Charset property to the desired charset. For example, if
// Charset is set to "iso-8859-1", the input string is first implicitly converted
// to iso-8859-1 (1 byte per character) before hashing. The full list fo supported
// charsets is listed in the EncryptString method description.
// 
// IMPORTANT: Hash algorithms hash bytes. Changing the bytes passed to a hash
// algorithm changes the result. A character (i.e. a visible glyph) can have
// different byte representations. The byte representation is defined by the
// Charset. For example, 'A' in us-ascii is a single byte 0x41, whereas in utf-16
// it is 2 bytes (0x41 0x00). The byte representation should be explicitly
// specified, otherwise unexpected results may occur.
//
func (z *Crypt2) HashString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_HashString(z.ckObj, cstr1, ckOutBytes)

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


// Hashes a string and returns the hash bytes as an encoded string.
// 
// The hash algorithm is specified by the HashAlgorithm property, The encoding is
// controlled by the EncodingMode property, which can be set to "base64", "hex",
// "base64url", or any of the encodings listed at the link below.
// 
// The Charset property controls the character encoding of the string that is
// hashed. Languages such as VB.NET, C#, and Visual Basic work with Unicode
// strings. If it is desired to hash Unicode directly (2 bytes/char) then set the
// Charset property to "Unicode". To implicitly convert to another charset before
// hashing, set the Charset property to the desired charset. For example, if
// Charset is set to "iso-8859-1", the input string is first implicitly converted
// to iso-8859-1 (1 byte per character) before hashing. The full list of supported
// charsets is listed in the EncryptString method description.
//
// return a string or nil if failed.
func (z *Crypt2) HashStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_hashStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// This method can be called after a digital signature has been verified by one of
// the Verify* methods. Returns true if a signing time for the Nth certificate is
// available and can be retrieved by either the GetSignatureSigningTime or
// GetSignatureSigningTimeStr methods.
func (z *Crypt2) HasSignatureSigningTime(arg1 int) bool {

    v := C.CkCrypt2_HasSignatureSigningTime(z.ckObj, C.int(arg1))


    return v != 0
}


// Implements RFC 4226: HOTP: An HMAC-Based One-Time Password Algorithm. The
// arguments to this method are:
//     secret: The shared secret in an enocded representation such as base64, hex,
//     ascii, etc.
//     secretEnc: The encoding of the shared secret, such as "base64"
//     counterHex: The 8-byte counter in hexidecimal format.
//     numDigits: The number of decimal digits to return.
//     truncOffset: Normally set this to -1 for dynamic truncation. Otherwise can be set
//     in the range 0..15.
//     hashAlg: Normally set to "sha1". Can be set to other hash algorithms such as
//     "sha256", "sha512", etc.
// return a string or nil if failed.
func (z *Crypt2) Hotp(arg1 string, arg2 string, arg3 string, arg4 int, arg5 int, arg6 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr6 := C.CString(arg6)

    retStr := C.CkCrypt2_hotp(z.ckObj, cstr1, cstr2, cstr3, C.int(arg4), C.int(arg5), cstr6)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr6))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Decompresses data that was compressed with CompressBytes.
// 
// This is a legacy method that should not be used in new development. It will not
// be marked as deprecated or removed from future APIs because existing
// applications may have data already compressed using CompressBytes.
// 
// This method expects the input to begin with an 8-byte header composed of a
// 4-byte magic number (0xB394A7E1) and the 4-byte length of the uncompressed data.
//
func (z *Crypt2) InflateBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_InflateBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// The opposite of CompressBytesENC. The EncodingMode and CompressionAlgorithm
// properties should match what was used when compressing.
func (z *Crypt2) InflateBytesENC(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_InflateBytesENC(z.ckObj, cstr1, ckOutBytes)

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


// The opposite of CompressString. The Charset and CompressionAlgorithm properties
// should match what was used when compressing.
// return a string or nil if failed.
func (z *Crypt2) InflateString(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCrypt2_inflateString(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The opposite of CompressStringENC. The Charset, EncodingMode, and
// CompressionAlgorithm properties should match what was used when compressing.
// return a string or nil if failed.
func (z *Crypt2) InflateStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_inflateStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if the component is unlocked.
func (z *Crypt2) IsUnlocked() bool {

    v := C.CkCrypt2_IsUnlocked(z.ckObj)


    return v != 0
}


// Provides information about what transpired in the last method called. For many
// methods, there is no information. For some methods, details about what
// transpired can be obtained via LastJsonData. For example, after calling a method
// to verify a signature, the LastJsonData will return JSON with details about the
// algorithms used for signature verification.
func (z *Crypt2) LastJsonData() *JsonObject {

    retObj := C.CkCrypt2_LastJsonData(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &JsonObject{retObj}
}


// Loads the caller of the task's async method.
func (z *Crypt2) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkCrypt2_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Computes a Message Authentication Code on the bytes contained in bd, using the
// MAC algorithm specified in the MacAlgorithm property. The result is encoded to a
// string using the encoding (base64, hex, etc.) specified by the EncodingMode
// property.
// return a string or nil if failed.
func (z *Crypt2) MacBdENC(arg1 *BinData) *string {

    retStr := C.CkCrypt2_macBdENC(z.ckObj, arg1.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Computes a Message Authentication Code using the MAC algorithm specified in the
// MacAlgorithm property.
func (z *Crypt2) MacBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_MacBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Computes a Message Authentication Code using the MAC algorithm specified in the
// MacAlgorithm property. The result is encoded to a string using the encoding
// (base64, hex, etc.) specified by the EncodingMode property.
// return a string or nil if failed.
func (z *Crypt2) MacBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCrypt2_macBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Computes a Message Authentication Code using the MAC algorithm specified in the
// MacAlgorithm property.
func (z *Crypt2) MacString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_MacString(z.ckObj, cstr1, ckOutBytes)

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


// Computes a Message Authentication Code using the MAC algorithm specified in the
// MacAlgorithm property. The result is encoded to a string using the encoding
// (base64, hex, etc.) specified by the EncodingMode property.
// return a string or nil if failed.
func (z *Crypt2) MacStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_macStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Matches MySQL's AES_DECRYPT function. strEncryptedHex is a hex-encoded string of the AES
// encrypted data. The return value is the original unencrypted string.
// return a string or nil if failed.
func (z *Crypt2) MySqlAesDecrypt(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkCrypt2_mySqlAesDecrypt(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Matches MySQL's AES_ENCRYPT function. The return value is a hex-encoded string
// of the encrypted data. The equivalent call in MySQL would look like this:
// HEX(AES_ENCRYPT('The quick brown fox jumps over the lazy dog','password'))
// return a string or nil if failed.
func (z *Crypt2) MySqlAesEncrypt(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkCrypt2_mySqlAesEncrypt(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// In-place signs the contents of bd. The contents of bd is replaced with the
// PKCS7/CMS format signature that embeds the data that was signed.
func (z *Crypt2) OpaqueSignBd(arg1 *BinData) bool {

    v := C.CkCrypt2_OpaqueSignBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Asynchronous version of the OpaqueSignBd method
func (z *Crypt2) OpaqueSignBdAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkCrypt2_OpaqueSignBdAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a byte array and returns a PKCS7/CMS format signature. This is a
// signature that contains both the original data as well as the signature. A
// certificate must be set by calling SetSigningCert prior to calling this method.
func (z *Crypt2) OpaqueSignBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_OpaqueSignBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the OpaqueSignBytes method
func (z *Crypt2) OpaqueSignBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCrypt2_OpaqueSignBytesAsync(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a byte array and returns a PKCS7/CMS format signature in encoded
// string format (such as Base64 or hex). This is a signature that contains both
// the original data as well as the signature. A certificate must be set by calling
// SetSigningCert prior to calling this method. The EncodingMode property controls
// the output encoding, which can be "Base64", "QP","Hex", etc. (See the
// EncodingMode property.)
// return a string or nil if failed.
func (z *Crypt2) OpaqueSignBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCrypt2_opaqueSignBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the OpaqueSignBytesENC method
func (z *Crypt2) OpaqueSignBytesENCAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCrypt2_OpaqueSignBytesENCAsync(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a string and returns a PKCS7/CMS format signature. This is a
// signature that contains both the original data as well as the signature. A
// certificate must be set by calling SetSigningCert prior to calling this method.
// The Charset property controls the character encoding of the string that is
// signed. (Languages such as VB.NET, C#, and Visual Basic work with Unicode
// strings.) To sign Unicode data (2 bytes per char), set the Charset property to
// "Unicode". To implicitly convert the string to a mutlibyte charset such as
// "iso-8859-1", "Shift_JIS", "utf-8", or something else, then set the Charset
// property to the name of the charset before signing. The complete list of
// charsets is listed in the EncryptString method description.
func (z *Crypt2) OpaqueSignString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_OpaqueSignString(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the OpaqueSignString method
func (z *Crypt2) OpaqueSignStringAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCrypt2_OpaqueSignStringAsync(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a string and returns a PKCS7/CMS format signature in encoded
// string format (such as base64 or hex). This is a signature that contains both
// the original data as well as the signature. A certificate must be set by calling
// SetSigningCert prior to calling this method. The Charset property controls the
// character encoding of the string that is signed. (Languages such as VB.NET, C#,
// and Visual Basic work with Unicode strings.) To sign Unicode data (2 bytes per
// char), set the Charset property to "Unicode". To implicitly convert the string
// to a mutlibyte charset such as "iso-8859-1", "Shift_JIS", "utf-8", or something
// else, then set the Charset property to the name of the charset before signing.
// The complete list of charsets is listed in the EncryptString method description.
// 
// The EncodingMode property controls the output encoding, which can be "Base64",
// "QP","Hex", etc. (See the EncodingMode property.)
//
// return a string or nil if failed.
func (z *Crypt2) OpaqueSignStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_opaqueSignStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the OpaqueSignStringENC method
func (z *Crypt2) OpaqueSignStringENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCrypt2_OpaqueSignStringENCAsync(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// In-place verifies and unwraps the PKCS7/CMS contents of bd. If the signature
// is verified, the contents of bd will be replaced with the original data, and
// the method returns true. If the signature is not verified, then the contents
// of bd remain unchanged and the method returns false.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) OpaqueVerifyBd(arg1 *BinData) bool {

    v := C.CkCrypt2_OpaqueVerifyBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Verifies an opaque signature and returns the original data. If the signature
// verification fails, the returned data will be 0 bytes in length.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) OpaqueVerifyBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_OpaqueVerifyBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Verifies an opaque signature (encoded in string form) and returns the original
// data. If the signature verification fails, the returned data will be 0 bytes in
// length.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) OpaqueVerifyBytesENC(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_OpaqueVerifyBytesENC(z.ckObj, cstr1, ckOutBytes)

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


// Verifies an opaque signature and returns the original string. If the signature
// verification fails, the returned string will be 0 characters in length.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
// return a string or nil if failed.
func (z *Crypt2) OpaqueVerifyString(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCrypt2_opaqueVerifyString(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Verifies an opaque signature (encoded in string form) and returns the original
// data string. If the signature verification fails, the returned string will be 0
// characters in length.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
// return a string or nil if failed.
func (z *Crypt2) OpaqueVerifyStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_opaqueVerifyStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Implements the PBKDF1 algorithm (Password Based Key Derivation Function #1). The
// password is converted to the character encoding represented by charset before being
// passed (internally) to the key derivation function. The hashAlg may be "md5",
// "sha1", "md2", etc. The salt should be random data at least 8 bytes (64 bits) in
// length. (The GenRandomBytesENC method is good for generating a random salt
// value.) The iterationCount should be no less than 1000. The length (in bits) of the
// derived key output by this method is controlled by outputKeyBitLen. The encoding argument may
// be "base64", "hex", etc. It controls the encoding of the output, and the
// expected encoding of the salt. The derived key is returned.
// 
// Note: Starting in version 9.5.0.47, if the charset is set to one of the keywords
// "hex" or "base64", then the password will be considered binary data that is hex
// or base64 encoded. The bytes will be decoded and used directly as a binary
// password.
//
// return a string or nil if failed.
func (z *Crypt2) Pbkdf1(arg1 string, arg2 string, arg3 string, arg4 string, arg5 int, arg6 int, arg7 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr7 := C.CString(arg7)

    retStr := C.CkCrypt2_pbkdf1(z.ckObj, cstr1, cstr2, cstr3, cstr4, C.int(arg5), C.int(arg6), cstr7)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr7))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Implements the PBKDF2 algorithm (Password Based Key Derivation Function #2). The
// password is converted to the character encoding represented by charset before being
// passed (internally) to the key derivation function. The hashAlg may be "sha256",
// "sha384", "sha512", "md5", "sha1", "md2", or any hash algorithm listed in the
// HashAlgorithm property. The salt should be random data at least 8 bytes (64
// bits) in length. (The GenRandomBytesENC method is good for generating a random
// salt value.) The iterationCount should be no less than 1000. The length (in bits) of the
// derived key output by this method is controlled by outputKeyBitLen. The encoding argument may
// be "base64", "hex", etc. It controls the encoding of the output, and the
// expected encoding of the salt. The derived key is returned.
// 
// Note: The PBKDF2 function (internally) utilizes a PRF that is a pseudorandom
// function that is a keyed HMAC. The hash algorithm specified by hashAlg determines
// this PRF. If hashAlg is "SHA256", then HMAC-SHA256 is used for the PRF. Likewise,
// if the hash function is "SHA1", then HMAC-SHA1 is used. HMAC can be used with
// any hash algorithm.
// 
// Note: Starting in version 9.5.0.47, if the charset is set to one of the keywords
// "hex" or "base64", then the password will be considered binary data that is hex
// or base64 encoded. The bytes will be decoded and used directly as a binary
// password.
//
// return a string or nil if failed.
func (z *Crypt2) Pbkdf2(arg1 string, arg2 string, arg3 string, arg4 string, arg5 int, arg6 int, arg7 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr7 := C.CString(arg7)

    retStr := C.CkCrypt2_pbkdf2(z.ckObj, cstr1, cstr2, cstr3, cstr4, C.int(arg5), C.int(arg6), cstr7)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr7))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sets the initialization vector to a random value.
func (z *Crypt2) RandomizeIV()  {

    C.CkCrypt2_RandomizeIV(z.ckObj)



}


// Sets the secret key to a random value.
func (z *Crypt2) RandomizeKey()  {

    C.CkCrypt2_RandomizeKey(z.ckObj)



}


// Convenience method to read an entire file and return as a byte array.
func (z *Crypt2) ReadFile(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_ReadFile(z.ckObj, cstr1, ckOutBytes)

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


// Provides a means for converting from one encoding to another (such as base64 to
// hex). This is helpful for programming environments where byte arrays are a real
// pain-in-the-***. The fromEncoding and toEncoding may be (case-insensitive) "Base64",
// "modBase64", "Base32", "Base58", "UU", "QP" (for quoted-printable), "URL" (for
// url-encoding), "Hex", "Q", "B", "url_oauth", "url_rfc1738", "url_rfc2396", and
// "url_rfc3986".
// return a string or nil if failed.
func (z *Crypt2) ReEncode(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkCrypt2_reEncode(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sets the digital certificate to be used for decryption when the CryptAlgorithm
// property is set to "PKI". A private key is required for decryption. Because this
// method only specifies the certificate, a prerequisite is that the certificate w/
// private key must have been pre-installed on the computer. Private keys are
// stored in the Windows Protected Store (either a user account specific store, or
// the system-wide store). The Chilkat component will automatically locate and find
// the certificate's corresponding private key from the protected store when
// decrypting.
func (z *Crypt2) SetDecryptCert(arg1 *Cert) bool {

    v := C.CkCrypt2_SetDecryptCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the digital certificate to be used for decryption when the CryptAlgorithm
// property is set to "PKI". The private key is supplied in the 2nd argument to
// this method, so there is no requirement that the certificate be pre-installed on
// a computer before decrypting (if this method is called).
func (z *Crypt2) SetDecryptCert2(arg1 *Cert, arg2 *PrivateKey) bool {

    v := C.CkCrypt2_SetDecryptCert2(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Sets the authenticated additional data from an encoded string. The authenticated
// additional data (AAD), if any, is used in authenticated encryption modes such as
// GCM. The aadStr argument can be set to any of the following strings: "base64",
// "hex", "quoted-printable", "ascii", or "url".
// 
// The Aad is used when the CipherMode is "gcm" (Galois/Counter Mode), which is a
// mode valid for symmetric ciphers that have a block size of 16 bytes, such as AES
// or Twofish.
//
func (z *Crypt2) SetEncodedAad(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_SetEncodedAad(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the expected authenticated tag from an encoded string. The authenticated
// tag is used in authenticated encryption modes such as GCM. An application would
// set the expected authenticated tag prior to decrypting. The authTagStr argument can be
// set to any of the following strings: "base64", "hex", "quoted-printable",
// "ascii", or "url".
// 
// The authenticated tag plays a role when the CipherMode is "gcm" (Galois/Counter
// Mode), which is a mode valid for symmetric block ciphers that have a block size
// of 16 bytes, such as AES or Twofish.
//
func (z *Crypt2) SetEncodedAuthTag(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_SetEncodedAuthTag(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the initialization vector from an encoded string. The encoding argument can
// be set to any of the following strings: "base64", "hex", "quoted-printable",
// "ascii", or "url".
func (z *Crypt2) SetEncodedIV(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkCrypt2_SetEncodedIV(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Sets the secret key from an encoded string. The encoding argument can be set to
// any of the following strings: "base64", "hex", "quoted-printable", "ascii", or
// "url".
func (z *Crypt2) SetEncodedKey(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkCrypt2_SetEncodedKey(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Sets the password-based encryption (PBE) salt bytes from an encoded string. The
// encoding argument can be set to any of the following strings: "base64", "hex",
// "quoted-printable", "ascii", or "url".
func (z *Crypt2) SetEncodedSalt(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkCrypt2_SetEncodedSalt(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Tells the encryption library to use a specific digital certificate for
// public-key encryption. To encrypt with multiple certificates, call
// AddEncryptCert once for each certificate. (Calling this method is the equivalent
// of calling ClearEncryptCerts followed by AddEncryptCert.)
func (z *Crypt2) SetEncryptCert(arg1 *Cert) bool {

    v := C.CkCrypt2_SetEncryptCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the MAC key to be used for one of the Mac methods.
func (z *Crypt2) SetMacKeyBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkCrypt2_SetMacKeyBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Sets the MAC key to be used for one of the Mac methods. The encoding can be set to
// any of the following strings: "base64", "hex", "quoted-printable", or "url".
func (z *Crypt2) SetMacKeyEncoded(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_SetMacKeyEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the MAC key to be used for one of the Mac methods.
func (z *Crypt2) SetMacKeyString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCrypt2_SetMacKeyString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Accepts a password string and (internally) generates a binary secret key of the
// appropriate bit length and sets the SecretKey property. This method should only
// be used if you are using Chilkat for both encryption and decryption because the
// password-to-secret-key algorithm would need to be identical for the decryption
// to match the encryption.
// 
// There is no minimum or maximum password length. The password string is
// transformed to a binary secret key by computing the MD5 digest (of the utf-8
// password) to obtain 16 bytes. If the KeyLength is greater than 16 bytes, then
// the MD5 digest of the Base64 encoding of the utf-8 password is added. A max of
// 32 bytes of key material is generated, and this is truncated to the actual
// KeyLength required. The example below shows how to manually duplicate the
// computation.
//
func (z *Crypt2) SetSecretKeyViaPassword(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkCrypt2_SetSecretKeyViaPassword(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Specifies a certificate to be used when creating PKCS7 digital signatures.
// Signing requires both a certificate and private key. In this case, the private
// key is implicitly specified if the certificate originated from a PFX that
// contains the corresponding private key, or if on a Windows-based computer where
// the certificate and corresponding private key are pre-installed. (If a PFX file
// is used, it is provided via the AddPfxSourceFile or AddPfxSourceData methods.)
func (z *Crypt2) SetSigningCert(arg1 *Cert) bool {

    v := C.CkCrypt2_SetSigningCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Specifies a digital certificate and private key to be used for creating PKCS7
// digital signatures.
func (z *Crypt2) SetSigningCert2(arg1 *Cert, arg2 *PrivateKey) bool {

    v := C.CkCrypt2_SetSigningCert2(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Sets the timestamp authority (TSA) options for cases where a CAdES-T signature
// is to be created. The http is used to send the requests, and it allows for
// connection related settings and timeouts to be set. For example, if HTTP or
// SOCKS proxies are required, these features can be specified on the http.
func (z *Crypt2) SetTsaHttpObj(arg1 *Http)  {

    C.CkCrypt2_SetTsaHttpObj(z.ckObj, arg1.ckObj)



}


// Sets the digital certificate to be used in verifying a signature.
func (z *Crypt2) SetVerifyCert(arg1 *Cert) bool {

    v := C.CkCrypt2_SetVerifyCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Digitally signs the contents of dataToSign and returns the detached digital signature
// in an encoded string (according to the EncodingMode property setting).
// return a string or nil if failed.
func (z *Crypt2) SignBdENC(arg1 *BinData) *string {

    retStr := C.CkCrypt2_signBdENC(z.ckObj, arg1.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the SignBdENC method
func (z *Crypt2) SignBdENCAsync(arg1 *BinData, c chan *Task) {

    hTask := C.CkCrypt2_SignBdENCAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a byte array and returns the detached digital signature. A
// certificate must be set by calling SetSigningCert prior to calling this method.
func (z *Crypt2) SignBytes(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_SignBytes(z.ckObj, ckbyd1, ckOutBytes)

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


// Asynchronous version of the SignBytes method
func (z *Crypt2) SignBytesAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCrypt2_SignBytesAsync(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a byte array and returns the detached digital signature encoded
// as a printable string. A certificate must be set by calling SetSigningCert prior
// to calling this method. The EncodingMode property controls the output encoding,
// which can be "Base64", "QP", or "Hex".
// return a string or nil if failed.
func (z *Crypt2) SignBytesENC(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCrypt2_signBytesENC(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the SignBytesENC method
func (z *Crypt2) SignBytesENCAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkCrypt2_SignBytesENCAsync(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a the contents of sb and returns the PKCS7 detached digital
// signature as an encoded string according to the EncodingMode property setting.
// return a string or nil if failed.
func (z *Crypt2) SignSbENC(arg1 *StringBuilder) *string {

    retStr := C.CkCrypt2_signSbENC(z.ckObj, arg1.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the SignSbENC method
func (z *Crypt2) SignSbENCAsync(arg1 *StringBuilder, c chan *Task) {

    hTask := C.CkCrypt2_SignSbENCAsync(z.ckObj, arg1.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a string and returns the detached digital signature. A
// certificate must be set by calling SetSigningCert prior to calling this method.
// The Charset property controls the character encoding of the string that is
// signed. (Languages such as VB.NET, C#, and Visual Basic work with Unicode
// strings.) To sign Unicode data (2 bytes per char), set the Charset property to
// "Unicode". To implicitly convert the string to a mutlibyte charset such as
// "iso-8859-1", "Shift_JIS", "utf-8", or something else, then set the Charset
// property to the name of the charset before signing. The complete list of
// charsets is listed in the EncryptString method description.
func (z *Crypt2) SignString(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_SignString(z.ckObj, cstr1, ckOutBytes)

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


// Asynchronous version of the SignString method
func (z *Crypt2) SignStringAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCrypt2_SignStringAsync(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Digitally signs a string and returns the PKCS7 detached digital signature as an
// encoded string. A certificate must be set by calling SetSigningCert prior to
// calling this method. The Charset property controls the character encoding of the
// string that is signed. (Languages such as VB.NET, C#, and Visual Basic work with
// Unicode strings.) To sign Unicode data (2 bytes per char), set the Charset
// property to "Unicode". To implicitly convert the string to a mutlibyte charset
// such as "iso-8859-1", "Shift_JIS", "utf-8", or something else, then set the
// Charset property to the name of the charset before signing. The complete list of
// charsets is listed in the EncryptString method description.
// 
// The encoding of the output string is controlled by the EncodingMode property,
// which can be set to "Base64", "QP", or "Hex".
//
// return a string or nil if failed.
func (z *Crypt2) SignStringENC(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCrypt2_signStringENC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Asynchronous version of the SignStringENC method
func (z *Crypt2) SignStringENCAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkCrypt2_SignStringENCAsync(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Convert a string to a byte array where the characters are encoded according to
// the charset specified.
func (z *Crypt2) StringToBytes(arg1 string, arg2 string) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCrypt2_StringToBytes(z.ckObj, cstr1, cstr2, ckOutBytes)

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


// Implements RFC 6238: TOTP: Time-Based One-Time Password Algorithm. The arguments
// to this method are:
//     secret: The shared secret in an enocded representation such as base64, hex,
//     ascii, etc.
//     secretEnc: The encoding of the shared secret, such as "base64"
//     t0: The Unix time to start counting time steps. It is a number in decimal
//     string form. A Unix time is the number of seconds elapsed since midnight UTC of
//     January 1, 1970. "0" is a typical value used for this argument.
//     tNow: The current Unix time in decimal string form. To use the current
//     system date/time, pass an empty string for this argument.
//     tStep: The time step in seconds. A typical value is 30. Note: Both client and
//     server must pre-agree on the secret, the t0, and the tStep.
//     numDigits: The number of decimal digits to return.
//     truncOffset: Normally set this to -1 for dynamic truncation. Otherwise can be set
//     in the range 0..15.
//     hashAlg: Normally set to "sha1". Can be set to other hash algorithms such as
//     "sha256", "sha512", etc.
// return a string or nil if failed.
func (z *Crypt2) Totp(arg1 string, arg2 string, arg3 string, arg4 string, arg5 int, arg6 int, arg7 int, arg8 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr8 := C.CString(arg8)

    retStr := C.CkCrypt2_totp(z.ckObj, cstr1, cstr2, cstr3, cstr4, C.int(arg5), C.int(arg6), C.int(arg7), cstr8)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr8))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Trim a string ending with a specific substring until the string no longer ends
// with that substring.
// return a string or nil if failed.
func (z *Crypt2) TrimEndingWith(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkCrypt2_trimEndingWith(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Unlocks the component. This must be called once prior to calling any other
// method.
func (z *Crypt2) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCrypt2_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds an XML certificate vault to the object's internal list of sources to be
// searched for certificates and private keys when encrypting/decrypting or
// signing/verifying. Unlike the AddPfxSourceData and AddPfxSourceFile methods,
// only a single XML certificate vault can be used. If UseCertVault is called
// multiple times, only the last certificate vault will be used, as each call to
// UseCertVault will replace the certificate vault provided in previous calls.
func (z *Crypt2) UseCertVault(arg1 *XmlCertVault) bool {

    v := C.CkCrypt2_UseCertVault(z.ckObj, arg1.ckObj)


    return v != 0
}


// Verifies a digital signature against the original data contained in data.
// Returns true if the signature is verified.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) VerifyBdENC(arg1 *BinData, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_VerifyBdENC(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Verifies a byte array against a digital signature and returns true if the byte
// array is unaltered.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) VerifyBytes(arg1 []byte, arg2 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkCrypt2_VerifyBytes(z.ckObj, ckbyd1, ckbyd2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Verifies a byte array against a string-encoded digital signature and returns
// true if the byte array is unaltered. This method can be used to verify a
// signature produced by SignBytesENC. The EncodingMode property must be set prior
// to calling to match the encoding of the digital signature string ("Base64",
// "QP", or "Hex").
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) VerifyBytesENC(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_VerifyBytesENC(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Verifies a .p7s (PKCS #7 Signature) against the original file (or exact copy of
// it). If the inFilename has not been modified, the return value is true, otherwise it
// is false. This method is equivalent to VerifyP7S.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) VerifyDetachedSignature(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_VerifyDetachedSignature(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Verifies a .p7m file and extracts the original file from the .p7m. Returns
// true if the signature is valid and the contents are unchanged. Otherwise
// returns false.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) VerifyP7M(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_VerifyP7M(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Verifies a .p7s (PKCS #7 Signature) against the original file (or exact copy of
// it). If the inFilename has not been modified, the return value is true, otherwise it
// is false.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) VerifyP7S(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_VerifyP7S(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Verifies a digital signature against the original data contained in sb.
// Returns true if the signature is verified.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) VerifySbENC(arg1 *StringBuilder, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_VerifySbENC(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Verifies a string against a binary digital signature and returns true if the
// string is unaltered. This method can be used to verify a signature produced by
// SignString. The Charset property must be set to the charset that was used when
// creating the signature.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) VerifyString(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkCrypt2_VerifyString(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Verifies a string against a string-encoded digital signature and returns true if
// the string is unaltered. This method can be used to verify a signature produced
// by SignStringENC. The Charset and EncodingMode properties must be set to the
// same values that were used when creating the signature.
// 
// Note: The signer certificates can be retrieved after any Verify* method call by
// using the NumSignerCerts property and the GetSignerCert method.
//
func (z *Crypt2) VerifyStringENC(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCrypt2_VerifyStringENC(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Convenience method to write an entire byte array to a file.
func (z *Crypt2) WriteFile(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkCrypt2_WriteFile(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


