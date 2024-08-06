# English Proficiency Testing System

## Overview

The **English Proficiency Testing System** is designed to assess pupils' English skills, provide true/false responses to
test questions, and assign proficiency levels. Additionally, the system offers the capability to create new English
groups for an educational center.

## Features

- **Pupil Testing**: Allows pupils to take English proficiency tests with multiple-choice questions.
- **True/False Responses**: Provides immediate feedback on whether answers are correct or incorrect.
- **Proficiency Levels**: Assigns proficiency levels such as Beginner, Elementary, Pre-Intermediate, Intermediate, Upper
  Intermediate, Advanced, and Proficient based on test performance.
- **Group Management**: Offers functionality to create and manage new English groups for educational centers.

## Database Schema

The system uses the following database tables:

- **students**: Stores information about pupils.
- **collections**: Contains collections of questions.
- **questions**: Stores questions linked to specific collections.
- **answers**: Contains true/false answers to the questions.
- **groups**: Manages English groups with details about the group and the teacher.

### Table Structures

- **students**
    - `id`: Big serial primary key
    - `phone_number`: Phone number of the student
    - `full_name`: Full name of the student
    - `created_at`: Timestamp of record creation
    - `updated_at`: Timestamp of last update

- **collections**
    - `id`: Big serial primary key
    - `name`: Name of the collection
    - `created_at`: Timestamp of record creation
    - `updated_at`: Timestamp of last update

- **questions**
    - `id`: Big serial primary key
    - `question_field`: Text of the question
    - `collection_id`: Foreign key referencing `collections(id)`
    - `created_at`: Timestamp of record creation
    - `updated_at`: Timestamp of last update

- **answers**
    - `id`: Big serial primary key
    - `is_true`: Boolean indicating if the answer is correct
    - `question_id`: Foreign key referencing `questions(id)`
    - `created_at`: Timestamp of record creation
    - `updated_at`: Timestamp of last update

- **groups**
    - `id`: Big serial primary key
    - `name`: Name of the group
    - `teacher_name`: Name of the teacher
    - `level`: Proficiency level (e.g., Beginner, Elementary, etc.)
    - `created_at`: Timestamp of record creation
    - `updated_at`: Timestamp of last update

## Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/golanguzb71/edu_check_graph.git

2. **Create tables**:
   ```bash
   make migrate-up

3. **Add mock infos**:
   ```bash
   make mock
4. **Drop tables**:
   ```bash
   make migrate-down
