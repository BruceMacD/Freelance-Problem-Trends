package visualizer

import (
	"os"

	"github.com/BruceMacD/Freelance-Problem-Trends/analyzer"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

// DrawTopNOccuranceBarChart draws a bar chart of the top 10 values in a word list by occurances
func DrawTopNOccuranceBarChart(wl analyzer.WordList, title, file string, n int) {
	var items []chart.Value

	for i := 1; i <= n; i++ {
		items = append(items, chart.Value{
			Value: float64(wl[i].Occurances),
			Label: wl[i].Value})
	}

	graph := chart.BarChart{
		Title: title,
		Background: chart.Style{
			FillColor: drawing.ColorBlue,
		},
		Height:   512,
		BarWidth: 60,
		Bars:     items,
	}

	f, _ := os.Create(title + ".png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
