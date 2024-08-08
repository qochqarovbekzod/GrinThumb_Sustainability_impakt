 CURRENT_DIR=$(shell pwd)
DBURL := postgres://postgres:03212164@localhost:5432/sustainability_impact_service?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}


mig-up:
	migrate -path databases/migrations -database '${DBURL}' -verbose up

mig-down:
	migrate -path databases/migrations -database '${DBURL}' -verbose down

mig-force:
	migrate -path databases/migrations -database '${DBURL}' -verbose force 1

mig-create-type-enum:
	migrate create -ext sql -dir databases/migrations -seq create_type_enum

mig-create-impactlogs:
	migrate create -ext sql -dir databases/migrations -seq create_impact_logs_table

mig-create-sustainabilitychallenges:
	migrate create -ext sql -dir databases/migrations -seq create_user_sustainability_challenges_table

mig-create-userchallenges:
	migrate create -ext sql -dir databases/migrations -seq create_user_user_challenges_table