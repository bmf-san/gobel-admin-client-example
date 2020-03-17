package session

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/bmf-san/gobel-admin-client-example/app/csrf"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

const (
	sessionPrefix   = "session_"
	csrfTokenPrefix = "csrf_"
	inputOldPrefix  = "input_old_"
)

// RedisHandler represents the singular of response.
type RedisHandler struct {
	Conn *redis.Client
}

//NewRedisHandler creates a RedisHandler.
func NewRedisHandler() (*RedisHandler, error) {
	dataSourceName := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	conn := redis.NewClient(&redis.Options{
		Addr:     dataSourceName,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := conn.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisHandler{
		Conn: conn,
	}, nil
}

// SetSession sets a session.
func (r *RedisHandler) SetSession() (string, error) {
	token := uuid.NewV4().String()

	if err := r.Conn.Set(sessionPrefix+token, token, 3600*24*7*time.Second).Err(); err != nil {
		return "", err
	}

	return token, nil
}

// GetSession gets a session
func (r *RedisHandler) GetSession(sessID string) (string, error) {
	token, err := r.Conn.Get(sessionPrefix + sessID).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetLoginSession gets a login session.
func (r *RedisHandler) GetLoginSession(loginSessID string) (string, error) {
	userID, err := r.Conn.Get(loginSessID).Result()
	if err != nil {
		return "", err
	}

	return userID, nil
}

// HasCSRFToken gets a csrf token.
func (r *RedisHandler) HasCSRFToken(token string) error {
	i, err := r.Conn.Exists(csrfTokenPrefix + token).Result()
	if err != nil {
		return err
	}

	if i == 0 {
		return errors.New("CSRF Token is not found")
	}

	return nil
}

// SetCSRFToken sets a csrf token.
func (r *RedisHandler) SetCSRFToken() (string, error) {
	token, err := csrf.GenerateToken()
	if err != nil {
		return "", err
	}

	if err := r.Conn.Set(csrfTokenPrefix+token, token, 30*time.Second).Err(); err != nil {
		return "", err
	}

	return token, err
}

// SetInputOld sets form input values.
func (r *RedisHandler) SetInputOld(inputs url.Values, sessID string, excludes ...string) error {
	key := inputOldPrefix + sessID

	for _, e := range excludes {
		inputs.Del(e)
	}

	for k, v := range inputs {
		if err := r.Conn.HSet(key, k, strings.Join(v, "")).Err(); err != nil {
			return err
		}
	}

	return nil
}

// GetInputOld gets form input values.
func (r *RedisHandler) GetInputOld(sessID string) (map[string]string, error) {
	v, err := r.Conn.HGetAll(inputOldPrefix + sessID).Result()
	if err != nil {
		return nil, err
	}

	return v, nil
}

// DelInputOld deletes form input values.
func (r *RedisHandler) DelInputOld(sessID string) error {
	if err := r.Conn.Del(inputOldPrefix + sessID).Err(); err != nil {
		return err
	}

	return nil
}
