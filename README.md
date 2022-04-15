# pi-wasm-go
go wasm support




## Build

```shell
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" assets 
```

## Run

```shell
go run cmd/main.go -dir assets/
```
Request:
```shell
curl http://127.0.0.1:9927
```
Out put:
```shell
<html>
<head>
    <meta charset="utf-8"/>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });
    </script>
</head>
<body>
<h1>WebAssembly Demo</h1>
</body>
</html>%  
```