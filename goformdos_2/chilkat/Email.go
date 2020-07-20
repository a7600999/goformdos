// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkEmail.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewEmail() *Email {
	return &Email{C.CkEmail_Create()}
}

func (z *Email) DisposeEmail() {
    if z != nil {
	C.CkEmail_Dispose(z.ckObj)
	}
}




// property: Body
// The body of the email. If the email has both HTML and plain-text bodies, this
// property returns the HTML body. The GetHtmlBody and GetPlainTextBody methods can
// be used to access a specific body. The HasHtmlBody and HasPlainTextBody methods
// can be used to determine the presence of a body.
func (z *Email) Body() string {
    return C.GoString(C.CkEmail_body(z.ckObj))
}
// property setter: Body
func (z *Email) SetBody(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putBody(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: BounceAddress
// The "return-path" address of the email to be used when the email is sent.
// Bounces (i.e. delivery status notifications, or DSN's) will go to this address.
// 
// Note: This is not the content of the "return-path" header for emails that are
// downloaded from a POP3 or IMAP server. The BounceAddress is the email address to
// be used in the process of sending the email via SMTP. (See the "SMTP Protocol in
// a Nutshell" link below.) The BounceAddress is the email address passed in the
// "MAIL FROM" SMTP command which becomes the "return-path" header in the email
// when received.
// 
// Note: The Sender and BounceAddress properties are identical and perform the same
// function. Setting the Sender property also sets the BounceAddress property, and
// vice-versa. The reason for the duplication is that BounceAddress existed first,
// and developers typically searched for a "Sender" property without realizing that
// the BounceAddress property served this function.
//
func (z *Email) BounceAddress() string {
    return C.GoString(C.CkEmail_bounceAddress(z.ckObj))
}
// property setter: BounceAddress
func (z *Email) SetBounceAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putBounceAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Charset
// Sets the charset for the entire email. The header fields and plain-text/HTML
// bodies will be converted and sent in this charset. (This includes parsing and
// updating the HTML with the appropriate META tag specifying the charset.) All
// formatting and encoding of the email MIME is handled automatically by the
// Chilkat Mail component. If your application wants to send a Shift_JIS email, you
// simply set the Charset property to "Shift_JIS". Note: If a charset property is
// not explicitly set, the Chilkat component automatically detects the charset and
// chooses the appropriate charset. If all characters are 7bit (i.e. us-ascii) the
// charset is left blank. If the email contain a mix of languages such that no one
// charset can be chosen, or if the language cannot be determined without
// ambiguity, then the "utf-8" charset will be chosen.
func (z *Email) Charset() string {
    return C.GoString(C.CkEmail_charset(z.ckObj))
}
// property setter: Charset
func (z *Email) SetCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putCharset(z.ckObj,cStr)
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
func (z *Email) DebugLogFilePath() string {
    return C.GoString(C.CkEmail_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Email) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Decrypted
// true if the email arrived encrypted and was successfully decrypted, otherwise
// false. This property is only meaningful when the ReceivedEncrypted property is
// equal to true.
func (z *Email) Decrypted() bool {
    v := int(C.CkEmail_getDecrypted(z.ckObj))
    return v != 0
}

// property: EmailDateStr
// The date/time from the "Date" header in the UTC/GMT timezone in RFC822 string
// form.
func (z *Email) EmailDateStr() string {
    return C.GoString(C.CkEmail_emailDateStr(z.ckObj))
}
// property setter: EmailDateStr
func (z *Email) SetEmailDateStr(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putEmailDateStr(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: EncryptedBy
// If the email was received encrypted, this contains the details of the
// certificate used for encryption.
func (z *Email) EncryptedBy() string {
    return C.GoString(C.CkEmail_encryptedBy(z.ckObj))
}

// property: FileDistList
// Set this property to send an email to a list of recipients stored in a plain
// text file. The file format is simple: one recipient per line, no comments
// allowed, blank lines are ignored.Setting this property is equivalent to adding a
// "CKX-FileDistList"header field to the email. Chilkat Mail treats header fields
// beginning with "CKX-"specially in that these fields are never transmitted with
// the email when sent. However, CKX fields are saved and restored when saving to
// XML or loading from XML (or MIME). When sending an email containing a
// "CKX-FileDistList"header field, Chilkat Mail will read the distribution list
// file and send the email to each recipient. Emails can be sent individually, or
// with BCC, 100 recipients at a time. (see the MailMan.SendIndividual property).
func (z *Email) FileDistList() string {
    return C.GoString(C.CkEmail_fileDistList(z.ckObj))
}
// property setter: FileDistList
func (z *Email) SetFileDistList(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putFileDistList(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: From
// The combined name and email address of the sender, such as "John Smith" . This
// is the content that will be placed in the From: header field. If the actual
// sender is to be different, then set the Sender property to a different email
// address.
func (z *Email) From() string {
    return C.GoString(C.CkEmail_from(z.ckObj))
}
// property setter: From
func (z *Email) SetFrom(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putFrom(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: FromAddress
// The email address of the sender.
func (z *Email) FromAddress() string {
    return C.GoString(C.CkEmail_fromAddress(z.ckObj))
}
// property setter: FromAddress
func (z *Email) SetFromAddress(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putFromAddress(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: FromName
// The name of the sender.
func (z *Email) FromName() string {
    return C.GoString(C.CkEmail_fromName(z.ckObj))
}
// property setter: FromName
func (z *Email) SetFromName(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putFromName(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Header
// The complete MIME header of the email.
func (z *Email) Header() string {
    return C.GoString(C.CkEmail_header(z.ckObj))
}

// property: Language
// A read-only property that identifies the primary language group for the email.
// Possible values are:
// 
//         "latin1" (for English and all Western European languages)
//         "central" (for Central European languages such as Polish, Czech,
//         Hungarian, etc.)
//         "russian" (for Cyrillic languages)
//         "greek"
//         "turkish"
//         "hebrew"
//         "arabic"
//         "thai"
//         "vietnamese"
//         "chinese"
//         "japanese"
//         "korean"
//         "devanagari"
//         "bengali"
//         "gurmukhi"
//         "gujarati"
//         "oriya"
//         "tamil"
//         "telugu"
//         "kannada"
//         "malayalam"
//         "sinhala"
//         "lao"
//         "tibetan"
//         "myanmar"
//         "georgian"
//         "unknown"
// 
// The language group determination is made soley on the subject and
// plain-text/HTML email bodies. Characters in the FROM, TO, CC, and other header
// fields are not considered.
// 
// The primary determining factor is the characters found in the Subject header
// field. For example, if an email contains Japanese in the Subject, but the body
// contains Russian characters, it will be considered "japanese".
// 
// The language is determined by where the Unicode chars fall in various blocks in
// the Unicode Basic Multilingual Plane. For more information, see the book
// "Unicode Demystified" by Richard Gillam.
//
func (z *Email) Language() string {
    return C.GoString(C.CkEmail_language(z.ckObj))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Email) LastErrorHtml() string {
    return C.GoString(C.CkEmail_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Email) LastErrorText() string {
    return C.GoString(C.CkEmail_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Email) LastErrorXml() string {
    return C.GoString(C.CkEmail_lastErrorXml(z.ckObj))
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
func (z *Email) LastMethodSuccess() bool {
    v := int(C.CkEmail_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Email) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmail_putLastMethodSuccess(z.ckObj,v)
}

// property: LocalDateStr
// The date/time found in the "Date" header field returned in the local timezone in
// RFC822 string form.
func (z *Email) LocalDateStr() string {
    return C.GoString(C.CkEmail_localDateStr(z.ckObj))
}
// property setter: LocalDateStr
func (z *Email) SetLocalDateStr(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putLocalDateStr(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Mailer
// Identifies the email software that sent the email.
func (z *Email) Mailer() string {
    return C.GoString(C.CkEmail_mailer(z.ckObj))
}
// property setter: Mailer
func (z *Email) SetMailer(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putMailer(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: NumAlternatives
// The number of alternative bodies present in the email. An email that is not
// "multipart/alternative"will return 0 alternatives. An email that is
// "multipart/alternative" will return a number greater than or equal to 1.
func (z *Email) NumAlternatives() int {
    return int(C.CkEmail_getNumAlternatives(z.ckObj))
}

// property: NumAttachedMessages
// Returns the number of embedded emails. Some mail clients will embed an email
// that is to be forwarded into a new email as a "message/rfc822" subpart of the
// MIME message structure. This property tells how many emails have been embedded.
// The original email can be retrieved by calling GetAttachedMessage.
func (z *Email) NumAttachedMessages() int {
    return int(C.CkEmail_getNumAttachedMessages(z.ckObj))
}

// property: NumAttachments
// The number of attachments contained in the email.
// 
// Note: If an email is downloaded from an IMAP server without attachments, then
// the number of attachments should be obtained by calling the IMAP object's
// GetMailNumAttach method. This property indicates the actual number of
// attachments already present within the email object.
//
func (z *Email) NumAttachments() int {
    return int(C.CkEmail_getNumAttachments(z.ckObj))
}

// property: NumBcc
// The number of blind carbon-copy email recipients.
func (z *Email) NumBcc() int {
    return int(C.CkEmail_getNumBcc(z.ckObj))
}

// property: NumCC
// The number of carbon-copy email recipients.
func (z *Email) NumCC() int {
    return int(C.CkEmail_getNumCC(z.ckObj))
}

// property: NumDaysOld
// Returns the number of days old from the current system date/time. The email's
// date is obtained from the "Date" header field. If the Date header field is
// missing, or invalid, then -9999 is returned. A negative number may be returned
// if the Date header field contains a future date/time. (However, -9999 represents
// an error condition.)
func (z *Email) NumDaysOld() int {
    return int(C.CkEmail_getNumDaysOld(z.ckObj))
}

// property: NumDigests
// Returns the number of message/rfc822 parts contained within the multipart/digest
// enclosure. If no multipart/digest enclosure exists, then this property has the
// value of 0. The GetDigest method is called to get the Nth digest as an email
// object.
func (z *Email) NumDigests() int {
    return int(C.CkEmail_getNumDigests(z.ckObj))
}

// property: NumHeaderFields
// The number of header fields. When accessing a header field by index, the 1st
// header field is at index 0, and the last is at NumHeaderFields-1. (Chilkat
// indexing is always 0-based.)
func (z *Email) NumHeaderFields() int {
    return int(C.CkEmail_getNumHeaderFields(z.ckObj))
}

// property: NumRelatedItems
// The number of related items present in this email. Related items are typically
// image files (JPEGs or GIFs) or style sheets (CSS files) that are included with
// HTML formatted messages with internal "CID"hyperlinks.
func (z *Email) NumRelatedItems() int {
    return int(C.CkEmail_getNumRelatedItems(z.ckObj))
}

// property: NumReplacePatterns
// Returns the number of replacement patterns previously set by calling the
// SetReplacePattern method 1 or more times. If replacement patterns are set, the
// email bodies and header fields are modified by applying the search/replacement
// strings during the message sending process.
func (z *Email) NumReplacePatterns() int {
    return int(C.CkEmail_getNumReplacePatterns(z.ckObj))
}

// property: NumReports
// (For multipart/report emails that have sub-parts with Content-Types such as
// message/feedback-report.) Any MIME sub-part within the email that has a
// Content-Type of "message/*", but is not a "message/rfc822", is considered to be
// a "report" and is included in this count. (A "message/rfc822" is considered an
// attached message and is handled by the NumAttachedMessages property and the
// GetAttachedMessage method.) Any MIME sub-part having a Content-Type equal to
// "text/rfc822-headers" is also considered to be a "report". The GetReport method
// may be called to get the body content of each "report" contained within a
// multipart/report email.
func (z *Email) NumReports() int {
    return int(C.CkEmail_getNumReports(z.ckObj))
}

// property: NumTo
// The number of direct email recipients.
func (z *Email) NumTo() int {
    return int(C.CkEmail_getNumTo(z.ckObj))
}

// property: OaepHash
// Selects the hash algorithm for use within OAEP padding when encrypting email
// using RSAES-OAEP. The valid choices are "sha1", "sha256", "sha384", "sha512",
func (z *Email) OaepHash() string {
    return C.GoString(C.CkEmail_oaepHash(z.ckObj))
}
// property setter: OaepHash
func (z *Email) SetOaepHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putOaepHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OaepMgfHash
// Selects the MGF hash algorithm for use within OAEP padding when encrypting email
// using RSAES-OAEP. The valid choices are "sha1", "sha256", "sha384", "sha512",
func (z *Email) OaepMgfHash() string {
    return C.GoString(C.CkEmail_oaepMgfHash(z.ckObj))
}
// property setter: OaepMgfHash
func (z *Email) SetOaepMgfHash(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putOaepMgfHash(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: OaepPadding
// Selects the RSA encryption scheme when encrypting email. The default value is
// false, which selects RSAES_PKCS1-V1_5. If set to true, then RSAES_OAEP is
// used.
func (z *Email) OaepPadding() bool {
    v := int(C.CkEmail_getOaepPadding(z.ckObj))
    return v != 0
}
// property setter: OaepPadding
func (z *Email) SetOaepPadding(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmail_putOaepPadding(z.ckObj,v)
}

// property: OverwriteExisting
// When true (the default) the methods to save email attachments and related
// items will overwrite files if they already exist. If false, then the methods
// that save email attachments and related items will append a string of 4
// characters to create a unique filename if a file already exists. The filename of
// the attachment (or related item) within the email object is updated and can be
// retrieved by the program to determine the actual file(s) created.
func (z *Email) OverwriteExisting() bool {
    v := int(C.CkEmail_getOverwriteExisting(z.ckObj))
    return v != 0
}
// property setter: OverwriteExisting
func (z *Email) SetOverwriteExisting(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmail_putOverwriteExisting(z.ckObj,v)
}

// property: Pkcs7CryptAlg
// When an email is sent encrypted (using PKCS7 public-key encryption), this
// selects the underlying symmetric encryption algorithm. Possible values are:
// "aes", "des", "3des", and "rc2". The default value is "aes".
func (z *Email) Pkcs7CryptAlg() string {
    return C.GoString(C.CkEmail_pkcs7CryptAlg(z.ckObj))
}
// property setter: Pkcs7CryptAlg
func (z *Email) SetPkcs7CryptAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putPkcs7CryptAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Pkcs7KeyLength
// When the email is sent encrypted (using PKCS7 public-key encryption), this
// selects the key length of the underlying symmetric encryption algorithm. The
// possible values allowed depend on the Pkcs7CryptAlg property. For "aes", the key
// length may be 128, 192, or 256. For "3des" the key length must be 192. For "des"
// the key length must be 40. For "rc2" the key length can be 40, 56, 64, or 128.
func (z *Email) Pkcs7KeyLength() int {
    return int(C.CkEmail_getPkcs7KeyLength(z.ckObj))
}
// property setter: Pkcs7KeyLength
func (z *Email) SetPkcs7KeyLength(value int) {
    C.CkEmail_putPkcs7KeyLength(z.ckObj,C.int(value))
}

// property: PreferredCharset
// Only applies when building an email with non-English characters where the
// charset is not explicitly set. The Chilkat email component will automatically
// choose a charset based on the languages found within an email (if the charset is
// not already specified within the MIME or explicitly specified by setting the
// Charset property). The default charset chosen for each language is:
// 
// Chinese: gb2312
// Japanese: shift_JIS
// Korean: ks_c_5601-1987
// Thai: windows-874
// All others: iso-8859-*
// 
// This allows for charsets such as iso-2022-jp to be chosen instead of the
// default. If the preferred charset does not apply to the situation, it is not
// used. For example, if the preferred charset is iso-2022-jp, but the email
// contains Greek characters, then the preferred charset is ignored.
//
func (z *Email) PreferredCharset() string {
    return C.GoString(C.CkEmail_preferredCharset(z.ckObj))
}
// property setter: PreferredCharset
func (z *Email) SetPreferredCharset(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putPreferredCharset(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: PrependHeaders
// If true, then header fields added via the AddHeaderField or AddHeaderField2
// methods are prepended to the top of the header as opposed to appended to the
// bottom. The default value is false.
func (z *Email) PrependHeaders() bool {
    v := int(C.CkEmail_getPrependHeaders(z.ckObj))
    return v != 0
}
// property setter: PrependHeaders
func (z *Email) SetPrependHeaders(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmail_putPrependHeaders(z.ckObj,v)
}

// property: ReceivedEncrypted
// true if this email was originally received with encryption, otherwise false.
func (z *Email) ReceivedEncrypted() bool {
    v := int(C.CkEmail_getReceivedEncrypted(z.ckObj))
    return v != 0
}

// property: ReceivedSigned
// true if this email was originally received with a digital signature, otherwise
// false.
func (z *Email) ReceivedSigned() bool {
    v := int(C.CkEmail_getReceivedSigned(z.ckObj))
    return v != 0
}

// property: ReplyTo
// Sets the "Reply-To" header field to the specified email address.
func (z *Email) ReplyTo() string {
    return C.GoString(C.CkEmail_replyTo(z.ckObj))
}
// property setter: ReplyTo
func (z *Email) SetReplyTo(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putReplyTo(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: ReturnReceipt
// Set to true if you want the email to request a return-receipt when received by
// the recipient. The default value is false.
func (z *Email) ReturnReceipt() bool {
    v := int(C.CkEmail_getReturnReceipt(z.ckObj))
    return v != 0
}
// property setter: ReturnReceipt
func (z *Email) SetReturnReceipt(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmail_putReturnReceipt(z.ckObj,v)
}

// property: SendEncrypted
// Set to true if this email should be sent encrypted.
func (z *Email) SendEncrypted() bool {
    v := int(C.CkEmail_getSendEncrypted(z.ckObj))
    return v != 0
}
// property setter: SendEncrypted
func (z *Email) SetSendEncrypted(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmail_putSendEncrypted(z.ckObj,v)
}

// property: Sender
// The sender's address for this email message.
// 
// This is the address of the actual sender acting on behalf of the author listed
// in the From: field.
// 
// Note: The Sender and BounceAddress properties are identical and perform the same
// function. Setting the Sender property also sets the BounceAddress property, and
// vice-versa. The reason for the duplication is that BounceAddress existed first,
// and developers typically searched for a "Sender" property without realizing that
// the BounceAddress property served this function.
// 
// Important Note: This property does not contain the value of the "Sender" header
// field, if one exists, for a received email. It is a property that is used when
// sending email. To get the value of the "Sender" header field (which may not
// always exist), call the GetHeaderField method instead.
//
func (z *Email) Sender() string {
    return C.GoString(C.CkEmail_sender(z.ckObj))
}
// property setter: Sender
func (z *Email) SetSender(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putSender(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SendSigned
// Set to true if this email should be sent with a digital signature.
func (z *Email) SendSigned() bool {
    v := int(C.CkEmail_getSendSigned(z.ckObj))
    return v != 0
}
// property setter: SendSigned
func (z *Email) SetSendSigned(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmail_putSendSigned(z.ckObj,v)
}

// property: SignaturesValid
// true if the email was received with one or more digital signatures, and if all
// the signatures were validated indicating that the email was not altered.
// Otherwise this property is set to false. (This property is only meaningful
// when the ReceivedSigned property is equal to true.)
func (z *Email) SignaturesValid() bool {
    v := int(C.CkEmail_getSignaturesValid(z.ckObj))
    return v != 0
}

// property: SignedBy
// If the email was received digitally signed, this property contains the fields of
// the cert's SubjectDN.
// 
// For example: US, 60187, Illinois, Wheaton, 1719 E Forest Ave, "Chilkat Software,
// Inc.", "Chilkat Software, Inc."
// 
// It is like the DN (Distinguished Name), but without the "AttrName=" before each
// attribute.
//
func (z *Email) SignedBy() string {
    return C.GoString(C.CkEmail_signedBy(z.ckObj))
}

// property: SigningAlg
// Selects the signature algorithm to be used when sending signed (PKCS7) email.
// The default value is "PKCS1-v1_5". This can be set to "RSASSA-PSS" (or simply
// "pss") to use the RSASSA-PSS signature scheme.
// 
// Note: This property only applies when signing with an RSA private key. It does
// not apply for ECC or DSA private keys.
//
func (z *Email) SigningAlg() string {
    return C.GoString(C.CkEmail_signingAlg(z.ckObj))
}
// property setter: SigningAlg
func (z *Email) SetSigningAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putSigningAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: SigningHashAlg
// Selects the underlying hash algorithm used when sending signed (PKCS7) email.
// Possible values are "sha1", "sha256", "sha384", "sha512", "md5", and "md2".
func (z *Email) SigningHashAlg() string {
    return C.GoString(C.CkEmail_signingHashAlg(z.ckObj))
}
// property setter: SigningHashAlg
func (z *Email) SetSigningHashAlg(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putSigningHashAlg(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Size
// The size in bytes of the email, including all parts and attachments.
// 
// Note: This property is only valid if the full email was downloaded. If only the
// header was downloaded, then this property will contain the size of just the
// header.
//
func (z *Email) Size() int {
    return int(C.CkEmail_getSize(z.ckObj))
}

// property: Subject
// The email subject.
func (z *Email) Subject() string {
    return C.GoString(C.CkEmail_subject(z.ckObj))
}
// property setter: Subject
func (z *Email) SetSubject(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putSubject(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: Uidl
// This is the unique ID assigned by the POP3 server. Emails can be retrieved or
// deleted from the POP3 server via the UIDL. The header field for this property is
// "X-UIDL".
// 
// Important: Emails downloaded via the IMAP protocol do not have UIDL's. UIDL's
// are specific to the POP3 protocol. IMAP servers use UID's (notice the spelling
// difference -- "UIDL" vs. "UID"). An email downloaded via the Chilkat IMAP
// component will contain a "ckx-imap-uid" header field that contains the UID. This
// may be accessed via the GetHeaderField method.
//
func (z *Email) Uidl() string {
    return C.GoString(C.CkEmail_uidl(z.ckObj))
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. This property
// defaults to the empty string, and should typically remain empty.
// 
// As of v9.5.0.79, the only possible value is:
//     "NO_FORMAT_FLOWED" - Don't automatically add "format=flowed" to any
//     Content-Type header.
//
func (z *Email) UncommonOptions() string {
    return C.GoString(C.CkEmail_uncommonOptions(z.ckObj))
}
// property setter: UncommonOptions
func (z *Email) SetUncommonOptions(goStr string) {
    cStr := C.CString(goStr)
    C.CkEmail_putUncommonOptions(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: UnpackUseRelPaths
// Applies to the UnpackHtml method. If true, then relative paths are used within
// the HTML for the links to the related files (images and style sheets) that were
// unpacked to the filesystem. Otherwise absolute paths are used. The default value
// is true.
func (z *Email) UnpackUseRelPaths() bool {
    v := int(C.CkEmail_getUnpackUseRelPaths(z.ckObj))
    return v != 0
}
// property setter: UnpackUseRelPaths
func (z *Email) SetUnpackUseRelPaths(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmail_putUnpackUseRelPaths(z.ckObj,v)
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Email) VerboseLogging() bool {
    v := int(C.CkEmail_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Email) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkEmail_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Email) Version() string {
    return C.GoString(C.CkEmail_version(z.ckObj))
}
// Adds an attachment using the contents of a BinData object. If contentType is empty,
// then the content-type will be inferred from the filename extension.
func (z *Email) AddAttachmentBd(arg1 string, arg2 *BinData, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)

    v := C.CkEmail_AddAttachmentBd(z.ckObj, cstr1, arg2.ckObj, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds or replaces a MIME header field in one of the email attachments. If the
// header field does not exist, it is added. Otherwise it is replaced.
func (z *Email) AddAttachmentHeader(arg1 int, arg2 string, arg3 string)  {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    C.CkEmail_AddAttachmentHeader(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))


}


// Adds a recipient to the blind carbon-copy list. address is required, but name
// may be empty.
func (z *Email) AddBcc(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddBcc(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a recipient to the carbon-copy list. address is required, but name may be
// empty.
func (z *Email) AddCC(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddCC(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds an attachment directly from data in memory to the email.
func (z *Email) AddDataAttachment(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkEmail_AddDataAttachment(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Adds an attachment to an email from in-memory data. Same as AddDataAttachment
// but allows the content-type to be specified.
func (z *Email) AddDataAttachment2(arg1 string, arg2 []byte, arg3 string) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))
    cstr3 := C.CString(arg3)

    v := C.CkEmail_AddDataAttachment2(z.ckObj, cstr1, ckbyd2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Allows for certificates to be explicitly specified for sending encrypted email
// to one or more recipients. Call this method once per certificate to be used. The
// ClearEncryptCerts method may be called to clear the list of explicitly-specified
// certificates.
// 
// Note: It is possible to send encrypted email without explicitly specifying the
// certificates. On Windows computers, the registry-based Current-User and
// Local-Machine certificate stores are automatically searched for certs matching
// each of the recipients (To, CC, and BCC recipients).
// 
// Note: The SetEncryptCert method is equivalent to calling ClearEncryptCerts
// followed by AddEncryptCert.
//
func (z *Email) AddEncryptCert(arg1 *Cert) bool {

    v := C.CkEmail_AddEncryptCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds a file as an attachment to the email. Returns the MIME content-type of the
// attachment, which is inferred based on the filename extension.
// return a string or nil if failed.
func (z *Email) AddFileAttachment(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkEmail_addFileAttachment(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Same as AddFileAttachment, but the content type can be explicitly specified.
func (z *Email) AddFileAttachment2(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddFileAttachment2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Any standard or non-standard (custom) header field can be added to the email
// with this method. One interesting feature is that all header fields whose name
// begins with "CKX-" will not be included in the header when an email is sent.
// These fields will be included when saved to or loaded from XML. This makes it
// easy to include persistent meta-data with an email which your programs can use
// in any way it chooses.
// 
// Important: This method will replace an already-existing header field. To allow
// for adding duplicate header fields, call AddHeaderField2 (see below).
//
func (z *Email) AddHeaderField(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkEmail_AddHeaderField(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// This method is the same as AddHeaderField, except that if the header field
// already exists, it is not replaced. A duplicate header will be added.
func (z *Email) AddHeaderField2(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkEmail_AddHeaderField2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// Sets the HTML body of the email. Use this method if there will be multiple
// versions of the body, but in different formats, such as HTML and plain text.
// Otherwise, set the body by calling the SetHtmlBody method.
func (z *Email) AddHtmlAlternativeBody(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_AddHtmlAlternativeBody(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds an iCalendar (text/calendar) alternative body to the email. The icalContent
// contains the content of the iCalendar data. A sample is shown here:
// BEGIN:VCALENDAR
// VERSION:2.0
// PRODID:-//hacksw/handcal//NONSGML v1.0//EN
// BEGIN:VEVENT
// UID:uid1@example.com
// DTSTAMP:19970714T170000Z
// ORGANIZER;CN=John Doe:MAILTO:john.doe@example.com
// DTSTART:19970714T170000Z
// DTEND:19970715T035959Z
// SUMMARY:Bastille Day Party
// END:VEVENT
// END:VCALENDAR
// The methodName is the "method" attribute used in the Content-Type header field in the
// alternative body. For example, if set to "REQUEST", then the alternative body's
// header would look like this:
// Content-Type: text/calendar; method=REQUEST
// Content-Transfer-Encoding: base64
func (z *Email) AddiCalendarAlternativeBody(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddiCalendarAlternativeBody(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds multiple recipients to the blind carbon-copy list. The parameter is a
// string containing a comma separated list of full email addresses. Returns True
// if successful.
func (z *Email) AddMultipleBcc(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_AddMultipleBcc(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds multiple recipients to the carbon-copy list. The parameter is a string
// containing a comma separated list of full email addresses. Returns True if
// successful.
func (z *Email) AddMultipleCC(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_AddMultipleCC(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds multiple recipients to the "to" list. The parameter is a string containing
// a comma separated list of full email addresses. Returns True if successful.
func (z *Email) AddMultipleTo(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_AddMultipleTo(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds a PFX to the object's internal list of sources to be searched for
// certificates and private keys when decrypting. Multiple PFX sources can be added
// by calling this method once for each. (On the Windows operating system, the
// registry-based certificate stores are also automatically searched, so it is
// commonly not required to explicitly add PFX sources.)
// 
// The pfxBytes contains the bytes of a PFX file (also known as PKCS12 or .p12).
//
func (z *Email) AddPfxSourceData(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddPfxSourceData(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds a PFX file to the object's internal list of sources to be searched for
// certificates and private keys when decrypting. Multiple PFX files can be added
// by calling this method once for each. (On the Windows operating system, the
// registry-based certificate stores are also automatically searched, so it is
// commonly not required to explicitly add PFX sources.)
// 
// The pfxFilePath contains the bytes of a PFX file (also known as PKCS12 or .p12).
//
func (z *Email) AddPfxSourceFile(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddPfxSourceFile(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the plain-text body of the email. Use this method if there will be multiple
// versions of the body, but in different formats, such as HTML and plain text.
// Otherwise, simply set the Body property.
func (z *Email) AddPlainTextAlternativeBody(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_AddPlainTextAlternativeBody(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds a related item using the contents of a BinData object. Returns the
// Content-ID for the newly added relted item.
// return a string or nil if failed.
func (z *Email) AddRelatedBd(arg1 string, arg2 *BinData) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkEmail_addRelatedBd(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Adds a related item using the contents of a BinData object. The fileNameInHtml should be
// set to the filename/path/url used in the corresponding HTML IMG tag's "src"
// attribute.
func (z *Email) AddRelatedBd2(arg1 *BinData, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddRelatedBd2(z.ckObj, arg1.ckObj, cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds the memory data as a related item to the email and returns the Content-ID.
// Emails formatted in HTML can include images with this call and internally
// reference the image through a "cid"hyperlink. (Chilkat Email.NET fully supports
// the MHTML standard.)
// return a string or nil if failed.
func (z *Email) AddRelatedData(arg1 string, arg2 []byte) *string {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    retStr := C.CkEmail_addRelatedData(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Adds a related item to the email from in-memory byte data. Related items are
// things such as images and style sheets that are embedded within an HTML email.
// They are not considered attachments because their sole purpose is to participate
// in the display of the HTML. This method differs from AddRelatedData in that it
// does not use or return a Content-ID. The filename argument should be set to the
// filename used in the HTML img tag's src attribute (if it's an image), or the URL
// referenced in an HTML link tag for a stylesheet.
func (z *Email) AddRelatedData2(arg1 []byte, arg2 string)  {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    C.CkEmail_AddRelatedData2(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))


}


// Adds the contents of a file to the email and returns the Content-ID. Emails
// formatted in HTML can include images with this call and internally reference the
// image through a "cid" hyperlink. (Chilkat Email.NET fully supports the MHTML
// standard.)
// return a string or nil if failed.
func (z *Email) AddRelatedFile(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkEmail_addRelatedFile(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Adds a related item to the email from a file. Related items are things such as
// images and style sheets that are embedded within an HTML email. They are not
// considered attachments because their sole purpose is to participate in the
// display of the HTML. This method differs from AddRelatedFile in that it does not
// use or return a Content-ID. The filenameInHtml argument should be set to the filename used
// in the HTML img tag's src attribute (if it's an image), or the URL referenced in
// an HTML link tag for a stylesheet. The filenameOnDisk is the path in the local filesystem
// of the file to be added.
// 
// Note: Outlook.com will not properly display embedded HTMl images when the filenameInHtml
// includes a path part. Apparently, Outlook.com is only capable of correctly
// displaying images when the filenameInHtml is a only a filename. Other email clients, such
// as Mozilla Thunderbird, have no trouble when the filenameInHtml includes a path part.
//
func (z *Email) AddRelatedFile2(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddRelatedFile2(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds or replaces a MIME header field in one of the email's related items. If the
// header field does not exist, it is added. Otherwise it is replaced.
func (z *Email) AddRelatedHeader(arg1 int, arg2 string, arg3 string)  {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    C.CkEmail_AddRelatedHeader(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))


}


// Adds a related item to the email. A related item is typically an image or style
// sheet referenced by an HTML tag within the HTML email body. The contents of the
// related item are passed str. nameInHtml specifies the filename that should be used
// within the HTML, and not an actual filename on the local filesystem. charset
// specifies the charset that should be used for the text content of the related
// item. Returns the content-ID generated for the added item.
// return a string or nil if failed.
func (z *Email) AddRelatedString(arg1 string, arg2 string, arg3 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkEmail_addRelatedString(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Adds a related item to the email from an in-memory string. Related items are
// things such as images and style sheets that are embedded within an HTML email.
// They are not considered attachments because their sole purpose is to participate
// in the display of the HTML. The filenameInHtml argument should be set to the
// filename used in the HTML img tag's src attribute (if it's an image), or the URL
// referenced in an HTML link tag for a stylesheet. The charset argument indicates
// that the content should first be converted to the specified charset prior to
// adding to the email. It should hava a value such as "iso-8859-1", "utf-8",
// "Shift_JIS", etc.
func (z *Email) AddRelatedString2(arg1 string, arg2 string, arg3 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    C.CkEmail_AddRelatedString2(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))


}


// Adds an attachment directly from a string in memory to the email.
func (z *Email) AddStringAttachment(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddStringAttachment(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Adds an attachment to an email. The path specifies the filename to be used for
// the attachment and is not an actual filename existing on the local filesystem.
// The content contains the text data for the attachment. The string will be converted
// to the charset specified in charset before being added to the email.
// 
// Note: Beginning in v9.5.0.48, the charset may be prepended with "bom-" or "no-bom-"
// to include or exclude the BOM (preamble) for charsets such as utf-16 or utf-8.
// For example: "no-bom-utf-8" or "bom-utf-8".
//
func (z *Email) AddStringAttachment2(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkEmail_AddStringAttachment2(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Adds a recipient to the "to" list. address is required, but name may be empty.
// Emails that have no "To" recipients will be sent to
// _LT_undisclosed-recipients_GT_.
func (z *Email) AddTo(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkEmail_AddTo(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Decrypts and restores an email message that was previously encrypted using
// AesEncrypt. The password must match the password used for encryption.
func (z *Email) AesDecrypt(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_AesDecrypt(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Encrypts the email body, all alternative bodies, all message sub-parts and
// attachments using 128-bit AES CBC encryption. Decrypting is achieved by calling
// AesDecrypt with the same password. The AesEncrypt/Decrypt methods use symmetric
// password-based AES encryption and greatly simplify sending and receiving
// encrypted emails because certificates and private keys are not used. However,
// the sending and receiving applications must both use Chilkat, and the password
// must be pre-known on both ends.
func (z *Email) AesEncrypt(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_AesEncrypt(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Appends a string to the plain-text body.
func (z *Email) AppendToBody(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkEmail_AppendToBody(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Sometimes emails created by other software or systems are not formatted
// according to typical conventions. This method provides a means to automatically
// fix certain problems.
// 
// The fixups set to a comma-separated list of keywords that identify the fixups to
// be applied. At the moment, there is only one fixup defined ("FixRelated") as
// described here:
//     FixRelated: Fixes the email so that HTML related items (images for example)
//     are properly located in the email MIME structure. This prevents them from being
//     seen as attachments.
//
func (z *Email) ApplyFixups(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_ApplyFixups(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Please see the examples at the following pages for detailed information:
func (z *Email) AspUnpack(arg1 string, arg2 string, arg3 string, arg4 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    v := C.CkEmail_AspUnpack(z.ckObj, cstr1, cstr2, cstr3, b4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Please see the examples at the following pages for detailed information:
func (z *Email) AspUnpack2(arg1 string, arg2 string, arg3 string, arg4 bool) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    ckOutBytes := C.CkByteData_Create()

    v := C.CkEmail_AspUnpack2(z.ckObj, cstr1, cstr2, cstr3, b4, ckOutBytes)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
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


// Attaches a MIME message to the email object. The attached MIME will be
// encapsulated in an message/rfc822 sub-part. To attach one email object to
// another, pass the output of GetMimeBinary to the input of this method.
func (z *Email) AttachMessage(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkEmail_AttachMessage(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Takes a byte array of multibyte (non-Unicode) data and returns a Unicode
// B-Encoded string.
// return a string or nil if failed.
func (z *Email) BEncodeBytes(arg1 []byte, arg2 string) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_bEncodeBytes(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Takes a Unicode string, converts it to the charset specified in the 2nd
// parameter, B-Encodes the converted multibyte data, and returns the encoded
// Unicode string.
// return a string or nil if failed.
func (z *Email) BEncodeString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_bEncodeString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Clears the email object of all information to the state as if the object was
// just created, which will have default headers such as Mime-Version, Date,
// Message-ID, Content-Type, Content-Transfer-Encoding, and X-Priority.
func (z *Email) Clear()  {

    C.CkEmail_Clear(z.ckObj)



}


// Clears the list of blind carbon-copy recipients.
func (z *Email) ClearBcc()  {

    C.CkEmail_ClearBcc(z.ckObj)



}


// Clears the list of carbon-copy recipients.
func (z *Email) ClearCC()  {

    C.CkEmail_ClearCC(z.ckObj)



}


// Clears the internal list of explicitly specified certificates to be used for
// this encrypted email.
func (z *Email) ClearEncryptCerts()  {

    C.CkEmail_ClearEncryptCerts(z.ckObj)



}


// Clears the list of "to" recipients.
func (z *Email) ClearTo()  {

    C.CkEmail_ClearTo(z.ckObj)



}


// Creates and returns an identical copy of the Email object.
func (z *Email) Clone() *Email {

    retObj := C.CkEmail_Clone(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Computes a global unique key for the email. The key is created by a digest-MD5
// hash of the concatenation of the following:
// messageID + CRLF + subject + CRLF + from + CRLF + date + CRLF + recipientAddrs
// 
// messageID contains the contents of the Message-ID header field.
// subject contains the contents of the Subject header field, trimmed of whitespace from both ends, 
//     where TAB chars are converted to SPACE chars, and internal whitespace is trimmed so that 
//    no more than one SPACE char in a row exists.
// from contains the lowercase FROM header email address.
// date contains the contents of the DATE header field.
// toAddrs contains lowercase TO and CC recipient email addresses, comma separated, with duplicates removed, and sorted 
//     in ascending order.  The BCC addresses are NOT included.
// 
// (After calling this method, the LastErrorText property can be examined to see the string that was hashed.)
// The 16-byte MD5 hash is returned as an encoded string. The encoding determines the
// encoding: base64, hex, url, etc. If bFold is true, then the 16-byte MD5 hash is
// folded to 8 bytes with an XOR to produce a shorter key.
// return a string or nil if failed.
func (z *Email) ComputeGlobalKey2(arg1 string, arg2 bool) *string {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    retStr := C.CkEmail_computeGlobalKey2(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Converts images embedded inline within HTML to multipart/related MIME parts
// referenced from the HTML by CID.
func (z *Email) ConvertInlineImages() bool {

    v := C.CkEmail_ConvertInlineImages(z.ckObj)


    return v != 0
}


// Creates a new DSN (Delivery Status Notification) email having the format as
// specified in RFC 3464. See the example (below) for more detailed information.
func (z *Email) CreateDsn(arg1 string, arg2 string, arg3 bool) *Email {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retObj := C.CkEmail_CreateDsn(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Returns a copy of the Email object with the body and header fields changed so
// that the newly created email can be forwarded. After calling CreateForward,
// simply add new recipients to the created email, and call MailMan.SendEmail.
func (z *Email) CreateForward() *Email {

    retObj := C.CkEmail_CreateForward(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Creates a new MDN (Message Disposition Notification) email having the format as
// specified in RFC 3798. See the example (below) for more detailed information.
func (z *Email) CreateMdn(arg1 string, arg2 string, arg3 bool) *Email {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    retObj := C.CkEmail_CreateMdn(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Returns a copy of the Email object with the body and header fields changed so
// that the newly created email can be sent as a reply. After calling CreateReply,
// simply prepend additional information to the body, and call MailMan.SendEmail.
// 
// Note: Attachments are not included in the returned reply email. However,
// attached messages are included. If an application does not wish to include the
// attached messages in a reply email, they can be removed by calling
// RemoveAttachedMessages on the reply email object.
//
func (z *Email) CreateReply() *Email {

    retObj := C.CkEmail_CreateReply(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Saves the email to a temporary MHT file so that a WebBrowser control can
// navigate to it and display it. If fileName is empty, a temporary filename is
// generated and returned. If fileName is non-empty, then it will be created or
// overwritten, and the input filename is simply returned.The MHT file that is
// created will not contain any of the email's attachments, if any existed. Also,
// if the email was plain-text, the MHT file will be saved such that the plain-text
// is converted to HTML using pre-formatted text ("pre" HTML tags) allowing it to
// be displayed correctly in a WebBrowser.
// return a string or nil if failed.
func (z *Email) CreateTempMht(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkEmail_createTempMht(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Removes all attachments from the email.
func (z *Email) DropAttachments()  {

    C.CkEmail_DropAttachments(z.ckObj)



}


// A related item is typically an embedded image referenced from the HTML in the
// email via a "CID" hyperlink. This method removes the Nth embedded image from the
// email. Note: If the HTML tries to reference the removed image, it will be
// displayed as a broken image link.
func (z *Email) DropRelatedItem(arg1 int)  {

    C.CkEmail_DropRelatedItem(z.ckObj, C.int(arg1))



}


// A related item is typically an embedded image referenced from the HTML in the
// email via a "CID" hyperlink. This method removes all the embedded images from
// the email.
func (z *Email) DropRelatedItems()  {

    C.CkEmail_DropRelatedItems(z.ckObj)



}


// Drops a single attachment from the email. Returns True if successful.
func (z *Email) DropSingleAttachment(arg1 int) bool {

    v := C.CkEmail_DropSingleAttachment(z.ckObj, C.int(arg1))


    return v != 0
}


// Digitally signed and/or encrypted emails are automatically "unwrapped" when
// received from a POP3 or IMAP server, or when loaded from any source such as a
// MIME string, in-memory byte data, or a .eml file. The results of the signature
// verification / decryption are stored in the properties such as ReceivedSigned,
// ReceivedEncrypted, SignaturesValid, etc. The signing certificate can be obtained
// via the GetSigningCert function. If the signature contained more certificates in
// the chain of authentication, this method provides a means to access them.
// 
// During signature verification, the email object collects the certs found in the
// signature and holds onto them internally. To get the issuing certificate of the
// signing certificate, call this method passing the cert returned by
// GetSigningCert. If the issuing cert is available, it is returned. Otherwise
// _NULL_ is returned. If the cert passed in is the root (i.e. a self-signed
// certificate), then the cert object returned is a copy of the cert passed in.
// 
// To traverse the chain to the root, one would write a loop that on first
// iteration passes the cert returned by GetSignedByCert (not GetSignerCert), and
// then on each subsequent iteration passes the cert from the previous iteration.
// The loop would exit when a cert is returned that has the same SubjectDN and
// SerialNumber as what was passed in (or when FindIssuer returns _NULL_).
//
func (z *Email) FindIssuer(arg1 *Cert) *Cert {

    retObj := C.CkEmail_FindIssuer(z.ckObj, arg1.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Generates a unique filename for this email. The filename will be different each
// time the method is called.
// return a string or nil if failed.
func (z *Email) GenerateFilename() *string {

    retStr := C.CkEmail_generateFilename(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth alternative body. The NumAlternatives property tells the number
// of alternative bodies present. Use the GetHtmlBody and GetPlainTextBody methods
// to easily get the HTML or plain text alternative bodies.
// return a string or nil if failed.
func (z *Email) GetAlternativeBody(arg1 int) *string {

    retStr := C.CkEmail_getAlternativeBody(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns contents of the Nth alternative body to binData. The 1st alternative body
// is at index 0. This method should only be called when the NumAlternatives
// property has a value greater than 0.
func (z *Email) GetAlternativeBodyBd(arg1 int, arg2 *BinData) bool {

    v := C.CkEmail_GetAlternativeBodyBd(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Returns the alternative body by content-type, such as "text/plain", "text/html",
// "text/xml", etc.
// return a string or nil if failed.
func (z *Email) GetAlternativeBodyByContentType(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkEmail_getAlternativeBodyByContentType(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the content type of the Nth alternative body. The NumAlternatives
// property tells the number of alternative bodies present.
// return a string or nil if failed.
func (z *Email) GetAlternativeContentType(arg1 int) *string {

    retStr := C.CkEmail_getAlternativeContentType(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of a header field within the Nth alternative body's MIME
// sub-part.
// return a string or nil if failed.
func (z *Email) GetAltHeaderField(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_getAltHeaderField(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns an embedded "message/rfc822" subpart as an email object. (Emails are
// embedded as "message/rfc822" subparts by some mail clients when forwarding an
// email.) This method allows the original email to be accessed.
func (z *Email) GetAttachedMessage(arg1 int) *Email {

    retObj := C.CkEmail_GetAttachedMessage(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// Returns a header field attribute value for the Nth attached (embedded) email.
// For example, to get the value of the "name" attribute in the Content-Type header
// for the 1st attached message:
// Content-Type: message/rfc822; name="md75000024149.eml"
// then the method arguments should contain the values 0, "Content-Type", "name".
// return a string or nil if failed.
func (z *Email) GetAttachedMessageAttr(arg1 int, arg2 string, arg3 string) *string {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkEmail_getAttachedMessageAttr(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the filename of the Nth attached (embedded) email. The filename is the
// "filename" attribute of the content-disposition header field found within the
// Nth message/rfc822 sub-part of the calling email object.
// 
// Important: The attached message filename is only present if the
// Content-Disposition header exists AND contains a "filename" attribute. If
// questions arise, one could open the email in a text editor to examine the MIME
// sub-header for the attached message (where the Content-Type = "message/rfc822").
// For example, here is a sub-header that has a filename:
// Content-Type: message/rfc822; name="GetAttachedMessageAttr.eml"
// Content-Transfer-Encoding: 7bit
// Content-Disposition: attachment; filename="GetAttachedMessageAttr.eml"
// Here is an attached message sub-header that does NOT have a filename:
// Content-Type: message/rfc822
// Content-Transfer-Encoding: 7bit
// Content-Disposition: attachment
//
// return a string or nil if failed.
func (z *Email) GetAttachedMessageFilename(arg1 int) *string {

    retStr := C.CkEmail_getAttachedMessageFilename(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a header field attribute value from the header field of the Nth
// attachment.
// return a string or nil if failed.
func (z *Email) GetAttachmentAttr(arg1 int, arg2 string, arg3 string) *string {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkEmail_getAttachmentAttr(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Copies an attachment's binary data into binData. The first attachment is at index
// 0.
func (z *Email) GetAttachmentBd(arg1 int, arg2 *BinData) bool {

    v := C.CkEmail_GetAttachmentBd(z.ckObj, C.int(arg1), arg2.ckObj)


    return v != 0
}


// Returns the ContentID header field for the Nth attachment. The first attachment
// is at index 0.
// return a string or nil if failed.
func (z *Email) GetAttachmentContentID(arg1 int) *string {

    retStr := C.CkEmail_getAttachmentContentID(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Content-Type header field for the Nth attachment. Indexing of
// attachments begins at 0.
// return a string or nil if failed.
func (z *Email) GetAttachmentContentType(arg1 int) *string {

    retStr := C.CkEmail_getAttachmentContentType(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Retrieves an attachment's binary data for in-memory access.
func (z *Email) GetAttachmentData(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkEmail_GetAttachmentData(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Retrieves an attachment's filename.
// return a string or nil if failed.
func (z *Email) GetAttachmentFilename(arg1 int) *string {

    retStr := C.CkEmail_getAttachmentFilename(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of a header field (by name) of an attachment.
// return a string or nil if failed.
func (z *Email) GetAttachmentHeader(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_getAttachmentHeader(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the size (in bytes) of the Nth attachment. The 1st attachment is at
// index 0. Returns -1 if there is no attachment at the specified index.
func (z *Email) GetAttachmentSize(arg1 int) int {

    v := C.CkEmail_GetAttachmentSize(z.ckObj, C.int(arg1))


    return int(v)
}


// Retrieves an attachment's data as a String. All CRLF sequences will be
// translated to single newline characters. The charset indicates how to interpret the
// bytes of the attachment. For example, if the attachment is a text file using the
// utf-8 byte representation for characters, pass "utf-8".
// return a string or nil if failed.
func (z *Email) GetAttachmentString(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_getAttachmentString(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Retrieves an attachment's data as a String. All end-of-lines will be translated
// to CRLF sequences.
// return a string or nil if failed.
func (z *Email) GetAttachmentStringCrLf(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_getAttachmentStringCrLf(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a blind carbon-copy recipient's full email address.
// return a string or nil if failed.
func (z *Email) GetBcc(arg1 int) *string {

    retStr := C.CkEmail_getBcc(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth BCC address (only the address part, not the friendly-name part).
// return a string or nil if failed.
func (z *Email) GetBccAddr(arg1 int) *string {

    retStr := C.CkEmail_getBccAddr(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth BCC name (only the friendly-name part, not the address part).
// return a string or nil if failed.
func (z *Email) GetBccName(arg1 int) *string {

    retStr := C.CkEmail_getBccName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a carbon-copy recipient's full email address.
// return a string or nil if failed.
func (z *Email) GetCC(arg1 int) *string {

    retStr := C.CkEmail_getCC(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth CC address (only the address part, not the friendly-name part).
// return a string or nil if failed.
func (z *Email) GetCcAddr(arg1 int) *string {

    retStr := C.CkEmail_getCcAddr(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth CC name (only the friendly-name part, not the address part).
// return a string or nil if failed.
func (z *Email) GetCcName(arg1 int) *string {

    retStr := C.CkEmail_getCcName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// If the email is a multipart/report, then it is a delivery status notification.
// This method can be used to get individual pieces of information from the
// message/delivery-status part of the email. This method should only be called if
// the IsMultipartReport method returns true.
// 
// The fieldName should be set a string such as "Final-Recipient", "Status", "Action",
// "Reporting-MTA", etc.
// Reporting-MTA: dns; XYZ.abc.nl
// 
// Final-recipient: RFC822; someEmailAddr@doesnotexist123.nl
// Action: failed
// Status: 5.4.4
// X-Supplementary-Info: 
//
// return a string or nil if failed.
func (z *Email) GetDeliveryStatusInfo(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkEmail_getDeliveryStatusInfo(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns an digest contained within a multipart/digest as an email object. The
// 1st digest is at index 0. Use the NumDigests property to get the number of
// digests available.
// 
// Note: This example requires Chilkat v9.5.0.66 or greater.
//
func (z *Email) GetDigest(arg1 int) *Email {

    retObj := C.CkEmail_GetDigest(z.ckObj, C.int(arg1))


    if retObj == nil {
            return nil
            }
    return &Email{retObj}
}


// If the email is a multipart/report, then it is a delivery status notification.
// This method can be used to get Final-Recipient values from the
// message/delivery-status part of the email. This method should only be called if
// the IsMultipartReport method returns true.
func (z *Email) GetDsnFinalRecipients() *StringArray {

    retObj := C.CkEmail_GetDsnFinalRecipients(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Returns the date/time found in the "Date" header field as a date/time object.
func (z *Email) GetDt() *CkDateTime {

    retObj := C.CkEmail_GetDt(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Returns the certificate that was previously set by SetEncryptCert.
func (z *Email) GetEncryptCert() *Cert {

    retObj := C.CkEmail_GetEncryptCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns the certificate associated with a received encrypted email.
func (z *Email) GetEncryptedByCert() *Cert {

    retObj := C.CkEmail_GetEncryptedByCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Reads a file and returns the contents as a String. This is here purely for
// convenience.
func (z *Email) GetFileContent(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkEmail_GetFileContent(z.ckObj, cstr1, ckOutBytes)

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


// Returns the value of a header field.
// return a string or nil if failed.
func (z *Email) GetHeaderField(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkEmail_getHeaderField(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Return the name of the Nth header field. The NumHeaderFields() method can be
// used to get the number of header fields. The GetHeaderField() method can be used
// to get the value of the field given the field name.
// 
// The 1st header field is at index 0. (All Chilkat indexing is 0-based.)
//
// return a string or nil if failed.
func (z *Email) GetHeaderFieldName(arg1 int) *string {

    retStr := C.CkEmail_getHeaderFieldName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the value of the Nth header field. (Indexing begins at 0) The number of
// header fields can be obtained from the NumHeaderFields property.
// 
// The 1st header field is at index 0, the last header field is at index
// NumHeaderFields-1. (All Chilkat indexing is 0-based.)
//
// return a string or nil if failed.
func (z *Email) GetHeaderFieldValue(arg1 int) *string {

    retStr := C.CkEmail_getHeaderFieldValue(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the body having the "text/html" content type.
// return a string or nil if failed.
func (z *Email) GetHtmlBody() *string {

    retStr := C.CkEmail_getHtmlBody(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// When email headers are downloaded from an IMAP server (using Chilkat IMAP), a
// "ckx-imap-uid" header field is added. The content of this header is the UID or
// sequence number of the email on the IMAP server. In addition, a "ckx-imap-isUid"
// header field is added, and this will have the value YES or NO. If the value is
// YES, then ckx-imap-uid contains a UID, if the value is NO, then ckx-imap-uid
// contains the sequence number. This method returns the UID if ckx-imap-uid exists
// and contains a UID, otherwise it returns -1.
// 
// An application that wishes to download the full email would use this UID and
// then call the Chilkat IMAP object's FetchSingle or FetchSingleAsMime methods.
// 
// Note:If an email was downloaded from the IMAP server in a way such that the UID
// is not received, then there will be no "ckx-imap-uid" header field and this
// method would return -1. For example, if emails are downloaded by sequence
// numbers via the Imap.FetchSequence method, then UIDs are not used and therefore
// the email object will not contain this information.
//
func (z *Email) GetImapUid() int {

    v := C.CkEmail_GetImapUid(z.ckObj)


    return int(v)
}


// Parses an HTML email and returns the set of domain names that occur in
// hyperlinks within the HTML body.
func (z *Email) GetLinkedDomains() *StringArray {

    retObj := C.CkEmail_GetLinkedDomains(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &StringArray{retObj}
}


// Returns a header field's data in a byte array. If the field was Q or B encoded,
// this is automatically decoded, and the raw bytes of the field are returned. Call
// GetHeaderField to retrieve the header field as a Unicode string.
// 
// The 1st header field is at index 0. (All Chilkat indexing is 0-based.)
//
func (z *Email) GetMbHeaderField(arg1 string, arg2 string) []byte {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkEmail_GetMbHeaderField(z.ckObj, cstr1, cstr2, ckOutBytes)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the HTML body converted to a specified charset. If no HTML body exists,
// the returned byte array is empty. The returned data will be such that not only
// is the character data converted (if necessary) to the convertToCharset, but the
// HTML is edited to add or modify the META tag that specifies the charset within
// the HTML.
func (z *Email) GetMbHtmlBody(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkEmail_GetMbHtmlBody(z.ckObj, cstr1, ckOutBytes)

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


// Returns the plain-text body converted to a specified charset. The return value
// is a byte array containing multibyte character data.
func (z *Email) GetMbPlainTextBody(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkEmail_GetMbPlainTextBody(z.ckObj, cstr1, ckOutBytes)

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


// Return the email as MIME text containing the email header, body (or bodies),
// related items (if any), and all attachments
// return a string or nil if failed.
func (z *Email) GetMime() *string {

    retStr := C.CkEmail_getMime(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Return the email as binary MIME containing the email header, body (or bodies),
// related items (if any), and all attachments. The MIME is appended to the
// existing contents (if any) of bindat.
func (z *Email) GetMimeBd(arg1 *BinData) bool {

    v := C.CkEmail_GetMimeBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the full MIME of an email.
func (z *Email) GetMimeBinary() []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkEmail_GetMimeBinary(z.ckObj, ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Return the email as MIME text containing the email header, body (or bodies),
// related items (if any), and all attachments. The MIME is appended to the
// existing contents (if any) of sb.
func (z *Email) GetMimeSb(arg1 *StringBuilder) bool {

    v := C.CkEmail_GetMimeSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the binary bytes of the Nth MIME sub-part having a specified content
// type (such as "application/pdf". Indexing begins at 0. Call GetNumPartsOfType to
// find out how many MIME sub-parts exist for any given content type. If inlineOnly is
// true, then only MIME sub-parts having a content-disposition of "inline" are
// included. If excludeAttachments is true, then MIME sub-parts having a content-disposition of
// "attachment" are excluded.
// 
// Note: If the email was downloaded as header-only, it will not contain all the
// parts of the full email. Also, if downloaded from IMAP excluding attachments,
// those parts that are the attachments will (of course) be missing.
//
func (z *Email) GetNthBinaryPartOfType(arg1 int, arg2 string, arg3 bool, arg4 bool) []byte {
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }
    ckOutBytes := C.CkByteData_Create()

    v := C.CkEmail_GetNthBinaryPartOfType(z.ckObj, C.int(arg1), cstr2, b3, b4, ckOutBytes)

    C.free(unsafe.Pointer(cstr2))

    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the text of the Nth MIME sub-part having a specified content type (such
// as "text/plain". Indexing begins at 0. Call GetNumPartsOfType to find out how
// many MIME sub-parts exist for any given content type. If inlineOnly is true, then
// only MIME sub-parts having a content-disposition of "inline" are included. If
// excludeAttachments is true, then MIME sub-parts having a content-disposition of "attachment"
// are excluded.
// 
// Note: If the email was downloaded as header-only, it will not contain all the
// parts of the full email. Also, if downloaded from IMAP excluding attachments,
// those parts that are the attachments will (of course) be missing.
//
// return a string or nil if failed.
func (z *Email) GetNthTextPartOfType(arg1 int, arg2 string, arg3 bool, arg4 bool) *string {
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }
    b4 := C.int(0)
    if arg4 {
        b4 = C.int(1)
    }

    retStr := C.CkEmail_getNthTextPartOfType(z.ckObj, C.int(arg1), cstr2, b3, b4)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the number of MIME sub-parts within the email having a specified content
// type (such as "text/plain"). If inlineOnly is true, then only MIME sub-parts having
// a content-disposition of "inline" are included. If excludeAttachments is true, then MIME
// sub-parts having a content-disposition of "attachment" are excluded.
// 
// Note: If the email was downloaded as header-only, it will not contain all the
// parts of the full email. Also, if downloaded from IMAP excluding attachments,
// those parts that are the attachments will (of course) be missing.
//
func (z *Email) GetNumPartsOfType(arg1 string, arg2 bool, arg3 bool) int {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkEmail_GetNumPartsOfType(z.ckObj, cstr1, b2, b3)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the email body having the "text/plain" content type.
// return a string or nil if failed.
func (z *Email) GetPlainTextBody() *string {

    retStr := C.CkEmail_getPlainTextBody(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a header field attribute value from the header field of the Nth related
// item.
// return a string or nil if failed.
func (z *Email) GetRelatedAttr(arg1 int, arg2 string, arg3 string) *string {
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    retStr := C.CkEmail_getRelatedAttr(z.ckObj, C.int(arg1), cstr2, cstr3)

    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the content ID of a related item contained with the email. Related items
// are typically images and style-sheets embedded within HTML emails.
// return a string or nil if failed.
func (z *Email) GetRelatedContentID(arg1 int) *string {

    retStr := C.CkEmail_getRelatedContentID(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Content-Location of a related item contained with the email. Related
// items are typically images and style-sheets embedded within HTML emails.
// return a string or nil if failed.
func (z *Email) GetRelatedContentLocation(arg1 int) *string {

    retStr := C.CkEmail_getRelatedContentLocation(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the content-type of the Nth related content item in an email message.
// return a string or nil if failed.
func (z *Email) GetRelatedContentType(arg1 int) *string {

    retStr := C.CkEmail_getRelatedContentType(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the content of a related item contained with the email. Related items
// are typically images and style-sheets embedded within HTML emails.
func (z *Email) GetRelatedData(arg1 int) []byte {
    ckOutBytes := C.CkByteData_Create()

    v := C.CkEmail_GetRelatedData(z.ckObj, C.int(arg1), ckOutBytes)


    if v == 0 {
        C.CkByteData_Dispose(ckOutBytes)
        return nil
    }
    p := C.CkByteData_getData(ckOutBytes)
    retBytes := C.GoBytes(unsafe.Pointer(p) ,C.int(C.CkByteData_getSize(ckOutBytes)))
    C.CkByteData_Dispose(ckOutBytes)
    return retBytes

}


// Returns the filename of a related item contained with the email. Related items
// are typically images and style-sheets embedded within HTML emails.
// return a string or nil if failed.
func (z *Email) GetRelatedFilename(arg1 int) *string {

    retStr := C.CkEmail_getRelatedFilename(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the text with CR line-endings of a related item contained with the
// email. Related items are typically images and style-sheets embedded within HTML
// emails.
// return a string or nil if failed.
func (z *Email) GetRelatedString(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_getRelatedString(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the text with CRLF line-endings of a related item contained with the
// email. Related items are typically images and style-sheets embedded within HTML
// emails.
// return a string or nil if failed.
func (z *Email) GetRelatedStringCrLf(arg1 int, arg2 string) *string {
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_getRelatedStringCrLf(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a replacement pattern previously defined for mail-merge operations.
// return a string or nil if failed.
func (z *Email) GetReplacePattern(arg1 int) *string {

    retStr := C.CkEmail_getReplacePattern(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a replacement string for a previously defined pattern/replacement string
// pair. (This is a mail-merge feature.)
// return a string or nil if failed.
func (z *Email) GetReplaceString(arg1 int) *string {

    retStr := C.CkEmail_getReplaceString(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns a replacement string for a previously defined pattern/replacement string
// pair. (This is a mail-merge feature.)
// return a string or nil if failed.
func (z *Email) GetReplaceString2(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkEmail_getReplaceString2(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// (See the NumReports property.) Returns the body content of the Nth report within
// a multipart/report email.
// 
// Multipart/report is a message type that contains data formatted for a mail
// server to read. It is split between a text/plain (or some other content/type
// easily readable) and a message/delivery-status, which contains the data
// formatted for the mail server to read.
// 
// It is defined in RFC 3462
//
// return a string or nil if failed.
func (z *Email) GetReport(arg1 int) *string {

    retStr := C.CkEmail_getReport(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Return the certificate used to digitally sign this email.
func (z *Email) GetSignedByCert() *Cert {

    retObj := C.CkEmail_GetSignedByCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Return the full certificate chain of the certificate used to digitally sign this
// email.
func (z *Email) GetSignedByCertChain() *CertChain {

    retObj := C.CkEmail_GetSignedByCertChain(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &CertChain{retObj}
}


// Return the certificate that will be used to digitally sign this email. This is
// the cerficate that was previously set by calling the SetSigningCert method.
func (z *Email) GetSigningCert() *Cert {

    retObj := C.CkEmail_GetSigningCert(z.ckObj)


    if retObj == nil {
            return nil
            }
    return &Cert{retObj}
}


// Returns a "to" recipient's full email address.
// return a string or nil if failed.
func (z *Email) GetTo(arg1 int) *string {

    retStr := C.CkEmail_getTo(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth To address (only the address part, not the friendly-name part).
// return a string or nil if failed.
func (z *Email) GetToAddr(arg1 int) *string {

    retStr := C.CkEmail_getToAddr(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the Nth To name (only the friendly-name part, not the address part).
// return a string or nil if failed.
func (z *Email) GetToName(arg1 int) *string {

    retStr := C.CkEmail_getToName(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Convert the email object to an XML document in memory
// return a string or nil if failed.
func (z *Email) GetXml() *string {

    retStr := C.CkEmail_getXml(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if the email has a header field with the specified fieldName with a
// value matching valuePattern. Case sensitivity is controlled by caseSensitive. The valuePattern may
// include 0 or more asterisk (wildcard) characters which match 0 or more of any
// character.
func (z *Email) HasHeaderMatching(arg1 string, arg2 string, arg3 bool) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    b3 := C.int(0)
    if arg3 {
        b3 = C.int(1)
    }

    v := C.CkEmail_HasHeaderMatching(z.ckObj, cstr1, cstr2, b3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Returns true if the email has an HTML body.
func (z *Email) HasHtmlBody() bool {

    v := C.CkEmail_HasHtmlBody(z.ckObj)


    return v != 0
}


// Returns true if the email has a plain-text body.
func (z *Email) HasPlainTextBody() bool {

    v := C.CkEmail_HasPlainTextBody(z.ckObj)


    return v != 0
}


// Returns true if the email is a multipart/report email.
func (z *Email) IsMultipartReport() bool {

    v := C.CkEmail_IsMultipartReport(z.ckObj)


    return v != 0
}


// Loads a complete email from a .EML file. (EML files are simply RFC822 MIME text
// files.)
// 
// Note: This replaces the entire contents of the email object, including the To/CC
// recipients.
//
func (z *Email) LoadEml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_LoadEml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads the email object from a completed asynchronous task.
func (z *Email) LoadTaskResult(arg1 *Task) bool {

    v := C.CkEmail_LoadTaskResult(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads an email with the contents of an XML email file.
// 
// Note: This replaces the entire contents of the email object, including the To/CC
// recipients.
//
func (z *Email) LoadXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_LoadXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads an email from an XML string (previously obtained by calling the GetXml
// method). The contents of the calling email object are erased and replaced with
// the email contained within the XML string.
func (z *Email) LoadXmlString(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_LoadXmlString(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Takes a byte array of multibyte (non-Unicode) data and returns a Unicode
// Q-Encoded string.
// return a string or nil if failed.
func (z *Email) QEncodeBytes(arg1 []byte, arg2 string) *string {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_qEncodeBytes(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Takes a Unicode string, converts it to the charset specified in the 2nd
// parameter, Q-Encodes the converted multibyte data, and returns the encoded
// Unicode string.
// return a string or nil if failed.
func (z *Email) QEncodeString(arg1 string, arg2 string) *string {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    retStr := C.CkEmail_qEncodeString(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Removes the Nth message/rfc822 sub-part of the email. Indexing begins at 0.
func (z *Email) RemoveAttachedMessage(arg1 int)  {

    C.CkEmail_RemoveAttachedMessage(z.ckObj, C.int(arg1))



}


// Removes all message/rfc822 sub-parts of the email object.
func (z *Email) RemoveAttachedMessages()  {

    C.CkEmail_RemoveAttachedMessages(z.ckObj)



}


// Removes path information from all attachment filenames.
func (z *Email) RemoveAttachmentPaths()  {

    C.CkEmail_RemoveAttachmentPaths(z.ckObj)



}


// Removes by name all occurrences of a header field.
func (z *Email) RemoveHeaderField(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkEmail_RemoveHeaderField(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Removes the HTML body from the email (if an HTML body exists).
func (z *Email) RemoveHtmlAlternative()  {

    C.CkEmail_RemoveHtmlAlternative(z.ckObj)



}


// Removes the plain-text body from the email (if a plain-text body exists).
func (z *Email) RemovePlainTextAlternative()  {

    C.CkEmail_RemovePlainTextAlternative(z.ckObj)



}


// Save all the attachments of an email to files in a directory specified by dirPath.
// The OverwriteExisting property controls whether existing files are allowed to be
// overwritten.
// 
// Note: Email attachment filenames can be renamed or modified prior to saving. The
// number of attachments is available in the NumAttachments property. An
// application can loop over the attachments to get the filename for each by
// calling GetAttachmentFilename(index). Each attachment's filename can be set by
// calling SetAttachmentFilename(index, newFilename).
//
func (z *Email) SaveAllAttachments(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_SaveAllAttachments(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the Nth email attachment to the directory specified by dirPath. The 1st
// attachment is at index 0. The OverwriteExisting property controls whether
// existing files are allowed to be overwritten.
func (z *Email) SaveAttachedFile(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkEmail_SaveAttachedFile(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Convert this email object to EML and save it to a file.
func (z *Email) SaveEml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_SaveEml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Saves the Nth related item to the directory specified by dirPath. (The 1st related
// item is at index 0) Related content items are typically image or style-sheets
// embedded within an HTML email. The OverwriteExisting property controls whether
// existing files are allowed to be overwritten.
func (z *Email) SaveRelatedItem(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkEmail_SaveRelatedItem(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Convert this email object to XML and save it to a file.
func (z *Email) SaveXml(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_SaveXml(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the charset attribute of the content-type header field for a specified
// attachment. This can be used if the attachment is a text file that contains text
// in a non us-ascii charset such as Shift_JIS, iso-8859-2, big5, iso-8859-5, etc.
func (z *Email) SetAttachmentCharset(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkEmail_SetAttachmentCharset(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Set's an attachment's disposition. The default disposition of an attachment is
// "attachment". This method is typically called to change the disposition to
// "inline". The 1st attachment is at index 0.
func (z *Email) SetAttachmentDisposition(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkEmail_SetAttachmentDisposition(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Renames a email attachment's filename. The 1st attachment is at index 0.
func (z *Email) SetAttachmentFilename(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkEmail_SetAttachmentFilename(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Sets the main body of the email to binary content of any type. The disposition
// can be an empty string, "inline", or "attachment". If a filename is specified,
// the disposition must be non-empty because the filename is an attribute of the
// content-disposition header field.
func (z *Email) SetBinaryBody(arg1 []byte, arg2 string, arg3 string, arg4 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkEmail_SetBinaryBody(z.ckObj, ckbyd1, cstr2, cstr3, cstr4)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Allows for a certificate to be explicity provided for decryption. When an email
// object is loaded via any method, such as LoadEml, SetFromMimeText,
// SetFromMimeBytes, etc., security layers (signatures and encryption) are
// automatically unwrapped. This method could be called prior to calling a method
// that loads the email.
func (z *Email) SetDecryptCert(arg1 *Cert) bool {

    v := C.CkEmail_SetDecryptCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Allows for a certificate and private key to be explicity specified for
// decryption. When an email object is loaded via any method, such as LoadEml,
// SetFromMimeText, SetFromMimeBytes, etc., security layers (signatures and
// encryption) are automatically unwrapped. Decryption requires a private key. On
// Windows-based systems, the private key is often pre-installed and nothing need
// be done to provide it because Chilkat will automatically find it and use it.
// However, if not on a Windows system, or if the private key was not
// pre-installed, then it can be provided by this method, or via the
// AddPfxSourceFile / AddPfxSourceData methods.
func (z *Email) SetDecryptCert2(arg1 *Cert, arg2 *PrivateKey) bool {

    v := C.CkEmail_SetDecryptCert2(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Sets the "Date" header field of the email to have the value of the date/time
// object provided.
func (z *Email) SetDt(arg1 *CkDateTime) bool {

    v := C.CkEmail_SetDt(z.ckObj, arg1.ckObj)


    return v != 0
}


// Creates a typical email used to send EDIFACT messages. Does the following:
//     Sets the email body to the EDIFACT message passed in message.
//     Sets the Content-Transfer-Encoding to Base64.
//     Set the Content-Type equal to "application/EDIFACT".
//     Sets the Content-Type header's name attribute to name.
//     Sets the Content-Disposition equal to "attachment".
//     Sets the Content-Disposition's "filename" attribute equal to filename.
//     The EDIFACT message is converted to the charset indicated by charset, and
//     encoded using Base64 in the email body.
// The email's subject, recipients, FROM address, and other headers are left
// unmodified.
func (z *Email) SetEdifactBody(arg1 string, arg2 string, arg3 string, arg4 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    C.CkEmail_SetEdifactBody(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))


}


// Set the encryption certificate to be used in encryption. Use the CreateCS,
// CertStore, and Cert classes to create a Cert object by either locating a
// certificate in a certificate store or loading one from a file.
func (z *Email) SetEncryptCert(arg1 *Cert) bool {

    v := C.CkEmail_SetEncryptCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads an email with the MIME stored in a BinData object.
// 
// Note: This replaces the entire contents of the email object, including the To/CC
// recipients.
//
func (z *Email) SetFromMimeBd(arg1 *BinData) bool {

    v := C.CkEmail_SetFromMimeBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads the email object with the mimeBytes. If the email object already contained an
// email, it is entirely replaced. The character encoding (such as "utf-8",
// "iso-8859-1", etc.) of the bytes is automatically inferred from the content. If
// for some reason it is not possible to determine the character encoding, the
// SetFromMimeBytes2 method may be called to explicitly specify the charset.
func (z *Email) SetFromMimeBytes(arg1 []byte) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))

    v := C.CkEmail_SetFromMimeBytes(z.ckObj, ckbyd1)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)

    return v != 0
}


// Loads the email object with the mimeBytes. If the email object already contained an
// email, it is entirely replaced.
// 
// The charset specifies the character encoding of the MIME bytes (such as "utf-8",
// "iso-8859-1", etc.).
//
func (z *Email) SetFromMimeBytes2(arg1 []byte, arg2 string) bool {
    ckbyd1 := C.CkByteData_Create()
    pData1 := C.CBytes(arg1)
    C.CkByteData_borrowData(ckbyd1, (*C.uchar)(pData1), C.ulong(len(arg1)))
    cstr2 := C.CString(arg2)

    v := C.CkEmail_SetFromMimeBytes2(z.ckObj, ckbyd1, cstr2)

    C.free(pData1)
    C.CkByteData_Dispose(ckbyd1)
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Loads an email with the MIME stored in a StringBuilder object.
// 
// Note: This replaces the entire contents of the email object, including the To/CC
// recipients.
//
func (z *Email) SetFromMimeSb(arg1 *StringBuilder) bool {

    v := C.CkEmail_SetFromMimeSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads an email with the contents of a .eml (i.e. MIME) contained in a string.
// 
// Note: This replaces the entire contents of the email object, including the To/CC
// recipients.
//
func (z *Email) SetFromMimeText(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_SetFromMimeText(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Loads an email from an XML string.
// 
// Note: This replaces the entire contents of the email object, including the To/CC
// recipients.
//
func (z *Email) SetFromXmlText(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_SetFromXmlText(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Sets the HTML body of an email.
func (z *Email) SetHtmlBody(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkEmail_SetHtmlBody(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Sets the HTML email body from a byte array containing character data in the
// specified character set. This method also updates the email "content-type"header
// to properly reflect the content type of the body.
func (z *Email) SetMbHtmlBody(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkEmail_SetMbHtmlBody(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Sets the plain-text email body from a byte array containing character data in
// the specified character set. This method also updates the email
// "content-type"header to properly reflect the content type of the body.
func (z *Email) SetMbPlainTextBody(arg1 string, arg2 []byte) bool {
    cstr1 := C.CString(arg1)
    ckbyd2 := C.CkByteData_Create()
    pData2 := C.CBytes(arg2)
    C.CkByteData_borrowData(ckbyd2, (*C.uchar)(pData2), C.ulong(len(arg2)))

    v := C.CkEmail_SetMbPlainTextBody(z.ckObj, cstr1, ckbyd2)

    C.free(unsafe.Pointer(cstr1))
    C.free(pData2)
    C.CkByteData_Dispose(ckbyd2)

    return v != 0
}


// Sets the filename for a related item within the email.
func (z *Email) SetRelatedFilename(arg1 int, arg2 string) bool {
    cstr2 := C.CString(arg2)

    v := C.CkEmail_SetRelatedFilename(z.ckObj, C.int(arg1), cstr2)

    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Create a pattern/replacement-text pair for mail-merge. When the email is sent
// via the MailMan's SendEmail method, or any other mail-sending method, the
// patterns are replaced with the replacement strings during the sending process.
// Do define multiple replacement patterns, simply call SetReplacePattern once per
// pattern/replacement string. (Note: The MailMan's RenderToMime method will also
// do pattern replacements. Methods such as SaveEml or GetMime do not replace
// patterns.)
// 
// Note: Replacement patterns may be placed in any header field, and in both HTML
// and plain-text email bodies.
//
func (z *Email) SetReplacePattern(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkEmail_SetReplacePattern(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Set the certificate to be used in creating a digital signature. Use the
// CreateCS, CertStore, and Cert classes to create a Cert object by either locating
// a certificate in a certificate store or loading one from a file.
func (z *Email) SetSigningCert(arg1 *Cert) bool {

    v := C.CkEmail_SetSigningCert(z.ckObj, arg1.ckObj)


    return v != 0
}


// Explicitly sets the certificate and private key to be used for sending digitally
// signed email. If the certificate's private key is already installed on the
// computer, then one may simply call SetSigningCert because the Chilkat component
// will automatically locate and use the corresponding private key (stored in the
// Windows Protected Store). In most cases, if the digital certificate is already
// installed w/ private key on the computer, it is not necessary to explicitly set
// the signing certificate at all. The Chilkat component will automatically locate
// and use the certificate containing the FROM email address (from the
// registry-based certificate store where it was installed).
func (z *Email) SetSigningCert2(arg1 *Cert, arg2 *PrivateKey) bool {

    v := C.CkEmail_SetSigningCert2(z.ckObj, arg1.ckObj, arg2.ckObj)


    return v != 0
}


// Sets the body of the email and also sets the Content-Type header field of the
// contentType. If the email is already multipart/alternative, an additional alternative
// with the indicated Content-Type will be added. If an alternative with the same
// Content-Type already exists, it is replaced.
func (z *Email) SetTextBody(arg1 string, arg2 string)  {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    C.CkEmail_SetTextBody(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))


}


// True if the caller email has a UIDL that equals the email passed in the
// argument.
func (z *Email) UidlEquals(arg1 *Email) bool {

    v := C.CkEmail_UidlEquals(z.ckObj, arg1.ckObj)


    return v != 0
}


// Unpacks an HTML email into an HTML file and related files (images and style
// sheets). The links within the HTML are updated to point to the files unpacked
// and saved to disk.
func (z *Email) UnpackHtml(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkEmail_UnpackHtml(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Unobfuscates emails by undoing what spammers do to obfuscate email. It removes
// comments from HTML bodies and unobfuscates hyperlinked URLs.
func (z *Email) UnSpamify()  {

    C.CkEmail_UnSpamify(z.ckObj)



}


// Unzips and replaces any Zip file attachments with the expanded contents. As an
// example, if an email contained a single Zip file containing 3 GIF image files as
// an attachment, then after calling this method the email would contain 3 GIF file
// attachments, and the Zip attachment would be gone.If an email contains multiple
// Zip file attachments, each Zip is expanded and replaced with the contents.
func (z *Email) UnzipAttachments() bool {

    v := C.CkEmail_UnzipAttachments(z.ckObj)


    return v != 0
}


// Adds an XML certificate vault to the object's internal list of sources to be
// searched for certificates and private keys when encrypting/decrypting or
// signing/verifying. Unlike the AddPfxSourceData and AddPfxSourceFile methods,
// only a single XML certificate vault can be used. If UseCertVault is called
// multiple times, only the last certificate vault will be used, as each call to
// UseCertVault will replace the certificate vault provided in previous calls.
func (z *Email) UseCertVault(arg1 *XmlCertVault) bool {

    v := C.CkEmail_UseCertVault(z.ckObj, arg1.ckObj)


    return v != 0
}


// Replaces all the attachments of an email with a single Zip file attachment
// having the filename specified.
func (z *Email) ZipAttachments(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkEmail_ZipAttachments(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


