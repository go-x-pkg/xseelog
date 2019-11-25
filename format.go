package xseelog

import "strings"

const (
	formatDate string = `[%Date(2/Jan/2006 15:04:05)]`
)

const ( // not used
	_formatFile string = formatDate + ` {{.Prefix}} [%l] %Msg%n`
)

// not used
const ( // 8bit colors
	_formatTrace string = `%EscM(0)` + formatDate +
		` %EscM(38;5;39){{.Prefix}}%EscM(0)` +
		` %EscM(38;5;39)[%l]%EscM(0)` +
		` %EscM(38;5;39)%Msg%EscM(0)%n`

	_formatDebug string = `%EscM(0)` + formatDate +
		` %EscM(38;5;39){{.Prefix}}%EscM(0)` +
		` %EscM(38;5;69)[%l]%EscM(0)` +
		` %EscM(38;5;69)%Msg%EscM(0)%n`

	_formatInfo string = `%EscM(0)` + formatDate +
		` %EscM(38;5;39){{.Prefix}}%EscM(0)` +
		` %EscM(32;1)[%l]%EscM(0)` +
		` %EscM(38;5;113)%Msg%EscM(0)%n`

	_formatWarn string = `%EscM(0)` + formatDate +
		` %EscM(38;5;39){{.Prefix}}%EscM(0)` +
		` %EscM(33;1)[%l]%EscM(0)` +
		` %EscM(38;5;220)%Msg%EscM(0)%n`

	_formatError string = `%EscM(0)` + formatDate +
		` %EscM(38;5;39){{.Prefix}}%EscM(0)` +
		` %EscM(38;5;196)[%l]%EscM(0)` +
		` %EscM(38;5;196)%Msg%EscM(0)%n`

	_formatCritical string = `%EscM(0)` + formatDate +
		` %EscM(38;5;39){{.Prefix}}%EscM(0)` +
		` %EscM(38;5;201)[%l]%EscM(0)` +
		` %EscM(38;5;201)%Msg%EscM(0)%n`
)

const ( // 4bit colors
	_formatTrace4bit string = `%EscM(0)` + formatDate +
		` %EscM(36){{.Prefix}}%EscM(0)` +
		` %EscM(36;1)[%l]%EscM(0)` +
		` %EscM(36)%Msg%EscM(0)%n`

	_formatDebug4bit string = `%EscM(0)` + formatDate +
		` %EscM(36){{.Prefix}}%EscM(0)` +
		` %EscM(34)[%l]%EscM(0)` +
		` %EscM(34;1)%Msg%EscM(0)%n`

	_formatInfo4bit string = `%EscM(0)` + formatDate +
		` %EscM(36){{.Prefix}}%EscM(0)` +
		` %EscM(32;1)[%l]%EscM(0)` +
		` %EscM(32)%Msg%EscM(0)%n`

	_formatWarn4bit string = `%EscM(0)` + formatDate +
		` %EscM(36){{.Prefix}}%EscM(0)` +
		` %EscM(33;1)[%l]%EscM(0)` +
		` %EscM(33)%Msg%EscM(0)%n`

	_formatError4bit string = `%EscM(0)` + formatDate +
		` %EscM(36){{.Prefix}}%EscM(0)` +
		` %EscM(35;1)[%l]%EscM(0)` +
		` %EscM(35)%Msg%EscM(0)%n`

	_formatCritical4bit string = `%EscM(0)` + formatDate +
		` %EscM(36){{.Prefix}}%EscM(0)` +
		` %EscM(31;1)[%l]%EscM(0)` +
		` %EscM(31)%Msg%EscM(0)%n`
)

func formatFile(prefix string) string {
	var b strings.Builder
	b.WriteString(formatDate)

	if prefix != "" {
		b.WriteByte(' ')
		b.WriteString(prefix)
	}

	b.WriteByte(' ')
	b.WriteString(`[%l]`)

	b.WriteByte(' ')
	b.WriteString(`%Msg`)

	b.WriteString(`%n`)

	return b.String()
}

