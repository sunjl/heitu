package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

var (
	etcdClient     *clientv3.Client
	requestTimeout time.Duration
	err            error
)

func InitConfig(endpointsIn []string, dialTimeoutIn int, requestTimeoutIn int) {
	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   endpointsIn,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	requestTimeout = 5 * time.Second
}
