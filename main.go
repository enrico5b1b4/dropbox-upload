package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	accessToken := MustGetEnv("ACCESS_TOKEN")
	filename := "example.json"
	pathOnDropbox := fmt.Sprintf("/test%s/%s", RandStringRunes(5), filename)
	config := dropbox.Config{
		Token:    accessToken,
		LogLevel: dropbox.LogOff,
	}

	// init clients
	dbxFilesClient := files.New(config)
	dbxSharingClient := sharing.New(config)

	// read file content
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	// upload
	commitInfo := files.NewCommitInfo(pathOnDropbox)
	resUL, err := dbxFilesClient.Upload(commitInfo, bufio.NewReader(file))
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Uploaded file: %s", resUL.Name)

	// create shareable link with default settings (public visibility)
	sharedLinkSettings := sharing.NewCreateSharedLinkWithSettingsArg(pathOnDropbox)
	resSharedLink, err := dbxSharingClient.CreateSharedLinkWithSettings(sharedLinkSettings)
	if err != nil {
		log.Fatalln(err)
	}

	sharedFileMetadata := resSharedLink.(*sharing.FileLinkMetadata)
	log.Printf("Created shareable link for file: %s (%s)", sharedFileMetadata.Name, sharedFileMetadata.Url)

	// download
	downloadArg := files.NewDownloadArg(pathOnDropbox)
	resDL, contentDL, err := dbxFilesClient.Download(downloadArg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Download file: %s", resDL.Name)

	// output downloaded file content
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(contentDL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(buf.String())
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func MustGetEnv(key string) string {
	envVar := os.Getenv(key)
	if envVar == "" {
		log.Fatalf("%s env var must not be empty", key)
	}
	return envVar
}
