package mgserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func (e *Engine) WithETCD(id, address string) *Engine {
	e.ETCDAddress = append(e.ETCDAddress, address)
	e.ETCDID = id
	return e
}

func (e *Engine) RegisterETCD(port string) error {

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{e.ETCDAddress[0]}, // Change to your etcd endpoint
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return err
	}

	defer etcdClient.Close()
	serverIP, err := getLocalIP()

	if err != nil {
		log.Fatalf("Failed to get local IP: %v", err)
	}

	if err := saveIPToEtcd(etcdClient, "server/"+e.ETCDID, serverIP+port); err != nil {
		log.Fatalf("Failed to save IP to etcd: %v", err)
	}

	return err
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no valid IP address found")
}

func saveIPToEtcd(client *clientv3.Client, key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Put(ctx, key, value)

	if err != nil {
		return fmt.Errorf("failed to put key-value in etcd: %w", err)
	}

	log.Printf("Successfully saved to ETCD %s: %s to etcd", key, value)
	return nil
}
