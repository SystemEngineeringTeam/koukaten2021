package camera

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type guest struct {
	People int    `json:"people"`
	Status string `json:"status"`
}

//Getで返す関数
func StartCamera() (int, error) { //カメラを起動して人数を返す関数
	out, err := http.Get("http://host.docker.internal:8082/")
	if err != nil {
		fmt.Println("Can't run detect.py\n", err)
		return -1, err
	}
	defer out.Body.Close()

	jsonBytes, err := ioutil.ReadAll(out.Body)
	if err != nil {
		fmt.Println("Can't catch people(io error)\n", err)
		return -1, err
	}

	var guest guest

	if err := json.Unmarshal(jsonBytes, &guest); err != nil {
		fmt.Println("Can't catch people(JSON Unmarshal error)\n", err)
		return -1, err
	}
	return guest.People, nil
}