PROJECTNAME = splitwise
PROGNAME = splitwise
SRC = pkg/splitwise.go
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64
#GOROOT=${GOROOT}

GO_TEST_COVERAGE_OUTPUT = coverage.out
GO_TEST_COVERAGE_ARGS = -coverprofile=${GO_TEST_COVERAGE_OUTPUT}

default: build

all: clean build test

build: $(SRC)
	go build -o $(PROGNAME) $(SRC)

#run: $(SRC)
#	go run $(PROGNAME) $(SRC)

#build_all:
#	$(foreach GOOS, $(PLATFORMS),\
#	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -v -o $(BINARY)-$(GOOS)-$(GOARCH))))

test: mockgen_files common_test

mockgen_files: test/mock/users.go test/mock/emanager.go

test/mock/users.go: pkg/user_management/users.go
	mockgen -source=$< -destination=$@ -package=mock

test/mock/emanager.go: pkg/expense_manager/emanager.go
	mockgen -source=$< -destination=$@ -package=mock

common_test:
	go test -v ${GO_TEST_ARGS} ${GO_TEST_COVERAGE_ARGS} -timeout 20m ./...
	go tool cover -func=${GO_TEST_COVERAGE_OUTPUT}
ifeq ($(GENERATE_COVERAGE),1)
	go tool cover -html=${GO_TEST_COVERAGE_OUTPUT} -o coverage.html
endif

clean:
	rm -rf $(PROGNAME)
	rm -rf test/mock/users.go test/mock/emanager.go
