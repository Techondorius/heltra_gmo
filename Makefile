DOCKER=docker compose

run:
	$(DOCKER) down
	$(DOCKER) up --build -d
down:
	$(DOCKER) down
dev-run:
	$(DOCKER) down
	$(DOCKER) -f ./compose.dev.yml up --build -d
dev-rerun:
	$(DOCKER) -f ./compose.dev.yml restart -d