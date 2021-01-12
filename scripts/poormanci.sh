#!/bin/bash

# I am poor and don't have time to invest in CI.
#
# Just `source` this file in your `.bashrc` to verify you code before `git push`.

# Wraps your git command and runs checks before pushing to remote.
function gitWrapper {
	# Behave as usual if not inside the `gitminer repository`.
	if ! grep -q "gitminer.go" <(ls); then
		git "$@"
		return $?
	fi

	echo "gitminer repository detected."

	if [ $1 == "push" ]; then
		echo "Executing checks before pushing..."
		go vet && go test ./... || return 1
		echo "All checks done, pushing..."
	fi
	git "$@"
}

alias git=gitWrapper
