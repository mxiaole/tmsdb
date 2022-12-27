package db

import "sync"

// Metrics 时序数据
type Metrics struct {
	MetricName string   `json:"metric_name"`
	Labels     []Label  `json:"labels"`
	Samples    []Sample `json:"sample"`
}

// Label 标签
type Label struct {
	K string `json:"k"`
	V string `json:"v"`
}

// Sample 样本值
type Sample struct {
	Timestamp int64       `json:"timestamp"`
	Value     interface{} `json:"value"`
}

// TODO 存在并发问题
// TODO 数据如何持久化

type DB struct {
	hashMap map[string][]Sample
	lock    *sync.RWMutex
}

func New() *DB {
	return &DB{
		hashMap: make(map[string][]Sample, 0),
		lock:    &sync.RWMutex{},
	}
}

func (db *DB) Save(metricName string, labels []Label, values []Sample) {

	// 指标名称
	name := metricName
	// 样本值
	sample := values

	for _, l := range labels {
		// 如果不存在就创建
		k := name + l.K + l.V
		db.lock.Lock()
		if _, ok := db.hashMap[name+l.K+l.V]; !ok {
			db.hashMap[k] = sample
		} else {
			db.hashMap[k] = append(db.hashMap[k], sample...)
		}
		db.lock.Unlock()
	}
}

func (db *DB) Query(metricName string, labels []Label) []Sample {

	name := metricName

	respData := make([]Sample, 0)

	// TODO 遍历查询太慢怎么优化
	// 思路: 倒排索引
	for _, l := range labels {
		key := name + l.K + l.V
		db.lock.RLock()
		if v, ok := db.hashMap[key]; ok {
			respData = append(respData, v...)
		}
		db.lock.RUnlock()
	}

	return respData
}
