package main

import (
	"fmt"
	"os"
	"strings"

	textinput "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// These imports will be used later on the tutorial. If you save the file
// now, Go might complain they are unused, but that's fine.
// You may also need to run `go mod tidy` to download bubbletea and its
// dependencies.

type model struct {
	result    string
	textInput textinput.Model
	err       error
	palette   string
	shade     string
}

type DraculaColors = map[string]map[string]string

var dracula_colors = DraculaColors{
	"darker": {
		"50":      "#cdd0e4",
		"100":     "#b5bad6",
		"200":     "#9ea4c8",
		"300":     "#888fb8",
		"400":     "#727aa8",
		"500":     "#5f6795",
		"600":     "#53597c",
		"700":     "#454a64",
		"800":     "#373a4d",
		"900":     "#282a36",
		"DEFAULT": "#282a36",
	},
	"dark": {
		"50":      "#f7f7fb",
		"100":     "#dfe1ed",
		"200":     "#c7cadf",
		"300":     "#b1b5d0",
		"400":     "#9ba0c0",
		"500":     "#858bb0",
		"600":     "#71779f",
		"700":     "#606689",
		"800":     "#525771",
		"900":     "#44475a",
		"DEFAULT": "#44475a",
	},
	"light": {
		"50":      "#f8f8f2",
		"100":     "#eaeada",
		"200":     "#dcdcc3",
		"300":     "#cdcdac",
		"400":     "#bdbd96",
		"500":     "#adad81",
		"600":     "#9c9c6c",
		"700":     "#85855d",
		"800":     "#6d6d4f",
		"900":     "#565641",
		"DEFAULT": "#f8f8f2",
	},
	"blue": {
		"50":      "#f0f2f8",
		"100":     "#d7dcec",
		"200":     "#bec6df",
		"300":     "#a6b0d1",
		"400":     "#8f9bc3",
		"500":     "#7886b4",
		"600":     "#6272a4",
		"700":     "#54628c",
		"800":     "#485273",
		"900":     "#3b425a",
		"DEFAULT": "#6272a4",
	},
	"cyan": {
		"50":      "#fafeff",
		"100":     "#d4f8ff",
		"200":     "#aff0fe",
		"300":     "#8be9fd",
		"400":     "#67e1fb",
		"500":     "#44d9f8",
		"600":     "#22d0f5",
		"700":     "#0dbfe5",
		"800":     "#0ca0bf",
		"900":     "#0c819a",
		"DEFAULT": "#8be9fd",
	},
	"green": {
		"50":      "#e2ffe9",
		"100":     "#bcffcd",
		"200":     "#97feb1",
		"300":     "#73fc96",
		"400":     "#50fa7b",
		"500":     "#2df760",
		"600":     "#0cf346",
		"700":     "#0ccd3d",
		"800":     "#0ba833",
		"900":     "#0a8329",
		"DEFAULT": "#50fa7b",
	},
	"orange": {
		"50":      "#ffefdd",
		"100":     "#ffdcb7",
		"200":     "#ffca92",
		"300":     "#ffb86c",
		"400":     "#fda648",
		"500":     "#fb9325",
		"600":     "#f38107",
		"700":     "#cd6d08",
		"800":     "#a75a08",
		"900":     "#824707",
		"DEFAULT": "#ffb86c",
	},
	"pink": {
		"50":      "#ffeaf6",
		"100":     "#ffc4e6",
		"200":     "#ff9fd6",
		"300":     "#ff79c6",
		"400":     "#fd55b6",
		"500":     "#fb31a5",
		"600":     "#f80e95",
		"700":     "#d90880",
		"800":     "#b3086b",
		"900":     "#8e0855",
		"DEFAULT": "#ff79c6",
	},
	"purple": {
		"50":      "#fefeff",
		"100":     "#e9dafe",
		"200":     "#d3b6fc",
		"300":     "#bd93f9",
		"400":     "#a771f6",
		"500":     "#924ff2",
		"600":     "#7d2eed",
		"700":     "#6916e0",
		"800":     "#5914bb",
		"900":     "#491298",
		"DEFAULT": "#bd93f9",
	},
	"red": {
		"50":      "#ffecec",
		"100":     "#ffc6c6",
		"200":     "#ffa0a0",
		"300":     "#ff7b7b",
		"400":     "#ff5555",
		"500":     "#fd3131",
		"600":     "#fb0e0e",
		"700":     "#dd0606",
		"800":     "#b70707",
		"900":     "#910707",
		"DEFAULT": "#ff5555",
	},
	"yellow": {
		"50":      "#fefff8",
		"100":     "#fafed3",
		"200":     "#f6fcaf",
		"300":     "#f1fa8c",
		"400":     "#ebf769",
		"500":     "#e5f347",
		"600":     "#dfef26",
		"700":     "#ccdd12",
		"800":     "#abb811",
		"900":     "#899410",
		"DEFAULT": "#f1fa8c",
	},
	"nosferatu": {
		"50":      "#cdd0e4",
		"100":     "#b5bad6",
		"200":     "#9ea4c8",
		"300":     "#888fb8",
		"400":     "#727aa8",
		"500":     "#5f6795",
		"600":     "#53597c",
		"700":     "#454a64",
		"800":     "#373a4d",
		"900":     "#282a36",
		"DEFAULT": "#282a36",
	},
	"aro": {
		"50":      "#f7f7fb",
		"100":     "#dfe1ed",
		"200":     "#c7cadf",
		"300":     "#b1b5d0",
		"400":     "#9ba0c0",
		"500":     "#858bb0",
		"600":     "#71779f",
		"700":     "#606689",
		"800":     "#525771",
		"900":     "#44475a",
		"DEFAULT": "#44475a",
	},
	"cullen": {
		"50":      "#f8f8f2",
		"100":     "#eaeada",
		"200":     "#dcdcc3",
		"300":     "#cdcdac",
		"400":     "#bdbd96",
		"500":     "#adad81",
		"600":     "#9c9c6c",
		"700":     "#85855d",
		"800":     "#6d6d4f",
		"900":     "#565641",
		"DEFAULT": "#f8f8f2",
	},
	"vonCount": {
		"50":      "#f0f2f8",
		"100":     "#d7dcec",
		"200":     "#bec6df",
		"300":     "#a6b0d1",
		"400":     "#8f9bc3",
		"500":     "#7886b4",
		"600":     "#6272a4",
		"700":     "#54628c",
		"800":     "#485273",
		"900":     "#3b425a",
		"DEFAULT": "#6272a4",
	},
	"vanHelsing": {
		"50":      "#fafeff",
		"100":     "#d4f8ff",
		"200":     "#aff0fe",
		"300":     "#8be9fd",
		"400":     "#67e1fb",
		"500":     "#44d9f8",
		"600":     "#22d0f5",
		"700":     "#0dbfe5",
		"800":     "#0ca0bf",
		"900":     "#0c819a",
		"DEFAULT": "#8be9fd",
	},
	"blade": {
		"50":      "#e2ffe9",
		"100":     "#bcffcd",
		"200":     "#97feb1",
		"300":     "#73fc96",
		"400":     "#50fa7b",
		"500":     "#2df760",
		"600":     "#0cf346",
		"700":     "#0ccd3d",
		"800":     "#0ba833",
		"900":     "#0a8329",
		"DEFAULT": "#50fa7b",
	},
	"morbius": {
		"50":      "#ffefdd",
		"100":     "#ffdcb7",
		"200":     "#ffca92",
		"300":     "#ffb86c",
		"400":     "#fda648",
		"500":     "#fb9325",
		"600":     "#f38107",
		"700":     "#cd6d08",
		"800":     "#a75a08",
		"900":     "#824707",
		"DEFAULT": "#ffb86c",
	},
	"buffy": {
		"50":      "#ffeaf6",
		"100":     "#ffc4e6",
		"200":     "#ff9fd6",
		"300":     "#ff79c6",
		"400":     "#fd55b6",
		"500":     "#fb31a5",
		"600":     "#f80e95",
		"700":     "#d90880",
		"800":     "#b3086b",
		"900":     "#8e0855",
		"DEFAULT": "#ff79c6",
	},
	"dracula": {
		"50":      "#fefeff",
		"100":     "#e9dafe",
		"200":     "#d3b6fc",
		"300":     "#bd93f9",
		"400":     "#a771f6",
		"500":     "#924ff2",
		"600":     "#7d2eed",
		"700":     "#6916e0",
		"800":     "#5914bb",
		"900":     "#491298",
		"DEFAULT": "#bd93f9",
	},
	"marcelin": {
		"50":      "#ffecec",
		"100":     "#ffc6c6",
		"200":     "#ffa0a0",
		"300":     "#ff7b7b",
		"400":     "#ff5555",
		"500":     "#fd3131",
		"600":     "#fb0e0e",
		"700":     "#dd0606",
		"800":     "#b70707",
		"900":     "#910707",
		"DEFAULT": "#ff5555",
	},
	"lincoln": {
		"50":      "#fefff8",
		"100":     "#fafed3",
		"200":     "#f6fcaf",
		"300":     "#f1fa8c",
		"400":     "#ebf769",
		"500":     "#e5f347",
		"600":     "#dfef26",
		"700":     "#ccdd12",
		"800":     "#abb811",
		"900":     "#899410",
		"DEFAULT": "#f1fa8c",
	},
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "palette,shade"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func process_result(input string) string {
	split := strings.Split(input, ",")

	var palette, shade, result string = "", "", ""

	if len(split) > 1 {
		palette = strings.Trim(split[0], " ")
		shade = strings.Trim(split[1], " ")
	} else {
		palette = strings.Trim(input, " ")
	}

	if dracula_colors[palette] == nil {
		return "error"
	}

	if shade != "" {
		result = string(dracula_colors[palette][shade])
	} else {
		result = fmt.Sprintf("%v", dracula_colors[strings.Trim(input, " ")])
	}

	return result
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.result = process_result(m.textInput.Value())
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Tell me what dracula color you need:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func main() {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()

	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.result != "" {
		fmt.Println(m.result)
	}
}
