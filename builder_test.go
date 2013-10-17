package main

import (
  "testing"
)

func TestGeneratesPrelude(t *testing.T) {
  t.Skip("Generator includes prelude template")
}

func TestGeneratesEpilogue(t *testing.T) {
  t.Skip("Generator includes epilogue template")
}

func TestGeneratesBase(t *testing.T) {
  t.Skip("Includes base image template")
}

func TestGeneratesPackages(t *testing.T) {
  t.Skip("Generates install line for each specified package")
}

func TestGeneratesConfigs(t *testing.T) {
  t.Skip("Adds configs tarball for base & each specified package")
}

func TestGeneratesExtras(t *testing.T) {
  t.Skip("Includes extras template for each specified package")
}
