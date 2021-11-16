package main

import (
	"errors"
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
		switch v {
		case FREQUENCY:
			setters[i] = FrequencySetter
		case CALL:
			setters[i] = CallSetter
		case DATE:
			setters[i] = DateSetter
		case TIME:
			setters[i] = TimeSetter
		case MODE:
			setters[i] = ModeSetter
		case SRX:
			setters[i] = SrxSetter
		case STX:
			setters[i] = StxSetter
		case PREC:
			setters[i] = PrecSetter
		case CK:
			setters[i] = CkSetter
		case SECT:
			setters[i] = SectSetter
		default:
			return nil, fmt.Errorf("Unknown verb: %v", v)
		}
	}
	return setters, nil
}

func parseWritingTemplate(line string) ([]FieldGetter, error) {

	var skccDb *SkccDB = nil

	verbs := strings.Split(line, " ")
	getters := make([]FieldGetter, len(verbs))
	for i, v := range verbs {
		switch v {
		case FREQUENCY:
			getters[i] = &FrequencyGetter{}
		case CALL:
			getters[i] = &CallGetter{}
		case DATE:
			getters[i] = &DateGetter{}
		case TIME:
			getters[i] = &TimeGetter{}
		case MODE:
			getters[i] = &ModeGetter{}
		case BAND:
			getters[i] = &BandGetter{}
		case SKCC:
			if skccDb == nil {
				skccDb = DownloadSkccRoster()
			}
			getters[i] = &SkccGetter{skccDb, ""}
		case NAME:
			if skccDb == nil {
				skccDb = DownloadSkccRoster()
			}
			getters[i] = &NameGetter{skccDb, ""}
		case SPC:
			if skccDb == nil {
				skccDb = DownloadSkccRoster()
			}
			getters[i] = &SpcGetter{skccDb, ""}
		case SRX:
			getters[i] = &SrxGetter{}
		case STX:
			getters[i] = &StxGetter{}
		case PREC:
			getters[i] = &PrecGetter{}
		case CK:
			getters[i] = &CkGetter{}
		case SECT:
			getters[i] = &SectGetter{}
		default:
			return nil, errors.New("Unknown verb: " + v)
		}
	}
	return getters, nil
}
