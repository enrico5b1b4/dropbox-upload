# Dropbox Upload

A very simple example of using the Dropbox API to upload, create a shareable link and download a file.  

## Getting Started

### Prerequisites
You will need an Access Token for interacting with the Dropbox API.  
Navigate to https://www.dropbox.com/developers/apps to create one.

### Install and run
Build the binary
```
go build
```
and run it providing the previously generated access token
```
ACCESS_TOKEN=<ACCESS_TOKEN> ./dropbox-upload
```
