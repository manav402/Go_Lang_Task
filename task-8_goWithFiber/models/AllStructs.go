package models

import (
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// a model for students which has department
type Student struct {
	Name           string             `json:"name" bson:"name"`
	DOB            primitive.DateTime `json:"dob" bson:"dob"`
	DepartmentId   primitive.ObjectID `json:"departmentid" bson:"departmentid"`
	DepartmentName string             `json:"departmentname" bson:"departmentname"`
	Enrollment     string             `json:"enrollment" bson:"enrollment"`
}

// a model for department which will have list of students
type Department struct {
	DepartmentId primitive.ObjectID `json:"departmentid" bson:"departmentid"`
	Name         string             `json:"name" bson:"name"`
	Intake       int                `json:"intake" bson:"intake"`
	Hod          string             `json:"hod" bson:"hod"`
	Subjects     primitive.A        `json:"subjects" bson:"subjects"`
}

// a model for storing subject id
type Subjects struct {
	Subjectid    primitive.ObjectID `json:"subjectid" bson:"subjectid"`
	Name         string             `json:"name" bson:"name" validate:"required"`
	Departmentid primitive.ObjectID `json:"departmentid" bson:"departmentid"`
}

// var StudentSchema = bson.M{
	
// }

// var DepartmentSchema = bson.M{

// }

// var SubjectSchema = bson.M{

// }
