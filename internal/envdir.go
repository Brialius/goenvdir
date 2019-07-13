package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// EnvDir runs another program with environment modified according to files in a specified directory
func EnvDir(dir string, child []string) (int, error) {
	envVars, err := getEnvVarsFromDir(dir)
	if err != nil {
		return -1, err
	}
	exitCode, err := execChild(child, mergeEnvVars(envVars, os.Environ()))
	return exitCode, err
}

func mergeEnvVars(envVars map[string]string, environ []string) []string {
	res := make([]string, 0, len(envVars)+len(environ))
	for _, osVariable := range environ {
		// Ignore windows cmd legacy variables
		// https://devblogs.microsoft.com/oldnewthing/?p=14133
		if strings.HasPrefix(osVariable, "=") {
			continue
		}
		v := strings.Split(osVariable, "=")
		if len(v) != 2 {
			log.Fatalf("Unexpected env variable format: `%s`", osVariable)
		}
		key, value := v[0], v[1]
		if _, ok := envVars[key]; !ok {
			res = append(res, key+"="+value)
		}
	}
	for key, value := range envVars {
		if value != "" {
			res = append(res, key+"="+value)
		}
	}
	return res
}

func getEnvVarsFromDir(dir string) (map[string]string, error) {
	res := map[string]string{}
	if stat, err := os.Stat(dir); err != nil {
		return nil, fmt.Errorf("can't open `%s`: %s", dir, err)
	} else if !stat.IsDir() {
		return nil, fmt.Errorf("`%s` is not a dir", dir)
	}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.Contains(info.Name(), "=") {
			f, err := os.Open(path)
			defer func() {
				if err := f.Close(); err != nil {
					log.Fatal(err)
				}
			}()
			if err != nil {
				return err
			}

			// Read first line only
			r := bufio.NewReader(f)
			buf, _, er := r.ReadLine()
			if er != nil && er != io.EOF {
				return er
			}
			res[info.Name()] = string(buf)
		}
		return nil
	})
	return res, err
}

func execChild(child []string, envVars []string) (int, error) {
	c := exec.Cmd{
		Path:   child[0],
		Args:   child,
		Env:    envVars,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := c.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode(), err
		}
		return -1, err
	}
	return 0, nil
}
