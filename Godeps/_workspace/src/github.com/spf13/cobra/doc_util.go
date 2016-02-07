// Copyright 2015 Red Hat Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cobra

// Test to see if we have a reason to print See Also information in docs
// Basically this is a test for a parent commend or a subcommand which is
// both not deprecated and not the autogenerated help command.
func (cmd *Command) hasSeeAlso() bool {
	if cmd.HasParent() {
		return true
	}
	children := cmd.Commands()
	if len(children) == 0 {
		return false
	}
	for _, c := range children {
		if len(c.Deprecated) != 0 || c == cmd.helpCommand {
			continue
		}
		return true
	}
	return false
}
