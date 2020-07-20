// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkBounce.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewBounce() *Bounce {
	return &Bounce{C.CkBounce_Create()}
}

func (z *Bounce) DisposeBounce() {
    if z != nil {
	C.CkBounce_Dispose(z.ckObj)
	}
}




// property: BounceAddress
// The bounced email address when a bounced email is recognized.
func (z *Bounce) BounceAddress() string {
    return C.GoString(C.CkBounce_bounceAddress(z.ckObj))
}

// property: BounceData
// The raw body of the bounced mail.
func (z *Bounce) BounceData() string {
    return C.GoString(C.CkBounce_bounceData(z.ckObj))
}

// property: BounceType
// A number representing the type of bounce that was recognized.
// A value of 0 indicates "No Bounce". Other values are:
// 
// 1. Hard Bounce. The email could not be delivered and BounceAddress contains the
// failed email address.
// 2. Soft Bounce. A temporary condition exists causing the email delivery to fail.
// The BounceAddress property contains the failed email address.
// 3. General Bounced Mail, cannot determine if it is hard or soft, and the email
// address is not available.
// 4. General Bounced Mail, cannot determine if it is hard or soft, but an email
// address is available.
// 5. Mail Block. A bounce occured because the sender was blocked.
// 6. Auto-Reply/Out-of-Office email.
// 7. Transient message, such as "Delivery Status / No Action Required".
// 8. Subscribe request.
// 9. Unsubscribe request.
// 10. Virus email notification.
// 11. Suspected Bounce, but no other information is available
// 12. Challenge/Response - Auto-reply message sent by SPAM software where only
// verified email addresses are accepted.
// 13. Address Change Notification Messages.
// 14. Success DSN indicating that the message was successfully relayed.
// 15. Abuse/fraud feedback report.
func (z *Bounce) BounceType() int {
    return int(C.CkBounce_getBounceType(z.ckObj))
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
func (z *Bounce) DebugLogFilePath() string {
    return C.GoString(C.CkBounce_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Bounce) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkBounce_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Bounce) LastErrorHtml() string {
    return C.GoString(C.CkBounce_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Bounce) LastErrorText() string {
    return C.GoString(C.CkBounce_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Bounce) LastErrorXml() string {
    return C.GoString(C.CkBounce_lastErrorXml(z.ckObj))
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
func (z *Bounce) LastMethodSuccess() bool {
    v := int(C.CkBounce_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Bounce) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkBounce_putLastMethodSuccess(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Bounce) VerboseLogging() bool {
    v := int(C.CkBounce_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Bounce) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkBounce_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Bounce) Version() string {
    return C.GoString(C.CkBounce_version(z.ckObj))
}
// Examines an email and sets the properties (BounceType, BounceAddress,
// BounceData) according to how the email is classified. This feature can only be
// used if Chilkat Mail is downloaded and installed, and it also requires Chilkat
// Mail to be licensed in addition to Chilkat Bounce.
func (z *Bounce) ExamineEmail(arg1 *Email) bool {

    v := C.CkBounce_ExamineEmail(z.ckObj, arg1.ckObj)


    return v != 0
}


// Examines an email from a .eml file and sets the properties (BounceType,
// BounceAddress, BounceData) according to how the email is classified.
func (z *Bounce) ExamineEml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkBounce_ExamineEml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Examines an email represented as raw MIME text and sets the properties
// (BounceType, BounceAddress, BounceData) according to how the email is
// classified. The return value is 1 for a successful classification, or 0 for a
// failure.
func (z *Bounce) ExamineMime(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkBounce_ExamineMime(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unlocks the component. This must be called once at the beginning of your program
// to unlock the component. A purchased unlock code is provided when the Bounce
// component is licensed.
// 
// A purchased unlock code for the bounce component/library will included the
// substring "BOUNCE", or can be a Bundle unlock code.
//
func (z *Bounce) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkBounce_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


