package server

import (
	"fmt"

	"github.com/go-cmd/cmd"
)

type StreamingCommand struct {
	Binary string
	Args   []string
	Output func(string)
	Error  func(string)
	Errors []error
	Status cmd.Status
}

func (c *StreamingCommand) Run() error {
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options
	envCmd := cmd.NewCmdOptions(cmdOptions, c.Binary, c.Args...)

	// Print STDOUT and STDERR lines streaming from Cmd
	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		// Done when both channels have been closed
		// https://dave.cheney.net/2013/04/30/curious-channels
		for envCmd.Stdout != nil || envCmd.Stderr != nil {
			select {
			case line, open := <-envCmd.Stdout:
				if !open {
					envCmd.Stdout = nil
					continue
				}
				c.Output(line)
			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}
				c.Error(line)
			}
		}
	}()

	// Run and wait for Cmd to return, discard Status
	status := <-envCmd.Start()

	// Wait for goroutine to print everything
	<-doneChan

	c.Status = status
	return nil
}

func (c *Workers) streamingCommandJob(url string) JobFunc {
	return func() error {
		s := &StreamingCommand{
			Binary: "yt-dlp",
			Args:   []string{url},
		}
		s.Output = func(line string) {
		}
		s.Error = func(line string) {
		}

		err := s.Run()
		if err != nil {
			return err
		}

		if s.Status.Exit != 0 {
			return fmt.Errorf("yt-dlp: exit code: %d", s.Status.Exit)
		}
		if len(s.Errors) > 0 {
			return fmt.Errorf("yt-dlp: %s", s.Errors)
		}

		return nil
	}
}

// CommandJob runs a command and wraps the output in a job
func (c *Workers) CommandJob(binary string, args ...string) JobFunc {
	return func() error {
		// see: https://github.com/go-cmd/cmd/blob/master/examples/blocking-streaming/main.go
		// Disable output buffering, enable streaming
		cmdOptions := cmd.Options{
			Buffered:  false,
			Streaming: true,
		}

		// Create Cmd with options
		envCmd := cmd.NewCmdOptions(cmdOptions, binary, args...)

		// Print STDOUT and STDERR lines streaming from Cmd
		doneChan := make(chan struct{})
		go func() {
			defer close(doneChan)
			// Done when both channels have been closed
			// https://dave.cheney.net/2013/04/30/curious-channels
			for envCmd.Stdout != nil || envCmd.Stderr != nil {
				select {
				case line, open := <-envCmd.Stdout:
					if !open {
						envCmd.Stdout = nil
						continue
					}
					c.ProcessLine("CommandJob", line)
				case line, open := <-envCmd.Stderr:
					if !open {
						envCmd.Stderr = nil
						continue
					}
					c.ProcessError("CommandJob", line)
				}
			}
		}()

		// Run and wait for Cmd to return, discard Status
		status := <-envCmd.Start()

		// Wait for goroutine to print everything
		<-doneChan

		c.ProcessLine("CommandJob", fmt.Sprintf("status: %+v", status))
		return nil
	}
}

func (c *Workers) ProcessLine(name, line string) {
	// Do something with the line here.
	c.logger.Warnf("[%s] %s", name, line)
}

func (c *Workers) ProcessError(name, line string) {
	// Do something with the line here.
	c.logger.Errorf("[%s] %s", name, line)
}

// https://www.dolthub.com/blog/2022-11-28-go-os-exec-patterns/
/*
	cmd := exec.Command("ls", "/usr/local/bin")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	err = cmd.Start()
	if err != nil {
		return err
	}
	for scanner.Scan() {
		// Do something with the line here.
		ProcessLine(scanner.Text())
	}
	if scanner.Err() != nil {
		cmd.Process.Kill()
		cmd.Wait()
		return scanner.Err()
	}
	return cmd.Wait()
*/
