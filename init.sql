CREATE TYPE role_enum AS ENUM ('student', 'teacher', 'admin');

CREATE TABLE  Users (
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Username VARCHAR(64) NOT NULL,
    Email VARCHAR(64) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Role role_enum
);


