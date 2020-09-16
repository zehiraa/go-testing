package calculation

import (
	"net/http"
	"testing"
)
// These constant gives us checkboxes for visualization.
const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// calculation örneği de var ancak zaman farkını görebilelim diye ekliyorum
// TestParallelize validates the http Get function can download content and handles different status conditions properly but runs the tests in parallel.
func TestParallelize(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{"statusok", "http://rss.cnn.com/rss/edition_world.rss", http.StatusOK},
		{"statusnotfound", "http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for _, tt := range tests {
			tf := func(t *testing.T) {
				// The only difference here is that we call Parallel function inside each of these
				// individual sub test function.
				t.Parallel()

				t.Logf("When checking %q for status code %d", tt.url, tt.statusCode)
				{
					resp, err := http.Get(tt.url)
					if err != nil {
						t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
					}
					t.Logf("\t%s\tShould be able to make the Get call.", succeed)

					defer resp.Body.Close()

					if resp.StatusCode == tt.statusCode {
						t.Logf("\t%s\tShould receive a %d status code.", succeed, tt.statusCode)
					} else {
						t.Errorf("\t%s\tShould receive a %d status code : %v", failed, tt.statusCode, resp.StatusCode)
					}
				}
			}

			t.Run(tt.name, tf)
		}
	}
}