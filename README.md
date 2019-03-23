# TRT MARCH 2019

## NAME
Torrent RSS Tracker

## SYNOPSIS
trt [url]

## DESCRIPTION
Trt will scrape the provided url for all RSS item titles that match the set
of regexs that are found in the filter.txt file. If no url is provided, trt
will read the RSS data from stdin. Trt will print each link of the torrent 
that matches at least one regex from filter.txt followed by a newline.
