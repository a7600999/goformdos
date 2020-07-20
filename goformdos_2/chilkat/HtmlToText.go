// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkHtmlToText.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewHtmlToText() *HtmlToText {
	return &HtmlToText{C.CkHtmlToText_Create()}
}

func (z *HtmlToText) DisposeHtmlToText() {
    if z != nil {
	C.CkHtmlToText_Dispose(z.ckObj)
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
func (z *HtmlToText) DebugLogFilePath() string {
    return C.GoString(C.CkHtmlToText_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *HtmlToText) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkHtmlToText_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DecodeHtmlEntities
// If true, then HTML entities are automatically decoded. For example _AMP_amp;
// is automatically decoded to _AMP_. If this property is set to false, then HTML
// entities are not decoded. The default value is true.
func (z *HtmlToText) DecodeHtmlEntities() bool {
    v := int(C.CkHtmlToText_getDecodeHtmlEntities(z.ckObj))
    return v != 0
}
// property setter: DecodeHtmlEntities
func (z *HtmlToText) SetDecodeHtmlEntities(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHtmlToText_putDecodeHtmlEntities(z.ckObj,v)
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *HtmlToText) LastErrorHtml() string {
    return C.GoString(C.CkHtmlToText_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *HtmlToText) LastErrorText() string {
    return C.GoString(C.CkHtmlToText_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *HtmlToText) LastErrorXml() string {
    return C.GoString(C.CkHtmlToText_lastErrorXml(z.ckObj))
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
func (z *HtmlToText) LastMethodSuccess() bool {
    v := int(C.CkHtmlToText_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *HtmlToText) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHtmlToText_putLastMethodSuccess(z.ckObj,v)
}

// property: RightMargin
// Used to control wrapping of text. The default value is 80. When the text gets
// close to this margin, the converter will try to break the line at a SPACE
// character. Set this property to 0 for no right margin.
func (z *HtmlToText) RightMargin() int {
    return int(C.CkHtmlToText_getRightMargin(z.ckObj))
}
// property setter: RightMargin
func (z *HtmlToText) SetRightMargin(value int) {
    C.CkHtmlToText_putRightMargin(z.ckObj,C.int(value))
}

// property: SuppressLinks
// If false, then link URL's are preserved inline. For example, the following
// HTML fragment:
// 
// _LT_p>Test _LT_a href="http://www.chilkatsoft.com/">chilkat_LT_/a>_LT_/p>
// 
// converts to:
// Test chilkat _LT_http://www.chilkatsoft.com/>
// If this property is true, the above HTML would convert to:
// Test chilkat
// The default value of this property is true.
//
func (z *HtmlToText) SuppressLinks() bool {
    v := int(C.CkHtmlToText_getSuppressLinks(z.ckObj))
    return v != 0
}
// property setter: SuppressLinks
func (z *HtmlToText) SetSuppressLinks(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHtmlToText_putSuppressLinks(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *HtmlToText) VerboseLogging() bool {
    v := int(C.CkHtmlToText_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *HtmlToText) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHtmlToText_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *HtmlToText) Version() string {
    return C.GoString(C.CkHtmlToText_version(z.ckObj))
}
// Returns true if the component is already unlocked. Otherwise returns false.
func (z *HtmlToText) IsUnlocked() bool {

    v := C.CkHtmlToText_IsUnlocked(z.ckObj)


    return v != 0
}


// Convenience method for reading a text file into a string. The character encoding
// of the text file is specified by srcCharset. Valid values, such as "iso-8895-1" or
// "utf-8" are listed at:List of Charsets
// <http://blog.chilkatsoft.com/?p=463>.
// return a string or nil if failed.
func (z *HtmlToText) ReadFileToString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkHtmlToText_readFileToString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Converts HTML to plain-text.
// return a string or nil if failed.
func (z *HtmlToText) ToText(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHtmlToText_toText(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Unlocks the component. An arbitrary unlock code may be passed to automatically
// begin a 30-day trial.
// 
// This class is included with the Chilkat HTML-to-XML conversion component
// license.
//
func (z *HtmlToText) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHtmlToText_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Convenience method for saving a string to a file. The character encoding of the
// output text file is specified by charset (the string is converted to this charset
// when writing). Valid values, such as "iso-8895-1" or "utf-8" are listed at:List
// of Charsets
// <http://blog.chilkatsoft.com/?p=463>.
func (z *HtmlToText) WriteStringToFile(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkHtmlToText_WriteStringToFile(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


