// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkDateTime.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCkDateTime() *CkDateTime {
	return &CkDateTime{C.CkDateTime_Create()}
}

func (z *CkDateTime) DisposeCkDateTime() {
    if z != nil {
	C.CkDateTime_Dispose(z.ckObj)
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
func (z *CkDateTime) DebugLogFilePath() string {
    return C.GoString(C.CkDateTime_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *CkDateTime) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkDateTime_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: IsDst
// This is the Daylight Saving Time flag. It can have one of three possible values:
// 1, 0, or -1. It has the value 1 if Daylight Saving Time is in effect, 0 if
// Daylight Saving Time is not in effect, and -1 if the information is not
// available.
// 
// Note: This is NOT the DST for the current system time. It is the DST that was in
// effect at the date value contained in this object.
//
func (z *CkDateTime) IsDst() int {
    return int(C.CkDateTime_getIsDst(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *CkDateTime) LastErrorHtml() string {
    return C.GoString(C.CkDateTime_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *CkDateTime) LastErrorText() string {
    return C.GoString(C.CkDateTime_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *CkDateTime) LastErrorXml() string {
    return C.GoString(C.CkDateTime_lastErrorXml(z.ckObj))
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
func (z *CkDateTime) LastMethodSuccess() bool {
    v := int(C.CkDateTime_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *CkDateTime) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDateTime_putLastMethodSuccess(z.ckObj,v)
}

// property: UtcOffset
// For the current system's timezone, returns the number of seconds offset from UTC
// for this date/time. The offset includes daylight savings adjustment. Local
// timezones west of UTC return a negative offset.
func (z *CkDateTime) UtcOffset() int {
    return int(C.CkDateTime_getUtcOffset(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *CkDateTime) VerboseLogging() bool {
    v := int(C.CkDateTime_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *CkDateTime) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkDateTime_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *CkDateTime) Version() string {
    return C.GoString(C.CkDateTime_version(z.ckObj))
}
// Adds an integer number of days to the date/time. To subtract days, pass a
// negative integer.
func (z *CkDateTime) AddDays(arg1 int) bool {

    v := C.CkDateTime_AddDays(z.ckObj, C.int(arg1))


    return v != 0
}


// Adds an integer number of seconds to the date/time. To subtract seconds, pass a
// negative integer.
func (z *CkDateTime) AddSeconds(arg1 int) bool {

    v := C.CkDateTime_AddSeconds(z.ckObj, C.int(arg1))


    return v != 0
}


// Loads the date/time with a string having the format as produced by the Serialize
// method, which is a string of SPACE separated integers containing (in this order)
// year, month, day, hour, minutes, seconds, and a UTC flag having the value of
// 1/0.
func (z *CkDateTime) DeSerialize(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkDateTime_DeSerialize(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Returns the difference in seconds between the dateTimeArg and this date/time. The value
// returned is this object's date/time - dateTimeArg's date/time. For example, if the
// returned value is positive, then this object's date/time is more recent than
// dateTimeArg's date/time. If the return value is negative, then this object's date/time
// is older than dateTimeArg's date/time.
func (z *CkDateTime) DiffSeconds(arg1 *CkDateTime) int {

    v := C.CkDateTime_DiffSeconds(z.ckObj, arg1.ckObj)


    return int(v)
}


// Returns true if the date/time is within n seconds/minutes/hours/days of the
// current system date/time. Otherwise returns false. The units can be "seconds",
// "minutes", "hours", or "days" (plural or singular).
func (z *CkDateTime) ExpiresWithin(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkDateTime_ExpiresWithin(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns the date/time as a 64-bit integer .NET DateTime value.
// 
// bLocal indicates whether a local or UTC time is returned.
// 
// This is a date and time expressed in the number of 100-nanosecond intervals that
// have elapsed since January 1, 0001 at 00:00:00.000 in the Gregorian calendar.
// 
// The DateTime value type represents dates and times with values ranging from
// 12:00:00 midnight, January 1, 0001 Anno Domini (Common Era) through 11:59:59
// P.M., December 31, 9999 A.D. (C.E.).
// 
// Time values are measured in 100-nanosecond units called ticks, and a particular
// date is the number of ticks since 12:00 midnight, January 1, 0001 A.D. (C.E.) in
// the GregorianCalendar calendar (excluding ticks that would be added by leap
// seconds). For example, a ticks value of 31241376000000000L represents the date,
// Friday, January 01, 0100 12:00:00 midnight. A DateTime value is always expressed
// in the context of an explicit or default calendar.
//
func (z *CkDateTime) GetAsDateTimeTicks(arg1 bool) int64 {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkDateTime_GetAsDateTimeTicks(z.ckObj, b1)


    return int64(v)
}


// Returns the date/time in a compatible ISO 8601 format according to the format
// specified in formatStr.. Examples of ISO 8601 formats include the following:
//     YYYY-MM-DD
// 
//     YYYY-MM-DDThh:mmTZD
// 
//     YYYY-MM-DDThh:mm:ssTZD
// For the date portion of these formats, YYYY is a four-digit year representation,
// MM is a two-digit month representation, and DD is a two-digit day
// representation. For the time portion, hh is the hour representation in 24-hour
// notation, mm is the two-digit minute representation, and ss is the two-digit
// second representation. A time designator T separates the date and time portions
// of the string, while a time zone designator TZD specifies a time zone (UTC).
// 
// bLocal indicates whether a local or UTC time is returned.
// 
// Note: The bLocal argument is interpreted as the reverse of what is intended . The
// problem was discovered just after releasing v9.5.0.65. It will be fixed in the
// next version update.
//
// return a string or nil if failed.
func (z *CkDateTime) GetAsIso8601(arg1 string, arg2 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkDateTime_getAsIso8601(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the date/time in a Windows OLE "DATE" format.
// 
// bLocal indicates whether a local or UTC time is returned.
// 
// The OLE automation date format is a floating point value, counting days since
// midnight 30 December 1899. Hours and minutes are represented as fractional days.
//
func (z *CkDateTime) GetAsOleDate(arg1 bool) float64 {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkDateTime_GetAsOleDate(z.ckObj, b1)


    return float64(v)
}


// Returns the date/time as an RFC822 formatted string. (An RFC822 format string is
// what is found in the "Date" header field of an email, such as "Wed, 18 Oct 2017
// 09:08:21 GMT".)
// 
// bLocal indicates whether a local or UTC time is returned.
//
// return a string or nil if failed.
func (z *CkDateTime) GetAsRfc822(arg1 bool) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    retStr := C.CkDateTime_getAsRfc822(z.ckObj, b1)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the date/time as an RFC 3339 formatted string, such as
// "1990-12-31T23:59:60Z". (This is an ISO 8061 format like the following:
// YYYY-MM-DDThh:mm:ssTZD)
// 
// bLocal indicates whether a local or UTC time is returned.
//
// return a string or nil if failed.
func (z *CkDateTime) GetAsTimestamp(arg1 bool) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    retStr := C.CkDateTime_getAsTimestamp(z.ckObj, b1)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the date/time as a 32-bit Unix time.
// 
// bLocal indicates whether the date/time returned is local or UTC.
// 
// Note: With this format, there is a Y2038 problem that pertains to 32-bit signed
// integers. There are approx 31.5 million seconds per year. The Unix time is
// number of seconds since the Epoch, 1970-01-01 00:00:00 +0000 (UTC). In 2012,
// it's 42 years since 1/1/1970, so the number of seconds is approx 1.3 billion. A
// 32-bit signed integer ranges from -2,147,483,648 to 2,147,483,647 Therefore, if
// a 32-bit signed integer is used, it turns negative in 2038.
// 
// The GetAsUnixTime64 and GetAsUnixTimeDbl methods are provided as solutions to
// the Y2038 problem.
// 
// (Note: The ActiveX Chilkat implementation omits methods that use 64-bit integers
// because there is no means for passing or returning 64-bit integers in ActiveX.)
//
func (z *CkDateTime) GetAsUnixTime(arg1 bool) uint {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkDateTime_GetAsUnixTime(z.ckObj, b1)


    return uint(v)
}


// The same as GetUnixTime, except returns the date/time as a 64-bit integer.
// 
// bLocal indicates whether a local or UTC time is returned.
//
func (z *CkDateTime) GetAsUnixTime64(arg1 bool) int64 {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkDateTime_GetAsUnixTime64(z.ckObj, b1)


    return int64(v)
}


// Returns the time in Unix format (in seconds since the epoch: 00:00:00 UTC on 1
// January 1970).
// 
// bLocal indicates whether the date/time returned is local or UTC.
//
// return a string or nil if failed.
func (z *CkDateTime) GetAsUnixTimeStr(arg1 bool) *string {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    retStr := C.CkDateTime_getAsUnixTimeStr(z.ckObj, b1)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Gets the date/time as a Chilkat "Dt" object.
func (z *CkDateTime) GetDtObj(arg1 bool) *DtObj {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    retObj := C.CkDateTime_GetDtObj(z.ckObj, b1)


    if retObj == nil {
            return nil
            }
    return &DtObj{retObj}
}


// Loads the date/time from a completed asynchronous task.
func (z *CkDateTime) LoadTaskResult(arg1 *Task) bool {

    v := C.CkDateTime_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns true if the date/time is older than the current system date/time by
// n seconds/minutes/hours/days. Otherwise returns false. The units can be
// "seconds", "minutes", "hours", or "days" (plural or singular).
func (z *CkDateTime) OlderThan(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkDateTime_OlderThan(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Serializes the date/time to a us-ascii string that can be imported at a later
// time via the DeSerialize method. The format of the string returned by this
// method is not intended to match any published standard. It is formatted to a
// string with SPACE separated integers containing (in this order) year, month,
// day, hour, minutes, seconds, and a UTC flag having the value of 1 or 0.
// return a string or nil if failed.
func (z *CkDateTime) Serialize() *string {

    retStr := C.CkDateTime_serialize(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sets the date/time from the current system time.
func (z *CkDateTime) SetFromCurrentSystemTime() bool {

    v := C.CkDateTime_SetFromCurrentSystemTime(z.ckObj)


    return v != 0
}


// Sets the date/time from a .NET DateTime value represented in ticks. See
// GetAsDateTimeTicks for more information.
// 
// bLocal indicates whether the passed in date/time is local or UTC.
//
func (z *CkDateTime) SetFromDateTimeTicks(arg1 bool, arg2 int64) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkDateTime_SetFromDateTimeTicks(z.ckObj, b1, C.__int64(arg2))


    return v != 0
}


// Sets the date/time from a Chilkat "Dt" object.
func (z *CkDateTime) SetFromDtObj(arg1 *DtObj) bool {

    v := C.CkDateTime_SetFromDtObj(z.ckObj, arg1.ckObj)


    return v != 0
}


// Sets the date/time from a 32-bit NTP time value. ntpSeconds is the number of seconds
// since 00:00 (midnight) 1 January 1900 GMT.
func (z *CkDateTime) SetFromNtpTime(arg1 int) bool {

    v := C.CkDateTime_SetFromNtpTime(z.ckObj, C.int(arg1))


    return v != 0
}


// Sets the date/time from a Windows OLE "DATE" value.
// 
// bLocal indicates whether the passed in date/time is local or UTC.
// 
// Note: This method was not working correctly. The problem was discovered just
// after releasing v9.5.0.65. It will be fixed in the next version update.
//
func (z *CkDateTime) SetFromOleDate(arg1 bool, arg2 float64) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkDateTime_SetFromOleDate(z.ckObj, b1, C.double(arg2))


    return v != 0
}


// Sets the date/time from an RFC822 date/time formatted string.
func (z *CkDateTime) SetFromRfc822(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDateTime_SetFromRfc822(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the date/time from an RFC 3339 timestamp format. (such as
// "1990-12-31T23:59:60Z:")
// 
// (This is an ISO 8061 format like the following: YYYY-MM-DDThh:mm:ssTZD)
// 
// Note: Starting in v9.5.0.77, strings formatted as "YYMMDDhhmmssZ", such as
// "181221132225Z", can also be passed to this method.
//
func (z *CkDateTime) SetFromTimestamp(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkDateTime_SetFromTimestamp(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the date/time from a 32-bit UNIX time value. (See GetAsUnixTime for
// information about the Y2038 problem.)
// 
// bLocal indicates whether the passed in date/time is local or UTC.
//
func (z *CkDateTime) SetFromUnixTime(arg1 bool, arg2 uint) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkDateTime_SetFromUnixTime(z.ckObj, b1, C.ulong(arg2))


    return v != 0
}


// The same as SetFromUnixTime, except that it uses a 64-bit integer to solve the
// Y2038 problem. (See GetAsUnixTime for more information about Y2038).
// 
// bLocal indicates whether the passed in date/time is local or UTC.
//
func (z *CkDateTime) SetFromUnixTime64(arg1 bool, arg2 int64) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }

    v := C.CkDateTime_SetFromUnixTime64(z.ckObj, b1, C.__int64(arg2))


    return v != 0
}


