#!/usr/bin/python
# -*- coding: utf-8 -*-

import requests
import logging

###############
# basicConfig #
###############
logging.basicConfig(format='%(asctime)s %(message)s', filename='/var/log/elastic_drop.log', level=logging.INFO)
HOST = 'elk-102.ix.km'
exludeIndicies = ['logstash-2014.05.24', 'logstash-2014.05.25', 'logstash-2014.05.27', 'logstash-2014.05.29',\
                  'logstash-2014.05.31', 'logstash-2014.05.30', '.kibana', 'grafana-dash', 'kibana-int']
################

if len(exludeIndicies) == 0:
    logging.info('exludeIndicies is empty: do nothing')
    exit(1)

try:
    response = requests.get('http://%s:9200/_cluster/health' % HOST, timeout=5)
except requests.exceptions.Timeout as t:
    logging.info('Connection timeout, %s' % t)
    exit(1)

if response.status_code != 200:
    logging.info('Wrong response, %s' % response.status_code)
    exit(1)
logging.info('HTTP answer is: %s -> doing work' % response.status_code)

responseBody = response.json()
if responseBody['status'] != 'green':
    logging.info('Cluster status is not green: %s -> stopped' % responseBody['status'])
    exit(1)
logging.info('cluster Status is: %s -> doing work' % responseBody['status'])

try:
    response = requests.get('http://%s:9200/_cat/indices' % HOST)
except requests.exceptions.Timeout as t:
    logging.info('Connection timeout, %s' % t)
    exit(1)
if response.status_code != 200:
    logging.info('Wrong response, %s' % response.status_code)
    exit(1)
logging.info('http answer code is: %s -> doing work' % response.status_code)

responseBody = response.text.rstrip()
responseBody = responseBody.split('\n')
clusterIndicies = list()
for i in responseBody:
    k = i.split(' ')
    clusterIndicies.append(k[2])
clusterIndicies.sort()

resultIndicies = list()
for index in clusterIndicies:
    if index not in exludeIndicies:
        resultIndicies.append(index)

if len(resultIndicies) - 60 == 0:
    logging.info('Nothing to delete')
    exit(0)

for index in range(len(resultIndicies) - 60):
    logging.info('Trying to delete index: %s' % resultIndicies[index])
    try:
        response = requests.delete(('http://%s:9200/' % HOST)+resultIndicies[index], timeout=5)
    except requests.exceptions.Timeout:
        logging.info('Could not delete index, %s' % resultIndicies[index])
    if response.status_code != 200:
        logging.info('Could not delete index: %s httpcode: %s' % resultIndicies[index], response.status_code)
    logging.info('Index: %s deleted' % resultIndicies[index])
