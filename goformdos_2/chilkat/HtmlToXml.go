// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkHtmlToXml.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewHtmlToXml() *HtmlToXml {
	return &HtmlToXml{C.CkHtmlToXml_Create()}
}

func (z *HtmlToXml) DisposeHtmlToXml() {
    if z != nil {
	C.CkHtmlToXml_Dispose(z.ckObj)
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
func (z *HtmlToXml) DebugLogFilePath() string {
    return C.GoString(C.CkHtmlToXml_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *HtmlToXml) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkHtmlToXml_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DropCustomTags
// If set to true, then any non-standard HTML tags will be dropped when converting
// to XML.
func (z *HtmlToXml) DropCustomTags() bool {
    v := int(C.CkHtmlToXml_getDropCustomTags(z.ckObj))
    return v != 0
}
// property setter: DropCustomTags
func (z *HtmlToXml) SetDropCustomTags(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHtmlToXml_putDropCustomTags(z.ckObj,v)
}

// property: Html
// The HTML to be converted by the ToXml method. To convert HTML to XML, first set
// this property to the HTML string and then call ToXml. The ConvertFile method can
// do file-to-file conversions.
func (z *HtmlToXml) Html() string {
    return C.GoString(C.CkHtmlToXml_html(z.ckObj))
}
// property setter: Html
func (z *HtmlToXml) SetHtml(goStr string) {
    cStr := C.CString(goStr)
    C.CkHtmlToXml_putHtml(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *HtmlToXml) LastErrorHtml() string {
    return C.GoString(C.CkHtmlToXml_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *HtmlToXml) LastErrorText() string {
    return C.GoString(C.CkHtmlToXml_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *HtmlToXml) LastErrorXml() string {
    return C.GoString(C.CkHtmlToXml_lastErrorXml(z.ckObj))
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
func (z *HtmlToXml) LastMethodSuccess() bool {
    v := int(C.CkHtmlToXml_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *HtmlToXml) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHtmlToXml_putLastMethodSuccess(z.ckObj,v)
}

// property: Nbsp
// Determines how to handle   HTML entities. The default value, 0 will cause
// _AMP_nbsp; entites to be convert to normal space characters (ASCII value 32). If
// this property is set to 1, then _AMP_nbsp;'s will be converted to _AMP_#160. If
// set to 2, then _AMP_nbps;'s are dropped. If set to 3, then _AMP_nbsp's are left
// unmodified.
func (z *HtmlToXml) Nbsp() int {
    return int(C.CkHtmlToXml_getNbsp(z.ckObj))
}
// property setter: Nbsp
func (z *HtmlToXml) SetNbsp(value int) {
    C.CkHtmlToXml_putNbsp(z.ckObj,C.int(value))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *HtmlToXml) VerboseLogging() bool {
    v := int(C.CkHtmlToXml_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *HtmlToXml) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHtmlToXml_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *HtmlToXml) Version() string {
    return C.GoString(C.CkHtmlToXml_version(z.ckObj))
}

// property: XmlCharset
// The charset, such as "utf-8" or "iso-8859-1" of the XML to be created. If
// XmlCharset is empty, the XML is created in the same character encoding as the
// HTML. Otherwise the HTML is converted XML and converted to this charset.
func (z *HtmlToXml) XmlCharset() string {
    return C.GoString(C.CkHtmlToXml_xmlCharset(z.ckObj))
}
// property setter: XmlCharset
func (z *HtmlToXml) SetXmlCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkHtmlToXml_putXmlCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}
// Converts an HTML file to a well-formed XML file that can be parsed for the
// purpose of programmatically extracting information.
func (z *HtmlToXml) ConvertFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHtmlToXml_ConvertFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Allows for any specified tag to be dropped from the output XML. To drop more
// than one tag, call this method once for each tag type to be dropped.
func (z *HtmlToXml) DropTagType(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkHtmlToXml_DropTagType(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Causes text formatting tags to be dropped from the XML output. Text formatting
// tags are: b, font, i, u, br, center, em, strong, big, tt, s, small, strike, sub,
// and sup.
func (z *HtmlToXml) DropTextFormattingTags()  {

    C.CkHtmlToXml_DropTextFormattingTags(z.ckObj)



}


// Returns true if the component is already unlocked. Otherwise returns false.
func (z *HtmlToXml) IsUnlocked() bool {

    v := C.CkHtmlToXml_IsUnlocked(z.ckObj)


    return v != 0
}


// Convenience method for reading a complete file into a byte array.
func (z *HtmlToXml) ReadFile(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkHtmlToXml_ReadFile(z.ckObj, cstr1, ckOutBytes)

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


// Convenience method for reading a text file into a string. The character encoding
// of the text file is specified by srcCharset. Valid values, such as "iso-8895-1" or
// "utf-8" are listed at:List of Charsets
// <http://blog.chilkatsoft.com/?p=463>.
// return a string or nil if failed.
func (z *HtmlToXml) ReadFileToString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkHtmlToXml_readFileToString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sets the Html property from the contents of bd.
func (z *HtmlToXml) SetHtmlBd(arg1 *BinData) bool {

    v := C.CkHtmlToXml_SetHtmlBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the Html property from a byte array.
func (z *HtmlToXml) SetHtmlBytes(arg1 []byte)  {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    C.CkHtmlToXml_SetHtmlBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)


}


// Sets the Html property by loading the HTML from a file.
func (z *HtmlToXml) SetHtmlFromFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHtmlToXml_SetHtmlFromFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Converts the HTML in the "Html" property to XML and returns the XML string.
// return a string or nil if failed.
func (z *HtmlToXml) ToXml() *string {

    retStr := C.CkHtmlToXml_toXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Converts the HTML in the "Html" property to XML and appends the XML to sb.
func (z *HtmlToXml) ToXmlSb(arg1 *StringBuilder) bool {

    v := C.CkHtmlToXml_ToXmlSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Causes a specified type of tag to NOT be dropped in the output XML.
func (z *HtmlToXml) UndropTagType(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkHtmlToXml_UndropTagType(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Causes text formatting tags to NOT be dropped from the XML output. Text
// formatting tags are: b, font, i, u, br, center, em, strong, big, tt, s, small,
// strike, sub, and sup.
// 
// Important: Text formatting tags are dropped by default. Call this method to
// prevent text formatting tags from being dropped.
//
func (z *HtmlToXml) UndropTextFormattingTags()  {

    C.CkHtmlToXml_UndropTextFormattingTags(z.ckObj)



}


// Unlocks the component. An arbitrary unlock code may be passed to automatically
// begin a 30-day trial.
func (z *HtmlToXml) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHtmlToXml_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Convenience method for saving a byte array to a file.
func (z *HtmlToXml) WriteFile(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkHtmlToXml_WriteFile(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Convenience method for saving a string to a file. The character encoding of the
// output text file is specified by charset (the string is converted to this charset
// when writing). Valid values, such as "iso-8895-1" or "utf-8" are listed at:List
// of Charsets
// <http://blog.chilkatsoft.com/?p=463>.
func (z *HtmlToXml) WriteStringToFile(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkHtmlToXml_WriteStringToFile(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


