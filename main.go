/*
Pay attention in order to run this script first
you need to give this script a module name
the format is {REMOTE}/{USERNAME}/{module name}
we give this module a name with:
$ go mod init {REMOTE}/{USERNAME}/{module name}
after that we need to download the specified modules
and list them in go.mod go dose that automaticlly by
$ go mod tidy
To download and cache dependencies in the vendor directory
$ go mod vendor
finaly we can run the program by
$ go run <gofile>
To generate a binary file
$ go build <gofile>
tip1 A package named main has an entrypoint at the main() function. A main package is compiled into an executable program.
tip2 A package by any other name is a library package. Libraries have no entry point.
tip3 Go programs are organized into packages. A package is a directory of Go code that's all compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package (directory).
tip4 The go install command compiles and installs a package or packages on your local machine for your personal usage. It installs the package's compiled binary in the GOBIN directory.
tip5 to import packages in other libaries first you need to create a fodler with the name of the libary then create the script in that folder Then  you need to set the package in the script to the folder name now you create you libaray finally to use the libary in other projects you need to use import statement with the followiing format => import "{REMOTE}/{USERNAME}/{module name}/{libary name}"
$ go install
*/
/*
Package Documentation
*/
package main // main is for executable programs, for libaries use any other name and put the script in a folder with the same name as the package

// import packages
import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ANSI color codes for terminal output
const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	// Text Colors
	FgBlack   = "\033[30m"
	FgRed     = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue    = "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan    = "\033[36m"
	FgWhite   = "\033[37m"
	// Background Colors
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

/*
SetupLogger initializes zerolog to write to both console and a file.
*/
func SetupLogger(logFilePath string, logLevel zerolog.Level) (zerolog.Logger, error) {
	// Open or create the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return zerolog.Logger{}, err
	}

	// Console writer with human-friendly formatting
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Set global log level
	zerolog.SetGlobalLevel(logLevel)

	// Combine both writers
	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	// Set global time format
	zerolog.TimeFieldFormat = time.RFC3339

	// Create the logger
	logger := zerolog.New(multi).With().Caller().Timestamp().Logger()

	// Set as the global logger
	log.Logger = logger

	return logger, nil
}

/*
# What it dose

ScrapeDataXpath scrapes data from a webpage using the provided XPath and URL.
It initializes a new Colly collector, listens for elements matching the XPath,
and returns the first matched element or an error if the scraping fails.

# Parameters

- logger: A zerolog.Logger instance for logging.

- xpath: A string representing the XPath query to locate elements on the webpage.

- url: A string representing the URL of the webpage to scrape.

# Returns

- *colly.XMLElement: A pointer to the first matched XMLElement.

- error: An error if the scraping process encounters an issue.
*/
func ScrapeDataXpath(logger zerolog.Logger, xpath string, url string) (*colly.XMLElement, error) {
	// Create a new collector
	logger.Info().Str("FunctionName:", "ScrapeDataXpath").Msg("Starting to scrape data")
	c := colly.NewCollector()
	element := &colly.XMLElement{}
	logger.Debug().Str("FunctionName:", "ScrapeDataXpath").Msg("Created c colly collector and element colley XMLElement")
	c.OnXML(xpath, func(e *colly.XMLElement) {
		element = e
	})
	logger.Debug().Str("FunctionName:", "ScrapeDataXpath").Str("xpath", xpath).Str("URL", url).Msg("Going to start to extract element using the provided XPath and URL")
	err := c.Visit(url)
	if err != nil {
		logger.Error().Str("FunctionName:", "ScrapeDataXpath").Err(err).Msg("Error visiting URL")
		return nil, err
	}
	logger.Info().Str("FunctionName:", "ScrapeDataXpath").Msg("Finished scraping data")
	return element, nil
}

/*
# What it dose

Gets a map in the form of map[string][]string{"URL","XPATH"} and extracts the price from the URL using the XPATH provided

# Parameters

- logger: A zerolog.Logger instance for logging.
- currency_list: map in the form of map[string][]string{"URL","XPATH"}

# Returns

- map in the form of map[string]string{"currency":"price"}

- error: An error if the scraping process encounters an issue.
*/
func extractPrice(logger zerolog.Logger, url string, xpath string) (map[string]string, error) {
	//TODO implement this function
	fmt.Println("TODO")
	return nil, nil
}

// The function that will be executed
func main() {
	logger, err := SetupLogger("app.log", zerolog.InfoLevel)
	startTime := time.Now() // Record start time
	if err != nil {
		panic(err)
	}
	logger.Info().Str("FunctionName:", "main").Msg("Main function started")
	defer func() {
		logger.Info().Str("FunctionName:", "main").TimeDiff("Duration (ms)", time.Now(), startTime).Msg("Main function ended.")
	}()
	// Create a new map with string key and list values and put "amon":["aa","ads"]
	currencies := make(map[string][]string)
	currencies["usdolor"] = []string{"https://www.tgju.org/profile/price_dollar_rl", "/html/body/main/div[1]/div[1]/div[1]/div/div[1]/div[3]/div[1]/span[2]"}

	XMLElement, err := ScrapeDataXpath(logger, currencies["usdolor"][1], currencies["usdolor"][0])

	if err != nil {
		logger.Error().Str("FunctionName:", "main").Err(err).Msg("Error in ScrapeDataXpath")
	}
	logger.Info().Str("FunctionName:", "main").Str("Extracted Element Text", XMLElement.Text).Msg("Successfully extracted element text")
	price := XMLElement.Text
	logger.Info().Str("FunctionName:", "main").Str("Price", price).Msg("The extracted price is")
}
