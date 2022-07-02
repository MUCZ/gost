package gist

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Gist struct {
	Uid         uuid.UUID
	CreatedTime time.Time
	Msg         []byte
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
	uuid := uuid.MustParse(uid)
	if g, ok := AllGist[uuid]; ok {
		return g, nil
	}
	return nil, errors.New("gist not found, uid: " + string(uid))
}

func Post(msg []byte) (uuid.UUID, error) {
	mutex.Lock()
	defer mutex.Unlock()
	g := &Gist{
		Uid:         uuid.New(),
		CreatedTime: time.Now(),
		Msg:         msg,
	}

	if _, ok := AllGist[g.Uid]; ok {
		return uuid.Nil, errors.New("gist already exists, uid: " + g.Uid.String())
	}
	AllGist[g.Uid] = g
	return g.Uid, nil
}

func Remove(uid string) error {
	mutex.Lock()
	defer mutex.Unlock()
	uuid := uuid.MustParse(uid)
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

func Describe(uid string) string {
	mutex.Lock()
	defer mutex.Unlock()
	uuid := uuid.MustParse(uid)
	if _, ok := AllGist[uuid]; !ok {
		return "gist not found, uid: " + string(uid)
	}
	return "Created At: " + AllGist[uuid].CreatedTime.String() + " : \n" + string(AllGist[uuid].Msg)
}
