package leveldb

import (
	"fmt"
	utils "imovie/utils"
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

	Leveldb, Errdb = leveldb.OpenLevelDB(utils.GetDBPath("level", false))
	if Errdb != nil {
		fmt.Println(Errdb)
	}
}
