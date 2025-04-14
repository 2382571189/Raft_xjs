module example.com/m/v2

go 1.18
require (
    course/shardkv v0.0.0
)

replace course/shardkv => /path/to/shardkv
replace github.com/pingcap/go-ycsb => /path/to/local/go-ycsb
