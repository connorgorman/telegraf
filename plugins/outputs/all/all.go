package all

import (
	_ "github.com/influxdata/telegraf/plugins/outputs/file"
	_ "github.com/influxdata/telegraf/plugins/outputs/influxdb"
	_ "github.com/influxdata/telegraf/plugins/outputs/influxdb_v2"
	_ "github.com/influxdata/telegraf/plugins/outputs/instrumental"
	_ "github.com/influxdata/telegraf/plugins/outputs/kafka"
	_ "github.com/influxdata/telegraf/plugins/outputs/kinesis"
	_ "github.com/influxdata/telegraf/plugins/outputs/librato"
	_ "github.com/influxdata/telegraf/plugins/outputs/mqtt"
	_ "github.com/influxdata/telegraf/plugins/outputs/nats"
	_ "github.com/influxdata/telegraf/plugins/outputs/nsq"
	_ "github.com/influxdata/telegraf/plugins/outputs/opentsdb"
	_ "github.com/influxdata/telegraf/plugins/outputs/prometheus_client"
	_ "github.com/influxdata/telegraf/plugins/outputs/riemann"
	_ "github.com/influxdata/telegraf/plugins/outputs/riemann_legacy"
	_ "github.com/influxdata/telegraf/plugins/outputs/socket_writer"
	_ "github.com/influxdata/telegraf/plugins/outputs/stackdriver"
	_ "github.com/influxdata/telegraf/plugins/outputs/wavefront"
)
