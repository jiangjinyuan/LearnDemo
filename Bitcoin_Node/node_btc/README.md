## 部署
### build
```shell
docker build -t bitcoind:v23.0 . --build-arg VERSION="23.0"
```

### run
```shell
 docker run -d -p 8332:8332 -p 8333:8333 -p 18332:18332 -p 18333:18333 -v bitcoin_data:/bitcoin/.bitcoin --name=node_bitcoind bitcoind:v23.0 supervisord -n -c /etc/supervisord.conf
```

### 验证
```shell
curl --user 123456 --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getblockcount", "params": [] }' -H 'content-type: text/plain;' http://localhost:18332/
```