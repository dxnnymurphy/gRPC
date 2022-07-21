from utils import ESClient, ADUtils
from models import MyEllipticEnvelope
from datetime import date
from kafka import KafkaProducer
from concurrent.futures import ThreadPoolExecutor
import logging
import grpc
import numpy as np
from anomaly_pb2 import AnomalyResponse
from anomaly_pb2_grpc import AnomalyDetectionServicer, add_AnomalyDetectionServicer_to_server

#define global variables
kafka_server = ['10.150.49.141:9092', '10.150.49.142:9092', '10.150.49.143:9092']
elasticsearch_server = ['10.150.49.141:9200', '10.150.49.142:9200', '10.150.49.143:9200']

class AnomalyServer(AnomalyDetectionServicer):

    def Predict(self, request, context):
        #if refresh = on reload model
        envelope = MyEllipticEnvelope()
        if request.metrics.refreshbool == True:
            envelope.refresh()
        else:
            envelope.loadModel()
        
        if request.metrics.topic == None:
            topic = "dev_tradingexpert_fixlogtracer_vertex_anomalydetector_executionvolume"
        
        #think this may be defined in the proxy not in the server
        today = str(date.today())
        es = ESClient(elasticsearch_server)
        #atm querys all of day up till now need to change (waste of memory), also may need option for index to query
        query = "SELECT * FROM  \"dev_tradingexpert_fixlogtracer_source_rawfixlogs-*\" WHERE \"spec.fix.17\" IS NOT NULL AND \"spec.fix.37\" IS NOT NULL AND \"spec.fix.35\" != 0 AND \"spec.fix.EXECTYPE\" IS NOT NULL AND \"@timestamp\" >=" + today
        df_X_predict = es.Query(query)
        X_predict = ADUtils.calculateExecutionVolume(df_X_predict)
        results = envelope.predict(X_predict)
        df_anomalies = ADUtils.getAnomalies(df_X_predict, results)
        producer = KafkaProducer(bootstrap_servers=kafka_server)
        df_anomalies = df_anomalies.to_json(orient='records', lines=True).split("\n")
        for item in df_anomalies:
            producer.send(topic, key=b'', value=item)
        string = "Model has been ran"
        resp = AnomalyResponse(response = string)
        return resp

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
    server = grpc.server(ThreadPoolExecutor())
    add_AnomalyDetectionServicer_to_server(AnomalyServer(), server)
    #change
    port = 9999
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    logging.info('server ready on port %r', port)
    server.wait_for_termination()
