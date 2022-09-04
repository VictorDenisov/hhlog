#!/bin/bash

# Based on this page: http://www.arrl.org/section-abbreviations

section_choice()
{
        PS3='Choose your section. Specify your section call sign area: '
        select cs_area in 'Call Sign Area 1' 'Call Sign Area 2' 'Call Sign Area 3' 'Call Sign Area 4' 'Call Sign Area 5' 'Call Sign Area 6' 'Call Sign Area 7' 'Call Sign Area 8' 'Call Sign Area 9' 'Call Sign Area 0' 'Canadian Area Call Sign'
        do
                break
        done
        PS3='Choose your section: '
        case $REPLY in
                1)
                        sections=('Connecticut' 'Eastern Massachusetts' 'Maine' 'New Hampshire' 'Rhode Island' 'Vermont' 'Western Massachusetts')
                        abbrevs=('CT' 'EMA' 'ME' 'NH' 'RI' 'VT' 'WMA')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                2)
                        sections=('Eastern New York' 'New York City - Long Island' 'Northern New Jersey' 'Northern New York' 'Southern New Jersey' 'Western New York'
                        )
                        abbrevs=('ENY' 'NLI' 'NNJ' 'NNY' 'SNJ' 'WNY')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                3)
                        sections=('Delaware' 'Eastern Pennsylvania' 'Maryland-DC' 'Western Pennsylvania')
                        abbrevs=('DE' 'EPA' 'MDC' 'WPA')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                4)
                        sections=('Alabama' 'Georgia' 'Kentucky' 'North Carolina' 'Northern Florida' 'South Carolina' 'Southern Florida' 'West  Central Florida' 'Tennessee' 'Virginia' 'Puerto Rico' 'Virgin Islands')
                        abbrevs=('AL' 'GA' 'KY' 'NC' 'NFL' 'SC' 'SFL' 'WCF' 'TN' 'VA' 'PR' 'VI')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                5)
                        sections=('Arkansas' 'Louisiana' 'Mississippi' 'New Mexico' 'North Texas' 'Oklahoma' 'South Texas' 'West Texas')
                        abbrevs=('AR' 'LA' 'MS' 'NM' 'NTX' 'OK' 'STX' 'WTX')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                6)
                        sections=('East Bay' 'Los Angeles' 'Orange' 'Santa Barbara' 'Santa Clara Valley' 'San Diego' 'San Francisco' 'San Joaquin Valley' 'Sacramento Valley' 'Pacific')
                        abbrevs=('EB' 'LAX' 'ORG' 'SB' 'SCV' 'SDG' 'SF' 'SJV' 'SV' 'PAC')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                7)
                        sections=('Arizona' 'Eastern Washington' 'Idaho' 'Montana' 'Nevada' 'Oregon' 'Utah' 'Western Washington' 'Wyoming' 'Alaska')
                        abbrevs=('AZ' 'EWA' 'ID' 'MT' 'NV' 'OR' 'UT' 'WWA' 'WY' 'AK')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                8)
                        sections=('Michigan' 'Ohio' 'West Virginia')
                        abbrevs=('MI' 'OH' 'WV')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                9)
                        sections=('Illinois' 'Indiana' 'Wisconsin')
                        abbrevs=('IL' 'IN' 'WI')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                10)
                        sections=('Colorado' 'Iowa' 'Kansas' 'Minnesota' 'Missouri' 'Nebraska' 'North Dakota' 'South Dakota')
                        abbrevs=('CO' 'IA' 'KS' 'MN' 'MO' 'NE' 'ND' 'SD')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
                11)
                        sections=('Maritime' 'Newfoundland/Labrador' 'Prince Edward Island' 'Quebec' 'Ontario East' 'Ontario North' 'Ontario South' 'Greater Toronto Area' 'Manitoba' 'Saskatchewan' 'Alberta' 'British Columbia' 'Northern Territories')
                        abbrevs=('MAR' 'NL' 'PE' 'QC' 'ONE' 'ONN' 'ONS' 'GTA' 'MB' 'SK' 'AB' 'BC' 'NT')
                        select section in "${sections[@]}"; do
                                echo ${abbrevs[$REPLY - 1]}
                                break
                        done
                        ;;
        esac
}

category_mode_choice()
{
        PS3='Select mode: '
        select category_mode in CW DIGI FM RTTY SSB MIXED; do
                echo $category_mode
                break
        done
}
