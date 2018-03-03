package main

import "os"

func main() {
	pathToWebRoot := os.Getenv("HOME") + "/Projects/codes4coffee.github.io"
	linkName := os.Args[1]
	fullURL := os.Args[2]
	err := os.Mkdir(pathToWebRoot+"/"+linkName, os.ModePerm)

	if err != nil {
		panic(err)
	}

	f, err := os.Create(pathToWebRoot + "/" + linkName + "/index.html")

	if err != nil {
		panic(err)
	}

	f.WriteString("<meta http-equiv='refresh' content='0; url=https://" + fullURL + "' />")
}
