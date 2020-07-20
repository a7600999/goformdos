// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkHashtable.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewHashtable() *Hashtable {
	return &Hashtable{C.CkHashtable_Create()}
}

func (z *Hashtable) DisposeHashtable() {
    if z != nil {
	C.CkHashtable_Dispose(z.ckObj)
	}
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
func (z *Hashtable) LastMethodSuccess() bool {
    v := int(C.CkHashtable_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *Hashtable) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkHashtable_putLastMethodSuccess(z.ckObj,v)
}
// Adds to the hash table from XML previously created by calling ToXmlSb.
func (z *Hashtable) AddFromXmlSb(arg1 *StringBuilder) bool {

    v := C.CkHashtable_AddFromXmlSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Adds or replaces an entry with the given key and integer value to the hash
// table. Returns true if a new hash entry was added or replaced. Returns false
// if an out-of-memory condition occurred.
func (z *Hashtable) AddInt(arg1 string, arg2 int) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHashtable_AddInt(z.ckObj, cstr1, C.int(arg2))

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds URL query parameters into the hashtable. The queryParams has the form:
// "field1=value1&field2=value2&field3=value3...". It is assumed that the values
// are URL encoded, and this method automatically URL decodes the values prior to
// inserting into the hashtable.
func (z *Hashtable) AddQueryParams(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHashtable_AddQueryParams(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Adds or replaces an entry with the given key and string value to the hash table.
// Returns true if a new hash entry was added or replaced. Returns false if an
// out-of-memory condition occurred.
func (z *Hashtable) AddStr(arg1 string, arg2 string) bool {
    cstr1 := C.CString(arg1)
    cstr2 := C.CString(arg2)

    v := C.CkHashtable_AddStr(z.ckObj, cstr1, cstr2)

    C.free(unsafe.Pointer(cstr1))
    C.free(unsafe.Pointer(cstr2))

    return v != 0
}


// Removes all elements from the Hashtable.
func (z *Hashtable) Clear()  {

    C.CkHashtable_Clear(z.ckObj)



}


// Removes all elements from the Hashtable and re-sizes with the specified capacity.
// 
// The capacity is the number of buckets in the hash table. In the case of a "hash
// collision", a single bucket stores multiple entries, which must be searched
// sequentially. One rule of thumb is to set the capacity to twice the number of
// expected items to be hashed. It's also preferable to set the capacity to a prime
// number. (The 1st 10,000 prime numbers are listed here:
// https://primes.utm.edu/lists/small/10000.txt )
// 
// The initial default capacity of the hash table is 521.
//
func (z *Hashtable) ClearWithNewCapacity(arg1 int) bool {

    v := C.CkHashtable_ClearWithNewCapacity(z.ckObj, C.int(arg1))


    return v != 0
}


// Determines if a given key is contained within the hash table. Returns true if
// the key exists, otherwise returns false.
func (z *Hashtable) Contains(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHashtable_Contains(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Determines if a given key is contained within the hash table. Returns true if
// the key exists, otherwise returns false.
func (z *Hashtable) ContainsIntKey(arg1 int) bool {

    v := C.CkHashtable_ContainsIntKey(z.ckObj, C.int(arg1))


    return v != 0
}


// Appends the complete set of hashtable key strings to strTable.
func (z *Hashtable) GetKeys(arg1 *StringTable) bool {

    v := C.CkHashtable_GetKeys(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the integer value associated with the specified key. If the key is not
// in the hash table, the return value is 0.
func (z *Hashtable) LookupInt(arg1 string) int {
    cstr1 := C.CString(arg1)

    v := C.CkHashtable_LookupInt(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return int(v)
}


// Returns the string value associated with the specified key.
// return a string or nil if failed.
func (z *Hashtable) LookupStr(arg1 string) *string {
    cstr1 := C.CString(arg1)

    retStr := C.CkHashtable_lookupStr(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Removes the entry with the specified key from the hash table. Returns true if
// the key existed and was removed. Returns false if the key did not already
// exist.
func (z *Hashtable) Remove(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkHashtable_Remove(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Serializes the hash table to XML format. The XML is appended to sbXml.
func (z *Hashtable) ToXmlSb(arg1 *StringBuilder) bool {

    v := C.CkHashtable_ToXmlSb(z.ckObj, arg1.ckObj)


    return v != 0
}


