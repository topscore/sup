package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/topscore/sup/crypt"

	"github.com/andybons/hipchat"
	"github.com/garyburd/redigo/redis"
)

var RedisURL string
var Config ConfigType

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ConfigType struct {
	TwilioSID        string   `gforms:"twilio_sid"`
	TwilioAuthToken  string   `gforms:"twilio_auth_token"`
	URL              string   `gforms:"url"`
	CallFrom         string   `gforms:"call_from"`
	Phones           []string `gforms:"phones"`
	HipchatAuthToken string   `gforms:"hipchat_auth_token"`
	HipchatRoom      string   `gforms:"hipchat_room"`
}

type StatusType struct {
	Disabled   bool
	LastStatus int
	LastRunAt  time.Time
	NumErrors  int
}

func HipchatMessage(message string) {
	if Config.HipchatAuthToken == "" || Config.HipchatRoom == "" {
		return
	}

	client := hipchat.NewClient(Config.HipchatAuthToken)

	err := client.PostMessage(hipchat.MessageRequest{
		RoomId:        Config.HipchatRoom,
		From:          "SUP",
		Message:       message,
		Color:         hipchat.ColorRed,
		MessageFormat: hipchat.FormatText,
		Notify:        true,
	})

	check(err)
}

func LoadConfig(configFile string, key []byte) ConfigType {
	file, err := ioutil.ReadFile(configFile)
	check(err)

	if key != nil && len(key) > 0 {
		file, err = crypt.Decrypt(key, file)
		check(err)
	}

	var conf ConfigType
	json.Unmarshal(file, &conf)
	return conf
}

func GetStatus() StatusType {
	var status StatusType

	c, err := getRedis()
	check(err)
	defer c.Close()

	statusData, err := redis.Bytes(c.Do("GET", "sup:status"))
	if err != nil {
		if err.Error() == "redigo: nil returned" {
			statusData = []byte("")
		} else {
			check(err)
		}
	}

	json.Unmarshal(statusData, &status)

	return status
}

func SetStatus(status StatusType) {
	json, err := json.Marshal(status)
	check(err)

	c, err := getRedis()
	check(err)
	defer c.Close()

	_, err = c.Do("SET", "sup:status", json)
	check(err)
}

func getRedis() (redis.Conn, error) {
	urlParts, err := url.Parse(RedisURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing Redis url: %s", err)
	}

	auth := ""
	if urlParts.User != nil {
		if password, ok := urlParts.User.Password(); ok {
			auth = password
		}
	}

	c, err := redis.Dial("tcp", urlParts.Host)
	if err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %s", err)
	}

	if len(auth) > 0 {
		_, err = c.Do("AUTH", auth)
		if err != nil {
			return nil, fmt.Errorf("error authenticating with Redis: %s", err)
		}
	}

	if len(urlParts.Path) > 1 {
		db := strings.TrimPrefix(urlParts.Path, "/")
		_, err = c.Do("SELECT", db)
		if err != nil {
			return nil, fmt.Errorf("error selecting Redis db: %s", err)
		}
	}

	return c, nil
}
