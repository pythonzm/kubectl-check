package table

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

var headers = []string{"KIND", "KIND_NAME", "FIELD", "FILED_VALUE"}

func GenTable(data [][]string) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader(headers)

	for _, value := range data {
		t.Append(value)
	}
	t.Render()
}
