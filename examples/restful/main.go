// RESTful API Controller Example.

package main

import (
	"github.com/headwindfly/clevergo"
	"html/template"
)

var (
	html = `<html>
	<head></head>
	<body>
		<h3>RESTful API Controller Example.</h3>

		<h4>Requests</h4>
		<ul>
			<li><a target="_blank" href="javascript:verifyByPOST('POST');">POST</a></li>
			<li><a target="_blank" href="javascript:verifyByPOST('DELETE');">DELETE</a></li>
			<li><a target="_blank" href="javascript:verifyByPOST('PUT');">PUT</a></li>
			<li><a target="_blank" href="javascript:verifyByPOST('HEAD');">HEAD</a></li>
			<li><a target="_blank" href="javascript:verifyByPOST('OPTIONS');">OPTIONS</a></li>
			<li><a target="_blank" href="javascript:verifyByPOST('PATCH');">PATCH</a></li>
		</ul>

		<h4>Result:</h4>
		<textarea rows="5" cols="100" id="result"></textarea>
		<br>

		<script>
			var resultEle = document.getElementById("result");
			var verifyByPOST = function(type){
				resultEle.value = 'Pending';
				var url = '/';
				switch(type){
					case 'POST':
						break;
					case 'DELETE':
						url += '?_method=DELETE';
						break;
					case 'PUT':
						url += '?_method=PUT';
						break;
					case 'HEAD':
						url += '?_method=HEAD';
						break;
					case 'OPTIONS':
						url += '?_method=OPTIONS';
						break;
					case 'PATCH':
						url += '?_method=PATCH';
						break;
				}

				xmlHttp = new XMLHttpRequest();
    				xmlHttp.open("POST", url);
				xmlHttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
				xmlHttp.send();
    				xmlHttp.onreadystatechange = function () {
        				resultEle.value = type + ": " + xmlHttp.responseText;
    				}
			}
		</script>
	</body>
	</html>`
	tpl = template.Must(template.New("").Parse(html))
)

type UserController struct {
	clevergo.Controller
	allowOrigin  string
	allowMethods string
}

func NewUserController() UserController {
	return UserController{
		allowOrigin:  "*",
		allowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD",
	}
}

func (c UserController) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		// Do anything what you want.
		ctx.Text("Prepare.\n")

		// Set Access-Control-Allow-Origin and Access-Control-Allow-Methods for ajax request.
		ctx.Response.Header.Set("Access-Control-Allow-Origin", c.allowOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", c.allowMethods)

		// Using param named '_method' to simulate the other request, such as PUT, DELETE etc.
		if !ctx.IsGet() {
			switch string(ctx.FormValue("_method")) {
			case "PUT":
				c.PUT(ctx)
				return
			case "DELETE":
				c.DELETE(ctx)
				return
			case "HEAD":
				c.HEAD(ctx)
				return
			case "OPTIONS":
				c.OPTIONS(ctx)
				return
			case "PATCH":
				c.PATCH(ctx)
				return
			}
		}

		// Invoke the request handler.
		next.Handle(ctx)
	})
}

func (c UserController) GET(ctx *clevergo.Context) {
	ctx.SetContentTypeToHTML()
	tpl.Execute(ctx, nil)
}

func (c UserController) POST(ctx *clevergo.Context) {
	ctx.Text("POST REQUEST.\n")
}

func (c UserController) DELETE(ctx *clevergo.Context) {
	ctx.Text("DELETE REQUEST.\n")
}

func (c UserController) PUT(ctx *clevergo.Context) {
	ctx.Text("PUT REQUEST.\n")
}

func (c UserController) OPTIONS(ctx *clevergo.Context) {
	ctx.Text("OPTIONS REQUEST.\n")
}

func (c UserController) PATCH(ctx *clevergo.Context) {
	ctx.Text("PATCH REQUEST.\n")
}

func (c UserController) HEAD(ctx *clevergo.Context) {
	ctx.Text("HEAD REQUEST.\n")
}

func main() {
	app := clevergo.NewApplication()

	// Create a router instance.
	router := clevergo.NewRouter()
	app.AddRouter("", router)

	// Register route handler.
	router.RegisterController("/", NewUserController())

	// Start server.
	app.Run()
}
