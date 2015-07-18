# Go - ScanDir

## Get It
Use go get to be able to import this code into your project

```bash
go get github.com/drpain/ScanDir
```

## Usage Example
Using it is pretty simple. With  2 required parameters, and 2 optional.

#### Example 1
Minimum required options, path and file extension

```go
scandir.Find("/home/user/", "jpg")
```


#### Example 2
Explicitly look for a string in the filepath with the third option. ie: /home/rudi/```Pictures```/Wallpapers/wallpaper1.jpg
```go
scandir.Find("/home/user/", "jpg", "Pictures")
```


#### Example 3
Let's make sure that we ignore specific files. For example ```thumbs.db``` and stuff like that. But really this would be any string in the path to ignore. Try to not make this too short as this may match things you don't expect.
```go
scandir.Find("/home/user/", "jpg", "Pictures", "thumbs.db")
```


#### Full example
In this example I look for CSS files, in a root path. And I ensure that the path contains ```assets/css/out```. And lastly I ensure to exclude already minified css files ```min.css```
```go
package main

import (
    "fmt"
    "github.com/drpain/ScanDir"
)

func main() {
    var files = scandir.Find("/home/rudi/DockIt/src/public/", "css", "assets/css/out", "min.css")
    for _, file := range files {
        fmt.Println(file)
    }
}
```

## Why the hell would I need it?
#### I wrote this as a helper for the filepath.Walk method:

My use-case may not be the same as most people, but maybe this helps someone out there.

#### My usecase (CSS, JS Asset Generation)
* My website uses dynamic asset generation based on the page you are on.
* It has a list of assets, which is joined into one file.
* Then I get the unique name of the asset by MD5 of the file contents.
* Which ensures that browsers will only cache content which is valid for the page.
* And this means I do not have to do any other cache busting.

Now this has been working great for me, but it's not minified at the moment. It just takes too much processing at time of rendering a page, and no body wants to hang around for 5 seconds just for me to minify CSS and JS.

This way I can have a routine which runs periodically. Scans the directories and minifies the files.

