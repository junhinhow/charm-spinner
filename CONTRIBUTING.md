# Contributing to charm-spinner

Thank you for your interest in contributing!

## How to Contribute

1. Fork the repository
2. Create a feature branch (`git checkout -b feat/my-spinner`)
3. Add your spinner to `spinner.go` and register it in the `All()` function
4. Write tests in `spinner_test.go`
5. Run tests: `go test ./...`
6. Commit using conventional commits (`feat:`, `fix:`, `docs:`)
7. Open a Pull Request

## Adding a New Spinner

Each spinner needs:
- Unique Unicode animation frames (at least 2 frames)
- An appropriate interval (typically 70-250ms)
- A descriptive variable name

## Code Style

- Run `gofmt` before committing
- Keep comments clear and concise
- All exported symbols must have doc comments

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
