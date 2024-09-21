package v1

import "context"
import "crawl-microservice/internal/models"
import "crawl-microservice/internal/dto"

func toRenewResponse(
	ctx context.Context,
	tokens *models.Tokens,
) *dto.RenewResponse {
	return &dto.RenewResponse{
		Tokens: &dto.Tokens{
			Access:  tokens.Access,
			Refresh: tokens.Refresh,
		},
	}
}
func toLoginResponse(
	ctx context.Context,
	emp *models.User,
	tokens *models.Tokens,
) *dto.LoginResponse {
	return &dto.LoginResponse{
		FirstName: emp.FirstName,
		LastName:  emp.LastName,
		Tokens: &dto.Tokens{
			Access:  tokens.Access,
			Refresh: tokens.Refresh,
		},
	}
}
