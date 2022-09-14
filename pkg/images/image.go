package images

import (
	"fmt"
	"regexp"
)

type Image struct {
	Registry string
	User     string
	Name     string
	Tag      string
}

func (I *Image) Marshal() (string, error) {
	if I.User != "" {
		if I.Tag != "" {
			return fmt.Sprintf("%s/%s/%s:%s", I.Registry, I.User, I.Name, I.Tag), nil
		} else {
			return "", fmt.Errorf("no Tag value found")
		}
	} else {
		if I.Tag != "" {
			return fmt.Sprintf("%s/%s:%s", I.Registry, I.Name, I.Tag), nil
		} else {
			return "", fmt.Errorf("no Tag value found")
		}
	}
}

func ParseImage(image string) (*Image, error) {
	img := &Image{}
	r := regexp.MustCompile(`(?P<Registry>[a-z0-9\-.]+\.[a-z0-9\-]+:?[0-9]*)?/?((?P<Name>[a-zA-Z0-9-_]+?)|(?P<UserName>[a-zA-Z0-9-_]+?)/(?P<ImageName>[a-zA-Z0-9-_/]+?))(:(?P<Tag>[a-zA-Z0-9-_\\.]+?)|)$`)
	if r.MatchString(image) {
		matches := r.FindStringSubmatch(image)
		if matches[1] != "" {
			img.Registry = matches[1]
		} else {
			img.Registry = "docker.io"
		}
		if matches[3] != "" {
			img.Name = matches[3]
		} else {
			img.User = matches[4]
			img.Name = matches[5]
		}
		if matches[7] != "" {
			img.Tag = matches[7]
		} else {
			img.Tag = "latest"
		}
		return img, nil
	} else {
		return nil, fmt.Errorf("%s does not appear to be a valid container image: %+v", image, img)
	}
}
