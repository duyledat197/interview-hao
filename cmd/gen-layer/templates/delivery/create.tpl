{{define "create"}}
func (d *{{.CamelCase}}Delivery) Create{{.PascalCase}}(ctx context.Context, req *pb.Create{{.PascalCase}}Request) (*pb.Create{{.PascalCase}}Response, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	if err := d.{{.CamelCase}}Service.Create{{.PascalCase}}(ctx, transform.PbTo{{.PascalCase}}Ptr(req.Get{{.PascalCase}}())); err != nil {
		return nil, status.Errof(codes.Internal,fmt.Errorf("Create: ",err).Error())
	}
	return &pb.Create{{.PascalCase}}Response{}, nil
}

{{end}}