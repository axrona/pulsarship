package init

func ZshInit() string {
	return `
setopt promptsubst

export PULSARSHIP_SHELL="zsh"

function pulsarship_prompt() {
  /usr/bin/pulsarship prompt
}

function pulsarship_right_prompt() {
  /usr/bin/pulsarship right
}

PROMPT='$(pulsarship_prompt)'
RPROMPT='$(pulsarship_right_prompt)'
`
}
