package opt_test

import (
	"net/url"
	"reflect"
	"split_download/opt"
	"testing"
)

func TestMain_parse(t *testing.T) {
	cases := map[string]struct {
		args                []string
		expectedParallelism int
		expectedOutput      string
		expectedURLString   string
	}{
		"no options": {args: []string{"http://example.com/foo.png"}, expectedParallelism: 8, expectedOutput: "foo.png", expectedURLString: "http://example.com/foo.png"},
		"-p=2":       {args: []string{"-p=2", "http://example.com/foo.png"}, expectedParallelism: 2, expectedOutput: "foo.png", expectedURLString: "http://example.com/foo.png"},
		"-o=bar.png": {args: []string{"-o=bar.png", "http://example.com/foo.png"}, expectedParallelism: 8, expectedOutput: "bar.png", expectedURLString: "http://example.com/foo.png"},
		"index.html": {args: []string{"http://example.com/"}, expectedParallelism: 8, expectedOutput: "index.html", expectedURLString: "http://example.com/"},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			args := c.args
			expectedParallelism := c.expectedParallelism
			expectedOutput := c.expectedOutput
			expectedURL, err := url.ParseRequestURI(c.expectedURLString)
			if err != nil {
				t.Fatalf("err %s", err)
			}

			opts, err := opt.Parse(args...)
			if err != nil {
				t.Fatalf("err %s", err)
			}

			actualParallelism := opts.Parallelism
			actualOutput := opts.Output
			actualURL := opts.URL

			if actualParallelism != expectedParallelism {
				t.Errorf(`unexpected parallelism: expected: %d actual: %d`, expectedParallelism, actualParallelism)
			}

			if actualOutput != expectedOutput {
				t.Errorf(`unexpected output: expected: "%s" actual: "%s"`, expectedOutput, actualOutput)
			}

			if !reflect.DeepEqual(actualURL, expectedURL) {
				t.Errorf(`unexpected URL: expected: "%s" actual: "%s"`, expectedURL, actualURL)
			}
		})
	}
}

func TestMain_parse_InvalidURL(t *testing.T) {
	t.Parallel()

	expected := "parse \"%\": invalid URI for request"

	_, err := opt.Parse([]string{"%"}...)
	if err == nil {
		t.Fatal("Unexpectedly err was nil")
	}

	actual := err.Error()
	if actual != expected {
		t.Errorf(`unexpected error: expected: "%s" actual: "%s"`, expected, actual)
	}
}
