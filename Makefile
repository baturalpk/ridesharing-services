#
SERVICE=ridesharing-services

up:
	@echo ⚙️🔥 $(SERVICE) is being deployed...
	@echo "\n"
	docker compose up -d
	@echo ✅ Ready to go...

start:
	@echo 🚀 $(SERVICE) is being started...
	docker compose start

stop:
	@echo ✋ $(SERVICE) is being stopped...
	docker compose stop

remove:
	@echo 🗑 $(SERVICE) is being removed permanently...
	docker compose rm -fsv
