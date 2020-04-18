package visualizer

import (
	"fmt"
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

// DrawHeatMap visualizes correlations between some data against a key
func DrawHeatMap(hmd [][3]interface{}, title, file string) {
	// TODO: use hmd
	fmt.Printf("not using hmd: %v", hmd)
	// TODO: remove
	var hours = [...]string{
		"12a", "1a", "2a", "3a", "4a", "5a", "6a", "7a", "8a", "9a", "10a", "11a",
	}
	var days = [...]string{
		"Saturday", "Friday", "Thursday", "Wednesday", "Tuesday", "Monday", "Sunday"}
	testData := [][3]int{{0, 0, 5}, {0, 1, 1}, {0, 2, 0}, {0, 3, 0}, {0, 4, 0}, {0, 5, 0}}

	hm := charts.NewHeatMap()
	hm.SetGlobalOptions(charts.TitleOpts{Title: title})
	hm.AddXAxis(hours).AddYAxis("heatmap", testData)
	hm.SetGlobalOptions(
		charts.YAxisOpts{Data: days, Type: "category", SplitArea: charts.SplitAreaOpts{Show: true}},
		charts.XAxisOpts{Type: "category", SplitArea: charts.SplitAreaOpts{Show: true}},
		charts.VisualMapOpts{Calculable: true, Max: 10, Min: 0,
			InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
	)

	f, err := os.Create(file + ".html")
	if err != nil {
		log.Println(err)
	}
	hm.Render(f)
}

func convertToInterfaceMap(msi map[string]int) (kv map[string]interface{}) {
	kv = make(map[string]interface{})
	for k, v := range msi {
		kv[k] = v
	}
	return kv
}
