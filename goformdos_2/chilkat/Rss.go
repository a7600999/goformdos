// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkRss.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewRss() *Rss {
	return &Rss{C.CkRss_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *Rss) DisposeRss() {
    if z != nil {
	C.CkRss_Dispose(z.ckObj)
	}
}




func (z *Rss) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkRss_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *Rss) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkRss_setExternalProgress(z.ckObj,1)
}

func (z *Rss) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkRss_setExternalProgress(z.ckObj,1)
}

func (z *Rss) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkRss_setExternalProgress(z.ckObj,1)
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
func (z *Rss) DebugLogFilePath() string {
    return C.GoString(C.CkRss_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Rss) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkRss_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Rss) LastErrorHtml() string {
    return C.GoString(C.CkRss_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Rss) LastErrorText() string {
    return C.GoString(C.CkRss_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Rss) LastErrorXml() string {
    return C.GoString(C.CkRss_lastErrorXml(z.ckObj))
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
func (z *Rss) LastMethodSuccess() bool {
    v := int(C.CkRss_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Rss) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRss_putLastMethodSuccess(z.ckObj,v)
}

// property: NumChannels
// The number of channels in the RSS document.
func (z *Rss) NumChannels() int {
    return int(C.CkRss_getNumChannels(z.ckObj))
}

// property: NumItems
// The number of items in the channel.
func (z *Rss) NumItems() int {
    return int(C.CkRss_getNumItems(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Rss) VerboseLogging() bool {
    v := int(C.CkRss_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Rss) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkRss_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Rss) Version() string {
    return C.GoString(C.CkRss_version(z.ckObj))
}
// Adds a new channel to the RSS document. Returns the Rss object representing the
// Channel which can then be edited.
func (z *Rss) AddNewChannel() *Rss {

    retObj := C.CkRss_AddNewChannel(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Rss{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Adds a new image to the RSS document. Returns the Rss object representing the
// image, which can then be edited.
func (z *Rss) AddNewImage() *Rss {

    retObj := C.CkRss_AddNewImage(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Rss{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Adds a new Item to an Rss channel. Returns the Rss object representing the item
// which can then be edited.
func (z *Rss) AddNewItem() *Rss {

    retObj := C.CkRss_AddNewItem(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Rss{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Downloads an RSS document from the Internet and populates the Rss object with
// the contents.
func (z *Rss) DownloadRss(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRss_DownloadRss(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the DownloadRss method
func (z *Rss) DownloadRssAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkRss_DownloadRssAsync(z.ckObj, cstr1)

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


// Returns the value of a sub-element attribute. For example, to get the value of
// the "isPermaLink" attribute of the "guid" sub-element, call
// item.GetAttr("guid","isPermaLink").
// return a string or nil if failed.
func (z *Rss) GetAttr(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkRss_getAttr(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth channel of an RSS document. Usually there is only 1 channel per
// document, so the index argument should be set to 0.
func (z *Rss) GetChannel(arg1 int) *Rss {

    retObj := C.CkRss_GetChannel(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Rss{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Return the number of sub-elements with a specific tag.
func (z *Rss) GetCount(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkRss_GetCount(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// The same as GetDate, except the date/time is returned in RFC822 string format.
// return a string or nil if failed.
func (z *Rss) GetDateStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkRss_getDateStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Return the image associated with the channel.
func (z *Rss) GetImage() *Rss {

    retObj := C.CkRss_GetImage(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Rss{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Return the value of a numeric sub-element as an integer.
func (z *Rss) GetInt(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkRss_GetInt(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Return the Nth item of a channel as an RSS object.
func (z *Rss) GetItem(arg1 int) *Rss {

    retObj := C.CkRss_GetItem(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Rss{retObj,ChilkatCallbacks{nil,nil,nil,nil}}
}


// Return the value of an sub-element as a string.
// return a string or nil if failed.
func (z *Rss) GetString(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkRss_getString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Load an RSS document from a file.
func (z *Rss) LoadRssFile(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRss_LoadRssFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads an RSS feed document from an in-memory string.
func (z *Rss) LoadRssString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkRss_LoadRssString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the caller of the task's async method.
func (z *Rss) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkRss_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Get an attribute value for the Nth sub-element having a specific tag. As an
// example, an RSS item may have several "category" sub-elements. To get the value
// of the "domain" attribute for the 3rd category, call
// MGetAttr("category",2,"domain").
// return a string or nil if failed.
func (z *Rss) MGetAttr(arg1 string, arg2 int, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    retStr := C.CkRss_mGetAttr(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Get the value of the Nth occurrence of a sub-element. Indexing begins at 0.
// return a string or nil if failed.
func (z *Rss) MGetString(arg1 string, arg2 int) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkRss_mGetString(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Set an attribute on the Nth occurrence of a sub-element.
func (z *Rss) MSetAttr(arg1 string, arg2 int, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkRss_MSetAttr(z.ckObj, cstr1, C.int(arg2), cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Set the value of the Nth occurrence of a sub-element. Indexing begins at 0.
func (z *Rss) MSetString(arg1 string, arg2 int, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    v := C.CkRss_MSetString(z.ckObj, cstr1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Clears the RSS document.
func (z *Rss) NewRss()  {

    C.CkRss_NewRss(z.ckObj)



}


// Removes a sub-element from the RSS document.
func (z *Rss) Remove(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkRss_Remove(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Sets the value of a sub-element attribute.
func (z *Rss) SetAttr(arg1 string, arg2 string, arg3 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    C.CkRss_SetAttr(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))


}


// Sets the value of a date/time sub-element to the current system date/time.
func (z *Rss) SetDateNow(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkRss_SetDateNow(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// The same as SetDate, except the date/time is passed as an RFC822 string.
func (z *Rss) SetDateStr(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkRss_SetDateStr(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Sets the value of an integer sub-element.
func (z *Rss) SetInt(arg1 string, arg2 int)  {
    cstr1 := C.CString(arg1)

    C.CkRss_SetInt(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))


}


// Sets the value of a sub-element.
func (z *Rss) SetString(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkRss_SetString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Returns the RSS document as an XML string.
// return a string or nil if failed.
func (z *Rss) ToXmlString() *string {

    retStr := C.CkRss_toXmlString(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


