package postgres

import (
	"fmt"
	"reflect"
	pb "sustainability-service/generated/sustainability"
	"testing"
)

func TestLogImpact(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)

	reslog, err := sus.LogImpact(&pb.LogImpactRequest{
		UserId:   "a85370e6-2259-4067-b072-4f457033285e",
		Category: "water_saved",
		Unit:     "ANY",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitlog := pb.LogImpactResponse{
		Success: true,
	}
	if !reflect.DeepEqual(reslog, &waitlog) {
		t.Errorf("have %v , wont %v", reslog, &waitlog)
	}
}

func TestGetUserImpact(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)
	resget, err := sus.GetUserImpact("a85370e6-2259-4067-b072-4f457033285e")
	if err != nil {
		fmt.Println(err)
	}
	waitget := pb.GetUserImpactResponse{
		UserImpact: []*pb.UserImpact{
			{
				UserId:   "a85370e6-2259-4067-b072-4f457033285e",
				Category: "water_saved",
				Amount:   "12.00",
				Unit:     "ANY",
			},
			{
				UserId:   "a85370e6-2259-4067-b072-4f457033285e",
				Category: "water_saved",
				Amount:   "12.00",
				Unit:     "ANY",
			},
		},
	}
	if !reflect.DeepEqual(resget, &waitget) {
		t.Errorf("have %v , wont %v", resget, &waitget)
	}
}

func TestCreateSustainability(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)
	rescreate, err := sus.CreateSustainability(&pb.CreateSustainabilityRequest{
		Title:       "ANY",
		Description: "ANY",
		GoalAmount:  12.00,
		GoalUnit:    "ANY",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitcreate := pb.CreateSustainabilityResponse{
		Success: true,
	}
	if !reflect.DeepEqual(rescreate, &waitcreate) {
		t.Errorf("have %v , wont %v", rescreate, &waitcreate)
	}
}

func TestJoinChallenge(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)
	resjoin, err := sus.JoinChallenge(&pb.JoinChallengeRequest{
		UserId:      "a85370e6-2259-4067-b072-4f457033285e",
		ChallengeId: "63ba7b6f-1f3c-4095-982a-da90d0961a17",
		Progress:    0,
	})
	if err != nil {
		fmt.Println(err)
	}
	waitjoin := pb.JoinChallengeResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resjoin, &waitjoin) {
		t.Errorf("have %v , wont %v", resjoin, &waitjoin)
	}
}

func TestUpdateChallengeProgress(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)
	resupdate, err := sus.UpdateChallengeProgress(&pb.UpdateChallengeProgressRequest{
		UserId:      "a85370e6-2259-4067-b072-4f457033285e",
		ChallengeId: "63ba7b6f-1f3c-4095-982a-da90d0961a17",
		Progress:    12.00,
	})
	if err != nil {
		fmt.Println(err)
	}
	waitupdate := pb.UpdateChallengeProgressResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resupdate, &waitupdate) {
		t.Errorf("have %v , wont %v", resupdate, &waitupdate)
	}
}

func TestGetCommunityImpact(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)
	resget, err := sus.GetCommunityImpact("a85370e6-2259-4067-b072-4f457033285e")
	if err != nil {
		fmt.Println(err)
	}
	waitget := []*pb.CommunityImpact{
		{
			Id: "8af45dac-05ea-4ac8-aaa2-ce86afac3be5",
			UserId:     "a85370e6-2259-4067-b072-4f457033285e",
			Category:   "water_saved",
			GoalAmount: 12.00,
			GoalUnit:   "ANY",
			LoggedAt:   "2024-07-04T19:42:55.923531+05:00",
		},
	}
	if !reflect.DeepEqual(resget, &waitget) {
		t.Errorf("have %v , wont %v", resget, &waitget)
	}
}

func TestGetChallenges(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)
	resget, err := sus.GetChallenges()
	if err != nil {
		fmt.Println(err)
	}
	waitget := pb.GetChallengesResponse{
		Challanges: []*pb.Challenge{
			{
				Id:          "63ba7b6f-1f3c-4095-982a-da90d0961a17",
				Title:       "ANY",
				Description: "ANY",
				GoalAmount:  12.00,
				GoalUnit:    "ANY",
			},
		},
	}
	if !reflect.DeepEqual(resget, &waitget) {
		t.Errorf("have %v , wont %v", resget, &waitget)
	}
}

func TestGetUsersChallenges(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)
	resget, err := sus.GetUsersChallenges("a85370e6-2259-4067-b072-4f457033285e")
	if err != nil {
		fmt.Println(err)
	}
	waitget := pb.GetUserChallengesResponse{
		UserChallenges: []*pb.UserChallenge{
			{
				UserId:      "a85370e6-2259-4067-b072-4f457033285e",
				ChallengeId: "63ba7b6f-1f3c-4095-982a-da90d0961a17",
				Progress:    12.00,
			},
		},
	}
	if !reflect.DeepEqual(resget, &waitget) {
		t.Errorf("have %v , wont %v", resget, &waitget)
	}
}

func TestGetUsersLeaderboard(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)
	resget, err := sus.GetUsersLeaderboard()
	if err != nil {
		fmt.Println(err)
	}
	waitget := []*pb.UsersLeaderboard{
		{
			UserId:       "a85370e6-2259-4067-b072-4f457033285e",
			TotalProgres: 12.00,
		},
	}
	if !reflect.DeepEqual(resget, waitget) {
		t.Errorf("have %v , wont %v", resget, waitget)
	}
}

func TestGetUsersCommonProgres(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sus := NewSustainabilityRepo(db)
	resget, err := sus.GetUsersCommonProgres("a85370e6-2259-4067-b072-4f457033285e")
	if err != nil {
		fmt.Println(err)
	}
	var waitget float32
	waitget = 12.00
	if !reflect.DeepEqual(resget, waitget) {
		t.Errorf("have %v , wont %v", resget, waitget)
	}
}
