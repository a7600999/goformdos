// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkHttpRequest.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewHttpRequest() *HttpRequest {
	return &HttpRequest{C.CkHttpRequest_Create()}
}

func (z *HttpRequest) DisposeHttpRequest() {
    if z != nil {
	C.CkHttpRequest_Dispose(z.ckObj)
	}
}




// property: Boundary
// Sets an explicit boundary string to be used in multipart/form-data requests. If
// no Boundary is set, then a boundary string is automaticaly generated as needed
// during the sending of a request.
func (z *HttpRequest) Boundary() string {
    return C.GoString(C.CkHttpRequest_boundary(z.ckObj))
}
// property setter: Boundary
func (z *HttpRequest) SetBoundary(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttpRequest_putBoundary(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Charset
// Controls the character encoding used for HTTP request parameters for POST
// requests. The default value is the ANSI charset of the computer. The charset
// should match the charset expected by the form target.
// 
// The "charset" attribute is only included in the Content-Type header of the
// request if the SendCharset property is set to true.
//
func (z *HttpRequest) Charset() string {
    return C.GoString(C.CkHttpRequest_charset(z.ckObj))
}
// property setter: Charset
func (z *HttpRequest) SetCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttpRequest_putCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ContentType
// The ContentType property sets the "Content-Type" header field, and identifies
// the content-type of the HTTP request body. Common values are:
// 
//         
// application/x-www-form-urlencoded    
// multipart/form-data    
// application/json    
// application/xml    
//     
// 
// If ContentType is set equal to the empty string, then no Content-Type header is
// included in the HTTP request.
func (z *HttpRequest) ContentType() string {
    return C.GoString(C.CkHttpRequest_contentType(z.ckObj))
}
// property setter: ContentType
func (z *HttpRequest) SetContentType(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttpRequest_putContentType(z.ckObj,cStr)
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
func (z *HttpRequest) DebugLogFilePath() string {
    return C.GoString(C.CkHttpRequest_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *HttpRequest) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttpRequest_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EntireHeader
// Composes and returns the entire MIME header of the HTTP request.
func (z *HttpRequest) EntireHeader() string {
    return C.GoString(C.CkHttpRequest_entireHeader(z.ckObj))
}
// property setter: EntireHeader
func (z *HttpRequest) SetEntireHeader(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttpRequest_putEntireHeader(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpVerb
// The HttpVerb property should be set to the name of the HTTP method that appears
// on the "start line" of an HTTP request, such as GET, POST, PUT, DELETE, etc. It
// is also possible to use the various WebDav verbs such as PROPFIND, PROPPATCH,
// MKCOL, COPY, MOVE, LOCK, UNLOCK, etc. In general, the HttpVerb may be set to
// anything, even custom verbs recognized by a custom server-side app.
func (z *HttpRequest) HttpVerb() string {
    return C.GoString(C.CkHttpRequest_httpVerb(z.ckObj))
}
// property setter: HttpVerb
func (z *HttpRequest) SetHttpVerb(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttpRequest_putHttpVerb(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: HttpVersion
// The HTTP version in the request header. Defaults to "1.1".
func (z *HttpRequest) HttpVersion() string {
    return C.GoString(C.CkHttpRequest_httpVersion(z.ckObj))
}
// property setter: HttpVersion
func (z *HttpRequest) SetHttpVersion(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttpRequest_putHttpVersion(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *HttpRequest) LastErrorHtml() string {
    return C.GoString(C.CkHttpRequest_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *HttpRequest) LastErrorText() string {
    return C.GoString(C.CkHttpRequest_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *HttpRequest) LastErrorXml() string {
    return C.GoString(C.CkHttpRequest_lastErrorXml(z.ckObj))
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
func (z *HttpRequest) LastMethodSuccess() bool {
    v := int(C.CkHttpRequest_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *HttpRequest) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttpRequest_putLastMethodSuccess(z.ckObj,v)
}

// property: NumHeaderFields
// Returns the number of request header fields.
func (z *HttpRequest) NumHeaderFields() int {
    return int(C.CkHttpRequest_getNumHeaderFields(z.ckObj))
}

// property: NumParams
// Returns the number of query parameters.
func (z *HttpRequest) NumParams() int {
    return int(C.CkHttpRequest_getNumParams(z.ckObj))
}

// property: Path
// The path of the resource requested. A path of "/" indicates the default document
// of a domain.
// 
// Explaining the Parts of a URL
// 
// http://example.com:8042/over/there?name=ferret#nose
// \__/   \______________/\_________/ \________/ \__/
//  |           |            |            |        |
// scheme   domain+port     path        query   fragment
// 
// This property should be set to the path part of the URL. You may also include
// the query part in this property value. If the Content-Type of the request is NOT
// application/x-www-form-urlencoded, then you would definitely want to include
// query parameters in the path. If the Content-Type of the request IS
// application/x-www-form-urlencoded, the query parameters are passed in the body
// of the request. It is also possible to pass some query parameters via the path,
// and some in the body of a application/x-www-form-urlencoded request, but you
// shouldn't include the same parameter in both places. You would never need to
// include the fragment part. The fragment is nothing more than an instruction for
// a browser to automatically navigate to a particular location in the HTML page
// (assuming the request returns HTML, otherwise a fragment makes no sense).
//
func (z *HttpRequest) Path() string {
    return C.GoString(C.CkHttpRequest_path(z.ckObj))
}
// property setter: Path
func (z *HttpRequest) SetPath(goStr string) {
    cStr := C.CString(goStr)
    C.CkHttpRequest_putPath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SendCharset
// Controls whether the charset is explicitly included in the content-type header
// field of the HTTP POST request. The default value of this property is false.
func (z *HttpRequest) SendCharset() bool {
    v := int(C.CkHttpRequest_getSendCharset(z.ckObj))
    return v != 0
}
// property setter: SendCharset
func (z *HttpRequest) SetSendCharset(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttpRequest_putSendCharset(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *HttpRequest) VerboseLogging() bool {
    v := int(C.CkHttpRequest_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *HttpRequest) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHttpRequest_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *HttpRequest) Version() string {
    return C.GoString(C.CkHttpRequest_version(z.ckObj))
}
// Adds a file to an upload request where the contents of the file come from byteData.
// 
// name is an arbitrary name. (In HTML, it is the form field name of the input
// tag.)
// remoteFilename is the name of the file to be created on the HTTP server.
// byteData contains the bytes to be uploaded.
// contentType contains is the value of the Content-Type header. An empty string may be
// passed to allow Chilkat to automatically determine the Content-Type based on the
// filename extension.
//
func (z *HttpRequest) AddBdForUpload(arg1 string, arg2 string, arg3 *BinData, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr4 := C.CString(arg4)

    v := C.CkHttpRequest_AddBdForUpload(z.ckObj, cstr1, cstr2, arg3.ckObj, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Adds a file to an upload request where the contents of the file come from an
// in-memory byte array. To create a file upload request, call UseUpload and then
// call AddBytesForUpload, AddStringForUpload, or AddFileForUpload for each file to
// be uploaded.
// 
// name is an arbitrary name. (In HTML, it is the form field name of the input
// tag.)
// remoteFileName is the name of the file to be created on the HTTP server.
// byteData contains the contents (bytes) to be uploaded.
//
func (z *HttpRequest) AddBytesForUpload(arg1 string, arg2 string, arg3 []byte) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkHttpRequest_AddBytesForUpload(z.ckObj, cstr1, cstr2, ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// Same as AddBytesForUpload, but allows the Content-Type header field to be
// directly specified. (Otherwise, the Content-Type header is automatically
// determined based on the remoteFileName's file extension.)
func (z *HttpRequest) AddBytesForUpload2(arg1 string, arg2 string, arg3 []byte, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))
    cstr4 := C.CString(arg4)

    v := C.CkHttpRequest_AddBytesForUpload2(z.ckObj, cstr1, cstr2, ckbyd3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Adds a file to an upload request. To create a file upload request, call
// UseUpload and then call AddFileForUpload, AddBytesForUpload, or
// AddStringForUpload for each file to be uploaded. This method does not read the
// file into memory. When the upload occurs, the data is streamed directly from the
// file, thus allowing for very large files to be uploaded without consuming large
// amounts of memory.
// 
// name is an arbitrary name. (In HTML, it is the form field name of the input
// tag.)
// filePath is the path to an existing file in the local filesystem.
//
func (z *HttpRequest) AddFileForUpload(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttpRequest_AddFileForUpload(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Same as AddFileForUpload, but allows the Content-Type header field to be
// directly specified. (Otherwise, the Content-Type header is automatically
// determined based on the file extension.)
// 
// name is an arbitrary name. (In HTML, it is the form field name of the input
// tag.)
// filePath is the path to an existing file in the local filesystem.
//
func (z *HttpRequest) AddFileForUpload2(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkHttpRequest_AddFileForUpload2(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds a request header to the HTTP request. If a header having the same field
// name is already present, this method replaces it.
// 
// Note: Never explicitly set the Content-Length header field. Chilkat will
// automatically compute the correct length and add the Content-Length header to
// all POST, PUT, or any other request where the Content-Length needs to be
// specified. (GET requests always have a 0 length body, and therefore never need a
// Content-Length header field.)
//
func (z *HttpRequest) AddHeader(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkHttpRequest_AddHeader(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Computes the Amazon MWS signature using the mwsSecretKey and adds the "Signature"
// parameter to the request. This method should be called for all Amazon
// Marketplace Web Service (Amazon MWS) HTTP requests. It should be called after
// all request parameters have been added.
// 
// Important: The Chilkat v9.5.0.75 release accidentally breaks Amazon MWS (not
// AWS) authentication. If you need MWS with 9.5.0.75, send email to
// support@chilkatsoft.com for a hotfix, or revert back to v9.5.0.73, or update to
// a version after 9.5.0.75.
// 
// The domain should be the domain of the request, such as one of the following:
//     mws.amazonservices.com
//     mws-eu.amazonservices.com
//     mws.amazonservices.in
//     mws.amazonservices.com.cn
//     mws.amazonservices.jp
// 
// Note: This method automatically adds or replaces the existing Timestamp
// parameter to the current system date/time.
//
func (z *HttpRequest) AddMwsSignature(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttpRequest_AddMwsSignature(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a request query parameter (name/value pair) to the HTTP request. The name
// and value strings passed to this method should not be URL encoded.
func (z *HttpRequest) AddParam(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkHttpRequest_AddParam(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Same as AddFileForUpload, but the upload data comes from an in-memory string
// instead of a file.
func (z *HttpRequest) AddStringForUpload(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkHttpRequest_AddStringForUpload(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Same as AddStringForUpload, but allows the Content-Type header field to be
// directly specified. (Otherwise, the Content-Type header is automatically
// determined based on the filename's file extension.)
func (z *HttpRequest) AddStringForUpload2(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    v := C.CkHttpRequest_AddStringForUpload2(z.ckObj, cstr1, cstr2, cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))

    return v != 0
}


// Adds a request header to the Nth sub-header of the HTTP request. If a header
// having the same field name is already present, this method replaces it.
func (z *HttpRequest) AddSubHeader(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkHttpRequest_AddSubHeader(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// The same as GenerateRequestText, except the generated request is written to the
// file specified by path.
func (z *HttpRequest) GenerateRequestFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttpRequest_GenerateRequestFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the request text that would be sent if Http.SynchronousRequest was
// called.
// return a string or nil if failed.
func (z *HttpRequest) GenerateRequestText() *string {

    retStr := C.CkHttpRequest_generateRequestText(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of a request header field.
// return a string or nil if failed.
func (z *HttpRequest) GetHeaderField(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttpRequest_getHeaderField(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth request header field name. Indexing begins at 0, and the number
// of request header fields is specified by the NumHeaderFields property.
// return a string or nil if failed.
func (z *HttpRequest) GetHeaderName(arg1 int) *string {

    retStr := C.CkHttpRequest_getHeaderName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth request header field value. Indexing begins at 0, and the number
// of request header fields is specified by the NumHeaderFields property.
// return a string or nil if failed.
func (z *HttpRequest) GetHeaderValue(arg1 int) *string {

    retStr := C.CkHttpRequest_getHeaderValue(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a request query parameter value by name.
// return a string or nil if failed.
func (z *HttpRequest) GetParam(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHttpRequest_getParam(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth request query parameter field name. Indexing begins at 0, and
// the number of request query parameter fields is specified by the NumParams
// property.
// return a string or nil if failed.
func (z *HttpRequest) GetParamName(arg1 int) *string {

    retStr := C.CkHttpRequest_getParamName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth request query parameter field value. Indexing begins at 0, and
// the number of request query parameter fields is specified by the NumParams
// property.
// return a string or nil if failed.
func (z *HttpRequest) GetParamValue(arg1 int) *string {

    retStr := C.CkHttpRequest_getParamValue(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the request parameters in URL encoded form (i.e. in the exact form that
// would be sent if the ContentType property was
// application/x-www-form-urlencoded). For example, if a request has two params:
// param1="abc 123" and param2="abc-123", then GetUrlEncodedParams would return
// "abc+123
// return a string or nil if failed.
func (z *HttpRequest) GetUrlEncodedParams() *string {

    retStr := C.CkHttpRequest_getUrlEncodedParams(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Uses the contents of the requestBody as the HTTP request body.
func (z *HttpRequest) LoadBodyFromBd(arg1 *BinData) bool {

    v := C.CkHttpRequest_LoadBodyFromBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// The HTTP protocol is such that all HTTP requests are MIME. For non-multipart
// requests, this method may be called to set the MIME body of the HTTP request to
// the exact contents of the byteData.
// Note: A non-multipart HTTP request consists of (1) the HTTP start line, (2) MIME
// header fields, and (3) the MIME body. This method sets the MIME body.
func (z *HttpRequest) LoadBodyFromBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkHttpRequest_LoadBodyFromBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// The HTTP protocol is such that all HTTP requests are MIME. For non-multipart
// requests, this method may be called to set the MIME body of the HTTP request to
// the exact contents of filePath.
// Note: A non-multipart HTTP request consists of (1) the HTTP start line, (2) MIME
// header fields, and (3) the MIME body. This method sets the MIME body.
func (z *HttpRequest) LoadBodyFromFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttpRequest_LoadBodyFromFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Uses the contents of the requestBody as the HTTP request body. The charset indicates the
// binary representation of the string, such as "utf-8", "utf-16", "iso-8859-*",
// "windows-125*", etc. Any of the character encodings supported at the link below
// are valid.
func (z *HttpRequest) LoadBodyFromSb(arg1 *StringBuilder, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkHttpRequest_LoadBodyFromSb(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// The HTTP protocol is such that all HTTP requests are MIME. For non-multipart
// requests, this method may be called to set the MIME body of the HTTP request to
// the exact contents of bodyStr.
// Note: A non-multipart HTTP request consists of (1) the HTTP start line, (2) MIME
// header fields, and (3) the MIME body. This method sets the MIME body.
// 
// charset indicates the charset, such as "utf-8" or "iso-8859-1", to be used. The
// HTTP body will contain the bodyStr converted to this character encoding.
//
func (z *HttpRequest) LoadBodyFromString(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHttpRequest_LoadBodyFromString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Removes all request parameters.
func (z *HttpRequest) RemoveAllParams()  {

    C.CkHttpRequest_RemoveAllParams(z.ckObj)



}


// Removes all occurrences of a HTTP request header field. Always returns true.
func (z *HttpRequest) RemoveHeader(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttpRequest_RemoveHeader(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Removes a single HTTP request parameter by name.
func (z *HttpRequest) RemoveParam(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkHttpRequest_RemoveParam(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Parses a URL and sets the Path and query parameters (NumParams, GetParam,
// GetParamName, GetParamValue).
func (z *HttpRequest) SetFromUrl(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkHttpRequest_SetFromUrl(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Useful for sending HTTP requests where the body is a very large file. For
// example, to send an XML HttpRequest containing a very large XML document, one
// would set the HttpVerb = "POST", the ContentType = "text/xml", and then call
// StreamBodyFromFile to indicate that the XML body of the request is to be
// streamed directly from a file. When the HTTP request is actually sent, the body
// is streamed directly from the file, and thus the file never needs to be loaded
// in its entirety in memory.
func (z *HttpRequest) StreamBodyFromFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHttpRequest_StreamBodyFromFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// This method is the same as StreamBodyFromFile, but allows for an offset and
// number of bytes to be specified. The offset and numBytes are integers passed as
// strings.
func (z *HttpRequest) StreamChunkFromFile(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkHttpRequest_StreamChunkFromFile(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


