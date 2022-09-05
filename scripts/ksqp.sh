#!/bin/bash

log_file=$1
file=$(basename $log_file .hhl)

source commons.sh
section=$(section_choice)
echo "Your location is: $section"

echo "Enter your call sign"
read callsign

mode=$(category_mode_choice)
echo "Mode: $mode"

op_category=$(category_operator_choice)
echo "Operator: $op_category"

power=$(power_choice)
echo "Power: $power"

station=$(station_choice)
echo "Station: $station"

echo "Enter your full name"
read full_name

echo "Enter your address"
read address

echo "Enter your city"
read city

echo "Enter your two letter state"
read state

echo "Enter your postal code"
read postal_code

echo "Enter your country"
read country

echo "Enter your email"
read email

cat << _end_of_text_ > $file.cab
START-OF-LOG: 3.0
CONTEST: KS-QSO-PARTY
LOCATION: $section
CALLSIGN: $callsign
CATEGORY-BAND: ALL
CATEGORY-MODE: $mode
CATEGORY-OPERATOR: $op_category
CATEGORY-POWER: $power
CATEGORY-STATION: $station
CLAIMED-SCORE:
CREATED-BY: hhlog
NAME: $full_name
ADDRESS: $address
ADDRESS-CITY: $city
ADDRESS-STATE-PROVINCE: $state
ADDRESS-POSTALCODE: $postal_code
ADDRESS-COUNTRY: $country
EMAIL: $email
OPERATORS:
SOAPBOX: 
_end_of_text_

hhlog -in $log_file -out cbr -tpl "%b %m %d %t %c %rst_rcvd %cnty" >> $file.cab

cat << _end_of_text_ >> $file.cab
END-OF-LOG:
_end_of_text_

