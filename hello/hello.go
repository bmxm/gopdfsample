package main

import (
	"bytes"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
	chart "github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	err := main2()
	if err != nil {
		log.Fatal(err)
	}
}

func main2() error {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()

	// 设置标题
	pdf.SetX(230)
	pdf.SetY(80)
	err := pdf.AddTTFFont("bold", "/Users/wangxu/codefile/go/gopdfsample/ttf/HarmonyOS_Sans_SC_Bold.ttf")
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
	err = pdf.AddTTFFont("black", "/Users/wangxu/codefile/go/gopdfsample/ttf/SourceHanSerifCN-VF.ttf")
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
		err = contractInfoTableItem(pdf, 80, float64(i*30+270), fieldName)
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

	// 审查结果统计图
	err = drawImage(pdf, 100, 600)
	if err != nil {
		return errors.WithStack(err)
	}

	// 风险点审核情况
	pdf.AddPage()
	err = drawReviewStatisticTable(pdf, 85, 80)
	if err != nil {
		return errors.WithStack(err)
	}

	// 部门审核情况
	err = drawDeptStatisticTable(pdf, 85, 80+120)
	if err != nil {
		return errors.WithStack(err)
	}

	err = pdf.WritePdf("/Users/wangxu/codefile/go/gopdfsample/hello/hello.pdf")
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

type GroupStatistic struct {
	GroupName       string `json:"groupName"`
	GroupID         int64  `json:"groupId"`
	PassedCount     int64  `json:"passedCount"`
	RejectCount     int64  `json:"rejectCount"`
	ReviewCount     int64  `json:"reviewCount"`
	SuggestionCount int64  `json:"suggestionCount"`
}

type TotalStatistic struct {
	NotHandleCount int64 `json:"notHandleCount"`
	PassedCount    int64 `json:"passedCount"`
	RejectCount    int64 `json:"rejectCount"`
	ReviewCount    int64 `json:"reviewCount"`
}

type ReviewResultStatistic struct {
	Total     TotalStatistic   `json:"total"`
	GroupList []GroupStatistic `json:"groupList"`
}

var reviewResultStatistic = ReviewResultStatistic{
	Total: TotalStatistic{},
	GroupList: []GroupStatistic{
		{
			GroupName:       "承办部门",
			GroupID:         0,
			PassedCount:     0,
			RejectCount:     0,
			ReviewCount:     0,
			SuggestionCount: 0,
		},
		{
			GroupName:       "业务部门",
			GroupID:         0,
			PassedCount:     0,
			RejectCount:     0,
			ReviewCount:     0,
			SuggestionCount: 0,
		},
		{
			GroupName:       "财务部门",
			GroupID:         0,
			PassedCount:     0,
			RejectCount:     0,
			ReviewCount:     0,
			SuggestionCount: 0,
		},
		{
			GroupName:       "法规部门",
			GroupID:         0,
			PassedCount:     0,
			RejectCount:     0,
			ReviewCount:     0,
			SuggestionCount: 0,
		},
	},
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

func drawImage(pdf *gopdf.GoPdf, x, y float64) error {
	b, err := getPieImg()
	if err != nil {
		return errors.WithStack(err)
	}
	imageHolder, err := gopdf.ImageHolderByBytes(b.Bytes())
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.ImageByHolder(imageHolder, 132, 615, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	pdf.SetX(160)
	pdf.SetY(625)
	err = pdf.SetFont("black", "", 9)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.Cell(nil, "风险点审核情况")
	if err != nil {
		return errors.WithStack(err)
	}

	// 图例
	drawLegend(pdf, 85, 650, true)

	pdf.SetX(85)
	pdf.SetY(610)
	err = pdf.SetFont("bold", "", 13)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.CellWithOption(&gopdf.Rect{
		W: 217,
		H: 200,
	}, "", gopdf.CellOption{Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top})
	if err != nil {
		return errors.WithStack(err)
	}

	// x y 是第二个方框左上角的坐标
	err = drawBarImage(pdf, 302, 610)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func drawBarImage(pdf *gopdf.GoPdf, x, y float64) error {
	// 框
	pdf.SetX(x)
	pdf.SetY(y)
	err := pdf.SetFont("bold", "", 13)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.CellWithOption(&gopdf.Rect{
		W: 217,
		H: 200,
	}, "", gopdf.CellOption{Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top})
	if err != nil {
		return errors.WithStack(err)
	}

	// 标题
	pdf.SetX(x + 72)
	pdf.SetY(y + 15)
	err = pdf.SetFont("black", "", 9)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.Cell(nil, "部门审核情况")
	if err != nil {
		return errors.WithStack(err)
	}

	// 横线
	pdf.SetLineWidth(1)
	pdf.SetStrokeColor(228, 228, 228)
	pdf.Line(x+30, y+40, x+200, y+40)
	pdf.Line(x+30, y+65, x+200, y+65)
	pdf.Line(x+30, y+90, x+200, y+90)
	pdf.Line(x+30, y+115, x+200, y+115)
	pdf.Line(x+30, y+140, x+200, y+140)

	// 先写死一个用户组个数
	//groupLength := 4
	//perGroupWidth := 200 / groupLength
	//perGroupBar := perGroupWidth / 4
	//
	//pdf.SetFillColor(0, 176, 243)
	//pdf.RectFromUpperLeftWithStyle(x+float64(perGroupBar), y+40, 5, 5, "F")
	//
	//pdf.SetFillColor(0, 0, 0)

	// 图例
	drawLegend(pdf, x+5, y+180, false)

	return nil
}

func getPieImg() (*bytes.Buffer, error) {
	font, err := getFont()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pie := chart.PieChart{
		Font:   font,
		Width:  217,
		Height: 400,
		Values: []chart.Value{
			{Value: 5, Label: "待处理", Style: chart.Style{
				FillColor: drawing.Color{
					R: 0,
					G: 176,
					B: 243,
					A: 255,
				},
			}},
			{Value: 5, Label: "未通过", Style: chart.Style{
				FillColor: drawing.Color{
					R: 255,
					G: 0,
					B: 0,
					A: 255,
				},
			}},
			{Value: 4, Label: "已通过", Style: chart.Style{
				FillColor: drawing.Color{
					R: 60,
					G: 179,
					B: 113,
					A: 255,
				},
			}},
			{Value: 4, Label: "待复核", Style: chart.Style{
				FillColor: drawing.Color{
					R: 255,
					G: 185,
					B: 15,
					A: 255,
				},
			}},
		},
	}

	b := bytes.NewBuffer([]byte{})
	err = pie.Render(chart.PNG, b)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return b, nil
}

// drawLegend 图例, 饼图和直方图一个函数
func drawLegend(pdf *gopdf.GoPdf, x, y float64, isPie bool) error {
	pdf.SetFillColor(0, 176, 243)
	pdf.RectFromUpperLeftWithStyle(x+35, y, 5, 5, "F")
	pdf.SetX(x + 40)
	pdf.SetY(y - 3)
	pdf.SetFillColor(0, 0, 0)
	err := pdf.Cell(nil, "待处理")
	if err != nil {
		return errors.WithStack(err)
	}

	pdf.SetFillColor(60, 179, 113)
	pdf.RectFromUpperLeftWithStyle(x+75, y, 5, 5, "F")
	pdf.SetX(x + 80)
	pdf.SetY(y - 3)
	pdf.SetFillColor(0, 0, 0)
	err = pdf.Cell(nil, "已通过")
	if err != nil {
		return errors.WithStack(err)
	}

	pdf.SetFillColor(255, 0, 0)
	pdf.RectFromUpperLeftWithStyle(x+115, y, 5, 5, "F")
	pdf.SetX(x + 120)
	pdf.SetY(y - 3)
	pdf.SetFillColor(0, 0, 0)
	err = pdf.Cell(nil, "未通过")
	if err != nil {
		return errors.WithStack(err)
	}

	pdf.SetFillColor(255, 185, 15)
	pdf.RectFromUpperLeftWithStyle(x+155, y, 5, 5, "F")
	pdf.SetX(x + 160)
	pdf.SetY(y - 3)
	pdf.SetFillColor(0, 0, 0)
	if isPie {
		err = pdf.Cell(nil, "待复核")
	} else {
		err = pdf.Cell(nil, "补充意见")
	}
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// getFont 加载字体
func getFont() (*truetype.Font, error) {

	//fontFile := "/Users/wangxu/codefile/go/gopdfsample/ttf/HarmonyOS_Sans_SC_Bold.ttf"
	fontFile := "/Users/wangxu/codefile/go/gopdfsample/ttf/SourceHanSerifCN-VF.ttf"
	//fontFile := "/Library/Fonts/AppleMyungjo.ttf"

	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return font, nil
}

// 风险点审核表格
func drawReviewStatisticTable(pdf *gopdf.GoPdf, x, y float64) error {
	pdf.SetX(x)
	pdf.SetY(y)
	err := pdf.SetFont("bold", "", 12)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.Cell(nil, "2.1 风险点审核情况")
	if err != nil {
		return errors.WithStack(err)
	}

	pdf.SetX(x - 5)
	pdf.SetY(y + 30)
	err = pdf.SetFont("black", "", 12)
	if err != nil {
		return errors.WithStack(err)
	}
	pdf.SetLineWidth(0.5)
	tableNameArray := [...]string{"", "全部", "待处理", "已通过", "未通过", "待复核"}
	for i := range tableNameArray {
		err = pdf.CellWithOption(&gopdf.Rect{
			W: 65,
			H: 32,
		}, tableNameArray[i], gopdf.CellOption{Align: gopdf.Middle | gopdf.Center,
			Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
		})
		if err != nil {
			return errors.WithStack(err)
		}
	}
	pdf.SetX(x - 5)
	pdf.SetY(y + 62)
	tableValueArray := [...]string{"智能审查", "0", "0", "0", "0", "0"}
	for i := range tableValueArray {
		err = pdf.CellWithOption(&gopdf.Rect{
			W: 65,
			H: 32,
		}, tableValueArray[i], gopdf.CellOption{Align: gopdf.Middle | gopdf.Center,
			Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
		})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// 部门审核
func drawDeptStatisticTable(pdf *gopdf.GoPdf, x, y float64) error {
	pdf.SetX(x)
	pdf.SetY(y)
	err := pdf.SetFont("bold", "", 12)
	if err != nil {
		return errors.WithStack(err)
	}
	err = pdf.Cell(nil, "2.2 部门审核情况")
	if err != nil {
		return errors.WithStack(err)
	}

	pdf.SetX(x - 5)
	pdf.SetY(y + 30)
	err = pdf.SetFont("black", "", 12)
	if err != nil {
		return errors.WithStack(err)
	}
	pdf.SetLineWidth(0.5)
	tableNameArray := [...]string{"", "已通过", "未通过", "待复核", "补充意见"}
	for i := range tableNameArray {
		err = pdf.CellWithOption(&gopdf.Rect{
			W: 79,
			H: 32,
		}, tableNameArray[i], gopdf.CellOption{Align: gopdf.Middle | gopdf.Center,
			Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
		})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	for i := range reviewResultStatistic.GroupList {
		pdf.SetX(x - 5)
		pdf.SetY(y + 62 + float64(i*32))

		tableValueArray := [...]string{
			reviewResultStatistic.GroupList[i].GroupName,
			strconv.Itoa(int(reviewResultStatistic.GroupList[i].PassedCount)),
			strconv.Itoa(int(reviewResultStatistic.GroupList[i].RejectCount)),
			strconv.Itoa(int(reviewResultStatistic.GroupList[i].ReviewCount)),
			strconv.Itoa(int(reviewResultStatistic.GroupList[i].SuggestionCount)),
		}
		for j := range tableValueArray {
			err = pdf.CellWithOption(&gopdf.Rect{
				W: 79,
				H: 32,
			}, tableValueArray[j], gopdf.CellOption{Align: gopdf.Middle | gopdf.Center,
				Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,
			})
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}
