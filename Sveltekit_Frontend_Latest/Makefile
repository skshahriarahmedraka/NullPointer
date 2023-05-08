dev :
	npm run dev
build:
	pnpm run build
preview:
	pnpm run preview
preview2:
	pnpm run build && pnpm run preview
test:
	npm run test

docker_test: docker-build_test 
	sudo docker compose up -d 
	# sudo docker compsoe exec -T http go test ./...
	sudo docker compose down 
docker_build:
	sudo docker build -t shahriarraka/null_pointer_frontend:latest .
	# docker build -t shahriarraka/null_pointer_frontend:latest .
docker_push:
	sudo docker push shahriarraka/null_pointer_frontend:latest 
docker_build_test:
	docker build . -t service_test --target=test

docker_run:
	sudo docker stop nullpointer_sveltekit || true && sudo docker rm nullpointer_sveltekit || true
	sudo docker run --publish 4173:4173 --name nullpointer_sveltekit shahriarraka/null_pointer_frontend:latest  