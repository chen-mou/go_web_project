from centos:centos7
copy server server
copy config ./config
cmd chmod 700 server && nohup ./server --env=dev >log.log &