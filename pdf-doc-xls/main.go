package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddUTF8Font("Chinese", "", "NotoSansSC-Regular.ttf")
	pdf.SetFont("Chinese", "", 15)
	pdf.SetTextColor(255, 0, 0)
	pdf.AddPage()

        // 将html转为pdf
        // todo: 如果html中含有b或i标签，会报错。其它场景暂未发现问题
	htmlStr := `测试一下中文: , ` +
		`, <u>带下划线的文本</u>, or <u>all at once</u>!<br><br>` +
		`<center>此处文本居中</center>` +
		`<right>文本右侧排列.</right>` +
		`也可以插入链接，比如：` +
		`<a href="https://www.test.com">https://www.test.com</a>, or 或者图片: click on the logo.`
	html := pdf.HTMLBasicNew()
	html.Write(15, htmlStr)
	err := pdf.OutputFileAndClose("write_html.pdf")
	if err != nil {
		panic(err.Error())
	}
}
