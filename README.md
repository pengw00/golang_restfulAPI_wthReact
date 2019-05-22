# golang_restfulAPI_wthReact

#go version go1.12.5 darwin/amd64
#Setup Go compile in VSCode

#deploy in heroku: 
home: https://blooming-brushlands-29652.herokuapp.com
endpoint1: /articles
endpoint2: /article/1
endpoint3: /article/{id} DELETE 
endpoint4: /article/{id} POST


#go project folder must be $GOPATH sub directory

use the following command before deploy
1. go get -u github.com/tools/godeps
2. godep save 

setup the port to env.port with priority!
