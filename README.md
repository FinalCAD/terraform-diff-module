# terraform-diff-module

Terraform module to extract and show differences in secrets before applying.

## Terraform module

Directory `module` contains the terraform module itself.
At the root of the directory `test.tf` can be used to launch a local test with real values.

[Documentation](./module/README.md)

## Diff command

Terraform module use a command `diff` build with golang.

[Documentation](./diff/README.md)

## Release

Update `version.txt` before releasing a new version (`version.txt` is a symlink of `module/version.txt`)
