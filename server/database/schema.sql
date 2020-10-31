CREATE TABLE tasks (
    id text PRIMARY KEY NOT NULL,
    taskName TEXT NOT NULL UNIQUE,
    taskDescription  TEXT NOT NULL,
    done BOOLEAN NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE
);
