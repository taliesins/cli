/*
Package cgapp includes a powerful CLI for the Create Go App project.

Create a new production-ready project with backend (Golang),
frontend (JavaScript, TypeScript) and deploy automation
(Ansible, Docker) by running one CLI command.

-> Focus on writing code and thinking of business logic!
<- The Create Go App CLI will take care of the rest.

A helpful documentation and next steps -> https://create-go.app/

Copyright © 2019-present Vic Shóstak <truewebartisans@gmail.com> (https://1wa.co)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cgapp

import (
	"bufio"
	"bytes"
	"os/exec"
)

// ExecCommand function to execute a given command.
func ExecCommand(command string, options []string) error {
	//
	if command == "" {
		return throwError("No command to execute!")
	}

	// Create buffer for stderr.
	stderr := &bytes.Buffer{}

	// Collect command line
	cmd := exec.Command(command, options...) // #nosec G204

	// Set buffer for stderr from cmd
	cmd.Stderr = stderr

	// Create a new reader
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return throwError(err.Error())
	}

	// Start executing command.
	if err := cmd.Start(); err != nil {
		return throwError(stderr.String())
	}

	// Create a new scanner and run goroutine func with output.
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			SendMsg(false, "*", scanner.Text(), "cyan", false)
		}
	}()

	// Wait for executing command.
	if err := cmd.Wait(); err != nil {
		return throwError(stderr.String())
	}

	return nil
}
