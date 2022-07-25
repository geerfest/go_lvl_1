package conf_reade_json

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type JsonUrl url.URL

func (j JsonUrl) String() string {

	return j.Scheme + "://" + j.Host
}

func (j *JsonUrl) UnmarshalJSON(bytes []byte) error {
	some := string(bytes)
	some = strings.Trim(some, "\"") // При чтении почему то получалась строка ""https://example.com""
	v, err := url.Parse(some)
	if err != nil {
		return err
	}

	*j = JsonUrl(*v)

	return nil
}

type Conf struct {
	Email string `json:"email"`
	Host  string `json:"host"`
	Jwt   struct {
		AppId     string `json:"appId"`
		AppSecret string `json:"appSecret"`
	} `json:"jwt"`
	PublicUrl JsonUrl `json:"publicUrl"`
}

func (conf *Conf) Print() {
	fmt.Println("Json Config:")
	fmt.Printf("  Email: %s\n", conf.Email)
	fmt.Printf("  Host: %s\n", conf.Host)
	fmt.Printf("  Jwt: { \n")
	fmt.Printf("    AppId: %s\n", conf.Jwt.AppId)
	fmt.Printf("    AppSecret: %s\n", conf.Jwt.AppSecret)
	fmt.Printf("  }\n")
	fmt.Printf("  PublicUrl: %s\n", conf.PublicUrl)
}

func GetConfig(path string) Conf {
	conf := Conf{}

	data, err := os.Open(path)
	if err != nil {
		return conf
	}
	defer data.Close()

	if err = json.NewDecoder(data).Decode(&conf); err != nil {
		panic(err)
	}

	return conf
}
