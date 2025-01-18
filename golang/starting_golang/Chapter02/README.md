# memo
## go mod
環境：
```
$ go version
go version go1.23.4 darwin/amd64
```
で本のコードが動かなかった：
```
$ go run main.go
"./animals" is relative, but relative import paths are not supported in module mode
```
以下が必要。
```
$ go mod init Chapter02
```
とすると、`go.mod`ファイルができる：
```
module Chapter02

go 1.23.4
```
次に、`main.go`のimport文を書き換える。
```
import (
    "fmt"

    "Chapter02/animals"
)
```
すると、動く。

## 実行方法
### 実行
```
$ go run main.go
# command-line-arguments
./main.go:10:17: undefined: AppName
```
これは失敗する。`go run main.go`は`main.go`しか見ないから。`app.go`も含めるには、
```
$ go run main.go app.go
```
とする。

### ビルド
```
$ go build
$ ./Chapter02
```
`go build`は、カレントディレクトリ全てが対象になる。

### テスト

```
$ go test ./animals
ok      Chapter02/animals       0.362s

$ go test -v ./animals
=== RUN   TestElephantFeed
--- PASS: TestElephantFeed (0.00s)
=== RUN   TestMonkeyFeed
--- PASS: TestMonkeyFeed (0.00s)
=== RUN   TestRabbitFeed
--- PASS: TestRabbitFeed (0.00s)
PASS
ok      Chapter02/animals       0.366s
```
mainパッケージのテスト：
```
$ go test
PASS
ok      Chapter02       0.358s

$ go test -v
=== RUN   TestAppName
--- PASS: TestAppName (0.00s)
PASS
ok      Chapter02       0.363s
```