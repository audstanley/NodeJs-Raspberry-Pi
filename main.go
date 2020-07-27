package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

// The Color object is to pretty print in the linux terminal.
type Color string

// Declaring the color codes for linux terminal
const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorCyan         = "\u001b[36m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

// NodeElement stores the full version path as well as the direct link to the download
type NodeElement struct {
	BigVersion   int
	MedVersion   int
	SmallVersion int
	A7Link       string
	A6Link       string
	idx          int
}

// used to print colorized test in terminal
func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func colorizeWithoutNewLine(color Color, message string) {
	fmt.Print(string(color), message, string(ColorReset))
}

// PrintLogo prints the application logo as a concurrent process
func PrintLogo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\n\n")
	colorize(ColorCyan, "                                   ;                ")
	colorize(ColorCyan, "                                   +++              ")
	colorize(ColorCyan, "                                   +++              ")
	colorize(ColorCyan, "                                   +++              ")
	colorize(ColorCyan, "    ''++''       :;;',        ,+++;+++    ''++''    ")
	colorize(ColorCyan, " :++++++++++: ;;;;;;;''''  '++++'+++++ :++++++++++: ")
	colorize(ColorCyan, " +++.    .++' ';;;;;;;;''  +++     +++ '++. '' .+++ ")
	colorize(ColorCyan, " +++.    .++' '';;;;;;;;'  +++     +++ '++. '' ''   ")
	colorize(ColorCyan, " +++.    .++' ''';;;;;;;;  '++++'+++++ :+++++,      ")
	colorize(ColorCyan, " :          ,   '';;;,        ,+++;       ''++'.    ")
	colorize(ColorGreen, "                    #+''#';+''+'                   ")
	colorize(ColorGreen, "                    ';;;+#'+;;'                    ")
	colorize(ColorRed, "                     .###''@#@                       ")
	colorize(ColorRed, "                     '@++@@'+#@                      ")
	colorize(ColorRed, "                   :'@'''@'''''+                     ")
	colorize(ColorRed, "                    ##@@'''#@+#@                     ")
	colorize(ColorRed, "                    .'''#''@''#                      ")
	colorize(ColorRed, "                      '#'''#'                        ")
	fmt.Print("\n\n")
	colorizeWithoutNewLine(ColorYellow, "  Developed By:")
	colorize(ColorCyan, "Richard Stanley")
	colorize(ColorCyan, " https://www.audstanley.com")
	colorizeWithoutNewLine(ColorCyan, "  This installer also works on ")
	colorizeWithoutNewLine(ColorYellow, "x86_64")
	colorizeWithoutNewLine(ColorCyan, ", and ")
	colorizeWithoutNewLine(ColorYellow, "arm64")
	colorize(ColorCyan, "processors.")
	colorize(ColorCyan, " 👍 The only nodeJs installer that you really need for Linux")
	fmt.Print("\n\n")

}

// RequestForArchetectureOfficial makes the GET request (as a concurrent waitgroup) for the archetecture passed and returns []NodeBigVersion struct
func RequestForArchetectureOfficial(n *map[string]NodeElement, a7BigV *map[int][]int, latestVersionArm7 *int, w *sync.WaitGroup, arch *string) {
	// defer the waitgroup, which will end at the end of the function.  This is for concurrency, and saves
	// time to not make the https requests in series.
	defer w.Done()

	resp, err := http.Get("https://nodejs.org/dist/")
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the Body of the get request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Match all the versions that are avaible for downloading
	re := regexp.MustCompile(`<a href="v((\d{1,})\.(\d{1,})\.(\d{1,}))`)
	info := re.FindAllStringSubmatch(string(body), -1)

	// Populate the map with all the Node Elements for quick searching.
	for _, v := range info {
		link := v[1]
		bigV, _ := strconv.Atoi(v[2])
		medV, _ := strconv.Atoi(v[3])
		smallV, _ := strconv.Atoi(v[4])

		if bigV != 0 {
			(*n)[link] = NodeElement{bigV, medV, smallV, `https://nodejs.org/dist/v` + link + `/node-v` + link + `-linux-` + *arch + `.tar.gz`, `https://nodejs.org/dist/v` + link + `/node-v` + link + `-linux-armv6l.tar.gz`, -1}
			(*a7BigV)[bigV] = append((*a7BigV)[bigV], medV)
			if *latestVersionArm7 < bigV {
				*latestVersionArm7 = bigV
			}
		}
	}
}

