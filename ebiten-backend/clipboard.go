package ebitenbackend

// Text implements imgui clipboard
func (e *EbitenBackend) Text() (string, error) {
	return e.manager.cliptxt, nil
}

// SetText implements imgui clipboard
func (e *EbitenBackend) SetText(text string) {
	e.manager.cliptxt = text
}
