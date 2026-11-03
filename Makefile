PROGRAM_NAME = "gluster.git"
PID_FILE = "sample.pid"
LOG_FILE = "sample.log"

$(PROGRAM_NAME): cmd/main.go
	go build -o $(PROGRAM_NAME) cmd/main.go

clean:
	$(RM) $(PROGRAM_NAME)
	$(RM) $(LOG_FILE)
	$(RM) $(PID_FILE)
