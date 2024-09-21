package repo

import "context"

type IViolationRepo interface {
	CreateViolation(ctx context.Context,Id string) (string,error)
}