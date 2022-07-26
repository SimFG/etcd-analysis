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
## Function List
### distribute
View data distribution in etcd.
```shell
$ etcdctl+ distribute
```

### look
Get all data in etcd, and you can use system tools to search.
```shell
$ etcdctl+ look | more
```
- Case 1:  Get all the data continuously and display it on the console
```shell
$ etcdctl+ look --write-out=file --hang=true

# New Terminal
$ vim analysis.txt
# update the file in vim, using `:e`
```