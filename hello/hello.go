package main

import (
	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

func main() {
	main2()
}

func main2() error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()

	// 设置标题
	pdf.SetX(230)
	pdf.SetY(80)
	err := pdf.AddTTFFont("bold", "/Users/wangxu/golangcode/gopdfsample/ttf/HarmonyOS_Sans_SC_Bold.ttf")
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.SetFont("bold", "", 16)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.Cell(nil, "合同智能审核报告")
	if err != nil {
		return errors.WithStack(err)
	}

	// 设置审查结论
	pdf.SetX(80)
	pdf.SetY(142)
	err = pdf.SetFont("bold", "", 13)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.CellWithOption(&gopdf.Rect{
		W: 435,
		H: 90,
	}, "", gopdf.CellOption{Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top})
	if err != nil {
		return errors.WithStack(err)
	}
	pdf.SetX(85)
	pdf.SetY(152)
	err = pdf.Cell(nil, "审查结论")
	if err != nil {
		return errors.WithStack(err)
	}

	pdf.SetX(85)
	pdf.SetY(185)
	err = pdf.AddTTFFont("black", "/Users/wangxu/golangcode/gopdfsample/ttf//SourceHanSerifCN-VF.ttf")
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.SetFont("black", "", 12)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.Cell(nil, "智能审查工作完成，存在 1 条低风险、1 条高风险审查点未处理，存在 2 条低风")
	if err != nil {
		return errors.WithStack(err)
	}
	pdf.SetX(85)
	pdf.SetY(210)
	err = pdf.Cell(nil, "险审查点未通过，存在 1 条低风险审查点待复核，建议处理高风险审查点")
	if err != nil {
		return errors.WithStack(err)
	}

	// 设置合同基本信息
	pdf.SetX(95)
	pdf.SetY(245)
	err = pdf.SetFont("bold", "", 13)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.Cell(nil, "一、 合同基本信息")
	if err != nil {
		return errors.WithStack(err)
	}

	contractInfoFieldNameList := []string{
		"合同名称 ",
		"合同编号 ",
		"资金流向 ",
		"合同总金额（元）",
		"承办部门 ",
		"合同承办人 ",
		"签约对方当事人法定名称 ",
	}
	for i, fieldName := range contractInfoFieldNameList {
		err = contractInfoTableItem(&pdf, 80, float64(i*30+270), fieldName)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// 设置审查统计结果
	pdf.SetX(95)
	pdf.SetY(580)
	err = pdf.SetFont("bold", "", 13)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.Cell(nil, "二、 审查统计结果")
	if err != nil {
		return errors.WithStack(err)
	}

	err = pdf.WritePdf("/Users/wangxu/golangcode/gopdfsample/hello/hello.pdf")
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func contractInfoTableItem(pdf *gopdf.GoPdf, x, y float64, text string) error {
	pdf.SetFillColor(241, 241, 241)
	pdf.RectFromUpperLeftWithStyle(x, y, 155, 30, "FD")

	pdf.SetX(x)
	pdf.SetY(y)
	pdf.SetFillColor(0, 0, 0)
	err := pdf.SetFont("bold", "", 12)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.CellWithOption(&gopdf.Rect{
		W: 155,
		H: 30,
	}, text, gopdf.CellOption{Align: gopdf.Middle | gopdf.Right,
		Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	pdf.SetX(x + 155)
	pdf.SetY(y)
	err = pdf.SetFont("black", "", 12)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.CellWithOption(&gopdf.Rect{
		W: 280,
		H: 30,
	}, " Center", gopdf.CellOption{Align: gopdf.Left | gopdf.Middle,
		Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
