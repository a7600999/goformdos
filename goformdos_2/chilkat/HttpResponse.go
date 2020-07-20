// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkHttpResponse.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewHttpResponse() *HttpResponse {
	return &HttpResponse{C.CkHttpResponse_Create()}
}

func (z *HttpResponse) DisposeHttpResponse() {
    if z != nil {
	C.CkHttpResponse_Dispose(z.ckObj)
	}
}




// property: Body
// The response body returned as a byte array.
func (z *HttpResponse) Body() []byte {
    ckbyd := C.CkByteData_Create()
    C.CkHttpResponse_getBody(z.ckObj,ckbyd)
    p := C.CkByteData_getData(ckbyd);
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckbyd)) )
    C.CkByteData_Dispose(ckbyd)
    return retBytes
}


// property: BodyQP
// The same as the Body property, but returned as a quoted-printable encoded
// string.
func (z *HttpResponse) BodyQP() string {
    return C.GoString(C.CkHttpResponse_bodyQP(z.ckObj))
}

// property: BodyStr
// The response body returned as a string.
func (z *HttpResponse) BodyStr() string {
    return C.GoString(C.CkHttpResponse_bodyStr(z.ckObj))
}

// property: Charset
// The response charset, such as "iso-8859-1", if applicable. Obviously, responses
// for GIF and JPG files will not have a charset.
func (z *HttpResponse) Charset() string {
    return C.GoString(C.CkHttpResponse_charset(z.ckObj))
}

// property: ContentLength
// The content length of the response, in bytes.
func (z *HttpResponse) ContentLength() uint {
    return uint(C.CkHttpResponse_getContentLength(z.ckObj))
}

