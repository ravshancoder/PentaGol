CREATE TABLE IF NOT EXISTS admins (
    id            SERIAL,
    name          VARCHAR(50) NOT NULL,
    email         VARCHAR(100) NOT NULL UNIQUE,
    password      TEXT NOT NULL,
    refresh_token TEXT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIME
);

INSERT INTO admins(name, email, password, refresh_token) VALUES('ravshan', 'mavlonovr555@gmail.com', '$2a$14$synOItEs8oaSXPgA13wFpetq1pJZBNf9mqTvd4XyIhsZy5R98v0UG', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.exw43QM2_SlQ_MC3zIZ1XyUvnOldQy7CxqRYsvUI4A0');
