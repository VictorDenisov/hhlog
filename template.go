package main

import (
	"fmt"
	"strings"
)

const (
	FREQUENCY = "%f"
	CALL      = "%c"
	DATE      = "%d"
	TIME      = "%t"
	BAND      = "%b"
	MODE      = "%m"
	SKCC      = "%skcc"
	NAME      = "%n"
	SPC       = "%spc"
	SRX       = "%srx"
	STX       = "%stx"
	PREC      = "%prec"
	CK        = "%ck"
	SECT      = "%sect"
	RST_RCVD  = "%rst_rcvd"
	RST_SENT  = "%rst_sent"
	STATE     = "%state"
)

type FieldHandlers struct {
	setter FieldSetterConstructor
	getter FieldGetterConstructor
	doc    string
}

var (
	skccDb           = DownloadSkccRoster()
	templateHandlers = map[string]FieldHandlers{
		FREQUENCY: FieldHandlers{
			func() FieldSetter { return FrequencySetter },
			func() FieldGetter { return &FrequencyGetter{} },
			fmt.Sprintf("%v\t- frequency in megahertz", FREQUENCY),
		},
		CALL: FieldHandlers{
			func() FieldSetter { return CallSetter },
			func() FieldGetter { return &CallGetter{} },
			fmt.Sprintf("%v\t- call sign", CALL),
		},
		DATE: FieldHandlers{
			func() FieldSetter { return DateSetter },
			func() FieldGetter { return &DateGetter{} },
			fmt.Sprintf("%v\t- eight digits of date without spaces: year month day", DATE),
		},
		TIME: FieldHandlers{
			func() FieldSetter { return TimeSetter },
			func() FieldGetter { return &TimeGetter{} },
			fmt.Sprintf("%v\t- four digits of UTC time", TIME),
		},
		MODE: FieldHandlers{
			func() FieldSetter { return ModeSetter },
			func() FieldGetter { return &ModeGetter{} },
			fmt.Sprintf("%v\t- band", BAND),
		},
		BAND: FieldHandlers{
			//func() FieldSetter { return BandSetter },
			nil,
			func() FieldGetter { return &BandGetter{} },
			fmt.Sprintf("%v\t- band", MODE),
		},
		SKCC: FieldHandlers{
			nil,
			//func() FieldSetter { return ModeSetter },
			func() FieldGetter { return &SkccGetter{skccDb, ""} },
			fmt.Sprintf("%v\t- skcc number", SKCC),
		},
		NAME: FieldHandlers{
			nil,
			//func() FieldSetter { return ModeSetter },
			func() FieldGetter { return &NameGetter{skccDb, ""} },
			fmt.Sprintf("%v\t- the contacted station's operator name", NAME),
		},
		SPC: FieldHandlers{
			nil,
			//func() FieldSetter { return ModeSetter },
			func() FieldGetter { return &SpcGetter{skccDb, ""} },
			fmt.Sprintf("%v\t- skcc spc", SPC),
		},
		SRX: FieldHandlers{
			func() FieldSetter { return SrxSetter },
			func() FieldGetter { return &SrxGetter{} },
			fmt.Sprintf("%v\t- contest QSO received serial number with a value greater than or equal to 0", SRX),
		},
		STX: FieldHandlers{
			func() FieldSetter { return StxSetter },
			func() FieldGetter { return &StxGetter{} },
			fmt.Sprintf("%v\t- contest QSO transmitted serial number with a value greater than or equal to 0", STX),
		},
		PREC: FieldHandlers{
			func() FieldSetter { return PrecSetter },
			func() FieldGetter { return &PrecGetter{} },
			fmt.Sprintf("%v\t- contest precedence", PREC),
		},
		CK: FieldHandlers{
			func() FieldSetter { return CkSetter },
			func() FieldGetter { return &CkGetter{} },
			fmt.Sprintf("%v\t- contest check", CK),
		},
		SECT: FieldHandlers{
			func() FieldSetter { return SectSetter },
			func() FieldGetter { return &SectGetter{} },
			fmt.Sprintf("%v\t- the contacted station's ARRL section", SECT),
		},
	}
)

func templateDoc() (res string) {
	for _, h := range templateHandlers {
		res += h.doc + "\n"
	}
	return
}

func isTemplateString(line string) bool {
	if len(line) == 0 {
		return false
	}
	tokens := strings.Split(line, "\t")
	n := len(tokens)
	verbCount := 0
	for _, v := range tokens {
		if _, ok := templateHandlers[v]; ok {
			verbCount++
		}
	}
	if n > 3 {
		return n-verbCount <= 2
	} else {
		return n == verbCount
	}
}

func parseReadingTemplate(line string) ([]FieldSetter, error) {

	verbs := strings.Split(line, "\t")
	setters := make([]FieldSetter, len(verbs))
	for i, v := range verbs {
		templateHanlder, ok := templateHandlers[v]
		if !ok {
			return nil, fmt.Errorf("Unknown verb: %v", v)
		}
		setters[i] = templateHanlder.setter()
	}
	return setters, nil
}

func parseWritingTemplate(line string) ([]FieldGetter, error) {

	verbs := strings.Split(line, " ")
	getters := make([]FieldGetter, len(verbs))
	for i, v := range verbs {
		templateHanlder, ok := templateHandlers[v]
		if !ok {
			return nil, fmt.Errorf("Unknown verb: %v", v)
		}
		getters[i] = templateHanlder.getter()
	}
	return getters, nil
}

type ValueVisitor struct {
	val string
}

func (v *ValueVisitor) visitFrequency(g *FrequencyGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitCall(g *CallGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitDate(g *DateGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitTime(g *TimeGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitMode(g *ModeGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitBand(g *BandGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitSkcc(g *SkccGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitName(g *NameGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitSpc(g *SpcGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitSrx(g *SrxGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitStx(g *StxGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitPrec(g *PrecGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitCk(g *CkGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitSect(g *SectGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitRstRcvd(g *RstRcvdGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitRstSent(g *RstSentGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitState(g *StateGetter) {
	v.val = string(g.val)
}
