// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkAsn.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewAsn() *Asn {
	return &Asn{C.CkAsn_Create()}
}

func (z *Asn) DisposeAsn() {
    if z != nil {
	C.CkAsn_Dispose(z.ckObj)
	}
}




// property: BoolValue
// The ASN.1 item's boolean value if it is a boolean item.
func (z *Asn) BoolValue() bool {
    v := int(C.CkAsn_getBoolValue(z.ckObj))
    return v != 0
}
// property setter: BoolValue
func (z *Asn) SetBoolValue(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAsn_putBoolValue(z.ckObj,v)
}

// property: Constructed
// true if this ASN.1 item is a constructed item. Sequence and Set items are
// constructed and can contain sub-items. All other tags (boolean, integer, octets,
// utf8String, etc.) are primitive (non-constructed).
func (z *Asn) Constructed() bool {
    v := int(C.CkAsn_getConstructed(z.ckObj))
    return v != 0
}

// property: ContentStr
// The ASN.1 item's content if it is an ASN.1 string type (such as Utf8String,
// BmpString, PrintableString, VisibleString, T61String, IA5String, NumericString,
// or UniversalString).
func (z *Asn) ContentStr() string {
    return C.GoString(C.CkAsn_contentStr(z.ckObj))
}
// property setter: ContentStr
func (z *Asn) SetContentStr(goStr string) {
    cStr := C.CString(goStr)
    C.CkAsn_putContentStr(z.ckObj,cStr)
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
func (z *Asn) DebugLogFilePath() string {
    return C.GoString(C.CkAsn_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Asn) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkAsn_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IntValue
// The ASN.1 item's integer value if it is a small integer item.
func (z *Asn) IntValue() int {
    return int(C.CkAsn_getIntValue(z.ckObj))
}
// property setter: IntValue
func (z *Asn) SetIntValue(value int) {
    C.CkAsn_putIntValue(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Asn) LastErrorHtml() string {
    return C.GoString(C.CkAsn_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Asn) LastErrorText() string {
    return C.GoString(C.CkAsn_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Asn) LastErrorXml() string {
    return C.GoString(C.CkAsn_lastErrorXml(z.ckObj))
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
func (z *Asn) LastMethodSuccess() bool {
    v := int(C.CkAsn_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Asn) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAsn_putLastMethodSuccess(z.ckObj,v)
}

// property: NumSubItems
// The number of sub-items contained within this ASN.1 item. Only constructed
// items, such as Sequence and Set will contain sub-iitems. Primitive items such as
// OIDs, octet strings, integers, etc. will never contain sub-items.
func (z *Asn) NumSubItems() int {
    return int(C.CkAsn_getNumSubItems(z.ckObj))
}

// property: Tag
// The ASN.1 item's tag as a descriptive string. Possible values are:
// boolean
// integer
// bitString
// octets
// null
// oid
// utf8String
// relativeOid
// sequence
// set
// numericString
// printableString
// t61String
// ia5String
// utcTime
// bmpString
func (z *Asn) Tag() string {
    return C.GoString(C.CkAsn_tag(z.ckObj))
}

// property: TagValue
// The ASN.1 item's tag as a integer value. The integer values for possible tags
// are as follows:
// boolean (1)
// integer (2)
// bitString (3)
// octets (4)
// null (5)
// oid (6)
// utf8String (12)
// relativeOid (13)
// sequence (16)
// set (17)
// numericString (18)
// printableString (19)
// t61String (20)
// ia5String (22)
// utcTime (23)
// bmpString (30)
func (z *Asn) TagValue() int {
    return int(C.CkAsn_getTagValue(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Asn) VerboseLogging() bool {
    v := int(C.CkAsn_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Asn) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAsn_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Asn) Version() string {
    return C.GoString(C.CkAsn_version(z.ckObj))
}
// Appends an ASN.1 integer, but one that is a big (huge) integer that is too large
// to be represented by an integer variable. The bytes composing the integer are
// passed in encoded string format (such as base64, hex, etc.). The byte order must
// be big-endian. The encoding may be any of the following encodings: "Base64", "Hex",
// "Base58", "modBase64", "Base32", "UU", "QP" (for quoted-printable), "URL" (for
// url-encoding), "Q", "B", "url_oath", "url_rfc1738", "url_rfc2396", and
// "url_rfc3986". The encoding name is case insensitive (for example, both "Base64" and
// "base64" are treated the same).
func (z *Asn) AppendBigInt(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAsn_AppendBigInt(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends an ASN.1 bit string to the caller's sub-items. The bytes containing the
// bits are passed in encoded string format (such as base64, hex, etc.). The byte
// order must be big-endian (MSB first). The encoding may be any of the following
// encodings: "Base64", "Hex", "Base58", "modBase64", "Base32", "UU", "QP" (for
// quoted-printable), "URL" (for url-encoding), "Q", "B", "url_oath",
// "url_rfc1738", "url_rfc2396", and "url_rfc3986". The encoding name is case
// insensitive (for example, both "Base64" and "base64" are treated the same).
func (z *Asn) AppendBits(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAsn_AppendBits(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends an ASN.1 boolean item to the caller's sub-items. Items may only be
// appended to constructed data types such as Sequence and Set.
func (z *Asn) AppendBool(arg1 bool) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkAsn_AppendBool(z.ckObj, b1)


    return v != 0
}


// Appends an ASN.1 context-specific constructed item to the caller's sub-items.
func (z *Asn) AppendContextConstructed(arg1 int) bool {

    v := C.CkAsn_AppendContextConstructed(z.ckObj, C.int(arg1))


    return v != 0
}


// Appends an ASN.1 context-specific primitive item to the caller's sub-items. The
// bytes are passed in encoded string format (such as base64, hex, etc.). The encoding
// may be any of the following encodings: "Base64", "Hex", "Base58", "modBase64",
// "Base32", "UU", "QP" (for quoted-printable), "URL" (for url-encoding), "Q", "B",
// "url_oath", "url_rfc1738", "url_rfc2396", and "url_rfc3986". The encoding name is
// case insensitive (for example, both "Base64" and "base64" are treated the same).
func (z *Asn) AppendContextPrimitive(arg1 int, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkAsn_AppendContextPrimitive(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Appends an ASN.1 integer item to the caller's sub-items. Items may only be
// appended to constructed data types such as Sequence and Set.
func (z *Asn) AppendInt(arg1 int) bool {

    v := C.CkAsn_AppendInt(z.ckObj, C.int(arg1))


    return v != 0
}


// Appends an ASN.1 null item to the caller's sub-items. Items may only be appended
// to constructed data types such as Sequence and Set.
func (z *Asn) AppendNull() bool {

    v := C.CkAsn_AppendNull(z.ckObj)


    return v != 0
}


// Appends an ASN.1 octet string to the caller's sub-items. The bytes are passed in
// encoded string format (such as base64, hex, etc.). The encoding may be any of the
// following encodings: "Base64", "Hex", "Base58", "modBase64", "Base32", "UU",
// "QP" (for quoted-printable), "URL" (for url-encoding), "Q", "B", "url_oath",
// "url_rfc1738", "url_rfc2396", and "url_rfc3986". The encoding name is case
// insensitive (for example, both "Base64" and "base64" are treated the same).
func (z *Asn) AppendOctets(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAsn_AppendOctets(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends an ASN.1 OID (object identifier) to the caller's sub-items. The OID is
// passed in string form, such as "1.2.840.113549.1.9.1".
func (z *Asn) AppendOid(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkAsn_AppendOid(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends an ASN.1 sequence item to the caller's sub-items.
func (z *Asn) AppendSequence() bool {

    v := C.CkAsn_AppendSequence(z.ckObj)


    return v != 0
}


// Appends an ASN.1 sequence item to the caller's sub-items, and updates the
// internal reference to point to the newly appended sequence item.
func (z *Asn) AppendSequence2() bool {

    v := C.CkAsn_AppendSequence2(z.ckObj)


    return v != 0
}


// Appends an ASN.1 sequence item to the caller's sub-items, and returns the newly
// appended sequence item.
func (z *Asn) AppendSequenceR() *Asn {

    retObj := C.CkAsn_AppendSequenceR(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Asn{retObj}
}


// Appends an ASN.1 set item to the caller's sub-items.
func (z *Asn) AppendSet() bool {

    v := C.CkAsn_AppendSet(z.ckObj)


    return v != 0
}


// Appends an ASN.1 set item to the caller's sub-items, and updates the internal
// reference to point to the newly appended set item.
func (z *Asn) AppendSet2() bool {

    v := C.CkAsn_AppendSet2(z.ckObj)


    return v != 0
}


// Appends an ASN.1 set item to the caller's sub-items, and returns the newly
// appended set item.
func (z *Asn) AppendSetR() *Asn {

    retObj := C.CkAsn_AppendSetR(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Asn{retObj}
}


// Appends a string item to the caller's sub-items. The strType specifies the type of
// string to be added. It may be "utf8", "ia5", "t61", "printable", "visible",
// "numeric", "universal", or "bmp". The value must conform to the ASN.1
// restrictions imposed for a given string type. The "utf8", "bmp", and "universal"
// types have no restrictions on what characters are allowed. In general, unless a
// specific type of string is required, choose the "utf8" type.
func (z *Asn) AppendString(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAsn_AppendString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends a UTCTime item to the caller's sub-items. The timeFormat specifies the format
// of the dateTimeStr. It should be set to "utc". (In the future, this method will be
// expanded to append GeneralizedTime items by using "generalized" for timeFormat.) To
// append the current date/time, set dateTimeStr equal to the empty string or the keyword
// "now". Otherwise, the dateTimeStr should be in the UTC time format "YYMMDDhhmm[ss]Z" or
// "YYMMDDhhmm[ss](+|-)hhmm".
func (z *Asn) AppendTime(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAsn_AppendTime(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Converts ASN.1 to XML and returns the XML string.
// return a string or nil if failed.
func (z *Asn) AsnToXml() *string {

    retStr := C.CkAsn_asnToXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Discards the Nth sub-item. (The 1st sub-item is at index 0.)
func (z *Asn) DeleteSubItem(arg1 int) bool {

    v := C.CkAsn_DeleteSubItem(z.ckObj, C.int(arg1))


    return v != 0
}


// Returns the ASN.1 in binary DER form.
func (z *Asn) GetBinaryDer() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkAsn_GetBinaryDer(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the content of the ASN.1 item in encoded string form. The encoding may be
// any of the following encodings: "Base64", "Hex", "Base58", "modBase64",
// "Base32", "UU", "QP" (for quoted-printable), "URL" (for url-encoding), "Q", "B",
// "url_oath", "url_rfc1738", "url_rfc2396", and "url_rfc3986". The encoding name is
// case insensitive (for example, both "Base64" and "base64" are treated the same).
// return a string or nil if failed.
func (z *Asn) GetEncodedContent(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkAsn_getEncodedContent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the binary DER in encoded string form. The encoding indicates the encoding
// and can be "base64", "hex", "uu", "quoted-printable", "base32", or "modbase64".
// return a string or nil if failed.
func (z *Asn) GetEncodedDer(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkAsn_getEncodedDer(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the last ASN.1 sub-item. This method can be called immediately after any
// Append* method to access the appended item.
func (z *Asn) GetLastSubItem() *Asn {

    retObj := C.CkAsn_GetLastSubItem(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Asn{retObj}
}


// Returns the Nth ASN.1 sub-item. The 1st sub-item is at index 0.
func (z *Asn) GetSubItem(arg1 int) *Asn {

    retObj := C.CkAsn_GetSubItem(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Asn{retObj}
}


// Loads ASN.1 from the XML representation (such as that created by the AsnToXml
// method).
func (z *Asn) LoadAsnXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkAsn_LoadAsnXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads ASN.1 from the binary DER contained in bd.
func (z *Asn) LoadBd(arg1 *BinData) bool {

    v := C.CkAsn_LoadBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads ASN.1 from binary DER.
func (z *Asn) LoadBinary(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkAsn_LoadBinary(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Loads ASN.1 from a binary DER file.
func (z *Asn) LoadBinaryFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkAsn_LoadBinaryFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads ASN.1 from an encoded string. The encoding can be "base64", "hex", "uu",
// "quoted-printable", "base32", or "modbase64".
func (z *Asn) LoadEncoded(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAsn_LoadEncoded(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the content of this primitive ASN.1 item. The encoding may be any of the
// following encodings: "Base64", "Hex", "Base58", "modBase64", "Base32", "UU",
// "QP" (for quoted-printable), "URL" (for url-encoding), "Q", "B", "url_oath",
// "url_rfc1738", "url_rfc2396", and "url_rfc3986". The encoding name is case
// insensitive (for example, both "Base64" and "base64" are treated the same).
func (z *Asn) SetEncodedContent(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAsn_SetEncodedContent(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Appends the ASN.1 in binary DER format to bd.
func (z *Asn) WriteBd(arg1 *BinData) bool {

    v := C.CkAsn_WriteBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Writes the ASN.1 in binary DER form to a file.
func (z *Asn) WriteBinaryDer(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkAsn_WriteBinaryDer(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


