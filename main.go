package main

import (
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/sqweek/dialog"
)

func main() {
	content, clipError := clipboard.ReadAll()
	if clipError != nil {
		dialog.
			Message("Error escaping your clipboard content: %s", clipError.Error()).
			Title("Escaping clipboard content").
			Error()
		return
	}

	clipboard.WriteAll(strings.TrimRight(strings.TrimLeft(strconv.QuoteToASCII(content), "\""), "\""))
}
