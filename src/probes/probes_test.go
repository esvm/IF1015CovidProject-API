package probes_test

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/esvm/if1015covidproject-api/src/probes"
	"github.com/stretchr/testify/assert"
)

func TestProbes(t *testing.T) {
	port := os.Getenv("K8S_PROBES_PORT")

	p := probes.New(port)
	go p.Serve()

	time.Sleep(300 * time.Millisecond)

	// Liveness
	assertStatus(t, "http://localhost:"+port+probes.LivenessRoute, http.StatusInternalServerError) // Default should return 500

	p.SetLivenessCheck(func() bool {
		return true
	})

	assertStatus(t, "http://localhost:"+port+probes.LivenessRoute, http.StatusOK)

	// Readiness
	assertStatus(t, "http://localhost:"+port+probes.ReadinessRoute, http.StatusInternalServerError) // Default should return 500

	p.SetReadinessCheck(func() bool {
		return true
	})

	assertStatus(t, "http://localhost:"+port+probes.ReadinessRoute, http.StatusOK)

}

func assertStatus(t *testing.T, route string, status int) {
	resp, err := http.Get(route)
	assert.NoError(t, err)
	assert.Equal(t, status, resp.StatusCode)
	resp.Body.Close()
}
