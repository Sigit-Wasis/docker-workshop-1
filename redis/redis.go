package redis

import (
	"log"
	"os"

	"github.com/keimoon/gore"
)

// User Struct is the model for the app
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

const (
	localDial = "localhost:6379"
	redisURL  = "REDIS_URL"
	redis     = "redis"
	users     = "users"
)

var (
	dockerDial = os.Getenv(redisURL)
)

// Users is a slice of User
type Users []User

var connect *gore.Conn

// SetKey sets a redis key
func SetKey(key string, data interface{}) {
	connect := Connect()
	if data == nil {
		log.Println("should not be nil")
	}
	rep, err := gore.NewCommand("SET", key, data).Run(connect)
	if err != nil || !rep.IsOk() {
		log.Fatal(err, "not ok")
	}
	defer connect.Close()
}

// GetKey returns a redis value
func GetKey(key string) (*gore.Reply, error) {
	connect := Connect()
	reply, err := gore.NewCommand("GET", key).Run(connect)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// Connect returns redis connection
func Connect() *gore.Conn {
	var err error
	var dial string
	if redisURL == redis {
		dial = dockerDial
	} else {
		dial = localDial
	}
	connect, err = gore.Dial(dial)
	if err != nil {
		return nil
	}
	return connect
}
