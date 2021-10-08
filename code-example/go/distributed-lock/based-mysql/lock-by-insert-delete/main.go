package main

import (
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	mclient "github.com/shonn-study/code-example/go/common/mysql/client"
	mconfig "github.com/shonn-study/code-example/go/common/mysql/config"
	"github.com/shonn-study/code-example/go/common/util/logger"
	utilptr "github.com/shonn-study/code-example/go/common/util/ptr"
	"github.com/shonn-study/code-example/go/distributed-lock/client"
	"github.com/shonn-study/code-example/go/distributed-lock/model"
)

func lockByInsertDelete(threadName string) {
	lockCli := client.NewMysqlInsertDeleteLockClient(mclient.New(mconfig.LocalConfig()), &model.MysqlInsterDeleteLock{
		LockName:   utilptr.StringPtr("lockByInsertDelete"),
		LockRemark: utilptr.StringPtr("threadName"),
	})
	defer lockCli.Close()
	err := lockCli.Lock()
	if err != nil {
		logger.Error("lock fail", logger.Field("threadName", threadName), logger.Field("err", err))
		return
	} else {
		logger.Info("lock success", logger.Field("threadName", threadName))
		sleepSec := time.Duration(rand.Intn(5)+1) * time.Second
		time.Sleep(sleepSec)
		logger.Info("sleep", logger.Field("threadName", threadName), logger.Field("sec", sleepSec))
		err = lockCli.UnLock()
		if err != nil {
			logger.Error("unlock fail", logger.Field("threadName", threadName), logger.Field("err", err))
			return
		}
		logger.Info("unlock success", logger.Field("threadName", threadName))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	c := make(chan os.Signal)
	exit := make(chan struct{})
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		for {
			select {
			case val := <-c:
				switch val {
				case os.Interrupt, os.Kill:
					logger.Info("receive exit signal", logger.Field("signal", val))
					close(exit)
					return
				default:
					logger.Info("receive signal unknown", logger.Field("signal", val))
				}
			}
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		threadName := "thread1"
		defer wg.Done()
		for {
			select {
			case <-exit:
				logger.Info("exit", logger.Field("threadName", threadName))
				return
			default:
				lockByInsertDelete(threadName)
				time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
			}
		}
	}()
	go func() {
		threadName := "thread2"
		defer wg.Done()
		for {
			select {
			case <-exit:
				logger.Info("exit", logger.Field("threadName", threadName))
				return
			default:
				lockByInsertDelete(threadName)
				time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
			}
		}
	}()
	wg.Wait()
}
