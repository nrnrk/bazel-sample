module github.com/nrnrk/bazel-sample/childapp

go 1.17

require github.com/nrnrk/bazel-sample/library v0.1.0

replace github.com/nrnrk/bazel-sample/library => ../library
