.PHONY: docker templ
docker: templ
	docker buildx build --builder mybuilder --push --platform linux/amd64,linux/arm64 -t mitaka8/filehost:latest .
templ:
	templ generate
