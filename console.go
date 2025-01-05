package consolelogrus

// Just a logger for console that format things the way i like it
// Also provide options to change format and level colors
import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// CustomFormatter is a custom logrus formatter to set specific colors for
// each log level
type CustomFormatter struct {
	PaddingEnabled  bool
	TimestampFormat string
	LevelColors     map[logrus.Level]int
}

var (
	defaultFormat = "2006/01/02 15:04:05"
	defaultColors = map[logrus.Level]int{
		logrus.DebugLevel: 96, // Cyan
		logrus.InfoLevel:  32, // Green
		logrus.WarnLevel:  33, // Yellow
		logrus.ErrorLevel: 31, // Red
		logrus.FatalLevel: 95, // Magenta
		logrus.PanicLevel: 34, // Blue
	}

	maxLevelLength = len("WARNING")
)

// InitNewLogger initializes and returns a new instance of logrus.Logger
func InitNewLogger(cFormat *CustomFormatter) *logrus.Logger {

	if cFormat == nil {
		cFormat = NewCustomFormatter()
	}

	if cFormat.TimestampFormat == "" {
		cFormat.TimestampFormat = defaultFormat
	}

	if cFormat.LevelColors == nil {
		cFormat.LevelColors = defaultColors
	}

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(cFormat)
	return logger
}

// NewCustomFormatter creates a new instance of CustomFormatter
// with default settings
func NewCustomFormatter() *CustomFormatter {

	return &CustomFormatter{
		TimestampFormat: defaultFormat,
		LevelColors:     defaultColors,
		PaddingEnabled:  false,
	}
}

// NewColors creates a new map of colors for log levels
func NewColors(d, i, w, e, f, p string) map[logrus.Level]int {

	colors := make(map[logrus.Level]int)
	colors[logrus.DebugLevel] = getColorOrDefault(d,
		defaultColors[logrus.DebugLevel])
	colors[logrus.InfoLevel] = getColorOrDefault(i,
		defaultColors[logrus.InfoLevel])
	colors[logrus.WarnLevel] = getColorOrDefault(w,
		defaultColors[logrus.WarnLevel])
	colors[logrus.ErrorLevel] = getColorOrDefault(e,
		defaultColors[logrus.ErrorLevel])
	colors[logrus.FatalLevel] = getColorOrDefault(f,
		defaultColors[logrus.FatalLevel])
	colors[logrus.PanicLevel] = getColorOrDefault(p,
		defaultColors[logrus.PanicLevel])
	return colors
}

// getColorOrDefault returns the color if provided, otherwise
// returns the default color
func getColorOrDefault(color string, defaultColor int) int {

	if color == "" {
		return defaultColor
	}

	return parseColor(color)
}

// parseColor converts a color string to its corresponding ANSI code
func parseColor(color string) int {

	switch strings.ToLower(color) {
	case "black":
		return 30
	case "red":
		return 31
	case "green":
		return 32
	case "yellow":
		return 33
	case "blue":
		return 34
	case "magenta":
		return 35
	case "cyan":
		return 36
	case "white":
		return 37
	case "gray":
		return 90
	case "light red":
		return 91
	case "light green":
		return 92
	case "light yellow":
		return 93
	case "light blue":
		return 94
	case "light magenta":
		return 95
	case "light cyan":
		return 96
	case "bright white":
		return 97
	default:
		return 37 // Default to white if unknown
	}
}

// Format formats the log entry
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := time.Now().Format(f.TimestampFormat)
	levelColor := f.getColorByLevel(entry.Level)
	levelText := strings.ToUpper(entry.Level.String())
	padding := ""
	if f.PaddingEnabled {
		padding = strings.Repeat(" ", maxLevelLength-len(levelText))
	}

	// Adjust the format to align the timestamp vertically with spaces
	fmt.Fprintf(b, "\x1b[%dm[%s]%s\x1b[0m [%s] %s\n", levelColor, levelText,
		padding, timestamp, entry.Message)
	return b.Bytes(), nil
}

// SetColorByLevel sets the color for a specific log level
func (f *CustomFormatter) SetColorByLevel(level logrus.Level, color int) {
	f.LevelColors[level] = color
}

// getColorByLevel returns the color code for the given log level
func (f *CustomFormatter) getColorByLevel(level logrus.Level) int {

	if color, ok := f.LevelColors[level]; ok {
		return color
	}

	return 37 // Default to white
}
