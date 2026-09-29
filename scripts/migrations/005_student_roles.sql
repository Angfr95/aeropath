ALTER TABLE students
    ADD COLUMN role TEXT NOT NULL DEFAULT 'student'
    CHECK (role IN ('student', 'admin'));
