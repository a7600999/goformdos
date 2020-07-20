// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkXmp.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewXmp() *Xmp {
	return &Xmp{C.CkXmp_Create()}
}

func (z *Xmp) DisposeXmp() {
    if z != nil {
	C.CkXmp_Dispose(z.ckObj)
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
func (z *Xmp) DebugLogFilePath() string {
    return C.GoString(C.CkXmp_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Xmp) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmp_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Xmp) LastErrorHtml() string {
    return C.GoString(C.CkXmp_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Xmp) LastErrorText() string {
    return C.GoString(C.CkXmp_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Xmp) LastErrorXml() string {
    return C.GoString(C.CkXmp_lastErrorXml(z.ckObj))
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
func (z *Xmp) LastMethodSuccess() bool {
    v := int(C.CkXmp_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Xmp) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmp_putLastMethodSuccess(z.ckObj,v)
}

// property: NumEmbedded
// The number of XMP metadata documents found within the JPG or TIFF file loaded by
// LoadAppFile.
func (z *Xmp) NumEmbedded() int {
    return int(C.CkXmp_getNumEmbedded(z.ckObj))
}

// property: StructInnerDescrip
// Determines whether structures are stored with rdf:parseType="Resource", or
// within an "rdf:Description" sub-node.
func (z *Xmp) StructInnerDescrip() bool {
    v := int(C.CkXmp_getStructInnerDescrip(z.ckObj))
    return v != 0
}
// property setter: StructInnerDescrip
func (z *Xmp) SetStructInnerDescrip(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmp_putStructInnerDescrip(z.ckObj,v)
}

// property: UncommonOptions
// A property to be used for future uncommon needs. The default value is the empty
// string.
func (z *Xmp) UncommonOptions() string {
    return C.GoString(C.CkXmp_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Xmp) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkXmp_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Xmp) VerboseLogging() bool {
    v := int(C.CkXmp_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Xmp) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXmp_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Xmp) Version() string {
    return C.GoString(C.CkXmp_version(z.ckObj))
}
// Adds or replaces an XMP property array. The XMP metadata to be updated is
// contained in the XML object passed in the 1st argument. The 2nd argument
// specifies the array type, which can be "bag", "seq", or "alt". The property name
// should be prefixed with the namespace, such as "dc:subject".
func (z *Xmp) AddArray(arg1 *Xml, arg2 string, arg3 string, arg4 *StringArray) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkXmp_AddArray(z.ckObj, arg1.ckObj, cstr2, cstr3, arg4.ckObj)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds a namespace to URI mapping. When a property is added via AddSimpleString or
// any of the other methods, the property name is namespace qualified. When adding
// the first property in a namespace, the rdf:Description is automatically added
// and the URI is obtained from the namespace-to-URI mappings. The standard (and
// commonly used) namespace mappings are defined by default. This is only used if
// the namespace is custom or not already handled.
func (z *Xmp) AddNsMapping(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkXmp_AddNsMapping(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Adds or updates an XMP integer property. The XMP metadata to be updated is
// contained in the XML object passed in the 1st argument. The property name should
// be prefixed with the namespace, such as "tiff:XResolution".
func (z *Xmp) AddSimpleInt(arg1 *Xml, arg2 string, arg3 int) bool {
    cstr2 := C.CString(arg2)

    v := C.CkXmp_AddSimpleInt(z.ckObj, arg1.ckObj, cstr2, C.int(arg3))

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds or updates a simple XMP string property. The XMP metadata to be updated is
// contained in the XML object passed in the 1st argument. The property name should
// be prefixed with the namespace, such as "photoshop:Credit".
func (z *Xmp) AddSimpleStr(arg1 *Xml, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkXmp_AddSimpleStr(z.ckObj, arg1.ckObj, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds or updates an XMP structured property value. The XMP metadata to be updated
// is contained in the XML object passed in the 1st argument. The structure name
// should be prefixed with the namespace, such as
// "Iptc4xmpCore:CreatorContactInfo". The property name within the structure should
// also be prefixed with the namespace, such as "Iptc4xmpCore:CiAdrCity".
func (z *Xmp) AddStructProp(arg1 *Xml, arg2 string, arg3 string, arg4 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkXmp_AddStructProp(z.ckObj, arg1.ckObj, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Appends a new XMP metadata file to the XMP object. Any XMPs appended via this
// method will be present in the file when SaveAppFile is called. Files containing
// XMP metadata typically only include a single XMP document, so this method is
// usually only called when adding XMP metadata to a file for the first time.
func (z *Xmp) Append(arg1 *Xml) bool {

    v := C.CkXmp_Append(z.ckObj, arg1.ckObj)


    return v != 0
}


// Finds and returns an XMP array property. The property name should be prefixed
// with the namespace, such as "dc:subject".
func (z *Xmp) GetArray(arg1 *Xml, arg2 string) *StringArray {
    cstr2 := C.CString(arg2)

    retObj := C.CkXmp_GetArray(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Returns the Nth embedded XMP document as a Chilkat XML object.
func (z *Xmp) GetEmbedded(arg1 int) *Xml {

    retObj := C.CkXmp_GetEmbedded(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}



func (z *Xmp) GetProperty(arg1 *Xml, arg2 string) *Xml {
    cstr2 := C.CString(arg2)

    retObj := C.CkXmp_GetProperty(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Finds and returns an XMP integer property. The property name should be prefixed
// with the namespace, such as "tiff:ResolutionUnit".
func (z *Xmp) GetSimpleInt(arg1 *Xml, arg2 string) int {
    cstr2 := C.CString(arg2)

    v := C.CkXmp_GetSimpleInt(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Finds and returns an XMP simple string property. The property name should be
// prefixed with the namespace, such as "photoshop:Source".
// return a string or nil if failed.
func (z *Xmp) GetSimpleStr(arg1 *Xml, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkXmp_getSimpleStr(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the property names used by an exsting structure within an XMP document.
// The contents of the structure can be retrieved by calling GetStructValue for
// each property name returned by GetStructPropNames.
func (z *Xmp) GetStructPropNames(arg1 *Xml, arg2 string) *StringArray {
    cstr2 := C.CString(arg2)

    retObj := C.CkXmp_GetStructPropNames(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Returns the value of a single item within an XMP structure property. Property
// names should always be prefixed with the namespace.
// return a string or nil if failed.
func (z *Xmp) GetStructValue(arg1 *Xml, arg2 string, arg3 string) *string {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkXmp_getStructValue(z.ckObj, arg1.ckObj, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads a TIFF or JPG file into the XMP object.
func (z *Xmp) LoadAppFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXmp_LoadAppFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads a JPG or TIFF from an byte buffer containing the image file data.
func (z *Xmp) LoadFromBuffer(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkXmp_LoadFromBuffer(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Creates and returns a new/empty XMP metadata document as a Chilkat XML object.
func (z *Xmp) NewXmp() *Xml {

    retObj := C.CkXmp_NewXmp(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Removes all XMP metadata documents from an XMP object. After calling this
// method, call SaveAppFile to rewrite the JPG or TIFF file with the XMP metadata
// removed.
func (z *Xmp) RemoveAllEmbedded() bool {

    v := C.CkXmp_RemoveAllEmbedded(z.ckObj)


    return v != 0
}


// Removes an XMP array property from the XMP document.
func (z *Xmp) RemoveArray(arg1 *Xml, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkXmp_RemoveArray(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Removes a namespace-to-URI mapping.
func (z *Xmp) RemoveNsMapping(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkXmp_RemoveNsMapping(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Removes a simple XMP property from the XMP document.
func (z *Xmp) RemoveSimple(arg1 *Xml, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkXmp_RemoveSimple(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Removes an XMP structure property from the XMP document.
func (z *Xmp) RemoveStruct(arg1 *Xml, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkXmp_RemoveStruct(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Removes a single member from an XMP structured property.
func (z *Xmp) RemoveStructProp(arg1 *Xml, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkXmp_RemoveStructProp(z.ckObj, arg1.ckObj, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Persists all changes made to the XMP document(s) by saving the XMP object to a
// file. Changes made by adding, updating, or removing properties are not persisted
// to the filesystem until this is called.
func (z *Xmp) SaveAppFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXmp_SaveAppFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves a JPG or TIFF image with updated XMP to a byte buffer.
func (z *Xmp) SaveToBuffer() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkXmp_SaveToBuffer(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Unlocks the XMP component at runtime. This must be called once at the beginning
// of your application. Passing an arbitrary value initiates a fully-functional
// 30-day trial. A purchased unlock code is required to use the component beyond 30
// days.
func (z *Xmp) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXmp_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


