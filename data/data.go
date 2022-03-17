// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package data 处理数据
package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"runtime/metrics"
	"sort"
	"strconv"
	"testing"

	"github.com/caixw/go-http-routers-testing/apis"
	"github.com/caixw/go-http-routers-testing/routers"
)

// JSON 输出 JSON 数据，数据按路由进行分组。
func JSON(dir string, log io.Writer) error {
	sort.SliceStable(apis.APIS, func(i, j int) bool {
		return apis.APIS[i].Name > apis.APIS[j].Name
	})

	sort.SliceStable(routers.Routers, func(i, j int) bool {
		return routers.Routers[i].Name < routers.Routers[j].Name
	})

	env := getEnv()
	env.Data = make([]string, 0, len(apis.APIS))
	for apiIndex, a := range apis.APIS {
		fmt.Fprintf(log, "开始测试 %v \n", a.Name)

		apiFile := &api{
			Name:    a.Name,
			Desc:    a.Desc,
			Count:   len(a.APIS),
			Routers: make([]*router, 0, len(routers.Routers)),
		}

		for routerIndex, r := range routers.Routers {
			fmt.Fprint(log, "    ", r.Name, "......")

			data := single(a, r)

			if data.Misses > 0 {
				filename := strconv.Itoa(apiIndex) + "-" + strconv.Itoa(routerIndex) + ".json"
				data.MissFile = filename
				if err := writeJSON(filepath.Join(dir, filename), data.missesData); err != nil {
					return err
				}
				data.missesData = nil
			}

			apiFile.Routers = append(apiFile.Routers, data)

			fmt.Fprintln(log, "[OK]")
		}

		filename := strconv.Itoa(apiIndex) + ".json"
		env.Data = append(env.Data, filename)
		if err := writeJSON(filepath.Join(dir, filename), apiFile); err != nil {
			return err
		}

		fmt.Fprintf(log, "完成 %v 测试\n\n", a.Name)
	} // end for routers.Routers

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

func getEnv() *env {
	return &env{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
		CPU:  runtime.NumCPU(),
		Go:   runtime.Version(),
	}
}

func single(c *apis.Collection, r *routers.Router) *router {
	ret := &router{
		RouterName: r.Name,
		RouterURL:  r.URL,
	}

	// 加载
	s, size := loadAPIS(c.APIS, r.Load)
	ret.MemBytes = size

	// 命中情况
	ret.missesData = testMiss(c.APIS, s)
	ret.Misses = len(ret.missesData)

	// bench
	rslt := testing.Benchmark(func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			api := c.APIS[i%len(c.APIS)]
			s(api)
		}
	})
	ret.NsPerOp = rslt.NsPerOp()
	ret.AllocsPerOp = rslt.AllocsPerOp()
	ret.AllocedBytesPerOp = rslt.AllocedBytesPerOp()

	return ret
}

func testMiss(apis []*apis.API, s routers.ServeFunc) []*miss {
	w := httptest.NewRecorder()
	misses := make([]*miss, 0, len(apis))

	for _, api := range apis {
		r, err := http.NewRequest(api.Method, api.Test, nil)
		if err != nil {
			panic(err)
		}

		body := s(api)
		if body != r.URL.Path {
			misses = append(misses, &miss{
				Method: r.Method,
				Path:   api.Test,
				Want:   api.Brace,
				Body:   body,
			})
		}

		w.Body.Reset()
	}

	return misses
}

// 加载一组 API 到指定的路由中，返回该路由的 http.Handler 接口和所消耗的内存
func loadAPIS(apis []*apis.API, load routers.Load) (routers.ServeFunc, uint64) {
	sample := make([]metrics.Sample, 1)
	sample[0].Name = "/gc/heap/allocs:bytes"

	runtime.GC()
	metrics.Read(sample)
	before := sample[0].Value

	s := load(apis)

	metrics.Read(sample)
	after := sample[0].Value

	return s, after.Uint64() - before.Uint64()
}
