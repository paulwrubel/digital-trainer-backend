package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type workout struct {
	workoutID  string
	activityID string
	timestamp  time.Time
	calories   int
	duration   time.Duration
}

func (a *workout) Type() string {
	return "workout"
}

func (w *workout) Save(baseLog *logrus.Entry, appData *appData) error {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "workout",
		"event":  "save",
	})
	log.Trace("database event initiated")

	tag, err := appData.db.Exec(context.Background(), `
		INSERT INTO workouts (
			workout_id,
			activity_id,
			timestamp,
			calorie,
			duration
		) VALUES ($1,$2,$3,$4,$5)`,
		w.workoutID,
		w.activityID,
		w.timestamp,
		w.calories,
		w.duration,
	)
	if err != nil || tag.RowsAffected() != 1 {
		return err
	}

	log.Trace("database event completed")
	return nil
}

func (w *workout) Get(baseLog *logrus.Entry, appData *appData) error {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "workout",
		"event":  "get",
	})
	log.Trace("database event initiated")

	err := appData.db.QueryRow(context.Background(), `
		SELECT 
			workout_id,
			activity_id,
			timestamp,
			calorie,
			duration
		FROM workouts
		WHERE workout_id = $1`, w.workoutID).Scan(
		&w.workoutID,
		&w.activityID,
		&w.timestamp,
		&w.calories,
		&w.duration,
	)
	if err != nil {
		return err
	}

	log.Trace("database event completed")
	return nil
}

func (w *workout) Update(baseLog *logrus.Entry, appData *appData) error {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "workout",
		"event":  "update",
	})
	log.Trace("database event initiated")

	tag, err := appData.db.Exec(context.Background(), `
		UPDATE workouts SET (
			workout_id,
			activity_id,
			timestamp,
			calorie,
			duration
		) = ($1,$2,$3,$4,$5)
		WHERE workout_id = $1`,
		w.workoutID,
		w.activityID,
		w.timestamp,
		w.calories,
		w.duration,
	)
	if err != nil || tag.RowsAffected() != 1 {
		return err
	}

	log.Trace("database event completed")
	return nil
}

func (w *workout) Delete(baseLog *logrus.Entry, appData *appData) error {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "workout",
		"event":  "delete",
	})
	log.Trace("database event initiated")

	tag, err := appData.db.Exec(context.Background(), `
		DELETE FROM workouts
		WHERE workout_id = $1`,
		w.workoutID,
	)
	if err != nil || tag.RowsAffected() != 1 {
		return err
	}

	log.Trace("database event completed")
	return nil
}

func (w *workout) Exists(baseLog *logrus.Entry, appData *appData) (bool, error) {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "workout",
		"event":  "exist",
	})
	log.Trace("database event initiated")

	var count int
	err := appData.db.QueryRow(context.Background(), `
		SELECT count(*)
		FROM workouts
		WHERE workout_id = $1`, w.workoutID).Scan(&count)
	if err != nil {
		return false, err
	}

	log.Trace("database event completed")
	return count == 1, nil
}

func getAllWorkouts(baseLog *logrus.Entry, appData *appData) ([]persistenceObject, error) {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "workout",
		"event":  "get all",
	})
	log.Trace("database event initiated")

	rows, err := appData.db.Query(context.Background(), `
		SELECT 
			workout_id,
			activity_id,
			timestamp,
			calorie,
			duration
		FROM workouts`)
	if err != nil {
		return nil, err
	}

	var workouts []persistenceObject
	for rows.Next() {
		w := &workout{}
		err = rows.Scan(
			&w.workoutID,
			&w.activityID,
			&w.timestamp,
			&w.calories,
			&w.duration,
		)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, w)
	}

	log.Trace("database event completed")
	return workouts, nil
}
