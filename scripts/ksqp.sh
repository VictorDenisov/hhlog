#!/bin/bash

source commons.sh
section=$(section_choice)
echo "Your location is: $section"

echo "Enter your call sign"
read callsign

mode=$(category_mode_choice)
echo "Mode: $mode"

cat << _end_of_text_ > log.txt
START-OF-LOG: 3.0
CONTEST: KS-QSO-PARTY
LOCATION: $section
CALLSIGN: $callsign
CATEGORY-BAND: ALL
CATEGORY-MODE: CW
CATEGORY-OPERATOR: SINGLE-OP
CATEGORY-POWER: LOW
CATEGORY-STATION: PORTABLE
CLAIMED-SCORE:
CREATED-BY: hhlog
NAME: Victor Denisov
ADDRESS: 1948 PO BOX
ADDRESS-CITY: novato
ADDRESS-STATE-PROVINCE: CA
ADDRESS-POSTALCODE: 94948
ADDRESS-COUNTRY: USA
EMAIL: denisovenator@gmail.com
OPERATORS:
SOAPBOX: 
QSO: 14000 CW 2022-08-27 1400 N6DVS         559 CA     W0BH          59  JAC    
END-OF-LOG:
_end_of_text_

