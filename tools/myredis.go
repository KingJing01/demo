package tools

import (
	"log"
	"time"

	redis "github.com/chasex/redis-go-cluster"
)

var Globalcluster *redis.Cluster

//初始化 redis
func InitRedis() {
	cluster, err := redis.NewCluster(
		&redis.Options{
			//StartNodes: []string{"10.10.10.208:7010", "10.10.10.208:7011", "10.10.10.197:7012"},
			//StartNodes:   []string{"192.168.2.109:7001", "192.168.2.109:7002", "192.168.2.109:7003"},
			StartNodes:   []string{"10.10.10.208:7010", "10.10.10.208:7011"},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	if err != nil {
		cluster.Do("SET", "foo", "bar")
		log.Fatalf("redis.New error: %s", err.Error())
	}
	Globalcluster = cluster
}
