package conf

type TransConf struct {
	KafkaConf `ini:"kafka"`
	ESConf `ini:"elasticSearch"`
}

type KafkaConf struct {
	Addr []string `ini:"addr,omitempty,allowshadow"`
	Topic string `ini:"topic"`
}

type ESConf struct {
	Addr []string `ini:"addr,omitempty,allowshadow"`
	ChanMaxSize int `ini:"chanmaxsize"`
}

