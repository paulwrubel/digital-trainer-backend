package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func listenAndServe(log *logrus.Logger, appData *appData) {
	router := mux.NewRouter().PathPrefix("/v1").Subrouter()

	// /activity
	router.Path("/activity/{id}").HandlerFunc(getActivitiesGetHandlerFunc(log, appData)).Methods("GET")
	router.Path("/activity").HandlerFunc(getActivitiesGetAllHandlerFunc(log, appData)).Methods("GET")
	router.Path("/activity").HandlerFunc(getActivitiesPostHandlerFunc(log, appData)).Methods("POST")
	router.Path("/activity/{id}").HandlerFunc(getActivitiesPutHandlerFunc(log, appData)).Methods("PUT")
	router.Path("/activity/{id}").HandlerFunc(getActivitiesDeleteHandlerFunc(log, appData)).Methods("DELETE")

	// /workouts
	router.Path("/workouts/{id}").HandlerFunc(getWorkoutsGetHandlerFunc(log, appData)).Methods("GET")
	router.Path("/workouts").HandlerFunc(getWorkoutsGetAllHandlerFunc(log, appData)).Methods("GET")
	router.Path("/workouts").HandlerFunc(getWorkoutsPostHandlerFunc(log, appData)).Methods("POST")
	router.Path("/workouts/{id}").HandlerFunc(getWorkoutsPutHandlerFunc(log, appData)).Methods("PUT")
	router.Path("/workouts/{id}").HandlerFunc(getWorkoutsDeleteHandlerFunc(log, appData)).Methods("DELETE")

	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.WithError(err).Error("error in http.ListenAndServer()")
		}
	}()
}
