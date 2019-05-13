package etcd

import (
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/constant"
	"golang.org/x/net/context"
	"log"
)

func AddHost(host model.Host) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	key := constant.Host + constant.Slash + host.Name
	bytes, err := json.Marshal(host)
	if err != nil {
		log.Print(err)
	}
	value := string(bytes)
	_, err = etcdClient.Put(ctx, key, value)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
}

func GetHost(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	key := constant.Host + constant.Slash + name
	resp, err := etcdClient.Get(ctx, key)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func WatchHost() {
	watchChan := etcdClient.Watch(context.Background(), constant.Host, clientv3.WithPrefix())
	for resp := range watchChan {
		for _, event := range resp.Events {
			fmt.Printf("%s %q : %q\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
}