// RequestForArm6Unofficial makes the GET request (as a concurrent waitgroup) for ARM6 processors and returns []NodeBigVersion struct
func RequestForArm6Unofficial(n *map[string]NodeElement, a6BigV *map[int][]int, latestVersionArm6 *int, w *sync.WaitGroup) {
	// defer the waitgroup, which will end at the end of the function.  This is for concurrency, and saves
	// time to not make the https requests in series.
	defer w.Done()

	resp, err := http.Get("https://unofficial-builds.nodejs.org/download/release/")
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the Body of the get request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Match all the versions that are avaible for downloading
	re := regexp.MustCompile(`<a href="v((\d{1,})\.(\d{1,})\.(\d{1,}))`)
	info := re.FindAllStringSubmatch(string(body), -1)

	// Populate the map with all the Node Elements for quick searching.
	for _, v := range info {
		link := v[1]
		bigV, _ := strconv.Atoi(v[2])
		medV, _ := strconv.Atoi(v[3])
		smallV, _ := strconv.Atoi(v[4])

		// For ARM6, we only need versions 12 and above for the unofficial releases
		if bigV >= 12 {
			(*n)[link] = NodeElement{bigV, medV, smallV, ``, `https://unofficial-builds.nodejs.org/download/release/v` + link + `/node-v` + link + `-linux-armv6l.tar.gz`, -1}
			(*a6BigV)[bigV] = append((*a6BigV)[bigV], medV)
			if *latestVersionArm6 < bigV {
				*latestVersionArm6 = bigV
			}
		}
	}

}

func deleteTheNodeJsDirectory(w *sync.WaitGroup) {
	defer w.Done()
	os.RemoveAll("/opt/nodejs/")
}

func deleteFolder(folder *string, w *sync.WaitGroup) {
	defer w.Done()
	os.RemoveAll(*folder)
}

func deleteFile(file *string, w *sync.WaitGroup) {
	defer w.Done()
	os.Remove(*file)

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// Untar takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func Untar(dst string, r io.Reader) error {

	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}

// CopyFile copies a single file from src to dst
func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// CopyDirectory copies a whole directory recursively
func CopyDirectory(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDirectory(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			} else {
				colorizeWithoutNewLine(ColorGreen, "Copying_File: "+srcfp+"\n    To: "+dstfp+"\n")
			}
		}
	}
	return nil
}

// SelectionDataFromUser prompts the user for which specific version of NodeJs they wish to install
func SelectionDataFromUser(a7 *map[string]NodeElement, version *string) string {
	versionMatched, _ := strconv.Atoi(regexp.MustCompile(`^([0-9]+)$`).FindStringSubmatch(*version)[1])
	var populatedVersions []NodeElement

	for _, element := range *a7 {
		if versionMatched == element.BigVersion {
			populatedVersions = append(populatedVersions, element)
		}
	}
	sort.SliceStable(populatedVersions, func(i, j int) bool {
		return populatedVersions[i].MedVersion < populatedVersions[j].MedVersion
	})

	for i, v := range populatedVersions {
		colorize(ColorYellow, "  "+strconv.Itoa(i)+": NodeJs version "+strconv.Itoa(v.BigVersion)+"."+strconv.Itoa(v.MedVersion)+"."+strconv.Itoa(v.SmallVersion)+" : "+v.A7Link)
	}

	reader := bufio.NewReader(os.Stdin)
	colorize(ColorCyan, "Please select a subversion from the list of integers")
	colorize(ColorCyan, "---------------------")

	selection := ""
	for {
		colorizeWithoutNewLine(ColorCyan, "-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if regexp.MustCompile(`^[0-9]+$`).MatchString(text) {
			userInput, _ := strconv.Atoi(regexp.MustCompile(`^([0-9]+)$`).FindStringSubmatch(text)[1])

			if userInput >= 0 && userInput < len(populatedVersions) {
				colorize(ColorGreen, "Selection is Valid")
				selection = strconv.Itoa(populatedVersions[userInput].BigVersion) + "." +
					strconv.Itoa(populatedVersions[userInput].MedVersion) + "." +
					strconv.Itoa(populatedVersions[userInput].SmallVersion)
				break

			} else {
				colorize(ColorRed, "Selection is invalid, Please input and integer from 0 to "+strconv.Itoa(len(populatedVersions)-1))
			}

		} else {
			colorize(ColorRed, "Invalid Selection")
		}
	}
	return selection
}

