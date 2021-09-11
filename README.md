# bazel-sample

bazel プロジェクトのサンプル

## how to build

```shell
# Generate BUILD.bazel for each go package
bazel run //:gazelle
bazel build //parentapp //childapp
```
