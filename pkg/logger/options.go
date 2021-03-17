package logger

// MaxSize maxsize
func MaxSize(size int) Option {
	return func(l *Logger) {
		l.maxsize = size
	}
}

// MaxAge maxage
func MaxAge(age int) Option {
	return func(l *Logger) {
		l.maxage = age
	}
}

// MaxBackups backups
func MaxBackups(b int) Option {
	return func(l *Logger) {
		l.maxbackups = b
	}
}

// Compress need compress
func Compress(c bool) Option {
	return func(l *Logger) {
		l.compress = c
	}
}

// Filename output filename
func Filename(name string) Option {
	return func(l *Logger) {
		l.filename = name
	}
}
