package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"clevergo.tech/captchas"
	"clevergo.tech/captchas/drivers"
	"clevergo.tech/captchas/stores/memcachedstore"
	"clevergo.tech/captchas/stores/memstore"
	"clevergo.tech/captchas/stores/redisstore"
	"clevergo.tech/clevergo"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v7"
)

var (
	addr      = flag.String("addr", "localhost:8080", "address")
	storeStr  = flag.String("store", "memory", "store: memory, redis, memcached")
	store     captchas.Store
	managers  map[string]*captchas.Manager
	indexTmpl = template.Must(template.ParseFiles("layout.tmpl", "index.tmpl"))
	apiTmpl   = template.Must(template.ParseFiles("layout.tmpl", "api.tmpl"))
)

func initStore() (err error) {
	switch *storeStr {
	case "memcached":
		memcachedClient := memcache.New("localhost:11211")
		store = memcachedstore.New(
			memcachedClient,
			memcachedstore.Expiration(int32(600)), // captcha expiration, optional.
			memcachedstore.Prefix("captchas"),     // key prefix, optional.
		)
	case "redis":
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		if _, err := client.Ping().Result(); err != nil {
			return err
		}
		store = redisstore.New(
			client,
			redisstore.Expiration(10*time.Minute), // key expiration, optional.
			redisstore.Prefix("captchas"),         // key prefix, optional.
		)
	default:
		store = memstore.New(
			memstore.Expiration(10*time.Minute), // expiration, optional.
			memstore.GCInterval(time.Minute),    // garbage collection interval, optional.
		)
	}
	return
}

func main() {
	flag.Parse()
	if err := initStore(); err != nil {
		log.Fatal(err)
	}

	managerOpts := []captchas.Option{
		// disable case sensitive, enabled by default, it will effects on string driver.
		captchas.CaseSensitive(false),
	}
	managers = map[string]*captchas.Manager{
		"digit":   captchas.New(store, drivers.NewDigit(), managerOpts...),
		"audio":   captchas.New(store, drivers.NewAudio(), managerOpts...),
		"math":    captchas.New(store, drivers.NewMath(), managerOpts...),
		"string":  captchas.New(store, drivers.NewString(), managerOpts...),
		"chinese": captchas.New(store, drivers.NewChinese(), managerOpts...),
	}

	app := clevergo.New()
	app.Any("/", index)
	app.Get("/api", api)
	app.Post("/validate", validate)
	app.Any("/generate", generate)
	app.Run(":8080")
}

func index(c *clevergo.Context) error {
	manager, err := getManager(c.Request)
	if err != nil {
		return err
	}

	ctx := c.Context()
	captcha, err := manager.Generate(ctx)
	if err != nil {
		return err
	}

	alert := ""
	valid := false
	if c.IsPost() {
		captchaID := c.PostFormValue("captcha_id")
		captchaVal := c.PostFormValue("captcha")
		if err := manager.Verify(ctx, captchaID, captchaVal, true); err != nil {
			alert = err.Error()
		} else {
			valid = true
			alert = "captcha is valid"
		}
	}

	return render(indexTmpl, c.Response, map[string]interface{}{
		"driver":  c.QueryParam("driver"),
		"captcha": captcha,
		"alert":   alert,
		"valid":   valid,
	})
}

func api(c *clevergo.Context) error {
	return render(apiTmpl, c.Response, map[string]interface{}{
		"driver": c.QueryParam("driver"),
	})
}

func generate(c *clevergo.Context) error {
	manager, err := getManager(c.Request)
	if err != nil {
		return err
	}

	captcha, err := manager.Generate(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, clevergo.Map{
		"id":   captcha.ID(),
		"data": captcha.EncodeToString(),
	})
}

func validate(c *clevergo.Context) error {
	manager, err := getManager(c.Request)
	if err != nil {
		return err
	}

	captchaID := c.PostFormValue("captcha_id")
	captchaValue := c.PostFormValue("captcha")
	err = manager.Verify(c.Context(), captchaID, captchaValue, true)
	if err != nil {
		return c.JSON(http.StatusOK, clevergo.Map{
			"msg": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, clevergo.Map{
		"msg": "success",
	})
}

func getManager(r *http.Request) (*captchas.Manager, error) {
	driver := r.URL.Query().Get("driver")
	if driver == "" {
		driver = "digit"
	}

	if m, ok := managers[driver]; ok {
		return m, nil
	}

	return nil, fmt.Errorf("unsupported driver: %s", driver)
}

func render(tmpl *template.Template, w http.ResponseWriter, data interface{}) error {
	return tmpl.Execute(w, data)
}
