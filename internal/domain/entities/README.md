# Entities

This directory contains the domain entities. These are the objects that are used to represent the domain model. They are the core of the domain model and are used to represent the state of the application.

## Guidelines

- Entities should be simple objects that contain only data and methods that operate on the data.
- Entities should not contain any business logic.
- Entities should be immutable.
- Entities should be serializable.
- Entities should be simple and small.
- Entities should be named after the domain concept they represent.

## Explanation

```go
type Student struct {
	User
	Name       string  `json:"name"`
	NISN       string  `json:"nisn"`
	Address    string  `json:"address"`
	Birthplace string  `json:"birthplace"`
	Birthdate  string  `json:"birthdate"`
	ParentID   *string `json:"parent_id"`
	Parent     *Parent `json:"parent"`
}
```

`*string` and `*Parent` are used to represent the optional field (nullable). If the field is not required, it should be a pointer.

```go
type Subject struct {
	ID          string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Semester    string    `json:"semester"`
	Teachers    []Teacher `gorm:"many2many:teacher_subjects;"`
}

type TeacherSubject struct {
	TeacherID string `gorm:"primaryKey"`
	SubjectID string `gorm:"primaryKey"`
	Teacher   Teacher
	Subject   Subject
}
```

`[]Teacher` and `TeacherSubject` are used to represent the many-to-many relationship. `TeacherSubject` is used to represent the join table.
