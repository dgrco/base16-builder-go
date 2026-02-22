package main

import (
	. "base16-builder/colors"
	"bufio"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Load a scheme from standard input and unmarshal the data into the provided
// dict.
func loadScheme() map[string]string {
	var scheme map[string]string
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	typeErr := yaml.Unmarshal(data, &scheme)
	if typeErr != nil {
		fmt.Fprintln(os.Stderr, typeErr)
	}
	return scheme
}

// Create a dictionary with tags.
func getTemplateTags(s map[string]string) map[string]string {
	tags := make(map[string]string)

	tags["scheme-name"] = s["scheme"]
	tags["scheme-author"] = s["author"]
	tags["scheme-slug"] = slugify(s["name"])

	delete(s, "scheme")
	delete(s, "author")

	for key, val := range s {

		tags[key+"-hex"] = val
		tags[key+"-bgr"] = ReverseRgb(val)

		tags[key+"-hex-r"] = val[:2]
		tags[key+"-hex-g"] = val[2:4]
		tags[key+"-hex-b"] = val[4:6]

		rgb := HexToRgb(val)

		tags[key+"-rgb-r"] = strconv.FormatUint(rgb.R, 10)
		tags[key+"-rgb-g"] = strconv.FormatUint(rgb.G, 10)
		tags[key+"-rgb-b"] = strconv.FormatUint(rgb.B, 10)

		tags[key+"-dec-r"] = strconv.FormatFloat(float64(rgb.R)/255.0, 'f', 2, 64)
		tags[key+"-dec-g"] = strconv.FormatFloat(float64(rgb.G)/255.0, 'f', 2, 64)
		tags[key+"-dec-b"] = strconv.FormatFloat(float64(rgb.B)/255.0, 'f', 2, 64)

		hsl := RgbToHsl(rgb)

		tags[key+"-hsl-h"] = strconv.FormatFloat(hsl.H, 'f', 2, 64)
		tags[key+"-hsl-s"] = strconv.FormatFloat(hsl.S, 'f', 2, 64)
		tags[key+"-hsl-l"] = strconv.FormatFloat(hsl.L, 'f', 2, 64)

	}

	return tags
}

func slugify(name string) string {
	slug := strings.ToLower(name)
	slug = regexp.MustCompile(`\s+`).ReplaceAllString(slug, "-")
	return slug
}

func main() {
	var templatePath = flag.String("template", "", "Path to template")
	flag.Parse()

	scheme := loadScheme()

	tags := getTemplateTags(scheme)

	file, err := os.Open(*templatePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result strings.Builder

	var tagRegex = regexp.MustCompile(`\{\{.+?\}\}`)

	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := tagRegex.ReplaceAllStringFunc(line, func(match string) string {
			key := strings.Trim(match, "{{}}")
			if val, ok := tags[key]; ok {
				return val
			}
			return match
		})
		result.WriteString(modifiedLine + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Fprint(os.Stdout, result.String())
}
