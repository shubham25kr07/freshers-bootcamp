package Redis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"go-day-4-5/Config"
)

//
//var (
//	pool        *redis.Pool
//	redisServer = flag.String("redisServer", ":6379", "")
//)
//
//func init() {
//	flag.Parse()
//	pool = Config.NewPool(*redisServer)
//}

const (
	lockScript = `
		return redis.call('SET', KEYS[1], ARGV[1], 'NX', 'PX', ARGV[2])
	`
	unlockScript = `
		if redis.call("get",KEYS[1]) == ARGV[1] then
		    return redis.call("del",KEYS[1])
		else
		    return 0
		end
	`
)

func Lock(key, value string, timeoutMs int, c chan<- int) error {
	r := Config.Pool.Get()
	defer r.Close()

	cmd := redis.NewScript(1, lockScript)
	if _, err := cmd.Do(r, key, value, timeoutMs); err != nil {
		return err
	} else {
		c <- 1
		return nil
	}
}

func Unlock(key, value string, c <-chan int) error {
	r := Config.Pool.Get()
	defer r.Close()

	cmd := redis.NewScript(1, unlockScript)
	if res, err := redis.Int(cmd.Do(r, key, value)); err != nil {
		return err
	} else if res != 1 {
		return errors.New("Unlock failed, key or secret incorrect")
	}
	<-c
	return nil
}
