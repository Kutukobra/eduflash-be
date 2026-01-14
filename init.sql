CREATE TYPE role_enum AS ENUM ('student', 'teacher', 'admin');

CREATE TABLE  Users (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Email VARCHAR(64) NOT 3NULL UNIQUE,
    Username VARCHAR(64) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Role role_enum
);

CREATE TABLE Rooms (
    id TEXT PRIMARY KEY NOT NULL UNIQUE,
    room_name TEXT,
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE Quizzes (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Questions TEXT NOT NULL
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
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_name TEXT NOT NULL,
    quiz_id UUID NOT NULL REFERENCES quizzes(id) ON DELETE CASCADE,
    room_id TEXT REFERENCES rooms(id) ON DELETE CASCADE,
    score REAL NOT NULL,
    submitted_at TIMESTAMP DEFAULT now()
);
