package service

import (
	"context"
	"database/sql"
	"errors"
	"sort"
	pb "sustainability-service/generated/sustainability"
	"sustainability-service/storage/postgres"
)

type SustainabilityService struct {
	pb.UnimplementedSustainabilityimpactServiceServer

	Sustainability *postgres.SustainabilityRepo
}

func (s *SustainabilityService) LogImpact(ctx context.Context, in *pb.LogImpactRequest) (*pb.LogImpactResponse, error) {
	resp, err := s.Sustainability.LogImpact(in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SustainabilityService) GetUserImpact(ctx context.Context, in *pb.GetUserImpactRequest) (*pb.GetUserImpactResponse, error) {
	impacts, err := s.Sustainability.GetUserImpact(in.UserId)
	if err != nil {
		return nil, err
	}

	return impacts, nil
}

func (s *SustainabilityService) GetCommunityImpact(ctx context.Context, in *pb.GetCommunityImpactRequest) (*pb.GetCommunityImpactResponse, error) {

	var commmunityImpact []*pb.CommunityImpact

	for _, member := range in.Members {
		resp, err := s.Sustainability.GetCommunityImpact(member)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		commmunityImpact=append(commmunityImpact, resp)
	}

	return &pb.GetCommunityImpactResponse{CommunityImpacts: commmunityImpact}, nil
}

func (s *SustainabilityService) GetChallenges(ctx context.Context, in *pb.GetChallengesRequest) (*pb.GetChallengesResponse, error) {
	challenges, err := s.Sustainability.GetChallenges()
	if err != nil {
		return nil, err
	}

	return challenges, nil
}

func (s *SustainabilityService) JoinChallenge(ctx context.Context, in *pb.JoinChallengeRequest) (*pb.JoinChallengeResponse, error) {
	resp, err := s.Sustainability.JoinChallenge(in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SustainabilityService) UpdateChallengeProgress(ctx context.Context, in *pb.UpdateChallengeProgressRequest) (*pb.UpdateChallengeProgressResponse, error) {

	resp, err := s.Sustainability.UpdateChallengeProgress(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SustainabilityService) GetUserChallenges(ctx context.Context, in *pb.GetUserChallengesRequest) (*pb.GetUserChallengesResponse, error) {
	userChallenges, err := s.Sustainability.GetUsersChallenges(in.UserId)

	if err != nil {
		return nil, err
	}

	return userChallenges, err
}

func (s *SustainabilityService) GetUsersLeaderboard(ctx context.Context, in *pb.GetUsersLeaderboardRequest) (*pb.GetUsersLeaderboardResponse, error) {
	leaderboard, err := s.Sustainability.GetUsersLeaderboard()

	if err != nil {
		return nil, err
	}

	return &pb.GetUsersLeaderboardResponse{Leaderboard: leaderboard}, nil
}

func (s *SustainabilityService) GetCommunitiesLeaderboard(ctx context.Context, in *pb.GetCommunitiesLeaderboardRequest) (*pb.GetCommunitiesLeaderboardResponse, error) {

	var ledaderboard []*pb.CommunitiesLeaderboard

	for _, community := range in.CommunityMembers {
		var totalProgres float32
		for _, member := range community.Members {
			progres, err := s.Sustainability.GetUsersCommonProgres(member)
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}

			totalProgres += progres
		}

		ledaderboard = append(ledaderboard, &pb.CommunitiesLeaderboard{
			CommunityName: community.CommunityName,
			CommunityId:   community.CommunityId,
			TotalProgres:  totalProgres,
		})
	}

	sort.Slice(ledaderboard, func(i, j int) bool {
		return ledaderboard[i].TotalProgres > ledaderboard[j].TotalProgres
	})

	return &pb.GetCommunitiesLeaderboardResponse{CommunitiesLeaderboards: ledaderboard[:10]}, nil
}

func (s *SustainabilityService) CreateSustainability(ctx context.Context, in *pb.CreateSustainabilityRequest) (*pb.CreateSustainabilityResponse, error) {
	return s.Sustainability.CreateSustainability(in)
}
