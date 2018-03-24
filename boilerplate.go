package main

import (
	"io"
	"os"
)

func main() {
	//do cool stuff here
	makeAllTheFolders("tester")
}

//create file structure
func makeAllTheFolders(fileName string) {
	os.MkdirAll("../"+fileName, os.ModePerm)
	// os.MkdirAll("../" + fileName + "/public", os.ModePerm)
	os.MkdirAll("../"+fileName+"/src", os.ModePerm)
	os.MkdirAll("../"+fileName+"/src/components", os.ModePerm)
	os.MkdirAll("../"+fileName+"/src/components/header", os.ModePerm)
	os.MkdirAll("../"+fileName+"/src/css", os.ModePerm)

	makeAllTheFiles(fileName)
}

//create files
func makeAllTheFiles(fileName string) {
	//declare items to use for creation below
	//gitignore
	const GITIGNORE = `# dependencies
/node_modules

# testing
/coverage

# production
/build
/dist

# misc
.DS_Store
.env.local
.env.development.local
.env.test.local
.env.production.local
.env

npm-debug.log*
yarn-debug.log*
yarn-error.log*

	`

	//readme
	const README = `
	This was created with a go builder, more information should probably go here or something...
	
	
	`

	//package.json
	const PACKAGEJSON = `{
		"name": "placeholder",
		"version": "0.0.1",
		"description": "this should probably be updated",
		"main": "/src/index.js",
		"scripts": {
		  "test": "echo \"Error: no test specified\" && exit 1",
		  "build": "webpack",
		  "start": "webpack-dev-server --inline --hot"
		},
		"repository": {
		  "type": "git",
		  "url": "https://github.com/FoxChrisRealTheGit/go-basic-react-webpack-builder"
		},
		"author": "Christopher Fox",
		"license": "ISC",
		"dependencies": {
		  "babel-core": "^6.26.0",
		  "babel-loader": "^7.1.4",
		  "babel-preset-react": "^6.24.1",
		  "react": "^16.2.0",
		  "react-dom": "^16.2.0",
		  "webpack": "^4.1.1",
		  "react-stylux": "^0.4.2",
		  "webpack-cli": "^2.0.13"
		}
	  }
	
	
	`

	//webpack config
	const WEBPACKCONFIG = `module.exports = {
		entry: "./src/Index.js",
		output: {
			filename: 'bundle.js'
		},
		module: {
			rules: [
				{
					exclude: /node_modules/,
					test: /\.jsx?$/,
					use: [
						{
							loader: "babel-loader",
							query:
							{
								presets:['react']
							}
						}
					]
				}
			]
		}
	}
	
	
	`

	//index.html
	const INDEXHTML = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Placeholder</title>

</head>

<body>

    <div id="root"></div>
    <script src="./bundle.js"></script>
</body>
</html>
	
	
	`

	//index.js
	const INDEXJS = `import React from 'react';
import ReactDOM from 'react-dom';
import App from './components/App.jsx';

ReactDOM.render(
    <App />,
    document.getElementById("root"));
	
	
	`

	//app.js
	const APPJS = `import React from 'react';
	export default class App extends React.Component{

		render(){
			return(
				<p>hi</p>
			)
		}
	}
	
	`

	//headercomponent
	const HEADERCOMP = `import React from 'react';

		export default function Header (){
			return(
				<p>Hello</p>
			)
		} 
	
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
	indexhtml, _ := os.Create("../" + fileName + "/index.html")
	//checkError(err)
	defer indexhtml.Close()
	io.WriteString(indexhtml, INDEXHTML)
	//checkError(err)

	//mke the index.js file
	indexjs, _ := os.Create("../" + fileName + "/src/Index.js")
	//checkError(err)
	defer indexjs.Close()
	io.WriteString(indexjs, INDEXJS)
	//checkError(err)

	//make the app file
	appjs, _ := os.Create("../" + fileName + "/src/components/App.jsx")
	//checkError(err)
	defer appjs.Close()
	io.WriteString(appjs, APPJS)
	//checkError(err)

	// //make the headercomponent file
	// headercomp, _ := os.Create("../" + fileName + "/components/header/header.jsx")
	// //checkError(err)
	// defer appjs.Close()
	// io.WriteString(headercomp, HEADERCOMP)
	// //checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
