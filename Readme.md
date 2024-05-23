Simple CLI tool that I intend to develop as an exercise

# Commands

* `add <task>` - adds task to listing 
* `ls` - displays all the tasks stored
* `del <id>` - deletes given task

# Reading documentation
- First, install godoc with: `go install golang.org/x/tools/cmd/godoc@latest`
- Then, in the project's root folder, run godoc -http=localhost:6060 
- All the documentation for Public methods will be available at `http://localhost:6060/pkg/clitest`