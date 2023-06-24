package cmd

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"ssorry/internal/http/idp"
	"ssorry/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{Use: "server", Run: start}
	port      *int
	hostname  string
)

func init() {
	port = serverCmd.Flags().IntP("port", "p", 3333, "Port to run the server on")
	rootCmd.AddCommand(serverCmd)
}

func start(cmd *cobra.Command, args []string) {
	hostname = fmt.Sprintf("http://localhost:%d", *port)
	rtr := router()
	log.Printf("Server listening on %s", hostname)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}

func router() *gin.Engine {
	router := gin.Default()
	gob.Register(map[string]interface{}{})
	// router.GET("/callback", client.FetchCallback(callbackChan))
	// router.GET("/sso/login", client.IdpCallback(callbackChan))
	router.GET("/discovery", idp.ServeDiscovery(util.BuildDiscovery(hostname)))
	router.GET("/authorize", idp.Authorize())
	router.GET("/token", idp.ServeToken())
	return router
}
