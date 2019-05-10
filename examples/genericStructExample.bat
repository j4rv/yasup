:: Example that shows how to generate YASUP code for the color.Color struct.
go install "github.com/j4rv/yasup/yasupGen"

cd ..

yasupGen.exe -type color.Color -typeCased "Color" -typePackage image/color --test ^
-singleVal "color.RGBA{255,0,0,255}" ^
-multipleVals "color.RGBA{255,0,0,255}, color.RGBA{1,0,0,255}, color.RGBA{2,0,0,255}, color.RGBA{3,0,0,255}, color.RGBA{4,0,0,255}, color.RGBA{5,0,0,255}, color.RGBA{6,0,0,255}, color.RGBA{7,0,0,255}, color.RGBA{8,0,0,255}" 

gofmt -s -w .
go test .