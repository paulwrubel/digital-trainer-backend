package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type GetActivitiesResponse struct {
	ActivityID string `json:"activity_id"`
	Name       string `json:"name"`
}

func getActivitiesGetHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/activities/{id}.GET",
			"request_id": requestID,
		})
		log.Debug("request received")

		activityID := mux.Vars(r)["id"]
		activity := &activity{
			activityID: activityID,
		}

		// check if row exists
		err := controllerCheckExists(rw, activity, log, appData)
		if err != nil {
			return
		}

		// get from db
		err = controllerDatabaseFunc(rw, activity, activity.Get, log, appData)
		if err != nil {
			return
		}

		response := GetActivitiesResponse{
			ActivityID: activity.activityID,
			Name:       activity.name,
		}
		controllerEncodeResponse(rw, log, http.StatusOK, response)

		log.Debug("request completed")

	}
}

type GetAllActivitiesResponse []GetAllActivitiesResponseItem
type GetAllActivitiesResponseItem struct {
	ActivityID string `json:"activity_id"`
	Name       string `json:"name"`
}

func getActivitiesGetAllHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/activities.GET",
			"request_id": requestID,
		})
		log.Debug("request received")

		// get from db
		persistenceObjects, err := controllerDatabaseGetAll(rw, "activity", log, appData)
		if err != nil {
			return
		}

		var response GetAllActivitiesResponse
		for _, object := range persistenceObjects {
			activity := object.(*activity)
			response = append(response, GetAllActivitiesResponseItem{
				ActivityID: activity.activityID,
				Name:       activity.name,
			})
		}

		err = controllerEncodeResponse(rw, log, http.StatusOK, response)
		if err != nil {
			return
		}

		log.Debug("request completed")

	}
}

type PostActivitiesRequest struct {
	Name *string `json:"name"`
}

type PostActivitiesResponse struct {
	ActivityID string `json:"activity_id"`
	Name       string `json:"name"`
}

func getActivitiesPostHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/activities.POST",
			"request_id": requestID,
		})
		log.Debug("request received")

		var postActivityRequest PostActivitiesRequest
		err := controllerDecodeRequest(rw, log, r.Body, &postActivityRequest)
		if err != nil {
			return
		}

		err = controllerCheckMissingFields(rw, log, postActivityRequest.Name)
		if err != nil {
			return
		}

		activity := &activity{
			activityID: uuid.NewString(),
			name:       *postActivityRequest.Name,
		}

		// check if row exists
		err = controllerCheckExists(rw, activity, log, appData)
		if err != nil {
			return
		}

		// get from db
		err = controllerDatabaseFunc(rw, activity, activity.Get, log, appData)
		if err != nil {
			return
		}

		response := PostActivitiesResponse{
			ActivityID: activity.activityID,
			Name:       activity.name,
		}
		err = controllerEncodeResponse(rw, log, http.StatusCreated, response)
		if err != nil {
			return
		}

		log.Debug("request completed")
	}
}

type PutActivitiesRequest struct {
	Name *string `json:"name"`
}

func getActivitiesPutHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/activities/{id}.PUT",
			"request_id": requestID,
		})
		log.Debug("request received")

		activityID := mux.Vars(r)["id"]
		activity := &activity{
			activityID: activityID,
		}

		// check if row exists
		err := controllerCheckExists(rw, activity, log, appData)
		if err != nil {
			return
		}

		var putActivityRequest PutActivitiesRequest
		err = controllerDecodeRequest(rw, log, r.Body, &putActivityRequest)
		if err != nil {
			return
		}

		err = controllerCheckMissingFields(rw, log, putActivityRequest.Name)
		if err != nil {
			return
		}

		activity.name = *putActivityRequest.Name

		// update in db
		err = controllerDatabaseFunc(rw, activity, activity.Update, log, appData)
		if err != nil {
			return
		}

		rw.WriteHeader(http.StatusNoContent)

		log.Debug("request completed")
	}
}

func getActivitiesDeleteHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/activities/{id}.DELETE",
			"request_id": requestID,
		})
		log.Debug("request received")

		activityID := mux.Vars(r)["id"]
		activity := &activity{
			activityID: activityID,
		}

		// check if row exists
		err := controllerCheckExists(rw, activity, log, appData)
		if err != nil {
			return
		}

		// delete from db
		err = controllerDatabaseFunc(rw, activity, activity.Delete, log, appData)
		if err != nil {
			return
		}

		rw.WriteHeader(http.StatusNoContent)

		log.Debug("request completed")
	}
}
