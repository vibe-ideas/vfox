/*
 *    Copyright 2025 Han Li and contributors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package commands

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/version-fox/vfox/internal"
	"github.com/version-fox/vfox/internal/base"
)

var Path = &cli.Command{
	Name:      "path",
	Usage:     "Show the absolute path of an installed SDK",
	UsageText: "vfox path <sdk>@<version>",
	Action:    pathCmd,
	Category:  CategorySDK,
}

func pathCmd(ctx *cli.Context) error {
	args := ctx.Args()
	if args.First() == "" {
		return cli.Exit("sdk name and version are required (e.g. nodejs@24)", 1)
	}

	sdkArg := args.First()
	argArr := strings.Split(sdkArg, "@")
	if len(argArr) != 2 {
		return cli.Exit("invalid format, expected <sdk>@<version> (e.g. nodejs@24)", 1)
	}

	name := strings.ToLower(argArr[0])
	version := base.Version(argArr[1])

	manager := internal.NewSdkManager()
	defer manager.Close()

	source, err := manager.LookupSdk(name)
	if err != nil {
		fmt.Println("notfound")
		return nil
	}

	if source.CheckExists(version) {
		fmt.Println(source.VersionPath(version))
	} else {
		fmt.Println("notfound")
	}

	return nil
}