package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/blaskovicz/go-cryptkeeper"
	"github.com/spf13/pflag"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	op, k, v, err := getFlags()
	if err != nil {
		return err
	}

	if err := cryptkeeper.SetCryptKey([]byte(k)); err != nil {
		return err
	}

	var result string

	switch op {
	case "encrypt":
		result, err = cryptkeeper.Encrypt(v)
	case "decrypt":
		result, err = cryptkeeper.Decrypt(v)
	default:
		return usageError(fmt.Sprintf("Invalid operation: %s", op))
	}

	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func getFlags() (string, string, string, error) {
	o := pflag.String("o", "", "operation: encrypt or decrypt")
	k := pflag.String("k", "", "encryption key")
	v := pflag.String("v", "", "target value of the operation")

	var errs []string
	if empty(o) {
		errs = append(errs, "operation flag missing")
	}
	if empty(k) {
		errs = append(errs, "key flag missing")
	}
	if empty(v) {
		errs = append(errs, "value flag missing")
	}
	if len(errs) > 0 {
		return "", "", "", usageError(strings.Join(errs, "\n"))
	}

	return *o, *k, *v, nil
}

func empty(s *string) bool {
	if s == nil || *s == "" {
		return true
	}

	return false
}

type usageError string

func (ue usageError) Error() string {
	return fmt.Sprintf("%s\n\nUsage of %s:\n%s", string(ue), os.Args[0], pflag.CommandLine.FlagUsages())
}
