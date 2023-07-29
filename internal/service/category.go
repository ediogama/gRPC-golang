package service

import (
	"context"

	"github.com/ediogama/gRPC-golang/internal/database"
	"github.com/ediogama/gRPC-golang/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, request *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.CreateCategory(request.Name, request.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}
