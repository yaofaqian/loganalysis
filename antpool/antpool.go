package antpool

import "github.com/panjf2000/ants/v2"

var P *ants.Pool

func init() {
	// 根据你程序最大承受量  设置goruntime大小
	pool, _ := ants.NewPool(500, ants.WithPreAlloc(true))
	P = pool
}
