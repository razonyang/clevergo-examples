package main

import (
	"github.com/clevergo/sessions"
	"github.com/headwindfly/clevergo"
	"github.com/headwindfly/clevergo/middlewares/csrf"
	"github.com/headwindfly/clevergo/middlewares/session"
	"html/template"
)

var (
	sessionMiddleware = sessionmiddleware.NewSessionMiddleware("GOSESSION")
	csrfMiddleware    = csrfmiddleware.NewCSRFMiddleware()
	html              = `<html>
	<head></head>
	<body>
		<h3>CRSF Protection</h3>

		<h4>Encoded Token:</h4>
		<input size="100" value="{{ .Token }}">
		<br>

		<h4>Verify</h4>
		<ul>
			<li><a target="_blank" href="javascript:verifyByPOST(1,'{{ .Token }}');">Verify (POST FORM)</a></li>
			<li><a target="_blank" href="javascript:verifyByPOST(2,'{{ .Token }}');">Verify (HEADER)</a></li>
			<li><a target="_blank" href="javascript:verifyByPOST(1,'Invalid token....');">Verify (Invalid Token)</a></li>
		</ul>

		<h4>Result:</h4>
		<textarea rows="3" cols="80" id="result"></textarea>
		<br>

		<script>
			var url = '/';
			var formKey = '{{ .FormKey }}';
			var headerKey = '{{ .HeaderKey }}';
			var resultEle = document.getElementById("result");

			var verifyByPOST = function(type,token){
				resultEle.value = 'Pending';
				xmlHttp = new XMLHttpRequest();
    				xmlHttp.open("POST", url);
				switch(type){
					case 1:
						typeText = "POST FORM";
						xmlHttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
						xmlHttp.send(formatParams({"{{ .FormKey }}":token}));
					break;
					case 2:
						typeText = "HEADER"
						xmlHttp.setRequestHeader("{{ .HeaderKey }}", token);
						xmlHttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
						xmlHttp.send();
				}
    				xmlHttp.onreadystatechange = function () {
        				resultEle.value = typeText + ": " + xmlHttp.responseText;
    				}
			}


    			function formatParams(data) {
        			var arr = [];
        			for (var name in data) {
            				arr.push(encodeURIComponent(name) + "=" + encodeURIComponent(data[name]));
        			}
        			arr.push(("v=" + Math.random()).replace(".",""));
        			return arr.join("&");
    			}
		</script>
	</body>
	</html>`
	tpl = template.Must(template.New("").Parse(html))
)

func get(ctx *clevergo.Context) {
	ctx.SetContentTypeToHTML()
	tpl.Execute(ctx, map[string]interface{}{
		"Token":     ctx.UserValue(csrfMiddleware.Key()),
		"HeaderKey": csrfMiddleware.HeaderKey(),
		"FormKey":   csrfMiddleware.FormKey(),
	})
}

func post(ctx *clevergo.Context) {
	ctx.HTML("Congratulation! Your token is valid.")
}

func main() {
	app := clevergo.NewApplication()

	// Create a router instance.
	router := clevergo.NewRouter()

	// Set the router as application default router.
	app.AddRouter("", router)

	// Set session store.
	// Create a redis session store.
	store := sessions.NewCookieStore([]byte("SecretKey"))

	// Set session store.
	router.SetSessionStore(store)

	// Add Session middleware, this middleware is necessary.
	router.AddMiddleware(sessionMiddleware)

	// Add CSRF middleware.
	csrfMiddleware.SetErrorHandler(func(ctx *clevergo.Context) {
		ctx.SetStatusCode(403)
		ctx.SetBodyString("Invalid token.")
	})
	router.AddMiddleware(csrfMiddleware)

	// Register route handler.
	router.GET("/", clevergo.HandlerFunc(get))
	router.POST("/", clevergo.HandlerFunc(post))

	// Start server.
	app.Run()
}
