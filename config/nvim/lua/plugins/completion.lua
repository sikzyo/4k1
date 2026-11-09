return {
	{
		"saghen/blink.cmp",
		opts = {
			keymap = {
				preset = "super-tab",
				["<CR>"] = { "fallback" },
				["<Tab>"] = { "select_and_accept", "snippet_forward", "fallback" },
				["<C-l>"] = { "snippet_forward", "fallback" },
				["<S-Tab>"] = { "snippet_backward", "fallback" },
			},
		},
	},
}
