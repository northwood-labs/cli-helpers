# CLI Helpers

Patterns codified as to reduce yak-shaving.

## Long help text for Cobra

```go
var rootCmd = &cobra.Command{
  Use:   "command",
  Short: "Does a thing!",

  // This. ↓↓↓
  Long: clihelpers.LongHelpText(`
  Does an even longer thing!

  Longer description that is indented the same as the code.`),
  Run: func(cmd *cobra.Command, args []string) {
    // ...
  },
}

```

## Version screen for Cobra

```go
var versionCmd = clihelpers.VersionScreen()

func init() {
  rootCmd.AddCommand(versionCmd)
}
```

## Default table styles for Bubbletea

```go
t := table.New(
  table.WithColumns(columns),
  table.WithRows(rows),
  table.WithFocused(true),
  table.WithHeight(height),
)

t.SetStyles(
  clihelpers.DefaultTableStyles(), // <-- This.
)
```

## Default keymap for Bubbletea

```go
model struct {
  help     help.Model
  keys     clihelpers.KeyMap // <-- This.
  // ...
}
```

Then…

```go
var keys = clihelpers.KeyBindings // <-- This.
```

Then handle the key-presses in your `model.Update()` method.

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    // If we set a width on the help menu it can gracefully truncate
    // its view as needed.
    m.help.Width = msg.Width

  case tea.KeyMsg:
    switch {
    case key.Matches(msg, m.keys.Up):
      m.lastKey = "↑"
    case key.Matches(msg, m.keys.Down):
      m.lastKey = "↓"
    case key.Matches(msg, m.keys.Help):
      m.help.ShowAll = !m.help.ShowAll
    case key.Matches(msg, m.keys.Quit):
      m.quitting = true

      return m, tea.Quit
    }
    // ...
  }
  // ...

  return m, cmd
}
```

## Base outer table style for Bubbletea

```go
func (m model) View() string {
  if m.quitting {
    return ""
  }

  helpView := m.help.View(m.keys)
  baseStyle := clihelpers.BaseOuterTableStyle() // <-- Get the standard style.

  return baseStyle.Render(m.table.View()) + "\n" + helpView
}
```

## Spinners for Bubbletea

No library code for this. Just an example of how to spin on an async process, then pass the result back to the outer scope.

```go
import "github.com/charmbracelet/huh/spinner"

var (
  apiToken string // <-- (1) Defined at this level of the scope.

  logger = log.NewWithOptions(os.Stderr, log.Options{
    ReportTimestamp: true,
    TimeFormat:      time.Kitchen,
    Prefix:          "app-name",
  })

  // ...

  rootCmd = &cobra.Command{
    Use:   "...",
    Short: "...",
    Long:  "...",
    Run: func(cmd *cobra.Command, args []string) {

      err := spinner.New().
        Title("Fetching OAuth Bearer Token...").
        Type(spinner.Dots).
        Action(func(apiToken *string) func() { // <-- (3) Receive the variable with *.
          return func() {
            // Do async stuff...
            token := myAsyncProcess()

            *apiToken = token // <-- (4) Assign the result to the variable we passed into the closure.
          }
        }(&apiToken)). // <-- (2) Pass the variable from the outer scope into the closure with &.
        Run()
      if err != nil {
        logger.Fatal(err)
      }

      // ...
    }

    apiToken // <-- (5) Now contains the value from the async process.
  }
)
```
