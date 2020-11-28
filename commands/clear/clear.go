// Copyright (c) 2020 rookie-ninja
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package rk_clear

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rookie-ninja/rk/common"
	"github.com/urfave/cli/v2"
	"os"
)

func Clear() *cli.Command {
	command := &cli.Command{
		Name:      "clear",
		Usage:     "clear target folder generated by rk build",
		UsageText: "rk clear",
		Action:    ClearAction,
	}

	return command
}

func ClearAction(ctx *cli.Context) error {
	event := rk_common.GetEvent("clear")

	// 1: removing target folder
	color.Cyan("[Action] clearing...")
	// just try our best to remove it
	if err := os.RemoveAll("target"); err != nil {
		color.Red(err.Error())
		event.AddPair("clear-target", "fail")
		return rk_common.Error(event,
			rk_common.NewScriptError(
				fmt.Sprintf("failed to clear target folder \n[err] %v", err)))
	}

	event.AddPair("clear-target", "success")
	rk_common.Success()

	return nil
}
