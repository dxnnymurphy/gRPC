from elasticsearch import Elasticsearch
import pandas as pd


class ESClient:
    def __init__(self, ip):
        #change this from hardcoding
        self.es = Elasticsearch(ip, http_auth=("dmurphy", "Queensmead11"), use_ssl=True, verify_certs=False,timeout=60)

    def parseResponse(self,response):
        parsed = []
        rowcounter = 0
        for row in response['rows']:
            parsed.append({})
            columncounter = 0
            for column in response['columns']:
                columnName = response['columns'][columncounter]['name']
                parsed[rowcounter][columnName] = response['rows'][rowcounter][columncounter]
                columncounter += 1
            rowcounter += 1
        return parsed

    def Query(self, query):
        res = self.es.sql.query(body={'query': query, 'fetch_size':10000})
        df=pd.DataFrame(self.parseResponse(res))
        while res.get('cursor'):
            res_new = self.es.sql.query(body={"cursor": res['cursor']})
            res_new['columns'] = res['columns']
            df_new = pd.DataFrame(self.parseResponse(res_new))
            df = pd.concat([df, df_new])
            res = res_new
        return df

class ADUtils:
    def __init__(self):
        pass

    def calculateExecutionVolume(df):
        df['spec.fix.PRICE'] = df['spec.fix.PRICE'].astype(float)
        df['spec.fix.ORDERQTY'] = df['spec.fix.ORDERQTY'].astype(float)
        df['exec_volume'] = (df['spec.fix.ORDERQTY'] * df['spec.fix.PRICE'])
        df_tidy = df.dropna(subset=['exec_volume'])
        X_values = df_tidy['exec_volume'].values.reshape(-1,1)
        return df_tidy, X_values
    
    def getAnomalies(df, results):
        df_anomalies = df.loc[df['anomaly'] == -1].drop('anomaly', axis=1)
        return df_anomalies