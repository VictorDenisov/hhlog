package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type SkccMember struct {
	Skcc string
	Call string
	Name string
	City string
	Spc  Spc
}

const (
	tableTag    = "<table>"
	tableEndTag = "</table>"
)

type SkccDB struct {
	list      []SkccMember
	callIndex map[string]SkccMember
}

const (
	skccCacheFile = "skcc-cache.html"
)

func DownloadSkccRoster() *SkccDB {
	rs, err := retrieveRoster()
	if err != nil {
		return nil
	}

	tableStart := strings.Index(rs, tableTag)
	tableEnd := strings.Index(rs, tableEndTag)
	table := rs[tableStart+len(tableTag) : tableEnd]
	table = StripHtmlComments(table)

	lr := &TagReader{table, "tr"}

	members := make([]SkccMember, 0)
	for {
		line := lr.ReadTag()
		if line == "" {
			break
		}
		members = append(members, ParseLine(line))
	}
	return &SkccDB{members, buildCallIndex(members)}
}

func retrieveRoster() (string, error) {
	skccCachePath := "." + skccCacheFile
	userHomeDir, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(userHomeDir, configDir)

		err = ensureConfigPathExists(configPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create config path: %v\n", err)
			fmt.Fprintf(os.Stderr, "Using current directory for cache")
		} else {
			skccCachePath = filepath.Join(configPath, skccCacheFile)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Failed to determine home directory: %v\n", err)
		fmt.Fprintf(os.Stderr, "Using current directory for cache")
	}

	file, err := os.Open(skccCachePath)
	if errors.Is(err, os.ErrNotExist) {
		html, err := http.Get("https://www.skccgroup.com/membership_data/membership_roster.php")
		if err != nil {
			return "", nil
		}
		roster, err := ioutil.ReadAll(html.Body)
		html.Body.Close()
		if err != nil {
			return "", nil
		}
		ioutil.WriteFile(skccCachePath, roster, 0666)
		return string(roster), nil
	}
	if err != nil {
		return "", err
	}
	roster, err := ioutil.ReadAll(file)
	return string(roster), nil
}

func buildCallIndex(members []SkccMember) (res map[string]SkccMember) {
	res = make(map[string]SkccMember)
	for _, v := range members {
		res[v.Call] = v
	}
	return
}

type TagReader struct {
	s   string
	tag string
}

func (r *TagReader) ReadTag() (res string) {
	var (
		trS = "<" + r.tag + ">"
		trE = "</" + r.tag + ">"
	)
	trStart := strings.Index(r.s, trS)
	if trStart == -1 {
		return ""
	}
	trEnd := strings.Index(r.s, trE)
	if trEnd == -1 {
		return ""
	}
	res = r.s[trStart+len(trS) : trEnd]
	r.s = r.s[trEnd+len(trE):]
	return
}

func ParseLine(s string) (res SkccMember) {
	tr := &TagReader{s, "td"}
	res.Skcc = tr.ReadTag()
	res.Call = strings.ToLower(tr.ReadTag())
	res.Name = tr.ReadTag()
	res.City = tr.ReadTag()
	res.Spc = Spc(tr.ReadTag())
	return res
}

// This function assumes that there is no open and close tags anywhere other then in the comments.
func StripHtmlComments(s string) (res string) {
	const (
		openTag  = "<!--"
		closeTag = "-->"
	)
	res = s
	for {
		start := strings.Index(res, openTag)
		if start == -1 {
			break
		}
		end := strings.Index(res, closeTag)
		if end == -1 {
			break
		}
		res = res[:start] + res[end+len(closeTag):]
	}
	return
}

func ensureConfigPathExists(configPath string) error {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(configPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
