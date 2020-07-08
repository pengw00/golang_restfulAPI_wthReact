# golang_restfulAPI_wthReact

#go version go1.12.5 darwin/amd64
#Setup Go compile in VSCode

#deploy in heroku: 
home: https://blooming-brushlands-29652.herokuapp.com
endpoint1: /articles
endpoint2: /article/1
endpoint3: /article/{id} DELETE 
endpoint4: /article/{id} POST

//those endpoint above is not autherized!


#go project folder must be $GOPATH sub directory

use the following command before deploy
1. go get -u github.com/tools/godeps
2. godep save 

setup the port to env.port with priority!

website: https://mholt.github.io/json-to-go/ to get json convert to go!

#debug hell: forget the create database when create item in db DB.create(accoount)!!!!!

#  not auth endpoints:api/user/new && api/user/login && auth: api/me/contact

#Update: 
Go Mod && go Sum with sove the dependencies management
