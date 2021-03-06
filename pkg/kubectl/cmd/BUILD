package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
    "go_test",
    "cgo_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "annotate.go",
        "apiversions.go",
        "apply.go",
        "attach.go",
        "autoscale.go",
        "clusterinfo.go",
        "clusterinfo_dump.go",
        "cmd.go",
        "completion.go",
        "convert.go",
        "cp.go",
        "create.go",
        "create_configmap.go",
        "create_deployment.go",
        "create_namespace.go",
        "create_quota.go",
        "create_secret.go",
        "create_service.go",
        "create_serviceaccount.go",
        "delete.go",
        "describe.go",
        "drain.go",
        "edit.go",
        "exec.go",
        "explain.go",
        "expose.go",
        "get.go",
        "help.go",
        "label.go",
        "logs.go",
        "options.go",
        "patch.go",
        "portforward.go",
        "proxy.go",
        "replace.go",
        "rollingupdate.go",
        "run.go",
        "scale.go",
        "stop.go",
        "taint.go",
        "top.go",
        "top_node.go",
        "top_pod.go",
        "version.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/api/annotations:go_default_library",
        "//pkg/api/errors:go_default_library",
        "//pkg/api/meta:go_default_library",
        "//pkg/api/unversioned:go_default_library",
        "//pkg/api/v1:go_default_library",
        "//pkg/api/validation:go_default_library",
        "//pkg/apimachinery/registered:go_default_library",
        "//pkg/apis/batch/v1:go_default_library",
        "//pkg/apis/extensions/v1beta1:go_default_library",
        "//pkg/apis/policy:go_default_library",
        "//pkg/client/clientset_generated/internalclientset:go_default_library",
        "//pkg/client/clientset_generated/internalclientset/typed/core/internalversion:go_default_library",
        "//pkg/client/restclient:go_default_library",
        "//pkg/client/unversioned:go_default_library",
        "//pkg/client/unversioned/clientcmd:go_default_library",
        "//pkg/client/unversioned/portforward:go_default_library",
        "//pkg/client/unversioned/remotecommand:go_default_library",
        "//pkg/fields:go_default_library",
        "//pkg/kubectl:go_default_library",
        "//pkg/kubectl/cmd/config:go_default_library",
        "//pkg/kubectl/cmd/rollout:go_default_library",
        "//pkg/kubectl/cmd/set:go_default_library",
        "//pkg/kubectl/cmd/templates:go_default_library",
        "//pkg/kubectl/cmd/util:go_default_library",
        "//pkg/kubectl/cmd/util/editor:go_default_library",
        "//pkg/kubectl/metricsutil:go_default_library",
        "//pkg/kubectl/resource:go_default_library",
        "//pkg/kubelet/server/remotecommand:go_default_library",
        "//pkg/kubelet/types:go_default_library",
        "//pkg/labels:go_default_library",
        "//pkg/runtime:go_default_library",
        "//pkg/util/crlf:go_default_library",
        "//pkg/util/errors:go_default_library",
        "//pkg/util/exec:go_default_library",
        "//pkg/util/flag:go_default_library",
        "//pkg/util/interrupt:go_default_library",
        "//pkg/util/intstr:go_default_library",
        "//pkg/util/sets:go_default_library",
        "//pkg/util/strategicpatch:go_default_library",
        "//pkg/util/term:go_default_library",
        "//pkg/util/validation:go_default_library",
        "//pkg/util/validation/field:go_default_library",
        "//pkg/util/wait:go_default_library",
        "//pkg/util/yaml:go_default_library",
        "//pkg/watch:go_default_library",
        "//vendor:github.com/daviddengcn/go-colortext",
        "//vendor:github.com/docker/distribution/reference",
        "//vendor:github.com/docker/docker/pkg/term",
        "//vendor:github.com/evanphx/json-patch",
        "//vendor:github.com/golang/glog",
        "//vendor:github.com/jonboulle/clockwork",
        "//vendor:github.com/renstrom/dedent",
        "//vendor:github.com/spf13/cobra",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "annotate_test.go",
        "apply_test.go",
        "attach_test.go",
        "clusterinfo_dump_test.go",
        "cmd_test.go",
        "cp_test.go",
        "create_configmap_test.go",
        "create_deployment_test.go",
        "create_namespace_test.go",
        "create_quota_test.go",
        "create_secret_test.go",
        "create_service_test.go",
        "create_serviceaccount_test.go",
        "create_test.go",
        "delete_test.go",
        "describe_test.go",
        "drain_test.go",
        "exec_test.go",
        "expose_test.go",
        "get_test.go",
        "label_test.go",
        "logs_test.go",
        "patch_test.go",
        "portforward_test.go",
        "replace_test.go",
        "rollingupdate_test.go",
        "run_test.go",
        "taint_test.go",
        "top_node_test.go",
        "top_pod_test.go",
        "top_test.go",
    ],
    data = [
        "//examples:config",
        "//test/fixtures",
    ],
    library = "go_default_library",
    tags = [
        "automanaged",
        "skip",
    ],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/api/annotations:go_default_library",
        "//pkg/api/errors:go_default_library",
        "//pkg/api/meta:go_default_library",
        "//pkg/api/resource:go_default_library",
        "//pkg/api/testapi:go_default_library",
        "//pkg/api/testing:go_default_library",
        "//pkg/api/unversioned:go_default_library",
        "//pkg/api/v1:go_default_library",
        "//pkg/apimachinery/registered:go_default_library",
        "//pkg/apis/batch:go_default_library",
        "//pkg/apis/extensions:go_default_library",
        "//pkg/apis/policy:go_default_library",
        "//pkg/client/restclient:go_default_library",
        "//pkg/client/restclient/fake:go_default_library",
        "//pkg/client/typed/dynamic:go_default_library",
        "//pkg/conversion:go_default_library",
        "//pkg/kubectl:go_default_library",
        "//pkg/kubectl/cmd/testing:go_default_library",
        "//pkg/kubectl/cmd/util:go_default_library",
        "//pkg/kubectl/resource:go_default_library",
        "//pkg/runtime:go_default_library",
        "//pkg/runtime/serializer/json:go_default_library",
        "//pkg/runtime/serializer/streaming:go_default_library",
        "//pkg/types:go_default_library",
        "//pkg/util/intstr:go_default_library",
        "//pkg/util/strings:go_default_library",
        "//pkg/util/term:go_default_library",
        "//pkg/util/wait:go_default_library",
        "//pkg/watch:go_default_library",
        "//pkg/watch/versioned:go_default_library",
        "//vendor:github.com/spf13/cobra",
        "//vendor:github.com/stretchr/testify/assert",
        "//vendor:k8s.io/heapster/metrics/apis/metrics/v1alpha1",
    ],
)
