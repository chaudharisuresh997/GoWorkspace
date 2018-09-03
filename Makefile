PROJECTS=onebigproject/... mvctemplate/...

all:
	@cd src/ && go install $(PROJECTS)

test:
	@cd src/ && go test -coverprofile=coverage.out $(PROJECTS)
	@go tool cover -html=src/coverage.out

# generate XML report in Cobertura format
test.xml:
	@cd src/ && gocov test $(PROJECTS) | gocov-xml > coverage.xml

run:
	@./bin/cmd1 &
	@./bin/cmd2 &

stop:
	@pkill -f "./bin/cmd1" &
	@pkill -f "./bin/cmd2" &

clean:

clean-all: