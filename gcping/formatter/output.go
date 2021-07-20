package formatter

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/maborosii/gcping/pinger"
)

// type define list of ping as Formatter
type Formatter []pinger.Ping

//table format filed name
var FieldNameList = table.Row{"host", "port", "successed", "failed", "success_rate", "minimum", "maximum", "average"}

// raw format output
func (raw Formatter) Setraw() string {
	statistics_group := make([]string, 10)
	for _, row := range raw {

		total, rate := row.Successedrate()
		statistics_header := fmt.Sprintf("\n--- %s[:%d] tcping statistics ---", row.Socket.Host, row.Socket.Port)
		statistics_body := fmt.Sprintf("\n%d connections, %d successed, %d failed, %s success rate", total, row.Successed, row.Failed, rate)

		mininum, maxinum, average := row.Statstics()
		statistics_footer := fmt.Sprintf("\nminimum = %f, maximum = %f, average = %f", mininum, maxinum, average)

		statistics := statistics_header + statistics_body + statistics_footer
		statistics_group = append(statistics_group, statistics)
	}
	return strings.Join(statistics_group, "")
}

// table format output
func (tb Formatter) Settb() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(FieldNameList)
	for _, row := range tb {
		_, rate := row.Successedrate()
		mininum, maxinum, average := row.Statstics()
		t.AppendRow([]interface{}{row.Socket.Host, row.Socket.Port, row.Successed, row.Failed, rate, mininum, maxinum, average})
	}
	t.Render()
}

// add new row
func (stat *Formatter) Addstat(row pinger.Ping) *Formatter {
	*stat = append(*stat, row)
	return stat
}

func NewFormatter() *Formatter {
	return &Formatter{}
}
