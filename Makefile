SRC_YAML="swagger.yml"

update-vendor:
	glide update

run: terraform-server
	./terraform-server --scheme=http

validate-swagger:
	swagger validate $(SRC_YAML)

terraform-server: validate-swagger
	swagger generate server \
		--target=server \
		--principal=models.Principal \
		--name=TerraformServer \
		--spec=$(SRC_YAML) && \
	go build \
		-o ./terraform-server \
		github.com/zeebox/terraform-server/server/cmd/terraform-server-server
