package id


import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

// Init initializes the Snowflake node (call once per service)
func Init(nodeID int64) error {
	var err error
	once.Do(func() {
		node, err = snowflake.NewNode(nodeID)
	})
	return err
}

// New generates a new Snowflake ID (uint64)
func New() uint64 {
	if node == nil {
		panic("snowflake not initialized")
	}
	return uint64(node.Generate().Int64())
}
