// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkXml.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewXml() *Xml {
	return &Xml{C.CkXml_Create()}
}

func (z *Xml) DisposeXml() {
    if z != nil {
	C.CkXml_Dispose(z.ckObj)
	}
}




// property: Cdata
// When True, causes an XML node's content to be encapsulated in a CDATA section.
func (z *Xml) Cdata() bool {
    v := int(C.CkXml_getCdata(z.ckObj))
    return v != 0
}
// property setter: Cdata
func (z *Xml) SetCdata(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXml_putCdata(z.ckObj,v)
}

// property: Content
// The content of the XML node. It is the text between the open and close tags, not
// including child nodes. For example:
// _LT_tag1_GT_This is the content_LT_/tag1_GT_
// 
// _LT_tag2_GT__LT_child1_GT_abc_LT_/child1_GT__LT_child2_GT_abc_LT_/child2_GT_This is the content_LT_/tag2_GT_
// Because the child nodes are not included, the content of "tag1" and "tag2" are
// both equal to "This is the content".
func (z *Xml) Content() string {
    return C.GoString(C.CkXml_content(z.ckObj))
}
// property setter: Content
func (z *Xml) SetContent(goStr string) {
    cStr := C.CString(goStr)
    C.CkXml_putContent(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ContentInt
// Set/get the content as an integer.
func (z *Xml) ContentInt() int {
    return int(C.CkXml_getContentInt(z.ckObj))
}
// property setter: ContentInt
func (z *Xml) SetContentInt(value int) {
    C.CkXml_putContentInt(z.ckObj,C.int(value))
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
func (z *Xml) DebugLogFilePath() string {
    return C.GoString(C.CkXml_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Xml) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkXml_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DocType
// The DOCTYPE declaration (if any) for the XML document.
func (z *Xml) DocType() string {
    return C.GoString(C.CkXml_docType(z.ckObj))
}
// property setter: DocType
func (z *Xml) SetDocType(goStr string) {
    cStr := C.CString(goStr)
    C.CkXml_putDocType(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EmitBom
// If true, then emit the BOM (byte order mark, also known as a preamble) for
// encodings such as utf-8, utf-16, etc. The defautl value is false. This only
// applies when writing XML files. It does not apply when getting the XML as a
// string via the GetXml method.
func (z *Xml) EmitBom() bool {
    v := int(C.CkXml_getEmitBom(z.ckObj))
    return v != 0
}
// property setter: EmitBom
func (z *Xml) SetEmitBom(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXml_putEmitBom(z.ckObj,v)
}

// property: EmitCompact
// If true, then GetXml and GetXmlSb emit compact XML. The XML emitted has no
// unnecessary whitespace, incuding no end-of-lines (CR's and/or LF's). The default
// value is false, which maintains backward compatibility.
func (z *Xml) EmitCompact() bool {
    v := int(C.CkXml_getEmitCompact(z.ckObj))
    return v != 0
}
// property setter: EmitCompact
func (z *Xml) SetEmitCompact(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXml_putEmitCompact(z.ckObj,v)
}

// property: EmitXmlDecl
// If true, then the XML declaration is emitted for methods (such as GetXml or
// SaveXml) where the XML is written to a file or string. The default value of this
// property is true. If set to false, the XML declaration is not emitted. (The
// XML declaration is the 1st line of an XML document starting with "_LT_?xml ...".
func (z *Xml) EmitXmlDecl() bool {
    v := int(C.CkXml_getEmitXmlDecl(z.ckObj))
    return v != 0
}
// property setter: EmitXmlDecl
func (z *Xml) SetEmitXmlDecl(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXml_putEmitXmlDecl(z.ckObj,v)
}

// property: Encoding
// This is the encoding attribute in the XML declaration, such as "utf-8" or
// "iso-8859-1". The default is "utf-8". This property can be set from any node in
// the XML document and when set, causes the encoding property to be added to the
// XML declaration. Setting this property does not cause the document to be
// converted to a different encoding.
// 
// Calling any of the LoadXml* methods causes this property to be set to the
// charset found within the XML, if any. If no charset is specified within the
// loaded XML, then the LoadXml method resets this property to its default value of
// "utf-8".
//
func (z *Xml) Encoding() string {
    return C.GoString(C.CkXml_encoding(z.ckObj))
}
// property setter: Encoding
func (z *Xml) SetEncoding(goStr string) {
    cStr := C.CString(goStr)
    C.CkXml_putEncoding(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: I
// Used in tagPaths (and ChilkatPath). The value of this property is substituted
// for "i" in "[i]". See the example below..
func (z *Xml) I() int {
    return int(C.CkXml_getI(z.ckObj))
}
// property setter: I
func (z *Xml) SetI(value int) {
    C.CkXml_putI(z.ckObj,C.int(value))
}

// property: IsBase64
// Returns true if the content contains only those characters allowed in the
// base64 encoding. A base64 string is composed of characters 'A'..'Z', 'a'..'z',
// '0'..'9', '+', '/' and it is often padded at the end with up to two '=', to make
// the length a multiple of 4. Whitespace is ignored.
func (z *Xml) IsBase64() bool {
    v := int(C.CkXml_getIsBase64(z.ckObj))
    return v != 0
}

// property: J
// Used in tagPaths (and ChilkatPath). The value of this property is substituted
// for "j" in "[j]". See the example below..
func (z *Xml) J() int {
    return int(C.CkXml_getJ(z.ckObj))
}
// property setter: J
func (z *Xml) SetJ(value int) {
    C.CkXml_putJ(z.ckObj,C.int(value))
}

// property: K
// Used in tagPaths (and ChilkatPath). The value of this property is substituted
// for "k" in "[k]". See the example below..
func (z *Xml) K() int {
    return int(C.CkXml_getK(z.ckObj))
}
// property setter: K
func (z *Xml) SetK(value int) {
    C.CkXml_putK(z.ckObj,C.int(value))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Xml) LastErrorHtml() string {
    return C.GoString(C.CkXml_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Xml) LastErrorText() string {
    return C.GoString(C.CkXml_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Xml) LastErrorXml() string {
    return C.GoString(C.CkXml_lastErrorXml(z.ckObj))
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
func (z *Xml) LastMethodSuccess() bool {
    v := int(C.CkXml_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Xml) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXml_putLastMethodSuccess(z.ckObj,v)
}

// property: NumAttributes
// The number of attributes. For example, the following node has 2 attributes:
// _LT_tag attr1="value1" attr2="value2"> This is the content_LT_/tag>
func (z *Xml) NumAttributes() int {
    return int(C.CkXml_getNumAttributes(z.ckObj))
}

// property: NumChildren
// The number of direct child nodes contained under this XML node.
func (z *Xml) NumChildren() int {
    return int(C.CkXml_getNumChildren(z.ckObj))
}

// property: SortCaseInsensitive
// If true (or 1 for ActiveX), then all Sort* methods use case insensitive sorting.
func (z *Xml) SortCaseInsensitive() bool {
    v := int(C.CkXml_getSortCaseInsensitive(z.ckObj))
    return v != 0
}
// property setter: SortCaseInsensitive
func (z *Xml) SetSortCaseInsensitive(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXml_putSortCaseInsensitive(z.ckObj,v)
}

// property: Standalone
// This is the standalone attribute in the XML declaration. This property can be
// set from any node in the XML document. A value of true adds a standalone="yes"
// to the XML declaration:
// _LT_?xml ... standalone="yes">
func (z *Xml) Standalone() bool {
    v := int(C.CkXml_getStandalone(z.ckObj))
    return v != 0
}
// property setter: Standalone
func (z *Xml) SetStandalone(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXml_putStandalone(z.ckObj,v)
}

// property: Tag
// The XML node's tag, including the namespace prefix.
func (z *Xml) Tag() string {
    return C.GoString(C.CkXml_tag(z.ckObj))
}
// property setter: Tag
func (z *Xml) SetTag(goStr string) {
    cStr := C.CString(goStr)
    C.CkXml_putTag(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TagNsPrefix
// The node's namespace prefix, if present. For example, if the tag is
// "soapenv:Envelope", then this property will return "soapenv".
func (z *Xml) TagNsPrefix() string {
    return C.GoString(C.CkXml_tagNsPrefix(z.ckObj))
}
// property setter: TagNsPrefix
func (z *Xml) SetTagNsPrefix(goStr string) {
    cStr := C.CString(goStr)
    C.CkXml_putTagNsPrefix(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TagPath
// Returns the path to reach this element from the XML document root. If the caller
// is the document root, then the empty string is returned.
func (z *Xml) TagPath() string {
    return C.GoString(C.CkXml_tagPath(z.ckObj))
}

// property: TagUnprefixed
// The node's tag without the namespace prefix. For example, if the tag is
// "soapenv:Envelope", then this property will return "Envelope".
func (z *Xml) TagUnprefixed() string {
    return C.GoString(C.CkXml_tagUnprefixed(z.ckObj))
}
// property setter: TagUnprefixed
func (z *Xml) SetTagUnprefixed(goStr string) {
    cStr := C.CString(goStr)
    C.CkXml_putTagUnprefixed(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TreeId
// Each tree (or XML document) has a unique TreeId. This is the ID of the tree, and
// can be used to determine if two Xml objects belong to the same tree.
func (z *Xml) TreeId() int {
    return int(C.CkXml_getTreeId(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Xml) VerboseLogging() bool {
    v := int(C.CkXml_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Xml) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkXml_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Xml) Version() string {
    return C.GoString(C.CkXml_version(z.ckObj))
}
// Accumulates the content of all nodes having a specific tag into a single result
// string. SkipTags specifies a set of subtrees to be avoided. The skipTags are
// formatted as a string of tags delimited by vertical bar characters. All nodes in
// sub-trees rooted with a tag appearing in skipTags are not included in the
// result.
// return a string or nil if failed.
func (z *Xml) AccumulateTagContent(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkXml_accumulateTagContent(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Adds an attribute to the calling node in the XML document. Returns True for
// success, and False for failure.
func (z *Xml) AddAttribute(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXml_AddAttribute(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds an integer attribute to a node.
func (z *Xml) AddAttributeInt(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_AddAttributeInt(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds an entire subtree as a child. If the child was a subtree within another Xml
// document then the subtree is effectively transferred from one XML document to
// another.
func (z *Xml) AddChildTree(arg1 *Xml) bool {

    v := C.CkXml_AddChildTree(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds an attribute to an XML node. If an attribute having the specified name
// already exists, the value is updated.
func (z *Xml) AddOrUpdateAttribute(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkXml_AddOrUpdateAttribute(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Adds an integer attribute to an XML node. If an attribute having the specified
// name already exists, the value is updated.
func (z *Xml) AddOrUpdateAttributeI(arg1 string, arg2 int)  {
    cstr1 := C.CString(arg1)

    C.CkXml_AddOrUpdateAttributeI(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Adds a style sheet declaration to the XML document. The styleSheet should be a string
// such as:
// _LT_?xml-stylesheet href="mystyle.css" title="Compact" type="text/css"?>
func (z *Xml) AddStyleSheet(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkXml_AddStyleSheet(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Adds an integer amount to an integer attribute's value. If the attribute does
// not yet exist, this method behaves the same as AddOrUpdateAttributeI.
func (z *Xml) AddToAttribute(arg1 string, arg2 int)  {
    cstr1 := C.CString(arg1)

    C.CkXml_AddToAttribute(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Adds an integer value to the content of a child node.
func (z *Xml) AddToChildContent(arg1 string, arg2 int)  {
    cstr1 := C.CString(arg1)

    C.CkXml_AddToChildContent(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Adds an integer amount to the node's content.
func (z *Xml) AddToContent(arg1 int)  {

    C.CkXml_AddToContent(z.ckObj, C.int(arg1))



}


// Appends text to the content of an XML node
func (z *Xml) AppendToContent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_AppendToContent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the node's content with 8bit data that is in a specified multibyte
// character encoding such as utf-8, shift-jis, big5, etc. The data is first
// B-encoded and the content is set to be the B-encoded string. For example, if
// called with "Big5"for the charset, you would get a string that looks something
// like this: "=?Big5?B?pHCtsw==?=". The data is Base64-encoded and stored between
// the last pair of "?" delimiters. Use the DecodeContent method to retrieve the
// byte data from a B encoded string.
func (z *Xml) BEncodeContent(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkXml_BEncodeContent(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Return true if a child at the specified tagPath contains content that matches a
// wildcarded pattern. Otherwise returns false.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) ChildContentMatches(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkXml_ChildContentMatches(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Follows a series of commands to navigate through an XML document to return a
// piece of data or update the caller's reference to a new XML document node.
// 
// Note: This method not related to the XPath (XML Path) standard in any way.
// 
// The pathCmd is formatted as a series of commands separated by vertical bar
// characters, and terminated with a return-command:
//     command|command|command|...|returnCommand
// 
// A command can be any of the following:
//     TagName -- Navigate to the 1st direct child with the given tag.
//     TagName[n] -- Navigate to the Nth direct child with the given tag.
//     ".." -- Navigate up to the parent.
//     "++" -- Navigate to the next sibling. (next/previous sibling feature added
//     in v9.5.0.76)
//     "--" -- Navigate to the previous sibling.
//     TagName{Content} -- Navigate to the 1st direct child with the given tag
//     having the exact content.
//     /T/TagName -- Traverse the XML DOM tree (rooted at the caller) and navigate
//     to the 1st node having the given tag.
//     /C/TagName,ContentPattern -- Traverse the XML DOM tree (rooted at the
//     caller) and navigate to the 1st node having the given tag with content that
//     matches the ContentPattern. ContentPattern may use one or more asterisk ('*")
//     characters to represent 0 or more of any character.
//     /C/ContentPattern -- Traverse the XML DOM tree (rooted at the caller) and
//     navigate to the 1st node having any tag with content that matches the
//     ContentPattern. ContentPattern may use one or more asterisk ('*") characters to
//     represent 0 or more of any character.
//     /A/TagName,AttrName,AttrValuePattern -- Traverse the XML DOM tree (rooted at
//     the caller) and navigate to the 1st node having the given tag, and attribute,
//     with the attribute value that matches the AttrValuePattern. AttrValuePattern may
//     use one or more asterisk ('*") characters to represent 0 or more of any
//     character.
// The returnCommand can be any of the following:
//     * -- Return the Content of the node.
//     (AttrName) -- Return the value of the given attribute.
//     $ -- Update the caller's internal reference to be the node (arrived at by
//     following the series of commands). Returns an empty string.
//
// return a string or nil if failed.
func (z *Xml) ChilkatPath(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkXml_chilkatPath(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Removes all children, attributes, and content from the XML node. Resets the tag
// name to "unnamed".
func (z *Xml) Clear()  {

    C.CkXml_Clear(z.ckObj)



}


// Return true if the node's content matches a wildcarded pattern.
func (z *Xml) ContentMatches(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkXml_ContentMatches(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Copies the tag, content, and attributes to the calling node.
func (z *Xml) Copy(arg1 *Xml)  {

    C.CkXml_Copy(z.ckObj, arg1.ckObj)



}


// Discards the caller's current internal reference and copies the internal
// reference from copyFromNode. Effectively updates the caller node to point to the same
// node in the XML document as copyFromNode.
func (z *Xml) CopyRef(arg1 *Xml)  {

    C.CkXml_CopyRef(z.ckObj, arg1.ckObj)



}


// Decodes a node's Q or B-encoded content string and returns the byte data.
func (z *Xml) DecodeContent() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkXml_DecodeContent(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Utility method to decode HTML entities. It accepts a string containing
// (potentially) HTML entities and returns a string with the entities decoded.
// return a string or nil if failed.
func (z *Xml) DecodeEntities(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkXml_decodeEntities(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Decrypts the content of an XML node that was previously 128-bit AES encrypted
// with the EncryptContent method.
func (z *Xml) DecryptContent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_DecryptContent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Encrypts the content of the calling XML node using 128-bit CBC AES encryption.
// The base64-encoded encrypted content replaces the original content.
func (z *Xml) EncryptContent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_EncryptContent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Removes and returns the Nth child of an XML node. The first child is at index 0.
func (z *Xml) ExtractChildByIndex(arg1 int) *Xml {

    retObj := C.CkXml_ExtractChildByIndex(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Removes and returns the first child node at the specified tag or tag path. The
// attrName and attrValue may be empty, in which case the first child matching the tag is
// removed and returned. If attrName is specified, then the first child having a tag
// equal to tagPath, and an attribute with attrName is returned. If attrValue is also
// specified, then only a child having a tag equal to tagPath, and an attribute named
// attrName, with a value equal to attrValue is returned.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) ExtractChildByName(arg1 string, arg2 string, arg3 string) *Xml {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkXml_ExtractChildByName(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Returns the child with the given tag or at the specified tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) FindChild(arg1 string) *Xml {
    cstr1 := C.CString(arg1)

    retObj := C.CkXml_FindChild(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Updates the Xml object's internal reference to point to a child at the specified
// tag or tagPath.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) FindChild2(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_FindChild2(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the next record node where the child with a specific tag matches a
// wildcarded pattern. This method makes it easy to iterate over high-level
// records.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) FindNextRecord(arg1 string, arg2 string) *Xml {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkXml_FindNextRecord(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// First checks for a child at tagPath, and if found, returns it. Otherwise creates a
// new child with empty content.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) FindOrAddNewChild(arg1 string) *Xml {
    cstr1 := C.CString(arg1)

    retObj := C.CkXml_FindOrAddNewChild(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Returns the first child. A program can step through the children by calling
// FirstChild, and then NextSibling repeatedly.
func (z *Xml) FirstChild() *Xml {

    retObj := C.CkXml_FirstChild(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Updates the internal reference of the caller to point to its first child.
func (z *Xml) FirstChild2() bool {

    v := C.CkXml_FirstChild2(z.ckObj)


    return v != 0
}


// Returns the name of the Nth attribute of an XML node. The first attribute is at
// index 0.
// return a string or nil if failed.
func (z *Xml) GetAttributeName(arg1 int) *string {

    retStr := C.CkXml_getAttributeName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of the Nth attribute of an XML node. The first attribute is at
// index 0.
// return a string or nil if failed.
func (z *Xml) GetAttributeValue(arg1 int) *string {

    retStr := C.CkXml_getAttributeValue(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns an attribute as an integer. Returns 0 if the attribute does not exist.
func (z *Xml) GetAttributeValueInt(arg1 int) int {

    v := C.CkXml_GetAttributeValueInt(z.ckObj, C.int(arg1))


    return int(v)
}


// Find and return the value of an attribute having a specified name.
// return a string or nil if failed.
func (z *Xml) GetAttrValue(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkXml_getAttrValue(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns an attribute as an integer. Returns 0 if the attribute does not exist.
func (z *Xml) GetAttrValueInt(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkXml_GetAttrValueInt(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns binary content of an XML node as a byte array. The content may have been
// Zip compressed, AES encrypted, or both. Unzip compression and AES decryption
// flags should be set appropriately.
func (z *Xml) GetBinaryContent(arg1 bool, arg2 bool, arg3 string) []byte {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkXml_GetBinaryContent(z.ckObj, b1, b2, cstr3, ckOutBytes)

    C.free(unsafe.Pointer(cstr3))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the Nth child of an XML node
func (z *Xml) GetChild(arg1 int) *Xml {

    retObj := C.CkXml_GetChild(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Updates the calling object's internal reference to the Nth child node.
func (z *Xml) GetChild2(arg1 int) bool {

    v := C.CkXml_GetChild2(z.ckObj, C.int(arg1))


    return v != 0
}


// Returns the content of a descendant child having a specified attribute. The tagPath
// can be a tag or a tag path.
// return a string or nil if failed.
func (z *Xml) GetChildAttrValue(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkXml_getChildAttrValue(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns false if the node's content is "0", otherwise returns true if the
// node contains a non-zero integer. The tagPath can be a tag or a tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "colors|primary|red".
//
func (z *Xml) GetChildBoolValue(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_GetChildBoolValue(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the content of a child having a specified tag. The tagPath can be a tag or
// a tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "colors|primary|red".
//
// return a string or nil if failed.
func (z *Xml) GetChildContent(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkXml_getChildContent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the content of the Nth child node.
// return a string or nil if failed.
func (z *Xml) GetChildContentByIndex(arg1 int) *string {

    retStr := C.CkXml_getChildContentByIndex(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the content of the XML element at the tagPath. The XML element's content is
// appended to the sb.
func (z *Xml) GetChildContentSb(arg1 string, arg2 *StringBuilder) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_GetChildContentSb(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the child having the exact tag and content.
func (z *Xml) GetChildExact(arg1 string, arg2 string) *Xml {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkXml_GetChildExact(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Returns the child integer content for a given tag. The tagPath can be a tag or a
// tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "colors|primary|red".
//
func (z *Xml) GetChildIntValue(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkXml_GetChildIntValue(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the tag name of the Nth child node.
// return a string or nil if failed.
func (z *Xml) GetChildTag(arg1 int) *string {

    retStr := C.CkXml_getChildTag(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the tag name of the Nth child node.
// return a string or nil if failed.
func (z *Xml) GetChildTagByIndex(arg1 int) *string {

    retStr := C.CkXml_getChildTagByIndex(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Finds and returns the XML child node having both a given tag and an attribute
// with a given name and value.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) GetChildWithAttr(arg1 string, arg2 string, arg3 string) *Xml {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkXml_GetChildWithAttr(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Returns the first child found having the exact content specified.
func (z *Xml) GetChildWithContent(arg1 string) *Xml {
    cstr1 := C.CString(arg1)

    retObj := C.CkXml_GetChildWithContent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Returns the child at the specified tag or tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) GetChildWithTag(arg1 string) *Xml {
    cstr1 := C.CString(arg1)

    retObj := C.CkXml_GetChildWithTag(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Returns the Nth child having a tag that matches exactly with the tagName. Use
// the NumChildrenHavingTag method to determine how many children have a particular
// tag.
func (z *Xml) GetNthChildWithTag(arg1 string, arg2 int) *Xml {
    cstr1 := C.CString(arg1)

    retObj := C.CkXml_GetNthChildWithTag(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Updates the calling object's internal reference to the Nth child node having a
// specific tag.
func (z *Xml) GetNthChildWithTag2(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_GetNthChildWithTag2(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the parent of this XML node, or NULL if the node is the root of the
// tree.
func (z *Xml) GetParent() *Xml {

    retObj := C.CkXml_GetParent(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Updates the internal reference of the caller to its parent.
func (z *Xml) GetParent2() bool {

    v := C.CkXml_GetParent2(z.ckObj)


    return v != 0
}


// Returns the root node of the XML document
func (z *Xml) GetRoot() *Xml {

    retObj := C.CkXml_GetRoot(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Updates the internal reference of the caller to the document root.
func (z *Xml) GetRoot2()  {

    C.CkXml_GetRoot2(z.ckObj)



}


// Returns a new XML object instance that references the same XML node.
func (z *Xml) GetSelf() *Xml {

    retObj := C.CkXml_GetSelf(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Generate the XML text document for the XML tree rooted at this node. If called
// from the root node of the XML document, then the XML declarator ("_LT_?xml
// version="1.0" encoding="utf-8" ?>") is included at the beginning of the XML.
// Otherwise, it is not included.
// return a string or nil if failed.
func (z *Xml) GetXml() *string {

    retStr := C.CkXml_getXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Emits (appends) the XML to the contents of bd.
func (z *Xml) GetXmlBd(arg1 *BinData) bool {

    v := C.CkXml_GetXmlBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Emits the XML to a StringBuilder object. (Appends to the existing contents of
// sb.)
func (z *Xml) GetXmlSb(arg1 *StringBuilder) bool {

    v := C.CkXml_GetXmlSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns true if the node contains an attribute with the specified name.
func (z *Xml) HasAttribute(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_HasAttribute(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if the node contains attribute with the name and value.
func (z *Xml) HasAttrWithValue(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXml_HasAttrWithValue(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns true if the node has a direct child node containing the exact content
// string specified.
func (z *Xml) HasChildWithContent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_HasChildWithContent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if the node has a child with the given tag (or tag path).
// Otherwise returns false.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) HasChildWithTag(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_HasChildWithTag(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if the node contains child with the given tag (or tag path) and
// content specified.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) HasChildWithTagAndContent(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXml_HasChildWithTagAndContent(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds an entire subtree as a child. If the child was a subtree within another Xml
// document then the subtree is effectively transferred from one XML document to
// another. The child tree is inserted in a position after the Nth child (of the
// calling node).
func (z *Xml) InsertChildTreeAfter(arg1 int, arg2 *Xml)  {

    C.CkXml_InsertChildTreeAfter(z.ckObj, C.int(arg1), arg2.ckObj)



}


// Adds an entire subtree as a child. If the child was a subtree within another Xml
// document then the subtree is effectively transferred from one XML document to
// another. The child tree is inserted in a position before the Nth child (of the
// calling node).
func (z *Xml) InsertChildTreeBefore(arg1 int, arg2 *Xml)  {

    C.CkXml_InsertChildTreeBefore(z.ckObj, C.int(arg1), arg2.ckObj)



}


// Returns the last Xml child node. A node's children can be enumerated by calling
// LastChild and then repeatedly calling PreviousSibling, until a NULL is returned.
func (z *Xml) LastChild() *Xml {

    retObj := C.CkXml_LastChild(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Updates the internal reference of the caller to its last child.
func (z *Xml) LastChild2() bool {

    v := C.CkXml_LastChild2(z.ckObj)


    return v != 0
}


// Loads XML from the contents of bd. If autoTrim is true, then each element's text
// content is trimmed of leading and trailing whitespace.
func (z *Xml) LoadBd(arg1 *BinData, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkXml_LoadBd(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Loads XML from the contents of a StringBuilder object.
func (z *Xml) LoadSb(arg1 *StringBuilder, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkXml_LoadSb(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Loads an XML document from a memory buffer and returns true if successful. The
// contents of the calling node are replaced with the root node of the XML document
// loaded.
func (z *Xml) LoadXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_LoadXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Same as LoadXml, but an additional argument controls whether or not
// leading/trailing whitespace is auto-trimmed from each leaf node's content.
func (z *Xml) LoadXml2(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkXml_LoadXml2(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads an XML document from a file and returns true if successful. The contents
// of the calling node are replaced with the root node of the XML document loaded.
func (z *Xml) LoadXmlFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_LoadXmlFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Same as LoadXmlFile, but an additional argument controls whether or not
// leading/trailing whitespace is auto-trimmed from each leaf node's content.
func (z *Xml) LoadXmlFile2(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkXml_LoadXmlFile2(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Creates a new child having tag and content. The new child is created even if a
// child with a tag equal to tagPath already exists. (Use FindOrAddNewChild to prevent
// creating children having the same tags.)
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "colors|primary|red". See the example below for details.
//
func (z *Xml) NewChild(arg1 string, arg2 string) *Xml {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkXml_NewChild(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Creates a new child node, but does not return the node that is created. The tagPath
// can be a tag or a tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "colors|primary|red". See the example below for details.
//
func (z *Xml) NewChild2(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkXml_NewChild2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Inserts a new child in a position after the Nth child node.
func (z *Xml) NewChildAfter(arg1 int, arg2 string, arg3 string) *Xml {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkXml_NewChildAfter(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Inserts a new child in a position before the Nth child node.
func (z *Xml) NewChildBefore(arg1 int, arg2 string, arg3 string) *Xml {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkXml_NewChildBefore(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Inserts a new child having an integer for content. The tagPath can be a tag or a
// tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "colors|primary|red". See the example below for details.
//
func (z *Xml) NewChildInt2(arg1 string, arg2 int)  {
    cstr1 := C.CString(arg1)

    C.CkXml_NewChildInt2(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Updates to Xml object's internal reference to the next node in a depth-first
// traversal. (This method name, NextInTraversal2, ends with "2" to signify that
// the internal reference is updated. There is no "NextInTraversal" method.)
// 
// The sbState contains the current state of the traversal. sbState should be empty when
// beginning a traversal.
//
func (z *Xml) NextInTraversal2(arg1 *StringBuilder) bool {

    v := C.CkXml_NextInTraversal2(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the nodes next sibling, or NULL if there are no more.
func (z *Xml) NextSibling() *Xml {

    retObj := C.CkXml_NextSibling(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Updates the internal reference of the caller to its next sibling.
func (z *Xml) NextSibling2() bool {

    v := C.CkXml_NextSibling2(z.ckObj)


    return v != 0
}


// Returns the number of children for the node indicated by tagPath. Returns -1 if the
// node at tagPath does not exist.
func (z *Xml) NumChildrenAt(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkXml_NumChildrenAt(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the number of children having a specific tag name.
func (z *Xml) NumChildrenHavingTag(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkXml_NumChildrenHavingTag(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the Xml object that is the node's previous sibling, or NULL if there are
// no more.
func (z *Xml) PreviousSibling() *Xml {

    retObj := C.CkXml_PreviousSibling(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Updates the internal reference of the caller to its previous sibling.
func (z *Xml) PreviousSibling2() bool {

    v := C.CkXml_PreviousSibling2(z.ckObj)


    return v != 0
}


// Recursively descends the XML from this node and removes all occurrences of the
// specified attribute. Returns the number of attribute occurrences removed.
func (z *Xml) PruneAttribute(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkXml_PruneAttribute(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Recursively descends the XML from this node and removes all occurrences of the
// specified tag, including all descendents of each removed node. Returns the
// number of tag occurrences removed.
func (z *Xml) PruneTag(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkXml_PruneTag(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Sets the node's content with 8bit data that is in a specified multibyte
// character encoding such as utf-8, shift-jis, big5, etc. The data is first
// Q-encoded and the content is set to be the Q-encoded string. For example, if
// called with "gb2312"for the charset, you would get a string that looks something
// like this: "=?gb2312?Q?=C5=B5=BB=F9?=". Character that are not 7bit are
// represented as "=XX" where XX is the hexidecimal value of the byte. Use the
// DecodeContent method to retrieve the byte data from a Q encoded string.
func (z *Xml) QEncodeContent(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkXml_QEncodeContent(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Removes all attributes from an XML node. Should always return True.
func (z *Xml) RemoveAllAttributes() bool {

    v := C.CkXml_RemoveAllAttributes(z.ckObj)


    return v != 0
}


// Removes all children from the calling node.
func (z *Xml) RemoveAllChildren()  {

    C.CkXml_RemoveAllChildren(z.ckObj)



}


// Removes an attribute by name from and XML node.
func (z *Xml) RemoveAttribute(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_RemoveAttribute(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Removes all children with a given tag or tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) RemoveChild(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkXml_RemoveChild(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Removes the Nth child from the calling node.
func (z *Xml) RemoveChildByIndex(arg1 int)  {

    C.CkXml_RemoveChildByIndex(z.ckObj, C.int(arg1))



}


// Removes all children having the exact content specified.
func (z *Xml) RemoveChildWithContent(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkXml_RemoveChildWithContent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Removes the calling object and its sub-tree from the XML document making it the
// root of its own tree.
func (z *Xml) RemoveFromTree()  {

    C.CkXml_RemoveFromTree(z.ckObj)



}


// Removes all XML stylesheets having an attribute with attrName equal to attrValue. Returns
// the number of stylesheets removed, or -1 if there was an error.
func (z *Xml) RemoveStyleSheet(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXml_RemoveStyleSheet(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Saves a node's binary content to a file.
func (z *Xml) SaveBinaryContent(arg1 string, arg2 bool, arg3 bool, arg4 string) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    cstr4 := C.CString(arg4)

    v := C.CkXml_SaveBinaryContent(z.ckObj, cstr1, b2, b3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Generates XML representing the tree or subtree rooted at this node and writes it
// to a file.
func (z *Xml) SaveXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_SaveXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Recursively traverses the XML rooted at the caller and scrubs according to the
// instructions in the comma separated directives. The currently defined directives are:
//     "AttrTrimEnds" - Leading and trailing whitespace removed from attribute
//     values.
//     "AttrTrimInside" - Replace all tabs, CR's, and LF's with SPACE chars, and
//     removes extra SPACE chars so that no more than one SPACE char in a row exists.
//     "ContentTrimEnds" - Same as AttrTrimEnds but for content.
//     "ContentTrimInside" - Same as AttrTrimInside but for content.
//     "LowercaseAttrs" - Convert all attribute names to lowercase.
//     "LowercaseTags" - Convert all tags to lowercase.
//     "RemoveCtrl" - Remove non-printable us-ascii control chars (us-ascii values
// 
// If you have other ideas for useful XML scrubbing directives, send email to
// support@chilkatsoft.com. It must be general enough such that many developers
// will find it useful.
//
func (z *Xml) Scrub(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkXml_Scrub(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Returns the first node having content matching the contentPattern. The contentPattern is a
// case-sensitive string that may contain any number of '*'s, each representing 0
// or more occurrences of any character. The search is breadth-first over the
// sub-tree rooted at the caller. A match is returned only after the search has
// traversed past the node indicated by afterPtr. To find the 1st occurrence, set afterPtr
// equal to _NULL_. (For the ActiveX implementation, the afterPtr should never be
// _NULL_. A reference to the caller's node should be passed instead.)
// 
// To iterate over matching nodes, the returned node can be passed in afterPtr for the
// next call to SearchAllForContent, until the method returns _NULL_.
//
func (z *Xml) SearchAllForContent(arg1 *Xml, arg2 string) *Xml {
    cstr2 := C.CString(arg2)

    retObj := C.CkXml_SearchAllForContent(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Same as SearchAllForContent except the internal reference of the caller is
// updated to point to the search result (instead of returning a new object).
func (z *Xml) SearchAllForContent2(arg1 *Xml, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkXml_SearchAllForContent2(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the first node having a tag equal to tag, an attribute named attr,
// whose value matches valuePattern. The valuePattern is a case-sensitive string that may contain
// any number of '*'s, each representing 0 or more occurrences of any character.
// The search is breadth-first over the sub-tree rooted at the caller. A match is
// returned only after the search has traversed past the node indicated by afterPtr. To
// find the 1st occurrence, set afterPtr equal to _NULL_. (For the ActiveX
// implementation, the afterPtr should never be _NULL_. A reference to the caller's
// node should be passed instead.)
// 
// To iterate over matching nodes, the returned node can be passed in afterPtr for the
// next call to SearchForAttribute, until the method returns _NULL_.
//
func (z *Xml) SearchForAttribute(arg1 *Xml, arg2 string, arg3 string, arg4 string) *Xml {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    retObj := C.CkXml_SearchForAttribute(z.ckObj, arg1.ckObj, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Same as SearchForAttribute except the internal reference of the caller is
// updated to point to the search result (instead of returning a new object).
func (z *Xml) SearchForAttribute2(arg1 *Xml, arg2 string, arg3 string, arg4 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkXml_SearchForAttribute2(z.ckObj, arg1.ckObj, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Returns the first node having a tag equal to tag, whose content matches contentPattern.
// The contentPattern is a case-sensitive string that may contain any number of '*'s, each
// representing 0 or more occurrences of any character. The search is breadth-first
// over the sub-tree rooted at the caller. A match is returned only after the
// search has traversed past the node indicated by afterPtr. To find the 1st
// occurrence, set afterPtr equal to _NULL_. (For the ActiveX implementation, the afterPtr
// should never be _NULL_. A reference to the caller's node should be passed
// instead.)
// 
// To iterate over matching nodes, the returned node can be passed in afterPtr for the
// next call to SearchForContent, until the method returns _NULL_.
//
func (z *Xml) SearchForContent(arg1 *Xml, arg2 string, arg3 string) *Xml {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retObj := C.CkXml_SearchForContent(z.ckObj, arg1.ckObj, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Same as SearchForContent except the internal reference of the caller is updated
// to point to the search result (instead of returning a new object).
func (z *Xml) SearchForContent2(arg1 *Xml, arg2 string, arg3 string) bool {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkXml_SearchForContent2(z.ckObj, arg1.ckObj, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Returns the first node having a tag equal to tag. The search is breadth-first
// over the sub-tree rooted at the caller. A match is returned only after the
// search has traversed past the node indicated by afterPtr. To find the 1st
// occurrence, set afterPtr equal to _NULL_. (For the ActiveX implementation, the afterPtr
// should never be _NULL_. A reference to the caller's node should be passed
// instead.)
// 
// To iterate over matching nodes, the returned node can be passed in afterPtr for the
// next call to SearchForTag, until the method returns _NULL_.
//
func (z *Xml) SearchForTag(arg1 *Xml, arg2 string) *Xml {
    cstr2 := C.CString(arg2)

    retObj := C.CkXml_SearchForTag(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Xml{retObj}
}


// Same as SearchForTag except the internal reference of the caller is updated to
// point to the search result (instead of returning a new object).
func (z *Xml) SearchForTag2(arg1 *Xml, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkXml_SearchForTag2(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the node's content to a block of binary data with optional Zip compression
// and/or AES encryption. The binary data is automatically converted to base64
// format whenever XML text is generated. If the zipFlag is True, the data is first
// compressed. If the encryptFlag is True, the data is AES encrypted using the
// Rijndael 128-bit symmetric-encryption algorithm.
func (z *Xml) SetBinaryContent(arg1 []byte, arg2 bool, arg3 bool, arg4 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    cstr4 := C.CString(arg4)

    v := C.CkXml_SetBinaryContent(z.ckObj, ckbyd1, b2, b3, cstr4)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Sets the node's content with binary (or text) data from a file. The file
// contents can be Zip compressed and/or encrypted, and the result is base-64
// encoded.
func (z *Xml) SetBinaryContentFromFile(arg1 string, arg2 bool, arg3 bool, arg4 string) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    cstr4 := C.CString(arg4)

    v := C.CkXml_SetBinaryContentFromFile(z.ckObj, cstr1, b2, b3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Sorts the direct child nodes by the value of a specified attribute.
func (z *Xml) SortByAttribute(arg1 string, arg2 bool)  {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    C.CkXml_SortByAttribute(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))


}


// Sorts the direct child nodes by the value of a specified attribute interpreted
// as an integer (not lexicographically as strings).
func (z *Xml) SortByAttributeInt(arg1 string, arg2 bool)  {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    C.CkXml_SortByAttributeInt(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))


}


// Sorts the direct child nodes by content.
func (z *Xml) SortByContent(arg1 bool)  {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    C.CkXml_SortByContent(z.ckObj, b1)



}


// Sorts the direct child nodes by tag.
func (z *Xml) SortByTag(arg1 bool)  {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    C.CkXml_SortByTag(z.ckObj, b1)



}


// Sorts the direct child nodes by the content of an attribute in the grandchild
// nodes.
func (z *Xml) SortRecordsByAttribute(arg1 string, arg2 string, arg3 bool)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    C.CkXml_SortRecordsByAttribute(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Sorts the direct child nodes by the content of the grandchild nodes.
func (z *Xml) SortRecordsByContent(arg1 string, arg2 bool)  {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    C.CkXml_SortRecordsByContent(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))


}


// Sorts the direct child nodes by the content of the grandchild nodes. For sorting
// purposes, the content is interpreted as an integer (not lexicographically as for
// strings).
func (z *Xml) SortRecordsByContentInt(arg1 string, arg2 bool)  {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    C.CkXml_SortRecordsByContentInt(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))


}


// Swaps another node's tag, content, and attributes with this one.
func (z *Xml) SwapNode(arg1 *Xml) bool {

    v := C.CkXml_SwapNode(z.ckObj, arg1.ckObj)


    return v != 0
}


// Swaps another node's tag, content, attributes, and children with this one.
func (z *Xml) SwapTree(arg1 *Xml) bool {

    v := C.CkXml_SwapTree(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the content of the 1st node found in the sub-tree rooted at the caller
// that has a given tag. (Note: The search for the node having tag ARG is not
// limited to the direct children of the caller.)
// return a string or nil if failed.
func (z *Xml) TagContent(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkXml_tagContent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if the node's tag equals the specified string.
func (z *Xml) TagEquals(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_TagEquals(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns the index of the XML node with the given tag. Returns -1 if no node
// having the specified tag is found at the tagPath.
func (z *Xml) TagIndex(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkXml_TagIndex(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns true if the node's tag namespace prefix equals the specified ns.
func (z *Xml) TagNsEquals(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_TagNsEquals(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Returns true if the node's unprefixed tag equals the specified string. For
// example, if the tag is "soapenv:Body", the unprefixed tag is "Body".
func (z *Xml) TagUnpEquals(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_TagUnpEquals(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unzip the content of the XML node replacing it's content with the decompressed
// data.
func (z *Xml) UnzipContent() bool {

    v := C.CkXml_UnzipContent(z.ckObj)


    return v != 0
}


// Unzips and recreates the XML node and the entire subtree, restoring it to the
// state before it was zip compressed.
func (z *Xml) UnzipTree() bool {

    v := C.CkXml_UnzipTree(z.ckObj)


    return v != 0
}


// Updates the content for the node indicated by tagPath. If autoCreate is true, then
// nodes along tagPath are auto-created as needed.
func (z *Xml) UpdateAt(arg1 string, arg2 bool, arg3 string) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    v := C.CkXml_UpdateAt(z.ckObj, cstr1, b2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Updates or adds the attribute value for the node indicated by tagPath. If autoCreate is
// true, then nodes along tagPath are auto-created as needed.
func (z *Xml) UpdateAttrAt(arg1 string, arg2 bool, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkXml_UpdateAttrAt(z.ckObj, cstr1, b2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Adds an attribute to the node if it doesn't already exist. Otherwise it updates
// the existing attribute with the new value.
func (z *Xml) UpdateAttribute(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkXml_UpdateAttribute(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Updates an attribute value. (Call UpdateAttribute if the attribute value is a
// string.)
func (z *Xml) UpdateAttributeInt(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkXml_UpdateAttributeInt(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Replaces the content of a child node. The tagPath can be a tag or tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) UpdateChildContent(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkXml_UpdateChildContent(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Replaces the content of a child node where the content is an integer. The tagPath
// can be a tag or tag path.
// 
// Beginning in version 9.5.0.64, the tagPath can be a tag path. A tag path is a
// series of tags separated by vertical bar characters. For example:
// "tagA|tagB|tagC".
//
func (z *Xml) UpdateChildContentInt(arg1 string, arg2 int)  {
    cstr1 := C.CString(arg1)

    C.CkXml_UpdateChildContentInt(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Applies Zip compression to the content of an XML node and replaces the content
// with base64-encoded compressed data.
func (z *Xml) ZipContent() bool {

    v := C.CkXml_ZipContent(z.ckObj)


    return v != 0
}


// Zip compresses the content and entire subtree rooted at the calling XML node and
// replaces the current content with base64-encoded Zip compressed data. The node
// and subtree can be restored by calling UnzipTree. Note that the node name and
// attributes are unaffected.
func (z *Xml) ZipTree() bool {

    v := C.CkXml_ZipTree(z.ckObj)


    return v != 0
}


