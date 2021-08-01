CREATE TABLE activities (
    activity_id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE workouts (
    workout_id TEXT PRIMARY KEY,
    activity_id TEXT NOT NULL REFERENCES activities(activity_id),
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    calories INTEGER NOT NULL,
    duration INTERVAL NOT NULL
);