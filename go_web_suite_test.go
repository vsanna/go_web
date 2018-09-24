package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoWeb Suite")
}
