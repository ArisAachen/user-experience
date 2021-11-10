package queue

import "sync"

// body use to store every push data
type body struct {
	handler BaseQueueHandler
	msg     string
}

type queue struct {
	// pool is sync safe, dont need to use lock
	pool sync.Pool
}

