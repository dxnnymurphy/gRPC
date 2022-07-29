import pandas as pd
from sklearn.covariance import EllipticEnvelope
from elasticsearch import Elasticsearch
import elasticsearch.helpers
import json
import numpy as np
from kafka import KafkaProducer
from datetime import date
import base64
from utils import ESClient, ADUtils

#define global variables
kafka_server = ['10.150.49.141:9092', '10.150.49.142:9092', '10.150.49.143:9092']
elasticsearch_server = ['10.150.49.141:9200', '10.150.49.142:9200', '10.150.49.143:9200']




class MyEllipticEnvelope(EllipticEnvelope):
    
    def __init__(self, contamination=0.02, X_train=None):
        EllipticEnvelope.__init__(self, contamination=contamination)
        self.X_train = X_train
    
    def saveModel(self):
        dict_ = {}
        dict_['contamination'] = self.contamination
        dict_['X_train'] = self.X_train.tolist() if self.X_train is not None else 'None'
        json_text = json.dumps(dict_, indent=4)
        encoded = base64.b64encode(bytes(json_text, "utf-8"))
        producer = KafkaProducer(bootstrap_servers=kafka_server)
        producer.send('dev_tradingexpert_fixlogtracer_vertex_anomalydetector_executionvolume_model', key=b'model', value=encoded)
    
    def loadModel(self):
        es = ESClient(elasticsearch_server)
        df_model = es.Query("SELECT TOP 1 \"@timestamp\", event.original FROM \"dev_tradingexpert_fixlogtracer_vertex_anomalydetector_model\" ORDER BY \"@timestamp\" DESC ")
        decoded = base64.b64decode(df_model['event.original'].values[0])
        dict_ = json.loads(decoded)
        self.contamination = dict_['contamination']
        self.X_train = np.asarray(dict_['X_train']) if dict_['X_train'] != 'None' else None
        self.fit(self.X_train)
    
    def train(self, time=None):
        #by default will take data from the last 24 hours and retrain the model
        if time:
            query = "SELECT \"@timestamp\", \"spec.fix.PRICE\", \"spec.fix.ORDERQTY\" FROM  \"dev_tradingexpert_fixlogtracer_source_rawfixlogs-*\" WHERE \"spec.fix.EXECTYPE\" IS NOT NULL AND \"@timestamp\" BETWEEN '" + time[0] +"' AND '" + time[1] +"' AND spec.fix.EXECTYPE = 'eliminate'"
        else:
            query = "SELECT \"@timestamp\", \"spec.fix.PRICE\", \"spec.fix.ORDERQTY\" FROM  \"dev_tradingexpert_fixlogtracer_source_rawfixlogs-*\" WHERE \"spec.fix.EXECTYPE\" IS NOT NULL AND \"@timestamp\" >= NOW() - INTERVAL 1 DAY"
        es = ESClient(elasticsearch_server)
        df = es.Query(query)
        df_new, X_train_new = ADUtils.calculateExecutionVolume(df)
        self.X_train = X_train_new
        self.saveModel()