// GetTheMostLatestVersionOfNodeJs will return the most up to date version of NodeJs
func GetTheMostLatestVersionOfNodeJs(a7 *map[string]NodeElement, version *string, archetecture *string) string {
	versionMatched, _ := strconv.Atoi(regexp.MustCompile(`^([0-9]+)$`).FindStringSubmatch(*version)[1])
	var populatedVersions []NodeElement

	for _, element := range *a7 {
		if versionMatched == element.BigVersion {
			populatedVersions = append(populatedVersions, element)
		}
	}

	sort.SliceStable(populatedVersions, func(i, j int) bool {
		return populatedVersions[i].MedVersion < populatedVersions[j].MedVersion
	})

	idx := len(populatedVersions) - 1
	return strconv.Itoa(populatedVersions[idx].BigVersion) + "." + strconv.Itoa(populatedVersions[idx].MedVersion) + "." + strconv.Itoa(populatedVersions[idx].SmallVersion)

}

// RunInstallation will take the specific version number of NodeJs, and install for the appropriate archetecture
func RunInstallation(n *map[string]NodeElement, version *string, archetecture *string) string {
	if val, ok := (*n)[*version]; ok {
		link := val.A7Link
		if *archetecture == "armv6l" {
			link = val.A6Link
		}
		colorize(ColorGreen, "Downloading: "+link)
		err := DownloadFile("/tmp/node-v"+*version+"-linux-"+*archetecture+".tar.gz", link)

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		f, _ := os.Open("/tmp/node-v" + *version + "-linux-" + *archetecture + ".tar.gz")
		defer f.Close()
		colorize(ColorGreen, "Untaring: "+"/tmp/node-v"+*version+"-linux-"+*archetecture+".tar.gz")
		Untar("/tmp/node-v"+*version+"-linux-"+*archetecture, f)
		colorize(ColorGreen, "Copying files over to the appropriate location, and creating symlinks")
		err = CopyDirectory("/tmp/node-v"+*version+"-linux-"+*archetecture+"/node-v"+*version+"-linux-"+*archetecture, "/opt/nodejs")

		if err != nil {
			colorize(ColorRed, "There was a problem with making a copy of the nodeJs directory: node-v"+*version)
			fmt.Println(err)
			os.Exit(1)
		}

		err = CopyDirectory("/opt/nodejs/lib/node_modules", "/usr/lib/node_modules")
		err = CopyFile("/opt/nodejs/bin/node", "/usr/bin/node")
		currentDir, _ := os.Getwd()
		os.Chdir("/usr/bin")
		exec.Command("ln", "-sf", "../lib/node_modules/npm/bin/npm", "/usr/bin/npm").Run()
		exec.Command("ln", "-sf", "../lib/node_modules/npm/bin/npm-cli.js", "/usr/bin/npm-cli.js").Run()
		exec.Command("ln", "-sf", "../lib/node_modules/npm/bin/npx", "/usr/bin/npx").Run()
		exec.Command("ln", "-sf", "../lib/node_modules/", "/usr/bin/node_modules").Run()
		os.Chdir(currentDir)

	} else {
		colorize(ColorRed, *version+" is not an available version for NodeJs on "+*archetecture)
	}
	return "/tmp/node" + *version + "-linux-" + *archetecture

}

