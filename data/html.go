// SPDX-License-Identifier: MIT

package data

import (
	"encoding/json"
	"html/template"
	"io/fs"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"time"
)

func HTML(docs string, l *log.Logger) error {
	e, err := readEnv(docs)
	if err != nil {
		return err
	}

	tpl, err := template.New("index").Funcs(template.FuncMap{
		"kb": func(v uint64) float64 {
			return math.Trunc(float64(v)/1024*1e2) * 1e-2
		},
		"year": func() string {
			return time.Now().Format("2006")
		},
	}).ParseFiles(filepath.Join(docs, "index.tpl"))
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(docs, "index.html"))
	if err != nil {
		return err
	}
	defer f.Close()

	return tpl.ExecuteTemplate(f, "index", e)
}

func readEnv(docs string) (*env, error) {
	dataFS := os.DirFS(path.Join(docs, dataDirName))

	e := &env{}
	if err := readJSON(dataFS, envFilenameJSON, e); err != nil {
		return nil, err
	}

	for _, file := range e.APIFiles {
		a := &api{}
		if err := readJSON(dataFS, file, a); err != nil {
			return nil, err
		}
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
