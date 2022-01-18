# smms-tiny
A tiny go library for uploading local file to `sm.ms`,which is a free picture bad

## Install

```go
$ go get github.com/Stitch-Zhang/smms
```

## Usage

To upload local images to `smms` you need to generate a `smms` API TOKEN

[generate token](https://sm.ms/home/apitoken)



```go
package main

import (
	"fmt"
    
	"github.com/Stitch-Zhang/smms"
)

func main() {
	uploadURL, err := smms.UploadImg("image_path", "API TOKEN")
	if err != nil {
		fmt.Printf("upload failed : %s", err.Error())
		return
	}
	fmt.Printf("uploaded successfully \n url: %s", uploadURL)
}
// if this works uploadURL should be the direct link of your local image in sm.ms
```

