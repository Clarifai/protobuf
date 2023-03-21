package typedeclimport

import subpkg "github.com/Clarifai/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
