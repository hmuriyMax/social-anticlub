package main

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/date"
	"math/rand"
	"socialanticlub/internal/helpers"
	"socialanticlub/internal/pb/user_service"
	"strings"
	"time"
)

func ParseToRequest(row []string) (*user_service.RegRequest, error) {
	if len(row) != 3 {
		return nil, fmt.Errorf("invalid row length: expected %d, but got %d", 3, len(row))
	}

	fullName := row[0]
	birthDate := row[1]
	homeTown := row[2]

	names := strings.Split(fullName, " ")
	if len(names) != 2 {
		return nil, fmt.Errorf("invalid name token length: expected %d, but got %d", 2, len(names))
	}

	sName := names[0]
	name := names[1]

	bday, err := time.Parse("2006-01-02", birthDate)
	if err != nil {
		return nil, err
	}

	genders := helpers.Keys(user_service.UserInfo_Gender_name)
	userNameAsUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &user_service.RegRequest{
		Pass: "qwerty",
		Info: &user_service.UserInfo{
			Name:    name,
			Surname: &sName,
			Birthday: &date.Date{
				Year:  int32(bday.Year()),
				Month: int32(bday.Month()),
				Day:   int32(bday.Day()),
			},
			Gender:   helpers.Ptr(user_service.UserInfo_Gender(genders[rand.Intn(len(genders))])),
			Hometown: &homeTown,
			Nickname: string([]rune(userNameAsUUID.String())[:15]),
		},
	}, nil
}
