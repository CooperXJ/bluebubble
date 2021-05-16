package snowflake

import (
	"time"
	"web_app/settings"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

/**
	雪花算法  生成全局唯一Id
	其特点是利用不同的位代表不同的含义  从startTime开始可以使用69年
 */

func Init(cfg *settings.SnowFlakeConfig) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", cfg.StartTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(cfg.MachineId)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
