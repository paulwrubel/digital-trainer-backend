package main

import (
	"context"

	"github.com/sirupsen/logrus"
)

type activity struct {
	activityID string
	name       string
}

func (a *activity) Type() string {
	return "activity"
}

func (a *activity) Save(baseLog *logrus.Entry, appData *appData) error {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "activity",
		"event":  "save",
	})
	log.Trace("database event initiated")

	tag, err := appData.db.Exec(context.Background(), `
		INSERT INTO activities (
			activity_id,
			name
		) VALUES ($1,$2)`,
		a.activityID,
		a.name,
	)
	if err != nil || tag.RowsAffected() != 1 {
		return err
	}

	log.Trace("database event completed")
	return nil
}

func (a *activity) Get(baseLog *logrus.Entry, appData *appData) error {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "activity",
		"event":  "get",
	})
	log.Trace("database event initiated")

	err := appData.db.QueryRow(context.Background(), `
		SELECT 
			activity_id,
			name
		FROM activities
		WHERE activity_id = $1`, a.activityID).Scan(
		&a.activityID,
		&a.name,
	)
	if err != nil {
		return err
	}

	log.Trace("database event completed")
	return nil
}

func (a *activity) Update(baseLog *logrus.Entry, appData *appData) error {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "activity",
		"event":  "update",
	})
	log.Trace("database event initiated")

	tag, err := appData.db.Exec(context.Background(), `
		UPDATE activities SET (
			activity_id,
			name
		) = ($1,$2)
		WHERE activity_id = $1`,
		a.activityID,
		a.name,
	)
	if err != nil || tag.RowsAffected() != 1 {
		return err
	}

	log.Trace("database event completed")
	return nil
}

func (a *activity) Delete(baseLog *logrus.Entry, appData *appData) error {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "activity",
		"event":  "delete",
	})
	log.Trace("database event initiated")

	tag, err := appData.db.Exec(context.Background(), `
		DELETE FROM activities
		WHERE activity_id = $1`,
		a.activityID,
	)
	if err != nil || tag.RowsAffected() != 1 {
		return err
	}

	log.Trace("database event completed")
	return nil
}

func (a *activity) Exists(baseLog *logrus.Entry, appData *appData) (bool, error) {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "activity",
		"event":  "exist",
	})
	log.Trace("database event initiated")

	var count int
	err := appData.db.QueryRow(context.Background(), `
		SELECT count(*)
		FROM activities
		WHERE activity_id = $1`, a.activityID).Scan(&count)
	if err != nil {
		return false, err
	}

	log.Trace("database event completed")
	return count == 1, nil
}

func getAllActivities(baseLog *logrus.Entry, appData *appData) ([]persistenceObject, error) {
	log := baseLog.WithFields(logrus.Fields{
		"entity": "activity",
		"event":  "get all",
	})
	log.Trace("database event initiated")

	rows, err := appData.db.Query(context.Background(), `
		SELECT 
			activity_id,
			name
		FROM activities`)
	if err != nil {
		return nil, err
	}

	var activities []persistenceObject
	for rows.Next() {
		a := &activity{}
		err = rows.Scan(
			&a.activityID,
			&a.name,
		)
		if err != nil {
			return nil, err
		}
		activities = append(activities, a)
	}

	log.Trace("database event completed")
	return activities, nil
}
