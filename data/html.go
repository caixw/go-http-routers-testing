// SPDX-License-Identifier: MIT

package data

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"time"
)

//go:embed *.html
var html embed.FS

func HTML(docs string, l *log.Logger) error {
	l.Println("加载数据...")
	e, err := readEnv(docs)
	if err != nil {
		return err
	}

	l.Println("加载模板...")
	tpl, err := template.New("index").Funcs(template.FuncMap{
		"kb": func(v uint64) string {
			return fmt.Sprintf("%.2f", float64(v)/1024)
		},
		"year": func() string {
			return time.Now().Format("2006")
		},
	}).ParseFS(html, "index.html")
	if err != nil {
		return err
	}

	l.Println("应用模板...")
	f, err := os.Create(filepath.Join(docs, "index.html"))
	if err != nil {
		return err
	}
	defer f.Close()

	return tpl.ExecuteTemplate(f, "index", e)
}

func readEnv(docs string) (*env, error) {
	docsFS := os.DirFS(docs)

	e := &env{}
	if err := readJSON(docsFS, path.Join(dataDirName, envFilenameJSON), e); err != nil {
		return nil, err
	}

	for _, file := range e.APIFiles {
		a := &api{}
		file = path.Join(dataDirName, file)
		if err := readJSON(docsFS, file, a); err != nil {
			return nil, err
		}
		a.File = file

		min := &testData{
			MemBytes:          math.MaxUint64,
			NsPerOp:           math.MaxInt64,
			AllocsPerOp:       math.MaxInt64,
			AllocedBytesPerOp: math.MaxInt64,
		}
		for _, r := range a.Routers {
			if min.MemBytes > r.MemBytes {
				min.MemBytes = r.MemBytes
			}
			if min.NsPerOp > r.NsPerOp {
				min.NsPerOp = r.NsPerOp
			}
			if min.AllocsPerOp > r.AllocsPerOp {
				min.AllocsPerOp = r.AllocsPerOp
			}
			if min.AllocedBytesPerOp > r.AllocedBytesPerOp {
				min.AllocedBytesPerOp = r.AllocedBytesPerOp
			}
		}
		a.Min = min

		e.APIs = append(e.APIs, a)
	}

	return e, nil
}

func readJSON(fsys fs.FS, name string, v any) error {
	data, err := fs.ReadFile(fsys, name)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
