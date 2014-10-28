// PubSub implements a simple Redis backed publish-subscribe client.
package pubsub

import "github.com/garyburd/redigo/redis"

// PubSub is a Redis pub/sub client. It uses a connection pool for
// communicating with Redis.
type Conn struct {
	pool *redis.Pool
}

// Dial connects to the Redis server with the given network and address.
func Dial(network string, address string, idle int, active int) Conn {
	pool := &redis.Pool{
		MaxIdle:   idle,
		MaxActive: active,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, address)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	return Conn{pool: pool}
}

// Publish a message with the given topic.
func (c Conn) Publish(topic string, msg []byte) {
	con := c.pool.Get()
	defer con.Close()
	con.Do("PUBLISH", topic, msg)
}

// Publish a message with the given topic.
func (c Conn) PublishX(topic string, msg []byte) {
	con := c.pool.Get()
	defer con.Close()
	con.Do("PUBLISH", topic, msg)
}

// Subscribe to topic and get messages received for this topic sent to the
// given channel.
func (c Conn) Subscribe(topic string, chn chan<- []byte) {
	con := redis.PubSubConn{Conn: c.pool.Get()}
	con.Subscribe(topic)

	go func(con redis.PubSubConn) {
		for {
			switch v := con.Receive().(type) {
			case redis.Message:
				chn <- v.Data
			case redis.Subscription:
			case error:
				break
			}
		}
		con.Unsubscribe(topic)
	}(con)
}

// Closes the connection to Redis. This closes the connection pool.
func (c Conn) Close() {
	if c.pool == nil {
		return
	}
	c.pool.Close()
}
