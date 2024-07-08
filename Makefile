

.PHONY: setup-tools c-proto commit evans

setup-tools:
	@if ! command -v buf &> /dev/null; then brew install bufbuild/buf/buf; fi
	
c-proto:
	chmod +x ./scripts/create_proto.sh
	./scripts/create_proto.sh ${DIR}

commit:
	npx git-cz

lint:
	buf dep update
	buf lint

push: lint
	buf push


