package store

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	server    = "0.0.0.0:6379"
	password  = "admin"
	indexName = "name"
)

var (
	objectTables = map[string]*ObjectConfig{}
)

func genID(table, id string) string {
	return fmt.Sprintf("%s:%s", table, id)
}

func genIndexKey(table, index, key string) string {
	return fmt.Sprintf("%s.%s:%s", table, index, key)
}

func genPrefixIndexKey(table, index, key string) string {
	return fmt.Sprintf("%s.%s:*%s*", table, index, key)
}

// RedisStore ...
type RedisStore struct {
	pool *redis.Pool
}

// NewRedisStore ...
func NewRedisStore() *RedisStore {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			// if _, err := c.Do("AUTH", password); err != nil {
			// 	c.Close()
			// 	return nil, err
			// }
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return &RedisStore{pool}
}

func registry(objectConfig *ObjectConfig) {
	objectTables[objectConfig.Name] = objectConfig
}

func (rs *RedisStore) lookup(table, index, id string) Object {
	conn := rs.pool.Get()
	defer conn.Close()
	return nil
}

func (rs *RedisStore) get(table, id string) Object {
	conn := rs.pool.Get()
	defer conn.Close()
	return nil
}

func (rs *RedisStore) find(table string, appendResult func(Object), bys ...By) error {
	conn := rs.pool.Get()
	defer conn.Close()

	var results []string
	var err error

	if len(bys) == 0 {
		results, err = scanAll(conn, table)
	} else {
		results, err = scanWithBys(conn, table, bys...)
	}

	if err != nil {
		return err
	}

	ids := make(map[string]struct{})
	for _, result := range results {
		if _, exists := ids[result]; exists {
			continue
		}

		jsonData, err := redis.String(conn.Do("GET", result))
		if err != nil {
			return err
		}

		if oc, ok := objectTables[table]; ok {
			o := oc.Objecter()
			o.SetContent(jsonData)
			appendResult(o)
		}
	}

	return nil
}

func scan(conn redis.Conn, match string) ([]string, error) {
	var (
		err     error
		cursor  int64
		items   []string
		results []string
		values  []interface{}
	)

	for {
		if match == "" {
			match = ":"
		}
		values, err = redis.Values(conn.Do("SCAN", cursor, "match", match))

		if err != nil {
			return nil, err
		}

		_, err = redis.Scan(values, &cursor, &items)
		if err != nil {
			return nil, err
		}

		results = append(results, items...)

		if cursor == 0 {
			break
		}
	}
	return results, nil
}

func scanAll(conn redis.Conn, table string) ([]string, error) {
	return scan(conn, fmt.Sprintf("%s:*", table))
}

func scanWithBys(conn redis.Conn, table string, bys ...By) ([]string, error) {
	var (
		results []string
		items   []string
		err     error
	)

	for _, by := range bys {
		switch v := by.(type) {
		case byName:
			items, err = scan(conn, genPrefixIndexKey(table, indexName, string(v)))
			if err != nil {
				return nil, err
			}
		}
		for item := range items {
			members, err := redis.Strings(conn.Do("SMEMBERS", item))
			if err != nil {
				return nil, err
			}
			results = append(results, members...)
		}
	}
	return results, nil
}

func (rs *RedisStore) create(table string, o Object) error {
	conn := rs.pool.Get()
	defer conn.Close()
	return nil
}

func (rs *RedisStore) update(table string, o Object) error {
	conn := rs.pool.Get()
	defer conn.Close()
	return nil
}

func (rs *RedisStore) delete(table, id string) error {
	conn := rs.pool.Get()
	defer conn.Close()
	return nil
}
