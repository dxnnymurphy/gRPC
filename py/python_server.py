from utils import ESClient, ADUtils
from models import MyEllipticEnvelope
from datetime import date
from kafka import KafkaProducer
from concurrent.futures import ThreadPoolExecutor
import logging
import grpc
import numpy as np
from grpc_reflection.v1alpha import reflection
from anomaly_pb2 import AnomalyResponse, DESCRIPTOR
from anomaly_pb2_grpc import AnomalyDetectionServicer, add_AnomalyDetectionServicer_to_server

#define global variables
kafka_server = ['10.150.49.141:9092', '10.150.49.142:9092', '10.150.49.143:9092']
elasticsearch_server = ['10.150.49.141:9200', '10.150.49.142:9200', '10.150.49.143:9200']

class AnomalyServer(AnomalyDetectionServicer):

    def Predict(self, request, context):
        #if refresh = on reload model
        envelope = MyEllipticEnvelope()
        for m in request.metrics:
            if m.refreshbool != "false":
                envelope.refresh()
            else:
                envelope.loadModel()
        
            if m.topic == "default":
                topic = "dev_tradingexpert_fixlogtracer_vertex_anomalydetector_executionvolume"
            else:
                topic = m.topic
        
        #will be set to date.today but for debug purposes
        today = '2022-07-01'
        es = ESClient(elasticsearch_server)
        #atm querys all of day up till now need to change (waste of memory), also may need option for index to query
        query = "SELECT \"@timestamp\", \"spec.fix.PRICE\",  \"spec.fix.EXECID\",  \"spec.fix.ORDERQTY\",  \"spec.fix.EXECTYPE\" FROM  \"dev_tradingexpert_fixlogtracer_source_rawfixlogs-*\" WHERE \"spec.fix.EXECTYPE\" IS NOT NULL AND \"@timestamp\" >= '" + today + "T00:00:00' AND spec.fix.EXECTYPE = 'eliminate'"
        df_X_predict = es.Query(query)
        logging.info('model x train: %r', envelope.X_train)
        logging.info('model contamintation: %r', envelope.contamination)
        df_X_predict, X_predict = ADUtils.calculateExecutionVolume(df_X_predict)
        results = envelope.predict(X_predict)
        df_X_predict['anomaly'] = results
        df_anomalies = ADUtils.getAnomalies(df_X_predict, results)
        producer = KafkaProducer(bootstrap_servers=kafka_server)
        df_anomalies = df_anomalies.to_json(orient='records', lines=True).split("\n")
        logging.info('%r', len(df_anomalies))
        for item in df_anomalies:
            producer.send(topic, key=b'', value=item.encode('utf-8'))
        string = "Model has been ran"
        resp = AnomalyResponse(response = string)
        return resp
    
    def Train(self, request, context):
        for m in request.metrics:
            time = []
            if m.startTime:
                time.append(m.startTime)
            if m.endTime:
                time.append(m.endTime)
            envelope = MyEllipticEnvelope()
            envelope.train(time)



if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
    server = grpc.server(ThreadPoolExecutor())
    add_AnomalyDetectionServicer_to_server(AnomalyServer(), server)
    SERVICE_NAMES = (
        DESCRIPTOR.services_by_name['AnomalyDetection'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    #change
    port = 9999
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    logging.info('server ready on port %r', port)
    server.wait_for_termination()
