package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"penna/config"
)

var snode *snowflake.Node

func Generate() (int64, error) {
	if snode == nil {
		node, err := snowflake.NewNode(config.Config().Server.Node)
		if err != nil {
			return 0, err
		}
		snode = node
	}

	return snode.Generate().Int64(), nil
}
