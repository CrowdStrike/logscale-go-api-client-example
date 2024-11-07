SCHEMA_CLUSTER?=https://cloud.community.humio.com
SCHEMA_CLUSTER_API_TOKEN?=${TOKEN}

generate:
	go generate

update-schema:
ifndef TOKEN
$(error Environment variable TOKEN must be set.)
endif
	go run github.com/suessflorian/gqlfetch/gqlfetch@607d6757018016bba0ba7fd1cb9fed6aefa853b5 --endpoint ${SCHEMA_CLUSTER}/graphql --header "Authorization=Bearer ${SCHEMA_CLUSTER_API_TOKEN}" > schema/schema.graphql
	printf "# Fetched from version %s" $$(curl --location '${SCHEMA_CLUSTER}/api/v1/status' | jq -r ".version") >> schema/schema.graphql
