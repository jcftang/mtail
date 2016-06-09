package exporter

import (
	"expvar"
	"flag"
	"fmt"

	"github.com/google/mtail/metrics"
)

var (
	logentriesHostPort = flag.String("logentries_hostport", "data.logentries.com:10000",
		"Host:port to logentries server to write metrics to.")
	logentriesToken = flag.String("logentries_token", "",
		"Log token to send data to")

	logentriesExportTotal   = expvar.NewInt("logentries_export_total")
	logentriesExportSuccess = expvar.NewInt("logentries_export_success")
)

func metricToLogentries(hostname string, m *metrics.Metric, l *metrics.LabelSet) string {
	m.RLock()
	defer m.RUnlock()
	var t string
	switch m.Kind {
	case metrics.Counter:
		t = "c" // logentries Counter
	case metrics.Gauge:
		t = "g" // logentries Gauge
	case metrics.Timer:
		t = "ms" // logentries Timer
	}
	return fmt.Sprintf("%s %s.%s=%d|%s\n",
		*logentriesToken,
		m.Program,
		formatLabels(m.Name, l.Labels, ".", "."),
		l.Datum.Get(), t)
}
