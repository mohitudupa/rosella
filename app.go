package main

import (
	"log"
	"net/http"

	"githib.com/mohitudupa/rosella/data"
	"githib.com/mohitudupa/rosella/handlers"
	"githib.com/mohitudupa/rosella/middlewares"
	"githib.com/mohitudupa/rosella/utils"
)

func getRepository(applicationConfig *utils.ApplicationConfig) data.Repository {
	switch applicationConfig.RepositoryBackend {
	case utils.FileRepositoryBackend:
		return data.NewFileRepository(applicationConfig.FileRepository.Path)
	default:
		if applicationConfig.Server.ReadOnly {
			log.Printf("WARN: server will be set to readOnly with an empty repository")
		}
		return data.NewDefaultRepository()
	}
}

func getServer(repository data.Repository, readOnly bool) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc(handlers.HealthyUrl, handlers.Healthy)

	flagsHandler := handlers.NewFlagsHandler(repository)
	router.HandleFunc(flagsHandler.ListURL, flagsHandler.List)
	router.HandleFunc(flagsHandler.GetURL, flagsHandler.Get)

	limitsHandler := handlers.NewLimitsHandler(repository)
	router.HandleFunc(limitsHandler.ListURL, limitsHandler.List)
	router.HandleFunc(limitsHandler.GetURL, limitsHandler.Get)

	valuesHandler := handlers.NewValuesHandler(repository)
	router.HandleFunc(valuesHandler.ListURL, valuesHandler.List)
	router.HandleFunc(valuesHandler.GetURL, valuesHandler.Get)

	configsHandler := handlers.NewConfigsHandler(repository)
	router.HandleFunc(configsHandler.ListURL, configsHandler.List)
	router.HandleFunc(configsHandler.GetURL, configsHandler.Get)

	if !readOnly {
		router.HandleFunc(flagsHandler.PostURL, flagsHandler.Post)
		router.HandleFunc(flagsHandler.DeleteURL, flagsHandler.Delete)

		router.HandleFunc(limitsHandler.PostURL, limitsHandler.Post)
		router.HandleFunc(limitsHandler.DeleteURL, limitsHandler.Delete)

		router.HandleFunc(valuesHandler.PostURL, valuesHandler.Post)
		router.HandleFunc(valuesHandler.DeleteURL, valuesHandler.Delete)

		router.HandleFunc(configsHandler.PostURL, configsHandler.Post)
		router.HandleFunc(configsHandler.DeleteURL, configsHandler.Delete)
	}

	return middlewares.LoggerMiddleware(router)
}

func getServerLocation(applicationConfig *utils.ApplicationConfig) string {
	return applicationConfig.Server.Host + ":" + applicationConfig.Server.Port
}
