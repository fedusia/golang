#!/usr/bin/python

import requests

indices = []
x = 9

response = requests.get('http://elk-102.ix.km:9200/_cluster/health')
if response.status_code != 200:
	# TODO: Do error message to log file
	print ('Wrong response, %s') % response.status_code
	exit (1)

print response.status_code

responseBody  = response.json()
if responseBody['status'] != 'green':
	# TODO: Do error message to log file
	print 'Status not green do nothing'
	exit (1)

print ('Cluster Status is: %s') % responseBody['status']

# http_req = requests.get('http://elk-102.ix.km:9200/_cat/indices?v')
# http_text = http_req.text
# text = http_text.split()
# print text
# i = len(text)
# while x < i:
# indices.append(text[x])
# x = x + 8
# print indices
# print len(indices)
# llen = len(indices) - 60
# print llen
# if llen >= 9:
# indices.sort()
# print indices
# for elem in range(9,llen):
# print 'Delete: ', indices[elem]
#     r = requests.delete('http://elk-102.ix.km:9200/'+indices[elem])
#     print r
