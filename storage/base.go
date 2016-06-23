package storage

type Store interface {
	Get(key string) *Service
	Set(key string, addr string, port int)
	List(label string) []*Service
}

type Service struct {
	Label string `json:"label"` // 用于负载均衡，若一个服务允许负载均衡，则设message的loadbalance＝true
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Port  int    `json:"port"`
}

type BaseStore struct {
}
