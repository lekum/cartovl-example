package main

import (
	"github.com/lekum/cartovl"
)

func main() {
	myMap := cartovl.NewMap(
		&cartovl.MapOptions{
			Container:       "map",
			Style:           "https://basemaps.cartocdn.com/gl/dark-matter-gl-style/style.json",
			Center:          []interface{}{-3.7038, 40.4168},
			Zoom:            11,
			ScrollZoom:      true,
			DragRotate:      false,
			TouchZoomRotate: false,
		},
	)
	cartovl.SetDefaultAuth("aguirao-carto", "Nmn6bkPv_A4LQj6rQFAZeg")
	stationsDs := cartovl.NewDataset("metro_madrid")
	stationsViz := cartovl.NewViz(`color: ramp($line,Prism) width: 10`)
	stationsLayer := cartovl.NewLayer("layer1", stationsDs, stationsViz)
	stationsLayer.AddTo(myMap, "watername_ocean")

	linesDs := cartovl.NewDataset("metro_trace")
	linesViz := cartovl.NewViz(`color: ramp($name,Prism) width: 2`)
	linesLayer := cartovl.NewLayer("layer2", linesDs, linesViz)
	linesLayer.AddTo(myMap, "watername_ocean")
}
