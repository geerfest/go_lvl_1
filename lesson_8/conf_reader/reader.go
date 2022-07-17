package conf_reader

import (
	"flag"
	"fmt"
	"net/url"
)

type MyUrl url.URL

// String implementing flag.Value
func (myUrl *MyUrl) String() string {

	return myUrl.RawPath
}

// Set implementing flag.Value
func (myUrl *MyUrl) Set(s string) error {
	v, err := url.Parse(s)
	if err != nil {
		return err
	}

	*myUrl = MyUrl(*v)

	return nil
}

type CustomConf struct {
	Port       int
	DbUrl      MyUrl
	DbUserName string
	JaegerUrl  MyUrl
	SentryUrl  MyUrl
	AppId      int
	AppKey     string
}

func (conf *CustomConf) Print() {
	fmt.Println("Config:")
	fmt.Printf("  Port: %d\n", conf.Port)
	fmt.Printf("  DbUrl: %s\n", conf.DbUrl.String())
	fmt.Printf("  DbUserName: %s\n", conf.DbUserName)
	fmt.Printf("  JaegerUrl: %s\n", conf.JaegerUrl.String())
	fmt.Printf("  SentryUrl: %s\n", conf.SentryUrl.String())
	fmt.Printf("  AppId: %d\n", conf.AppId)
	fmt.Printf("  AppKey: %s\n", conf.AppKey)
}

func GetConfig() CustomConf {
	conf := CustomConf{}

	flag.IntVar(&conf.Port, "port", 80, "your application port")
	flag.Var(&conf.DbUrl, "db_url", "url to connect db")
	flag.StringVar(&conf.DbUserName, "db_user_name", "root", "db user name")
	flag.Var(&conf.JaegerUrl, "jaeger_url", "url to connect jaeger")
	flag.Var(&conf.SentryUrl, "sentry_url", "url to connect sentry")
	flag.IntVar(&conf.AppId, "app_id", 0, "your application id")
	flag.StringVar(&conf.AppKey, "app_key", "", "your application key")

	flag.Parse()

	return conf
}
