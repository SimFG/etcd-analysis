# etcd-analysis
etcd is generally used to store system metadata or service discovery, and is suitable for storing small key-value pairs. At the same time, etcd is sensitive to the size of key-value pairs. When storing large key-value pairs, if the number is too large, it will bring many adverse effects, such as the stability of the watch function is reduced, and a large amount of memory is occupied.

When testing the stability of the system, it may be necessary to pay attention to the size distribution of the data currently stored in etcd by the system. That's why this project came about.

# Getting started
## Getting the source code
Clone this code repository
```shell
$ git clone https://github.com/SimFG/etcd-analysis.git
```
## Build
Compile code into executable
```shell
$ go build -o etcdctl+
```
## Usage
Get help with functions
```shell
$ etcdctl+ distribute -h
```
### Auto complete config
Download the [etcdctl+.ts](ts/etcdctl+.ts), and [config it](https://simfg.github.io/fig). You can also use the [etcdctl.ts](https://simfg.github.io/etcdctl.ts) config the `etcdctl` command.

## Function List

1. **distribute** View data distribution according to data size
2. **look** Show or export all the etcd data, and be used with terminal or loki
3. **find** Get key based on certain characters
4. **leader** Get the leader node info
5. **clear** Clear all the etcd data
6. **decode** decode the etcd value that is encoded 

### distribute
View data distribution in etcd according to the `key` size , `value` size or `key + value` size by setting the `type` command param.
```shell
$ etcdctl+ distribute

Summary:
  Count:        116.
  Total:        7.3 KB.
  Smallest:     22.0 B.
  Largest:      85.0 B.
  Average:      64.0 B.

Size histogram:
  22.0 B [1]    |
  34.0 B [6]    |∎∎∎
  46.0 B [29]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  58.0 B [13]   |∎∎∎∎∎∎∎
  70.0 B [1]    |
  85.0 B [66]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎

Size distribution:
  10% in 38.0 B.
  25% in 39.0 B.
  50% in 76.0 B.
  75% in 83.0 B.
  90% in 85.0 B.
```

### look
Get all data in etcd, and you can use system tools to search.
```shell
$ etcdctl+ look | more

Current Stage
  cluster_id:2037210783374497686 member_id:13195394291058371180 revision:254946 raft_term:9 
Kv List
| Key | Value | CreateRevision | ModRevision | Version | Lease |
| by-dev/kv/gid/idTimestamp | - | 253775 | 254802 | 12 | 0 |
```
- Case 1:  Get all data continuously and display it on the console
```shell
$ etcdctl+ look --write-out=file --hang=true

# New Terminal
$ vim analysis.txt
# update the file in vim, using `:e`
```
- Case 2:  Get the kv data of the specified size range.
```shell
$ etcdctl+ look --filter=key --filter-min=74 --filter-max=100

Current Stage
  cluster_id:2037210783374497686 member_id:13195394291058371180 revision:326021 raft_term:14 
Kv List
| Key | Value | CreateRevision | ModRevision | Version | Lease |
| by-dev/meta/channelwatch/-9223372036854775808/by-dev-rootcoord-dml_4_435191634150817793v0 | - | 326013 | 326013 | 1 | 0 |
```

### find
Get key based on certain characters
```shell
$ ./etcdctl+ find --key=index
Kv List
| Key | Value |
| by-dev/meta/field-index/438660758500016136/438660903999573339 |  |
| by-dev/meta/segment-index/438660758500016136/438660758500016137/438660758500216145/438660903999573340 |  |
| by-dev/meta/segment-index/438660758500016136/438660758500016137/438660758500216146/438660903999573341 |  |
| by-dev/meta/segment-index/438660758500016136/438660758500016137/438672571137461581/438672571137461597 |  |
| by-dev/meta/segment-index/438660758500016136/438660758500016137/438672571137461582/438672571137461598 |  |
```

### leader
Get the leader node info
```shell
$ etcdctl+ leader

Name: default
ClientUrls: [http://127.0.0.1:2379]
```

### clear
Clear all etcd data
```shell
$ etcdctl+ clear

Clear All Data, (Y/n):
```