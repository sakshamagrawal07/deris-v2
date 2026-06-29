package config

var Host string = "0.0.0.0"
var Port int = 7379
var KeysLimit int = 100

var EvictionRatio float64 = 0.40

const MaxClients int = 20000

var EvictionStrategy string = "allkeys-random"
var AOFFile string = "./deris.aof"