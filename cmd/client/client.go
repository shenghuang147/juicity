package main

import "os"

func main() {
        if _, ok := os.LookupEnv("QUIC_GO_DISABLE_GSO"); !ok {
                os.Setenv("QUIC_GO_DISABLE_GSO", "1")
        }

	if err := Execute(); err != nil {
		os.Exit(1)
	}
}
