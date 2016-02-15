#!/usr/bin/env bash

#/usr/bin/curl -XGET 'http://elk-102.ix.km:9200/logstash-2015.10.27/scribe_sv_maps_fullsearch_prod/1'
/usr/bin/curl -XGET 'http://elk-102.ix.km:9200/logstash-2015.10.27/scribe_sv_maps_fullsearch_prod/_search?q=*&size=11152080'
