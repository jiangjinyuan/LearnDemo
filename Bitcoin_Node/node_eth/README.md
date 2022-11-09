# ethereum 文档

## 配置要求

Minimum:

CPU with 2+ cores
4GB RAM
1TB free storage space to sync the Mainnet
8 MBit/sec download Internet service
Recommended:

Fast CPU with 4+ cores
16GB+ RAM
High Performance SSD with at least 1TB free space
25+ MBit/sec download Internet service

## 配置说明

1. 现在默认为快照同步

## 部署

在项目目录的时候的 docker 编译命令：

```
docker build . -f Dockerfile -t geth --builf-arg VERSION=v1.10.25
```

```
docker build . -f Dockerfile -t prysm
```


## 运行

～～～～重要～～～～

### 生成 jwt secret 的方式

```
openssl rand -hex 32 | tr -d "\n" | sudo tee ./jwtsecret
sudo chmod +r ./jwtsecret
```

### docker compose

```
docker-compose up -d
```

lodestar 客户端信标链配置说明

--execution.urls ，执行层的 url，可配置多个，但目前建议 1 个，注意这个要配置是 8551 端口的，用来运行 engine 的 API 的

--eth1.providerUrls，执行层的 rpc url，可配置多个，这个是 8545 的端口，如果是 docker-compose 是 geth 服务名

需要及时查看 beacon 链的同步日志，以确定跟浏览器一致。

beacon 浏览器：https://beaconscan.com/

### 执行层 docker 快速启动

```
docker run -d --name ethereum-node -v /data/ethereum:/root \
           -p 8545:8545 -p 30303:30303 \
           geth
```

copy 二进制

## The Merge 合并

这个过程称之为 Pow 共识要向 Pos 转变，geth 之后作为执行层，beacon 链要作为共识层。

合并的 EIP：

https://eips.ethereum.org/EIPS/eip-3675

https://eips.ethereum.org/EIPS/eip-4399

### 时间点

8.22，客户端发布新版本

9.6，Bellatrix 升级

触发合并的 Terminal Total Difficulty 值

![](https://www.ethereum.cn/static/d8305048645eb50421784e5f1cd79180/62feafd3d92d6ac32e7bc7b7342dc6a1.png)

### 共识层客户端

共识层

```
名字	版本
Lighthouse	v3.0.0
Lodestar	v1.0.0
Nimbus	v22.8.0
Prysm	v3.0.0
Teku	22.8.1
```

docker-compose 里主要是采用 Lodestar 作为共识层，也可以切换到其他共识层客户端。

以 docker-compose 的方式的话，共识层直接通过 服务名 来找到对应的服务。

https://github.com/ChainSafe/lodestar

### 共识层同步优化

Optimistic Sync 在共识规范仓库的 [sync](https://github.com/ethereum/consensus-specs/blob/dev/sync/optimistic.md) 文件夹里说明了，它被共识层用于在执行层客户端同步时导入区块，并给执行层提供共识层链头的部分视域。

同时 infura 也提供了一些共识层优化同步的 checkpoint。

TESTNET USERS: The URL for a Prater endpoint --checkpointSyncUrl should end in @eth2-beacon-prater.infura.io

The Ethereum Foundation has also setup public testnet endpoints for weak subjectivity syncing:

Goerli/Prater: https://goerli.checkpoint-sync.ethdevops.io
Ropsten: https://ropsten.checkpoint-sync.ethdevops.io
Sepolia: https://sepolia.checkpoint-sync.ethdevops.io

### Fork Choice Rule

https://medium.com/taipei-ethereum-meetup/ethereum-casper-fork-choice-rule-%E4%B9%8Bghost-%E8%88%87-rpj-34c902083491