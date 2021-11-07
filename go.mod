module post-pulumi

go 1.14

require (
	github.com/cheggaaa/pb v1.0.27 // indirect
	github.com/eminetto/post-pulumi/iac v0.0.0-00010101000000-000000000000
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
)

replace github.com/eminetto/post-pulumi/iac => ./iac
