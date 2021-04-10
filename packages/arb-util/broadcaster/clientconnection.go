package broadcaster

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// ClientConnection represents client connection.
// It contains logic of receiving and sending messages.
// That is, there are no active reader or writer. Some other layer of the
// application should call Receive() to read client's incoming message.
type ClientConnection struct {
	io   sync.Mutex
	conn io.ReadWriteCloser

	id            uint
	name          string
	clientManager *ClientManager
}

// Receive reads next message from client's underlying connection.
// It blocks until full message received.
func (cc *ClientConnection) Receive() error {
	req, err := cc.readRequest()
	if err != nil {
		_ = cc.conn.Close()
		return err
	}
	if req == nil {
		// Handled some control message.
		return nil
	}
	switch req.Method {
	case "sequence":
		return cc.writeResultTo(req, Object{
			"lastSequence": "not yet implemented",
		})
	case "rename":
		return cc.writeResultTo(req, nil)
	case "publish":
		req.Params["author"] = cc.name
		req.Params["time"] = timestamp()
		err = cc.clientManager.Broadcast("publish", req.Params)
		if err != nil {
			logger.Warn().Err(err).Msg("error sending publish broadcast")
		}
	default:
		return cc.writeErrorTo(req, Object{
			"error": "not implemented",
		})
	}
	return nil
}

// readRequests reads json-rpc request from connection.
// It takes io mutex.
func (cc *ClientConnection) readRequest() (*Request, error) {
	cc.io.Lock()
	defer cc.io.Unlock()

	h, r, err := wsutil.NextReader(cc.conn, ws.StateServerSide)
	if err != nil {
		return nil, err
	}
	if h.OpCode.IsControl() {
		return nil, wsutil.ControlFrameHandler(cc.conn, ws.StateServerSide)(h, r)
	}

	req := &Request{}
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (cc *ClientConnection) writeErrorTo(req *Request, err Object) error {
	return cc.write(Error{
		ID:    req.ID,
		Error: err,
	})
}

func (cc *ClientConnection) writeResultTo(req *Request, result Object) error {
	return cc.write(Response{
		ID:     req.ID,
		Result: result,
	})
}

func (cc *ClientConnection) writeNotice(method string, params Object) error {
	return cc.write(Request{
		Method: method,
		Params: params,
	})
}

func (cc *ClientConnection) write(x interface{}) error {
	w := wsutil.NewWriter(cc.conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(w)

	cc.io.Lock()
	defer cc.io.Unlock()

	if err := encoder.Encode(x); err != nil {
		return err
	}

	return w.Flush()
}

func (cc *ClientConnection) writeRaw(p []byte) error {
	cc.io.Lock()
	defer cc.io.Unlock()

	_, err := cc.conn.Write(p)

	return err
}
