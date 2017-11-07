package snowflake

import (
	"sync"
	"time"
	"errors"
)
//node 最大值
const MAX_NODE = 1023
// seq 最大值
const MAX_SEQ = 4095

const EPOCH = 1288834974657

type SnowFlake struct {
	node     uint16
	seq      uint16
	lastTime int64
}
//construct
func getNewSnowFlake(node uint16) (*SnowFlake, error) {
	var err error
	if node > MAX_NODE {
		err = errors.New("node ")
	}

	return &SnowFlake{
		node:node,
		seq:0,
		lastTime:0,
	}, err

}

// 普通版本，根据当前时间来生成
func (s *SnowFlake)Generate() int64 {
	//加锁
	mu.Lock()
	defer mu.Unlock()
	//毫秒
	now := time.Now().UnixNano() / 1e6

	//时间相同，使用seq区分
	if now == s.lastTime {
		s.seq += 1
		if s.seq == MAX_SEQ {
			s.seq = 0
			//强制等一毫秒
			time.Sleep(1e6)
			now = time.Now().UnixNano() / 1e6
		}

	} else {
		s.seq = 0
	}

	s.lastTime = now
	//生成id
	id := ((s.lastTime - EPOCH) << 22) | int64((s.node << 12 )) | int64(s.seq)
	return id
}
// 特殊版本，可以根据指定的时间，node,seq生成
func (s *SnowFlake)CustomGenerate(givenTime uint64, node uint16, seq uint16) {

}

var instance *SnowFlake
var mu sync.Mutex

//singleton
func GetInstance(node uint16) (*SnowFlake, error) {

	mu.Lock()
	defer mu.Unlock()
	var err error
	if instance == nil {
		instance, err = getNewSnowFlake(node)

	}
	return instance, err
}

