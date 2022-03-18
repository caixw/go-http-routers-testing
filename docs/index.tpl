{{- define "index" -}}
<html lang="zh-cnm-Hans">
<head>
    <meta charset="utf-8"/>
    <title>Go 路由功能测试数据</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <meta name="generator" content="https://github.com/caixw/go-http-routers-testing"/>
    <meta name="author" content="https://caixw.io"/>
    <meta name="keywords" content="go,golang,benchmark,testing,route,routes,router,mux"/>
    <meta name="description" content="Google Go 路由的性能测试工具。"/>
    <link rel="icon" href="./favicon.ico"/>
    <link href="./style.css" type="text/css" rel="stylesheet"/>
</head>

<body>
<article @vue:mounted="mounted()">
    <h1>Go 路由功能测试数据</h1>

    <p>数据根据测试环境的不同，会有不同的呈现，仅可作为参考用。</p>

    <section>
        <h3>测试环境</h3>
        <ul>
            <li>系统: {{.OS}}</li>
            <li>架构: {{.Arch}}</li>
            <li>CPU: {{.CPU}} 核</li>
            <li>版本: {{.Go}}</li>
        </ul>
    </section>

    {{- range .APIs}}
    <section id="{{.ID}}">
        <h3>{{.Name}}({{.Count}})</h3>
        <p>{{.Desc}}</p>
        <table>
            <thead>
                <tr>
                    <th>路由</th>
                    <th>ns/op</th>
                    <th>B/op</th>
                    <th>allocs/op</th>
                    <th>占用内存(KB)</th>
                    <th>未命中</th>
                </tr>
            </thead>
            <tbody>
                {{- range .Routers}}
                <tr>
                    <th><a href="{{.URL}}">{{.Name}}</a></th>
                    <td>{{.NsPerOp}}</td>
                    <td>{{.AllocedBytesPerOp}}</td>
                    <td>{{.AllocsPerOp}}</td>
                    <td>{{kb .MemBytes}}</td>
                    <td>{{if gt .Misses 0}}<a href="miss.html?{{.MissFile}}">{{.Misses}}</a>{{else}}无{{end}}</td>
                </tr>
                {{- end}}
            </tbody>
        </table>
    </section>
    {{- end -}}

</article>

<footer>
    <p>&copy; 2017-{{year}} by
        <a href="https://caixw.io">caixw</a> &middot;
        <a href="https://github.com/caixw/go-http-routers-testing">Github</a>
    </p>
</footer>

</body>
</html>
{{- end -}}
