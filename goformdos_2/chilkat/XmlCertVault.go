// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkXmlCertVault.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewXmlCertVault() *XmlCertVault {
	return &XmlCertVault{C.CkXmlCertVault_Create()}
}

func (z *XmlCertVault) DisposeXmlCertVault() {
    if z != nil {
	C.CkXmlCertVault_Dispose(z.ckObj)
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
func (z *XmlCertVault) DebugLogFilePath() string {
    return C.GoString(C.CkXmlCertVault_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *XmlCertVault) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlCertVault_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *XmlCertVault) LastErrorHtml() string {
    return C.GoString(C.CkXmlCertVault_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *XmlCertVault) LastErrorText() string {
    return C.GoString(C.CkXmlCertVault_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *XmlCertVault) LastErrorXml() string {
    return C.GoString(C.CkXmlCertVault_lastErrorXml(z.ckObj))
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
func (z *XmlCertVault) LastMethodSuccess() bool {
    v := int(C.CkXmlCertVault_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *XmlCertVault) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmlCertVault_putLastMethodSuccess(z.ckObj,v)
}

// property: MasterPassword
// The master password for the vault. Certificates are stored unencrypted, but
// private keys are stored 256-bit AES encrypted using the individual PFX
// passwords. The PFX passwords are stored 256-bit AES encrypted using the master
// password. The master password is required to acces any of the private keys
// stored within the XML vault.
func (z *XmlCertVault) MasterPassword() string {
    return C.GoString(C.CkXmlCertVault_masterPassword(z.ckObj))
}
// property setter: MasterPassword
func (z *XmlCertVault) SetMasterPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmlCertVault_putMasterPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *XmlCertVault) VerboseLogging() bool {
    v := int(C.CkXmlCertVault_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *XmlCertVault) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmlCertVault_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *XmlCertVault) Version() string {
    return C.GoString(C.CkXmlCertVault_version(z.ckObj))
}
// Adds a certificate to the XML vault.
func (z *XmlCertVault) AddCert(arg1 *Cert) bool {

    v := C.CkXmlCertVault_AddCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds a certificate to the XML vault from any binary format, such as DER.
func (z *XmlCertVault) AddCertBinary(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkXmlCertVault_AddCertBinary(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Adds a chain of certificates to the XML vault.
func (z *XmlCertVault) AddCertChain(arg1 *CertChain) bool {

    v := C.CkXmlCertVault_AddCertChain(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds a certificate to the XML vault where certificate is passed directly from
// encoded bytes (such as Base64, Hex, etc.). The encoding is indicated by encoding.
func (z *XmlCertVault) AddCertEncoded(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXmlCertVault_AddCertEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a certificate to the XML vault.
func (z *XmlCertVault) AddCertFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXmlCertVault_AddCertFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds a certificate from any string representation format such as PEM.
func (z *XmlCertVault) AddCertString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXmlCertVault_AddCertString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds the contents of a PEM file to the XML vault. The PEM file may be encrypted,
// and it may contain one or more certificates and/or private keys. The password is
// optional and only required if the PEM file contains encrypted content that
// requires a password.
func (z *XmlCertVault) AddPemFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXmlCertVault_AddPemFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a PFX (PKCS12) to the XML vault.
func (z *XmlCertVault) AddPfx(arg1 *Pfx) bool {

    v := C.CkXmlCertVault_AddPfx(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds a PFX to the XML vault where PFX is passed directly from in-memory binary
// bytes.
func (z *XmlCertVault) AddPfxBinary(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkXmlCertVault_AddPfxBinary(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a PFX to the XML vault where PFX is passed directly from encoded bytes
// (such as Base64, Hex, etc.). The encoding is indicated by encoding.
func (z *XmlCertVault) AddPfxEncoded(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkXmlCertVault_AddPfxEncoded(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds a PFX file to the XML vault.
func (z *XmlCertVault) AddPfxFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXmlCertVault_AddPfxFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the contents of the cert vault as an XML string.
// return a string or nil if failed.
func (z *XmlCertVault) GetXml() *string {

    retStr := C.CkXmlCertVault_getXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads from an XML string.
func (z *XmlCertVault) LoadXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXmlCertVault_LoadXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads from an XML file.
func (z *XmlCertVault) LoadXmlFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXmlCertVault_LoadXmlFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the contents to an XML file.
func (z *XmlCertVault) SaveXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXmlCertVault_SaveXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


