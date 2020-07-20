// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkCharset.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCharset() *Charset {
	return &Charset{C.CkCharset_Create()}
}

func (z *Charset) DisposeCharset() {
    if z != nil {
	C.CkCharset_Dispose(z.ckObj)
	}
}




// property: AltToCharset
// If the ErrorAction property is set to 6, then this property controls how errors
// are handled. It specifies an alternate "To" charset. When a character in the
// input data cannot be converted to the target charset, an attempt is made to
// convert it to the AltToCharset. If that fails, the input character is dropped.
func (z *Charset) AltToCharset() string {
    return C.GoString(C.CkCharset_altToCharset(z.ckObj))
}
// property setter: AltToCharset
func (z *Charset) SetAltToCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkCharset_putAltToCharset(z.ckObj,cStr)
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
func (z *Charset) DebugLogFilePath() string {
    return C.GoString(C.CkCharset_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Charset) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkCharset_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ErrorAction
// Controls how errors are handled. When a character in the input data cannot be
// converted to the target charset, the action taken is controlled by this
// property. The possible settings are: (0) drop the error characters, (1)
// substitute the data set by the SetErrorBytes or SetErrorString method, (2)
// convert to a hex-escaped string (
func (z *Charset) ErrorAction() int {
    return int(C.CkCharset_getErrorAction(z.ckObj))
}
// property setter: ErrorAction
func (z *Charset) SetErrorAction(value int) {
    C.CkCharset_putErrorAction(z.ckObj,C.int(value))
}

// property: FromCharset
// Tells the charset converter the charset of the input data for a conversion.
// Possible values are:
// us-ascii
// unicode  (also known as UTF16LE or simply UTF16)
// unicodefffe  (also known as UTF16BE)
// ebcdic
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
// ibm01140
// ibm01141
// ibm01142
// ibm01143
// ibm01144
// ibm01145
// ibm01146
// ibm01147
// ibm01148
// ibm01149
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
func (z *Charset) FromCharset() string {
    return C.GoString(C.CkCharset_fromCharset(z.ckObj))
}
// property setter: FromCharset
func (z *Charset) SetFromCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkCharset_putFromCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Charset) LastErrorHtml() string {
    return C.GoString(C.CkCharset_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Charset) LastErrorText() string {
    return C.GoString(C.CkCharset_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Charset) LastErrorXml() string {
    return C.GoString(C.CkCharset_lastErrorXml(z.ckObj))
}

// property: LastInputAsHex
// If SaveLast is set to true, then the input and output of a conversion is saved
// to allow the exact bytes that are sent to the converter to be seen (for
// debugging purposes). This property shows the last input data in a
// hexidecimalized string.
func (z *Charset) LastInputAsHex() string {
    return C.GoString(C.CkCharset_lastInputAsHex(z.ckObj))
}

// property: LastInputAsQP
// If SaveLast is set to true, then the input and output of a conversion is saved
// to allow the exact bytes that are sent to the converter to be seen (for
// debugging purposes). This property shows the last input data in a
// quoted-printable string.
func (z *Charset) LastInputAsQP() string {
    return C.GoString(C.CkCharset_lastInputAsQP(z.ckObj))
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
func (z *Charset) LastMethodSuccess() bool {
    v := int(C.CkCharset_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Charset) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCharset_putLastMethodSuccess(z.ckObj,v)
}

// property: LastOutputAsHex
// If SaveLast is set to true, then the input and output of a conversion is saved
// to allow the exact bytes that are sent to the converter to be seen (for
// debugging purposes). This property shows the last output data in a
// hexidecimalized string.
func (z *Charset) LastOutputAsHex() string {
    return C.GoString(C.CkCharset_lastOutputAsHex(z.ckObj))
}

// property: LastOutputAsQP
// If SaveLast is set to true, then the input and output of a conversion is saved
// to allow the exact bytes that are sent to the converter to be seen (for
// debugging purposes). This property shows the last output data in a
// quoted-printable string.
func (z *Charset) LastOutputAsQP() string {
    return C.GoString(C.CkCharset_lastOutputAsQP(z.ckObj))
}

// property: SaveLast
// Tells the component to keep the input/output byte data in memory after a
// conversion is complete so the data can be examined via the LastInputAsHex/QP and
// LastOutputAsHex/QP properties. (for debugging purposes)
func (z *Charset) SaveLast() bool {
    v := int(C.CkCharset_getSaveLast(z.ckObj))
    return v != 0
}
// property setter: SaveLast
func (z *Charset) SetSaveLast(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCharset_putSaveLast(z.ckObj,v)
}

// property: ToCharset
// Tells the charset converter the target charset for a conversion. Possible values
// are:
// us-ascii
// unicode  (also known as UTF16LE or simply UTF16)
// unicodefffe  (also known as UTF16BE)
// ebcdic
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
// ibm01140
// ibm01141
// ibm01142
// ibm01143
// ibm01144
// ibm01145
// ibm01146
// ibm01147
// ibm01148
// ibm01149
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
func (z *Charset) ToCharset() string {
    return C.GoString(C.CkCharset_toCharset(z.ckObj))
}
// property setter: ToCharset
func (z *Charset) SetToCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkCharset_putToCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Charset) VerboseLogging() bool {
    v := int(C.CkCharset_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Charset) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCharset_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Charset) Version() string {
    return C.GoString(C.CkCharset_version(z.ckObj))
}
// Converts a charset name to a code page number. For example, "iso-8859-1"
// converts to code page 28591.
func (z *Charset) CharsetToCodePage(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkCharset_CharsetToCodePage(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Converts a code page number to a charset name. For example, 65001 converts to
// "utf-8".
// return a string or nil if failed.
func (z *Charset) CodePageToCharset(arg1 int) *string {

    retStr := C.CkCharset_codePageToCharset(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Converts character data from one charset to another. Before calling ConvertData,
// the FromCharset and ToCharset properties must be set to the source and
// destination charset names, such as "iso-8859-1" or "Shift_JIS".
func (z *Charset) ConvertData(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCharset_ConvertData(z.ckObj, ckbyd1, ckOutBytes)

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


// Converts a file from one character encoding to another. The FromCharset and
// ToCharset properties specify the source and destination character encodings. If
// the ToCharset is utf-16 or utf-8, then the preamble (also known as BOM) is
// included in the output. (Call ConvertFileNoPreamble to suppress the output of
// the BOM.)
func (z *Charset) ConvertFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCharset_ConvertFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Converts a file from one character encoding to another. The FromCharset and
// ToCharset properties specify the source and destination character encodings. No
// preamble (also known as BOM) is included in the output.
func (z *Charset) ConvertFileNoPreamble(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCharset_ConvertFileNoPreamble(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Converts Unicode (utf-16) text to the charset specified by the ToCharset
// property.
func (z *Charset) ConvertFromUnicode(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCharset_ConvertFromUnicode(z.ckObj, cstr1, ckOutBytes)

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


// Converts utf-16 text to the charset specified by the ToCharset property.
func (z *Charset) ConvertFromUtf16(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCharset_ConvertFromUtf16(z.ckObj, ckbyd1, ckOutBytes)

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


// Converts HTML text from one character encoding to another. The FromCharset and
// ToCharset properties must be set prior to calling this method. This method
// automatically edits the META tag within the HTML that indicates the charset.
func (z *Charset) ConvertHtml(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCharset_ConvertHtml(z.ckObj, ckbyd1, ckOutBytes)

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


// Converts an HTML file from one character encoding to another. The ToCharset
// properties must be set prior to calling this method. If the FromCharset is not
// set, it is obtained from the HTML META tag that indicates the charset. This
// method automatically edits the META tag within the HTML that indicates the
// charset.
func (z *Charset) ConvertHtmlFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCharset_ConvertHtmlFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Converts multibyte character data to a Unicode string. The FromCharset property
// should be set before calling this method.
// return a string or nil if failed.
func (z *Charset) ConvertToUnicode(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCharset_convertToUnicode(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// To be documented soon.
func (z *Charset) ConvertToUtf16(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCharset_ConvertToUtf16(z.ckObj, ckbyd1, ckOutBytes)

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


// Converts non-US-ASCII characters to Unicode decimal entities (_AMP_#xxxxx;)
// return a string or nil if failed.
func (z *Charset) EntityEncodeDec(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCharset_entityEncodeDec(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Converts non-US-ASCII characters to Unicode hex entities (_AMP_#xXXXX;)
// return a string or nil if failed.
func (z *Charset) EntityEncodeHex(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCharset_entityEncodeHex(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Examines HTML text and extracts the charset name specified by the META tag, if
// present.
// return a string or nil if failed.
func (z *Charset) GetHtmlCharset(arg1 []byte) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    retStr := C.CkCharset_getHtmlCharset(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Examines an HTML file and extracts the charset name specified by the META tag,
// if present.
// return a string or nil if failed.
func (z *Charset) GetHtmlFileCharset(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCharset_getHtmlFileCharset(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Converts HTML entities to Unicode characters.
// return a string or nil if failed.
func (z *Charset) HtmlDecodeToStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCharset_htmlDecodeToStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Decodes HTML entities. See http://www.w3.org/TR/REC-html40/sgml/entities.html
// for information on HTML entities. Examples of HTML entities are _AMP_lt; ,
// _AMP_#229; , _AMP_#xE5; , _AMP_#x6C34; , _AMP_Iacute; , etc.
func (z *Charset) HtmlEntityDecode(arg1 []byte) []byte {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCharset_HtmlEntityDecode(z.ckObj, ckbyd1, ckOutBytes)

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


// Decodes HTML entities in a file and creates a new HTML file with the entities
// decoded. See http://www.w3.org/TR/REC-html40/sgml/entities.html for information
// on HTML entities. Examples of HTML entities are _AMP_lt; , _AMP_#229; ,
// _AMP_#xE5; , _AMP_#x6C34; , _AMP_Iacute; , etc.
func (z *Charset) HtmlEntityDecodeFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCharset_HtmlEntityDecodeFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns true if the component is unlocked.
func (z *Charset) IsUnlocked() bool {

    v := C.CkCharset_IsUnlocked(z.ckObj)


    return v != 0
}


// Converts a string to lowercase.
// return a string or nil if failed.
func (z *Charset) LowerCase(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCharset_lowerCase(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Convenience method for reading the entire contents of a file into a byte array.
func (z *Charset) ReadFile(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCharset_ReadFile(z.ckObj, cstr1, ckOutBytes)

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


// Reads a text file and returns the text converted to a Unicode string. The
// filename is specified by the first method argument, and the charset of the text
// data is specified by the 2nd method argument.
// return a string or nil if failed.
func (z *Charset) ReadFileToString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkCharset_readFileToString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// If the ErrorAction property is set to 1, the bytes passed to this method are
// used as the result for any characters that cannot be converted during a
// conversion.
func (z *Charset) SetErrorBytes(arg1 []byte)  {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    C.CkCharset_SetErrorBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)


}


// If the ErrorAction property is set to 1, the string passed to this method is
// used as the result for any characters that cannot be converted during a
// conversion.
func (z *Charset) SetErrorString(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkCharset_SetErrorString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Unlocks the component. This method must be called once at the beginning of the
// program. Properties can be get/set without unlocking, but methods will not work
// unless the component has been unlocked.
func (z *Charset) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCharset_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Converts a string to uppercase.
// return a string or nil if failed.
func (z *Charset) UpperCase(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCharset_upperCase(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// URL decodes a string.
// return a string or nil if failed.
func (z *Charset) UrlDecodeStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCharset_urlDecodeStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if the byte data conforms to the charset passed in the first
// argument.
func (z *Charset) VerifyData(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkCharset_VerifyData(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Returns true if the file contains character data that conforms to the charset
// passed in the 1st argument.
func (z *Charset) VerifyFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCharset_VerifyFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Convenience method for saving an entire byte array to a file.
func (z *Charset) WriteFile(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkCharset_WriteFile(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Converts a Unicode string to a multibyte charset and writes the multibyte text
// data to a file. The destination charset is specified in the 2nd method argument.
func (z *Charset) WriteStringToFile(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkCharset_WriteStringToFile(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


