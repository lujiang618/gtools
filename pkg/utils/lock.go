package utils

import (
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type DistributedLock struct {
	session *concurrency.Session
	*concurrency.Mutex
}

func NewDistributedLock(endPoints []string, name string) (*DistributedLock, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endPoints,
		DialTimeout: time.Second,
	})
	if err != nil {
		return nil, err
	}

	//创建一个session，并根据业务情况设置锁的ttl
	session, err := concurrency.NewSession(cli, concurrency.WithTTL(3))
	if err != nil {
		return nil, err
	}

	return &DistributedLock{
		session: session,
		Mutex:   concurrency.NewMutex(session, name),
	}, nil

}

func (dl *DistributedLock) Close() {
	fmt.Println("start close.....")
	dl.session.Close()
}
