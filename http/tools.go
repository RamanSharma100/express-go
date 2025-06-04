package http

type LoggerType interface {
	Log(message string)
	Debug(message string)
	Error(message string)
	Info(message string)
	Warn(message string)
	Fatal(message string)
	Trace(message string)
	WithFields(fields map[string]any) LoggerType
	WithField(key string, value any) LoggerType
	WithError(err error) LoggerType
}

type LoggerFunc func(message string)

func Logger() LoggerType {
	return logger{}
}

type logger struct {
	message string
	fields  map[string]any
}

func (l logger) Log(message string) {
	if l.message != "" {
		message = l.message + ": " + message
	}

	println(message)
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			println(k+":", v)
		}
	}
}

func (l logger) Debug(message string) {
	if l.message != "" {
		message = l.message + ": " + message
	}

	println("DEBUG:", message)
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			println(k+":", v)
		}
	}
}

func (l logger) Error(message string) {
	if l.message != "" {
		message = l.message + ": " + message
	}

	println("ERROR:", message)
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			println(k+":", v)
		}
	}
}

func (l logger) Info(message string) {
	if l.message != "" {
		message = l.message + ": " + message
	}

	println("INFO:", message)
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			println(k+":", v)
		}
	}
}

func (l logger) Warn(message string) {
	if l.message != "" {
		message = l.message + ": " + message
	}

	println("WARN:", message)
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			println(k+":", v)
		}
	}
}

func (l logger) Fatal(message string) {
	if l.message != "" {
		message = l.message + ": " + message
	}

	println("FATAL:", message)
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			println(k+":", v)
		}

	}
}

func (l logger) Trace(message string) {
	if l.message != "" {
		message = l.message + ": " + message
	}

	println("TRACE:", message)
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			println(k+":", v)
		}
	}
}

func (l logger) WithFields(fields map[string]any) LoggerType {
	newFields := make(map[string]any)
	for k, v := range l.fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}
	return logger{
		message: l.message,
		fields:  newFields,
	}
}

func (l logger) WithField(key string, value any) LoggerType {
	newFields := make(map[string]any)
	for k, v := range l.fields {
		newFields[k] = v
	}
	newFields[key] = value
	return logger{
		message: l.message,
		fields:  newFields,
	}
}

func (l logger) WithError(err error) LoggerType {
	return l.WithField("error", err)
}
