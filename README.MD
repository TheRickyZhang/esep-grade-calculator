# ESEP Grade Calculator

This library is a simple utility to calculate your grades.
The final letter grade comes from three categories: **assignments**, **exams**, **essays**.
Final numeric = weighted sum of category means, then mapped to a letter grade.

## Requirements
- Golang 1.24+

## Usage
Typical flow:
1) Add item scores to each category
2) Compute weighted numeric final
3) Map to letter grade

See *grade_calculator.go* for exported API and current weights/scale.

## Tests
```bash
cd grade_calculator
go test ./...
# with coverage
go test -coverprofile=coverage.txt ./...
go tool cover -html=coverage.txt
```

## Notes
- Categories: assignments, exams, essays
- Final = sum(mean(category) * weight(category))
- Letter grade from fixed scale

## Planned / verify
- Configurable weights/scale
- Empty-category handling and validation
- Rounding policy docs
- CI: run `go test` on push/PR

## License
TODO
