package clientutil

import (
	"log"
	"sync"

	nmsg "github.com/op/go-nanomsg"
	"github.com/opentarock/service-api/go/proto"
)

type ReqClient struct {
	socket *nmsg.ReqSocket
	lock   *sync.Mutex
}

func NewReqClient() *ReqClient {
	socket, err := nmsg.NewReqSocket()
	if err != nil {
		log.Panicf("Error creating req socket: %s", err)
	}
	return &ReqClient{
		socket: socket,
		lock:   new(sync.Mutex),
	}
}

type Request struct {
	done       chan *proto.Message
	cancel     func()
	cancelOnce *sync.Once
	err        error
	lock       *sync.RWMutex
}

func newRequest(f func(done chan<- *proto.Message) error, cancel func()) *Request {
	done := make(chan *proto.Message)
	req := &Request{
		done:       done,
		cancel:     cancel,
		cancelOnce: new(sync.Once),
		lock:       new(sync.RWMutex),
	}
	go func() {
		req.setErr(f(done))
	}()
	return req
}

func (r *Request) setErr(err error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.err = err
}

func (r *Request) Cancel() {
	r.cancelOnce.Do(r.cancel)
}

func (r *Request) Err() error {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.err
}

func (r *Request) Done() <-chan *proto.Message {
	return r.done
}

func (s *ReqClient) Request(
	request proto.ProtobufMessage,
	headers ...proto.ProtobufMessage) (*Request, error) {

	msg, err := proto.MarshalHeaders(request, headers)
	if err != nil {
		return nil, err
	}

	s.lock.Lock()

	req := newRequest(func(done chan<- *proto.Message) error {
		err = s.send(msg)
		if err != nil {
			return err
		}
		responseMsg, err := s.recv()
		if err != nil {
			return err
		}
		done <- responseMsg
		return nil
	}, func() {
		s.lock.Unlock()
	})

	return req, nil
}

func (c *ReqClient) Connect(address string) error {
	endpoint, err := c.socket.Connect(address)
	if err != nil {
		return err
	}
	log.Printf("Connected to endpoint: %s", endpoint.Address)
	return nil
}

func (s *ReqClient) send(msg *proto.Message) error {
	packed, err := msg.Pack()
	if err != nil {
		return err
	}
	_, err = s.socket.Send(packed, 0)
	return err
}

func (s *ReqClient) recv() (*proto.Message, error) {
	responseData, err := s.socket.Recv(0)
	if err != nil {
		return nil, err
	}
	responseMsg, err := proto.Parse(responseData)
	if err != nil {
		return nil, err
	}
	return responseMsg, err
}

func (s *ReqClient) Close() error {
	return s.socket.Close()
}
