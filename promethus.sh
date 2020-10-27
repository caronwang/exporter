docker run -d \
    --name prom \
    -p 9090:9090 \
    -v /Users/qiangwang/Desktop/go/src/exporter/conf/promethus/prometheus.yml:/etc/prometheus/prometheus.yml \
    -v /Users/qiangwang/Desktop/go/src/exporter/conf/promethus/fileconfig:/etc/prometheus/fileconfig \
    prom/prometheus:v2.20.1