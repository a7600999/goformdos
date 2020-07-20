// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkCache.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewCache() *Cache {
	return &Cache{C.CkCache_Create()}
}

func (z *Cache) DisposeCache() {
    if z != nil {
	C.CkCache_Dispose(z.ckObj)
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
func (z *Cache) DebugLogFilePath() string {
    return C.GoString(C.CkCache_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *Cache) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkCache_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Cache) LastErrorHtml() string {
    return C.GoString(C.CkCache_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *Cache) LastErrorText() string {
    return C.GoString(C.CkCache_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *Cache) LastErrorXml() string {
    return C.GoString(C.CkCache_lastErrorXml(z.ckObj))
}

// property: LastEtagFetched
// The ETag of the last item fetched from cache.
func (z *Cache) LastEtagFetched() string {
    return C.GoString(C.CkCache_lastEtagFetched(z.ckObj))
}

// property: LastExpirationFetchedStr
// Expiration date/time of the last item fetched from cache in RFC822 string
// format.
func (z *Cache) LastExpirationFetchedStr() string {
    return C.GoString(C.CkCache_lastExpirationFetchedStr(z.ckObj))
}

// property: LastHitExpired
// true if the LastExpirationFetched is before the current date/time. Otherwise
// false.
func (z *Cache) LastHitExpired() bool {
    v := int(C.CkCache_getLastHitExpired(z.ckObj))
    return v != 0
}

// property: LastKeyFetched
// The key of the last item fetched from cache. (For web pages, the key is
// typically the canonicalized URL. Otherwise, the key is a unique identifer used
// to access the cached item.)
func (z *Cache) LastKeyFetched() string {
    return C.GoString(C.CkCache_lastKeyFetched(z.ckObj))
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
func (z *Cache) LastMethodSuccess() bool {
    v := int(C.CkCache_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Cache) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCache_putLastMethodSuccess(z.ckObj,v)
}

// property: Level
// The number of directory levels in the cache. Possible values are:
// 
//     0: All cache files are in a single directory (the cache root).
// 
//     1: Cache files are located in 256 sub-directories numbered 0 .. 255 directly
//     under the cache root.
// 
//     2: There are two levels of sub-directories under the cache root. The 1st
//     level has 256 sub-directories numbered 0 .. 255 directly under the cache root.
//     The 2nd level allows for up to 256 sub-directories (0..255) under each level-1
//     directory. Cache files are stored in the leaf directories.
//
func (z *Cache) Level() int {
    return int(C.CkCache_getLevel(z.ckObj))
}
// property setter: Level
func (z *Cache) SetLevel(value int) {
    C.CkCache_putLevel(z.ckObj,C.int(value))
}

// property: NumRoots
// The number of root directories composing the cache. A typical multi-root cache
// would place each root on a separate hard drive.
func (z *Cache) NumRoots() int {
    return int(C.CkCache_getNumRoots(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *Cache) VerboseLogging() bool {
    v := int(C.CkCache_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *Cache) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkCache_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *Cache) Version() string {
    return C.GoString(C.CkCache_version(z.ckObj))
}
// Must be called once for each cache root. For example, if the cache is spread
// across D:\cacheRoot, E:\cacheRoot, and F:\cacheRoot, an application would setup
// the cache object by calling AddRoot three times -- once with "D:\cacheRoot",
// once with "E:\cacheRoot", and once with "F:\cacheRoot".
func (z *Cache) AddRoot(arg1 string)  {
    cstr1 := C.CString(arg1)

    C.CkCache_AddRoot(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))


}


// Deletes all items in the cache. This method completely clears the cache. All
// files in the cache are deleted. (If the cache is multi-level, existing
// sub-directories are not deleted.)
// 
// Returns the number of items (i.e. cache files) deleted.
//
func (z *Cache) DeleteAll() int {

    v := C.CkCache_DeleteAll(z.ckObj)


    return int(v)
}


// Deletes all expired items from the cache.
// 
// Returns the number of items (i.e. cache files) deleted.
//
func (z *Cache) DeleteAllExpired() int {

    v := C.CkCache_DeleteAllExpired(z.ckObj)


    return int(v)
}


// Deletes a single item from the disk cache. Returns false if the item exists in
// cache but could not be deleted. Otherwise returns true.
func (z *Cache) DeleteFromCache(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCache_DeleteFromCache(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// The same as DeleteOlder, except the dateTime is passed as a CkDateTime.
func (z *Cache) DeleteOlderDt(arg1 *CkDateTime) int {

    v := C.CkCache_DeleteOlderDt(z.ckObj, arg1.ckObj)


    return int(v)
}


// The same as DeleteOlder, except the dateTimeStr is passed as a date/time in RFC822
// format.
func (z *Cache) DeleteOlderStr(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkCache_DeleteOlderStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Fetches an item from cache.
// 
// The key may be any length and may include any characters. It should uniquely
// identify the cached item. (No two items in the cache should have the same key.)
//
func (z *Cache) FetchFromCache(arg1 string) []byte {
    cstr1 := C.CString(arg1)
    ckOutBytes := C.CkByteData_Create()

    v := C.CkCache_FetchFromCache(z.ckObj, cstr1, ckOutBytes)

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


// Fetches a text item from the cache and returns it's text content.
// 
// The key may be any length and may include any characters. It should uniquely
// identify the cached item. (No two items in the cache should have the same key.)
//
// return a string or nil if failed.
func (z *Cache) FetchText(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCache_fetchText(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the eTag for an item in the cache.
// return a string or nil if failed.
func (z *Cache) GetEtag(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCache_getEtag(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the expiration date/time for an item in the cache as a CkDateTime
// object.
func (z *Cache) GetExpirationDt(arg1 string) *CkDateTime {
    cstr1 := C.CString(arg1)

    retObj := C.CkCache_GetExpirationDt(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    if retObj == nil {
            return nil
            }
    return &CkDateTime{retObj}
}


// Returns the expiration date/time for an item in the cache in RFC822 string
// format.
// return a string or nil if failed.
func (z *Cache) GetExpirationStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCache_getExpirationStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the absolute file path of the cache file associated with the key.
// return a string or nil if failed.
func (z *Cache) GetFilename(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkCache_getFilename(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the directory path of the Nth cache root. (Indexing begins at 0.)
// return a string or nil if failed.
func (z *Cache) GetRoot(arg1 int) *string {

    retStr := C.CkCache_getRoot(z.ckObj, C.int(arg1))


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns true if the item is found in the cache, otherwise returns false.
func (z *Cache) IsCached(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCache_IsCached(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// The same as SaveText, except the expire date/time is passed as a CkDateTime
// object.
func (z *Cache) SaveTextDt(arg1 string, arg2 *CkDateTime, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkCache_SaveTextDt(z.ckObj, cstr1, arg2.ckObj, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// Inserts or replaces an text item in the cache with no expiration date/time. The
// eTag is optional and may be set to a zero-length string. Applications may use it
// as a place to save additional information about the cached item.
func (z *Cache) SaveTextNoExpire(arg1 string, arg2 string, arg3 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)

    v := C.CkCache_SaveTextNoExpire(z.ckObj, cstr1, cstr2, cstr3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// The same as SaveText, except the expire date/time is passed as a string in
// RFC822 format.
func (z *Cache) SaveTextStr(arg1 string, arg2 string, arg3 string, arg4 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    cstr4 := C.CString(arg4)

    v := C.CkCache_SaveTextStr(z.ckObj, cstr1, cstr2, cstr3, cstr4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(unsafe.Pointer(cstr4))

    return v != 0
}


// The same as SaveToCache, except the expire date/time is passed as a CkDateTime
// object.
func (z *Cache) SaveToCacheDt(arg1 string, arg2 *CkDateTime, arg3 string, arg4 []byte) bool {
    cstr1 := C.CString(arg1)
    cstr3 := C.CString(arg3)
    ckbyd4 := C.CkByteData_Create()
    pData4 := C.CBytes(arg4)
    C.CkByteData_borrowData(ckbyd4, (*C.uchar)(pData4), C.ulong(len(arg4)))

    v := C.CkCache_SaveToCacheDt(z.ckObj, cstr1, arg2.ckObj, cstr3, ckbyd4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr3))
    C.free(pData4)
    C.CkByteData_Dispose(ckbyd4)

    return v != 0
}


// Inserts or replaces an item in the cache. The cached item will have no
// expiration. The eTag is optional and may be set to a zero-length string.
// Applications may use it as a place to save additional information about the
// cached item.
func (z *Cache) SaveToCacheNoExpire(arg1 string, arg2 string, arg3 []byte) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    ckbyd3 := C.CkByteData_Create()
    pData3 := C.CBytes(arg3)
    C.CkByteData_borrowData(ckbyd3, (*C.uchar)(pData3), C.ulong(len(arg3)))

    v := C.CkCache_SaveToCacheNoExpire(z.ckObj, cstr1, cstr2, ckbyd3)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(pData3)
    C.CkByteData_Dispose(ckbyd3)

    return v != 0
}


// The same as SaveToCache, except the expire date/time is passed in RFC822 string
// format.
func (z *Cache) SaveToCacheStr(arg1 string, arg2 string, arg3 string, arg4 []byte) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)
    cstr3 := C.CString(arg3)
    ckbyd4 := C.CkByteData_Create()
    pData4 := C.CBytes(arg4)
    C.CkByteData_borrowData(ckbyd4, (*C.uchar)(pData4), C.ulong(len(arg4)))

    v := C.CkCache_SaveToCacheStr(z.ckObj, cstr1, cstr2, cstr3, ckbyd4)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))
    C.free(unsafe.Pointer(cstr3))
    C.free(pData4)
    C.CkByteData_Dispose(ckbyd4)

    return v != 0
}


// The same as UpdateExpiration, except the expireDateTime is passed as a CkDateTime.
func (z *Cache) UpdateExpirationDt(arg1 string, arg2 *CkDateTime) bool {
    cstr1 := C.CString(arg1)

    v := C.CkCache_UpdateExpirationDt(z.ckObj, cstr1, arg2.ckObj)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// The same as UpdateExpiration, except the expireDateTime is passed in RFC822 string format.
func (z *Cache) UpdateExpirationStr(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkCache_UpdateExpirationStr(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


