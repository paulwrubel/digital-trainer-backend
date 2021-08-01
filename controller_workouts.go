package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type GetWorkoutsResponse struct {
	WorkoutID      string `json:"workout_id"`
	ActivityID     string `json:"activity_id"`
	Timestamp      string `json:"timestamp"`
	CaloriesBurned int    `json:"calories_burned"`
	Duration       int64  `json:"duration"`
}

func getWorkoutsGetHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/workouts/{id}.GET",
			"request_id": requestID,
		})
		log.Debug("request received")

		workoutID := mux.Vars(r)["id"]
		workout := &workout{
			workoutID: workoutID,
		}

		// check if row exists
		err := controllerCheckExists(rw, workout, log, appData)
		if err != nil {
			return
		}

		// get from db
		err = controllerDatabaseFunc(rw, workout, workout.Get, log, appData)
		if err != nil {
			return
		}

		response := GetWorkoutsResponse{
			WorkoutID:      workout.workoutID,
			ActivityID:     workout.activityID,
			Timestamp:      workout.timestamp.Format(time.RFC3339),
			CaloriesBurned: workout.caloriesBurned,
			Duration:       workout.duration.Milliseconds(),
		}
		controllerEncodeResponse(rw, log, http.StatusOK, response)

		log.Debug("request completed")

	}
}

type GetAllWorkoutsResponse []GetAllWorkoutsResponseItem
type GetAllWorkoutsResponseItem struct {
	WorkoutID      string `json:"workout_id"`
	ActivityID     string `json:"activity_id"`
	Timestamp      string `json:"timestamp"`
	CaloriesBurned int    `json:"calories_burned"`
	Duration       int64  `json:"duration"`
}

func getWorkoutsGetAllHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/workouts.GET",
			"request_id": requestID,
		})
		log.Debug("request received")

		// get from db
		persistenceObjects, err := controllerDatabaseGetAll(rw, "workout", log, appData)
		if err != nil {
			return
		}

		response := GetAllWorkoutsResponse{}
		for _, object := range persistenceObjects {
			workout := object.(*workout)
			response = append(response, GetAllWorkoutsResponseItem{
				WorkoutID:      workout.workoutID,
				ActivityID:     workout.activityID,
				Timestamp:      workout.timestamp.Format(time.RFC3339),
				CaloriesBurned: workout.caloriesBurned,
				Duration:       workout.duration.Milliseconds(),
			})
		}

		err = controllerEncodeResponse(rw, log, http.StatusOK, response)
		if err != nil {
			return
		}

		log.Debug("request completed")

	}
}

type PostWorkoutsRequest struct {
	ActivityID     *string `json:"activity_id"`
	Timestamp      *string `json:"timestamp"`
	CaloriesBurned *int    `json:"calories_burned"`
	Duration       *int64  `json:"duration"`
}

type PostWorkoutsResponse struct {
	WorkoutID      string `json:"workout_id"`
	ActivityID     string `json:"activity_id"`
	Timestamp      string `json:"timestamp"`
	CaloriesBurned int    `json:"calories_burned"`
	Duration       int64  `json:"duration"`
}

func getWorkoutsPostHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/workouts.POST",
			"request_id": requestID,
		})
		log.Debug("request received")

		var postWorkoutRequest PostWorkoutsRequest
		err := controllerDecodeRequest(rw, log, r.Body, &postWorkoutRequest)
		if err != nil {
			return
		}

		err = controllerCheckMissingFields(rw, log,
			postWorkoutRequest.ActivityID,
			postWorkoutRequest.CaloriesBurned,
			postWorkoutRequest.Duration,
			postWorkoutRequest.Timestamp)
		if err != nil {
			return
		}

		parsedTime, err := time.Parse(time.RFC3339, *postWorkoutRequest.Timestamp)
		if err != nil {
			writeErrorResponse(rw, http.StatusBadRequest, "invalid timestamp format", err)
		}

		workout := &workout{
			workoutID:      uuid.NewString(),
			activityID:     *postWorkoutRequest.ActivityID,
			timestamp:      parsedTime,
			caloriesBurned: *postWorkoutRequest.CaloriesBurned,
			duration:       time.Duration(*postWorkoutRequest.Duration) * time.Millisecond,
		}

		// check if row exists
		err = controllerCheckExists(rw, workout, log, appData)
		if err != nil {
			return
		}

		// get from db
		err = controllerDatabaseFunc(rw, workout, workout.Get, log, appData)
		if err != nil {
			return
		}

		response := PostWorkoutsResponse{
			WorkoutID:      workout.workoutID,
			ActivityID:     workout.activityID,
			Timestamp:      workout.timestamp.Format(time.RFC3339),
			CaloriesBurned: workout.caloriesBurned,
			Duration:       workout.duration.Milliseconds(),
		}
		err = controllerEncodeResponse(rw, log, http.StatusCreated, response)
		if err != nil {
			return
		}

		log.Debug("request completed")
	}
}

type PutWorkoutsRequest struct {
	ActivityID     *string `json:"activity_id"`
	Timestamp      *string `json:"timestamp"`
	CaloriesBurned *int    `json:"calories_burned"`
	Duration       *int64  `json:"duration"`
}

func getWorkoutsPutHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/workouts/{id}.PUT",
			"request_id": requestID,
		})
		log.Debug("request received")

		workoutID := mux.Vars(r)["id"]
		workout := &workout{
			workoutID: workoutID,
		}

		// check if row exists
		err := controllerCheckExists(rw, workout, log, appData)
		if err != nil {
			return
		}

		var putWorkoutRequest PutWorkoutsRequest
		err = controllerDecodeRequest(rw, log, r.Body, &putWorkoutRequest)
		if err != nil {
			return
		}

		err = controllerCheckMissingFields(rw, log,
			putWorkoutRequest.ActivityID,
			putWorkoutRequest.Timestamp,
			putWorkoutRequest.CaloriesBurned,
			putWorkoutRequest.Duration)
		if err != nil {
			return
		}

		parsedTime, err := time.Parse(time.RFC3339, *putWorkoutRequest.Timestamp)
		if err != nil {
			writeErrorResponse(rw, http.StatusBadRequest, "invalid timestamp format", err)
		}

		workout.activityID = *putWorkoutRequest.ActivityID
		workout.timestamp = parsedTime
		workout.caloriesBurned = *putWorkoutRequest.CaloriesBurned
		workout.duration = time.Duration(*putWorkoutRequest.Duration) * time.Millisecond

		// update in db
		err = controllerDatabaseFunc(rw, workout, workout.Update, log, appData)
		if err != nil {
			return
		}

		rw.WriteHeader(http.StatusNoContent)

		log.Debug("request completed")
	}
}

func getWorkoutsDeleteHandlerFunc(baseLog *logrus.Logger, appData *appData) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()
		log := baseLog.WithFields(logrus.Fields{
			"endpoint":   "/activities/{id}.DELETE",
			"request_id": requestID,
		})
		log.Debug("request received")

		workoutID := mux.Vars(r)["id"]
		workout := &workout{
			workoutID: workoutID,
		}

		// check if row exists
		err := controllerCheckExists(rw, workout, log, appData)
		if err != nil {
			return
		}

		// delete from db
		err = controllerDatabaseFunc(rw, workout, workout.Delete, log, appData)
		if err != nil {
			return
		}

		rw.WriteHeader(http.StatusNoContent)

		log.Debug("request completed")
	}
}
