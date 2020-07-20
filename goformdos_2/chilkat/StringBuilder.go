// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkStringBuilder.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{C.CkStringBuilder_Create()}
}

func (z *StringBuilder) DisposeStringBuilder() {
    if z != nil {
	C.CkStringBuilder_Dispose(z.ckObj)
	}
}




// property: IntValue
// Returns the content of the string converted to an integer.
func (z *StringBuilder) IntValue() int {
    return int(C.CkStringBuilder_getIntValue(z.ckObj))
}
// property setter: IntValue
func (z *StringBuilder) SetIntValue(value int) {
    C.CkStringBuilder_putIntValue(z.ckObj,C.int(value))
}

// property: IsBase64
// Returns true if the content contains only those characters allowed in the
// base64 encoding. A base64 string is composed of characters 'A'..'Z', 'a'..'z',
// '0'..'9', '+', '/' and it is often padded at the end with up to two '=', to make
// the length a multiple of 4. Whitespace is ignored.
func (z *StringBuilder) IsBase64() bool {
    v := int(C.CkStringBuilder_getIsBase64(z.ckObj))
    return v != 0
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
func (z *StringBuilder) LastMethodSuccess() bool {
    v := int(C.CkStringBuilder_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *StringBuilder) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkStringBuilder_putLastMethodSuccess(z.ckObj,v)
}

// property: Length
// The number of characters of the string contained within this instance.
func (z *StringBuilder) Length() int {
    return int(C.CkStringBuilder_getLength(z.ckObj))
}
// Appends a copy of the specified string to this instance.
func (z *StringBuilder) Append(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringBuilder_Append(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends the contents of binData. The charset specifies the character encoding of the
// bytes contained in binData. The charset can be any of the supported encodings listed
// atChilkat Supported Character Encodings
// <http://cknotes.com/chilkat-charsets-character-encodings-supported/>. To append
// the entire contents of binData, set offset and numBytes equal to zero. To append a range
// of binData, set the offset and numBytes to specify the range.
func (z *StringBuilder) AppendBd(arg1 *BinData, arg2 string, arg3 int, arg4 int) bool {
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_AppendBd(z.ckObj, arg1.ckObj, cstr2, C.int(arg3), C.int(arg4))

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends binary data using the encoding specified by encoding, such as "base64",
// "hex", etc.
func (z *StringBuilder) AppendEncoded(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_AppendEncoded(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends the string representation of a specified 32-bit signed integer to this
// instance.
func (z *StringBuilder) AppendInt(arg1 int) bool {

    v := C.CkStringBuilder_AppendInt(z.ckObj, C.int(arg1))


    return v != 0
}


// Appends the string representation of a specified 64-bit signed integer to this
// instance.
func (z *StringBuilder) AppendInt64(arg1 int64) bool {

    v := C.CkStringBuilder_AppendInt64(z.ckObj, C.__int64(arg1))


    return v != 0
}


// Appends the value followed by a CRLF or LF to the end of the curent StringBuilder
// object. If crlf is true, then a CRLF line ending is used. Otherwise a LF line
// ending is used.
func (z *StringBuilder) AppendLine(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkStringBuilder_AppendLine(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends the contents of another StringBuilder to this instance.
func (z *StringBuilder) AppendSb(arg1 *StringBuilder) bool {

    v := C.CkStringBuilder_AppendSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Removes all characters from the current StringBuilder instance.
func (z *StringBuilder) Clear()  {

    C.CkStringBuilder_Clear(z.ckObj)



}


// Returns true if the str is contained within this object. For case sensitive
// matching, set caseSensitive equal to true. For case-insensitive, set caseSensitive equal to
// false.
func (z *StringBuilder) Contains(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkStringBuilder_Contains(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if the word is contained within this object, but only if it is a
// whole word. This method is limited to finding whole words in strings that only
// contains characters in the Latin1 charset (i.e. iso-8859-1 or Windows-1252). A
// whole word can only contain alphanumeric chars where the alpha chars are
// restricted to those of the Latin1 alpha chars. (The underscore character is also
// considered part of a word.)
// 
// For case sensitive matching, set caseSensitive equal to true. For case-insensitive, set
// caseSensitive equal to false.
//
func (z *StringBuilder) ContainsWord(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkStringBuilder_ContainsWord(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if the contents of this object equals the str. Returns false
// if unequal. For case insensitive equality, set caseSensitive equal to false.
func (z *StringBuilder) ContentsEqual(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkStringBuilder_ContentsEqual(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if the contents of this object equals the sb. Returns false
// if unequal. For case insensitive equality, set caseSensitive equal to false.
func (z *StringBuilder) ContentsEqualSb(arg1 *StringBuilder, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkStringBuilder_ContentsEqualSb(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Decodes and replaces the contents with the decoded string. The encoding can be set
// to any of the following strings: "base64", "hex", "quoted-printable" (or "qp"),
// "url", "base32", "Q", "B", "url_rc1738", "url_rfc2396", "url_rfc3986",
// "url_oauth", "uu", "modBase64", or "html" (for HTML entity encoding). The full
// up-to-date list of supported binary encodings is available at the link entitled
// "Supported Binary Encodings" below.
// 
// Note: This method can only be called if the encoded content decodes to a string.
// The charset indicates the charset to be used in intepreting the decoded bytes. For
// example, the charset can be "utf-8", "utf-16", "iso-8859-1", "shift_JIS", etc.
//
func (z *StringBuilder) Decode(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_Decode(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Encodes to base64, hex, quoted-printable, URL encoding, etc. The encoding can be set
// to any of the following strings: "base64", "hex", "quoted-printable" (or "qp"),
// "url", "base32", "Q", "B", "url_rc1738", "url_rfc2396", "url_rfc3986",
// "url_oauth", "uu", "modBase64", or "html" (for HTML entity encoding). The full
// up-to-date list of supported binary encodings is available at the link entitled
// "Supported Binary Encodings" below.
func (z *StringBuilder) Encode(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_Encode(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns true if the string ends with substr. Otherwise returns false. The
// comparison is case sensitive if caseSensitive is true, and case insensitive if caseSensitive is
// false.
func (z *StringBuilder) EndsWith(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkStringBuilder_EndsWith(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Decodes HTML entities. SeeHTML entities
// <https://duckduckgo.com/?q=html+entities&bext=wfp&ia=web> for more information
// about HTML entities.
func (z *StringBuilder) EntityDecode() bool {

    v := C.CkStringBuilder_EntityDecode(z.ckObj)


    return v != 0
}


// Begin searching after the 1st occurrence of searchAfter is found, and then return the
// substring found between the next occurrence of beginMark and the next occurrence of
// endMark.
// return a string or nil if failed.
func (z *StringBuilder) GetAfterBetween(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkStringBuilder_getAfterBetween(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the substring found after the final occurrence of marker. If removeFlag is
// true, the marker and the content that follows is removed from this content.
// 
// If the marker is not present, then the entire string is returned. In this case, if
// removeFlag is true, this object is also cleared.
//
// return a string or nil if failed.
func (z *StringBuilder) GetAfterFinal(arg1 string, arg2 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkStringBuilder_getAfterFinal(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the contents as a string.
// return a string or nil if failed.
func (z *StringBuilder) GetAsString() *string {

    retStr := C.CkStringBuilder_getAsString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the substring found before the 1st occurrence of marker. If removeFlag is
// true, the content up to and including the marker is removed from this object's
// contents.
// 
// If the marker is not present, then the entire string is returned. In this case, if
// removeFlag is true, this object is also cleared.
//
// return a string or nil if failed.
func (z *StringBuilder) GetBefore(arg1 string, arg2 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkStringBuilder_getBefore(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the substring found between the 1st occurrence of beginMark and the next
// occurrence of endMark.
// return a string or nil if failed.
func (z *StringBuilder) GetBetween(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkStringBuilder_getBetween(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Decodes and returns the decoded bytes. The encoding can be set to any of the
// following strings: "base64", "hex", "quoted-printable" (or "qp"), "url",
// "base32", "Q", "B", "url_rc1738", "url_rfc2396", "url_rfc3986", "url_oauth",
// "uu", "modBase64", or "html" (for HTML entity encoding). The full up-to-date
// list of supported binary encodings is available at the link entitled "Supported
// Binary Encodings" below.
func (z *StringBuilder) GetDecoded(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkStringBuilder_GetDecoded(z.ckObj, cstr1, ckOutBytes)

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


// Returns the string contents encoded in an encoding such as base64, hex,
// quoted-printable, or URL-encoding. The encoding can be set to any of the following
// strings: "base64", "hex", "quoted-printable" (or "qp"), "url", "base32", "Q",
// "B", "url_rc1738", "url_rfc2396", "url_rfc3986", "url_oauth", "uu", "modBase64",
// or "html" (for HTML entity encoding). The full up-to-date list of supported
// binary encodings is available at the link entitled "Supported Binary Encodings"
// below.
// 
// Note: The Encode method modifies the content of this object. The GetEncoded
// method leaves this object's content unmodified.
//
// return a string or nil if failed.
func (z *StringBuilder) GetEncoded(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkStringBuilder_getEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth substring in string that is a list delimted by delimiterChar. The first
// substring is at index 0. If exceptDoubleQuoted is true, then the delimiter char found
// between double quotes is not treated as a delimiter. If exceptEscaped is true, then an
// escaped (with a backslash) delimiter char is not treated as a delimiter.
// return a string or nil if failed.
func (z *StringBuilder) GetNth(arg1 int, arg2 string, arg3 bool, arg4 bool) *string {
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    retStr := C.CkStringBuilder_getNth(z.ckObj, C.int(arg1), cstr2, b3, b4)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the last N lines of the text. If fewer than numLines lines exists, then all
// of the text is returned. If bCrlf is true, then the line endings of the
// returned string are converted to CRLF, otherwise the line endings are converted
// to LF-only.
// return a string or nil if failed.
func (z *StringBuilder) LastNLines(arg1 int, arg2 bool) *string {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkStringBuilder_lastNLines(z.ckObj, C.int(arg1), b2)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads the contents of a file.
func (z *StringBuilder) LoadFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_LoadFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Obfuscates the string. (The Unobfuscate method can be called to reverse the
// obfuscation to restore the original string.)
// 
// The Chilkat string obfuscation algorithm works by taking the utf-8 bytes of the
// string, base64 encoding it, and then scrambling the letters of the base64
// encoded string. It is deterministic in that the same string will always
// obfuscate to the same result. It is NOT a secure way of encrypting a string. It
// is only meant to be a simple means of transforming a string into something
// unintelligible.
//
func (z *StringBuilder) Obfuscate()  {

    C.CkStringBuilder_Obfuscate(z.ckObj)



}


// Prepends a copy of the specified string to this instance.
func (z *StringBuilder) Prepend(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringBuilder_Prepend(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// In-place decodes the string from punycode.
func (z *StringBuilder) PunyDecode() bool {

    v := C.CkStringBuilder_PunyDecode(z.ckObj)


    return v != 0
}


// In-place encodes the string to punycode.
func (z *StringBuilder) PunyEncode() bool {

    v := C.CkStringBuilder_PunyEncode(z.ckObj)


    return v != 0
}


// Removes the substring found after the final occurrence of the marker. Also removes
// the marker. Returns true if the marker was found and content was removed.
// Otherwise returns false.
func (z *StringBuilder) RemoveAfterFinal(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringBuilder_RemoveAfterFinal(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Removes the substring found before the 1st occurrence of the marker. Also removes
// the marker. Returns true if the marker was found and content was removed.
// Otherwise returns false.
func (z *StringBuilder) RemoveBefore(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringBuilder_RemoveBefore(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Replaces all occurrences of a specified string in this instance with another
// specified string. Returns the number of replacements.
func (z *StringBuilder) Replace(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_Replace(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Replaces the content found after the final occurrence of marker with replacement.
func (z *StringBuilder) ReplaceAfterFinal(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_ReplaceAfterFinal(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Replaces the first occurrence of ALL the content found between beginMark and endMark
// with replacement. The beginMark and endMark are included in what is replaced if replaceMarks is true.
func (z *StringBuilder) ReplaceAllBetween(arg1 string, arg2 string, arg3 string, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkStringBuilder_ReplaceAllBetween(z.ckObj, cstr1, cstr2, cstr3, b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Replaces all occurrences of value with replacement, but only where value is found between
// beginMark and endMark. Returns the number of replacements made.
func (z *StringBuilder) ReplaceBetween(arg1 string, arg2 string, arg3 string, arg4 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkStringBuilder_ReplaceBetween(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return int(v)
}


// Replaces the first occurrence of a specified string in this instance with
// another string. Returns true if the value was found and replaced. Otherwise
// returns false.
func (z *StringBuilder) ReplaceFirst(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_ReplaceFirst(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Replaces all occurrences of value with the decimal integer replacement. Returns the
// number of replacements.
func (z *StringBuilder) ReplaceI(arg1 string, arg2 int) int {
    cstr1 := C.CString(arg1)

    v := C.CkStringBuilder_ReplaceI(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Replaces all occurrences of value with replacement (case insensitive). Returns the
// number of replacements.
func (z *StringBuilder) ReplaceNoCase(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_ReplaceNoCase(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Replaces all word occurrences of a specified string in this instance with
// another specified string. Returns the number of replacements made.
// 
// Important: This method is limited to replacing whole words in strings that only
// contains characters in the Latin1 charset (i.e. iso-8859-1 or Windows-1252). A
// whole word can only contain alphanumeric chars where the alpha chars are
// restricted to those of the Latin1 alpha chars. (The underscore character is also
// considered part of a word.)
//
func (z *StringBuilder) ReplaceWord(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkStringBuilder_ReplaceWord(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Removes all characters from the current StringBuilder instance, and write zero
// bytes to the allocated memory before deallocating.
func (z *StringBuilder) SecureClear()  {

    C.CkStringBuilder_SecureClear(z.ckObj)



}


// Sets the Nth substring in string in a list delimted by delimiterChar. The first substring
// is at index 0. If exceptDoubleQuoted is true, then the delimiter char found between double
// quotes is not treated as a delimiter. If exceptEscaped is true, then an escaped (with a
// backslash) delimiter char is not treated as a delimiter.
func (z *StringBuilder) SetNth(arg1 int, arg2 string, arg3 string, arg4 bool, arg5 bool) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    b5 := C.int(0)
    if arg5 {
        b5 = C.int(1)
    }

    v := C.CkStringBuilder_SetNth(z.ckObj, C.int(arg1), cstr2, cstr3, b4, b5)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Sets this instance to a copy of the specified string.
func (z *StringBuilder) SetString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkStringBuilder_SetString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if the string starts with substr. Otherwise returns false. The
// comparison is case sensitive if caseSensitive is true, and case insensitive if caseSensitive is
// false.
func (z *StringBuilder) StartsWith(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkStringBuilder_StartsWith(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Converts line endings to CRLF (Windows) format.
func (z *StringBuilder) ToCRLF() bool {

    v := C.CkStringBuilder_ToCRLF(z.ckObj)


    return v != 0
}


// Converts line endings to LF-only (UNIX) format.
func (z *StringBuilder) ToLF() bool {

    v := C.CkStringBuilder_ToLF(z.ckObj)


    return v != 0
}


// Converts the contents to lowercase.
func (z *StringBuilder) ToLowercase() bool {

    v := C.CkStringBuilder_ToLowercase(z.ckObj)


    return v != 0
}


// Converts the contents to uppercase.
func (z *StringBuilder) ToUppercase() bool {

    v := C.CkStringBuilder_ToUppercase(z.ckObj)


    return v != 0
}


// Trims whitespace from both ends of the string.
func (z *StringBuilder) Trim() bool {

    v := C.CkStringBuilder_Trim(z.ckObj)


    return v != 0
}


// Replaces all tabs, CR's, and LF's, with SPACE chars, and removes extra SPACE's
// so there are no occurances of more than one SPACE char in a row.
func (z *StringBuilder) TrimInsideSpaces() bool {

    v := C.CkStringBuilder_TrimInsideSpaces(z.ckObj)


    return v != 0
}


// Unobfuscates the string.
// 
// The Chilkat string obfuscation algorithm works by taking the utf-8 bytes of the
// string, base64 encoding it, and then scrambling the letters of the base64
// encoded string. It is deterministic in that the same string will always
// obfuscate to the same result. It is not a secure way of encrypting a string. It
// is only meant to be a simple means of transforming a string into something
// unintelligible.
//
func (z *StringBuilder) Unobfuscate()  {

    C.CkStringBuilder_Unobfuscate(z.ckObj)



}


// Writes the contents to a file. If emitBom is true, then the BOM (also known as a
// preamble), is emitted for charsets that define a BOM (such as utf-8, utf-16,
// utf-32, etc.)
func (z *StringBuilder) WriteFile(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkStringBuilder_WriteFile(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Writes the contents to a file, but only if it is a new file or if the contents
// are different than the existing file. If emitBom is true, then the BOM (also
// known as a preamble), is emitted for charsets that define a BOM (such as utf-8,
// utf-16, utf-32, etc.)
func (z *StringBuilder) WriteFileIfModified(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkStringBuilder_WriteFileIfModified(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


