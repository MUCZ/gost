package gist

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Gist struct {
	CreatedTime time.Time
	Msg         string
}

func (g *Gist) String() string {
	return string(g.Msg)
}

var mutex *sync.Mutex
var AllGist map[uuid.UUID]*Gist

func init() {
	mutex = &sync.Mutex{}
	AllGist = make(map[uuid.UUID]*Gist)
}

func Get(uid string) (*Gist, error) {
	mutex.Lock()
	defer mutex.Unlock()
	uuid, err := uuid.Parse(uid)
	if err != nil {
		return nil, errors.New("invalid uid: " + uid)
	}
	if g, ok := AllGist[uuid]; ok {
		return g, nil
	}
	return nil, errors.New("gist not found, uid: " + string(uid))
}

func Post(msg []byte) (uuid.UUID, error) {
	mutex.Lock()
	defer mutex.Unlock()
	g := &Gist{
		CreatedTime: time.Now(),
		Msg:         string(msg),
	}

	Uid := uuid.New()
	if _, ok := AllGist[Uid]; ok {
		return uuid.Nil, errors.New("gist already exists, uid: " + Uid.String())
	}
	AllGist[Uid] = g
	return Uid, nil
}

func Remove(uid string) error {
	mutex.Lock()
	defer mutex.Unlock()
	uuid, err := uuid.Parse(uid)
	if err != nil {
		return err
	}
	if _, ok := AllGist[uuid]; !ok {
		return errors.New("gist not found, uid: " + string(uid))
	}
	delete(AllGist, uuid)
	return nil
}

func GetAllKeys() []string {
	mutex.Lock()
	defer mutex.Unlock()
	keys := make([]string, 0)
	for k := range AllGist {
		keys = append(keys, k.String())
	}
	return keys
}

func GetAllKV() map[string]*Gist {
	mutex.Lock()
	defer mutex.Unlock()
	kv := make(map[string]*Gist)
	for k, v := range AllGist {
		kv[k.String()] = v
	}
	return kv
}

func Describe(uid string) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()
	uuid, err := uuid.Parse(uid)
	if err != nil {
		return "", err
	}
	if _, ok := AllGist[uuid]; !ok {
		return "", errors.New("gist not found, uid: " + string(uid))
	}
	return "Created At: " + AllGist[uuid].CreatedTime.String() + " : \n" + string(AllGist[uuid].Msg) + "\n", nil
}
