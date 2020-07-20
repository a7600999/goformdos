// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkUrl.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewUrl() *Url {
	return &Url{C.CkUrl_Create()}
}

func (z *Url) DisposeUrl() {
    if z != nil {
	C.CkUrl_Dispose(z.ckObj)
	}
}




// property: Frag
// Contains any text following a fragment marker (#) in the URL, excluding the
// fragment marker. Given the URI http://www.contoso.com/index.htm#main, the
// fragment is "main".
func (z *Url) Frag() string {
    return C.GoString(C.CkUrl_frag(z.ckObj))
}

// property: Host
// The DNS host name or IP address part of the URL. For example, if the URL is
// "http://www.contoso.com:8080/", the Host is "www.contoso.com". If the URL is
// "https://192.168.1.124/test.html", the Host is "192.168.1.124".
func (z *Url) Host() string {
    return C.GoString(C.CkUrl_host(z.ckObj))
}

// property: HostType
// The type of the host name specified in the URL. Possible values are:
//     "dns": The host name is a domain name system (DNS) style host name.
//     "ipv4": The host name is an Internet Protocol (IP) version 4 host address.
//     "ipv6": The host name is an Internet Protocol (IP) version 6 host address.
func (z *Url) HostType() string {
    return C.GoString(C.CkUrl_hostType(z.ckObj))
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
func (z *Url) LastMethodSuccess() bool {
    v := int(C.CkUrl_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Url) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkUrl_putLastMethodSuccess(z.ckObj,v)
}

// property: Login
// If the URL contains a login and password, this is the login part. For example,
// if the URL is "http://user:password@www.contoso.com/index.htm ", then the login
// is "user".
func (z *Url) Login() string {
    return C.GoString(C.CkUrl_login(z.ckObj))
}

// property: Password
// If the URL contains a login and password, this is the password part. For
// example, if the URL is "http://user:password@www.contoso.com/index.htm ", then
// the password is "password".
func (z *Url) Password() string {
    return C.GoString(C.CkUrl_password(z.ckObj))
}

// property: Path
// The path (and params) part of the URL, excluding the query and fragment. If the
// URL is:
// "http://www.amazon.com/gp/product/1476752842/ref=s9_psimh_gw_p14_d0_i3?pf_rd_m=AT
// VPDKIKX0DER&pf_rd_s=desktop-1", then the path is
// "/gp/product/1476752842/ref=s9_psimh_gw_p14_d0_i3".
func (z *Url) Path() string {
    return C.GoString(C.CkUrl_path(z.ckObj))
}

// property: PathWithQueryParams
// The path (and params) part of the URL, including the query params, but excluding
// the fragment. If the URL is:
// "http://www.amazon.com/gp/product/1476752842/ref=s9_psimh_gw_p14_d0_i3?pf_rd_m=AT
// VPDKIKX0DER&pf_rd_s=desktop-1", then then this property returns
// "/gp/product/1476752842/ref=s9_psimh_gw_p14_d0_i3?pf_rd_m=ATVPDKIKX0DER&pf_rd_s=d
// esktop-1".
func (z *Url) PathWithQueryParams() string {
    return C.GoString(C.CkUrl_pathWithQueryParams(z.ckObj))
}

// property: Port
// The port number of the URL.
// 
// For example, if the URL is "http://www.contoso.com:8080/", the port number is
// 8080.
// If the URL is "https://192.168.1.124/test.html", the port number is the default
// 80.
// If the URL is "https://www.amazon.com/", the port number is the default SSL/TLS
// port 443.
//
func (z *Url) Port() int {
    return int(C.CkUrl_getPort(z.ckObj))
}

// property: Query
// The query part of the URL, excluding the fragment. If the URL is:
// "http://www.amazon.com/gp/product/1476752842/ref=s9_psimh_gw_p14_d0_i3?pf_rd_m=AT
// VPDKIKX0DER&pf_rd_s=desktop-1#frag", then the query is
// "pf_rd_m=ATVPDKIKX0DER&pf_rd_s=desktop-1".
func (z *Url) Query() string {
    return C.GoString(C.CkUrl_query(z.ckObj))
}

// property: Ssl
// true if the URL indicates SSL/TLS, otherwise false. A URL beginning with
// "https://" indicates SSL/TLS.
func (z *Url) Ssl() bool {
    v := int(C.CkUrl_getSsl(z.ckObj))
    return v != 0
}
// Parses a full URL. After parsing, the various parts of the URL are available iin
// the properties.
func (z *Url) ParseUrl(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkUrl_ParseUrl(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


