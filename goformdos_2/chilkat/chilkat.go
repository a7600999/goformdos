// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

import "C"
import "unsafe"


type ChilkatProgressInfo func(name,value string, userData interface{})
type ChilkatPercentDone func(percentDone int, userData interface{}) bool
type ChilkatAbortCheck func(userData interface{}) bool

type ChilkatCallbacks struct {
	ProgressInfo ChilkatProgressInfo
	PercentDone ChilkatPercentDone
	AbortCheck ChilkatAbortCheck
	UserData interface{}
}


//export ExtProgressInfoCb
func ExtProgressInfoCb(name,value *C.char, pChilkatCallbacks unsafe.Pointer) {
	cb := (*ChilkatCallbacks)(pChilkatCallbacks)
	if cb == nil {
		return
	} 
   	if cb.ProgressInfo != nil {
		cb.ProgressInfo(C.GoString(name),C.GoString(value),cb.UserData)
	}
}

//export ExtAbortCheckCb
func ExtAbortCheckCb(pChilkatCallbacks unsafe.Pointer) C.int {
	cb := (*ChilkatCallbacks)(pChilkatCallbacks)
	if cb == nil {
		return C.int(0)
	} 
   	if cb.AbortCheck != nil {
		b := cb.AbortCheck(cb.UserData)
		if b == true {
			return C.int(1)
		}
	}
	return C.int(0)
}

//export ExtPercentDoneCb
func ExtPercentDoneCb(pctDone C.int, pChilkatCallbacks unsafe.Pointer) C.int {
	cb := (*ChilkatCallbacks)(pChilkatCallbacks)
	if cb == nil {
		return C.int(0)
	} 
   	if cb.PercentDone != nil {
		b := cb.PercentDone(int(pctDone),cb.UserData)
		if b == true {
			return C.int(1)
		}
	}
	return C.int(0)
}

// The ExtTaskCompletedCb is not actually used in Chilkat/Go, but we must have it here for linkage reasons.

//export ExtTaskCompletedCb
func ExtTaskCompletedCb(cTaskObj, pChilkatCallbacks unsafe.Pointer) {
}


type Asn struct {
	ckObj unsafe.Pointer
}

type Atom struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type AuthAws struct {
	ckObj unsafe.Pointer
}

type AuthAzureAD struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type AuthAzureSAS struct {
	ckObj unsafe.Pointer
}

type AuthAzureStorage struct {
	ckObj unsafe.Pointer
}

type AuthGoogle struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type AuthUtil struct {
	ckObj unsafe.Pointer
}

type BinData struct {
	ckObj unsafe.Pointer
}

type Bounce struct {
	ckObj unsafe.Pointer
}

type Bz2 struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Cache struct {
	ckObj unsafe.Pointer
}

type Cert struct {
	ckObj unsafe.Pointer
}

type CertChain struct {
	ckObj unsafe.Pointer
}

type CertStore struct {
	ckObj unsafe.Pointer
}

type Charset struct {
	ckObj unsafe.Pointer
}

type CkDateTime struct {
	ckObj unsafe.Pointer
}

type Compression struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Crypt2 struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Csr struct {
	ckObj unsafe.Pointer
}

type Csv struct {
	ckObj unsafe.Pointer
}

type Dh struct {
	ckObj unsafe.Pointer
}

type DirTree struct {
	ckObj unsafe.Pointer
}

type Dkim struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Dsa struct {
	ckObj unsafe.Pointer
}

type DtObj struct {
	ckObj unsafe.Pointer
}

type Ecc struct {
	ckObj unsafe.Pointer
}

type EdDSA struct {
	ckObj unsafe.Pointer
}

type Email struct {
	ckObj unsafe.Pointer
}

type EmailBundle struct {
	ckObj unsafe.Pointer
}

type FileAccess struct {
	ckObj unsafe.Pointer
}

type Ftp2 struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Global struct {
	ckObj unsafe.Pointer
}

type Gzip struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Hashtable struct {
	ckObj unsafe.Pointer
}

type HtmlToText struct {
	ckObj unsafe.Pointer
}

type HtmlToXml struct {
	ckObj unsafe.Pointer
}

type Http struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type HttpRequest struct {
	ckObj unsafe.Pointer
}

type HttpResponse struct {
	ckObj unsafe.Pointer
}

type Imap struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type JavaKeyStore struct {
	ckObj unsafe.Pointer
}

type JsonArray struct {
	ckObj unsafe.Pointer
}

type JsonObject struct {
	ckObj unsafe.Pointer
}

type Jwe struct {
	ckObj unsafe.Pointer
}

type Jws struct {
	ckObj unsafe.Pointer
}

type Jwt struct {
	ckObj unsafe.Pointer
}

type Log struct {
	ckObj unsafe.Pointer
}

type MailMan struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Mailboxes struct {
	ckObj unsafe.Pointer
}

type MessageSet struct {
	ckObj unsafe.Pointer
}

type Mht struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Mime struct {
	ckObj unsafe.Pointer
}

type Ntlm struct {
	ckObj unsafe.Pointer
}

type OAuth1 struct {
	ckObj unsafe.Pointer
}

type OAuth2 struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Pem struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Pfx struct {
	ckObj unsafe.Pointer
}

type PrivateKey struct {
	ckObj unsafe.Pointer
}

type Prng struct {
	ckObj unsafe.Pointer
}

type PublicKey struct {
	ckObj unsafe.Pointer
}

type Rest struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Rsa struct {
	ckObj unsafe.Pointer
}

type Rss struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type SFtp struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type SFtpDir struct {
	ckObj unsafe.Pointer
}

type SFtpFile struct {
	ckObj unsafe.Pointer
}

type Scp struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type SecureString struct {
	ckObj unsafe.Pointer
}

type ServerSentEvent struct {
	ckObj unsafe.Pointer
}

type Socket struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Spider struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Ssh struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type SshKey struct {
	ckObj unsafe.Pointer
}

type SshTunnel struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Stream struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type StringArray struct {
	ckObj unsafe.Pointer
}

type StringBuilder struct {
	ckObj unsafe.Pointer
}

type StringTable struct {
	ckObj unsafe.Pointer
}

type Tar struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Task struct {
	ckObj unsafe.Pointer
}

type TrustedRoots struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type UnixCompress struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Upload struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Url struct {
	ckObj unsafe.Pointer
}

type WebSocket struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type Xml struct {
	ckObj unsafe.Pointer
}

type XmlCertVault struct {
	ckObj unsafe.Pointer
}

type XmlDSig struct {
	ckObj unsafe.Pointer
}

type XmlDSigGen struct {
	ckObj unsafe.Pointer
}

type Xmp struct {
	ckObj unsafe.Pointer
}

type Zip struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type ZipCrc struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

type ZipEntry struct {
	ckObj unsafe.Pointer
	callbacks ChilkatCallbacks
}

