build:
	docker buildx build --platform=linux/amd64,linux/arm64 . -t whzy1990/http-client:latest --push