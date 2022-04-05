package wsbroadcastserver

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type chainedReader struct {
	readers []io.Reader
}

func (cr *chainedReader) Read(b []byte) (n int, err error) {
	for len(cr.readers) > 0 {
		n, err = cr.readers[0].Read(b)
		if errors.Is(err, io.EOF) {
			cr.readers = cr.readers[1:]
			if n == 0 {
				continue // EOF and empty, skip to next
			} else {
				// The Read interface specifies some data can be returned along with an EOF.
				if len(cr.readers) != 1 {
					// If this isn't the last reader, return the data without the EOF since this
					// may not be the end of all the readers.
					return n, nil
				} else {
					return
				}
			}
		}
		break
	}
	return
}

func (cr *chainedReader) add(r io.Reader) *chainedReader {
	if r != nil {
		cr.readers = append(cr.readers, r)
	}
	return cr
}

func logError(err error, msg string) {
	if !strings.Contains(err.Error(), "use of closed network connection") {
		logger.Error().Err(err).Msg(msg)
	}
}

func ReadData(ctx context.Context, conn net.Conn, earlyFrameData io.Reader, idleTimeout time.Duration, state ws.State) ([]byte, ws.OpCode, error) {
	controlHandler := wsutil.ControlFrameHandler(conn, state)
	reader := wsutil.Reader{
		Source:          (&chainedReader{}).add(earlyFrameData).add(conn),
		State:           state,
		CheckUTF8:       true,
		SkipHeaderCheck: false,
		OnIntermediate:  controlHandler,
	}

	// Remove timeout when leaving this function
	defer func(conn net.Conn) {
		err := conn.SetReadDeadline(time.Time{})
		if err != nil {
			logError(err, "error removing read deadline")
		}
	}(conn)

	for {
		select {
		case <-ctx.Done():
			return nil, 0, nil
		default:
		}

		err := conn.SetReadDeadline(time.Now().Add(idleTimeout))
		if err != nil {
			return nil, 0, err
		}

		// Control packet may be returned even if err set
		header, err := reader.NextFrame()
		if header.OpCode.IsControl() {
			// Control packet may be returned even if err set
			if err2 := controlHandler(header, &reader); err2 != nil {
				return nil, 0, err2
			}

			// Discard any data after control packet
			if err2 := reader.Discard(); err2 != nil {
				return nil, 0, err2
			}

			return nil, 0, nil
		}
		if err != nil {
			return nil, 0, err
		}

		if header.OpCode != ws.OpText &&
			header.OpCode != ws.OpBinary {
			if err := reader.Discard(); err != nil {
				return nil, 0, err
			}
			continue
		}

		data, err := ioutil.ReadAll(&reader)

		return data, header.OpCode, err
	}
}
