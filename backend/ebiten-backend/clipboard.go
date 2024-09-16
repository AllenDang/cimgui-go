package ebitenbackend

// Text implements imgui clipboard
func (e *EbitenBackend) Text() (string, error) {
	return e.cliptxt, nil
}

// SetText implements imgui clipboard
func (e *EbitenBackend) SetText(text string) {
	e.cliptxt = text
}
