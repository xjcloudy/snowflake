package snowflake

import (
	"testing"
	"sync"
	"fmt"
	"strconv"
	"time"
)

var lc sync.Mutex

func TestGetTime(t *testing.T) {
	timestamp := GetTime(927852876200415233)
	fmt.Printf("timestamp:%d date:%s \n", timestamp, time.Unix(int64(timestamp / 1e3), (timestamp % 100) * 1e7), )
}
func TestGetInstance(t *testing.T) {
	ins, _ := GetInstance(1)
	fmt.Printf("node :%d", ins.GetNode())
}
func TestGetInstanceMulti(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	ins, _ := GetInstance(1)
	bucket := make(map[string]int)

	go func() {

		for i := 1; i < 100; i++ {
			wg.Add(1)
			go func() {
				id := ins.Generate()
				mapkey := strconv.Itoa(int(id))

				lc.Lock()
				defer lc.Unlock()
				if _, ok := bucket[mapkey]; ok {
					bucket[mapkey]++
					fmt.Printf("%d 发现重复 ，重复次数为%d \n", id, bucket[mapkey])

				} else {
					bucket[mapkey] = 1

				}

				wg.Done()
			}()
		}
		wg.Done()

	}()
	wg.Add(1)

	go func() {
		for i := 1; i < 100; i++ {
			wg.Add(1)
			go func() {
				id := ins.Generate()
				mapkey := strconv.Itoa(int(id))
				lc.Lock()
				defer lc.Unlock()
				if _, ok := bucket[mapkey]; ok {
					bucket[mapkey]++
					fmt.Printf("%d 发现重复 ，重复次数为%d \n", id, bucket[mapkey])

				} else {
					bucket[mapkey] = 1

				}
				wg.Done()
			}()
		}
		wg.Done()
	}()

	wg.Wait()

}
