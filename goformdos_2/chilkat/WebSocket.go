// The Chilkat v9.5.0.83 API for the Go programming language.
package chilkat

/*
#include <stdlib.h>
#cgo CPPFLAGS: -DCK_GO_LANG
#include "c_includes/C_CkWebSocket.h"


typedef unsigned char UBYTE;
#include "c_includes/C_CkByteData.h"


// Set CGO_LDFLAGS environment variable for linker options.
// for example (on Windows):  -LC:/ckBuildSystems/go/chilkatLib/tdm-5.1.0-64 -lchilkatExt-9.5.0 -lws2_32 -lstdc++

*/
import "C"
import "unsafe"

func NewWebSocket() *WebSocket {
	return &WebSocket{C.CkWebSocket_Create(),ChilkatCallbacks{nil,nil,nil,nil}}
}

func (z *WebSocket) DisposeWebSocket() {
    if z != nil {
	C.CkWebSocket_Dispose(z.ckObj)
	}
}




func (z *WebSocket) SetCallbackUserData(userData interface{}) {
    z.callbacks.UserData = userData;
    C.CkWebSocket_setCallbackContext(z.ckObj,unsafe.Pointer(&z.callbacks))
}

func (z *WebSocket) SetProgressInfo(f ChilkatProgressInfo) {
    z.callbacks.ProgressInfo = f
    C.CkWebSocket_setExternalProgress(z.ckObj,1)
}

func (z *WebSocket) SetAbortCheck(f ChilkatAbortCheck) {
    z.callbacks.AbortCheck = f
    C.CkWebSocket_setExternalProgress(z.ckObj,1)
}

func (z *WebSocket) SetPercentDone(f ChilkatPercentDone) {
    z.callbacks.PercentDone = f
    C.CkWebSocket_setExternalProgress(z.ckObj,1)
}




// property: CloseAutoRespond
// If true, then a Close control frame is automatically sent in response to
// receiving a Close control frame (assuming that we did not initiate the Close in
// the first place). When the Close frame has both been received and sent, the
// underlying connection is automatically closed (as per the WebSocket protocol RFC
// specifications). Thus, if this property is true, then two things automatically
// happen when a Close frame is received: (1) a Close frame is sent in response,
// and (2) the connection is closed.
// 
// The default value of this property is true.
//
func (z *WebSocket) CloseAutoRespond() bool {
    v := int(C.CkWebSocket_getCloseAutoRespond(z.ckObj))
    return v != 0
}
// property setter: CloseAutoRespond
func (z *WebSocket) SetCloseAutoRespond(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkWebSocket_putCloseAutoRespond(z.ckObj,v)
}

// property: CloseReason
// The reason string received with the Close frame, if any.
func (z *WebSocket) CloseReason() string {
    return C.GoString(C.CkWebSocket_closeReason(z.ckObj))
}

// property: CloseReceived
// If true, then a Close frame was already received on this websocket connection.
// If CloseAutoRespond is false, then an application can check this property
// value to determine if a Close frame should be sent in response.
func (z *WebSocket) CloseReceived() bool {
    v := int(C.CkWebSocket_getCloseReceived(z.ckObj))
    return v != 0
}

