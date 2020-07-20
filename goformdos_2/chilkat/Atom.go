// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkAtom.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewAtom() *Atom {
	return &Atom{C.CkAtom_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Atom) DisposeAtom() {
    if z != nil {
	C.CkAtom_Dispose(z.ckObj)
	}
}




func (z *Atom) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkAtom_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Atom) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkAtom_setExternalProgress(z.ckObj,1)
}

func (z *Atom) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkAtom_setExternalProgress(z.ckObj,1)
}

func (z *Atom) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkAtom_setExternalProgress(z.ckObj,1)
}




// property: AbortCurrent
// When set to true, causes the currently running method to abort. Methods that
// always finish quickly (i.e.have no length file operations or network
// communications) are not affected. If no method is running, then this property is
// automatically reset to false when the next method is called. When the abort
// occurs, this property is reset to false. Both synchronous and asynchronous
// method calls can be aborted. (A synchronous method call could be aborted by
// setting this property from a separate thread.)
func (z *Atom) AbortCurrent() bool {
    v := int(C.CkAtom_getAbortCurrent(z.ckObj))
    return v != 0
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
func (z *Atom) DebugLogFilePath() string {
    return C.GoString(C.CkAtom_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Atom) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkAtom_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Atom) LastErrorHtml() string {
    return C.GoString(C.CkAtom_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Atom) LastErrorText() string {
    return C.GoString(C.CkAtom_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Atom) LastErrorXml() string {
    return C.GoString(C.CkAtom_lastErrorXml(z.ckObj))
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
func (z *Atom) LastMethodSuccess() bool {
    v := int(C.CkAtom_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Atom) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAtom_putLastMethodSuccess(z.ckObj,v)
}

// property: NumEntries
// Number of entries in the Atom document.
func (z *Atom) NumEntries() int {
    return int(C.CkAtom_getNumEntries(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Atom) VerboseLogging() bool {
    v := int(C.CkAtom_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Atom) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkAtom_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Atom) Version() string {
    return C.GoString(C.CkAtom_version(z.ckObj))
}
// Adds a new element to the Atom document. The tag is a string such as "title",
// "subtitle", "summary", etc. Returns the index of the element added, or -1 for
// failure.
func (z *Atom) AddElement(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAtom_AddElement(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Adds a new date-formatted element to the Atom document. The tag is a string
// such as "created", "modified", "issued", etc. The dateTimeStr should be an RFC822
// formatted date/time string such as "Tue, 25 Sep 2012 12:25:32 -0500". Returns
// the index of the element added, or -1 for failure.
func (z *Atom) AddElementDateStr(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAtom_AddElementDateStr(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Adds a new date-formatted element to the Atom document. The tag is a string such
// as "created", "modified", "issued", etc. Returns the index of the element added,
// or -1 for failure.
func (z *Atom) AddElementDt(arg1 string, arg2 *CkDateTime) int {
    cstr1 := C.CString(arg1)

    v := C.CkAtom_AddElementDt(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Adds a new HTML formatted element to the Atom document. Returns the index of the
// element added, or -1 for failure.
func (z *Atom) AddElementHtml(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAtom_AddElementHtml(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Adds a new XHTML formatted element to the Atom document. Returns the index of
// the element added, or -1 for failure.
func (z *Atom) AddElementXHtml(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAtom_AddElementXHtml(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Adds a new XML formatted element to the Atom document. Returns the index of the
// element added, or -1 for failure.
func (z *Atom) AddElementXml(arg1 string, arg2 string) int {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkAtom_AddElementXml(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return int(v)
}


// Adds an "entry" Atom XML document to the caller's Atom document.
func (z *Atom) AddEntry(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkAtom_AddEntry(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Adds a link to the Atom document.
func (z *Atom) AddLink(arg1 string, arg2 string, arg3 string, arg4 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    C.CkAtom_AddLink(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))


}


// Adds a person to the Atom document. The tag should be a string such as "author",
// "contributor", etc. If a piece of information is not known, an empty string or
// NULL value may be passed.
func (z *Atom) AddPerson(arg1 string, arg2 string, arg3 string, arg4 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    C.CkAtom_AddPerson(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))


}


// Removes the Nth occurrence of a given element from the Atom document. Indexing
// begins at 0. For example, to remove the 2nd category, set tag = "category" and
// index = 1.
func (z *Atom) DeleteElement(arg1 string, arg2 int)  {
    cstr1 := C.CString(arg1)

    C.CkAtom_DeleteElement(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Remove an attribute from an element.The index should be 0 unless there are
// multiple elements having the same tag, in which case it selects the Nth
// occurrence based on the index ( 0 = first occurrence ).
func (z *Atom) DeleteElementAttr(arg1 string, arg2 int, arg3 string)  {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    C.CkAtom_DeleteElementAttr(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))


}


// Deletes a person from the Atom document. The tag is a string such as "author".
// The index should be 0 unless there are multiple elements having the same tag, in
// which case it selects the Nth occurrence based on the index. For example,
// DeletePerson("author",2) deletes the 3rd author.
func (z *Atom) DeletePerson(arg1 string, arg2 int)  {
    cstr1 := C.CString(arg1)

    C.CkAtom_DeletePerson(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Download an Atom feed from the Internet and load it into the Atom object.
func (z *Atom) DownloadAtom(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkAtom_DownloadAtom(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the DownloadAtom method
func (z *Atom) DownloadAtomAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkAtom_DownloadAtomAsync(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Returns the content of the Nth element having a specified tag.
// return a string or nil if failed.
func (z *Atom) GetElement(arg1 string, arg2 int) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkAtom_getElement(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of an element's attribute. The element is selected by the tag
// name and the index (the Nth element having a specific tag) and the attribute is
// selected by name.
// return a string or nil if failed.
func (z *Atom) GetElementAttr(arg1 string, arg2 int, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    retStr := C.CkAtom_getElementAttr(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The number of elements having a specific tag.
func (z *Atom) GetElementCount(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkAtom_GetElementCount(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns an element's value as a date/time in an RFC822 formatted string, such as
// such as "Tue, 25 Sep 2012 12:25:32 -0500".
// return a string or nil if failed.
func (z *Atom) GetElementDateStr(arg1 string, arg2 int) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkAtom_getElementDateStr(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns an element's value as a date/time object.
func (z *Atom) GetElementDt(arg1 string, arg2 int) *CkDateTime {
    cstr1 := C.CString(arg1)

    retObj := C.CkAtom_GetElementDt(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Returns the Nth entry as an Atom object. (Indexing begins at 0)
func (z *Atom) GetEntry(arg1 int) *Atom {

    retObj := C.CkAtom_GetEntry(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Atom{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Returns the href attribute of the link having a specified "rel" attribute (such
// as "service.feed", "alternate", etc.).
// return a string or nil if failed.
func (z *Atom) GetLinkHref(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkAtom_getLinkHref(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a piece of information about a person. To get the 2nd author's name,
// call GetPersonInfo("author",1,"name").
// return a string or nil if failed.
func (z *Atom) GetPersonInfo(arg1 string, arg2 int, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    retStr := C.CkAtom_getPersonInfo(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of an attribute on the top-level XML node. The tag of a
// top-level Atom XML node is typically "feed" or "entry", and it might have
// attributes such as "xmlns" and "xml:lang".
// return a string or nil if failed.
func (z *Atom) GetTopAttr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkAtom_getTopAttr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// True (1) if the element exists in the Atom document. Otherwise 0.
func (z *Atom) HasElement(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkAtom_HasElement(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the caller of the task's async method.
func (z *Atom) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkAtom_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads the Atom document from an XML string.
func (z *Atom) LoadXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkAtom_LoadXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Initializes the Atom document to be a new "entry".
func (z *Atom) NewEntry()  {

    C.CkAtom_NewEntry(z.ckObj)



}


// Initializes the Atom document to be a new "feed".
func (z *Atom) NewFeed()  {

    C.CkAtom_NewFeed(z.ckObj)



}


// Adds or replaces an attribute on an element.
func (z *Atom) SetElementAttr(arg1 string, arg2 int, arg3 string, arg4 string)  {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    C.CkAtom_SetElementAttr(z.ckObj, cstr1, C.int(arg2), cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))


}


// Adds or replaces an attribute on the top-level XML node of the Atom document.
func (z *Atom) SetTopAttr(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkAtom_SetTopAttr(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Serializes the Atom document to an XML string.
// return a string or nil if failed.
func (z *Atom) ToXmlString() *string {

    retStr := C.CkAtom_toXmlString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Replaces the content of an element.
func (z *Atom) UpdateElement(arg1 string, arg2 int, arg3 string)  {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    C.CkAtom_UpdateElement(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))


}


// Replaces the content of a date-formatted element. The index should be an RFC822
// formatted date/time string.
func (z *Atom) UpdateElementDateStr(arg1 string, arg2 int, arg3 string)  {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    C.CkAtom_UpdateElementDateStr(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))


}


// Replaces the content of a date-formatted element.
func (z *Atom) UpdateElementDt(arg1 string, arg2 int, arg3 *CkDateTime)  {
    cstr1 := C.CString(arg1)

    C.CkAtom_UpdateElementDt(z.ckObj, cstr1, C.int(arg2), arg3.ckObj)

    C.free(unsafe.Pointer(cstr1))


}


// Replaces the content of an HTML element.
func (z *Atom) UpdateElementHtml(arg1 string, arg2 int, arg3 string)  {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    C.CkAtom_UpdateElementHtml(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))


}


// Replaces the content of an XHTML element.
func (z *Atom) UpdateElementXHtml(arg1 string, arg2 int, arg3 string)  {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    C.CkAtom_UpdateElementXHtml(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))


}


// Replaces the content of an XML element.
func (z *Atom) UpdateElementXml(arg1 string, arg2 int, arg3 string)  {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    C.CkAtom_UpdateElementXml(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))


}


// Replaces the content of a person. To update the 3rd author, call
// UpdatePerson("author",2,"new name","new URL","new email"). If a piece of
// information is not known, pass an empty string or a NULL.
func (z *Atom) UpdatePerson(arg1 string, arg2 int, arg3 string, arg4 string, arg5 string)  {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)
    cstr5 := C.CString(arg5)

    C.CkAtom_UpdatePerson(z.ckObj, cstr1, C.int(arg2), cstr3, cstr4, cstr5)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))
    C.free(unsafe.Pointer(cstr5))


}


