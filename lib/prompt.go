package lib

import (
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/dcaiafa/nitro"
)

func init() {
	survey.InputQuestionTemplate = `
{{- if .ShowHelp }}{{- color .Config.Icons.Help.Format }}{{ .Config.Icons.Help.Text }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- color .Config.Icons.Question.Format }}{{ .Config.Icons.Question.Text }} {{color "reset"}}
{{- color "default+hb"}}{{ .Message }} {{color "reset"}}
{{- if .ShowAnswer}}
  {{- color "cyan"}}{{.Answer}}{{color "reset"}}{{"\n"}}
{{- else if .PageEntries -}}
  {{- .Answer}}
  {{- "\n"}}
  {{- range $ix, $choice := .PageEntries}}
    {{- if eq $ix $.SelectedIndex }}{{color $.Config.Icons.SelectFocus.Format }}{{ $.Config.Icons.SelectFocus.Text }} {{else}}{{color "default"}}  {{end}}
    {{- $choice.Value}}
    {{- color "reset"}}{{"\n"}}
  {{- end}}
{{- else }}
  {{- if .Default}}{{color "white"}}({{.Default}}) {{color "reset"}}{{end}}
  {{- .Answer -}}
{{- end}}`

	survey.SelectQuestionTemplate = `
{{- if .ShowHelp }}{{- color .Config.Icons.Help.Format }}{{ .Config.Icons.Help.Text }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- color .Config.Icons.Question.Format }}{{ .Config.Icons.Question.Text }} {{color "reset"}}
{{- color "default+hb"}}{{ .Message }}{{ .FilterMessage }}{{color "reset"}}
{{- if .ShowAnswer}}{{color "cyan"}} {{.Answer}}{{color "reset"}}{{"\n"}}
{{- else}}
  {{- "\n"}}
  {{- range $ix, $choice := .PageEntries}}
    {{- if eq $ix $.SelectedIndex }}{{color $.Config.Icons.SelectFocus.Format }}{{ $.Config.Icons.SelectFocus.Text }} {{else}}{{color "default"}}  {{end}}
    {{- $choice.Value}}
    {{- color "reset"}}{{"\n"}}
  {{- end}}
{{- end}}`

}

var errPromptUsage = errors.New(
	`invalid usage. Expected prompt(string?)`)

type promptOptions struct {
	Type          string      `nitro:"type"`
	Message       string      `nitro:"message"`
	Default       string      `nitro:"default"`
	Help          string      `nitro:"help"`
	Options       []string    `nitro:"options"`
	PageSize      int         `nitro:"pagesize"`
	FilterMessage string      `nitro:"filtermessage"`
	Validate      nitro.Value `nitro:"validate"`
	VimMode       bool        `nitro:"vimmode"`
}

var promptOptionsConv Value2Structer

func prompt(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 {
		return nil, errPromptUsage
	}

	askOptions := []survey.AskOpt{
		survey.WithIcons(func(icons *survey.IconSet) {
			icons.Question.Text = ""
		}),
		survey.WithShowCursor(true),
	}

	var opt promptOptions
	var p survey.Prompt
	if arg, ok := args[0].(nitro.String); ok {
		opt.Message = arg.String()
	} else {
		err := promptOptionsConv.Convert(args[0], &opt)
		if err != nil {
			return nil, err
		}
	}

	if opt.Validate != nil {
		validator, err := wrapSurveyValidator(m, opt.Validate)
		if err != nil {
			return nil, fmt.Errorf("invalid validator: %v", err)
		}
		askOptions = append(askOptions, survey.WithValidator(validator))
	}

	if opt.Type == "" {
		opt.Type = "input"
	}

	switch opt.Type {
	case "input":
		p = &survey.Input{
			Message: opt.Message,
			Default: opt.Default,
			Help:    opt.Help,
		}

	case "select":
		p = &survey.Select{
			Message:       opt.Message,
			Options:       opt.Options,
			Default:       opt.Default,
			Help:          opt.Help,
			PageSize:      opt.PageSize,
			FilterMessage: opt.FilterMessage,
			VimMode:       opt.VimMode,
		}

	default:
		return nil, fmt.Errorf("invalid type %q", opt.Type)
	}

	var resp string
	err := survey.AskOne(p, &resp, askOptions...)
	if err != nil {
		return nil, err
	}

	return []nitro.Value{nitro.NewString(resp)}, nil
}

func wrapSurveyValidator(vm *nitro.VM, fn nitro.Value) (survey.Validator, error) {
	call, ok := fn.(nitro.Callable)
	if !ok {
		return nil, fmt.Errorf("value type %v is not callable", nitro.TypeName(fn))
	}
	return func(ans interface{}) error {
		ansStr := ans.(string)
		_, err := vm.Call(call, []nitro.Value{nitro.NewString(ansStr)}, 0)
		if err != nil {
			var rerr *nitro.RuntimeError
			if errors.As(err, &rerr) {
				err = fmt.Errorf("%v", rerr.Message())
			}
			return err
		}
		return nil
	}, nil
}
