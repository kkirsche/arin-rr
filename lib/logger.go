package irr

import "log"

// NewLogger returns a new logger instance
func NewLogger(verbose bool) *Logger {
	return &Logger{
		Verbose: verbose,
	}
}

// Verbosefln prints the statement if in verbose mode
func (l *Logger) Verbosefln(format string, v ...interface{}) {
	if l.Verbose {
		log.Printf(format+"\n", v...)
	}
}

// Verbosef prints the statement if in verbose mode
func (l *Logger) Verbosef(format string, v ...interface{}) {
	if l.Verbose {
		log.Printf(format, v...)
	}
}

// Verboseln prints the statement if in verbose mode
func (l *Logger) Verboseln(v ...interface{}) {
	if l.Verbose {
		log.Println(v...)
	}
}

// Printfln prints the statement if in verbose mode
func (l *Logger) Printfln(format string, v ...interface{}) {
	log.Printf(format+"\n", v...)
}

// Printf prints the statement if in verbose mode
func (l *Logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Println prints the statement if in verbose mode
func (l *Logger) Println(v ...interface{}) {
	log.Println(v...)
}
