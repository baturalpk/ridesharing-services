# ridesharing-services Makefile

SERVICE=ridesharing-services

up:
	@echo ⚙️🔥 $(SERVICE) is being deployed...
	@echo "\n"
	docker compose up -d
	@echo ✅ Ready to go...

down:
	@echo ✋ $(SERVICE) is being stopped and destroyed...
	docker compose down -v --rmi local

.PHONY: up, down
