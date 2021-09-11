module github.com/nrnrk/bazel-sample/childapp

go 1.17

require (
	github.com/nrnrk/bazel-sample/library v0.1.0
	google.golang.org/grpc v1.40.0
	google.golang.org/grpc/examples v0.0.0-20210910232509-03268c8ed29e
)

replace github.com/nrnrk/bazel-sample/library => ../library
