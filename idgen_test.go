package idgen_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/johntech-o/idgen"
)

func TestGenId(t *testing.T) {
	t.Log("time now: ", time.Now().UnixNano()/1000000)
	version := 0
	shardId := 1000
	if (idgen.SetVersion(version) != nil) || (idgen.SetShardId(shardId) != nil) {
		t.Error("set version | set shardId failed")
	}
	id := idgen.GenId()
	if idgen.GetVersion(id) != version || idgen.GetShardId(id) != shardId {
		t.Error("id generation error", idgen.GetVersion(id), idgen.GetShardId(id))
	}
	hash := map[uint64]int{}
	st := time.Now()
	count := 1000000
	for i := 0; i < count; i++ {
		id := idgen.GenId()
		if _, ok := hash[id]; ok {
			hash[id]++
		} else {
			hash[id] = 1
		}
		if i == count-1 {
			t.Log("example id is: ", id)
		}
	}
	t.Log("gen id amount: "+strconv.Itoa(count)+" timecost: ", time.Now().Sub(st))
	if len(hash) != count {
		for k, v := range hash {
			if v > 1 {
				t.Error("idgen reduplicate id ", k, v, idgen.GetSequence(k), idgen.GetTimeUnixNano(k))
				break
			}
		}
	}

	t.Log("success")
}

// version 0 shardid 100 sequence 3685 timeMillSecond+epoch 1533549445093000
var testId uint64 = 10390179810917
var testTimeNano int64 = 1533549445093000000

func TestGetVersion(t *testing.T) {
	version := idgen.GetVersion(testId)
	if version != 0 {
		t.Error("getVersion failed", version)
		return
	}
	t.Log("get version success", version)
}

func TestGetShardId(t *testing.T) {
	shardId := idgen.GetShardId(testId)
	if shardId != 1000 {
		t.Error("get shardid failed")
		return
	}
	t.Log("get shardId success", shardId)
}

func TestGetTime(t *testing.T) {
	time := idgen.GetTime(testId)
	if time.UnixNano() != testTimeNano {
		t.Error("get time failed expect 1533549445093000000 got", time)
		return
	}
	t.Log("get time success", time)
}

func TestGetTimeUnixNano(t *testing.T) {
	timeUnixNano := idgen.GetTimeUnixNano(testId)
	if timeUnixNano != testTimeNano {
		t.Error("get Time Unix Nano failed expect 1533549445093000000 got: ", timeUnixNano)
		return
	}
	t.Log("get time unix success", timeUnixNano)
}

func TestGetSequence(t *testing.T) {
	sequence := idgen.GetSequence(testId)
	if sequence != 3685 {
		t.Error("get sequence error expect 1 got sequence: ", sequence)
	}
}
