package main

import (
	"fmt"
	"github.com/scottsbergeron/qforge"
)

func main() {
	server := qforge.PostgresDatabaseServer{
		Databases: []qforge.Database{
			{
				Name: "school_db",
				Schemas: []qforge.DatabaseSchema{
					{
						Name: "public",
						Tables: []qforge.Table{
							{
								Name: "students",
								Columns: []qforge.Column{
									{
										Id:       "student.id",
										Name:     "id",
										DataType: qforge.IntegerDataType,
									},
									{
										Id:       "student.first_name",
										Name:     "first_name",
										DataType: qforge.StringDataType,
									},
									{
										Id:       "student.last_name",
										Name:     "last_name",
										DataType: qforge.StringDataType,
									},
									{
										Id:       "student.birth_date",
										Name:     "birth_date",
										DataType: qforge.DateDataType,
									},
									{
										Id:       "student.grade_level",
										Name:     "grade_level",
										DataType: qforge.IntegerDataType,
									},
								},
							},
							{
								Name: "teachers",
								Columns: []qforge.Column{
									{
										Id:       "teacher.id",
										Name:     "id",
										DataType: qforge.IntegerDataType,
									},
									{
										Id:       "teacher.first_name",
										Name:     "first_name",
										DataType: qforge.StringDataType,
									},
									{
										Id:       "teacher.last_name",
										Name:     "last_name",
										DataType: qforge.StringDataType,
									},
									{
										Id:       "teacher.hire_date",
										Name:     "hire_date",
										DataType: qforge.DateDataType,
									},
								},
							},
						},
					},
					{
						Name: "school_data",
						Tables: []qforge.Table{
							{
								Name: "courses",
								Columns: []qforge.Column{
									{
										Id:       "course.id",
										Name:     "id",
										DataType: qforge.IntegerDataType,
									},
									{
										Id:       "course.name",
										Name:     "name",
										DataType: qforge.StringDataType,
									},
									{
										Id:       "teacher.id",
										Name:     "teacher_id",
										DataType: qforge.IntegerDataType,
									},
								},
							},
							{
								Name: "enrollments",
								Columns: []qforge.Column{
									{
										Id:       "enrollment.id",
										Name:     "id",
										DataType: qforge.IntegerDataType,
									},
									{
										Id:       "student.id",
										Name:     "student_id",
										DataType: qforge.IntegerDataType,
									},
									{
										Id:       "course.id",
										Name:     "course_id",
										DataType: qforge.IntegerDataType,
									},
									{
										Id:       "enrollment.date",
										Name:     "date",
										DataType: qforge.DateDataType,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	forge := qforge.QForge{
		Server: &server,
	}

	fmt.Println(forge.BuildQuery([]string{"course.id"}))
}
