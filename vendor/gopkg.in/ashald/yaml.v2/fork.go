package yaml


func (e *encoder) setWidth(width int) {
	yaml_emitter_set_width(&e.emitter, width)
}

// SetLineWidth sets the preferred line width.
// To disable long line breaks set width lower than zero.
// By default, line width is set to 80.
func (e *Encoder) SetLineWidth(width int) {
	e.encoder.setWidth(width)
}


func (e *Encoder) SetFlowStyle(flow bool) {
	e.encoder.flow = flow
}
