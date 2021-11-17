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
)

type FieldHandlers struct {
	setter FieldSetterConstructor
	getter FieldGetterConstructor
}

var (
	skccDb           = DownloadSkccRoster()
	templateHandlers = map[string]FieldHandlers{
		FREQUENCY: FieldHandlers{
			func() FieldSetter { return FrequencySetter },
			func() FieldGetter { return &FrequencyGetter{} },
		},
		CALL: FieldHandlers{
			func() FieldSetter { return CallSetter },
			func() FieldGetter { return &CallGetter{} },
		},
		DATE: FieldHandlers{
			func() FieldSetter { return DateSetter },
			func() FieldGetter { return &DateGetter{} },
		},
		TIME: FieldHandlers{
			func() FieldSetter { return TimeSetter },
			func() FieldGetter { return &TimeGetter{} },
		},
		MODE: FieldHandlers{
			func() FieldSetter { return ModeSetter },
			func() FieldGetter { return &ModeGetter{} },
		},
		BAND: FieldHandlers{
			//func() FieldSetter { return BandSetter },
			nil,
			func() FieldGetter { return &BandGetter{} },
		},
		SKCC: FieldHandlers{
			nil,
			//func() FieldSetter { return ModeSetter },
			func() FieldGetter { return &SkccGetter{skccDb, ""} },
		},
		NAME: FieldHandlers{
			nil,
			//func() FieldSetter { return ModeSetter },
			func() FieldGetter { return &NameGetter{skccDb, ""} },
		},
		SPC: FieldHandlers{
			nil,
			//func() FieldSetter { return ModeSetter },
			func() FieldGetter { return &SpcGetter{skccDb, ""} },
		},
		SRX: FieldHandlers{
			func() FieldSetter { return SrxSetter },
			func() FieldGetter { return &SrxGetter{} },
		},
		STX: FieldHandlers{
			func() FieldSetter { return StxSetter },
			func() FieldGetter { return &StxGetter{} },
		},
		PREC: FieldHandlers{
			func() FieldSetter { return PrecSetter },
			func() FieldGetter { return &PrecGetter{} },
		},
		CK: FieldHandlers{
			func() FieldSetter { return CkSetter },
			func() FieldGetter { return &CkGetter{} },
		},
		SECT: FieldHandlers{
			func() FieldSetter { return SectSetter },
			func() FieldGetter { return &SectGetter{} },
		},
	}
)

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
