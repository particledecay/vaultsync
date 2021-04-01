package plugin

// Plugin is a platform-specific layer that knows how to communicate
// with the platform and retrieve the target data.
type Plugin interface {
	Name() string
}
