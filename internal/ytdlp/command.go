package ytdlp

// import "github.com/go-cmd/cmd"
//
// type StreamingCommand struct {
// 	Binary string
// 	Args   []string
// 	Output func(string)
// 	Error  func(string)
// 	Errors []error
// 	Status cmd.Status
// }
//
// func (c *StreamingCommand) Run() error {
// 	cmdOptions := cmd.Options{
// 		Buffered:  false,
// 		Streaming: true,
// 	}
//
// 	// Create Cmd with options
// 	envCmd := cmd.NewCmdOptions(cmdOptions, c.Binary, c.Args...)
//
// 	// Print STDOUT and STDERR lines streaming from Cmd
// 	doneChan := make(chan struct{})
// 	go func() {
// 		defer close(doneChan)
// 		// Done when both channels have been closed
// 		// https://dave.cheney.net/2013/04/30/curious-channels
// 		for envCmd.Stdout != nil || envCmd.Stderr != nil {
// 			select {
// 			case line, open := <-envCmd.Stdout:
// 				if !open {
// 					envCmd.Stdout = nil
// 					continue
// 				}
// 				c.Output(line)
// 			case line, open := <-envCmd.Stderr:
// 				if !open {
// 					envCmd.Stderr = nil
// 					continue
// 				}
// 				c.Error(line)
// 			}
// 		}
// 	}()
//
// 	// Run and wait for Cmd to return, discard Status
// 	status := <-envCmd.Start()
//
// 	// Wait for goroutine to print everything
// 	<-doneChan
//
// 	c.Status = status
// 	return nil
// }
//
// func SimpleCommand(binary string, args ...string) ([]string, []string, cmd.Status, error) {
// 	stdout := []string{}
// 	stderr := []string{}
//
// 	stdoutFunc := func(line string) {
// 		stdout = append(stdout, line)
// 	}
// 	stderrFunc := func(line string) {
// 		stderr = append(stderr, line)
// 	}
//
// 	c := &StreamingCommand{
// 		Binary: binary,
// 		Args:   args,
// 		Output: stdoutFunc,
// 		Error:  stderrFunc,
// 	}
//
// 	err := c.Run()
// 	return stdout, stderr, c.Status, err
// }
