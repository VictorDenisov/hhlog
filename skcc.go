package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type SkccMember struct {
	Skcc string
	Call string
	Name string
	City string
	Spc  string
}

const (
	tableTag    = "<table>"
	tableEndTag = "</table>"
)

type SkccDB struct {
	list      []SkccMember
	callIndex map[string]SkccMember
}

func DownloadSkccRoster() *SkccDB {
	html, err := http.Get("https://www.skccgroup.com/membership_data/membership_roster.php")
	if err != nil {
		return nil
	}
	roster, err := ioutil.ReadAll(html.Body)
	html.Body.Close()
	if err != nil {
		return nil
	}
	rs := string(roster)
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
	res.Call = tr.ReadTag()
	res.Name = tr.ReadTag()
	res.City = tr.ReadTag()
	res.Spc = tr.ReadTag()
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
