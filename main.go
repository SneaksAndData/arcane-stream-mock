package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/SneaksAndData/arcane-stream-mock/pkg/apis/streaming/v1"
)

// +kubebuilder:rbac:groups=streaming.sneaksanddata.com,resources=teststreamdefinitions,verbs=get;list;watch;create;update;patch;delete
func main() {

	specEnv := os.Getenv("STREAMCONTEXT__SPEC")
	if specEnv == "" {
		fmt.Println("Error: STREAMCONTEXT__SPEC environment variable is not defined")
		os.Exit(1)
	}
	var spec v1.TestsStreamDefinitionSpec
	err := json.Unmarshal([]byte(specEnv), &spec)
	if err != nil {
		fmt.Println("Failed to deserialize STREAMCONTEXT__SPEC:", err)
		os.Exit(1)
	}

	fmt.Printf("Deserialized STREAMCONTEXT__SPEC: %+v\n", spec)

	// Parse the run duration
	duration, err := time.ParseDuration(spec.RunDuration)
	if err != nil {
		fmt.Println("Failed to parse RunDuration:", err)
		os.Exit(1)
	}

	// Generate random duration between 0 and runDuration
	randomFailureDuration := time.Duration(rand.Int64N(int64(duration)))

	if spec.ShouldFail {
		fmt.Printf("Simulating failure after %v\n", randomFailureDuration)
		time.Sleep(randomFailureDuration)
		panic("Simulated failure as per ShouldFail=true")
	}

	fmt.Printf("Running stream for %v\n", duration)
	time.Sleep(duration)
	fmt.Println("Stream run completed")
}
