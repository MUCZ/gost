package client

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	serverAddr = "http://localhost:8080"
)

func List() (string, error) {
	resp, err := http.Get(serverAddr + "/gist")

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Get(uid string) (string, error) {
	resp, err := http.Get(serverAddr + "/gist/" + uid)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return string(body), errors.New("gist not found, uid: " + uid)
	} else {
		return string(body), nil
	}

}

func Post(msg string) (string, error) {
	buf := bytes.NewBuffer([]byte(msg))

	resp, err := http.Post(serverAddr+"/gist", "text/plain", buf)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return string(body), errors.New("post failed, status code: " + resp.Status)
	} else {
		return string(body), nil
	}
}

func Describe(uid string) (string, error) {
	resp, err := http.Get(serverAddr + "/gist/" + uid + "/describe")

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(string(body))
	} else {
		return string(body), nil
	}
}

func Delete(uid string) (string, error) {
	req, err := http.NewRequest("DELETE", serverAddr+"/gist/"+uid, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("delete failed, status code: " + string(body))
	} else {
		return string(body), nil
	}

}

func Check() (bool, error) {
	resp, err := http.Get(serverAddr + "/health")

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	return string(body) == "OK", nil

}
