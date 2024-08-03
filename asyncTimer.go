package utils

import (
	"context"
	"errors"
	"sync"
	"time"
)

func ChokeWithWaitGroup(fun func(), maTime time.Duration) time.Duration {
	start := time.Now()
	wg := sync.WaitGroup{}
	// 增加阻塞计数器
	wg.Add(1)
	go func() {
		time.Sleep(maTime)
		// 扣减阻塞计数器
		wg.Done()
	}()
	go func() {
		fun()
		// 扣减阻塞计数器
		wg.Done()
	}()
	// 等待阻塞计数器到 0
	wg.Wait()
	d := time.Since(start)
	return d
}

// 利用 time.After 启动了一个异步的定时器，返回一个 channel，当超过指定的时间后，该 channel 将会接受到信号。
// 启动了子协程，函数执行结束后，将向 channel ch 发送结束信号。
// 使用 select 阻塞等待 done 或 time.After 的信息，若超时，输出timeout，若没有超时，则输出done。
func TimerTimeAfter(fun func(), maTime time.Duration) error {
	ch := make(chan struct{}, 1)
	go func() {
		fun()
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		return nil
	case <-time.After(maTime):
		return errors.New("超时")
	}
}

// 利用 context.WithTimeou 返回启动了一个异步的定时器，返回一个 channel，当超过指定的时间后，该 channel 将会接受到信号。
// 启动了子协程，函数执行结束后，将向 channel ch 发送结束信号。
// 使用 select 阻塞等待 done 或 time.Done的信息，若超时，输出timeout，若没有超时，则输出done。
func TimerWithTimeout(fun func(), maTime time.Duration) (time.Duration, error) {
	start := time.Now()

	ch := make(chan string)
	timeout, cancel := context.WithTimeout(context.Background(), maTime)
	defer cancel()
	go func() {
		fun()
		ch <- "done"
	}()
	select {
	case <-ch:
		//logrus.Println(res)
	case <-timeout.Done():
		//logrus.Println("timout", timeout.Err())
	}

	err := timeout.Err()

	d := time.Since(start)

	if err != nil {
		if err.Error() == "context canceled" {
			return d, nil
		}
		if err.Error() == "context deadline exceeded" {
			return d, errors.New("超时")
		}
	}

	return d, nil
}
