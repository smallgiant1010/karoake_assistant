package mapper

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"karoake_assistant/backend/internal/domains"
	"karoake_assistant/backend/internal/data/sqlc"
)

func UserModelToDomain(model *sqlc.User) *domains.User {
	return domains.NewUser(
		int32(model.Userid),
		model.Username,
		model.Password,
		model.Generatecount.Int32,
	)
}

func UserDomainToModel(domain *domains.User) *sqlc.User {
	var generateCount pgtype.Int4
	if err := generateCount.Scan(domain.GenerateCount); err != nil {
		fmt.Printf("error occured casting generateCount to pgtype: %v", err)
		return nil
	}
	return &sqlc.User{
		Userid: int64(domain.UserID),
		Username: domain.Username,
		Password: domain.Password,
		Generatecount: generateCount,
	}
}
