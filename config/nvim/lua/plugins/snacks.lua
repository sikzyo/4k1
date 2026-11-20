local exclude = {
	".git",
	".cache",
	"tmp",
	"temp",

	"node_modules",
	"coverage",
	".next",
	".nuxt",
	".output",
	".turbo",
	".svelte-kit",

	"vendor",

	".dart_tool",
	".pub-cache",
	".flutter-plugins",
	".gradle",
	"Pods",
	"DerivedData",
}

return {
	{
		"folke/snacks.nvim",
		opts = {
			picker = {
				sources = {
					files = {
						hidden = true,
						ignored = true,
						exclude = exclude,
					},
					explorer = {
						hidden = true,
						ignored = true,
						exclude = exclude,
					},
				},
			},
		},
	},
}
