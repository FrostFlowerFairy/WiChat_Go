package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWichat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wichat Suite")
}
