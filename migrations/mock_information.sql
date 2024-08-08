-- Insert data into students table
INSERT INTO students (phone_number, full_name)
VALUES
    ('123-456-7890', 'Alice Johnson'),
    ('234-567-8901', 'Bob Smith'),
    ('345-678-9012', 'Charlie Brown'),
    ('456-789-0123', 'David Williams'),
    ('567-890-1234', 'Eva Davis'),
    ('678-901-2345', 'Frank Miller'),
    ('789-012-3456', 'Grace Wilson'),
    ('890-123-4567', 'Hannah Moore'),
    ('901-234-5678', 'Ian Taylor'),
    ('012-345-6789', 'Jane Anderson');

-- Insert data into collections table
INSERT INTO collections (name)
VALUES
    ('Collection A'),
    ('Collection B'),
    ('Collection C'),
    ('Collection D'),
    ('Collection E'),
    ('Collection F'),
    ('Collection G'),
    ('Collection H'),
    ('Collection I'),
    ('Collection J');

-- Insert data into questions table
INSERT INTO questions (question_field, collection_id)
VALUES
    ('What is 2 + 2?', 2),
    ('What is the capital of France?', 2),
    ('What is the boiling point of water?', 2),
    ('Who wrote "To Kill a Mockingbird"?', 2),
    ('What is the largest planet in our solar system?', 2),
    ('What is the chemical symbol for gold?', 2),
    ('What is the speed of light?', 2),
    ('Who painted the Mona Lisa?', 2),
    ('What is the hardest natural substance on Earth?', 2),
    ('What is the currency of Japan?', 2);

-- Insert data into answers table
INSERT INTO answers (is_true, question_id, answer_field)
VALUES
    (TRUE, 1, 'two chewing gums'),
    (FALSE, 1, 'two chewing gums'),
    (FALSE, 1, 'two chewing gums'),
    (FALSE, 2, 'two chewing gums'),
    (TRUE, 2, 'two chewing gums'),
    (FALSE, 2, 'two chewing gums'),
    (TRUE, 3, 'two chewing gums'),
    (FALSE, 3, 'two chewing gums'),
    (FALSE, 3, 'two chewing gums'),
    (FALSE, 4, 'two chewing gums'),
    (TRUE, 4, 'two chewing gums'),
    (FALSE, 4, 'two chewing gums');

-- Insert data into groups table
INSERT INTO groups (name, teacher_name, level)
VALUES
    ('Group 1', 'Mr. Adams', 'BEGINNER'),
    ('Group 2', 'Ms. Brown', 'ELEMENTARY'),
    ('Group 3', 'Mr. Clark', 'PRE_INTERMEDIATE'),
    ('Group 4', 'Ms. Davis', 'INTERMEDIATE'),
    ('Group 5', 'Mr. Evans', 'UPPER_INTERMEDIATE'),
    ('Group 6', 'Ms. Foster', 'ADVANCED'),
    ('Group 7', 'Mr. Gray', 'PROFICIENT'),
    ('Group 8', 'Ms. Harris', 'BEGINNER'),
    ('Group 9', 'Mr. Ives', 'ELEMENTARY'),
    ('Group 10', 'Ms. Jones', 'PRE_INTERMEDIATE');
