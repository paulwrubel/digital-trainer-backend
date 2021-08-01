package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ErrorInfo struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type ErrorResponse struct {
	ErrorInfo *ErrorInfo `json:"error"`
}

func writeErrorResponse(response http.ResponseWriter, statusCode int, errorMessage string, err error) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	errorString := ""
	if err != nil {
		errorString = err.Error()
	}

	json.NewEncoder(response).Encode(
		ErrorResponse{
			ErrorInfo: &ErrorInfo{
				Message: errorMessage,
				Error:   errorString,
			},
		},
	)
}

func controllerCheckExists(rw http.ResponseWriter, o persistenceObject, log *logrus.Entry, appData *appData) error {
	// check if row exists
	exists, err := o.Exists(log, appData)
	if err != nil {
		errorMessage := "error checking " + o.Type() + " existence in database"
		errorStatusCode := http.StatusInternalServerError

		log.WithError(err).Error(errorMessage)
		writeErrorResponse(rw, errorStatusCode, errorMessage, err)
		return fmt.Errorf("error checking existence: %w", err)
	}
	if !exists {
		errorMessage := o.Type() + " does not exist"
		errorStatusCode := http.StatusNotFound

		log.Error(errorMessage)
		writeErrorResponse(rw, errorStatusCode, errorMessage, nil)
		return fmt.Errorf("does not exist")
	}
	return nil
}

func controllerCheckMissingFields(rw http.ResponseWriter, log *logrus.Entry, fields ...interface{}) error {
	// check for missing fields
	for _, field := range fields {
		if field == nil {
			errorMessage := "missing field from request"
			errorStatusCode := http.StatusBadRequest

			log.Error(errorMessage)
			writeErrorResponse(rw, errorStatusCode, errorMessage, nil)
			return fmt.Errorf("missing field from request")
		}
	}
	return nil
}

func controllerDatabaseFunc(rw http.ResponseWriter, o persistenceObject, oFunc func(*logrus.Entry, *appData) error, log *logrus.Entry, appData *appData) error {
	err := oFunc(log, appData)
	if err != nil {
		errorMessage := "error getting " + o.Type() + " from database"
		errorStatusCode := http.StatusInternalServerError

		log.WithError(err).Error(errorMessage)
		writeErrorResponse(rw, errorStatusCode, errorMessage, err)
		return fmt.Errorf("error getting from existence: %w", err)
	}
	return nil
}

func controllerDatabaseGetAll(rw http.ResponseWriter, persistenceObjectType string, log *logrus.Entry, appData *appData) ([]persistenceObject, error) {
	var persistenceObjects []persistenceObject
	var err error
	switch persistenceObjectType {
	case "activity":
		persistenceObjects, err = getAllActivities(log, appData)
	case "workout":
		persistenceObjects, err = getAllWorkouts(log, appData)
	default:
		err = fmt.Errorf("unknown persistence object type: this is a server error and reflects no invalid client action")
	}
	if err != nil {
		errorMessage := "error getting all of type " + persistenceObjectType + " from database"
		errorStatusCode := http.StatusInternalServerError

		log.WithError(err).Error(errorMessage)
		writeErrorResponse(rw, errorStatusCode, errorMessage, err)
		return nil, fmt.Errorf("error getting from existence: %w", err)
	}
	return persistenceObjects, nil
}

func controllerDecodeRequest(rw http.ResponseWriter, log *logrus.Entry, rc io.ReadCloser, v interface{}) error {
	// decode request
	err := json.NewDecoder(rc).Decode(v)
	if err != nil {
		errorMessage := "error decoding request body"
		errorStatusCode := http.StatusBadRequest

		log.WithError(err).Error(errorMessage)
		writeErrorResponse(rw, errorStatusCode, errorMessage, err)
		return fmt.Errorf("error decoding request body")
	}
	return nil
}

func controllerEncodeResponse(rw http.ResponseWriter, log *logrus.Entry, statusCode int, v interface{}) error {
	// encode response
	rw.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(v)
	if err != nil {
		errorMessage := "error encoding response body"
		errorStatusCode := http.StatusInternalServerError

		log.WithError(err).Error(errorMessage)
		writeErrorResponse(rw, errorStatusCode, errorMessage, err)
		return fmt.Errorf("error encoding response body")
	}
	return nil
}
