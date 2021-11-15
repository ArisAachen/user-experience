package queue

import (
	"sync"

	"github.com/ArisAachen/experience/abstract"
	"github.com/ArisAachen/experience/define"
)

type body struct {
	handler abstract.BaseQueueHandler
	msg     define.RequestMsg
}

// body use to store every push data
type node struct {
	// body
	body *body

	// next body
	next *node
}

type queue struct {
	// lock
	lock sync.Mutex

	// head and tail consist for all message
	num  int
	head *node
	tail *node
}

func newQueue() *queue {
	que := &queue{
		head: nil,
		tail: nil,
	}
	return que
}

// elems count
func (que *queue) count() int {
	que.lock.Lock()
	defer que.lock.Unlock()
	n := que.num
	return n
}

// check is is empty
func (que *queue) empty() bool {
	que.lock.Lock()
	defer que.lock.Unlock()
	return que.num == 0
}

// reset queue
func (que *queue) reset() {
	que.lock.Lock()
	defer que.lock.Unlock()
	que.head = nil
	que.tail = nil
	que.num = 0
}

// push body to queue
func (que *queue) push(handler abstract.BaseQueueHandler, msg define.RequestMsg) {
	// lock
	que.lock.Lock()
	defer que.lock.Unlock()

	// create node
	elem := &node{
		body: &body{handler: handler, msg: msg},
	}

	// either head or tail is nil, saying now has no elem
	if que.num == 0 {
		// store elem
		que.head = elem
		que.tail = elem
		// now head is tail, tail is head
		que.head.next = que.tail
		// add count
		que.num = 1
		return
	}

	// if is has already elem
	que.tail.next = elem
	que.tail = elem
	que.num++
}

// push head data, in some case, emergency data should be sent as quick as possible
// so these data should push to head
func (que *queue) pushHead(handler abstract.BaseQueueHandler, msg define.RequestMsg) {
	// lock
	que.lock.Lock()
	defer que.lock.Unlock()

	// create node
	elem := &node{
		body: &body{handler: handler, msg: msg},
	}

	// either head or tail is nil, saying now has no elem
	if que.num == 0 {
		// store elem
		que.head = elem
		que.tail = elem
		// now head is tail, tail is head
		que.head.next = que.tail
		// add count
		que.num = 1
		return
	}

	// if is has already elem
	elem.next = que.head
	que.head = elem
	que.num++
}

// pop one elem
func (que *queue) pop() *body {
	// lock
	que.lock.Lock()
	defer que.lock.Unlock()

	// now has no elem, return nil directly
	if que.num == 0 {
		return nil
	}

	// get head
	elem := que.head
	que.head = que.head.next
	que.num--

	// if queue is empty
	if que.num == 0 {
		que.head = nil
		que.tail = nil
	}
	return elem.body
}
