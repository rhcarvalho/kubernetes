//+build ignore

// exec_demo is used to observe the behavior of containers and exec sessions.
// Useful commands to inspect the Docker host:
//
// 1. watch containers created by exec_demo:
//
//	watch -d -txn1 docker ps -f name=exec-demo --format 'table {{.Names}}\t{{.ID}}\t{{.CreatedAt}}\t{{.RunningFor}}'
//
// 2. watch processes of the most recent container created by exec_demo:
//
//	watch -d -tn1 'docker top "$(docker ps -f name=exec-demo -n 1 -q)" -eo pid,cmd'
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	dockertypes "github.com/docker/engine-api/types"
	dockercontainer "github.com/docker/engine-api/types/container"
	"k8s.io/kubernetes/pkg/kubelet/dockertools"
)

var timeout = flag.Duration("timeout", 2*time.Second, "container exec timeout")

func main() {
	rand.Seed(time.Now().UnixNano())

	client := dockertools.CreateDockerClientOrDie("", 0)

	ccc := dockertypes.ContainerCreateConfig{
		Name: fmt.Sprintf("exec-demo-%08x", rand.Uint32()),
		Config: &dockercontainer.Config{
			Image:      "gcr.io/google_containers/busybox:1.24",
			Entrypoint: []string{"tail", "-f", "/dev/null"},
		},
	}
	createResp, err := client.CreateContainer(ccc)
	if err != nil {
		log.Fatalf("Failed to create container %q: %v", ccc.Name, err)
	}
	log.Printf("Created container %q", ccc.Name)
	if len(createResp.Warnings) != 0 {
		log.Printf("Container created with warnings: %v", createResp.Warnings)
	}

	id := createResp.ID

	if err = client.StartContainer(id); err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}
	log.Print("Started container")

	stdin, stdout, stderr := os.Stdin, os.Stdout, os.Stderr

	ec := dockertypes.ExecConfig{
		Cmd:          []string{"tail", "-f", "/dev/null"},
		AttachStdin:  stdin != nil,
		AttachStdout: stdout != nil,
		AttachStderr: stderr != nil,
	}
	execObj, err := client.CreateExec(id, ec)
	if err != nil {
		log.Fatalf("Failed to create exec: %v", err)
	}
	log.Print("Created exec")

	log.Print("Starting exec...")

	esc := dockertypes.ExecStartCheck{}
	streamOpts := dockertools.StreamOptions{
		InputStream:  stdin,
		OutputStream: stdout,
		ErrorStream:  stderr,
	}
	err = client.StartExec(execObj.ID, esc, streamOpts)
	if err != nil {
		log.Fatalf("Failed to start exec: %v", err)
	}
}
