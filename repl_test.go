package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello     world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "hello, world!",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "hello , world !",
			expected: []string{"hello", ",", "world", "!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected slice of length %d, got: %d",
				len(c.expected), len(actual))

			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("mismatch at index %d: expected %s, got: %s",
					i, expectedWord, word)

				return
			}
		}
	}
}

func cmpCliCmds(c1 cliCmd, c2 cliCmd) bool {
	return c1.argsCount == c2.argsCount &&
		c1.description == c2.description &&
		c1.name == c2.name
}

func TestGetCliCommand(t *testing.T) {
	cmd0 := cliCmd{
		name:        "cmd0",
		description: "desc0",
		callback:    nil,
		argsCount:   0,
	}

	cmd1 := cliCmd{
		name:        "cmd1",
		description: "desc1",
		callback:    nil,
		argsCount:   1,
	}

	cmd2 := cliCmd{
		name:        "cmd2",
		description: "desc2",
		callback:    nil,
		argsCount:   2,
	}

	cmdsMap := cliCmdsMap{"cmd0": cmd0, "cmd1": cmd1, "cmd2": cmd2}

	cases := []struct {
		input []string

		expTrue   bool
		expCliCmd cliCmd
		expArgs   []string
	}{
		{
			input:     []string{"cmd0"},
			expTrue:   true,
			expCliCmd: cmd0,
			expArgs:   nil,
		},
		{
			input:     []string{"cmd1", "arg1"},
			expTrue:   true,
			expCliCmd: cmd1,
			expArgs:   []string{"arg1"},
		},
		{
			input:     []string{"cmd2", "arg1", "arg2"},
			expTrue:   true,
			expCliCmd: cmd2,
			expArgs:   []string{"arg1", "arg2"},
		},
		{
			input:     []string{"not", "a", "valid", "command"},
			expTrue:   false,
			expCliCmd: cliCmd{},
			expArgs:   nil,
		},
	}

	for _, c := range cases {
		actCliCmd, actArgs, actTrue := getCliCommand(c.input, &cmdsMap)

		if actTrue != c.expTrue {
			t.Errorf("mismatch at truthness of %s: expected %t, got: %t",
				c.expCliCmd.name, c.expTrue, actTrue)
			return
		}

		if !cmpCliCmds(actCliCmd, c.expCliCmd) {
			t.Errorf("mismatch at cli commands: expected %v, got: %v",
				c.expCliCmd, actCliCmd)
			return
		}

		if len(actArgs) != len(c.expArgs) {
			t.Errorf("expected %d args, got: %d", len(c.expArgs), len(actArgs))
			return
		}

		for i, arg := range actArgs {
			if arg != c.expArgs[i] {
				t.Errorf("arguments at index %d do not match: expected %s, got: %s",
					i, arg, c.expArgs[i])
				return
			}
		}
	}
}
