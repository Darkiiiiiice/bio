# bio

Get one sentence from [hitokoto](https://hitokoto.cn) and update github bio

## usage

* Build

``` bash
$ make OS=linux ARCH=amd64 build
go fmt ./...
go vet ./...
GOOS=linux GOARCH=amd64 go build -ldflags  \
    "                                        \
    -X 'main.GoVersion=go version go1.20.2 darwin/arm64'        \
    -X 'main.BuildVersion=1.0.0'  \
    -X 'main.BuildTime=2023-04-11 09:43:47'        \
    -X 'main.CommitID=ec0869361628d358e59eb9bad846b0ca68dae9a1'        \
    "                                        \
 -o build/bin/bio
```

* Run

``` bash
$ GITHUB_TOKEN=your_token ./build/bin/bio
BuildInfo:
-- BuildVersion: 1.0.0
-- BuildTime:    2023-04-11 09:45:55
-- BuildWith:    go version go1.20.2 darwin/arm64
-- CommitID:     ec0869361628d358e59eb9bad846b0ca68dae9a1

================================
Hitokoto:
-- ID: 4281
-- Hitokoto: 上了陆地的鱼就不能再叫做鱼了
-- Type: d
-- From: 三体:黑暗森林
-- FromWho:
-- Creator: smallxu
-- CreatorUID: 2639
-- Reviewer: 0
-- UUID: 44220b75-4e12-4e7f-9cbb-2f31f7fbcdcd
-- CommitFrom: web
-- Length: 14
-- CreatedAt: 2019-02-11T19:22:11+08:00
```
