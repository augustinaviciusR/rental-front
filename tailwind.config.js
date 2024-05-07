module.exports = {
    purge: {
        content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
        safelist: [
            // Add dynamic classes here
            'text-red-300',
            'text-4xl',
            'font-bold',
            'text-emerald-300'
        ]
    },
    theme: { extend: {} },
    plugins: [

    ],
}