// property: DateStr
// Returns the content of the Date response header field in RFC822 date/time string
// format.
func (z *HttpResponse) DateStr() string {
    return C.GoString(C.CkHttpResponse_dateStr(z.ckObj))
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
func (z *HttpResponse) DebugLogFilePath() string {
    return C.GoString(C.CkHttpResponse_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *HttpResponse) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttpResponse_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Domain
// The domain of the HTTP server that created this response.
func (z *HttpResponse) Domain() string {
    return C.GoString(C.CkHttpResponse_domain(z.ckObj))
}

// property: FinalRedirectUrl
// Returns the redirect URL for 301/302 responses.
func (z *HttpResponse) FinalRedirectUrl() string {
    return C.GoString(C.CkHttpResponse_finalRedirectUrl(z.ckObj))
}

// property: FullMime
// Returns the full MIME (header + body) of the HTTP response.
func (z *HttpResponse) FullMime() string {
    return C.GoString(C.CkHttpResponse_fullMime(z.ckObj))
}

// property: Header
// The full text of the response header.
func (z *HttpResponse) Header() string {
    return C.GoString(C.CkHttpResponse_header(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *HttpResponse) LastErrorHtml() string {
    return C.GoString(C.CkHttpResponse_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *HttpResponse) LastErrorText() string {
    return C.GoString(C.CkHttpResponse_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *HttpResponse) LastErrorXml() string {
    return C.GoString(C.CkHttpResponse_lastErrorXml(z.ckObj))
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
func (z *HttpResponse) LastMethodSuccess() bool {
    v := int(C.CkHttpResponse_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *HttpResponse) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttpResponse_putLastMethodSuccess(z.ckObj,v)
}

// property: NumCookies
// The number of cookies included in the response.
func (z *HttpResponse) NumCookies() int {
    return int(C.CkHttpResponse_getNumCookies(z.ckObj))
}

// property: NumHeaderFields
// The number of response header fields.
func (z *HttpResponse) NumHeaderFields() int {
    return int(C.CkHttpResponse_getNumHeaderFields(z.ckObj))
}

// property: StatusCode
// The status code (as an integer) from the first line of an HTTP response. If the
// StatusLine = "HTTP/1.0 200 OK", the response status code returned is 200.
func (z *HttpResponse) StatusCode() int {
    return int(C.CkHttpResponse_getStatusCode(z.ckObj))
}

// property: StatusLine
// The first line of an HTTP response, such as "HTTP/1.0 200 OK".
func (z *HttpResponse) StatusLine() string {
    return C.GoString(C.CkHttpResponse_statusLine(z.ckObj))
}

// property: StatusText
// The text that follows the status code in the 1st line of the HTTP response. For
// example, i the first line of an HTTP response is "HTTP/1.0 200 OK", then this
// property contains "OK".
func (z *HttpResponse) StatusText() string {
    return C.GoString(C.CkHttpResponse_statusText(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *HttpResponse) VerboseLogging() bool {
    v := int(C.CkHttpResponse_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *HttpResponse) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttpResponse_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *HttpResponse) Version() string {
    return C.GoString(C.CkHttpResponse_version(z.ckObj))
}
// Copies the response body to a BinData object.
func (z *HttpResponse) GetBodyBd(arg1 *BinData) bool {

    v := C.CkHttpResponse_GetBodyBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Copies the response body to a Chilkat StringBuilder object.
func (z *HttpResponse) GetBodySb(arg1 *StringBuilder) bool {

    v := C.CkHttpResponse_GetBodySb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the domain of the Nth cookie in the response. Indexing begins at 0. The
// number of response cookies is specified in the NumCookies property.
// return a string or nil if failed.
func (z *HttpResponse) GetCookieDomain(arg1 int) *string {

    retStr := C.CkHttpResponse_getCookieDomain(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the expiration date/time of the Nth cookie in the response. Indexing
// begins at 0. The number of response cookies is specified in the NumCookies
// property.
// return a string or nil if failed.
func (z *HttpResponse) GetCookieExpiresStr(arg1 int) *string {

    retStr := C.CkHttpResponse_getCookieExpiresStr(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the name of the Nth cookie returned in the response. Indexing begins at
// 0. The number of response cookies is specified in the NumCookies property.
// return a string or nil if failed.
func (z *HttpResponse) GetCookieName(arg1 int) *string {

    retStr := C.CkHttpResponse_getCookieName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the path of the Nth cookie returned in the response. Indexing begins at
// 0. The number of response cookies is specified in the NumCookies property.
// return a string or nil if failed.
func (z *HttpResponse) GetCookiePath(arg1 int) *string {

    retStr := C.CkHttpResponse_getCookiePath(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of the Nth cookie returned in the response. Indexing begins at
// 0. The number of response cookies is specified in the NumCookies property.
// return a string or nil if failed.
func (z *HttpResponse) GetCookieValue(arg1 int) *string {

    retStr := C.CkHttpResponse_getCookieValue(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of a response header field accessed by field name.
// return a string or nil if failed.
func (z *HttpResponse) GetHeaderField(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttpResponse_getHeaderField(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a response header field attribute. As an example, the response charset
// is simply the GetHeaderFieldAttr("content-type","charset")
// return a string or nil if failed.
func (z *HttpResponse) GetHeaderFieldAttr(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkHttpResponse_getHeaderFieldAttr(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets the name of the Nth response header field. Indexing begins at 0. The number
// of response headers is specified by the NumHeaderFields property.
// return a string or nil if failed.
func (z *HttpResponse) GetHeaderName(arg1 int) *string {

    retStr := C.CkHttpResponse_getHeaderName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets the value of the Nth response header field. Indexing begins at 0. The
// number of response headers is specified by the NumHeaderFields property.
// return a string or nil if failed.
func (z *HttpResponse) GetHeaderValue(arg1 int) *string {

    retStr := C.CkHttpResponse_getHeaderValue(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads the HTTP response from a completed asynchronous task. A failed return
// value indicates that no HTTP response was received in the HTTP asynchronous
// method call (i.e. the asynchronous HTTP request failed in such a way that no
// response was received).
func (z *HttpResponse) LoadTaskResult(arg1 *Task) bool {

    v := C.CkHttpResponse_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


// Saves the body of the HTTP response to a file.
func (z *HttpResponse) SaveBodyBinary(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttpResponse_SaveBodyBinary(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the HTTP response body to a file. This method provides control over CRLF
// vs bare-LF line-endings. If bCrlf is true, then line endings are automatically
// converted to CRLF if necessary. If bCrlf is false, then line-endings are
// automatically converted to bare-LF's (Unix style) if necessary.
// 
// To save the HTTP response body exactly as-is (with no line-ending manipulation),
// then call SaveBodyBinary.
//
func (z *HttpResponse) SaveBodyText(arg1 bool, arg2 string) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr2 := C.CString(arg2)

    v := C.CkHttpResponse_SaveBodyText(z.ckObj, b1, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Convenience method for parsing a param's value out of a URL-encoded param
// string. For example, if a caller passes the following string in
// encodedParamString:oauth_token=ABC&oauth_token_secret=123&oauth_callback_confirmed=true and
// "oauth_token_secret" in paramName, then the return value would be "123".
// return a string or nil if failed.
func (z *HttpResponse) UrlEncParamValue(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkHttpResponse_urlEncParamValue(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


