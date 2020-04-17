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
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: title},
		charts.ToolboxOpts{Show: true},
		charts.InitOpts{Width: "1400px", Height: "600px"},
	)
	bar.AddXAxis(nameItems).AddYAxis("", valItems)
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})

	f, err := os.Create(file + ".html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}

// DrawWordCloud outputs a pie chart of the data
func DrawWordCloud(msi map[string]int, title, file string) {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(charts.TitleOpts{Title: title})
	wc.Add(title, convertToInterfaceMap(msi), charts.WordCloudOpts{SizeRange: []float32{14, 80}})

	f, err := os.Create(file + ".html")
	if err != nil {
		log.Println(err)
	}
	wc.Render(f)
}

func convertToInterfaceMap(msi map[string]int) (kv map[string]interface{}) {
	kv = make(map[string]interface{})
	for k, v := range msi {
		kv[k] = v
	}
	return kv
}
