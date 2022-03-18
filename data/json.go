// SPDX-License-Identifier: MIT

package data

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/caixw/go-http-routers-testing/apis"
	"github.com/caixw/go-http-routers-testing/routers"
)

const envFilenameJSON = "env.json"

// JSON 输出 JSON 数据
func JSON(dir string, l *log.Logger) error {
	dir = filepath.Join(dir, dataDirName)

	rs := routers.Routers()
	as := apis.APIs()

	env := getEnv()
	env.APIFiles = make([]string, 0, len(as))
	for _, a := range as {
		l.Printf("开始测试 %s \n", a.Name)

		apiFile := &api{
			ID:      a.ID,
			Name:    a.Name,
			Desc:    a.Desc,
			Count:   len(a.APIS),
			Routers: make([]*router, 0, len(rs)),
		}

		for _, r := range rs {
			l.Print("    ", r.Name, "...")

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
		}

		filename := a.ID + ".json"
		env.APIFiles = append(env.APIFiles, filename)
		if err := writeJSON(filepath.Join(dir, filename), apiFile); err != nil {
			return err
		}

		l.Printf("完成 %s 测试\n\n", a.Name)
	}

	return writeJSON(filepath.Join(dir, envFilenameJSON), env)
}

func writeJSON(path string, obj any) error {
	bs, err := json.MarshalIndent(obj, "", "\t")
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
