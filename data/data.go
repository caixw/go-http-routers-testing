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

	for cindex, c := range apis.APIS {
		cdata := make([]*item, 0, len(routers.Routers))

		fmt.Fprintf(log, "开始测试 %v \n", c.Name)

		for rindex, r := range routers.Routers {
			fmt.Fprint(log, "    ", r.Name, "......")

			filename := strconv.Itoa(cindex) + "-" + strconv.Itoa(rindex) + ".json"
			path := filepath.Join(dir, filename)
			data := single(c, r)
			data.HitFile = filename
			if err := writeJSON(path, data.Hits); err != nil {
				return err
			}

			data.Hits = nil
			cdata = append(cdata, data)

			fmt.Fprintln(log, "[OK]")
		}

		path := filepath.Join(dir, strconv.Itoa(cindex)+".json")
		if err := writeJSON(path, cdata); err != nil {
			return err
		}

		fmt.Fprintf(log, "完成 %v 测试\n\n", c.Name)
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
	var cnt int
	for _, hit := range ret.Hits {
		if hit.OK {
			cnt++
		}
	}
	ret.HitPrecent = cnt / len(ret.Hits) * 100

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