package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func listenAndServe(log *logrus.Logger, appData *appData) {
	router := mux.NewRouter().PathPrefix("/v1").Subrouter()

	// /activity
	activitiesRouter := router.Path("/activities").Subrouter()
	activitiesRouter.Path("/{id}").HandlerFunc(getActivitiesGetHandlerFunc(log, appData)).Methods("GET")
	activitiesRouter.Path("").HandlerFunc(getActivitiesGetAllHandlerFunc(log, appData)).Methods("GET")
	activitiesRouter.Path("").HandlerFunc(getActivitiesPostHandlerFunc(log, appData)).Methods("POST")
	activitiesRouter.Path("/{id}").HandlerFunc(getActivitiesPutHandlerFunc(log, appData)).Methods("PUT")
	activitiesRouter.Path("/{id}").HandlerFunc(getActivitiesDeleteHandlerFunc(log, appData)).Methods("DELETE")

	// /workouts
	workoutsRouter := router.Path("/workouts").Subrouter()
	workoutsRouter.Path("/{id}").HandlerFunc(getWorkoutsGetHandlerFunc(log, appData)).Methods("GET")
	workoutsRouter.Path("").HandlerFunc(getWorkoutsGetAllHandlerFunc(log, appData)).Methods("GET")
	workoutsRouter.Path("").HandlerFunc(getWorkoutsPostHandlerFunc(log, appData)).Methods("POST")
	workoutsRouter.Path("/{id}").HandlerFunc(getWorkoutsPutHandlerFunc(log, appData)).Methods("PUT")
	workoutsRouter.Path("/{id}").HandlerFunc(getWorkoutsDeleteHandlerFunc(log, appData)).Methods("DELETE")

	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.WithError(err).Error("error in http.ListenAndServer()")
		}
	}()
}
