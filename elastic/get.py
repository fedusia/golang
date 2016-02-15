#!/usr/bin/python
# -*- coding: utf-8 -*-

import requests
import json

host = 'elk-102.ix.km'
port = '9200'
response = requests.get('http://%s:%s/_cluster/health' % (host, port))
#print response
if response.status_code != 200:
	# TODO: Do error message to log file
	print ('Wrong response, %s') % response.status_code
	exit (1)

responseBody  = response.json()
if responseBody['status'] != 'green':
	# TODO: Do error message to log file
	print 'Status not green do nothing'
	exit (1)

#print ('Cluster Status is: %s') % responseBody['status']

#Get list indices
response = requests.get('http://%s:%s/_cat/indices' % (host, port))
responseBody = response.text
responseBody = responseBody.split('\n')
del responseBody[-1]
query = {'query': {' match_all': {}}}
#http://$address/_all/$type/_search
countIndex = dict();
for line in responseBody:
  elasticIndex = line.split(' ')[1]
  response = requests.get('http://%s:%s/%s/scribe_sv_maps_fullsearch_prod/_count' % (host, port, elasticIndex), query)
  countEvents = response.json()['count']
  countIndex[elasticIndex] = countEvents
  print (countEvents)

#for key, value in countIndex.iteritems() :
#  print ("indexName: %s, countEvents: %s" % (key, value))
#  response = requests.get('http://%s:%s/%s/scribe_sv_maps_fullsearch_prod/_search?scroll=1m&search_type=scan&pretty' % (host, port, key), query)
    
#print('qqq')
#response = requests.get('http://%s:%s/logstash-2015.10.12/scribe_sv_maps_fullsearch_prod/_search?pretty&size=988236' % (host, port), query)
#  data = response.text
#  print (data.encode("utf-8"))

