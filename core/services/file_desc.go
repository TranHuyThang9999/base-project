package services

import (
	"context"
	customerrors "demo_time_sheet_server/core/custom_errors"
	"demo_time_sheet_server/core/domain"
)

type FileDescriptorsService struct {
	file domain.RepositoryFileDescriptors
}

func NewFileDescriptorsService(file domain.RepositoryFileDescriptors,
) *FileDescriptorsService {
	return &FileDescriptorsService{}
}

func (u *FileDescriptorsService) DeleteFileById(ctx context.Context, userID, fileID int64) *customerrors.CustomError {
	return nil
}
func (u *FileDescriptorsService) AddListFileByObjectID(ctx context.Context, object_id int64, urls []*string) *customerrors.CustomError {
	return nil
}
