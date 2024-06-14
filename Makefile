# build is a target that compiles your Go code and outputs an executable named app in the bin directory
build: templ
	# '@' is used to suppress the command from being printed to the terminal
	# 'go build -o bin/app' builds the app and outputs the executable to bin/app
	@GOOS=linux GOARCH=amd64 go build -o ./tmp/app ./app/app.go

# run is a target that depends on the tailwindcss, templ and build targets
# it executes the app binary produced by the build target
run:
	# './bin/app' runs the compiled executable
	@go run ./app/app.go

# tailwindcss is a target that generates your Tailwind CSS styles
#css:
#	# '@bun run tailwindcss ...' runs TailwindCSS with the given parameters
#	# '--config configs/tailwind.config.js' specifies the Tailwind configuration file
#	# '-i configs/input.css -o static/css/styles.css' specifies the input CSS file and output file
#	@bun run tailwindcss --config tailwind.config.js -i static/css/input.css -o static/css/styles.css

# templ is a target that generates files from your templates using the Templ tool
templ:
	# '@templ generate' runs the 'generate' command in Templ
	@templ generate
	@templ fmt templates/landing/
	@templ fmt templates/inquiry_form/
	@templ fmt templates/common/