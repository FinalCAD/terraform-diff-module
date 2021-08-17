module github.com/FinalCAD/terraform-diff-module/diff

go 1.14

replace github.com/FinalCAD/terraform-diff-module/diff/cmd => ./cmd

replace github.com/FinalCAD/terraform-diff-module/diff/internal/diff => ./internal/diff

replace github.com/FinalCAD/terraform-diff-module/diff/internal/types => ./internal/types

require (
	github.com/FinalCAD/terraform-diff-module/diff/cmd v0.0.0-00010101000000-000000000000
	github.com/FinalCAD/terraform-diff-module/diff/internal/diff v0.0.0-00010101000000-000000000000 // indirect
	github.com/FinalCAD/terraform-diff-module/diff/internal/types v0.0.0-00010101000000-000000000000 // indirect
)
