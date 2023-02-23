package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func GetFileMD5(pathName string) string {
	f, err := os.Open(pathName)
	if err != nil {
		fmt.Println("Open", err)
		return ""
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		fmt.Println("Copy", err)
		return ""
	}
	has := md5hash.Sum(nil)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
func GenerateM3u8(filename, output string) {

	cmdArguments := []string{"-i", filename, "-c:v", "libx264",
		"-c:a", "aac", "-strict", "-2", "-f", "hls", "-hls_time", "10", "-hls_list_size", "2", output,
	}

	cmd := exec.Command("ffmpeg", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("command output: %q", out.String())
}