// property: CloseStatusCode
// The status code received with the Close frame. If no status code was provided,
// or no Close frame has yet been received, then this property will be 0.
func (z *WebSocket) CloseStatusCode() int {
    return int(C.CkWebSocket_getCloseStatusCode(z.ckObj))
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
func (z *WebSocket) DebugLogFilePath() string {
    return C.GoString(C.CkWebSocket_debugLogFilePath(z.ckObj))
}
// property setter: DebugLogFilePath
func (z *WebSocket) SetDebugLogFilePath(goStr string) {
    cStr := C.CString(goStr)
    C.CkWebSocket_putDebugLogFilePath(z.ckObj,cStr)
    C.free(unsafe.Pointer(cStr))
}

// property: FinalFrame
// true if the last data frame received by calling ReadFrame was a final frame.
// Otherwise false.
func (z *WebSocket) FinalFrame() bool {
    v := int(C.CkWebSocket_getFinalFrame(z.ckObj))
    return v != 0
}

// property: FrameDataLen
// The number of bytes accumulated from one or more calls to ReadFrame. Accumulated
// incoming frame data can be retrieved by calling GetFrameData, GetFrameDataSb, or
// GetFrameDataBd.
func (z *WebSocket) FrameDataLen() int {
    return int(C.CkWebSocket_getFrameDataLen(z.ckObj))
}

// property: FrameOpcode
// Indicates the type of frame received in the last call to ReadFrame. Possible
// values are "Continuation", "Text", "Binary", "Close", "Ping", or "Pong".
// Initially this property is set to the empty string because nothing has yet been
// received.
func (z *WebSocket) FrameOpcode() string {
    return C.GoString(C.CkWebSocket_frameOpcode(z.ckObj))
}

// property: FrameOpcodeInt
// The integer value of the opcode (type of frame) received in the last call to
// ReadFrame. Possible values are:
// 0 - Continuation
// 1 - Text
// 2 - Binary
// 8 - Close
// 9 - Ping
// 10 - Pong
func (z *WebSocket) FrameOpcodeInt() int {
    return int(C.CkWebSocket_getFrameOpcodeInt(z.ckObj))
}

// property: IdleTimeoutMs
// The maximum amount of time to wait for additional incoming data when receiving,
// or the max time to wait to send additional data. The default value is 30000 (30
// seconds). This is not an overall max timeout. Rather, it is the maximum time to
// wait when receiving or sending has halted.
func (z *WebSocket) IdleTimeoutMs() int {
    return int(C.CkWebSocket_getIdleTimeoutMs(z.ckObj))
}
// property setter: IdleTimeoutMs
func (z *WebSocket) SetIdleTimeoutMs(value int) {
    C.CkWebSocket_putIdleTimeoutMs(z.ckObj,C.int(value))
}

// property: IsConnected
// Returns true if the websocket is connected. Otherwise returns false.
func (z *WebSocket) IsConnected() bool {
    v := int(C.CkWebSocket_getIsConnected(z.ckObj))
    return v != 0
}

// property: LastErrorHtml
// Provides information in HTML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *WebSocket) LastErrorHtml() string {
    return C.GoString(C.CkWebSocket_lastErrorHtml(z.ckObj))
}

// property: LastErrorText
// Provides information in plain-text format about the last method/property called.
// If a method call returns a value indicating failure, or behaves unexpectedly,
// examine this property to get more information.
func (z *WebSocket) LastErrorText() string {
    return C.GoString(C.CkWebSocket_lastErrorText(z.ckObj))
}

