# SQL TABLES

## Facilities Manage table

| Field       | Type             | Null | Key | Default | Extra          |
|-------------|------------------|------|-----|---------|----------------|
| id          | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at  | datetime         | YES  |     | NULL    |                |
| updated_at  | datetime         | YES  |     | NULL    |                |
| deleted_at  | datetime         | YES  | MUL | NULL    |                |
| room_id     | int(11)          | NO   |     | NULL    |                |
| facility_id | int(11)          | NO   |     | NULL    |                |
| quantity    | int(11)          | NO   |     | 0       |                |
| default     | int(11)          | NO   |     | 0       |                |
| logs        | text             | YES  |     | NULL    |                |

## Facility table

| Field       | Type             | Null | Key | Default | Extra          |
|-------------|------------------|------|-----|---------|----------------|
| id          | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at  | datetime         | YES  |     | NULL    |                |
| updated_at  | datetime         | YES  |     | NULL    |                |
| deleted_at  | datetime         | YES  | MUL | NULL    |                |
| name        | varchar(255)     | NO   | UNI | NULL    |                |
| description | text             | YES  |     | NULL    |                |

## Money Manage table

| Field       | Type             | Null | Key | Default | Extra          |
|-------------|------------------|------|-----|---------|----------------|
| id          | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at  | datetime         | YES  |     | NULL    |                |
| updated_at  | datetime         | YES  |     | NULL    |                |
| deleted_at  | datetime         | YES  | MUL | NULL    |                |
| student_id  | int(11)          | NO   | MUL | NULL    |                |
| month       | int(11)          | NO   |     | NULL    |                |
| year        | int(11)          | NO   |     | NULL    |                |
| status      | text             | NO   |     | NULL    |                |
| description | text             | YES  |     | NULL    |                |
| money       | int(11)          | NO   |     | NULL    |                |

## Notification table

| Field      | Type             | Null | Key | Default | Extra          |
|------------|------------------|------|-----|---------|----------------|
| id         | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | datetime         | YES  |     | NULL    |                |
| updated_at | datetime         | YES  |     | NULL    |                |
| deleted_at | datetime         | YES  | MUL | NULL    |                |
| title      | text             | NO   |     | NULL    |                |
| content    | text             | NO   |     | NULL    |                |

## Request table

| Field      | Type             | Null | Key | Default | Extra          |
|------------|------------------|------|-----|---------|----------------|
| id         | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | datetime         | YES  |     | NULL    |                |
| updated_at | datetime         | YES  |     | NULL    |                |
| deleted_at | datetime         | YES  | MUL | NULL    |                |
| status     | text             | NO   |     | NULL    |                |
| student_id | int(11)          | NO   |     | NULL    |                |
| title      | text             | NO   |     | NULL    |                |
| message    | text             | NO   |     | NULL    |                |
| reply      | text             | YES  |     | NULL    |                |
| note       | text             | YES  |     | NULL    |                |

## Room table

| Field       | Type     | Null | Key | Default | Extra |
|-------------|----------|------|-----|---------|-------|
| created_at  | datetime | YES  |     | NULL    |       |
| updated_at  | datetime | YES  |     | NULL    |       |
| deleted_at  | datetime | YES  |     | NULL    |       |
| room_id     | int(11)  | NO   | PRI | NULL    |       |
| price       | int(11)  | NO   |     | 0       |       |
| occupied    | int(11)  | NO   |     | 0       |       |
| max         | int(11)  | NO   |     | 0       |       |
| description | text     | YES  |     | NULL    |       |

## Student table

| Field      | Type     | Null | Key | Default | Extra |
|------------|----------|------|-----|---------|-------|
| created_at | datetime | YES  |     | NULL    |       |
| updated_at | datetime | YES  |     | NULL    |       |
| deleted_at | datetime | YES  |     | NULL    |       |
| student_id | int(11)  | NO   | PRI | NULL    |       |
| name       | text     | NO   |     | NULL    |       |
| dob        | text     | YES  |     | NULL    |       |
| contact    | text     | NO   |     | NULL    |       |
| address    | text     | YES  |     | NULL    |       |
| room_id    | int(11)  | NO   |     | NULL    |       |
| priority   | int(11)  | YES  |     | 0       |       |

## User table

| Field      | Type         | Null | Key | Default | Extra |
|------------|--------------|------|-----|---------|-------|
| created_at | datetime     | YES  |     | NULL    |       |
| updated_at | datetime     | YES  |     | NULL    |       |
| deleted_at | datetime     | YES  |     | NULL    |       |
| user_id    | int(11)      | NO   | PRI | 0       |       |
| role       | int(11)      | NO   |     | 0       |       |
| username   | varchar(255) | NO   | UNI | NULL    |       |
| password   | text         | NO   |     | NULL    |       |
