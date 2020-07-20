// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkNtlm.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewNtlm() *Ntlm {
	return &Ntlm{C.CkNtlm_Create()}
}

func (z *Ntlm) DisposeNtlm() {
    if z != nil {
	C.CkNtlm_Dispose(z.ckObj)
	}
}




// property: ClientChallenge
// The ClientChallenge is passed in the Type 3 message from the client to the
// server. It must contain exactly 8 bytes. Because this is a string property, the
// bytes are get/set in encoded form (such as hex or base64) based on the value of
// the EncodingMode property. For example, if the EncodingMode property = "hex",
// then a hex representation of 8 bytes should be used to set the ClientChallenge.
// 
// Note: Setting the ClientChallenge is optional. If the ClientChallenge remains
// unset, it will be automatically set to 8 random bytes when the GenType3 method
// is called.
//
func (z *Ntlm) ClientChallenge() string {
    return C.GoString(C.CkNtlm_clientChallenge(z.ckObj))
}
// property setter: ClientChallenge
func (z *Ntlm) SetClientChallenge(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putClientChallenge(z.ckObj,cStr)
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
func (z *Ntlm) DebugLogFilePath() string {
    return C.GoString(C.CkNtlm_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Ntlm) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DnsComputerName
// Optional. This is information that would be set by the server for inclusion in
// the "Target Info" internal portion of the Type 2 message. Note: If any optional
// "Target Info" fields are provided, then both NetBiosComputerName and
// NetBiosDomainName must be provided.
func (z *Ntlm) DnsComputerName() string {
    return C.GoString(C.CkNtlm_dnsComputerName(z.ckObj))
}
// property setter: DnsComputerName
func (z *Ntlm) SetDnsComputerName(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putDnsComputerName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: DnsDomainName
// Optional. This is information that would be set by the server for inclusion in
// the "Target Info" internal portion of the Type 2 message. Note: If any optional
// "Target Info" fields are provided, then both NetBiosComputerName and
// NetBiosDomainName must be provided.
func (z *Ntlm) DnsDomainName() string {
    return C.GoString(C.CkNtlm_dnsDomainName(z.ckObj))
}
// property setter: DnsDomainName
func (z *Ntlm) SetDnsDomainName(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putDnsDomainName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Domain
// Optional. May be set by the client for inclusion in the Type 1 message.
func (z *Ntlm) Domain() string {
    return C.GoString(C.CkNtlm_domain(z.ckObj))
}
// property setter: Domain
func (z *Ntlm) SetDomain(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putDomain(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EncodingMode
// Determines the encoding mode used for getting/setting various properties, such
// as ClientChallenge. The valid case-insensitive modes are "Base64", "modBase64",
// "Base32", "UU", "QP" (for quoted-printable), "URL" (for url-encoding), "Hex",
// "Q", "B", "url_oath", "url_rfc1738", "url_rfc2396", and "url_rfc3986".
func (z *Ntlm) EncodingMode() string {
    return C.GoString(C.CkNtlm_encodingMode(z.ckObj))
}
// property setter: EncodingMode
func (z *Ntlm) SetEncodingMode(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putEncodingMode(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Flags
// The negotiate flags that are set in the Type 1 message generated by the client
// and sent to the server. These flags have a default value and should ONLY be set
// by a programmer that is an expert in the NTLM protocol and knows what they mean.
// In general, this property should be left at it's default value.
// 
// The flags are represented as a string of letters, where each letter represents a
// bit. The full set of possible flags (bit values) are shown below:
// NegotiateUnicode               0x00000001
// NegotiateOEM                   0x00000002
// RequestTarget                  0x00000004
// NegotiateSign                  0x00000010
// NegotiateSeal                  0x00000020
// NegotiateDatagramStyle         0x00000040
// NegotiateLanManagerKey         0x00000080
// NegotiateNetware               0x00000100
// NegotiateNTLMKey               0x00000200
// NegotiateDomainSupplied        0x00001000
// NegotiateWorkstationSupplied   0x00002000
// NegotiateLocalCall             0x00004000
// NegotiateAlwaysSign            0x00008000
// TargetTypeDomain               0x00010000
// TargetTypeServer               0x00020000
// TargetTypeShare                0x00040000
// NegotiateNTLM2Key              0x00080000
// RequestInitResponse            0x00100000
// RequestAcceptResponse          0x00200000
// RequestNonNTSessionKey         0x00400000
// NegotiateTargetInfo            0x00800000
// Negotiate128                   0x20000000
// NegotiateKeyExchange           0x40000000
// Negotiate56                    0x80000000
// The mapping of letters to bit values are as follows:
// 0x01 - "A"
// 0x02 - "B"
// 0x04 - "C"
// 0x10 - "D"
// 0x20 - "E"
// 0x40 - "F"
// 0x80 - "G"
// 0x200 - "H"
// 0x400 - "I"
// 0x800 - "J"
// 0x1000 - "K"
// 0x2000 - "L"
// 0x8000 - "M"
// 0x10000 - "N"
// 0x20000 - "O"
// 0x40000 - "P"
// 0x80000 - "Q"
// 0x100000 - "R"
// 0x400000 - "S"
// 0x800000 - "T"
// 0x2000000 - "U"
// 0x20000000 - "V"
// 0x40000000 - "W"
// 0x80000000 - "X"
// The default Flags value has the following flags set: NegotiateUnicode,
// NegotiateOEM, RequestTarget, NegotiateNTLMKey, NegotiateAlwaysSign,
// NegotiateNTLM2Key. The corresponds to the string "ABCHMQ".
//
func (z *Ntlm) Flags() string {
    return C.GoString(C.CkNtlm_flags(z.ckObj))
}
// property setter: Flags
func (z *Ntlm) SetFlags(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putFlags(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Ntlm) LastErrorHtml() string {
    return C.GoString(C.CkNtlm_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Ntlm) LastErrorText() string {
    return C.GoString(C.CkNtlm_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Ntlm) LastErrorXml() string {
    return C.GoString(C.CkNtlm_lastErrorXml(z.ckObj))
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
func (z *Ntlm) LastMethodSuccess() bool {
    v := int(C.CkNtlm_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Ntlm) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkNtlm_putLastMethodSuccess(z.ckObj,v)
}

// property: NetBiosComputerName
// Optional. This is information that would be set by the server for inclusion in
// the "Target Info" internal portion of the Type 2 message. Note: If any optional
// "Target Info" fields are provided, then both NetBiosComputerName and
// NetBiosDomainName must be provided.
func (z *Ntlm) NetBiosComputerName() string {
    return C.GoString(C.CkNtlm_netBiosComputerName(z.ckObj))
}
// property setter: NetBiosComputerName
func (z *Ntlm) SetNetBiosComputerName(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putNetBiosComputerName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: NetBiosDomainName
// Optional. This is information that would be set by the server for inclusion in
// the "Target Info" internal portion of the Type 2 message. Note: If any optional
// "Target Info" fields are provided, then both NetBiosComputerName and
// NetBiosDomainName must be provided.
func (z *Ntlm) NetBiosDomainName() string {
    return C.GoString(C.CkNtlm_netBiosDomainName(z.ckObj))
}
// property setter: NetBiosDomainName
func (z *Ntlm) SetNetBiosDomainName(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putNetBiosDomainName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: NtlmVersion
// The version of the NTLM protocol to be used. Must be set to either 1 or 2. The
// default value is 1 (for NTLMv1). Setting this property equal to 2 selects
// NTLMv2.
func (z *Ntlm) NtlmVersion() int {
    return int(C.CkNtlm_getNtlmVersion(z.ckObj))
}
// property setter: NtlmVersion
func (z *Ntlm) SetNtlmVersion(value int) {
    C.CkNtlm_putNtlmVersion(z.ckObj,C.int(value))
}

// property: OemCodePage
// If the "A" flag is unset, then Unicode strings are not used internally in the
// NTLM messages. Strings are instead represented using the OEM code page (i.e.
// charset, or character encoding) as specified here. In general, given that the
// Flags property should rarely be modified, and given that the "A" flag is set by
// default (meaning that Unicode is used), the OemCodePage property will not apply.
// The default value is the default OEM code page of the local computer.
func (z *Ntlm) OemCodePage() int {
    return int(C.CkNtlm_getOemCodePage(z.ckObj))
}
// property setter: OemCodePage
func (z *Ntlm) SetOemCodePage(value int) {
    C.CkNtlm_putOemCodePage(z.ckObj,C.int(value))
}

// property: Password
// The password corresponding to the username of the account to be authenticated.
// This must be set by the client prior to generating and sending the Type 3
// message.
func (z *Ntlm) Password() string {
    return C.GoString(C.CkNtlm_password(z.ckObj))
}
// property setter: Password
func (z *Ntlm) SetPassword(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putPassword(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ServerChallenge
// This is similar to the ClientChallenge in that it must contain 8 bytes.
// 
// The ServerChallenge is passed in the Type 2 message from the server to the
// client. Because this is a string property, the bytes are get/set in encoded form
// (such as hex or base64) based on the value of the EncodingMode property. For
// example, if the EncodingMode property = "hex", then a hex representation of 8
// bytes should be used to set the ServerChallenge.
// 
// Note: Setting the ServerChallenge is optional. If the ServerChallenge remains
// unset, it will be automatically set to 8 random bytes when the GenType2 method
// is called.
//
func (z *Ntlm) ServerChallenge() string {
    return C.GoString(C.CkNtlm_serverChallenge(z.ckObj))
}
// property setter: ServerChallenge
func (z *Ntlm) SetServerChallenge(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putServerChallenge(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: TargetName
// The authentication realm in which the authenticating account has membership,
// such as a domain for domain accounts, or a server name for local machine
// accounts. The TargetName is used in the Type2 message sent from the server to
// the client.
func (z *Ntlm) TargetName() string {
    return C.GoString(C.CkNtlm_targetName(z.ckObj))
}
// property setter: TargetName
func (z *Ntlm) SetTargetName(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putTargetName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UserName
// The username of the account to be authenticated. This must be set by the client
// prior to generating and sending the Type 3 message.
func (z *Ntlm) UserName() string {
    return C.GoString(C.CkNtlm_userName(z.ckObj))
}
// property setter: UserName
func (z *Ntlm) SetUserName(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putUserName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Ntlm) VerboseLogging() bool {
    v := int(C.CkNtlm_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Ntlm) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkNtlm_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Ntlm) Version() string {
    return C.GoString(C.CkNtlm_version(z.ckObj))
}

// property: Workstation
// The value to be used in the optional workstation field in Type 1 message.
func (z *Ntlm) Workstation() string {
    return C.GoString(C.CkNtlm_workstation(z.ckObj))
}
// property setter: Workstation
func (z *Ntlm) SetWorkstation(goStr string) {
    cStr := C.CString(goStr)
    C.CkNtlm_putWorkstation(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}
// Compares the internal contents of two Type3 messages to verify that the LM and
// NTLM response parts match. A server would typically compute the Type3 message by
// calling GenType3, and then compare it with the Type3 message received from the
// client. The method returns true if the responses match, and false if they do
// not.
func (z *Ntlm) CompareType3(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkNtlm_CompareType3(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Generates the Type 1 message. The Type 1 message is sent from Client to Server
// and initiates the NTLM authentication exchange.
// return a string or nil if failed.
func (z *Ntlm) GenType1() *string {

    retStr := C.CkNtlm_genType1(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Generates a Type2 message from a received Type1 message. The server-side
// generates the Type2 message and sends it to the client. This is the 2nd step in
// the NTLM protocol. The 1st step is the client generating the initial Type1
// message which is sent to the server.
// return a string or nil if failed.
func (z *Ntlm) GenType2(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkNtlm_genType2(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Generates the final message in the NTLM authentication exchange. This message is
// sent from the client to the server. The Type 2 message received from the server
// is passed to GenType3. The Username and Password properties are finally used
// here in the generation of the Type 3 message. Note, the Password is never
// actually sent. It is used to compute a binary response that the server can then
// check, using the password it has on file, to verify that indeed the client
// must've used the correct password.
// return a string or nil if failed.
func (z *Ntlm) GenType3(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkNtlm_genType3(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// The server-side should call this method with the Type 3 message received from
// the client. The LoadType3 method sets the following properties: Username,
// Domain, Workstation, and ClientChallenge, all of which are embedded within the
// Type 3 message.
// 
// The server-side code may then use the Username to lookup the associated password
// and then it will itself call the GenType3 method to do the same computation as
// done by the client. The server then compares it's computed Type 3 message with
// the Type 3 message received from the client. If the Type 3 messages are exactly
// the same, then it must be that the client used the correct password, and
// therefore the client authentication is successful.
//
func (z *Ntlm) LoadType3(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkNtlm_LoadType3(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// For informational purposes only. Allows for the server-side to parse a Type 1
// message to get human-readable information about the contents.
// return a string or nil if failed.
func (z *Ntlm) ParseType1(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkNtlm_parseType1(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// For informational purposes only. Allows for the client-side to parse a Type 2
// message to get human-readable information about the contents.
// return a string or nil if failed.
func (z *Ntlm) ParseType2(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkNtlm_parseType2(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// For informational purposes only. Allows for the server-side to parse a Type 3
// message to get human-readable information about the contents.
// return a string or nil if failed.
func (z *Ntlm) ParseType3(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkNtlm_parseType3(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Sets one of the negotiate flags to be used in the Type 1 message sent by the
// client. It should normally be unnecessary to modify the default flag settings.
// For more information about flags, see the description for the Flags property
// above.
func (z *Ntlm) SetFlag(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkNtlm_SetFlag(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Unlocks the component. This must be called once prior to calling any other
// method.
func (z *Ntlm) UnlockComponent(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkNtlm_UnlockComponent(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