func formatTrace(prefix string) string {
	var b strings.Builder
	b.WriteString(`%EscM(0)`)
	b.WriteString(formatDate)

	if prefix != "" {
		b.WriteByte(' ')
		b.WriteString(`%EscM(38;5;39)`)
		b.WriteString(prefix)
		b.WriteString(`%EscM(0)`)
	}

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;39)`)
	b.WriteString(`[%l]`)
	b.WriteString(`%EscM(0)`)

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;39)`)
	b.WriteString(`%Msg`)
	b.WriteString(`%EscM(0)`)

	b.WriteString(`%n`)

	return b.String()
}

func formatDebug(prefix string) string {
	var b strings.Builder
	b.WriteString(`%EscM(0)`)
	b.WriteString(formatDate)

	if prefix != "" {
		b.WriteByte(' ')
		b.WriteString(`%EscM(38;5;39)`)
		b.WriteString(prefix)
		b.WriteString(`%EscM(0)`)
	}

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;69)`)
	b.WriteString(`[%l]`)
	b.WriteString(`%EscM(0)`)

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;69)`)
	b.WriteString(`%Msg`)
	b.WriteString(`%EscM(0)`)

	b.WriteString(`%n`)

	return b.String()
}

func formatInfo(prefix string) string {
	var b strings.Builder
	b.WriteString(`%EscM(0)`)
	b.WriteString(formatDate)

	if prefix != "" {
		b.WriteByte(' ')
		b.WriteString(`%EscM(38;5;39)`)
		b.WriteString(prefix)
		b.WriteString(`%EscM(0)`)
	}

	b.WriteByte(' ')
	b.WriteString(`%EscM(32;1)`)
	b.WriteString(`[%l]`)
	b.WriteString(`%EscM(0)`)

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;113)`)
	b.WriteString(`%Msg`)
	b.WriteString(`%EscM(0)`)

	b.WriteString(`%n`)

	return b.String()
}

func formatWarn(prefix string) string {
	var b strings.Builder
	b.WriteString(`%EscM(0)`)
	b.WriteString(formatDate)

	if prefix != "" {
		b.WriteByte(' ')
		b.WriteString(`%EscM(38;5;39)`)
		b.WriteString(prefix)
		b.WriteString(`%EscM(0)`)
	}

	b.WriteByte(' ')
	b.WriteString(`%EscM(33;1)`)
	b.WriteString(`[%l]`)
	b.WriteString(`%EscM(0)`)

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;220)`)
	b.WriteString(`%Msg`)
	b.WriteString(`%EscM(0)`)

	b.WriteString(`%n`)

	return b.String()
}

func formatError(prefix string) string {
	var b strings.Builder
	b.WriteString(`%EscM(0)`)
	b.WriteString(formatDate)

	if prefix != "" {
		b.WriteByte(' ')
		b.WriteString(`%EscM(38;5;39)`)
		b.WriteString(prefix)
		b.WriteString(`%EscM(0)`)
	}

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;196)`)
	b.WriteString(`[%l]`)
	b.WriteString(`%EscM(0)`)

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;196)`)
	b.WriteString(`%Msg`)
	b.WriteString(`%EscM(0)`)

	b.WriteString(`%n`)

	return b.String()
}

func formatCritical(prefix string) string {
	var b strings.Builder
	b.WriteString(`%EscM(0)`)
	b.WriteString(formatDate)

	if prefix != "" {
		b.WriteByte(' ')
		b.WriteString(`%EscM(38;5;39)`)
		b.WriteString(prefix)
		b.WriteString(`%EscM(0)`)
	}

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;201)`)
	b.WriteString(`[%l]`)
	b.WriteString(`%EscM(0)`)

	b.WriteByte(' ')
	b.WriteString(`%EscM(38;5;201)`)
	b.WriteString(`%Msg`)
	b.WriteString(`%EscM(0)`)

	b.WriteString(`%n`)

	return b.String()
}
