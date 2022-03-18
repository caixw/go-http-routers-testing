// SPDX-License-Identifier: MIT

package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/caixw/go-http-routers-testing/apis"
	"github.com/caixw/go-http-routers-testing/routers"
)

// JSON 输出 JSON 数据
func JSON(dir string, log io.Writer) error {
	sort.SliceStable(apis.APIS, func(i, j int) bool {
		return apis.APIS[i].Name > apis.APIS[j].Name
	})

	sort.SliceStable(routers.Routers, func(i, j int) bool {
		return routers.Routers[i].Name < routers.Routers[j].Name
	})

	env := getEnv()
	env.APIs = make([]string, 0, len(apis.APIS))
	for _, a := range apis.APIS {
		fmt.Fprintf(log, "开始测试 %s \n", a.Name)

		apiFile := &api{
			ID:      a.ID,
			Name:    a.Name,
			Desc:    a.Desc,
			Count:   len(a.APIS),
			Routers: make([]*router, 0, len(routers.Routers)),
		}

		for _, r := range routers.Routers {
			fmt.Fprint(log, "    ", r.Name, "......")

			data := testRouter(a, r)

			if data.Misses > 0 {
				filename := a.ID + "_" + r.ID + ".json"
				data.MissFile = filename
				if err := writeJSON(filepath.Join(dir, filename), data.missesData); err != nil {
					return err
				}
				data.missesData = nil
			}

			apiFile.Routers = append(apiFile.Routers, data)

			fmt.Fprintln(log, "[OK]")
		}

		filename := a.ID + ".json"
		env.APIs = append(env.APIs, filename)
		if err := writeJSON(filepath.Join(dir, filename), apiFile); err != nil {
			return err
		}

		fmt.Fprintf(log, "完成 %s 测试\n\n", a.Name)
	}

	return writeJSON(filepath.Join(dir, "env.json"), env)
}

func writeJSON(path string, obj interface{}) error {
	bs, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bs)
	return err
}