// property: LastErrorXml
// Provides information in XML format about the last method/property called. If a
// method call returns a value indicating failure, or behaves unexpectedly, examine
// this property to get more information.
func (z *WebSocket) LastErrorXml() string {
    return C.GoString(C.CkWebSocket_lastErrorXml(z.ckObj))
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
func (z *WebSocket) LastMethodSuccess() bool {
    v := int(C.CkWebSocket_getLastMethodSuccess(z.ckObj))
    return v != 0
}
// property setter: LastMethodSuccess
func (z *WebSocket) SetLastMethodSuccess(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkWebSocket_putLastMethodSuccess(z.ckObj,v)
}

// property: NeedSendPong
// If true, then a Ping frame was received, but no Pong frame has yet been sent
// in response. The application should send a Pong frame by calling SendPong as
// soon as possible.
func (z *WebSocket) NeedSendPong() bool {
    v := int(C.CkWebSocket_getNeedSendPong(z.ckObj))
    return v != 0
}

// property: PingAutoRespond
// If true, then a Pong frame is automatically sent when a Ping frame is
// received. If set to false, then the application may check the NeedSendPong
// property to determine if a Pong response is needed, and if so, may call the
// SendPong method to send a Pong.
// 
// Note: If this property is true, then the ReadFrame method will auto-consume
// incoming Ping frames. In other words, ReadFrame will continue with reading the
// next incoming frame (thus Ping frames will never be returned to the
// application). This relieves the application from having to worry about receiving
// and handling spurious Ping frames.
// 
// The default value is true.
//
func (z *WebSocket) PingAutoRespond() bool {
    v := int(C.CkWebSocket_getPingAutoRespond(z.ckObj))
    return v != 0
}
// property setter: PingAutoRespond
func (z *WebSocket) SetPingAutoRespond(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkWebSocket_putPingAutoRespond(z.ckObj,v)
}

// property: PongAutoConsume
// If true, then incoming Pong frames are automatically consumed, and a call to
// ReadFrame will continue reading until it receives a non-Pong frame. The
// PongConsumed property can be checked to see if the last ReadFrame method call
// auto-consumed a Pong frame.
// 
// The default value is true.
//
func (z *WebSocket) PongAutoConsume() bool {
    v := int(C.CkWebSocket_getPongAutoConsume(z.ckObj))
    return v != 0
}
// property setter: PongAutoConsume
func (z *WebSocket) SetPongAutoConsume(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkWebSocket_putPongAutoConsume(z.ckObj,v)
}

// property: PongConsumed
// Is true if the last call to ReadFrame auto-consumed a Pong frame. This
// property is reset to false each time a ReadFrame method is called, and will
// get set to true if (1) the PongAutoConsume property is true and (2) a Pong
// frame was consumed within the ReadFrame method.
// 
// The purpose of PongAutoConsume and PongConsumed is to eliminate the concern for
// unanticipated Pong frames in the stream. In the websocket protocol, both sides
// (client and server) may send Pong frames at any time. In addition, if a Ping
// frame is sent, the corresponding Pong response frame can arrive at some
// unanticipated point later in the conversation. It's also possible, if several
// Ping frames are sent, that a Pong response frame is only sent for the most
// recent Ping frame. The default behavior of Chilkat's WebSocket API is to
// auto-consume incoming Pong frames and set this property to true. This allows
// the application to call a ReadFrame method for whatever application data frame
// it may be expecting, without needing to be concerned if the next incoming frame
// is a Pong frame.
//
func (z *WebSocket) PongConsumed() bool {
    v := int(C.CkWebSocket_getPongConsumed(z.ckObj))
    return v != 0
}

// property: ReadFrameFailReason
// If the ReadFrame method returns false, this property holds the fail reason. It
// can have one of the following values:
// 0 - No failure.
// 1 - Read Timeout.
// 2 - Aborted by Application Callback.
// 3 - Fatal Socket Error (Lost Connection).
// 4 - Received invalid WebSocket frame bytes.
// 99 - A catch-all for any unknown failure.  (Should not ever occur.  If it does, contact Chilkat.)
func (z *WebSocket) ReadFrameFailReason() int {
    return int(C.CkWebSocket_getReadFrameFailReason(z.ckObj))
}

// property: UncommonOptions
// This is a catch-all property to be used for uncommon needs. This property
// defaults to the empty string and should typically remain empty. Can be set to a
// list of the following comma separated keywords:
//     "ProtectFromVpn" - Introduced in v9.5.0.80. On Android systems, will bypass
//     any VPN that may be installed or active.
func (z *WebSocket) UncommonOptions() string {
    return C.GoString(C.CkWebSocket_uncommonOptions(z.ckObj))
}

// property: VerboseLogging
// If set to true, then the contents of LastErrorText (or LastErrorXml, or
// LastErrorHtml) may contain more verbose information. The default value is
// false. Verbose logging should only be used for debugging. The potentially
// large quantity of logged information may adversely affect peformance.
func (z *WebSocket) VerboseLogging() bool {
    v := int(C.CkWebSocket_getVerboseLogging(z.ckObj))
    return v != 0
}
// property setter: VerboseLogging
func (z *WebSocket) SetVerboseLogging(value bool) {
    v := C.int(0)
    if value {
            v = C.int(1)
    }
    C.CkWebSocket_putVerboseLogging(z.ckObj,v)
}

// property: Version
// Version of the component/library, such as "9.5.0.63"
func (z *WebSocket) Version() string {
    return C.GoString(C.CkWebSocket_version(z.ckObj))
}
// Adds the required WebSocket client-side open handshake headers. The headers
// specifically added to the previously specified REST object (in the call to
// UseConnection) are:
// Upgrade: websocket
// Connection: Upgrade
// Sec-WebSocket-Key: ...
// Sec-WebSocket-Version: 13
func (z *WebSocket) AddClientHeaders() bool {

    v := C.CkWebSocket_AddClientHeaders(z.ckObj)


    return v != 0
}


// Forcibly closes the underlying connection. This is a non-clean way to close the
// connection, but may be used if needed. The clean way to close a websocket is to
// send a Close frame, and then receive the Close response.
func (z *WebSocket) CloseConnection() bool {

    v := C.CkWebSocket_CloseConnection(z.ckObj)


    return v != 0
}


// Returns the accumulated received frame data as a string. Calling GetFrameData
// clears the internal receive buffer.
// return a string or nil if failed.
func (z *WebSocket) GetFrameData() *string {

    retStr := C.CkWebSocket_getFrameData(z.ckObj)


        if retStr == nil {
            return nil
            }
        retGoStr := C.GoString(retStr)
        return &retGoStr
}


// Returns the accumulated received frame data in a BinData object. The received
// data is appended to the binData.
// 
// Calling this method clears the internal receive buffer.
//
func (z *WebSocket) GetFrameDataBd(arg1 *BinData) bool {

    v := C.CkWebSocket_GetFrameDataBd(z.ckObj, arg1.ckObj)


    return v != 0
}


// Returns the accumulated received frame data in a StringBuilder object. The
// received data is appended to the sb.
// 
// Calling this method clears the internal receive buffer.
//
func (z *WebSocket) GetFrameDataSb(arg1 *StringBuilder) bool {

    v := C.CkWebSocket_GetFrameDataSb(z.ckObj, arg1.ckObj)


    return v != 0
}


// Loads the caller of the task's async method.
func (z *WebSocket) LoadTaskCaller(arg1 *Task) bool {

    v := C.CkWebSocket_LoadTaskCaller(z.ckObj, arg1.ckObj)


    return v != 0
}


// Check to see if data is available for reading on the websocket. Returns true
// if data is waiting and false if no data is waiting to be read.
func (z *WebSocket) PollDataAvailable() bool {

    v := C.CkWebSocket_PollDataAvailable(z.ckObj)


    return v != 0
}


// Reads a single frame from the connected websocket. If a frame was successfuly
// received, then the following properties are set: FrameOpcode, FrameDataLen,
// FinalFrame, and the received frame data can be retrieved by calling
// GetFrameData, GetFrameDataSb, or GetFrameDataBd.
func (z *WebSocket) ReadFrame() bool {

    v := C.CkWebSocket_ReadFrame(z.ckObj)


    return v != 0
}


// Asynchronous version of the ReadFrame method
func (z *WebSocket) ReadFrameAsync(c chan *Task) {

    hTask := C.CkWebSocket_ReadFrameAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a Close control frame. If includeStatus is true, then the statusCode is sent in the
// application data part of the Close frame. A Close reason may be provided only if
// includeStatus is true. If this Close was sent to satisfy an already-received Close
// frame, then the underlying connection will also automatically be closed.
// 
// Note: If a status code and reason are provided, the utf-8 representation of the
// reason string must be 123 bytes or less. Chilkat will automatically truncate the
// reason to 123 bytes if necessary. Also, the status code must be an integer in the
// range 0 to 16383.
// 
// The WebSocket protocol specifies some pre-defined status codes atWebSocket
// Status Codes
// <https://tools.ietf.org/html/rfc6455#section-7.4.1>. For a normal closure, a
// status code value of 1000 should be used. The reason can be any string, as long
// as it is 123 bytes or less.
//
func (z *WebSocket) SendClose(arg1 bool, arg2 int, arg3 string) bool {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    v := C.CkWebSocket_SendClose(z.ckObj, b1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

    return v != 0
}


// Asynchronous version of the SendClose method
func (z *WebSocket) SendCloseAsync(arg1 bool, arg2 int, arg3 string, c chan *Task) {
    b1 := C.int(0)
    if arg1 {
        b1 = C.int(1)
    }
    cstr3 := C.CString(arg3)

    hTask := C.CkWebSocket_SendCloseAsync(z.ckObj, b1, C.int(arg2), cstr3)

    C.free(unsafe.Pointer(cstr3))

    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a single data frame containing a string. If this is the final frame in a
// message, then finalFrame should be set to true. Otherwise set finalFrame equal to false.
func (z *WebSocket) SendFrame(arg1 string, arg2 bool) bool {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkWebSocket_SendFrame(z.ckObj, cstr1, b2)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SendFrame method
func (z *WebSocket) SendFrameAsync(arg1 string, arg2 bool, c chan *Task) {
    cstr1 := C.CString(arg1)
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkWebSocket_SendFrameAsync(z.ckObj, cstr1, b2)

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


// Sends a single data frame containing binary data (the contents of bdToSend). If this
// is the final frame in a message, then finalFrame should be set to true. Otherwise
// set finalFrame equal to false.
func (z *WebSocket) SendFrameBd(arg1 *BinData, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkWebSocket_SendFrameBd(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Asynchronous version of the SendFrameBd method
func (z *WebSocket) SendFrameBdAsync(arg1 *BinData, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkWebSocket_SendFrameBdAsync(z.ckObj, arg1.ckObj, b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a single data frame containing a string (the contents of sbToSend). If this is
// the final frame in a message, then finalFrame should be set to true. Otherwise set
// finalFrame equal to false.
func (z *WebSocket) SendFrameSb(arg1 *StringBuilder, arg2 bool) bool {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    v := C.CkWebSocket_SendFrameSb(z.ckObj, arg1.ckObj, b2)


    return v != 0
}


// Asynchronous version of the SendFrameSb method
func (z *WebSocket) SendFrameSbAsync(arg1 *StringBuilder, arg2 bool, c chan *Task) {
    b2 := C.int(0)
    if arg2 {
        b2 = C.int(1)
    }

    hTask := C.CkWebSocket_SendFrameSbAsync(z.ckObj, arg1.ckObj, b2)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Sends a Ping control frame, optionally including text data. If pingData is
// non-empty, the utf-8 representation of the string must be 125 bytes or less.
// Chilkat will automatically truncate the pingData to 125 bytes if necessary.
func (z *WebSocket) SendPing(arg1 string) bool {
    cstr1 := C.CString(arg1)

    v := C.CkWebSocket_SendPing(z.ckObj, cstr1)

    C.free(unsafe.Pointer(cstr1))

    return v != 0
}


// Asynchronous version of the SendPing method
func (z *WebSocket) SendPingAsync(arg1 string, c chan *Task) {
    cstr1 := C.CString(arg1)

    hTask := C.CkWebSocket_SendPingAsync(z.ckObj, cstr1)

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


// Sends a Pong control frame. If this Pong frame is sent to satisfy an
// unresponded-to Ping frame, then the previously received Ping data is
// automatically sent in this Pong frame.
func (z *WebSocket) SendPong() bool {

    v := C.CkWebSocket_SendPong(z.ckObj)


    return v != 0
}


// Asynchronous version of the SendPong method
func (z *WebSocket) SendPongAsync(c chan *Task) {

    hTask := C.CkWebSocket_SendPongAsync(z.ckObj)


    if hTask == nil {
        c <- NewTask()
        return
        }
    task := Task{hTask}
    c <- &task
    task.RunSynchronously()
    close(c)

}


// Initializes the connection for a WebSocket session. All WebSocket sessions begin
// with a call to UseConnection. A Chilkat REST object is used for the connection
// because the WebSocket handshake begins with an HTTP GET request. The Chilkat
// REST API provides the ability to add custom headers, authentication, etc. to the
// opening GET handshake. It also provides the ability to establish connections
// over TLS or SSH and to benefit from the rich set of features already present
// relating to HTTP proxies, SOCKS proxies, bandwidth throttling, IPv6, socket
// options, etc.
func (z *WebSocket) UseConnection(arg1 *Rest) bool {

    v := C.CkWebSocket_UseConnection(z.ckObj, arg1.ckObj)


    return v != 0
}


// Called after sending the opening handshake from the Rest object. Validates the
// server's response to the opening handshake message. If validation is successful,
// the application may begin sending and receiving data and control frames.
func (z *WebSocket) ValidateServerHandshake() bool {

    v := C.CkWebSocket_ValidateServerHandshake(z.ckObj)


    return v != 0
}


