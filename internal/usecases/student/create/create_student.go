package create

import (
	"context"
	"errors"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/entities"
	"github.com/dmarins/student-api/internal/domain/repositories"
	"github.com/dmarins/student-api/internal/domain/usecases"
	"github.com/google/uuid"
)

type CreateStudentUseCase struct {
	StudentRepository repositories.IStudentRepository
}

func NewCreateStudentUseCase(studentRepository repositories.IStudentRepository) usecases.ICreateStudentUseCase {
	return &CreateStudentUseCase{
		StudentRepository: studentRepository,
	}
}

func (uc *CreateStudentUseCase) Execute(ctx context.Context, student entities.Student) (*dtos.StudentOutput, error) {

	exists, err := uc.StudentRepository.ExistsByName(ctx, student.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("student already exists")
	}

	newStudent := entities.Student{
		ID:   uuid.New().String(),
		Name: student.Name,
	}

	err = uc.StudentRepository.Save(ctx, &newStudent)
	if err != nil {
		return nil, err
	}

	output := &dtos.StudentOutput{
		ID:   newStudent.ID,
		Name: newStudent.Name,
	}

	return output, nil
}
