package init

func BashInit() string {
	return `
export PULSARSHIP_SHELL="bash"

__pulsarship_prompt() {
    PS1="$(/usr/bin/pulsarship prompt)"
}

PROMPT_COMMAND=__pulsarship_prompt
`
}
