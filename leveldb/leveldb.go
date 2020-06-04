package leveldb

import (
	"fmt"
	"sync"

	"github.com/alash3al/redix/kvstore/leveldb"
)

// LevelDB a
type LevelDB struct {
	sync.RWMutex
	// contains filtered or unexported fields
}

var (
	// Leveldb leveldb
	Leveldb *leveldb.LevelDB

	// Errdb error
	Errdb error
)

func init() {

	Leveldb, Errdb = leveldb.OpenLevelDB("/home/sun/Works/go/src/newapp/Cache")
	if Errdb != nil {
		fmt.Println(Errdb)
	}
}
