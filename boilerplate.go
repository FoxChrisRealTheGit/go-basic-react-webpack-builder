package main

import (
	"io"
	"os"
)

func main(){
	//do cool stuff here
	makeAllTheFolders("tester");
}


//create file structure
func makeAllTheFolders(fileName string){
	os.MkdirAll("../" + fileName, os.ModePerm)
	os.MkdirAll("../" + fileName + "/public", os.ModePerm)
	os.MkdirAll("../" + fileName + "/src", os.ModePerm)
	os.MkdirAll("../" + fileName + "/src/components", os.ModePerm)
	os.MkdirAll("../" + fileName + "/src/components/header", os.ModePerm)
	os.MkdirAll("../" + fileName + "/src/css", os.ModePerm)

	makeAllTheFiles(fileName)
}


//create files
func makeAllTheFiles(fileName string){
	//declare items to use for creation below
	//gitignore
	const GITIGNORE=`
	.env
	.node_modules
	
	`

	//readme
	const README = `
	
	
	
	`

	//package.json
	const PACKAGEJSON = `
	
	
	
	`

	//webpack config
	const WEBPACKCONFIG=`
	
	
	
	`

	//index.html
	const INDEXHTML=`
	
	
	
	`

	//index.js
	const INDEXJS =`
	
	
	
	`

	//app.js
	const APPJS =`
	
	
	`

	//headercomponent
	const HEADERCOMP=`
	
	
	`

	/* below are the os create/write opperations for file creation */

	//make the .gitignore
	gitIgnore, _ := os.Create("../" + fileName + "/.gitignore.txt")
	//checkError(err)
	defer gitIgnore.Close()
	io.WriteString(gitIgnore, GITIGNORE)
	//checkError(err)
	
	//make the readme
	readme, _ := os.Create("../" + fileName + "/README.md")
	//checkError(err)
	defer readme.Close()
	io.WriteString(readme, README)
	//checkError(err)


	//make the package.json
	packagejson, _ := os.Create("../" + fileName + "/package.json")
	//checkError(err)
	defer packagejson.Close()
	io.WriteString(packagejson, PACKAGEJSON)
	//checkError(err)


	//make the webpackconfig
	webpackconfig, _ := os.Create("../" + fileName + "/webpack.config.js")
	//checkError(err)
	defer webpackconfig.Close()
	io.WriteString(webpackconfig, WEBPACKCONFIG)
	//checkError(err)

	//make the index.html file
	indexhtml, _ := os.Create("../" + fileName + "/public/index.html")
	//checkError(err)
	defer indexhtml.Close()
	io.WriteString(indexhtml, INDEXHTML)
	//checkError(err)

	//mke the index.js file
	indexjs, _ := os.Create("../" + fileName + "/src/index.js")
	//checkError(err)
	defer indexjs.Close()
	io.WriteString(indexjs, INDEXJS)
	//checkError(err)

	//make the app file
	appjs, _ := os.Create("../" + fileName + "/src/app.js")
	//checkError(err)
	defer appjs.Close()
	io.WriteString(appjs, APPJS)
	//checkError(err)

	//make the headercomponent file
	headercomp, _ := os.Create("../" + fileName + "/components/header/header.js")
	//checkError(err)
	defer appjs.Close()
	io.WriteString(headercomp, HEADERCOMP)
	//checkError(err)
}

func checkError(err error){
	if err != nil{
		panic(err)
	}
}