func runCommandConcurrent(command *string, args *[]string, w *sync.WaitGroup) {
	defer w.Done()
	cmd := exec.Command(*command, *args...)
	cmd.Run()
}

func main() {
	var wg sync.WaitGroup

	// We should varify that the user is root.
	user, _ := user.Current()

	if user.Uid != "0" {
		wg.Add(1)
		go PrintLogo(&wg)
		wg.Wait()
		colorize(ColorRed, "\n\n    You need to run node-install as root")
		os.Exit(1)
	}

	var latestVersionArm7 int
	var latestVersionArm6 int
	a6 := make(map[string]NodeElement)
	a7 := make(map[string]NodeElement)
	a6BigV := make(map[int][]int)
	a7BigV := make(map[int][]int)

	version := flag.String("v", "", "to install a specific version of NodeJs")
	latest := flag.Bool("a", false, "to install the latest version of NodeJs")

	var uname syscall.Utsname
	syscall.Uname(&uname)

	// When cross compiling for arm, we need to use a slightly different arrayToString function.
	// The compiler will work with the "-tags arm" argument that we assign to go build.
	// so "go build -tags arm", will use the arm.go file's function for the arrayToString function,
	// and go build (with no tag argument) will use the x64.go version of the arrayToString function.
	// There is a difference in the way the CPU archetecture deals with ascii integer values (as a [65]uint8 - unsigned, and not a [65]int8).
	archetecture := arrayToString(uname.Machine)

	if archetecture == "x86_64" {
		archetecture = "x64"
	} else if archetecture == "aarch64" {
		archetecture = "arm64"
	}

	nodeJsSymlinks := []string{"/usr/bin/node", "/usr/bin/nodejs", "/usr/lib/nodejs", "/usr/sbin/node", "/sbin/node", "/sbin/node", "/usr/local/bin/node", "/usr/bin/npm", "/usr/sbin/npm", "/sbin/npm", "/usr/local/bin/npm", "/usr/bin/node_modules"}
	updateAlternatives := "/usr/bin/update-alternatives"
	// nodeAndNpmSymlinks := [][]string{
	// 	{"--install", "/usr/bin/node", "node", "/opt/nodejs/bin/node", "1"},
	// 	{"--install", "/usr/bin/npm", "npm", "/opt/nodejs/lib/node_modules/npm/bin/npm", "1"},
	// 	{"--install", "/usr/bin/npm-cli.js", "npm-cli.js", "/opt/nodejs/lib/node_modules/npm/bin/npm-cli.js", "1"},
	// 	{"--install", "/usr/bin/npx", "npx", "/opt/nodejs/lib/node_modules/npm/bin/npx", "1"},
	// 	{"--install", "/usr/bin/npx-cli.js", "npx-cli.js", "/opt/nodejs/lib/node_modules/npm/bin/npx-cli.js", "1"},
	// }
	nodeAndNpmSymlinks := [][]string{
		{"--install", "/usr/bin/node", "node", "/opt/nodejs/bin/node", "1"},
		// {"--install", "/usr/bin/npm", "npm", "../lib/node_modules/npm/bin/npm", "1"},
		// {"--install", "/usr/bin/npm-cli.js", "npm-cli.js", "../lib/node_modules/npm/bin/npm-cli.js", "1"},
		// {"--install", "/usr/bin/npx", "npx", "../lib/node_modules/npm/bin/npx", "1"},
		// {"--install", "/usr/bin/npx-cli.js", "npx-cli.js", "../lib/node_modules/npm/bin/npx-cli.js", "1"},
	}

	nodeJsDirectory := "/opt/nodejs/"

	wgCount := 4 + len(nodeJsSymlinks)
	wg.Add(wgCount)
	go PrintLogo(&wg)
	for i := range nodeJsSymlinks {
		go deleteFile(&nodeJsSymlinks[i], &wg)
	}
	go deleteFolder(&nodeJsDirectory, &wg)
	go RequestForArchetectureOfficial(&a7, &a7BigV, &latestVersionArm7, &wg, &archetecture) // Making the body requests as a concurrent task
	go RequestForArm6Unofficial(&a6, &a6BigV, &latestVersionArm6, &wg)                      // Making the body requests as a concurrent task
	wg.Wait()

	colorize(ColorGreen, "Obtaining NodeJs for archetecture: "+archetecture)
	flag.Parse()

	// If the archetecture is arm6, we are going to use the a7 map of elements, but overwrite the links in the NodeElement.A6Link
	// This way, official builds up to version 12 will be installed for arm6, and unofficial builds will be installed for versions 12+
	// Ultimately, this will add arm6 support for the pi zero for as long as the unofficial builds are released.
	if archetecture == "armv6l" {

		for key := range a6 {
			a7[key] = a6[key]
		}

	}

	if *version != "" && regexp.MustCompile(`^[0-9]+$`).MatchString(*version) && !(*latest) {
		selection := SelectionDataFromUser(&a7, version)
		tmpFile := RunInstallation(&a7, &selection, &archetecture)
		tarFile := tmpFile + ".tar.gz"
		nodeDownloadedFolder := tmpFile + "/"
		wgCount = 2 + len(nodeAndNpmSymlinks)
		wg.Add(wgCount)
		for i := range nodeAndNpmSymlinks {
			go runCommandConcurrent(&updateAlternatives, &nodeAndNpmSymlinks[i], &wg)
		}
		go deleteFile(&tarFile, &wg)
		go deleteFolder(&nodeDownloadedFolder, &wg)
		wg.Wait()
		colorize(ColorGreen, "👍 good to go 👍")

	} else if *version != "" && regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`).MatchString(*version) && !(*latest) {
		tmpFile := RunInstallation(&a7, version, &archetecture)
		tarFile := tmpFile + ".tar.gz"
		nodeDownloadedFolder := tmpFile + "/"
		wgCount = 2 + len(nodeAndNpmSymlinks)
		wg.Add(wgCount)
		for i := range nodeAndNpmSymlinks {
			go runCommandConcurrent(&updateAlternatives, &nodeAndNpmSymlinks[i], &wg)
		}
		go deleteFile(&tarFile, &wg)
		go deleteFolder(&nodeDownloadedFolder, &wg)
		wg.Wait()
		colorize(ColorGreen, "👍 good to go 👍")
	} else if *latest && *version == "" {
		// install the latest version of NodeJs
		colorize(ColorGreen, "Installing latest version of NodeJs")
		latestVersion := strconv.Itoa(latestVersionArm7)
		latestVersionAsString := GetTheMostLatestVersionOfNodeJs(&a7, &latestVersion, &archetecture)
		colorize(ColorGreen, "  version: "+latestVersionAsString)
		tmpFile := RunInstallation(&a7, &latestVersionAsString, &archetecture)
		tarFile := tmpFile + ".tar.gz"
		nodeDownloadedFolder := tmpFile + "/"
		wgCount = 2 + len(nodeAndNpmSymlinks)
		wg.Add(wgCount)
		for i := range nodeAndNpmSymlinks {
			go runCommandConcurrent(&updateAlternatives, &nodeAndNpmSymlinks[i], &wg)
		}
		go deleteFile(&tarFile, &wg)
		go deleteFolder(&nodeDownloadedFolder, &wg)
		wg.Wait()
		colorize(ColorGreen, "👍 good to go 👍")

	} else if !*latest && *version == "" {
		colorize(ColorRed, "You need to at least specify one option")
		os.Exit(1)
	} else if *latest && *version != "" {
		colorize(ColorRed, "You cannot run the latest install flag : -a, and a version selection: -v at the same time. Use one or the other")
		os.Exit(1)
	}

}
