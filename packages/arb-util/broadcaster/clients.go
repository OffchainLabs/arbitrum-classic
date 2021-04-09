package broadcaster

import (
	"bytes"
	"encoding/json"
	"github.com/gobwas/ws-examples/src/gopool"
	"io"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// User represents user connection.
// It contains logic of receiving and sending messages.
// That is, there are no active reader or writer. Some other layer of the
// application should call Receive() to read user's incoming message.
type User struct {
	io   sync.Mutex
	conn io.ReadWriteCloser

	id   uint
	name string
	chat *Chat
}

// Receive reads next message from user's underlying connection.
// It blocks until full message received.
func (u *User) Receive() error {
	req, err := u.readRequest()
	if err != nil {
		_ = u.conn.Close()
		return err
	}
	if req == nil {
		// Handled some control message.
		return nil
	}
	switch req.Method {
	case "rename":
		name, ok := req.Params["name"].(string)
		if !ok {
			return u.writeErrorTo(req, Object{
				"error": "bad params",
			})
		}
		prev, ok := u.chat.Rename(u, name)
		if !ok {
			return u.writeErrorTo(req, Object{
				"error": "already exists",
			})
		}
		_ = u.chat.Broadcast("rename", Object{
			"prev": prev,
			"name": name,
			"time": timestamp(),
		})
		return u.writeResultTo(req, nil)
	case "publish":
		req.Params["author"] = u.name
		req.Params["time"] = timestamp()
		_ = u.chat.Broadcast("publish", req.Params)
	default:
		return u.writeErrorTo(req, Object{
			"error": "not implemented",
		})
	}
	return nil
}

// readRequests reads json-rpc request from connection.
// It takes io mutex.
func (u *User) readRequest() (*Request, error) {
	u.io.Lock()
	defer u.io.Unlock()

	h, r, err := wsutil.NextReader(u.conn, ws.StateServerSide)
	if err != nil {
		return nil, err
	}
	if h.OpCode.IsControl() {
		return nil, wsutil.ControlFrameHandler(u.conn, ws.StateServerSide)(h, r)
	}

	req := &Request{}
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (u *User) writeErrorTo(req *Request, err Object) error {
	return u.write(Error{
		ID:    req.ID,
		Error: err,
	})
}

func (u *User) writeResultTo(req *Request, result Object) error {
	return u.write(Response{
		ID:     req.ID,
		Result: result,
	})
}

func (u *User) writeNotice(method string, params Object) error {
	return u.write(Request{
		Method: method,
		Params: params,
	})
}

func (u *User) write(x interface{}) error {
	w := wsutil.NewWriter(u.conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(w)

	u.io.Lock()
	defer u.io.Unlock()

	if err := encoder.Encode(x); err != nil {
		return err
	}

	return w.Flush()
}

func (u *User) writeRaw(p []byte) error {
	u.io.Lock()
	defer u.io.Unlock()

	_, err := u.conn.Write(p)

	return err
}

// Chat contains logic of user interaction.
type Chat struct {
	mu  sync.RWMutex
	seq uint
	us  []*User
	ns  map[string]*User

	pool *gopool.Pool
	out  chan []byte
}

func NewChat(pool *gopool.Pool) *Chat {
	chat := &Chat{
		pool: pool,
		ns:   make(map[string]*User),
		out:  make(chan []byte, 1),
	}

	go chat.writer()

	return chat
}

// Register registers new connection as a User.
func (c *Chat) Register(conn net.Conn) *User {
	user := &User{
		chat: c,
		conn: conn,
	}

	c.mu.Lock()
	{
		user.id = c.seq
		user.name = strconv.Itoa(int(c.seq))

		c.us = append(c.us, user)
		c.ns[user.name] = user

		c.seq++
	}
	c.mu.Unlock()

	_ = user.writeNotice("hello", Object{
		"name": user.name,
	})
	_ = c.Broadcast("greet", Object{
		"name": user.name,
		"time": timestamp(),
	})

	return user
}

// Remove removes user from chat.
func (c *Chat) Remove(user *User) {
	c.mu.Lock()
	removed := c.remove(user)
	c.mu.Unlock()

	if !removed {
		return
	}

	_ = c.Broadcast("goodbye", Object{
		"name": user.name,
		"time": timestamp(),
	})
}

// Rename renames user.
func (c *Chat) Rename(user *User, name string) (prev string, ok bool) {
	c.mu.Lock()
	{
		if _, has := c.ns[name]; !has {
			ok = true
			prev, user.name = user.name, name
			delete(c.ns, prev)
			c.ns[name] = user
		}
	}
	c.mu.Unlock()

	return prev, ok
}

// Broadcast sends message to all alive users.
func (c *Chat) Broadcast(method string, params Object) error {
	var buf bytes.Buffer

	w := wsutil.NewWriter(&buf, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(w)

	r := Request{Method: method, Params: params}
	if err := encoder.Encode(r); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}

	c.out <- buf.Bytes()

	return nil
}

// writer writes broadcast messages from chat.out channel.
func (c *Chat) writer() {
	for bts := range c.out {
		c.mu.RLock()
		us := c.us
		c.mu.RUnlock()

		toRemove := make([]*User, 0)
		for _, u := range us {
			u := u // For closure.
			c.pool.Schedule(func() {
				err := u.writeRaw(bts)
				if err != nil {
					//logger.Error().Err(err).Str("connection_name", u.name).Msg("Error broadcasting")
					toRemove = append(toRemove, u)
				}
			})
		}

		// Remove any connections that had errors
		for _, u := range toRemove {
			logger.Info().Str("connection_name", u.name).Msg("Removing connection")
			c.remove(u)
		}
	}
}

// mutex must be held.
func (c *Chat) remove(user *User) bool {
	if _, has := c.ns[user.name]; !has {
		return false
	}

	delete(c.ns, user.name)

	i := sort.Search(len(c.us), func(i int) bool {
		return c.us[i].id >= user.id
	})
	if i >= len(c.us) {
		panic("chat: inconsistent state")
	}

	without := make([]*User, len(c.us)-1)
	copy(without[:i], c.us[:i])
	copy(without[i:], c.us[i+1:])
	c.us = without

	return true
}

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
