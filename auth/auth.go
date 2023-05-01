package auth

import (
	"github.com/fazarrahman/cognotiv/lib"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func Init() {
	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set(lib.GetEnv("AUTH_CLIENT_ID"), &models.Client{
		ID:     lib.GetEnv("AUTH_CLIENT_ID"),
		Secret: lib.GetEnv("AUTH_SECRET"),
		Domain: lib.GetEnv("APP_DOMAIN"),
	})
	manager.MapClientStorage(clientStore)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)
}
