package ns

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

func (r Rename) output(ctx context.Context, resources []unstructured.Unstructured) error {
	for _, resource := range resources {
		select {
		case <-ctx.Done():
			return ctx.Err() //nolint:wrapcheck
		default:
		}
		bytes, err := yaml.Marshal(resource.Object)
		if err != nil {
			return err //nolint:wrapcheck
		}
		_, err = r.Output.Write([]byte("---\n"))
		if err != nil {
			return err //nolint:wrapcheck
		}
		_, err = r.Output.Write(bytes)
		if err != nil {
			return err //nolint:wrapcheck
		}
	}
	return nil
}
