DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;


CREATE TABLE  Users (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Email VARCHAR(64) NOT NULL UNIQUE,
    Username VARCHAR(64) NOT NULL,
    Password VARCHAR(255) NOT NULL
);

CREATE TABLE Rooms (
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    room_name TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE Quizzes (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Content JSONB
);

CREATE TABLE Room_Student (
    room_id TEXT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    student_name TEXT NOT NULL UNIQUE
);

CREATE TABLE Room_Quiz (
    room_id TEXT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    quiz_id UUID NOT NULL REFERENCES quizzes(id) ON DELETE CASCADE,
    PRIMARY KEY (room_id, quiz_id)
);

CREATE TABLE student_scores (
    quiz_id UUID NOT NULL REFERENCES quizzes(id) ON DELETE CASCADE,
    student_name TEXT NOT NULL,
    score REAL NOT NULL,
    submitted_at TIMESTAMP DEFAULT now(),
    PRIMARY KEY (quiz_id, student_name)
);
