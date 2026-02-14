package main
//TODO: want to add a border that hugs the wall of the terminal screen but not right on the edge

import (
	"flag"
	"fmt"
	"log"
	"time"

	"GoStatus/ui/fonts"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// sessionState is used to track which model is focused
type sessionState uint

const (
	defaultTime              = time.Minute
	timerView   sessionState = iota
	spinnerView
)

var colorMap = map[string]string{
	"red":  "\033[91m",
	"green":  "\033[92m",
	"yellow":  "\033[93m",
	"blue":  "\033[94m",
	"magenta":  "\033[95m",
	"cyan":  "\033[96m",
	"white":  "\033[97m",
}

var (
	modelStyle = lipgloss.NewStyle().
			Width(14).
			Height(5).
			Align(lipgloss.Center, lipgloss.Center, lipgloss.Center).
			// BorderStyle(lipgloss.NormalBorder()). BorderForeground(lipgloss.Color("69"))
			BorderStyle(lipgloss.HiddenBorder())
	timeStyle = lipgloss.NewStyle().
			Width(30).
			Height(5).
			Align(lipgloss.Center, lipgloss.Center, lipgloss.Center).
			// BorderStyle(lipgloss.NormalBorder()). BorderForeground(lipgloss.Color("69"))
			BorderStyle(lipgloss.HiddenBorder())
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	// windowStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Align(lipgloss.Center).Padding(12)
)

type mainModel struct {
	state   sessionState
	timer   timer.Model
	contentMap map[string]fonts.Font
	content string
	index   int
	hour1, hour2, minute1, minute2 byte
	timeM bool
	timeString []string
	width, height int
	color string // using escape codes 
}

func newModel(timeout time.Duration) mainModel {
	m := mainModel{state: timerView}

	t := time.Now().Local().Format("03:04PM")
	colorFlag := flag.String("C", "", "color flag needs a value of red, green, yellow, blue, magenta, cyan, white")
	flag.Parse()
	fmt.Println(*colorFlag)
	m.color = colorMap[*colorFlag]

	m.hour1 = t[0]
	m.hour2 = t[1]
	m.minute1 = t[3]
	m.minute2 = t[4]
	m.timeM = t[5] == 'P'
	m.timer = timer.New(timeout)
	return m
}

type tickMsg time.Time

func tick() tea.Cmd {
  return tea.Tick(time.Second * 5, func(t time.Time) tea.Msg {
    return tickMsg(t)
  })
}

func (m mainModel) Init() tea.Cmd {
	return tick()
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {

  case tickMsg:
		t := time.Now().Local().Format("03:04PM")

		m.hour1 = t[0]
		m.hour2 = t[1]
		m.minute1 = t[3]
		m.minute2 = t[4]
		m.timeM = t[5] == 'P' // just checks if true or not

    return m, tick()
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// windowStyle.Padding(msg.Width*2)

  case tea.KeyMsg:
    switch msg.String() {
    case "q", "ctrl+c":
      return m, tea.Quit
    }
  }

  return m, nil
}

func (m mainModel) View() string {
	var s string
	s += helpStyle.Render("\nI am a mesage from the top\n")
	s += lipgloss.JoinHorizontal(lipgloss.Top, 
	modelStyle.Render(m.color + fonts.Rebels[m.hour1]), // 0 - 8
	modelStyle.Render(m.color + fonts.Rebels[m.hour2]),
	modelStyle.Render(m.color + fonts.Rebels[':']),// always the :
	modelStyle.Render(m.color + fonts.Rebels[m.minute1]),
	modelStyle.Render(m.color + fonts.Rebels[m.minute2]),
	timeStyle.Render(m.color + func () string{
		if m.timeM {
			return fonts.Rebels['P']
		}else{
			return fonts.Rebels['A']
		}
	}())) // 11 is am 12 is pm

	s += helpStyle.Render("\nq: exit\n")
	// s += lipgloss.Place(0, 0, lipgloss.Center, lipgloss.Center, s)
	// return s
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, s)
}

func main() {
	p := tea.NewProgram(newModel(defaultTime), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
