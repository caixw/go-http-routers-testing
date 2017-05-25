// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package data

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"fmt"

	"github.com/caixw/go-http-routes-testing/apis"
	"github.com/caixw/go-http-routes-testing/routers"
	"github.com/issue9/utils"
)

// JSON 输出 JSON 数据，数据按路由进行分组。
func JSON(dir string, log io.Writer) error {
	if !utils.FileExists(dir) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return err
		}
	}

	writeJSON(filepath.Join(dir, "env.json"), getEnv())

	for _, r := range routers.Routers {
		fmt.Fprintf(log, "开始测试 %v \n", r.Name)

		// 以路由名创建目录
		routerDir := filepath.Join(dir, r.Name)
		if !utils.FileExists(routerDir) {
			if err := os.Mkdir(routerDir, os.ModePerm); err != nil {
				return err
			}
		}

		for index, c := range apis.APIS {
			fmt.Fprintf(log, "    %v......", c.Name)

			path := filepath.Join(routerDir, strconv.Itoa(index)+".json")
			if err := writeJSON(path, single(c, r)); err != nil {
				return err
			}

			fmt.Fprintln(log, "[OK]")
		}
		fmt.Fprintf(log, "完成 %v 测试\n\n", r.Name)
	} // end for routers.Routers

	return nil
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
	f.Write(bs)
	f.Close()

	return nil
}

func getEnv() *env {
	return &env{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
		CPU:  runtime.NumCPU(),
	}
}

func single(c *apis.Collection, r *routers.Router) *item {
	ret := &item{
		RouterName: r.Name,
		APIName:    c.Name,
	}

	// 加载
	h, size := loadAPIS(c.APIS, r.Load)
	ret.MemBytes = size

	// 命中情况
	ret.Hits = testHit(c.APIS, h)

	// bench
	rslt := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			api := c.APIS[i%len(c.APIS)]

			w := httptest.NewRecorder()
			r := httptest.NewRequest(api.Method, api.Test, nil)
			h.ServeHTTP(w, r)
		}
	})
	ret.Bench = &bench{
		NsPerOp:           rslt.NsPerOp(),
		AllocsPerOp:       rslt.AllocsPerOp(),
		AllocedBytesPerOp: rslt.AllocedBytesPerOp(),
	}

	return ret
}

func testHit(apis []*apis.API, h http.Handler) []*hit {
	w := httptest.NewRecorder()
	hits := make([]*hit, 0, len(apis))

	for _, api := range apis {
		r := httptest.NewRequest(api.Method, api.Test, nil)
		h.ServeHTTP(w, r)

		hits = append(hits, &hit{
			OK:     w.Body.String() == r.URL.Path,
			Path:   api.Test,
			Want:   api.Brace,
			Actual: w.Body.String(),
		})

		w.Body.Reset()
	}

	return hits
}

// 加载一组 API 到指定的路由中，返回该路由的 http.Handler 接口和所消耗的内存
func loadAPIS(apis []*apis.API, load routers.Load) (http.Handler, uint64) {
	stats := &runtime.MemStats{}

	runtime.GC()
	runtime.ReadMemStats(stats)
	before := stats.HeapAlloc

	h := load(apis)

	runtime.GC()
	runtime.ReadMemStats(stats)
	after := stats.HeapAlloc

	return h, after - before
}
