// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkCertStore.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCertStore() *CertStore {
	return &CertStore{C.CkCertStore_Create()}
}

func (z *CertStore) DisposeCertStore() {
    if z != nil {
	C.CkCertStore_Dispose(z.ckObj)
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
func (z *CertStore) DebugLogFilePath() string {
    return C.GoString(C.CkCertStore_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *CertStore) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkCertStore_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *CertStore) LastErrorHtml() string {
    return C.GoString(C.CkCertStore_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *CertStore) LastErrorText() string {
    return C.GoString(C.CkCertStore_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *CertStore) LastErrorXml() string {
    return C.GoString(C.CkCertStore_lastErrorXml(z.ckObj))
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
func (z *CertStore) LastMethodSuccess() bool {
    v := int(C.CkCertStore_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *CertStore) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCertStore_putLastMethodSuccess(z.ckObj,v)
}

// property: NumCertificates
// The number of certificates held in the certificate store.
func (z *CertStore) NumCertificates() int {
    return int(C.CkCertStore_getNumCertificates(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *CertStore) VerboseLogging() bool {
    v := int(C.CkCertStore_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *CertStore) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCertStore_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *CertStore) Version() string {
    return C.GoString(C.CkCertStore_version(z.ckObj))
}
// Locates and returns a certificate by its RFC 822 name.
// 
// If multiple certificates match, then non-expired certificates will take
// precedence over expired certificates. In other words, Chilkat will aways return
// the non-expired certificate over the expired certificate.
//
func (z *CertStore) FindCertByRfc822Name(arg1 string) *Cert {
    cstr1 := C.CString(arg1)

    retObj := C.CkCertStore_FindCertByRfc822Name(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Finds and returns the certificate that has the matching serial number.
// 
// Returns _NULL_ on failure.
//
func (z *CertStore) FindCertBySerial(arg1 string) *Cert {
    cstr1 := C.CString(arg1)

    retObj := C.CkCertStore_FindCertBySerial(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Finds a certificate by it's SHA-1 thumbprint. The thumbprint is a hexidecimal
// string.
// 
// Returns _NULL_ on failure.
//
func (z *CertStore) FindCertBySha1Thumbprint(arg1 string) *Cert {
    cstr1 := C.CString(arg1)

    retObj := C.CkCertStore_FindCertBySha1Thumbprint(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Finds a certificate where one of the Subject properties (SubjectCN, SubjectE,
// SubjectO, SubjectOU, SubjectL, SubjectST, SubjectC) matches exactly (but case
// insensitive) with the passed string. A match in SubjectCN will be tried first,
// followed by SubjectE, and SubjectO. After that, the first match found in
// SubjectOU, SubjectL, SubjectST, or SubjectC, but in no guaranteed order, is
// returned. All matches are case insensitive.
// 
// Returns _NULL_ on failure.
//
func (z *CertStore) FindCertBySubject(arg1 string) *Cert {
    cstr1 := C.CString(arg1)

    retObj := C.CkCertStore_FindCertBySubject(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Finds a certificate where the SubjectCN property (common name) matches exactly
// (but case insensitive) with the passed string
// 
// Returns _NULL_ on failure.
//
func (z *CertStore) FindCertBySubjectCN(arg1 string) *Cert {
    cstr1 := C.CString(arg1)

    retObj := C.CkCertStore_FindCertBySubjectCN(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Finds a certificate where the SubjectE property (email address) matches exactly
// (but case insensitive) with the passed string. This function differs from
// FindCertForEmail in that the certificate does not need to match the
// ForSecureEmail property.
// 
// Returns _NULL_ on failure.
//
func (z *CertStore) FindCertBySubjectE(arg1 string) *Cert {
    cstr1 := C.CString(arg1)

    retObj := C.CkCertStore_FindCertBySubjectE(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Finds a certificate where the SubjectO property (organization) matches exactly
// (but case insensitive) with the passed string.
// 
// Returns _NULL_ on failure.
//
func (z *CertStore) FindCertBySubjectO(arg1 string) *Cert {
    cstr1 := C.CString(arg1)

    retObj := C.CkCertStore_FindCertBySubjectO(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the Nth certificate in the store. The first certificate is at index 0.
// 
// Returns _NULL_ on failure.
//
func (z *CertStore) GetCertificate(arg1 int) *Cert {

    retObj := C.CkCertStore_GetCertificate(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Loads the certificates contained within a PEM formatted file.
func (z *CertStore) LoadPemFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCertStore_LoadPemFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the certificates contained within an in-memory PEM formatted string.
func (z *CertStore) LoadPemStr(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCertStore_LoadPemStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a PFX from an in-memory image of a PFX file. Once loaded, the certificates
// within the PFX may be searched via the Find* methods. It is also possible to
// iterate from 0 to NumCertficates-1, calling GetCertificate for each index, to
// retrieve each certificate within the PFX.
func (z *CertStore) LoadPfxData(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkCertStore_LoadPfxData(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads a PFX file. Once loaded, the certificates within the PFX may be searched
// via the Find* methods. It is also possible to iterate from 0 to
// NumCertficates-1, calling GetCertificate for each index, to retrieve each
// certificate within the PFX.
func (z *CertStore) LoadPfxFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCertStore_LoadPfxFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


