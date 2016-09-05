package main

import (
	"crypto"
	"fmt"
	"github.com/clevergo/jwt"
	"github.com/headwindfly/clevergo"
	"github.com/headwindfly/clevergo/middlewares/jwt"
	"html/template"
	"math/rand"
	"strconv"
)

var (
	ttl = int64(60) // Time-to-live of token: 60 seconds.

	issuer = "CleverGo" // JWT issuer.

	j = jwt.NewJWT(issuer, ttl) //JWT manager.

	jwtKey = "_jwt" // The key for getting token from request.

	// JWT Middleware error handler.
	errorHandler = func(ctx *clevergo.Context) {
		ctx.HTML("Sorry, your token is invlaid!")
	}

	// JWT Middleware success handler.
	successHandler = func(ctx *clevergo.Context, token *jwt.Token) {
		// Set audience.
		ctx.SetUserValue("jwt_audience", token.Payload.Aud)
	}

	html = `<html>
	<head></head>
	<body>
		<h3>{{ .Audience }}</h3>

		<h4>JWT token (effective within {{ .TTL }} seconds):</h4>
		<textarea rows="5" cols="100">{{ .Token }}</textarea>
		<br>

		<h4>Verify</h4>
		<ul>
			<li><a target="_blank" href="javascript:verifyByGET('{{ .Token }}');">Verify (GET)</a></li>
			<li><a target="_blank" href="javascript:verifyByPOST(1);">Verify (POST FORM)</a></li>
			<li><a target="_blank" href="javascript:verifyByPOST(2);">Verify (HEADER Authorization)</a></li>
			<li><a target="_blank" href="javascript:verifyByGET('Invalid token....');">Verify (Invalid Token)</a></li>
		</ul>

		<h4>Result:</h4>
		<textarea rows="3" cols="100" id="result"></textarea>
		<br>

		<script>
			var url = '{{ .Url }}';
			var key = '{{ .Key }}';
			var token = '{{ .Token }}';
			var resultEle = document.getElementById("result");

			var verifyByGET = function(token){
				resultEle.value = 'Pending';
				xmlHttp = new XMLHttpRequest();
    				xmlHttp.open("GET", url + "?" + key + "=" + token);
    				xmlHttp.send(null);
    				xmlHttp.onreadystatechange = function () {
        				resultEle.value = "GET: " + xmlHttp.responseText;
    				}
			}

			var verifyByPOST = function(type){
				resultEle.value = 'Pending';
				xmlHttp = new XMLHttpRequest();
    				xmlHttp.open("POST", url);
				xmlHttp.setRequestHeader("Content-Type","application/x-www-form-urlencoded");
				if (type == 1) {
					typeText = "POST FORM"
					xmlHttp.send(formatParams({"{{ .Key }}":token}));
				} else {
					typeText = "HEADER Authorization"
					xmlHttp.setRequestHeader("Authorization","Bearer " + token);
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

func init() {
	// Add HMACAlgorithm.
	hs256, err := jwt.NewHMACAlgorithm(crypto.SHA256, []byte("secrey key"))
	if err != nil {
		panic(err)
	}
	j.AddAlgorithm("HS256", hs256)
}

func index(ctx *clevergo.Context) {
	audience := "Audience_" + strconv.Itoa(rand.Intn(100))
	token, err := j.NewToken("HS256", "CleverGO", audience)
	if err != nil {
		ctx.Textf("Failed to generate token: %s", err.Error())
		return
	}

	// Parse the token in order to get raw token.
	token.Parse()

	ctx.SetContentTypeToHTML()
	tpl.Execute(ctx, map[string]interface{}{
		"Token":    token.Raw.Token(),
		"TTL":      ttl,
		"Audience": audience,
		"Key":      jwtKey,
		"Url":      "/verify",
	})

}

func verify(ctx *clevergo.Context) {
	// Get audience.
	ctx.HTML(fmt.Sprintf("Hello %s.", ctx.UserValue("jwt_audience")))
}

func main() {
	// Ceeate a application instance.
	app := clevergo.NewApplication()

	// Create a router instance.
	router := clevergo.NewRouter()

	// Add and set this router as application's default router.
	app.AddRouter("", router)

	// Note that.
	// Before registering middleware, we should register index handler first.
	// In order to cross over the JWT Middleware,
	// otherwise the index handler will be blocked by the JWT Middleware.
	router.GET("/", clevergo.HandlerFunc(index))

	// Add JWT Middleware
	jwtMiddleware := jwtmiddleware.NewJWTMiddleware(j) // JWT middleware.
	jwtMiddleware.SetKey(jwtKey)
	jwtMiddleware.SetErrorHandler(errorHandler)
	jwtMiddleware.SetSuccessHandler(successHandler)

	router.AddMiddleware(jwtMiddleware)

	// Register verify handler.
	router.GET("/verify", clevergo.HandlerFunc(verify))
	router.POST("/verify", clevergo.HandlerFunc(verify))

	// Start application.
	app.Run()
}
