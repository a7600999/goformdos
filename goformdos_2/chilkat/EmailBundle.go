// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkEmailBundle.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewEmailBundle() *EmailBundle {
	return &EmailBundle{C.CkEmailBundle_Create()}
}

func (z *EmailBundle) DisposeEmailBundle() {
    if z != nil {
	C.CkEmailBundle_Dispose(z.ckObj)
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
func (z *EmailBundle) DebugLogFilePath() string {
    return C.GoString(C.CkEmailBundle_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *EmailBundle) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmailBundle_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *EmailBundle) LastErrorHtml() string {
    return C.GoString(C.CkEmailBundle_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *EmailBundle) LastErrorText() string {
    return C.GoString(C.CkEmailBundle_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *EmailBundle) LastErrorXml() string {
    return C.GoString(C.CkEmailBundle_lastErrorXml(z.ckObj))
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
func (z *EmailBundle) LastMethodSuccess() bool {
    v := int(C.CkEmailBundle_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *EmailBundle) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmailBundle_putLastMethodSuccess(z.ckObj,v)
}

// property: MessageCount
// The number of emails in this bundle.
func (z *EmailBundle) MessageCount() int {
    return int(C.CkEmailBundle_getMessageCount(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *EmailBundle) VerboseLogging() bool {
    v := int(C.CkEmailBundle_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *EmailBundle) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmailBundle_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *EmailBundle) Version() string {
    return C.GoString(C.CkEmailBundle_version(z.ckObj))
}
// Adds an email object to the bundle.
func (z *EmailBundle) AddEmail(arg1 *Email) bool {

    v := C.CkEmailBundle_AddEmail(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the first email having a header field matching the headerFieldName and headerFieldValue exactly
// (case sensitive). If no matching email is found, returns _NULL_.
func (z *EmailBundle) FindByHeader(arg1 string, arg2 string) *Email {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retObj := C.CkEmailBundle_FindByHeader(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Returns the Nth Email in the bundle. The email returned is a copy of the email
// in the bundle. Updating the email object returned by GetEmail has no effect on
// the email within the bundle. To update/replace the email in the bundle, your
// program should call GetEmail to get a copy, make modifications, call
// RemoveEmailByIndex to remove the email (passing the same index used in the call
// to GetEmail), and then call AddEmail to insert the new/modified email into the
// bundle.
// 
// IMPORTANT: This method does NOT communicate with any mail server to download the
// email. It simply returns the Nth email object that exists within it's in-memory
// collection of email objects.
//
func (z *EmailBundle) GetEmail(arg1 int) *Email {

    retObj := C.CkEmailBundle_GetEmail(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Returns a StringArray object containing UIDLs for all Email objects in the
// bundle. UIDLs are only valid for emails retrieved from POP3 servers. An email on
// a POP3 server has a "UIDL", an email on IMAP servers has a "UID". If the email
// was retrieved from an IMAP server, the UID will be accessible via the
// "ckx-imap-uid" header field.
func (z *EmailBundle) GetUidls() *StringArray {

    retObj := C.CkEmailBundle_GetUidls(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Converts the email bundle to an XML document in memory. Returns the XML document
// as a string.
// return a string or nil if failed.
func (z *EmailBundle) GetXml() *string {

    retStr := C.CkEmailBundle_getXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Loads the email bundle from a completed asynchronous task.
func (z *EmailBundle) LoadTaskResult(arg1 *Task) bool {

    v := C.CkEmailBundle_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads an email bundle from an XML file.
func (z *EmailBundle) LoadXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmailBundle_LoadXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads an email bundle from an XML string.
func (z *EmailBundle) LoadXmlString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmailBundle_LoadXmlString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Removes an email from the bundle. This does not remove the email from the mail
// server.
func (z *EmailBundle) RemoveEmail(arg1 *Email) bool {

    v := C.CkEmailBundle_RemoveEmail(z.ckObj, arg1.ckObj)


    return v != 0
}


// Removes the Nth email in a bundle. (Indexing begins at 0.)
func (z *EmailBundle) RemoveEmailByIndex(arg1 int) bool {

    v := C.CkEmailBundle_RemoveEmailByIndex(z.ckObj, C.int(arg1))


    return v != 0
}


// Converts each email to XML and persists the bundle to an XML file. The email
// bundle can later be re-instantiated by calling MailMan.LoadXmlFile
func (z *EmailBundle) SaveXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmailBundle_SaveXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sorts emails in the bundle by date.
func (z *EmailBundle) SortByDate(arg1 bool)  {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    C.CkEmailBundle_SortByDate(z.ckObj, b1)



}


// Sorts emails in the bundle by recipient.
func (z *EmailBundle) SortByRecipient(arg1 bool)  {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    C.CkEmailBundle_SortByRecipient(z.ckObj, b1)



}


// Sorts emails in the bundle by sender.
func (z *EmailBundle) SortBySender(arg1 bool)  {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    C.CkEmailBundle_SortBySender(z.ckObj, b1)



}


// Sorts emails in the bundle by subject.
func (z *EmailBundle) SortBySubject(arg1 bool)  {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    C.CkEmailBundle_SortBySubject(z.ckObj, b1)



}


