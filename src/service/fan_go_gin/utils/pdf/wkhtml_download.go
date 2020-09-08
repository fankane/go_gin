package pdf

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func WkhtmltoPDF(reqURL, pdfFile string) error {
	if reqURL == "" || pdfFile == "" {
		return fmt.Errorf("[WkHtmlToPDF] reqURL(%s)或pdfFile(%s)为空.", reqURL, pdfFile)
	}

	cmdTpl := `/usr/local/bin/wkhtmltopdf --orientation Portrait --page-size A4 --encoding utf-8 -R 0 -L 0 -T 0 -B 0 --quiet page %s %s`
	cmdCreate := fmt.Sprintf(cmdTpl, reqURL, pdfFile)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := exec.CommandContext(ctx, "sh", "-c", cmdCreate).Run()
	if err != nil {
		return fmt.Errorf(" err:%s cmd:%s", err, cmdCreate)
	}
	return nil
}
