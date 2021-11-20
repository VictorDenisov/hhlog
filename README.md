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

This logger allows you to type the fields that you have in a tab separated file.
Vim editor is very efficient in dealing with tab separated files. Then you
invoke the logger with the template that you want to serialize into an ADIF file
and you are good to upload your logs to POTA or LOTW.

There are plans to add command line interface to this logger for logging during
contests. Again no jumping between text fields is expected. Type in only the
values that you have separated by space or tab and the logger will do the rest.

Usage
=====

Right the main purpose of hhlog is to convert tab separated files into ADIF
files. ADIF files are very nicely formatted, but they are not very convenient
to be typed in by a user.

You can enter the values that are variable - call sign, time. It takes little
time to type in these values from your field note book. Then you can add
missing columns using Vim's vertical selection feature - Ctrl - v.

Specify in the header of your file the format of the file.
The header is a commented line with tab separated flags.
Currently only these flags are supported:

 - %f - frequency in megahertz
 - %c - call sign
 - %d - date. It's an eight digit value: year month day.
 - %t - time. It's a four digit UTC time: 24 hour, minute.
 - %b - band. It doesn't need to be specified, but can be used in the output template.
 - %m - mode.

Sample input file:
```
" %c	%t
q1bro	1020 " comment
q2bro	1120
" comment
q3bro	1130
```

After this you can excute hhlog to generate your ADIF file:

```
./hhlog -in input.hhl -out "adi" -tpl "%c %t"
```

The input files can have extension hhl. They are plain text files and you can
check them into a git repo(github or bitbucket). Text files are convenient to be
inspected without any extra app.

Currently two output types are supported: adi and cbr.

Logs can be stored as plain files in hhl format and then converted to adi or cbr
as necessary.

The format of hhl file can be changed in the middle of the file by adding
another comment line that contains only tab separated field flags.

How to Organize Logs
--------------------

A recommended way for keeping your log would be to store hhl files in a
hierarchy of directories. The whole hierarchy can be stored in a version
control system. When a file or a set of files needs to be exported to LOTW or
cabrillo files you can convert hhl files to the required format. You only keep
hhl files that are just plain text files.
