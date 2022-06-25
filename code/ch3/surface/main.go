package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	output := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%c' height='%c'>", width, height)
	output_file(output)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			if i > int(dy)/2 && j > int(dy)/2 {
				output := fmt.Sprintf("<polygon points='%f,%f %f,%f %f,%f %f,%f' style='stroke: #f00'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
				output_file(output)
			}else{
				output := fmt.Sprintf("<polygon points='%f,%f %f,%f %f,%f %f,%f' style='stroke: blue'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
				output_file(output)
			}
		}
	}
	output_file("</svg>")

	// web section 
	// http.HandleFunc("/", webHost)
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// func webHost(w http.ResponseWriter,r *http.Request){

// }

func corner(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Compute egglot
	//z := eggplot(x, y)
	
	// check
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, fmt.Errorf("invalid value")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}
func eggplot(x, y float64) float64 {
    return (math.Sin(x) + math.Sin(y)) * 0.25
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func create_file(){
	// create folder
	creating_doc,err := os.Create("output.txt")

	if err != nil {
		fmt.Println("done")
	}
	creating_doc.Close() 
}

func output_file(output string){

	writing_doc,err1 := os.OpenFile("output.txt", os.O_WRONLY|os.O_APPEND, 0666 )

	if err1 !=nil {
		create_file()
	} 

	k, err := writing_doc.WriteString(output) // k yazılan dosyanın boyutu
	errCheck(err)
	fmt.Println(k) 
	writing_doc.Close()
}

func errCheck(err error){
	if err != nil{
		panic(err)
	}
}