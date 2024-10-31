package apps

import (
	"fmt"

	"github.com/kareka-gb/cphelper/utils"
)

type GetterApp struct{}

func (g *GetterApp) Run(contestLink string) {
	parsedResponse := utils.NewContestParser(contestLink)
	fmt.Println(parsedResponse)
}
