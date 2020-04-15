package visualizer

import (
	"log"
	"os"

	"github.com/BruceMacD/Freelance-Problem-Trends/analyzer"
	"github.com/go-echarts/go-echarts/charts"
)

// DrawTopNOccuranceBarChart draws a bar chart of the top 10 values in a word list by occurances
func DrawTopNOccuranceBarChart(wl analyzer.WordList, title, file string, n int) {
	var nameItems []string
	var valItems []int

	for i := 1; i <= n; i++ {
		nameItems = append(nameItems, wl[i].Value)
		valItems = append(valItems, wl[i].Occurances)
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: title}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).AddYAxis("", valItems)
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})

	f, err := os.Create(file + ".html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}
