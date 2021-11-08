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
		default:
			return nil, errors.New("Unknown verb: " + v)
		}
	}
	return getters, nil
}
