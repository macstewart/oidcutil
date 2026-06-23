package cmd

import (
	"embed"
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"ssorry/internal/http/client"
	"ssorry/internal/http/idp"
	"ssorry/internal/store"
	"ssorry/internal/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var resources embed.FS

var (
	serverCmd     = &cobra.Command{Use: "server", Run: start}
	port          *int
	hostname      *string
	local         *bool
	tokenCallback = make(chan string)
)

func StoreResources(embedded embed.FS) {
	resources = embedded
}

func init() {
	port = serverCmd.Flags().IntP("port", "p", 3333, "Port to run the server on")
	hostname = serverCmd.Flags().StringP("hostname", "H", "http://localhost:3333", "Hostname publish in the discovery")
	local = serverCmd.Flags().BoolP("local", "l", false, "Run the service with local authorize page")
	rootCmd.AddCommand(serverCmd)
	store.SetCallback(tokenCallback)
}

func start(cmd *cobra.Command, args []string) {
	rtr := router()
	log.Printf("Server listening on %s", *hostname)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}

func router() *gin.Engine {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(resources, "resources/html/*.tmpl"))
	router.SetHTMLTemplate(templ)
	gob.Register(map[string]interface{}{})
	router.SetFuncMap(template.FuncMap{
		"StringsJoin": strings.Join,
	})
	// router.GET("/callback", client.FetchCallback(callbackChan))
	// router.GET("/sso/login", client.IdpCallback(callbackChan))
	router.GET("/discovery", idp.ServeDiscovery(util.BuildDiscovery(*local, *hostname)))
	router.GET("/authorize", idp.Authorize())
	router.GET("/keys", idp.HandleJwks())
	router.POST("/sso/login", client.Login())
	router.GET("/gettoken", client.FetchTokenCallback(tokenCallback))
	router.POST("/token", idp.ServeToken())
	router.POST("/user", idp.UpdateUser())
	return router
}
