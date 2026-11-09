vim.keymap.set({ "n", "t" }, "<C-g>", function()
	Snacks.terminal.toggle(nil, { cwd = LazyVim.root() })
end, { desc = "Toggle terminal (Root Dir)" })
