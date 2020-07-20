// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkAuthAws.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewAuthAws() *AuthAws {
	return &AuthAws{C.CkAuthAws_Create()}
}

func (z *AuthAws) DisposeAuthAws() {
    if z != nil {
	C.CkAuthAws_Dispose(z.ckObj)
	}
}




// property: AccessKey
// The AWS access key.
func (z *AuthAws) AccessKey() string {
    return C.GoString(C.CkAuthAws_accessKey(z.ckObj))
}
// property setter: AccessKey
func (z *AuthAws) SetAccessKey(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAws_putAccessKey(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: CanonicalizedResourceV2
// If AWS Signature Version V2 is used, then this property must be set. The rules
// for setting the canonicalized resource for the V2 signature method is described
// here:Constructing the CanonicalizedResource Element
// <http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html#Construct
// ingTheCanonicalizedResourceElement>.
func (z *AuthAws) CanonicalizedResourceV2() string {
    return C.GoString(C.CkAuthAws_canonicalizedResourceV2(z.ckObj))
}
// property setter: CanonicalizedResourceV2
func (z *AuthAws) SetCanonicalizedResourceV2(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAws_putCanonicalizedResourceV2(z.ckObj,cStr)
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
func (z *AuthAws) DebugLogFilePath() string {
    return C.GoString(C.CkAuthAws_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *AuthAws) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAws_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthAws) LastErrorHtml() string {
    return C.GoString(C.CkAuthAws_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *AuthAws) LastErrorText() string {
    return C.GoString(C.CkAuthAws_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *AuthAws) LastErrorXml() string {
    return C.GoString(C.CkAuthAws_lastErrorXml(z.ckObj))
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
func (z *AuthAws) LastMethodSuccess() bool {
    v := int(C.CkAuthAws_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *AuthAws) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthAws_putLastMethodSuccess(z.ckObj,v)
}

// property: PrecomputedMd5
// This property can optionally be set for AWS requests that have a non-empty
// request body. This should be the base64 encoding of the 16 bytes of the MD5
// hash. The most common need for this is if doing an S3 upload from a stream. (If
// the pre-computed MD5 is not provided, then Chilkat is forced to stream the
// entire file into memory so that it can calculate the MD5 for authentication.)
// 
// Note: AWS Signature Version 2 uses the MD5, whereas Signature Version 4 uses
// SHA256.
//
func (z *AuthAws) PrecomputedMd5() string {
    return C.GoString(C.CkAuthAws_precomputedMd5(z.ckObj))
}
// property setter: PrecomputedMd5
func (z *AuthAws) SetPrecomputedMd5(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAws_putPrecomputedMd5(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: PrecomputedSha256
// This property can optionally be set for AWS requests that have a non-empty
// request body. This should be the lowercase hex encoding of the 32-bytes of the
// SHA256 hash. The most common need for this is if doing an S3 upload from a
// stream. (If the pre-computed SHA-256 is not provided, then Chilkat is forced to
// stream the entire file into memory so that it can calculate the SHA-256 for
// authentication.)
// 
// Note: AWS Signature Version 4 uses the SHA256 hash. (AWS Signature Version 2
// uses MD5)
//
func (z *AuthAws) PrecomputedSha256() string {
    return C.GoString(C.CkAuthAws_precomputedSha256(z.ckObj))
}
// property setter: PrecomputedSha256
func (z *AuthAws) SetPrecomputedSha256(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAws_putPrecomputedSha256(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Region
// The AWS region, such as "us-east-1", "us-west-2", "eu-west-1", "eu-central-1",
// etc. The default is "us-east-1". It is only used when the SignatureVersion
// property is set to 4. This property is unused when the SignatureVersion property
// is set to 2.
func (z *AuthAws) Region() string {
    return C.GoString(C.CkAuthAws_region(z.ckObj))
}
// property setter: Region
func (z *AuthAws) SetRegion(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAws_putRegion(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SecretKey
// The AWS secret key.
func (z *AuthAws) SecretKey() string {
    return C.GoString(C.CkAuthAws_secretKey(z.ckObj))
}
// property setter: SecretKey
func (z *AuthAws) SetSecretKey(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAws_putSecretKey(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ServiceName
// The AWS service namespace, such as "s3", "ses", etc. See
// http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-
// aws-service-namespaces
// 
// This property is unused when the SignatureVersion property is set to 2.
//
func (z *AuthAws) ServiceName() string {
    return C.GoString(C.CkAuthAws_serviceName(z.ckObj))
}
// property setter: ServiceName
func (z *AuthAws) SetServiceName(goStr string) {
    cStr := C.CString(goStr)
    C.CkAuthAws_putServiceName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SignatureVersion
// The AWS Signature Version to be used in authentication. The default value is 4.
// This can optionally be set to the value 2 to use the older V2 signature version.
func (z *AuthAws) SignatureVersion() int {
    return int(C.CkAuthAws_getSignatureVersion(z.ckObj))
}
// property setter: SignatureVersion
func (z *AuthAws) SetSignatureVersion(value int) {
    C.CkAuthAws_putSignatureVersion(z.ckObj,C.int(value))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *AuthAws) VerboseLogging() bool {
    v := int(C.CkAuthAws_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *AuthAws) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAuthAws_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *AuthAws) Version() string {
    return C.GoString(C.CkAuthAws_version(z.ckObj))
}
