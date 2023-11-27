hhlog
=====

The name stands for hacker ham log. This is a ham logger for hams who are also
hackers. If you don't want to waste your time on dealing with graphical
interfaces this is the tool for you.

Right now the functionality is very limited. It addresses my personal
aspirations in ham radio logging. I log my pota and sota contacts using
Rite-in-the-Rain books. However when I come back home I need to type the
contacts into a computer. Jumping between lots of text fields that usually make
up interfaces of most of ham logging software is not fun and is not exactly the
fastest way to do it.

This logger allows you to organize your contacts into a tab separated file with
flexible format. Vim editor is very efficient in dealing with tab separated
files. Then you invoke the logger with the template that you want to serialize
into an ADIF file and you are good to upload your logs to POTA or LOTW.

There are plans to add command line interface to this logger for logging during
contests. Again no jumping between text fields is expected. Type in only the
values that you have separated by space or tab and the logger will do the rest.

Contacts
========

https://groups.io/g/hhlog - group on groups.io for hhlog realted communications.
New features are announced in this group.

Usage
=====

Hhlog is best suited for converting tab separated files into various formats:
ADIF, cabrillo, TSV(sota submission format).

On the first line of your file specify the header - the format of the file.
Place fields strategically from most frequently changed to the least frequently
changed. If your format specifies 10 fields and a line contains only 5 then the
logger assumes that those 5 are the first 5 and fills the remaining fields with
the last specified values from the lines above. The header is a commented line
with tab separated flags. Currently the following flags are supported:

 - %skcc	- skcc number
 - %stx	- contest QSO transmitted serial number with a value greater than or equal to 0
 - %my_sota_ref	- the logging station's International SOTA Reference.
 - %cnty	- the contacted station's Secondary Administrative Subdivision (e.g. US county)
 - %d	- eight digits of date without spaces: year month day
 - %b	- band
 - %srx	- contest QSO received serial number with a value greater than or equal to 0
 - %my_call	- the logging station's Call Sign
 - %rst_rcvd	- signal report from the contacted station
 - %rst_sent	- signal report sent to the contacted station
 - %c	- call sign
 - %t	- four digits of UTC time
 - %m	- mode
 - %spc	- skcc spc
 - %ck	- contest check
 - %sect	- the contacted station's ARRL section
 - %my_state	- the logging station's state.
 - %f	- frequency in megahertz
 - %n	- the contacted station's operator name
 - %prec	- contest precedence
 - %state	- the code for the contacted station's Primary Administrative Subdivision (e.g. US State, JA Island, VE Province)
 - %my_pota_ref	- the logging station's POTA reference


Sample input file:
```
" %c	%t
q1bro	1020 " comment
q2bro	1120
" comment
q3bro	1130
```

After this you can execute hhlog to generate your ADIF file:

```
./hhlog -in input.hhl -out "adi" -tpl "%c %t"
```

Conventionally input files have extension hhl, though it's not important. They
are plain text files and you can check them into a git repo(github or
bitbucket). Text files are convenient to be inspected without any extra app.

Currently four output types are supported:

 - adi - adif format
 - cbr - cabrillo format
 - hhl - hhlog file format
 - tsv - tab separated file for sota submissions

Logs can be stored as plain files in hhl format and then converted to adi or cbr
as necessary.

The format of hhl file can be changed in the middle of the file by adding
another comment line that contains only tab separated field flags.

Submitting POTA Reports
-----------------------

This logger simplifies submitting your logs to POTA and WWFF coordinators.
References from both programs are similar and have the same numbers up to 4446.
Instead of sending two emails to the coordinators you can instruct hhlog to
submit one logs separately formatted for POTA and WWFF coordinator.

Create .hhlog.conf file

How to Organize Logs
--------------------

A recommended way for keeping your log would be to store hhl files in a
hierarchy of directories. The whole hierarchy can be stored in a version
control system. When a file or a set of files needs to be exported to LOTW or
cabrillo files you can convert hhl files to the required format. You only keep
hhl files that are just plain text files.
