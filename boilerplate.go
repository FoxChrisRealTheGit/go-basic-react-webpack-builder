package main

import (
	"io"
	"os"
)

func main() {
	//do cool stuff here
	makeAllTheFolders("reduxtester", false)
}

//create file structure
func makeAllTheFolders(fileName string, redux bool) {
	os.MkdirAll("../"+fileName, os.ModePerm)
	// os.MkdirAll("../" + fileName + "/public", os.ModePerm)
	os.MkdirAll("../"+fileName+"/src", os.ModePerm)
	os.MkdirAll("../"+fileName+"/src/components", os.ModePerm)
	os.MkdirAll("../"+fileName+"/src/components/header", os.ModePerm)
	os.MkdirAll("../"+fileName+"/src/css", os.ModePerm)
	if redux{
		os.MkdirAll("../"+fileName+"/src/ducks", os.ModePerm)
	}
	makeAllTheFiles(fileName, redux)
}

//create files
func makeAllTheFiles(fileName string, redux bool) {
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
	var PACKAGEJSON = `{
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
	  }`

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
</html>`

	//index.js
	var INDEXJS = `import React from 'react';
import ReactDOM from 'react-dom';
import App from './components/App.jsx';

ReactDOM.render(
    <App />,
	document.getElementById("root"));`

	//app.js
	var APPJS = `import React from 'react';
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
			/*need for if redux is true and stuff */
			const STORE=`import {createStore, applyMiddleware} from 'redux';
			import promisemiddleware from 'redux-promise-middleware';
			import reducer from './ducks/reducer';
			
			export default createStore(reducer, applyMiddleware(promisemiddleware()))
			`

			const REDUCER=`import axios from 'axios';

			const initialState = {
				something: "sort of works",
			}

			const TEST = 'TEST';

			export function test() {
				    return {
				        type: TEST,
				        payload: "Testing Complete"
				    }
				}

				export default function reducer(state = initialState, action) {
					switch (action.type) {
					case TEST:
					let test = action.payload
					return  Object.assign({}, state, {something: test})
					default:
						return state;
				}
			}
			`


	/*below triggers if redux is true */
	if redux {
		 PACKAGEJSON = `{
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
			  "webpack-cli": "^2.0.13",
			  "redux-promise-middleware":"^5.0.0",
			  "redux":"^3.7.2",
			  "react-redux": "^5.0.7",
			  "axios": "^0.18.0"
			}
		  }`

		 INDEXJS = `import React from 'react';
import ReactDOM from 'react-dom';
import App from './components/App.jsx';
import { Provider } from 'react-redux';
import store from './store';

ReactDOM.render(
	<Provider store={store}>
<App />
</Provider>,
document.getElementById("root"));`

		 APPJS = `import React from 'react';
		 import {connect} from 'react-redux';
		 import {test} from '../ducks/reducer';
	class App extends React.Component{
		render(){
			// this.props.test()
			return(
				<p>{this.props.something}</p>
			)
		}
	}
		function mapStateToProps(state) {
			return state
		}
		export default connect(mapStateToProps, { test })(App);
	`
	}
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
	if redux{
		store, _ := os.Create("../"+ fileName + "/src/store.js")
		//checkError(err)
		defer store.Close()
		io.WriteString(store, STORE)

		reducer, _ := os.Create("../"+ fileName + "/src/ducks/reducer.js")
		//checkError(err)
		defer reducer.Close()
		io.WriteString(reducer, REDUCER)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
