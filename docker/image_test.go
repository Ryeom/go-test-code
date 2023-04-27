package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestImage(t *testing.T) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	// Ubuntu 이미지 가져오기
	imageName := "ubuntu"
	imageTag := "latest"
	pullOptions := types.ImagePullOptions{}
	out, err := cli.ImagePull(ctx, imageName+":"+imageTag, pullOptions)
	fmt.Println(cli.ClientVersion())
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	// Ubuntu 컨테이너 생성 및 시작
	containerConfig := &container.Config{
		Image: imageName + ":" + imageTag,
	}
	resp, err := cli.ContainerCreate(ctx, containerConfig, nil, nil, nil, "")
	if err != nil {
		log.Fatal(err)
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatal(err)
	}

	// 파일 생성
	containerPath := "/root/test.txt"
	fileContent := []byte("Hello, World!\n")
	if err := ioutil.WriteFile("test.txt", fileContent, 0644); err != nil {
		log.Fatal(err)
	}

	// 파일을 컨테이너로 복사
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	copyToContainerOptions := types.CopyToContainerOptions{
		AllowOverwriteDirWithFile: false,
	}
	if err := cli.CopyToContainer(ctx, resp.ID, containerPath, file, copyToContainerOptions); err != nil {
		log.Fatal(err)
	}

	fmt.Println("File created and copied to the container successfully")
}
