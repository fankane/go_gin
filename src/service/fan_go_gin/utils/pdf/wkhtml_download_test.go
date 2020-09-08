package pdf

import (
	"fmt"
	"testing"
)

func TestWkhtmltoPDF(t *testing.T) {
	err := WkhtmltoPDF("https://www.puroland.jp/qrticket/e/?p=3000098763000038003000000000930000000000000009011100033062085000090640002021900000000000020203021081", "./klook.pdf")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("success")
}
