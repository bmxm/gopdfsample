package main

import (
	"fmt"
	"github.com/signintech/gopdf"
	"log"
	"testing"
)

func TestXY(t *testing.T) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("HDZB_5", "/Users/wangxu/Downloads/SourceHanSerif-VF.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("HDZB_5", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetY(20)
	pdf.SetX(40)
	pdf.Cell(nil, "沃尔夫就")

	pdf.SetY(0)
	pdf.SetX(20)
	pdf.Cell(nil, "1")

	pdf.SetY(60)
	s := ""
	i := 0
	for i*14 <= 595 {
		i++
		s += "1"
	}
	fmt.Println(i)

	pdf.MultiCell(&gopdf.Rect{
		W: 20,
		H: 200,
	}, s+"djflasdjfla;ksdjflkasjd;lfjasldjflajldkjflasjdf地方拉时代峻峰垃圾收代理费就按大沙发")

	pdf.WritePdf("/Users/wangxu/golangcode/gopdfsample/hello/hello.pdf")

}
