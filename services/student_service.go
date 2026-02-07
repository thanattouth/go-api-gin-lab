package services

import (
	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	// (Optional) คุณสามารถเพิ่ม Logic ตรวจสอบข้อมูลตรงนี้ได้ เช่น GPA ห้ามติดลบ
	return s.Repo.Create(student)
}

func (s *StudentService) UpdateStudent(id string, student models.Student) error {
	// ตรงนี้อาจจะเพิ่ม Logic เช็คว่า id นี้มีจริงไหมก่อนก็ได้
	return s.Repo.Update(id, student)
}
