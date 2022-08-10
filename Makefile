DOCKER=docker compose

run:
	$(DOCKER) up --build -d
down:
	$(DOCKER) down
dev-run:
	$(DOCKER) -f ./compose.dev.yml up --build -d
dev-rerun:
	$(DOCKER) -f ./compose.dev.yml restart -d