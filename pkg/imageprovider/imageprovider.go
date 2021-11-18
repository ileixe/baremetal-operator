package imageprovider

import metal3 "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1"

type ImageProvider interface {
	// SupportsFormat returns whether the ImageProvider can provide images in
	// the given format.
	SupportsFormat(metal3.ImageFormat) bool

	// BuildImage requests the ImageProvider to build an image in the given
	// format.
	BuildImage([]metal3.ImageFormat) (string, metal3.ImageFormat, string)
}
