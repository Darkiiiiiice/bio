BUILD_VERSION   := 1.0.1
BUILD_TIME      := $(shell date "+%F %T")
BUILD_TAG       := $(shell date "+%Y%m%d-%H%M")
GO_VERSION      := $(shell go version)
COMMIT_SHA1     := $(shell git rev-parse HEAD )

BUILD_DIR = build
TARGET_PATH = ${BUILD_DIR}/bin
TARGET_NAME = bio
OS=linux
ARCH=amd64

build: fmt vet
	GOOS=${OS} GOARCH=${ARCH} go build -ldflags  \
    "                                        \
    -X 'main.GoVersion=${GO_VERSION}'        \
    -X 'main.BuildVersion=${BUILD_VERSION}'  \
    -X 'main.BuildTime=${BUILD_TIME}'        \
    -X 'main.CommitID=${COMMIT_SHA1}'        \
    "                                        \
	-o ${TARGET_PATH}/${TARGET_NAME}

fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...