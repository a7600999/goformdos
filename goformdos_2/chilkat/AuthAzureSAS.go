// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkAuthAzureSAS.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewAuthAzureSAS() *AuthAzureSAS {
	return &AuthAzureSAS{C.CkAuthAzureSAS_Create()}
}

func (z *AuthAzureSAS) DisposeAuthAzureSAS() {
    if z != nil {
	C.CkAuthAzureSAS_Dispose(z.ckObj)
	}
}




// property: AccessKey
// This is the signing key (access key) that must be kept private. It is a base64
// string such as
// "abdTvCZFFoWUyre6erlNN+IOb9qhXgDsyhrxmZvpmxqFDwpl9oD0X9Fy0hIQa6L5UohznRLmkCtUYySO
// 4Y2eaw=="
func (z *AuthAzureSAS) AccessKey() string {
    return C.GoString(C.CkAuthAzureSAS_accessKey(z.ckObj))
}
// property setter: AccessKey
func (z *AuthAzureSAS) SetAccessKey(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureSAS_putAccessKey(z.ckObj,cStr)
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
func (z *AuthAzureSAS) DebugLogFilePath() string {
    return C.GoString(C.CkAuthAzureSAS_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *AuthAzureSAS) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureSAS_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthAzureSAS) LastErrorHtml() string {
    return C.GoString(C.CkAuthAzureSAS_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *AuthAzureSAS) LastErrorText() string {
    return C.GoString(C.CkAuthAzureSAS_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthAzureSAS) LastErrorXml() string {
    return C.GoString(C.CkAuthAzureSAS_lastErrorXml(z.ckObj))
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
func (z *AuthAzureSAS) LastMethodSuccess() bool {
    v := int(C.CkAuthAzureSAS_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *AuthAzureSAS) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthAzureSAS_putLastMethodSuccess(z.ckObj,v)
}

// property: StringToSign
// Defines the format of the string to sign.
// 
// The format is specified as a comma-separated list of names. For example:
// 
// signedpermissions,signedstart,signedexpiry,canonicalizedresource,signedidentifier,signedIP,signedProtocol,signedversion,rscc,rscd,rsce,rscl,rsct
// This will result in an actual string-to-sign that is composed of the values for
// each name separated by newline (LF) chars. For example:
// signedpermissions + "\n" +  
// signedstart + "\n" +  
// signedexpiry + "\n" +  
// canonicalizedresource + "\n" +  
// signedidentifier + "\n" +  
// signedIP + "\n" +  
// signedProtocol + "\n" +  
// signedversion + "\n" +  
// rscc + "\n" +  
// rscd + "\n" +  
// rsce + "\n" +  
// rscl + "\n" +  
// rsct
func (z *AuthAzureSAS) StringToSign() string {
    return C.GoString(C.CkAuthAzureSAS_stringToSign(z.ckObj))
}
// property setter: StringToSign
func (z *AuthAzureSAS) SetStringToSign(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAzureSAS_putStringToSign(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *AuthAzureSAS) VerboseLogging() bool {
    v := int(C.CkAuthAzureSAS_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *AuthAzureSAS) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthAzureSAS_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *AuthAzureSAS) Version() string {
    return C.GoString(C.CkAuthAzureSAS_version(z.ckObj))
}
// Clears all params set by the methods SetNonTokenParam and SetTokenParam.
func (z *AuthAzureSAS) Clear()  {

    C.CkAuthAzureSAS_Clear(z.ckObj)



}


// Generates and returns the SAS token based on the defined StringToSign and
// params.
// return a string or nil if failed.
func (z *AuthAzureSAS) GenerateToken() *string {

    retStr := C.CkAuthAzureSAS_generateToken(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Adds a non-token parameter name/value. This is a value that is included in the
// string to sign, but is NOT included as a parameter in the Authorization header.
func (z *AuthAzureSAS) SetNonTokenParam(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAuthAzureSAS_SetNonTokenParam(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a token parameter name/value. This is a value that is included in the
// string to sign, and is also included as a parameter in the Authorization header.
func (z *AuthAzureSAS) SetTokenParam(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkAuthAzureSAS_SetTokenParam(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


