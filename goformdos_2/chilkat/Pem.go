// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkPem.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewPem() *Pem {
	return &Pem{C.CkPem_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Pem) DisposePem() {
    if z != nil {
	C.CkPem_Dispose(z.ckObj)
	}
}




func (z *Pem) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkPem_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Pem) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkPem_setExternalProgress(z.ckObj,1)
}

func (z *Pem) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkPem_setExternalProgress(z.ckObj,1)
}

func (z *Pem) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkPem_setExternalProgress(z.ckObj,1)
}




// property: AppendMode
// When set to true, each of the Load* methods appends to the current contents of
// this PEM object. When set to false, a Load* method replaces the contents of
// this PEM object. The default is false.
func (z *Pem) AppendMode() bool {
    v := int(C.CkPem_getAppendMode(z.ckObj))
    return v != 0
}
// property setter: AppendMode
func (z *Pem) SetAppendMode(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPem_putAppendMode(z.ckObj,v)
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
func (z *Pem) DebugLogFilePath() string {
    return C.GoString(C.CkPem_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Pem) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkPem_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HeartbeatMs
// This is the number of milliseconds between each AbortCheck event callback. The
// AbortCheck callback allows an application to abort any event-enabled methods
// prior to completion. If HeartbeatMs is 0 (the default), no AbortCheck event
// callbacks will fire.
func (z *Pem) HeartbeatMs() int {
    return int(C.CkPem_getHeartbeatMs(z.ckObj))
}
// property setter: HeartbeatMs
func (z *Pem) SetHeartbeatMs(value int) {
    C.CkPem_putHeartbeatMs(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Pem) LastErrorHtml() string {
    return C.GoString(C.CkPem_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Pem) LastErrorText() string {
    return C.GoString(C.CkPem_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Pem) LastErrorXml() string {
    return C.GoString(C.CkPem_lastErrorXml(z.ckObj))
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
func (z *Pem) LastMethodSuccess() bool {
    v := int(C.CkPem_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Pem) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPem_putLastMethodSuccess(z.ckObj,v)
}

// property: NumCerts
// The number of certificates in the loaded PEM. To get the 1st certificate, call
// GetCert(0).
func (z *Pem) NumCerts() int {
    return int(C.CkPem_getNumCerts(z.ckObj))
}

// property: NumCrls
// The number of certificate revocation lists (CRLs) in the loaded PEM.
func (z *Pem) NumCrls() int {
    return int(C.CkPem_getNumCrls(z.ckObj))
}

// property: NumCsrs
// The number of certificate signing requests (CSRs) in the loaded PEM.
func (z *Pem) NumCsrs() int {
    return int(C.CkPem_getNumCsrs(z.ckObj))
}

// property: NumPrivateKeys
// The number of private keys in the loaded PEM. To get the 1st private key, call
// GetPrivateKey(0).
func (z *Pem) NumPrivateKeys() int {
    return int(C.CkPem_getNumPrivateKeys(z.ckObj))
}

// property: NumPublicKeys
// The number of public keys in the loaded PEM. To get the 1st public key, call
// GetPublicKey(0).
func (z *Pem) NumPublicKeys() int {
    return int(C.CkPem_getNumPublicKeys(z.ckObj))
}

// property: PrivateKeyFormat
// Controls the format to be used for unencrypted private keys when writing a PEM.
// Possible values are "pkcs1" and "pkcs8". (OpenSSL typically uses the "pkcs8"
// format.) When writing encrypted private keys to PEM, the format is always PKCS8,
// and the PEM header is "BEGIN ENCRYPTED PRIVATE KEY". The default is "pkcs8".
// 
// The PKCS1 format uses the PEM header: BEGIN RSA PRIVATE KEY.
// The PKCS8 format uses the PEM header: BEGIN PRIVATE KEY.
//
func (z *Pem) PrivateKeyFormat() string {
    return C.GoString(C.CkPem_privateKeyFormat(z.ckObj))
}
// property setter: PrivateKeyFormat
func (z *Pem) SetPrivateKeyFormat(goStr string) {
    cStr := C.CString(goStr)
    C.CkPem_putPrivateKeyFormat(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: PublicKeyFormat
// Controls the format to be used for public keys when writing a PEM. Possible
// values are "pkcs1" and "pkcs8". (OpenSSL typically uses the "pkcs8" format.) The
// default is "pkcs8".
// 
// The PKCS1 format uses the PEM header: BEGIN RSA PUBLIC KEY.
// The PKCS8 format uses the PEM header: BEGIN PUBLIC KEY.
//
func (z *Pem) PublicKeyFormat() string {
    return C.GoString(C.CkPem_publicKeyFormat(z.ckObj))
}
// property setter: PublicKeyFormat
func (z *Pem) SetPublicKeyFormat(goStr string) {
    cStr := C.CString(goStr)
    C.CkPem_putPublicKeyFormat(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Pem) VerboseLogging() bool {
    v := int(C.CkPem_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Pem) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkPem_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Pem) Version() string {
    return C.GoString(C.CkPem_version(z.ckObj))
}
// Adds a certificate, and potentially the certs in its chain of authentication to
// the PEM. If includeChain is true, then certificates in the cert's chain of
// authentication up to and including the root are automatically added.
func (z *Pem) AddCert(arg1 *Cert, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkPem_AddCert(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Adds a certificate, private key, public key, or csr to the PEM. The possible
// values for itemType are "certificate" (or "cert"), "privateKey", "publicKey", or
// "csr". The encoding can be "Base64", "modBase64", "Base32", "QP" (for
// quoted-printable), "URL" (for url-encoding), "Hex", "url_oauth", "url_rfc1738",
// "url_rfc2396", and "url_rfc3986". The itemData contains the ASN.1 data in string
// format according to the encoding specified in encoding.
func (z *Pem) AddItem(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkPem_AddItem(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds a private key to the PEM object.
func (z *Pem) AddPrivateKey(arg1 *PrivateKey) bool {

    v := C.CkPem_AddPrivateKey(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds a private key and it's associated certificate chain to the PEM object.
func (z *Pem) AddPrivateKey2(arg1 *PrivateKey, arg2 *CertChain) bool {

    v := C.CkPem_AddPrivateKey2(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Adds a public key to the PEM object.
func (z *Pem) AddPublicKey(arg1 *PublicKey) bool {

    v := C.CkPem_AddPublicKey(z.ckObj, arg1.ckObj)


    return v != 0
}


// Removes all content from this PEM object.
func (z *Pem) Clear() bool {

    v := C.CkPem_Clear(z.ckObj)


    return v != 0
}


// Returns the Nth certificate from the PEM. The first certificate is at index 0.
func (z *Pem) GetCert(arg1 int) *Cert {

    retObj := C.CkPem_GetCert(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the encoded contents of the Nth item of a particular type (0-based
// index). The possible values for itemType are "certificate" (or "cert"), "privateKey",
// "publicKey", or "csr". Input string args are case-insensitive. If the itemType is
// "privateKey", the itemSubType may be "der" or "pkcs8". If the itemType is "publicKey", the
// itemSubType may be "der" or "pkcs1". The itemSubType is ignored for other values of itemType. The
// valid encoding modes are "Base64", "modBase64", "Base32", "Base58", "QP" (for
// quoted-printable), "URL" (for url-encoding), "Hex", "url_oauth", "url_rfc1738",
// "url_rfc2396", and "url_rfc3986".
// return a string or nil if failed.
func (z *Pem) GetEncodedItem(arg1 string, arg2 string, arg3 string, arg4 int) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkPem_getEncodedItem(z.ckObj, cstr1, cstr2, cstr3, C.int(arg4))

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth private key from the PEM. The first private key is at index 0.
func (z *Pem) GetPrivateKey(arg1 int) *PrivateKey {

    retObj := C.CkPem_GetPrivateKey(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &PrivateKey{retObj}
}


// Returns the Nth public key from the PEM. The first public key is at index 0.
func (z *Pem) GetPublicKey(arg1 int) *PublicKey {

    retObj := C.CkPem_GetPublicKey(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &PublicKey{retObj}
}


// Loads the PEM from the contents of an in-memory PKCS7 container (.p7b).
func (z *Pem) LoadP7b(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkPem_LoadP7b(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Asynchronous version of the LoadP7b method
func (z *Pem) LoadP7bAsync(arg1 []byte, c chan *Task) {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    hTask := C.CkPem_LoadP7bAsync(z.ckObj, ckbyd1)

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


// Loads the contents of a PKCS7 container (.p7b file).
func (z *Pem) LoadP7bFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkPem_LoadP7bFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the LoadP7bFile method
func (z *Pem) LoadP7bFileAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkPem_LoadP7bFileAsync(z.ckObj, cstr1)

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


// Loads the PEM from a PEM string. If encrypted, then the password is required for
// decryption. Otherwise, an empty string (or any string) may be passed for the
// password.
func (z *Pem) LoadPem(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPem_LoadPem(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the LoadPem method
func (z *Pem) LoadPemAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkPem_LoadPemAsync(z.ckObj, cstr1, cstr2)

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


// Loads the PEM from a PEM file. If encrypted, then the password is required for
// decryption. Otherwise, an empty string (or any string) may be passed for the
// password.
func (z *Pem) LoadPemFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkPem_LoadPemFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Asynchronous version of the LoadPemFile method
func (z *Pem) LoadPemFileAsync(arg1 string, arg2 string, c chan *Task) {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    hTask := C.CkPem_LoadPemFileAsync(z.ckObj, cstr1, cstr2)

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


// Loads the caller of the task's async method.
func (z *Pem) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkPem_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Removes the Nth certificate from the PEM. The first certificate is at index 0.
func (z *Pem) RemoveCert(arg1 int) bool {

    v := C.CkPem_RemoveCert(z.ckObj, C.int(arg1))


    return v != 0
}


// Removes the Nth private key from the PEM. The first private key is at index 0.
func (z *Pem) RemovePrivateKey(arg1 int) bool {

    v := C.CkPem_RemovePrivateKey(z.ckObj, C.int(arg1))


    return v != 0
}


// Converts the PEM to JKS and returns the Java KeyStore object. If the alias is
// non-empty, the 1st object (private key or certificate) will use the alias, and
// all others (if any) will receive auto-generated aliases. The JKS returned will
// be encrypted using the provided password. If the PEM contains only certificates (no
// private keys), then the password is unused.
func (z *Pem) ToJks(arg1 string, arg2 string) *JavaKeyStore {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkPem_ToJks(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &JavaKeyStore{retObj}
}


// Write the PFX to a PEM formatted string. The resultant PEM will contain the
// private key, as well as the certs in the chain of authentication (or whatever
// certs are available in the PFX). For example:
//  -----BEGIN RSA PRIVATE KEY-----
// ...
// ... the private key associated with the main certificate.
// ...
// -----END RSA PRIVATE KEY-----
// -----BEGIN CERTIFICATE-----
// ...
// ... the main certificate
// ...
// -----END CERTIFICATE-----
// -----BEGIN CERTIFICATE-----
// ...
// ... an intermediate CA certificate (if present)
// ...
// -----END CERTIFICATE-----
// -----BEGIN CERTIFICATE-----
// ...
// ... the root CA certificate
// ...
// -----END CERTIFICATE-----
// return a string or nil if failed.
func (z *Pem) ToPem() *string {

    retStr := C.CkPem_toPem(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Write the PFX to a PEM formatted string. If extendedAttrs is true, then extended
// properties (Bag Attributes and Key Attributes) are output. If noKeys is true,
// then no private keys are output. If noCerts is true, then no certificates are
// output. If noCaCerts is true, then no CA certs or intermediate CA certs are output.
// If encryptAlg is not empty, it indicates the encryption algorithm to be used for
// encrypting the private keys (otherwise the private keys are output unencrypted).
// The possible choices for the encryptAlg are "des3", "aes128", "aes192", and "aes256".
// (All encryption algorithm choices use CBC mode.) If the private keys are to be
// encrypted, then password is the password to be used. Otherwise, password may be left
// empty. For example:
// Bag Attributes
//     Microsoft Local Key set: localKeyID: 01 00 00 00 
//     friendlyName: le-2b09a3d2-9037-4a05-95cc-4d44518e8607
//     Microsoft CSP Name: Microsoft RSA SChannel Cryptographic Provider
// Key Attributes
//     X509v3 Key Usage: 10 
//  -----BEGIN RSA PRIVATE KEY-----
// ...
// ... the private key associated with the main certificate.
// ...
// -----END RSA PRIVATE KEY-----
// Bag Attributes
//     localKeyID: 01 00 00 00 
//     1.3.6.1.4.1.311.17.3.92: 00 08 00 00 
//     1.3.6.1.4.1.311.17.3.20: C2 53 54 F3 ...
//     1.3.6.1.4.1.311.17.3.71: 49 00 43 00 ...
//     1.3.6.1.4.1.311.17.3.75: 31 00 42 00 ...
// subject=/OU=Domain Control Validated/OU=PositiveSSL/CN=something.com
// issuer=/C=GB/ST=Greater Manchester/L=Salford/O=COMODO CA Limited/CN=COMODO RSA Domain Validation Secure Server CA
// -----BEGIN CERTIFICATE-----
// ...
// ... the main certificate
// ...
// -----END CERTIFICATE-----
// ...
// -----BEGIN CERTIFICATE-----
// ...
// ... an intermediate CA certificate (if present)
// ...
// -----END CERTIFICATE-----
// ...
// -----BEGIN CERTIFICATE-----
// ...
// ... the root CA certificate
// ...
// -----END CERTIFICATE-----
// return a string or nil if failed.
func (z *Pem) ToPemEx(arg1 bool, arg2 bool, arg3 bool, arg4 bool, arg5 string, arg6 string) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
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
    cstr5 := C.CString(arg5)
    cstr6 := C.CString(arg6)

    retStr := C.CkPem_toPemEx(z.ckObj, b1, b2, b3, b4, cstr5, cstr6)

    C.free(unsafe.Pointer(cstr5))
    C.free(unsafe.Pointer(cstr6))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Converts the PEM to PKCS12 and returns the PFX object. The PFX object has method
// for saving to a file, exporting to an encoded string, converting to a JKS (Java
// Keystore), or even converting back to PEM.
// 
// Note: The PEM must contain at least one private key to convert to PKCS12. The
// typical case is that a PKCS12 contains a single private key, along with the
// associated certificate and the certificates in the chain of authentication.
//
func (z *Pem) ToPfx() *Pfx {

    retObj := C.CkPem_ToPfx(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Pfx{retObj}
}


