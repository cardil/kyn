package ns

import (
	"context"

	"github.com/manifestival/manifestival"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func (r Rename) injectNamespace(m manifestival.Manifest) (manifestival.Manifest, error) {
	for _, namespaceRename := range r.Namespaces {
		tm, err := m.Transform(manifestival.InjectNamespace(namespaceRename.To))
		if err != nil {
			return manifestival.Manifest{}, err //nolint:wrapcheck
		}
		m = tm
	}
	return m, nil
}

func (r Rename) readManifest(ctx context.Context) (manifestival.Manifest, error) {
	m, _ := manifestival.ManifestFrom(manifestival.Slice([]unstructured.Unstructured{}))
	for _, in := range r.Inputs {
		im, err := r.readInput(ctx, in, m)
		if err != nil {
			return manifestival.Manifest{}, err
		}
		m = im
	}
	return m, nil
}

func (r Rename) readInput(ctx context.Context, input Input, manifest manifestival.Manifest) (manifestival.Manifest, error) {
	select {
	case <-ctx.Done():
		return manifestival.Manifest{}, ctx.Err() //nolint:wrapcheck
	default:
	}
	reads, err := input.Read()
	defer func() {
		for _, read := range reads {
			_ = read.Close()
		}
	}()
	if err != nil {
		return manifestival.Manifest{}, err //nolint:wrapcheck
	}
	for _, read := range reads {
		m, err := manifestival.ManifestFrom(
			manifestival.Reader(read))
		if err != nil {
			return manifestival.Manifest{}, err //nolint:wrapcheck
		}
		manifest = manifest.Append(m)
	}
	return manifest, nil
}
