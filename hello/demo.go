package main

import (
	"fmt"
	"github.com/signintech/gopdf"
	"github.com/signintech/gopdf/fontmaker/core"
	"log"
)

func main1() {
	//file, _ := exec.LookPath(os.Args[0])
	//path, _ := filepath.Abs(file)
	//index := strings.LastIndex(path, string(os.PathSeparator))

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4
	pdf.AddPage()
	// err := pdf.AddTTFFont("wts", "/Users/wangxu/codefile/go/gopdfsample/hello/wts11.ttf")
	// err := pdf.AddTTFFont("wts", "/Users/wangxu/Downloads/HarmonyOS Sans/HarmonyOS_Sans/HarmonyOS_Sans_Black.ttf")
	// err := pdf.AddTTFFont("wts", "/Users/wangxu/Downloads/HarmonyOS Sans/HarmonyOS_Sans_Condensed_Italic/HarmonyOS_Sans_Condensed_Black_Italic.ttf")
	// err := pdf.AddTTFFont("wts", "/Users/wangxu/Downloads/HarmonyOS Sans/HarmonyOS_Sans_Naskh_Arabic/HarmonyOS_Sans_Naskh_Arabic_Black.ttf")

	fontSize := 16
	var parser core.TTFParser
	err := parser.Parse("/Users/wangxu/Downloads/SourceHanSerif-VF.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	//Measure Height
	//get  CapHeight (https://en.wikipedia.org/wiki/Cap_height)
	cap := float64(float64(parser.CapHeight()) * 1000.00 / float64(parser.UnitsPerEm()))
	//convert
	realHeight := cap*(float64(fontSize)/1000.0) + 5
	fmt.Printf("realHeight = %f", realHeight)

	pdf.Br(realHeight)
	_ = pdf.AddTTFFont("black", "/Users/wangxu/Downloads/SourceHanSerif-VF.ttf")
	_ = pdf.SetFont("black", "", fontSize)
	pdf.Cell(nil, "                    222   bbb  ccc                       智能审查合同报告")

	pdf.Br(realHeight)
	_ = pdf.AddTTFFont("black2", "/Users/wangxu/Downloads/SourceHanSerif-VF.ttf")
	_ = pdf.SetFont("black2", "B", 10)
	pdf.Cell(nil, "倒计时了房间爱丽丝的房间阿里斯顿")

	//pdf.Br(realHeight)
	//_ = pdf.AddTTFFont("bold", "/Users/wangxu/Downloads/HarmonyOS Sans/HarmonyOS_Sans_TC/HarmonyOS_Sans_TC_Black.ttf")
	//_ = pdf.SetFont("bold", "", fontSize)
	//pdf.Cell(nil, "    智能审查合同报告")
	//
	//pdf.Br(realHeight)
	//_ = pdf.AddTTFFont("light", "/Users/wangxu/Downloads/HarmonyOS Sans/HarmonyOS_Sans_Condensed/HarmonyOS_Sans_Condensed_Light.ttf")
	//_ = pdf.SetFont("light", "", fontSize)
	//pdf.Cell(nil, "    智能审查合同报告")
	//
	//pdf.Br(realHeight)
	//_ = pdf.AddTTFFont("medium", "/Users/wangxu/Downloads/HarmonyOS Sans/HarmonyOS_Sans_Condensed/HarmonyOS_Sans_Condensed_Medium.ttf")
	//_ = pdf.SetFont("medium", "", fontSize)
	//pdf.Cell(nil, "    智能审查合同报告")
	//
	//pdf.Br(realHeight)
	//_ = pdf.AddTTFFont("regular", "/Users/wangxu/Downloads/HarmonyOS Sans/HarmonyOS_Sans_Condensed/HarmonyOS_Sans_Condensed_Regular.ttf")
	//_ = pdf.SetFont("regular", "", fontSize)
	//pdf.Cell(nil, "     智能审查合同报告")
	//
	//pdf.Br(realHeight)
	//_ = pdf.AddTTFFont("thin", "/Users/wangxu/Downloads/HarmonyOS Sans/HarmonyOS_Sans_Condensed/HarmonyOS_Sans_Condensed_Thin.ttf")
	//_ = pdf.SetFont("thin", "", fontSize)
	//pdf.Cell(nil, "    智能审查合同报告")

	pdf.WritePdf("/Users/wangxu/golangcode/gopdfsample/hello/hello.pdf")
}
